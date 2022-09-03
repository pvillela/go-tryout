/*
 * Copyright © 2022 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the MIT license
 * that can be found in the LICENSE file.
 */

package mod

import "github.com/pvillela/go-tryout/module-config/push-to-file/fwk"

func barBfCfgAdapter(appCfg fwk.AppCfgInfo) BarBfCfgInfo {
	return BarBfCfgInfo{
		Z: appCfg.Y,
	}
}

func BarBfBoot(appCfg func() fwk.AppCfgInfo) BarBfT {
	return BarBfC(BarBfCfgSrc{
		Get: func() BarBfCfgInfo { return barBfCfgAdapter(appCfg()) },
	})
}
