package main

// Example based on https://pkg.go.dev/github.com/pkg/errors#example-Cause and
// https://dave.cheney.net/2016/06/12/stack-traces-and-the-errors-package

import (
	"fmt"

	"github.com/pkg/errors"
)

func fn0() error {
	e1 := errors.New("error")
	e2 := errors.Wrap(e1, "inner")
	e3 := errors.Wrap(e2, "middle")
	return errors.Wrap(e3, "outer")
}

func fn() error { return fn0() }

func main() {
	err := fn()
	fmt.Println("\n*** err ->", err)
	fmt.Println("\n*** errors.Cause(err) ->", errors.Cause(err))
	fmt.Println("\n**** stack trace below")
	fmt.Printf("%+v", err)
}
