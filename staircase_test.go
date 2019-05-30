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
	totalSteps := 100
	steps := []int{1, 3, 5, 7, 9}
	memo := make([]uint, totalSteps+1)

	var expected uint
	expected = 10767769711533586674
	actual := staircaseMemoize(totalSteps, steps, memo)

	fmt.Printf("staircaseMemoize: %d ways of climb\n", actual)
	assert.Equal(t, expected, actual, "they should be equal")
}

func TestStaircaseBottomUp(t *testing.T) {
	// t.Skip()
	totalSteps := 100
	steps := []int{1, 3, 5, 7, 9}
	var expected uint
	expected = 10767769711533586674
	actual := staircaseBottomUp(totalSteps, steps)

	fmt.Printf("staircaseBottomUp: %d ways of climb\n", actual)
	assert.Equal(t, expected, actual, "they should be equal")
}
func BenchmarkStaircase(b *testing.B) {
	// t.Skip()
	totalSteps := 4
	steps := []int{1, 2}

	for i := 0; i < b.N; i++ {
		staircase(totalSteps, steps)
	}
}

func BenchmarkStaircaseMemoize(b *testing.B) {
	// t.Skip()
	totalSteps := 100
	steps := []int{1, 3, 5, 7, 9}

	for i := 0; i < b.N; i++ {
		memo := make([]uint, totalSteps+1)
		staircaseMemoize(totalSteps, steps, memo)
	}
}

func BenchmarkStaircaseBottomUp(b *testing.B) {
	// t.Skip()
	totalSteps := 100
	steps := []int{1, 3, 5, 7, 9}

	for i := 0; i < b.N; i++ {
		staircaseBottomUp(totalSteps, steps)
	}
}
