# Runeflow

[![Build Status](https://travis-ci.org/runeflow/runeflow.svg?branch=master)](https://travis-ci.org/runeflow/runeflow)
[![Godoc Reference](https://godoc.org/github.com/runeflow/runeflow?status.svg)](https://godoc.org/github.com/runeflow/runeflow)

## Install the Runeflow CLI

### Ubuntu

```
$ wget -O - https://download.runeflow.com/linux/ubuntu/gpg | apt-key add -
$ codename=$(awk -F '=' '$1=="VERSION_CODENAME"{print $2}' < /etc/os-release)
$ echo "deb https://dl.bintray.com/runeflow/debian $codename main" > /etc/apt/sources.list.d/runeflow.list
```

## Sign up for a Runeflow account

```
$ runeflow register
Email: me@example.com
First Name: Ben
Last Name: Burwell
```

## Add your server to Runeflow

You can authenticate interactively by running `runeflow auth`. You'll be
prompted for the email address of the user you'd like to associate to, and
you'll need to confirm the new server on the dashboard.

```
$ runeflow auth
Email: me@example.com
```

You can also run the `auth` command by passing the email address on the command
line. This is more convenient for scripting.

```
$ runeflow auth --email me@example.com
```

##### Coming soon!

> Rather than using your email address and needing to confirm the new server, you
> can generate a token for your account that can be used to register new servers.
