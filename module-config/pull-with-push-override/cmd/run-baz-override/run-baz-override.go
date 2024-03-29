/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package main

import "github.com/pvillela/go-tryout/module-config/pull-with-push-override/fs"

func main() {
	fs.BazCfgSrcV = func() fs.BazCfgInfo {
		return fs.BazCfgInfo{X: "bar"}
	}
	fs.Baz()
}
