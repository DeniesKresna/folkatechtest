package main

import (
	"fmt"
	"math"

	"github.com/DeniesKresna/gohelper/utlog"
)

func main() {
	//firstQuestion()
	//secondQuestion()
}

// ==================================================================
// ============= First Question =====================================
// ==================================================================
func firstQuestion() {

	// TODO: change this for test the first question
	prices := []int{4, 11, 2, 20, 59, 80}
	i := 2

	res := MaxProfit(prices, i)
	utlog.Infof("prices: %+v", prices)
	utlog.Infof("total transactions: %d", i)
	utlog.Infof("result: %d", res)
	return
}

// Calculate maximal profit can be gained from the stocks prices
//
// prices: the list of stock price daily
//
// i: total transaction can be done
func MaxProfit(prices []int, i int) (total int) {
	// validation
	if len(prices) <= 0 || i <= 0 {
		return
	}

	diffs := []int{}
	smallest, largest, now, prev := prices[0], prices[0], 0, prices[0]
	for m := 0; m < len(prices); m++ {
		now = prices[m]

		if now < prev || m == len(prices)-1 {
			if largest < now {
				largest = now
			}

			diff := largest - smallest
			diffs = append(diffs, diff)
			smallest = now
			largest = now
			prev = now
			continue
		}

		if smallest > now {
			smallest = now
		}
		if largest < now {
			largest = now
		}

		prev = now
	}

	for p := 0; p < len(diffs)-1; p++ {
		for q := 1; q < len(diffs); q++ {
			if diffs[q] > diffs[p] {
				c := diffs[p]
				diffs[p] = diffs[q]
				diffs[q] = c
			}
		}
	}

	for _, v := range diffs {
		if i == 0 {
			return
		}
		total += v
		i--
	}

	return
}

// ==================================================================
// ============= Second Question ====================================
// ==================================================================
func secondQuestion() {
	// TODO: change this for test the first question
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	i := 2

	utlog.Infof("array: %+v\n", array)
	res := ShiftArray(array, i)
	utlog.Infof("amount of shift: %d\n", i)
	PrintArray(res)
	return
}

func isRealInteger(val float64) bool {
	return val == float64(int(val))
}

func PrintArray(array []int) {
	arrayLength := len(array)
	sqrtWidth := math.Sqrt(float64(arrayLength))
	if !isRealInteger(sqrtWidth) {
		utlog.Errorf("amount of array data not satisfy the array requirement")
	}
	sideWidth := int(sqrtWidth)

	fmt.Printf("\n")
	for index, v := range array {
		cIndex := index + 1
		fmt.Printf("%d", v)
		if cIndex%sideWidth == 0 {
			fmt.Printf("\n")
		}
	}
	fmt.Printf("\n")
}

// Shift outer side data of the 2d array
//
// array: the list of 2d data
//
// i: amount of shift
func ShiftArray(array []int, i int) []int {
	for i > 0 {
		arrayLength := len(array)
		if arrayLength == 0 {
			return array
		}

		if arrayLength == 1 {
			return array
		}

		sqrtWidth := math.Sqrt(float64(arrayLength))
		if !isRealInteger(sqrtWidth) {
			utlog.Errorf("amount of array data not satisfy the array requirement")
			return []int{}
		}
		sideWidth := int(sqrtWidth)

		copyOfArray := make([]int, arrayLength)
		copy(copyOfArray, array)

		for index, _ := range array {
			cIndex := index + 1
			if cIndex%sideWidth == 1 {
				if float64(cIndex)/float64(sideWidth) > float64(sideWidth-1) {
					array[index] = copyOfArray[index+1]
				} else {
					array[index] = copyOfArray[index+sideWidth]
				}
			} else if cIndex%sideWidth == 0 {
				if float64(cIndex)/float64(sideWidth) == 1 {
					array[index] = copyOfArray[index-1]
				} else {
					array[index] = copyOfArray[index-sideWidth]
				}
			} else if float64(cIndex)/float64(sideWidth) < 1 {
				array[index] = copyOfArray[index-1]
			} else if float64(cIndex)/float64(sideWidth) > float64(sideWidth-1) {
				array[index] = copyOfArray[index+1]
			}
		}

		i--
	}

	return array
}
