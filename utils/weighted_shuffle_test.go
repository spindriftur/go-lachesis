package utils

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

// Test average distribution of the shuffle
func Test_WeightedPermutation_distribution(t *testing.T) {
	weightsArr := getTestWeightsIncreasing(30)

	weightHits := make(map[int]int) // weight -> number of occurrences
	for roundSeed := 0; roundSeed < 3000; roundSeed++ {
		seed := hashOf(common.Hash{}, uint32(roundSeed))
		perm := WeightedPermutation(len(weightsArr)/10, weightsArr, seed)
		for _, p := range perm {
			weight := weightsArr[p]
			weightFactor := int(weight / 1000)

			_, ok := weightHits[weightFactor]
			if !ok {
				weightHits[weightFactor] = 0
			}
			weightHits[weightFactor]++
		}
	}

	assertar := assert.New(t)
	for weightFactor, hits := range weightHits {
		//fmt.Printf("Test_RandomElection_distribution: %d \n", hits/weightFactor)
		assertar.Equal((hits/weightFactor) > 20-8, true)
		assertar.Equal((hits/weightFactor) < 20+8, true)
		if t.Failed() {
			return
		}
	}
}

func Test_WeightedPermutation_correctness(t *testing.T) {
	testCorrectPermutation(WeightedPermutation, t, getTestWeightsIncreasing(1))
	testCorrectPermutation(WeightedPermutation, t, getTestWeightsIncreasing(30))
	testCorrectPermutation(WeightedPermutation, t, getTestWeightsEqual(1000))
}

func Test_WeightedPermutation_determinism(t *testing.T) {
	assertar := assert.New(t)

	weightsArr := getTestWeightsIncreasing(5)

	assertar.Equal([]int{3, 2, 4, 1, 0}, WeightedPermutation(len(weightsArr), weightsArr, hashOf(common.Hash{}, 0)))
	assertar.Equal([]int{0, 4, 2, 1, 3}, WeightedPermutation(len(weightsArr), weightsArr, hashOf(common.Hash{}, 1)))
	assertar.Equal([]int{3, 4, 2, 1, 0}, WeightedPermutation(len(weightsArr), weightsArr, hashOf(common.Hash{}, 2)))
	assertar.Equal([]int{4, 2, 1, 3, 0}, WeightedPermutation(len(weightsArr), weightsArr, hashOf(common.Hash{}, 3)))
	assertar.Equal([]int{1, 4}, WeightedPermutation(len(weightsArr)/2, weightsArr, hashOf(common.Hash{}, 4)))
}

func Test_WeightedPermutation_concurency(t *testing.T) {
	testPermutationConcurency(WeightedPermutation, t)
}
