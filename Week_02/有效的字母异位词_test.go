package Week_02_test

import (
	"fmt"
	"testing"
)

// https://leetcode-cn.com/problems/valid-anagram/

//【题意】
//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//示例 1:
//
//输入: s = "anagram", t = "nagaram"
//输出: true

//【解题方式】
//1、哈希表，先把字符串s压如哈希表中，在循环t字符串，如果哈希表中的的每项值都为0则返回true

func isAnagram(s string, t string) bool {

	s_hash := map[string]int{}
	for _, v := range s {
		if _, ok := s_hash[string(v)]; ok {
			s_hash[string(v)] += 1
		} else {
			s_hash[string(v)] = 1
		}
	}
	for _, v := range t {
		if _, ok := s_hash[string(v)]; ok {
			s_hash[string(v)] -= 1
			if s_hash[string(v)] < 0 {
				fmt.Println(2)
				return false
			}
		} else {
			fmt.Println(1)
			return false
		}
	}
	for _, v := range s_hash {
		if v > 0 {
			fmt.Println(3)
			return false
		}
	}
	return true
}

func isAnagramS(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	s_hash := map[string]int{}
	for k, v := range s {
		s_hash[string(v)] += 1
		s_hash[string(t[k])] -= 1
	}
	fmt.Println(s_hash)
	for _, v := range s_hash {
		if v > 0 || v < 0 {
			return false
		}
	}
	return true
}

func TestIsAnagram(t *testing.T) {
	s := "anagram"
	tt := "nagaram"
	s = "a"
	tt = "b"
	fmt.Println(isAnagramS(s, tt))
}
