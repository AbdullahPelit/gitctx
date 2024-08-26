#!/bin/bash

CONFIG_FILE=~/.gitctx_config

# Konfigürasyon dosyasının var olup olmadığını kontrol edin, yoksa oluşturun
if [ ! -f "$CONFIG_FILE" ]; then
    touch "$CONFIG_FILE"
fi

function list_accounts() {
    echo "Available accounts:"
    cat "$CONFIG_FILE"
}

function generate_ssh_key() {
    read -p "Enter email for SSH key: " email
    read -p "Enter a name for your SSH key (e.g., id_rsa_github): " key_name

    ssh-keygen -t ed25519 -C "$email" -f ~/.ssh/"$key_name"

    echo "SSH key generated at ~/.ssh/$key_name"

    echo "Add the following public key to your Git provider (GitHub, GitLab, etc.):"
    cat ~/.ssh/"$key_name".pub
    echo ""
    echo "Alternatively, you can copy it to your clipboard using the following command:"
    echo "xclip -sel clip < ~/.ssh/$key_name.pub"
    echo ""
    echo "GitHub: https://github.com/settings/keys"
    echo "GitLab: https://gitlab.com/profile/keys"
}

function add_account() {
    read -p "Enter account name (e.g., work, personal): " account_name
    read -p "Do you want to generate a new SSH key for this account? (y/n): " generate_key

    if [ "$generate_key" == "y" ]; then
        generate_ssh_key
        read -p "Enter the SSH key file name (e.g., id_rsa_github): " key_name
        ssh_key="~/.ssh/$key_name"
    else
        read -p "Enter the existing SSH key path (e.g., ~/.ssh/id_rsa): " ssh_key
    fi

    echo "$account_name:$ssh_key" >> "$CONFIG_FILE"
    echo "Account $account_name added."
}

function remove_account() {
    read -p "Enter account name to remove: " account_name
    sed -i "/^$account_name:/d" "$CONFIG_FILE"
    echo "Account $account_name removed."
}

function switch_account() {
    account_name=$1
    ssh_key=$(grep "^$account_name:" "$CONFIG_FILE" | cut -d':' -f2)

    if [ -n "$ssh_key" ]; then
        git config --global user.name "$account_name"
        git config --global user.email "$account_name@example.com"
        export GIT_SSH_COMMAND="ssh -i $ssh_key"
        echo "Switched to $account_name account."
    else
        echo "Account $account_name does not exist."
        list_accounts
    fi
}

case "$1" in
    add)
        add_account
        ;;
    remove)
        remove_account
        ;;
    switch)
        switch_account "$2"
        ;;
    list)
        list_accounts
        ;;
    *)
        echo "Usage: gitctx {add|remove|switch|list}"
        ;;
esac
