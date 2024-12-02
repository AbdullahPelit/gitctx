# gitctx - GitHub/GitLab Account Management Tool

[![Go Version](https://img.shields.io/github/go-mod/go-version/AbdullahPelit/gitctx)](https://golang.org/)
[![License](https://img.shields.io/github/license/AbdullahPelit/gitctx)](LICENSE)
[![Issues](https://img.shields.io/github/issues/AbdullahPelit/gitctx)](https://github.com/AbdullahPelit/gitctx/issues)
[![Pull Requests](https://img.shields.io/github/issues-pr/AbdullahPelit/gitctx)](https://github.com/AbdullahPelit/gitctx/pulls)
[![Go Report Card](https://goreportcard.com/badge/github.com/AbdullahPelit/gitctx)](https://goreportcard.com/report/github.com/AbdullahPelit/gitctx)

`gitctx` is a tool that allows you to easily manage multiple GitHub or GitLab accounts. You can dynamically generate SSH keys, switch between accounts, and use separate SSH keys for different accounts.

## Features

- Create and manage separate SSH keys for accounts
- Quick switching between accounts
- List registered accounts with status indicators
- Add and remove accounts
- Support for different SSH key types (ED25519, RSA)
- Use existing SSH keys
- Automatic SSH config management
- Backup and restore functionality

## Installation

### Option 1: Using Go Install

```bash
go install github.com/AbdullahPelit/gitctx@latest
```

### Option 2: From Source

1. Clone the Repository:

```bash
git clone https://github.com/AbdullahPelit/gitctx.git
cd gitctx
```

2. Install Dependencies:

```bash
go mod tidy
```

3. Build the Project:

```bash
go build -o gitctx
```

4. Add to PATH:

```bash
mv gitctx ~/bin/
```

Add to shell configuration (~/.bashrc, ~/.zshrc):

```bash
export PATH=$PATH:~/bin
source ~/.bashrc  # or ~/.zshrc
```

### Option 3: Download Binary

Download the latest binary for your platform from the [releases page](https://github.com/AbdullahPelit/gitctx/releases).

## Usage

### Adding an Account

Basic usage:

```bash
gitctx add personal user@example.com
```

With options:

```bash
# Use RSA key
gitctx add work work@company.com --key-type rsa --bits 4096

# Use existing key
gitctx add github github@email.com --use-existing

# Force overwrite
gitctx add gitlab gitlab@email.com --force
```

### Removing an Account

```bash
gitctx remove <account_name>
```

### Switching Between Accounts

```bash
gitctx switch <account_name>
```

### Listing Accounts

```bash
gitctx list
```

### Version Information

```bash
gitctx -v
```

## Configuration

The tool stores its configuration in the following locations:
- SSH Keys: `~/.ssh/`
- Config File: `~/.gitctx_config`
- Current Account: `~/.gitctx_current`

## Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

### Development Setup

1. Fork the repository
2. Clone your fork
3. Install dependencies
4. Create a feature branch
5. Make your changes
6. Submit a pull request

## Security

For security concerns, please see our [Security Policy](SECURITY.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

- 📫 Report bugs via [GitHub Issues](https://github.com/AbdullahPelit/gitctx/issues)
- 💡 Request features via [GitHub Issues](https://github.com/AbdullahPelit/gitctx/issues)
- 💬 Ask questions in [GitHub Discussions](https://github.com/AbdullahPelit/gitctx/discussions)

## Changelog

See [CHANGELOG.md](CHANGELOG.md) for release history.

## Acknowledgments

Thanks to all [contributors](https://github.com/AbdullahPelit/gitctx/graphs/contributors) who have helped improve this project.

---

<p align="center">
  Made with ❤️ by <a href="https://github.com/AbdullahPelit">Abdullah Pelit</a>
</p>



