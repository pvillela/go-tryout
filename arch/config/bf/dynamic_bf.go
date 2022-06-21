/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package bf

// DynamicBfT is a business function type
type DynamicBfT = func(i int) int

// DynamicBfCfg is the config data type for DynamicBfT
type DynamicBfCfg struct {
	K int
}

// DynamicBfCfgPvdr is the type of functions that provide
// the required config data for DynamicBfT.
type DynamicBfCfgPvdr func() DynamicBfCfg

// DynamicBfC is the higher-order function that constructs instances of DynamicBfT.
// The constructed function gets fresh configuration data whenever the config
// properties are updated.
func DynamicBfC(
	cfgPvdr DynamicBfCfgPvdr,
) DynamicBfT {
	return func(i int) int {
		cfg := cfgPvdr()
		return i + cfg.K
	}
}
