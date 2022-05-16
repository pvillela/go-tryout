/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package errx

import "github.com/pvillela/go-tryout/arch/errx/errx3"

type Errx = errx3.Errx

var NewErrx = errx3.NewErrx
var ErrxOf = errx3.ErrxOf
var StackTraceOf = errx3.StackTraceOf

type Kind = errx3.Kind

var KindOf = errx3.KindOf
var NewKind = errx3.NewKind
