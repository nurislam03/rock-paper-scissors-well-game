package game

type Well struct {
}

func (_ Well) Name() string {
	return "paper"
}

func (w Well) Beats(c Character) string {
	if c.Name() == "rock" || c.Name() == "scissors" {
		return "win"
	}

	if c.Name() == w.Name() {
		return "tie"
	}

	return "lose"
}
