package controller

import (
	"net/http"

	l "github.com/chrisbradleydev/go-dependency-injection/pkg/logger"
	sl "github.com/chrisbradleydev/go-dependency-injection/pkg/simplelogic"
)

type Controller struct {
	l     l.Logger
	logic sl.Logic
}

func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	c.l.Log("Controller SayHello")
	userID := r.URL.Query().Get("user_id")
	message, err := c.logic.SayHello(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}

func NewController(l l.Logger, logic sl.Logic) Controller {
	return Controller{
		l:     l,
		logic: logic,
	}
}
