package main

import (
    "container/heap"
    "errors"
    "fmt"
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

func (s *smedian) print() {
    fmt.Println("lHeap:")
    for s.lHeap.Len() > 0 {
        elem := heap.Pop(s.lHeap)
        fmt.Printf("%d ", elem)
    } 

    fmt.Println("\nrHeap:")
    for s.rHeap.Len() > 0 {
        elem := heap.Pop(s.rHeap)
        fmt.Printf("%d ", elem)
    } 
    fmt.Println("")
}

func (s *smedian) median() float64 {
    if s.lHeap.Len() == 0 && s.rHeap.Len() == 0 {
        return 0
    }

    if s.lHeap.Len() < s.rHeap.Len() {
        med := heap.Pop(s.rHeap).(int)
        heap.Push(s.rHeap, med)
        return float64(med)
    } else if s.lHeap.Len() > s.rHeap.Len() {
        med := heap.Pop(s.lHeap).(int)
        heap.Push(s.lHeap, med)
        return float64(med)
    } else {
        lMed := heap.Pop(s.lHeap).(int)
        heap.Push(s.lHeap, lMed)
        rMed := heap.Pop(s.rHeap).(int)
        heap.Push(s.rHeap, rMed)
        return float64(lMed + rMed) / 2
    }
}

func (s *smedian) add(elem int) {

    // terms:
    // - lMax: max elem in lHeap
    // - rMin: min elem in rHeap 
    //
    // cases:

    
    if s.lHeap.Len() == 0 {
        heap.Push(s.lHeap, elem)
        return
    }

    lMax := heap.Pop(s.lHeap).(int)

    if s.rHeap.Len() == 0 {
        if elem >= lMax {
            heap.Push(s.rHeap, elem)
            heap.Push(s.lHeap, lMax)
        } else {
            heap.Push(s.lHeap, elem)
            heap.Push(s.rHeap, lMax)
        }

        return
    }

    rMin := heap.Pop(s.rHeap).(int)

    if s.lHeap.Len() < s.rHeap.Len() {
        if elem <= rMin {
            heap.Push(s.lHeap, elem)
            heap.Push(s.rHeap, rMin)
        } else {
            heap.Push(s.rHeap, elem)
            heap.Push(s.lHeap, rMin)
        }

        heap.Push(s.lHeap, lMax)
    } else if s.lHeap.Len() > s.rHeap.Len() {
        if elem >= lMax {
            heap.Push(s.rHeap, elem)
            heap.Push(s.lHeap, lMax)
        } else {
            heap.Push(s.lHeap, elem)
            heap.Push(s.rHeap, lMax)
        }

        heap.Push(s.rHeap, rMin)
    } else {
        // same size heaps
        if elem <= lMax {
            heap.Push(s.lHeap, elem)
        } else {
            heap.Push(s.rHeap, elem)
        }

        heap.Push(s.lHeap, lMax)
        heap.Push(s.rHeap, rMin)
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
