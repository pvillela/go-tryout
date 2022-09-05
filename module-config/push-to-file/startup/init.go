/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package startup

import (
	"github.com/pvillela/go-tryout/module-config/push-to-file/fs"
	"github.com/pvillela/go-tryout/module-config/push-to-file/fwk"
)

func Initialize() struct{} {
	c := fwk.GetAppConfiguration
	fs.FooSflAdapterCfgSrc.Set(c)
	fs.BarBfAdapterCfgSrc.Set(c)
	return struct{}{}
}

var _ = Initialize()
