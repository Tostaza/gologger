FROM golang:1.23

WORKDIR /usr/src/gologger

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/gologger ./...

ENV GOLOGGER_LISTEN_ADDR="0.0.0.0"
ENV GOLOGGER_LISTEN_PORT="8080"

CMD ["gologger"]
