package Week_02_test

import (
    "fmt"
    "testing"
)

func isAnagram(s string, t string) bool {

    map_ := map[string]int{}
    for _,v := range s {
        s_v := string(v)
        if _,ok := map_[s_v]; !ok {
            map_[s_v] = 1
        } else {
            map_[s_v] ++
        }
    }

    for _,v := range t {
        s_v := string(v)
        if _, ok := map_[s_v]; ok {
            map_[s_v] --
        } else {
            return  false
        }
    }
    for _, v := range map_ {
        if v > 0 {
            return  false
        }
    }
    return true
}

func TestIsAnagram(t *testing.T)  {
    s := "anagram"
    tt := "nagaram"
    fmt.Println(isAnagram(s, tt))
}