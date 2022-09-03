/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package mod

import (
	"fmt"
	"github.com/pvillela/go-tryout/module-config/push-to-file/fwk"
)

type FooSflCfgInfo struct {
	X string
}

var FooSflCfgSrc = fwk.MakeConfigSource[FooSflCfgInfo]()

func FooSfl() {
	fmt.Println(FooSflCfgSrc.Get().X)
	BarBf()
}
