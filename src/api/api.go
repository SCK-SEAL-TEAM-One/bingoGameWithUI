package api

import(
	"net/http"
	"encoding/json"
	"bingogame"

)
const (
	numberOfGrid = 5
	numberInBox = 75
)
type  StartGameRequest struct{
	PlayerOne string `json:"playerOne"`
	PlayerTwo string `json:"playerTwo"`	
}

type Api struct{
	Game bingogame.Game
} 

func (a Api)StartGameHandler(w http.ResponseWriter, r *http.Request) {
	request := StartGameRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	ticket := bingogame.NewTicket(numberOfGrid)
	ticket1 := bingogame.MockTicketNumber(ticket, 1)
	ticket2 := bingogame.MockTicketNumber(ticket, 2)
	player1 := bingogame.NewPlayer(request.PlayerOne, ticket1)
	player2 := bingogame.NewPlayer(request.PlayerTwo, ticket2)
	numberBox := bingogame.MockNumberBox()
	allPlayer := []bingogame.Player{player1, player2}
	a.Game = bingogame.NewGame(allPlayer, numberBox)
	w.WriteHeader(http.StatusOK)
}