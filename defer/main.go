/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package main

import "fmt"

func thirdOrder(x string) func(y string) func(z string) {
	fmt.Println("third order:", x)
	return func(y string) func(z string) {
		fmt.Println("second order:", y)
		return func(z string) {
			fmt.Println("first order:", z)
		}
	}
}

func fThirdOrder() {
	defer thirdOrder("fThirdOrder")
	fmt.Println("Body of fThirdOrder")
}

func fSecondOrder() {
	defer thirdOrder("fThirdOrder")("fSecondOrder")
	fmt.Println("Body of fSecondOrder")
}

func fFirstOrder() {
	defer thirdOrder("fThirdOrder")("fSecondOrder")("fFirstOrder")
	fmt.Println("Body of fFirstOrder")
}

func trace(name string) func() {
	fmt.Printf("%s entered\n", name)
	return func() {
		fmt.Printf("%s returned\n", name)
	}

}

func fTrace0() {
	defer trace("fTrace")
	fmt.Println("Body of fTrace")
}

func fTrace1() {
	defer trace("fTrace")()
	fmt.Println("Body of fTrace")
}

func foo(x string) {
	fmt.Println("foo:", x)
}

func fFoo() {
	defer foo("fFoo")
	fmt.Println("Body of fFoo")
}

func bar() {
	fmt.Println("bar")
}

func fBar() {
	defer bar()
	fmt.Println("Body of fBar")
}

func main() {
	defer fmt.Println("baz")

	fThirdOrder()
	fmt.Println()
	fSecondOrder()
	fmt.Println()
	fFirstOrder()
	fmt.Println()

	fTrace0()
	fmt.Println()
	fTrace1()
	fmt.Println()

	fFoo()
	fmt.Println()

	fBar()
	fmt.Println()
}
