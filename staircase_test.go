package staircase

import (
	"fmt"
	"testing"

	// "github.com/kr/pretty"
	"github.com/stretchr/testify/assert"
)

func TestStaircase(t *testing.T) {
	// t.Skip()
	totalSteps := 4
	steps := []int{1, 2}
	expected := 5
	actual := staircase(totalSteps, steps)

	fmt.Printf("staircase: %d ways of climb\n", actual)
	assert.Equal(t, expected, actual, "they should be equal")
}

func TestStaircaseMemoize(t *testing.T) {
	// t.Skip()
	totalSteps := 4
	steps := []int{1, 2}
	memo := make([]int, totalSteps+1)

	expected := 5
	actual := staircaseMemoize(totalSteps, steps, memo)

	fmt.Printf("staircaseMemoize: %d ways of climb\n", actual)
	assert.Equal(t, expected, actual, "they should be equal")
}

func TestStaircaseBottomUp(t *testing.T) {
	// t.Skip()
	totalSteps := 4
	steps := []int{1, 2}
	expected := 5
	actual := staircaseBottomUp(totalSteps, steps)

	fmt.Printf("staircaseBottomUp: %d ways of climb\n", actual)
	assert.Equal(t, expected, actual, "they should be equal")
}

// func BenchmarkEggDrop(b *testing.B) {
// 	// t.Skip()
// 	egg := 2
// 	floor := 10

// 	for i := 0; i < b.N; i++ {
// 		eggDrop(egg, floor)
// 	}
// }

// func BenchmarkEggMemoize(b *testing.B) {
// 	// t.Skip()
// 	egg := 3
// 	floor := 1000

// 	eggIdx := egg + 1
// 	floorIdx := floor + 1

// 	// create and initialize memo
// 	memo := make([][]int, eggIdx)
// 	for i := 0; i < eggIdx; i++ {
// 		memo[i] = make([]int, floorIdx)
// 	}

// 	for i := 0; i < eggIdx; i++ {
// 		for j := 0; j < floorIdx; j++ {
// 			memo[i][j] = MinVal
// 		}
// 	}

// 	for i := 0; i < b.N; i++ {
// 		eggMemoize(egg, floor, memo)
// 	}
// }

// func BenchmarkEggBottomUp(b *testing.B) {
// 	// t.Skip()
// 	egg := 3
// 	floor := 1000

// 	eggIdx := egg + 1
// 	floorIdx := floor + 1

// 	// create and initialize memo
// 	memo := make([][]int, eggIdx)
// 	for i := 0; i < eggIdx; i++ {
// 		memo[i] = make([]int, floorIdx)
// 	}

// 	for i := 0; i < eggIdx; i++ {
// 		for j := 0; j < floorIdx; j++ {
// 			memo[i][j] = MinVal
// 		}
// 	}

// 	for i := 0; i < b.N; i++ {
// 		eggBottomUp(egg, floor, memo)
// 	}
// }
