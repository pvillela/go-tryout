/*
 * Copyright © 2022 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the MIT license
 * that can be found in the LICENSE file.
 */

import { AppCfg } from "./app-cfg";
import { Cfg } from "./module1";

func adapter(appCfg AppCfg) Cfg {
  return {
    z appCfg.y
  };
}
