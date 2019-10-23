# base64url
Command-line base64url encoder / decoder

# Development Setup

## Install Goimports

This is used for code formatting.

```bash
$ go get -u golang.org/x/tools/cmd/goimports
```

## Install Golint

This is used for code linting.

```bash
$ go get -u golang.org/x/lint/golint
```

## Enable Commit Signing

The repository has been configured to enforce signature verification.

```bash
$ git config commit.gpgsign true
$ git config alias.tag 'tag -s'
```

## Set Up the Git Pre-Commit Hook

This pre-commit hook makes sure all code is properly formatted and linted before it is committed.

```bash
$ cp githooks/pre-commit .git/hooks/pre-commit
```

Install `goimports`, `golint` and `go vet` support in your editor of choice as needed.
