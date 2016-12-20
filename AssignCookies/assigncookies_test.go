package AssignCookies
import (
    "testing"
    "fmt"
)

func TestA(t *testing.T){
    g, s := []int{1,2,3}, []int{1,1}
    ret := findContentChildren(g, s)
    if ret != 1 {
        t.Error(fmt.Sprintf("wrong:%d", ret))
    }
}

func TestB(t *testing.T){
    g, s := []int{1,2}, []int{1,2,3}
    ret := findContentChildren(g, s)
    if ret != 2 {
        t.Error(fmt.Sprintf("wrong:%d", ret))
    }
}
