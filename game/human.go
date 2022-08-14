package game

type Human struct {
	move string
}

// NewHuman receives a move/choice for the character and returns a human player
func NewHuman(move string) Player {
	return &Human{move: move}
}

// Choice returns a character (e.g. rock, paper) based on their choice/move
func (h Human) Choice() Character {
	return CharacterFromName(h.move)
}
