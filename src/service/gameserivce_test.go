package service_test

import (
	. "bingogame"
	. "service"
	"testing"
)

func Test_PlayGame_Should_Be_Number_9_Winer_Empty(t *testing.T) {

	ticketBlankPlayerA := NewTicket(5)
	ticketBlankPlayerB := NewTicket(5)
	ticketWithNumberA := MockTicketNumber(ticketBlankPlayerA, 1)
	ticketWithNumberB := MockTicketNumber(ticketBlankPlayerB, 2)
	playerA := NewPlayer("A", ticketWithNumberA)
	playerB := NewPlayer("B", ticketWithNumberB)
	numberBox := NewNumberBox(75)
	numberBox = MockNumberBox()
	allPlayer := []Player{playerA, playerB}
	game := NewGame(allPlayer, numberBox)

	playGameService := PlayGameService{Game: game}

	expected := PlayResponse{
		Number: 9,
		Winner: "",
	}

	actual := playGameService.PlayGame()

	if expected != actual {
		t.Errorf("expect %v but got %v")
	}
}
