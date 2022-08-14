package game

type Rock struct {
}

func (_ Rock) Name() string {
	return "rock"
}

func (r Rock) Beats(c Character) string {
	if c.Name() == "scissors" {
		return "win"
	}

	if c.Name() == r.Name() {
		return "tie"
	}

	return "lose"
}
