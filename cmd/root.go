package cmd

import (
	"github.com/spf13/cobra"
)

func RootCmd() {

	var rootCmd = &cobra.Command{
		Use:   "gobase64",
		Short: "GoBase64 is a programm for decode/encode file to base64",
		Long: `Use it when you need decode config files/secrets etc. 
									  to base64 and using it in k8s`,
		Run: func(cmd *cobra.Command, args []string) {
		  // Do Stuff Here
		},
	  }
	  
	  func Execute() {
		if err := rootCmd.Execute(); err != nil {
		  fmt.Println(err)
		  os.Exit(1)
		}
}
