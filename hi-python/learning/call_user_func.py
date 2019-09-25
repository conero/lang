# 2019年9月25日 星期三
# 实现类似 PHP 中的 call_user_func
# version - 3.7
# 参考:
#       http://www.itdaan.com/blog/2009/05/13/992b030eb45c086c6330bba45a76eedf.html

class Aa:
    def __init__(self):
        self.name = 'Base class'

    def getName(self):
        return self.name

    def age(self, age):
        return f'{self.name} age is {age}!'


# 测试
aa = Aa()
print(getattr(aa, 'age')(16))
