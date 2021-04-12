/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Foo struct {
	X1   int    `json:"x1"`
	X2   int    `json:"x2Foo"`
	YFoo string `json:"yFoo"`
}

type Bar struct {
	X1 int `json:"x1"`
	X2 int `json:"x2Bar"`
	Foo
}

type Baz struct {
	Foo
	X1 int `json:"x1"`
	X2 int `json:"x2Baz"`
}

type Xyz interface {
	M1() int
}

type Fuz struct {
	Xyz `json:",omitempty"`
	X1  int `json:"x1"`
	X2  int `json:"x2Fuz"`
}

type fuzAlias Fuz // defined to avoid infinite recursion in unmarshalling

func (s Foo) M1() int { return s.X1 }

func (s Bar) M1() int { return s.X1 }

// FuzUnmarshal -- see https://endophage.com/post/golang-parse-to-interface/
// My original solution that I converted to a method below.
// This is a simple example with poor error handling.
func FuzUnmarshal(data []byte, fuz *Fuz) {
	err := json.Unmarshal(data, fuz)
	if err != nil {
		// Ignore error on purpose as other fields have been unmarshalled.
		fmt.Println(err)
	}

	m := make(map[string]json.RawMessage)
	_ = json.Unmarshal(data, &m)
	fmt.Println(m["Xyz"])

	xyzSer, ok := m["Xyz"]
	if !ok {
		return
	}
	fmt.Println("fuz before XyzUnmarshal: ", fuz)
	_ = XyzUnmarshal(xyzSer, &fuz.Xyz)
	fmt.Println("fuz after XyzUnmarshal: ", fuz)
}

//  UnmarshalJSON -- see https://endophage.com/post/golang-parse-to-interface/
func (fuz *Fuz) UnmarshalJSON(data []byte) error {
	fuzAlias := (*fuzAlias)(fuz) // to avoid infinite recursion
	err := json.Unmarshal(data, fuzAlias)
	if err != nil {
		// Ignore error on purpose as other fields have been unmarshalled.
		fmt.Println(err)
	}

	m := make(map[string]json.RawMessage)
	err = json.Unmarshal(data, &m)
	if err != nil {
		return err
	}

	fmt.Println(m["Xyz"])

	xyzSer, ok := m["Xyz"]
	if !ok {
		return errors.New("not a valid Fuz JSON representation")
	}

	fmt.Println("fuz before XyzUnmarshal: ", fuz)
	err = XyzUnmarshal(xyzSer, &fuz.Xyz)
	if err != nil {
		return err
	}
	fmt.Println("fuz after XyzUnmarshal: ", fuz)

	return nil
}

// XyzUnmarshal -- see https://endophage.com/post/golang-parse-to-interface/
func XyzUnmarshal(data []byte, xyz *Xyz) error {
	m := make(map[string]json.RawMessage)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	_, ok := m["x2Bar"]
	if ok {
		bar := Bar{}
		err = json.Unmarshal(data, &bar)
		if err != nil {
			return err
		}
		*xyz = bar
	} else {
		foo := Foo{}
		err = json.Unmarshal(data, &foo)
		if err != nil {
			return err
		}
		*xyz = foo
	}
	return nil
}

func SerializeAndPrint(label string, object interface{}) []byte {
	serialized, _ := json.Marshal(object)
	fmt.Printf("%s as JSON:\n  %s\n", label, serialized)
	fmt.Printf("%s as string:\n  %v\n", label, object)
	return serialized
}

func main() {
	foo := Foo{
		X1:   1,
		X2:   2,
		YFoo: "y1",
	}

	func() {
		fooSer := SerializeAndPrint("foo", foo)
		fooDes := Bar{}
		_ = json.Unmarshal(fooSer, &fooDes)
		SerializeAndPrint("fooDes", fooDes)
		fmt.Println()
	}()

	bar := Bar{
		X1:  100,
		X2:  200,
		Foo: foo,
	}

	func() {
		barSer := SerializeAndPrint("bar", bar)
		barDes := Bar{}
		_ = json.Unmarshal(barSer, &barDes)
		SerializeAndPrint("barDes", barDes)
		fmt.Println()
	}()

	func() {
		baz := Baz{
			X1:  91,
			X2:  92,
			Foo: foo,
		}
		bazSer := SerializeAndPrint("baz", baz)

		bazDes := Baz{}
		_ = json.Unmarshal(bazSer, &bazDes)
		SerializeAndPrint("bazDes", bazDes)
		fmt.Println()
	}()

	func() {
		fuz1 := Fuz{
			X1:  91,
			X2:  92,
			Xyz: foo,
		}
		fuz1Ser := SerializeAndPrint("fuz1", fuz1)

		fuz1Des := Fuz{}
		_ = json.Unmarshal(fuz1Ser, &fuz1Des)
		SerializeAndPrint("fuz1Des", fuz1Des)
		fmt.Println()
	}()

	func() {
		fuz2 := Fuz{
			X1:  91,
			X2:  92,
			Xyz: foo,
		}
		fuz2Ser := SerializeAndPrint("fuz2", fuz2)

		fuz2Des := Fuz{}
		//FuzUnmarshal(fuz2Ser, &fuz2Des)
		_ = json.Unmarshal(fuz2Ser, &fuz2Des)
		SerializeAndPrint("fuz2Des", fuz2Des)
		fmt.Println()
	}()

	func() {
		fuz3 := Fuz{
			X1:  91,
			X2:  92,
			Xyz: bar,
		}
		fuz3Ser := SerializeAndPrint("fuz3", fuz3)

		fuz3Des := Fuz{}
		//FuzUnmarshal(fuz3Ser, &fuz3Des)
		_ = json.Unmarshal(fuz3Ser, &fuz3Des)
		SerializeAndPrint("fuz3Des", fuz3Des)
		fmt.Println()
	}()
}
