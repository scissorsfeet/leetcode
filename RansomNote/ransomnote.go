package RansomNote

/*
Given an arbitrary ransom note string and another string containing letters from all the magazines,

write a function that will return true if the ransom note can be constructed from the magazines ; otherwise, it will return false.

Each letter in the magazine string can only be used once in your ransom note.
*/

func canConstruct(ransomNote string, magazine string) bool {
    m := make(map[uint8]uint)
    for i, _ := range magazine{
        m[magazine[i]]++
    }
    ret := true
    for i, _ := range ransomNote{
        if _, ok := m[ransomNote[i]];ok{
            if m[ransomNote[i]] > 0{
                m[ransomNote[i]]--
            }else{
                ret = false
                break
            }
        }else{
            ret = false
            break
        }
    }
    return ret
}