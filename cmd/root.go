package cmd

import (
	"fmt"
	"happcmd/internal/profile"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "happcmd",
	Short: "CLI для генерации Happ routing профилей",
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Сгенерировать Routing профиль",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		p := profile.NewProfile(name)
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
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringP("mode", "m", "add", "Режим импорта: add или onadd")
	generateCmd.Flags().StringP("name", "n", "DefaultProfile", "Название профиля в Happ")
	generateCmd.Flags().StringArray("")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
