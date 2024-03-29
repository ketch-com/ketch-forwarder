#!/usr/bin/env sh

# Code generated by shipbuilder init 1.21.7. DO NOT EDIT.

check_installed() {
  installed=$(which "$1")

  if [ -z "$installed" ]; then
    echo "ERROR: $1 is not installed."
    if [ "$CI" = "true" ]; then
      exit 1
    fi

    echo "Do you want to install and continue? Y/N"
    read install
    install=$(echo "$install" | tr "[:lower:]" "[:upper:]")
    if [ "$install" = "Y" -o "$install" = "YES" ]; then
      sh -c "$2"
    fi

    export installed=$(which "$1")
    if [ -z "$installed" ]; then
      echo "ERROR: $1 is not installed."
      echo "Install $1 using '$2'."
      exit 1
    fi
  fi
}

brew_installed() {
  if [ "$CI" != "true" ]; then
    check_installed brew '/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"'
    brew=$installed
  fi
}

brew_install() {
  brew_installed
  check_installed $1 "$brew install ${2:-$1}"
}

node_installed() {
  if [ -f "/usr/local/opt/nvm/nvm.sh" ]; then
    . /usr/local/opt/nvm/nvm.sh
    nvm use --lts
  fi

  if [ -z "$installed" ]; then
    installed=$(which node)
  fi

  if [ -z "$installed" ]; then
    echo "ERROR: node is not installed."
    if [ "$CI" = "true" ]; then
      exit 1
    fi

    echo "Do you want to install and continue? Y/N"
    read install
    install=$(echo "$install" | tr "[:lower:]" "[:upper:]")
    if [ "$install" = "Y" -o "$install" = "YES" ]; then
      brew_installed

      $brew install nvm
      . /usr/local/opt/nvm/nvm.sh
      nvm install --lts --latest-npm
      nvm use --lts
      installed=$(which node)
    fi

    if [ -z "$installed" ]; then
      echo "ERROR: node is not installed."
      echo "Install node using '$brew install nvm'."
      exit 1
    fi
  fi

  nvm=nvm
  node="$installed"
}

npm_installed() {
  node_installed
  check_installed npm "$nvm install-latest-npm"
  npm="$installed"
}

npm_install() {
  npm_installed
  check_installed "$1" "$npm install --location=global ${2:-$1}"
}

yq_installed() {
  brew_install yq
  yq="$installed"
}

swagger_cli_installed() {
  npm_install swagger-cli "@apidevtools/swagger-cli"
  swagger_cli="$installed"
}

check() {
  for i in $*; do
    "${i}_installed"
  done
}
