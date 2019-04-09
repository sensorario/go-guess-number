package main

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
