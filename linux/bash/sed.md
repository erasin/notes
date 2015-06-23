# sed

直接修改文件

	sed -i 's/xxx/ddd/' filename

mac 下需要g

	sed -ig 's/sss/sdd/' filename


1. Sed简介

sed 是一种在线编辑器，它一次处理一行内容。处理时，把当前处理的行存储在临时缓冲区中，称为“模式空间”（pattern space），接着用sed命令处理缓冲区中的内容，处理完成后，把缓冲区的内容送往屏幕。接着处理下一行，这样不断重复，直到文件末尾。文件内容并没有改变，除非你使用重定向存储输出。Sed主要用来自动编辑一个或多个文件；简化对文件的反复操作；编写转换程序等。以下介绍的是Gnu版本的Sed 3.02。

2. 定址

可以通过定址来定位你所希望编辑的行，该地址用数字构成，用逗号分隔的两个行数表示以这两行为起止的行的范围（包括行数表示的那两行）。如1，3表示1，2，3行，美元符号($)表示最后一行。范围可以通过数据，正则表达式或者二者结合的方式确定 。

3. Sed命令

调用sed命令有两种形式：

sed [options] 'command' file(s)
sed [options] -f scriptfile file(s)

a\ 在当前行后面加入一行文本。
b lable 分支到脚本中带有标记的地方，如果分支不存在则分支到脚本的末尾。
c\ 用新的文本改变本行的文本。
d 从模板块（Pattern space）位置删除行。
D 删除模板块的第一行。
i\ 在当前行上面插入文本。
h 拷贝模板块的内容到内存中的缓冲区。
H 追加模板块的内容到内存中的缓冲区
g 获得内存缓冲区的内容，并替代当前模板块中的文本。
G 获得内存缓冲区的内容，并追加到当前模板块文本的后面。
l 列表不能打印字符的清单。
n 读取下一个输入行，用下一个命令处理新的行而不是用第一个命令。
N 追加下一个输入行到模板块后面并在二者间嵌入一个新行，改变当前行号码。
p 打印模板块的行。
P（大写） 打印模板块的第一行。
q 退出Sed。
r file 从file中读行。
t label if分支，从最后一行开始，条件一旦满足或者T，t命令，将导致分支到带有标号的命令处，或者到脚本的末尾。
T label 错误分支，从最后一行开始，一旦发生错误或者T，t命令，将导致分支到带有标号的命令处，或者到脚本的末尾。
w file 写并追加模板块到file末尾。
W file 写并追加模板块的第一行到file末尾。
! 表示后面的命令对所有没有被选定的行发生作用。
s/re/string 用string替换正则表达式re。
= 打印当前行号码。
# 把注释扩展到下一个换行符以前。

以下的是替换标记

g表示行内全面替换。
p表示打印行。
w表示把行写入一个文件。
x表示互换模板块中的文本和缓冲区中的文本。
y表示把一个字符翻译为另外的字符（但是不用于正则表达式）

4. 选项

-e command, --expression=command
允许多台编辑。

-h, --help
打印帮助，并显示bug列表的地址。

-n, --quiet, --silent
取消默认输出。

-f, --filer=script-file
引导sed脚本文件名。

-V, --version
打印版本和版权信息。

5. 元字符集

^
锚定行的开始 如：/^sed/匹配所有以sed开头的行。

$
锚定行的结束 如：/sed$/匹配所有以sed结尾的行。

.
匹配一个非换行符的字符 如：/s.d/匹配s后接一个任意字符，然后是d。

*
匹配零或多个字符 如：/*sed/匹配所有模板是一个或多个空格后紧跟sed的行。

[]
匹配一个指定范围内的字符，如/[Ss]ed/匹配sed和Sed。

[^]
匹配一个不在指定范围内的字符，如：/[^A-RT-Z]ed/匹配不包含A-R和T-Z的一个字母开头，紧跟ed的行。

\(..\)
保存匹配的字符，如s/\(love\)able/\1rs，loveable被替换成lovers。

&
保存搜索字符用来替换其他字符，如s/love/**&**/，love这成**love**。

\<
锚定单词的开始，如:/\<love/匹配包含以love开头的单词的行。

\>
锚定单词的结束，如/love\>/匹配包含以love结尾的单词的行。

x\{m\}
重复字符x，m次，如：/0\{5\}/匹配包含5个o的行。

x\{m,\}
重复字符x,至少m次，如：/o\{5,\}/匹配至少有5个o的行。

x\{m,n\}
重复字符x，至少m次，不多于n次，如：/o\{5,10\}/匹配5--10个o的行。

6. 实例

删除：d命令
$ sed '2d' example-----删除example文件的第二行。

