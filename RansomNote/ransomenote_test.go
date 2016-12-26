package RansomNote
import "testing"

func TestA(t *testing.T){
    ret := canConstruct("aaa", "aab")
    if ret {
        t.Error("wrong")
    }
}

func TestB(t *testing.T){
    ret := canConstruct("aaa", "aaab")
    if !ret {
        t.Error("wrong")
    }
}

func TestC(t *testing.T){
    ret := canConstruct("cdae", "aaaccclasdjflsjde")
    if !ret {
        t.Error("wrong")
    }
}

func TestD(t *testing.T){
    ret := canConstruct("dewsssx", "de")
    if ret {
        t.Error("wrong")
    }
}