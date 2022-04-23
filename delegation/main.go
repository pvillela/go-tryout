/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package main

import "fmt"

type Foo struct {
	x int
	y string
}

func (s Foo) PrintX() {
	fmt.Println("x =", s.x)
}

func FooPrintX(foo Foo) {
	fmt.Println("foo.x =", foo.x)
}

type Bar struct {
	Foo
	z int
}

type BarG[T any] struct {
	Entity T
	z      int
}

func main() {
	bar := Bar{
		Foo: Foo{9, "bar"},
		z:   42,
	}
	bar.PrintX()
	// FooPrintX(bar) // compilation error

	barg := BarG[Foo]{
		Entity: Foo{9, "bar"},
		z:      42,
	}

	// barg.PrintX() // compilation error
	barg.Entity.PrintX()
}
