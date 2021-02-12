package main

import (
	"github.com/lanDeleih/gobase64/app/command"
	"go.uber.org/zap"
	"log"
	"os"
)

var VERSION = "dev"

func main() {
	zapLog, err := zap.NewProduction()
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	sugar := zapLog.Sugar()

	app := command.NewGoBase64Command(VERSION, sugar)

	if err := app.Run(os.Args); err != nil {
		sugar.Fatal(err)
	}
}
