package main

import "fmt"

/*
	在一个小镇里，按从 1 到 N 标记了 N 个人。传言称，这些人中有一个是小镇上的秘密法官。
	如果小镇的法官真的存在，那么：

		小镇的法官不相信任何人。
		每个人（除了小镇法官外）都信任小镇的法官。
		只有一个人同时满足属性 1 和属性 2 。
		给定数组 trust，该数组由信任对 trust[i] = [a, b] 组成，表示标记为 a 的人信任标记为 b 的人。

	如果小镇存在秘密法官并且可以确定他的身份，请返回该法官的标记。否则，返回 -1。



	示例 1：
		输入：N = 2, trust = [[1,2]]
		输出：2

	示例 2：
		输入：N = 3, trust = [[1,3],[2,3]]
		输出：3

	示例 3：
		输入：N = 3, trust = [[1,3],[2,3],[3,1]]
		输出：-1

	示例 4：
		输入：N = 3, trust = [[1,2],[2,3]]
		输出：-1

	示例 5：
		输入：N = 4, trust = [[1,3],[1,4],[2,3],[2,4],[4,3]]
		输出：3


	提示：
		1 <= N <= 1000
		trust.length <= 10000
		trust[i] 是完全不同的
		trust[i][0] != trust[i][1]
		1 <= trust[i][0], trust[i][1] <= N
*/

// 解
func findJudge(N int, trust [][]int) int {
	psMap := map[int][]int{}
	trustList := map[int]bool{}
	judge := -1
	if len(trust) == 0 && N == 1{
		judge = 1
	}
	for _, ts := range trust {
		p, t := ts[0], ts[1]
		dd, has := psMap[t]
		if has {
			psMap[t] = append(dd, p)
		} else {
			psMap[t] = []int{p}
		}
		trustList[p] = true
		// 检查到法官
		//fmt.Println(len(psMap[t]), psMap[t])
		//if len(psMap[t]) == N-1 {
		//	judge = t
		//	_, noTs := trustList[t]
		//	if noTs{
		//		judge = -1
		//	}
		//}
	}
	//fmt.Println(psMap)
	for jd, ps := range psMap {
		if len(ps) == N-1 {
			judge = jd
			_, noTs := trustList[jd]
			if noTs {
				judge = -1
			}
		}
	}
	return judge
}

func main() {
	var pp, comp int

	// 1
	pp = findJudge(2, [][]int{[]int{1, 2}})
	comp = 2
	fmt.Println(comp, pp, pp == comp)

	// 2
	pp = findJudge(3, [][]int{[]int{1, 3}, []int{2, 3}})
	comp = 3
	fmt.Println(comp, pp, pp == comp)

	//
	pp = findJudge(3, [][]int{[]int{1, 2}, []int{2, 3}, []int{3, 1}})
	comp = -1
	fmt.Println(comp, pp, pp == comp)

	//
	pp = findJudge(3, [][]int{[]int{1, 2}, []int{2, 3}})
	comp = -1
	fmt.Println(comp, pp, pp == comp)

	//
	pp = findJudge(4, [][]int{[]int{1, 3}, []int{1, 4}, []int{2, 3}, []int{2, 4}, []int{4, 3}})
	comp = 3
	fmt.Println(comp, pp, pp == comp)

	//
	pp = findJudge(3, [][]int{[]int{1, 3}, []int{2, 3}, []int{3, 1}})
	comp = -1
	fmt.Println(comp, pp, pp == comp)

	//
	pp = findJudge(1, [][]int{})
	comp = 1
	fmt.Println(comp, pp, pp == comp)
}
