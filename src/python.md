# Python(py)

> 2018年8月9日 星期四
>
> Joshua Conero



# Python 学习笔记

- 官网:  https://www.python.org/





## 概述

​	日期： 2017年2月23日 星期四

​	特点： “优雅”、“明确”、“简单”



### 安装

安装 python 运行包以后，可以在安装的目录下查看语言的核心库等源码



## 基础

```python
# -*- coding: encoding -*-
# 首行申明编码格式: utf-8
# 当行注释
```

### numbers/数字

> 使用 python 解析作为计算器

- _支持简单运算符号: `+,-,*,/`(加减乘除)，以及括号运算_
- `//` 整数除法
- `%` 余运算
- `**` 平方(x^n)





```powershell
# 进入 python 交互模式
python

>>> 7 // 3  # =2 ; 整数除法
>>> 7 % 3   # = 1; 余运算(x % y = (x - x/y))
>>> 5**3	# = 125; 5^3

```



### strings/字符串

_单引号或双引号包裹的字符, 支持`\` 转移符；`print`函数，字符串默认解析转移字符。`+`_ 拼接字符串

```python
>>> "I'm Joshua Conero."		# I'm Joshua Conero
>>> 'I\'m Joshua Conero.'		# I'm Joshua Conero
>>> 'To learning "python".'		# To learning "python".

>>> print('C:\some\name')
C:\some
ame

# 使用 r'/r" 标识: raw
>>> print(r'C:\some\name')
>>> print(r"C:\some\name")
C:\some\name
    
# 使用 * 重复字符互串    
>>> 'g' + 3*'o' + 'd'		# goood

# 自动连接字符串
>>> 'go' 'od' ' new' 's'    # good news    
```



> 字符串

```
+---+---+---+---+---+---+
| P | y | t | h | o | n |
+---+---+---+---+---+---+
0   1   2   3   4   5   6
-6  -5  -4  -3  -2  -1
```



```powershell

py = 'Python'
py[0]		# =P; 获取第一个字符
py[:]		# =Python; 相当于整个字符串
py[0:2]		# =Py;
py[-4:]		# =thon;
```



### Lists/集合

```python
squares = [1, 4, 9, 16, 25]
squares			# =[1, 4, 9, 16, 25]
squares[:]		# ; 同上
squares[-2:]	# ; [16, 25]
squares + [-25, -16]	# 数组append运算
squares.append(36)		# append 操作
squares
len(squares)
```





## 流程控制

### if

```python
x = int(input("Please enter an integer: "))
if x < 0:
    print('您输入一个负数')
elif x == 0:
    print('您输入零值')
elif x > 0:
    print('您输入一个正数')
else:
    print('位置您输入数的范围.')

# 函数中的 if-elif    
def ask_ok(prompt, retries=4, reminder='Please try again!'):
    while True:
        ok = input(prompt)
        if ok in ('y', 'ye', 'yes'):
            return True
        if ok in ('n', 'no', 'nop', 'nope'):
            return False
        retries = retries - 1
        if retries < 0:
            raise ValueError('invalid user response')
        print(reminder)   
```



### for

```python
words = ['w', 'w1', 'w2', '...', 'wn']

# 循环
# for w in words[:]:
for w in words:
    # w,w1,w2
    print(w)
    
# 通过长度
for i in range(len(words)):
    # i 系统处理
    print(words[i])
```



*`break`、`continue` 中断/跳过当前的步骤*



### function

```python
# 申明函数
# 可使用默认值
def f(param, d=-1):
    # function body
    # 返回值
    return d

# 使用函数
t = f(..)
```



_函数默认值，在函数定义的作用域中有效。_

```python
i = 5

def f(arg=i):
    print(arg)

i = 6
f()
```



_默认值有限性仅仅单次，单可变对象如：`list, dictionary, or instances of most classes`多次有效。_

```python
def f(a, L=[]):
    L.append(a)
    return L

print(f(1))		# [1]
print(f(2))		# [1, 2]
print(f(3))		# [1, 2, 3]

# 或使用不同的函数写法避免
def f(a, L=None):
    if L is None:
        L = []
    L.append(a)
    return L
