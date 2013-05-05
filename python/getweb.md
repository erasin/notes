返回小站
参考文献
RSS 归档
<http://zhan.renren.com/reference?gid=3602888497995300069&checked=true>

Python下载网页和处理正则的几个要点
这是最简洁的一种，当然也是Get的方法 

    fd    =  urllib2.urlopen(url_link) 
    data =  fd.read()

匹配任意字符（包括换行符） 
(.*)    #无法获得换行之后的文本，匹配的是除了换行符“\n”以外的所有字符
([\s\S]*)  #可以匹配任意字符，包括换行符
Python读写文件 
 
open，使用open打开文件后一定要记得调用文件对象的close()方法。比如可以用try/finally语句来确保最后能关闭文件： 

    file_object = open('thefile.txt')
    try:
         all_the_text = file_object.read( )
    finally:
         file_object.close( )

注：不能把open语句放在try块里，因为当打开文件出现异常时，文件对象file_object无法执行close()方法。
 
2
通过GET的方法
def GetHtmlSource(url):
    try:
    
        htmSource = ''
        
        req = urllib2.Request(url)
       
        fd = urllib2.urlopen(req,"")
        
        while 1:
            data = fd.read(1024)
            if not len(data):
                break
            
            htmSource += data

        fd.close()
  
        del fd
        del req
        
        htmSource = htmSource.decode('cp936')
        htmSource = formatStr(htmSource)
         
        return htmSource
        
    except socket.error, err:
        
        str_err =  "%s" % err
        return ""
3
 通过GET的方法
def GetHtmlSource_Get(htmurl):
    htmSource = ""
    
    try:
        
        urlx = httplib.urlsplit(htmurl)
        conn = httplib.HTTPConnection(urlx.netloc)
        conn.connect()
        conn.putrequest("GET", htmurl, None)
        
        conn.putheader("Content-Length", 0)
        conn.putheader("Connection", "close")
        conn.endheaders()
        
        res = conn.getresponse()
        htmSource = res.read()
        
    except Exception(), err:
        trackback.print_exec()
        conn.close()
        
        
    return htmSource
 
通过POST的方法 
def GetHtmlSource_Post(getString):
    htmSource = ""
    
    try:
        url = httplib.urlsplit("http://app.sipo.gov.cn:8080")
        conn = httplib.HTTPConnection(url.netloc)
        conn.connect()
        conn.putrequest("POST", "/sipo/zljs/hyjs-jieguo.jsp")
        conn.putheader("Content-Length", len(getString))
        conn.putheader("Content-Type", "application/x-www-form-urlencoded")
        conn.putheader("Connection", " Keep-Alive")
        conn.endheaders()
        conn.send(getString)
        f = conn.getresponse()
        if not f:
            raise socket.error, "timed out"
        
        htmSource = f.read()
        f.close()
        conn.close()
         
        return htmSource
        
    except Exception(), err:
        
        trackback.print_exec()
        conn.close()
        
        
    return htmSource
 
2.读文件
读文本文件
input = open('data', 'r')
#第二个参数默认为r 
 input = open('data')
 
读二进制文件
input = open('data', 'rb')
 
读取所有内容
file_object = open('thefile.txt')
try:
     all_the_text = file_object.read( )
finally:
     file_object.close( )
 
读固定字节
file_object = open('abinfile', 'rb')
try:
    while True:
         chunk = file_object.read(100)
        ifnot chunk:
            break 
         do_something_with(chunk)
finally:
     file_object.close( )
 
读每行
list_of_all_the_lines = file_object.readlines( )

如果文件是文本文件，还可以直接遍历文件对象获取每行：
for line in file_object:
     process line
 
3.写文件
写文本文件
    output = open('data', 'w')
 
写二进制文件
    output = open('data', 'wb')
 
追加写文件
    output = open('data', 'w+')
 
写数据
    file_object = open('thefile.txt', 'w')
    file_object.write(all_the_text)
    file_object.close( )
 
写入多行
    file_object.writelines(list_of_text_strings)

注意，调用writelines写入多行在性能上会比使用write一次性写入要高。
在处理日志文件的时候，常常会遇到这样的情况：日志文件巨大，不可能一次性把整个文件读入到内存中进行处理，例如需要在一台物理内存为 2GB 的机器上处理一个 2GB 的日志文件，我们可能希望每次只处理其中 200MB 的内容。
在 Python 中，内置的 File 对象直接提供了一个 readlines(sizehint) 函数来完成这样的事情。以下面的代码为例：
 
    file = open('test.log', 'r')
    sizehint =209715200   # 200M 
    position = 0
    lines = file.readlines(sizehint)
    whilenot file.tell() - position < 0:
         position = file.tell()
         lines = file.readlines(sizehint)
 

