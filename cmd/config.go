package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

type ConfigSettings struct {
	IPAddress string
	Port      string
	Shell     string
	Encoding  string
}

var configPaths = []string{
	".config/revshell/config",
}

var configCmd = &cobra.Command{
	Use:     "config",
	Short:   "Create a sample config file",
	GroupID: "utility",
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting home directory:", err)
			return
		}
		configDir := filepath.Join(homeDir, ".config", "revshell")
		err = os.MkdirAll(configDir, 0755)
		if err != nil {
			fmt.Println("Error creating config directory:", err)
			return
		}
		configPath := filepath.Join(configDir, "config")
		ips := getIP()
		suggestedIP := "10.10.10.10"
		if len(ips) > 0 {
			suggestedIP = ips[0]
		}
		content := `# revshell configuration
# Uncomment and modify the following lines as needed

# Primary IP address to use for reverse shells
#ip=` + suggestedIP + `

# Default port to use
# port=9001

# Default shell
# shell=bash
`
		err = os.WriteFile(configPath, []byte(content), 0644)
		if err != nil {
			fmt.Println("Error writing config file:", err)
			return
		}
		fmt.Println("Config file created at:", configPath)
		fmt.Println("Edit this file to customize your reverse shell settings.")
	},
}

func readConfigFromFile() ConfigSettings {
	config := ConfigSettings{}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return config
	}
	for _, configPath := range configPaths {
		fullPath := filepath.Join(homeDir, configPath)
		data, err := os.ReadFile(fullPath)
		if err == nil {
			lines := strings.Split(string(data), "\n")
			for _, line := range lines {
				line = strings.TrimSpace(line)
				if strings.HasPrefix(line, "#") || line == "" {
					continue
				}
				parts := strings.SplitN(line, "=", 2)
				if len(parts) != 2 {
					continue
				}
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				switch key {
				case "ip":
					config.IPAddress = value
				case "port":
					config.Port = value
				case "shell":
					config.Shell = value
				case "encoding":
					config.Encoding = value
				}
			}
			break
		}
	}
	return config
}

func initConfig() {
	config := readConfigFromFile()
	if config.Port != "" && rootCmd.Flag("port") != nil && rootCmd.Flag("port").Value.String() == "" {
		rootCmd.Flag("port").Value.Set(config.Port)
	}
}

func init() {
	rootCmd.AddCommand(configCmd)
}