$ sed '2,$d' example-----删除example文件的第二行到末尾所有行。

$ sed '$d' example-----删除example文件的最后一行。

$ sed '/test/'d example-----删除example文件所有包含test的行。

替换：s命令
$ sed 's/test/mytest/g' example-----在整行范围内把test替换为mytest。如果没有g标记，则只有每行第一个匹配的test被替换成mytest。

$ sed -n 's/^test/mytest/p' example-----(-n)选项和p标志一起使用表示只打印那些发生替换的行。也就是说，如果某一行开头的test被替换成mytest，就打印它。

$ sed 's/^192.168.0.1/&localhost/' example-----&符号表示替换换字符串中被找到的部份。所有以192.168.0.1开头的行都会被替换成它自已加 localhost，变成192.168.0.1localhost。

$ sed -n 's/\(love\)able/\1rs/p' example-----love被标记为1，所有loveable会被替换成lovers，而且替换的行会被打印出来。

$ sed 's#10#100#g' example-----不论什么字符，紧跟着s命令的都被认为是新的分隔符，所以，“#”在这里是分隔符，代替了默认的“/”分隔符。表示把所有10替换成100。

选定行的范围：逗号
$ sed -n '/test/,/check/p' example-----所有在模板test和check所确定的范围内的行都被打印。

$ sed -n '5,/^test/p' example-----打印从第五行开始到第一个包含以test开始的行之间的所有行。

$ sed '/test/,/check/s/$/sed test/' example-----对于模板test和west之间的行，每行的末尾用字符串sed test替换。

多点编辑：e命令
$ sed -e '1,5d' -e 's/test/check/' example-----(-e)选项允许在同一行里执行多条命令。如例子所示，第一条命令删除1至5行，第二条命令用check替换test。命令的执行顺序对结果有影响。如果两个命令都是替换命令，那么第一个替换命令将影响第二个替换命令的结果。

$ sed --expression='s/test/check/' --expression='/love/d' example-----一个比-e更好的命令是--expression。它能给sed表达式赋值。

从文件读入：r命令
$ sed '/test/r file' example-----file里的内容被读进来，显示在与test匹配的行后面，如果匹配多行，则file的内容将显示在所有匹配行的下面。

写入文件：w命令
$ sed -n '/test/w file' example-----在example中所有包含test的行都被写入file里。

追加命令：a命令
$ sed '/^test/a\\--->this is a example' example<-----'this is a example'被追加到以test开头的行后面，sed要求命令a后面有一个反斜杠。

插入：i命令
$ sed '/test/i\\

new line

-------------------------' example

如果test被匹配，则把反斜杠后面的文本插入到匹配行的前面。

下一个：n命令
$ sed '/test/{ n; s/aa/bb/; }' example-----如果test被匹配，则移动到匹配行的下一行，替换这一行的aa，变为bb，并打印该行，然后继续。

变形：y命令
$ sed '1,10y/abcde/ABCDE/' example-----把1--10行内所有abcde转变为大写，注意，正则表达式元字符不能使用这个命令。

退出：q命令
$ sed '10q' example-----打印完第10行后，退出sed。

保持和获取：h命令和G命令
$ sed -e '/test/h' -e '$G example-----在sed处理文件的时候，每一行都被保存在一个叫模式空间的临时缓冲区中，除非行被删除或者输出被取消，否则所有被处理的行都将打印在屏幕上。接着模式空间被清空，并存入新的一行等待处理。在这个例子里，匹配test的行被找到后，将存入模式空间，h命令将其复制并存入一个称为保持缓存区的特殊缓冲区内。第二条语句的意思是，当到达最后一行后，G命令取出保持缓冲区的行，然后把它放回模式空间中，且追加到现在已经存在于模式空间中的行的末尾。在这个例子中就是追加到最后一行。简单来说，任何包含test的行都被复制并追加到该文件的末尾。

保持和互换：h命令和x命令
$ sed -e '/test/h' -e '/check/x' example -----互换模式空间和保持缓冲区的内容。也就是把包含test与check的行互换。

7. 脚本

Sed脚本是一个sed的命令清单，启动Sed时以-f选项引导脚本文件名。Sed对于脚本中输入的命令非常挑剔，在命令的末尾不能有任何空白或文本，如果在一行中有多个命令，要用分号分隔。以#开头的行为注释行，且不能跨行。

8. 小技巧

在sed的命令行中引用shell变量时要使用双引号，而不是通常所用的单引号。下面是一个根据name变量的内容来删除named.conf文件中zone段的脚本：

name='zone\ "localhost"'
sed "/$name/,/};/d" named.conf
原BLOG地址 http://blog.chinaunix.net/u/15010/showart_243352.html



