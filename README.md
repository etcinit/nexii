# Nexus CLI client

CLI client for the [Nexus Configuration Server](https://github.com/etcinit/nexus) written in Go

## Installing

- Setup your `$GOPATH`
- Make sure `$GOPATH/bin` is in your `$PATH`
- `$ go install github.com/etcinit/nexii`

## Usage

```
NAME:
   nexii - A Nexus Configuration Server CLI client

USAGE:
   nexii [global options] command [command options] [arguments...]

VERSION:
   0.1.0

AUTHOR(S):
   Eduardo Trujillo <ed@chromabits.com>

COMMANDS:
   list, ls	List all keys
   get, g, cat	Get and output a specific key
   download, d	Get and save a specific key
   info, i, app	Get information about the application
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --endpoint 		Nexus server endpoint [$NEXUS_SERVER, $NEXUS_ENDPOINT]
   --token 		Nexus server API token [$NEXUS_APIKEY, $NEXUS_TOKEN]
   --help, -h		show help
   --version, -v	print the version
```
