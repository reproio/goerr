package main

import (
	"errors"
	"github.com/reproio/goerr"
	"go.uber.org/zap"
	"log"
)

type Nested struct {
	NestedData string
}
type MyStruct struct {
	Num    int
	Str    string
	Nested *Nested
}

func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	st := MyStruct{
		Num: 1,
		Str: "str",
		Nested: &Nested{
			NestedData: "nested",
		},
	}

	err := goerr.Wrap(errors.New("some error")).
		WithValue("key", "value").
		WithValue("num_key", 1).
		WithValue("struct", st)

	if goErr := goerr.Unwrap(err); goErr != nil {
		log.Printf("%s\n", goErr.LogValue())
	}

	logger.Error("something happening", zap.Object("error", err))

	if goErr := goerr.Unwrap(err); goErr != nil {
		err2 := goErr.Clone()
		err2 = err2.WithValue("___2", "This is err2")
		logger.Error("something happening2", zap.Object("error2", err2))
	}
}
