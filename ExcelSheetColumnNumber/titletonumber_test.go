package ExcelSheetColumnNumber
import "testing"

func TestA(t *testing.T) {
    ret := titleToNumber("AB")
    if ret != 28{
        t.Error(ret)
    }
}

func TestB(t *testing.T) {
    ret := titleToNumber("C")
    if ret != 3{
        t.Error(ret)
    }
}
