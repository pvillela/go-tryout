/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package main

import "fmt"

func zero[T any]() T {
	var zero T
	return zero
}

type barT struct {
	x int
	y string
}

func main() {
	i := zero[int]()
	fmt.Println(i)

	bar := zero[barT]()
	fmt.Println(bar)
}
