package Stack

import "testing"

func TestMergeIntervals(t *testing.T) {
	s :=  merge([]Interval{Interval{1,3}, Interval{2,4}, Interval{5,6}})
	t.Log(s)

	s =  merge([]Interval{Interval{1,4}, Interval{0,4}, Interval{5,6}})
	t.Log(s)
}