# Runeflow

[![Build Status](https://travis-ci.org/runeflow/runeflow.svg?branch=master)](https://travis-ci.org/runeflow/runeflow)
[![Godoc Reference](https://godoc.org/github.com/runeflow/runeflow?status.svg)](https://godoc.org/github.com/runeflow/runeflow)

## About Runeflow

Runeflow ✨ automagically sets up monitoring for your self-hosted Wordpress
sites. It monitors your server usage, like CPU, memory, and disk space, and
will send you an email so you can fix problems that occur. It can also
routinely fix common problems on your server (like restarting Apache, or MySQL
when memory usage gets too high). This repo is the server agent that works with
the [Runeflow](https://runeflow.com) web service.

**Please note**: We take Runeflow's security and our users' trust very
seriously.  If you believe you have found a security issue in Runeflow, please
responsibly disclose by contacting us at security@runeflow.com.

![Runeflow dashboard](https://assets.runeflow.com/product.png)

## Installation

To get started with Runeflow, install the CLI on your servers. The CLI will
automatically pickup your Wordpress sites and also collect CPU, memory, swap,
and disk usage metrics. These important metrics allow you to better understand
what the root cause of a problem actually is so you can triage it faster.

There are several easy ways to install Runeflow. Keep in mind we currently only
support Ubuntu and Debian Linux distributions.

### Recommended Installation

These commands will import our code signing key, add our [apt repository on
Bintray](https://bintray.com/runeflow/debian), and install the CLI.

```
$ wget -qO - 'https://bintray.com/user/downloadSubjectPublicKey?username=bintray' | sudo apt-key add -
$ echo "deb https://dl.bintray.com/runeflow/debian unstable main" | sudo tee /etc/apt/sources.list.d/runeflow.list
$ sudo apt-get update
$ sudo apt-get install runeflow
```

If you don't have a Runeflow account, you'll need to create one. Note that you
only ever need to do this once, e.g. on the first server you want to authorize.

```
$ runeflow register
Email: me@example.com
First Name: Ben
Last Name: Burwell
```

After you have a Runeflow account, simply run `runeflow auth` on each of your
servers to connect them to your account. Servers will show up as Pending on your
dashboard where you must go to claim them before they will start monitoring your
sites.

```
$ runeflow auth
Email: me@example.com
```

You can also run the `auth` command by passing the email address on the command
line. This is more convenient for scripting.

```
$ runeflow auth --email me@example.com
```

### Alternate one-liner setup

We do not recommend using "curl to bash" as a means to install software on
production servers due to the security risks of running unauthenticated code.
However, this can be a useful tool for quickly testing out Runeflow on staging
servers.

```
bash -c "$(curl -s https://i.runeflow.com)"
```

We encourage you to examine the script before running it! You can also download
the script locally and then run it to ensure you know what is being executed.
The script executes the steps above with some extra checks for different
distributions.

### Install from source

While you will not get the benefits of updates for new functionality with
`apt-get upgrade`, you can also build the Runeflow agent from source. You may
want to do this if you are developing Runeflow, or if you have exceptional
security concerns. To do this requires a functioning [Go development
environment](https://golang.org) (1.10 or later) and a Linux build environment.

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

You may also want to move the agent to `/usr/bin/runeflow`. You can then use
the
[Systemd](https://github.com/runeflow/runeflow/blob/master/release/runeflow/etc/systemd/system/runeflow.service)
service to load runeflow as a system service.

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
RUNEFLOW_ENDPOINT=https://my-server.com/runeflow ./build/runeflow
```

# Contributing

We're glad you want to contribute! Nothing would make us happier to see more
people get involved.

First: if you're unsure or afraid of anything, just ask or submit the issue or
pull request anyways. You won't be yelled at for giving your best effort. The
worst that can happen is that you'll be politely asked to change something. We
appreciate any sort of contributions, and don't want a wall of rules to get in
the way of that. Here are some great ideas to get you started:

- Report potential bugs.
- Suggest product enhancements.
- Increase our test coverage.
- Fix a [bug](https://github.com/runeflow/runeflow/labels/bug).
