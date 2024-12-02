# gitctx - GitHub/GitLab Account Management Tool

`gitctx` is a tool that allows you to easily manage multiple GitHub or GitLab accounts. You can dynamically generate SSH keys, switch between accounts, and use separate SSH keys for different accounts.

## Features

- Create and manage separate SSH keys for accounts
- Quick switching between accounts
- List registered accounts
- Add and remove accounts

## Installation

### 1. Clone the Repository

First, you need to download the `gitctx` tool. Clone the repository using the following command:

```bash
git clone https://github.com/AbdullahPelit/gitctx.git
cd gitctx
```

### 2. Install Dependencies

The project uses the cobra package. Run the following command in the terminal to install dependencies:

```bash
go mod tidy
```

### 3. Build the Project

You can build an executable version of the project using the following command:

```bash
go build -o gitctx
```

### 4. Add gitctx to PATH

To use the gitctx command from anywhere, add the executable to your system's PATH. To do this:

```bash
mv gitctx ~/bin/
```

Then, add this directory to PATH by adding the following line to ~/.bashrc, ~/.zshrc, or the appropriate shell configuration file:

```bash
export PATH=$PATH:~/bin
source ~/.bashrc  # or run this command if you're using ~/.zshrc
```

## Usage

### Adding an Account

To add a new account:

```bash
gitctx add
```

This command gives you the option to create a new SSH key or use an existing one. If you create a new SSH key, you'll need to add it to your Git provider (GitHub, GitLab, etc.). The SSH key will be printed to the screen.

### Removing an Account

To remove an existing account:

```bash
gitctx remove <account_name>
```

### Switching Between Accounts

To switch between added accounts:

```bash
gitctx switch <account_name>
```

### Listing Accounts

To list all registered accounts:

```bash
gitctx list
```

## Contributing
If you'd like to contribute, please fork the repository and submit a pull request. We're open to all kinds of contributions!



