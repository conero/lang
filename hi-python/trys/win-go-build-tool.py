# encoding = utf8
# 2019年8月28日 星期三
# Windows 系统下go语言的自动编译工具

# 
# 主要看点，可利用的特性
#   .通过遍历目录以及go脚本，找到对应的包命名
#   .多个文件中筛选go脚本
#   .python 启动 Windows 的 powershell 进行go语言编译工具（多进程）


import sys
import os
import re
# import multiprocessing
import subprocess

# 用于测试的变量
# 目标地址
vdir = "C:/conero/repository/golang/iris/_examples"

# 参照的文件名称
refname = "main.go"
refname = ""
# refname = None


# 目标目录获取
pref = '  '
print('Windows 系统下go语言的自动编译工具!')

if vdir == "":
    vdir = input(f'{pref}.输入需要编译的go项目模板目录：')
if refname == None:
    refname = input(f'{pref}.输入快速的参照文件名称：')

print(f'正遍历目标目录({vdir})……')


def goBuildDir(vp):
    ''''''
    print(f'正在编译{vp}')

    # cmds = [r'powershell', r'echo', '7*2']
    cmds = [r'powershell', r'cd', vp, '|', 'go', 'build']
    p = subprocess.Popen(cmds, stdout=subprocess.PIPE)
    dt = p.stdout.read()
    print(dt)

    # os.system('powershell')
# go build 命令
def scandirTagert(vp):
    files = os.listdir(vp)
    for fl in files:
        tvp = os.path.join(vp, fl)
        if os.path.isdir(tvp):
            scandirTagert(tvp)
        else:
            tfl = fl.lower()
            # clsprint(re.match(r'.*(\.go)+$', tfl))
            if re.match(r'.*(\.go)$', tfl):
                if refname.lower() == tfl:
                    # 参照文件匹配为目标地址
                    # 运行编译命令
                    goBuildDir(vp)
                else:
                    # 通过文本读取，查找
                    for lines in open(tvp, 'r', encoding='UTF-8'):
                        lines = lines.strip()
                        # 空行
                        if lines == "" or len(lines) == 0:
                            continue
                        # 注释行
                        if lines[:2] == "//":
                            continue
                        # package 定义行，找到报名后直接退出文本遍历
                        if 'package' in lines:
                            lines = lines.replace('package', '').strip()
                            if lines == "main":
                                # 目标编译
                                goBuildDir(vp)
                                break
                            break

if os.path.exists(vdir):
    scandirTagert(vdir)
else:
    print(f'({vdir}) 不是有效的目录！')