```



**Keyword Arguments/ 关键字参数**

_格式： `kwarg=value`_

```python
# 申明函数
def parrot(voltage, state='a stiff', action='voom', type='Norwegian Blue'):
    print("-- This parrot wouldn't", action, end=' ')
    print("if you put", voltage, "volts through it.")
    print("-- Lovely plumage, the", type)
    print("-- It's", state, "!")

# 使用函数    
parrot(1000)                                          # 1 positional argument
parrot(voltage=1000)                                  # 1 keyword argument
parrot(voltage=1000000, action='VOOOOOM')             # 2 keyword arguments
parrot(action='VOOOOOM', voltage=1000000)             # 2 keyword arguments
parrot('a million', 'bereft of life', 'jump')         # 3 positional arguments
parrot('a thousand', state='pushing up the daisies')  # 1 positional, 1 keyword
```



**形参**

```python
def cheeseshop(kind, *arguments, **keywords):
    print("-- Do you have any", kind, "?")
    print("-- I'm sorry, we're all out of", kind)
    for arg in arguments:
        print(arg)
    print("-" * 40)
    for kw in keywords:
        print(kw, ":", keywords[kw])
        
cheeseshop("Limburger", "It's very runny, sir.",
           "It's really very, VERY runny, sir.",
           shopkeeper="Michael Palin",
           client="John Cleese",
           sketch="Cheese Shop Sketch")

# 输出
'''
-- Do you have any Limburger ?
-- I'm sorry, we're all out of Limburger
It's very runny, sir.
It's really very, VERY runny, sir.
----------------------------------------
shopkeeper : Michael Palin
client : John Cleese
sketch : Cheese Shop Sketch
'''
```

*任意参数列表：*

```powershell
>>> def concat(*args, sep="/"):
...     return sep.join(args)
...
>>> concat("earth", "mars", "venus")
'earth/mars/venus'
>>> concat("earth", "mars", "venus", sep=".")
'earth.mars.venus'
```

*使用任意参数*

```python
# 集合
list(range(3, 6))  
args = [3, 6]
list(range(*args))

# 字典
def parrot(voltage, state='a stiff', action='voom'):
    print("-- This parrot wouldn't", action, end=' ')
    print("if you put", voltage, "volts through it.", end=' ')
    print("E's", state, "!")

d = {"voltage": "four million", "state": "bleedin' demised", "action": "VOOM"}
parrot(**d)
```



*`lambda` 表达式*

```python
def make_incrementor(n):
    return lambda x: x + n

f = make_incrementor(42)
f(0)
f(1)
```



**文档字符**

```python
def my_function():
    """Do nothing, but document it.

    No, really, it doesn't do anything.
    """
    pass

print(my_function.__doc__)
```



**函数注解 (Function Annotations)**

```python
# 注解符号:  ->
def f(ham: str, eggs: str = 'eggs') -> str:
    print("Annotations:", f.__annotations__)
    print("Arguments:", ham, eggs)
    return ham + ' and ' + eggs

f('spam')
```



## 数据结构

### list/列表对象

> `Built-In` 通过编辑器如“Ctrl + ” 追踪脚本

```python

class list(object):
    """
    list() -> new empty list
    list(iterable) -> new list initialized from iterable's items
    """

    def append(self, p_object):  # real signature unknown; restored from __doc__
        """ L.append(object) -> None -- append object to end """
        pass

    def clear(self):  # real signature unknown; restored from __doc__
        """ L.clear() -> None -- remove all items from L """
        pass

    def copy(self):  # real signature unknown; restored from __doc__
        """ L.copy() -> list -- a shallow copy of L """
        return []

    def count(self, value):  # real signature unknown; restored from __doc__
        """ L.count(value) -> integer -- return number of occurrences of value """
        return 0

    def extend(self, iterable):  # real signature unknown; restored from __doc__
        """ L.extend(iterable) -> None -- extend list by appending elements from the iterable """
        pass

    def index(self, value, start=None, stop=None):  # real signature unknown; restored from __doc__
        """
        L.index(value, [start, [stop]]) -> integer -- return first index of value.
        Raises ValueError if the value is not present.
        """
        return 0

    def insert(self, index, p_object):  # real signature unknown; restored from __doc__
        """ L.insert(index, object) -- insert object before index """
        pass

    def pop(self, index=None):  # real signature unknown; restored from __doc__
        """
        L.pop([index]) -> item -- remove and return item at index (default last).
        Raises IndexError if list is empty or index is out of range.
        """
        pass

    def remove(self, value):  # real signature unknown; restored from __doc__
        """
        L.remove(value) -> None -- remove first occurrence of value.
        Raises ValueError if the value is not present.
        """
        pass

    def reverse(self):  # real signature unknown; restored from __doc__
        """ L.reverse() -- reverse *IN PLACE* """
        pass

    def sort(self, key=None, reverse=False):  # real signature unknown; restored from __doc__
        """ L.sort(key=None, reverse=False) -> None -- stable sort *IN PLACE* """
        pass
