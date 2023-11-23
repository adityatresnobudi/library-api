# syntax=docker/dockerfile:1
FROM golang:1.18-alpine as buildStage

WORKDIR /library-api
COPY ./ ./
RUN go build -ldflags "-s -w" -o /output

FROM alpine:latest
WORKDIR /apps
COPY .env .
COPY --from=buildStage /output /output
EXPOSE 8080
CMD [ "/output" ]