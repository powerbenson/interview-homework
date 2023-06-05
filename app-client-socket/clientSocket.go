package main

import (
	"github.com/powerbenson/interview-homework/client"
)

func main()  {
	client := client.New(&client.Config{
		Host: "localhost",
		Port: "8080",
	})
	client.Run("client1", "socket")
	defer client.Stop()
	client.Receive()
}