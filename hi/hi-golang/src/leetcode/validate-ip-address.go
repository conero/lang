package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

/**
@date 2018年11月1日 星期四
@description 编写一个函数来验证输入的字符串是否是有效的 IPv4 或 IPv6 地址。
@link https://leetcode-cn.com/problems/validate-ip-address/
@name validate-ip-address (via)

	IPv4 地址由十进制数和点来表示，每个地址包含4个十进制数，其范围为 0 - 255， 用(".")分割。比如，172.16.254.1；
	同时，IPv4 地址内的数不会以 0 开头。比如，地址 172.16.254.01 是不合法的。
	IPv6 地址由8组16进制的数字来表示，每组表示 16 比特。这些组数字通过 (":")分割。比如,  2001:0db8:85a3:0000:0000:8a2e:0370:7334 是一个有效的地址。而且，我们可以加入一些以 0 开头的数字，字母可以使用大写，也可以是小写。所以， 2001:db8:85a3:0:0:8A2E:0370:7334 也是一个有效的 IPv6 address地址 (即，忽略 0 开头，忽略大小写)。
	然而，我们不能因为某个组的值为 0，而使用一个空的组，以至于出现 (::) 的情况。 比如， 2001:0db8:85a3::8A2E:0370:7334 是无效的 IPv6 地址。
	同时，在 IPv6 地址中，多余的 0 也是不被允许的。比如， 02001:0db8:85a3:0000:0000:8a2e:0370:7334 是无效的。
	说明: 你可以认为给定的字符串里没有空格或者其他特殊字符。

	示例 1:
	输入: "172.16.254.1"
	输出: "IPv4"
	解释: 这是一个有效的 IPv4 地址, 所以返回 "IPv4"。

	示例 2:
	输入: "2001:0db8:85a3:0:0:8A2E:0370:7334"
	输出: "IPv6"
	解释: 这是一个有效的 IPv6 地址, 所以返回 "IPv6"。

	示例 3:
	输入: "256.256.256.256"
	输出: "Neither"
	解释: 这个地址既不是 IPv4 也不是 IPv6 地址。
*/

func main() {
	type TestData struct {
		input  string
		output string
	}
	fmt.Println("validIPAddress 测试如下：")
	testVia := func(datas ...TestData) {
		for _, dd := range datas {
			out := validIPAddress(dd.input)
			fmt.Printf("%t) 需校验IP: %s, 实际输出: %s == %s\n", out == dd.output, dd.input, out, dd.output)
		}
	}

	testVia(
		TestData{"0.0.255.1", "IPv4"},
		TestData{"172.16.254.01", "Neither"},
		TestData{"12.12.12.12.12", "Neither"},
		TestData{"172.16.254.1", "IPv4"},
		TestData{"2001:0db8:85a3:0:0:8A2E:0370:7334", "IPv6"},
		TestData{"f:f:f:f:f:f:f:f", "IPv6"},
		TestData{"256.256.256.256", "Neither"},
		TestData{"2001:0db8:85a3:0:0:8A2E:0370:7334:", "Neither"},
		TestData{"2001:0db8:85a3:00000:0:8A2E:0370:7334", "Neither"},
		TestData{"2001:db8:85a3:0::8a2E:0370:7334", "Neither"},
		TestData{"1e1.4.5.6", "Neither"},
		TestData{"1.1.1.1.", "Neither"},
	)
}

func validIPAddress(IP string) string {
	const Neither = "Neither"
	var ipType string = Neither
	if matched, _ := regexp.MatchString("[\\d]{1,3}\\.[\\d]{1,3}\\.[\\d]{1,3}\\.[\\d]{1,3}", IP); matched {
		matched2, _ := regexp.MatchString("[a-zA-Z]+", IP)
		isValidV4 := !matched2
		if isValidV4 && strings.Count(IP, ".") > 3 {
			isValidV4 = false
		}
		for _, s := range strings.Split(IP, ".") {
			if !isValidV4 {
				break
			}
			// 大于 255
			if n, _ := strconv.Atoi(s); n > 255 {
				isValidV4 = false
				break
			}
			if cmatched, _ := regexp.MatchString("^0{1,}.*", s); cmatched && len(s) > 1 {
				isValidV4 = false
				break
			}
		}
		if isValidV4 {
			ipType = "IPv4"
		}
	}
	// ipv6 本题不允许 "::"
	// X:X:X:X:X:X:X:X(4*8X=128bit)
	const ipv6MaxLen = 128
	if matched, _ := regexp.MatchString("^[0-9A-Fa-f]+(:{1}[0-9A-Fa-f])+[:0-9A-Fa-f]+$", IP); matched &&
		ipType == Neither &&
		len(IP) <= ipv6MaxLen {
		isValidV6 := true
		if cmatched, _ := regexp.MatchString("[:]{2}", IP); cmatched {
			isValidV6 = false
		}
		if isValidV6 && strings.Count(IP, ":") > 7 {
			isValidV6 = false
		}
		for _, ss := range strings.Split(IP, ":") {
			if !isValidV6 {
				break
			}
			if len(ss) > 4 {
				isValidV6 = false
				break
			}
		}
		if isValidV6 {
			ipType = "IPv6"
		}
	}
	//fmt.Println(regexp.MatchString("^[0-9A-Fa-f:]+", IP))
	return ipType
}
