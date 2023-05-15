# frider
Manage scripts and interact with applications using frida.

`frider` is a tool that allows you to manage your scripts and keep them in one place. 
Also, using [frida-go-bindings](https://github.com/frida/frida-go) it allows you to interact with the application 
using Frida.

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
or from some other file.

## backup

`backup` allows you to quickly export data(`backup export`) from the database so that it can be transferred to another device by using `backup import` subcommand.

## script

`script` subcommand is the main command that interacts with the database, stores the script, prints them to the screen as well as deleting them.