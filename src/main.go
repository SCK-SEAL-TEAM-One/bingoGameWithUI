package main

import (
	apiPackage "api"
	"log"
	"net/http"
	"service"
)

func main() {
	api := apiPackage.Api{GameService: &service.MockGameService{}}
	http.HandleFunc("/bingo/start", api.StartGameHandler)
	http.HandleFunc("/bingo/info", api.GetPlayersInfoHandler)
	http.HandleFunc("/bingo/play", api.PlayHandler)
	http.HandleFunc("/bingo/ticket/change", api.ChangeTicketHandler)
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/bingogame/", http.StripPrefix("/bingogame/", fs))

	log.Fatal(http.ListenAndServe(":3000", nil))
}
