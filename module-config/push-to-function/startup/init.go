/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package startup

import (
	"github.com/pvillela/go-tryout/module-config/push-to-function/fs/boot"
	"github.com/pvillela/go-tryout/module-config/push-to-var/config"
)

var FooSfl = boot.FooSflBoot(config.GetAppConfiguration)
