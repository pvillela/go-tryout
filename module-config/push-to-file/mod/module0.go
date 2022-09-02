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

type Module0CfgT struct {
	X string
}

var Module0Cfg = fwk.MakeConfigSource[Module0CfgT]()

func Foo() {
	fmt.Println(Module0Cfg.Get().X)
	Bar()
}
