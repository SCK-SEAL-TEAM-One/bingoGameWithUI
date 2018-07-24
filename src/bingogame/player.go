package bingogame

func NewPlayer(name string, ticket Ticket) Player {
	return Player{
		Name:   name,
		Ticket: ticket,
	}
}

func (p *Player) Mark(positionX, positionY int) bool {
	positionX--
	positionY--
	p.Ticket.SetGridStatus(positionX, positionY, true)

	if p.Ticket.Grid[positionX][positionY].Status == false {
		return false
	}
	return true
}

func (p Player) CheckNumber(number int) (int, int) {
	for rowIndex := range p.Ticket.Grid {
		for columnIndex := range p.Ticket.Grid[rowIndex] {
			if number == p.Ticket.GetGridNumber(rowIndex, columnIndex) {
				return rowIndex + 1, columnIndex + 1
			}
		}
	}
	return -1, -1
}

func (p Player) GetBingo(positionX, positionY int) bool {
	return p.CheckHorizontal(positionX) || p.CheckDiagonal(positionX, positionY) || p.CheckVertical(positionY)
}

func (p Player) CheckVertical(positionY int) bool {
	positionY--
	var number int
	for rowIndex := range p.Ticket.Grid {
		if p.Ticket.GetGridStatus(rowIndex, positionY) {
			number++
			if number == p.Ticket.SizeX {
				return true
			}
		}
	}
	return false
}

func (p Player) CheckHorizontal(positionX int) bool {
	var number int
	positionX--
	for columnIndex := 0; columnIndex < p.Ticket.SizeY; columnIndex++ {
		if p.Ticket.GetGridStatus(positionX, columnIndex) {
			number++
			if number == p.Ticket.SizeY {
				return true
			}
		}
	}
	return false
}

func (p Player) CheckDiagonal(positionX, positionY int) bool {
	isDiagonalLeftToRight := true
	isDiagonalRightToLeft := true
	for indexRowColumn := 0; indexRowColumn < 5; indexRowColumn++ {
		if !p.Ticket.GetGridStatus(indexRowColumn, indexRowColumn) {
			isDiagonalLeftToRight = false
			break
		}
	}
	for indexRow, indexColumn := 4, 0; indexRow >= 0; indexRow, indexColumn = indexRow-1, indexColumn+1 {
		if !p.Ticket.GetGridStatus(indexRow, indexColumn) {
			isDiagonalRightToLeft = false
			break
		}
	}
	return isDiagonalLeftToRight || isDiagonalRightToLeft
}
