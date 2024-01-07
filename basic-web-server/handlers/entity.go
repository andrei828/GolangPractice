package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type Entity struct {
  logger *log.Logger  
}

func NewEntity(l *log.Logger) *Entity {
  return &Entity{l}
}

func (entity *Entity) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
  fmt.Println("got here")
}
