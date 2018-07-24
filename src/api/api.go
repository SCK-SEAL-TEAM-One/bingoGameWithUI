package api

import (
	"bingogame"
	"encoding/json"
	"net/http"
	"service"
)

type StartGameRequest struct {
	PlayerOne string `json:"playerOne"`
	PlayerTwo string `json:"playerTwo"`
}

type Api struct {
	GameService service.Service
}

type PlayerInfoResponse struct {
	PlayerOne     bingogame.Player `json:"playerOne"`
	PlayerTwo     bingogame.Player `json:"playerTwo"`
	HistoryPickUp []int            `json:"historyPickUp"`
}

func (a *Api) StartGameHandler(writer http.ResponseWriter, request *http.Request) {
	requestGame := StartGameRequest{}
	err := json.NewDecoder(request.Body).Decode(&requestGame)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	err = a.GameService.NewGame(requestGame.PlayerOne, requestGame.PlayerTwo)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (a *Api) GetPlayersInfoHandler(writer http.ResponseWriter, request *http.Request) {
	playerInfoResponse := a.GameService.GetPlayerInfo()
	playerInfoResponseJson, _ := json.Marshal(playerInfoResponse)
	writer.Write(playerInfoResponseJson)
}

func (a *Api) PlayHandler(writer http.ResponseWriter, request *http.Request) {
	playResponse := a.GameService.PlayGame()
	playJson, _ := json.Marshal(playResponse)
	writer.Write(playJson)
}

func (a *Api) ChangeTicketHandler(writer http.ResponseWriter, request *http.Request) {
	playerName := request.URL.Query().Get("playerName")
	player := a.GameService.ChangeTicket(playerName)
	playerJson, _ := json.Marshal(player)
	writer.Write(playerJson)
}
