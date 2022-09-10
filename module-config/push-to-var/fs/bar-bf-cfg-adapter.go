/*
 * Copyright © 2022 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the MIT license
 * that can be found in the LICENSE file.
 */

package fs

import (
	"github.com/pvillela/go-tryout/module-config/push-to-var/config"
	"github.com/pvillela/go-tryout/module-config/push-to-var/fwk"
)

func barBfCfgAdapter(appCfgSrc config.AppCfgSrc) BarBfCfgSrc {
	return func() BarBfCfgInfo {
		return BarBfCfgInfo{
			Z: appCfgSrc().Y,
		}
	}
}

var BarBfCfgAdaptation = fwk.MakeCfgSrcAdaptation[config.AppCfgInfo, BarBfCfgInfo](
	&BarBfCfgSrcV,
	barBfCfgAdapter,
)