```



*主要方法名*

| 方法名  | 功能描述           | 用法                                       |
| ------- | ------------------ | ------------------------------------------ |
| append  | 尾部追加元素       | L.append(object) -> None                   |
| clear   | 清空元素           | L.clear() -> None                          |
| copy    | 浅复制             | L.copy() -> list                           |
| count   | 统计元素次数       | L.count(value) -> integer                  |
| extend  | 迭代添加元素       | L.extend(iterable) -> None                 |
| index   | 首次出现的位置     | L.index(value, [start, [stop]]) -> integer |
| insert  | 在index前插入元素  | L.insert(index, object) -> None            |
| pop     | 移除元素(默认最后) | L.pop([index]) -> item                     |
| remove  | 移除指定的第一个值 | L.remove(value) -> None                    |
| reverse | 翻转元素           | L.reverse() -> None                        |
| sort    | 元素排序           | L.sort(key=None, reverse=False) -> None    |



> _**集合推导**_

```python
# 使用 for 添加元素
squares = []
for x in range(10):
    squares.append(x**2)
squares

# 等同
squares = list(map(lambda x: x**2, range(10)))

# 等同
squares = [x**2 for x in range(10)]
```

_使用 `for if` 组合_

```python
combs = [(x, y) for x in [1,2,3] for y in [3,1,4] if x != y]

# 等价于
combs = []
for x in [1,2,3]:
    for y in [3,1,4]:
        if x != y:
            combs.append((x, y))

combs
```



### `del` 语句

*可删除 `list` 的元素。*

```powershell
>>> a = [-1, 1, 66.25, 333, 333, 1234.5]
>>> del a[0]
>>> a
[1, 66.25, 333, 333, 1234.5]
>>> del a[2:4]
>>> a
[1, 66.25, 1234.5]
>>> del a[:]
>>> a
[]
```



### 元组 和 序列

> _Tuples and Sequences_

_元组是由逗号分隔的多个值组成的变量。它们可创建，但是不可对其赋值。元组和list类似，但是两者的运用场景不一致，前者不可变，后者则反之。_

```powershell
t = 12, 5.4, 'Joshua Conero'    # 创建元组
t[2]							# 'Joshua Conero'; 获取指定元组
# t[2] = 7.5	Error: 不可赋值
t								# (12, 5.4, 'Joshua Conero')

# 元组嵌套
u = t, (1, 2), (False, True)
u		# ((12, 5.4, 'Joshua Conero'), (1, 2), (False, True))

empty = ()			# 定义空元组
empty = 8,			# (8,)

# 解构元组
x, y , z = t
y					# 5.4
```



### Sets/集合

`Sets`是一个无序的集合，没有重复的元素。它也支持集合运算，如: 并集、交集、差集等

```python
basket = {'apple', 'orange', 'apple', 'pear', 'orange', 'banana'}
print(basket)                      # show that duplicates have been removed
# -> {'orange', 'banana', 'pear', 'apple'}

'orange' in basket                 # fast membership testing
# -> True
'crabgrass' in basket
# -> False

# Demonstrate set operations on unique letters from two words

a = set('abracadabra')
b = set('alacazam')
a                                  # unique letters in a
# -> {'a', 'r', 'b', 'c', 'd'}
a - b                              # letters in a but not in b(a 与 b 的差集)
# -> {'r', 'd', 'b'}
a | b                              # letters in a or b or both(a 与 b 的并集)
# -> {'a', 'c', 'r', 'd', 'b', 'm', 'z', 'l'}
a & b                              # letters in both a and b(a 与 b 的交集)
# -> {'a', 'c'}
a ^ b                              # letters in a or b but not both(a 与 b 的余集)
# -> {'r', 'd', 'b', 'm', 'z', 'l'}
```



支持与 list 对应的运算符

```python
a = {x for x in 'abracadabra' if x not in 'abc'}
a		# {'r', 'd'}; 
```



### 字典/dick

> `map 类型字典`(k-v 类型)

*键值(key)作为索引，由不可变类型充当，如：字符串、数字、仅仅包含字符串/数组的__元组__。*

```python
tel = {'jack': 4098, 'sape': 4139}
tel['guido'] = 4127
tel		# {'sape': 4139, 'guido': 4127, 'jack': 4098}

