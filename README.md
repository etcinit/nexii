# Nexus CLI client

CLI client for the [Nexus Configuration Server](https://github.com/etcinit/nexus) written in Go

[![wercker status](https://app.wercker.com/status/d9243f7ee424999e29466af3c3e0b060/m "wercker status")](https://app.wercker.com/project/bykey/d9243f7ee424999e29466af3c3e0b060)

## Installing

- Setup your `$GOPATH`
- Make sure `$GOPATH/bin` is in your `$PATH`
- `$ go get github.com/etcinit/nexii`

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
