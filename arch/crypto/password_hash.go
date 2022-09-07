/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package crypto

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/alexedwards/argon2id"
	"github.com/pvillela/go-tryout/arch/errx"
	"golang.org/x/crypto/bcrypt"
)

// bcryptCost is the cost parameter for the bcrypt hashing algorithm.
// On an intel Intel(R) Core(TM) i7-7500U CPU @ 2.70GHz with 16 GB of memory, hash creation
// requires ~500ms.
var bcryptCost = 13

func BcryptPasswordHash(password string) string {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	errx.PanicOnError(err)
	return string(hashBytes)
}

func BcryptPasswordCheck(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// argonParamsProd defines the cost parameters for the argon2id hashing algorithm in production.
// On an intel Intel(R) Core(TM) i7-7500U CPU @ 2.70GHz with 16 GB of memory, hash creation
// requires ~400-500ms.
var argonParamsProd = &argon2id.Params{
	Memory:      1 * 1024 * 1024,
	Iterations:  1,
	Parallelism: 3,
	SaltLength:  16,
	KeyLength:   32,
}

// argonParamsDev defines the cost parameters for the argon2id hashing algorithm in development.
// Unlike production, the parameters are chosen for fast execution.
var argonParamsDev = &argon2id.Params{
	Memory:      10 * 1024,
	Iterations:  1,
	Parallelism: 1,
	SaltLength:  8,
	KeyLength:   8,
}

// argonParams defines the cost parameters for the ArgonPasswordHash function.
var argonParams = argonParamsDev

func ArgonPasswordHash(password string) string {
	hash, err := argon2id.CreateHash(password, argonParams)
	errx.PanicOnError(err)
	return hash
}

func ArgonPasswordCheck(password string, hash string) bool {
	match, err := argon2id.ComparePasswordAndHash(password, hash)
	errx.PanicOnError(err)
	return match
}

// RandomString generates a hex string corresponding to a random byte slice of length n.
func RandomString(n int) string {
	buf := make([]byte, n)
	_, _ = rand.Reader.Read(buf)
	return hex.EncodeToString(buf)
}
