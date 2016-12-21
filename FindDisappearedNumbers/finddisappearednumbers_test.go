package FindDisappearedNumbers
import "testing"

func TestA(t *testing.T) {
    s := []int{4,3,2,7,8,2,3,1}
    ret := findDisappearedNumbers(s)
    t.Log(ret)
}