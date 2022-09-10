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

var FooSflCfgAdapter = func(appCfgSrc config.AppCfgSrc) func() fs.FooSflCfgInfo {
	return func() fs.FooSflCfgInfo {
		return fs.FooSflCfgInfo{
			X: appCfgSrc().X,
		}
	}
}

func FooSflBoot(appCfgSrc config.AppCfgSrc) fs.FooSflT {
	return fs.FooSflC(fs.FooSflCfgSrc{
		Get:   FooSflCfgAdapter(appCfgSrc),
		BarBf: BarBfBoot(appCfgSrc),
	})
}
