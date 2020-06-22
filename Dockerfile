FROM golang:alpine AS build

EXPOSE 5000

ARG POSTGRES_CONN
ENV CONN_STR=$POSTGRES_CONN

ENV GO111MODULE=on

WORKDIR /go/src/web-services-101

RUN mkdir cors
COPY cors/* ./cors/

RUN mkdir database
COPY database/* ./database/

RUN mkdir topping
COPY topping/* ./topping/

COPY main.go .
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go build -o donutStoreAPI
ENTRYPOINT ["./donutStoreAPI"]
CMD ["-PORT=5000", "-DB_CONN_STR=''"]
