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
	ticketWithNumberOne := MockTicketNumber(ticketOne, 1)
	ticketWithNumberTwo := MockTicketNumber(ticketTwo, 2)
	playerOne := bingogame.NewPlayer("A", ticketWithNumberOne)
	playerTwo := bingogame.NewPlayer("B", ticketWithNumberTwo)
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

	api.PlayHandler(responseRecorder, request)

	response := responseRecorder.Result()
	body, _ := ioutil.ReadAll(response.Body)
	var actual bingogame.PlayResponse
	json.Unmarshal(body, &actual)
	//expected2, _ := json.Marshal(expected)

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
	ticketWithNumberOne := MockTicketNumber(ticketOne, 1)
	ticketWithNumberTwo := MockTicketNumber(ticketTwo, 2)
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

	api.GetPlayersInfoHandler(responseRecorder, request)

	response := responseRecorder.Result()
	body, _ := ioutil.ReadAll(response.Body)

	if string(expectedResponseString) != string(body) {
		t.Errorf("Should be %s but it got %s", string(expectedResponseString), string(body))
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
