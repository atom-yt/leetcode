package main

import (
	"fmt"
)
/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长 
子串
 的长度。

 

示例 1:

输入: s = "abcabcbb"
输出: 3 
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
     
解题思路：

设置start标志位和index标志位，index从前往后依次遍历
每次出现的字母，都缓存其最后一次出现的索引
如果index指向的字母的lastLndex在start前面，也就是小于start
如果当前字母的的lastIndex大于或者等于start，则start标志位移动到index的前面一位
每次遍历都更新lastIndex和maxLength
*/
func lengthOfLongestSubstring(s string) int {
    rus := []rune(s)
    if len(rus) == 0 {
        return 0
    }

    charIndex := make(map[rune]int)
    var start, maxLen, length int

    for i, v := range rus {
        last, ok := charIndex[v]
        if ok && last >= start {
            start = last + 1
        }
        charIndex[v] = i

        length = i - start + 1
        if maxLen < length {
            maxLen = length
        }
    }
    return maxLen
}


func reverseStr(s string) string {
	sr := []rune(s)
	slen := len(sr)
	for i:=0; i<slen/2;i++ {
		sr[i], sr[slen-i-1] = sr[slen-i-1], sr[i]
	}
	return string(sr)
}

func main() {
	// fmt.Println(reverseStr("abcdf"))
	fmt.Println( lengthOfLongestSubstring("aaaaabcdaabcdeeee"))
}
