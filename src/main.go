package main

import (
	"api"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/bingo/play", api.PlayHandler)
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/bingogame/", http.StripPrefix("/bingogame/", fs))

	log.Fatal(http.ListenAndServe(":3000", nil))
}
