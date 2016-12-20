package AssignCookies
import "sort"

/*
Assume you are an awesome parent and want to give your children some cookies.
假设你想要给你的孩子们一些饼干
But, you should give each child at most one cookie.
但是，你只能给每个孩子最多一个饼干
Each child i has a greed factor gi, which is the minimum size of a cookie that the child will be content with; and each cookie j has a size sj.
孩子i的贪婪因素值为gi，代表着满足孩子的最小数量。sj代表着饼干的大小。
If sj >= gi, we can assign the cookie j to the child i, and the child i will be content.
如果sj大于等于gi，就把饼干j给孩子i，孩子i奖被满足
Your goal is to maximize the number of your content children and output the maximum number.
你的目标获取到最大满足的孩子数量
You may assume the greed factor is always positive.
可以认为贪婪因数都是正数
You cannot assign more than one cookie to one child.
最多只能给单个孩子一个饼干
*/

func findContentChildren(g []int, s []int) int {
    sort.Sort(sort.IntSlice(g))
    sort.Sort(sort.IntSlice(s))

    gLen, sLen := len(g), len(s)

    ret := 0
    for i, j:=0, 0;i<gLen && j<sLen;{
        if s[j] >= g[i] {
            ret++
            i++
            j++
        }else{
            j++
        }
    }
    return ret
}
