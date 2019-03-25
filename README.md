# language

[![CircleCI](https://circleci.com/gh/davidsbond/language/tree/master.svg?style=shield&circle-token=d306e9788fef6101b49b0b66b356117d0da9fa69)](https://circleci.com/gh/davidsbond/language/tree/master)
[![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/davidsbond/language.svg)](https://hub.docker.com/r/davidsbond/language)
[![Coverage Status](https://coveralls.io/repos/github/davidsbond/language/badge.svg?branch=master)](https://coveralls.io/github/davidsbond/language?branch=master)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://godoc.org/github.com/davidsbond/language?status.svg)](http://godoc.org/github.com/davidsbond/language)
[![Go Report Card](https://goreportcard.com/badge/github.com/davidsbond/language)](https://goreportcard.com/report/github.com/davidsbond/language)

## Getting started

This section assumes you have [go](https://golang.org/) 1.11+

```bash
# Download the source
go get github.com/davidsbond/language
# Build the binary
go install github.com/davidsbond/language

# Start the repl
./language

> 1 + 1
2
>
```

## Project structure

```text
.
├── ast             # Types representing the abstract syntax tree
├── builtin         # Built-in methods
├── evaluator       # Methods for evaluating expressions
├── lexer           # Lexical analysis methods
├── object          # Object definitions
├── parser          # The token parser
└── token           # Token definitions
```