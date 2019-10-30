# base64url
Command-line base64url encoder / decoder

# Installation

```bash
$ go get -u github.com/oscarayoy/base64url
```

# Usage

```bash
$ base64url -h
Usage of base64url:
  -d	decodes input (default is to encode)
  -i string
    	input file or "-" for stdin (default "-")
  -o string
    	output file or "-" for stdout (default "-")
```

Encode from stdin:

```bash
$ echo "Hello, world" | base64url
SGVsbG8sIHdvcmxkCg
```

Decode from stdin:

```bash
$ echo "SGVsbG8sIHdvcmxkCg" | base64url -d
Hello, world
```

Encode to and from file:

```bash
$ base64url -i plain.txt -o b64.txt
```

Decode to and from file:

```bash
$ base64url -d -i b64.txt -o plain.txt
```

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
