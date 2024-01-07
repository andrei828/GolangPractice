package main

import (
	handlers "handlers/entity.go"
	"net/http"
)

func main() {
  logger = log.i
  router := http.NewServeMux()
  router.Handle("/", handlers.NewEntity())
}
