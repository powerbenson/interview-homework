FROM golang:1.20

WORKDIR /app

COPY . .

RUN mkdir /tmp/conn-file

RUN go mod download

RUN go build ./app-client-socket/clientSocket.go

CMD ["/app/serverSender"]
