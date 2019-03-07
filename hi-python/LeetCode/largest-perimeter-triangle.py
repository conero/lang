# 2019年1月18日 星期五
# @link: https://leetcode-cn.com/problems/largest-perimeter-triangle/
# python 3.5+
'''
--------------------------------------------------
------| FAILURE ERROR            |----------------
------| @TODO PLEASE HANDLE ME!! |----------------
--------------------------------------------------
'''

'''
给定由一些正数（代表长度）组成的数组 A，返回由其中三个长度组成的、面积不为零的三角形的最大周长。
如果不能形成任何面积不为零的三角形，返回 0。

 

示例 1：
    输入：[2,1,2]
    输出：5
    
示例 2：
    输入：[1,2,1]
    输出：0
    
示例 3：
    输入：[3,2,3,4]
    输出：10
    
示例 4：
    输入：[3,6,2,3]
    输出：8
 

提示：
    3 <= A.length <= 10000
    1 <= A[i] <= 10^6
'''


class Solution:
    def largestPerimeter(self, A):
        """
        :type A: List[int]
        :rtype: int
        """
        if len(A) < 3:
            return 0
        # 三角形三边关系： 任意两边之和大于第三边，任意两边之差小于第三边         



# -------------------------------------------------------------------
# 单元测试
import unittest

# 测试用例
class TestCase(unittest.TestCase):
    def test_largestPerimeter(self):
        A = [2,1,2]
        self.assertEqual(2, Solution().largestPerimeter(A))

        A = [1,2,1]
        self.assertEqual(0, Solution().largestPerimeter(A))

        A = [3,2,3,4]
        self.assertEqual(10, Solution().largestPerimeter(A))

        A = [3,6,2,3]
        self.assertEqual(8, Solution().largestPerimeter(A))

# 运行测试用例
if __name__ == '__main__':
    unittest.main()

