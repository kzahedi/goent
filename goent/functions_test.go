package main

import (
	"fmt"
	"math"
	"testing"
)

func TestMerge3Data(t *testing.T) {
	t.Log("Testing Mutual Information")
	data := [][]float64{
		{1.1, 2.1, 3.1, 4.1},
		{1.2, 2.2, 3.2, 4.2},
		{1.3, 2.3, 3.3, 4.3},
		{1.4, 2.4, 3.4, 4.4},
		{1.5, 2.5, 3.5, 4.5},
		{1.6, 2.6, 3.6, 4.6},
		{1.7, 2.7, 3.7, 4.7},
		{1.8, 2.8, 3.8, 4.8},
		{1.9, 2.9, 3.9, 4.9}}

	test1 := [][]float64{
		{1.1, 2.1},
		{1.2, 2.2},
		{1.3, 2.3},
		{1.4, 2.4},
		{1.5, 2.5},
		{1.6, 2.6},
		{1.7, 2.7}}

	test2 := [][]float64{
		{2.2, 3.2},
		{2.3, 3.3},
		{2.4, 3.4},
		{2.5, 3.5},
		{2.6, 3.6},
		{2.7, 3.7},
		{2.8, 3.8}}

	test3 := [][]float64{
		{3.3, 4.3},
		{3.4, 4.4},
		{3.5, 4.5},
		{3.6, 4.6},
		{3.7, 4.7},
		{3.8, 4.8},
		{3.9, 4.9}}

	m := merge3Data(data, []int64{0, 1}, 0, []int64{1, 2}, 1, []int64{2, 3}, 2, false)

	if len(m) != 7 {
		t.Errorf(fmt.Sprintf("Length of the array should be 7 but it is %d", len(m)))
	}

	for i := 0; i < 7; i++ {

		if math.Abs(test1[i][0]-m[i][0]) > 0.0000001 {
			t.Errorf(fmt.Sprintf("Test 1: Value missmatch %f vs %f", test1[i][0], m[i][0]))
		}
		if math.Abs(test1[i][1]-m[i][1]) > 0.0000001 {
			t.Errorf(fmt.Sprintf("Test 1: Value missmatch %f vs %f", test1[i][0], m[i][1]))
		}

		if math.Abs(test2[i][0]-m[i][2]) > 0.0000001 {
			t.Errorf(fmt.Sprintf("Test 2: Value missmatch %f vs %f", test2[i][0], m[i][2]))
		}
		if math.Abs(test2[i][1]-m[i][3]) > 0.0000001 {
			t.Errorf(fmt.Sprintf("Test 2: Value missmatch %f vs %f", test2[i][1], m[i][3]))
		}

		if math.Abs(test3[i][0]-m[i][4]) > 0.0000001 {
			t.Errorf(fmt.Sprintf("Test 3: Value missmatch %f vs %f", test3[i][0], m[i][4]))
		}
		if math.Abs(test3[i][1]-m[i][5]) > 0.0000001 {
			t.Errorf(fmt.Sprintf("Test 3: Value missmatch %f vs %f", test3[i][1], m[i][5]))
		}

	}
}
