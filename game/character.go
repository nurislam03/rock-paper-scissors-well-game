package game

type Character interface {
	Beats(Character) string
	Name() string
}

func CharacterFromName(name string) Character {
	switch name {
	case "rock":
		return Rock{}
	case "paper":
		return Paper{}
	case "scissors":
		return Scissors{}
	case "well":
		return Well{}
	}

	return nil
}
