/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package bf

// StaticBfT is a business function type
type StaticBfT = func(i int) int

// StaticBfCfg is the config data type for StaticBfT
type StaticBfCfg struct {
	K int
}

// StaticBfCfgPvdr is the type of functions that provide
// the required config data for StaticBfT.
type StaticBfCfgPvdr func() StaticBfCfg

// StaticBfC is the higher-order function that constructs instances of StaticBfT.
// The constructed function gets fresh configuration data whenever the config
// properties are updated.
func StaticBfC(
	cfgPvdr StaticBfCfgPvdr,
) StaticBfT {
	cfg := cfgPvdr()
	return func(i int) int {
		return i + cfg.K
	}
}
