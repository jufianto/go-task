package lib

import (
	"fmt"
	"io"
	"log"
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

// CreateCSV buat file csv berdasarkan bulan dan tahun, jika tidak ada buat baru
func CreateCSV() io.Writer {
	today := time.Now()
	year, month, _ := today.Date()
	filename := fmt.Sprintf("kegiatan/kegiatan-bulan-%02d-%d.csv", month, year)

	if err := fileExists(filename); err {
		fmt.Println("file tidak ada, buat baru")
		file, err := os.Create(filename)
		CheckError("error create file", err)
		fmt.Printf("file csv dengan nama %s telah dibuat \n", filename)
		return file
	} else {
		fmt.Println("file telah ada")
		file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			log.Fatalf("failed opening file: %s", err)
		}
		return file
	}
}

// fileExists check file csv on folder kegiatan
func fileExists(filename string) bool {
	path := fmt.Sprintf("kegiatan/%s", filename)
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}
