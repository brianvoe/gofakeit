/*
Package gofakeit is a random data generator written in go

Every function has an example and a benchmark

See the full list here https://godoc.org/github.com/brianvoe/gofakeit

Examples:
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
  // 80+ more!!!
*/
package gofakeit
