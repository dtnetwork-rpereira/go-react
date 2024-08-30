FROM golang:alpine

COPY ./backend /server

WORKDIR /server
# RUN go build

ENTRYPOINT go run .
# ENTRYPOINT ./go-react
