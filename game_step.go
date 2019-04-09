package main

type GameStep interface {
	play(g *Game) bool
	name() string
}
