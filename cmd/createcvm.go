package cmd

import (
	"Txops/cvm"
	"github.com/spf13/cobra"
)

var createcvmCmd = &cobra.Command{
	Use:   "create"  ,
	Short: " Create your txcloud  cvm",
	Long:  `txops  cvm  create`,
	Run: func(cmd *cobra.Command, args []string) {
		 cvm.Createcvm()

	},
}

func init()  {
	cvmCmd.AddCommand(createcvmCmd)
	createcvmCmd.Flags().StringVar(&cvm.Ipaddress, "ip", "", "ipadress  for create cvm")
	createcvmCmd.Flags().StringVar(&cvm.HostName, "hostname", "", "hostname for create cvm")
	createcvmCmd.Flags().StringVar(&cvm.ImageName, "imagename", "base-image", "imagename for create cvm")
	createcvmCmd.Flags().StringVar(&cvm.Zone,"zone","guangzhou","zone for show cvm ")
	createcvmCmd.Flags().StringVar(&cvm.DryRun, "dryrun", "true", "try for create cvm")
}
