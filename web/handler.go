package web

import (
	"encoding/json"
	"net/http"

	"github.com/jufianto/go-task/lib"
)

type HandlerInterface interface {
	IndexHandler(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
}

func (h Handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch method {
	case "GET":
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte("NICE"))

	case "POST":
		task := lib.Task{}
		err := json.NewDecoder(r.Body).Decode(&task)
		lib.CheckError("error when read Body : ", err)
		defer r.Body.Close()
		task.TambahTask()
		respone := lib.Response{
			IsSuccess:  true,
			HttpStatus: 200,
			Message:    "Berhasil Menambahkan Data ke CSV",
		}
		task.ResponeJson(w, respone)
	}
}
