/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package main

import "github.com/pvillela/go-tryout/module-config/pull-with-push-override/mod"

func initialize() {
	mod.Module2Cfg.Set(func() mod.Module2CfgT {
		return mod.Module2CfgT{X: "bar"}
	})
}

func main() {
	initialize()
	mod.Foo2()
}