package cmd

import (
	"github.com/GoBase64/cmd/decode"
	"github.com/GoBase64/cmd/encode"
	"github.com/spf13/cobra"
)

//NewBase64Command execute decode or encode function
func NewBase64Command(command *cobra.Command) {

	var encode = &cobra.Command{
		Use:   "encode",
		Short: "Encode file to base64",
		Run: func(cmd *cobra.Command, args []string) {
			encode.Base64()
		},
	}

	var decode = &cobra.Command{
		Use:   "decode",
		Short: "Decode string from base64",
		Run: func(cmd *cobra.Command, args []string) {
			decode.Base64()
		},
	}
	// decode.Base64()
	// encode.Base64()

}
