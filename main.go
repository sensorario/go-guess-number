package main

import "math/rand"
import "fmt"
import "os"
import "time"
import "github.com/fatih/color"

type StartGameStep struct{}

func (s *StartGameStep) play(k *Game) bool {
	k.nextStep = &StartTurnStep{}
	fmt.Println("")
	return true
}

func (s *StartGameStep) name() string {
	return "game started"
}

type StartTurnStep struct{}

func (s *StartTurnStep) play(k *Game) bool {
	if k.GameFinished == true {
		return false
	}

	k.nextStep = &QuestionStep{}
	return true
}

func (s *StartTurnStep) name() string {
	return "turn started"
}

type EndTurnStep struct{}

func (s *EndTurnStep) play(k *Game) bool {
	if k.GameFinished == true {
		return false
	}

	k.nextStep = &StartTurnStep{}
	return true
}

func (s *EndTurnStep) name() string {
	return "turn ended"
}

type Game struct {
	nextStep     GameStep
	Number       int
	GameFinished bool
}

func (a *Game) logStep() {
	fmt.Println(color.GreenString(a.nextStep.name()))
	if a.nextStep.name() == "turn ended" {
		fmt.Println("")
	}
}

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
