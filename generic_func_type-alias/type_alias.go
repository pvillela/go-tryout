/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package main

import (
	"fmt"
	"strconv"
)

type FunT[S any, T any] func(in S) T

// Can be used as type FunT[int, string]
func foo(x int) string {
	return strconv.Itoa(x)
}

func apply[S, T any](f FunT[S, T], s S) T {
	return f(s)
}

func main() {
	str := apply(foo, 1)
	fmt.Println(str)
}
