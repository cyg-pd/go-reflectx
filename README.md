# go-reflectx

[![tag](https://img.shields.io/github/tag/cyg-pd/go-reflectx.svg)](https://github.com/cyg-pd/go-reflectx/releases)
![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.18-%23007d9c)
[![GoDoc](https://godoc.org/github.com/cyg-pd/go-reflectx?status.svg)](https://pkg.go.dev/github.com/cyg-pd/go-reflectx)
![Build Status](https://github.com/cyg-pd/go-reflectx/actions/workflows/test.yml/badge.svg)
[![Go report](https://goreportcard.com/badge/github.com/cyg-pd/go-reflectx)](https://goreportcard.com/report/github.com/cyg-pd/go-reflectx)
[![Coverage](https://img.shields.io/codecov/c/github/cyg-pd/go-reflectx)](https://codecov.io/gh/cyg-pd/go-reflectx)
[![Contributors](https://img.shields.io/github/contributors/cyg-pd/go-reflectx)](https://github.com/cyg-pd/go-reflectx/graphs/contributors)
[![License](https://img.shields.io/github/license/cyg-pd/go-reflectx)](./LICENSE)

## 🚀 Install

```sh
go get github.com/cyg-pd/go-reflectx@v1
```

This library is v1 and follows SemVer strictly.

No breaking changes will be made to exported APIs before v2.0.0.

This library has no dependencies outside the Go standard library.

## 💡 Usage

You can import `go-reflectx` using:

```go
package main

import (
	"log/slog"

	"github.com/cyg-pd/go-reflectx"
)

type SomeStruct struct{}

func main() {
	name := reflectx.FullName(SomeStruct{})
	slog.Info(name) // main.SomeStruct
}
```
