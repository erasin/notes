# linux shell 常用技巧


## 一.    特殊文件: /dev/null和/dev/tty

Linux系统提供了两个对Shell编程非常有用的特殊文件，/dev/null和/dev/tty。其中/dev/null将会丢掉所有写入它的数据，换句换说，当程序将数据写入到此文件时，会认为它已经成功完成写入数据的操作，但实际上什么事都没有做。如果你需要的是命令的退出状态，而非它的输出，此功能会非常有用，见如下Shell代码：

    /> vi test_dev_null.sh
    
    #!/bin/bash
    if grep hello TestFile > /dev/null
    then
        echo "Found"
    else
        echo "NOT Found"
    fi

在vi中保存并退出后执行以下命令：

    /> chmod +x test_dev_null.sh  #使该文件成为可执行文件
    /> cat > TestFile
    hello my friend
    CTRL + D                             #退出命令行文件编辑状态
    /> ./test_dev_null.sh
    Found                                 #这里并没有输出grep命令的执行结果。

将以上Shell脚本做如下修改：

    /> vi test_dev_null.sh
    
    #!/bin/bash
    if grep hello TestFile
    then
        echo "Found"
    else
        echo "NOT Found"
    fi

在vi中保存退出后，再次执行该脚本：

    /> ./test_dev_null.sh
    hello my friend                      #grep命令的执行结果被输出了。
    Found
    
下面我们再来看/dev/tty的用途。当程序打开此文件是，Linux会自动将它重定向到一个终端窗口，因此该文件对于读取人工输入时特别有用。见如下Shell代码：

    /> vi test_dev_tty.sh
    
    #!/bin/bash
    printf "Enter new password: "    #提示输入
    stty -echo                               #关闭自动打印输入字符的功能
    read password < /dev/tty         #读取密码
    printf "\nEnter again: "             #换行后提示再输入一次
    read password2 < /dev/tty       #再读取一次以确认
    printf "\n"                               #换行
    stty echo                                #记着打开自动打印输入字符的功能
    echo "Password = " $password #输出读入变量
    echo "Password2 = " $password2
    echo "All Done"

在vi中保存并退出后执行以下命令：

    /> chmod +x test_dev_tty.sh #使该文件成为可执行文件
    /> ./test_dev_tty
    Enter new password:             #这里密码的输入被读入到脚本中的password变量
    Enter again:                          #这里密码的输入被读入到脚本中的password2变量
    Password = hello
    Password2 = hello
    All Done

## 二.    简单的命令跟踪:

Linux Shell提供了两种方式来跟踪Shell脚本中的命令，以帮助我们准确的定位程序中存在的问题。下面的代码为第一种方式，该方式会将Shell脚本中所有被执行的命令打印到终端，并在命令前加"+"：加号的后面还跟着一个空格。

    /> cat > trace_all_command.sh
    who | wc -l                          #这两条Shell命令将输出当前Linux服务器登录的用户数量
    CTRL + D                            #退出命令行文件编辑状态
    /> chmod +x trace_all_command.sh
    /> sh -x ./trace_all_command.sh #Shell执行器的-x选项将打开脚本的执行跟踪功能。
    + wc -l                               #被跟踪的两条Shell命令
    + who
    2                                       #实际输出结果。

Linux Shell提供的另一种方式可以只打印部分被执行的Shell命令，该方法在调试较为复杂的脚本时，显得尤为有用。

    /> cat > trace_patial_command.sh
    #! /bin/bash
    set -x                                #从该命令之后打开跟踪功能
    echo 1st echo                     #将被打印输出的Shell命令
    set +x                               #该Shell命令也将被打印输出，然而在该命令被执行之后，所有的命令将不再打印输出
    echo 2nd echo                    #该Shell命令将不再被打印输出。
    CTRL + D                           #退出命令行文件编辑状态
    /> chmod +x trace_patial_command.sh
    /> ./trace_patial_command.sh
    + echo 1st echo
    1st echo
    + set +x
    2nd echo
   
## 三.    正则表达式基本语法描述:

Linux Shell环境下提供了两种正则表达式规则，一个是基本正则表达式(BRE)，另一个是扩展正则表达式(ERE)。

下面是这两种表达式的语法列表，需要注意的是，如果没有明确指出的Meta字符，其将可同时用于BRE和ERE，否则将尽适用于指定的模式。

正则元字符	模式含义	用例
\	通常用于关闭其后续字符的特殊意义，恢复其原意。	\(...\)，这里的括号仅仅表示括号。
.	匹配任何单个字符。	a.b，将匹配abb、acb等
*	匹配它之前的0-n个的单个字符。	a*b，将匹配ab、aab、aaab等。
^	匹配紧接着的正则表达式，在行的起始处。	^ab，将匹配abc、abd等，但是不匹配cab。
$	匹配紧接着的正则表达式，在行的结尾处。	ab$，将匹配ab、cab等，但是不匹配abc。
[...]	方括号表达式，匹配其内部任何字符。其中-表示连续字符的范围，^符号置于方括号里第一个字符则有反向的含义，即匹配不在列表内(方括号)的任何字符。如果想让]和-表示其原意，需要将其放置在方括号的首字符位置，如[]ab]或[-ab]，如这两个字符同时存在，则将]放置在首字符位置，-放置在最尾部，如[]ab-]。	[a-bA-Z0-9!]表示所有的大小写字母，数字和感叹号。[^abc]表示a、b、c之外的所有字符。[Tt]om，可以匹配Tom和tom。
\{n,m\}	区间表达式，匹配在它前面的单个字符重复出现的次数区间，\{n\}表示重复n次；\{n,\}表示至少重复n次；\{n,m\}表示重复n到m次。	ab\{2\}表示abb；ab\{2,\}表示abb、abbb等。ab\{2,4\}表示abb、abbb和abbbb。
\(...\)	将圆括号之间的模式存储在特殊“保留空间”。最多可以将9个独立的子模式存储在单个模式中。匹配于子模式的文本，可以通过转义序列\1到\9，被重复使用在相同模式里。	\(ab\).*\1表示ab组合出现两次，两次之间可存在任何数目的任何字符，如abcdab、abab等。
{n,m}(ERE)	其功能等同于上面的\{n,m\}，只是不再写\转义符了。	ab+匹配ab、abbb等，但是不匹配a。
+(ERE)	和前面的星号相比，+匹配的是前面正则表达式的1-n个实例。	 
?(ERE)	匹配前面正则表达式的0个或1个。	ab?仅匹配a或ab。
|(ERE)	匹配于|符号前后的正则表达式。	(ab|cd)匹配ab或cd。
[:alpha:]	匹配字母字符。	[[:alpha:]!]ab$匹配cab、dab和!ab。
[:alnum:]	匹配字母和数字字符。	[[:alnum:]]ab$匹配1ab、aab。
[:blank:]	匹配空格(space)和Tab字符。	[[:alnum:]]ab$匹配1ab、aab。
[:cntrl:]	匹配控制字符。	 
[:digit:]	匹配数字字符。	 
[:graph:]	匹配非空格字符。	 
[:lower:]	匹配小写字母字符。	 
[:upper:]	匹配大写字母字符。	 
[:punct:]	匹配标点字符。	 
[:space:]	匹配空白(whitespace)字符。	 
[:xdigit:]	匹配十六进制数字。	 
\w	匹配任何字母和数字组成的字符，等同于[[:alnum:]_]	 
\W	匹配任何非字母和数字组成的字符，等同于[^[:alnum:]_]	 
\<\>	匹配单词的起始和结尾。	\<read匹配readme，me\>匹配readme。

下面的列表给出了Linux Shell中常用的工具或命令分别支持的正则表达式的类型。

 	grep	sed	vi	egrep	awk
BRE	*	*	*	 	 
ERE	 	 	 	*	*

## 四.使用cut命令选定字段:

 cut命令是用来剪下文本文件里的数据，文本文件可以是字段类型或是字符类型。下面给出应用实例：

    /> cat /etc/passwd
    root:x:0:0:root:/root:/bin/bash
    bin:x:1:1:bin:/bin:/sbin/nologin
    daemon:x:2:2:daemon:/sbin:/sbin/nologin
    adm:x:3:4:adm:/var/adm:/sbin/nologin
    ... ...
    /> cut -d : -f 1,5 /etc/passwd     #-d后面的冒号表示字段之间的分隔符，-f表示取分割后的哪些字段
    root:root                                 #这里取出的是第一个和第五个字段。
    bin:bin
    daemon:daemon
    adm:adm
    ... ...
    /> cut -d: -f 3- /etc/passwd       #从第三个字段开始显示，直到最后一个字段。
    0:0:root:/root:/bin/bash
    1:1:bin:/bin:/sbin/nologin
    2:2:daemon:/sbin:/sbin/nologin
    3:4:adm:/var/adm:/sbin/nologin
    4:7:lp:/var/spool/lpd:/sbin/nologin
    ... ...    

这里需要进一步说明的是，使用cut命令还可以剪切以字符数量为标量的部分字符，该功能通过-c选项实现，其不能与-d选项共存。

    /> cut -c 1-4 /etc/passwd          #取每行的前1-4个字符。
    /> cut -c-4 /etc/passwd            #取每行的前4个字符。 
    root
    bin:
    daem
    adm:
    ... ...
    /> cut -c4- /etc/passwd            #取每行的第4个到最后字符。
    t:x:0:0:root:/root:/bin/bash
    :x:1:1:bin:/bin:/sbin/nologin
    mon:x:2:2:daemon:/sbin:/sbin/nologin
    :x:3:4:adm:/var/adm:/sbin/nologin
    ... ...
    /> cut -c1,4 /etc/passwd           #取每行的第一个和第四个字符。
    rt
    b:
    dm
    a:
    ... ...
    /> cut -c1-4,5 /etc/passwd        #取每行的1-4和第5个字符。
    root:
    bin:x
    daemo
    adm:x

## 五.    计算行数、字数以及字符数:

Linux提供了一个简单的工具wc用于完成该功能，见如下用例：

    /> echo This is a test of the emergency broadcast system | wc
    1    9    49                              #1行，9个单词，49个字符
    /> echo Testing one two three | wc -c
    22                                         #22个字符
    /> echo Testing one two three | wc -l
    1                                           #1行
    /> echo Testing one two three | wc -w
    4                                           #4个单词
    /> wc /etc/passwd /etc/group    #计算两个文件里的数据。
    39   71  1933  /etc/passwd
    62   62  906    /etc/group
    101 133 2839  总用量

## 六.    提取开头或结尾数行:

有时，你会需要从文本文件里把几行字，多半是靠近开头或结尾的几行提取出来。如查看工作日志等操作。Linux Shell提供head和tail两个命令来完成此项工作。见如下用例：

    /> head -n 5 /etc/passwd           #显示输入文件的前五行。
    root:x:0:0:root:/root:/bin/bash
    bin:x:1:1:bin:/bin:/sbin/nologin
    daemon:x:2:2:daemon:/sbin:/sbin/nologin
    adm:x:3:4:adm:/var/adm:/sbin/nologin
    lp:x:4:7:lp:/var/spool/lpd:/sbin/nologin

    /> tail -n 5 /etc/passwd             #显示输入文件的最后五行。
    sshd:x:74:74:Privilege-separated SSH:/var/empty/sshd:/sbin/nologin
    mysql:x:27:27:MySQL Server:/var/lib/mysql:/bin/bash
    pulse:x:496:494:PulseAudio System Daemon:/var/run/pulse:/sbin/nologin
    gdm:x:42:42::/var/lib/gdm:/sbin/nologin
    stephen:x:500:500:stephen:/home/stephen:/bin/bash

如果使用者想查看不间断增长的日志(如服务程序输出的)，可以使用tail的-f选项，这样可以让tail命令不会自动退出，必须通过CTRL+C命令强制退出，因此该选项不适合用于Shell脚本中，见如下用例：

    /> tail -f -n 5 my_server_log
    ... ...
    ^C                                         #CTRL+C退出到命令行提示符状态。


## 七.　grep家族:
    
1、  grep退出状态：

    0: 表示成功；
    1: 表示在所提供的文件无法找到匹配的pattern；
    2: 表示参数中提供的文件不存在。

见如下示例：

    /> grep 'root' /etc/passwd
    root:x:0:0:root:/root:/bin/bash
    operator:x:11:0:operator:/root:/sbin/nologin
    /> echo $?
    0
    
    /> grep 'root1' /etc/passwd  #用户root1并不存在
    /> echo $?
    1
    
    /> grep 'root' /etc/passwd1  #这里的/etc/passwd1文件并不存在
    grep: /etc/passwd1: No such file or directory
    /> echo $?
    2
    
2、  grep中应用正则表达式的实例：

需要说明的是下面所涉及的正则表达式在上一篇中已经给出了详细的说明，因此在看下面例子的时候，可以与前一篇的正则说明部分结合着看。

    /> cat testfile
    northwest        NW      Charles Main           3.0     .98     3       34
    western           WE       Sharon Gray          5.3     .97     5       23
    southwest       SW       Lewis Dalsass         2.7     .8       2       18
    southern         SO       Suan Chin               5.1     .95     4       15
    southeast       SE        Patricia Hemenway    4.0     .7       4       17
    eastern           EA        TB Savage              4.4     .84     5       20
    northeast        NE        AM Main Jr.              5.1     .94     3       13
    north              NO       Margot Weber          4.5     .89     5       9
    central            CT        Ann Stephens          5.7     .94     5       13

    
    /> grep NW testfile     #打印出testfile中所有包含NW的行。
    northwest       NW      Charles Main        3.0     .98     3       34
    
    /> grep '^n' testfile   #打印出以n开头的行。
    northwest       NW      Charles Main        3.0     .98     3       34
    northeast        NE       AM Main Jr.          5.1     .94     3       13
    north              NO      Margot Weber      4.5     .89     5       9
    
    /> grep '4$' testfile   #打印出以4结尾的行。
    northwest       NW      Charles Main        3.0     .98     3       34
    
    /> grep '5\..' testfile #打印出第一个字符是5，后面跟着一个.字符，在后面是任意字符的行。
    western         WE      Sharon Gray         5.3     .97     5       23
    southern        SO      Suan Chin             5.1     .95     4       15
    northeast       NE      AM Main Jr.            5.1     .94     3       13
    central           CT      Ann Stephens        5.7     .94     5       13
    
    /> grep '\.5' testfile  #打印出所有包含.5的行。
    north           NO      Margot Weber        4.5     .89     5       9

    /> grep '^[we]' testfile #打印出所有以w或e开头的行。
    western         WE      Sharon Gray         5.3     .97     5       23
    eastern          EA      TB Savage            4.4     .84     5       20
    
    /> grep '[^0-9]' testfile #打印出所有不是以0-9开头的行。
    northwest       NW     Charles Main             3.0     .98      3       34
    western          WE      Sharon Gray             5.3     .97     5       23
    southwest       SW     Lewis Dalsass           2.7     .8       2       18
    southern         SO      Suan Chin                5.1     .95     4       15
    southeast        SE      Patricia Hemenway     4.0     .7      4       17
    eastern           EA      TB Savage                4.4     .84     5       20
    northeast        NE      AM Main Jr.                5.1     .94     3       13
    north              NO      Margot Weber           4.5     .89     5       9
    central            CT      Ann Stephens            5.7     .94     5       13

    /> grep '[A-Z][A-Z] [A-Z]' testfile #打印出所有包含前两个字符是大写字符，后面紧跟一个空格及一个大写字母的行。
    eastern          EA      TB Savage       4.4     .84     5       20
    northeast       NE      AM Main Jr.      5.1     .94     3       13

> 注：在执行以上命令时，如果不能得到预期的结果，即grep忽略了大小写，导致这一问题的原因很可能是当前环境的本地化的设置问题。对于以上命令，如果我将当前语言设置为en_US的时候，它会打印出所有的行，当我将其修改为中文环境时，就能得到我现在的输出了。

    /> export LANG=zh_CN  #设置当前的语言环境为中文。
    /> export LANG=en_US  #设置当前的语言环境为美国。
    /> export LANG=en_Br  #设置当前的语言环境为英国。
    
    /> grep '[a-z]\{9\}' testfile #打印所有包含每个字符串至少有9个连续小写字符的字符串的行。
    northwest        NW      Charles Main          3.0     .98     3       34
    southwest       SW      Lewis Dalsass         2.7     .8       2       18
    southeast        SE      Patricia Hemenway   4.0     .7       4       17
    northeast        NE      AM Main Jr.              5.1     .94     3       13
    
> 第一个字符是3，紧跟着一个句点，然后是任意一个数字，然后是任意个任意字符，然后又是一个3，然后是制表符，然后又是一个3，需要说明的是，下面正则中的\1表示\(3\)。

    /> grep '\(3\)\.[0-9].*\1    *\1' testfile 
    northwest       NW      Charles Main        3.0     .98     3       34
    
    /> grep '\<north' testfile    #打印所有以north开头的单词的行。
    northwest       NW      Charles Main          3.0     .98     3       34
    northeast        NE       AM Main Jr.            5.1     .94     3       13
    north              NO      Margot Weber        4.5     .89     5       9
    
    /> grep '\<north\>' testfile  #打印所有包含单词north的行。
    north           NO      Margot Weber        4.5     .89     5       9
    
    /> grep '^n\w*' testfile      #第一个字符是n，后面是任意字母或者数字。
    northwest       NW     Charles Main          3.0     .98     3       34
    northeast        NE      AM Main Jr.            5.1     .94     3       13
    north             NO      Margot Weber        4.5     .89     5       9
    
 3、  扩展grep(grep -E 或者 egrep)：

使用扩展grep的主要好处是增加了额外的正则表达式元字符集。下面我们还是继续使用实例来演示扩展grep。

    /> egrep 'NW|EA' testfile     #打印所有包含NW或EA的行。如果不是使用egrep，而是grep，将不会有结果查出。
    northwest       NW      Charles Main        3.0     .98     3       34
    eastern         EA      TB Savage           4.4     .84     5       20
    
    /> grep 'NW\|EA' testfile     #对于标准grep，如果在扩展元字符前面加\，grep会自动启用扩展选项-E。
    northwest       NW      Charles Main        3.0     .98     3       34
    eastern           EA       TB Savage           4.4     .84     5       20
    
    /> egrep '3+' testfile
    /> grep -E '3+' testfile
    /> grep '3\+' testfile        #这3条命令将会打印出相同的结果，即所有包含一个或多个3的行。
    northwest       NW      Charles Main         3.0     .98     3       34
    western          WE      Sharon Gray         5.3     .97     5       23
    northeast        NE       AM Main Jr.           5.1     .94     3       13
    central            CT       Ann Stephens       5.7     .94     5       13
    
    /> egrep '2\.?[0-9]' testfile 
    /> grep -E '2\.?[0-9]' testfile
    /> grep '2\.\?[0-9]' testfile #首先含有2字符，其后紧跟着0个或1个点，后面再是0和9之间的数字。
    western         WE       Sharon Gray          5.3     .97     5       23
    southwest      SW      Lewis Dalsass         2.7     .8      2       18
    eastern          EA       TB Savage             4.4     .84     5       20
    
    /> egrep '(no)+' testfile
    /> grep -E '(no)+' testfile
    /> grep '\(no\)\+' testfile   #3个命令返回相同结果，即打印一个或者多个连续的no的行。
    northwest       NW      Charles Main        3.0     .98     3       34
    northeast        NE       AM Main Jr.          5.1     .94     3       13
    north              NO      Margot Weber      4.5     .89     5       9
    
    /> grep -E '\w+\W+[ABC]' testfile #首先是一个或者多个字母，紧跟着一个或者多个非字母数字，最后一个是ABC中的一个。
    northwest       NW     Charles Main       3.0     .98     3       34
    southern        SO      Suan Chin           5.1     .95     4       15
    northeast       NE      AM Main Jr.          5.1     .94     3       13
    central           CT      Ann Stephens      5.7     .94     5       13
    
    /> egrep '[Ss](h|u)' testfile
    /> grep -E '[Ss](h|u)' testfile
    /> grep '[Ss]\(h\|u\)' testfile   #3个命令返回相同结果，即以S或s开头，紧跟着h或者u的行。
    western         WE      Sharon Gray       5.3     .97     5       23
    southern        SO      Suan Chin          5.1     .95     4       15
    
    /> egrep 'w(es)t.*\1' testfile    #west开头，其中es为\1的值，后面紧跟着任意数量的任意字符，最后还有一个es出现在该行。
    northwest       NW      Charles Main        3.0     .98     3       34
   

4、  grep选项：

 这里先列出grep常用的命令行选项：

选项	说明
-c	只显示有多少行匹配，而不具体显示匹配的行。
-h	不显示文件名。
-i	在字符串比较的时候忽略大小写。
-l	只显示包含匹配模板的行的文件名清单。
-L	只显示不包含匹配模板的行的文件名清单。
-n	在每一行前面打印改行在文件中的行数。
-v	反向检索，只显示不匹配的行。
-w	只显示完整单词的匹配。
-x	只显示完整行的匹配。
-r/-R	如果文件参数是目录，该选项将递归搜索该目录下的所有子目录和文件。

    /> grep -n '^south' testfile  #-n选项在每一个匹配行的前面打印行号。
    3:southwest     SW      Lewis Dalsass         2.7     .8      2       18
    4:southern       SO      Suan Chin               5.1     .95     4       15
    5:southeast      SE      Patricia Hemenway    4.0     .7      4       17

    /> grep -i 'pat' testfile     #-i选项关闭了大小写敏感。
    southeast       SE      Patricia Hemenway       4.0     .7      4       17

    /> grep -v 'Suan Chin' testfile #打印所有不包含Suan Chin的行。
    northwest       NW      Charles Main          3.0     .98     3       34
    western          WE      Sharon Gray           5.3     .97    5       23
    southwest       SW      Lewis Dalsass        2.7     .8      2       18
    southeast        SE      Patricia Hemenway   4.0     .7      4       17
    eastern           EA      TB Savage              4.4     .84     5       20
    northeast        NE      AM Main Jr.             5.1     .94     3       13
    north              NO      Margot Weber        4.5     .89     5       9
    central            CT      Ann Stephens         5.7     .94     5       13

    /> grep -l 'ss' testfile  #-l使得grep只打印匹配的文件名，而不打印匹配的行。
    testfile

    /> grep -c 'west' testfile #-c使得grep只打印有多少匹配模板的行。
    3

    /> grep -w 'north' testfile #-w只打印整个单词匹配的行。
    north           NO      Margot Weber    4.5     .89     5       9

    /> grep -C 2 Patricia testfile #打印匹配行及其上下各两行。
    southwest      SW     Lewis Dalsass         2.7     .8       2       18
    southern        SO      Suan Chin              5.1     .95     4       15
    southeast       SE      Patricia Hemenway   4.0     .7      4       17
    eastern          EA      TB Savage              4.4     .84     5       20
    northeast       NE      AM Main Jr.             5.1     .94     3       13

    /> grep -B 2 Patricia testfile #打印匹配行及其前两行。
    southwest      SW      Lewis Dalsass         2.7     .8      2       18
    southern        SO      Suan Chin               5.1     .95    4       15
    southeast       SE      Patricia Hemenway   4.0     .7      4       17

    /> grep -A 2 Patricia testfile #打印匹配行及其后两行。
    southeast       SE      Patricia Hemenway   4.0     .7      4       17
    eastern           EA      TB Savage              4.4     .84     5       20
    northeast       NE       AM Main Jr.             5.1     .94     3       13

## 八.　流编辑器sed:

sed一次处理一行文件并把输出送往屏幕。sed把当前处理的行存储在临时缓冲区中，称为模式空间(pattern space)。一旦sed完成对模式空间中的行的处理，模式空间中的行就被送往屏幕。行被处理完成之后，就被移出模式空间，程序接着读入下一行，处理，显示，移出......文件输入的最后一行被处理完以后sed结束。通过存储每一行在临时缓冲区，然后在缓冲区中操作该行，保证了原始文件不会被破坏。
    
1、  sed的命令和选项：

命令	功能描述
a\	 在当前行的后面加入一行或者文本。
c\	 用新的文本改变或者替代本行的文本。
d	 从pattern space位置删除行。
i\	 在当前行的上面插入文本。
h	 拷贝pattern space的内容到holding buffer(特殊缓冲区)。
H	 追加pattern space的内容到holding buffer。
g	 获得holding buffer中的内容，并替代当前pattern space中的文本。
G	 获得holding buffer中的内容，并追加到当前pattern space的后面。
n	 读取下一个输入行，用下一个命令处理新的行而不是用第一个命令。
p	 打印pattern space中的行。
P	 打印pattern space中的第一行。
q	 退出sed。
w file	 写并追加pattern space到file的末尾。
!	 表示后面的命令对所有没有被选定的行发生作用。
s/re/string	 用string替换正则表达式re。
=	 打印当前行号码。
替换标记	 
g	 行内全面替换，如果没有g，只替换第一个匹配。
p	 打印行。
x	 互换pattern space和holding buffer中的文本。
y	 把一个字符翻译为另一个字符(但是不能用于正则表达式)。
选项	 
-e	 允许多点编辑。
-n	 取消默认输出。
     需要说明的是，sed中的正则和grep的基本相同，完全可以参照本系列的第一篇中的详细说明。

