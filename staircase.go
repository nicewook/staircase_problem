package staircase

// import "strconv"

// const MaxIntVal = int(^uint(0) >> 1)
// const MinIntVal = -MaxIntVal - 1
// const MinVal = -1

// func MaxInt(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }

func staircase(totalSteps int, steps []int) int {
	if totalSteps == 0 {
		return 1
	}

	var result int
	for _, step := range steps {
		stepsLeft := totalSteps - step
		if stepsLeft < 0 {
			continue
		}
		if stepsLeft == 0 {
			result++
			continue
		}
		result += staircase(stepsLeft, steps)
	}
	return result
}

func staircaseMemoize(totalSteps int, steps []int, memo []uint) uint {
	if memo[totalSteps] != 0 {
		return memo[totalSteps]
	}
	if totalSteps == 0 {
		memo[totalSteps] = 1
		return 1
	}

	var result uint
	for _, step := range steps {
		stepsLeft := totalSteps - step
		if stepsLeft < 0 {
			continue
		}
		if stepsLeft == 0 {
			result++
			continue
		}
		result += staircaseMemoize(stepsLeft, steps, memo)
	}
	memo[totalSteps] = result
	return result
}

func staircaseBottomUp(totalSteps int, steps []int) uint {

	if totalSteps == 0 {
		return 1
	}

	memo := make([]uint, totalSteps+1)
	memo[0] = 1

	for tSteps := 1; tSteps <= totalSteps; tSteps++ {
		var result uint
		for _, step := range steps {
			stepsLeft := tSteps - step
			if stepsLeft >= 0 {
				result += memo[stepsLeft]
			}
		}
		memo[tSteps] = result
	}
	return memo[totalSteps]
}
