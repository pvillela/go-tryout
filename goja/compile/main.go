/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

// See https://prasanthmj.github.io/go/javascript-parser-in-go/

package main

import (
	"fmt"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/require"
)

func SimpleJS() {
	vm := goja.New()

	new(require.Registry).Enable(vm)
	console.Enable(vm)

	script := `
		console.log("Hello world - from Javascript inside Go! ")
		function foo(x, y) {
			return x + y
		}
		foo(40, 2)
	`
	fmt.Println("Compiling ... ")
	prog, err := goja.Compile("", script, true)
	if err != nil {
		fmt.Printf("Error compiling the script %v ", err)
		return
	}
	fmt.Println("Running ... \n ")
	v, err := vm.RunProgram(prog)
	fmt.Println("v:", v)

	// Get foo from vm as Go function and invoke it
	{
		var foo func(int, int) int
		err = vm.ExportTo(vm.Get("foo"), &foo)
		if err != nil {
			fmt.Printf("Error exporting the function %v", err)
			return
		}

		res := foo(40, 2)
		fmt.Println("res:", res)
	}
}

func main() {
	SimpleJS()
}
