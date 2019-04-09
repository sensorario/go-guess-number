package main

import "fmt"

import "github.com/fatih/color"

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
