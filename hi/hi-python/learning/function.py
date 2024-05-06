# 2018年12月16日 星期日
# python 函数
# >= 3.6


# 默认列表
def quque(*args):
    for v in args:
        print(v)

# k-v 结构参数
def kv(**kvs):
    for k in kvs.keys():
        print(k, kvs[k])

# 多返回值
def mutReturn():
    return 1, 2, 3, 4, 'Joshua Conero'

# ----------------------------------------------------
# 数据测试
# quque(1, 5, 'ddd', 'dd')        
# quque()

kv(test="222", t=5)

# 多行，返回 元组
print(mutReturn())
a,b, c, b, e = mutReturn()
print(a)