tel['jack']      # 4098; 获取值
del tel['jack']  # 删除 jack 键值
list(tel.keys()) # ['sape', 'guido']
'jack' in tel	 # False; 键值存在判断


# dick 构造字典
dict([('sape', 4139), ('guido', 4127), ('jack', 4098)])
# {'sape': 4139, 'jack': 4098, 'guido': 4127}

{x: x**2 for x in (2, 4, 6)}
# {2: 4, 4: 16, 6: 36}

# 使用函数参数
dict(sape=4139, guido=4127, jack=4098)
# {'sape': 4139, 'jack': 4098, 'guido': 4127}
```



> **遍历字典**

*使用 `items()` k-v 值；*

```python
knights = {'gallahad': 'the pure', 'robin': 'the brave'}
for k, v in knights.items():
    print(k, v)
```



*`enumerate()`函数变量 list， index-value*

```python
for i, v in enumerate(['tic', 'tac', 'toe']):
    print(i, v)
```



*同时变量两个序列*

```python
>>> questions = ['name', 'quest', 'favorite color']
>>> answers = ['lancelot', 'the holy grail', 'blue']
>>> for q, a in zip(questions, answers):
...     print('What is your {0}?  It is {1}.'.format(q, a))
...
What is your name?  It is lancelot.
What is your quest?  It is the holy grail.
What is your favorite color?  It is blue.
```



## Modules/模块

_**模块** 是指包含的 python 语句或表达式，且以`.py`后缀结尾的文件。模块的名字(文件的名字,字符串)可通过全局变量`__name__`获取。_

模块通过`import` 引入到指定的模块中，各个模块之间的变量为其内部私有，但可通过`modname.itemname`访问。

```python
# 从 A模块中导出对象
from A from a1,a2
# 导出所有对象，处以“_”开头的变量
from B from *
# 别名导出
import A as Ax
from C import a1 as c1
```



*python 解析器执行的脚本，其内部的`__name__="__main__"`*

```python
# python fibo.py <arguments>
if __name__ == "__main__":
    print(__name__ + "被设置为： __main__")
```



*导入模块是的目录搜索顺序，首先是built-in里面搜索，为找到时搜索对应的模块名字：`built-in`, `<name>.py`。*



**python 编译文件**

为加快模块的加载速度，python 提供模块下`__pycache__`的缓存目录，缓存的文件格式`module.version.pyc`，如： python 3.6 缓存 suba模块的缓存文件 `suba.cpython-36.pyc`。python会自动根据编译的版本决定是否需要重新缓存，另外其缓存的文件是平台独立(跨平台)的。

两种情况下python不会检测/编译缓存文件；

1. 直接加载到命令行(python解析器直接读取的文件/入口文件)
2. 不存在源模块时，即非标准库。



### 标准模块

> Standard Modules

标准库包括 *built-in* 和 语言核心库。



### package/包

包为不同的模块，多模块集合提供“点操作规则”的命名空间。

包中必须包含`__init__.py` 文件，类似于go语言中包的`func init()` 函数，初始化时会调用包。

```python
from package import item
import item.subitem.subsubitem

# 与前面的模块导入不同
# 起导出的仅仅是 package/__init__.py 中 __all__ = ['mod1', 'mod2', 'mod3'] 定义的包
# 若 __all__ 未定义将不会导出如何模块
from package import *
```



使用“.”调用父级模块,如：

- A
  - ac
  - at
  - B
    - bb
    - b1
  - C
    - c1

```python
# bb -> b1
from . import bb
# ac -> b1
from .. import ac
# c1 -> b1
from ..C import c1
```







// @TODO [7.Input and Output](https://docs.python.org/3.6/tutorial/inputoutput.html)

