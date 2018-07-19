package main

import (
	"api"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/bingo/play", api.PlayHandeler)
	log.Fatal(http.ListenAndServe(":3000", http.FileServer("./public")))
}
