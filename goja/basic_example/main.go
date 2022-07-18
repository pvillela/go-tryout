/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

// See https://github.com/dop251/goja#basic-example

package main

import (
	"fmt"
	"github.com/dop251/goja"
)

func main() {
	vm := goja.New()
	v, err := vm.RunString("40 + 2")
	if err != nil {
		panic(err)
	}
	num := v.Export().(int64)
	fmt.Println(num)
}
