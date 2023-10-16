package controller

import (
	"mymodule/golang-book/idioms/ch7/di/logic"
	"net/http"
)

type Logic interface {
	SayHello(id string) (string, error)
}

type Controller struct {
	l logic.Logger
	logic Logic
}

func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	c.l.Log("In Controller SayHello")
	id := r.URL.Query().Get("id")
	msg, err := c.logic.SayHello(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(msg))
}

func NewController(l logic.Logger, logic Logic) Controller {
	return Controller{
		l: l,
		logic: logic,
	}
}