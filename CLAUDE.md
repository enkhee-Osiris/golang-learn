# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Purpose

Personal Go learning project. Code here is meant to be readable and instructive — prioritize clarity over cleverness.

## Commands

```bash
make test       # run all tests
make fmt        # format all files with gofmt
make lint       # run golangci-lint (install via: brew install golangci-lint)

# Run a specific example
go run files/hello-world/main.go

# Run a single test
go test ./... -run TestFunctionName

# Build
go build ./...
```

## Structure

Each concept or topic lives in its own subdirectory under `files/` (e.g. `files/hello-world/main.go`). Each subdirectory is its own Go package — this avoids `main redeclared` errors that occur when multiple `package main` files share the same directory.

```
files/
  hello-world/main.go
  values/main.go
```

Run a specific example with `go run files/<topic>/main.go`.

## Style

- Keep each file focused on one concept
- Prefer explicit over implicit — this is a learning repo, so verbose is better than terse
- Add a short comment at the top of each new file explaining what concept it demonstrates
