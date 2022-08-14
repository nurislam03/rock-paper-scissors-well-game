package game

type Game struct {
	p1, p2 Player
}

func NewGame(p1, p2 Player) Game {
	return Game{p1, p2}
}

type GameResult struct {
	Choice1, Choice2 Character
	Result           string
}

func (g Game) Play() GameResult {
	choice1, choice2 := g.p1.Choice(), g.p2.Choice()
	return GameResult{
		Choice1: choice1,
		Choice2: choice2,
		Result:  choice1.Beats(choice2),
	}
}
