package bingogame

func NewTicket(numberOfGrid int) Ticket {
	grid := make([][]State, numberOfGrid)
	for index := 0; index < numberOfGrid; index++ {
		grid[index] = make([]State, numberOfGrid)
	}
	centerPosition := numberOfGrid / 2
	grid[centerPosition][centerPosition].Status = true
	return Ticket{
		SizeX: numberOfGrid,
		SizeY: numberOfGrid,
		Grid:  grid,
	}
}

func GenerateTicketNumber(ticket Ticket) Ticket {
	for indexRow := range ticket.Grid {
		startNumber := 1 + (indexRow * 15)
		endNumber := 15 + (indexRow * 15)
		suffleNumber := Shuffle(startNumber, endNumber)
		for indexColumn := range ticket.Grid[indexRow] {
			var value int
			value, suffleNumber = suffleNumber[0], suffleNumber[1:]
			ticket.SetGridNumber(indexRow, indexColumn, value)
		}
	}
	return ticket
}
func (t *Ticket) SetGridNumber(row, column, value int) {
	t.Grid[row][column].Number = value
}
func (t *Ticket) SetGridStatus(row, column int, status bool) {
	t.Grid[row][column].Status = status
}

func (t Ticket) GetGridNumber(row, column int) int {
	return t.Grid[row][column].Number
}

func (t Ticket) GetGridStatus(Row, Column int) bool {
	return t.Grid[Row][Column].Status
}
