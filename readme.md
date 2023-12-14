# SSMinator

[![Go Report Card](https://goreportcard.com/badge/github.com/scharissis/ssminator)](https://goreportcard.com/report/github.com/scharissis/ssminator)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/scharissis/ssminator?tab=doc)


Makes using AWS SSM simpler, with a focus on easily running commands across your entire fleet.

`Warning: This is very early in development; I do not recommend using this package as of yet.`

## Features
- Simpler return structs (without pointers)

## Features Under Consideration
- Converting async calls to sync, or providing sync variants

## CLI API
```
cmd
 - run
 - check
```

## TODO
- Decide on high-level cli api
- Decide on high-level pkg api
- Document missing methods we wish to expose
- Tests
