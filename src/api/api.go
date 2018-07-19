package api

import (
	"bingogame"
	"encoding/json"
	"net/http"
)

type Api struct {
	Game bingogame.Game
}

type PlayerInfoResponse struct {
	PlayerOne     bingogame.Player `json:"playerOne"`
	PlayerTwo     bingogame.Player `json:"playerTwo"`
	HistoryPickUp []int            `json:"historyPickUp"`
}

func (a Api) GetPlayersInfoHandler(writer http.ResponseWriter, request *http.Request) {
	ticketBlankPlayerA := bingogame.NewTicket(5)
	ticketBlankPlayerB := bingogame.NewTicket(5)
	playerOne := a.Game.Players[0]
	playerOne.Ticket = MockTicketNumber(ticketBlankPlayerA, 1)
	playerTwo := a.Game.Players[1]
	playerTwo.Ticket = MockTicketNumber(ticketBlankPlayerB, 2)

	playerInfoResponse := PlayerInfoResponse{
		PlayerOne:     a.Game.Players[0],
		PlayerTwo:     a.Game.Players[1],
		HistoryPickUp: a.Game.HistoryPickUp,
	}
	playerInfoResponseJson, _ := json.Marshal(playerInfoResponse)
	writer.Write(playerInfoResponseJson)
}

func MockTicketNumber(ticket bingogame.Ticket, mockId int) bingogame.Ticket {
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
