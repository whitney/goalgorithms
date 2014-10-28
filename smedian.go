package main

import (
    "container/heap"
    "errors"
    "log"
)

/*
Data structure to store a set with median
*/

// smedian mainatins:
// a MaxHeap maxH of elems and a 
// a MinHeap minH of elems which satisfy:
// - all elems in maxH are <= all elems in minH
// - the the cardinality of minH and maxH can 
//   differ by at most one.     
type smedian struct {
    lHeap *MaxHeap
    rHeap *MinHeap
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
    // Push and Pop use pointer receivers because they modify the slice's length,
    // not just its contents.
    *h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

type MaxHeap struct {
    MinHeap
}

func (h MaxHeap) Less(i, j int) bool { return h.MinHeap[i] > h.MinHeap[j] }

func NewSMedian() *smedian {
    lHeap := &MaxHeap{[]int{}}
    heap.Init(lHeap)

    rHeap := &MinHeap{}
    heap.Init(rHeap)

    return &smedian{lHeap: lHeap, rHeap: rHeap}
}

func (s *smedian) median() float64 {
    if s.lHeap.Len() == 0 && s.rHeap.Len() == 0 {
        return 0
    }

    if s.lHeap.Len() < s.rHeap.Len() {
        med := s.rHeap.Pop()
        s.rHeap.Push(med)
        return med.(float64)
    } else if s.lHeap.Len() > s.rHeap.Len() {
        med := s.lHeap.Pop()
        s.lHeap.Push(med)
        return med.(float64)
    } else {
        lMed := s.lHeap.Pop()
        s.lHeap.Push(lMed)
        rMed := s.rHeap.Pop()
        s.rHeap.Push(rMed)
        return (lMed.(float64) + rMed.(float64)) / 2
    }
}

func (s *smedian) add(elem int) {

    // terms:
    // - lMax: max elem in lHeap
    // - rMin: min elem in rHeap 
    //
    // cases:

    lMax := s.lHeap.Pop().(int)
    rMin := s.rHeap.Pop().(int)

    if s.lHeap.Len() < s.rHeap.Len() {
        if elem <= rMin {
            s.lHeap.Push(elem)
            s.rHeap.Push(rMin)
        } else {
            s.rHeap.Push(elem)
            s.lHeap.Push(rMin)
        }

        s.lHeap.Push(lMax)
    } else if s.lHeap.Len() > s.rHeap.Len() {
        if elem >= lMax {
            s.rHeap.Push(elem)
            s.lHeap.Push(lMax)
        } else {
            s.lHeap.Push(elem)
            s.rHeap.Push(lMax)
        }

        s.rHeap.Push(rMin)
    } else {
        // same size heaps
        if elem <= lMax {
            s.lHeap.Push(elem)
        } else {
            s.rHeap.Push(elem)
        }

        s.lHeap.Push(lMax)
        s.rHeap.Push(rMin)
    }


    // assert heap size invariant
    if !heapInvarient(s.lHeap, s.rHeap) {
        log.Fatal(errors.New("invarient violated: heap size mis-match"))
    }
}

func heapInvarient(maxH *MaxHeap, minH *MinHeap) bool {
    if maxH.Len() >= minH.Len() {
        return maxH.Len() - minH.Len() <= 1
    }

    return minH.Len() - maxH.Len() == 1
}
