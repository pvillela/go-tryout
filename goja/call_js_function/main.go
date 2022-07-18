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
function f(param) {
    return +param + 2;
}
`

	vm := goja.New()
	_, err := vm.RunString(SCRIPT)
	if err != nil {
		panic(err)
	}

	var fn func(string) string
	err = vm.ExportTo(vm.Get("f"), &fn)
	if err != nil {
		panic(err)
	}

	fmt.Println(fn("40")) // note, _this_ value in the function will be undefined.
	// Output: 42
}
