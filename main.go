package main

import (
	"log"
	"time"
)

func main() {
	for {
		log.Println(message())
		time.Sleep(time.Second)
	}
}

func message() string {
	return "Hello World"
}