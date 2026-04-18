package cmd

import (
	"fmt"
	"happcmd/internal/profile"

	"github.com/spf13/cobra"
)

var DecodeCommand = &cobra.Command{
	Use:   "decode",
	Short: "Decode a happ:// link",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		decodedUrl, err := profile.DecodeUrl(url)

		if err != nil {
			fmt.Println("error:", err)
			return
		}
		fmt.Println(decodedUrl)
	},
}
