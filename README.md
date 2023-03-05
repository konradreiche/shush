# 🤫 shush
[![Go Reference](https://pkg.go.dev/badge/github.com/konradreiche/shush.svg)](https://pkg.go.dev/github.com/konradreiche/shush) [![actions](https://github.com/konradreiche/shush/actions/workflows/test.yaml/badge.svg)](https://github.com/konradreiche/shush/actions) [![codecov](https://codecov.io/gh/konradreiche/shush/branch/main/graph/badge.svg?token=VIY0XN5FF0)](https://codecov.io/gh/konradreiche/shush)

A Go linter to report `fmt.Println` which may have been unintentionally left in the code for debugging purposes.

## Install

```
go install github.com/konradreiche/shush/cmd/shush
```

## Usage

```
shush ./...
```

You can also run it as part of `go vet`.

```
go vet -vettool=$(which shush)
```
