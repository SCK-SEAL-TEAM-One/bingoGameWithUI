package api

import (
	"bingo/service"
	"encoding/json"
	"net/http"
)

type StartGameRequest struct {
	PlayerNames []string `json:"players"`
}

type Api struct {
	GameService service.Service
}

func (a *Api) StartGameHandler(writer http.ResponseWriter, request *http.Request) {
	requestGame := StartGameRequest{}
	err := json.NewDecoder(request.Body).Decode(&requestGame)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	err = a.GameService.NewGame(requestGame.PlayerNames)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (a *Api) GetPlayersInfoHandler(writer http.ResponseWriter, request *http.Request) {
	playerInfoResponse := a.GameService.GetPlayerInfo()
	playerInfoResponseJson, _ := json.Marshal(playerInfoResponse)
	_, err := writer.Write(playerInfoResponseJson)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
}

func (a *Api) PlayHandler(writer http.ResponseWriter, request *http.Request) {
	playResponse := a.GameService.PlayGame()
	playJson, _ := json.Marshal(playResponse)
	_, err := writer.Write(playJson)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
}

func (a *Api) ChangeTicketHandler(writer http.ResponseWriter, request *http.Request) {
	playerName := request.URL.Query().Get("playerName")
	player := a.GameService.ChangeTicket(playerName)
	playerJson, _ := json.Marshal(player)
	_, err := writer.Write(playerJson)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
}
