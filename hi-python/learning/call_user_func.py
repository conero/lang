# 2019年9月25日 星期三
# 实现类似 PHP 中的 call_user_func
# version - 3.7
# 参考:
#       http://www.itdaan.com/blog/2009/05/13/992b030eb45c086c6330bba45a76eedf.html

class Aa:
    # 作用域名称与属性名不可重合
    vConst = '类内部变量'
    def __init__(self):
        self.name = 'Base class'

    def getName(self):
        return self.name

    def age(self, age):
        return f'{self.name} age is {age}!'

    def changeVc(self, name):
        self.vConst = name
        return self


# 测试
aa = Aa()
print(getattr(aa, 'age')(16))

print(Aa.vConst)
Aa.vConst = '外部改写'
print(Aa.vConst)

# 类方法存在性检测
print(hasattr(aa, 'hereNoneThisMethod'))

# 内部变量改变
print(f"对象实例内部变量: {aa.changeVc('ChangeName').vConst}, 对象内部变量: {Aa.vConst}")
Aa.vConst = 'ChangeFromTheOuter'
print(f"对象实例内部变量: {aa.changeVc('ChangeName2').vConst}, 对象内部变量: {Aa.vConst}")
