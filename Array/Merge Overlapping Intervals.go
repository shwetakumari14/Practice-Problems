package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	start, end int
}

func main() {
	arr := []Interval{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	//arr := []Interval{{1, 10}, {2, 9}, {3, 8}, {4, 7}, {5, 6}, {6, 6}}
	result := merge(arr)
	fmt.Println(result)
}

func merge(A []Interval) []Interval {
	var ans []Interval

	sort.Slice(A, func(i, j int) bool {
		return A[i].start < A[j].start
	})

	for i := 0; i < len(A); i++ {
		if i+1 < len(A) && isOverlapping(A[i], A[i+1]) {
			A[i+1] = mergeInterval(A[i], A[i+1])
		} else {
			ans = append(ans, A[i])
		}
	}

	return ans
}

func isOverlapping(interval1, interval2 Interval) bool {

	if interval1.end >= interval2.start && interval2.start >= interval1.start {
		return true
	}
	if interval2.end >= interval1.start && interval2.end <= interval1.start {
		return true
	}

	return false
}

func mergeInterval(interval1, interval2 Interval) Interval {
	var newInterval Interval
	newInterval.start = min(interval1.start, interval2.start)
	newInterval.end = max(interval1.end, interval2.end)

	return newInterval
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
