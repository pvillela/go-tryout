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

func fooSflCfgAdapter(appCfgSrc config.AppCfgSrc) FooSflCfgSrc {
	return func() FooSflCfgInfo {
		return FooSflCfgInfo{
			X: appCfgSrc().X,
		}
	}
}

var FooSflCfgAdaptation = fwk.MakeCfgSrcAdaptation[config.AppCfgInfo, FooSflCfgInfo](
	&FooSflCfgSrcV,
	fooSflCfgAdapter,
)
