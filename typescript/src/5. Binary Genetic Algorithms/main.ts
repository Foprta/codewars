/**
 * {@link https://www.codewars.com/kata/526f35b9c103314662000007/train/typescript}
 */
export class GeneticAlgorithm {
  private readonly mutateMap: Record<string, string> = { "0": "1", "1": "0" };

  generate(length: number): string {
    if (length === 0) return "";

    return Math.floor(Math.random() * Math.pow(2, length))
      .toString(2)
      .padStart(length, "0");
  }

  select(population: string[], fitnesses: number[]): string {
    let fitnessesSum = fitnesses.reduce((a, f) => (a += f), 0);

    for (let i = 0; i < population.length; i++) {
      const random = Math.random();
      if (random <= fitnesses[i] / fitnessesSum) {
        return population[i];
      }
      fitnessesSum -= fitnesses[i];
    }

    return "";
  }

  mutate(chromosome: string, p: number): string {
    return chromosome
      .split("")
      .map((v) => (Math.random() <= p ? this.mutateMap[v] : v))
      .join("");
  }

  crossover(chromosome1: string, chromosome2: string, index: number): string[] {
    const left1 = chromosome1.slice(0, index);
    const right1 = chromosome1.slice(index);
    const left2 = chromosome2.slice(0, index);
    const right2 = chromosome2.slice(index);
    return [left1.concat(right2), left2.concat(right1)];
  }

  run(
    fitness: (chromosome: string) => number,
    length: number,
    p_c: number,
    p_m: number,
    iterations = 100,
    popLength = 100
  ): string {
    let pop: string[] = new Array(popLength)
      .fill("")
      .map(() => this.generate(length));

    for (let i = 0; i < iterations; i++) {
      const newPop = [];

      const fitnesses = pop.map((c) => fitness(c));

      while (newPop.length < popLength) {
        const selected1 = this.select(pop, fitnesses);
        const selected2 = this.select(pop, fitnesses);

        if (Math.random() < p_c) {
          newPop.push(
            ...this.crossover(
              selected1,
              selected2,
              Math.floor(Math.random() * length)
            )
          );
        } else {
          newPop.push(selected1, selected2);
        }
      }

      pop = newPop.map((v) => this.mutate(v, p_m));
    }

    return pop.reduce((r, c) => (fitness(c) > fitness(r) ? c : r), pop[0]);
  }
}
