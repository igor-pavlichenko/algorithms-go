package recursion

type Point struct {
	x int
	y int
}

/*
example maze:

	      ⬇ end here
	"#####  #"
	"#      #"
	"#  #####"
	  ⬆ start here

use recursion... define base cases:
 1. off the map
 2. it's a wall
 3. it's the end
 4. if we have seen it

remember that the entrance to the recursive function is often not the one we are going to recurse in

when to use recursion: when it's not easy to do with simple for loops
especially when there is a branching factor, like in the coordinates case,
there are 4 directions we can go at any point

@returns list of points from the start to the end
*/
func Solve(maze []string, wall string, start Point, end Point) []Point {
	var seen [][]bool = make([][]bool, len(maze))
	var path []Point = make([]Point, 0)

	for i := range seen {
		seen[i] = make([]bool, len(maze[0]))
	}

	walk(maze, wall, start, end, seen, &path)

	return path
}

func walk(
	maze []string,
	wall string,
	curr Point,
	end Point,
	seen [][]bool,
	path *[]Point, // we need to keep track of the steps we are taking
	// ps: path has to be a pointer otherwise the context is lost in every walk
	// not sure it doesn't complain about seen array tho
	// probably because we are changing values inside of the slice,
	// but not modifying the slice itself
) bool {

	// 1. Base case
	// off the map
	if curr.x < 0 || curr.x >= len(maze[0]) || curr.y < 0 || curr.y >= len(maze) {
		return false
	}

	// it's a wall
	if string(maze[curr.y][curr.x]) == wall {
		return false
	}

	// it's the end
	if curr.x == end.x && curr.y == end.y {
		// pre
		*path = append(*path, end)
		return true
	}

	// we have seen it
	// we could mutate the point to add a flag but that would be weird
	// instead, we'll pass around a 2D array with all coordinates we visited, true if we've seen it, false if we didn't
	if seen[curr.y][curr.x] {
		return false
	}

	// 2. Recurse - 3 steps
	// pre
	seen[curr.y][curr.x] = true
	*path = append(*path, curr)
	// recurse
	for _, dir := range directions {
		x := dir[0]
		y := dir[1]
		if walk(maze, wall, Point{x: curr.x + x, y: curr.y + y}, end, seen, path) {
			return true
		}
	}
	// post
	*path = (*path)[:len(*path)-1]
	return false
}

var directions = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}
