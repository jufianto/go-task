package web

import (
	"fmt"
	"net/http"
)

type HandlerInterface interface {
	IndexHandler(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
}

func (h Handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("not get")
	method := r.Method
	fmt.Println(method)
	switch method {
	case "GET":
		w.Write([]byte("hello dolli"))
		fmt.Println("running")
	}
}
