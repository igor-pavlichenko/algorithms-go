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

	var path []Point = make([]Point, 0)

	return path
}
