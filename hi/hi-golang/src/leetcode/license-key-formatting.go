package main

import (
	"fmt"
	"strings"
)

// @name 482. 密钥格式化
// @link https://leetcode-cn.com/problems/license-key-formatting/
// @data 2019年3月8日 星期五
// @TODO 算法错误，需要重新改良

/*
	给定一个密钥字符串S，只包含字母，数字以及 '-'（破折号）。N 个 '-' 将字符串分成了 N+1 组。给定一个数字 K，重新格式化字符串，除了第一个分组以外，每个分组要包含 K 个字符，第一个分组至少要包含 1 个字符。两个分组之间用 '-'（破折号）隔开，并且将所有的小写字母转换为大写字母。
给定非空字符串 S 和数字 K，按照上面描述的规则进行格式化。

	示例 1：
		输入：S = "5F3Z-2e-9-w", K = 4
		输出："5F3Z-2E9W"
		解释：字符串 S 被分成了两个部分，每部分 4 个字符；
			 注意，两个额外的破折号需要删掉。

	示例 2：
		输入：S = "2-5g-3-J", K = 2
		输出："2-5G-3J"
		解释：字符串 S 被分成了 3 个部分，按照前面的规则描述，第一部分的字符可以少于给定的数量，其余部分皆为 2 个字符。


	提示:
		S 的长度不超过 12,000，K 为正整数
		S 只包含字母数字（a-z，A-Z，0-9）以及破折号'-'
		S 非空
*/

// 解答
func licenseKeyFormatting(S string, K int) string {
	sep := "-"
	vq := []string{}

	tLen, tQ := 0, []string{}
	for i, w := range strings.Split(S, sep) {
		if i > 0 {
			iLen := len(w)
			var ts string
			// 第二组特殊处理
			if i == 1 && iLen > K {
				vq = append(vq, strings.ToUpper(w))
				ts = strings.Join(vq, "")
				vq = []string{ts[0:K]}

				ts = ts[K:]
				tLen, tQ = len(ts), []string{ts}
				//fmt.Println(vq, tLen, tQ)
				continue
			}

			tAllLen := tLen + iLen
			// 根据长度进行判断
			if tAllLen == K {
				tQ = append(tQ, strings.ToUpper(w))
				vq = append(vq, strings.Join(tQ, ""))
				tLen, tQ = 0, []string{}
			} else if tAllLen < K {
				tLen += iLen
				tQ = append(tQ, strings.ToUpper(w))
			} else {
				tQ = append(tQ, strings.ToUpper(w))
				ts = strings.Join(tQ, "")
				vq = append(vq, ts[0:K])

				ts = ts[K+1:]
				tLen, tQ = len(ts), []string{ts}
			}
		} else {
			vq = append(vq, strings.ToUpper(w))
		}
	}
	return strings.Join(vq, sep)
}

// 运行
func main() {
	var ps, out, comv string
	var pk int
	var fmtS string = "%v) [%s, %v] => [%s] VS [%s]\n"

	// CASE
	ps, pk = "5F3Z-2e-9-w", 4
	out, comv = licenseKeyFormatting(ps, pk), "5F3Z-2E9W"
	//out = licenseKeyFormatting(ps, pk)
	//comv = "5F3Z-2E9W"

	fmt.Printf(fmtS, out == comv, ps, pk, out, comv)

	// CASE
	ps, pk = "2-5g-3-J", 2
	out, comv = licenseKeyFormatting(ps, pk), "2-5G-3J"
	fmt.Printf(fmtS, out == comv, ps, pk, out, comv)

	// CASE
	ps, pk = "2-4A0r7-4k", 4
	out, comv = licenseKeyFormatting(ps, pk), "24A0-R74K"
	fmt.Printf(fmtS, out == comv, ps, pk, out, comv)
}
