package main

import (
	"fmt"
	"strconv"
)

/**
 * @DATE        2019/4/26
 * @NAME        Joshua Conero
 * @DESCRIPIT   443. 压缩字符串
 * @LINK 		https://leetcode-cn.com/problems/string-compression/
 **/

// 题目
// @TODO 提交失败，本地输出的几个与服务器不一致，本地通过服务器不通过（本地 go1.12.1  VS 服务器 Go 1.10.3）。

/*
给定一组字符，使用原地算法将其压缩。
压缩后的长度必须始终小于或等于原数组长度。
数组的每个元素应该是长度为1 的字符（不是 int 整数类型）。
在完成原地修改输入数组后，返回数组的新长度。


进阶：
你能否仅使用O(1) 空间解决问题？


示例 1：
	输入：
		["a","a","b","b","c","c","c"]
	输出：
		返回6，输入数组的前6个字符应该是：["a","2","b","2","c","3"]
	说明：
		"aa"被"a2"替代。"bb"被"b2"替代。"ccc"被"c3"替代。


示例 2：
	输入：
		["a"]
	输出：
		返回1，输入数组的前1个字符应该是：["a"]
	说明：
		没有任何字符串被替代。

示例 3：
	输入：
		["a","b","b","b","b","b","b","b","b","b","b","b","b"]
	输出：
		返回4，输入数组的前4个字符应该是：["a","b","1","2"]。

	说明：
		由于字符"a"不重复，所以不会被压缩。"bbbbbbbbbbbb"被“b12”替代。
		注意每个数字在数组中都有它自己的位置。

注意：
	所有字符都有一个ASCII值在[35, 126]区间内。
	1 <= len(chars) <= 1000。
*/

// 解答
func compress(chars []byte) int {
	var pStr string
	var pCtt int

	var vChars []byte = []byte{}
	// 中间循环
	for i, c := range chars {
		cs := string(c)
		if cs != pStr {
			// 起始
			if pStr == "" {
				pStr = cs
				pCtt = 1
			} else if i > 0 && string(chars[i-1]) == pStr {
				vChars = append(vChars, chars[i-1])
				if pCtt > 1 {
					cByte := []byte(strconv.Itoa(pCtt))
					vChars = append(vChars, cByte[0])
				}
				pStr = cs
				pCtt = 1
			}
		} else {
			pCtt += 1
		}
	}
	// 末尾处理
	if pStr != "" {
		vChars = append(vChars, []byte(pStr)[0])
		if pCtt > 1 {
			cByte := []byte(strconv.Itoa(pCtt))
			vChars = append(vChars, cByte...)
		}
	}
	bs := len(vChars)
	//chars = vChars
	//fmt.Println(string(vChars))
	return bs
}

// 控制台
func main() {
	const format = "%v)    {%s} => {%v} VS {%v}\n"
	var vin []byte
	var vout, vneed int

	// case
	vin = []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	vout, vneed = compress(vin), 6
	fmt.Printf(format, vout == vneed, vin, vout, vneed)

	// case
	vin = []byte{'a'}
	vout, vneed = compress(vin), 1
	fmt.Printf(format, vout == vneed, vin, vout, vneed)

	// case
	vin = []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
	vout, vneed = compress(vin), 4
	fmt.Printf(format, vout == vneed, vin, vout, vneed)

	vin = []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	vout, vneed = compress(vin), 6
	fmt.Printf(format, vout == vneed, vin, vout, vneed)

}
