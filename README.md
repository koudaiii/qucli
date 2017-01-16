# dockerepos

[![Build Status](https://travis-ci.org/koudaiii/dockerepos.svg?branch=master)](https://travis-ci.org/koudaiii/ddockerepos)
[![Docker Repository on Quay](https://quay.io/repository/koudaiii/dockerepos/status "Docker Repository on Quay")](https://quay.io/repository/koudaiii/dockerepos)
[![GitHub release](https://img.shields.io/github/release/koudaiii/dockerepos.svg)](https://github.com/koudaiii/dockerepos/releases)

## Description

Manage repository in Quay

## Table of Contents

* [Requirements](#requirements)
* [Installation](#installation)
  + [Using Homebrew (OS X only)](#using-homebrew-os-x-only)
  + [Precompiled binary](#precompiled-binary)
  + [From source](#from-source)
  + [Run in a Docker container](#run-in-a-docker-container)
* [Usage](#usage)
  + [Options](#options)
* [Development](#development)
* [Author](#author)
* [License](#license)

## Requirements

- Enviroment QUAY_API_TOKEN
- Get Api Token. [Applications and Tokens](https://docs.quay.io/api/)

```bash
$ export QUAY_API_TOKEN=foobar
```

## Installation

### Using Homebrew (OS X only)

Formula is available at [koudaiii/homebrew-tools](https://github.com/koudaiii/homebrew-tools).

```bash
$ brew tap koudaiii/tools
$ brew install dockerepos
```

### Precompiled binary

Precompiled binaries for Windows, OS X, Linux are available at [Releases](https://github.com/koudaiii/dockerepos/releases).

### From source
To install, use `go get`:

```bash
$ go get -d github.com/koudaiii/dockerepos
$ cd $GOPATH/src/github.com/koudaiii/dockerepos
$ make deps
$ make install
```

### Run in a Docker container

docker image is available at [quay.io/koudaiii/dockerepos](https://quay.io/repository/koudaiii/dockerepos).

```bash
# -t is required to colorize logs
$ docker run \
    --rm \
    -t \
    -e QUAY_API_TOKEN=foobar \
    quay.io/koudaiii/dockerepos:latest
```

## Usage

```bash
usage: dockerepos [--version] [--help] <command> [<args>]

Available commands are:
    add-team       Add team in repository
    add-user       Add user in repository
    create         Create repository in Quay
    delete         Delete repository in Quay
    delete-team    Delete team in repository
    delete-user    Delete user in repository
    get            Get repository and Permissions in Quay
    version        Print dockerepos version and quit
```

### `create`

Create repository in Quay

With `--visibility` option, you can `public` or `private`

```bash
$ dockerepos create quay.io/wantedly/test --visibility private
Created! quay.io/wantedly/test
```

### `get`

Get repository and Permissions in Quay

```bash
$ dockerepos get quay.io/wantedly/test
Repository:
	quay.io/wantedly/test
Visibility:
	private
Permissions:
	koudaiii(admin)
```

### `add-user`

Add user in repository

With `--role` option, you can `read` or `write` or `admin`

```bash
$ dockerepos add-user quay.io/wantedly/test dtan4 --role write
Added! dtan4(write) in quay.io/wantedly/test
```

```bash
$ dockerepos get quay.io/wantedly/test
Repository:
	quay.io/wantedly/test
Visibility:
	private
Permissions:
	koudaiii(admin)
	dtan4(write)
```

### `add-team`

Add team in repository

With `--role` option, you can `read` or `write` or `admin`

```bash
$ dockerepos add-team quay.io/wantedly/test infrastructure --role write
Added! infrastructure(write) in quay.io/wantedly/test
```

```bash
$ dockerepos get quay.io/wantedly/test
Repository:
	quay.io/wantedly/test
Visibility:
	private
Permissions:
	koudaiii(admin)
	dtan4(write)
	infrastructure(write)
```

### `delete-user`

Delete user from repository

```bash
$ dockerepos delete-user quay.io/wantedly/test dtan4
Deleted! dtan4 in quay.io/wantedly/test
```

```bash
$ dockerepos get quay.io/wantedly/test
Repository:
	quay.io/wantedly/test
Visibility:
	private
Permissions:
	koudaiii(admin)
	infrastructure(write)
```

### `delete-team`

Delete team from repository

```bash
$ dockerepos delete-team quay.io/wantedly/test infrastructure
Deleted! infrastructure in quay.io/wantedly/test
```

```bash
$ dockerepos get quay.io/wantedly/test
Repository:
	quay.io/wantedly/test
Visibility:
	private
Permissions:
	koudaiii(admin)
```

### `delete`

Delete repository in Quay

```bash
$ dockerepos delete quay.io/wantedly/test
Deleted! quay.io/wantedly/test
```

```bash
$ dockerepos get quay.io/wantedly/test
err: HTTP error!
URL: https://quay.io/api/v1/repository/wantedly/test
status code: 404
body:
{"status": 404, "error_message": "Not Found", "title": "not_found", "error_type": "not_found", "detail": "Not Found", "type": "https://quay.io/api/v1/error/not_found"}
```

### Options

|Option|Description|Required|Default|
|---------|-----------|-------|-------|
|`--visibility=VISIBILITY`| "visibility set to 'public' or 'private'||`public`|
|`--role=ROLE`|role to use for the user or team=  ['read', 'write', 'admin']||`read`|
|`--help`|Print command line usage|||
|`-v`, `--version`|Print version|||

## Development

Clone this repository and build using `make`.

```bash
$ go get -d github.com/koudaiii/dockerepos
$ cd $GOPATH/src/github.com/koudaiii/dockerepos
$ make
```

## Contribution

1. Fork ([https://github.com/koudaiii/dockerepos/fork](https://github.com/koudaiii/dockerepos/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[koudaiii](https://github.com/koudaiii)

## License

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
