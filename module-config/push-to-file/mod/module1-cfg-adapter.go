/*
 * Copyright © 2022 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the MIT license
 * that can be found in the LICENSE file.
 */

package mod

import "github.com/pvillela/go-tryout/module-config/push-to-file/fwk"

func Module1Adapter(appCfg fwk.AppCfg) Module1CfgT {
	return Module1CfgT{
		Z: appCfg.Y,
	}
}
