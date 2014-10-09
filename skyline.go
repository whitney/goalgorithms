package main

import (
    //"fmt"
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

func skyline(rects []rect) {

}
