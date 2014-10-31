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

    s.add(3)
    if med := s.median(); med != 6 {
        t.Errorf("med expected: %f, actual: %f\n", 6, med)
    }

    s.add(3)
    if med := s.median(); med != 3 {
        t.Errorf("med expected: %f, actual: %f\n", 3, med)
    }

    s.add(6)
    if med := s.median(); med != 4.5 {
        t.Errorf("med expected: %f, actual: %f\n", 4.5, med)
    }

    s.add(12)
    if med := s.median(); med != 6 {
        t.Errorf("med expected: %f, actual: %f\n", 6, med)
    }

}
