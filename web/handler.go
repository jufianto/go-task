package web

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/jufianto/go-task/lib"
	"net/http"
	"time"
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
		w.Write([]byte("hello dolli"))
		fmt.Println("running")

	case "POST":
		file := lib.CreateCSV()
		writer := csv.NewWriter(file)
		defer writer.Flush()
		task := lib.Task{}
		//format := "2020-01-10"
		waktu := time.Now().Format("Monday, 02 January 2006")
		task.Date = waktu
		err := json.NewDecoder(r.Body).Decode(&task)
		defer r.Body.Close()
		lib.CheckError("error reading body json: ", err)
		fmt.Printf("%+v \n", task)
		dt := []string{task.Date, task.Issuer, task.DescriptionJob}
		err = writer.Write(dt)
		lib.CheckError("Error write file >>> ", err)
	}
}
