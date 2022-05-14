/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package main

import "fmt"

type fooT interface {
	bar() int
}

var err error
var foo fooT
var slice []int = nil // assignment just for emphasis as the default (zero) is nil
var pAny *any
var pErr *error
var pFoo *fooT

func isNilAny(x any) bool {
	if x == nil {
		return true
	}
	return false
}

// A nil interface is a nil any, but a nil slice and a nil pointer are not.
func main() {
	fmt.Println("isNilAny(err):", isNilAny(err))
	fmt.Println("isNilAny(foo):", isNilAny(foo))
	fmt.Println("isNilAny(slice):", isNilAny(slice))
	fmt.Println("isNilAny(pAny):", isNilAny(pAny))
	fmt.Println("isNilAny(pErr):", isNilAny(pErr))
	fmt.Println("isNilAny(pFoo):", isNilAny(pFoo))
}
