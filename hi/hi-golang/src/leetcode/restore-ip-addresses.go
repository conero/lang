package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

/**
@date 2018年10月31日 星期三
@description 给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。
@link https://leetcode-cn.com/problems/restore-ip-addresses/
@name restore-ip-addresses (ripa)


示例:
输入: "25525511135"
输出: ["255.255.11.135", "255.255.111.35"]

*/

func main() {
	s := ""

	s = "25525511135"
	fmt.Printf("输入字符: %s => %s\n", s, restoreIpAddresses(s))

	s = "010010"
	fmt.Printf("输入字符: %s => %s\n", s, restoreIpAddresses(s))
}

// IPv4
// 数据区间: [0-225]			数字位数: 3*4 = 12
// 数字扫描，1-3 位扫描
// [[1,2,3], [1,2,3], [1,2,3], [1,2,3]]
// 组合排列 3
func restoreIpAddresses(s string) []string {
	reg := regexp.MustCompile("\\s")
	s = reg.ReplaceAllString(s, "")
	addrs := []string{}
	sLen := len(s)

	// 位扫描
	// i->j->m-n
	//bit := 12
	topMax := 255
	isGoodStrFn := func(ipx string) bool {
		// 以 0 开头的数字无效; 如: 01, 001, 011
		if matched, _ := regexp.MatchString("^0+[^0]+", ipx); matched {
			return true
		}
		// 多个零无效, 如: 00, 000
		if matched, _ := regexp.MatchString("^0{2,}.*", ipx); matched {
			return true
		}
		ip1N, err := strconv.Atoi(ipx)
		if err != nil || ip1N > topMax {
			return true
		}
		return false
	}
	for i := 1; i <= 3 && i <= sLen; i++ {
		ip1 := s[0:i]
		if isGoodStrFn(ip1) {
			continue
		}
		for j := 1; j <= 3 && (i+j) <= sLen; j++ {
			ip2 := s[i : i+j]
			if isGoodStrFn(ip2) {
				continue
			}
			for m := 1; m <= 3 && (i+j+m) <= sLen; m++ {
				ip3 := s[i+j : i+j+m]
				//fmt.Println(ip3, i, j, m)
				if isGoodStrFn(ip3) {
					continue
				}
				for n := 1; n <= 3 && (i+j+m+n) <= sLen; n++ {
					if (i + j + m + n) != sLen {
						continue
					}
					ip4 := s[i+j+m : i+j+m+n]
					if isGoodStrFn(ip4) {
						continue
					}
					addrs = append(addrs, strings.Join([]string{ip1, ip2, ip3, ip4}, "."))
				}
			}
		}
	}

	return addrs
}
