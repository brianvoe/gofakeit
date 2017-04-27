# gofakeit [![Go Report Card](https://goreportcard.com/badge/github.com/brianvoe/gofakeit)](https://goreportcard.com/report/github.com/brianvoe/gofakeit) [![Build Status](https://travis-ci.org/brianvoe/gofakeit.svg?branch=master)](https://travis-ci.org/brianvoe/gofakeit)
Random data generator written in go

## Documentation
[![GoDoc](https://godoc.org/github.com/brianvoe/gofakeit?status.svg)](https://godoc.org/github.com/brianvoe/gofakeit)

## Example
```go
import "github.com/brianvoe/gofakeit"

gofakeit.Name() // Markus Moen
gofakeit.Email() // alaynawuckert@kozey.biz
gofakeit.Phone() // (570)245-7485
gofakeit.Address().Address // 75776 Lake Viewland, Sterlingstad, New Hampshire 82250-2868
```
