package cmd

import (
	"Txops/cvm"
	"github.com/spf13/cobra"
)

var deletecvmCmd = &cobra.Command{
	Use:   "delete",
	Short: " Delete your txcloud  cvm",
	Long:  `txops  cvm  delete`,
	Run: func(cmd *cobra.Command, args []string) {
		cvm.Deletecvm()

	},
}

func init()  {
	cvmCmd.AddCommand(deletecvmCmd)
	deletecvmCmd.Flags().StringVar(&cvm.Ipaddress, "ip", "", "ipadress  for delete cvm")
	deletecvmCmd.Flags().StringVar(&cvm.Zone,"zone","guangzhou","zone for show cvm ")
}
