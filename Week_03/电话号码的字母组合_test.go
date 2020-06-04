package Week_03_test

import (
	"fmt"
	"testing"
)

// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/
//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
//
//给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

//示例:
//
//输入："23"
//输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].

func letterCombinations(digits string) []string {
	h_map := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	res := []string{}
	if len(digits) < 1 {
		return res
	}
	search2("", 0, digits, &res, h_map)
	return res
}

//s 每次递归的结果
//i 递归的层数
//res 最终返回结果
//digits 输入
//h_map 哈希映射表

func search2(s string, index int, digits string, res *[]string, h_map map[string]string) {

	if len(digits) == index {
		*res = append(*res, s)
		return
	}

	t_num := string(digits[index])
	letter := h_map[t_num]
	for i := 0; i < len(letter); i++ {
		search2(s+string(letter[i]), index+1, digits, res, h_map)
	}
}

func TestLetter(t *testing.T) {
	fmt.Println(letterCombinations("23"))
}

func search(s string, i int, digits string, res *[]string, h_map map[string]string) {
	if i == len(digits) {
		*res = append(*res, s)
		return
	}

	//获取本次递归的号码
	tmp_s := string(digits[i])
	//获取本次递归号码对应的字母
	letters := h_map[tmp_s]
	//循环号码对应的所有字母
	for j := 0; j < len(letters); j++ {
		search(s+string(letters[j]), i+1, digits, res, h_map)
	}
}
