# dynamodb-clone

[![Build Status](https://travis-ci.org/kishaningithub/dynamodb-clone.svg?branch=master)](https://travis-ci.org/kishaningithub/dynamodb-clone)
[![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)
[![Go Report Card](https://goreportcard.com/badge/github.com/kishaningithub/dynamodb-clone)](https://goreportcard.com/report/github.com/kishaningithub/dynamodb-clone)
[![Downloads](https://img.shields.io/github/downloads/kishaningithub/dynamodb-clone/latest/total.svg)](https://github.com/kishaningithub/dynamodb-clone/releases)
[![Latest release](https://img.shields.io/github/release/kishaningithub/dynamodb-clone.svg)](https://github.com/kishaningithub/dynamodb-clone/releases)

Clone your dynamo db to local

## Table of Contents

- [dynamodb-clone](#dynamodb-clone)
  - [Table of Contents](#table-of-contents)
  - [Install](#install)
    - [Using Homebrew](#using-homebrew)
    - [Binary Installation](#binary-installation)
  - [Example](#example)
  - [Usage](#usage)
  - [Maintainers](#maintainers)
  - [Contribute](#contribute)
  - [License](#license)

## Install

### Using Homebrew

```bash
brew tap kishaningithub/homebrew-tap
brew install dynamodb-clone
```

### Binary Installation

```bash
# binary will be $GOPATH/bin/dynamodb-clone
curl -sfL https://raw.githubusercontent.com/kishaningithub/dynamodb-clone/master/install.sh | sh -s -- -b $GOPATH/bin

# or install it into ./bin/
curl -sfL https://raw.githubusercontent.com/kishaningithub/dynamodb-clone/master/install.sh | sh -s

# In alpine linux (as it does not come with curl by default)
wget -O - -q https://raw.githubusercontent.com/kishaningithub/dynamodb-clone/master/install.sh | sh -s
```

## Example 

```bash
$ AWS_REGION=eu-west-1 dynamodb-clone -t employee-details -e http://localhost:4569
```
To load the shared config, `export AWS_SDK_LOAD_CONFIG="true"` before running the above command. 
This is useful to load the shared credentials file from `~/.aws/credentials`.

## Usage

```bash
dynamodb-clone -h
Usage:
  main [OPTIONS]

Application Options:
  -t, --table-name=   Name of the dynamo db table
  -e, --endpoint-url= Endpoint url of destination dynamodb instance

Help Options:
  -h, --help          Show this help message
```

## Maintainers

- [@kishaningithub](https://github.com/kishaningithub)

## Contribute

PRs accepted.

Small note: If editing the README, please conform to the [standard-readme](https://github.com/RichardLitt/standard-readme) specification.

## License

MIT © 2018 Kishan B
