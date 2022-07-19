/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

// See https://github.com/dop251/goja#calling-js-functions-from-go

package main

import (
	"fmt"
	"github.com/dop251/goja"
)

func main() {
	const SCRIPT = `
function sum(a, b) {
    return a+b;
}
`

	vm := goja.New()
	_, err := vm.RunString(SCRIPT)
	if err != nil {
		panic(err)
	}

	v := vm.Get("sum")

	// Using AssertFunction
	{
		sum, ok := goja.AssertFunction(v)
		if !ok {
			panic("Not a function")
		}

		res, err := sum(goja.Undefined(), vm.ToValue(40), vm.ToValue(2))
		if err != nil {
			panic(err)
		}

		fmt.Println(res)
		// Output: 42
	}

	// Using ExportTo
	{
		var fn func(int, int) int
		err = vm.ExportTo(v, &fn)
		if err != nil {
			panic(err)
		}

		res := fn(40, 2)

		fmt.Println(res) // note, _this_ value in the function will be undefined.
		// Output: 42
	}

	// ExportTo is not type-safe
	{
		var fn func(struct{ a int }, int) int
		err = vm.ExportTo(v, &fn)
		if err != nil {
			panic(err)
		}

		res := fn(struct{ a int }{40}, 2)

		fmt.Println(res)
	}
}
