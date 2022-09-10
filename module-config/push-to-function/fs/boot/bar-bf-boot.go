/*
 * Copyright © 2022 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the MIT license
 * that can be found in the LICENSE file.
 */

package boot

import (
	"github.com/pvillela/go-tryout/module-config/push-to-function/fs"
	"github.com/pvillela/go-tryout/module-config/push-to-var/config"
)

func barBfCfgAdapter(appCfg config.AppCfgInfo) fs.BarBfCfgInfo {
	return fs.BarBfCfgInfo{
		Z: appCfg.Y,
	}
}

func BarBfBoot(appCfg func() config.AppCfgInfo) fs.BarBfT {
	return fs.BarBfC(fs.BarBfCfgSrc{
		Get: func() fs.BarBfCfgInfo { return barBfCfgAdapter(appCfg()) },
	})
}
