package main

import (
    "testing"
)

func TestSMEdian(t *testing.T) {

    s := NewSMedian()

    if med := s.median(); med != 0 {
        t.Errorf("med expected: %f, actual: %f\n", 0, med)
    }

    s.add(9)
    if med := s.median(); med != 9 {
        t.Errorf("med expected: %f, actual: %f\n", 9, med)
    }
}
