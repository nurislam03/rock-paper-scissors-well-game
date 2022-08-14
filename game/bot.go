package game

import (
	"math/rand"
	"time"
)

type RandGenerator interface {
	Rand(end int) int
}

type Bot struct {
	randGen RandGenerator
}

func NewBot(randGen RandGenerator) Player {
	return Bot{randGen}
}

func (b Bot) Choice() Character {
	c := b.randGen.Rand(4)
	return CharacterFromName([]string{"rock", "paper", "scissors", "well"}[c])
}

type SimpleRandGenerator struct {
}

func NewSimpleRandGenerator() RandGenerator {
	rand.Seed(time.Now().UnixNano())
	return SimpleRandGenerator{}
}

func (r SimpleRandGenerator) Rand(end int) int {
	return rand.Intn(end)
}
