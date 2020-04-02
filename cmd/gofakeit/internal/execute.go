package internal

import (
	"time"

	"github.com/brianvoe/gofakeit/v4"
)

func init() {
	gofakeit.Seed(time.Now().UnixNano())
}

//go:generate go run ./namexec -o execute_gen.go -p internal ./../../..
