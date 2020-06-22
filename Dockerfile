# Stage -1 : vet and lint
FROM golang:alpine AS lint
ENV GO111MODULE=on
ENV CGO_ENABLED=1
RUN apk add build-base
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
RUN go vet ./...
RUN go get -u golang.org/x/lint/golint
RUN golint ./...

# Stage -2: Run Unit tests
FROM golang:alpine AS tests
ARG DB_CONN_STR
RUN apk add build-base
WORKDIR /go/src/web-services-101
COPY --from=lint /go/src/web-services-101 .
RUN go test -v ./... -DB_CONN_STR=$DB_CONN_STR

# Stage -3: Build
FROM golang:alpine AS build
RUN apk --no-cache add ca-certificates
WORKDIR /go/src/web-services-101
COPY --from=tests /go/src/web-services-101 .
RUN go build -o donutStoreAPI

# Stage -4: Ship
FROM scratch as ship
EXPOSE 5000
WORKDIR /go/src/app
COPY --from=build /go/src/web-services-101/donutStoreAPI ./
ENTRYPOINT ["./donutStoreAPI"]
CMD ["-PORT=5000", "-DB_CONN_STR=''"]
