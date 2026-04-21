package cmd

import (
	"encoding/json"
	"fmt"
	"happcmd/internal/profile"
	"happcmd/internal/storage"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var profileCommand = &cobra.Command{
	Use:   "profile",
	Short: "Manage profiles",
}

var profileListCommand = &cobra.Command{
	Use:   "list",
	Short: "List available profiles",
	Run: func(cmd *cobra.Command, args []string) {
		profiles, err := storage.List()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if len(profiles) == 0 {
			fmt.Println("No profiles available")
			return
		}
		fmt.Println("==========")
		for index, profile := range profiles {
			fmt.Printf("%d. %s\n", index+1, profile)
		}
		fmt.Println("==========")
	},
}

var profileDeleteCommand = &cobra.Command{
	Use:   "delete",
	Short: "Delete profile",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := filepath.Base(args[0])
		err := storage.Delete(name)
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
		profile, err := storage.Load(name)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		data, err := json.MarshalIndent(profile, "", "  ")
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

		exists, err := storage.Exists(name)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if exists {
			fmt.Fprintln(os.Stderr, "profile already exists:", name)
			os.Exit(1)
		}

		var p *profile.Profile

		profileType, err := cmd.Flags().GetString("type")

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

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

		err = storage.Save(name, p)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Println("Profile saved:", name)
	},
}

func init() {
	profileCommand.AddCommand(profileListCommand)
	profileCommand.AddCommand(profileDeleteCommand)
	profileCommand.AddCommand(loadProfileCommand)
	profileCommand.AddCommand(saveProfileCommand)
	saveProfileCommand.Flags().StringP("type", "t", "default", "Profile type: default|strict|bypass")
}
