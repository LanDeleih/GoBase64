package decode

import (
	"encoding/base64"
	"fmt"
	"github.com/spf13/cobra"
	"k8s.io/klog"
)

//NewDecode decode base64 function
func NewDecode() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "decode",
		Short: "Decode file from base64",
		Run: func(cmd *cobra.Command, args []string) {
			if err := Decode(args); err != nil {
				klog.Fatalf("Failed to decode: %s", err)
			}
		},
	}

	return cmd
}

func Decode(args []string) error {
	for _, v := range args {
		decode, err := base64.StdEncoding.DecodeString(v)
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", decode)
	}
	return nil
}