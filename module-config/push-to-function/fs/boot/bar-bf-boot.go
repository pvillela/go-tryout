/*
 * Copyright © 2022 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the MIT license
 * that can be found in the LICENSE file.
 */

package boot

import (
	"github.com/pvillela/go-tryout/module-config/push-to-function/config"
	"github.com/pvillela/go-tryout/module-config/push-to-function/fs"
)

var BarBfCfgAdapter = func(appCfgSrc config.AppCfgSrc) func() fs.BarBfCfgInfo {
	return func() fs.BarBfCfgInfo {
		return fs.BarBfCfgInfo{
			Z: appCfgSrc().Y,
		}
	}
}

func BarBfBoot(appCfg config.AppCfgSrc) fs.BarBfT {
	return fs.BarBfC(fs.BarBfCfgSrc{
		Get: BarBfCfgAdapter(appCfg),
	})
}
