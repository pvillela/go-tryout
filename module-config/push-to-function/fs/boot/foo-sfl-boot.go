/*
 * Copyright © 2022 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the MIT license
 * that can be found in the LICENSE file.
 */

package boot

import (
	"github.com/pvillela/go-tryout/module-config/push-to-file/fwk"
	"github.com/pvillela/go-tryout/module-config/push-to-function/fs"
)

func fooSflCfgAdapter(appCfg fwk.AppCfgInfo) fs.FooSflCfgInfo {
	return fs.FooSflCfgInfo{
		X: appCfg.X,
	}
}

func FooSflBoot(appCfg func() fwk.AppCfgInfo) fs.FooSflT {
	barBf := BarBfBoot(appCfg)
	return fs.FooSflC(fs.FooSflCfgSrc{
		Get:   func() fs.FooSflCfgInfo { return fooSflCfgAdapter(appCfg()) },
		BarBf: barBf,
	})
}
