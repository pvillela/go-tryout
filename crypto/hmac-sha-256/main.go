/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

// Based on https://golangcode.com/generate-sha256-hmac/

package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

func main() {

	// Creating a 512-bit secret by concatenating two 256-bit parts as in the SMART Health Cards
	// revocation specification
	hexSecretPart1 := strings.Repeat("5", 64)
	hexSecretPart2 := strings.Repeat("0", 32) + strings.Repeat("7", 32)
	hexSecret := hexSecretPart1 + hexSecretPart2

	data := "data"

	fmt.Printf("hexSecret: %s data: %s\n", hexSecret, data)

	byteSecret, err := hex.DecodeString(hexSecret)
	if err != nil {
		panic(err)
	}

	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, byteSecret)

	// Write Data to it
	h.Write([]byte(data))

	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))

	fmt.Println("Result: " + sha)
}
