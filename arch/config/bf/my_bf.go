/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package bf

// MyBfT is a business function type
type MyBfT = func(i int) int

// MyBfCfg is the config data type for MyBfT
type MyBfCfg struct {
	K int
}

// MyBfCfgProvider is the type of functions that provide
// the required config data for MyBfT.
type MyBfCfgProvider func() MyBfCfg

// MyBfC is the higher-order function that constructs instances of MyBfT.
func MyBfC(
	cfgProvider MyBfCfgProvider,
) MyBfT {
	return func(i int) int {
		cfg := cfgProvider()
		return i + cfg.K
	}
}
