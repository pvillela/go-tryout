/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package config

// Config is an example of a global configuration data structure.
type Config struct {
	X string
	Y int
}

var GlobalCfg = Config{
	X: "XYZ",
	Y: 9,
}

// GlobalCfgPvdr is an example of a global configuration provider that sources
// configuration properties from the environment.
func GlobalCfgPvdr() Config {
	return GlobalCfg
}
