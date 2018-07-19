package bingogame_test

import (
	. "bingogame"
	"testing"
)

func Test_NewPlayer_Input_Name_A_With_Ticket_Should_Be_Player_Name_A_With_Ticket(t *testing.T) {
	name := "A"
	ticket := Ticket{
		SizeX: 5,
		SizeY: 5,
		Grid: [][]State{
			[]State{State{Number: 1}, State{Number: 17}, State{Number: 35}, State{Number: 51}, State{Number: 74}},
			[]State{State{Number: 9}, State{Number: 21}, State{Number: 41}, State{Number: 38}, State{Number: 79}},
			[]State{State{Number: 2}, State{Number: 23}, State{Number: 0, Status: true}, State{Number: 47}, State{Number: 68}},
			[]State{State{Number: 14}, State{Number: 29}, State{Number: 32}, State{Number: 49}, State{Number: 66}},
			[]State{State{Number: 11}, State{Number: 30}, State{Number: 39}, State{Number: 56}, State{Number: 70}},
		},
	}
	expectedPlayer := Player{
		Name:   "A",
		Ticket: ticket,
	}
	actualPlayer := NewPlayer(name, ticket)

	if expectedPlayer.Name != actualPlayer.Name {
		t.Errorf("expected player is %v but it got %v", expectedPlayer, actualPlayer)
	}
}

func Test_NewPlayer_Input_Name_B_With_Ticket_Should_Be_Player_Name_B_With_Ticket(t *testing.T) {
	name := "B"
	ticket := Ticket{
		SizeX: 5,
		SizeY: 5,
		Grid: [][]State{
			[]State{State{Number: 3}, State{Number: 21}, State{Number: 39}, State{Number: 53}, State{Number: 75}},
			[]State{State{Number: 12}, State{Number: 29}, State{Number: 32}, State{Number: 54}, State{Number: 67}},
			[]State{State{Number: 11}, State{Number: 30}, State{Number: 0, Status: true}, State{Number: 49}, State{Number: 70}},
			[]State{State{Number: 9}, State{Number: 16}, State{Number: 41}, State{Number: 45}, State{Number: 68}},
			[]State{State{Number: 7}, State{Number: 19}, State{Number: 44}, State{Number: 52}, State{Number: 72}},
		},
	}
	expectedPlayer := Player{
		Name:   "B",
		Ticket: ticket,
	}
	actualPlayer := NewPlayer(name, ticket)

	if expectedPlayer.Name != actualPlayer.Name {
		t.Errorf("expected player is %v but it got %v", expectedPlayer, actualPlayer)
	}
}

func Test_Mark_Input_X_1_Y_4_Should_Be_True(t *testing.T) {
	x := 1
	y := 4
	ticket := Ticket{
		SizeX: 5,
		SizeY: 5,
		Grid: [][]State{
			[]State{State{Number: 3}, State{Number: 21}, State{Number: 39}, State{Number: 53}, State{Number: 75}},
			[]State{State{Number: 12}, State{Number: 29}, State{Number: 32}, State{Number: 54}, State{Number: 67}},
			[]State{State{Number: 11}, State{Number: 30}, State{Number: 0, Status: true}, State{Number: 49}, State{Number: 70}},
			[]State{State{Number: 9}, State{Number: 16}, State{Number: 41}, State{Number: 45}, State{Number: 68}},
			[]State{State{Number: 7}, State{Number: 19}, State{Number: 44}, State{Number: 52}, State{Number: 72}},
		},
	}
	player := Player{
		Name:   "B",
		Ticket: ticket,
	}
	expected := true
	actual := player.Mark(x, y)

	if expected != actual {
		t.Errorf("expect %t but got %t", expected, actual)
	}
}

func Test_CheckNumber_Input_9_Should_Be_2_1(t *testing.T) {
	number := 9
	expectedPositionX := 2
	expectedPositionY := 1
	player := Player{
		Name: "A",
		Ticket: Ticket{
			SizeX: 5,
			SizeY: 5,
			Grid: [][]State{
				[]State{State{Number: 1}, State{Number: 17}, State{Number: 35}, State{Number: 51}, State{Number: 74}},
				[]State{State{Number: 9}, State{Number: 21}, State{Number: 41}, State{Number: 38}, State{Number: 79}},
				[]State{State{Number: 2}, State{Number: 23}, State{Number: 0, Status: true}, State{Number: 47}, State{Number: 68}},
				[]State{State{Number: 14}, State{Number: 29}, State{Number: 32}, State{Number: 49}, State{Number: 66}},
				[]State{State{Number: 11}, State{Number: 30}, State{Number: 39}, State{Number: 56}, State{Number: 70}},
			},
		},
	}
	x, y := player.CheckNumber(number)

	if expectedPositionX != x || expectedPositionY != y {
		t.Errorf("expected Position is %v,%v but got %v,%v", expectedPositionX, expectedPositionY, x, y)
	}

}

