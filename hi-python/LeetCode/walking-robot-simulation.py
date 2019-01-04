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
                        if ox > x and ox < x1:
                            new_pot = [ox - 1, cy]
                            break
                        elif ox > x1 and ox < x:
                            new_pot = [ox + 1, cy]
                            break
                elif y != None:
                    oy = pot[1]
                    if cx == pot[0]:
                        if oy > y and oy < y1:
                            new_pot = [cx, oy-1]
                            break
                        elif oy > y1 and oy < y:
                            new_pot = [cx, oy + 1]
                            break
            # print(new_pot, x, y, x1, y1, cpoint)
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
            # print(cpoint, ':', direction, '; cmd=', cmd)
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


# -------------------------------------------------------------------
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
        #
        # commands = [-2, -1, -2, 3, 7]
        # obstacles = [[1, -3], [2, -3], [4, 0], [-2, 5], [-5, 2], [0, 0], [4, -4], [-2, -5], [-1, -2], [0, 2]]
        # self.assertEqual(100, Solution().robotSim(commands, obstacles))

        commands = [1,2,-2,5,-1,-2,-1,8,3,-1,9,4,-2,3,2,4,3,9,2,-1,-1,-2,1,3,-2,4,1,4,-1,1,9,-1,-2,5,-1,5,5,-2,6,6,7,7,
                    2,8,9,-1,7,4,6,9,9,9,-1,5,1,3,3,-1,5,9,7,4,8,-1,-2,1,3,2,9,3,-1,-2,8,8,7,5,-2,6,8,4,6,2,7,2,-1,7,
                    -2,3,3,2,-2,6,9,8,1,-2,-1,1,4,7]
        obstacles = [[-57,-58],[-72,91],[-55,35],[-20,29],[51,70],[-61,88],[-62,99],[52,17],[-75,-32],[91,-22],[54,33],
                     [-45,-59],[47,-48],[53,-98],[-91,83],[81,12],[-34,-90],[-79,-82],[-15,-86],[-24,66],[-35,35],[3,31],[87,93],
                     [2,-19],[87,-93],[24,-10],[84,-53],[86,87],[-88,-18],[-51,89],[96,66],[-77,-94],[-39,-1],[89,51],[-23,-72],[27,24],
                     [53,-80],[52,-33],[32,4],[78,-55],[-25,18],[-23,47],[79,-5],[-23,-22],[14,-25],[-11,69],[63,36],[35,-99],[-24,82],
                     [-29,-98],[-50,-70],[72,95],[80,80],[-68,-40],[65,70],[-92,78],[-45,-63],[1,34],[81,50],[14,91],[-77,-54],[13,-88],[24,37],
                     [-12,59],[-48,-62],[57,-22],[-8,85],[48,71],[12,1],[-20,36],[-32,-14],[39,46],[-41,75],[13,-23],[98,10],[-88,64],[50,37],
                     [-95,-32],[46,-91],[10,79],[-11,43],[-94,98],[79,42],[51,71],[4,-30],[2,74],[4,10],[61,98],[57,98],[46,43],[-16,72],[53,-69],
                     [54,-96],[22,0],[-7,92],[-69,80],[68,-73],[-24,-92],[-21,82],[32,-1],[-6,16],[15,-29],[70,-66],[-85,80],[50,-3],[6,13],[-30,-98],
                     [-30,59],[-67,40],[17,72],[79,82],[89,-100],[2,79],[-95,-46],[17,68],[-46,81],[-5,-57],[7,58],[-42,68],[19,-95],[-17,-76],[81,-86],
                     [79,78],[-82,-67],[6,0],[35,-16],[98,83],[-81,100],[-11,46],[-21,-38],[-30,-41],[86,18],[-68,6],[80,75],[-96,-44],[-19,66],[21,84],
                     [-56,-64],[39,-15],[0,45],[-81,-54],[-66,-93],[-4,2],[-42,-67],[-15,-33],[1,-32],[-74,-24],[7,18],[-62,84],[19,61],[39,79],
                     [60,-98],[-76,45],[58,-98],[33,26],[-74,-95],[22,30],[-68,-62],[-59,4],[-62,35],[-78,80],[-82,54],[-42,81],[56,-15],[32,-19],
                     [34,93],[57,-100],[-1,-87],[68,-26],[18,86],[-55,-19],[-68,-99],[-9,47],[24,94],[92,97],[5,67],[97,-71],[63,-57],[-52,-14],
                     [-86,-78],[-17,92],[-61,-83],[-84,-10],[20,13],[-68,-47],[7,28],[66,89],[-41,-17],[-14,-46],[-72,-91],[4,52],[-17,-59],[-85,-46],
                     [-94,-23],[-48,-3],[-64,-37],[2,26],[76,88],[-8,-46],[-19,-68]]
        self.assertEqual(5140, Solution().robotSim(commands, obstacles))

# 运行测试用例
if __name__ == '__main__':
    unittest.main()

