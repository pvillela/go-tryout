/*
 * Copyright © 2022 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the MIT license
 * that can be found in the LICENSE file.
 */

package mod

import "github.com/pvillela/go-tryout/module-config/pull-with-push-override/fwk"

func BarBfCfgAdapter(appCfg fwk.AppCfgInfo) BarBfCfgInfo {
	return BarBfCfgInfo{
		Z: appCfg.Y,
	}
}
