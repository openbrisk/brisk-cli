[![Build Status](https://api.travis-ci.org/openbrisk/brisk-cli.svg?branch=master)](https://travis-ci.org/openbrisk/brisk-cli)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)
[![OpenBrisk](https://img.shields.io/badge/openbrisk-serverless-blue.svg)](https://www.openbrisk.org)

# OpenBrisk CLI
A command-line interface for the OpenBrisk serverless framework.

## Setup Development Environment

- Install Docker

- Install Golang
  - https://golang.org/dl

Alter installation: put this in your `$HOME/.bash_profile`

```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

- Clone and build `brisk-cli`

```bash
mkdir -p $GOPATH/src/github.com/openbrisk
cd $GOPATH/src/github.com/openbrisk
git clone https://github.com/openbrisk/brisk-cli
cd brisk-cli
go build
```
- Install Dep

https://github.com/golang/dep

- Install Travis

`brew install travis`
