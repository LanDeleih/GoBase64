package command

import (
	"github.com/lanDeleih/gobase64/app/cmd"
	"github.com/urfave/cli"
	"go.uber.org/zap"
)

func NewGoBase64Command(version string, logger *zap.SugaredLogger) cli.App {
	o := cmd.CheckOpts{Logger: logger}
	return cli.App{
		Name:        "goBase64",
		Version:     version,
		Description: "decode and encode plain text or file",
		Commands: []cli.Command{
			cmd.NewDecodeCommand(o),
			cmd.NewEncodeCommand(o),
		},
	}
}
