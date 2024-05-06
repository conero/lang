# python 3.7
# 检测环境变量，是否正确 【环境变量】
# 2019年5月18日 星期六

import os

# 环境变量检查
# IED 下可以
from trys.common.cmd import SimpleApp, Cmds, Party
# from .common.cmd import SimpleApp, Cmds, Party

class Stool(Cmds):
    def before(self):
        pass

class cmdEnv(Party):
    # 检测是否正确
    def check(self):
        vpath = os.environ["path"]

        vpathQue = vpath.split(';')

        fixMk = mycmd.option('f', 'fix')
        fixData = []

        for v in vpathQue[:]:
            v = v.strip()
            # 跳过空字符串
            if v == "":
                continue
            v = v.replace('\\', '/')
            if os.path.exists(v) != True:
                print(" " + v)
            elif fixMk:
                fixData.append(v)

        # 修复数据
        if fixMk:
            ss = ';'.join(fixData)
            os.environ['path'] = ss


def cmdHelp():
    print(' . env 环境变量处理')
    print('   - check  -f[--fix] 环境变量有效目录检测, fix 可删除无效的path')
    print(' . exit 退出程序')

mycmd = Stool()

mycmd.on('help', cmdHelp).party(cmdEnv(), "env")

# 欢迎提示语
print(' 欢迎使用 winSysTool Made by Joshua Conero!')
print(' help 提示帮助, exit 退出')

SimpleApp(mycmd).run()
