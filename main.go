package main

import (
	"fmt"
	"github.com/jufianto/go-task/lib"
	"net/http"
)

func main() {
	fmt.Println("Initate Go Task")
	mux := new(lib.JufiMux)
	Route(mux)
	server := new(http.Server)
	port := "8000"
	server.Addr = fmt.Sprintf(":%s", port)
	server.Handler = mux
	fmt.Printf("listen and serve on localhost port %s \n", port)
	server.ListenAndServe()
}
