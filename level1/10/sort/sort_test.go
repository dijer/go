package sort

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// go test -bench=.
func BenchmarkInsertionSort(b *testing.B) {
	slice := []int{3, 5, 1, -2, 0, 10, 5}
	for i := 0; i < b.N; i++ {
		InsertionSort(slice)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	slice := []int{3, 5, 1, -2, 0, 10, 5}
	for i := 0; i < b.N; i++ {
		BubbleSort(slice)
	}
}

// example
func ExampleInsertionSort() {
	slice := []int{3, 6, 1}
	result, err := BubbleSort(slice)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

type testCase struct {
	name          string
	input         []int
	expected      []int
	expectedError error
}

// m test
func TestInsertionSort(t *testing.T) {
	testCases := []testCase{
		{
			"case with [0,1]",
			[]int{1, 0},
			[]int{0, 1},
			nil,
		},
		{
			"case with [3,1,-1]",
			[]int{3, 1, -1},
			[]int{-1, 1, 3},
			nil,
		},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.name, func(t *testing.T) {
			result, err := InsertionSort(cse.input)
			if !errors.Is(err, cse.expectedError) {
				t.Errorf("unexpected err: %w", err)
				return
			}

			if !reflect.DeepEqual(result, cse.expected) {
				t.Errorf("for %d expected %d, got %d", cse.input, cse.expected, result)
			}
		})
	}
}

func TestTestify(t *testing.T) {
	input := []int{1, 0}
	result, _ := BubbleSort(input)
	expected := []int{0, 1}
	assert.True(t, reflect.DeepEqual(result, expected), "{1,0}")
}
