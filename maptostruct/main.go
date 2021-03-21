/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package main

import (
	"encoding/json"
	"fmt"
)

func decode(in, out interface{}) error {
	b, err := json.Marshal(in)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, out)
	return err
}

type country struct {
	Name       string
	Population int
}

func main() {
	countryMap := map[string]interface{}{
		"Name": "Canada",
		"Population":  40_000_000,
	}
	var result country
	err := decode(countryMap, &result)
	fmt.Printf("result=%+v, err=%v\n", result, err)
}
