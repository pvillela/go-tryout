/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package main

import (
	"fmt"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

func main() {
	const SCRIPT = `
	var m = require("./goja/require/testdata/m.js");
	m.test();
	`

	vm := goja.New()

	registry := new(require.Registry)
	registry.Enable(vm)

	v, err := vm.RunString(SCRIPT)
	if err != nil {
		panic(err)
	}

	if !v.StrictEquals(vm.ToValue("passed")) {
		panic(fmt.Errorf("Unexpected result: %v", v))
	}

	fmt.Println(v)

	vGo := v.Export().(string)
	fmt.Println(vGo)
}
