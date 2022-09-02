/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package fwk

type CfgSrc[T any] interface {
	Set(func() T)
	Get() T
}

type cfgSrcImpl[T any] struct {
	cfgSrc func() T
}

func (cs *cfgSrcImpl[T]) Set(cfgSrc func() T) {
	cs.cfgSrc = cfgSrc
}

func (cs *cfgSrcImpl[T]) Get() T {
	if cs.cfgSrc == nil {
		panic("Module used before being initialized")
	}
	return cs.cfgSrc()
}

func MakeConfigSource[T any]() CfgSrc[T] {
	return &cfgSrcImpl[T]{}
}
