/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package errx

import "github.com/pvillela/go-tryout/arch/errx/errx4"

type Errx = errx4.Errx

var NewErrx = errx4.NewErrx
var ErrxOf = errx4.ErrxOf
var StackTraceOf = errx4.StackTraceOf

type Kind = errx4.Kind

var KindOf = errx4.KindOf
var NewKind = errx4.NewKind
