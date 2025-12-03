package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Individual struct {
	genes []float64
	fit   float64
}

func fitness(ind Individual) float64 {
	x := ind.genes[0]
	return -(math.Pow(x-3, 2)) + 10 // peak at x=3
}

func mutate(ind *Individual) {
	ind.genes[0] += rand.NormFloat64() * 0.3
}

func main() {
	rand.Seed(time.Now().UnixNano())

	pop := make([]Individual, 20)
	for i := range pop {
		pop[i] = Individual{genes: []float64{rand.Float64() * 10}}
	}

	for gen := 0; gen < 20; gen++ {
		done := make(chan Individual, len(pop))

		for _, ind := range pop {
			go func(ind Individual) {
				ind.fit = fitness(ind)
				done <- ind
			}(ind)
		}

		for i := range pop {
			pop[i] = <-done
		}

		// Select best
		best := pop[0]
		for _, ind := range pop {
			if ind.fit > best.fit {
				best = ind
			}
		}

		// Mutate to produce next generation
		for i := range pop {
			pop[i] = best
			mutate(&pop[i])
		}

		fmt.Println("Gen", gen, "Best =", best.genes[0], "Fit =", best.fit)
	}
}
