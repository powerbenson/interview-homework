package main

import (
	"bufio"
	"os"
	"github.com/powerbenson/interview-homework/server"
)

func main()  {
	server := server.New(&server.Config{
		Host: "localhost",
		Port: "8080",
	})
	go server.Run()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()

		if input == "Q" {
			server.Stop()
			break
		}
		server.SendMessage(input)
	}
}