2.  sed实例：

    /> cat testfile
    northwest       NW     Charles Main           3.0      .98      3       34
    western          WE      Sharon Gray           5.3      .97     5       23
    southwest       SW     Lewis Dalsass          2.7      .8      2       18
    southern         SO      Suan Chin               5.1     .95     4       15
    southeast       SE       Patricia Hemenway   4.0      .7      4       17
    eastern           EA      TB Savage               4.4     .84     5       20
    northeast        NE      AM Main Jr.              5.1     .94     3       13
    north              NO      Margot Weber         4.5     .89     5       9
    central            CT      Ann Stephens          5.7     .94     5       13

    /> sed '/north/p' testfile #如果模板north被找到，sed除了打印所有行之外，还有打印匹配行。
    northwest       NW      Charles Main           3.0     .98     3       34
    northwest       NW      Charles Main           3.0     .98     3       34
    western          WE      Sharon Gray           5.3     .97     5       23
    southwest      SW      Lewis Dalsass          2.7     .8       2       18
    southern        SO       Suan Chin               5.1     .95     4       15
    southeast       SE       Patricia Hemenway   4.0     .7       4       17
    eastern           EA      TB Savage               4.4     .84     5       20
    northeast        NE      AM Main Jr.              5.1     .94     3       13
    northeast        NE      AM Main Jr.              5.1     .94     3       13
    north              NO      Margot Weber         4.5     .89     5       9
    north              NO      Margot Weber         4.5     .89     5       9
    central            CT      Ann Stephens          5.7     .94     5       13

    #-n选项取消了sed的默认行为。在没有-n的时候，包含模板的行被打印两次，但是在使用-n的时候将只打印包含模板的行。
    /> sed -n '/north/p' testfile
    northwest       NW      Charles Main    3.0     .98     3       34
    northeast        NE      AM Main Jr.       5.1     .94     3       13
    north              NO      Margot Weber  4.5     .89     5       9

    /> sed '3d' testfile  #第三行被删除，其他行默认输出到屏幕。
    northwest       NW     Charles Main            3.0     .98     3       34
    western          WE      Sharon Gray           5.3     .97     5       23
    southern         SO      Suan Chin               5.1     .95     4       15
    southeast       SE       Patricia Hemenway   4.0     .7       4       17
    eastern           EA      TB Savage               4.4     .84     5       20
    northeast       NE       AM Main Jr.              5.1     .94     3       13
    north             NO       Margot Weber         4.5     .89     5       9
    central           CT       Ann Stephens          5.7     .94     5       13

    /> sed '3,$d' testfile  #从第三行删除到最后一行，其他行被打印。$表示最后一行。
    northwest       NW      Charles Main    3.0     .98     3       34
    western          WE      Sharon Gray    5.3     .97     5       23

    /> sed '$d' testfile    #删除最后一行，其他行打印。
    northwest       NW     Charles Main           3.0     .98     3       34
    western          WE     Sharon Gray           5.3     .97     5       23
    southwest       SW    Lewis Dalsass          2.7     .8      2       18
    southern         SO     Suan Chin              5.1     .95     4       15
    southeast       SE      Patricia Hemenway   4.0     .7      4       17
    eastern           EA      TB Savage             4.4     .84     5       20
    northeast       NE      AM Main Jr.             5.1     .94     3       13
    north             NO      Margot Weber        4.5     .89     5       9

    /> sed '/north/d' testfile #删除所有包含north的行，其他行打印。
    western           WE      Sharon Gray           5.3     .97     5       23
    southwest       SW      Lewis Dalsass          2.7     .8      2       18
    southern          SO      Suan Chin               5.1     .95     4       15
    southeast         SE      Patricia Hemenway   4.0     .7       4       17
    eastern            EA      TB Savage               4.4     .84     5       20
    central             CT      Ann Stephens          5.7     .94     5       13

    #s表示替换，g表示命令作用于整个当前行。如果该行存在多个west，都将被替换为north，如果没有g，则只是替换第一个匹配。
    /> sed 's/west/north/g' testfile
    northnorth      NW     Charles Main           3.0     .98    3       34
    northern         WE      Sharon Gray          5.3     .97    5       23
    southnorth      SW     Lewis Dalsass         2.7     .8      2       18
    southern         SO      Suan Chin              5.1     .95    4       15
    southeast       SE      Patricia Hemenway   4.0     .7      4       17
    eastern           EA      TB Savage             4.4     .84     5       20
    northeast       NE      AM Main Jr.              5.1     .94    3       13
    north             NO      Margot Weber        4.5     .89     5       9
    central            CT      Ann Stephens        5.7     .94     5       13

    /> sed -n 's/^west/north/p' testfile #-n表示只打印匹配行，如果某一行的开头是west，则替换为north。
    northern        WE      Sharon Gray     5.3     .97     5       23

    #&符号表示替换字符串中被找到的部分。所有以两个数字结束的行，最后的数字都将被它们自己替换，同时追加.5。
    /> sed 's/[0-9][0-9]$/&.5/' testfile
    northwest       NW      Charles Main          3.0     .98     3       34.5
    western          WE      Sharon Gray           5.3     .97     5       23.5
    southwest       SW      Lewis Dalsass        2.7     .8       2       18.5
    southern         SO      Suan Chin              5.1     .95     4       15.5
    southeast       SE      Patricia Hemenway   4.0     .7       4       17.5
    eastern           EA      TB Savage              4.4     .84     5       20.5
    northeast        NE      AM Main Jr.             5.1     .94     3       13.5
    north              NO      Margot Weber        4.5     .89     5       9
    central            CT      Ann Stephens         5.7     .94     5       13.5

    /> sed -n 's/Hemenway/Jones/gp' testfile  #所有的Hemenway被替换为Jones。-n选项加p命令则表示只打印匹配行。
    southeast       SE      Patricia Jones  4.0     .7      4       17

    #模板Mar被包含在一对括号中，并在特殊的寄存器中保存为tag 1，它将在后面作为\1替换字符串，Margot被替换为Marlianne。
    /> sed -n 's/\(Mar\)got/\1lianne/p' testfile
    north           NO      Marlianne Weber 4.5     .89     5       9

    #s后面的字符一定是分隔搜索字符串和替换字符串的分隔符，默认为斜杠，但是在s命令使用的情况下可以改变。不论什么字符紧跟着s命令都认为是新的分隔符。这个技术在搜索含斜杠的模板时非常有用，例如搜索时间和路径的时候。
    /> sed 's#3#88#g' testfile
    northwest       NW      Charles Main            88.0    .98     88     884
    western          WE       Sharon Gray           5.88    .97     5       288
    southwest       SW      Lewis Dalsass          2.7     .8       2       18
    southern         SO       Suan Chin               5.1     .95     4       15
    southeast       SE        Patricia Hemenway   4.0     .7        4       17
    eastern           EA       TB Savage               4.4     .84      5      20
    northeast        NE       AM Main Jr.              5.1     .94      88     188
    north              NO       Margot Weber         4.5     .89      5       9
    central            CT       Ann Stephens          5.7     .94      5       188

    #所有在模板west和east所确定的范围内的行都被打印，如果west出现在esst后面的行中，从west开始到下一个east，无论这个east出现在哪里，二者之间的行都被打印，即使从west开始到文件的末尾还没有出现east，那么从west到末尾的所有行都将打印。
    /> sed -n '/west/,/east/p' testfile
    northwest       NW      Charles Main           3.0     .98      3      34
    western          WE      Sharon Gray            5.3     .97     5      23
    southwest       SW     Lewis Dalsass          2.7     .8       2      18
    southern         SO      Suan Chin               5.1     .95     4      15
    southeast        SE      Patricia Hemenway    4.0     .7       4      17

    /> sed -n '5,/^northeast/p' testfile  #打印从第五行开始到第一个以northeast开头的行之间的所有行。
    southeast       SE      Patricia Hemenway   4.0     .7       4       17
    eastern           EA      TB Savage              4.4     .84     5       20
    northeast        NE      AM Main Jr.             5.1     .94     3       13

    #-e选项表示多点编辑。第一个编辑命令是删除第一到第三行。第二个编辑命令是用Jones替换Hemenway。
    /> sed -e '1,3d' -e 's/Hemenway/Jones/' testfile
    southern        SO      Suan Chin          5.1     .95     4       15
    southeast       SE      Patricia Jones      4.0     .7      4       17
    eastern          EA      TB Savage          4.4     .84     5       20
    northeast       NE      AM Main Jr.         5.1     .94     3       13
    north             NO      Margot Weber    4.5     .89     5       9
    central           CT      Ann Stephens     5.7     .94     5       13

    /> sed -n '/north/w newfile' testfile #将所有匹配含有north的行写入newfile中。
    /> cat newfile
    northwest       NW      Charles Main     3.0     .98     3       34
    northeast       NE      AM Main Jr.         5.1     .94     3       13
    north             NO      Margot Weber    4.5     .89     5       9

    /> sed '/eastern/i\ NEW ENGLAND REGION' testfile #i是插入命令，在匹配模式行前插入文本。
    northwest       NW      Charles Main          3.0     .98      3       34
    western          WE      Sharon Gray           5.3     .97     5       23
    southwest       SW      Lewis Dalsass         2.7     .8      2       18
    southern         SO      Suan Chin              5.1     .95     4       15
    southeast        SE      Patricia Hemenway   4.0     .7      4       17
    NEW ENGLAND REGION
    eastern          EA      TB Savage              4.4     .84     5       20
    northeast       NE      AM Main Jr.             5.1     .94     3       13
    north             NO      Margot Weber        4.5     .89     5       9
    central           CT      Ann Stephens         5.7     .94     5       13

    #找到匹配模式eastern的行后，执行后面花括号中的一组命令，每个命令之间用逗号分隔，n表示定位到匹配行的下一行，s/AM/Archie/完成Archie到AM的替换，p和-n选项的合用，则只是打印作用到的行。
    /> sed -n '/eastern/{n;s/AM/Archie/;p}' testfile
    northeast       NE      Archie Main Jr. 5.1     .94     3       13

-e表示多点编辑，第一个编辑命令y将前三行中的所有小写字母替换为大写字母，-n表示不显示替换后的输出，第二个编辑命令将只是打印输出转换后的前三行。注意y不能用于正则。

    /> sed -n -e '1,3y/abcdefghijklmnopqrstuvwxyz/ABCDEFGHIJKLMNOPQRSTUVWXYZ/' -e '1,3p' testfile
    NORTHWEST       NW      CHARLES MAIN     3.0     .98     3       34
    WESTERN           WE      SHARON GRAY      5.3     .97     5       23
    SOUTHWEST       SW      LEWIS DALSASS   2.7     .8      2       18

    /> sed '2q' testfile  #打印完第二行后退出。
    northwest       NW      Charles Main    3.0     .98     3       34
    western          WE      Sharon Gray     5.3     .97     5       23

当模板Lewis在某一行被匹配，替换命令首先将Lewis替换为Joseph，然后再用q退出sed。

     /> sed '/Lewis/{s/Lewis/Joseph/;q;}' testfile
    northwest       NW      Charles Main      3.0     .98     3       34
    western          WE      Sharon Gray      5.3     .97     5       23
    southwest       SW      Joseph Dalsass  2.7     .8      2       18

在sed处理文件的时候，每一行都被保存在pattern space的临时缓冲区中。除非行被删除或者输出被取消，否则所有被处理过的行都将打印在屏幕上。接着pattern space被清空，并存入新的一行等待处理。在下面的例子中，包含模板的northeast行被找到，并被放入pattern space中，h命令将其复制并存入一个称为holding buffer的特殊缓冲区内。在第二个sed编辑命令中，当达到最后一行后，G命令告诉sed从holding buffer中取得该行，然后把它放回到pattern space中，且追加到现在已经存在于模式空间的行的末尾。

     /> sed -e '/northeast/h' -e '$G' testfile
    northwest       NW     Charles Main            3.0    .98     3       34
    western          WE     Sharon Gray            5.3    .97     5       23
    southwest       SW    Lewis Dalsass          2.7     .8       2       18
    southern         SO     Suan Chin               5.1     .95     4       15
    southeast       SE      Patricia Hemenway   4.0     .7       4       17
    eastern           EA      TB Savage              4.4     .84     5       20
    northeast       NE      AM Main Jr.              5.1     .94     3       13
    north             NO      Margot Weber         4.5     .89     5       9
    central           CT      Ann Stephens          5.7     .94     5       13
    northeast       NE      AM Main Jr.              5.1     .94     3       13

    #如果模板WE在某一行被匹配，h命令将使得该行从pattern space中复制到holding buffer中，d命令在将该行删除，因此WE匹配行没有在原来的位置被输出。第二个命令搜索CT，一旦被找到，G命令将从holding buffer中取回行，并追加到当前pattern space的行末尾。简单的说，WE所在的行被移动并追加到包含CT行的后面。
    /> sed -e '/WE/{h;d;}' -e '/CT/{G;}' testfile
    northwest       NW    Charles Main           3.0     .98     3       34
    southwest       SW    Lewis Dalsass         2.7     .8      2       18
    southern         SO     Suan Chin              5.1     .95     4       15
    southeast       SE      Patricia Hemenway   4.0     .7      4       17
    eastern           EA     TB Savage              4.4     .84     5       20
    northeast       NE      AM Main Jr.              5.1     .94     3       13
    north             NO      Margot Weber         4.5     .89     5       9
    central           CT      Ann Stephens          5.7     .94     5       13
    western         WE      Sharon Gray           5.3     .97     5       23

    #第一个命令将匹配northeast的行从pattern space复制到holding buffer，第二个命令在读取的文件的末尾时，g命令告诉sed从holding buffer中取得行，并把它放回到pattern space中，以替换已经存在于pattern space中的。简单说就是包含模板northeast的行被复制并覆盖了文件的末尾行。
    /> sed -e '/northeast/h' -e '$g' testfile
    northwest       NW     Charles Main          3.0     .98     3       34
    western          WE      Sharon Gray         5.3     .97      5       23
    southwest       SW     Lewis Dalsass        2.7     .8       2       18
    southern         SO      Suan Chin             5.1     .95     4       15
    southeast       SE      Patricia Hemenway   4.0     .7      4       17
    eastern           EA      TB Savage             4.4     .84     5       20
    northeast       NE      AM Main Jr.             5.1     .94     3       13
    north             NO      Margot Weber        4.5     .89     5       9
    northeast       NE      AM Main Jr.             5.1     .94     3       13

    #模板WE匹配的行被h命令复制到holding buffer，再被d命令删除。结果可以看出WE的原有位置没有输出。第二个编辑命令将找到匹配CT的行，g命令将取得holding buffer中的行，并覆盖当前pattern space中的行，即匹配CT的行。简单的说，任何包含模板northeast的行都将被复制，并覆盖包含CT的行。    
    /> sed -e '/WE/{h;d;}' -e '/CT/{g;}' testfile
    northwest       NW    Charles Main           3.0     .98      3      34
    southwest       SW    Lewis Dalsass         2.7     .8       2       18
    southern         SO     Suan Chin              5.1     .95      4      15
    southeast       SE      Patricia Hemenway   4.0     .7       4      17
    eastern          EA      TB Savage              4.4     .84      5      20
    northeast       NE      AM Main Jr.              5.1     .94     3      13
    north             NO      Margot Weber        4.5     .89      5      9
    western         WE      Sharon Gray           5.3     .97     5      23

    #第一个编辑中的h命令将匹配Patricia的行复制到holding buffer中，第二个编辑中的x命令，会将holding buffer中的文本考虑到pattern space中，而pattern space中的文本被复制到holding buffer中。因此在打印匹配Margot行的地方打印了holding buffer中的文本，即第一个命令中匹配Patricia的行文本，第三个编辑命令会将交互后的holding buffer中的文本在最后一行的后面打印出来。
     /> sed -e '/Patricia/h' -e '/Margot/x' -e '$G' testfile
    northwest       NW      Charles Main           3.0      .98      3       34
    western           WE      Sharon Gray           5.3     .97      5       23
    southwest       SW      Lewis Dalsass         2.7      .8       2       18
    southern         SO      Suan Chin               5.1      .95     4       15
    southeast       SE       Patricia Hemenway    4.0      .7       4       17
    eastern           EA      TB Savage               4.4      .84     5       20
    northeast       NE       AM Main Jr.               5.1     .94      3       13
    southeast       SE      Patricia Hemenway      4.0     .7       4       17
    central            CT      Ann Stephens            5.7     .94     5       13

## 九.  awk实用功能:

    和sed一样，awk也是逐行扫描文件的，从第一行到最后一行，寻找匹配特定模板的行，并在这些行上运行“选择”动作。如果一个模板没有指定动作，这些匹配的行就被显示在屏幕上。如果一个动作没有模板，所有被动作指定的行都被处理。
    
   1.  awk的基本格式：
    /> awk 'pattern' filename
    /> awk '{action}' filename
    /> awk 'pattern {action}' filename
    
    具体应用方式分别见如下三个用例：
    /> cat employees
    Tom Jones         4424    5/12/66         543354
    Mary Adams      5346    11/4/63         28765
    Sally Chang       1654    7/22/54         650000
    Billy Black         1683    9/23/44         336500

    /> awk '/Mary/' employees   #打印所有包含模板Mary的行。
    Mary Adams      5346    11/4/63         28765

    #打印文件中的第一个字段，这个域在每一行的开始，缺省由空格或其它分隔符。
    /> awk '{print $1}' employees 
    Tom
    Mary
    Sally
    Billy
    
    /> awk '/Sally/{print $1, $2}' employees #打印包含模板Sally的行的第一、第二个域字段。
    Sally Chang
    
    2.  awk的格式输出：
    awk中同时提供了print和printf两种打印输出的函数，其中print函数的参数可以是变量、数值或者字符串。字符串必须用双引号引用，参数用逗号分隔。如果没有逗号，参数就串联在一起而无法区分。这里，逗号的作用与输出文件的分隔符的作用是一样的，只是后者是空格而已。下面给出基本的转码序列：

转码	含义
\n	换行
\r	回车
\t	制表符

    /> date | awk '{print "Month: " $2 "\nYear: ", $6}'
    Month: Oct
    Year:  2011

    /> awk '/Sally/{print "\t\tHave a nice day, " $1,$2 "\!"}' employees
                    Have a nice day, Sally Chang!

    在打印数字的时候你也许想控制数字的格式，我们通常用printf来完成这个功能。awk的特殊变量OFMT也可以在使用print函数的时候，控制数字的打印格式。它的默认值是"%.6g"----小数点后面6位将被打印。
    /> awk 'BEGIN { OFMT="%.2f"; print 1.2456789, 12E-2}'
    1.25  0.12

    现在我们介绍一下功能更为强大的printf函数，其用法和c语言中printf基本相似。下面我们给出awk中printf的格式化说明符列表：

格式化说明符	功能	示例	结果
%c	打印单个ASCII字符。	printf("The character is %c.\n",x)	The character is A.
%d	打印十进制数。	printf("The boy is %d years old.\n",y)	The boy is 15 years old.
%e	打印用科学记数法表示的数。	printf("z is %e.\n",z)	z is 2.3e+01.
%f	打印浮点数。	printf("z is %f.\n",z)	z is 2.300000
%o	打印八进制数。	printf("y is %o.\n",y)	y is 17.
%s	打印字符串。	printf("The name of the culprit is %s.\n",$1);	The name of the culprit is Bob Smith.
%x	打印十六进制数。	printf("y is %x.\n",y)	y is f.
    注：假设列表中的变脸值为x = A, y = 15, z = 2.3, $1 = "Bob Smith"

    /> echo "Linux" | awk '{printf "|%-15s|\n", $1}'  # %-15s表示保留15个字符的空间，同时左对齐。
    |Linux          |

    /> echo "Linux" | awk '{printf "|%15s|\n", $1}'   # %-15s表示保留15个字符的空间，同时右对齐。
    |          Linux|

    #%8d表示数字右对齐，保留8个字符的空间。
     /> awk '{printf "The name is %-15s ID is %8d\n", $1,$3}' employees
    The name is Tom             ID is     4424
    The name is Mary            ID is     5346
    The name is Sally            ID is     1654
    The name is Billy             ID is     1683

    3.  awk中的记录和域：
    awk中默认的记录分隔符是回车，保存在其内建变量ORS和RS中。$0变量是指整条记录。
    /> awk '{print $0}' employees #这等同于print的默认行为。
    Tom Jones        4424    5/12/66         543354
    Mary Adams      5346    11/4/63         28765
    Sally Chang       1654    7/22/54         650000
    Billy Black         1683    9/23/44         336500

    变量NR(Number of Record)，记录每条记录的编号。
    /> awk '{print NR, $0}' employees
    1 Tom Jones        4424    5/12/66         543354
    2 Mary Adams      5346    11/4/63         28765
    3 Sally Chang       1654    7/22/54         650000
    4 Billy Black         1683    9/23/44         336500

    变量NF(Number of Field)，记录当前记录有多少域。
    /> awk '{print $0,NF}' employees
    Tom Jones        4424    5/12/66          543354   5
    Mary Adams      5346    11/4/63         28765     5
    Sally Chang      1654    7/22/54          650000   5
    Billy Black        1683     9/23/44         336500    5

    #根据employees生成employees2。sed的用法可以参考上一篇blog。
    /> sed 's/[[:space:]]\+\([0-9]\)/:\1/g;w employees2' employees
    /> cat employees
    Tom Jones:4424:5/12/66:543354
    Mary Adams:5346:11/4/63:28765
    Sally Chang:1654:7/22/54:650000
    Billy Black:1683:9/23/44:336500

    /> awk -F: '/Tom Jones/{print $1,$2}' employees2  #这里-F选项后面的字符表示分隔符。
    Tom Jones 4424

    变量OFS(Output Field Seperator)表示输出字段间的分隔符，缺省是空格。
    />  awk -F: '{OFS = "?"};  /Tom/{print $1,$2 }' employees2 #在输出时，域字段间的分隔符已经是?(问号)了
    Tom Jones?4424

    对于awk而言，其模式部分将控制这动作部分的输入，只有符合模式条件的记录才可以交由动作部分基础处理，而模式部分不仅可以写成正则表达式(如上面的例子)，awk还支持条件表达式，如：
    /> awk '$3 < 4000 {print}' employees
    Sally Chang     1654    7/22/54         650000
    Billy Black       1683    9/23/44         336500

    在花括号内，用分号分隔的语句称为动作。如果模式在动作前面，模式将决定什么时候发出动作。动作可以是一个语句或是一组语句。语句之间用分号分隔，也可以用换行符，如：
    pattern { action statement; action statement; etc. } or
    pattern {
        action statement
        action statement
    }
    模式和动作一般是捆绑在一起的。需要注意的是，动作是花括号内的语句。模式控制的动作是从第一个左花括号开始到第一个右花括号结束，如下：
    /> awk '$3 < 4000 && /Sally/ {print}' employees
    Sally Chang     1654    7/22/54         650000

    4.  匹配操作符：
    " ~ " 用来在记录或者域内匹配正则表达式。
    /> awk '$1 ~ /[Bb]ill/' employees      #显示所有第一个域匹配Bill或bill的行。
    Billy Black     1683    9/23/44         336500

    /> awk '$1 !~ /[Bb]ill/' employees     #显示所有第一个域不匹配Bill或bill的行，其中!~表示不匹配的意思。
    Tom Jones        4424    5/12/66         543354
    Mary Adams      5346    11/4/63         28765
    Sally Chang       1654    7/22/54         650000

    5.  awk的基本应用实例：
    /> cat testfile
    northwest     NW        Charles Main            3.0        .98        3        34
    western        WE        Sharon Gray            5.3        .97        5        23
    southwest     SW        Lewis Dalsass          2.7        .8          2        18
    southern       SO        Suan Chin                5.1        .95        4        15
    southeast      SE        Patricia Hemenway    4.0        .7          4        17
    eastern         EA        TB Savage                4.4        .84        5        20
    northeast      NE        AM Main Jr.               5.1        .94        3        13
    north            NO        Margot Weber          4.5        .89        5        9
    central          CT        Ann Stephens           5.7        .94        5        13

    /> awk '/^north/' testfile            #打印所有以north开头的行。
    northwest      NW      Charles Main     3.0     .98     3       34
    northeast       NE      AM Main Jr.        5.1     .94     3       13
    north             NO      Margot Weber   4.5     .89     5       9

    /> awk '/^(no|so)/' testfile          #打印所有以so和no开头的行。
    northwest       NW      Charles Main                3.0     .98      3       34
    southwest       SW      Lewis Dalsass              2.7     .8       2       18
    southern         SO      Suan Chin                    5.1     .95     4       15
    southeast        SE      Patricia Hemenway        4.0     .7       4       17
    northeast        NE      AM Main Jr.                   5.1     .94     3       13
    north              NO      Margot Weber              4.5     .89     5       9

    /> awk '$5 ~ /\.[7-9]+/' testfile     #第五个域字段匹配包含.(点)，后面是7-9的数字。
    southwest       SW      Lewis Dalsass            2.7     .8      2       18
    central             CT      Ann Stephens            5.7     .94     5       13

    /> awk '$8 ~ /[0-9][0-9]$/{print $8}' testfile  #第八个域以两个数字结束的打印。
    34
    23
    18
    15
    17
    20
    13

## 十.  awk表达式功能:

    1.  比较表达式：
    比较表达式匹配那些只在条件为真时才运行的行。这些表达式利用关系运算符来比较数字和字符串。见如下awk支持的条件表达式列表：

运算符	含义	例子
<	小于	x < y
<=	小于等于	x <= y
==	等于	x == y
!=	不等于	x != y
>=	大于等于	x >= y
>	大于	x > y
~	匹配	x ~ /y/
!~	不匹配	x !~ /y/
    /> cat employees
    Tom Jones        4424    5/12/66         543354
    Mary Adams      5346    11/4/63         28765
    Sally Chang       1654    7/22/54         650000
    Billy Black         1683    9/23/44         336500

    /> awk '$3 == 5346' employees       #打印第三个域等于5346的行。
    Mary Adams      5346    11/4/63         28765

    /> awk '$3 > 5000 {print $1}' employees  #打印第三个域大于5000的行的第一个域字段。
    Mary

    /> awk '$2 ~ /Adam/' employess      #打印第二个域匹配Adam的行。
    Mary Adams      5346    11/4/63         28765

    2.  条件表达式：
    条件表达式使用两个符号--问号和冒号给表达式赋值： conditional expression1 ? expression2 : expressional3，其逻辑等同于C语言中的条件表达式。其对应的if/else语句如下：
    {
        if (expression1)
            expression2
        else
            expression3
    }
    /> cat testfile
    northwest     NW        Charles Main             3.0        .98        3        34
    western        WE        Sharon Gray             5.3        .97         5        23
    southwest     SW        Lewis Dalsass           2.7        .8          2        18
    southern       SO        Suan Chin                 5.1        .95        4        15
    southeast      SE        Patricia Hemenway     4.0        .7          4        17
    eastern         EA        TB Savage                 4.4        .84        5        20
    northeast      NE        AM Main Jr.                5.1       .94         3        13
    north            NO        Margot Weber           4.5       .89         5        9
    central          CT        Ann Stephens            5.7       .94         5        13

    /> awk 'NR <= 3 {print ($7 > 4 ? "high "$7 : "low "$7) }' testfile
    low 3
    high 5
    low 2

    3.  数学表达式：
    运算可以在模式内进行，其中awk将所有的运算都视为浮点运算，见如下列表：

运算符	含义	例子
+	加	x + y
-	减	x - y
*	乘	x * y
/	除	x / y
%	取余	x % y
^	乘方	x ^ y
    /> awk '/southern/{print $5 + 10}' testfile  #如果记录包含正则表达式southern，第五个域就加10并打印。
    15.1

    /> awk '/southern/{print $8 /2 }' testfile   #如果记录包含正则表达式southern，第八个域除以2并打印。
    7.5

    4.  逻辑表达式：
    见如下列表：

运算符	含义	例子
&&	逻辑与	a && b
||	逻辑或	a || b
!	逻辑非	!a
    /> awk '$8 > 10 && $8 < 17' testfile   #打印出第八个域的值大于10小于17的记录。
    southern        SO      Suan Chin               5.1     .95     4       15
    central            CT      Ann Stephens         5.7     .94     5       13

    #打印第二个域等于NW，或者第一个域匹配south的行的第一、第二个域。
    /> awk '$2 == "NW" || $1 ~ /south/ {print $1,$2}' testfile
    northwest  NW
    southwest  SW
    southern    SO
    southeast   SE

    /> awk '!($8 > 13) {print $8}' testfile  #打印第八个域字段不大于13的行的第八个域。
    3
    9
    13

    5.  范围模板：
    范围模板匹配从第一个模板的第一次出现到第二个模板的第一次出现，第一个模板的下一次出现到第一个模板的下一次出现等等。如果第一个模板匹配而第二个模板没有出现，awk就显示到文件末尾的所有行。
    /> awk '/^western/,/^eastern/ {print $1}' testfile #打印以western开头到eastern开头的记录的第一个域。
    western    WE
    southwest SW
    southern   SO
    southeast  SE
    eastern      EA    

    6.  赋值符号：
    #找到第三个域等于Ann的记录，然后给该域重新赋值为Christian，之后再打印输出该记录。
    /> awk '$3 == "Ann" { $3 = "Christian"; print}' testfile
    central CT Christian Stephens 5.7 .94 5 13

    /> awk '/Ann/{$8 += 12; print $8}' testfile #找到包含Ann的记录，并将该条记录的第八个域的值+=12，最后再打印输出。
    25

## 十一.  awk编程:

    1.  变量：
    在awk中变量无须定义即可使用，变量在赋值时即已经完成了定义。变量的类型可以是数字、字符串。根据使用的不同，未初始化变量的值为0或空白字符串" "，这主要取决于变量应用的上下文。下面为变量的赋值负号列表：

符号	含义	等价形式
=	a = 5	a = 5
+=	a = a + 5	a += 5
-=	a = a - 5	a -= 5
*=	a = a * 5	a *= 5
/=	a = a / 5	a /= 5
%=	a = a % 5	a %= 5
^=	a = a ^ 5	a ^= 5
    /> awk '$1 ~ /Tom/ {Wage = $2 * $3; print Wage}' filename
    该命令将从文件中读取，并查找第一个域字段匹配Tom的记录，再将其第二和第三个字段的乘积赋值给自定义的Wage变量，最后通过print命令将该变量打印输出。

    /> awk ' {$5 = 1000 * $3 / $2; print}' filename
    在上面的命令中，如果$5不存在，awk将计算表达式1000 * $3 / $2的值，并将其赋值给$5。如果第五个域存在，则用表达式覆盖$5原来的值。

    我们同样也可以在命令行中定义自定义的变量，用法如下：
    /> awk -F: -f awkscript month=4 year=2011 filename
    这里的month和year都是自定义变量，且分别被赋值为4和2000，在awk的脚本中这些变量将可以被直接使用，他们和脚本中定义的变量在使用上没有任何区别。

    除此之外，awk还提供了一组内建变量(变量名全部大写)，见如下列表：

