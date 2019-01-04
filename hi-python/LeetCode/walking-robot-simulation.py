# 2019年1月2日 星期三
# @link: https://leetcode-cn.com/problems/walking-robot-simulation/
# python 3.5+
'''
--------------------------------------------------
------| FAILURE ERROR            |----------------
------| @TODO PLEASE HANDLE ME!! |----------------
--------------------------------------------------
'''

'''
机器人在一个无限大小的网格上行走，从点 (0, 0) 处开始出发，面向北方。该机器人可以接收以下三种类型的命令：
    -2：向左转 90 度
    -1：向右转 90 度
    1 <= x <= 9：向前移动 x 个单位长度

在网格上有一些格子被视为障碍物。
第 i 个障碍物位于网格点  (obstacles[i][0], obstacles[i][1])
如果机器人试图走到障碍物上方，那么它将停留在障碍物的前一个网格方块上，但仍然可以继续该路线的其余部分。
返回从原点到机器人的最大欧式距离的平方。

示例 1：
    输入: commands = [4,-1,3], obstacles = []
    输出: 25
    解释: 机器人将会到达 (3, 4)

示例 2：
    输入: commands = [4,-1,4,-2,4], obstacles = [[2,4]]
    输出: 65
    解释: 机器人在左转走到 (1, 8) 之前将被困在 (1, 4) 处

提示：
    0 <= commands.length <= 10000
    0 <= obstacles.length <= 10000
    -30000 <= obstacle[i][0] <= 30000
    -30000 <= obstacle[i][1] <= 30000
    答案保证小于 2 ^ 31
'''




class Solution:
    def robotSim(self, commands, obstacles):
        """
        :type commands: List[int]
        :type obstacles: List[List[int]]
        :rtype: int
        """
        dcts = ['N', 'E', 'S', 'W']     # 方位，[北-东-南-西] 循环链表
        direction = 'N'                 # 方向，当前的方位
        cpoint = [0, 0]                 # 当前位置，坐标

        # 检测障碍物
        def check_obstacles(x=None, y=None, x1=None, y1=None):
            new_pot = None
            cx = cpoint[0]
            cy = cpoint[1]
            for pot in obstacles:
                if x != None: # x 轴坐标移动
                    ox = pot[0]
                    if cy == pot[1]:
                        if ox >= x and ox <= x1:
                            new_pot = [ox - 1, cy]
                            break
                        elif ox >= x1 and ox <= x:
                            new_pot = [ox + 1, cy]
                            break
                elif y != None:
                    oy = pot[1]
                    if cx == pot[0]:
                        if oy >= y and oy <= y1:
                            new_pot = [cx, oy-1]
                            break
                        elif oy >= y1 and oy <= y:
                            new_pot = [cx, oy + 1]
                            break
            print(new_pot, x, y, x1, y1, cpoint)
            return new_pot


        for cmd in commands:
            # 旋转
            # 右顺时针，左逆时针
            if cmd == -2: # <=
                idx = dcts.index(direction)
                vlen = len(dcts)
                if idx == 0:
                    direction = dcts[vlen - 1]
                else:
                    direction = dcts[idx - 1]
                continue
            elif cmd == -1: # =>
                idx = dcts.index(direction)
                vlen = len(dcts)
                if idx == vlen-1:
                    direction = dcts[0]
                else:
                    direction = dcts[idx+1]
                continue
            print(cpoint, ':', direction, '; cmd=', cmd)
            # 值转换
            if direction == 'N':
                new_cmd = cpoint[1] + cmd
                cmd = cpoint[1]
                cpoint[1] = new_cmd
                # 障碍物检测
                co = check_obstacles(y=cmd, y1=new_cmd)
                if co:
                    cpoint = co
            elif direction == 'E':
                new_cmd = cpoint[0] + cmd
                cmd = cpoint[0]
                cpoint[0] = new_cmd
                # 障碍物检测
                co = check_obstacles(x=cmd, x1=new_cmd)
                if co:
                    cpoint = co
            elif direction == 'S':
                new_cmd = cpoint[1] - cmd
                cmd = cpoint[1]
                cpoint[1] = new_cmd
                # 障碍物检测
                co = check_obstacles(y=cmd, y1=new_cmd)
                if co:
                    cpoint = co
            elif direction == 'W':
                new_cmd = cpoint[0] - cmd
                cmd = cpoint[0]
                cpoint[0] = new_cmd
                # 障碍物检测
                co = check_obstacles(x=cmd, x1=new_cmd)
                if co:
                    cpoint = co


        # 欧式距离的平方
        return abs(cpoint[0])**2 + abs(cpoint[1])**2



# 单元测试
import unittest

# 测试用例
class TestCase(unittest.TestCase):
    def test_robotSim(self):
        # commands = [4,-1,3]
        # obstacles = []
        # self.assertEqual(25, Solution().robotSim(commands, obstacles))
        #
        # commands = [4,-1,4,-2,4]
        # obstacles = [[2,4]]
        # self.assertEqual(65, Solution().robotSim(commands, obstacles))
        #
        # commands = [7, -2, -2, 7, 5]
        # obstacles = [[-3, 2], [-2, 1], [0, 1], [-2, 4], [-1, 0], [-2, -3], [0, -3], [4, 4], [-3, 3], [2, 2]]
        # self.assertEqual(4, Solution().robotSim(commands, obstacles))

        commands = [-2, -1, -2, 3, 7]
        obstacles = [[1, -3], [2, -3], [4, 0], [-2, 5], [-5, 2], [0, 0], [4, -4], [-2, -5], [-1, -2], [0, 2]]
        self.assertEqual(100, Solution().robotSim(commands, obstacles))


# 运行测试用例
if __name__ == '__main__':
    unittest.main()

