package main

import (
	"fmt"
	"os"

	"github.com/shounaklohokare/one2n-go-bootcamp/word-count/cmd"
)

func main() {

	run(cmd.RootCmd.Execute())

}

func run(err error) {

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
