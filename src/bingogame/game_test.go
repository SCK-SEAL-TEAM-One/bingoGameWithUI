package bingogame_test

import (
	. "bingogame"
	"testing"
)

func Test_NewNumberBox_Input_75_Should_Be_Number_1_To_75(t *testing.T) {
	number := 75
	expectedlength := 75
	actualNumbers := NewNumberBox(number)

	if expectedlength != len(actualNumbers) {
		t.Errorf("expected Array Length  %v but got %v", expectedlength, len(actualNumbers))
	}

	unique := map[int]bool{}
	for index := 0; index < number; index++ {
		value := actualNumbers[index]
		if unique[value] {
			t.Errorf("expected Array is Unique")
			break
		}
		unique[value] = true
	}
}

func Test_NewGame(t *testing.T) {
	newTicketA := NewTicket(5)
	newTicketB := NewTicket(5)
	newPlayerA := NewPlayer("A", newTicketA)
	newPlayerB := NewPlayer("B", newTicketB)
	allPlayer := []Player{newPlayerA, newPlayerB}
	newNumberBox := NewNumberBox(75)
	expected := Game{Players: allPlayer, NumberBox: newNumberBox}

	actualNewGame := NewGame(allPlayer, newNumberBox)

	for indexPlayer := 0; indexPlayer < len(actualNewGame.Players); indexPlayer++ {
		if expected.Players[indexPlayer].Name != actualNewGame.Players[indexPlayer].Name {
			t.Errorf("Should be %v but got %v", expected.Players[indexPlayer].Name, actualNewGame.Players[indexPlayer].Name)
		}
	}
	if len(expected.NumberBox) != len(actualNewGame.NumberBox) {
		t.Errorf("Should be %d but got %d", len(expected.NumberBox), len(actualNewGame.NumberBox))
	}
}

func Test_PickUpNumber_Should_Be_Number_In_NumberBox(t *testing.T) {
	numberbox := []int{9, 51, 47, 29, 56, 49, 39, 58}
	game := Game{NumberBox: numberbox}
	for index := 0; index < len(numberbox); index++ {
		pickUpNumber := game.PickUpNumber()
		if numberbox[index] != pickUpNumber {
			t.Errorf("expected is %v but got %v", numberbox[index], pickUpNumber)
		}
	}
}
