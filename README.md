# gofakeit [![Go Report Card](https://goreportcard.com/badge/github.com/brianvoe/gofakeit)](https://goreportcard.com/report/github.com/brianvoe/gofakeit) [![Build Status](https://travis-ci.org/brianvoe/gofakeit.svg?branch=master)](https://travis-ci.org/brianvoe/gofakeit) [![codecov.io](https://codecov.io/github/brianvoe/gofakeit/branch/master/graph/badge.svg)](https://codecov.io/github/brianvoe/gofakeit) [![GoDoc](https://godoc.org/github.com/brianvoe/gofakeit?status.svg)](https://godoc.org/github.com/brianvoe/gofakeit) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/icrowley/fake/master/LICENSE)
Random data generator written in go

Every function has an example and a benchmark

### 80+ Functions!!!
If there is something that is generic enough missing from this package [add an issue](https://github.com/brianvoe/gofakeit/issues) and let me know what you need.
Most of the time i'll add it!

## Documentation
[![GoDoc](https://godoc.org/github.com/brianvoe/gofakeit?status.svg)](https://godoc.org/github.com/brianvoe/gofakeit)

## Example
```go
import "github.com/brianvoe/gofakeit"

gofakeit.Name() // Markus Moen
gofakeit.Email() // alaynawuckert@kozey.biz
gofakeit.Phone() // (570)245-7485
gofakeit.BS() // front-end
gofakeit.BeerName() // Duvel
gofakeit.Color() // MediumOrchid
gofakeit.Company() // Moen, Pagac and Wuckert
gofakeit.CreditCardNumber() // 4287271570245748
gofakeit.HackerPhrase() // Connecting the array won't do anything, we need to generate the haptic COM driver!
gofakeit.JobTitle() // Director
gofakeit.Password(true, true, true, true, true, 32) // WV10MzLxq2DX79w1omH97_0ga59j8!kj
gofakeit.CurrencyShort() // USD
// 80+ more!!!
```
