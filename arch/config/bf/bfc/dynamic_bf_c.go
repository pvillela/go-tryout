/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package bfc

import (
	"github.com/pvillela/go-tryout/arch/config/bf"
	"github.com/pvillela/go-tryout/arch/config/config"
)

// dynamicBfCfgProdProvider is the config provider for production.
var dynamicBfCfgProdProvider bf.DynamicBfCfgPvdr = func() bf.DynamicBfCfg {
	gCfg := config.GlobalCfgPvdr()
	return bf.DynamicBfCfg{K: gCfg.Y}
}

// DynamicBf is the configured stereotype instance.
var DynamicBf bf.DynamicBfT = bf.DynamicBfC(dynamicBfCfgProdProvider)
