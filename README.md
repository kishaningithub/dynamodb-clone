# dynamodb-clone

[![Build Status](https://travis-ci.org/kishaningithub/dynamodb-clone.svg?branch=master)](https://travis-ci.org/kishaningithub/dynamodb-clone)
[![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)
[![Go Report Card](https://goreportcard.com/badge/github.com/kishaningithub/dynamodb-clone)](https://goreportcard.com/report/github.com/kishaningithub/dynamodb-clone)
[![Downloads](https://img.shields.io/github/downloads/kishaningithub/dynamodb-clone/latest/total.svg)](https://github.com/kishaningithub/dynamodb-clone/releases)
[![Latest release](https://img.shields.io/github/release/kishaningithub/dynamodb-clone.svg)](https://github.com/kishaningithub/dynamodb-clone/releases)

Clone dynamodb table to local

## Table of Contents

- [dynamodb-clone](#dynamodb-clone)
  - [Table of Contents](#table-of-contents)
  - [Install](#install)
    - [Using Homebrew](#using-homebrew)
    - [Using Binary](#using-binary)
  - [Example](#example)
  - [Usage](#usage)
  - [Maintainers](#maintainers)
  - [Contribute](#contribute)
  - [License](#license)

## Install

### Using Homebrew

```bash
brew tap kishaningithub/tap
brew install dynamodb-clone
```

### Using Binary

```bash
# All unix environments with curl
curl -sfL https://raw.githubusercontent.com/kishaningithub/dynamodb-clone/master/install.sh | sh -s -- -b /usr/local/bin

# In alpine linux (as it does not come with curl by default)
wget -O - -q https://raw.githubusercontent.com/kishaningithub/dynamodb-clone/master/install.sh | sudo sh -s -- -b /usr/local/bin
```

## Example

```bash
AWS_REGION=eu-west-1 AWS_SDK_LOAD_CONFIG=true dynamodb-clone -t employee-details -e http://localhost:4569
```

Run `export AWS_SDK_LOAD_CONFIG="true"` load the shared config, before running the above command.
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
