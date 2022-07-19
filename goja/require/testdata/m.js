const assert = require("./assert");

function test() {
  assert(true);
  assert.assert(true);
  assert.ok(true);
  assert.equal(1, 1, "all is good");
  console.log("Running test ...")
  return "passed";
}

module.exports = {
  test: test
}
