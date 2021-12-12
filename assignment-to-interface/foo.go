/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package main

type X interface {
	x() int
	y() int
	z() string
}

type XImpl struct{}

func (XImpl) x() int { return 1 }

func (XImpl) y() int { return 2 }

func (XImpl) z() string { return "zzz" }

type Y struct {
	X
	v string
}

type Z struct {
	XImpl
	w func() int
}

var x X
var y Y
var z Z

func F() {
	x = y
	x = z
}
