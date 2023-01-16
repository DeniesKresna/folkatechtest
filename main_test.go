package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	type testCase struct {
		Prices            []int
		TotalTransactions int
		Expectation       int
	}
	testCases := []testCase{
		{
			Prices:            []int{4, 11, 2, 20, 59, 80},
			TotalTransactions: 1,
			Expectation:       78,
		},
		{
			Prices:            []int{4, 11, 2, 20, 59, 80},
			TotalTransactions: 2,
			Expectation:       85,
		},
	}

	for _, v := range testCases {
		assert.Equal(t, MaxProfit(v.Prices, v.TotalTransactions), v.Expectation)
	}
}

func TestShiftArray(t *testing.T) {
	type testCase struct {
		Data        []int
		TotalShifts int
		Expectation []int
	}
	testCases := []testCase{
		{
			Data:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			TotalShifts: 1,
			Expectation: []int{4, 1, 2, 7, 5, 3, 8, 9, 6},
		},
		{
			Data:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			TotalShifts: 2,
			Expectation: []int{7, 4, 1, 8, 5, 2, 9, 6, 3},
		},
	}

	for _, v := range testCases {
		assert.Equal(t, ShiftArray(v.Data, v.TotalShifts), v.Expectation)
	}
}
