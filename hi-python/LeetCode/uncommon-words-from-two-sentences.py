# 2019年1月10日 星期四
# @link: https://leetcode-cn.com/problems/uncommon-words-from-two-sentences/
# python 3.5+

'''
    给定两个句子 A 和 B 。 （句子是一串由空格分隔的单词。每个单词仅由小写字母组成。）
    如果一个单词在其中一个句子中只出现一次，在另一个句子中却没有出现，那么这个单词就是不常见的。
    返回所有不常用单词的列表。
    您可以按任何顺序返回列表。

 

    示例 1：
        输入：A = "this apple is sweet", B = "this apple is sour"
        输出：["sweet","sour"]
        
    示例 2：
        输入：A = "apple apple", B = "banana"
        输出：["banana"]
 
    提示：
        0 <= A.length <= 200
        0 <= B.length <= 200
        A 和 B 都只包含空格和小写字母。
'''


class Solution:
    def uncommonFromSentences(self, A, B):
        """
        :type A: str
        :type B: str
        :rtype: List[str]
        """
        a_que,b_que = A.split(' '), B.split(' ')
        a_len, b_len = len(a_que), len(b_que)

        if a_len >= b_len:
            more, less = a_que, b_que
            m_len, s_len = a_len, b_len
        else:
            more, less = b_que, a_que
            m_len, s_len= b_len, a_len

        # print(more, less, a_len, b_len, a_que, b_que, (m_len, s_len))
        data = []
        for idx in range(len(more)):
            ms = more[idx]
        # for idx, ms in more:
        #     print(more.count(ms), less.count(ms))
            if more.count(ms) == 1 and less.count(ms) == 0:
                data.append(ms)
            if idx < s_len:
                ls = less[idx]
                # print(more.count(ls), less.count(ls), ls)
                if more.count(ls) == 0 and less.count(ls) == 1:
                    data.append(ls)

        return  data





# -------------------------------------------------------------------
# 单元测试
import unittest

# 测试用例
class TestCase(unittest.TestCase):
    def test_uncommonFromSentences(self):
        A = "this apple is sweet"
        B = "this apple is sour"
        self.assertEqual(["sweet","sour"], Solution().uncommonFromSentences(A, B))

        A = "apple apple"
        B = "banana"
        self.assertEqual(["banana"], Solution().uncommonFromSentences(A, B))

        A = "abcd def abcd xyz"
        B = "ijk def ijk"
        self.assertEqual(["xyz"], Solution().uncommonFromSentences(A, B))
# 运行测试用例
if __name__ == '__main__':
    unittest.main()

