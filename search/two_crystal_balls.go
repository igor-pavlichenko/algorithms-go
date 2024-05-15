package search

import (
	"errors"
	"math"
)

/*
Given 2 crystal balls and a 100 floor building. Find out the maximum height from which a ball
can be dropped without breaking in the most optimized way.

Worst-case complexity: O(√n)
Best complexity: O(1)
*/
func TwoCrystalBalls(floors []bool) (int, error) {
	// we want to reach sub-linear complexity, and we can't really reach log here, that's why we use square root
	jumpAmount := int(math.Floor(math.Sqrt(float64(len(floors)))))

	i := jumpAmount // start at index we jumped to
	for ; i < len(floors); i += jumpAmount {
		if floors[i] {
			break // means our first crystal ball has broken
		}
	}

	// so now we "jump" back one jump, and start dropping the crystal ball one floor at a time
	i -= jumpAmount
	// we could limit the amount of walking to "jumpAmount"
	// but we know for sure that it will break before reaching that anyway
	for ; i < len(floors); i++ {
		if floors[i] {
			return i, nil
		}
	}

	return -1, errors.New("balls never broke")
}
