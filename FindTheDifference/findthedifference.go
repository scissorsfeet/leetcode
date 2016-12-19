package FindTheDifference

/*
Given two strings s and t which consist of only lowercase letters.
给两个小写字符串
String t is generated by random shuffling string s and then add one more letter at a random position.
t是由s打乱顺序并增加了一个字符而来
Find the letter that was added in t.
找到增加的这个字符
*/
func findthedifference(s, t string) (differ string) {
    m := make(map[rune]int)

    for _, c := range s {
        m[c]++
    }

    for _, c := range t {
        if m[c] > 0 {
            m[c]--
        }else{
            differ = string(c)
            break
        }
    }

    return
}