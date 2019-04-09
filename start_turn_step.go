package main

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
