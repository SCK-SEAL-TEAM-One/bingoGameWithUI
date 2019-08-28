package bingogame_test

import (
	. "bingo/bingogame"
	"testing"
)

func Test_Shuffle_Input_1_And_75_Should_Be_Number_1_To_75_Are_Shuffled(t *testing.T) {
	startNumber := 1
	endNumber := 75
	expectedlength := 75
	actualNumbers := Shuffle(startNumber, endNumber)

	if expectedlength != len(actualNumbers) {
		t.Errorf("expected Array Length  %v but got %v", expectedlength, len(actualNumbers))
	}

	shuffle := map[int]bool{}
	for index := 0; index < endNumber; index++ {
		value := actualNumbers[index]
		if shuffle[value] {
			t.Errorf("expected Array is Unique")
			break
		}
		shuffle[value] = true
	}

}
