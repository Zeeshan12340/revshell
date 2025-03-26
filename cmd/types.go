package cmd

type Command struct {
	Name    string
	Method  string
	Command string
	Meta    []string
}

type CommandParams struct {
	Name      string
	Method    string
	IPAddress string
	Port      string
	Shell     string
	Encoding  string
}
