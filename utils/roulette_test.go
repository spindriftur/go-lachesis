package utils

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func Test_StochasticPermutation_correctness(t *testing.T) {
	testCorrectPermutation(StochasticPermutation, t, getTestWeightsIncreasing(1))
	testCorrectPermutation(StochasticPermutation, t, getTestWeightsIncreasing(30))
	testCorrectPermutation(StochasticPermutation, t, getTestWeightsEqual(1000))
}

func Test_StochasticPermutation_determinism(t *testing.T) {
	assertar := assert.New(t)

	weightsArr := getTestWeightsIncreasing(5)

	assertar.Equal([]int{2, 3, 4, 1, 0}, StochasticPermutation(len(weightsArr), weightsArr, hashOf(common.Hash{}, 0)))
	assertar.Equal([]int{2, 3, 4, 1, 0}, StochasticPermutation(len(weightsArr), weightsArr, hashOf(common.Hash{}, 1)))
	assertar.Equal([]int{2, 3, 4, 1, 0}, StochasticPermutation(len(weightsArr), weightsArr, hashOf(common.Hash{}, 3)))
	assertar.Equal([]int{2, 3}, StochasticPermutation(len(weightsArr)/2, weightsArr, hashOf(common.Hash{}, 4)))
}

func Test_StochasticPermutation_concurency(t *testing.T) {
	testPermutationConcurency(StochasticPermutation, t)
}
