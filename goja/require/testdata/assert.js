/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

function ok(pred, msg) {
  if (!pred) {
    throw new Error(msg);
  }
}

function equal(actual, expected, msg) {
  ok(actual === expected, msg);
}

module.exports = ok;

module.exports.assert = ok;
module.exports.ok = ok;
module.exports.equal = equal;
