package main

import (
	"mymodule/golang-book/idioms/ch7/di/controller"
	"mymodule/golang-book/idioms/ch7/di/logic"
	"net/http"
)

func main() {
	l := logic.LoggerAdapter(logic.LogOutput)
	ds := logic.NewSimpleDataStore()
	logic := logic.NewSimpleLogic(l, ds)
	c := controller.NewController(l, logic)
	l.Log("starting server on :8080")

	http.HandleFunc("/hello", c.SayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		l.Log(err.Error())
	}
}