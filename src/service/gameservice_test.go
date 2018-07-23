package service_test

import (
	"api"
	"bingogame"
	"service"
	"testing"
)

func Test_GetPlayerInfo_Should_Be_PlayerInfoResponse(t *testing.T) {
	ticketOne := bingogame.NewTicket(5)
	ticketTwo := bingogame.NewTicket(5)
	ticketWithNumberOne := api.MockTicketNumber(ticketOne, 1)
	ticketWithNumberTwo := api.MockTicketNumber(ticketTwo, 2)
	playerOne := bingogame.NewPlayer("A", ticketWithNumberOne)
	playerTwo := bingogame.NewPlayer("B", ticketWithNumberTwo)
	expectedResponse := api.PlayerInfoResponse{
		PlayerOne:     playerOne,
		PlayerTwo:     playerTwo,
		HistoryPickUp: []int{},
	}
	numberBox := []int{9, 2, 1, 3, 7, 3, 6, 2, 3, 6}
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

	actual := gameService.GetPlayerInfo()

	if expectedResponse.PlayerOne.Name != actual.PlayerOne.Name &&
		expectedResponse.PlayerTwo.Name != actual.PlayerTwo.Name {
		t.Errorf("expected %v but got %v", expectedResponse, actual)
	}
}
