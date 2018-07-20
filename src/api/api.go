package api

import (
	"bingogame"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	numberOfGrid = 5
	numberInBox  = 75
)

var ticket bingogame.Ticket

type StartGameRequest struct {
	PlayerOne string `json:"playerOne"`
	PlayerTwo string `json:"playerTwo"`
}

type Api struct {
	Game bingogame.Game
}

type PlayerInfoResponse struct {
	PlayerOne     bingogame.Player `json:"playerOne"`
	PlayerTwo     bingogame.Player `json:"playerTwo"`
	HistoryPickUp []int            `json:"historyPickUp"`
}

type PlayResponse struct {
	Number int    `json:"number"`
	Winner string `json:"winner"`
}

func (a *Api) GetPlayersInfoHandler(writer http.ResponseWriter, request *http.Request) {
	gameData, err := ioutil.ReadFile("./gamedata")
	if err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}
	err = json.Unmarshal(gameData, &a.Game)
	if err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}
	playerInfoResponse := PlayerInfoResponse{
		PlayerOne:     a.Game.Players[0],
		PlayerTwo:     a.Game.Players[1],
		HistoryPickUp: a.Game.HistoryPickUp,
	}
	playerInfoResponseJson, _ := json.Marshal(playerInfoResponse)
	writer.Write(playerInfoResponseJson)
}

func (a *Api) StartGameHandler(writer http.ResponseWriter, request *http.Request) {
	requestGame := StartGameRequest{}
	err := json.NewDecoder(request.Body).Decode(&requestGame)
	if err != nil {
		http.Error(writer, err.Error(), 400)
		return
	}
	ticket1 := bingogame.NewTicket(numberOfGrid)
	ticket2 := bingogame.NewTicket(numberOfGrid)
	ticket1 = bingogame.MockTicketNumber(ticket1, 1)
	ticket2 = bingogame.MockTicketNumber(ticket2, 2)
	player1 := bingogame.NewPlayer(requestGame.PlayerOne, ticket1)
	player2 := bingogame.NewPlayer(requestGame.PlayerTwo, ticket2)
	numberBox := bingogame.MockNumberBox()
	allPlayer := []bingogame.Player{player1, player2}
	a.Game = bingogame.NewGame(allPlayer, numberBox)
	gameData, err := json.Marshal(a.Game)
	if err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}
	err = ioutil.WriteFile("./gamedata", gameData, 0644)
	if err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (a Api) PlayHandler(writer http.ResponseWriter, request *http.Request) {
	gameData, err := ioutil.ReadFile("./gamedata")
	if err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}
	err = json.Unmarshal(gameData, &a.Game)
	playResponse := a.Game.Play()
	playJson, _ := json.Marshal(playResponse)
	gameData, err = json.Marshal(a.Game)
	if err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}
	err = ioutil.WriteFile("./gamedata", gameData, 0644)
	if err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}
	writer.Write(playJson)
}