1,sed介绍
    sed可删除(delete)、改变(change)、添加(append)、插入(insert)、合、交换文件中的资料行,或读入其它档的资料到文>件中,也可替换(substuite)它们其中的字串、或转换(tranfer)其中的字母等等。例如将文件中的连续空白行删成一行、"local"字串替换成"remote"、"t"字母转换成"T"、将第10行资料与第11资料合等.
    总合上述所言,当sed由标准输入读入一行资料并放入pattern space时,sed依照sed script 的编辑指令逐一对pattern space内的资料执行编辑,之後,再由pattern space内的结果送到标准输出,接着再将下一行资料读入.如此重执行上述动作,直至读>完所有资料行为止.
    小结,记住:
            (1)sed总是以行对输入进行处理
            (2)sed处理的不是原文件而是原文件的拷贝

命令行概述:
    sed 编辑指令的格式如下 :
              [address1[,address2]]function[argument]
其中 , 位址参数 address1 、address2 为行数或 regular expression 字串 , 表示所执行编辑的资料行; 函数参数 function[argument] 为 sed 的内定函数 , 表示执行的编辑动作。

有那些函数(function)参数
   下页表中介绍所有 sed 的函数参数(参照[chapter 4])的功能。
函数参数 功能
: label  建立 script file 内指令互相参考的位置。
#  建立解
    { }  集合有相同位址参数的指令。
    !  不执行函数参数。
    =  印出资料行数( line number )。
    a/  添加使用者输入的资料。
    b label  将执行的指令跳至由 : 建立的参考位置。
    c/  以使用者输入的资料取代资料。
    d  删除资料。
    D  删除 pattern space 内第一个 newline 字母 / 前的资料。
    g  拷贝资料从 hold space。
    G  添加资料从 hold space 至 pattern space 。
    h  拷贝资料从 pattern space 至 hold space 。
    H  添加资料从 pattern space 至 hold space 。
    l  印出 l 资料中的 nonprinting character 用 ASCII 码。
    i/  插入添加使用者输入的资料行。
    n  读入下一笔资料。
    N  添加下一笔资料到 pattern space。
    p  印出资料。
    P  印出 pattern space 内第一个 newline 字母 / 前的资料。
    q  跳出 sed 编辑。
    r  读入它档内容。
    s  替换字串。
    t label  先执行一替换的编辑指令 , 如果替换成牛p>则将编辑指令跳至 : label 处执行。
    w  写资料到它档内。
    x  交换 hold space 与 pattern space 内容。
    y  转换(transform)字元。
虽然 , sed 只有上表所述几个拥有基本编辑功能的函数 , 但由指令中位址参数和指令与指令间的配合 , 也能使sed 完成大部份的编辑任务。

2,1 删除
(1) sed -e '1d' inputfile (删除第一行)
    那么删除第x行呢?删除第x1,x2,x3行呢？
    sed -e 'xd' inputfile
    sed -e 'x1d' -e 'x2d' -e 'x3d' inputfile
    当然也许还有更好的办法。

(2) sed -e '1,3d' file (删除第一到第三行)
    思考：删除第n行到第m行?也就是
    sed -e 'n,md' file
    删除第一行到最后一行
    sed -e '1,$d' file     #$ 最后一行和一行的最后

(3) sed -e '/#/d' file  (删除含有'#'号的行)
    思考：删除含有字母xx的行
    sed -e '/xx/d' file
    思考: 删除除含有字符串xx的所有行
    sed -e '/xx/!d' file

(4) sed -e '/word1/, /word2/d' file  (删除从含有单词word1到含有单词word2的行)
    sed -e '10,/word1/d' file
    删除文件中从第10行到含有word1的行
    sed -e '/word1/,10/d' file
    和上面的匹配相反，删除从含有word1的行到第10行

(5) sed -e '/t.*t/d' file     (删除含有两个t的行)
    思考：删除含有指定正在表达式匹配的行。


2.2 替换
Sed 可替换文件中的字串、资料行、甚至资料区。其中,表示替换字串的指令中的函数参数为s;表示替换资料行、或资料区>的指令中的函数参数为c。上述情况以下面三个例子说明。

