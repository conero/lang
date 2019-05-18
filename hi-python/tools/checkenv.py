# python 3.7
# 检测环境变量，是否正确 【环境变量】
# 2019年5月18日 星期六

import os

vpath = os.environ["path"]

vpathQue = vpath.split(',')

for v in vpathQue[:]:
    print(v)
    

