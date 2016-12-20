package HammingDistance
import "testing"

func TestA(t *testing.T) {
    ret := hammingdistance(1, 4)
    if ret != 2{
        t.Error("wrong")
    }
}
