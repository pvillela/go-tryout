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

func fooSflCfgAdapter(appCfg config.AppCfgInfo) fs.FooSflCfgInfo {
	return fs.FooSflCfgInfo{
		X: appCfg.X,
	}
}

func FooSflBoot(appCfg func() config.AppCfgInfo) fs.FooSflT {
	return fs.FooSflC(fs.FooSflCfgSrc{
		Get:   func() fs.FooSflCfgInfo { return fooSflCfgAdapter(appCfg()) },
		BarBf: BarBfBoot(appCfg),
	})
}
