幂的表示
在计算机数学公式中，通常用`^`来表示，C、Java、Go 等大多数主流编程语言则使用`pow()`，Python 则使用 `**`。Python的`**`会自动根据整型或浮点型计算，`math.pow()`则只会输出浮点数，另外，`pow(x, y, z)`表示 (x 的 y 次方) 对 z 取模，即 `pow(2, 10000000, 13)` 等价于 `(2 ** 10000000) % 13`，这在密码学中非常方便。


身份运算符：`is`, `is not`
成员运算符：`in`, `not in`
逻辑运算符：`and`, `or`, `not`，而不是常用的 `&&`, `||`, `!`

Python 与 Go 的 `:=`
||Go |Python |
|---|---|---|
|官方名称|短变量声明|赋值表达式 (海象运算符)|
|主要目的|函数内创建新变量 (替代 var)|在判断/计算中顺便赋值|
|是否返回值|否 (它是语句)|是 (它是表达式)|
|作用域|严格限制在当前代码块 {} 内|通常会“泄漏”到当前函数或模块|

与PHP、JavaScript不同，Python中`==`即为严格比较，而非`===`

Python中的布尔值`True`和`False`也是首字母大写，而不是其他语言中的小写

除Python和Shell中使用`elif`外，其他主流编程语言如C、Java、Go、JavaScript等使用`else if`

Python 中默认的编码格式是 ASCII，因此通常在文件开头加入 `# -*- coding: UTF-8 -*-` 或 `# coding=utf-8`（注意=两边没有空格）来兼容非英文字符。
`#`表示单行注释，`"""`表示多行注释
没有`const`常量关键字，通过全大写变量名约定来实现常量
三元运算使用的是 `max = a if a > b else b` ，而非 `max = a > b ? a : b`
Python 不支持 `switch-case`，只支持 `match-case`
拥有高效的推导式`[表达式 for 迭代变量 in 可迭代对象 [if 条件表达式]]`
函数定义关键字是 `def`，而非 `function` 或 `func`

就像Linux中一切皆文件一样，Python中，一切皆对象

Python 中只有 `for-in` 和 `while`，没有普通的 `for`

在Python中，字符串前用 `r` 表示原生输出字符串，`f` 表示格式化输出，可解析字符串中的变量，如`f"Hello, {full_name.title()}!"`

Python中的序列分为
- 有序
  - 列表：可以修改，支持运算
  - 元组：不可修改
- 无序
  - 字典
  - 集合

|  | 特点 | 查 | 增 | 改 | 删 | 备注 |
| --- | --- | --- | --- | --- | --- | --- |
| 列表 | 有序 | ・`.index(value, start, end)` 查找元素的索引
・`.count()` 统计元素出现的次数 | ・`.append(value)`末尾整体追加元素
・`.extend(value)`末尾逐个追加元素
・`.insert(index, value)` 在指定位置整体插入元素
・`+` 拼接列表 | ・根据下标修改；
・通过切片修改； | ・`del listname[index]` 或`del listname[start : end]` 根据索引删除
・`.pop(index)`根据索引删除
・`.remove(value)` 根据元素值删除
・`.clear()` 清空元素 | `range()` 函数的返回值并不直接是列表，而是 range |
| 元组 | 有序、不可变 | ・通过索引
・通过切片 | ・`()`
・`tuple()` | 无 | `del tuplename` | 当创建的元组中只有一个字符串类型的元素时，该元素后面必须要加一个逗号 `,`，否则 Python 解释器会将它视为字符串。 |
| 字典 | 无序、键值对 | ・`dictname[key]`通过索引查询，键不存在时报错
・`dictname.get(key[, 键不存在时的默认值])` 键不存在时不报错，可设置默认值
・`in`或`not in` 查询是否存在值
・`for k in dictname.keys():` 或 `for k in dictname:` 获取所有键（`keys()`会返回一个列表）
・`for v in dictname.values():` 获取所有值 
・`for k, v in dictname.items():` 来遍历字典 | ・`{}`
・`dict()`
・`fromkeys` 创建带默认值的字典
・`dictname[key] = value` 通过键值对添加 | ・`dictname[key] = value` 通过键值对修改
・`update` 包含则更新，否则添加 | ・`del dictname[key]` 根据键删除
・`.pop()` 删除指定的键值对
・`popitem()` 移除末尾的一个键值对
 | `setdefault`为不存在的 key 设置一个默认值 |
| 集合 | 无序、唯一 | ・遍历读取 | ・`.add(value)` | 无 | ・`.remove()`不存在元素报错
・`.discard()`不存在元素不报错
・`.clear()` | `forzenset` 不允许变更元素，`set` 中变更元素的方法都不支持，不变更元素的方法都支持。 |

# 推导式

- 列表`[表达式 for 迭代变量 in 可迭代对象 [if 条件表达式]]`
- 元组`(表达式 for 迭代变量 in 可迭代对象 [if 条件表达式])` ，元组推导式生成的结果是生成器对象，使用 `for` 或 `__next__` 遍历生成器对象后，原值将被清理
- 字典`{表达式 for 迭代变量 in 可迭代对象 [if 条件表达式]}`
- 集合`{表达式 for 迭代变量 in 可迭代对象 [if 条件表达式]}`

在 Python 中，通过子类名后的括号标注父类来表示继承，如`class SubClass(ParentClass):`，这在PHP、Java、JavaScript中则是通过`extends`关键字实现。如果子类继承多个父类，父类间通过`,`分隔

在Python中，模块名是由文件名来确定的（Java用反向域名，如`com.example.myapp`；PHP用`namespace`，如`MyCompany\MyPackage`），但不一定与文件名完全相同（Python中允许一个文件中放多个类）。当在一个名为`moudle`的目录下放置一个`__init__.py`的空文件时，`moudle`的整个目录都会成为Python包，包内所有文件都被视为包的模块

- 导入整个模块：`import (as alias)`
- 导入模块特定的函数：`from module_name import function0, function1, function2 (as alias)`


以下划线开头的标识符有特殊含义，例如：

- 以单下划线开头的标识符（如 `_width`），表示不能直接访问的类属性，其无法通过 `from...import *` 的方式导入；
- 以双下划线开头的标识符（如`__add`）表示类的私有成员；
- 以双下划线作为开头和结尾的标识符（如 `__init__`），是专用标识符。

Python中通过 `if __name__ == '__main__':` 来指定整个程序的入口点，这里的`__name__`和`__main__`只是一个内部变量，前者表示当前模块，后者表示当前模块作为主程序入口，而非导入为一个模块，二者并非要满足相等条件，只是一种声明写法。
```py
if __name__ == '__main__':
    print_hi('PyCharm')
```

Python 的环境管理比较混乱，有 venv、poetry、uv 等，uv 则是目前最流行与方便的工具，不仅具有 rust 的高效，而且能自动管理虚拟环境，还有 ruff 生态。
Python有强大的绘图工具Matplotlib，数学计算工具NumPy，游戏制作工具PyGame