/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package config

import (
	"github.com/pvillela/go-tryout/arch/errx"
	"os"
	"strconv"
)

// Config is an example of a global configuration data structure.
type Config struct {
	X string
	Y int
}

// GlobalCfgPvdr is an example of a global configuration provider that sources
// configuration properties from the environment.
func GlobalCfgPvdr() Config {
	yStr := os.Getenv("Y")
	y, err := strconv.Atoi(yStr)
	errx.PanicOnError(err)
	return Config{
		X: os.Getenv("X"),
		Y: y,
	}
}
