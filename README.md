# wwwredirector

A simple HTTP(S) redirect service: prepends `www.` to any requested hostname

## Installation

Assuming a working go 1.6 environment:

```sh
go get github.com/dgoodlad/wwwredirector
```

## Usage

```sh
# By default, we listen on 0.0.0.0:8080
wwwredirector

# You can specify interface and port with environment variables
LISTEN_INTERFACE=127.0.0.1 PORT=3000 wwwredirector
```
