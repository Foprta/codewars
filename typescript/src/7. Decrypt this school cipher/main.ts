/**
 * {@link https://www.codewars.com/kata/5cd48cffaae6e30018943175/train/typescript}
 */
export const decrypt = (str: string): string => {
  return str
    .replace(/'\d+'/g, (match) =>
      String.fromCharCode(parseInt(match.slice(1, -1)))
    )
    .split("")
    .reverse()
    .join("");
};
