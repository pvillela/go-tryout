/*
 * Copyright © 2022 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the MIT license
 * that can be found in the LICENSE file.
 */

export type Cfg = {
  x string;
};

export type Dep = {
  bar () => void;
};

export type Foo = () => void;

func fooC(cfg () => Cfg, dep Dep) Foo {
  return function() {
    console.log(cfg().x);
    dep.bar()
  }
}
