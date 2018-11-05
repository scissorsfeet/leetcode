package Sort

import "sort"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	bs := byteSlice(s)
	bt := byteSlice(t)
	sort.Sort(bt)
	sort.Sort(bs)

	for i:=0;i<len(bs);i++ {
		if bs[i] != bt[i] {
			return false
		}
	}
	return true
}

type byteSlice []byte
func (p byteSlice) Len() int           { return len(p) }
func (p byteSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p byteSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
