package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/andrei828/golang-practice/basic-web-server/handlers"
)

var bindAddress = ":9090"

func main() {

  logger := log.New(os.Stdout, "webserver", log.LstdFlags)
  entity := handlers.NewEntity(logger)
  router := http.NewServeMux()
  router.Handle("/", entity)
  
  server := &http.Server {
    Addr:         bindAddress,
    Handler:      router,
    ErrorLog:     logger,
    ReadTimeout:  5 * time.Second,
    WriteTimeout: 10 * time.Second,
    IdleTimeout:  120 * time.Second,
  }

  go startServer(server, logger)
  handleShutdown(server, logger)
}

func startServer(server *http.Server, logger *log.Logger) {
  logger.Println("Starting the web server on port %s", bindAddress)

  err := server.ListenAndServe()
  if err != nil {
    logger.Printf("Error starting server: %s\n", err)
    os.Exit(1)
  }
}

func handleShutdown(server *http.Server, logger *log.Logger) {
  c := make(chan os.Signal, 1)
  signal.Notify(c, os.Interrupt)
  signal.Notify(c, os.Kill)

  sig := <- c
  logger.Println("Got signal", sig)

  ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
  server.Shutdown(ctx)
}
