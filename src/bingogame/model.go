package bingogame

type Ticket struct {
	SizeX int       `json:"sizeX"`
	SizeY int       `json:"sizeY"`
	Grid  [][]State `json:"grid"`
}

type State struct {
	Number int  `json:"number"`
	Status bool `json:"status"`
}

type Player struct {
	Name   string `json:"name"`
	Ticket Ticket `json:"ticket"`
}

type Game struct {
	Players       []Player `json:"players"`
	NumberBox     []int    `json:"numberBox"`
	HistoryPickUp []int    `json:"historyPickUp"`
}

type PlayResponse struct {
	Number int    `json:"number"`
	Winner string `json:"winner"`
}
