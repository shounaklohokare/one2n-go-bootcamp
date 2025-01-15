package game

import (
	"math/rand"
	"time"
)

type Dice struct{}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (dice Dice) RollDie() int {

	return rand.Intn(6) + 1
}
