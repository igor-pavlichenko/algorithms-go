package recursion

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertAt(t *testing.T) {
	assert := assert.New(t)
	maze := []string{
		"xxxxxxxxxx x",
		"x        x x",
		"x        x x",
		"x xxxxxxxx x",
		"x          x",
		"x xxxxxxxxxx",
	}
	expectedResult := []Point{
		{x: 10, y: 0},
		{x: 10, y: 1},
		{x: 10, y: 2},
		{x: 10, y: 3},
		{x: 10, y: 4},
		{x: 9, y: 4},
		{x: 8, y: 4},
		{x: 7, y: 4},
		{x: 6, y: 4},
		{x: 5, y: 4},
		{x: 4, y: 4},
		{x: 3, y: 4},
		{x: 2, y: 4},
		{x: 1, y: 4},
		{x: 1, y: 5},
	}

	actualResult := Solve(maze, "x", Point{x: 10, y: 0}, Point{x: 1, y: 5})

	assert.Equal(drawPath(maze, expectedResult), drawPath(maze, actualResult))

}

func drawPath(maze []string, path []Point) []string {
	var pathCanvas [][]string
	for _, row := range maze {
		characters := strings.Split(row, "")
		pathCanvas = append(pathCanvas, characters)
	}

	for _, point := range path {
		pathCanvas[point.y][point.x] = "*"
	}

	var pathDrawing []string
	for _, elem := range pathCanvas {
		pathDrawing = append(pathDrawing, strings.Join(elem, ""))
	}

	return pathDrawing
}