func Test_GetBingo_Input_X_1_Y_4_Should_Be_false(t *testing.T) {
	newTicket := NewTicket(5)
	newPlayer := NewPlayer("A", newTicket)
	expectedResult := false

	actualResult := newPlayer.GetBingo(1, 4)

	if expectedResult != actualResult {
		t.Errorf("Should be %t but got %t", expectedResult, actualResult)
	}
}
func Test_CheckVertical_Input_X_1_Y_4_Should_Be_False(t *testing.T) {
	positionX := 1
	positionY := 4
	player := Player{
		Name: "A",
		Ticket: Ticket{
			SizeX: 5,
			SizeY: 5,
			Grid: [][]State{
				[]State{State{Number: 1}, State{Number: 17}, State{Number: 35}, State{Number: 51}, State{Number: 74}},
				[]State{State{Number: 9}, State{Number: 21}, State{Number: 41}, State{Number: 38}, State{Number: 79}},
				[]State{State{Number: 2}, State{Number: 23}, State{Number: 0, Status: true}, State{Number: 47}, State{Number: 68}},
				[]State{State{Number: 14}, State{Number: 29}, State{Number: 32}, State{Number: 49}, State{Number: 66}},
				[]State{State{Number: 11}, State{Number: 30}, State{Number: 39}, State{Number: 56}, State{Number: 70}},
			},
		},
	}
	expectedBingo := false

	actualBingo := player.CheckVertical(positionX, positionY)
	if expectedBingo != actualBingo {
		t.Errorf("expected %v but got %v", expectedBingo, actualBingo)
	}
}

func Test_CheckHorizental_Input_X_1_Y_4_Should_Be_False(t *testing.T) {
	positionX := 1
	positionY := 4
	player := Player{
		Name: "A",
		Ticket: Ticket{
			SizeX: 5,
			SizeY: 5,
			Grid: [][]State{
				[]State{State{Number: 1}, State{Number: 17}, State{Number: 35}, State{Number: 51}, State{Number: 74}},
				[]State{State{Number: 9}, State{Number: 21}, State{Number: 41}, State{Number: 38}, State{Number: 79}},
				[]State{State{Number: 2}, State{Number: 23}, State{Number: 0, Status: true}, State{Number: 47}, State{Number: 68}},
				[]State{State{Number: 14}, State{Number: 29}, State{Number: 32}, State{Number: 49}, State{Number: 66}},
				[]State{State{Number: 11}, State{Number: 30}, State{Number: 39}, State{Number: 56}, State{Number: 70}},
			},
		},
	}
	expectedBingo := false

	actualBingo := player.CheckHorizental(positionX, positionY)

	if expectedBingo != actualBingo {
		t.Errorf("expected %v but got %v", expectedBingo, actualBingo)
	}
}
func Test_CheckHorizental_Input_X_2_Y_4_Should_Be_True(t *testing.T) {
	positionX := 2
	positionY := 4
	player := Player{
		Name: "A",
		Ticket: Ticket{
			SizeX: 5,
			SizeY: 5,
			Grid: [][]State{
				[]State{State{Number: 1}, State{Number: 17}, State{Number: 35}, State{Number: 51}, State{Number: 74}},
				[]State{State{Number: 9, Status: true}, State{Number: 21, Status: true}, State{Number: 41, Status: true}, State{Number: 38, Status: true}, State{Number: 79, Status: true}},
				[]State{State{Number: 2}, State{Number: 23}, State{Number: 0, Status: true}, State{Number: 47}, State{Number: 68}},
				[]State{State{Number: 14}, State{Number: 29}, State{Number: 32}, State{Number: 49}, State{Number: 66}},
				[]State{State{Number: 11}, State{Number: 30}, State{Number: 39}, State{Number: 56}, State{Number: 70}},
			},
		},
	}
	expectedBingo := true
	actualBingo := player.CheckHorizental(positionX, positionY)
	if expectedBingo != actualBingo {
		t.Errorf("expected %v but got %v", expectedBingo, actualBingo)
	}
}

