package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert an image to a different format",
	Run: func(cmd *cobra.Command, args []string) {
		if inputPath == "" || outputPath == "" || format == "" {
			fmt.Println("Usage: imgtool convert --input input.jpg --output output.png --format png")
			return
		}

		img, _, err := LoadImage(inputPath)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		err = SaveImage(img, outputPath, format)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Image successfully converted to", outputPath)
	},
}

var inputPath, outputPath, format string

func init() {
	rootCmd.AddCommand(convertCmd)
	convertCmd.Flags().StringVarP(&inputPath, "input", "i", "", "Input image file")
	convertCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Output image file")
	convertCmd.Flags().StringVarP(&format, "format", "f", "", "Output format (jpeg, png, webp)")
}
