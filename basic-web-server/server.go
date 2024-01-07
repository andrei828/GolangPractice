package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/andrei828/golang-practice/basic-web-server/handlers"
)

func main() {

  logger := log.New(os.Stdout, "webserver", log.LstdFlags)
  entity := handlers.NewEntity(logger)
  router := http.NewServeMux()
  router.Handle("/", entity)
  
  bindAddress := ":9090"
  server := &http.Server {
    Addr:         bindAddress,
    Handler:      router,
    ErrorLog:     logger,
    ReadTimeout:  5 * time.Second,
    WriteTimeout: 10 * time.Second,
    IdleTimeout:  120 * time.Second,
  }

  logger.Println("Starting the web server on %s", bindAddress)
}
