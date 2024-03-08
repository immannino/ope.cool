package main

import (
	"log"
	"net/http"

	handler "github.com/immannino/ope.cool/api"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := http.NewServeMux()

	r.HandleFunc("/api/fetch", handler.Fetch)
	r.HandleFunc("/api/mysql", handler.Mysql)

	fs := http.FileServer(http.Dir("./"))
	r.Handle("/", fs)

	http.ListenAndServe(":3000", r)
}
