package cmd

import (
	"Txops/image"
	"github.com/spf13/cobra"
)

var showimageCmd = &cobra.Command{
	Use:   "show",
	Short: " Show your txcloud  image",
	Long:  `txops  image  show`,
	Run: func(cmd *cobra.Command, args []string) {
		image.ShowImage()

	},
}

func init()  {
	imageCmd.AddCommand(showimageCmd)
	showimageCmd.Flags().StringVar(&image.All,"all","", "all  for show cvm")
	showimageCmd.Flags().StringVar(&image.Zone,"zone","guangzhou","zone for show image ")
	showimageCmd.Flags().StringVar(&image.ImageName, "imagename", "", "imagename for show image")
	showimageCmd.Flags().StringVar(&image.ImageId, "imageid", "", "imageid for show image")
}
