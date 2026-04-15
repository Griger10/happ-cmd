package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "happcmd",
	Short: "CLI tool for generating Happ routing profile links",
	Long: `happcmd — a command-line tool for generating Happ routing
profile import links.

Allows creating profiles with custom routing rules
and exporting them in happ:// format.`,
	Run: func(cmd *cobra.Command, args []string) {
		for {
			fmt.Printf("\nhappcmd v%s\n", Version)
			fmt.Println("1. Generate profile")
			fmt.Println("2. Exit")
			fmt.Print("\nSelect action: ")
			choiceStr := readChoice()
			var choice int
			_, err := fmt.Sscan(choiceStr, &choice)
			if err != nil {
				fmt.Println("Please enter a number")
				continue
			}

			switch choice {
			case 1:
				name := readLine("Profile name", "DefaultProfile")
				mode := readLine("Mode (add/onadd)", "add")
				directSites := readLines("Direct sites")
				blockSites := readLines("Blocked sites")
				directIPs := readLines("Direct IPs")
				runGenerate(name, mode, directSites, blockSites, directIPs)
			case 2:
				return
			default:
				fmt.Println("Unknown command")
			}
		}
	},
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a routing profile",
	Long: `Generates a routing profile and outputs an import link for Happ.

Examples:
  happcmd generate
  happcmd generate -n "My Profile"
  happcmd generate -m onadd
  happcmd generate --add-direct-site "domain:github.com"
  happcmd generate --add-block-site "geosite:gambling" --add-direct-ip "1.2.3.4/32"`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		mode, _ := cmd.Flags().GetString("mode")
		directSites, _ := cmd.Flags().GetStringSlice("add-direct-site")
		blockSites, _ := cmd.Flags().GetStringSlice("add-block-site")
		directIPs, _ := cmd.Flags().GetStringSlice("add-direct-ip")
		runGenerate(name, mode, directSites, blockSites, directIPs)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringP("mode", "m", "add", "Import mode: add or onadd")
	generateCmd.Flags().StringP("name", "n", "DefaultProfile", "Profile name in Happ")
	generateCmd.Flags().StringSlice("add-direct-site", []string{}, "Sites for direct routing")
	generateCmd.Flags().StringSlice("add-block-site", []string{}, "Sites to block")
	generateCmd.Flags().StringSlice("add-direct-ip", []string{}, "IPs for direct routing")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
