/*
 * Copyright © 2022 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the MIT license
 * that can be found in the LICENSE file.
 */

import { getAppConfiguration as c } from "./app-cfg";
import * as mod1 from "./module1"
import * as mod0 from "./module0";
import { adapter as mod1Adapter } from "./module1-cfg-adapter";

const bar = (() => {
  const cfg = () => mod1Adapter(c())
  return mod1.barC(cfg)
})()

export const foo = mod0.fooC(c, { bar });
