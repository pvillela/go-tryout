/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package main

import "fmt"

type Foo struct {
	x int
	y int
}

type Bar struct {
	x int
	y int
}

type IY interface {
	Py() *int
	Y() int
}

func (s Foo) Py() *int {
	return &s.y
}

func (s Foo) Y() int {
	return s.y
}

func (s *Bar) Py() *int {
	return &s.y
}

func (s *Bar) Y() int {
	return s.y
}

func main() {
	foo1 := Foo{
		x: 1,
		y: 2,
	}
	p := &foo1.y

	foo2 := Foo{
		x: 100,
		y: 200,
	}

	fmt.Println("foo1", foo1)
	fmt.Println("foo2", foo2)

	foo1 = foo2
	fmt.Println("foo1", foo1)

	*p = 42
	fmt.Println("foo1", foo1)

	// Struct poking doesn't work when the struct is passed as an interface
	foo := Foo{}
	var ifoo IY = foo
	*ifoo.Py() = 42
	fmt.Println("foo", foo)
	fmt.Println("ifoo", ifoo)
	fmt.Println("ifoo.Y()", ifoo.Y())
	fmt.Println("ifoo.Py()", ifoo.Py())
	fmt.Println("foo.Py() ", foo.Py())

	// Struct poking works when a pointer to the struct is passed as an interface
	bar := Bar{}
	var ibar IY = &bar
	*ibar.Py() = 42
	fmt.Println("bar", bar)
	fmt.Println("ibar", ibar)
	fmt.Println("ibar.Y()", ibar.Y())
	fmt.Println("ibar.Py()", ibar.Py())
	fmt.Println("bar.Py() ", bar.Py())
}