每 次调用 readlines(sizehint) 函数，会返回大约 200MB 的数据，而且所返回的必然都是完整的行数据，大多数情况下，返回的数据的字节数会稍微比 sizehint 指定的值大一点（除最后一次调用 readlines(sizehint) 函数的时候）。通常情况下，Python 会自动将用户指定的 sizehint 的值调整成内部缓存大小的整数倍。

file在python是一个特殊的类型，它用于在python程序中对外部的文件进行操作。在python中一切都是对象，file也不例外，file有file的方法和属性。下面先来看如何创建一个file对象：

    file(name[, mode[, buffering]])

file()函数用于创建一个file对象，它有一个别名叫open()，可能更形象一些，它们是内置函数。来看看它的参数。它参数都是以字符串的形式传递的。name是文件的名字。

mode 是打开的模式，可选的值为r w a U，分别代表读（默认） 写 添加支持各种换行符的模式。用w或a模式打开文件的话，如果文件不存在，那么就自动创建。此外，用w模式打开一个已经存在的文件时，原有文件的内容会被清 空，因为一开始文件的操作的标记是在文件的开头的，这时候进行写操作，无疑会把原有的内容给抹掉。由于历史的原因，换行符在不同的系统中有不同模式，比如 在 unix中是一个\n，而在windows中是‘\r\n’，用U模式打开文件，就是支持所有的换行模式，也就说‘\r’ '\n' '\r\n'都可表示换行，会有一个tuple用来存贮这个文件中用到过的换行符。不过，虽说换行有多种模式，读到python中统一用\n代替。在模式 字符的后面，还可以加上+ b t这两种标识，分别表示可以对文件同时进行读写操作和用二进制模式、文本模式（默认）打开文件。

buffering如果为0表示不进行缓冲;如果为1表示进行“行缓冲“;如果是一个大于1的数表示缓冲区的大小，应该是以字节为单位的。

file对象有自己的属性和方法。先来看看file的属性。

    closed #标记文件是否已经关闭，由close()改写
    encoding #文件编码
    mode #打开模式
    name #文件名
    newlines #文件中用到的换行模式，是一个tuple
    softspace #boolean型，一般为0，据说用于print

file的读写方法：

F.read([size]) 
:   size为读取的长度，以byte为单位

F.readline([size]) 
:   读一行，如果定义了size，有可能返回的只是一行的一部分

F.readlines([size]) 
:   把文件每一行作为一个list的一个成员，并返回这个list。其实它的内部是通过循环调用readline()来实现的。如果提供size参数，size是表示读取内容的总长，也就是说可能只读到文件的一部分。

F.write(str) 
:   把str写到文件中，write()并不会在str后加上一个换行符

F.writelines(seq) 
:   把seq的内容全部写到文件中。这个函数也只是忠实地写入，不会在每行后面加上任何东西。

file的其他方法：

F.close() 
:   关闭文件。python会在一个文件不用后自动关闭文件，不过这一功能没有保证，最好还是养成自己关闭的习惯。如果一个文件在关闭后还对其进行操作会产生ValueError

F.flush() 
:   把缓冲区的内容写入硬盘

F.fileno() 
:   返回一个长整型的”文件标签“

F.isatty() 
:   文件是否是一个终端设备文件（unix系统中的）

F.tell() 
:   返回文件操作标记的当前位置，以文件的开头为原点

F.next() 
:   返回下一行，并将文件操作标记位移到下一行。把一个file用于for ... in file这样的语句时，就是调用next()函数来实现遍历的。

F.seek(offset[,whence]) 
:   将文件打操作标 记移到offset的位置。这个offset一般是相对于文件的开头来计算的，一般为正数。但如果提供了whence参数就不一定了，whence可以为 0表示从头开始计算，1表示以当前位置为原点计算。2表示以文件末尾为原点进行计算。需要注意，如果文件以a或a+的模式打开，每次进行写操作时，文件操 作标记会自动返回到文件末尾。

