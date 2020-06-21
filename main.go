package main

import (
	"log"
	"net/http"
	"time"
	"github.com/heptiolabs/healthcheck"
)

func main() {
	health := healthcheck.NewHandler()
	health.AddLivenessCheck("liveness", func() error {
		return nil
	})

	go http.ListenAndServe("0.0.0.0:8080", health)

	for {
		log.Println(message())
		time.Sleep(time.Second)
	}
}

func message() string {
	return "Hello World"
}