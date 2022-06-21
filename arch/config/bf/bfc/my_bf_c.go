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

// myBfCfgProdProvider is the config provider for production.
var myBfCfgProdProvider bf.MyBfCfgProvider = func() bf.MyBfCfg {
	gCfg := config.GlobalConfigProvider()
	return bf.MyBfCfg{K: gCfg.Y}
}

// MyBf is the configured stereotype instance.
var MyBf bf.MyBfT = bf.MyBfC(myBfCfgProdProvider)
