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

// staticBfCfgProdProvider is the config provider for production.
var staticBfCfgProdProvider bf.StaticBfCfgPvdr = func() bf.StaticBfCfg {
	gCfg := config.GlobalCfgPvdr()
	return bf.StaticBfCfg{K: gCfg.Y}
}

// StaticBf is the configured stereotype instance.
var StaticBf bf.StaticBfT = bf.StaticBfC(staticBfCfgProdProvider)
