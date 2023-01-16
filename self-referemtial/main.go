/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package main

import (
	"fmt"
)

type Foo struct {
	x string
	y *string
}

func main() {
	foo1 := Foo{
		x: "foo1",
		y: nil,
	}

	foo1.y = &foo1.x

	foo2 := Foo{
		x: "foo2",
		y: nil,
	}

	foo2.y = &foo2.x

	fmt.Println("foo1.x:", foo1.x, ", foo1.y:", foo1.y, ", *foo1.y:", *foo1.y)
	fmt.Println("foo2.x:", foo2.x, ", foo2.y:", foo2.y, ", *foo2.y:", *foo2.y)

	foo1, foo2 = foo2, foo1

	fmt.Println("foo1.x:", foo1.x, ", foo1.y:", foo1.y, ", *foo1.y:", *foo1.y)
	fmt.Println("foo2.x:", foo2.x, ", foo2.y:", foo2.y, ", *foo2.y:", *foo2.y)
}