F.truncate([size]) 
:   把文件裁成规定的大小，默认的是裁到当前文件操作标记的位置。如果size比文件的大小还要大，依据系统的不同可能是不改变文件，也可能是用0把文件补到相应的大小，也可能是以一些随机的内容加上去。

 
python类型转换、数值操作
关键字: python类型转换、数值操作 python类型转换
 
 
Java代码
 
函数                      描述  
int(x [,base ])         将x转换为一个整数  
long(x [,base ])        将x转换为一个长整数  
float(x )               将x转换到一个浮点数  
complex(real [,imag ])  创建一个复数  
str(x )                 将对象 x 转换为字符串  
repr(x )                将对象 x 转换为表达式字符串  
eval(str )              用来计算在字符串中的有效Python表达式,并返回一个对象  
tuple(s )               将序列 s 转换为一个元组  
list(s )                将序列 s 转换为一个列表  
chr(x )                 将一个整数转换为一个字符  
unichr(x )              将一个整数转换为Unicode字符  
ord(x )                 将一个字符转换为它的整数值  
hex(x )                 将一个整数转换为一个十六进制字符串  
oct(x )                 将一个整数转换为一个八进制字符串  
 
函 数 描述int(x [,base ]) 将x转换为一个整数long(x [,base ]) 将x转换为一个长整数float(x ) 将x转换到一个浮点数complex(real [,imag ]) 创建一个复数str(x ) 将对象 x 转换为字符串repr(x ) 将对象 x 转换为表达式字符串eval(str ) 用来计算在字符串中的有效Python表达式,并返回一个对象tuple(s ) 将序列 s 转换为一个元组list(s ) 将序列 s 转换为一个列表chr(x ) 将一个整数转换为一个字符unichr(x ) 将一个整数转换为Unicode字符ord(x ) 将一个字符转换为它的整数值hex(x ) 将一个整数转换为一个十六进制字符串oct(x ) 将一个整数转换为一个八进制字符串

序列支持一下操作：
 
 
Python代码
 
操作                      描述  
s + r                   序列连接  
s * n , n * s           s的 n 次拷贝,n为整数  
s % d                   字符串格式化(仅字符串)  
s[i]                    索引  
s[i :j ]                切片  
x in s , x not in s     从属关系  
for x in s :            迭代  
len(s)                  长度  
min(s)                  最小元素  
max(s)                  最大元素  
s[i ] = x               为s[i]重新赋值  
s[i :j ] = r            将列表片段重新赋值  
del s[i ]               删除列表中一个元素  
del s[i :j ]            删除列表中一个片段  
 
操 作 描述s + r 序列连接s * n , n * s s的 n 次拷贝,n为整数s % d 字符串格式化(仅字符串)s[i] 索引s[i :j ] 切片x in s , x not in s 从属关系for x in s : 迭代len(s) 长度min(s) 最小元素max(s) 最大元素s[i ] = x 为s[i]重新赋值s[i :j ] = r 将列表片段重新赋值del s[i ] 删除列表中一个元素del s[i :j ] 删除列表中一个片段

