package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/adityatresnobudi/library-api/internal/db"
	handler "github.com/adityatresnobudi/library-api/internal/handler/handler_grpc"
	"github.com/adityatresnobudi/library-api/internal/interceptor"
	"github.com/adityatresnobudi/library-api/proto/pb"
	"github.com/adityatresnobudi/library-api/internal/repository"
	"github.com/adityatresnobudi/library-api/internal/usecase"
	"google.golang.org/grpc"
)

func main() {
	db, err := db.Connect()
	if err != nil {
		log.Println(err)
	}

	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	// setup the network
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Println("Failed to listen:", err)
		return
	}

	// setup the server
	server := grpc.NewServer(grpc.ChainUnaryInterceptor(interceptor.LoggerInterceptor, interceptor.ErrorInterceptor, interceptor.AuthInterceptor))

	// register the handler to server
	ur := repository.NewUserRepository(db)
	br := repository.NewBorrowRecordRepository(db)

	uu := usecase.NewUserUsecase(ur)
	bu := usecase.NewBorrowRecordUsecase(br)

	lh := handler.NewLoginHandler(uu)
	brh := handler.NewBorrowHandler(bu)

	pb.RegisterAuthServer(server, lh)
	pb.RegisterBorrowServer(server, brh)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		s := <-sigCh
		log.Printf("got signal %v, attempting graceful shutdown", s)
		cancel()
		server.GracefulStop()
		// grpc.Stop() // leads to error while receiving stream response: rpc error: code = Unavailable desc = transport is closing
		wg.Done()
	}()

	log.Println("starting grpc server")
	err = server.Serve(listener)
	if err != nil {
		log.Fatalf("could not serve: %v", err)
	}
	wg.Wait()
	log.Println("clean shutdown")
}
