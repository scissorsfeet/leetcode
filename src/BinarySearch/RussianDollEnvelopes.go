package BinarySearch

import (
	"sort"
)

func maxEnvelopes(envelopes [][]int) int {
	sort.Sort(EnvelopeSlice(envelopes))

	var res = 0
	for i:=0;i<len(envelopes);{
		r := searchFirstGT(envelopes[i+1:], envelopes[i])
		res++
		i += r
	}
	return res
}

func searchFirstGT(envelopes [][]int, v []int) int {
	low := 0
	high := len(envelopes) - 1

	for low <= high {
		mid := (low + high) >> 1
		if envelopes[mid][0] > v[0] && envelopes[mid][1] > v[1] {
			if mid == 0 {
				if envelopes[mid][0] < v[0] || envelopes[mid][1] < v[0] {
					return 2
				}
				return 1
			}
			if envelopes[mid-1][0] < v[0] || envelopes[mid-1][1] < v[0] {
				return mid
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return low + 1
}

type EnvelopeSlice [][]int

func (p EnvelopeSlice) Len() int           { return len(p) }
func (p EnvelopeSlice) Less(i, j int) bool {
	if p[i][0] == p[j][0] {
		return p[i][1] < p[j][1]
	}
	return p[i][0] < p[j][0]
}
func (p EnvelopeSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }