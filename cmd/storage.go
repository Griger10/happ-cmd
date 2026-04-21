package cmd

import (
	"encoding/json"
	"fmt"
	"happcmd/internal/profile"
	"happcmd/internal/storage"
	"os"
	"path/filepath"
	"sort"

	"github.com/spf13/cobra"
)

var store *storage.Storage

var profileCommand = &cobra.Command{
	Use:   "profile",
	Short: "Manage profiles",
}

var profileListCommand = &cobra.Command{
	Use:   "list",
	Short: "List available profiles",
	Run: func(cmd *cobra.Command, args []string) {
		profiles, err := store.List()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if len(profiles) == 0 {
			fmt.Println("No profiles available")
			return
		}

		sort.Strings(profiles)

		fmt.Println("Profiles:")
		for i, p := range profiles {
			fmt.Printf("  %d) %s\n", i+1, p)
		}
	},
}

var profileDeleteCommand = &cobra.Command{
	Use:   "delete",
	Short: "Delete profile",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := filepath.Base(args[0])

		err := store.Delete(name)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Println("Deleted:", name)
	},
}

var loadProfileCommand = &cobra.Command{
	Use:   "load",
	Short: "Load profile",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := filepath.Base(args[0])

		p, err := store.Load(name)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		data, err := json.MarshalIndent(p, "", "  ")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Println(string(data))
	},
}

var saveProfileCommand = &cobra.Command{
	Use:   "save",
	Short: "Save profile",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := filepath.Base(args[0])

		exists, err := store.Exists(name)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if exists {
			fmt.Fprintln(os.Stderr, "profile already exists:", name)
			os.Exit(1)
		}

		profileType, err := cmd.Flags().GetString("type")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		var p *profile.Profile

		switch profileType {
		case "default":
			p = profile.NewProfile(name, nil, nil, nil)
		case "strict":
			p = profile.NewStrictProfile(name)
		case "bypass":
			p = profile.NewBypassAllProfile(name)
		default:
			fmt.Fprintln(os.Stderr, "unknown profile type:", profileType)
			os.Exit(1)
		}

		if err := store.Save(name, p); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Println("Profile saved:", name)
	},
}

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	store = storage.New(home)

	profileCommand.AddCommand(profileListCommand)
	profileCommand.AddCommand(profileDeleteCommand)
	profileCommand.AddCommand(loadProfileCommand)
	profileCommand.AddCommand(saveProfileCommand)

	saveProfileCommand.Flags().StringP("type", "t", "default", "Profile type: default|strict|bypass")
}
