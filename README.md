# revshell

Generate revshells in your terminal

## Installation

```bash
go install -v github.com/gubarz/revshell@latest
```

## Usage

### Interactive Mode

The easiest way to generate a reverse shell is using the interactive mode:

```bash
revshell generate
```

This will guide you through selecting a shell type, method, IP address, port, and other options.

### Direct Shell Commands
For quick access to common shells:

```bash
# Basic bash reverse shell
revshell bash

# PowerShell reverse shell with default port
revshell powershell 10.10.14.20

# Python reverse shell with base64 encoding
revshell python 10.10.14.20 4444 -e base64

# PHP reverse shell
revshell php 10.10.14.20 4444
```

### Custom Shell Generation

For more control, use the custom command:

```bash
# Generate a specific shell type and method
revshell custom -t powershell -m base64 -i 10.10.14.20 -p 4444 -s powershell

# Use tab completion to see available options
revshell custom -t [TAB] -m [TAB]
```

### Utility Commands

```bash
# Show system information (IPs and config)
revshell info

# List available shell types
revshell list types

# List methods for a specific shell
revshell list methods bash

# Create a config file
revshell config
```

### Configuration

You can create a configuration file to set default values:

```bash
# Create a sample config file
revshell config

# Edit the config file
vim ~/.config/revshell/config
```

Configuration options:

- ip: Default IP address
- port: Default port (e.g., 9001)
- shell: Default shell (e.g., bash)
- encoding: Default encoding

## Contributing

Contributions are welcome! Please make a pull request.

## Credits

https://www.revshells.com / https://github.com/0dayCTF/reverse-shell-generator
