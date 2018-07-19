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

func Test_PlayHandle_Should_Be_Json_number_9_winner_Empty(t *testing.T) {
	expected := PlayResponse{
		Number: 9,
		Winner: "",
	}

	rwq := httptest.NewRequest("GET", "localhost:3000/bigo/play", nil)
	w := httptest.NewRecorder()

	PlayHandler(w, rwq)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var actaul PlayResponse
	json.Unmarshal(body, &actaul)

	if expected != actaul {
		t.Errorf("expected %v but got %v", expected, actaul)
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

func Test_StartGameHandler(t *testing.T) {
	request := StartGameRequest{PlayerOne: "A", PlayerTwo: "B"}
	requestByte, _ := json.Marshal(request)

	req := httptest.NewRequest("POST", "/bingo/start", bytes.NewBuffer(requestByte))
	w := httptest.NewRecorder()
	api := Api{}
	api.StartGameHandler(w, req)
	expectedRespondstatus := 200
	response := w.Result()
	if expectedRespondstatus != response.StatusCode {
		t.Errorf("expect %d but got %d", expectedRespondstatus, response.StatusCode)
	}
}
