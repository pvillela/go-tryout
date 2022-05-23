/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package main

import (
	"fmt"
	"github.com/pvillela/go-tryout/arch/errx/errx4"
)

var (
	ErrXxx = errx4.NewKind("ErrXxx")
	ErrYyy = errx4.NewKind("ErrYyy", ErrXxx, ErrXxx)
	ErrZzz = errx4.NewKind("ErrZzz", ErrYyy)
	ErrWww = errx4.NewKind("ErrWww", ErrYyy, ErrZzz)
)

var bazErr errx4.Errx

func baz() errx4.Errx {
	bazErr = ErrXxx.Make(nil, "baz %v", "BAZ")
	return bazErr
}

func bar() errx4.Errx {
	err := baz()
	err = ErrYyy.Make(err, "bar %v", "BAR")
	return err
}

func foo() errx4.Errx {
	err := bar()
	err = ErrZzz.Decorate(err, "foo %v", "FOO")
	return err
}

type errW error

func main() {
	{
		err := fmt.Errorf("Error zero")
		errX := errx4.NewErrx(err, "new error %v cause", "with")
		fmt.Println(errX)
		fmt.Println(errX.StackTrace())
	}
	{
		errX := errx4.NewErrx(nil, "new error %v cause", "without")
		fmt.Println(errX)
		fmt.Println(errX.StackTrace())
	}
	{
		fmt.Println("ErxOf error")
		err := fmt.Errorf("Error zero")
		errX := errx4.ErrxOf(err)
		fmt.Println(errX)
		fmt.Println(errX.StackTrace())
	}
	{
		fmt.Println("ErxOf string")
		errX := errx4.ErrxOf("a string")
		fmt.Println(errX)
		fmt.Println(errX.StackTrace())
	}

	fmt.Println(ErrXxx.Make(nil, ""))
	fmt.Println(ErrXxx.Make(nil, "foo"))
	fmt.Println(ErrXxx.Make(nil, "foo%v"))
	fmt.Println(ErrXxx.Make(nil, "foo%v", "bar"))

	fooErr := foo()

	fmt.Println("\n---errW(fooErr)--------------------------------------------")
	fmt.Printf("%+v\n", errW(fooErr))

	fmt.Println("\n---fmt.Println(fooErr)----------------------------------------------")
	fmt.Println(fooErr)

	fmt.Println("\n---fmt.Println(error(fooErr)--------------------------------------------")
	fmt.Println(error(fooErr))

	fmt.Println("\n---error(fooErr)--------------------------------------------")
	fmt.Printf("%+v\n", error(fooErr))

	fmt.Println("\n---fooErr.StackTrace()--------------------------------------------")
	fmt.Println(fooErr.StackTrace())

	fmt.Println("\n---errx4.StackTrace(fooErr)--------------------------------------------")
	fmt.Println(errx4.StackTraceOf(fooErr))

	fmt.Println("\n---errx4.StackTraceOf(error(fooErr))--------------------------------------------")
	fmt.Printf("%+v\n", errx4.StackTraceOf(error(fooErr)))

	fmt.Println("\n---fmt.Println(fooErr.Cause()--------------------------------------------")
	fmt.Println(fooErr.Unwrap())

	fmt.Println("\n---fmt.Println(fooErr.InnermostErrx()--------------------------------------------")
	fmt.Println(fooErr.InnermostErrx())

	fmt.Println("\n---fmt.Println(fooErr.InnermostCause()--------------------------------------------")
	fmt.Println(fooErr.InnermostCause())

	fmt.Println("\n---fooErr.CauseChain()----------------------------------------------")
	for _, err := range fooErr.CauseChain() {
		fmt.Println(err)
	}

	fmt.Println("\n---fooErr.ErrxChain()----------------------------------------------")
	for _, err := range fooErr.ErrxChain() {
		fmt.Println(err)
	}

	fmt.Println("\n===bazErr=====================================================================")

	fmt.Println("\n---errW(bazErr)--------------------------------------------")
	fmt.Printf("%+v\n", errW(bazErr))

	fmt.Println("\n---fmt.Println(bazErr)----------------------------------------------")
	fmt.Println(bazErr)

	fmt.Println("\n---fmt.Println(errx4.StackTraceOf(bazErr)--------------------------------------------")
	fmt.Println(errx4.StackTraceOf(bazErr))

	fmt.Println("\n---errx4.StackTraceOf(bazErr)----------------------------------------------")
	fmt.Printf("%+v\n", errx4.StackTraceOf(bazErr))

	fmt.Println("\n===SubKinds=====================================================================")
	fmt.Println()

	deref := func(m map[*errx4.Kind]struct{}) []errx4.Kind {
		slice := make([]errx4.Kind, 0, len(m))
		for kind, _ := range m {
			slice = append(slice, *kind)
		}
		return slice
	}

	fmt.Println("ErrXxx.SuperKinds()", deref(ErrXxx.SuperKinds()))
	fmt.Println("ErrXxx.IsSubKindOf(ErrXxx)", ErrXxx.IsSubKindOf(ErrXxx))
	fmt.Println("ErrXxx.IsSubKindOf(ErrYyy)", ErrXxx.IsSubKindOf(ErrYyy))
	fmt.Println()

	fmt.Println("ErrYyy.SuperKinds()", deref(ErrYyy.SuperKinds()))
	fmt.Println("ErrYyy.IsSubKindOf(ErrXxx)", ErrYyy.IsSubKindOf(ErrXxx))
	fmt.Println("ErrYyy.IsSubKindOf(ErrZzz)", ErrYyy.IsSubKindOf(ErrZzz))
	fmt.Println()

	fmt.Println("ErrZzz.SuperKinds()", deref(ErrZzz.SuperKinds()))
	fmt.Println("ErrZzz.IsSubKindOf(ErrXxx)", ErrZzz.IsSubKindOf(ErrXxx))
	fmt.Println("ErrZzz.IsSubKindOf(ErrYyy)", ErrZzz.IsSubKindOf(ErrYyy))
	fmt.Println("ErrZzz.IsSubKindOf(ErrZzz)", ErrZzz.IsSubKindOf(ErrZzz))
	fmt.Println("ErrZzz.IsSubKindOf(ErrWww)", ErrZzz.IsSubKindOf(ErrWww))
	fmt.Println()

	fmt.Println("ErrWww.SuperKinds()", deref(ErrWww.SuperKinds()))
	fmt.Println("ErrWww.IsSubKindOf(ErrXxx)", ErrWww.IsSubKindOf(ErrXxx))
	fmt.Println("ErrWww.IsSubKindOf(ErrYyy)", ErrWww.IsSubKindOf(ErrYyy))
	fmt.Println("ErrWww.IsSubKindOf(ErrZzz)", ErrWww.IsSubKindOf(ErrZzz))
	fmt.Println("ErrWww.IsSubKindOf(ErrWww)", ErrWww.IsSubKindOf(ErrWww))
}
