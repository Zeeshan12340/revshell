package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:     "info",
	Short:   "Shows configuration information",
	GroupID: "utility",
	Run: func(cmd *cobra.Command, args []string) {
		ipSources := getIPWithMetadata()
		if len(ipSources) == 0 {
			fmt.Println("No IP addresses found")
			return
		}
		config := readConfigFromFile()
		if config.IPAddress != "" || config.Port != "" || config.Shell != "" || config.Encoding != "" {
			fmt.Println("Configuration:")
			if config.IPAddress != "" {
				fmt.Printf("Configured IP: %s\n", config.IPAddress)
			}
			if config.Port != "" {
				fmt.Printf("Configured Port: %s\n", config.Port)
			}
			if config.Shell != "" {
				fmt.Printf("Configured Shell: %s\n", config.Shell)
			}
			if config.Encoding != "" {
				fmt.Printf("Configured Encoding: %s\n", config.Encoding)
			}
			homeDir, _ := os.UserHomeDir()
			configFile := "Not found"
			for _, path := range configPaths {
				fullPath := filepath.Join(homeDir, path)
				if _, err := os.Stat(fullPath); err == nil {
					configFile = fullPath
					break
				}
			}
			fmt.Printf("Config file: %s\n", configFile)
			fmt.Println()
		}

		fmt.Println("IP addresses in priority order:")
		for i, src := range ipSources {
			if src.Primary {
				fmt.Printf("%d. %s (primary, %s)\n", i+1, src.Address, src.Source)
			} else {
				fmt.Printf("%d. %s (%s)\n", i+1, src.Address, src.Source)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
