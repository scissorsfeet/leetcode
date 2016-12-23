package MinimumMovesToEqualArrayElements
import "testing"

func TestA(t *testing.T){
    arr := []int{1,2,3}
    ret := minMoves(arr)
    if ret != 3{
        t.Error(ret)
    }
}
