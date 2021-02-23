import assert from "assert";
import { decrypt } from "./main";

describe("decrypt", () => {
  it("should decrypt any messages", () => {
    assert.strictEqual(decrypt("'101''99''105''108''65'"), "Alice");
    assert.strictEqual(decrypt("'98''111''66'"), "Bob");
    assert.strictEqual(decrypt("30 71"), "17 03");
  });
});
