package main

import (
    "testing"
)

func TestTopology(t *testing.T) {
    a0 := []int{1, 1, 3, 2, 1}
    a1 := []int{2, 1, 2}
    a2 := []int{3, 2, 1, 2, 3, 2}
    a3 := []int{3, 2, 1, 2, 3, 1, 3}

    if t0 := topology(a0); t0 != 0 {
        t.Errorf("t0 expected: %f, actual: %f\n", 0, t0)
    }

    if t1 := topology(a1); t1 != 1 {
        t.Errorf("t1 expected: %f, actual: %f\n", 1, t1)
    }

    if t2 := topology(a2); t2 != 4 {
        t.Errorf("t2 expected: %f, actual: %f\n", 4, t2)
    }

    if t3 := topology(a3); t3 != 6 {
        t.Errorf("t3 expected: %f, actual: %f\n", 6, t3)
    }

    /*
    fmt.Printf("%v => %f\n", a0, solve(a0))
    fmt.Printf("%v => %f\n", a1, solve(a1))
    fmt.Printf("%v => %f\n", a2, solve(a2))
    fmt.Printf("%v => %f\n", a3, solve(a3))
    */
}
