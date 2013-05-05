#python学习笔记
**title:**python学习笔记  
**tags:**python,中文文档,中文笔记  
**info:**简单的整理下python的基本概念

##Python 数据类型

Python 类型 Number,String,List,Dictionary,Tuple,File  

###Python运算符和表达式优先级

* x or y                     逻辑或（仅在x为false时，对y求值）
* lambda args:表达式         匿名函数
* x and y                    逻辑与（仅在x为true时，对y求值）
* not x                      逻辑非
* <,<=,>,>=,= =,<>,!=        比较测试
* is,is not                  识别测试
* in,not in                  是否为成员测试
* x | y                      按位或
* x^y                        按位异或
* x&y                        按位与
* x<<y,x>>y                  x向左或向右移y位
* x+y,x-y                    相加/相连，相减
* `x*y,x/y,x%y`                相乘/重复，相除，求余/格式化
* -x,+x,～x    ，            同加，按位补取
* x[i],x[i:j],x.y,x(...)     索引，分片，限定，函数调用
* (...),[...],{...},'...'    Tuple，list，dictionary，转化为string


###数值函数

* abs(x)  
    返回数的绝对值，忽略正负，若x位复数，返回复数模
* coerce(x,y)  
    用普通表达式规则将x和y的变成常用类型，返回tuple值。例如语句coerce(2,3.5),返回(2.0,3.5)
* divmod(x,y)  
    用x除以y，返回包含商和余数的tuple值。等价于（a/b，a%b）
* pow(x,y[,z])  
    x的y次幕。返回值类型与x同
* round(x,[y])  
    将浮点数x舍入城想数为0位。加上参数y后摄入详述位y位。值依然为浮点数。

###字符串

字符串变量可以相'+'来连接字符串变量  
字符串与其他对象连接表达用 ',' 或这是用格式化  
可使用'*'来重复字符串，例如：'this is ha'*5   
是用len(string)来计算大小  

####字符串是数组

利用string[strat:end]来分片(slicing),  
所得结果不包括end很end后所有的字符串，  
若Start和end均为负数，则索引从字符串尾部开始，最后一个为-1.  

###格式化

使用'%'来格式化字符串和其他相同的基本符号的对象。  
运算符%的转换格式

* %%    百分号标志
* %c    字符及其ASCII码
* %s    字符串，在打印前将相对应的对象转换为字符串，因此%s可运行于任何对象
* %d    有符号整数（十进制）
* %u    无符号整数（十进制）
* %o    无符号整数（八进制）
* %x    无符号整数（十六进制）
* %X    无符号整数（十六进制大写）
* %e    浮点数字（科学计数法）
* %E    浮点数字（科学计数法用E代替e）
* %f    浮点数字（用小数点符号）
* %g    浮点数字（根据值的大小采用%e或%f）
* %G    浮点数字（类似%g，在合适的位置用E代替e）
* %p    指针（用十六进制打印值的内存地址）
* %n    存储输出字符的数量放进参数列表的下一个变量中

####可选格式化标记  
运算符%的转换格式 
 
* 空格：  用空格做证书的前缀
* +：     用加好做正数的前缀
* —：     区域左对齐
* 0：     用0而非空格进行右对齐
* #：     用0做做非零八进制数前缀，0x做十六进制的前缀
* 数字：  最小宽度
* 数字：  制定浮点精度（小数点后的位数）

>     >>> album={'title':'Flood','id':56}
>     >>> print 'Catalog Number %(id)05d is %(title)s ' % album
>     Catalog Number 00056 is Flood 

####转义字符

* \(在行尾) ： 继续
* \\：饭斜线符号
* \'：单引号
* \"：双引号
* \a：Bell
* \b：退格键
* \e：转义
* \000：空，Python字符串不以空结束
* \n：换行
* \v：纵向制表符
* \t：水平制表符
* \r：回车
* \f：换页
* \0yy：八进制数yy代表的字符（例如\012等价于新起一行）
* \xyy：十六进制yy代表字符（例如\x0a等价于新起一行）
* \y：以上未列出的任何字符y以普通格式输出

####原始字符串 
使转义字符无效 用 ‘r’和‘R’定义原始字符串  例如：r'\r\n\x99'

>     >>> print r'\r\n\x99'
    \r\n\x99

###List列表 
>list[]

* 赋值 list1[3:4]=[a,b]  
* len(list)长度  
* del list 删除对象  

列表对象支持的方法

* append(x)    尾部追加 单个对象x，使用多个对象会引起异常。
* count(x)     返回对象x在list中出现的次数
* extend(L)    将列表L中的项添加到表中
* index(x)     返回匹配对象x第一个表项的索引，无匹配时产生异常
* insert(i,x)  在索引‘i’的元素钱插入对象x
* pop(x)       删除列表中索引x的表项，并返回同该表项的值，无参数删除最后
* remove(x)    删除表匹配对象x的第一个元素，无匹配时异常
* reverse()    颠倒列表元素的顺序
* sort()       对列表排序

###元祖
>Tuples  ()

tuples是不变的，列表是可变的，一旦创建tuple便不可修改了，除非新建立。

对于序列  
__隶属__  
    x in list 在列表list中找到x对象。找到返回1，位匹配返回0

迭代  

>     >>> l=[14,22,12,456,53]
>     >>> for i in l:
>            print i

###词典
>dictionary{name:value,...}

