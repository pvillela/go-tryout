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

var BarBfAdapterCfgSrc = fwk.MakeConfigSource[fwk.AppCfgInfo]()

var _ = (func() struct{} {
	BarBfCfgSrc.Set(func() BarBfCfgInfo {
		return barBfCfgAdapter(BarBfAdapterCfgSrc.Get())
	})
	return struct{}{}
})()
