# python 3.7
# 检测环境变量，是否正确 【环境变量】
# 2019年5月18日 星期六

import os

# 环境变量检查
from trys.common.cmd import SimpleApp, Cmds, Party


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
    def before(self):
        pass

class cmdEnv(Party):
    ''''''


mycmd = Stool()
mycmd.party(cmdEnv(), "env")

# 欢迎提示语
print(' 欢迎使用 winSysTool Made by Joshua Conero!')
print(' help 提示帮助, exit 退出')

SimpleApp(mycmd).run()
