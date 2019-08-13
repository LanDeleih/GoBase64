package cmd

import (
	"io"

	"github.com/GoBase64/cmd/decode"
	"github.com/GoBase64/cmd/encode"
	"github.com/spf13/cobra"
)

// var encode = &cobra.Command{
// 	Use:   "encode",
// 	Short: "Encode file to base64",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		encode.ToBase64()
// 	},
// }

// var decode = &cobra.Command{
// 	Use:   "decode",
// 	Short: "Decode string from base64",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		decode.FromBase64()
// 	},
// }

//NewBase64Command execute decode or encode function
func NewBase64Command(in io.Reader) *cobra.Command {

	cmds := &cobra.Command{
		Use:   "base64",
		Short: "Use it to decode/encode base64 string or files",
	}
	cmds.AddCommand(encode.ToBase64(in))
	cmds.AddCommand(decode.FromBase64(in))

	return cmds
}
