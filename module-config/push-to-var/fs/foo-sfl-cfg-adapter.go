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

func fooSflCfgAdapter(appCfg config.AppCfgInfo) FooSflCfgInfo {
	return FooSflCfgInfo{
		X: appCfg.X,
	}
}

var FooSflAdapterCfgSrc = fwk.MakeConfigSource[config.AppCfgInfo]()

var _ = (func() struct{} {
	FooSflCfgSrc.Set(func() FooSflCfgInfo {
		return fooSflCfgAdapter(FooSflAdapterCfgSrc.Get())
	})
	return struct{}{}
})()
