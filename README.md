# frider
Manage scripts and interact with applications using frida.

`frider` is a tool that allows you to manage your scripts and keep them in one place. 
Also, using [frida-go-bindings](https://github.com/frida/frida-go) it allows you to interact with the application 
using Frida. Prior to installing `frider` you need to have devkit downloaded, according to the instructions inside [frida-go](https://github.com/frida/frida-go).

# Installation

```bash
$ go install github.com/nsecho/frider@latest
```

# Usage

```bash
$ frider --help
Frida helper tool

Usage:
  frider [command]

Available Commands:
  app         Interact with device applications using frida
  backup      Manage database backups
  help        Help about any command
  script      Manage database scripts

Flags:
  -h, --help   help for frider

Use "frider [command] --help" for more information about a command.
```

## app

`app` subcommand allows you to list applications, kill specific PID, load scripts from the database 
or use raw script file.

```bash
$ frider app --help
Interact with device applications using frida

Usage:
  frider app [command]

Available Commands:
  kill        Kill specific process or application
  load        Load script from the database to the application

Flags:
  -h, --help   help for app

Use "frider app [command] --help" for more information about a command.
```

## backup

`backup` allows you to quickly export data(`backup export`) from the database so that it can be transferred to another device by using `backup import` subcommand.

```bash
$ frider backup --help
Manage database backups

Usage:
  frider backup [command]

Available Commands:
  export      Export scripts
  import      Import exported .frider scripts

Flags:
  -h, --help   help for backup

Use "frider backup [command] --help" for more information about a command.
```

## script

`script` subcommand is the main command that interacts with the database, stores the script, prints them to the screen as well as deleting them.

```bash
$ frider script --help
Manage database scripts

Usage:
  frider script [command]

Available Commands:
  delete      Delete script from the database
  get         Get content of the script as file
  print       Print all the scripts
  save        Save script
  show        Show specific script

Flags:
  -h, --help   help for script

Use "frider script [command] --help" for more information about a command
```
