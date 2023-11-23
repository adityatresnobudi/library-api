package router

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/adityatresnobudi/library-api/internal/db"
	handler "github.com/adityatresnobudi/library-api/internal/handler/handler_rest"
	"github.com/adityatresnobudi/library-api/internal/logger"
	"github.com/adityatresnobudi/library-api/internal/middleware"
	"github.com/adityatresnobudi/library-api/internal/repository"
	"github.com/adityatresnobudi/library-api/internal/usecase"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func WithTimeout(c *gin.Context) {
	ctx := c.Request.Context()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)
	c.Next()
}

func NewRouter(h *handler.Handler) *gin.Engine {
	router := gin.Default()
	router.ContextWithFallback = true

	router.Use(requestid.New())
	router.Use(middleware.Logger(logger.NewLogger()))
	router.Use(middleware.GlobalErrorMiddleware())

	book := router.Group("/books", WithTimeout)
	book.GET("", h.GetBooks)
	book.POST("", h.AddBooks)

	user := router.Group("/users", WithTimeout)
	user.GET("", h.GetUsers)
	user.POST("", h.CreateUser)
	user.POST("/login", h.LoginUser)

	record := router.Group("/records", WithTimeout, middleware.Auth())
	record.POST("/borrow", h.BorrowBook)
	record.PUT("/return/:id", h.ReturnBook)

	return router
}

func Serve() {
	db, err := db.Connect()
	if err != nil {
		log.Println(err)
	}

	br := repository.NewBookRepository(db)
	bu := usecase.NewBookUsecase(br)

	ur := repository.NewUserRepository(db)
	uu := usecase.NewUserUsecase(ur)

	rr := repository.NewBorrowRecordRepository(db)
	ru := usecase.NewBorrowRecordUsecase(rr)

	h := handler.NewHandler(bu, uu, ru)
	router := NewRouter(h)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	log.Println("timeout of 5 seconds.")

	log.Println("Server exiting")
}
