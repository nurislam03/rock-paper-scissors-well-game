package game

type Character interface {
	Beats(Character) string
	Name() string
}
