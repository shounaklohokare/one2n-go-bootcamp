package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

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

		if strings.Contains(args[0], "-") && strings.Contains(args[1], "-") {
			rangeStartP1, rangeEndP1, err := getRange(args[0])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			rangeStartP2, rangeEndP2, err := getRange(args[1])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			executeStory3(rangeStartP1, rangeEndP1, rangeStartP2, rangeEndP2)
			return

		}

		p1, err := validateTarget(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if strings.Contains(args[1], "-") {
			rangeStart, rangeEnd, err := getRange(args[1])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			executeStory2(p1, rangeStart, rangeEnd)
			return
		}

		p2, err := validateTarget(args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		winsp1, winsp2 := executeStory1(p1, p2)

		displayResult(p1, p2, winsp1, winsp2)

	},
}

func displayResult(p1Target, p2Target, winsp1, winsp2 int) {

	winPercent := 100 * float32(winsp1) / 10.0

	fmt.Printf("Holding at %d vs Holding at %d : wins : %d/10 (%0.1f%%), losses: %d/10 (%0.1f%%)\n", p1Target, p2Target, winsp1, winPercent, winsp2, 100-winPercent)

}

func getRange(rangeInput string) (int, int, error) {

	parts := strings.Split(rangeInput, "-")

	x, err := validateTarget(parts[0])
	if err != nil {
		fmt.Println(err)
		return -1, -1, err
	}

	y, err := validateTarget(parts[1])
	if err != nil {
		fmt.Println(err)
		return -1, -1, err
	}

	return x, y, nil

}

func executeStory1(p1Target, p2Target int) (int, int) {

	p1 := game.Player{PlayerNumber: "p1", HoldTarget: p1Target}
	p2 := game.Player{PlayerNumber: "p2", HoldTarget: p2Target}

	dice := game.Dice{}

	wins := make(map[string]int)
	for i := 0; i < 10; i++ {

		p1.CurrentScore = 0
		p2.CurrentScore = 0

		p := &p1
		for {

			if p.ExecuteTurn(dice, 100) {
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

	return wins["p1"], wins["p2"]

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

func executeStory2(p1Target, rangeStart, rangeEnd int) {

	for p2Target := rangeStart; p2Target <= rangeEnd; p2Target++ {

		if p2Target == p1Target {
			continue
		}

		winsp1, winsp2 := executeStory1(p1Target, p2Target)

		displayResult(p1Target, p2Target, winsp1, winsp2)

	}

}

func executeStory3(rangeStartP1, rangeEndP1, rangeStartP2, rangeEndP2 int) {

	for p1Target := rangeStartP1; p1Target <= rangeEndP1; p1Target++ {

		totalWinsP1 := 0
		totalWinsP2 := 0

		for p2Target := rangeStartP2; p2Target <= rangeEndP2; p2Target++ {

			if p1Target == p2Target {
				continue
			}

			winsp1, winsp2 := executeStory1(p1Target, p2Target)

			totalWinsP1 += winsp1
			totalWinsP2 += winsp2

		}

		displayResultStory3(p1Target, totalWinsP1, totalWinsP2)

	}

}

func displayResultStory3(p1Target, totalWinsP1, totalWinsP2 int) {
	winPercent := 100 * float64(totalWinsP1) / 990.0

	fmt.Printf("Result: Wins, losses staying at k = %d: wins : %d/990 (%0.1f%%), losses: %d/990 (%0.1f%%)\n", p1Target, totalWinsP1, winPercent, totalWinsP2, 100-winPercent)
}
