package game

type Paper struct {
}

func (_ Paper) Name() string {
	return "paper"
}

func (p Paper) Beats(c Character) string {
	if c.Name() == "rock" || c.Name() == "well" {
		return "win"
	}

	if c.Name() == p.Name() {
		return "tie"
	}

	return "lose"
}
