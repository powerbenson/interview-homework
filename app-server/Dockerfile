FROM golang:1.20

WORKDIR /app

COPY . .

RUN mkdir /tmp/conn-file

RUN go mod download

RUN go build ./app-server/serverSender.go

CMD ["/app/serverSender"]
