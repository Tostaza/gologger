FROM golang:1.23.5-alpine3.21 AS build
WORKDIR /usr/src/gologger

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -o /usr/local/bin/gologger ./...


FROM alpine:3.21.2
COPY --from=build /usr/local/bin/gologger /usr/local/bin/gologger 
CMD ["/usr/local/bin/gologger"]
