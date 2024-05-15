package main

import "fmt"

// 20. 有效的括号: https://leetcode.cn/problems/valid-parentheses/description/
//
// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
// 有效字符串需满足：
//   - 左括号必须用相同类型的右括号闭合。
//   - 左括号必须以正确的顺序闭合。
//   - 每个右括号都有一个对应的相同类型的左括号。
func isValid(s string) bool {
	if len(s)%2 > 0 {
		return false
	}
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	var stack []rune
	for _, v := range s {
		switch v {
		case '(', '{', '[':
			stack = append(stack, v)
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[v] {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
			return false
		}
	}
	return len(stack) == 0
}

func main() {
	s := "()"
	fmt.Println(isValid(s)) // expect: true
	s = "()[]{}"
	fmt.Println(isValid(s)) // expect: true
	s = "(]"
	fmt.Println(isValid(s)) // expect: false
	s = "){"
	fmt.Println(isValid(s)) // expect: false
}
