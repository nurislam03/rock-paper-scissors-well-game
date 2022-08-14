package game

type Scissors struct {
}

func (_ Scissors) Name() string {
	return "scissors"
}

func (s Scissors) Beats(c Character) string {
	if c.Name() == "paper" {
		return "win"
	}

	if c.Name() == s.Name() {
		return "tie"
	}

	return "lose"
}
