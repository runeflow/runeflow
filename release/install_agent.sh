#!/bin/bash
# (C) Runeflow. 2018
# All rights reserved
# Licensed under MIT License (see LICENSE)
# Runeflow Agent installation script: install and set up the Agent on supported
# Linux distributions using the package manager and Runeflow repositories.

set -e
logfile="runeflow-install.log"
if [ $(command -v curl) ]; then
    dl_cmd="curl -f"
else
    dl_cmd="wget --quiet"
fi

# Set up a named pipe for logging
npipe=/tmp/$$.tmp
mknod $npipe p

# Log all output to a log for error checking
tee <$npipe $logfile &
exec 1>&-
exec 1>$npipe 2>&1
trap "rm -f $npipe" EXIT

function on_error() {
    printf "\033[31m$ERROR_MESSAGE
Oh no! It looks like you hit an issue when trying to install the Runeflow
Agent. Don't worry, we're going to do everything we can to make it right.

Please reach out to us with the contents of runeflow-install.log. You can find
us at:
  * https://www.reddit.com/r/runeflow
  * help@runeflow.com

\n\033[0m\n"
}
trap on_error ERR

# OS/Distro Detection
# Try lsb_release, fallback with /etc/issue then uname command
KNOWN_DISTRIBUTION="(Debian|Ubuntu|RedHat|CentOS|openSUSE|Amazon|Arista|SUSE)"
DISTRIBUTION=$(lsb_release -d 2>/dev/null | grep -Eo $KNOWN_DISTRIBUTION  || grep -Eo $KNOWN_DISTRIBUTION /etc/issue 2>/dev/null || grep -Eo $KNOWN_DISTRIBUTION /etc/Eos-release 2>/dev/null || grep -m1 -Eo $KNOWN_DISTRIBUTION /etc/os-release 2>/dev/null || uname -s)

if [ $DISTRIBUTION = "Darwin" ]; then
    printf "\033[31mThis script does not support installing on the Mac.\033[0m\n"
    exit 1;
elif [ -f /etc/debian_version -o "$DISTRIBUTION" == "Debian" -o "$DISTRIBUTION" == "Ubuntu" ]; then
    OS="Debian"
elif [ -f /etc/redhat-release -o "$DISTRIBUTION" == "RedHat" -o "$DISTRIBUTION" == "CentOS" -o "$DISTRIBUTION" == "Amazon" ]; then
    OS="RedHat"
# Some newer distros like Amazon may not have a redhat-release file
elif [ -f /etc/system-release -o "$DISTRIBUTION" == "Amazon" ]; then
    OS="RedHat"
# Arista is based off of Fedora14/18 but do not have /etc/redhat-release
elif [ -f /etc/Eos-release -o "$DISTRIBUTION" == "Arista" ]; then
    OS="RedHat"
# openSUSE and SUSE use /etc/SuSE-release or /etc/os-release
elif [ -f /etc/SuSE-release -o "$DISTRIBUTION" == "SUSE" -o "$DISTRIBUTION" == "openSUSE" ]; then
    OS="SUSE"
fi

# Root user detection
if [ $(echo "$UID") = "0" ]; then
    sudo_cmd=''
else
    sudo_cmd='sudo'
fi

# Install the necessary package sources
if [ $OS = "Debian" ]; then
    printf "\033[32m\n* Installing apt-transport-https\n\033[0m\n"
    $sudo_cmd apt-get update || printf "\033[31m'apt-get update' failed, the script will not install the latest version of apt-transport-https.\033[0m\n"
    $sudo_cmd apt-get install -y apt-transport-https

    printf "\033[32m\n* Installing APT package sources for Runeflow\n\033[0m\n"
    $sudo_cmd sh -c "wget -qO - 'https://bintray.com/user/downloadSubjectPublicKey?username=bintray' | apt-key add -"
    $sudo_cmd sh -c "echo 'deb https://dl.bintray.com/runeflow/debian unstable main' | tee /etc/apt/sources.list.d/runeflow.list"
    printf "\033[32m\n* Installing the Runeflow Agent package\n\033[0m\n"
    ERROR_MESSAGE="ERROR
Failed to update the sources after adding the Runeflow repository.
This may be due to any of the configured APT sources failing -
see the logs above to determine the cause.
If the failing repository is Runeflow, please contact Runeflow support.
*****
"
    $sudo_cmd apt-get update
    ERROR_MESSAGE="ERROR
Failed to install the Runeflow package, sometimes it may be
due to another APT source failing. See the logs above to
determine the cause.
If the cause is unclear, please contact Runeflow support.
*****
"
    $sudo_cmd apt-get install -y --force-yes runeflow
    ERROR_MESSAGE=""
else
    printf "\033[31mYour OS or distribution are not supported by this install script.
    Please email help@runeflow.com and let us know which OS you are trying to use
    Runeflow on.\033[0m\n"
    exit;
fi

## Sign up for a Runeflow account
printf "\033[32m* Register for Free Runeflow Account...\n\033[0m\n"
read -p "Email: " EMAIL
read -p "First Name: " FIRST_NAME
read -p "Last Name: " LAST_NAME
# Optional in this script
FIRST_NAME=${FIRST_NAME:-" "}
LAST_NAME=${LAST_NAME:-" "}
$sudo_cmd runeflow register --email=$EMAIL --first-name=$FIRST_NAME --last-name=$LAST_NAME

printf "\033[32m* Add Agent to account...\n\033[0m\n"
$sudo_cmd runeflow auth --email=$EMAIL

printf "\033[32m* Restarting agent with new configuration...\n\033[0m\n"
$sudo_cmd service runeflow restart

stop_instructions="$sudo_cmd service runeflow stop"
start_instructions="$sudo_cmd service runeflow start"

# Metrics are submitted, echo some instructions and exit
printf "\033[32m
Your Agent is running and functioning properly. It will continue to run in the
background and submit metrics to Runeflow.
If you ever want to stop the Agent, run:
    $stop_instructions
And to run it again run:
    $start_instructions
\033[0m"

