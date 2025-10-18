package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("🌐 BFF rodando na porta 8080...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("BFF Online 🚀"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
