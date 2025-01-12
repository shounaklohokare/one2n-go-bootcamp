package cmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a game of pig",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Welcome to A Game of Pig")

	},
}

func rollDie() int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(6) + 1
}

func turn() {

}

func init() {
	RootCmd.AddCommand(startCmd)
}
