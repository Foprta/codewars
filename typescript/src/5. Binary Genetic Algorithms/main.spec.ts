import assert from "assert";
import { GeneticAlgorithm } from "./main";

const ideal = "1111";

const fitness = (c: string) => {
  let r = 0;
  for (let i = 0; i < c.length; ++i) {
    if (c[i] === ideal[i]) {
      r++;
    }
  }
  return r / ideal.length;
};

xdescribe("GeneticAlgorithm", () => {
  let genAlg: GeneticAlgorithm;

  beforeEach(() => {
    genAlg = new GeneticAlgorithm();
  });

  it("should generate random chromosome", () => {
    assert.strictEqual(genAlg.generate(32).length, 32);
    assert.strictEqual(genAlg.generate(5).length, 5);
    assert.strictEqual(genAlg.generate(0).length, 0);
  });

  it("should mutate chromosome", () => {
    assert.strictEqual(genAlg.mutate("0101", 1), "1010");
    assert.strictEqual(genAlg.mutate("0101", 0), "0101");
    assert.strictEqual(/1/.test(genAlg.mutate("0000000000000000", 0.5)), true);
    assert.strictEqual(/0/.test(genAlg.mutate("1111111111111111", 0.5)), true);
  });

  it("should crossover 2 chromosomes", () => {
    assert.strictEqual(genAlg.crossover("110", "001", 2)[0], "111");
    assert.strictEqual(genAlg.crossover("110", "001", 2)[1], "000");
    assert.strictEqual(genAlg.crossover("111000", "000110", 3)[0], "111110");
    assert.strictEqual(genAlg.crossover("111000", "000110", 3)[1], "000000");
  });

  xit("should mostly select most fitness chromosome", () => {
    const pop = ["0", "1", "2", "3", "4", "5", "6", "7", "8", "9"];
    const probs = [0, 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9];

    const res: Record<string, number> = {};

    for (let i = 0; i < 1000; i++) {
      const selected = genAlg.select(pop, probs);
      res[selected] ? res[selected]++ : (res[selected] = 1);
    }

    // console.log(res);
  });

  it("should finally works", () => {
    assert.strictEqual(genAlg.run(fitness, ideal.length, 0.6, 0.02), ideal);
  });
});
