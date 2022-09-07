/*
 * Copyright © 2022 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package util

import (
	"github.com/pvillela/go-tryout/arch/types"
)

// Contextualizer represents a higher-order function that executes another function within
// a context.
// An example is a function that executes a target function while delimiting the target function
// within a transaction.
//
// - configCtx represents configuration data that is passed to the higher-orfer function.
// For example, a database object.
//
// - externalRuntimeCtx represents a runtime context input for use by both the Contextualizer
// and the target function.
// For example, context.Context.
//
// - internalRuntimeCtx is a runtime argument produced by the Contextualizer and passed to
// the target function.
// For example, a transaction object.
type Contextualizer[CC, ERC, IRC, T any] func(
	configCtx CC,
	externalRuntimeCtx ERC,
	block func(externalRuntimeCtx ERC, internalRuntimeCtx IRC) (T, error),
) (T, error)

// Contextualize returns a function that is the partial application of f
// within the execution context provided by contextualizer.
func Contextualize[CC, ERC, IRC, T any](
	contextualizer Contextualizer[CC, ERC, IRC, T],
	configCtx CC,
	f func(externalRuntimeCtx ERC, internalRuntimeCtx IRC) (T, error),
) func(externalRuntimeCtx ERC) (T, error) {
	return func(externalRuntimeCtx ERC) (T, error) {
		return contextualizer(configCtx, externalRuntimeCtx, f)
	}
}

// Contextualize1 returns a function that is the partial application of f
// within the execution context provided by contextualizer.
func Contextualize1[CC, ERC, IRC, S1, T any](
	contextualizer Contextualizer[CC, ERC, IRC, T],
	configCtx CC,
	f func(externalRuntimeCtx ERC, internalRuntimeCtx IRC, s1 S1) (T, error),
) func(externalRuntimeCtx ERC, s1 S1) (T, error) {
	return func(externalRuntimeCtx ERC, s1 S1) (T, error) {
		block := func(externalRuntimeCtx ERC, internalRuntimeCtx IRC) (T, error) {
			return f(externalRuntimeCtx, internalRuntimeCtx, s1)
		}
		return contextualizer(configCtx, externalRuntimeCtx, block)
	}
}

// Contextualize2 returns a function that is the partial application of f
// within the execution context provided by contextualizer.
func Contextualize2[CC, ERC, IRC, S1, S2, T any](
	contextualizer Contextualizer[CC, ERC, IRC, T],
	configCtx CC,
	f func(externalRuntimeCtx ERC, internalRuntimeCtx IRC, s1 S1, s2 S2) (T, error),
) func(externalRuntimeCtx ERC, s1 S1, s2 S2) (T, error) {
	return func(externalRuntimeCtx ERC, s1 S1, s2 S2) (T, error) {
		block := func(externalRuntimeCtx ERC, internalRuntimeCtx IRC) (T, error) {
			return f(externalRuntimeCtx, internalRuntimeCtx, s1, s2)
		}
		return contextualizer(configCtx, externalRuntimeCtx, block)
	}
}

// Contextualize1V returns a function that is the partial application of f
// within the execution context provided by contextualizer.
func Contextualize1V[CC, ERC, IRC, S1 any](
	contextualizer Contextualizer[CC, ERC, IRC, types.Unit],
	configCtx CC,
	f func(externalRuntimeCtx ERC, internalRuntimeCtx IRC, s1 S1),
) func(externalRuntimeCtx ERC, s1 S1) {
	return func(externalRuntimeCtx ERC, s1 S1) {
		block := func(externalRuntimeCtx ERC, internalRuntimeCtx IRC) (types.Unit, error) {
			f(externalRuntimeCtx, internalRuntimeCtx, s1)
			return types.UnitV, nil
		}
		_, err := contextualizer(configCtx, externalRuntimeCtx, block)
		if err != nil {
			panic(err)
		}
	}
}
