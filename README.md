# 🤫 shush
![golangci-lint](https://github.com/konradreiche/shush/actions/workflows/test.yaml/badge.svg) [![codecov](https://codecov.io/gh/konradreiche/shush/branch/main/graph/badge.svg?token=VIY0XN5FF0)](https://codecov.io/gh/konradreiche/shush)

A Go linter to report `fmt.Println` which may have been unintentionally left in the code for debugging purposes.

## Usage

```
go install github.com/konradreiche/shush
shush ./...
```
