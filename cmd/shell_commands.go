package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var bashCmd = &cobra.Command{
	Use:     "bash [ip] [port]",
	Short:   "Generate a bash reverse shell",
	GroupID: "shell",
	Args:    cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		config := readConfigFromFile()
		params := CommandParams{
			Name:     "bash",
			Method:   "i",
			Shell:    "bash",
			Encoding: "none",
		}
		if config.Port != "" {
			params.Port = config.Port
		} else {
			params.Port = DefaultPort
		}

		if config.Shell != "" {
			params.Shell = config.Shell
		}
		if len(args) > 0 {
			params.IPAddress = args[0]
		} else {
			if config.IPAddress != "" {
				params.IPAddress = config.IPAddress
			} else {
				ips := getIP()
				if len(ips) > 0 {
					params.IPAddress = ips[0]
				}
			}
		}
		if len(args) > 1 {
			params.Port = args[1]
		}
		shell := getCommand(params)
		encoded := setEncoding(params.Encoding, shell)
		fmt.Println(encoded)
	},
}

var powershellCmd = &cobra.Command{
	Use:     "powershell [ip] [port]",
	Short:   "Generate a PowerShell reverse shell",
	GroupID: "shell",
	Args:    cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		config := readConfigFromFile()
		params := CommandParams{
			Name:     "powershell",
			Method:   "base64",
			Encoding: "none",
		}
		if config.Port != "" {
			params.Port = config.Port
		} else {
			params.Port = DefaultPort
		}

		if config.Shell != "" {
			params.Shell = config.Shell
		}
		if len(args) > 0 {
			params.IPAddress = args[0]
		} else {
			if config.IPAddress != "" {
				params.IPAddress = config.IPAddress
			} else {
				ips := getIP()
				if len(ips) > 0 {
					params.IPAddress = ips[0]
				}
			}
		}
		if len(args) > 1 {
			params.Port = args[1]
		}
		shell := getCommand(params)
		fmt.Println(shell)
	},
}

var pythonCmd = &cobra.Command{
	Use:     "python [ip] [port]",
	Short:   "Generate a Python reverse shell",
	GroupID: "shell",
	Args:    cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		config := readConfigFromFile()
		params := CommandParams{
			Name:     "python",
			Method:   "1",
			Encoding: "none",
		}
		if config.Port != "" {
			params.Port = config.Port
		} else {
			params.Port = DefaultPort
		}

		if config.Shell != "" {
			params.Shell = config.Shell
		}
		if len(args) > 0 {
			params.IPAddress = args[0]
		} else {
			if config.IPAddress != "" {
				params.IPAddress = config.IPAddress
			} else {
				ips := getIP()
				if len(ips) > 0 {
					params.IPAddress = ips[0]
				}
			}
		}
		if len(args) > 1 {
			params.Port = args[1]
		}
		shell := getCommand(params)
		fmt.Println(shell)
	},
}

var phpCmd = &cobra.Command{
	Use:     "php [ip] [port]",
	Short:   "Generate a PHP reverse shell",
	GroupID: "shell",
	Args:    cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		config := readConfigFromFile()
		params := CommandParams{
			Name:     "php",
			Method:   "ivansincek",
			Encoding: "none",
		}
		if config.Port != "" {
			params.Port = config.Port
		} else {
			params.Port = DefaultPort
		}

		if config.Shell != "" {
			params.Shell = config.Shell
		}
		if len(args) > 0 {
			params.IPAddress = args[0]
		} else {
			if config.IPAddress != "" {
				params.IPAddress = config.IPAddress
			} else {
				ips := getIP()
				if len(ips) > 0 {
					params.IPAddress = ips[0]
				}
			}
		}
		if len(args) > 1 {
			params.Port = args[1]
		}
		shell := getCommand(params)
		fmt.Println(shell)
	},
}

func init() {
	rootCmd.AddCommand(bashCmd)
	rootCmd.AddCommand(powershellCmd)
	rootCmd.AddCommand(pythonCmd)
	rootCmd.AddCommand(phpCmd)
	for _, subCmd := range []*cobra.Command{bashCmd, powershellCmd, pythonCmd, phpCmd} {
		subCmd.Flags().StringP("encode", "e", "none", "Encoding (none/base64/url/doubleurl)")
	}
}
