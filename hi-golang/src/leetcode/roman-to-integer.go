package main

import "fmt"

/**
   @date 2018年10月26日 星期五
   @description 罗马数字转整数
   @link https://leetcode-cn.com/problems/roman-to-integer/
   @name 罗马数字转整数


		罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。

		字符          数值
		I             1
		V             5
		X             10
		L             50
		C             100
		D             500
		M             1000
		例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

		通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

		I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
		X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
		C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
		给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。

		示例 1:

		输入: "III"
		输出: 3
		示例 2:

		输入: "IV"
		输出: 4
		示例 3:

		输入: "IX"
		输出: 9
		示例 4:

		输入: "LVIII"
		输出: 58
		解释: L = 50, V= 5, III = 3.
		示例 5:

		输入: "MCMXCIV"
		输出: 1994
		解释: M = 1000, CM = 900, XC = 90, IV = 4.
*/
type rtICsYl struct {
	s string
	ref int
	success bool
}
func rtICs(rr ...rtICsYl)  {
	fmt.Println(" 正确 ) 罗马文 => 转换值, ==参考值")
	for _, r := range rr{
		s := r.s
		ref := r.ref

		v := romanToInt(s)
		ok := v == ref
		fmt.Printf(" %t ) %s => %d, == %d\n", ok, s, v, ref)

	}
}

type itRCsYl struct {
	n int
	ref string
}
func itRCs(rr ...itRCsYl)  {
	fmt.Println(" 正确 ) 数字 => 转换值, ==参考值")
	for _, r := range rr{
		n := r.n
		ref := r.ref

		v := intToRoman(n)
		ok := v == ref
		fmt.Printf(" %t ) %d => %s, == %s\n", ok, n, v, ref)

	}
}

func main() {
	//rtICs(
	//	rtICsYl{"III", 3, false},
	//	rtICsYl{"IV", 4, false},
	//	rtICsYl{"IX", 9, false},
	//	rtICsYl{"LVIII", 58, false},
	//	rtICsYl{"MCMXCIV", 1994, false},
	//	)

	itRCs(
		itRCsYl{3, "III"},
		itRCsYl{4, "IV"},
		itRCsYl{9, "IX"},
		itRCsYl{58, "LVIII"},
		itRCsYl{1994, "MCMXCIV"},
		itRCsYl{1, "I"},
		itRCsYl{5, "V"},
		itRCsYl{20, "XX"},
		)
}


// 罗马文转整形
func romanToInt(s string) int {
	var value int
	dick := map[string]int{
		"I":1,
		"V":5,
		"X":10,
		"L":50,
		"C":100,
		"D":500,
		"M":1000,
	}
	sLen := len(s)
	for i := 0; i<sLen; i++{
		c := s[i:i+1]
		//fmt.Println(i, c, value)
		if i+1 < sLen{
			c2 := s[i+1: i+2]
			//fmt.Println(i, c, c2, value)
			isSep := false
			switch {
			case c == "I" && (c2 == "V" || c2 == "X"):
				if c2 == "V"{				// IV = 4
					value += 4
				}else {
					value += 9				// IX = 9
				}
				isSep = true
			case c == "X" && (c2 == "L" || c2 == "C"):
				if c2 == "L"{		// XL = 40
					value += 40
				}else {
					value += 90		// XC = 90
				}
				isSep = true
			case c == "C" && (c2 == "D" || c2 == "M"):
				if c2 == "D"{		// CD = 400
					value += 400
				}else {
					value += 900		// CM = 900
				}
				isSep = true
			}
			if isSep{
				i += 1
				continue
			}
		}
		if nx, has := dick[c]; has{
			value += nx
		}
	}

	return value
}



/**
	@descript 整数转罗马数字
	@link https://leetcode-cn.com/problems/integer-to-roman/
 */
func intToRoman(num int) string {
	var s string
	dick := map[int]string{
		1000: "M",
		500: "D",
		100: "C",
		50: "L",
		10: "X",
		5: "V",
		1: "I",
	}
	queue := []int{1000, 500, 100, 50, 10, 5, 1}
	seqDick := map[int]string{
		900: "CM",
		400: "CD",
		90: "XC",
		40: "XL",
		9: "IX",
		4: "IV",
	}

	// 重复字符串
	var repeatFn = func(str string, n int) string {
		nStr := ""
		for i:=0; i<n; i++{
			nStr += str
		}
		return nStr
	}

	for i := num; i>0; {
		// 特殊字符换处理
		if seq, has := seqDick[i]; has{
			s += seq
			i = 0
			break
		}
		for _, e := range queue{
			//fmt.Println(i, e)
			if i >= e{
				// 特殊字符串处理
				for si, sv := range seqDick{
					if i>si && si > e && i - si < si{
						s += sv
						i -= si
					}
				}

				if 0 == i%e{
					i = i/e
					s += repeatFn(dick[e], i)
					i = 0
				}else {
					tn := i%e
					i = (i-tn)/e
					s += repeatFn(dick[e], i)
					i = tn
				}
				break
			}
		}
	}
	return s
}