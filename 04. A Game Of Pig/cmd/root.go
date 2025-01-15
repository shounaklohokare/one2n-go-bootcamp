package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/shounaklohokare/one2n/game_of_pig/game"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "pig",
	Short: "Pig is a two-player game with a die",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 2 {
			fmt.Println("Please provide two target strategies")
			os.Exit(1)
		}

		p1, err := validateTarget(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		p2, err := validateTarget(args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		executeStory1(p1, p2)

	},
}

func executeStory1(p1Target, p2Target int) {

	p1 := game.Player{PlayerNumber: "p1", HoldTarget: p1Target}
	p2 := game.Player{PlayerNumber: "p2", HoldTarget: p2Target}

	dice := game.Dice{}

	wins := make(map[string]int)
	for i := 0; i < 10; i++ {

		p1.CurrentScore = 0
		p2.CurrentScore = 0

		p := &p1
		for {

			if p.ExecuteTurn(dice, 200) {
				break
			}

			if p.PlayerNumber == "p1" {
				p = &p2
			} else {
				p = &p1
			}

		}
		wins[p.PlayerNumber]++

	}

	percent := 100 * float32(wins["p1"]) / 10.0

	fmt.Printf("Holding at %d vs Holding at %d : wins : %d/10 (%0.1f%%), losses: %d/10 (%0.1f%%)", p1Target, p2Target, wins["p1"], percent, 10-wins["p1"], 100-percent)

}

func validateTarget(target string) (int, error) {

	out, err := strconv.Atoi(target)
	if err != nil {
		return -1, fmt.Errorf("the hold target should be a valid number :- %w", err)
	}

	if out < 1 || out > 100 {
		return -1, fmt.Errorf("the hold target should be between 1 to 100 :- %w", err)
	}

	return out, nil

}
