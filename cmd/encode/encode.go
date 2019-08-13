package encode

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
)

//ToBase64 this function will encode files to base64
func ToBase64(in io.Reader) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "encode",
		Short: "Encode file to base64",
		Run: func(cmd *cobra.Command, args []string) {
			if len(os.Args) < 2 {
				fmt.Println("You missed file name")
			}

			file, err := ioutil.ReadFile(os.Args[2])
			if err != nil {
				log.Fatalln("Can't read file:", os.Args[2])
			}

			encode := base64.StdEncoding.EncodeToString(file)

			fmt.Println(encode)
		},
	}

	return cmd

}
