/*
Copyright © 2024 indrora
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// extractCmd represents the extract command
var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "Extract raw images from an LRI file",
	Long: `Extracts raw data from an LRI raw file. 
This includes images, JSON-formatted versions of the protobuf data,
and more. `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("extract called")
	},
}

func init() {
	rootCmd.AddCommand(extractCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// extractCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// extractCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
