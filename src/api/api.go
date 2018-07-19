package api

import (
	"encoding/json"
	"net/http"
)

type Play struct {
	number int    `json:"number"`
	winner string `json:"winner"`
}

func PlayHendle(w http.ResponseWriter, r *http.Request) {
	num := 9
	playerWin := ""
	playJson, _ := json.Marshal(Play{number: num, winner: playerWin})
	w.Write(playJson)
}
