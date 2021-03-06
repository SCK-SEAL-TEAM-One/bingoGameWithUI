package api_test

import (
	. "bingo/api"
	"bingo/bingogame"
	"bingo/service"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func Test_StartGameHandler(t *testing.T) {
	requestPlayers := StartGameRequest{
		PlayerNames: []string{"A", "B"},
	}
	requestByte, _ := json.Marshal(requestPlayers)
	expectedRespondstatus := 200
	request := httptest.NewRequest("POST", "/bingo/start", bytes.NewBuffer(requestByte))
	responseRecorder := httptest.NewRecorder()
	api := Api{GameService: &service.GameService{}}

	api.StartGameHandler(responseRecorder, request)

	response := responseRecorder.Result()
	if expectedRespondstatus != response.StatusCode {
		t.Errorf("expect %d but got %d", expectedRespondstatus, response.StatusCode)
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
	expectedResponse := service.PlayerInfoResponse{
		Players:       []bingogame.Player{playerOne, playerTwo},
		HistoryPickUp: []int{},
	}
	expectedResponseString, _ := json.Marshal(expectedResponse)
	numberBox := []int{9, 2, 1, 3, 7, 3, 6, 2, 3, 6}
	api := Api{
		GameService: &service.GameService{
			Game: bingogame.Game{
				Players: []bingogame.Player{
					playerOne,
					playerTwo,
				},
				NumberBox:     numberBox,
				HistoryPickUp: []int{},
			},
		},
	}
	api.GetPlayersInfoHandler(responseRecorder, request)

	response := responseRecorder.Result()
	body, _ := ioutil.ReadAll(response.Body)

	if string(expectedResponseString) != string(body) {
		t.Errorf("Should be \n%s but it got \n%s", string(expectedResponseString), string(body))
	}

}
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
	gameService := service.GameService{
		Game: bingogame.Game{
			Players: []bingogame.Player{
				playerOne,
				playerTwo,
			},
			NumberBox:     numberBox,
			HistoryPickUp: []int{},
		},
	}

	api := Api{
		GameService: &gameService,
	}
	api.PlayHandler(responseRecorder, request)

	response := responseRecorder.Result()
	body, _ := ioutil.ReadAll(response.Body)
	var actual bingogame.PlayResponse
	_ = json.Unmarshal(body, &actual)

	if expected != actual {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}

func Test_ChangeTicketHandler_Input_Name_A_Should_Be_New_Ticket(t *testing.T) {
	request := httptest.NewRequest("GET", "/bingo/ticket/change?playerName=A", nil)
	responseRecorder := httptest.NewRecorder()
	ticket := bingogame.NewTicket(5)
	oldTicket := bingogame.MockTicketNumber(ticket, 1)
	ticket = bingogame.MockTicketNumber(ticket, 3)
	expectedPlayer := bingogame.Player{
		Name:   "A",
		Ticket: ticket,
	}
	gameService := service.MockGameService{
		Game: bingogame.Game{
			Players: []bingogame.Player{
				bingogame.Player{Name: "A", Ticket: oldTicket},
			},
		},
	}
	api := Api{
		GameService: &gameService,
	}

	api.ChangeTicketHandler(responseRecorder, request)
	response := responseRecorder.Result()
	body, _ := ioutil.ReadAll(response.Body)
	var actualPlayer bingogame.Player
	_ = json.Unmarshal(body, &actualPlayer)
	if expectedPlayer.Name != actualPlayer.Name ||
		(expectedPlayer.Ticket.Grid[0][0] != actualPlayer.Ticket.Grid[0][0] ||
			expectedPlayer.Ticket.Grid[4][4] != actualPlayer.Ticket.Grid[4][4] ||
			expectedPlayer.Ticket.Grid[0][4] != actualPlayer.Ticket.Grid[0][4] ||
			expectedPlayer.Ticket.Grid[4][0] != actualPlayer.Ticket.Grid[4][0]) {
		t.Errorf("expected %v but go %v", expectedPlayer.Name, actualPlayer.Name)
	}
}
