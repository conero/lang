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

// @TODO [5. Data Structures](https://docs.python.org/3/tutorial/datastructures.html)

