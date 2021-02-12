package cmd

import (
	"encoding/base64"
	"fmt"
	"github.com/urfave/cli"
	"go.uber.org/zap"
)

type CheckOpts struct {
	Logger *zap.SugaredLogger
}

func NewDecodeCommand(o CheckOpts) cli.Command {
	return cli.Command{
		Name:    "decode",
		Usage:   "Decode base64 file or string",
		Aliases: []string{"d"},
		Action:  o.decode,
		Flags:   decodeFlags(),
	}
}

func decodeFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:     "file, f",
			Usage:    "Specify file to encoding",
			Required: false,
		},
	}
}

func (o *CheckOpts) decode(ctx *cli.Context) {
	switch {
	case ctx.String("file") != "":
		var pathToFile = ctx.String("file")
		data, err := getDataFromFile(pathToFile)
		if err != nil {
			o.Logger.Fatalf("failed to get data from file: \"%s\", %s", pathToFile, err)
		}
		o.decodeFile(data)
	default:
		for _, argument := range ctx.Args() {
			o.decodeString(argument)
		}
	}
}

func (o CheckOpts) decodeFile(data []byte) {
	decodedFile, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		o.Logger.Fatalf("failed to decode data from file: %s", err)
	}
	fmt.Printf("%s\n", decodedFile)
}

func (o CheckOpts) decodeString(data string) {
	decodedString, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		o.Logger.Errorf("failed to decode string: %s", err)
		return
	}
	fmt.Printf("%s\n", decodedString)
}
