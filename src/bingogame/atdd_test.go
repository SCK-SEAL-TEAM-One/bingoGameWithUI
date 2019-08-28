package bingogame_test

import (
	. "bingo/bingogame"
	"testing"
)

func Test_AcceptanceTest_VerticalBingoRule_Input_Player_A_And_B_PlayRound_8_Should_Be_Player_A_Bingo_With_Number_51_47_56_49_58(t *testing.T) {
	ticketBlankPlayerA := NewTicket(5)
	ticketBlankPlayerB := NewTicket(5)
	ticketWithNumberA := MockTicketNumber(ticketBlankPlayerA, 1)
	ticketWithNumberB := MockTicketNumber(ticketBlankPlayerB, 2)
	playerA := NewPlayer("A", ticketWithNumberA)
	playerB := NewPlayer("B", ticketWithNumberB)
	numberBox := MockNumberBox()
	allPlayer := []Player{playerA, playerB}
	game := NewGame(allPlayer, numberBox)
	game.Play()
	game.Play()
	game.Play()
	game.Play()
	game.Play()
	game.Play()
	game.Play()
	actualResult := game.Play()

	expectedBingoPlayer := "A"
	actualBingoPlayer := actualResult.Winner

	if expectedBingoPlayer != actualBingoPlayer {
		t.Errorf("expected player is %s but it got %s", expectedBingoPlayer, actualBingoPlayer)
	}
}

func Test_AcceptanceTest_VerticalCenterBingoRule_Input_Player_A_And_B_PlayRound_6_Should_Be_Player_B_Bingo_With_Number_45_14_41_32_11_36(t *testing.T) {
	ticketBlankPlayerA := NewTicket(5)
	ticketBlankPlayerB := NewTicket(5)
	ticketWithNumberA := MockTicketNumber(ticketBlankPlayerA, 4)
	ticketWithNumberB := MockTicketNumber(ticketBlankPlayerB, 3)
	playerA := NewPlayer("A", ticketWithNumberA)
	playerB := NewPlayer("B", ticketWithNumberB)
	numberBox := MockNumberBoxForCase2()
	allPlayer := []Player{playerA, playerB}
	game := NewGame(allPlayer, numberBox)
	game.Play()
	game.Play()
	game.Play()
	game.Play()
	game.Play()
	actualResult := game.Play()
	expectedBingoPlayer := "B"
	actualBingoPlayer := actualResult.Winner

	if expectedBingoPlayer != actualBingoPlayer {
		t.Errorf("expected player is %s but it got %s", expectedBingoPlayer, actualBingoPlayer)
	}

}
