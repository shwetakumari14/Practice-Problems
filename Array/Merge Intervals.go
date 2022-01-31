package main

import "fmt"

type Interval struct {
	start, end int
}

func main() {
	arr := []Interval{{1, 3}, {6, 9}}
	var newInterval Interval
	newInterval.start = 2
	newInterval.end = 6
	result := insertInterval(arr, newInterval)
	fmt.Println(result)
}

func insertInterval(intervals []Interval, newInterval Interval) []Interval {
	var ans []Interval
	var tempInterval Interval

	if newInterval.end < newInterval.start {
		tempInterval.start = newInterval.end
		tempInterval.end = newInterval.start
	} else {
		tempInterval = newInterval
	}

	count, min, max := 0, -1, -1

	for i := 0; i < len(intervals); i++ {
		if isOverlapping(intervals[i], tempInterval) {
			if count == 0 {
				min = i
				count++
			}
			max = i
		}
	}

	if min != -1 {
		count := 0
		for i := 0; i < len(intervals); i++ {
			if i < min || i > max {
				ans = append(ans, intervals[i])
			} else if count == 0 {
				var mergedInterval Interval
				mergedInterval.start = minInterval(intervals[min], tempInterval)
				mergedInterval.end = maxInterval(intervals[max], tempInterval)
				ans = append(ans, mergedInterval)
				count++
			}
		}
	} else {
		count := 0
		for i := 0; i < len(intervals); i++ {
			if intervals[i].start > tempInterval.start {
				if count == 0 {
					ans = append(ans, tempInterval)
					count++
				}
			}
			ans = append(ans, intervals[i])
		}
		if count == 0 {
			ans = append(ans, tempInterval)
		}
	}

	return ans
}

func isOverlapping(interval1, interval2 Interval) bool {
	if interval1.end > interval2.start && interval2.start > interval1.start {
		return true
	}
	if interval2.end > interval1.start && interval2.end < interval1.start {
		return true
	}

	return false
}

func maxInterval(a, b Interval) int {
	if a.end > b.end {
		return a.end
	}

	return b.end
}

func minInterval(a, b Interval) int {
	if a.start < b.start {
		return a.start
	}

	return b.start
}
