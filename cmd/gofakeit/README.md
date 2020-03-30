# gofakeit

This is a version of [gofakeit](../../README.md) that can be used directly from a CLI.

All functions except for these that require additional parameters and do not return a response as a single string are supported.

### Install

```shell script
$ go install https://github.com/brianvoe/gofakeit/cmd/gofakeit
```

### Usage

```shell script
$ gofakeit -h
Random fake data generator

Usage:
  gofakeit [command]

Available Commands:
  animal                  Generate random Animal
  animaltype              Generate random AnimalType
  beeralcohol             Generate random BeerAlcohol
...
``` 
