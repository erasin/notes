# BASH 的基本语法

最简单的例子 —— Hello World!

关于输入、输出和错误输出

BASH 中对变量的规定（与 C 语言的异同）

BASH 中的基本流程控制语法

函数的使用

## 最简单的例子 —— Hello World!

几乎所有的讲解编程的书给读者的第一个例子都是 Hello World 程序，那么我们今天也就从这个例子出发，来逐步了解 BASH。

用 vi 编辑器编辑一个 hello 文件如下：

    #!/bin/bash
    # This is a very simple example
    echo Hello World

这样最简单的一个 BASH 程序就编写完了。这里有几个问题需要说明一下：

1. 第一行的 #! 是什么意思
2. 第一行的 /bin/bash 又是什么意思
3. 第二行是注释吗
4. echo 语句
5. 如何执行该程序

**`#!`** 是说明 hello 这个文件的类型的，有点类似于 Windows 系统下用不同文件后缀来表示不同文件类型的意思（但不相同）。  

Linux 系统根据 "#!" 及该字串后面的信息确定该文件的类型，关于这一问题同学们回去以后可以通过 "man magic"命令 及 /usr/share/magic 文件来了解这方面的更多内容。

在 BASH 中 第一行的 `#!` 及后面的 `/bin/bash` 就表明该文件是一个 BASH 程序，需要由`/bin`目录下的 bash 程序来解释执行。BASH 这个程序一般是存放在`/bin`目录下，如果你的 Linux 系统比较特别，bash 也有可能被存放在`/sbin`,`/usr/local/bin` 、`/usr/bin` 、`/usr/sbin` 或 `/usr/local/sbin` 这样的目录下；如果还找不到，你可以用 "locate bash" "find / -name bash 2> /dev/null" 或 "whereis bash" 这三个命令找出 bash 所在的位置；如果仍然找不到，那你可能需要自己动手安装一个 BASH 软件包了。


第二行的 `# This is a ...` 就是 BASH 程序的注释，在 BASH 程序中从“#”号（注意：后面紧接着是`!`号的除外）开始到行尾的多有部分均被看作是程序的注释。的三行的 echo 语句的功能是把 echo 后面的字符串输出到标准输出中去。由于 echo 后跟的是 "Hello World" 这个字符串，因此 "Hello World"这个字串就被显示在控制台终端的屏幕上了。需要注意的是 BASH 中的绝大多数语句结尾处都没有分号。

如何执行该程序呢？有两种方法：一种是显式制定 BASH 去执行：

    $ bash hello #或
    $ sh hello   #（这里 sh 是指向 bash 的一个链接，“lrwxrwxrwx 1 root root 4 Aug 20 05:41 /bin/sh -> bash”）

或者可以先将 hello 文件改为可以执行的文件，然后直接运行它，此时由于 hello 文件第一行的 "#! /bin/bash" 的作用，系统会自动用/bin/bash 程序去解释执行 hello 文件的：

    $ chmod u+x hello
    $ ./hello

此处没有直接 “$ hello”是因为当前目录不是当前用户可执行文件的默认目录，而将当前目录“.”设为默认目录是一个不安全的设置。

需要注意的是，BASH 程序被执行后，实际上 Linux 系统是另外开设了一个进程来运行的。

## 关于输入、输出和错误输出

在字符终端环境中，标准输入/标准输出的概念很好理解。输入即指对一个应用程序或命令的输入，无论是从键盘输入还是从别的文件输入；输出即指应用程序或命令产生的一些信息；与 Windows 系统下不同的是，Linux 系统下还有一个标准错误输出的概念，这个概念主要是为程序调试和系统维护目的而设置的，错误输出于标准输出分开可以让一些高级的错误信息不干扰正常的输出信息，从而方便一般用户的使用。

在 Linux 系统中：标准输入(stdin)默认为键盘输入；标准输出(stdout)默认为屏幕输出；标准错误输出(stderr)默认也是输出到屏幕（上面的 std 表示 standard）。在 BASH 中使用这些概念时一般将标准输出表示为 1，将标准错误输出表示为 2。下面我们举例来说明如何使用他们，特别是标准输出和标准错误输出。

输入、输出及标准错误输出主要用于 I/O 的重定向，就是说需要改变他们的默认设置。先看这个例子：

    $ ls > ls_result
    $ ls -l >> ls_result

上面这两个命令分别将 ls 命令的结果输出重定向到 ls_result 文件中和追加到 ls_result 文件中，而不是输出到屏幕上。">"就是输出（标准输出和标准错误输出）重定向的代表符号，连续两个 ">" 符号，即 ">>" 则表示不清除原来的而追加输出。下面再来看一个稍微复杂的例子：

    $ find /home -name lost* 2> err_result

