package main

import (
	"github.com/jufianto/go-task/lib"
	"github.com/jufianto/go-task/web"
)

func Route(mux *lib.JufiMux) {
	handlerIndex := new(web.Handler)
	mux.HandleFunc("/", handlerIndex.IndexHandler)
}
