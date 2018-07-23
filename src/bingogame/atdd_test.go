package bingogame_test

import (
	. "bingogame"
	"testing"
)

func Test_AcceptanceTest_VerticalBingoRule_Input_Player_A_And_B_PlayRound_8_Should_Be_Player_A_Bingo_With_Number_51_47_56_49_58(t *testing.T) {
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
	bingoPlayer := ""
	pickupNumber := game.PickUpNumber()
	positionXPlayer1, positionYPlayer1 := game.Players[0].CheckNumber(pickupNumber)
	if positionXPlayer1 != -1 && positionYPlayer1 != -1 {
		game.Players[0].Mark(positionXPlayer1, positionYPlayer1)
		if game.Players[0].GetBingo(positionXPlayer1, positionYPlayer1) {
			bingoPlayer = game.Players[0].Name
		}
	}
	positionXPlayer2, positionYPlayer2 := game.Players[1].CheckNumber(pickupNumber)
	if positionXPlayer2 != -1 && positionYPlayer2 != -1 {
		game.Players[1].Mark(positionXPlayer2, positionYPlayer2)
		if game.Players[1].GetBingo(positionXPlayer2, positionYPlayer2) {
			bingoPlayer = game.Players[1].Name
		}
	}
	pickupNumber = game.PickUpNumber()
	positionXPlayer1, positionYPlayer1 = game.Players[0].CheckNumber(pickupNumber)
	if positionXPlayer1 != -1 && positionYPlayer1 != -1 {
		game.Players[0].Mark(positionXPlayer1, positionYPlayer1)
		if game.Players[0].GetBingo(positionXPlayer1, positionYPlayer1) {
			bingoPlayer = game.Players[0].Name
		}
	}
	positionXPlayer2, positionYPlayer2 = game.Players[1].CheckNumber(pickupNumber)
	if positionXPlayer2 != -1 && positionYPlayer2 != -1 {
		game.Players[1].Mark(positionXPlayer2, positionYPlayer2)
		if game.Players[1].GetBingo(positionXPlayer2, positionYPlayer2) {
			bingoPlayer = game.Players[1].Name
		}
	}
	pickupNumber = game.PickUpNumber()
	positionXPlayer1, positionYPlayer1 = game.Players[0].CheckNumber(pickupNumber)
	if positionXPlayer1 != -1 && positionYPlayer1 != -1 {
		game.Players[0].Mark(positionXPlayer1, positionYPlayer1)
		if game.Players[0].GetBingo(positionXPlayer1, positionYPlayer1) {
			bingoPlayer = game.Players[0].Name
		}
	}
	positionXPlayer2, positionYPlayer2 = game.Players[1].CheckNumber(pickupNumber)
	if positionXPlayer2 != -1 && positionYPlayer2 != -1 {
		game.Players[1].Mark(positionXPlayer2, positionYPlayer2)
		if game.Players[1].GetBingo(positionXPlayer2, positionYPlayer2) {
			bingoPlayer = game.Players[1].Name
		}
	}
	pickupNumber = game.PickUpNumber()
	positionXPlayer1, positionYPlayer1 = game.Players[0].CheckNumber(pickupNumber)
	if positionXPlayer1 != -1 && positionYPlayer1 != -1 {
		game.Players[0].Mark(positionXPlayer1, positionYPlayer1)
		if game.Players[0].GetBingo(positionXPlayer1, positionYPlayer1) {
			bingoPlayer = game.Players[0].Name
		}
	}
	positionXPlayer2, positionYPlayer2 = game.Players[1].CheckNumber(pickupNumber)
	if positionXPlayer2 != -1 && positionYPlayer2 != -1 {
		game.Players[1].Mark(positionXPlayer2, positionYPlayer2)
		if game.Players[1].GetBingo(positionXPlayer2, positionYPlayer2) {
			bingoPlayer = game.Players[1].Name
		}
	}
	pickupNumber = game.PickUpNumber()
	positionXPlayer1, positionYPlayer1 = game.Players[0].CheckNumber(pickupNumber)
	if positionXPlayer1 != -1 && positionYPlayer1 != -1 {
		game.Players[0].Mark(positionXPlayer1, positionYPlayer1)
		if game.Players[0].GetBingo(positionXPlayer1, positionYPlayer1) {
			bingoPlayer = game.Players[0].Name
		}
	}
	positionXPlayer2, positionYPlayer2 = game.Players[1].CheckNumber(pickupNumber)
	if positionXPlayer2 != -1 && positionYPlayer2 != -1 {
		game.Players[1].Mark(positionXPlayer2, positionYPlayer2)
		if game.Players[1].GetBingo(positionXPlayer2, positionYPlayer2) {
			bingoPlayer = game.Players[1].Name
		}
	}
	pickupNumber = game.PickUpNumber()
	positionXPlayer1, positionYPlayer1 = game.Players[0].CheckNumber(pickupNumber)
	if positionXPlayer1 != -1 && positionYPlayer1 != -1 {
		game.Players[0].Mark(positionXPlayer1, positionYPlayer1)
		if game.Players[0].GetBingo(positionXPlayer1, positionYPlayer1) {
			bingoPlayer = game.Players[0].Name
		}
	}
	positionXPlayer2, positionYPlayer2 = game.Players[1].CheckNumber(pickupNumber)
	if positionXPlayer2 != -1 && positionYPlayer2 != -1 {
		game.Players[1].Mark(positionXPlayer2, positionYPlayer2)
		if game.Players[1].GetBingo(positionXPlayer2, positionYPlayer2) {
			bingoPlayer = game.Players[1].Name
		}
	}
	pickupNumber = game.PickUpNumber()
	positionXPlayer1, positionYPlayer1 = game.Players[0].CheckNumber(pickupNumber)
	if positionXPlayer1 != -1 && positionYPlayer1 != -1 {
		game.Players[0].Mark(positionXPlayer1, positionYPlayer1)
		if game.Players[0].GetBingo(positionXPlayer1, positionYPlayer1) {
			bingoPlayer = game.Players[0].Name
		}
	}
	positionXPlayer2, positionYPlayer2 = game.Players[1].CheckNumber(pickupNumber)
	if positionXPlayer2 != -1 && positionYPlayer2 != -1 {
		game.Players[1].Mark(positionXPlayer2, positionYPlayer2)
		if game.Players[1].GetBingo(positionXPlayer2, positionYPlayer2) {
			bingoPlayer = game.Players[1].Name
		}
	}
	pickupNumber = game.PickUpNumber()
	positionXPlayer1, positionYPlayer1 = game.Players[0].CheckNumber(pickupNumber)
	if positionXPlayer1 != -1 && positionYPlayer1 != -1 {
		game.Players[0].Mark(positionXPlayer1, positionYPlayer1)
		if game.Players[0].GetBingo(positionXPlayer1, positionYPlayer1) {
			bingoPlayer = game.Players[0].Name
		}
	}
	positionXPlayer2, positionYPlayer2 = game.Players[1].CheckNumber(pickupNumber)
	if positionXPlayer2 != -1 && positionYPlayer2 != -1 {
		game.Players[1].Mark(positionXPlayer2, positionYPlayer2)
		if game.Players[1].GetBingo(positionXPlayer2, positionYPlayer2) {
			bingoPlayer = game.Players[1].Name
		}
	}

	expectedBingoPlayer := "A"
	actualBingoPlayer := bingoPlayer

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
	numberBox := NewNumberBox(75)
	numberBox = MockNumberBoxForCase2()
	allPlayer := []Player{playerA, playerB}
	game := NewGame(allPlayer, numberBox)
	bingoPlayer := ""
	pickupNumber := game.PickUpNumber()
	positionXPlayer1, positionYPlayer1 := game.Players[0].CheckNumber(pickupNumber)
	if positionXPlayer1 != -1 && positionYPlayer1 != -1 {
		game.Players[0].Mark(positionXPlayer1, positionYPlayer1)
		if game.Players[0].GetBingo(positionXPlayer1, positionYPlayer1) {
			bingoPlayer = game.Players[0].Name
		}
	}
	positionXPlayer2, positionYPlayer2 := game.Players[1].CheckNumber(pickupNumber)
	if positionXPlayer2 != -1 && positionYPlayer2 != -1 {
		game.Players[1].Mark(positionXPlayer2, positionYPlayer2)
		if game.Players[1].GetBingo(positionXPlayer2, positionYPlayer2) {
			bingoPlayer = game.Players[1].Name
		}
	}
	pickupNumber = game.PickUpNumber()
	positionXPlayer1, positionYPlayer1 = game.Players[0].CheckNumber(pickupNumber)
	if positionXPlayer1 != -1 && positionYPlayer1 != -1 {
		game.Players[0].Mark(positionXPlayer1, positionYPlayer1)
		if game.Players[0].GetBingo(positionXPlayer1, positionYPlayer1) {
			bingoPlayer = game.Players[0].Name
		}
	}
	positionXPlayer2, positionYPlayer2 = game.Players[1].CheckNumber(pickupNumber)
	if positionXPlayer2 != -1 && positionYPlayer2 != -1 {
		game.Players[1].Mark(positionXPlayer2, positionYPlayer2)
		if game.Players[1].GetBingo(positionXPlayer2, positionYPlayer2) {
			bingoPlayer = game.Players[1].Name
		}
	}
	pickupNumber = game.PickUpNumber()
	positionXPlayer1, positionYPlayer1 = game.Players[0].CheckNumber(pickupNumber)
	if positionXPlayer1 != -1 && positionYPlayer1 != -1 {
		game.Players[0].Mark(positionXPlayer1, positionYPlayer1)
		if game.Players[0].GetBingo(positionXPlayer1, positionYPlayer1) {
			bingoPlayer = game.Players[0].Name
		}
	}
	positionXPlayer2, positionYPlayer2 = game.Players[1].CheckNumber(pickupNumber)
	if positionXPlayer2 != -1 && positionYPlayer2 != -1 {
		game.Players[1].Mark(positionXPlayer2, positionYPlayer2)
		if game.Players[1].GetBingo(positionXPlayer2, positionYPlayer2) {
			bingoPlayer = game.Players[1].Name
		}
	}
	pickupNumber = game.PickUpNumber()
	positionXPlayer1, positionYPlayer1 = game.Players[0].CheckNumber(pickupNumber)
	if positionXPlayer1 != -1 && positionYPlayer1 != -1 {
		game.Players[0].Mark(positionXPlayer1, positionYPlayer1)
		if game.Players[0].GetBingo(positionXPlayer1, positionYPlayer1) {
			bingoPlayer = game.Players[0].Name
		}
	}
	positionXPlayer2, positionYPlayer2 = game.Players[1].CheckNumber(pickupNumber)
	if positionXPlayer2 != -1 && positionYPlayer2 != -1 {
		game.Players[1].Mark(positionXPlayer2, positionYPlayer2)
		if game.Players[1].GetBingo(positionXPlayer2, positionYPlayer2) {
			bingoPlayer = game.Players[1].Name
		}
	}
	pickupNumber = game.PickUpNumber()
	positionXPlayer1, positionYPlayer1 = game.Players[0].CheckNumber(pickupNumber)
	if positionXPlayer1 != -1 && positionYPlayer1 != -1 {
		game.Players[0].Mark(positionXPlayer1, positionYPlayer1)
		if game.Players[0].GetBingo(positionXPlayer1, positionYPlayer1) {
			bingoPlayer = game.Players[0].Name
		}
	}
	positionXPlayer2, positionYPlayer2 = game.Players[1].CheckNumber(pickupNumber)
	if positionXPlayer2 != -1 && positionYPlayer2 != -1 {
		game.Players[1].Mark(positionXPlayer2, positionYPlayer2)
		if game.Players[1].GetBingo(positionXPlayer2, positionYPlayer2) {
			bingoPlayer = game.Players[1].Name
		}
	}
	pickupNumber = game.PickUpNumber()
	positionXPlayer1, positionYPlayer1 = game.Players[0].CheckNumber(pickupNumber)
	if positionXPlayer1 != -1 && positionYPlayer1 != -1 {
		game.Players[0].Mark(positionXPlayer1, positionYPlayer1)
		if game.Players[0].GetBingo(positionXPlayer1, positionYPlayer1) {
			bingoPlayer = game.Players[0].Name
		}
	}
	positionXPlayer2, positionYPlayer2 = game.Players[1].CheckNumber(pickupNumber)
	if positionXPlayer2 != -1 && positionYPlayer2 != -1 {
		game.Players[1].Mark(positionXPlayer2, positionYPlayer2)
		if game.Players[1].GetBingo(positionXPlayer2, positionYPlayer2) {
			bingoPlayer = game.Players[1].Name
		}
	}
	expectedBingoPlayer := "B"
	actualBingoPlayer := bingoPlayer

	if expectedBingoPlayer != actualBingoPlayer {
		t.Errorf("expected player is %s but it got %s", expectedBingoPlayer, actualBingoPlayer)
	}

}
