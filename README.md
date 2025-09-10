# PingCLI

A lightweight command-line HTTP ping utility built with Go.

## Installation

`go install github.com/faizanfirdousi/pingcli@latest`

## Usage

Basic ping
`pingcli ping --url https://example.com`

With timeout
`pingcli ping --url https://example.com --timeout 10s`

Check version
pingcli --version

## Features

-[x] HTTP status checking
-[ ] Multiple URL support
-[ ] Metrics endpoint
-[ ] Docker support
-[ ] Configuration files

## Development

Run tests
`go test ./...`

Build locally
`go build -o bin/pingcli ./cmd/pingcli`
