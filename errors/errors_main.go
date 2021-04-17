/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package main

// Example based on https://pkg.go.dev/github.com/pkg/errors#example-Cause and
// https://dave.cheney.net/2016/06/12/stack-traces-and-the-errors-package

import (
	"fmt"

	"github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func fn0() error {
	e1 := errors.New("original")
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

	fmt.Println("\n\n********** with stackTracer")
	fmt.Printf("%+v", err.(stackTracer).StackTrace())
}
