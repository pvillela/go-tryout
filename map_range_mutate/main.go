/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package main

import "fmt"

func main() {
	m := map[string]int{"a": 1, "b": 2}
	fmt.Println(m)
	for key := range m {
		m[key] = m[key] + 1
	}
	fmt.Println(m)
}
