package bingogame

import (
	"math/rand"
	"time"
)

func Shuffle(startNumber, endNumber int) []int {
	length := endNumber - startNumber + 1
	shuffle := make([]int, length)
	for index := 0; index < length; index++ {
		shuffle[index] = index + 1
	}
	for index := len(shuffle) - 1; index > 0; index-- {
		secondsSince1970 := time.Now().UTC().UnixNano()
		rand.Seed(secondsSince1970)
		randomIndex := rand.Intn(index + 1)
		shuffle[index], shuffle[randomIndex] = shuffle[randomIndex], shuffle[index]
	}

	return shuffle

}
