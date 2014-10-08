package main

import (
    "fmt"
    "math"
)

/*
Given an array of numbers (some candidates find it easier if they are integers, 
but the general solution should work for floating points as well), assume these 
are altitudes -- that the array forms a 2D topology. For example, if the input 
is [1,2,3,2,3,1], the topology would look as follows:

     _   _
   _| |_| |
 _|       |_
|           |

The question is to compute how much water would be retained within this structure 
after a rain. In the above example, the answer should be 1. A few more examples:

[1, 1, 3, 2, 1] => 0
[2, 1, 2] => 1
[3, 2, 1, 2, 3, 2] => 4
[3, 2, 1, 2, 3, 1, 3] => 6

For the simplest solution, the key insight is that the potential water level at any 
index is the minimum of (height of highest point to the left) and (height of highest 
point to the right). If this is less than or equal to the height at the index, there 
is no standing water at the index; if it is greater, then (the height of the water 
at that index - the height of land at the index) units of water stand at that index. 
By summing this over the array, the total volume can be found. A simple two-pass 
solution can compute (highest point to the left of every index) and (highest point 
to the right of every index), and from there the water held at each index and the 
total water, in O(n) runtime with O(n) extra space.

Optimal solution has a linear time compexity.
*/

func solve(a []float64) float64 {
    // observation: the amount of water held above any given index
    // of the input array depends on the max height to the left of the index 
    // and the max height to the right of the index.
    // Preprocess two arrays l and r such that l[i] is the max value to the left
    // of i in a, and r[i] is the max value to the right of i in a.

    l := make([]float64, len(a))
    l[0] = 0
    for i := 1; i < len(a); i++ {
        l[i] = math.Max(a[i-1], l[i-1]) 
    }

    r := make([]float64, len(a))
    r[len(a)-1] = 0
    for j := len(a) - 2; j >= 0; j-- {
        r[j] = math.Max(a[j+1], r[j+1])
    } 

    //fmt.Printf("a: %v, l: %v, r: %v\n", a, l, r)

    // Finally, the total amount of water the topology can hold is equal to the sum of 
    // the amounts held at each index of a. Furthermore the amount of water held at a 
    // given index i is equal to the lower of the two global maximums (to the left and to the right), 
    // minus the value of a[i]:
    var w float64 = 0
    for k := 0; k < len(a); k++ {
        wi := math.Min(l[k], r[k]) - a[k]
        if wi > 0.0 {
            w += wi
        }
    }

    return float64(w)
}

func main() {
    a0 := []float64{1, 1, 3, 2, 1}
    a1 := []float64{2, 1, 2}
    a2 := []float64{3, 2, 1, 2, 3, 2}
    a3 := []float64{3, 2, 1, 2, 3, 1, 3}

    fmt.Printf("%v => %f\n", a0, solve(a0))
    fmt.Printf("%v => %f\n", a1, solve(a1))
    fmt.Printf("%v => %f\n", a2, solve(a2))
    fmt.Printf("%v => %f\n", a3, solve(a3))
}
