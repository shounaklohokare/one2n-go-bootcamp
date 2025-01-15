package game

type Player struct {
	PlayerNumber string
	HoldTarget   int
	CurrentScore int
}

func (player *Player) ExecuteTurn(dice Dice, winningScore int) bool {

	for {

		roll := dice.RollDie()

		if roll == 1 {
			return false
		}

		player.CurrentScore += roll

		if player.CurrentScore >= player.HoldTarget {
			return true
		}

	}

}