这个命令在 `>` 符号之前多了一个 `2`，`2>` 表示将标准错误输出重定向。由于 /home 目录下有些目录由于权限限制不能访问，因此会产生一些标准错误输出被存放在`err_result`文件中。大家可以设想一下` find /home -name lost* 2>>err_result `命令会产生什么结果？

如果直接执行`find /home -name lost* > all_result` ，其结果是只有标准输出被存入`all_result`文件中，要想让标准错误输出和标准输入一样都被存入到文件中，那该怎么办呢？看下面这个例子：

    $ find /home -name lost* > all_result 2>& 1

上面这个例子中将首先将标准错误输出也重定向到标准输出中，再将标准输出重定向到 `all_result` 这个文件中。这样我们就可以将所有的输出都存储到文件中了。为实现上述功能，还有一种简便的写法如下：

    $ find /home -name lost* >& all_result

如果那些出错信息并不重要，下面这个命令可以让你避开众多无用出错信息的干扰：

    $ find /home -name lost* 2> /dev/null

同学们回去后还可以再试验一下如下几种重定向方式，看看会出什么结果，为什么？

    $ find /home -name lost* > all_result 1>& 2
    $ find /home -name lost* 2> all_result 1>& 2
    $ find /home -name lost* 2>& 1 > all_result

另外一个非常有用的重定向操作符是 "-"，请看下面这个例子：

    $ (cd /source/directory && tar cf - . ) | (cd /dest/directory && tar xvfp -)

该命令表示把 /source/directory 目录下的所有文件通过压缩和解压，快速的全部移动到 /dest/directory 目录下去，这个命令在 /source/directory 和 /dest/directory 不处在同一个文件系统下时将显示出特别的优势。

下面还几种不常见的用法：

* **n<&-** 表示将 n 号输入关闭
* **<&-** 表示关闭标准输入（键盘）
* **n>&**- 表示将 n 号输出关闭
* **>&-** 表示将标准输出关闭

 
## BASH 中对变量的规定（与 C 语言的异同）

好了下面我们进入正题，先看看 BASH 中的变量是如何定义和使用的。对于熟悉 C 语言的程序员，我们将解释 BASH 中的定义和用法与 C 语言中有何不同。

### BASH 中的变量介绍

我们先来从整体上把握一下 BASH 中变量的用法，然后再去分析 BASH 中变量使用与 C 语言中的不同。BASH 中的变量都是不能含有保留字，不能含有 "-" 等保留字符，也不能含有空格。

#### 简单变量

在 BASH 中变量定义是不需要的，没有 "int i" 这样的定义过程。如果想用一个变量，只要他没有在前面被定义过，就直接可以用，当然你使用该变量的第一条语句应该是对他赋初值了，如果你不赋初值也没关系，只不过该变量是空（ 注意：是 NULL，不是 0 ）。不给变量赋初值虽然语法上不反对，但不是一个好的编程习惯。好了我们看看下面的例子：

首先用 vi 编辑下面这个文件 hello2：

    #!/bin/bash
    # give the initialize value to STR
    STR="Hello World"
    echo $STR

在上面这个程序中我们需要注意下面几点：

1. 变量赋值时，`=`左右两边都_不能有空格_；
2. BASH 中的语句结尾_不需要分号（";"）_；
3. 除了在变量赋值和在FOR循环语句头中，BASH 中的变量使用必须在变量前加"$"符号，同学们可以将上面程序中第三行改为 "echo STR" 再试试，看看会出什么结果。
4. 由于 BASH 程序是在一个新的进程中运行的，所以该程序中的变量定义和赋值不会改变其他进程或原始 Shell 中同名变量的值，也不会影响他们的运行。

更细致的文档甚至提到以但引号括起来的变量将不被 BASH 解释为变量，如 '$STR' ，而被看成为纯粹的字符串。而且更为**标准**的变量引用方式是 `${STR}` 这样的，$STR 自不过是对 ${STR} 的一种简化。在复杂情况下（即有可能产生歧义的地方）最好用带 {} 的表示方式。

BASH 中的变量既然不需要定义，也就没有类型一说，一个变量即可以被定义为一个字符串，也可以被再定义为整数。如果对该变量进行整数运算，他就被解释为整数；如果对他进行字符串操作，他就被看作为一个字符串。请看下面的例子：

    #!/bin/bash
    x=1999
    let "x = $x + 1"
    echo $x
    x="olympic'"$x
    echo $x

关于整数变量计算，有如下几种：` + - * / % `，他们的意思和字面意思相同。整数运算一般通过 `let` 和 `expr` 这两个指令来实现，如对变量 x 加 1 可以写作：`let "x = $x + 1"` 或者 x=`expr $x + 1`

在比较操作上，整数变量和字符串变量各不相同，详见下表：

