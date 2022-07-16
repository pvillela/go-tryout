/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package main

import "fmt"

func main() {
	anonymouslyTyped := struct {
		foo int
		bar string
		baz struct {
			a string
			b int
		}
	}{
		foo: 0,
		bar: "barStr",
		baz: struct {
			a string
			b int
		}{
			a: "aStr",
			b: 0,
		},
	}

	type bazT struct {
		a string
		b int
	}

	type xT struct {
		foo int
		bar string
		baz bazT
	}

	nominallyTyped := xT{
		foo: 0,
		bar: "barStr",
		baz: bazT{
			a: "aStr",
			b: 0,
		},
	}

	fmt.Println("anonymouslyTyped:", anonymouslyTyped)
	fmt.Println("nominallyTyped  :", nominallyTyped)
}
