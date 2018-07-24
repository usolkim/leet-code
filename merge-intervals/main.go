package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(merge([]Interval{Interval{1, 3}, Interval{2, 6}, Interval{8, 10}, Interval{15, 18}}))
	fmt.Println(merge([]Interval{Interval{1, 4}, Interval{4, 5}}))
	fmt.Println(merge([]Interval{}))
	fmt.Println(merge([]Interval{Interval{1, 4}}))
	fmt.Println(merge([]Interval{Interval{4, 5}, Interval{1, 4}, Interval{2, 3}}))
}

// Interval Definition for an interval.
type Interval struct {
	Start int
	End   int
}

// ByInterval Sort by Interval
type ByInterval []Interval

func (a ByInterval) Len() int      { return len(a) }
func (a ByInterval) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByInterval) Less(i, j int) bool {
	return a[i].Start < a[j].Start || (a[i].Start == a[j].Start && a[i].End < a[j].End)
}

func merge(intervals []Interval) []Interval {
	l := len(intervals)
	if l < 2 {
		return intervals
	}
	sort.Sort(ByInterval(intervals))

	res := make([]Interval, 0)
	s := intervals[0].Start
	e := intervals[0].End
	for i := 0; i < l-1; i++ {
		if intervals[i+1].Start <= e {
			if e < intervals[i+1].End {
				e = intervals[i+1].End
			}
		} else {
			res = append(res, Interval{Start: s, End: e})
			s = intervals[i+1].Start
			e = intervals[i+1].End
		}
	}
	res = append(res, Interval{Start: s, End: e})

	return res
}
