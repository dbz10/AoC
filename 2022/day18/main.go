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

	cursor := coord{xMin - 2, yMin - 2, zMin - 2}
	for {
		cursor.x++
		if cube[cursor] {
			break
		}
		cursor.y++
		if cube[cursor] {
			break
		}
		cursor.z++
		if cube[cursor] {
			break
		}
	}
	fmt.Println(cursor)

	outerEdge := map[coord]bool{}
	checked := map[coord]bool{}
	outerEdge[cursor] = true
	// breadth first search
	queue := unvisitedSurfaceNeighbors(cursor, checked, cube)
	fmt.Println(queue)

	for len(queue) > 0 {
		// time.Sleep(500 * time.Millisecond)
		head := queue[0]
		outerEdge[head] = true
		queue = append(queue, unvisitedSurfaceNeighbors(head, checked, cube)...)
		queue = queue[1:]
		fmt.Println(queue)
	}

	fmt.Println(outerEdge)
	fmt.Println(len(outerEdge))

	var ea int
	for key := range outerEdge {
		ea += countEmptyNeighbors(key, cube)
	}

	fmt.Println("size of the outer edge is ", ea)
}

func unvisitedSurfaceNeighbors(c coord, checked map[coord]bool, cube map[coord]bool) []coord {
	out := []coord{}
	allNeighbors := generateNeighbors(c)
	for _, n := range allNeighbors {
		if !checked[n] && countEmptyNeighbors(n, cube) > 0 && cube[n] {
			checked[n] = true
			out = append(out, n)
		}
	}
	checked[c] = true
	return out
}

func generateNeighbors(c coord) []coord {
	out := []coord{}
	for _, dx := range []int{-1, 0, 1} {
		for _, dy := range []int{-1, 0, 1} {
			for _, dz := range []int{-1, 0, 1} {
				out = append(out, coord{c.x + dx, c.y + dy, c.z + dz})

			}
		}
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
