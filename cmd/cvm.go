package cmd

import (
	"github.com/spf13/cobra"
)

var cvmCmd = &cobra.Command{
	Use:   "cvm",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}




func init()  {
	RootCmd.AddCommand(cvmCmd)

}




