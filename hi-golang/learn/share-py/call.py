# python 调用 dll库
# 2018年12月10日 星期一

import ctypes

lib = ctypes.CDLL("share.dll")


# 调用 joshua_conero
lib.joshua_conero()

# 调用
print(lib.cal(2, 1))
print(lib.cal(3, 2))
print(lib.cal(4, 3))
print(lib.cal(5, 6))

# 调用
print(lib.cal_about())

