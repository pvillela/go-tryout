/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package main

import (
	"fmt"
	"github.com/pvillela/go-tryout/arch/errx"
	"github.com/pvillela/go-tryout/arch/util"
)

func main() {
	{
		x := 5
		fmt.Println(util.SafeCast[int](x))
	}
	{
		x := 5
		fmt.Println(util.SafeCast[string](x))
	}
	{
		x := "abc"
		fmt.Println(util.SafeCast[string](x))
	}
	{
		x := "abc"
		fmt.Println(util.SafeCast[int](x))
	}
	{
		x := errx.NewErrx(nil, "FooError")
		fmt.Println(util.SafeCast[errx.Errx](x))
	}
	{
		x := errx.NewErrx(nil, "FooError")
		fmt.Println(util.SafeCast[error](x))
	}
	{
		x := errx.NewErrx(nil, "FooError")
		fmt.Println(util.SafeCast[int](x))
	}
}
