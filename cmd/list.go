package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var list = map[string][]string{
	"encodings": {"none", "base64", "url", "doubleurl"},
	"shells":    {"sh", "/bin/sh", "bash", "/bin/bash", "cmd", "powershell", "pwsh", "ash", "bsh", "csh", "ksh", "zsh", "pdksh", "tcsh", "mksh", "dash"},
	"ports":     {DefaultPort},
}

var listCmd = &cobra.Command{
	Use:   "list [resource]",
	Short: "List available shell types, methods, etc.",
	Long: `List available resources for the reverse shell generator.
    
Available resources:
  types     - List all available shell types
  methods   - List methods for a specific shell type
  ips       - List available IP addresses
  shells    - List available shells
  encodings - List available encodings
    `,
	GroupID: "utility",
	Args:    cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
		switch args[0] {
		case "types":
			types := getType()
			fmt.Println("Available shell types:")
			for _, t := range types {
				fmt.Println("-", t)
			}
		case "methods":
			if len(args) < 2 {
				fmt.Println("Please specify a shell type:")
				fmt.Println("  revshell list methods <type>")
				return
			}
			methods := getMethod(args[1])
			fmt.Printf("Available methods for %s:\n", args[1])
			for _, m := range methods {
				fmt.Println("-", m)
			}
		case "ips":
			ips := getIP()
			fmt.Println("Available IP addresses:")
			for _, ip := range ips {
				fmt.Println("-", ip)
			}
		case "shells":
			fmt.Println("Available shells:")
			for _, s := range list["shells"] {
				fmt.Println("-", s)
			}
		case "encodings":
			fmt.Println("Available encodings:")
			for _, e := range list["encodings"] {
				fmt.Println("-", e)
			}
		default:
			fmt.Printf("Unknown resource: %s\n", args[0])
			fmt.Println("Try 'revshell list' for available resources")
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
