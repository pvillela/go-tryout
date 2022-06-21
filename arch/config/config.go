/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package config

// ConfigProvider is the generic type of config providers.
type ConfigProvider[C any] func() C

// ProdConfigProvider is the specific type of the config provider for this example,
// exemplifying a production config provider.
func ProdConfigProvider() string {
	return "9"
}
