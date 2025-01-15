package main

import (
	"fmt"
	"os"

	"github.com/shounaklohokare/one2n/game_of_pig/cmd"
)

func main() {

	must(cmd.RootCmd.Execute())

}

func must(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
