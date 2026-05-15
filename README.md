# golang-learn

Personal Go learning project. Each topic lives in its own subdirectory under `files/`.

## Environment Setup

**Install Go**
```bash
brew install go
```

**Install gopls (Go LSP)**
```bash
go install golang.org/x/tools/gopls@latest
```

**Install golangci-lint**
```bash
brew install golangci-lint
```

**Verify installations**
```bash
go version
gopls version
golangci-lint --version
```

## Usage

```bash
make test    # run all tests
make fmt     # format code
make lint    # run linter

go run files/<topic>/main.go  # run a specific example
```
