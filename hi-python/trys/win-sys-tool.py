# python 3.7
# 检测环境变量，是否正确 【环境变量】
# 2019年5月18日 星期六

import os


# 环境变量检查
from trys.common.cmd import Cmds


def checkEnv():
    vpath = os.environ["path"]

    vpathQue = vpath.split(';')

    for v in vpathQue[:]:
        v = v.strip()
        # 跳过空字符串
        if v == "":
            continue
        v = v.replace('\\', '/')
        if os.path.exists(v) != True:
            print("!" + v)



class Stool(Cmds):
    ''''''

class OnCmd:
    ''''''


mycmd = Stool()

# 欢迎提示语
print(' 欢迎使用 winSysTool Made by Joshua Conero!')
print(' --help 提示帮助, --exit 退出')

# 循环等待输入
while (True):
    ipt = input('$ ')
    mycmd.router(ipt)
    if ipt == '--exit':
        break
    
    print('')
    if mycmd.isQuit():
        break

# 最后的退出确认
input(' 键入任何键退出！')
