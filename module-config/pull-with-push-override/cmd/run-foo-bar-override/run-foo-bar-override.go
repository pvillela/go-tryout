/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package main

import "github.com/pvillela/go-tryout/module-config/pull-with-push-override/fs"

func main() {
	fs.FooSflCfgSrc.Set(func() fs.FooSflCfgInfo {
		return fs.FooSflCfgInfo{X: "foo"}
	})

	fs.BarBfCfgSrc.Set(func() fs.BarBfCfgInfo {
		return fs.BarBfCfgInfo{Z: 99}
	})

	fs.FooSfl()
}
