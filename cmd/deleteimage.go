package cmd

import (
	"Txops/image"
	"github.com/spf13/cobra"
)

var deleteimageCmd = &cobra.Command{
	Use:   "delete",
	Short: " Delete your txcloud  image",
	Long:  `txops  image  delete`,
	Run: func(cmd *cobra.Command, args []string) {
		image.Deleteimage()

	},
}

func init()  {
	imageCmd.AddCommand(deleteimageCmd)
	deleteimageCmd.Flags().StringVar(&image.ImageName, "imagename", "", "ipadress  for delete image")
	deleteimageCmd.Flags().StringVar(&image.Zone,"zone","guangzhou","zone for show image ")
}
