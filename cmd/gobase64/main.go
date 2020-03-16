package main

import (
	"os"

	cmd "github.com/lanDeleih/gobase64/internal/command"
)

func main() {

	cmds := cmd.NewBase64Command()

	if err := cmds.Execute(); err != nil {
		os.Exit(1)
	}
}
