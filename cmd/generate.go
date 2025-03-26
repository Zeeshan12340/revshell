package cmd

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:     "generate",
	Short:   "Generate a reverse shell with interactive prompts",
	GroupID: "utility",
	Run: func(cmd *cobra.Command, args []string) {
		params := promptForParams()
		command := getCommand(params)
		encoded := setEncoding(params.Encoding, command)
		fmt.Println(encoded)
	},
}

func promptSelection(reader *bufio.Reader, options []string, prompt string) (string, int) {
	for i, opt := range options {
		fmt.Printf("%d. %s\n", i+1, opt)
	}
	fmt.Print(prompt)
	choice := readInput(reader)
	var num int
	if _, err := fmt.Sscanf(choice, "%d", &num); err == nil {
		index := num - 1
		if index >= 0 && index < len(options) {
			return options[index], index
		}
	}
	return choice, -1
}

func promptForParams() CommandParams {
	reader := bufio.NewReader(os.Stdin)
	params := CommandParams{}
	types := getType()
	fmt.Println("Available shell types:")
	params.Name, _ = promptSelection(reader, types, "\nSelect shell type (number): ")
	methods := getMethod(params.Name)
	fmt.Println("\nAvailable methods for", params.Name+":")
	params.Method, _ = promptSelection(reader, methods, "\nSelect method (number): ")
	ips := getIP()
	fmt.Println("\nChoose IP address:")
	params.IPAddress, _ = promptSelection(reader, ips, "\nSelect IP (number): ")
	if net.ParseIP(params.IPAddress) == nil {
		fmt.Println("Warning: IP address format may be invalid")
	}
	config := readConfigFromFile()
	defaultPort := "9001"
	if config.Port != "" {
		defaultPort = config.Port
	}
	fmt.Printf("\nPort [%s]: ", defaultPort)
	portChoice := readInput(reader)
	if _, err := fmt.Sscanf(portChoice, "%d", new(int)); err != nil && portChoice != "" {
		fmt.Println("Warning: Port should be a number, using default")
		portChoice = "9001"
	}
	if portChoice == "" {
		params.Port = "9001"
	} else {
		params.Port = portChoice
	}
	fmt.Print("\nShell [sh]: ")
	shellChoice := readInput(reader)
	if shellChoice == "" {
		params.Shell = "sh"
	} else {
		params.Shell = shellChoice
	}
	fmt.Println("\nEncoding:")
	for i, e := range list["encodings"] {
		fmt.Printf("%d. %s\n", i+1, e)
	}
	fmt.Print("\nSelect encoding (number) [1]: ")
	encChoice := readInput(reader)
	if encChoice == "" {
		params.Encoding = "none"
	} else if num, err := fmt.Sscanf(encChoice, "%d", new(int)); err == nil {
		index := num - 1
		if index >= 0 && index < len(list["encodings"]) {
			params.Encoding = list["encodings"][index]
		}
	} else {
		params.Encoding = encChoice
	}
	return params
}

func readInput(reader *bufio.Reader) string {
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}
	return strings.TrimSpace(input)
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
