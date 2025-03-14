package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// resizeCmd represents the resize command
var resizeCmd = &cobra.Command{
	Use:   "resize",
	Short: "Resize an image while maintaining aspect ratio",
	Run: func(cmd *cobra.Command, args []string) {
		if inputPath == "" || outputPath == "" || width == 0 || height == 0 {
			fmt.Println("Usage: imgtool resize --input input.jpg --output resized.jpg --width 800 --height 600")
			return
		}

		img, _, err := LoadImage(inputPath)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		resizedImg := ResizeImage(img, width, height)
		err = SaveImage(resizedImg, outputPath, "jpeg") // Change as needed
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Image successfully resized to", outputPath)
	},
}

var width, height int

func init() {
	rootCmd.AddCommand(resizeCmd)
	resizeCmd.Flags().StringVarP(&inputPath, "input", "i", "", "Input image file")
	resizeCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Output image file")
	resizeCmd.Flags().IntVarP(&width, "width", "w", 0, "Width of resized image")
	resizeCmd.Flags().IntVarP(&height, "height", "h", 0, "Height of resized image")
}
