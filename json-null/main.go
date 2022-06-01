/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package main

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

type Foo struct {
	X int  `json:"x"`
	Y *int `json:"y"`
	Z *int `json:"z,omitempty"`
}

func marshalAndDisplay(foo Foo) string {
	_, _ = spew.Printf("foo: %+v\n", foo)
	marshalled, _ := json.Marshal(foo)
	marshalledString := string(marshalled)
	fmt.Println("marshalledString: ", marshalledString)
	return marshalledString
}

func unmarshalAndDisplay(str string) Foo {
	marshalled := []byte(str)
	var unmarshalled Foo
	_ = json.Unmarshal(marshalled, &unmarshalled)
	_, _ = spew.Printf("unmarshalled: %+v\n", unmarshalled)
	return unmarshalled
}

func main() {
	{
		x := 1
		y := 2
		z := 3
		foo := Foo{
			X: x,
			Y: &y,
			Z: &z,
		}
		marshalledStr := marshalAndDisplay(foo)
		_ = unmarshalAndDisplay(marshalledStr)
		fmt.Println()
	}
	{
		x := 1
		foo := Foo{
			X: x,
			Y: nil,
			Z: nil,
		}
		marshalledStr := marshalAndDisplay(foo)
		_ = unmarshalAndDisplay(marshalledStr)
		fmt.Println()
	}
	{
		marshalledString := `{"x":1}`
		fmt.Println("marshalledString: ", marshalledString)
		_ = unmarshalAndDisplay(marshalledString)
		fmt.Println()
	}
}
