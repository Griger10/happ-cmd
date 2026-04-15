package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func readChoice() string {
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func readLine(prompt string, defaultValue string) string {
	fmt.Printf("%s (Enter для %q): ", prompt, defaultValue)
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())
	if input == "" {
		return defaultValue
	}
	return input
}

func readLines(prompt string) []string {
	fmt.Printf("%s (через запятую, Enter чтобы пропустить): ", prompt)
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())
	if input == "" {
		return []string{}
	}
	parts := strings.Split(input, ",")
	for i, p := range parts {
		parts[i] = strings.TrimSpace(p)
	}
	return parts
}
