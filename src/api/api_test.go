package api_test

import (
	. "api"
	"bingogame"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func Test_PlayHandler_Should_Be_Json_number_9_winner_Empty(t *testing.T) {

	request := httptest.NewRequest("GET", "/bigo/play", nil)
	responseRecorder := httptest.NewRecorder()
	expected := bingogame.PlayResponse{
		Number: 9,
		Winner: "",
	}

	ticketOne := bingogame.NewTicket(5)
	ticketTwo := bingogame.NewTicket(5)
	ticketWithNumberOne := bingogame.MockTicketNumber(ticketOne, 1)
	ticketWithNumberTwo := bingogame.MockTicketNumber(ticketTwo, 2)
	playerOne := bingogame.NewPlayer("A", ticketWithNumberOne)
	playerTwo := bingogame.NewPlayer("B", ticketWithNumberTwo)
	numberBox := bingogame.MockNumberBox()
	api := Api{
		Game: bingogame.Game{
			Players: []bingogame.Player{
				playerOne,
				playerTwo,
			},
			NumberBox:     numberBox,
			HistoryPickUp: []int{},
		},
	}
	gameData, _ := json.Marshal(api.Game)
	ioutil.WriteFile("./gamedata", gameData, 0644)
	api.PlayHandler(responseRecorder, request)

	response := responseRecorder.Result()
	body, _ := ioutil.ReadAll(response.Body)
	var actual bingogame.PlayResponse
	json.Unmarshal(body, &actual)

	if expected != actual {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}
func Test_GetPlayersInfoHandler_Should_Be_InfoResponse(t *testing.T) {
	url := "/bingo/info"
	request := httptest.NewRequest("GET", url, nil)
	responseRecorder := httptest.NewRecorder()
	ticketOne := bingogame.NewTicket(5)
	ticketTwo := bingogame.NewTicket(5)
	ticketWithNumberOne := bingogame.MockTicketNumber(ticketOne, 1)
	ticketWithNumberTwo := bingogame.MockTicketNumber(ticketTwo, 2)
	playerOne := bingogame.NewPlayer("A", ticketWithNumberOne)
	playerTwo := bingogame.NewPlayer("B", ticketWithNumberTwo)
	expectedResponse := PlayerInfoResponse{
		PlayerOne:     playerOne,
		PlayerTwo:     playerTwo,
		HistoryPickUp: []int{},
	}
	expectedResponseString, _ := json.Marshal(expectedResponse)
	numberBox := []int{9, 2, 1, 3, 7, 3, 6, 2, 3, 6}
	api := Api{
		Game: bingogame.Game{
			Players: []bingogame.Player{
				playerOne,
				playerTwo,
			},
			NumberBox:     numberBox,
			HistoryPickUp: []int{},
		},
	}
	gameData, _ := json.Marshal(api.Game)
	ioutil.WriteFile("./gamedata", gameData, 0644)
	api.GetPlayersInfoHandler(responseRecorder, request)

	response := responseRecorder.Result()
	body, _ := ioutil.ReadAll(response.Body)

	if string(expectedResponseString) != string(body) {
		t.Errorf("Should be \n%s but it got \n%s", string(expectedResponseString), string(body))
	}

}

func Test_StartGameHandler(t *testing.T) {
	requestPlayers := StartGameRequest{PlayerOne: "A", PlayerTwo: "B"}
	requestByte, _ := json.Marshal(requestPlayers)

	request := httptest.NewRequest("POST", "/bingo/start", bytes.NewBuffer(requestByte))
	responseRecorder := httptest.NewRecorder()
	api := Api{}
	api.StartGameHandler(responseRecorder, request)
	expectedRespondstatus := 200
	response := responseRecorder.Result()
	if expectedRespondstatus != response.StatusCode {
		t.Errorf("expect %d but got %d", expectedRespondstatus, response.StatusCode)
	}
}
