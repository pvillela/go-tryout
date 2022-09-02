/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package main

import "github.com/pvillela/go-tryout/module-config/pull-with-push-override/mod"

func initialize() {
	mod.Module0Cfg.Set(func() mod.Module0CfgT {
		return mod.Module0CfgT{X: "bar"}
	})

	mod.Module1Cfg.Set(func() mod.Module1CfgT {
		return mod.Module1CfgT{Z: 99}
	})
}

func main() {
	initialize()
	mod.Foo()
}