|对应的操作 |整数操作|字符串操作|
|-----------|--------|----------|
|相同       |   -eq  | =  |
|不同       |   -ne  | != |
|大于       |   -gt  | >  |
|小于       |   -lt  | <  |
|大于或等于 |   -ge  |    |
|小于或等于 |   -l   |    |
|为空       |        | -z |
|不为空     |        | -n |


比如：

比较字符串 a 和 b 是否相等就写作：if [ $a = $b ]  
判断字符串 a 是否为空就写作： if [ -z $a ]  
判断整数变量 a 是否大于 b 就写作：if [ $a -gt $b ]  

更细致的文档推荐在字符串比较时尽量不要使用 -n ，而用 ! -z 来代替。（其中符号 "!" 表示求反操作）

BASH 中的变量除了用于对 整数 和 字符串 进行操作以外，另一个作用是作为文件变量。BASH 是 Linux 操作系统的 Shell，因此系统的文件必然是 BASH 需要操作的重要对象，如 `if [ -x /root ]` 可以用于判断 /root 目录是否可以被当前用户进入。下表列出了 BASH 中用于判断文件属性的操作符：

|运算符         |含义（ 满足下面要求时返回 TRUE ）|
|---------------|---------------------------------|
|-e file        |文件 file 已经存在               |
|-f file        |文件 file 是普通文件             |
|-s file        |文件 file 大小不为零             |
|-d file        |文件 file 是一个目录             |
|-r file        |文件 file 对当前用户可以读取     |
|-w file        |文件 file 对当前用户可以写入     |
|-x file        |文件 file 对当前用户可以执行     |
|-g file        |文件 file 的 GID 标志被设置      |
|-u file        |文件 file 的 UID 标志被设置      |
|-O file        |文件 file 是属于当前用户的       |
|-G file        |文件 file 的组 ID 和当前用户相同 |
|file1 -nt file2|文件 file1 比 file2 更新         |
|file1 -ot file2|文件 file1 比 file2 更老         |

注意：上表中的 file 及 file1、file2 都是指某个文件或目录的路径。

###  关于局部变量

在 BASH 程序中如果一个变量被使用了，那么直到该程序的结尾，该变量都一直有效。为了使得某个变量存在于一个局部程序块中，就引入了局部变量的概念。BASH 中，在变量首次被赋初值时加上 local 关键字就可以声明一个局部变量，如下面这个例子：

    #!/bin/bash
    HELLO=Hello
    function hello {
    　　local HELLO=World
    　　echo $HELLO
    }
    echo $HELLO
    hello
    echo $HELLO

该程序的执行结果是：

    Hello
    World
    Hello

这个执行结果表明全局变量 $HELLO 的值在执行函数 hello 时并没有被改变。也就是说局部变量 $HELLO 的影响只存在于函数那个程序块中。

### BASH 中的变量与 C 语言中变量的区别

这里我们为原来不熟悉 BASH 编程，但是非常熟悉 C 语言的程序员总结一下在 BASH 环境中使用变量需要注意的问题。

1. BASH 中的变量在引用时都需要在变量前加上 "$" 符号（ 第一次赋值及在For循环的头部不用加 "$"符号 ）；
2. BASH 中没有浮点运算，因此也就没有浮点类型的变量可用；
3. BASH 中的整形变量的比较符号与 C 语言中完全不同，而且整形变量的算术运算也需要经过 let 或 expr 语句来处理；

## BASH 中的基本流程控制语法

BASH 中几乎含有 C 语言中常用的所有控制结构，如条件分支、循环等，下面逐一介绍。

### if...then...else

if 语句用于判断和分支，其语法规则和 C 语言的 if 非常相似。其几种基本结构为：

    if [ expression ]
    then
    　　statments
    fi

或者

    if [ expression ]
    then
    　　statments
    else
    　　statments
    fi

或者

    if [ expression ]
    then
    　　statments
    else if [ expression ]
    　　then
    　　　　statments
    　　else
    　　　　statments
    fi

或者

    if [ expression ]
    then
    　　statments
    elif [ expression ]
    　　then
    　　　　statments
    　　else
    　　　　statments
    fi

值得说明的是如果你将 if 和 then 简洁的写在一行里面，就必须在 then 前面加上分号，如：if [ expression ]; then ... 。下面这个例子说明了如何使用 if 条件判断语句：

    #!/bin/bash

    if [ $1 -gt 90 ]
    then
    　　echo "Good, $1"
    elif [ $1 -gt 70 ]
    　　then
    　　　　echo "OK, $1"
    　　else
    　　　　echo "Bad, $1"
    fi

    exit 0

上面例子中的 $1 是指命令行的第一个参数，这个会在后面的“BASH 中的特殊保留字”中讲解。
### for

