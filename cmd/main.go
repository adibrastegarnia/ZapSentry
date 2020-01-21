package main

import (
	"github.com/adibrastegarnia/ZapSentry/pkg/log"
)

func main() {
	dsn := "http://82001944a998481aac5336c28865b0b7@127.0.0.1:9000/3"
	log := log.GetLogger(dsn, "log_test")

	log.Info("Hello")
	log.Error("Error new 6 ")
	log.Error("Error new 7")

}
