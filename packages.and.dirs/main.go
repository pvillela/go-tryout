/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

/*
 * The imported path has a "." in the last segment. In addition, the package name of the
 * imported function Foo is not the last segment of the import path.
 */

package main

import (
	"fmt"
	"github.com/pvillela/go-tryout/packages.and.dirs/abc.x"
)

func main() {
	y.Foo()
	fmt.Println("success")
}
