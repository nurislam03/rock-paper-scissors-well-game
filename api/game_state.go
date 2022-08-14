package api

type gameState struct {
	gameState map[pair]string
}

type pair struct {
	first  string
	second string
}

func NewGameState() *gameState {

	return &gameState{
		map[pair]string{
			pair{"rock", "rock"}:         "tie",
			pair{"rock", "paper"}:        "lose",
			pair{"rock", "scissors"}:     "win",
			pair{"rock", "well"}:         "lose",
			pair{"paper", "rock"}:        "win",
			pair{"paper", "paper"}:       "tie",
			pair{"paper", "scissors"}:    "lose",
			pair{"paper", "well"}:        "win",
			pair{"scissors", "rock"}:     "lose",
			pair{"scissors", "paper"}:    "win",
			pair{"scissors", "scissors"}: "tie",
			pair{"scissors", "well"}:     "lose",
			pair{"well", "rock"}:         "win",
			pair{"well", "paper"}:        "lose",
			pair{"well", "scissors"}:     "win",
			pair{"well", "well"}:         "tie",
		},
	}
}

func (g *gameState) resultState(choiceA, choiceB string) string {
	return g.gameState[pair{
		choiceA,
		choiceB,
	}]
}
