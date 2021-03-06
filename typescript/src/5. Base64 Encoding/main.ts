/**
 * {@link https://www.codewars.com/kata/5270f22f862516c686000161/train/typescript}
 */
String.prototype.toBase64 = function () {
  let res = "";
  for (let i = 0; i < this.length; i += 3) {
    const sub = this.substr(i, 3)
      .split("")
      .map((s) => s.charCodeAt(0));

    res += map[sub[0] >>> 2];
    res += map[((sub[0] << 4) & 0b111111) | (sub[1] >>> 4)];
    if (!sub[1]) continue;
    res += map[((sub[1] << 2) & 0b111111) | (sub[2] >>> 6)];
    if (!sub[2]) continue;
    res += map[sub[2] & 0b00111111];
  }
  return res;
};

String.prototype.fromBase64 = function () {
  let res = "";
  for (let i = 0; i < this.length; i += 4) {
    const sub = this.substr(i, 4)
      .split("")
      .map((s) => map.findIndex((c) => s == c));

    res += String.fromCharCode((sub[0] << 2) | (sub[1] >>> 4));
    if (!sub[2]) continue;
    res += String.fromCharCode(((sub[1] << 4) & 0b11111111) | (sub[2] >>> 2));
    if (!sub[3]) continue;
    res += String.fromCharCode(((sub[2] << 6) & 0b11111111) | sub[3]);
  }
  return res;
};

const map = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/".split(
  ""
);

declare interface String {
  toBase64: () => string;
  fromBase64: () => string;
}
