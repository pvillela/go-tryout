/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package main

import (
	"fmt"
	"github.com/pvillela/go-tryout/arch/config/bf"
	"github.com/pvillela/go-tryout/arch/config/bf/bfc"
	"github.com/pvillela/go-tryout/arch/errx"
	"os"
)

func test() {

	// Config provider for testing bf.MyBf.
	testMyBfCfgProvider := func() bf.MyBfCfg {
		return bf.MyBfCfg{K: 3}
	}

	// Instance of business function for test.
	myBf := bf.MyBfC(testMyBfCfgProvider)

	res := myBf(2)
	fmt.Println("Test result:", res)
}

func prod() {
	// Instance of business function for production.
	myBf := bfc.MyBf

	res := myBf(2)
	fmt.Println("Prod result:", res)
}

func main() {
	test()

	// Set environment for prod config properties.
	err := os.Setenv("X", "xyz")
	errx.PanicOnError(err)
	err = os.Setenv("Y", "9")
	errx.PanicOnError(err)

	prod()
}
