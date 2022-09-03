/*
 * Copyright © 2022 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the MIT license
 * that can be found in the LICENSE file.
 */

package mod

import "github.com/pvillela/go-tryout/module-config/push-to-file/fwk"

func fooSflCfgAdapter(appCfg fwk.AppCfgInfo) FooSflCfgInfo {
	return FooSflCfgInfo{
		X: appCfg.X,
	}
}

func FooSflBoot(appCfg func() fwk.AppCfgInfo) FooSflT {
	barBf := BarBfBoot(appCfg)
	return FooSflC(FooSflCfgSrc{
		Get:   func() FooSflCfgInfo { return fooSflCfgAdapter(appCfg()) },
		BarBf: barBf,
	})
}
