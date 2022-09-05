/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package main

import "fmt"

type Itf interface {
	set(x int)
	get() int
}

type Foo struct {
	x int
}

func (s *Foo) set(x int) {
	s.x = x
}

func (s *Foo) get() int {
	return s.x
}

func updateItf(i Itf, x int) {
	i.set(x)
}

func main() {
	var foo Itf = &Foo{}
	fmt.Println(foo)

	updateItf(foo, 42)
	fmt.Println(foo)
}
