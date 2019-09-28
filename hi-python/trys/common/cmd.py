# -*- coding: utf8 -*-
# -*- python: 3.7 -*-
# 2019年9月18日 星期三

import time

'''
 Cmds 类型注册使用
 ```python
    # 实例化项目
    # 也可 Cmds 进行充重置
    cmd = Cmds()
    
    # 注册命令
    def exitCmd():
        # 退出项目
        cmd.quit()
        
    cmd.on('exit', exitCmd)
    # 通用类
    class PartyDemo:
        pass
    # 继承 Party 类，可是现在二级默认调用以及二级未知检测
    class PartyDemo2(Party):
        pass
        
    cmd.party(PartyDemo)
    cmd.party(PartyDemo2)
    # [$ cmd subCmd] 
    cmd.party(PartyDemo, "cmd")
    cmd.party(PartyDemo2, "cmd")
 ```
'''

# 常量
PARTY_CMD_DEFAULT = 'cmdDefault'
PARTY_CMD_UNFIND = 'cmdUnfind'


# 方法存在性检测
def methodExist(obj, method):
    has, mth = False, None
    try:
        if callable(getattr(obj, method)):
            has = True
            mth = getattr(obj, method)
    finally:
        return has, mth


# 通用的命令行解析类
class Cmds:
    def __init__(self, args=None):
        self.ctime = time.time()  # 当前的时间
        self.args = args
        self._quitMk = False  # 退出描述
        self.onKvQue = {}  # 命令配置
        self.partyData = {  # 类型测试
            "cmds": {},
            "instance": []
        }
        self._init_data()

    # 初始化数据
    def _init_data(self):
        self.cmd = ''  # 命令
        self.subCmd = ''  # 子命令
        self.data = {}  # 数据
        self.setting = []  # 配置

    # 命令行解析
    def router(self, args=None):
        self._init_data()
        if args is None:
            args = self.args
        if args is None:
            self.empty()
        else:
            args = args.split(" ")
            for arg in args[:]:
                arg = arg.strip()
                if arg == '':
                    continue
                # --setting
                if arg[:2] == "--":
                    arg = arg[2:]
                    if '=' in arg:
                        idx = arg.index('=')
                        self.data[arg[:idx + 1]] = arg[idx + 1:]
                    else:
                        self.setting.append(arg)
                # -xzy => -x -z -y  无效后顺序
                elif arg[:1] == "-":
                    arg = arg[1:]
                    for v in arg:
                        self.setting.append(v)
                elif self.cmd is None or self.cmd == '':
                    self.cmd = arg
                elif self.subCmd is None or self.subCmd == '':
                    self.subCmd = arg

            if self.before():  # 使用前操作来中断
                return
            if self.cmd == '' or self.cmd is None:  # 空方法
                self.empty()
            elif self.cmd in self.onKvQue:  # 绑定检查成功
                self.onKvQue[self.cmd]()
            elif self._partyCheck():
                pass
            else:
                self.unfind(self.cmd)

    # 命令绑定
    def on(self, cmd, cal):
        self.onKvQue[cmd] = cal
        return self

    # 数据获取
    # 多个注册
    def party(self, ins, cmd=None):
        if cmd != None:
            self.partyData["cmds"][cmd] = ins
        else:
            self.partyData["instance"].append(ins)
        return self

    # part 路由检测
    def _partyCheck(self):
        if self.cmd is None or self.cmd == '':
            return False

        party = self.partyData
        # party["instance"] 类直接监听
        # subCmd 为空
        if self.subCmd is None or self.subCmd == '':
            for ins in party["instance"][:]:
                has, mth = methodExist(ins, self.cmd)
                if has:
                    mth()
                    return True

        if self.cmd in party["cmds"]:
            ins = party["cmds"][self.cmd]

            if self.subCmd is None or self.subCmd == '':
                has, mth = methodExist(ins, PARTY_CMD_DEFAULT)
                if has:
                    mth()
                    return True
            elif self.subCmd != '' and self is not None:
                if self.cmd in party["cmds"]:
                    ins = party["cmds"][self.cmd]
                    has, mth = methodExist(ins, self.subCmd)
                    if has:
                        mth()
                        return True
                    else:
                        has, mth = methodExist(ins, PARTY_CMD_UNFIND)
                        if has:
                            mth(self.subCmd)
                            return True
        return False

    # 默认的空命令运行
    def empty(self):
        print('命令为空运行')

    # 未绑定命令
    def unfind(self, cmd=None):
        print("[" + cmd + "] 不存在")

    # 前置操作
    def before(self):
        ''''''

    # 退出应用
    def quit(self):
        self._quitMk = True

    # 是否退出
    def isQuit(self):
        return self._quitMk

    # 选项检测是否存在
    def option(self, *args):
        has = False
        for arg in args[:]:
            try:
                if self.setting.index(arg) >= 0:
                    has = True
            except:
                pass

            if has:
                break
        return has


# party 执行
class Party:
    # 默认命令
    def cmdDefault(self):
        print(" 默认二级命名执行")

    # 二级命令未知时
    def cmdUnfind(self, cmd=None):
        print(" [" + cmd + "] 二级命令不存在")


# 简单的应用
class SimpleApp:
    def __init__(self, cIns=None):
        self.pref = '$ '
        self.check = '键入任何键退出！'

        if cIns == None:
            cIns = Cmds()
        self.cIns = cIns

    # 运行
    def run(self):
        # 循环等待输入
        while (True):
            ipt = input(self.pref)

            if ipt == 'exit':
                break

            self.cIns.router(ipt)
            print('')
            if self.cIns.isQuit():
                break

        # 最后的退出确认
        input(self.check)
