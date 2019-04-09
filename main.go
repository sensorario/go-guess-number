package main

import "math/rand"
import "fmt"
import "os"
import "time"

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
