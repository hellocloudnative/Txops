package cmd

import (
	"Txops/image"
	"github.com/spf13/cobra"
)

var createimageCmd = &cobra.Command{
	Use:   "create"  ,
	Short: " Create your txcloud  image",
	Long:  `txops  image  create`,
	Run: func(cmd *cobra.Command, args []string) {
		image.MakeImage()

	},
}

func init()  {
	imageCmd.AddCommand(createimageCmd)
	createimageCmd.Flags().StringVar(&image.ImageName, "imagename", "", "imagename for make image")
	createimageCmd.Flags().StringVar(&image.Ipaddress, "ipaddress", "", "ipaddress for make image")
	createimageCmd.Flags().StringVar(&image.Zone,"zone","guangzhou","zone for show image ")
	createimageCmd.Flags().StringVar(&image.DryRun, "dryrun", "true", "try for create image")
}
