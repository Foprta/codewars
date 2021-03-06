import assert, { strict, strictEqual } from "assert";
import "./main";

xdescribe("Base64", () => {
  it("should encode to base64", () => {
    assert.strictEqual(
      "this is a string!!".toBase64(),
      "dGhpcyBpcyBhIHN0cmluZyEh"
    );
    assert.strictEqual("any carnal pleas".toBase64(), "YW55IGNhcm5hbCBwbGVhcw");
    assert.strictEqual(
      "any carnal pleasu".toBase64(),
      "YW55IGNhcm5hbCBwbGVhc3U"
    );
    assert.strictEqual(
      "any carnal pleasur".toBase64(),
      "YW55IGNhcm5hbCBwbGVhc3Vy"
    );
  });

  it("should decode from base64", () => {
    assert.strictEqual(
      "dGhpcyBpcyBhIHN0cmluZyEh".fromBase64(),
      "this is a string!!"
    );
    assert.strictEqual(
      "YW55IGNhcm5hbCBwbGVhcw".fromBase64(),
      "any carnal pleas"
    );
    assert.strictEqual(
      "YW55IGNhcm5hbCBwbGVhc3U".fromBase64(),
      "any carnal pleasu"
    );
    assert.strictEqual(
      "YW55IGNhcm5hbCBwbGVhc3Vy".fromBase64(),
      "any carnal pleasur"
    );
  });
});
