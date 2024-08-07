package main

import (
	"net/http"

	c "github.com/chrisbradleydev/go-dependency-injection/pkg/controller"
	s "github.com/chrisbradleydev/go-dependency-injection/pkg/datastore"
	l "github.com/chrisbradleydev/go-dependency-injection/pkg/logger"
	sl "github.com/chrisbradleydev/go-dependency-injection/pkg/simplelogic"
)

func main() {
	logger := l.LoggerAdapter(l.LogOutputMagenta)
	dataStore := s.NewSimpleDataStore()
	logic := sl.NewSimpleLogic(logger, dataStore)
	controller := c.NewController(logger, logic)
	http.HandleFunc("/hello", controller.SayHello)
	http.ListenAndServe(":8080", nil)
}
