package Stack

import "sort"

/**
* Definition for an interval.
*/
type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	sort.Sort(IntervalSlice(intervals))
	res := make([]Interval, 0, len(intervals))
	for i:=0;i<len(intervals);i++ {
		if i == 0 {
			res = append(res, intervals[i])
		} else {
			as := &res[len(res)-1]
			bs := &intervals[i]
			if as.Start <= bs.Start && as.End >= bs.Start {
				var newInterval Interval
				if as.End > bs.End {
					newInterval = Interval{Start: as.Start, End: as.End}
				} else {
					newInterval = Interval{Start: as.Start, End: bs.End}
				}
				res[len(res)-1] = newInterval
			} else {
				res = append(res, intervals[i])
			}
		}
	}
	return res
}

type IntervalSlice []Interval

func (p IntervalSlice) Len() int           { return len(p) }
func (p IntervalSlice) Less(i, j int) bool { return p[i].Start < p[j].Start }
func (p IntervalSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }