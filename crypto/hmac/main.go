// Adapted from the overview snippet of https://pkg.go.dev/crypto/hmac.

package main

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/pvillela/go-tryout/arch/util"
)

// generateKey generates a random 256 bit key.
func generateKey() ([]byte, error) {
	key := make([]byte, 32)
	_, err := rand.Reader.Read(key)
	return key, err
}

func sign(msg string, key []byte) []byte {
	msgBytes := []byte(msg)
	mac := hmac.New(sha256.New, key)
	mac.Write(msgBytes)
	return mac.Sum(nil)
}

func verify(msg string, key []byte, expecteHmac []byte) bool {
	sig := sign(msg, key)
	return hmac.Equal(sig, expecteHmac)
}

func main() {
	msg := "hello, world"
	badMsg := "hallo world"

	{
		key, err := generateKey()
		util.PanicOnError(err)
		fmt.Printf("\n********* generated key: %x\n", key)

		sig := sign(msg, key)
		fmt.Printf("signature: %x\n", sig)
		valid := verify(msg, key, sig)
		fmt.Println("signature verified - msg:", valid)
		valid = verify(badMsg, key, sig)
		fmt.Println("signature verified - badMsg:", valid)
	}

	{
		var givenHexKey = "5717e558d867a8f71b0e740bace172fa62ee6c25dfaaafe8937d2049244aaf16"
		var key, err = hex.DecodeString(givenHexKey)
		util.PanicOnError(err)
		fmt.Printf("\n********* given key: %x\n", key)

		sig := sign(msg, key)
		fmt.Printf("signature: %x\n", sig)
		valid := verify(msg, key, sig)
		fmt.Println("signature verified - msg:", valid)
		valid = verify(badMsg, key, sig)
		fmt.Println("signature verified - badMsg:", valid)
	}
}
