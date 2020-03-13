package lib

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Task struct {
	Date           string `json:"date"`
	Issuer         string `json:"issuer"`
	DescriptionJob string `json:"task"`
}

// CheckError untuk cek error jika terjadi error
func CheckError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func GenerateFileName() (fileName string) {
	year, month, _ := time.Now().Date()
	fileName = fmt.Sprintf("kegiatan/kegiatan-bulan-%02d-%d.csv", month, year)
	return
}

func (t *Task) TambahTask() {
	t.Date = time.Now().Format("Monday, 02 January 2006")
	data := []string{t.Date, t.Issuer, t.DescriptionJob}
	file, err := os.OpenFile(GenerateFileName(), os.O_APPEND|os.O_WRONLY, 0600)
	CheckError("error when open file:  ", err)
	writer := csv.NewWriter(file)
	writer.Write(data)
	defer writer.Flush()
}

func ReadCSV(nameFile string) [][]string {
	csvFile, err := os.Open(nameFile)
	CheckError("error when open csv file", err)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	taskList, err := reader.ReadAll()
	CheckError("error when read all csv data => ", err)
	return taskList
}

// CreateCSV buat file csv berdasarkan bulan dan tahun, jika tidak ada buat baru
func CreateCSV(fileName string) io.Writer {
	if err := fileExists(fileName); err {
		fmt.Println("file tidak ada, buat baru")
		file, err := os.Create(fileName)
		CheckError("error create file", err)
		fmt.Printf("file csv dengan nama %s telah dibuat \n", fileName)
		return file
	} else {
		fmt.Println("file telah ada")
		file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			log.Fatalf("failed opening file: %s", err)
		}
		return file
	}
}

type Response struct {
	IsSuccess  bool   `json:"isSuccess"`
	HttpStatus int    `json:"-"`
	Message    string `json:"message"`
}

func (t *Task) ResponeJson(w http.ResponseWriter, body Response) {
	w.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(body)
	w.Write(b)
}

// fileExists check file csv on folder kegiatan
func fileExists(filename string) bool {
	path := fmt.Sprintf("kegiatan/%s", filename)
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}
