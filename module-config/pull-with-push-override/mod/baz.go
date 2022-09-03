/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package mod

import (
	"fmt"
	"github.com/pvillela/go-tryout/module-config/pull-with-push-override/fwk"
)

type BazCfgInfo struct {
	X string
}

var BazCfgSrc = fwk.MakeConfigSource[BazCfgInfo](nil)

func Baz() {
	fmt.Println(BazCfgSrc.Get().X)
}
