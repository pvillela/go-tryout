/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package main

import (
	"fmt"
)

func main() {
	var x interface{}
	y := x
	x = 1
	fmt.Println(y) // y is different from x

	var v interface{}
	k := 1
	v = &k
	w := v
	*v.(*int) = 2
	fmt.Println(*w.(*int)) // w is equal to v
}
