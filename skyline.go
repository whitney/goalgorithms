package main

import (
    //"fmt"
    "sort"
)

/*
You are given a set of n rectangles in no particular order. They have varying widths and heights, 
but their bottom edges are collinear, so that they look like buildings on a skyline. For each 
rectangle, you’re given the x position of the left edge, the x position of the right edge, and 
the height. Your task is to draw an outline around the set of rectangles so that you can see what 
the skyline would look like when silhouetted at night.

See: https://briangordon.github.io/2014/08/the-skyline-problem.html 
*/

type rect struct {
    left, right, height int
}

type ByLeftCoord []rect

func (r ByLeftCoord) Len() int {
    return len(r)
}

func (r ByLeftCoord) Swap(i, j int) {
    r[i], r[j] = r[j], r[i]
}

func (r ByLeftCoord) Less(i, j int) bool {
    return r[i].left < r[j].left
}

func (r rect) intersects(s rect) bool {
    return false
}

// represents a vertical or horizontal
// line segment (all segments are in quardant I 
// of the Cartesian plane, ie all indices are >= 0)
type seg struct {
    x, y, x1, y1 int
}

// "brute force" O(n!) solution
func skyline0(rects []rect) []seg {
    var segs []seg

    sort.Sort(ByLeftCoord(rects))

    // for a []rect of [A, B, C], we look
    // for the intersections of the rects 
    // AB, AC, BC
    for i := 0; i < len(rects) - 1; i++ {
        for j := i + 1; j < len(rects); j++ {
            if (rects[i].intersects(rects[j])) {

            }
        }
    }

    

    return segs
}
