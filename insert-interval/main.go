package main

import "fmt"

func main() {
	fmt.Println(insert([]Interval{Interval{1, 5}, Interval{6, 8}}, Interval{0, 9}))
	fmt.Println(insert([]Interval{Interval{1, 3}, Interval{6, 9}}, Interval{2, 5}))
	fmt.Println(insert([]Interval{Interval{1, 2}, Interval{3, 5}, Interval{6, 7}, Interval{8, 10}, Interval{12, 16}}, Interval{4, 8}))
	fmt.Println(insert([]Interval{}, Interval{4, 8}))
	fmt.Println(insert([]Interval{Interval{1, 3}}, Interval{4, 8}))
	fmt.Println(insert([]Interval{Interval{1, 4}}, Interval{4, 8}))
	fmt.Println(insert([]Interval{Interval{1, 8}}, Interval{4, 8}))
	fmt.Println(insert([]Interval{Interval{1, 10}}, Interval{4, 8}))
}

// Interval Definition for an interval.
type Interval struct {
	Start int
	End   int
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	l := len(intervals)
	if l == 0 {
		return []Interval{newInterval}
	}
	res := make([]Interval, 0)

	begin := bsearch(intervals, 0, l-1, newInterval.Start)
	end := bsearch(intervals, 0, l-1, newInterval.End)

	ni := Interval{}

	ni.Start = newInterval.Start
	ni.End = newInterval.End

	if intervals[begin].End < newInterval.Start {
		begin++
	}
	if begin < l && ni.Start > intervals[begin].Start {
		ni.Start = intervals[begin].Start
	}

	if intervals[end].Start <= newInterval.End {
		end++
	}
	if end > 0 && ni.End < intervals[end-1].End {
		ni.End = intervals[end-1].End
	}

	res = append(res, intervals[:begin]...)
	res = append(res, ni)
	res = append(res, intervals[end:]...)

	return res
}

func bsearch(intervals []Interval, s, e, target int) int {
	if s >= e {
		return s
	}
	i := (s + e) / 2
	ii := intervals[i]

	if ii.Start <= target && target <= ii.End {
		return i
	} else if ii.Start > target {
		return bsearch(intervals, s, i-1, target)
	} else {
		return bsearch(intervals, i+1, e, target)
	}
}
