/*
 * Copyright © 2022 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the MIT license
 * that can be found in the LICENSE file.
 */

package mod

import "github.com/pvillela/go-tryout/module-config/pull-with-push-override/fwk"

func FooSflCfgAdapter(appCfg fwk.AppCfgInfo) FooSflCfgInfo {
	return FooSflCfgInfo{
		X: appCfg.X,
	}
}