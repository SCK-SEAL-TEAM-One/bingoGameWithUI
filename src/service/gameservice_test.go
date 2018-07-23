package service_test

import (
	. "service"
	"testing"
)

func Test_NewGame_Input_PlayerNameOne_A_And_PlayerNameTwo_B_Should_Be_Error_Nil(t *testing.T) {
	playerNameOne := "A"
	playerNameTwo := "B"
	gameService := GameService{}

	actual := gameService.NewGame(playerNameOne, playerNameTwo)

	if actual != nil && gameService.Game.Players[0].Name != playerNameOne && gameService.Game.Players[1].Name != playerNameTwo {
		t.Errorf("expectedError nil but got %v", actual)
	}
}
