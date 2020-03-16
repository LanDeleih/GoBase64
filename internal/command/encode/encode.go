package encode

import (
	"encoding/base64"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"k8s.io/klog"
)

//NewEncode this function will encode files to base64
func NewEncode() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "encode",
		Short: "Encode to base64",
		Run: func(cmd *cobra.Command, args []string) {
			if err := Encode(cmd, args); err != nil {
				klog.Fatalf("Failed to encode: %s", err)
			}
		},
	}
	cmd.Flags().StringP("file", "f", "", "Determine a file to encode")
	return cmd
}

func Encode(cmd *cobra.Command, args []string) error {
	a := len(args)
	switch {
	case a >= 1:
		fmt.Println(base64.StdEncoding.EncodeToString([]byte(args[0])))
		return nil
	case a == 0:
		f, err := cmd.Flags().GetString("file")
		if err != nil {
			return err
		}
		file, err := ioutil.ReadFile(f)
		if err != nil {
			return err
		}
		fmt.Println(base64.StdEncoding.EncodeToString(file))
		return nil
	}
	return nil
}
