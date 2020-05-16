package cmd

import (
	"Txops/cvm"
	"github.com/spf13/cobra"
)

var showcvmCmd = &cobra.Command{
	Use:   "show",
	Short: " show your txcloud  cvm",
	Long:  `txops  cvm  show`,
	Run: func(cmd *cobra.Command, args []string) {
		cvm.Showcvm()

	},
}

func init()  {
	cvmCmd.AddCommand(showcvmCmd)
	showcvmCmd.Flags().StringVar(&cvm.All, "all", "", "all  for show cvm")
	showcvmCmd.Flags().StringVar(&cvm.Zone,"zone","guangzhou","zone for show cvm ")
}
