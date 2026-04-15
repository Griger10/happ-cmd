package cmd

import (
	"fmt"
	"happcmd/internal/profile"
)

func runGenerate(name string, mode string, directSites, blockSites, directIPs []string) {
	p := profile.NewProfile(name, directSites, blockSites, directIPs)
	url, err := p.ToURL(profile.Mode(mode))
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(url)
}
