package cmd

import (
	"encoding/base64"
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
)

func NewEncodeCommand(o CheckOpts) cli.Command {
	return cli.Command{
		Name:      "encode",
		Usage:     "Encode base64 file or string",
		ShortName: "e",
		Action:    o.encode,
		Flags:     encodeFlags(),
	}
}

func encodeFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:     "file, f",
			Usage:    "Specify file to encoding",
			Required: false,
		},
	}
}

func (o CheckOpts) encode(ctx *cli.Context) {
	switch {
	case ctx.String("file") != "":
		var pathToFile = ctx.String("file")
		data, err := getDataFromFile(pathToFile)
		if err != nil {
			o.Logger.Fatalf("failed to get data from file: \"%s\", %s", pathToFile, err)
		}
		o.encodeFile(data)
	default:
		for _, argument := range ctx.Args() {
			o.encodeString([]byte(argument))
		}
	}
}

func (o CheckOpts) encodeFile(data []byte) {
	encodedFile := base64.StdEncoding.EncodeToString(data)
	fmt.Printf("%s\n", encodedFile)
}

func (o CheckOpts) encodeString(data []byte) {
	encodedString := base64.StdEncoding.EncodeToString(data)
	fmt.Printf("%s\n", encodedString)
}

func getDataFromFile(file string) ([]byte,error) {
	return ioutil.ReadFile(file)
}