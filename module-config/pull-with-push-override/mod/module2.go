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

type Module2CfgT struct {
	X string
}

var Module2Cfg = fwk.MakeConfigSource[Module2CfgT](nil)

func Foo2() {
	fmt.Println(Module2Cfg.Get().X)
}
