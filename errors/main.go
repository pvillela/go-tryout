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
	e4 := errors.WithMessage(e3, "decorated middle")
	e5 := errors.Wrap(e4, "outer")
	return e5
}

func fn() error { return fn0() }

func main() {
	err := fn()

	// This initial set of examples shows that Unwrap needs to call twice to undo Warp,
	// once to undo errors.WithMessage

	u4a := errors.Unwrap(err)
	u4 := errors.Unwrap(u4a)
	u3 := errors.Unwrap(u4)
	u2a := errors.Unwrap(u3)
	u2 := errors.Unwrap(u2a)
	u1a := errors.Unwrap(u2)
	u1 := errors.Unwrap(u1a)
	u0 := errors.Unwrap(u1)

	fmt.Println("\n*** err ->", err)
	fmt.Println("\n*** u4a ->", u4a)
	fmt.Println("\n*** u4 ->", u4)
	fmt.Println("\n*** u3 ->", u3)
	fmt.Println("\n*** u2a ->", u2a)
	fmt.Println("\n*** u2 ->", u2)
	fmt.Println("\n*** u1a ->", u1a)
	fmt.Println("\n*** u1 ->", u1)
	fmt.Println("\n*** u0 ->", u0)

	fmt.Printf("\n##### err -> \n%+v\n", err)
	fmt.Printf("\n##### u4a -> \n%+v\n", u4a)
	fmt.Printf("\n##### u4 -> \n%+v\n", u4)
	fmt.Printf("\n##### u3 -> \n%+v\n", u3)
	fmt.Printf("\n##### u2a -> \n%+v\n", u2a)
	fmt.Printf("\n##### u2 -> \n%+v\n", u2)
	fmt.Printf("\n##### u1a -> \n%+v\n", u1a)
	fmt.Printf("\n##### u1 -> \n%+v\n", u1)
	fmt.Printf("\n##### u0 -> \n%+v\n", u0)

	// These examples show behaviour according to the spec.

	fmt.Println("\n*** errors.Cause(err) ->", errors.Cause(err))
	fmt.Println("\n**** stack trace below")
	fmt.Printf("%+v", err)

	// This shows that when the StackTrace method is used then only the most recent outer trace
	// is output.

	fmt.Println("\n\n********** err via stackTracer")
	fmt.Printf("%+v", err.(stackTracer).StackTrace())

	fmt.Println("\n\n********** u4 via stackTracer")
	fmt.Printf("%+v", u4.(stackTracer).StackTrace())
}
