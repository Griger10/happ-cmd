package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "happcmd",
	Short: "CLI для генерации Happ routing профилей",
	Long: `happcmd — инструмент для генерации ссылок импорта
профилей маршрутизации для приложения Happ.

Позволяет создавать профили с кастомными правилами
маршрутизации и экспортировать их в формате happ://`,
	Run: func(cmd *cobra.Command, args []string) {
		for {
			fmt.Println("\nhappcmd v0.1.0")
			fmt.Println("1. Сгенерировать профиль")
			fmt.Println("2. Выход")
			fmt.Print("\nВыберите действие: ")
			choiceStr := readChoice()
			var choice int
			_, err := fmt.Sscan(choiceStr, &choice)
			if err != nil {
				fmt.Println("Введите число")
				continue
			}

			switch choice {
			case 1:
				name := readLine("Название профиля", "DefaultProfile")
				mode := readLine("Режим (add/onadd)", "add")
				directSites := readLines("Прямые сайты")
				blockSites := readLines("Заблокированные сайты")
				directIPs := readLines("Прямые IP")
				runGenerate(name, mode, directSites, blockSites, directIPs)
			case 2:
				fmt.Println("Выход...")
				return
			default:
				fmt.Println("Неизвестная команда")
			}
		}
	},
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Сгенерировать Routing профиль",
	Long: `Генерирует профиль маршрутизации и выводит ссылку для импорта в Happ.

Примеры:
  happcmd generate
  happcmd generate -n "Мой профиль"
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
	generateCmd.Flags().StringP("mode", "m", "add", "Режим импорта: add или onadd")
	generateCmd.Flags().StringP("name", "n", "DefaultProfile", "Название профиля в Happ")
	generateCmd.Flags().StringSlice("add-direct-site", []string{}, "Сайты для прямой передачи трафика")
	generateCmd.Flags().StringSlice("add-block-site", []string{}, "Сайты для блокировки трафика")
	generateCmd.Flags().StringSlice("add-direct-ip", []string{}, "Сайты для прямой передачи трафика по IP")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
