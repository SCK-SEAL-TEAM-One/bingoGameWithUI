package api

import (
	"bingogame"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"service"
)

var ticket bingogame.Ticket

type StartGameRequest struct {
	PlayerOne string `json:"playerOne"`
	PlayerTwo string `json:"playerTwo"`
}

type Api struct {
	GameService service.GameService
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
	playerInfoResponse := a.GameService.GetPlayerInfo()
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

	err = a.GameService.NewGame(requestGame.PlayerOne, requestGame.PlayerTwo)

	if err != nil {
		http.Error(writer, err.Error(), 400)
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
	playResponse := a.Service.Game.Play()
	playJson, _ := json.Marshal(playResponse)
	writer.Write(playJson)
}
