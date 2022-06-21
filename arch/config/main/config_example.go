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
	"github.com/pvillela/go-tryout/arch/config/config"
)

func testDynamic() {

	// Config provider for testing bf.MyBf.
	testMyBfCfgPvdr := func() bf.DynamicBfCfg {
		return bf.DynamicBfCfg{K: 3}
	}

	// Instance of business function for test.
	dynamicBf := bf.DynamicBfC(testMyBfCfgPvdr)

	res := dynamicBf(2)
	fmt.Println("Test result for DynamicBf:", res)
}

func prodDynamic() {
	// Instance of business function for production.
	dynamicBf := bfc.DynamicBf

	res := dynamicBf(2)
	fmt.Println("Prod result for DynamicBf:", res)
}

func prodStatic() {
	// Instance of business function for production.
	staticBf := bfc.StaticBf

	res := staticBf(2)
	fmt.Println("Prod result for StaticBf:", res)
}

func main() {
	testDynamic()

	fmt.Println("Initial value of config.GlobalCfg.Y:", config.GlobalCfg.Y)
	prodDynamic()
	prodStatic()

	// Update global config properties.
	config.GlobalCfg.Y = 99
	fmt.Println("Changed config.GlobalCfg.Y to", config.GlobalCfg.Y)

	prodDynamic()
	prodStatic()
}
