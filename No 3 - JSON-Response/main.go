package main

import (
	"log"
	"mahasiswa"
	"net/http"
)

func main() {
	http.HandleFunc("/mahasiswa", mahasiswa.CreateMahasiswa)
	http.HandleFunc("/mahasiswa/all", mahasiswa.ViewAllMahasiswa)

	log.Fatal(http.ListenAndServe(":5050", nil))
}
