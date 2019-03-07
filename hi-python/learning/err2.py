# 2018年12月17日 星期一
# python 错误异常处理
# >= 3.6
import sys

argv = sys.argv

j = 'none input the string data.'
if len(argv) == 1:
    j = input('give me the data: ')

try:
    i = i*1
    print(i)
except BaseException as identifier:
    print(' the script run error')
    print(' msg: {}'.format(identifier))
    pass
finally:
    print(' run right or wrong finnally.')   
    print(' data: {}'.format(j)) 