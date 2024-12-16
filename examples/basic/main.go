package main

import (
	"errors"
	"log"
	"time"

	"github.com/reproio/goerr"
)

func someAction(input string) error {
	if input != "OK" {
		return goerr.New("input is not OK").WithValue("input", input).WithValue("time", time.Now())
	}
	return nil
}

func main() {
	if err := someAction("ng"); err != nil {
		var goErr *goerr.Error
		if errors.As(err, &goErr) {
			for k, v := range goErr.Values() {
				log.Printf("%s = %v\n", k, v)
			}
		}
		log.Fatalf("Error: %+v\n", err)
	}
}
