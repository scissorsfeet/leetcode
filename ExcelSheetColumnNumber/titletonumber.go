package ExcelSheetColumnNumber
import (
    "strings"
    "math"
)

/*
Given a column title as appear in an Excel sheet, return its corresponding column number.
*/

func titleToNumber(s string) int {
    s = strings.ToUpper(s)
    bitNum := len(s)
    var ret int
    for i:=bitNum-1;i>=0;i-- {
        tmp := int(s[i])-64
        ret += tmp*int(math.Pow(float64(26),float64(bitNum-i-1)))
    }
    return ret
}