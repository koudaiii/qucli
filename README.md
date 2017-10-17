# qucli

[![Build Status](https://travis-ci.org/koudaiii/qucli.svg?branch=master)](https://travis-ci.org/koudaiii/dqucli)
[![Docker Repository on Quay](https://quay.io/repository/koudaiii/qucli/status "Docker Repository on Quay")](https://quay.io/repository/koudaiii/qucli)
[![GitHub release](https://img.shields.io/github/release/koudaiii/qucli.svg)](https://github.com/koudaiii/qucli/releases)

## Description

Manage repositories in Quay.io

## Table of Contents

* [qucli](#qucli)
  * [Description](#description)
  * [Table of Contents](#table-of-contents)
  * [Requirements](#requirements)
  * [Installation](#installation)
    * [Using Homebrew (OS X only)](#using-homebrew-os-x-only)
    * [Precompiled binary](#precompiled-binary)
    * [From source](#from-source)
    * [Run in a Docker container](#run-in-a-docker-container)
  * [Usage](#usage)
    * [list](#list)
    * [get](#get)
    * [create](#create)
    * [delete](#delete)
    * [add-user](#add-user)
    * [add-team](#add-team)
    * [delete-user](#delete-user)
    * [delete-team](#delete-team)
    * [add-notification](#add-notification)
    * [test-notification](#test-notification)
    * [delete-notification](#delete-notification)
    * [Options](#options)
  * [Development](#development)
  * [Contribution](#contribution)
  * [Author](#author)
  * [License](#license)

## Requirements

- Enviroment QUAY_API_TOKEN
- Get Api Token. [Applications and Tokens](https://docs.quay.io/api/)

```shell-session
$ export QUAY_API_TOKEN=foobar
```

if Quay Enterprise user, add Enviroment QUAY_HOSTNAME or `--hostname`

```shell-session
$ export QUAY_HOSTNAME=quay.example.com
or
$ qucli xxx --hostname=quay.example.com
```

## Installation

### Using Homebrew (OS X only)

Formula is available at [koudaiii/homebrew-tools](https://github.com/koudaiii/homebrew-tools).

```shell-session
$ brew tap koudaiii/tools
$ brew install qucli
```

### Precompiled binary

Precompiled binaries for Windows, OS X, Linux are available at [Releases](https://github.com/koudaiii/qucli/releases).

### From source
To install, use `go get`:

```shell-session
$ go get -d github.com/koudaiii/qucli
$ cd $GOPATH/src/github.com/koudaiii/qucli
$ make deps
$ make install
```

### Run in a Docker container

docker image is available at [quay.io/koudaiii/qucli](https://quay.io/repository/koudaiii/qucli).

```shell-session
# -t is required to colorize logs
$ docker run \
    --rm \
    -t \
    -e QUAY_API_TOKEN=foobar \
    quay.io/koudaiii/qucli:latest
```

## Usage

```shell-session
$ qucli
usage: qucli [--version] [--help] <command> [<args>]

Available commands are:
    add-notification       Add notification in repository
    add-team               Add team in repository
    add-user               Add user in repository
    create                 Create repository in Quay
    delete                 Delete repository in Quay
    delete-notification    Delete notification in repository
    delete-team            Delete team in repository
    delete-user            Delete user in repository
    get                    Get Repository and Permissions and Notifications in Quay
    list                   List repository and Permissions in Quay
    test-notification      Test notification in repository
    version                Print qucli version and quit
```

### `list`

List repository in namespace

With `--is-public` option, you can `true` or `false`

```bsah
$ qucli list koudaiii
NAME				isPublic	DESCRIPTION
quay.io/koudaiii/apig-sample	true
quay.io/koudaiii/qucli	true
quay.io/koudaiii/kubeps		true
quay.io/koudaiii/test		true
```

### `get`

Get repository and Permissions in Quay

```shell-session
$ qucli get koudaiii/test
Repository:
	quay.io/koudaiii/test
Visibility:
	private
Permissions:
	koudaiii(admin)
Notifications:
	b0736be9-f0eb-4c3c-8d33-2e331b1e5b0f	Some title	repo_push	map[]	slack	map[url:https://hooks.slack.com/service/some/token/here]
```

### `create`

Create repository in Quay

With `--visibility` option, you can `public` or `private`

```shell-session
$ qucli create koudaiii/test --visibility private
Created! quay.io/koudaiii/test
```

### `delete`

Delete repository in Quay

```shell-session
$ qucli delete koudaiii/test
Deleted! quay.io/koudaiii/test
```

```shell-session
$ qucli get koudaiii/test
err: HTTP error!
URL: https://quay.io/api/v1/repository/koudaiii/test
status code: 404
body:
{"status": 404, "error_message": "Not Found", "title": "not_found", "error_type": "not_found", "detail": "Not Found", "type": "https://quay.io/api/v1/error/not_found"}
```

### `add-user`

Add user in repository

With `--role` option, you can `read` or `write` or `admin`

```shell-session
$ qucli add-user koudaiii/test dtan4 --role write
Added! dtan4(write) in quay.io/koudaiii/test
```

```shell-session
$ qucli get koudaiii/test
Repository:
	quay.io/koudaiii/test
Visibility:
	private
Permissions:
	koudaiii(admin)
	dtan4(write)
```

### `add-team`

Add team in repository

With `--role` option, you can `read` or `write` or `admin`

```shell-session
$ qucli add-team koudaiii/test infrastructure --role write
Added! infrastructure(write) in quay.io/koudaiii/test
```

```shell-session
$ qucli get koudaiii/test
Repository:
	quay.io/koudaiii/test
Visibility:
	private
Permissions:
	koudaiii(admin)
	dtan4(write)
	infrastructure(write)
```

### `delete-user`

Delete user from repository

```shell-session
$ qucli delete-user koudaiii/test dtan4
Deleted! dtan4 in quay.io/koudaiii/test
```

```shell-session
$ qucli get koudaiii/test
Repository:
	quay.io/koudaiii/test
Visibility:
	private
Permissions:
	koudaiii(admin)
	infrastructure(write)
```

### `delete-team`

Delete team from repository

```shell-session
$ qucli delete-team koudaiii/test infrastructure
Deleted! infrastructure in quay.io/koudaiii/test
```

```shell-session
$ qucli get koudaiii/test
Repository:
	quay.io/koudaiii/test
Visibility:
	private
Permissions:
	koudaiii(admin)
```

### `add-notification`

Add notification in repository with some options.

- `webhook` method

```shell-session
$ qucli add-notification koudaiii/test --event="repo_push" --method="webhook" --url="http://url/goes/here"
Added! 	3c3c142c-2161-42ae-9414-39c787386b5c		repo_push	map[]	webhook	map[url:http://url/goes/here]	in quay.io/koudaiii/test
```

- `slack` method

```shell-session
$ qucli add-notification koudaiii/test --event="repo_push" --method="slack" --url="https://hooks.slack.com/service/{some}/{token}/{here}"
Added! 	61ae254f-89f0-4a36-a439-9b78004f2ab0		repo_push	map[]	slack	map[url:https://hooks.slack.com/service/{some}/{token}/{here}]	in quay.io/koudaiii/test
```

- options

```shell-session
$ qucli add-notification
qucli supported only Quay.io
Usage: add-notification
  qucli add-notification koudaiii/qucli --event="repo_push" --method="webhook" --url="http://url/goes/here"

Option:
  --event string        set 'evnet'.  ['repo_push', 'build_queued', 'build_start', 'build_success', 'build_failure', 'build_cancelled', 'vulnerability_found'].
  --level string        if you use 'vulnerability_found' evnet, A vulnerability must have a severity of the chosen level (highest level is 0).[0-6]
  --ref string          if you use event excluding 'repo_push' event, an optional regular expression for matching the git branch or tag git ref. If left blank, the notification will fire for all builds.(refs/heads/somebranch)|(refs/tags/sometag)
  --method string       set 'method'.  ['webhook', 'slack', 'email'].
  --email string        if you use 'email' method, set E-mail address. 'test@example.com'.
  --url string          if you use 'webhook' or 'slack' method, set url. 'http://url/goes/here' or 'https://hooks.slack.com/service/{some}/{token}/{here}'.
  --title string        The title for a notification is an optional field for a human-readable title for the notification.
```

### `test-notification`

Test notification from repository.

```shell-session
$ qucli test-notification koudaiii/qucli 0c91e746-9d9e-4845-8dff-3c0995976dfa
Test Notification! 0c91e746-9d9e-4845-8dff-3c0995976dfa notification in quay.io/koudaiii/qucli
```

### `delete-notification`

Delete notification from repository.

```shell-session
$ ./bin/qucli delete-notification koudaiii/test 3c3c142c-2161-42ae-9414-39c787386b5c
Deleted! 3c3c142c-2161-42ae-9414-39c787386b5c notification in quay.io/koudaiii/test
```

### Options

|Option|Description|Required|Default|
|---------|-----------|-------|-------|
|`--visibility=VISIBILITY`| "visibility set to 'public' or 'private'||`public`|
|`--role=ROLE`|role to use for the user or team ROLE='read' or 'write' or 'admin'||`read`|
|`--is-public=bool`| repository type is public. `true` or `false`||`true`|
|`--hostname=HOSTNAME`| if Quay Enterprise user, set hostname. ||`quay.io`|
|`--event=EVENT` | set 'evnet'. EVENT='repo_push' or 'build_queued' or 'build_start' or 'build_success' or 'build_failure' or 'build_cancelled' or 'vulnerability_found'. |true||
|`--level=LEVEL`| if you use 'vulnerability_found' evnet, A vulnerability must have a severity of the chosen level (highest level is 0).LEVEL=0-6 |||
|`--ref=REF`|if you use event excluding 'repo_push' event, an optional regular expression for matching the git branch or tag git ref. If left blank, the notification will fire for all builds.(refs/heads/somebranch)|(refs/tags/sometag) |||
|`--method=METHOD`|set 'method'.  METHOD='webhook' or 'slack' or 'email'.|true||
|`--email=EMAIL`|if you use 'email' method, set E-mail address. EMAIL='test@example.com'.|||
|`--url=URL`|if you use 'webhook' or 'slack' method, set url. 'http://url/goes/here' or 'https://hooks.slack.com/service/{some}/{token}/{here}'.|||
|`--title=TITLE`|The title for a notification is an optional field for a human-readable title for the notification.|||
|`--help`|Print command line usage|||
|`-v`, `--version`|Print version|||

## Development

Clone this repository and build using `make`.

```shell-session
$ go get -d github.com/koudaiii/qucli
$ cd $GOPATH/src/github.com/koudaiii/qucli
$ make
```

## Contribution

1. Fork ([https://github.com/koudaiii/qucli/fork](https://github.com/koudaiii/qucli/fork))
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
