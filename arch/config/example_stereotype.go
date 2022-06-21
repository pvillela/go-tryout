/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package config

import (
	"github.com/pvillela/go-tryout/arch/errx"
	"strconv"
)

// MyBfT is a business function type
type MyBfT = func(i int) int

// MyBfCfg is the config data structure for MyBfT
type MyBfCfg struct {
	K int
}

// MyBfCfgExtr is the type of functions that extract from a ConfigProvider
// the required config structure for MyBfT.
type MyBfCfgExtr[C any] func(c C) MyBfCfg

// MyBfC is the higher-order function that constructs instances of MyBfT.
func MyBfC[C any](
	cfgProvider ConfigProvider[C],
	cfgExtractor MyBfCfgExtr[C],
) MyBfT {
	return func(i int) int {
		cfg := cfgExtractor(cfgProvider())
		return i + cfg.K
	}
}

// myBfCfgProdExtr is the config extractor for ProdConfigProvider.
var myBfCfgProdExtr MyBfCfgExtr[string] = func(c string) MyBfCfg {
	i, err := strconv.Atoi(c)
	errx.PanicOnError(err)
	return MyBfCfg{K: i}
}

// MyBfI is the configured stereotype instance.
var MyBfI MyBfT = MyBfC(ProdConfigProvider, myBfCfgProdExtr)
