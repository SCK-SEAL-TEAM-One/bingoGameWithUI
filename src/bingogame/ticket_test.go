package bingogame_test

import (
	. "bingogame"
	"testing"
)

func Test_NewTicket_Input_Number_Of_Grid_5_Should_Be_Ticket_With_SizeX_And_SizeY_And_Grid(t *testing.T) {
	numberOfGrid := 5
	grid := make([][]State, numberOfGrid)
	for index := 0; index < numberOfGrid; index++ {
		grid[index] = make([]State, numberOfGrid)
	}
	centerPosition := (numberOfGrid) / 2
	grid[centerPosition][centerPosition].Status = true
	expected := Ticket{SizeX: 5, SizeY: 5, Grid: grid}

	actualTicket := NewTicket(numberOfGrid)

	for indexRow := 0; indexRow < 5; indexRow++ {
		for indexColumn := 0; indexColumn < 5; indexColumn++ {
			if expected.Grid[indexRow][indexColumn] != actualTicket.Grid[indexRow][indexColumn] {
				t.Errorf("expect %v but got %v", expected, actualTicket)
			}
		}
	}
	if actualTicket.SizeX != expected.SizeX || actualTicket.SizeY != expected.SizeY {
		t.Errorf("expect %v but got %v", expected, actualTicket)
	}
}

func Test_GenerateTicketNumber_Input_BlankTicket_Should_Be_TicketWithNumber(t *testing.T) {
	numberOfGrid := 5
	grid := make([][]State, numberOfGrid)
	for index := 0; index < numberOfGrid; index++ {
		grid[index] = make([]State, numberOfGrid)
	}
	centerPosition := (numberOfGrid) / 2
	grid[centerPosition][centerPosition].Status = true
	ticket := Ticket{SizeX: numberOfGrid, SizeY: numberOfGrid, Grid: grid}

	ticket = GenerateTicketNumber(ticket)

	for indexRow := 0; indexRow < numberOfGrid; indexRow++ {
		for indexColumn := 0; indexColumn < numberOfGrid; indexColumn++ {
			if ticket.Grid[indexRow][indexColumn].Number == 0 && indexColumn != 2 && indexRow != 2 {
				t.Errorf("GirdNumber Row %d Column %d is %d", indexRow, indexColumn, ticket.Grid[indexRow][indexColumn].Number)
			}
		}
	}
}

func MockTicketNumber(ticket Ticket, mockId int) Ticket {
	switch mockId {
	case 1:
		ticketDataId := []int{
			1, 17, 35, 51, 74,
			9, 21, 41, 58, 79,
			2, 23, 0, 47, 68,
			14, 29, 32, 49, 66,
			11, 30, 39, 56, 70}
		ticketDataIdIndex := 0
		for indexRow := 0; indexRow < 5; indexRow++ {
			for indexColumn := 0; indexColumn < 5; indexColumn++ {
				ticket.Grid[indexRow][indexColumn].Number = ticketDataId[ticketDataIdIndex]
				ticketDataIdIndex++
			}
		}
	case 2:
		ticketDataId := []int{
			3, 21, 39, 53, 55,
			12, 29, 32, 54, 67,
			11, 30, 0, 49, 70,
			9, 16, 41, 45, 68,
			7, 19, 44, 52, 72}
		ticketDataIdIndex := 0
		for indexRow := 0; indexRow < 5; indexRow++ {
			for indexColumn := 0; indexColumn < 5; indexColumn++ {
				ticket.Grid[indexRow][indexColumn].Number = ticketDataId[ticketDataIdIndex]
				ticketDataIdIndex++
			}
		}
	case 3:
		ticketDataId := []int{
			1, 16, 32, 46, 62,
			8, 21, 45, 55, 70,
			11, 26, 0, 51, 66,
			13, 20, 41, 49, 72,
			9, 23, 36, 59, 61}
		ticketDataIdIndex := 0
		for indexRow := 0; indexRow < 5; indexRow++ {
			for indexColumn := 0; indexColumn < 5; indexColumn++ {
				ticket.Grid[indexRow][indexColumn].Number = ticketDataId[ticketDataIdIndex]
				ticketDataIdIndex++
			}
		}
	case 4:
		ticketDataId := []int{
			2, 17, 31, 49, 71,
			6, 20, 44, 56, 62,
			7, 18, 0, 50, 69,
			10, 28, 40, 47, 73,
			14, 23, 35, 52, 65}
		ticketDataIdIndex := 0
		for indexRow := 0; indexRow < 5; indexRow++ {
			for indexColumn := 0; indexColumn < 5; indexColumn++ {
				ticket.Grid[indexRow][indexColumn].Number = ticketDataId[ticketDataIdIndex]
				ticketDataIdIndex++
			}
		}
	}
	return ticket
}
