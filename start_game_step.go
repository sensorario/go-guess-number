package main

import "fmt"

type StartGameStep struct{}

func (s *StartGameStep) play(k *Game) bool {
	k.nextStep = &StartTurnStep{}
	fmt.Println("")
	return true
}

func (s *StartGameStep) name() string {
	return "game started"
}
