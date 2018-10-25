package main

import "fmt"

/**
   @date 2018年10月25日 星期四
   @description 给定两个由小写字母构成的字符串 A 和 B ，只要我们可以通过交换 A 中的两个字母得到与 B 相等的结果，就返回 true ；否则返回 false 。
   @link https://leetcode-cn.com/problems/buddy-strings/
   @name 亲密字符串


	示例 1：

		输入： A = "ab", B = "ba"
		输出： true
		示例 2：

		输入： A = "ab", B = "ab"
		输出： false
		示例 3:

		输入： A = "aa", B = "aa"
		输出： true
		示例 4：

		输入： A = "aaaaaaabc", B = "aaaaaaacb"
		输出： true
		示例 5：

		输入： A = "", B = "aa"
		输出： false


		提示：

		0 <= A.length <= 20000
		0 <= B.length <= 20000
		A 和 B 仅由小写字母构成。
*/

func case1(datas ... []string)  {
	for _, data := range datas{
		fmt.Println(data, buddyStrings(data[0], data[1]))
	}
}
func main() {
	case1(
		[]string{"aaaaaaabc", "aaaaaaacb"},
		[]string{"ab", "ba"},
		[]string{"ab", "ab"},
		[]string{"aa", "aa"},
		[]string{"", "aa"},
		[]string{"abab", "abab"},
		[]string{"abc", "abc"},
		)
}




func buddyStrings(A string, B string) bool {
	isBuddy := false
	aLen := len(A)
	bLen := len(B)
	if bLen == aLen && aLen > 0{
		for i := 0; i<aLen; i++  {
			c1 := A[i:i+1]
			for j := 0; j<aLen; j++ {
				// 自身不用交换
				if i == j{
					continue
				}
				A1 := ""
				c2 := A[j:j+1]
				if c1 == c2{
					if A == B{
						isBuddy = true
					}
					continue
				}
				if i < j{
					A1 = A[:i] + c2
					if i <= aLen-1 && i + 1 <= j{
						A1 += A[i+1:j]
					}
					A1 += c1
					//if j < aLen - 1 {}
					if j < aLen - 1{
						A1 += A[j+1:]
					}
				}else{
					A1 = A[:j] + c2
					if j <= aLen-1 && j + 1 <= i{
						A1 += A[j+1:i]
					}
					A1 += c1
					//if j < aLen - 1 {}
					if i < aLen - 1{
						A1 += A[i+1:]
					}
				}
				//fmt.Println(A1, i, j, c1, c2)
				if A == A1{
					continue
				}
				if A1 == B{
					isBuddy = true
					break
				}
			}
			if isBuddy{
				break
			}
		}
	}

	return  isBuddy
}