func Test_CheckVertical_Input_X_2_Y_4_Should_Be_True(t *testing.T) {
	positionX := 2
	positionY := 4
	player := Player{
		Name: "A",
		Ticket: Ticket{
			SizeX: 5,
			SizeY: 5,
			Grid: [][]State{
				[]State{State{Number: 1}, State{Number: 17}, State{Number: 35}, State{Number: 51, Status: true}, State{Number: 74}},
				[]State{State{Number: 9}, State{Number: 21}, State{Number: 41}, State{Number: 38, Status: true}, State{Number: 79}},
				[]State{State{Number: 2}, State{Number: 23}, State{Number: 0, Status: true}, State{Number: 47, Status: true}, State{Number: 68}},
				[]State{State{Number: 14}, State{Number: 29}, State{Number: 32}, State{Number: 49, Status: true}, State{Number: 66}},
				[]State{State{Number: 11}, State{Number: 30}, State{Number: 39}, State{Number: 56, Status: true}, State{Number: 70}},
			},
		},
	}
	expectedBingo := true
	actualBingo := player.CheckVertical(positionX, positionY)
	if expectedBingo != actualBingo {
		t.Errorf("expected %v but got %v", expectedBingo, actualBingo)
	}
}
func Test_CheckDiagonal_Input_X_1_Y_4_Should_Be_False(t *testing.T) {
	positionX := 1
	positionY := 4
	player := Player{
		Name: "A",
		Ticket: Ticket{
			SizeX: 5,
			SizeY: 5,
			Grid: [][]State{
				[]State{State{Number: 1}, State{Number: 17}, State{Number: 35}, State{Number: 51}, State{Number: 74}},
				[]State{State{Number: 9}, State{Number: 21}, State{Number: 41}, State{Number: 38}, State{Number: 79}},
				[]State{State{Number: 2}, State{Number: 23}, State{Number: 0, Status: true}, State{Number: 47}, State{Number: 68}},
				[]State{State{Number: 14}, State{Number: 29}, State{Number: 32}, State{Number: 49}, State{Number: 66}},
				[]State{State{Number: 11}, State{Number: 30}, State{Number: 39}, State{Number: 56}, State{Number: 70}},
			},
		},
	}
	expectedBingo := false

	actualBingo := player.CheckDiagonal(positionX, positionY)

	if expectedBingo != actualBingo {
		t.Errorf("expected %v but got %v", expectedBingo, actualBingo)
	}
}

func Test_CheckDiagonal_Input_X_5_Y_5_Should_Be_True(t *testing.T) {
	positionX := 5
	positionY := 5
	player := Player{
		Name: "A",
		Ticket: Ticket{
			SizeX: 5,
			SizeY: 5,
			Grid: [][]State{
				[]State{State{Number: 1, Status: true}, State{Number: 17}, State{Number: 35}, State{Number: 51}, State{Number: 74}},
				[]State{State{Number: 9}, State{Number: 21, Status: true}, State{Number: 41}, State{Number: 38}, State{Number: 79}},
				[]State{State{Number: 2}, State{Number: 23}, State{Number: 0, Status: true}, State{Number: 47}, State{Number: 68}},
				[]State{State{Number: 14}, State{Number: 29}, State{Number: 32}, State{Number: 49, Status: true}, State{Number: 66}},
				[]State{State{Number: 11}, State{Number: 30}, State{Number: 39}, State{Number: 56}, State{Number: 70, Status: true}},
			},
		},
	}
	expectedBingo := true

	actualBingo := player.CheckDiagonal(positionX, positionY)
	if expectedBingo != actualBingo {
		t.Errorf("expected %v but got %v", expectedBingo, actualBingo)
	}
}

func Test_CheckDiagonal_Input_X_1_Y_5_Should_Be_True(t *testing.T) {
	positionX := 5
	positionY := 5
	player := Player{
		Name: "A",
		Ticket: Ticket{
			SizeX: 5,
			SizeY: 5,
			Grid: [][]State{

				[]State{State{Number: 1}, State{Number: 17}, State{Number: 35}, State{Number: 51}, State{Number: 74, Status: true}},
				[]State{State{Number: 9}, State{Number: 21}, State{Number: 41}, State{Number: 38, Status: true}, State{Number: 79}},
				[]State{State{Number: 2}, State{Number: 23}, State{Number: 0, Status: true}, State{Number: 47}, State{Number: 68}},
				[]State{State{Number: 14}, State{Number: 29, Status: true}, State{Number: 32}, State{Number: 49}, State{Number: 66}},
				[]State{State{Number: 11, Status: true}, State{Number: 30}, State{Number: 39}, State{Number: 56}, State{Number: 70}},
			},
		},
	}
	expectedBingo := true

	actualBingo := player.CheckDiagonal(positionX, positionY)

	if expectedBingo != actualBingo {
		t.Errorf("expected %v but got %v", expectedBingo, actualBingo)
	}
}
