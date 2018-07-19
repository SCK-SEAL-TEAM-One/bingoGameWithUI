package api_test

import (
	. "api"
	"bingogame"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

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
	api := Api{
		Game: bingogame.Game{
			Players: []bingogame.Player{
				playerOne,
				playerTwo,
			},
		},
	}

	api.GetPlayersInfoHandler(responseRecorder, request)
	response := responseRecorder.Result()
	body, _ := ioutil.ReadAll(response.Body)
	var actualResponse PlayerInfoResponse
	json.Unmarshal(body, &actualResponse)

	if expectedResponse.PlayerOne.Name != actualResponse.PlayerOne.Name {
		t.Errorf("Should be %s but it got %s", expectedResponse.PlayerOne.Name, actualResponse.PlayerOne.Name)
	}

}
