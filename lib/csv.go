package lib

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type task struct {
	date           string
	issues         string
	descriptionJob string
}

// CheckError untuk cek error jika terjadi error
func CheckError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

// CreateCSV buat file csv berdasarkan bulan dan tahun, jika tidak ada buat baru
func CreateCSV() io.Writer {
	today := time.Now()
	year, month, _ := today.Date()
	filename := fmt.Sprintf("kegiatan-bulan-%02d-%d.csv", month, year)

	if err := fileExists(filename); err {
		fmt.Println("file tidak ada, buat baru")
		file, err := os.Create(filename)
		CheckError("error create file", err)
		fmt.Printf("file csv dengan nama %s telah dibuat \n", filename)
		defer file.Close()
	} else {
		fmt.Println("file telah ada")
	}
}

// fileExists check file csv on folder kegiatan
func fileExists(filename string) bool {
	path := fmt.Sprintf("kegiatan/%s", filename)
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}
