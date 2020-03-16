package command

import (
	"github.com/lanDeleih/gobase64/internal/command/decode"
	"github.com/lanDeleih/gobase64/internal/command/encode"
	"github.com/spf13/cobra"
)


//NewBase64Command execute decode or encode function
func NewBase64Command() *cobra.Command {

	cmds := &cobra.Command{
		Use:   "base64",
		Short: "Use it to decode/encode base64 string or files",
	}
	cmds.AddCommand(encode.NewEncode())
	cmds.AddCommand(decode.NewDecode())

	return cmds
}
