# Gofakeit Command Line Tool
Run gofakeit via that command line. 
All functions are available to run in lowercase and if they require additional parameters you may pass by space seperation.

### Installation
```go
go get -u github.com/brianvoe/gofakeit/v6/cmd/gofakeit
```

### Example
```bash
gofakeit firstname // billy
gofakeit sentence 5 // Exercitationem occaecati sed dignissimos inventore.
gofakeit shufflestrings hello,world,whats,up // up,hello,whats,world
```

### List of available functions
```bash
gofakeit list
```

![](https://raw.githubusercontent.com/brianvoe/gofakeit/master/cmd/gofakeit/cmd.gif)