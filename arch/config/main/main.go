/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package main

import (
	"fmt"
	"github.com/pvillela/go-tryout/arch/config"
)

func test() {

	// Config provider for testing, different from config.ConfigProvider.
	testCfgProvider := func() int {
		return 3
	}

	// Config extractor to go along with the above test config provider.
	// Different from the production config extractor.
	var myBfCfgExtrForTest config.MyBfCfgExtr[int] = func(c int) config.MyBfCfg {
		return config.MyBfCfg{K: c}
	}

	// Instance of business function for test.
	myBf := config.MyBfC(
		testCfgProvider,
		myBfCfgExtrForTest,
	)

	res := myBf(2)
	fmt.Println("Test result:", res)
}

func prod() {
	// Instance of business function for production.
	myBf := config.MyBfI

	res := myBf(2)
	fmt.Println("Prod result:", res)
}

func main() {
	test()
	prod()
}
