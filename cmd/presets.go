package cmd

import (
	"fmt"
	"happcmd/internal/profile"

	"github.com/spf13/cobra"
)

var presetCommand = &cobra.Command{
	Use:   "preset",
	Short: "Manage presets",
}

var presetsInfoCommand = &cobra.Command{
	Use:   "list",
	Short: "List available presets",
	Run: func(cmd *cobra.Command, args []string) {
		for k, preset := range profile.Presets {
			fmt.Printf("%s: %s\n", k, preset.Name)
		}
	},
}

var presetsApplyCommand = &cobra.Command{
	Use:   "apply",
	Short: "Apply a preset",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		preset, ok := profile.Presets[name]
		if !ok {
			fmt.Println("Preset not found")
			return
		}
		p := preset.Factory(name)
		mode, _ := cmd.Flags().GetString("mode")
		url, err := p.ToURL(profile.Mode(mode))
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		fmt.Println(url)
	},
}

func init() {
	presetCommand.AddCommand(presetsInfoCommand)
	presetCommand.AddCommand(presetsApplyCommand)
	presetsApplyCommand.Flags().StringP("mode", "m", "add", "Import mode: add or onadd")
}
