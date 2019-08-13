package decode

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

//FromBase64 decode base64 function
func FromBase64(in io.Reader) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "decode",
		Short: "Decode file from base64",
		Run: func(cmd *cobra.Command, args []string) {
			read := os.Args[2:]

			for i := range read {
				decode, err := base64.StdEncoding.DecodeString(read[i])
				if err != nil {
					log.Fatalln("Cannot decode input", err)
				}

				fmt.Println(string(decode))
			}
		},
	}

	return cmd
}