for 循环结构与 C 语言中有所不同，在 BASH 中 for 循环的基本结构是：

    for $var in [list]
    do
    　　statments
    done

其中 $var 是循环控制变量，[list] 是 $var 需要遍历的一个集合，do/done 对包含了循环体，相当于 C 语言中的一对大括号。另外如果do 和 for 被写在同一行，必须在 do 前面加上 ";"。如： for $var in [list]; do 。下面是一个运用 for 进行循环的例子：

    #!/bin/bash

    for day in Sun Mon Tue Wed Thu Fri Sat
    do
    　　echo $day
    done

    # 如果列表被包含在一对双引号中，则被认为是一个元素
    for day in "Sun Mon Tue Wed Thu Fri Sat"
    do
    　　echo $day
    done

    exit 0

注意上面的例子中，在 for 所在那行的变量 day 是没有加 "$" 符号的，而在循环体内，echo 所在行变量 $day 是必须加上 "$" 符号的。另外如果写成 for day 而没有后面的 in [list] 部分，则 day 将取遍命令行的所有参数。如这个程序：

    #!/bin/bash

    for param
    do
    　　echo $param
    done

    exit 0

上面这个程序将列出所有命令行参数。for 循环结构的循环体被包含在 do/done 对中，这也是后面的 while、until 循环所具有的特点。

### while

while 循环的基本结构是：

    while [ condition ]
    do
    　　statments
    done

这个结构请大家自己编写一个例子来验证。
### until

until 循环的基本结构是：

    until [ condition is TRUE ]
    do
    　　statments
    done

这个结构也请大家自己编写一个例子来验证。
### case

BASH 中的 case 结构与 C 语言中的 switch 语句的功能比较类似，可以用于进行多项分支控制。其基本结构是：

    case "$var" in
    　condition1 )
    　　statments1;;
    　condition2 )
    　　statments2;;
    　...
    　* )
    　　default statments;;
    esac

下面这个程序是运用 case 结构进行分支执行的例子：

    #!/bin/bash

    echo "Hit a key, then hit return."
    read Keypress

    case "$Keypress" in
    　[a-z] ) echo "Lowercase letter";;
    　[A-Z] ) echo "Uppercase letter";;
    　[0-9] ) echo "Digit";;
    　* ) echo "Punctuation, whitespace, or other";;
    esac

    exit 0

上面例子中的第四行 "read Keypress" 一句中的 read 语句表示从键盘上读取输入。这个命令将在本讲义的 BASH 的其他高级问题中讲解。
### break/continue

熟悉 C 语言编程的都很熟悉 break 语句和 continue 语句。BASH 中同样有这两条语句，而且作用和用法也和 C 语言中相同，break 语句可以让程序流程从当前循环体中完全跳出，而 continue 语句可以跳过当次循环的剩余部分并直接进入下一次循环。

 
## 函数的使用

BASH 是一个相对简单的脚本语言，不过为了方便结构化的设计，BASH 中也提供了函数定义的功能。BASH 中的函数定义很简单，只要向下面这样写就可以了：

    function my_funcname {
    　code block
    }

或者

    my_funcname() {
    　code block
    }

上面的第二种写法更接近于 C 语言中的写法。BASH 中要求函数的定义必须在函数使用之前，这是和 C 语言用头文件说明函数方法的不同。

更进一步的问题是如何给函数传递参数和获得返回值。BASH 中函数参数的定义并不需要在函数定义处就制定，而只需要在函数被调用时用 BASH 的保留变量 $1 $2 ... 来引用就可以了；BASH 的返回值可以用 return 语句来指定返回一个特定的整数，如果没有 return 语句显式的返回一个返回值，则返回值就是该函数最后一条语句执行的结果（一般为 0，如果执行失败返回错误码）。函数的返回值在调用该函数的程序体中通过 $? 保留字来获得。下面我们就来看一个用函数来计算整数平方的例子：

    #!/bin/bash

    square() {
    　let "res = $1 * $1"
    　return $res
    }

    square $1
    result=$?
    echo $result

    exit 0

[source] [aka]
[aka]:http://www.aka.org.cn/Lectures/002/Lecture-2.1.2/bash-syntax.html

<script src="/static/syntax/scripts/shCore.js" type="text/javascript" charset="utf-8"></script>
<script src="/static/syntax/scripts/shBrushBash.js" type="text/javascript" charset="utf-8"></script>
<link rel="stylesheet" href="/static/syntax/styles/shCore.css" type="text/css" media="screen" title="no title" charset="utf-8">
<link rel="stylesheet" href="/static/syntax/styles/shCoreDefault.css" type="text/css" media="screen" title="no title" charset="utf-8">
<script type="text/javascript">
    $('pre').addClass('brush: bash;');
    SyntaxHighlighter.all();
</script>
