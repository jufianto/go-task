# go-task
sistem aplikasi hal kegiatan apa saja yang sudah dilakukan dan data tersebut disimpan kedalam data portabel seperti csv maupun sqlite

## Cara Penggunaan Dev Applikasi Golang
1. ketik perintah terminal `go get github.com/jufianto/go-task`
2. pergi ke folder `$GOPATH/src/github.com/jufianto/go-task`
3. jalankan perintah `go run main.go route.go` 
4. buka Postman dan setting URL postman ke `localhost:8000/` method `POST` dengan pameter raw `JSON` 
```json
{
	"issuer": "Jufianto",
	"task": "Buat Aplikasi Open Source"
}
```
5. Jika berhasil akan mendapat respon http status `200` dengan response body 
```json
{
    "isSuccess": true,
    "message": "Berhasil Menambahkan Data ke CSV"
}
```

## Cara Penggunaan Hasil Build Exe 
1. clone repository ini 
2. pergi ke folder build dan jalankan main.exe 
3. dan ikuti langkah 4 sampai dengan 5 pada bagian penggunaan aplikasi dev
