package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(10)

	game := Game{
		nextStep: &StartGameStep{},
		Number:   number,
	}

	game.logStep()

	for game.nextStep.play(&game) {
		game.logStep()
	}
}