*行的替换
(1) sed -e '1c/#!/bin/more' file (把第一行替换成#!/bin/more)
    思考: 把第n行替换成just do it
    sed -e 'nc/just do it' file

(2) sed -e '1,10c/I can do it' file  (把1到10行替换成一行:I can do it)
    思考: 换成两行(I can do it! Let's start)
    sed -e '1,10c/I can do it!/nLet'"/'"'s start' file

*字符的替换
(3) sed -e 's/word1/& word2/' file (将每一行的word1单词替换成s参数最多与两个位置参数相结合,函数参数s中有两个特殊的符号:
    & : 代表pattern
    /n : 代表 pattern 中被第 n 个 /( 、/)(参照[附录 A]) 所括起来的字串。例如

    sed -e 's/w1/& w2/' file  # w1的地方输出 w1 w2
    sed -e  's//(test/) /(my/) /(car/)/[/2 /3 /1]/' file   #结果: [my car test]

*flag 参数举例
    sed -e 's/w1/& w2/g' file
    g : 代表替换所有匹配项目;这里,文件中所有字符串w1都会被替换成字串w1 w2
    sed -e 's/w1/& w2/10' file
    m(10) : 替换行内第m个符合的字串; 记住，是行内的第m个匹配的字串
    sed -e 's/w1/& w2/p' file
    p : 替换第一个和w1匹配的字符串为w1 w2，并输出到标准输出.
    sed -e 's/w1/& w2/w w2file' file
    w filename : 该参数会将替换过的内容写入到文件w2file并输出替换后的整个文件。注意w2file里写的只是替换过的行。    sed 'e 's/w1/& w2/' file
    这里的flag 为空, 这样就只是将第一个w1匹配的字符串替换成w1 w2而后面的不进行替换。

*位置参数应用举例
    sed -e '/machine/s/phi/beta/g' file
    将文件中含"machine"字串的资料行中的"phi"字串,替换成为"beta"字串
    sed -e '1,10 s/w1/& w2/g' file
    把1到10内的w1字符串替换成w1 w2字符串。
    sed -e '1,/else/ s/w1/& w2/g' file
    把1到字符串else内的w1字符串替换成w1 w2字符串。

其它位置参数的应用与前面的相同。


2.3 内容的插入
i
    基本格式:
    [address] i/ 插入内容 filename
 word2)
说明:
函数参数 s 表示替换(substitute)文件内字串。其指令格式如下 :
[address1[ ,address2]] s/pattern/replacemen/[flag]

    sed -e '/#/i/words' file      #在#字符的前面插入一行words

说明：
    这里的函数参数是i，它只能有一个地址参数。
    sed -e '1/i/words' file
    在第一行前加一行words
    cat "word" | sed -e '/$/.doc/g'   #输出word.doc
    在word后面加上后缀名，从而输出word.doc
    i 参数正好与a参数相反，它是插入到所给内容的前面.

a
    a参数的使用格式如下：
    [address] a/ <插入内容> filename

    sed -e '/unix/a/ haha' test.txt   #在含有unix的行后添加"haha"
    #输出结果为:
        unix
        haha

    另外: sed -e '1 a/ hh' test.txt  #在第一行后添加hh字符.

2.4 文本的打印: p
    基本格式：
    [address1,[address2]] p

    (1) sed -e '/then/ p' filename  #打印所有行并重复打印含有then 的行
    (2) sed -n '/then/ p' filename  #只打印含有then的行
    (3) sed -e '1,3 p' filename     # 打印所有行并重复1-3行
    (4) sed -n '1,3 p' filename     # 打印1-3行
    (5) sed -n '/if/,/fi/ p' filename #打印字符if和fi之间的内容

    p函数为sed的打印函数，在这里要注意-e 和-n 参数的区别。一般使用-n参数。


2.5 字元的替换: y
    例如：
    (1)sed -e 'y/abc../xyz../' filename
    把文件中的a字母替换成x, b替换成y, c替换成z。
    (2) sed  -e 'y/abc/ABC' filename
    把小写的abc转换成大写的ABC

2.6 反相执行命令 : !
    基本格式:
    [address1[ , address2]] ! 函数参数

    sed -e '/1996/!d' filename
    删除除了含有1996的所有行。


2.7 改变文件中的资料: c
    基本格式：
    [address1[ ,address2]]c/ filename
    函数参数 c 紧接着 "/" 字元用来表示此行结束 , 使用者所输入的资料必须从下一行输入。如果资料超过一行 , 则须在>每行的结尾加入"/"

    sed -e '/zhengxh/c hhhh' filename
    表示把含有字符串zhengxh的行，该成hhhh。

2.8 读入下一行资料: n
    基本格式：
    [address1[ ,address2]] n

    sed -n -e '/echo/n' -e 'p' temp
    表示输出文件，但如果一行含有字符串echo，则输出包含该字符串的下一行。
    sed -n -e 'n' -e 'p' filename
    输出文中的偶数行

3, 命令的复用
    一次执行多个命令的方式有三种：
    (1) sed 's/w1/& w2/g; 1/i/words' filename   (使用;号把命令隔开，注意前面不加-e参数)
    (2) sed -e 'cmd1' -e 'cmd2'  filename     (使用多个-e参数)