变量名	变量内容
ARGC	命令行参数的数量。
ARGIND	命令行正在处理的当前文件的AGV的索引。
ARGV	命令行参数数组。
CONVFMT	转换数字格式。
ENVIRON	从shell中传递来的包含当前环境变量的数组。
ERRNO	当使用close函数或者通过getline函数读取的时候，发生的重新定向错误的描述信息就保存在这个变量中。
FIELDWIDTHS	在对记录进行固定域宽的分割时，可以替代FS的分隔符的列表。
FILENAME	当前的输入文件名。
FNR	当前文件的记录号。
FS	输入分隔符，默认是空格。
IGNORECASE	在正则表达式和字符串操作中关闭大小写敏感。
NF	当前文件域的数量。
NR	当前文件记录数。
OFMT	数字输出格式。
OFS	输出域分隔符。
ORS	输出记录分隔符。
RLENGTH	通过match函数匹配的字符串的长度。
RS	输入记录分隔符。
RSTART	通过match函数匹配的字符串的偏移量。
SUBSEP	下标分隔符。
    /> cat employees2
    Tom Jones:4424:5/12/66:543354
    Mary Adams:5346:11/4/63:28765
    Sally Chang:1654:7/22/54:650000
    Mary Black:1683:9/23/44:336500

    /> awk -F: '{IGNORECASE = 1}; $1 == "mary adams" { print NR, $1, $2, $NF}' employees2
    2 Mary Adams 5346 28765
    /> awk -F: ' $1 == "mary adams" { print NR, $1, $2, $NF}' employees2
    没有输出结果。
    当IGNORECASE内置变量的值为非0时，表示在进行字符串操作和处理正则表达式时关闭大小写敏感。这里的"mary adams"将匹配文件中的"Mary Admams"记录。最后print打印出第一、第二和最后一个域。需要说明的是NF表示当前记录域的数量，因此$NF将表示最后一个域的值。

    awk在动作部分还提供了BEGIN块和END块。其中BEGIN动作块在awk处理任何输入文件行之前执行。事实上，BEGIN块可以在没有任何输入文件的条件下测试。因为在BEGIN块执行完毕以前awk将不读取任何输入文件。BEGIN块通常被用来改变内建变量的值，如OFS、RS或FS等。也可以用于初始化自定义变量值，或打印输出标题。
    /> awk 'BEGIN {FS = ":"; OFS = "\t"; ORS = "\n\n"} { print $1,$2,$3} filename
    上例中awk在处理文件之前，已经将域分隔符(FS)设置为冒号，输出文件域分隔符(OFS)设置为制表符，输出记录分隔符(ORS)被设置为两个换行符。BEGIN之后的动作模块中如果有多个语句，他们之间用分号分隔。
    和BEGIN恰恰相反，END模块中的动作是在整个文件处理完毕之后被执行的。
    /> awk 'END {print "The number of the records is " NR }' filename
    awk在处理输入文件之后，执行END模块中的动作，上例中NR的值是读入的最后一个记录的记录号。

    /> awk '/Mary/{count++} END{print "Mary was found " count " times." }' employees2
    Mary was found 2 times.

    /> awk '/Mary/{count++} END{print "Mary was found " count " times." }' employees2
    Mary was found 2 times.
    
    /> cat testfile
    northwest       NW      Charles Main                3.0     .98     3       34
    western          WE      Sharon Gray                5.3     .97     5       23
    southwest       SW      Lewis Dalsass              2.7     .8      2       18
    southern         SO      Suan Chin                   5.1     .95     4       15
    southeast        SE      Patricia Hemenway        4.0     .7      4       17
    eastern           EA      TB Savage                   4.4     .84     5       20
    northeast        NE      AM Main Jr.                  5.1     .94     3       13
    north             NO       Margot Weber             4.5     .89     5       9
    central           CT       Ann Stephens              5.7     .94     5       13

    /> awk '/^north/{count += 1; print count}' testfile     #如记录以正则north开头，则创建变量count同时增一，再输出其值。
    1
    2
    3
    
    #这里只是输出前三个字段，其中第七个域先被赋值给变量x，在自减一，最后再同时打印出他们。
    /> awk 'NR <= 3 {x = $7--; print "x = " x ", $7 = " $7}' testfile
    x = 3, $7 = 2
    x = 5, $7 = 4
    x = 2, $7 = 1    
    
    #打印NR(记录号)的值在2--5之间的记录。
    /> awk 'NR == 2,NR == 5 {print "The record number is " NR}' testfile
    The record number is 2
    The record number is 3
    The record number is 4
    The record number is 5

    #打印环境变量USER和HOME的值。环境变量的值由父进程shell传递给awk程序的。
    /> awk 'BEGIN { print ENVIRON["USER"],ENVIRON["HOME"]}' 
    root /root
    
    #BEGIN块儿中对OFS内置变量重新赋值了，因此后面的输出域分隔符改为了\t。
    /> awk 'BEGIN { OFS = "\t"}; /^Sharon/{ print $1,$2,$7}' testfile
    western WE      5
    
    #从输入文件中找到以north开头的记录count就加一，最后在END块中输出该变量。
    /> awk '/^north/{count++}; END{print count}' testfile
    3

    2.  重新定向：
    在动作语句中使用shell通用的重定向输出符号">"就可以完成awk的重定向操作，当使用>的时候，原有文件将被清空，同时文件持续打开，直到文件被明确的关闭或者awk程序终止。来自后面的打印语句的输出会追加到前面内容的后面。符号">>"用来打开一个文件但是不清空原有文件的内容，重定向的输出只是被追加到这个文件的末尾。
    /> awk '$4 >= 70 {print $1,$2 > "passing_file"}' filename  #注意这里的文件名需要用双引号括起来。
    #通过两次cat的结果可以看出>和>>的区别。
    /> awk '/north/{print $1,$3,$4 > "districts" }' testfile
    /> cat districts
    northwest Joel Craig
    northeast TJ Nichols
    north Val Shultz
    /> awk '/south/{print $1,$3,$4 >> "districts" }' testfile
    /> cat districts
    northwest Joel Craig
    northeast TJ Nichols
    north Val Shultz
    southwest Chris Foster
    southern May Chin
    southeast Derek Jonhson

    
    awk中对于输入重定向是通过getline函数来完成的。getline函数的作用是从标准输入、管道或者当前正在处理的文件之外的其他输入文件获得输入。他负责从输入获得下一行的内容，并给NF、NR和FNR等内建变量赋值。如果得到一个记录，getline就返回1，如果达到文件末尾就返回0。如果出现错误，如打开文件失败，就返回-1。
    /> awk 'BEGIN { "date" | getline d; print d}'
    Tue Nov 15 15:31:42 CST 2011
    上例中的BEGIN动作模块中，先执行shell命令date，并通过管道输出给getline，然后再把输出赋值给自定义变量d并打印输出它。
    
    /> awk 'BEGIN { "date" | getline d; split(d,mon); print mon[2]}'
    Nov
    上例中date命令通过管道输出给getline并赋值给d变量，再通过内置函数split将d拆分为mon数组，最后print出mon数组的第二个元素。
    
    /> awk 'BEGIN { while("ls" | getline) print}'
    employees
    employees2
    testfile
    命令ls的输出传递给getline作为输入，循环的每个反复，getline都从ls的结果中读取一行输入，并把他打印到屏幕。
    
    /> awk 'BEGIN { printf "What is your name? "; \
        getline name < "/dev/tty"}\
        $1 ~ name {print "Found" name " on line ", NR "."}\
        END {print "See ya, " name "."}' employees2
    What is your name? Mary
    Found Mary on line  2.
    See ya, Mary.    
    上例先是打印出BEGIN块中的"What is your name? "，然后等待用户从/dev/tty输入，并将读入的数据赋值给name变量，之后再从输入文件中读取记录，并找到匹配输入变量的记录并打印出来，最后在END块中输出结尾信息。
    
    /> awk 'BEGIN { while(getline < "/etc/passwd" > 0) lc++; print lc}'
    32
    awk将逐行读取/etc/passwd文件中的内容，在达到文件末尾之前，计数器lc一直自增1，当到了末尾后打印lc的值。lc的值为/etc/passwd文件的行数。
    由于awk中同时打开的管道只有一个，那么在打开下一个管道之前必须关闭它，管道符号右边可以通过可以通过双引号关闭管道。如果不关闭，它将始终保持打开状态，直到awk退出。
    /> awk {print $1,$2,$3 | "sort -4 +1 -2 +0 -1"} END {close("sort -4 +1 -2 +0 -1") } filename
    上例中END模块中的close显示关闭了sort的管道，需要注意的是close中关闭的命令必须和当初打开时的完全匹配，否则END模块产生的输出会和以前的输出一起被sort分类。


    3.  条件语句：
    awk中的条件语句是从C语言中借鉴来的，见如下声明方式：
    if (expression) {
        statement;
        statement;
        ... ...
    }
    /> awk '{if ($6 > 50) print $1 "Too hign"}' filename
    /> awk '{if ($6 > 20 && $6 <= 50) { safe++; print "OK}}' filename

    if (expression) {
        statement;
    } else {
        statement2;
    }
    /> awk '{if ($6 > 50) print $1 " Too high"; else print "Range is OK" }' filename
    /> awk '{if ($6 > 50) { count++; print $3 } else { x = 5; print $5 }' filename

    if (expression) {
        statement1;
    } else if (expression1) {
        statement2;
    } else {
        statement3;
    }
    /> awk '{if ($6 > 50) print "$6 > 50" else if ($6 > 30) print "$6 > 30" else print "other"}' filename

   4.  循环语句：
    awk中的循环语句同样借鉴于C语言，支持while、do/while、for、break、continue，这些关键字的语义和C语言中的语义完全相同。

    5.  流程控制语句：
    next语句是从文件中读取下一行，然后从头开始执行awk脚本。
    exit语句用于结束awk程序。它终止对记录的处理。但是不会略过END模块，如果exit()语句被赋值0--255之间的参数，如exit(1)，这个参数就被打印到命令行，以判断退出成功还是失败。

    6.  数组：
    因为awk中数组的下标可以是数字和字母，数组的下标通常被称为关键字(key)。值和关键字都存储在内部的一张针对key/value应用hash的表格里。由于hash不是顺序存储，因此在显示数组内容时会发现，它们并不是按照你预料的顺序显示出来的。数组和变量一样，都是在使用时自动创建的，awk也同样会自动判断其存储的是数字还是字符串。一般而言，awk中的数组用来从记录中收集信息，可以用于计算总和、统计单词以及跟踪模板被匹配的次数等等。
    /> cat employees
    Tom Jones       4424    5/12/66         543354
    Mary Adams      5346    11/4/63         28765
    Sally Chang     1654    7/22/54         650000
    Billy Black     1683    9/23/44         336500

    /> awk '{name[x++] = $2}; END{for (i = 0; i < NR; i++) print i, name[i]}' employees    
    0 Jones
    1 Adams
    2 Chang
    3 Black
    在上例中，数组name的下标是变量x。awk初始化该变量的值为0，在每次使用后自增1，读取文件中的第二个域的值被依次赋值给name数组的各个元素。在END模块中，for循环遍历数组的值。因为下标是关键字，所以它不一定从0开始，可以从任何值开始。

    #这里是用内置变量NR作为数组的下标了。
    /> awk '{id[NR] = $3}; END {for (x = 1; x <= NR; x++) print id[x]}' employees
    4424
    5346
    1654
    1683

    awk中还提供了一种special for的循环，见如下声明：
    for (item in arrayname) {
        print arrayname[item]
    }

    /> cat db
    Tom Jones
    Mary Adams
    Sally Chang
    Billy Black
    Tom Savage
    Tom Chung
    Reggie Steel
    Tommy Tucker

    /> awk '/^Tom/{name[NR]=$1}; END {for(i = 1;i <= NR; i++) print name[i]}' db
    Tom



    Tom
    Tom

    Tommy
    从输出结果可以看出，只有匹配正则表达式的记录的第一个域被赋值给数组name的指定下标元素。因为用NR作为下标，所以数组的下标不可能是连续的，因此在END模块中用传统的for循环打印时，不存在的元素就打印空字符串了。下面我们看看用special for的方式会有什么样的输出。
    /> awk '/^Tom/{name[NR]=$1};END{for(i in name) print name[i]}' db
    Tom
    Tom
    Tommy
    Tom

    下面我们看一下用字符串作为下标的例子：(如果下标是字符串文字常量，则需要用双引号括起来)    
    /> cat testfile2
    tom
    mary
    sean
    tom
    mary
    mary
    bob
    mary
    alex
    /> awk '/tom/{count["tom"]++}; /mary/{count["mary"]++}; END{print "There are " count["tom"] \
        " Toms and " count["mary"] " Marys in the file."} testfile2
    There are 2 Toms and 4 Marys in the file.
    在上例中，count数组有两个元素，下标分别为tom和mary，每一个元素的初始值都是0，没有tom被匹配的时候，count["tom"]就会加一，count["mary"]在匹配mary的时候也同样如此。END模块中打印出存储在数组中的各个元素。

    /> awk '{count[$1]++}; END{for(name in count) printf "%-5s%d\n",name, count[name]}' testfile2
    mary 4
    tom  2
    alex 1
    bob  1
    sean 1
    在上例中，awk是以记录的域作为数组count的下标。

    /> awk '{count[$1]++; if (count[$1] > 1) name[$1]++}; END{print "The duplicates were "; for(i in name) print i}' testfile2
    The duplicates were
    mary
    tom
    在上例中，如count[$1]的元素值大于1的时候，也就是当名字出现多次的时候，一个新的数组name将被初始化，最后打印出那么数组中重复出现的名字下标。

    之前我们介绍的都是如何给数组添加新的元素，并赋予初值，现在我们需要介绍一下如何删除数组中已经存在的元素。要完成这一功能我们需要使用内置函数delete，见如下命令：
    /> awk '{count[$1]++}; \
        END{for(name in count) {\
                if (count[name] == 1)\
                    delete count[name];\
            } \
            for (name in count) \
                print name}' testfile2
    mary
    tom
    上例中的主要技巧来自END模块，先是变量count数组，如果数组中某个元素的值等于1，则删除该元素，这样等同于删除只出现一次的名字。最后用special for循环打印出数组中仍然存在的元素下标名称。

    最后我们来看一下如何使用命令行参数数组，见如下命令：
    /> awk 'BEGIN {for(i = 0; i < ARGC; i++) printf("argv[%d] is %s.\n",i,ARGV[i]); printf("The number of arguments, ARGC=%d\n",ARGC)}' testfile "Peter Pan" 12
    argv[0] is awk.
    argv[1] is testfile.
    argv[2] is Peter Pan.
    argv[3] is 12.
    The number of arguments, ARGC=4
    从输出结果可以看出，命令行参数数组ARGV是以0作为起始下标的，命令行的第一个参数为命令本身(awk)，这个使用方式和C语句main函数完全一致。

    /> awk 'BEGIN{name=ARGV[2]; print "ARGV[2] is " ARGV[2]}; $1 ~ name{print $0}' testfile2 "bob"    
    ARGV[2] is bob
    bob
    awk: (FILENAME=testfile2 FNR=9) fatal: cannot open file `bob' for reading (No such file or directory)
    先解释一下以上命令的含义，name变量被赋值为命令行的第三个参数，即bob，之后再在输入文件中找到匹配该变量值的记录，并打印出该记录。
    在输出的第二行报出了awk的处理错误信息，这主要是因为awk将bob视为输入文件来处理了，然而事实上这个文件并不存在，下面我们需要做进一步的处理来修正这个问题。
    /> awk 'BEGIN{name=ARGV[2]; print "ARGV[2] is " ARGV[2]; delete ARGV[2]}; $1 ~ name{print $0}' testfile2 "bob"    
    ARGV[2] is bob
    bob
    从输出结果中我们可以看到我们得到了我们想要的结果。需要注意的是delete函数的调用必要要在BEGIN模块中完成，因为这时awk还没有开始读取命令行参数中指定的文件。

    7.  内建函数：
    字符串函数
    sub(regular expression,substitution string);
    sub(regular expression,substitution string,target string);

    /> awk '{sub("Tom","Tommy"); print}' employees   #这里使用Tommy替换了Tom。
    Tommy Jones       4424    5/12/66         543354

    #当正则表达式Tom在第一个域中第一次被匹配后，他将被字符串"Tommy"替换，如果将sub函数的第三个参数改为$2，将不会有替换发生。
    /> awk '{sub("Tom","Tommy",$1); print}' employees
    Tommy Jones       4424    5/12/66         543354

    gsub(regular expression,substitution string);
    gsub(regular expression,substitution string,target string);
    和sub不同的是，如果第一个参数中正则表达式在记录中出现多次，那么gsub将完成多次替换，而sub只是替换第一次出现的。

    index(string,substring)
    该函数将返回第二个参数在第一个参数中出现的位置，偏移量从1开始。
    /> awk 'BEGIN{print index("hello","el")}'
    2

    length(string)
    该函数返回字符串的长度。
    /> awk 'BEGIN{print length("hello")}'
    5

    substr(string,starting position)
    substr(string,starting position,length of string)
    该函数返回第一个参数的子字符串，其截取起始位置为第二个参数(偏移量为1)，截取长度为第三个参数，如果没有该参数，则从第二个参数指定的位置起，直到string的末尾。
    />  awk 'BEGIN{name = substr("Hello World",2,3); print name}'
    ell

    match(string,regular expression)
    该函数返回在字符串中正则表达式位置的索引，如果找不到指定的正则表达式就返回0.match函数设置内置变量RSTART为字符串中子字符串的开始位置，RLENGTH为到字字符串末尾的字符个数。
    /> awk 'BEGIN{start=match("Good ole CHINA", /[A-Z]+$/); print start}'
    10
    上例中的正则表达式[A-Z]+$表示在字符串的末尾搜索连续的大写字母。在字符串"Good ole CHINA"的第10个位置找到字符串"CHINA"。

    /> awk 'BEGIN{start=match("Good ole CHINA", /[A-Z]+$/); print RSTART, RLENGTH}'
    10 5
    RSTART表示匹配时的起始索引，RLENGTH表示匹配的长度。

    /> awk 'BEGIN{string="Good ole CHINA";start=match(string, /[A-Z]+$/); print substr(string,RSTART, RLENGTH)}'
    CHINA
    这里将match、RSTART、RLENGTH和substr巧妙的结合起来了。

    toupper(string)
    tolower(string)
    以上两个函数分别返回参数字符串的大写和小写的形式。
    /> awk 'BEGIN {print toupper("hello"); print tolower("WORLD")}'
    HELLO
    world

    split(string,array,field seperator)
    split(string,array)
    该函数使用作为第三个参数的域分隔符把字符串分隔为一个数组。如果第三个参数没有提供，则使用当前默认的FS值。
    /> awk 'BEGIN{split("11/20/2011",date,"/"); print date[2]}'
    20

    variable = sprintf("string with format specifiers ",expr1,expr2,...)
    该函数和printf的差别等同于C语言中printf和sprintf的差别。前者将格式化后的结果输出到输出流，而后者输出到函数的返回值中。
    /> awk 'BEGIN{line = sprintf("%-15s %6.2f ", "hello",4.2); print line}'
    hello             4.20

    时间函数：
    systime()
    该函数返回当前时间距离1970年1月1日之间相差的秒数。
    /> awk 'BEGIN{print systime()}'
    1321369554

    strftime()
    时间格式化函数，其格式化规则等同于C语言中的strftime函数提供的规则，见以下列表：

数据格式	含义
%a	Abbreviated weekday name
%A	Full weekday name
%b	Abbreviated month name
%B	Full month name
%c	Date and time representation appropriate for locale
%d	Day of month as decimal number (01 – 31)
%H	Hour in 24-hour format (00 – 23)
%I	Hour in 12-hour format (01 – 12)
%j	Day of year as decimal number (001 – 366)
%m	Month as decimal number (01 – 12)
%M	Minute as decimal number (00 – 59)
%p	Current locale's A.M./P.M. indicator for 12-hour clock
%S	Second as decimal number (00 – 59)
%U	Week of year as decimal number, with Sunday as first day of week (00 – 53)
%w	Weekday as decimal number (0 – 6; Sunday is 0)
%W	Week of year as decimal number, with Monday as first day of week (00 – 53)
%x	Date representation for current locale
%X	Time representation for current locale
%y	Year without century, as decimal number (00 – 99)
%Y	Year with century, as decimal number
    /> awk 'BEGIN{ print strftime("%D",systime())}'
    11/15/11
    /> awk 'BEGIN{ now = strftime("%T"); print now}'
    23:17:29

    内置数学函数：

名称	返回值
atan2(x,y)	y,x范围内的余切
cos(x)	余弦函数
exp(x)	求幂
int(x)	取整
log(x)	自然对数
sin(x)	正弦函数
sqrt(x)	平方根
    /> awk 'BEGIN{print 31/3}'
    10.3333
    /> awk 'BEGIN{print int(31/3)}'
    10

    自定义函数：
    自定义函数可以放在awk脚本的任何可以放置模板和动作的地方。
    function name(parameter1,parameter2,...) {
        statements
        return expression
    }
    给函数中本地变量传递值。只使用变量的拷贝。数组通过地址或者指针传递，所以可以在函数内部直接改变数组元素的值。函数内部使用的任何没有作为参数传递的变量都被看做是全局变量，也就是这些变量对于整个程序都是可见的。如果变量在函数中发生了变化，那么就是在整个程序中发生了改变。唯一向函数提供本地变量的办法就是把他们放在参数列表中，这些参数通常被放在列表的最后。如果函数调用没有提供正式的参数，那么参数就初始化为空。return语句通常就返回程序控制并向调用者返回一个值。
    /> cat grades
    20 10
    30 20
    40 30

    /> cat add.sc
    function add(first,second) {
            return first + second
    }
    { print add($1,$2) }

    /> awk -f add.sc grades
    30
    50
    70

## 十二.   行的排序命令sort:

  1.  sort命令行选项：

选项	描述
-t	字段之间的分隔符
-f	基于字符排序时忽略大小写
-k	定义排序的域字段，或者是基于域字段的部分数据进行排序
-m	将已排序的输入文件，合并为一个排序后的输出数据流
-n	以整数类型比较字段
-o outfile	将输出写到指定的文件
-r	倒置排序的顺序为由大到小，正常排序为由小到大
-u	只有唯一的记录，丢弃所有具有相同键值的记录
-b	忽略前面的空格

   2.  sort使用实例：
    提示：在下面的输出结果中红色标注的为第一排序字段，后面的依次为紫、绿。
    /> sed -n '1,5p' /etc/passwd > users
    /> cat users
    root:x:0:0:root:/root:/bin/bash
    bin:x:1:1:bin:/bin:/sbin/nologin
    daemon:x:2:2:daemon:/sbin:/sbin/nologin
    adm:x:3:4:adm:/var/adm:/sbin/nologin
    lp:x:4:7:lp:/var/spool/lpd:/sbin/nologin

    #-t定义了冒号为域字段之间的分隔符，-k 2指定基于第二个字段正向排序(字段顺序从1开始)。
    /> sort -t':' -k 1 users
    adm:x:3:4:adm:/var/adm:/sbin/nologin
    bin:x:1:1:bin:/bin:/sbin/nologin
    daemon:x:2:2:daemon:/sbin:/sbin/nologin
    lp:x:4:7:lp:/var/spool/lpd:/sbin/nologin
    root:x:0:0:root:/root:/bin/bash

    #还是以冒号为分隔符，这次是基于第三个域字段进行倒置排序。
    /> sort -t':' -k 3r users
    lp:x:4:7:lp:/var/spool/lpd:/sbin/nologin
    adm:x:3:4:adm:/var/adm:/sbin/nologin
    daemon:x:2:2:daemon:/sbin:/sbin/nologin
    bin:x:1:1:bin:/bin:/sbin/nologin
    root:x:0:0:root:/root:/bin/bash

    #先以第六个域的第2个字符到第4个字符进行正向排序，在基于第一个域进行反向排序。
    /> sort -t':' -k 6.2,6.4 -k 1r users
    bin:x:1:1:bin:/bin:/sbin/nologin
    root:x:0:0:root:/root:/bin/bash
    daemon:x:2:2:daemon:/sbin:/sbin/nologin
    lp:x:4:7:lp:/var/spool/lpd:/sbin/nologin
    adm:x:3:4:adm:/var/adm:/sbin/nologin

    #先以第六个域的第2个字符到第4个字符进行正向排序，在基于第一个域进行正向排序。和上一个例子比，第4和第5行交换了位置。
    /> sort -t':' -k 6.2,6.4 -k 1 users
    bin:x:1:1:bin:/bin:/sbin/nologin
    root:x:0:0:root:/root:/bin/bash
    daemon:x:2:2:daemon:/sbin:/sbin/nologin
    adm:x:3:4:adm:/var/adm:/sbin/nologin
    lp:x:4:7:lp:/var/spool/lpd:/sbin/nologin

    #基于第一个域的第2个字符排序
    /> sort -t':' -k 1.2,1.2 users    
    daemon:x:2:2:daemon:/sbin:/sbin/nologin
    adm:x:3:4:adm:/var/adm:/sbin/nologin
    bin:x:1:1:bin:/bin:/sbin/nologin
    root:x:0:0:root:/root:/bin/bash
    lp:x:4:7:lp:/var/spool/lpd:/sbin/nologin

    #基于第六个域的第2个字符到第4个字符进行正向排序，-u命令要求在排序时删除键值重复的行。
    /> sort -t':' -k 6.2,6.4 -u users
    bin:x:1:1:bin:/bin:/sbin/nologin
    root:x:0:0:root:/root:/bin/bash
    daemon:x:2:2:daemon:/sbin:/sbin/nologin
    adm:x:3:4:adm:/var/adm:/sbin/nologin

    /> cat /etc/passwd | wc -l  #计算该文件中文本的行数。
    39
    /> sed -n '35,$p' /etc/passwd > users2  #取最后5行并输出到users2中。
    /> cat users2
    sshd:x:74:74:Privilege-separated SSH:/var/empty/sshd:/sbin/nologin
    mysql:x:27:27:MySQL Server:/var/lib/mysql:/bin/bash
    pulse:x:496:494:PulseAudio System Daemon:/var/run/pulse:/sbin/nologin
    gdm:x:42:42::/var/lib/gdm:/sbin/nologin
    stephen:x:500:500:stephen:/home/stephen:/bin/bash

    #基于第3个域字段以文本的形式排序
    /> sort -t':' -k 3 users2
    mysql:x:27:27:MySQL Server:/var/lib/mysql:/bin/bash
    gdm:x:42:42::/var/lib/gdm:/sbin/nologin
    pulse:x:496:494:PulseAudio System Daemon:/var/run/pulse:/sbin/nologin
    stephen:x:500:500:stephen:/home/stephen:/bin/bash
    sshd:x:74:74:Privilege-separated SSH:/var/empty/sshd:/sbin/nologin

    #基于第3个域字段以数字的形式排序
    /> sort -t':' -k 3n users2
    mysql:x:27:27:MySQL Server:/var/lib/mysql:/bin/bash
    gdm:x:42:42::/var/lib/gdm:/sbin/nologin
    sshd:x:74:74:Privilege-separated SSH:/var/empty/sshd:/sbin/nologin
    pulse:x:496:494:PulseAudio System Daemon:/var/run/pulse:/sbin/nologin
    stephen:x:500:500:stephen:/home/stephen:/bin/bash

    #基于当前系统执行进程的owner名排序，并将排序的结果写入到result文件中
    /> ps -ef | sort -k 1 -o result

## 十三. 删除重复行的命令uniq:

    uniq有3个最为常用的选项，见如下列表：

选项	命令描述
-c	可在每个输出行之前加上该行重复的次数
-d	仅显示重复的行
-u	显示为重复的行
    /> cat testfile
    hello
    world
    friend
    hello
    world
    hello

    #直接删除未经排序的文件，将会发现没有任何行被删除
    /> uniq testfile  
    hello
    world
    friend
    hello
    world
    hello

    #排序之后删除了重复行，同时在行首位置输出该行重复的次数
    /> sort testfile | uniq -c  
    1 friend
    3 hello
    2 world

    #仅显示存在重复的行，并在行首显示该行重复的次数
    /> sort testfile | uniq -dc
    3 hello
    2 world

    #仅显示没有重复的行
    /> sort testfile | uniq -u
    friend  


## 十四. 文件压缩解压命令tar:

   1.  tar命令行选项

选项	命令描述
-c	建立压缩档案
-x	解压
--delete	从压缩包中删除已有文件，如果该文件在包中出现多次，该操作其将全部删除。
-t	查看压缩包中的文件列表
-r	向压缩归档文件末尾追加文件
-u	更新原压缩包中的文件
-z	压缩为gzip格式，或以gzip格式解压
-j	压缩为bzip2格式，或以bzip2格式解压
-v	显示压缩或解压的过程，该选项一般不适于后台操作
-f	使用档案名字，这个参数是最后一个参数，后面只能接档案名。

    2.  tar使用实例：
    #将当前目录下所有文件压缩打包，需要说明的是很多人都习惯将tar工具压缩的文件的扩展名命名为.tar
    /> tar -cvf test.tar *
    -rw-r--r--. 1 root root   183 Nov 11 08:02 users
    -rw-r--r--. 1 root root   279 Nov 11 08:45 users2

    /> cp ../*.log .                  #从上一层目录新copy一个.log文件到当前目录。
    /> tar -rvf test.tar *.log     #将扩展名为.log的文件追加到test.tar包里。
    /> tar -tvf test.tar
    -rw-r--r-- root/root        183 2011-11-11 08:02 users
    -rw-r--r-- root/root        279 2011-11-11 08:45 users2
    -rw-r--r-- root/root     48217 2011-11-11 22:16 install.log

    /> touch install.log           #使原有的文件更新一下最新修改时间
    /> tar -uvf test.tar *.log    #重新将更新后的log文件更新到test.tar中
    /> tar -tvf test.tar             #从输出结果可以看出tar包中多出一个更新后install.log文件。
    -rw-r--r-- root/root         183 2011-11-11 08:02 users
    -rw-r--r-- root/root         279 2011-11-11 08:45 users2
    -rw-r--r-- root/root     48217 2011-11-11 22:16 install.log
    -rw-r--r-- root/root     48217 2011-11-11 22:20 install.log

    /> tar --delete install.log -f test.tar #基于上面的结果，从压缩包中删除install.log
    -rw-r--r-- root/root       183 2011-11-11 08:02 users
    -rw-r--r-- root/root       279 2011-11-11 08:45 users2

    /> rm -f users users2      #从当前目录将tar中的两个文件删除
    /> tar -xvf test.tar          #解压
    /> ls -l users*                 #仅列出users和users2的详细列表信息
    -rw-r--r--. 1 root root 183 Nov 11 08:02 users
    -rw-r--r--. 1 root root 279 Nov 11 08:45 users2

    #以gzip的格式压缩并打包，解压时也应该以同样的格式解压，需要说明的是以该格式压缩的包习惯在扩展名后加.gz
    /> tar -cvzf test.tar.gz *
    /> tar -tzvf test.tar.gz      #查看压缩包中文件列表时也要加z选项(gzip格式)
    -rw-r--r-- root/root     48217 2011-11-11 22:50 install.log
    -rw-r--r-- root/root         183 2011-11-11 08:02 users
    -rw-r--r-- root/root         279 2011-11-11 08:45 users2

    /> rm -f users users2 install.log
    /> tar -xzvf test.tar.gz     #以gzip的格式解压
    /> ls -l *.log users*
    -rw-r--r-- root/root     48217 2011-11-11 22:50 install.log
    -rw-r--r-- root/root         183 2011-11-11 08:02 users
    -rw-r--r-- root/root         279 2011-11-11 08:45 users2

    /> rm -f test.*                #删除当前目录下原有的压缩包文件
    #以bzip2的格式压缩并打包，解压时也应该以同样的格式解压，需要说明的是以该格式压缩的包习惯在扩展名后加.bz2
    /> tar -cvjf test.tar.bz2 *
    /> tar -tjvf test.tar.bz2    #查看压缩包中文件列表时也要加j选项(bzip2格式)
    -rw-r--r-- root/root     48217 2011-11-11 22:50 install.log
    -rw-r--r-- root/root         183 2011-11-11 08:02 users
    -rw-r--r-- root/root         279 2011-11-11 08:45 users2

    /> rm -f *.log user*
    /> tar -xjvf test.tar.bz2    #以bzip2的格式解压
    /> ls -l
    -rw-r--r--. 1 root root 48217 Nov 11 22:50 install.log
    -rw-r--r--. 1 root root 10530 Nov 11 23:08 test.tar.bz2
    -rw-r--r--. 1 root root     183 Nov 11 08:02 users
    -rw-r--r--. 1 root root     279 Nov 11 08:45 users2

## 十五. 大文件拆分命令split:

    下面的列表中给出了该命令最为常用的几个命令行选项：

选项	描述
-l	指定行数，每多少分隔成一个文件，缺省值为1000行。
-b	指定字节数，支持的单位为：k和m
-C	与-b参数类似，但切割时尽量维持每行的完整性
-d	生成文件的后缀为数字，如果不指定该选项，缺省为字母
    /> ls -l
    -rw-r--r--. 1 root root 10530 Nov 11 23:08 test.tar.bz2

    /> split -b 5k test.tar.bz2     #以每文件5k的大小切割test.tar.bz2
    /> ls -l                                #查看切割后的结果，缺省情况下拆分后的文件名为以下形式。
    -rw-r--r--. 1 root root 10530 Nov 11 23:08 test.tar.bz2
    -rw-r--r--. 1 root root   5120 Nov 11 23:34 xaa
    -rw-r--r--. 1 root root   5120 Nov 11 23:34 xab
    -rw-r--r--. 1 root root     290 Nov 11 23:34 xac

    /> rm -f x*                         #删除拆分后的小文件
    /> split -d -b 5k test.tar.bz2 #-d选项以后缀为数字的形式命名拆分后的小文件
    /> ls -l
    -rw-r--r--. 1 root root 10530 Nov 11 23:08 test.tar.bz2
    -rw-r--r--. 1 root root   5120 Nov 11 23:36 x00
    -rw-r--r--. 1 root root   5120 Nov 11 23:36 x01
    -rw-r--r--. 1 root root     290 Nov 11 23:36 x02

    /> wc install.log -l             #计算该文件的行数
    /> split -l 300 install.log     #每300行拆分成一个小文件
    /> ls -l x*
    -rw-r--r--. 1 root root 11184 Nov 11 23:42 xaa
    -rw-r--r--. 1 root root 10805 Nov 11 23:42 xab
    -rw-r--r--. 1 root root 12340 Nov 11 23:42 xac
    -rw-r--r--. 1 root root 11783 Nov 11 23:42 xad
    -rw-r--r--. 1 root root   2105 Nov 11 23:42 xae

## 十六. 文件查找命令find:

    下面给出find命令的主要应用示例：
    /> ls -l     #列出当前目录下所包含的测试文件
    -rw-r--r--. 1 root root 48217 Nov 12 00:57 install.log
    -rw-r--r--. 1 root root      37 Nov 12 00:56 testfile.dat
    -rw-r--r--. 1 root root 10530 Nov 11 23:08 test.tar.bz2
    -rw-r--r--. 1 root root     183 Nov 11 08:02 users
    -rw-r--r--. 1 root root     279 Nov 11 08:45 users2
    
    1. 按文件名查找：
    -name:  查找时文件名大小写敏感。
    -iname: 查找时文件名大小写不敏感。
    #该命令为find命令中最为常用的命令，即从当前目录中查找扩展名为.log的文件。需要说明的是，缺省情况下，find会从指定的目录搜索，并递归的搜索其子目录。
    /> find . -name "*.log"
     ./install.log
    /> find . -iname U*          #如果执行find . -name U*将不会找到匹配的文件
    users users2


    2. 按文件时间属性查找：
    -atime  -n[+n]: 找出文件访问时间在n日之内[之外]的文件。
    -ctime  -n[+n]: 找出文件更改时间在n日之内[之外]的文件。
    -mtime -n[+n]: 找出修改数据时间在n日之内[之外]的文件。
    -amin   -n[+n]: 找出文件访问时间在n分钟之内[之外]的文件。
    -cmin   -n[+n]: 找出文件更改时间在n分钟之内[之外]的文件。
    -mmin  -n[+n]: 找出修改数据时间在n分钟之内[之外]的文件。
    /> find -ctime -2        #找出距此时2天之内创建的文件
    .
    ./users2
    ./install.log
    ./testfile.dat
    ./users
    ./test.tar.bz2
    /> find -ctime +2        #找出距此时2天之前创建的文件
    没有找到                     #因为当前目录下所有文件都是2天之内创建的
    /> touch install.log     #手工更新install.log的最后访问时间，以便下面的find命令可以找出该文件
    /> find . -cmin  -3       #找出修改状态时间在3分钟之内的文件。
    install.log

    3. 基于找到的文件执行指定的操作：
    -exec: 对匹配的文件执行该参数所给出的shell命令。相应命令的形式为'command' {} \;，注意{}和\；之间的空格，同时两个{}之间没有空格
    -ok:   其主要功能和语法格式与-exec完全相同，唯一的差别是在于该选项更加安全，因为它会在每次执行shell命令之前均予以提示，只有在回答为y的时候，其后的shell命令才会被继续执行。需要说明的是，该选项不适用于自动化脚本，因为该提供可能会挂起整个自动化流程。
    #找出距此时2天之内创建的文件，同时基于find的结果，应用-exec之后的命令，即ls -l，从而可以直接显示出find找到文件的明显列表。
    /> find . -ctime -2 -exec ls -l {} \;
    -rw-r--r--. 1 root root      279 Nov 11 08:45 ./users2
    -rw-r--r--. 1 root root  48217 Nov 12 00:57 ./install.log
    -rw-r--r--. 1 root root        37 Nov 12 00:56 ./testfile.dat
    -rw-r--r--. 1 root root      183 Nov 11 08:02 ./users
    -rw-r--r--. 1 root root  10530 Nov 11 23:08 ./test.tar.bz2
    #找到文件名为*.log, 同时文件数据修改时间距此时为1天之内的文件。如果找到就删除他们。有的时候，这样的写法由于是在找到之后立刻删除，因此存在一定误删除的危险。
    /> ls
    install.log  testfile.dat  test.tar.bz2  users  users2
    /> find . -name "*.log" -mtime -1 -exec rm -f {} \; 
    /> ls
    testfile.dat  test.tar.bz2  users  users2
    在控制台下，为了使上面的命令更加安全，我们可以使用-ok替换-exec，见如下示例：
    />  find . -name "*.dat" -mtime -1 -ok rm -f {} \;
    < rm ... ./testfile.dat > ? y    #对于该提示，如果回答y，找到的*.dat文件将被删除，这一点从下面的ls命令的结果可以看出。
    /> ls
    test.tar.bz2  users  users2

    4. 按文件所属的owner和group查找：
    -user:      查找owner属于-user选项后面指定用户的文件。
    ! -user:    查找owner不属于-user选项后面指定用户的文件。
    -group:   查找group属于-group选项后面指定组的文件。
    ! -group: 查找group不属于-group选项后面指定组的文件。
    /> ls -l                            #下面三个文件的owner均为root
    -rw-r--r--. 1 root root 10530 Nov 11 23:08 test.tar.bz2
    -rw-r--r--. 1 root root     183 Nov 11 08:02 users
    -rw-r--r--. 1 root root     279 Nov 11 08:45 users2
    /> chown stephen users   #将users文件的owner从root改为stephen。
    /> ls -l
    -rw-r--r--. 1 root       root 10530 Nov 11 23:08 test.tar.bz2
    -rw-r--r--. 1 stephen root    183 Nov 11 08:02 users
    -rw-r--r--. 1 root       root     279 Nov 11 08:45 users2
    /> find . -user root          #搜索owner是root的文件
    .
    ./users2
    ./test.tar.bz2
    /> find . ! -user root        #搜索owner不是root的文件，注意!和-user之间要有空格。
    ./users
    /> ls -l                            #下面三个文件的所属组均为root
    -rw-r--r--. 1 root      root 10530 Nov 11 23:08 test.tar.bz2
    -rw-r--r--. 1 stephen root    183 Nov 11 08:02 users
    -rw-r--r--. 1 root      root    279 Nov 11 08:45 users2
    /> chgrp stephen users    #将users文件的所属组从root改为stephen
    /> ls -l
    -rw-r--r--. 1 root           root    10530 Nov 11 23:08 test.tar.bz2
    -rw-r--r--. 1 stephen stephen      183 Nov 11 08:02 users
    -rw-r--r--. 1 root            root       279 Nov 11 08:45 users2
    /> find . -group root        #搜索所属组是root的文件
    .
    ./users2
    ./test.tar.bz2
    /> find . ! -group root      #搜索所属组不是root的文件，注意!和-user之间要有空格。    
    ./users

    5. 按指定目录深度查找：
    -maxdepth: 后面的参数表示距当前目录指定的深度，其中1表示当前目录，2表示一级子目录，以此类推。在指定该选项后，find只是在找到指定深度后就不在递归其子目录了。下例中的深度为1，表示只是在当前子目录中搜索。如果没有设置该选项，find将递归当前目录下的所有子目录。
    /> mkdir subdir               #创建一个子目录，并在该子目录内创建一个文件
    /> cd subdir
    /> touch testfile
    /> cd ..
    #maxdepth后面的参数表示距当前目录指定的深度，其中1表示当前目录，2表示一级子目录，以此类推。在指定该选项后，find只是在找到指定深度后就不在递归其子目录了。下例中的深度为1，表示只是在当前子目录中搜索。如果没有设置该选项，find将递归当前目录下的所有子目录。
    /> find . -maxdepth 1 -name "*"
    .
    ./users2
    ./subdir
    ./users
    ./test.tar.bz2
    #搜索深度为子一级子目录，这里可以看出子目录下刚刚创建的testfile已经被找到
    /> find . -maxdepth 2 -name "*"  
    .
    ./users2
    ./subdir
    ./subdir/testfile
    ./users
    ./test.tar.bz2
    
    6. 排除指定子目录查找：
    -path pathname -prune:   避开指定子目录pathname查找。
    -path expression -prune:  避开表达中指定的一组pathname查找。
    需要说明的是，如果同时使用-depth选项，那么-prune将被find命令忽略。
    #为后面的示例创建需要避开的和不需要避开的子目录，并在这些子目录内均创建符合查找规则的文件。
    /> mkdir DontSearchPath  
    /> cd DontSearchPath
    /> touch datafile1
    /> cd ..
    /> mkdir DoSearchPath
    /> cd DoSearchPath
    /> touch datafile2
    /> cd ..
    /> touch datafile3
    #当前目录下，避开DontSearchPath子目录，搜索所有文件名为datafile*的文件。
    /> find . -path "./DontSearchPath" -prune -o -name "datafile*" -print
    ./DoSearchPath/datafile2
    ./datafile3
    #当前目录下，同时避开DontSearchPath和DoSearchPath两个子目录，搜索所有文件名为datafile*的文件。
    /> find . \( -path "./DontSearchPath" -o -path "./DoSearchPath" \) -prune -o -name "datafile*" -print
    ./datafile3
        
    7. 按文件权限属性查找：
    -perm mode:   文件权限正好符合mode(mode为文件权限的八进制表示)。
    -perm +mode: 文件权限部分符合mode。如命令参数为644(-rw-r--r--)，那么只要文件权限属性中有任何权限和644重叠，这样的文件均可以被选出。
    -perm -mode:  文件权限完全符合mode。如命令参数为644(-rw-r--r--)，当644中指定的权限已经被当前文件完全拥有，同时该文件还拥有额外的权限属性，这样的文件可被选出。
    /> ls -l
    -rw-r--r--. 1 root            root           0 Nov 12 10:02 datafile3
    -rw-r--r--. 1 root            root    10530 Nov 11 23:08 test.tar.bz2
    -rw-r--r--. 1 stephen stephen        183 Nov 11 08:02 users
    -rw-r--r--. 1 root            root        279 Nov 11 08:45 users2
    /> find . -perm 644      #查找所有文件权限正好为644(-rw-r--r--)的文件。
    ./users2
    ./datafile3
    ./users
    ./test.tar.bz2
    /> find . -perm 444      #当前目录下没有文件的权限属于等于444(均为644)。    
    /> find . -perm -444     #644所包含的权限完全覆盖444所表示的权限。
    .
    ./users2
    ./datafile3
    ./users
    ./test.tar.bz2
    /> find . -perm +111    #查找所有可执行的文件，该命令没有找到任何文件。
    /> chmod u+x users     #改变users文件的权限，添加owner的可执行权限，以便于下面的命令可以将其找出。
    /> find . -perm +111     
    .
    ./users    
    
    8. 按文件类型查找：
    -type：后面指定文件的类型。
    b - 块设备文件。
    d - 目录。
    c - 字符设备文件。
    p - 管道文件。
    l  - 符号链接文件。
    f  - 普通文件。
    /> mkdir subdir 
    /> find . -type d      #在当前目录下，找出文件类型为目录的文件。
    ./subdir
　 /> find . ! -type d    #在当前目录下，找出文件类型不为目录的文件。
    ./users2
    ./datafile3
    ./users
    ./test.tar.bz2
    /> find . -type f       #在当前目录下，找出文件类型为文件的文件
    ./users2
    ./datafile3
    ./users
    ./test.tar.bz2
    
    9. 按文件大小查找：
    -size [+/-]100[c/k/M/G]: 表示文件的长度为等于[大于/小于]100块[字节/k/M/G]的文件。
    -empty: 查找空文件。
    /> find . -size +4k -exec ls -l {} \;  #查找文件大小大于4k的文件，同时打印出找到文件的明细
    -rw-r--r--. 1 root root 10530 Nov 11 23:08 ./test.tar.bz2
    /> find . -size -4k -exec ls -l {} \;  #查找文件大小小于4k的文件。
    -rw-r--r--. 1 root            root 279 Nov 11 08:45 ./users2
    -rw-r--r--. 1 root             root    0 Nov 12 10:02 ./datafile3
    -rwxr--r--. 1 stephen stephen 183 Nov 11 08:02 ./users
    /> find . -size 183c -exec ls -l {} \; #查找文件大小等于183字节的文件。
    -rwxr--r--. 1 stephen stephen 183 Nov 11 08:02 ./users
    /> find . -empty  -type f -exec ls -l {} \;
    -rw-r--r--. 1 root root 0 Nov 12 10:02 ./datafile3
    
    10. 按更改时间比指定文件新或比文件旧的方式查找：
    -newer file1 ! file2： 查找文件的更改日期比file1新，但是比file2老的文件。
    /> ls -lrt   #以时间顺序(从早到晚)列出当前目录下所有文件的明细列表，以供后面的例子参考。
    -rwxr--r--. 1 stephen stephen   183 Nov 11 08:02 users1
    -rw-r--r--. 1 root           root    279 Nov 11 08:45 users2
    -rw-r--r--. 1 root           root 10530 Nov 11 23:08 test.tar.bz2
    -rw-r--r--. 1 root           root        0 Nov 12 10:02 datafile3
    /> find . -newer users1     #查找文件更改日期比users1新的文件，从上面结果可以看出，其余文件均符合要求。
    ./users2
    ./datafile3
    ./test.tar.bz2
    /> find . ! -newer users2   #查找文件更改日期不比users1新的文件。
    ./users2
    ./users
    #查找文件更改日期比users2新，但是不比test.tar.bz2新的文件。
    /> find . -newer users2 ! -newer test.tar.bz2
    ./test.tar.bz2
    
    细心的读者可能发现，关于find的说明，在我之前的Blog中已经给出，这里之所以拿出一个小节再次讲述该命令主要是因为以下三点原因：
    1. find命令在Linux Shell中扮演着极为重要的角色；
    2. 为了保证本系列的完整性；
    3. 之前的Blog是我多年之前留下的总结笔记，多少有些粗糙，这里给出了更为详细的举例。



## 十七. xargs命令:

    该命令的主要功能是从输入中构建和执行shell命令。       
    在使用find命令的-exec选项处理匹配到的文件时， find命令将所有匹配到的文件一起传递给exec执行。但有些系统对能够传递给exec的命令长度有限制，这样在find命令运行几分钟之后，就会出现溢出错误。错误信息通常是“参数列太长”或“参数列溢出”。这就是xargs命令的用处所在，特别是与find命令一起使用。  
    find命令把匹配到的文件传递给xargs命令，而xargs命令每次只获取一部分文件而不是全部，不像-exec选项那样。这样它可以先处理最先获取的一部分文件，然后是下一批，并如此继续下去。  
    在有些系统中，使用-exec选项会为处理每一个匹配到的文件而发起一个相应的进程，并非将匹配到的文件全部作为参数一次执行；这样在有些情况下就会出现进程过多，系统性能下降的问题，因而效率不高；  
    而使用xargs命令则只有一个进程。另外，在使用xargs命令时，究竟是一次获取所有的参数，还是分批取得参数，以及每一次获取参数的数目都会根据该命令的选项及系统内核中相应的可调参数来确定。
    /> ls -l
    -rw-r--r--. 1 root root        0 Nov 12 10:02 datafile3
    -rw-r--r--. 1 root root 10530 Nov 11 23:08 test.tar.bz2
    -rwxr--r--. 1 root root    183 Nov 11 08:02 users
    -rw-r--r--. 1 root root    279 Nov 11 08:45 users2
    #查找当前目录下的每一个普通文件，然后使用xargs命令来测试它们分别属于哪类文件。
    /> find . -type f -print | xargs file 
    ./users2:        ASCII text
    ./datafile3:      empty
    ./users:          ASCII text
    ./test.tar.bz2: bzip2 compressed data, block size = 900k
    #回收当前目录下所有普通文件的执行权限。
    /> find . -type f -print | xargs chmod a-x
    /> ls -l
    -rw-r--r--. 1 root root     0 Nov 12 10:02 datafile3
    -rw-r--r--. 1 root root 10530 Nov 11 23:08 test.tar.bz2
    -rw-r--r--. 1 root root   183 Nov 11 08:02 users
    -rw-r--r--. 1 root root   279 Nov 11 08:45 users2
    #在当面目录下查找所有普通文件，并用grep命令在搜索到的文件中查找hostname这个词
    /> find . -type f -print | xargs grep "hostname" 
    #在整个系统中查找内存信息转储文件(core dump) ，然后把结果保存到/tmp/core.log 文件中。
    /> find / -name "core" -print | xargs echo "" >/tmp/core.log 　

    /> pgrep mysql | xargs kill -9　　#直接杀掉mysql的进程
    [1]+  Killed                  mysql

## 十八.  和系统运行状况相关的Shell命令:

    1.  Linux的实时监测命令(watch):
    watch 是一个非常实用的命令，可以帮你实时监测一个命令的运行结果，省得一遍又一遍的手动运行。该命令最为常用的两个选项是-d和-n，其中-n表示间隔多少秒执行一次"command"，-d表示高亮发生变化的位置。下面列举几个在watch中常用的实时监视命令：
    /> watch -d -n 1 'who'   #每隔一秒执行一次who命令，以监视服务器当前用户登录的状况
    Every 1.0s: who       Sat Nov 12 12:37:18 2011
    
    stephen  tty1           2011-11-11 17:38 (:0)
    stephen  pts/0         2011-11-11 17:39 (:0.0)
    root       pts/1         2011-11-12 10:01 (192.168.149.1)
    root       pts/2         2011-11-12 11:41 (192.168.149.1)
    root       pts/3         2011-11-12 12:11 (192.168.149.1)
    stephen  pts/4         2011-11-12 12:22 (:0.0)
    此时通过其他Linux客户端工具以root的身份登录当前Linux服务器，再观察watch命令的运行变化。
    Every 1.0s: who       Sat Nov 12 12:41:09 2011
    
    stephen  tty1          2011-11-11 17:38 (:0)
    stephen  pts/0        2011-11-11 17:39 (:0.0)
    root       pts/1        2011-11-12 10:01 (192.168.149.1)
    root       pts/2        2011-11-12 11:41 (192.168.149.1)
    root       pts/3        2011-11-12 12:40 (192.168.149.1)
    stephen  pts/4        2011-11-12 12:22 (:0.0)
    root       pts/5        2011-11-12 12:41 (192.168.149.1)
    最后一行中被高亮的用户为新登录的root用户。此时按CTRL + C可以退出正在执行的watch监控进程。
    
    #watch可以同时运行多个命令，命令间用分号分隔。
    #以下命令监控磁盘的使用状况，以及当前目录下文件的变化状况，包括文件的新增、删除和文件修改日期的更新等。
    /> watch -d -n 1 'df -h; ls -l'
    Every 1.0s: df -h; ls -l     Sat Nov 12 12:55:00 2011
    
    Filesystem            Size  Used Avail Use% Mounted on
    /dev/sda1             5.8G  3.3G  2.2G  61% /
    tmpfs                 504M  420K  504M   1% /dev/shm
    total 20
    -rw-r--r--. 1 root root 10530 Nov 11 23:08 test.tar.bz2
    -rw-r--r--. 1 root root   183 Nov 11 08:02 users
    -rw-r--r--. 1 root root   279 Nov 11 08:45 users2
    此时通过另一个Linux控制台窗口，在watch监视的目录下，如/home/stephen/test，执行下面的命令
    /> touch aa         #在执行该命令之后，另一个执行watch命令的控制台将有如下变化
    Every 1.0s: df -h; ls -l                                Sat Nov 12 12:57:08 2011
    
    Filesystem            Size  Used Avail Use% Mounted on
    /dev/sda1             5.8G  3.3G  2.2G  61% /
    tmpfs                 504M  420K  504M   1% /dev/shm
    total 20
    -rw-r--r--. 1 root root        0 Nov 12 12:56 aa
    -rw-r--r--. 1 root root        0 Nov 12 10:02 datafile3
    -rw-r--r--. 1 root root 10530 Nov 11 23:08 test.tar.bz2
    -rw-r--r--. 1 root root     183 Nov 11 08:02 users
    -rw-r--r--. 1 root root     279 Nov 11 08:45 users2
    其中黄色高亮的部分，为touch aa命令执行之后watch输出的高亮变化部分。

    
    2.  查看当前系统内存使用状况(free)：
    free命令有以下几个常用选项：

选项	说明
-b	以字节为单位显示数据。
-k	以千字节(KB)为单位显示数据(缺省值)。
-m	以兆(MB)为单位显示数据。
-s delay	该选项将使free持续不断的刷新，每次刷新之间的间隔为delay指定的秒数，如果含有小数点，将精确到毫秒，如0.5为500毫秒，1为一秒。
    free命令输出的表格中包含以下几列：

列名	说明
total	总计物理内存的大小。
used	已使用的内存数量。
free	可用的内存数量。
Shared	多个进程共享的内存总额。
Buffers/cached	磁盘缓存的大小。

    见以下具体示例和输出说明：
    /> free -k
                        total         used          free     shared    buffers     cached
    Mem:       1031320     671776     359544          0      88796     352564
    -/+ buffers/cache:      230416     800904
    Swap:        204792              0     204792
    对于free命令的输出，我们只需关注红色高亮的输出行和绿色高亮的输出行，见如下具体解释：
    红色输出行：该行使从操作系统的角度来看待输出数据的，used(671776)表示内核(Kernel)+Applications+buffers+cached。free(359544)表示系统还有多少内存可供使用。
    绿色输出行：该行则是从应用程序的角度来看输出数据的。其free = 操作系统used + buffers + cached，既：
    800904 = 359544 + 88796 + 352564
    /> free -m
                      total        used        free      shared    buffers     cached
    Mem:          1007         656        351            0         86            344
    -/+ buffers/cache:        225        782
    Swap:          199             0        199
    /> free -k -s 1.5  #以千字节(KB)为单位显示数据，同时每隔1.5刷新输出一次，直到按CTRL+C退出
                      total        used       free     shared    buffers     cached
    Mem:          1007         655        351          0           86        344
    -/+ buffers/cache:        224        782
    Swap:          199             0        199

                      total        used       free     shared    buffers     cached
    Mem:          1007         655        351          0           86        344
    -/+ buffers/cache:        224        782
    Swap:          199             0        199

    3.  CPU的实时监控工具(mpstat)：
    该命令主要用于报告当前系统中所有CPU的实时运行状况。
    #该命令将每隔2秒输出一次CPU的当前运行状况信息，一共输出5次，如果没有第二个数字参数，mpstat将每隔两秒执行一次，直到按CTRL+C退出。
    /> mpstat 2 5  
    Linux 2.6.32-71.el6.i686 (Stephen-PC)   11/12/2011      _i686_  (1 CPU)

    04:03:00 PM  CPU    %usr   %nice    %sys %iowait    %irq   %soft  %steal  %guest   %idle
    04:03:02 PM  all    0.00    0.00    0.50    0.00    0.00    0.00    0.00    0.00   99.50
    04:03:04 PM  all    0.00    0.00    0.00    0.00    0.00    0.00    0.00    0.00  100.00
    04:03:06 PM  all    0.00    0.00    0.00    0.00    0.00    0.00    0.00    0.00  100.00
    04:03:08 PM  all    0.00    0.00    0.00    0.00    0.00    0.00    0.00    0.00  100.00
    04:03:10 PM  all    0.00    0.00    0.00    0.00    0.00    0.00    0.00    0.00  100.00
    Average:       all    0.00    0.00    0.10    0.00    0.00    0.00    0.00    0.00   99.90

    第一行的末尾给出了当前系统中CPU的数量。后面的表格中则输出了系统当前CPU的使用状况，以下为每列的含义：

列名	说明
%user	在internal时间段里，用户态的CPU时间(%)，不包含nice值为负进程  (usr/total)*100
%nice	在internal时间段里，nice值为负进程的CPU时间(%)   (nice/total)*100
%sys	在internal时间段里，内核时间(%)       (system/total)*100
%iowait	在internal时间段里，硬盘IO等待时间(%) (iowait/total)*100
%irq	在internal时间段里，硬中断时间(%)     (irq/total)*100
%soft	在internal时间段里，软中断时间(%)     (softirq/total)*100
%idle	在internal时间段里，CPU除去等待磁盘IO操作外的因为任何原因而空闲的时间闲置时间(%) (idle/total)*100
    计算公式：
    total_cur=user+system+nice+idle+iowait+irq+softirq
    total_pre=pre_user+ pre_system+ pre_nice+ pre_idle+ pre_iowait+ pre_irq+ pre_softirq
    user=user_cur – user_pre
    total=total_cur-total_pre
    其中_cur 表示当前值，_pre表示interval时间前的值。上表中的所有值可取到两位小数点。    

    /> mpstat -P ALL 2 3  #-P ALL表示打印所有CPU的数据，这里也可以打印指定编号的CPU数据，如-P 0(CPU的编号是0开始的)
    Linux 2.6.32-71.el6.i686 (Stephen-PC)   11/12/2011      _i686_  (1 CPU)

    04:12:54 PM  CPU    %usr   %nice    %sys %iowait    %irq   %soft  %steal  %guest   %idle
    04:12:56 PM    all      0.00      0.00     0.50    0.00      0.00    0.00    0.00      0.00     99.50
    04:12:56 PM      0     0.00      0.00     0.50    0.00      0.00    0.00    0.00      0.00     99.50

    04:12:56 PM  CPU    %usr   %nice    %sys %iowait    %irq   %soft  %steal  %guest   %idle
    04:12:58 PM    all     0.00      0.00     0.00    0.00      0.00    0.00    0.00      0.00    100.00
    04:12:58 PM     0     0.00      0.00     0.00    0.00      0.00    0.00    0.00      0.00    100.00

    04:12:58 PM  CPU    %usr   %nice    %sys %iowait    %irq   %soft  %steal  %guest   %idle
    04:13:00 PM    all      0.00     0.00    0.00    0.00      0.00    0.00     0.00      0.00    100.00
    04:13:00 PM     0      0.00     0.00    0.00    0.00      0.00    0.00     0.00      0.00    100.00

    Average:       CPU    %usr   %nice    %sys %iowait    %irq   %soft  %steal  %guest   %idle
    Average:         all      0.00     0.00    0.17    0.00      0.00    0.00     0.00      0.00     99.83
    Average:          0      0.00     0.00    0.17    0.00      0.00    0.00     0.00      0.00     99.83

    4.  虚拟内存的实时监控工具(vmstat)：
    vmstat命令用来获得UNIX系统有关进程、虚存、页面交换空间及CPU活动的信息。这些信息反映了系统的负载情况。vmstat首次运行时显示自系统启动开始的各项统计信息，之后运行vmstat将显示自上次运行该命令以后的统计信息。用户可以通过指定统计的次数和时间来获得所需的统计信息。
    /> vmstat 1 3    #这是vmstat最为常用的方式，其含义为每隔1秒输出一条，一共输出3条后程序退出。
    procs  -----------memory----------   ---swap-- -----io---- --system-- -----cpu-----
     r  b   swpd      free      buff   cache   si   so     bi    bo     in   cs  us  sy id  wa st
     0  0        0 531760  67284 231212  108  0     0  260   111  148  1   5 86   8  0
     0  0        0 531752  67284 231212    0    0     0     0     33   57   0   1 99   0  0
     0  0        0 531752  67284 231212    0    0     0     0     40   73   0   0 100 0  0

    /> vmstat 1       #其含义为每隔1秒输出一条，直到按CTRL+C后退出。

    下面将给出输出表格中每一列的含义说明：
    有关进程的信息有：(procs)
    r:  在就绪状态等待的进程数。
    b: 在等待状态等待的进程数。    
    有关内存的信息有：(memory)
    swpd:  正在使用的swap大小，单位为KB。
    free:    空闲的内存空间。
    buff:    已使用的buff大小，对块设备的读写进行缓冲。
    cache: 已使用的cache大小，文件系统的cache。
    有关页面交换空间的信息有：(swap)
    si:  交换内存使用，由磁盘调入内存。
    so: 交换内存使用，由内存调入磁盘。  
    有关IO块设备的信息有：(io)
    bi:  从块设备读入的数据总量(读磁盘) (KB/s)
    bo: 写入到块设备的数据总理(写磁盘) (KB/s)   
    有关故障的信息有：(system)
    in: 在指定时间内的每秒中断次数。
    sy: 在指定时间内每秒系统调用次数。
    cs: 在指定时间内每秒上下文切换的次数。   
    有关CPU的信息有：(cpu)
    us:  在指定时间间隔内CPU在用户态的利用率。
    sy:  在指定时间间隔内CPU在核心态的利用率。
    id:  在指定时间间隔内CPU空闲时间比。
    wa: 在指定时间间隔内CPU因为等待I/O而空闲的时间比。   
    vmstat 可以用来确定一个系统的工作是受限于CPU还是受限于内存：如果CPU的sy和us值相加的百分比接近100%，或者运行队列(r)中等待的进程数总是不等于0，且经常大于4，同时id也经常小于40，则该系统受限于CPU；如果bi、bo的值总是不等于0，则该系统受限于内存。

    5.  设备IO负载的实时监控工具(iostat)：
    iostat主要用于监控系统设备的IO负载情况，iostat首次运行时显示自系统启动开始的各项统计信息，之后运行iostat将显示自上次运行该命令以后的统计信息。用户可以通过指定统计的次数和时间来获得所需的统计信息。
    其中该命令中最为常用的使用方式如下：
    /> iostat -d 1 3    #仅显示设备的IO负载，其中每隔1秒刷新并输出结果一次，输出3次后iostat退出。
    Linux 2.6.32-71.el6.i686 (Stephen-PC)   11/16/2011      _i686_  (1 CPU)

    Device:            tps   Blk_read/s   Blk_wrtn/s   Blk_read   Blk_wrtn
    sda                 5.35       258.39        26.19     538210      54560

    Device:            tps   Blk_read/s   Blk_wrtn/s   Blk_read   Blk_wrtn
    sda                 0.00         0.00         0.00                  0          0

    Device:            tps   Blk_read/s   Blk_wrtn/s   Blk_read   Blk_wrtn
    sda                 0.00         0.00         0.00                  0          0

    Device:            tps   Blk_read/s   Blk_wrtn/s   Blk_read   Blk_wrtn
    sda                 0.00         0.00         0.00                  0          0
    /> iostat -d 1  #和上面的命令一样，也是每隔1秒刷新并输出一次，但是该命令将一直输出，直到按CTRL+C退出。
    下面将给出输出表格中每列的含义：

列名	说明
Blk_read/s	每秒块(扇区)读取的数量。
Blk_wrtn/s	每秒块(扇区)写入的数量。
Blk_read	总共块(扇区)读取的数量。
Blk_wrtn	总共块(扇区)写入的数量。
    iostat还有一个比较常用的选项-x，该选项将用于显示和io相关的扩展数据。
    /> iostat -dx 1 3
    Device:  rrqm/s wrqm/s  r/s   w/s  rsec/s wsec/s avgrq-sz avgqu-sz   await  svctm  %util
    sda            5.27   1.31 2.82 1.14 189.49  19.50    52.75     0.53     133.04  10.74   4.26

    Device:  rrqm/s wrqm/s  r/s   w/s  rsec/s wsec/s avgrq-sz avgqu-sz   await  svctm  %util
    sda            0.00   0.00 0.00 0.00   0.00   0.00        0.00     0.00         0.00   0.00   0.00

    Device:  rrqm/s wrqm/s  r/s   w/s  rsec/s wsec/s avgrq-sz avgqu-sz   await  svctm  %util
    sda            0.00   0.00 0.00 0.00   0.00   0.00        0.00     0.00         0.00   0.00   0.00
    还可以在命令行参数中指定要监控的设备名，如：
    /> iostat -dx sda 1 3   #指定监控的设备名称为sda，该命令的输出结果和上面命令完全相同。

    下面给出扩展选项输出的表格中每列的含义：

列名	说明
rrqm/s	队列中每秒钟合并的读请求数量
wrqm/s	队列中每秒钟合并的写请求数量
r/s	每秒钟完成的读请求数量
w/s	每秒钟完成的写请求数量
rsec/s	每秒钟读取的扇区数量
wsec/s	每秒钟写入的扇区数量
avgrq-sz	平均请求扇区的大小
avgqu-sz	平均请求队列的长度
await	平均每次请求的等待时间
util	设备的利用率
    下面是关键列的解释：
    util是设备的利用率。如果它接近100%，通常说明设备能力趋于饱和。
    await是平均每次请求的等待时间。这个时间包括了队列时间和服务时间，也就是说，一般情况下，await大于svctm，它们的差值越小，则说明队列时间越短，反之差值越大，队列时间越长，说明系统出了问题。
    avgqu-sz是平均请求队列的长度。毫无疑问，队列长度越短越好。                 

     6.  当前运行进程的实时监控工具(pidstat)：
     pidstat主要用于监控全部或指定进程占用系统资源的情况，如CPU，内存、设备IO、任务切换、线程等。pidstat首次运行时显示自系统启动开始的各项统计信息，之后运行pidstat将显示自上次运行该命令以后的统计信息。用户可以通过指定统计的次数和时间来获得所需的统计信息。
    在正常的使用，通常都是通过在命令行选项中指定待监控的pid，之后在通过其他具体的参数来监控与该pid相关系统资源信息。

选项	说明
-l	显示该进程和CPU相关的信息(command列中可以显示命令的完整路径名和命令的参数)。
-d	显示该进程和设备IO相关的信息。
-r	显示该进程和内存相关的信息。
-w	显示该进程和任务时间片切换相关的信息。
-t	显示在该进程内正在运行的线程相关的信息。
-p	后面紧跟着带监控的进程id或ALL(表示所有进程)，如不指定该选项，将监控当前系统正在运行的所有进程。
    #监控pid为1(init)的进程的CPU资源使用情况，其中每隔3秒刷新并输出一次，3次后程序退出。
    /> pidstat -p 1 2 3 -l
    07:18:58 AM       PID    %usr %system  %guest    %CPU   CPU  Command
    07:18:59 AM         1    0.00    0.00    0.00    0.00     0  /sbin/init
    07:19:00 AM         1    0.00    0.00    0.00    0.00     0  /sbin/init
    07:19:01 AM         1    0.00    0.00    0.00    0.00     0  /sbin/init
    Average:               1    0.00    0.00    0.00    0.00     -  /sbin/init
    %usr：      该进程在用户态的CPU使用率。
    %system：该进程在内核态(系统级)的CPU使用率。
    %CPU：     该进程的总CPU使用率，如果在SMP环境下，该值将除以CPU的数量，以表示每CPU的数据。
    CPU:         该进程所依附的CPU编号(0表示第一个CPU)。

    #监控pid为1(init)的进程的设备IO资源负载情况，其中每隔2秒刷新并输出一次，3次后程序退出。
    /> pidstat -p 1 2 3 -d    
    07:24:49 AM       PID   kB_rd/s   kB_wr/s kB_ccwr/s  Command
    07:24:51 AM         1      0.00      0.00      0.00  init
    07:24:53 AM         1      0.00      0.00      0.00  init
    07:24:55 AM         1      0.00      0.00      0.00  init
    Average:               1      0.00      0.00      0.00  init
    kB_rd/s:   该进程每秒的字节读取数量(KB)。
    kB_wr/s:   该进程每秒的字节写出数量(KB)。
    kB_ccwr/s: 该进程每秒取消磁盘写入的数量(KB)。

    #监控pid为1(init)的进程的内存使用情况，其中每隔2秒刷新并输出一次，3次后程序退出。
    /> pidstat -p 1 2 3 -r
    07:29:56 AM       PID  minflt/s  majflt/s     VSZ    RSS   %MEM  Command
    07:29:58 AM         1      0.00      0.00    2828   1368   0.13  init
    07:30:00 AM         1      0.00      0.00    2828   1368   0.13  init
    07:30:02 AM         1      0.00      0.00    2828   1368   0.13  init
    Average:               1      0.00      0.00    2828   1368   0.13  init
    %MEM:  该进程的内存使用百分比。

    #监控pid为1(init)的进程任务切换情况，其中每隔2秒刷新并输出一次，3次后程序退出。
    /> pidstat -p 1 2 3 -w
    07:32:15 AM       PID   cswch/s nvcswch/s  Command
    07:32:17 AM         1      0.00      0.00  init
    07:32:19 AM         1      0.00      0.00  init
    07:32:21 AM         1      0.00      0.00  init
    Average:            1      0.00      0.00  init
    cswch/s:    每秒任务主动(自愿的)切换上下文的次数。主动切换是指当某一任务处于阻塞等待时，将主动让出自己的CPU资源。
    nvcswch/s: 每秒任务被动(不自愿的)切换上下文的次数。被动切换是指CPU分配给某一任务的时间片已经用完，因此将强迫该进程让出CPU的执行权。

    #监控pid为1(init)的进程及其内部线程的内存(r选项)使用情况，其中每隔2秒刷新并输出一次，3次后程序退出。需要说明的是，如果-t选项后面不加任何其他选项，缺省监控的为CPU资源。结果中黄色高亮的部分表示进程和其内部线程是树状结构的显示方式。
    /> pidstat -p 1 2 3 -tr
    Linux 2.6.32-71.el6.i686 (Stephen-PC)   11/16/2011      _i686_  (1 CPU)

    07:37:04 AM      TGID       TID  minflt/s  majflt/s     VSZ    RSS   %MEM  Command
    07:37:06 AM         1         -      0.00      0.00        2828   1368      0.13  init
    07:37:06 AM         -         1      0.00      0.00        2828   1368      0.13  |__init

    07:37:06 AM      TGID       TID  minflt/s  majflt/s     VSZ    RSS   %MEM  Command
    07:37:08 AM         1         -      0.00      0.00        2828   1368      0.13  init
    07:37:08 AM         -         1      0.00      0.00        2828   1368      0.13  |__init

    07:37:08 AM      TGID       TID  minflt/s  majflt/s     VSZ    RSS   %MEM  Command
    07:37:10 AM         1         -      0.00      0.00        2828   1368      0.13  init
    07:37:10 AM         -         1      0.00      0.00        2828   1368      0.13  |__init

    Average:         TGID       TID  minflt/s  majflt/s     VSZ    RSS   %MEM  Command
    Average:            1         -      0.00      0.00        2828   1368      0.13  init
    Average:            -         1      0.00      0.00        2828   1368      0.13  |__init
    TGID: 线程组ID。
    TID： 线程ID。  

    以上监控不同资源的选项可以同时存在，这样就将在一次输出中输出多种资源的使用情况，如：pidstat -p 1 -dr。

    7.  报告磁盘空间使用状况(df):
    该命令最为常用的选项就是-h，该选项将智能的输出数据单位，以便使输出的结果更具可读性。
    /> df -h
    Filesystem             Size  Used   Avail Use% Mounted on
    /dev/sda1             5.8G  3.3G  2.2G  61%   /
    tmpfs                  504M  260K  504M   1%  /dev/shm

    8.  评估磁盘的使用状况(du)：

选项	说明
-a	包括了所有的文件，而不只是目录。
-b	以字节为计算单位。
-k	以千字节(KB)为计算单位。
-m	以兆字节(MB)为计算单位。
-h	是输出的信息更易于阅读。
-s	只显示工作目录所占总空间。
--exclude=PATTERN	排除掉符合样式的文件,Pattern就是普通的Shell样式，？表示任何一个字符，*表示任意多个字符。
--max-depth=N	从当前目录算起，目录深度大于N的子目录将不被计算，该选项不能和s选项同时存在。 

    #仅显示子一级目录的信息。
    /> du --max-depth=1 -h
    246M    ./stephen
    246M    .    
    /> du -sh ./*   #获取当前目录下所有子目录所占用的磁盘空间大小。
    352K    ./MemcachedTest
    132K    ./Test
    33M     ./thirdparty    
    #在当前目录下，排除目录名模式为Te*的子目录(./Test)，输出其他子目录占用的磁盘空间大小。
    /> du --exclude=Te* -sh ./*  
    352K    ./MemcachedTest
    33M     ./thirdparty

## 十九.  和系统运行进程相关的Shell命令:
    
   1.  进程监控命令(ps):
    要对进程进行监测和控制，首先必须要了解当前进程的情况，也就是需要查看当前进程，而ps命令就是最基本同时也是非常强大的进程查看命令。使用该命令可以确定有哪些进程正在运行和运行的状态、进程是否结束、进程有没有僵死、哪些进程占用了过多的资源等等。总之大部分信息都是可以通过执行该命令得到的。
    ps命令存在很多的命令行选项和参数，然而我们最为常用只有两种形式，这里先给出与它们相关的选项和参数的含义：

选项	说明
a	显示终端上的所有进程，包括其他用户的进程。
u	以用户为主的格式来显示程序状况。
x	显示所有程序，不以终端来区分。
-e	显示所有进程。
o	其后指定要输出的列，如user，pid等，多个列之间用逗号分隔。
-p	后面跟着一组pid的列表，用逗号分隔，该命令将只是输出这些pid的相关数据。
    /> ps aux

    root         1  0.0  0.1   2828  1400 ?        Ss   09:51   0:02 /sbin/init
    root         2  0.0  0.0      0          0 ?        S    09:51   0:00 [kthreadd]
    root         3  0.0  0.0      0          0 ?        S    09:51   0:00 [migration/0]
    ... ...   
    /> ps -eo user,pid,%cpu,%mem,start,time,command | head -n 4
    USER       PID %CPU %MEM  STARTED     TIME        COMMAND
    root         1         0.0    0.1   09:51:08     00:00:02  /sbin/init
    root         2         0.0    0.0   09:51:08     00:00:00  [kthreadd]
    root         3         0.0    0.0   09:51:08     00:00:00  [migration/0]
    这里需要说明的是，ps中存在很多和进程性能相关的参数，它们均以输出表格中的列的方式显示出来，在这里我们只是给出了非常常用的几个参数，至于更多参数，我们则需要根据自己应用的实际情况去看ps的man手册。
    #以完整的格式显示pid为1(init)的进程的相关数据
    /> ps -fp 1
    UID        PID  PPID  C STIME TTY          TIME   CMD
    root         1        0  0 05:16   ?        00:00:03 /sbin/init
    
    2.  改变进程优先级的命令(nice和renice):
    该Shell命令最常用的使用方式为：nice [-n <优先等级>][执行指令]，其中优先等级的范围从-20-19，其中-20最高，19最低，只有系统管理者可以设置负数的等级。
    #后台执行sleep 100秒，同时在启动时将其nice值置为19
    /> nice -n 19 sleep 100 &
    [1] 4661
    #后台执行sleep 100秒，同时在启动时将其nice值置为-19
    /> nice -n -19 sleep 100 &
    [2] 4664
    #关注ps -l输出中用黄色高亮的两行，它们的NI值和我们执行是设置的值一致。
    /> ps -l
    F S   UID   PID  PPID  C PRI  NI  ADDR  SZ    WCHAN  TTY       TIME        CMD
    4 S     0  2833  2829  0  80   0     -      1739     -         pts/2    00:00:00  bash
    0 S     0  4661  2833  0  99  19    -      1066     -         pts/2    00:00:00  sleep
    4 S     0  4664  2833  0  61 -19    -      1066     -         pts/2    00:00:00  sleep
    4 R     0  4665  2833  1  80   0     -      1231     -         pts/2    00:00:00  ps
    
    renice命令主要用于为已经执行的进程重新设定nice值，该命令包含以下几个常用选项：

选项	说明
-g	使用程序群组名称，修改所有隶属于该程序群组的程序的优先权。
-p	改变该程序的优先权等级，此参数为预设值。
-u	指定用户名称，修改所有隶属于该用户的程序的优先权。
    #切换到stephen用户下执行一个后台进程，这里sleep进程将在后台睡眠1000秒。
    /> su stephen
    /> sleep 1000&  
    [1] 4812
    /> exit   #退回到切换前的root用户
    #查看已经启动的后台sleep进程，其ni值为0，宿主用户为stephen
    /> ps -eo user,pid,ni,command | grep stephen
    stephen   4812   0 sleep 1000
    root        4821    0 grep  stephen
    #以指定用户的方式修改该用户下所有进程的nice值
    /> renice -n 5 -u stephen
    500: old priority 0, new priority 5
    #从再次执行ps的输出结果可以看出，该sleep后台进程的nice值已经调成了5
    /> ps -eo user,pid,ni,command | grep stephen
    stephen   4812   5 sleep 1000
    root         4826   0 grep  stephen
    #以指定进程pid的方式修改该进程的nice值
    /> renice -n 10 -p 4812
    4812: old priority 5, new priority 10
    #再次执行ps，该sleep后台进程的nice值已经从5变成了10
    /> ps -eo user,pid,ni,command | grep stephen
    stephen   4812  10 sleep 1000
    root        4829   0 grep  stephen


    3.  列出当前系统打开文件的工具(lsof):
    lsof(list opened files)，其重要功能为列举系统中已经被打开的文件，如果没有指定任何选项或参数，lsof则列出所有活动进程打开的所有文件。众所周知，linux环境中任何事物都是文件，如设备、目录、sockets等。所以，用好lsof命令，对日常的linux管理非常有帮助。下面先给出该命令的常用选项：

选项	说明
-a	该选项会使后面选项选出的结果列表进行and操作。
-c command_prefix	显示以command_prefix开头的进程打开的文件。
-p PID	显示指定PID已打开文件的信息
+d directory	从文件夹directory来搜寻(不考虑子目录)，列出该目录下打开的文件信息。
+D directory	从文件夹directory来搜寻(考虑子目录)，列出该目录下打开的文件信息。
-d num_of_fd	以File Descriptor的信息进行匹配，可使用3-10，表示范围，3,10表示某些值。
-u user	显示某用户的已经打开的文件，其中user可以使用正则表达式。
-i	监听指定的协议、端口、主机等的网络信息，格式为：[proto][@host|addr][:svc_list|port_list]
    #查看打开/dev/null文件的进程。
    /> lsof /dev/null | head -n 5
    COMMAND    PID      USER   FD   TYPE DEVICE SIZE/OFF NODE NAME
    init         1      root    0u   CHR    1,3      0t0 3671 /dev/null
    init         1      root    1u   CHR    1,3      0t0 3671 /dev/null
    init         1      root    2u   CHR    1,3      0t0 3671 /dev/null
    udevd 397      root    0u   CHR    1,3      0t0 3671 /dev/null

    #查看打开22端口的进程
    /> lsof -i:22
    COMMAND  PID USER   FD   TYPE DEVICE SIZE/OFF NODE NAME
    sshd    1582 root    3u  IPv4  11989      0t0  TCP *:ssh (LISTEN)
    sshd    1582 root    4u  IPv6  11991      0t0  TCP *:ssh (LISTEN)
    sshd    2829 root    3r   IPv4  19635      0t0  TCP bogon:ssh->bogon:15264 (ESTABLISHED)

    #查看init进程打开的文件
    />  lsof -c init
    COMMAND PID USER   FD   TYPE     DEVICE   SIZE/OFF   NODE    NAME
    init               1 root  cwd      DIR        8,2     4096              2        /
    init               1 root  rtd       DIR        8,2     4096              2        /
    init               1 root  txt       REG       8,2   136068       148567     /sbin/init
    init               1 root  mem    REG        8,2    58536      137507     /lib/libnss_files-2.12.so
    init               1 root  mem    REG        8,2   122232     186675     /lib/libgcc_s-4.4.4-20100726.so.1
    init               1 root  mem    REG        8,2   141492     186436     /lib/ld-2.12.so
    init               1 root  mem    REG        8,2  1855584    186631     /lib/libc-2.12.so
    init               1 root  mem    REG        8,2   133136     186632     /lib/libpthread-2.12.so
    init               1 root  mem    REG        8,2    99020      180422     /lib/libnih.so.1.0.0
    init               1 root  mem    REG        8,2    37304      186773     /lib/libnih-dbus.so.1.0.0
    init               1 root  mem    REG        8,2    41728      186633     /lib/librt-2.12.so
    init               1 root  mem    REG        8,2   286380     186634     /lib/libdbus-1.so.3.4.0
    init               1 root    0u     CHR        1,3      0t0           3671      /dev/null
    init               1 root    1u     CHR        1,3      0t0           3671      /dev/null
    init               1 root    2u     CHR        1,3      0t0           3671      /dev/null
    init               1 root    3r      FIFO       0,8      0t0           7969      pipe
    init               1 root    4w     FIFO       0,8      0t0           7969      pipe
    init               1 root    5r      DIR        0,10        0             1         inotify
    init               1 root    6r      DIR        0,10        0             1         inotify
    init               1 root    7u     unix   0xf61e3840  0t0       7970      socket
    init               1 root    9u     unix   0xf3bab280  0t0      11211     socket
    在上面输出的FD列中，显示的是文件的File Descriptor number，或者如下的内容：
    cwd:  current working directory;
    mem:  memory-mapped file;
    mmap: memory-mapped device;
    pd:   parent directory;
    rtd:  root directory;
    txt:  program text (code and data);
    文件的File Descriptor number显示模式有:
    r for read access;
    w for write access;
    u for read and write access;

    在上面输出的TYPE列中，显示的是文件类型，如：
    DIR:  目录
    LINK: 链接文件
    REG:  普通文件


    #查看pid为1的进程(init)打开的文件，其输出结果等同于上面的命令，他们都是init。
    /> lsof -p 1
    #查看owner为root的进程打开的文件。
    /> lsof -u root
    #查看owner不为root的进程打开的文件。
    /> lsof -u ^root
    #查看打开协议为tcp，ip为192.168.220.134，端口为22的进程。
    /> lsof -i tcp@192.168.220.134:22
    COMMAND  PID USER   FD   TYPE DEVICE SIZE/OFF NODE NAME
    sshd        2829 root     3r    IPv4  19635      0t0      TCP    bogon:ssh->bogon:15264 (ESTABLISHED)    
    #查看打开/root文件夹，但不考虑目录搜寻
    /> lsof +d /root
    #查看打开/root文件夹以及其子目录搜寻
    /> lsof +D /root
    #查看打开FD(0-3)文件的所有进程
    /> lsof -d 0-3
    #-a选项会将+d选项和-c选项的选择结果进行and操作，并输出合并后的结果。
    /> lsof +d .
    COMMAND  PID USER   FD   TYPE DEVICE SIZE/OFF  NODE NAME
    bash       9707  root  cwd    DIR    8,1     4096         39887 .
    lsof         9791  root  cwd    DIR    8,1     4096         39887 .
    lsof         9792  root  cwd    DIR    8,1     4096         39887 .
    /> lsof -a -c bash +d .
    COMMAND  PID USER   FD   TYPE DEVICE SIZE/OFF  NODE NAME
    bash        9707 root  cwd    DIR    8,1     4096         39887 .
    最后需要额外说明的是，如果在文件名的末尾存在(delete)，则说明该文件已经被删除，只是还存留在cache中。


    4.  进程查找/杀掉命令(pgrep/pkill)：
    查找和杀死指定的进程, 他们的选项和参数完全相同, 这里只是介绍pgrep。下面是常用的命令行选项：

选项	说明
-d	定义多个进程之间的分隔符, 如果不定义则使用换行符。
-n	表示如果该程序有多个进程正在运行，则仅查找最新的，即最后启动的。
-o	表示如果该程序有多个进程正在运行，则仅查找最老的，即最先启动的。
-G	其后跟着一组group id，该命令在搜索时，仅考虑group列表中的进程。
-u	其后跟着一组有效用户ID(effetive user id)，该命令在搜索时，仅考虑该effective user列表中的进程。
-U	其后跟着一组实际用户ID(real user id)，该命令在搜索时，仅考虑该real user列表中的进程。
-x	表示进程的名字必须完全匹配, 以上的选项均可以部分匹配。
-l	将不仅打印pid,也打印进程名。
-f	一般与-l合用, 将打印进程的参数。
    #手工创建两个后台进程
    /> sleep 1000&
    3456
    /> sleep 1000&
    3457

    #查找进程名为sleep的进程，同时输出所有找到的pid
    /> pgrep sleep
    3456
    3457
    #查找进程名为sleep的进程pid，如果存在多个，他们之间使用:分隔，而不是换行符分隔。
    /> pgrep -d: sleep
    3456:3457
    #查找进程名为sleep的进程pid，如果存在多个，这里只是输出最后启动的那一个。
    /> pgrep -n sleep
    3457
    #查找进程名为sleep的进程pid，如果存在多个，这里只是输出最先启动的那一个。
    /> pgrep -o  sleep
    3456
    #查找进程名为sleep，同时这个正在运行的进程的组为root和stephen。
    /> pgrep -G root,stephen sleep
    3456
    3457
    #查找有效用户ID为root和oracle，进程名为sleep的进程。
    /> pgrep -u root,oracle sleep
    3456
    3457
    #查找实际用户ID为root和oracle，进程名为sleep的进程。
    /> pgrep -U root,oracle sleep
    3456
    3457
    #查找进程名为sleep的进程，注意这里找到的进程名必须和参数中的完全匹配。
    /> pgrep -x sleep
    3456
    3457
    #-x不支持部分匹配，sleep进程将不会被查出，因此下面的命令没有结果。
    /> pgrep -x sle
    #查找进程名为sleep的进程，同时输出所有找到的pid和进程名。    
    /> pgrep -l sleep
    3456 sleep
    3457 sleep
    #查找进程名为sleep的进程，同时输出所有找到的pid、进程名和启动时的参数。
    /> pgrep -lf sleep
    3456 sleep 1000
    3457 sleep 1000
    #查找进程名为sleep的进程，同时以逗号为分隔符输出他们的pid，在将结果传给ps命令，-f表示显示完整格式，-p显示pid列表，ps将只是输出该列表内的进程数据。
    /> pgrep -f sleep -d, | xargs ps -fp
    UID        PID  PPID  C STIME TTY          TIME CMD
    root      3456  2138  0 06:11 pts/5    00:00:00 sleep 1000
    root      3457  2138  0 06:11 pts/5    00:00:00 sleep 1000  

## 二十. 通过管道组合Shell命令获取系统运行数据:


    1.  输出当前系统中占用内存最多的5条命令:
    #1) 通过ps命令列出当前主机正在运行的所有进程。
    #2) 按照第五个字段基于数值的形式进行正常排序(由小到大)。
    #3) 仅显示最后5条输出。
    /> ps aux | sort -k 5n | tail -5
    stephen   1861  0.2  2.0  96972 21596  ?  S     Nov11   2:24 nautilus
    stephen   1892  0.0  0.4 102108  4508  ?  S<sl Nov11   0:00 /usr/bin/pulseaudio 
    stephen   1874  0.0  0.9 107648 10124 ?  S     Nov11   0:00 gnome-volume
    stephen   1855  0.0  1.2 123776 13112 ?  Sl     Nov11   0:00 metacity
    stephen   1831  0.0  0.9 125432  9768  ?  Ssl   Nov11   0:05 /usr/libexec/gnome
    
    2.  找出cpu利用率高的20个进程:
    #1) 通过ps命令输出所有进程的数据，-o选项后面的字段列表列出了结果中需要包含的数据列。
    #2) 将ps输出的Title行去掉，grep -v PID表示不包含PID的行。
    #3) 基于第一个域字段排序，即pcpu。n表示以数值的形式排序。
    #4) 输出按cpu使用率排序后的最后20行，即占用率最高的20行。
    /> ps -e -o pcpu,pid,user,sgi_p,cmd | grep -v PID | sort -k 1n | tail -20

    3.  获取当前系统物理内存的总大小:
    #1) 以兆(MB)为单位输出系统当前的内存使用状况。
    #2) 通过grep定位到Mem行，该行是以操作系统为视角统计数据的。
    #3) 通过awk打印出该行的第二列，即total列。
    /> free -m | grep "Mem" | awk '{print $2, "MB"}'
    1007 MB

## 二十一. 通过管道组合Shell命令进行系统管理:

    1.  获取当前或指定目录下子目录所占用的磁盘空间，并将结果按照从大到小的顺序输出:
    #1) 输出/usr的子目录所占用的磁盘空间。
    #2) 以数值的方式倒排后输出。
    /> du -s /usr/* | sort -nr
    1443980 /usr/share
    793260   /usr/lib
    217584   /usr/bin
    128624   /usr/include
    60748    /usr/libexec
    45148    /usr/src
    21096    /usr/sbin
    6896      /usr/local
    4           /usr/games
    4           /usr/etc
    0           /usr/tmp
    
    2.  批量修改文件名:
    #1) find命令找到文件名扩展名为.output的文件。
    #2) sed命令中的-e选项表示流编辑动作有多次，第一次是将找到的文件名中相对路径前缀部分去掉，如./aa改为aa。
    #    流编辑的第二部分，是将20110311替换为mv & 20110310，其中&表示s命令的被替换部分，这里即源文件名。
    #    \1表示被替换部分中#的\(.*\)。
    #3) 此时的输出应为
    #    mv 20110311.output 20110310.output
    #    mv 20110311abc.output 20110310abc.output
    #    最后将上面的输出作为命令交给bash命令去执行，从而将所有20110311*.output改为20110311*.output
    /> find ./ -name "*.output" -print  | sed -e 's/.\///g' -e 's/20110311\(.*\)/mv & 20110310\1/g' | bash
    
    3.  统计当前目录下文件和目录的数量:
    #1) ls -l命令列出文件和目录的详细信息。
    #2) ls -l输出的详细列表中的第一个域字段是文件或目录的权限属性部分，如果权限属性部分的第一个字符为d，
    #    该文件为目录，如果是-，该文件为普通文件。
    #3) 通过wc计算grep过滤后的行数。
    /> ls -l * | grep "^-" | wc -l
    /> ls -l * | grep "^d" | wc -l
    
    4.  杀掉指定终端的所有进程:
    #1) 通过ps命令输出终端为pts/1的所有进程。
    #2) 将ps的输出传给grep，grep将过滤掉ps输出的Title部分，-v PID表示不包含PID的行。
    #3) awk打印输出grep查找结果的第一个字段，即pid字段。
    #4) 上面的三个组合命令是在反引号内被执行的，并将执行的结果赋值给数组变量${K}。
    #5) kill方法将杀掉数组${K}包含的pid。
    /> kill -9 ${K}=`ps -t pts/1 | grep -v PID | awk '{print $1}'`    

    5.  将查找到的文件打包并copy到指定目录:
    #1) 通过find找到当前目录下(包含所有子目录)的所有*.txt文件。
    #2) tar命令将find找到的结果压缩成test.tar压缩包文件。
    #3) 如果&&左侧括号内的命令正常完成，则可以执行&&右侧的shell命令了。
    #4) 将生成后的test.tar文件copy到/home/.目录下。
    /> (find . -name "*.txt" | xargs tar -cvf test.tar) && cp -f test.tar /home/.
    
    #1) cpio从find的结果中读取文件名，将其打包压缩后发送到./dest/dir(目标目录)。
    #2) cpio的选项介绍：
    #    -d：创建需要的目录。
    #    -a：重置源文件的访问时间。
    #    -m：保护新文件的修改时间。
    #    -p：将cpio设置为copy pass-through模式。 
    /> find . -name "*" | cpio -dampv ./dest/dir

　  最后需要说明的是，该篇Blog中绝大多数的示例来自于互联网，是本人经过一天左右的时间收集和整理之后筛选出来的，其中注释部分是我在后来添加的，以便于我们阅读时的理解。如果今后再发现更好更巧妙的Shell组合命令，本人将持续更新该Blog。如果您有确实非常不错的Shell命令组合，且愿意和我们在这里分享，可以直接放在回复中，本人将对该篇Blog始终保持重点关注。

## 二十二. 交互式使用Bash Shell:

    1.  用set命令设置bash的选项:
    下面为set主要选项的列表及其表述：

选项名	开关缩写	描述
allexport	-a	打开此开关，所有变量都自动输出给子Shell。
noclobber	-C	防止重定向时文件被覆盖。
noglob	-d	在路径和文件名中，关闭通配符。

    #打开该选项
    /> set -o allexport   #等同于set -a
    #关闭该选项
    /> set +o allexport  #等同于set +a
    #列出当前所有选项的当前值。
    /> set -o
    allexport         off
    braceexpand   on
    emacs             on
    errexit            off
    errtrace          off
    functrace        off
    hashall            on
    histexpand      on
    ... ...
    /> set -o noclobber     #打开noclobber选项，防止在重定向时原有文件被覆盖。
    /> date > outfile         #通过date命令先生成一个文件outfile。
    /> ls > outfile             #将ls命令的输出重定向到该文件outfile，shell将提示不能覆盖已经存在的文件。
    -bash: outfile: cannot overwrite existing file
    /> set +o noclobber    #关闭noclobber选项。
    /> ls > outfile             #重新将ls的输出重定向到outfile，成功。


    2.  变量:
    设置局部变量:
    /> name="stephen liu"  #注意等号两边不要有空格，如果变量值之间存在空格，则需要用双引号括起
    /> echo $name
    stephen liu
    /> name=                    #将变量设置为空时，等号后面也不要有空格，直接回车即可。
    /> echo $name             #name变量为空，因此echo不会有任何输出。
    注意：以上变量的声明方式均可替换为declare variable=value的形式。

    /> declare name="stephen liu"
    /> readonly name         #将name变量设置为只读。
    /> echo $name
    stephen liu
    /> name="my wife"      #如果针对只读变量重新赋值，将报错，提示name是只读变量。
    -bash: name: readonly variable
    /> unset name             #如果unset只读变量，将同样报错，提示不能unset只读变量。
    -bash: unset: name: cannot unset: readonly variable

    设置全局/环境变量:
    在当前Shell中创建的全局/环境变量可以直接传递给它所有的子Shell，当前创建环境变量的Shell被称为夫Shell。
    /> export allname=john         #利用export命令，将其后声明的变量置为环境变量
    /> bash                                #启动一个新的子Shell
    /> echo $allname                  #在子Shell中echo变量$allname，发现夫Shell中设置的值被传递到子Shell
    john
    /> declare -x allname2=peter #这里的功能和结果都是和上面的命令相同，只是利用declare -x命令设置环境变量
    /> bash
    /> echo $allname2
    peter

    下面的列表将给出常用的内置Shell环境变量:

变量名	含义
BASH	表示bash命令的完整路径名。
ENV	在启动新bash shell时执行的环境文件名。
HOME	主目录。
LANG	本地化语言。
PATH	命令搜索路径，彼此之间冒号分隔。
PPID	父进程PID。
PWD	当前工作目录，用cd命令设置。

    3.  echo命令:
    该命令主要用于将其参数打印到标准输出。其中-e选项使得echo命令可以无限制地使用转义序列控制输出的效果。下面的列表给出常用的转义序列。

转义序列	功能
\c	不换行打印
\n	换行
\t	制表符
\\	反斜杠
    echo还提供了一个常用的-n选项，其功能不输出换行符。

    /> echo The username is $LOGNAME
    The username is stephen
    #下面命令的输出中的“/>”表示命令行提示符。
    /> echo -e "\tHello World\c"
        Hello World />
    /> echo -n "Hello World"
    Hello World />

    4.  printf命令:
    该命令和C语言中的printf函数的功能相同，都用用来格式化输出的。格式包括字符串本身和描述打印效果的字符。定义格式的方法和C语言也是完全一样的，即在%后面跟一个说明符，如%f表示后面是一个浮点数，%d表示一个整数。printf命令也同样支持转义序列符，其常用转义序列如下：

转义序列	功能
\c	不换行打印
\n	换行
\t	制表符
\\	反斜杠
\"	双引号
    其常用的格式化说明符列表如下：

说明符	描述
%c	ASCII字符
%d,%i	十进制整数
%f	浮点格式
%o	不带正负号的八进制值
%s	字符串
%u	不带正负号的十进制值
%x	不带正负号的十六进制值，其中使用a-f表示10-15
%X	不带正负号的十六进制值，其中使用A-F表示10-15
%%	表示%本身
    下面是printf的一些常用使用方式：
    /> printf "The number is %.2f.\n" 100   这里.2f表示保留小数点后两位
    The number is 100.00.
    #%-20s表示一个左对齐、宽度为20个字符字符串格式，不足20个字符，右侧补充相应数量的空格符。
    #%-15s表示一个左对齐、宽度为15个字符字符串格式。
    #%10.2f表示右对齐、10个字符长度的浮点数，其中一个是小数点，小数点后面保留两位。
    /> printf "%-20s%-15s%10.2f\n" "Stephen" "Liu" 35
    Stephen             Liu                 35.00
    #%10s表示右对齐、宽度为10的字符串，如不足10个字符，左侧补充相应数量的空格符。
    /> printf "|%10s|\n" hello
    |     hello|
    在printf中还有一些常用的标志符，如上面例子中的符号(-)，这里我们在介绍一个比较常用的标识符"#"
    #如果#标志和%x/%X搭配使用，在输出十六进制数字时，前面会加0x/0X前缀。
    /> printf "%x %#x\n" 15 15
    f 0xf


   5.  变量替换运算符:
    bash中提供了一组可以同时检验和修改变量的特定修改符。这些修改符提供了一个快捷的方法来检验变量是不是被设置过，并把输出结果输出到一个变量中，见下表：

修改符	描述	用途
${variable:-word}	如variable被设置且非空，则返回该值，否则返回word，变量值不变。	如变量未定义，返回默认值。
${variable-word}	如variable未被设置，则返回word，变量值不变，如果设置变量，则返回变量值，即使变量的值为空值。	如变量未设置，返回默认值。
${variable:=word}	如variable被设置且非空，则返回该值，否则设置变量为word，同时返回word。	如果变量未定义，则设置其为默认值。
${variable=word}	如variable未设置，则设置变量为word，同时返回word，如果variable被设置且为空，将返回空值，同时variable不变。否则返回variable值，同时variable不变。	如果变量未设置，则设置其为默认值。
${variable:+word}	如variable被设置且非空，则返回word，否则返回null，变量值不变。	用于测试变量是否存在。
${variable+word}	如variable被设置(即使是空值)，则返回word，否则返回空。	用于测试变量是否设置。
${variable:?word}	如variable被设置且非空，则返回该值，否则显示word，然后退出Shell。如果word为空，打印"parameter null or not set"	为了捕捉由于变量未定义所导致的错误。
${variable:offset}	从variable的offset位置开始取，直到末尾。	 
${variable:offset:length}	从variable的offset位置开始取length个字符。	 

    #${variable:-word}的示例，其C语言表示形式为：
    #    if (NULL == variable)
    #        return word;
    #    else
    #        return $variable;
    /> unset var_name                        #将变量var_name置为空。
    /> var_name=
    /> echo ${var_name:-NewValue}    #var_name为空，因此返回NewValue
    NewValue
    /> echo $var_name                        #var_name的值未变化，仍然为空。

    /> var_name=OldValue                   #给var_name赋值。
    /> echo ${var_name:-NewValue}    #var_name非空，因此返回var_name的原有值。
    OldValue
    /> echo $var_name                        #var_name的值未变化，仍然OldValue。
    OldValue

    #${variable-word}的示例，其伪码表示形式为：
    #    if (variable is NOT set)
    #        return word;
    #    else
    #        return $variable;
    /> unset var_name                         #取消该变量var_name的设置。
    /> echo ${var_name-NewValue}    #var_name为空，因此返回NewValue
    NewValue
    /> echo $var_name                        #var_name的值未变化，仍然为空。

    /> var_name=OldValue                   #给var_name赋值，即便执行var_name=，其结果也是一样。
    /> echo ${var_name-NewValue}    #var_name非空，因此返回var_name的原有值。
    OldValue
    /> echo $var_name                        #var_name的值未变化，仍然OldValue。
    OldValue

    
    #${variable:=word}的示例，其C语言表示形式为：
    #    if (NULL == variable) {
    #        variable=world;
    #        return word;
    #    } else {
    #        return $variable;
    #    }
    /> unset var_name                        #将变量var_name置为空。
    /> var_name=
    /> echo ${var_name:=NewValue}   #var_name为空，设置变量为NewValue同时返回NewValue。
    NewValue
    /> echo $var_name                        #var_name的值已经被设置为NewValue。
    NewValue
    /> var_name=OldValue                  #给var_name赋值。
    /> echo ${var_name:=NewValue}   #var_name非空，因此返回var_name的原有值。
    OldValue
    /> echo $var_name                       #var_name的值未变化，仍然OldValue。
    OldValue
    
    #${variable=word}的示例，其伪码表示形式为：
    #    if (variable is NOT set) {
    #        variable=world;
    #        return word;
    #    } else if (variable == NULL) {
    #        return $variable;  //variable is NULL
    #    } else {
    #        return $variable;
    #    }
    /> unset var_name                        #取消该变量var_name的设置。
    /> echo ${var_name=NewValue}  #var_name未被设置，设置变量为NewValue同时返回NewValue。
    NewValue
    /> echo $var_name                        #var_name的值已经被设置为NewValue。
    NewValue
    /> var_name=                              #设置变量var_name，并给该变量赋空值。
    /> echo ${var_name=NewValue}  #var_name被设置，且为空值，返回var_name的原有空值。
    
    /> echo $var_name                       #var_name的值未变化，仍未空值。
    
    /> var_name=OldValue                  #给var_name赋值。
    /> echo ${var_name=NewValue}  #var_name非空，因此返回var_name的原有值。
    OldValue
    /> echo $var_name                       #var_name的值未变化，仍然OldValue。
    OldValue

    #${variable:+word}的示例，其C语言表示形式为：
    #    if (NULL != variable)
    #        return word;
    #    else
    #        return $variable;
    /> var_name=OldValue                  #设置变量var_name，其值为非空。
    /> echo ${var_name:+NewValue}   #由于var_name有值，因此返回NewValue
    NewValue
    /> echo $var_name                       #var_name的值仍然为远之OldValue
    OldValue
    /> unset var_name                        #将var_name置为空值。
    /> var_name=
    /> echo ${var_name:+NewValue}   #由于var_name为空，因此返回null。
    /> echo $var_name                       #var_name仍然保持原有的空值。

    #${variable+word}的示例，其伪码表示形式为：
    #    if (variable is set)
    #        return word;
    #    else
    #        return $variable;
    /> var_name=OldValue                  #设置变量var_name，其值为非空。
    /> echo ${var_name+NewValue}   #由于var_name有值，因此返回NewValue
    NewValue
    /> echo $var_name                       #var_name的值仍然为远之OldValue
    OldValue
    /> unset var_name                        #取消对变量var_name的设置。
    /> echo ${var_name+NewValue}   #返回空值。
    /> echo $var_name                       #var_name仍未被设置。

    #${variable:?word}的示例，其C语言表示形式为：
    #    if (NULL != variable) {
    #        return variable;
    #    } else {
    #        if (NULL != word)
    #            return "variable : word";
    #        else
    #            return "parameter null or not set";
    #    }
    /> var_name=OldValue                  #设置变量var_name，其值为非空。
    /> echo ${var_name:?NewValue}   #由于var_name有值，因此返回变量的原有值
    OldValue
    /> unset var_name                        #将var_name置为空值。
    /> var_name=
    /> echo ${var_name:?NewValue}   #由于var_name为空，因此返回word。
    -bash: var_name: NewValue
    /> echo $var_name                       #var_name仍然保持原有的空值。

    /> echo ${var_name:?}                #如果word为空，返回下面的输出。
    -bash: var_name: parameter null or not set

    #${variable:offset}示例：
    /> var_name=notebook
    /> echo ${var_name:2}
    tebook
    /> echo ${var_name:0}                #如果offset为0，则取var_name的全部值。
    notebook

    ${variable:offset:length}示例：
    /> var_name=notebook
    /> echo ${var_name:0:4}
    note
    /> echo ${var_name:4:4}
    book

    6.  变量模式匹配运算符:
    Shell中还提供了一组模式匹配运算符，见下表：

运算符	替换
${variable#pattern}	如果模式匹配变量值的开头，则删除匹配的最短部分，并返回剩下的部分，变量原值不变。
${variable##pattern}	如果模式匹配变量值的开头，则删除匹配的最长部分，并返回剩下的部分，变量原值不变。
${variable%pattern}	如果模式匹配变量值的结尾，则删除匹配的最短部分，并返回剩下的部分，变量原值不变。
${variable%%pattern}	如果模式匹配变量值的结尾，则删除匹配的最长部分，并返回剩下的部分，变量原值不变。
${#variable}	返回变量中字母的数量。

    #${variable#pattern}示例：
    /> pathname="/home/stephen/mycode/test.h"
    /> echo ${pathname#/home}        #变量pathname开始处匹配/home的最短部分被删除。
    /stephen/mycode/test.h
    /> echo $pathname                       #pathname的原值不变。
    /home/stephen/mycode/test.h

    #${variable##pattern}示例：
    /> pathname="/home/stephen/mycode/test.h"
    /> echo ${pathname##*/}            #变量pathname开始处匹配*/的最长部分被删除，*表示任意字符。
    test.h
    /> echo $pathname                       #pathname的原值不变。
    /home/stephen/mycode/test.h

    #${variable%pattern}示例：
    /> pathname="/home/stephen/mycode/test.h"
    /> echo ${pathname%/*}             #变量pathname结尾处匹配/*的最短部分被删除。
    /home/stephen/mycode
    /> echo $pathname                       #pathname的原值不变。
    /home/stephen/mycode/test.h

    #${variable%%pattern}示例：
    /> pathname="/home/stephen/mycode/test.h"
    /> echo ${pathname%%/*}          #变量pathname结尾处匹配/*的最长部分被删除，这里所有字符串均被删除。

    /> echo $pathname                       #pathname的原值不变。
    /home/stephen/mycode/test.h

    #${#variable}示例:
    /> name="stephen liu"
    /> echo ${#name}
    11

    7.  Shell中的内置变量:
    Shell中提供了一些以$开头的内置变量，见下表：

变量名	描述
$?	表示Shell命令的返回值
$$	表示当前Shell的pid
$-	表示当前Shell的命令行选项
$!	最后一个放入后台作业的PID值
$0	表示脚本的名字
$1--$9	表示脚本的第一到九个参数
${10}	表示脚本的第十个参数
$#	表示参数的个数
$*,$@	表示所有的参数，有双引号时除外，"$*"表示赋值到一个变量，"$@"表示赋值到多个。
    所有的内置变量都比较容易理解，因此这里仅给出$*和$@的区别用法：
    /> set 'apple pie' pears peaches
    /> for i in $*
    >  do
    >  echo $i
    >  done
    apple
    pie
    pears
    peaches

    /> set 'apple pie' pears peaches
    /> for i in $@
    >  do
    >  echo $i
    >  done
    apple
    pie
    pears
    peaches


    /> set 'apple pie' pears peaches
    /> for i in "$*"           #将所有参数变量视为一个
    >  do
    >  echo $i
    >  done
    apple pie pears    peaches
    
    /> set 'apple pie' pears peaches
    /> for i in "$@"
    >  do
    >  echo $i
    >  done
    apple pie                   #这里的单引号将两个单词合成一个.
    pears
    peaches    
    
    8.  引用:
    Shell中提供三种引用字符，分别是：反斜杠、单引号和双引号，它们可以使Shell中所有元字符失去其特殊功能，而还原其本意。见以下元字符列表：

元字符	描述
;	命令分隔符
&	后台处理Shell命令
()	命令组，创建一个子Shell
{}	命令组，但是不创建子Shell
|	管道
< >	输入输出重定向
$	变量前缀
*[]?	用于文件名扩展的Shell通配符
    注：单引号和双引号唯一的区别就是，双引号内可以包含变量和命令替换，而单引号则不会解释这些，见如下示例：
    /> name=Stephen
    /> echo "Hi $name, I'm glad to meet you! "  #name变量被替换
    Hi Stephen, I'm glad to meet you!
    /> echo 'Hi $name, I am glad to meet you! ' #name变量没有被替换
    Hi $name, I am glad to meet you!
    /> echo "Hey $name, the time is $(date)"      #name变量和date命令均被替换
    Hey Stephen, the time is Fri Nov 18 16:27:31 CST 2011
    /> echo 'Hey $name, the time is $(date)'
    Hey $name, the time is $(date)                      #name变量和date命令均未被替换

    9.  命令替换:
    同样我们需要把命令的输出结果赋值给一个变量或者需要用字符串替换变量的输出结果时，我们可以使用变量替换。在Shell中，通常使用反引号的方法进行命令替换。
    /> d=`date`                           #将date命令的执行结果赋值给d变量。
    /> echo $d
    Fri Nov 18 16:35:28 CST 2011
    /> pwd
    /home/stephen
    /> echo `basename \`pwd\``  #基于反引号的命令替换是可嵌入的，但是嵌入命令的反引号需要使用反斜杠转义。
    stephen
    除了反引号可以用于命令替换，这里我们也可以使用$(command)形式用于命令替换。
    /> d=$(date)
    /> echo $d
    Fri Nov 18 16:42:33 CST 2011
    /> dirname="$(basename $(pwd))"     #和之前的反引号一样，该方式也支持嵌套。
    /> echo $dirname
    stephen

   10.  数学扩展:
    Shell中提供了两种计算数学表达式的格式：$[ expression ]和$(( expression ))。
    /> echo $[5+4-2]
    7
    /> echo $[5+2*3]
    11
    /> echo $((5+4-2))
    7
    /> echo $((5+2*3))
    11    
    事实上，我们也可以在Shell中声明数值型的变量，这需要在declare命令的后面添加-i选项，如：
    /> declare -i num
    /> num=5+5      #注意在赋值的过程中，所有的符号之间均没有空格，如果需要空格，需要在表达式的外面加双引号
    /> echo $num    #如果没有声明declare -i num，该命令将返回5+5
    10
    /> num="5 * 5"
    /> echo $num
    25
    /> declare strnum
    /> strnum=5+5
    /> echo $strnum #由于并未将strnum声明为数值型，因此该输出将按字符串方式处理。
    5+5
    Shell还允许我们以不同进制的方式显示数值型变量的输出结果，其格式为：进制+#+变量。
    /> declare -i x=017 #017其格式为八进制格式
    /> echo $x           #缺省是十进制
    15
    /> x=2#101         #二进制
    /> echo $x
    5
    /> x=8#17           #八进制
    /> echo $x
    15
    /> x=16#b           #十六进制
    /> echo $x
    11
    在Shell中还提供了一个内置命令let，专门用于计算数学运算的，见如下示例：
    /> let i=5
    /> let i=i+1
    /> echo $i
    6
    /> let "i = i + 2"
    /> echo $i
    8
    /> let "i+=1"
    /> echo $i
    9    

    11.  数组:
    Shell中提供了创建一维数组的能力，你可以把一串数字、名字或者文件放在一个变量中。使用declare的-a选项即可创建它们，或者在变量后面增加下标操作符直接创建。和很多其它开发语言一样，Shell中的数组也是0开始的，然而不同的是Shell中数组的下标是可以不连续的。获取数组中某个元素的语法格式为: ${arrayname[index]}。见如下示例：
    /> declare -a friends                   #声明一个数组变量
    /> friends=(sheryl peter louise)   #给数组变量赋值
    /> echo ${friends[0]}                #通过数组下标的方式访问数组的元素
    sheryl
    /> echo ${friends[1]}
    peter
    /> echo ${friends[2]}
    louise
    /> echo ${friends[*]}                #下标中星号表示所有元素。
    shery1 peter louise
    # ${#array[*]}表示数组中元素的数量，而${#friend[0]}则表示第一个元素的长度。
    /> echo ${#friends[*]}    
    3
    /> unset friends                         #unset array清空整个数组，unset array[index]仅清空指定下标的元素。

    /> x[3]=100
    /> echo ${x[*]}
    100
    /> echo ${x[0]}                        #0下标的元素并没有被赋值过，因为输出为空。

    /> echo ${x[3]}
    100
    /> states=(ME [3]=CA [2]=CT)   #ME的下标为0。
    /> echo ${states[0]}
    ME
    /> echo ${states[1]}                 #数组下标为1的位置没有被赋值过，因此没有输出。

    /> echo ${states[2]}
    CT
    /> echo ${states[3]}
    CA

    12.  函数:
    和C语言一样，Shell中也可以创建自己的自定义函数。其格式如下：
    function_name () { commands; commands; }
    function function_name { commands; commands; }
    function function_name () { commands; commands; }
    函数的参数在函数内是以$[0-9]、${10}...，这种局部变量的方式来访问的。见下面的示例：

    #函数的左花括号和命令之间必须有至少一个空格。每个命令的后面都要有一个分号，即便是最后一个命令
    /> function greet { echo "Hello $LOGNAME, today is $(date)"; }
    #此时函数已经驻留在当前的bash shell中，因此使用函数效率更高。
    /> greet   
    Hello root, today is Fri Nov 18 20:45:10 CST 2011
    /> greet() { echo "Hello $LOGNAME, today is $(date)"; }
    /> greet
    Hello root, today is Fri Nov 18 20:46:40 CST 2011
    #welcome函数内部使用了函数参数。
    /> function welcome { echo "Hi $1 and $2"; }
    /> welcome stephen jane    
    Hi stephen and jane
    #declare -F选项将列出当前Shell中驻留的函数
    /> declare -F
    declare -f greet
    declare -f welcome
    #清空指定的函数，使其不在Shell中驻留。
    /> unset -f welcome

    13.  重定向:
    下面的列表为Shell中支持的重新定向操作符。

操作符	功能
<	重新定向输入
>	重新定向输出
>>	追加输出
2>	重新定向错误
&>	重新定向错误和输出
>&	重新定向错误和输出
2>&1	重新定向错误到标准输出
1>&2	重新定向标准输出到错误
>|	重新定向输出的时候覆盖noclobber选项

    #find命令将搜索结果输出到foundit文件，把错误信息输出到/dev/null
    /> find . -name "*.c" -print > foundit 2> /dev/null
    #将find命令的搜索结果和错误信息均输出到foundit文件中。
    /> find . -name "*.c" -print >& foundit
    #同上。
    /> find . -name "*.c" -print > foundit 2>&1
    #echo命令先将错误输出到errfile，再把信息发送到标准错误，该信息标准错误与标准输出合并在一起(errfile中)。
    /> echo "File needs an argument" 2> errfile 1>&2
    /> cat errfile
    File needs an argument

## 二十三. Bash Shell编程:

    1.  读取用户变量:
    read命令是用于从终端或者文件中读取输入的内建命令，read命令读取整行输入，每行末尾的换行符不被读入。在read命令后面，如果没有指定变量名，读取的数据将被自动赋值给特定的变量REPLY。下面的列表给出了read命令的常用方式：

命令格式	描述
read answer	从标准输入读取输入并赋值给变量answer。
read first last	从标准输入读取输入到第一个空格或者回车，将输入的第一个单词放到变量first中，并将该行其他的输入放在变量last中。
read	从标准输入读取一行并赋值给特定变量REPLY。
read -a arrayname	把单词清单读入arrayname的数组里。
read -p prompt	打印提示，等待输入，并将输入存储在REPLY中。
read -r line	允许输入包含反斜杠。

    见下面的示例(绿色高亮部分的文本为控制台手工输入信息)：
    /> read answer        #等待读取输入，直到回车后表示输入完毕，并将输入赋值给变量answer
    Hello                       #控制台输入Hello
    /> echo $answer      #打印变量
    Hello

    #等待一组输入，每个单词之间使用空格隔开，直到回车结束，并分别将单词依次赋值给这三个读入变量。
    /> read one two three
    1 2 3                      #在控制台输入1 2 3，它们之间用空格隔开。
    /> echo "one = $one, two = $two, three = $three"
    one = 1, two = 2, three = 3

    /> read                  #等待控制台输入，并将结果赋值给特定内置变量REPLY。
    This is REPLY          #在控制台输入该行。
    /> echo $REPLY      #打印输出特定内置变量REPLY，以确认是否被正确赋值。
    This is REPLY

    /> read -p "Enter your name: "    #输出"Enter your name: "文本提示，同时等待输入，并将结果赋值给REPLY。
    Enter you name: stephen            #在提示文本之后输入stephen
    /> echo $REPLY
    stephen

    #等待控制台输入，并将输入信息视为数组，赋值给数组变量friends，输入信息用空格隔开数组的每个元素
    /> read -a friends
    Tim Tom Helen
    /> echo "I have ${#friends} friends"
    I have 3 friends
    /> echo "They are ${friends[0]}, ${friends[1]} and ${friends[2]}."
    They are Tim, Tom and Helen.

   2.  状态判断:
    test是Shell中提供的内置命令，主要用于状态的检验，如果结果为0，表示成功，否则表示失败。见如下示例：
    /> name=stephen
    /> test $name != stephen
    /> echo $?
    1
    需要注意的是test命令不支持Shell中提供的各种通配符，如：
    /> test $name = [Ss]tephen
    /> echo $?
    1
    test命令还可以中括号予以替换，其语义保持不变，如：
    /> [ $name = stephen ]
    /> echo $?
    0    
    在Shell中还提供了另外一种用于状态判断的方式：[[ expr ]]，和test不同的是，该方式中的表达式支持通配符，如：
    /> name=stephen
    /> [[ $name == [Ss]tephen ]]
    /> echo $?
    0
    #在[[ expression ]]中，expression可以包含&&(逻辑与)和||(逻辑或)。
    /> [[ $name == [Ss]tephen && $friend == "Jose" ]]
    /> echo $?
    1
    /> shopt -s extglob   #打开Shell的扩展匹配模式。
    /> name=Tommy
    # "[Tt]o+(m)y"的含义为，以T或t开头，后面跟着一个o，再跟着一个或者多个m，最后以一个y结尾。
    /> [[ $name == [Tt]o+(m)y ]] 
    /> echo $?
    0
    在Shell中还提供了let命令的判断方式： (( expr ))，该方式的expr部分，和C语言提供的表达式规则一致，如：
    /> x=2
    /> y=3
    /> (( x > 2 ))
    /> echo $?
    1
    /> (( x < 2 ))
    /> echo $?
    0　　
    /> (( x == 2 && y == 3 ))
    /> echo $?
    0
    /> (( x > 2 || y < 3 ))
    /> echo $?
    1

    下面的表格是test命令支持的操作符：

判断操作符	判断为真的条件
字符串判断	 
[ stringA=stringB ]	stringA等于stringB
[ stringA==stringB ]	stringA等于stringB
[ stringA!=stringB ]	stringA不等于stringB
[ string ]	string不为空
[ -z string ]	string长度为0
[ -n string ]	string长度不为0
逻辑判断	 
[ stringA -a stringB ]	stringA和stringB都是真
[ stringA -o stringB ]	stringA或stringB是真
[ !string ]	string不为真
逻辑判断(复合判断)	 
[[ pattern1 && pattern2 ]]	pattern1和pattern2都是真
[[ pattern1 || pattern2 ]	pattern1或pattern2是真
[[ !pattern ]]	pattern不为真
整数判断	 
[ intA -eq intB ]	intA等于intB
[ intA -ne intB ]	intA不等于intB
[ intA -gt intB ]	intA大于intB
[ intA -ge intB ]	intA大于等于intB
[ intA -lt intB ]	intA小于intB
[ intA -le intB ]	intA小于等于intB
文件判断中的二进制操作	 
[ fileA -nt fileB ]	fileA比fileB新
[ fileA -ot fileB ]	fileA比fileB旧
[ fileA -ef fileB ]	fileA和fileB有相同的设备或者inode值
文件检验	 
[ -d $file ] or [[ -d $file ]]	file为目录且存在时为真
[ -e $file ] or [[ -e $file ]]	file为文件且存在时为真
[ -f $file ] or [[ -f $file ]]	file为非目录普通文件存在时为真
[ -s $file ] or [[ -s $file ]]	file文件存在, 且长度不为0时为真
[ -L $file ] or [[ -L $file ]]	file为链接符且存在时为真
[ -r $file ] or [[ -r $file ]]	file文件存在且可读时为真
[ -w $file ] or [[ -w $file ]]	file文件存在且可写时为真
[ -x $file ] or [[ -x $file ]]	file文件存在且可执行时为真
    注：在逻辑判断(复合判读中)，pattern可以包含元字符，在字符串的判断中，pattern2必须被包含在引号中。

    let命令支持的操作符和C语言中支持的操作符完全相同，如：
    +,-,*,/,%            加，减，乘，除，去模
    >>,<<                右移和左移
    >=,<=,==,!=      大于等于，小于等于，等于，不等于
    &,|,^                  按位与，或，非
    &&,||,!                逻辑与，逻辑或和取反
    还有其含义和C语言等同的快捷操作符，如=,*=,/=,%=,+=,-=,<<=,>>=,&=,|=,^=。

    3.  流程控制语句:
    if语句格式如下：
    #if语句的后面是Shell命令，如果该命令执行成功返回0，则执行then后面的命令。
    if command        
    then
        command
        command
    fi
    #用test命令测试其后面expression的结果，如果为真，则执行then后面的命令。
    if test expression
    then
        command
    fi
    #下面的格式和test expression等同
    if [ string/numeric expression ]
    then
        command
    fi
    #下面的两种格式也可以用于判断语句的条件表达式，而且它们也是目前比较常用的两种。
    if [[ string expression ]]
    then
        command
    fi

    if (( numeric expression ))           #let表达式
    then
        command
    fi
    见如下示例：
    /> cat > test1.sh                       #从命令行直接编辑test1.sh文件。
    echo -e "Are you OK(y/n)? \c"
    read answer
    #这里的$answer变量必须要用双引号扩住，否则判断将失败。当变量$answer等于y或Y时，支持下面的echo命令。
    if [ "$answer" = y -o "$answer" = Y ]   
    then
        echo "Glad to see it."
    fi
    CTRL+D  
    /> . ./test1.sh
    Are you OK(y/n)? y
    Glad to see it.
    上面的判断还可以替换为：
    /> cat > test2.sh
    echo -e "Are you OK(y/n or Maybe)? \c"
    read answer
    # [[ ]]复合命令操作符允许其中的表达式包含元字符，这里输入以y或Y开头的任意单词，或Maybe都执行then后面的echo。
    if [[ $answer == [yY]* || $answer = Maybe ]]  
    then
        echo "Glad to hear it.
    fi
    CTRL+D
    /> . ./test2.sh
    Are you OK(y/n or Maybe)? yes
    Glad to hear it.
    下面的例子将使用Shell中的扩展通配模式。
    /> shopt -s extglob        #打开该扩展模式
    /> answer="not really"
    /> if [[ $answer = [Nn]o?( way |t really) ]]
    > then
    >    echo "I am sorry."
    > fi
    I am sorry.
    对于本示例中的扩展通配符，这里需要给出一个具体的解释。[Nn]o匹配No或no，?( way|t really)则表示0个或1个( way或t really),因此answer变量匹配的字符串为No、no、Not really、not really、No way、no way。
    下面的示例使用了let命令操作符，如：
    /> cat > test3.sh
    if (( $# != 2 ))                    #等同于 [ $# -ne 2 ]
    then
        echo "Usage: $0 arg1 arg2" 1>&2
        exit 1                         #exit退出值为0-255之间，只有0表示成功。
    fi
    if (( $1 < 0 || $1 > 30 ))      #等同于 [ $1 -lt 0 -o $1 -gt 30 ]
    then
        echo "arg1 is out of range."
        exit 2
    fi
    if (( $2 <= 20 ))                  #等同于 [ $2 -le 20 ]
    then
        echo "arg2 is out of range."
    fi
    CTRL+D
    /> sh ./test3.sh
    Usage: ./test3.sh arg1 arg2
    /> echo $?                          #Shell脚本的退出值为exit的参数值。
    1
    /> sh ./test3.sh 40 30
    arg1 is out of range.
    /> echo $?
    2
    下面的示例为如何在if的条件表达式中检验空变量：
    /> cat > test4.sh
    if [ "$name" = "" ]                #双引号就表示空字符串。
    then
        echo "name is null."
    fi
    CTRL+D
    /> . ./test4.sh
    name is null.

    if/elif/else语句的使用方式和if语句极为相似，相信有编程经验的人都不会陌生，这里就不在赘述了，其格式如下：
    if command
    then 
        command
    elif command
    then
        command
    else
        command
    fi
    见如下示例脚本：
    /> cat > test5.sh
    echo -e "How old are you? \c"
    read age
    if [ $age -lt 0 -o $age -gt 120 ]                #等同于 (( age < 0 || age > 120 ))
    then
        echo "You are so old."
    elif [ $age -ge 0 -a $age -le 12 ]               #等同于 (( age >= 0 && age <= 12 ))
    then
        echo "You are child."
    elif [ $age -ge 13 -a $age -le 19 ]             #等同于 (( age >= 13 && age <= 19 ))
    then
        echo "You are 13--19 years old."
    elif [ $age -ge 20 -a $age -le 29 ]             #等同于 (( age >= 20 && age <= 29 ))
    then
        echo "You are 20--29 years old."
    elif [ $age -ge 30 -a $age -le 39 ]             #等同于 (( age >= 30 && age <= 39 ))
    then
        echo "You are 30--39 years old."
    else
        echo "You are above 40."
    fi
    CTRL+D
    /> . ./test5.sh
    How old are you? 50
    You are above 40.

    case语句格式如下：
    case variable in
    value1)
        command
        ;;            #相同于C语言中case语句内的break。
    value2)
        command
        ;;
    *)                #相同于C语言中switch语句内的default
       command
        ;;
    esac
    见如下示例脚本：
    /> cat > test6.sh
    #!/bin/sh
    echo -n "Choose a color: "
    read color
    case "$color" in
    [Bb]l??)
        echo "you select blue color."
        ;;
    [Gg]ree*)
        echo "you select green color."
        ;;
    red|orange)
        echo "you select red or orange."
        ;;
    *)
        echo "you select other color."
        ;;
    esac
    echo "Out of case command."
    /> . ./test6.sh
    Choose a color: green
    you select green color.
    Out of case command.

   4.  循环语句:
    Bash Shell中主要提供了三种循环方式：for、while和until。
    for循环声明格式：
    for variable in word_list
    do
        command
    done
    见如下示例脚本：
    /> cat > test7.sh
    for score in math english physics chemist   #for将循环读取in后面的单词列表，类似于Java的for-each。
    do
        echo "score = $score"
    done
    echo "out of for loop"
    CTRL+D
    /> . ./test7.sh
    score = math
    score = english
    score = physics
    score = chemist
    out of for loop

    /> cat > mylist   #构造数据文件
    tom
    patty
    ann
    jake
    CTRL+D
    /> cat > test8.sh
    #!/bin/sh
    for person in $(cat mylist)                 #for将循环读取cat mylist命令的执行结果。
    do
        echo "person = $person"
    done
    echo "out of for loop."
    CTRL+D
    /> . ./test8.sh
    person = tom
    person = patty
    person = ann
    person = jake
    out of for loop.

    /> cat > test9.sh
    for file in test[1-8].sh                        #for将读取test1-test8，后缀为.sh的文件
    do
        if [ -f $file ]                              #判断文件在当前目录是否存在。
        then
            echo "$file exists."
        fi
    done
    CTRL+D
    /> . ./test9.sh
    test2.sh exists.
    test3.sh exists.
    test4.sh exists.
    test5.sh exists.
    test6.sh exists.
    test7.sh exists.
    test8.sh exists.

    /> cat > test10.sh
    for name in $*                                  #读取脚本的命令行参数数组，还可以写成for name的简化形式。
    do
        echo "Hi, $name"
    done
    CTRL+D
    /> . ./test10.sh stephen ann
    Hi, stephen
    Hi, ann

    while循环声明格式：
    while command  #如果command命令的执行结果为0，或条件判断为真时，执行循环体内的命令。
    do
        command
    done
    见如下示例脚本：
    /> cat > test1.sh  
    num=0
    while (( num < 10 ))               #等同于 [ $num -lt 10 ]
    do
        echo -n "$num "
        let num+=1
    done
    echo -e "\nHere's out of loop."
    CTRL+D
    /> . ./test1.sh
    0 1 2 3 4 5 6 7 8 9 
    Here's out of loop.

    /> cat > test2.sh
    go=start
    echo Type q to quit.
    while [[ -n $go ]]                     #等同于[ -n "$go" ]，如使用该风格，$go需要被双引号括起。
    do
        echo -n How are you.
        read word
        if [[ $word == [Qq] ]]      #等同于[ "$word" = Q -o "$word" = q ]
        then
            echo Bye.
            go=                        #将go变量的值置空。
        fi
    done
    CTRL+D
    /> . ./test2.sh
    How are you. Hi
    How are you. q
    Bye.

    until循环声明格式:
    until command                         #其判断条件和while正好相反，即command返回非0，或条件为假时执行循环体内的命令。
    do
        command
    done
    见如下示例脚本：
    /> cat > test3.sh
    until who | grep stephen           #循环体内的命令将被执行，直到stephen登录，即grep命令的返回值为0时才退出循环。
    do
        sleep 1
        echo "Stephen still doesn't login."
    done
    CTRL+D

    shift命令声明格式:shift [n]
    shift命令用来把脚本的位置参数列表向左移动指定的位数(n)，如果shift没有参数，则将参数列表向左移动一位。一旦移位发生，被移出列表的参数就被永远删除了。通常在while循环中，shift用来读取列表中的参数变量。
    见如下示例脚本：
    /> set stephen ann sheryl mark #设置4个参数变量。
    /> shift                                    #向左移动参数列表一次，将stephen移出参数列表。
    /> echo $*
    ann sheryl mark
    /> shift 2                                 #继续向左移动两位，将sheryl和ann移出参数列表
    /> echo $*
    mark
    /> shift 2                                 #继续向左移动两位，由于参数列表中只有mark了，因此本次移动失败。
    /> echo $*
    mark

    /> cat > test4.sh
    while (( $# > 0 ))                    #等同于 [ $# -gt 0 ]
    do
        echo $*
        shift
    done
    CTRL+D
    /> . ./test4.sh a b c d e
    a b c d e
    b c d e
    c d e
    d e
    e        

    break命令声明格式：break [n]
    和C语言不同的是，Shell中break命令携带一个参数，即可以指定退出循环的层数。如果没有指定，其行为和C语言一样，即退出最内层循环。如果指定循环的层数，则退出指定层数的循环体。如果有3层嵌套循环，其中最外层的为1，中间的为2，最里面的是3。
    见如下示例脚本:
    /> cat > test5.sh
    while true
    do
        echo -n "Are you ready to move on?"
        read answer
        if [[ $answer == [Yy] ]]
        then
            break
        else
            echo "Come on."
        fi
    done
    echo "Here we are."
    CTRL+D
    /> . ./test5.sh
    Are you ready to move on? y
    Here we are

    continue命令声明格式：continue [n]
    和C语言不同的是，Shell中continue命令携带一个参数，即可以跳转到指定层级的循环顶部。如果没有指定，其行为和C语言一样，即跳转到最内层循环的顶部。如果指定循环的层数，则跳转到指定层级循环的顶部。如果有3层嵌套循环，其中最外层的为3，中间的为2，最里面的是1。
    /> cat  maillist                       #测试数据文件maillist的内容为以下信息。
    stephen
    ann
    sheryl
    mark

    /> cat > test6.sh
    for name in $(cat maillist)
    do
        if [[ $name == stephen ]]; then
            continue
        else
            echo "Hello, $name."
        fi
    done
    CTRL+D
    /> . ./test6.sh
    Hello, ann.
    Hello, sheryl.
    Hello, mark.

    I/O重新定向和子Shell：
    文件中的输入可以通过管道重新定向给一个循环，输出也可以通过管道重新定向给一个文件。Shell启动一个子Shell来处理I/O重新定向和管道。在循环终止时，循环内部定义的任何变量对于脚本的其他部分来说都是不看见的。
    /> cat > demodata                        #为下面的脚本构造册数数据
    abc
    def
    ghi
    CRTL+D
    /> cat > test7.sh
    if (( $# < 1 ))                                #如果脚本参数的数量小于1，则给出错误提示后退出。
    then
        echo "Usage: $0 filename " >&2
        exit 1
    fi
    count=1
    cat $1 | while read line                   #参数一中的文件被cat命令输出后，通过管道逐行输出给while read line。
    do
        let $((count == 1)) && echo "Processing file $1..." > /dev/tty  #该行的echo将输出到当前终端窗口。
        echo -e "$count\t$line"              #将输出行号和文件中该行的内容，中间用制表符隔开。
        let count+=1
    done > outfile                               #将while循环中所有的输出，除了>/dev/tty之外，其它的全部输出到outfile文件。
    CTRL+D
    /> . ./test7.sh demodata                #只有一行输出，其余的都输出到outfile中了。
    Processing file demodata...
    /> cat outfile
    1       abc
    2       def
    3       ghi

    /> cat > test8.sh
    for i in 9 7 2 3 5 4
    do
        echo $i
    done | sort -n                                #直接将echo的输出通过管道重定向sort命令。
    CTRL+D
    /> . ./test8.sh
    2
    3
    4
    5
    7
    9

    5.  IFS和循环：
    Shell的内部域分隔符可以是空格、制表符和换行符。它可以作为命令的分隔符用在例如read、set和for等命令中。如果在列表中使用不同的分隔符，用户可以自己定义这个符号。在修改之前将IFS原始符号的值保存在另外一个变量中，这样在需要的时候还可以还原。
    见如下示例脚本：
    /> cat > test9.sh
    names=Stephen:Ann:Sheryl:John   #names变量包含的值用冒号分隔。
    oldifs=$IFS                                   #保留原有IFS到oldifs变量，便于后面的还原。
    IFS=":"                            
    for friends in $names                     #这是遍历以冒号分隔的names变量值。    
    do
        echo Hi $friends
    done
    IFS=$oldifs                                   #将IFS还原为原有的值。
    set Jerry Tom Angela
    for classmates in $*                      #再以原有IFS的值变量参数列表。
    do
        echo Hello $classmates
    done
    CTRL+D
    /> . ./test9.sh
    Hi Stephen
    Hi Ann
    Hi Sheryl
    Hi John
    Hello Jerry
    Hello Tom
    Hello Angela

    6.  函数：
    Shell中函数的职能以及优势和C语言或其它开发语言基本相同，只是语法格式上的一些差异。下面是Shell中使用函数的一些基本规则：
    1) 函数在使用前必须定义。
    2) 函数在当前环境下运行，它和调用它的脚本共享变量，并通过位置参量传递参数。而该位置参量将仅限于该函数，不会影响到脚本的其它地方。
    3) 通过local函数可以在函数内建立本地变量，该变量在出了函数的作用域之后将不在有效。
    4) 函数中调用exit，也将退出整个脚本。
    5) 函数中的return命令返回函数中最后一个命令的退出状态或给定的参数值，该参数值的范围是0-256之间。如果没有return命令，函数将返回最后一个Shell的退出值。
    6) 如果函数保存在其它文件中，就必须通过source或dot命令把它们装入当前脚本。
    7) 函数可以递归。
    8) 将函数从Shell中清空需要执行：unset -f function_name。
    9) 将函数输出到子Shell需要执行：export -f function_name。
    10) 可以像捕捉Shell命令的返回值一样获取函数的返回值，如$(function_name)。
    Shell中函数的声明格式如下：
    function function_name { command; command; }
    见如下示例脚本：
    /> cat > test1.sh
    function increment() {            #定义函数increment。
        local sum                           #定义本地变量sum。
        let "sum=$1+1"    
        return $sum                      #返回值是sum的值。
    }
    echo -n "The num is "
    increment 5                          #increment函数调用。
    echo $?                                #输出increment函数的返回值。
    CTRL+D
    /> . ./test1.sh
    The num is 6

    7.  陷阱信号(trap)：
    在Shell程序运行的时候，可能收到各种信号，有的来自于操作系统，有的来自于键盘，而该Shell在收到信号后就立刻终止运行。但是在有些时候，你可能并不希望在信号到达时，程序就立刻停止运行并退出。而是他能希望忽略这个信号而一直在运行，或者在退出前作一些清除操作。trap命令就允许你控制你的程序在收到信号以后的行为。
    其格式如下：
    trap 'command; command' signal-number
    trap 'command; command' signal-name
    trap signal-number  
    trap signal-name
    后面的两种形式主要用于信号复位，即恢复处理该信号的缺省行为。还需要说明的是，如果trap后面的命令是使用单引号括起来的，那么该命令只有在捕获到指定信号时才被执行。如果是双引号，则是在trap设置时就可以执行变量和命令替换了。
    下面是系统给出的信号数字和信号名称的对照表：
    1)SIGHUP 2)SIGINT 3)SIGQUIT 4)SIGILL 5)SIGTRAP 6)SIGABRT 7)SIGBUS 8)SIGFPE
    9)SIGKILL 10) SIGUSR1 11)SIGEGV 12)SIGUSR2 13)SIGPIPE 14)SIGALRM 15)SIGTERM 17)SIGCHLD
    18)SIGCONT 19)SIGSTOP ... ...
    见如下示例脚本：
    /> trap 'rm tmp*;exit 1' 1 2 15      #该命令表示在收到信号1、2和15时，该脚本将先执行rm tmp*，然后exit 1退出脚本。
    /> trap 2                                      #当收到信号2时，将恢复为以前的动作，即退出。
    /> trap " " 1 2                              #当收到信号1和2时，将忽略这两个信号。
    /> trap -                                      #表示恢复所有信号处理的原始值。
    /> trap 'trap 2' 2                           #在第一次收到信号2时，执行trap 2，这时将信号2的处理恢复为缺省模式。在收到信号2时，Shell程序退出。
    /> cat > test2.sh
    trap 'echo "Control+C will not terminate $0."' 2   #捕获信号2，即在键盘上按CTRL+C。
    trap 'echo "Control+\ will not terminate $0."' 3   #捕获信号3，即在键盘上按CTRL+\。
    echo "Enter stop to quit shell."
    while true                                                        #无限循环。
    do
        echo -n "Go Go...."
        read
        if [[ $REPLY == [Ss]top ]]                            #直到输入stop或Stop才退出循环和脚本。
       then
            break
        fi
    done
    CTRL+D
    /> . ./test2.sh
    Enter stop to quit shell.
    Go Go....^CControl+C will not terminate -bash.
    ^\Control+\ will not terminate -bash.
    stop

    8.  用getopts处理命令行选项：
    这里的getopts命令和C语言中的getopt几乎是一致的，因为脚本的位置参量在有些时候是失效的，如ls -lrt等。这时候-ltr都会被保存在$1中，而我们实际需要的则是三个展开的选项，即-l、-r和-t。见如下带有getopts的示例脚本：
    /> cat > test3.sh
    #!/bin/sh
    while getopts xy options                           #x和y是合法的选项，并且将-x读入到变量options中，读入时会将x前面的横线去掉。
    do
        case $options in
        x) echo "you entered -x as an option" ;;       
        y) echo "you entered -y as an option" ;;
        esac
    done
    /> ./test3.sh -xy
    you entered -x as an option
    you entered -y as an option
    /> ./test3.sh -x
    you entered -x as an option
    /> ./test3.sh -b                                       #如果输入非法选项，getopts会把错误信息输出到标准错误。
    ./test3.sh: illegal option -- b
    /> ./test3.sh b                                        #该命令不会有执行结果，因为b的前面有没横线，因此是非法选项，将会导致getopts停止处理并退出。

    /> cat > test4.sh
    #!/bin/sh
    while getopts xy options 2>/dev/null         #如果再出现选项错误的情况，该重定向会将错误输出到/dev/null。
    do
        case $options in
        x) echo "you entered -x as an option" ;; 
        y) echo "you entered -y as an option" ;;
        \?) echo "Only -x and -y are valid options" 1>&2 # ?表示所有错误的选项，即非-x和-y的选项。
    esac
    done
    /> . ./test4.sh -g                                     #遇到错误的选项将直接执行\?)内的代码。
    Only -x and -y are valid options
    /> . ./test4.sh -xg
    you entered -x as an option
    Only -x and -y are valid options

    /> cat > test5.sh
    #!/bin/sh
    while getopts xyz: arguments 2>/dev/null #z选项后面的冒号用于提示getopts，z选项后面必须有一个参数。
    do
        case $arguments in
        x) echo "you entered -x as an option." ;;
        y) echo "you entered -y as an option." ;;
        z) echo "you entered -z as an option."  #z的后面会紧跟一个参数，该参数保存在内置变量OPTARG中。
            echo "\$OPTARG is $OPTARG.";
           ;;
        \?) echo "Usage opts4 [-xy] [-z argument]"
            exit 1 ;;
        esac
    done
    echo "The number of arguments passed was $(( $OPTIND - 1 ))" #OPTIND保存一下将被处理的选项的位置，他是永远比实际命令行参数多1的数。
    /> ./test5.sh -xyz foo
    you entered -x as an option.
    you entered -y as an option.
    you entered -z as an option.
    $OPTARG is foo.
    The number of arguments passed was 2
    /> ./test5.sh -x -y -z boo
    you entered -x as an option.
    you entered -y as an option.
    you entered -z as an option.
    $OPTARG is boo.
    The number of arguments passed was 4

    9.  eval命令与命令行解析：
    eval命令可以对命令行求值，做Shell替换，并执行命令行，通常在普通命令行解析不能满足要求时使用。
    /> set a b c d
    /> echo The last argument is \$$#
    The last argument is $4
    /> eval echo The last argument is \$$#    #eval命令先进行了变量替换，之后再执行echo命令。
    The last argument is d



