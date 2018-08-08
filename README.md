# legacy_exec_history

osquery table extension that presents the history of 32-bit applications run on a 10.14 machine

## Introduction

At the 2018 WWDC State of the Union event, following the announcement of macOS Mojave (macOS 10.14) in the keynote earlier in the day, Apple vice president of software Sebastien Marineau revealed Mojave will be "the last release to support 32-bit at all.". Since macOS 10.13.4, Apple has provided the ability to set your machine to 64-bit only mode for testing. For most users this is not a very convenient way to test. As of 10.14 the System Information application has a new "Legacy Software" section that shows you all of the 32-bit applications that have been run on the machine.

This new "Legacy Software" information provides great insight for Mac Admins into what 32-bit applications their users are running so that they can work with vendors to get software updated prior to the release of macOS 10.15 in the fall of 2019.

This legacy information is just stored in a sqlite database located in `/var/db/SystemPolicyConfiguration/` called `ExecPolicy`. You can actually query this information yourself from the command line but it does require root access.

```
sudo sqlite3 /var/db/SystemPolicyConfiguration/ExecPolicy 'select * from legacy_exec_history_v3'
```

In order to provide a more convienant method to get this information this osquery table plugin was created. It does not require any special privileges because just like the System Information application it uses the `SystemPolicy.framework` to talk directly to `syspolicyd` to get this information. By making this an osquery extension it allows us to query this information uniformly across an entire fleet of macOS machines.

## Extension Overview

This table extension provides a new read only table called `legacy_exec_history` with the following layout:

| Column           | Type | Description                                                      |
|------------------|------|------------------------------------------------------------------|
| exec_path        | TEXT | Path of the application if it was executed directly              |
| mmap_path        | TEXT | Path of the application if it was loaded into memory             |
| signing_id       | TEXT | The Signing ID from the code signing blob                        |
| team_id          | TEXT | The Team ID from the code signing blob                           |
| cd_hash          | TEXT | The primary CD Hash from the code signing blob                   |
| responsible_path | TEXT | Path of the application that launched or mmaped this application |
| developer_name   | TEXT | Name from the signing certificate                                |
| last_seen        | TEXT | The timestamp of the last time the applications was used         |

## Building

This extension will only compile on macOS 10.14. Go 1.10.3 currently has a [bug](https://github.com/golang/go/issues/25908) compiling on macOS 10.14 so you will need to use the Go 1.11 beta until 1.10.4 is released which has the fix backported in it. [Dep](https://github.com/golang/dep) is used for dependency management and building is fairly straigth forward. Just follow the steps below.

```
# if GOPATH is unset:
# export GOPATH=${HOME}/go

# clone repo into GOPATH:
git clone https://github.com/knightsc/legacy_exec_history.git $GOPATH/src/github.com/knightsc/legacy_exec_history
cd $GOPATH/src/github.com/knightsc/legacy_exec_history

# get dependencies and build
dep ensure
go build
```

## Testing

To test this code, start an osquery shell and find the path of the osquery extension socket:

```
osqueryi --nodisable_extensions
osquery> select value from osquery_flags where name = 'extensions_socket';
+-----------------------------------+
| value                             |
+-----------------------------------+
| /Users/USERNAME/.osquery/shell.em |
+-----------------------------------+
```

Then start the Go extension and have it communicate with osqueryi via the extension socket that you retrieved above:

```
./legacy_exec_history -socket /Users/USERNAME/.osquery/shell.em
```

Alternatively, you can also autoload your extension when starting an osquery shell:

```
osqueryi --extension ./legacy_exec_history
```

## Installing

Installation of this extension should be similar to any other osquery extension. Start by reviewing the [Using Extensions](https://osquery.readthedocs.io/en/stable/deployment/extensions/) guide from the osquery project.

Essentially you will want to rename the binary to legacy_exec_history.ext (or use the pre-built release version) and put it somewhere on your machine with the correct permissions. You can then add an entry for this extension in your autoload file.
