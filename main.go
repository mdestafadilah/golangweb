package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hai", haiHandler)

	log.Println("Menjalankan server web")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func haiHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hai, saya belajar golang"))
}
