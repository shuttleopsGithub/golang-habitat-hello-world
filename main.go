package main

import (
	"log"
	"os"
	"time"
)

func main() {
	for {
		log.Println(message(os.Args[1]))
		time.Sleep(time.Second)
	}
}

func message(message string) string {
	return message
}