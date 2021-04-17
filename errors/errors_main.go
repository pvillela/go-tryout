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

	// This initial set of examples shows that Unwrap is buggy

	u3a := errors.Unwrap(err)
	u3b := errors.Unwrap(u3a)
	u2a := errors.Unwrap(u3b)
	u2b := errors.Unwrap(u2a)
	u1a := errors.Unwrap(u2b)
	u1b := errors.Unwrap(u1a)
	u0a := errors.Unwrap(u1b)
	u0b := errors.Unwrap(u0a)

	fmt.Println("\n*** err ->", err)
	fmt.Println("\n*** u3a ->", u3a)
	fmt.Println("\n*** u3b ->", u3b)
	fmt.Println("\n*** u2a ->", u2a)
	fmt.Println("\n*** u2b ->", u2b)
	fmt.Println("\n*** u1a ->", u1a)
	fmt.Println("\n*** u1b ->", u1b)
	fmt.Println("\n*** u0a ->", u0a)
	fmt.Println("\n*** u0b ->", u0b)

	fmt.Printf("\n##### err -> \n%+v\n", err)
	fmt.Printf("\n##### u3a -> \n%+v\n", u3a)
	fmt.Printf("\n##### u3b -> \n%+v\n", u3b)
	fmt.Printf("\n##### u2a -> \n%+v\n", u2a)
	fmt.Printf("\n##### u2b -> \n%+v\n", u2b)
	fmt.Printf("\n##### u1a -> \n%+v\n", u1a)
	fmt.Printf("\n##### u1b -> \n%+v\n", u1b)
	fmt.Printf("\n##### u0a -> \n%+v\n", u0a)
	fmt.Printf("\n##### u0b -> \n%+v\n", u0b)

	// These examples show behaviour according to the spec.

	fmt.Println("\n*** errors.Cause(err) ->", errors.Cause(err))
	fmt.Println("\n**** stack trace below")
	fmt.Printf("%+v", err)

	// This shows that when the StackTrace method is used then only the most recent outer trace
	// is output.

	fmt.Println("\n\n********** with stackTracer")
	fmt.Printf("%+v", err.(stackTracer).StackTrace())
}
