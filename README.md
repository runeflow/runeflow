# Runeflow

[![Build Status](https://travis-ci.org/runeflow/runeflow.svg?branch=master)](https://travis-ci.org/runeflow/runeflow)
[![Godoc Reference](https://godoc.org/github.com/runeflow/runeflow?status.svg)](https://godoc.org/github.com/runeflow/runeflow)

# About Runeflow
Runeflow ✨ automagically sets up monitoring for your self-hosted Wordpress
sites. It monitors your server usage, like CPU, memory, and disk space, and
will send you an email so you can fix problems that occur. This repo is the
server agent that works with the [Runeflow](https://runeflow.com) web service. 

Please note: We take Runeflow's security and our users' trust very seriously.
If you believe you have found a security issue in Runeflow, please responsibly
disclose by contacting us at security@runeflow.com.

## Quickstart: Install the Runeflow CLI
The easy 1-liner to install the Runeflow agent. Supports Ubuntu and Debian.
```
curl -sS https://i.runeflow.com | bash
```
We encourage you to examine the script before running it! You can also download
the script locally and then run it to ensure you know what is being executed.
Below are step by step instructions to also install the agent.

### Install on Ubuntu / Debian
First, add the repository.
```
$ wget -qO - 'https://bintray.com/user/downloadSubjectPublicKey?username=bintray' | sudo apt-key add -
$ echo "deb https://dl.bintray.com/runeflow/debian unstable main" | sudo tee /etc/apt/sources.list.d/runeflow.list
$ sudo apt-get update
$ sudo apt-get install runeflow
```

## Sign up for a Runeflow account
Once installed, the Agent needs to be registed to the
[Runeflow](https://runeflow.com) webservice.

```
$ runeflow register
Email: me@example.com
First Name: Ben
Last Name: Burwell
```

## Add your server to Runeflow
Finally, add your agent to your account. You can authenticate interactively by
running `runeflow auth`. You'll be prompted for the email address of the user
you'd like to associate to, and you'll need to confirm the new server on the
dashboard.

```
$ runeflow auth
Email: me@example.com
```

You can also run the `auth` command by passing the email address on the command
line. This is more convenient for scripting.

```
$ runeflow auth --email me@example.com
```

# Development
Building the agent from source requires a [Go](https://golang.org/) dev
environment.

* Go 1.10 or later
* A Linux build environment

```
$ git clone git@github.com:runeflow/runeflow.git $GOPATH/src/github.com/runeflow/runeflow
$ cd $GOPATH/src/github.com/runeflow/runeflow
$ go get -u ./...
$ make
```

The make command will only work in a Linux build environment. If you want to
build on another system, run a more generic build command:
```
$ go build -o "build/runeflow" ./cli
```

You can run the Agent with `./build/runeflow`. You may want to change the
endpoints it communicates with, see Configuration.

# Testing
To run the tests, run:
```
go test -v ./...
```
Contributed code should have tests.

# Configuration
The Runeflow agent looks for a configuration file at
`/etc/runeflow/runeflow.yml`. The available configurations can be seen
[here](https://github.com/runeflow/runeflow/blob/master/config/config.go).
These values can also be set with environment variables. For example:

```
RUNEFLOW_ENDPOING=https://my-server.com/runeflow ./build/runeflow
```

# Contributing
We're glad you want to contribute! Nothing would make us happier to see more
people get involved.

First: if you're unsure or afraid of anything, just ask or submit the issue or
pull request anyways. You won't be yelled at for giving your best effort. The
worst that can happen is that you'll be politely asked to change something. We
appreciate any sort of contributions, and don't want a wall of rules to get in
the way of that. Here are some great ideas to get you started:

* Report potential bugs.
* Suggest product enhancements.
* Increase our test coverage.
* Fix a [bug](https://github.com/runeflow/runeflow/labels/bug).

