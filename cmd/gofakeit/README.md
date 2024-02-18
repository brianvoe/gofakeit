# Gofakeit Command Line Tool

Run gofakeit via that command line.
All functions are available to run in lowercase and if they require additional parameters you may pass by space seperation.

### Installation

```go
go install -v github.com/brianvoe/gofakeit/v7/cmd/gofakeit@latest
```

### Example

```bash
gofakeit firstname // billy
gofakeit sentence 5 // Exercitationem occaecati sed dignissimos inventore.
gofakeit shufflestrings hello,world,whats,up // up,hello,whats,world

// You can also pass in a length flag to output a specific function multiple times
gofakeit sentence -loop=5
```

### List of available functions

```bash
gofakeit list

// You may also pass a category and/or function to get a specific function data
gofakeit list [category] [function]
```

![](https://raw.githubusercontent.com/brianvoe/gofakeit/master/cmd/gofakeit/cmd.gif)
