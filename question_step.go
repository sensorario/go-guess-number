package main

import (
	"fmt"
	"os"
)

type QuestionStep struct{}

func (s *QuestionStep) play(k *Game) bool {
	var n int
	fmt.Print(">> ")
	fmt.Fscanf(os.Stdin, "%v", &n)
	if n == k.Number {
		fmt.Println("")
		fmt.Println(">> you win !!!")
		fmt.Println("")
		k.GameFinished = true
	} else {
		if n > k.Number {
			fmt.Println(">> you number is greater")
		} else {
			fmt.Println(">> you number is lower")
		}
	}
	k.nextStep = &EndTurnStep{}
	return true
}

func (s *QuestionStep) name() string {
	return "guess the number ... "
}
