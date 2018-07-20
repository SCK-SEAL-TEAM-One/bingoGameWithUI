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

func (a *Api) GetPlayersInfoHandler(w http.ResponseWriter, r *http.Request) {
	gameData, err := ioutil.ReadFile("./gamedata")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	err = json.Unmarshal(gameData, &a.Game)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	playerInfoResponse := PlayerInfoResponse{
		PlayerOne:     a.Game.Players[0],
		PlayerTwo:     a.Game.Players[1],
		HistoryPickUp: a.Game.HistoryPickUp,
	}
	playerInfoResponseJson, _ := json.Marshal(playerInfoResponse)
	w.Write(playerInfoResponseJson)
}

func (a *Api) StartGameHandler(w http.ResponseWriter, r *http.Request) {
	request := StartGameRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	ticket1 := bingogame.NewTicket(numberOfGrid)
	ticket2 := bingogame.NewTicket(numberOfGrid)
	ticket1 = bingogame.MockTicketNumber(ticket1, 1)
	ticket2 = bingogame.MockTicketNumber(ticket2, 2)
	player1 := bingogame.NewPlayer(request.PlayerOne, ticket1)
	player2 := bingogame.NewPlayer(request.PlayerTwo, ticket2)
	numberBox := bingogame.MockNumberBox()
	allPlayer := []bingogame.Player{player1, player2}
	a.Game = bingogame.NewGame(allPlayer, numberBox)
	gameData, err := json.Marshal(a.Game)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	err = ioutil.WriteFile("./gamedata", gameData, 0644)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (a Api) PlayHandler(w http.ResponseWriter, r *http.Request) {
	gameData, err := ioutil.ReadFile("./gamedata")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	err = json.Unmarshal(gameData, &a.Game)
	playResponse := a.Game.Play()
	playJson, _ := json.Marshal(playResponse)
	gameData, err = json.Marshal(a.Game)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	err = ioutil.WriteFile("./gamedata", gameData, 0644)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(playJson)
}

func MockTicketNumber(ticket bingogame.Ticket, mockId int) bingogame.Ticket {
	switch mockId {
	case 1:
		ticketDataId := []int{
			1, 17, 35, 51, 74,
			9, 21, 41, 58, 79,
			2, 23, 0, 47, 68,
			14, 29, 32, 49, 66,
			11, 30, 39, 56, 70}
		ticketDataIdIndex := 0
		for indexRow := 0; indexRow < 5; indexRow++ {
			for indexColumn := 0; indexColumn < 5; indexColumn++ {
				ticket.Grid[indexRow][indexColumn].Number = ticketDataId[ticketDataIdIndex]
				ticketDataIdIndex++
			}
		}
	case 2:
		ticketDataId := []int{
			3, 21, 39, 53, 55,
			12, 29, 32, 54, 67,
			11, 30, 0, 49, 70,
			9, 16, 41, 45, 68,
			7, 19, 44, 52, 72}
		ticketDataIdIndex := 0
		for indexRow := 0; indexRow < 5; indexRow++ {
			for indexColumn := 0; indexColumn < 5; indexColumn++ {
				ticket.Grid[indexRow][indexColumn].Number = ticketDataId[ticketDataIdIndex]
				ticketDataIdIndex++
			}
		}
	case 3:
		ticketDataId := []int{
			1, 16, 32, 46, 62,
			8, 21, 45, 55, 70,
			11, 26, 0, 51, 66,
			13, 20, 41, 49, 72,
			9, 23, 36, 59, 61}
		ticketDataIdIndex := 0
		for indexRow := 0; indexRow < 5; indexRow++ {
			for indexColumn := 0; indexColumn < 5; indexColumn++ {
				ticket.Grid[indexRow][indexColumn].Number = ticketDataId[ticketDataIdIndex]
				ticketDataIdIndex++
			}
		}
	case 4:
		ticketDataId := []int{
			2, 17, 31, 49, 71,
			6, 20, 44, 56, 62,
			7, 18, 0, 50, 69,
			10, 28, 40, 47, 73,
			14, 23, 35, 52, 65}
		ticketDataIdIndex := 0
		for indexRow := 0; indexRow < 5; indexRow++ {
			for indexColumn := 0; indexColumn < 5; indexColumn++ {
				ticket.Grid[indexRow][indexColumn].Number = ticketDataId[ticketDataIdIndex]
				ticketDataIdIndex++
			}
		}
	}
	return ticket
}
