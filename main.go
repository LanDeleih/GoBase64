package main

import (
	"os"

	"github.com/GoBase64/cmd"
)

func main() {
	cmds := cmd.NewBase64Command(os.Args[1:])

	if err := cmds.Execute(); err != nil {
		os.Exit(1)
	}
}
