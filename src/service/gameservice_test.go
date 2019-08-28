package service_test

import (
	"bingo/bingogame"
	"bingo/service"
	"testing"
)

func Test_GetPlayerInfo_Should_Be_PlayerInfoResponse(t *testing.T) {
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

	if expectedResponse.Players[0].Name != actual.Players[0].Name &&
		expectedResponse.Players[1].Name != actual.Players[1].Name {
		t.Errorf("expected %v but got %v", expectedResponse, actual)

	}
}

func Test_NewGame_Input_PlayerNameOne_A_And_PlayerNameTwo_B_Should_Be_Error_Nil(t *testing.T) {
	playerNames := []string{"A", "B"}
	gameService := service.GameService{}

	actual := gameService.NewGame(playerNames)

	if actual != nil &&
		gameService.Game.Players[0].Name != playerNames[0] &&
		gameService.Game.Players[1].Name != playerNames[1] {
		t.Errorf("expectedError nil but got %v", actual)
	}
}
func Test_PlayGame_Should_Be_Number_9_Winer_Empty(t *testing.T) {

	ticketBlankPlayerA := bingogame.NewTicket(5)
	ticketBlankPlayerB := bingogame.NewTicket(5)
	ticketWithNumberA := bingogame.MockTicketNumber(ticketBlankPlayerA, 1)
	ticketWithNumberB := bingogame.MockTicketNumber(ticketBlankPlayerB, 2)
	playerA := bingogame.NewPlayer("A", ticketWithNumberA)
	playerB := bingogame.NewPlayer("B", ticketWithNumberB)
	numberBox := bingogame.MockNumberBox()
	allPlayer := []bingogame.Player{playerA, playerB}
	game := bingogame.NewGame(allPlayer, numberBox)
	gameService := service.GameService{Game: game}
	expected := bingogame.PlayResponse{
		Number: 9,
		Winner: "",
	}

	actual := gameService.PlayGame()

	if expected != actual {
		t.Errorf("expect %v but got %v", expected, actual)
	}
}
func Test_ChangeTicket_Input_PlayerName_A_Should_Be_Player_A_With_New_Ticket(t *testing.T) {

	ticketBlankPlayerA := bingogame.NewTicket(5)
	ticketBlankPlayerB := bingogame.NewTicket(5)
	oldTicket := bingogame.NewTicket(5)
	oldTicket = bingogame.MockTicketNumber(oldTicket, 1)
	ticketWithNumberA := bingogame.MockTicketNumber(ticketBlankPlayerA, 1)
	ticketWithNumberB := bingogame.MockTicketNumber(ticketBlankPlayerB, 2)
	playerA := bingogame.NewPlayer("A", ticketWithNumberA)
	playerB := bingogame.NewPlayer("B", ticketWithNumberB)
	numberBox := bingogame.MockNumberBox()
	allPlayer := []bingogame.Player{playerA, playerB}
	game := bingogame.NewGame(allPlayer, numberBox)
	gameService := service.GameService{Game: game}
	actual := gameService.ChangeTicket("A")

	if oldTicket.Grid[0][0] == actual.Ticket.Grid[0][0] &&
		oldTicket.Grid[0][4] == actual.Ticket.Grid[0][4] &&
		oldTicket.Grid[4][0] == actual.Ticket.Grid[4][0] &&
		oldTicket.Grid[4][4] == actual.Ticket.Grid[4][4] {
		t.Errorf("expect \n%v but got \n%v", oldTicket, actual)
	}
}
