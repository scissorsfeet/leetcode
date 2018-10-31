package Sort

import "testing"

func TestSortColors(t *testing.T) {
	a := []int{2,2,2,1,1,1,0}
	sortColors(a)
	t.Log(a)

	a = []int{2,0,2,1,1,0}
	sortColors(a)
	t.Log(a)

	a = []int{2,0}
	sortColors(a)
	t.Log(a)

	a = []int{2,1}
	sortColors(a)
	t.Log(a)

	a = []int{0,0,0}
	sortColors(a)
	t.Log(a)

	a = []int{1,1,1}
	sortColors(a)
	t.Log(a)

	a = []int{2,2,2}
	sortColors(a)
	t.Log(a)

	a = []int{0,0,1}
	sortColors(a)
	t.Log(a)

	a = []int{2,0,0,0,0,1,2,2}
	sortColors(a)
	t.Log(a)
}