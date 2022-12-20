package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type coord struct {
	x, y, z int
}

func dot(c1, c2 coord) int {
	return c1.x*c2.x + c1.y*c2.y + c1.z*c2.z
}

func main() {
	contents, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	cube := map[coord]bool{}
	for _, line := range strings.Split(string(contents), "\n") {
		xyz := strings.Split(line, ",")
		x, _ := strconv.Atoi(xyz[0])
		y, _ := strconv.Atoi(xyz[1])
		z, _ := strconv.Atoi(xyz[2])
		cube[coord{x, y, z}] = true
	}

	var sa int
	for key := range cube {
		sa += countEmptyNeighbors(key, cube)
	}

	fmt.Printf("Part One: Cube has an exposed surface area of %d\n", sa)

	// Part two:

	var xs, ys, zs []int
	for v := range cube {
		xs = append(xs, v.x)
		ys = append(ys, v.y)
		zs = append(zs, v.z)
	}
	sort.Ints(xs)
	sort.Ints(ys)
	sort.Ints(zs)

	xMin := xs[0]
	yMin := ys[0]
	zMin := zs[0]
	xMax := xs[len(xs)-1]
	yMax := ys[len(ys)-1]
	zMax := zs[len(zs)-1]

	cursor := coord{xMin - 2, yMin - 2, zMin - 2}
	for {
		cursor.x++
		if cube[cursor] {
			cursor.x--
			break
		}
		cursor.y++
		if cube[cursor] {
			cursor.y--
			break
		}
		cursor.z++
		if cube[cursor] {
			cursor.z--
			break
		}
	}

	outerEdge := map[coord]bool{}
	checked := map[coord]bool{}
	outerEdge[cursor] = true

	// breadth first search
	queue := unvisitedSurfaceNeighbors(cursor, checked, cube)
	for counter := 0; len(queue) > 0; counter++ {
		if counter%100000 == 0 {
			fmt.Println("current queue length:", len(queue))
		}
		head := queue[0]
		outerEdge[head] = true
		queue = append(queue, unvisitedSurfaceNeighbors(head, checked, cube)...)
		queue = queue[1:]
	}

	// Rendering 2d slices of the rock for
	// fun and debugging
	for z := zMin - 1; z <= zMax+1; z++ {
		fmt.Println("z=", z)
		for y := yMin - 1; y <= yMax+1; y++ {
			for x := xMin - 1; x <= xMax+1; x++ {
				if cube[coord{x, y, z}] {
					fmt.Print("#")
				} else if outerEdge[coord{x, y, z}] {
					fmt.Print("@")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
		fmt.Print("\n\n")
	}

	var ea int
	for key := range outerEdge {
		for _, n := range generateNeighbors(key) {
			if cube[n] {
				ea++
			}
		}
	}

	fmt.Println("size of the outer edge is ", ea)
}

func unvisitedSurfaceNeighbors(c coord, checked map[coord]bool, cube map[coord]bool) []coord {
	out := []coord{}
	for _, n := range generateNeighbors(c) {
		if !checked[n] && countEmptyNeighbors(n, cube) < 6 && !cube[n] {
			out = append(out, n)
			// we added it to the queue so it will have been visited
			// by the time we get back to this function
			checked[n] = true

		}
		if !cube[n] {
			// yeah this ended up being kind of ugly
			// we have to explore next nearest neighbors
			// but only through air, not through rock.
			for _, nn := range generateNeighbors(n) {
				if !checked[nn] && countEmptyNeighbors(nn, cube) < 6 && !cube[nn] {
					out = append(out, nn)
					checked[nn] = true
				}
			}
		}
	}
	checked[c] = true
	return out
}

func generateNeighbors(c coord) []coord {
	out := []coord{}
	for _, d := range []int{-1, 1} {
		out = append(out, coord{c.x + d, c.y, c.z})
		out = append(out, coord{c.x, c.y + d, c.z})
		out = append(out, coord{c.x, c.y, c.z + d})
	}
	return out
}

func countEmptyNeighbors(c coord, cube map[coord]bool) int {
	base := 6
	for _, dx := range []int{-1, 1} {
		if cube[coord{c.x + dx, c.y, c.z}] {
			base--
		}
	}
	for _, dy := range []int{-1, 1} {
		if cube[coord{c.x, c.y + dy, c.z}] {
			base--
		}
	}
	for _, dz := range []int{-1, 1} {
		if cube[coord{c.x, c.y, c.z + dz}] {
			base--
		}
	}
	return base
}

func countSignedEmptyNeighbors(c coord, cube map[coord]bool) int {
	signedSa := 0
	for _, dx := range []int{-1, 1} {
		dr := coord{dx, 0, 0}
		if !cube[coord{c.x + dx, c.y, c.z}] {
			signedSa += sgn(dot(c, dr))
		}
	}
	for _, dy := range []int{-1, 1} {
		dr := coord{0, dy, 0}
		if !cube[coord{c.x, c.y + dy, c.z}] {
			signedSa += sgn(dot(c, dr))
		}
	}
	for _, dz := range []int{-1, 1} {
		dr := coord{0, 0, dz}
		if !cube[coord{c.x, c.y, c.z + dz}] {
			signedSa += sgn(dot(c, dr))
		}
	}

	return signedSa
}

func sgn(i int) int {
	if i > 0 {
		return 1
	}
	return -1
}
