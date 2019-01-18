# system_policy

osquery table extension that allows querying of information from the macOS private SystemPolicy.framework.

## Introduction

Beginning with macOS 10.13 Apple started enforcing certain system policies. One of those policies is the requirement that loaded kernel extensions need to be approved by the user. These system policies are also reponsible for keeping track of what 32-bit applications have been run. As announced at the 2018 WWDC, macOS 10.14 will be the last version of macOS to support 32-bit applications. This information can be viewed in the System Information application.

This information is just stored in sqlite databases located in `/var/db/SystemPolicyConfiguration/`. You can actually query this information yourself from the command line but it does require root access.

```
sudo sqlite3 /var/db/SystemPolicyConfiguration/ExecPolicy 'select * from legacy_exec_history_v4'
```

In order to provide a more convienant method to get this information this osquery table plugin was created. It does not require any special privileges because just like the System Information application it uses the `SystemPolicy.framework` to talk directly to `syspolicyd` to get this information. By making this an osquery extension it allows us to query this information uniformly across an entire fleet of macOS machines.

## Extension Overview

This table extension provides two new read only tables with the following layout:

### kext\_policy
Data only returned on macOS 10.13 and higher

| Column           | Type    | Description                                         |
|------------------|---------|-----------------------------------------------------|
| developer_name   | TEXT    | Name of the developer who signed the kext           |
| application_name | TEXT    | Name of the application that tried to load the kext |
| application_path | TEXT    | Path to the application that tried to load the kext |
| team_id          | TEXT    | The Team ID from the code signing blob              |
| bundle_id        | TEXT    | The Bundle ID of the kext                           |
| allowed          | INTEGER | Whether the kext has been user/mdm approved or not  |
| reboot_required  | INTEGER | If the kext requires a reboot                       |
| modified         | INTEGER | Unknown                                             |

### legacy\_exec\_history
Data only returned on macOS 10.14 and higher

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

Since this project makes use of the new [modules](https://github.com/golang/go/wiki/Modules) functionality Go 1.11 or above will be required to compile. It should compile on all versions of macOS. Building is as simple as the following:

```
git clone https://github.com/knightsc/system_policy.git 
cd system_policy
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
./system_policy -socket /Users/USERNAME/.osquery/shell.em
```

Alternatively, you can also autoload your extension when starting an osquery shell:

```
osqueryi --extension ./system_policy
```

## Installing

Installation of this extension should be similar to any other osquery extension. Start by reviewing the [Using Extensions](https://osquery.readthedocs.io/en/stable/deployment/extensions/) guide from the osquery project.

Essentially you will want to rename the binary to system_policy.ext (or use the pre-built release version) and put it somewhere on your machine with the correct permissions. You can then add an entry for this extension in your autoload file.
