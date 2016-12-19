package FindTheDifference
import "testing"

func TestA(t *testing.T) {
    s := "abcd"
    q := "abecd"

    differ := findthedifference(s, q)
    if differ != "e"{
        t.Error("wrong")
    }else{
        t.Log("succ")
    }
}

func TestB(t *testing.T) {
    s := ""
    q := "a"

    differ := findthedifference(s, q)
    if differ != "a"{
        t.Error("fail")
    }else{
        t.Log("succ")
    }
}