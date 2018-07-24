package bingogame

const half = 2

func NewTicket(numberOfGrid int) Ticket {
	grid := make([][]State, numberOfGrid)
	for index := 0; index < numberOfGrid; index++ {
		grid[index] = make([]State, numberOfGrid)
	}
	centerPosition := numberOfGrid / half
	grid[centerPosition][centerPosition].Status = true
	return Ticket{
		SizeX: numberOfGrid,
		SizeY: numberOfGrid,
		Grid:  grid,
	}
}

func GenerateTicketNumber(ticket Ticket) Ticket {
	centerColumn := ticket.SizeY / half
	centerRow := ticket.SizeX / half

	for indexRow := range ticket.Grid {
		startNumber := 1 + (indexRow * 15)
		endNumber := 15 + (indexRow * 15)
		shuffleNumber := Shuffle(startNumber, endNumber)
		for indexColumn := range ticket.Grid[indexRow] {
			var value int
			value, shuffleNumber = shuffleNumber[0], shuffleNumber[1:]
			if !(indexRow == centerRow && indexColumn == centerColumn) {
				ticket.SetGridNumber(indexColumn, indexRow, value)
			}
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
