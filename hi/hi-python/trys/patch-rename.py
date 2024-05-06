import os
import re

def rename_file(vpath, index, prex='', like=None, limit=None, suffix=None, onlyshow=None):
    count = 1

    likeReg = None
    if like:
        likeReg = re.compile(like)
    for name in os.listdir(vdir):
        if limit:
            if count > limit:
                break
            count += 1
        _, ve = os.path.splitext(name)

        # 后缀筛选
        if suffix and ve != suffix:
            continue
        
        # 正则匹配
        if like and likeReg:
            if not likeReg.match(name):
                continue

        rNameStr = f'{prex}{index}{ve}'
        vd = f'{vpath}{name}'

        print(f'{vd} --> {rNameStr}')
        if not onlyshow:
            os.rename(vd, f'{vpath}{rNameStr}')

        index += 1


def input_number(text, vdef=0):
    try:
        vnum = int(input(text))
    except Exception:
        vnum = 0
    return vnum

if __name__ == "__main__":
    # vdir = './'
    # rename_file(vdir, 297, suffix='.jpg', limit=20, prex='p', like='IMG_2020')
    print('欢迎使用批量文件名称修改修改器，根据索引修改，适用于图片等二进制文件')
    print('')
    vdir = input('清输入目标地址(./)：')
    if vdir == '' or not vdir:
        vdir = './'
    
    index = input_number('清输入索引名称(数字)：')
    prex = input('清输入前缀：')
    like = input('清输入正则匹配：')
    limit = input_number('清输入修改数量（数字）：')
    suffix = input('清输入过滤后缀(如 .jpg)：')
    onlyshow = input('清输入仅仅显示结果，不执行：')
    
    print(f'（目录： {vdir} ， 索引： {index} ， 前缀：{prex} ， 正则过滤：{like} ， 限制数：{limit} ， 后缀：{suffix} ， 不执行：{onlyshow} ？）')
    check = input(f'您确定要进行批量文件修改吗？(y)')

    if check and '' != check:
        rename_file(vdir, index, prex=prex, like=like, limit=limit, suffix=suffix, onlyshow=onlyshow)
