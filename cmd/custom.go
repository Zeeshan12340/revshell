package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var customCmd = &cobra.Command{
	Use:   "custom",
	Short: "Generate a custom shell (scriptable)",
	Example: `  revshell custom -t bash -m i -i 10.10.14.20 -p 4444
  revshell custom -t powershell -m base64
  revshell custom -t python -m 1 -s /bin/bash -e base64`,
	GroupID: "utility",
	Run: func(cmd *cobra.Command, args []string) {
		params := CommandParams{}
		shellType, _ := cmd.Flags().GetString("type")
		method, _ := cmd.Flags().GetString("method")
		ip, _ := cmd.Flags().GetString("ip")
		port, _ := cmd.Flags().GetString("port")
		shell, _ := cmd.Flags().GetString("shell")
		encoding, _ := cmd.Flags().GetString("encoding")
		if shellType == "" {
			fmt.Println("Error: --type is required")
			return
		}
		if method == "" {
			methods := getMethod(shellType)
			if len(methods) > 0 {
				method = methods[0]
			}
		}
		if ip == "" {
			ips := getIP()
			if len(ips) > 0 {
				ip = ips[0]
			}
		}
		config := readConfigFromFile()
		if port == "" {
			if config.Port != "" {
				port = config.Port
			} else {
				port = DefaultPort
			}
		}
		if shell == "" {
			if config.Shell != "" {
				shell = config.Shell
			} else {
				shell = DefaultShell
			}
		}
		params.Name = shellType
		params.Method = method
		params.IPAddress = ip
		params.Port = port
		params.Shell = shell
		params.Encoding = encoding
		command := getCommand(params)
		encoded := setEncoding(params.Encoding, command)
		fmt.Println(encoded)
	},
}

func init() {
	rootCmd.AddCommand(customCmd)
	customCmd.Flags().StringP("type", "t", "", "Shell type (bash/python/powershell/etc)")
	customCmd.Flags().StringP("method", "m", "", "Method to use")
	customCmd.Flags().StringP("ip", "i", "", "IP address")
	customCmd.Flags().StringP("port", "p", "", "Port")
	customCmd.Flags().StringP("shell", "s", "", "Shell executable")
	customCmd.Flags().StringP("encoding", "e", "none", "Encoding type")
	customCmd.ValidArgsFunction = func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	customCmd.MarkFlagRequired("type")
	customCmd.MarkFlagRequired("method")
	customCmd.RegisterFlagCompletionFunc("type", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getType(), cobra.ShellCompDirectiveNoFileComp
	})
	customCmd.RegisterFlagCompletionFunc("method", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		t, _ := cmd.Flags().GetString("type")
		if t == "" {
			return []string{}, cobra.ShellCompDirectiveNoFileComp
		}
		return getMethod(t), cobra.ShellCompDirectiveNoFileComp
	})
	customCmd.RegisterFlagCompletionFunc("ip", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getIP(), cobra.ShellCompDirectiveNoFileComp
	})
	customCmd.RegisterFlagCompletionFunc("shell", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return list["shells"], cobra.ShellCompDirectiveNoFileComp
	})
	customCmd.RegisterFlagCompletionFunc("encoding", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return list["encodings"], cobra.ShellCompDirectiveNoFileComp
	})
}