###字典的方法
* has_keys(x)       若字典中有x返回true
* keys()            返回键的列表
* values()          返回值的列表
* dict.items()      返回tuples的列表。每个tuple有字典的dict的键和相应的值组成
* clear()           删除词典的所有条目
* copy()            返回字典的高层结构的拷贝，但不复制嵌入结构，而复制那些结构的引用。
* update(x)         用字典x中的键/值对更新字典的内容。
* get(x[,y])        返回键x。若未找到返回None

###内置对象类型转换

* str(x)              将对象x翻译为字符串
* list(x)             将对象序列x作为列表返回。例如‘hello’返回['h','e','l','l','o'],将tuple转换为列表
* tuple(x)            将对象序列x作为tuple返回
* int(x)              将字符串和数字转换为整数,对浮点进行舍位而非舍入
* long(x)             将字符串和数字转换为长整形
* float(x)            将str和num转换为浮点对象
* complex(x,y)        将x做实部，y做虚部创建复数
* hex(x)              将整数或长整数转换为十六进制字符串
* oct(x)              将整数或长整数转换为八进制字符串
* ord(x)              返回字符x的ASCII值
* chr(x)              返回ASCII码x代表的字符
* min(x[,...])        返回序列中最小的元素
* max(x[,...])        返回序列中最大的元素

###比较运算符

* x is y              指向同一对象
* x is not y          不同对象
* not y               取反
* x or y；x and y     逻辑取值
* x<y<z               链式比较

###对象和常量的真、假值
* ''            false
* 'string'      true
* 0             false
* >1            true
* <-1           true
* ()            false
* []            false
* {}            false
* None          false

###赋值

基本赋值、tuple赋值、列表赋值和多目标赋值。

##语句

###控制语句

    if expression:
        block
    elif expression2:
        block
    else:
        block

    while expression:
        block
    else:
        block

    for target in object:
        block
    else:
        block

###范围

range函数产生值列表  
    range([start,]stop[,step])

###循环控制语句

* break  
    退出当前循环，并忽略任何else语句，而继续执行循环语句块最后一行后面的行
* continue  
    强制循环立即进入下一次的循环，护绿当前愉快中的其他语句，对循环语句表达式重新求职
* pass  
    不做任何事情

###普通陷阱

* 变量名
* 使用前不需事先声明，根据分配的值设置对象类型
* 变量名为字母和数字的混排
* 不需要用特殊字符对变量限定
* 有些变量有操作对象内容的内嵌方法
* 块和缩排
* 块以前面的冒号开始
* 在创建新的块以前，或者船舷标志块结束的缩排格式以前，该块一直继续
* 当用tab和空格缩进时，不要混合使用两种方法，建议使用空格

##函数

__定义__ 

>    def name(arg1[,...]):
        startement block
        [return value]

定义全局变量和局域变量 global 来定义全局变量

###LGB 规则
名字引用一次搜寻3个作用域：__局部（Local）__，__全局（Global）__，__内置（Built-in）__（LGB）   
在局部作用域内对名字赋值时会创建新对象或更新对象。若在局部作用域对全局对象赋值鼻血使用关键字global  

###参数

函数中定义的指定的参数按照其定义的顺去依次被接收。

参数传递过程遵循的规则如下：  
通过引用将参数复制到局部作用域的对象中。这以为着用来访问函数的变量于提供给函数的对象无关，并且修改局部对象不会改变原始参数。  
可以适当位置改变对象。当人们复制列表或字典时，就复制了对象列表的引用同，如果改变引用的值，则修改了原始的参数。

####参数默认值
使用 arg=value   为参数定义默认的value，并且尽量将有默认value的对象后置，这样可以有效的减少错误。

####参数Tuple

接受长参数，不想使用默认值，可以采用tuple参数*arg可将接受参数作为tuple使用  
例如 
    def function(list,*items):...

####参数字典

使用　**arg　来接收参数  
例如：
>     cdict(**item):
>     ...     for field in item.keys():
>     ...             print field+":",item[field]
>     ...     print
>     >>> cdict(f="wo",t="you",msg="ceshineirong")
>     msg: ceshineirong
>     t: you
>     f: wo

####返回 return

返回值使用 __return vars__;

####函数规则：

* 默认参数值必须放在非默认参数后
* 单个函数只能使用一个tuple和一个字典函数
* tuple参数必须在连续参数和默认参数之后
* 字典参数必须是最后的参数。

####调用函数：

* 函数的调用必须带有圆括号
* 按照顺序为关键字参数赋值
* 按照参数名字的顺序为参数赋值
* 将额外的参数赋值给tuple参数或字典参数
* 将默认值赋予任何未赋值的参数

###高级函数的调用

* __apply语句__   
    apply(function,tuple)  调用函数，将参数tuple应用户函数function和显现式调用一致
* __map语句__  
    使用map函数程序员完成循环操作和重复复制的操作。  
    例子：将一个数字列表转化为其立方的列表。

>     ---普通方法---
>     >>> list = [1,2,3,4]
>     >>> for index in range(len(list)):
>     ...     list[index] = pow(list[index],3)
>     ... 
>     >>> list
>     [1, 8, 27, 64]
>
>     --- map ---
>
>     >>> list=[1,2,3,4]
>     >>> def cube(x):
>     ...     return pow(x,3)
>     ... 
>     >>> list = map(cube,list)
>     >>> list
>     [1, 8, 27, 64]
>
>    --- map & lambda ----
>
>     >>> list = [1,2,3,4]
>     >>> list = map(lambda x:pow(x,3),list)
>     >>> list
>     [1, 8, 27, 64]
>
>     ------= over ------

####间接函数调用

x=functionName 调用x()就可以调用函数function

####匿名函数lambda

lambda arg[,arg,...]:expr