数值操作：
python进行文件读写的函数是open或file
file_handler = open(filename,,mode）
Table mode
模式
描述
r
以读方式打开文件，可读取文件信息。
w
以写方式打开文件，可向文件写入信息。如文件存在，则清空该文件，再写入新内容
a
以追加模式打开文件（即一打开文件，文件指针自动移到文件末尾），如果文件不存在则创建
r+
以读写方式打开文件，可对文件进行读和写操作。
w+
消除文件内容，然后以读写方式打开文件。
a+
以读写方式打开文件，并把文件指针移到文件尾。
b
以二进制模式打开文件，而不是以文本模式。该模式只对Windows或Dos有效，类Unix的文件是用二进制模式进行操作的。
 
Table 文件对象方法
方法
描述
f.close()
关闭文件，记住用open()打开文件后一定要记得关闭它，否则会占用系统的可打开文件句柄数。
f.fileno()
获得文件描述符，是一个数字
f.flush()
刷新输出缓存
f.isatty()
如果文件是一个交互终端，则返回True，否则返回False。
f.read([count])
读出文件，如果有count，则读出count个字节。
f.readline()
读出一行信息。
f.readlines()
读出所有行，也就是读出整个文件的信息。
f.seek(offset[,where])
把文件指针移动到相对于where的offset位置。where为0表示文件开始处，这是默认值 ；1表示当前位置；2表示文件结尾。
f.tell()
获得文件指针位置。
f.truncate([size])
截取文件，使文件的大小为size。
f.write(string)
把string字符串写入文件。
f.writelines(list)
把list中的字符串一行一行地写入文件，是连续写入文件，没有换行。
 
例子如下：
读文件
 
 
 
Python代码
read = open(result)  
       line=read.readline()  
       while line:  
             print line  
             line=read.readline()#如果没有这行会造成死循环  
       read.close  
read = open(result) line=read.readline() while line: print line line=read.readline()#如果没有这行会造成死循环 read.close
 写文件
 
 
 
Python代码
read = file(result,'a+')  
        read.write("\r\n")  
        read.write("thank you")  
        read.close  
read = file(result,'a+') read.write("\r\n") read.write("thank you") read.close
 其它
 
 
 
Python代码
#-*- encoding:UTF-8 -*-  
filehandler = open('c:\\111.txt','r')    #以读方式打开文件，rb为二进制方式(如图片或可执行文件等)  
  
print 'read() function:'              #读取整个文件  
print filehandler.read()  
  
print 'readline() function:'          #返回文件头，读取一行  
filehandler.seek(0)  
print filehandler.readline()  
  
print 'readlines() function:'         #返回文件头，返回所有行的列表  
filehandler.seek(0)  
print filehandler.readlines()  
  
print 'list all lines'                #返回文件头，显示所有行  
filehandler.seek(0)  
textlist = filehandler.readlines()  
for line in textlist:  
    print line,  
print   
print  
  
print 'seek(15) function'               #移位到第15个字符，从16个字符开始显示余下内容  
filehandler.seek(15)  
print 'tell() function'  
print filehandler.tell()              #显示当前位置  
print filehandler.read()  
  
filehandler.close()                   #关闭文件句柄 
 
Python代码
 
x << y                  左移  
x >> y                  右移  
x & y                   按位与  
x | y                   按位或  
x ^ y                   按位异或 (exclusive or)  
~x                      按位翻转  
x + y                   加  
x - y                   减  
x * y                   乘  
x / y                   常规除  
x // y                  地板除  
x ** y                  乘方 (xy )  
x % y                   取模 (x mod y )  
-x                      改变操作数的符号位  
+x                      什么也不做  
~x                      ~x=-(x+1)  
abs(x )                 绝对值  
divmod(x ,y )           返回 (int(x / y ), x % y )  
pow(x ,y [,modulo ])    返回 (x ** y ) x % modulo  
round(x ,[n])           四舍五入，n为小数点位数  
x < y                   小于  
x > y                   大于  
x == y                  等于  
x != y                  不等于(与<>相同)  
x >= y                  大于等于  
x <= y                  小于等于  
 
[代码] [Python]代码
 
view sourceprint?
 
01	#-*- encoding: gb2312 -*-
02	import os, sys, string
03	import MySQLdb
04	 
05	# 连接数据库
06	try:
07	    conn = MySQLdb.connect(host='localhost',user='root',passwd='xxxx',db='test1')
08	except Exception, e:
09	    print e
10	    sys.exit()
11	 
12	# 获取cursor对象来进行操作
13	 
14	cursor = conn.cursor()
15	# 创建表
16	sql = "create table if not exists test1(name varchar(128) primary key, age int(4))"
17	cursor.execute(sql)
18	# 插入数据
19	sql = "insert into test1(name, age) values ('%s', %d)" % ("zhaowei", 23)
20	try:
21	    cursor.execute(sql)
22	except Exception, e:
23	    print e
24	 
25	sql = "insert into test1(name, age) values ('%s', %d)" % ("张三", 21)
26	try:
27	    cursor.execute(sql)
28	except Exception, e:
29	    print e
30	# 插入多条
31	 
32	sql = "insert into test1(name, age) values (%s, %s)"
33	val = (("李四", 24), ("王五", 25), ("洪六", 26))
34	try:
35	    cursor.executemany(sql, val)
36	except Exception, e:
37	    print e
38	 
39	#查询出数据
40	sql = "select * from test1"
41	cursor.execute(sql)
42	alldata = cursor.fetchall()
43	# 如果有数据返回，就循环输出, alldata是有个二维的列表
44	if alldata:
45	    for rec in alldata:
46	        print rec[0], rec[1]
47	 
48	 
49	cursor.close()
50	 
51	conn.close()
# python
6
分享
喜欢
 同时分享 
李斯特 2011-10-24发布 回复
 
