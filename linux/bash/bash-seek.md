# 速查表


零、shell中的内部变量:


1.    $?:    表示shell命令的返回值.
2.    $$:    表示当前shell的pid.
3.    $!:    最后一个放入后台作业的PID值.
4.    $0:    表示脚本的名字.
5.    $1--$9,${10}: 表示脚本的第一到九个参数,和第十个参数.
6.    $#:    表示参数的个数.
7.    $*,$@: 表示所有的参数. 

    两者的区别如下: //都是双引号惹的祸^-^
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
    /> for i in "$*"
    >  do
    >  echo $i
    >  done
    apple pie pears    peaches
    
    /> set 'apple pie' pears peaches
    /> for i in "$@"
    >  do
    >  echo $i
    >  done
    apple pie    //这里的单引号将两个单词合成一个.
    pears
    peaches

 

一、正则表达式在vi中的用法：


1.    ^:      如/^love,表示所有以love开头的行.
2.    $:      如/love$,表示所有以love结尾的行.
3.    .:    如/l..e, dot表示任意字符，如love,l22e,live等.
4.    *:      如/*love, *表示0多多个字符,这里表示love前面可以有0个多任意多个空格字符，如/go*gle,可以表示ggle,gogle,google,goooooooogle.
5.    []:     如/[Ll]ove,[]中的任意一个字符都可能成为候选者,如Love和love.
6.    [x-y]: 如/[A-Z]t, 表示[]中指定范围内的字符都可能成为候选者,如At, It等, 也可表示多个区间段如：[a-zA-TV-Z]表示所有除V之外的所有大小写英文字符.
7.    [^]:   如/[^A-Z]ove,表示A-Z之内的任意字符都是非法的, 如Love,Dove等.
8.    \:      转义符,    如果想表示任何meta字符的原义, 需使用在meta字符前加转义符\, 如\.将只表示dot，而不能在表示任何其他字符了.
9.    \<:    如/\<love, 表示任何单词的开始, 如love和lover, 但是glove将非法.
10.  \>:   如/love\>, 表示任何单词的结束, 如love和glove, 但是lover将非法.
11.  \(..\):      如/\(love\)able/\1rs/, 这里的\1表示love, 这种标签替代最多达到\9, 该例子表示用lovers代替loveable.
12.  x\{m\}:   如x\{5\}, 表示x被重复5次,如xxxxx.
13.  x\{m,\}:  如x\{5,\}, 表示x被至少重复5次,如xxxxx,xxxxxxxx.
14.  x\{m,n\}:如x\{5,10\}, 表示x被重复5-10次,如xxxxx,xxxxxxxx.
以下为grep的正则表示式用法:
15.  \w和\W: 等同于[a-zA-Z0-9].
16.  \b: 等同于\<和\>,均表示单词的边界.
以下为grep的正则表示式的扩展用法(grep -E或egrep):
17.  +:   如/lo+ve, +表示1个或者多个先前的字符,这里表示love,loove,但是lve非法.
18.  ?:   如/lo?ve, ?表示0个或者1个先前的字符, 这里只表示love和lve.
19.  (a|b|c):  如/l(o|i)ve, 表示或的意思,这里表示love和live. (o|i)和[oi]的主要区别就是(word|word)可以表示单词之间或的关系,[]只能表示字符.
20.  x{m},x{m,},x{m,n}  等同于grep普通模式中的x\{m\},x\{m,\},x\{m,n\}.

 

二、grep家族:


1.    家族成员:
    egrep: 执行带有扩展正则表达式元字符的grep搜索.
    fgrep:  将关闭grep的所有正则功能, 即搜索字符串中所有正则元字符都将只是表示其字符本意.
2.    返回值: 
    0: 表示成功
    1: 表示搜索字符串不存在
    2: 表示搜索文件不存在.    
3.    grep的选项规则:
    -#,-A#和-B#: 表示在输出匹配内容的时候同时也输出其上下指定数量的行数, 如grep -2 "love" *, 该例输出匹配love的上下两行, 
    grep -A2 "love" * 该例输出匹配love的后两行, grep -B2 "love" * 该例输出匹配love的前两行. 这里A表示after,B表示before.
    -F: 等同于fgrep, 这个选项将关闭所有正则功能,即所有正则的元字符均表示其本身含义.
    -c: 不输出找到的内容,只是输出在该文件中有多少匹配的行数.
    -h: 不输出匹配搜索字符串的文件的文件名,只是输出内容.
    -i:  搜索时忽略大小写.
    -l:  只显示匹配搜索内容的文件名, 不显示具体的内容.
    -L: 只显示没有包含搜索内容的文件名.
    -n: 输出匹配内容的同时也输出其所在的行号.
    -v: 反向搜索,输出不匹配搜索字符串的行.
    -w:只打印以完整单词形式匹配的行, 如果该搜索字符为某个单词的部分内容,将不会被输出.
    -x: 只打印以行形式匹配的行, 如果该搜索字符为行的部分内容,将不会被输出.
    -q: 不会输出任何信息, 该选项主要用于测试某个搜索字符或搜索pattern在执行grep命令之后的返回值.
    -r: 表示递归的搜索当前目录的子目录中的文件.  
4.    对于普通模式的grep,如果搜索的字符中普通字符前面加入\,则该字符按照扩展grep(egrep或者grep -E)的正则规则进行查找.如grep "love\|live" filename, 
    将等同于egrep "love|live" filename,这里的\|将按照egrep中的|元字符处理, 再如, egrep "3+" filename等同于grep "3\+" filename.

 

三、sed:


1.    sed命令:
    ,:  表示范围.
    1) sed -n '/west/,/east/p' datafile 表示打印所有从包含west开始到包含east的行,如果直到文件的结尾都没有包含east的行,将打印west后面的所有行.
        其实逻辑很简单, 就是sed在发现包含west行之后开发打印该行,直到发现包含east的行打印才结束,否则一直打印直到文件的末尾.
    2) sed -n '5,/^northeast/p' datafile 表示从第五行开始打印,直到遇到以northeast开始的行结束打印.
     
    !:  表示对匹配结果取反.
    1) sed '/north/!d' datafile 将删除所有不包含north的行.

    a: 追加命令.
    1) sed '/^north/a first line \
        second line \
        third line' datafile 将会在所有包含north行的后面追加first line \r second line \n third line. 其中\表示下一行还有内容的连词. 如果是c-shell:
        sed '/^north/a first line \\
        second line \\
        third line' datafile 其中多出来的\是转义符.
     
    d: 表示删除.
    1) sed '/north/d' datafile 将删除所有包含north的行.
    2) sed '3d' datafile    将删除第三行.
    3) sed '3,$d' datafile    将删除第三行到文件的结尾行.
    4) sed 'd' datafile 将删除所有行.
     
    e: 表示多点编辑.
    1) sed -e '1,3d' -e 's/Hemenway/Jones/' datafile    一个sed语句执行多条编辑命令, 因此命令的顺序会影响其最终结果.
    2) sed -e 's/Hemenway/Jones/' -e 's/Jones/Max/' datafile 先用Jones替换Hemenway, 再用Max替换Jones.

    h和g/G: 保持和获取命令.
    1) sed -e '/northeast/h' -e '$G' datafile sed将把所有包含northeast的行轮流缓存到其内部缓冲区, 最后将只是保留最后一个匹配的行, 
        $G是将缓冲区的行输出到$G匹配行的后面, 该例表示将最后一个包含northeast的行追加到文件的末尾.
    2) sed -e '/WE/{h; d;}' -e '/CT/{G;}' datafile 表示将包含WE的行保存到缓冲区, 然后删除该行,最后将缓冲区中保存的那份输出到CT行的后面.
    3) sed -e '/northeast/h' -e '$g' datafile 表示将包含northeast的行保存到缓冲区, 再将缓冲区中保存的那份替换文件的最后一行并输出.
        再与h合用时, g表示替换, G表示追加到匹配行后面.
    4) sed -e '/WE/{h; d;}' -e '/CT/{g;}' datafile 保留包含WE的行到缓冲区, 如果有新的匹配行出现将会替换上一个存在缓冲区中的行, 如果此时发现有
        包含CT的行出现, 就用缓冲区中的当前行替换这个匹配CT的行, 之后如果有新的WE出现, 将会用该新行替换缓冲区中数据, 当前再次遇到CT的时候,将用最
        新的缓冲区数据替换该CT行.
    
    i: 表示插入.
    1) sed '/north/i first line \
        second line \
        third line' datafile    其规则和a命令基本相同, 只是a是将额外的信息输出到匹配行的后面, i是将额外信息输出到匹配行的前面.
     
    p: 表示打印.
    1) sed '/north/p' datafile 将打印所有包含north的行.
    2) sed '3p' datafile    将打印第三行.
    3) sed '3,$p' datafile    将打印第三行到文件的结尾行.
    4) sed 'p' datafile 将打印所有行.
    注: 使用p的时候sed将会输出指定打印的行和所有行, 当其与-n选项组合时候,将只是打印输出匹配的行.
     
    n: 下一行命令.
    1) sed '/north/ {n; s/Chin/Joseph/}' datafile 将先定位包含north的行, 然后取其下一行作为目标行, 再在该目标行上执行s/Chin/Joseph/的替换操作.
    2) sed '/north/ {n; n; s/Chin/Joseph/}' datafile 将取north包含行的后两行作为目标行.
    注: {}作为嵌入的脚本执行.
     
    q: 退出命令.
    1) sed '5q' datafile 到第五行退出(输出第五行).
    2) sed '/north/q' datafile 输出到包含north的行退出(输出包含north的行).
    3) sed '/Lewis/ {s/Lewis/Joseph/; q}' datafile 将先定位包含Lewis的行, 然后用Joseph替换Lewis,最后退出sed操作.
     
    r: 文件读入.
    1) sed '/Suan/r newfile' datafile    在输出时,将newfile的文件内容跟随在datafile中包含Suan的行后面输出,如果多行都包含Suan,则文件被多次输出.
    
    s: 表示替换.
    1) sed 's/west/north/g' datafile    将所有west替换为north, g表示如果一行之内多次出现west,将全部替换, 如果没有g命令,将只是替换该行的第一个匹配.
    2) sed -n 's/^west/north/p' datafile    将所有以west开头的行替换为north, 同时只是输出替换匹配的行.
    3) sed -n '1,5 s/\(Mar\)got/\1ianne/p' datafile    将从第一行到第五行中所有的Margot替换为Marianne, \1是\(Mar\)的变量替代符.
     
    w: 文件写入.
    1) sed -n '/north/w newfile2' datafile    将datafile中所有包含north的行都写入到newfile2中.
     
    x: 互换命令.
    1) sed -e '/pat/h' -e '/Margot/x' datafile x命令表示当定位到包含Margot行,互换缓冲区和该匹配Margot行的数据, 即缓冲区中的数据替换该匹配行显示, 
        该匹配行进入缓冲区, 如果在交换时缓冲区是空, 则该匹配行被换入缓冲区, 空行将替换该行显示, 后面依此类推. 如果交换后, 再次出现匹配pat的行, 该
        行将仍然会按照h命令的规则替换(不是交换, 交换只是发生在发现匹配Margot的时候)缓冲区中的数据.
    
    y: 变形命令.
    1) sed '1,3y/abcd/ABCD/' datafile 将1到3行中的小写abcd对应者替换为ABCD,注意abcd和ABCD是一一对应的. 如果他们的长度不匹配,sed将报错.
    2) sed 'y/abcdefghijklmnopqrstuvwxyz/ABCDEFGHIJKLMNOPQRSTUVWXYZ/' datafile 将datafile中所有的小写字符替换为大写字母.


四、awk家族:


1.    执行方式:

      1) awk 'pattern' filename 如awk '/Mary/' employees
      2) awk '{action}' filename 如awk '{print $1}' employees
      3) awk 'pattern {action}' filename 如awk '/Mary/ {print $1}' employees 
      注: 模式/Mary/对action的作用范围是从其后面的第一个左花括号开始,到第一个右花括号结束. 其后的pattern将不会影响前面的action.
    
2.    内置变量:

    $0:    表示一整行(相当于数据库中一条记录).
    NR:    当前行号.
    NF:    当前记录的域(相当于数据库中的字段)数量
    RS:    行分隔符(缺省为回车).
    FS:    域分隔符,缺省为\t. awk -F: '{print $1,$2,$3}' employees 这里FS等于":".
    OFS:输出域分隔符, awk  -F: '{print $1,$2,$3}' employees 这里OFS等于" "空格, 因为在$1和$2之间是空格分开的.
    ARGC: 命令行参数的数量.
    ARGV: 命令行参数数组.
    ENVIRON: 从shell传递来的包含当前环境变量的数组.
    ERRNO: 错误号.
    FILENAME: 当前的输入文件名.

3.    格式化输出:

转义码:

    \b:    Backspace.
    \n:    换行.
    \r:    回车.
    \t:    制表符.    
    
格式化说明符:

    %c:    单个ASCII字符.
    %d:    十进制数字.
    %e:    科学记数法表示的数字.
    %f:    浮点数.
    %o:    八进制数字.
    %s:    打印字符串.
    %x:    十六进制数字.
    -:    表示左对齐,如%-15d, 在十进制数字的后面会有一些空格,同时该数字是左对齐的. %+15d或%15d表示右对齐,当数字不足15位的时候.
    #:    如%#o或%#x, 会在八进制的数字前面加入0,十六进制前加0x.

4.    操作符:
    ~:    匹配运算符. 如awk '$1~/Mary/' employees, 表示第一个域($1)中包含Mary的被打印, 如果其他域包含,第一个域没有,则仍然视为无效.
    !~:    不匹配运算符. 如awk '$1!~/Mary/' employees, 表示第一个域($1)中不包含Mary的被打印, 如果其他域包含,第一个域没有,则仍然视为有效.
    <,>,<=,>=,!=,==: 关系运算符. awk '$3>5000 {print $3}' datafile
    cond ? expr1 : expr2 条件表达式 awk '{max = $1 > $2 ? $1 : $2; print max}' datafile
    =,+=,-=,*=,/=,%=: 赋值运算符.
    -,+,*,/,%,^(x^y[乘方]): 数学运算符.
    &&, ||, !: 逻辑运算符.
    ,: 表示范围, awk '/Tom/,/Mary/' datafile 其规则可参照sed中逗号运算符.
    
5.    选项:
    -F:    指定特定的分隔符,而不是缺省的\t, 如-F:,这里分隔符是":".    

6.    awk编程:
    1) BEGIN: 其后紧跟着动作块, 该块将会在任何输入文件被读入之前执行, 如一些初始化工作, 或者打印一些输出标题.
    awk 'BEGIN{FS=":"; OFS="\t";ORS="\n\n"} {print $1,$2,$3}' file
    即使输入文件不存在, BEGIN块动作仍然会被执行.
     
    2) END: 其后也紧随动作块, 该动作模块将在整个输入文件处理完毕之后被处理, 但是END需要有文件名的输入.
    awk 'END {print "The end\n"} filename.
     
    3) 输入输出重新定向:
    awk 'BEGIN {print "Hello" > "newfile"}' datafile 文件名一定要用双引号扩起来, > 如果文件存在,则清空后重写新文件.
    awk 'BEGIN {print "Hello" >> "newfile"}' datafile 文件名一定要用双引号扩起来, > 如果文件存在, 则在文件末尾追加写入.
    awk 'BEGIN {getline name < "/dev/tty"; print name}' getline是awk的内置函数, 就像c语言的gets, 将输入赋值给name变量.
     
    4) system函数可以执行shell中的命令,这些命令必须用双引号扩起.
    awk 'END { system("clear"); system ("cat " FILENAME)}' filename
     
    5) 条件语句:
    if (expr) { stat; } else { stat; }
    if (expr) { stat; } else if { stat; } else { stat; }
    awk '{ if ($7 <= 2) { print "less than 2", $7 } else if ($7 <= 4) { print "less than 4", $7 } else { print "the others", $7 } }' datafile
     
    6) 循环语句:
    while (expr) { stat; }
    for (i = 1; i <= NF; i++) { stat; }
    break;
    continue;
    exit(exitcode);    awk 将退出. 退出后的$?将会是这里的exitcode.
    next; 读取下一条记录. awk '{ if ($7 == 3) { next } else { print $0 }}' datafile 将不会输出$7等于3的记录.
     
    7) 数组:
    awk的数组和pl/sql中数组有些类似, 都是通过哈希表来实现的,其下标可以是数字, 也可以是字符串.
    awk '{name[x++]=$3};END{for(i = 0; i < NR; i++) { print i, name[i]}}' employees
    awk '{id[NR]=$3};END{for (x = 1; x <= NR; x++) { print id[x]} }' employees
    awk '/^Tom/{name[NR]=$1}; END{for (i in name) { print name[i]}}' employees 特殊的for语句
    awk '/Tom/{count["tom"]++}; /Mary/{count["mary"]++}; END{print "count[tom] = ",count["tom"]; print "count[mary] = ", count["mary"]}' employees
    awk '{count[$2]++};END{for (name in count) {print name,count[name]}}' datafile 域变量也可以作为数组的下标.

7.    内置函数:
    1) sub/gsub(regexp, substitution string, [target string]); gsub和sub的差别是sub只是替换每条记录中第一个匹配正则的, gsub则替换该记录中所有匹配
    正则的, 就是vi中s/src/dest/ 和s/src/dest/g的区别, 如果target string没有输入, 其缺省值是$0.
    awk '{sub(/Tom/,"Thomas"); print}' employees
    awk '{sub(/Tom/,"Thomas",$1); print}' employees
     
    awk '{gsub(/Tom/,"Thomas"); print}' employees
    awk '{gsub(/Tom/,"Thomas",$1); print}' employees
     
    2) index(string ,substring) 返回子字符串第一次被匹配的位置(1开始)
    awk 'BEGIN{print index("hollow", "low") }'
     
    3) length(string) 返回字符串的长度.
    awk 'BEGIN{print length("hello")}'
     
    4) substr(string, starting position, [length])
    awk 'BEGIN{print substr("Santa Claus",7,6)}'
    awk 'BEGIN{print substr("Santa Claus",7)}'
     
    5) match(string, regexp) 返回正则表示在string中的位置, 没有定位返回0
    awk 'BEGIN{print match("Good ole USA",/[A-Z]+$/)}'

    6) toupper(string)和tolower(string) 仅仅gawk有效.
    awk 'BEGIN{print toupper("linux"), tolower("BASH")}'

    7) split(string, array, [field seperator]) 如果不输入field seperator, FS内置变量作为其缺省值.
    awk 'BEGIN{split("12/24/99",date,"/"); for (i in date) {print date[i]} }'
     
    8) variable = sprintf(format, ...) 和printf的最大区别就是他返回格式化后的字符串.
    awk '{line = sprintf("%-15s %6.2f ",$5,$6); print line}' datafile
     
    9) systime() 返回1970/1/1到当前时间的整秒数.
     
    10) variable = strftime(format, [timestamp])
     
    11) 数学函数: atan2(x,y), cos(x), exp(x)[求幂], int(x)[求整数], log(x), rand()[随机数], sin(x), sqrt(x), srand(x)

五、BASH SHELL编程:


1.    初始化顺序: /etc/profile    ( ~/.bash_profile | ~/.bash_login | ~/.profile )    ~/.bashrc

2.    set -o allexport 当前shell变量对其所有子shell都有效.
    set +o allexport 当前shell变量对其所有子shell都无效.
    set -o noclobber 重定向输出时,如果输出文件已经存在则提示输出失败, date > out; date > out, 第二次操作失败
    set +o noclobber 缺省shell行为. date > out; date > out, 第二次操作成功
    shopt -s extglob 使用扩展通配符,如 abc?(2|9)K, abc*([0-9]), abc+([0-9]), no@(thing|body), no!(thing|body) 
    其中?,*,+,@和!都是用于修饰后面的()的.
    
3.    变量声明: declare, 在赋值的时候等号的两边不需要空格. variable=value.    
    declare -r variable=value 声明只读变量.
    declare -x variable=value 相当于export variable=value.
    数组声明:
    declare -a variable=(1 2 3 4)
    or
    name=(tom tim helen)
    or
    x[0]=5
    x[4]=10
    数组的声明可以不是连续的, 这一点和awk中的数组比较类似.
    e.g.
    /> declare -a friends
    /> friends=(sheryl peter louise)
    /> echo ${friends[0]}
    sheryl
    /> echo ${friends[1]}
    peter
    /> echo ${friends[2]}
    louise
    /> echo ${friends[*]}
    shery1 peter louise
    /> echo ${#friends[*]}
    3
    unset friends
    
    /> declare -a states=(ME [3]=CA [2]=CT)
    /> echo ${states[*]}
    ME CA CT
    /> echo ${#states[*]}
    3
    /> echo ${states[0]}
    ME
    /> echo ${states[1]}
    
    /> echo ${states[2]}
    CT
    /> echo ${states[3]}
    CA
    unset states
    
4.    函数声明:
    function greeting 
    {
        echo "Hi $1 and $2";
    }
    /> greeting tom joe
    Hi tom and joe
    unset -f greeting
    
5.    printf其参数类型类似于awk的printf.

6.    变量扩展修改符:
    ${variable:+word}
    if (NULL != variable)
        echo word
    else
        echo $variable
     
    e.g.
    /> unset var_name
    /> var_name=
    /> echo ${var_name:+AA}
      
    /> var_name=BB
    /> echo ${var_name:+AA}
    AA
    /> echo $var_name
    BB

    ${variable:-word}
    if (NULL == variable)
       echo word
    else
       echo $variable

    e.g.
    /> unset var_name
    /> var_name=
    /> echo ${var_name:-AA}
    /> AA
    /> var_name=BB
    /> echo ${var_name:-AA}
    BB
    /> echo $var_name
    BB
      
    ${variable:=word}
    if (NULL == variable)
    {
        variable=word
        echo word
    }
    else
    {
        echo $variable
    }
    
    e.g.
    /> unset var_name
    /> echo ${var_name:=AA}
    AA
    /> echo $var_name
    AA
    /> echo ${var_name:=CC}
    AA
    /> echo $var_name
    AA

    ${variable:offset}/${variable:offset:length}

    e.g.
    /> var_name=notebook
    /> echo ${var_name:0:4}
    /> note
    /> echo ${var_name:4:4}
    /> book
    /> echo ${var_name:2}
    /> tebook
    /> echo ${var_name:0}
    /> notebook
    
    ${variable%pattern} 从variable尾部开始,最小化的删除pattern
    e.g.
    /> variable="/usr/bin/local/bin"
    /> echo ${variable%/bin*}
    /usr/bin/local
    
    ${variable%%pattern} 从variable尾部开始,最大化的删除pattern
    /> variable="/usr/bin/local/bin"
    /> echo ${variable%%/bin*}
    /usr

    ${variable#pattern} 从variable头部开始,最小化的删除pattern
    /> variable="/home/lilliput/jake/.bashrc
    /> echo ${variable#/home}
    /lilliput/jake/.bashrc
    
    ${variable##pattern} 从variable头部开始,最大化的删除pattern
    /> variable="/home/lilliput/jake/.bashrc
    /> echo ${variable##*/}
    .bashrc
    
    ${#pattern} 返回patter的字符数量.
    /> variable=abc123
    6
    
7.    引用:
    \: 可以史shell中的元字符无效, 如?, < >, \, $, *, [ ], |, ( ), ;, &, { }
    ': 单引号可以史其内的所有元字符无效, 也包括\.
    ": 双引号也可以史其内的所有元字符无效, 变量和命令替换除外. 如 echo "What's time? $(date)", 这里date命令将被执行.
    
8.    命令替换:
    variable=$(date) or variable=`date`
    variable=`basename \`pwd\`` or variable=$(basename $(pwd))    命令替换可以嵌套, 第一种方法中,嵌套的命令必须使用\进行转义.
    eval: 可以进行命令行求值操作.
    e.g.
    /> set a b c d
    /> echo The last argument is \$$#
    /> The last argument is $4
    /> eval echo The last argument is \$$#
    /> The last argument is d
    
9.    数学计算:
    /> echo $[5+4-2]
    7
    /> echo $[5+2*3]
    11
    /> echo $((5+4-2))
    7
    /> echo $((5+2*3))
    11
    
    /> declare -i num //必须声明-i, 以表示整型变量.
    /> num=5+5
    /> echo $num
    10            //如果没有声明declare -i, 则返回5+5.
    /> num=5 + 5
    -bash: + 5: command not found.
    /> num="5 + 5"
    /> echo $num
    10
    /> num=4*6
    /> echo $num
    24
    
    使用不同进制(2~36)表示数字.
    /> declare -i x=017
    /> echo $x
    15
    /> x=2#101
    /> echo $x
    5
    /> x=8#17
    /> echo $x
    15
    /> x=16#b
    /> echo $x
    11
    
    let专门用于数学运算的bash内置命令.
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
    
10.    读取用户输入(read)命令:
    /> read answer
    yes
    /> echo "$answer is the right response."
    yes is the right reponse.
    
    /> read first middle last
    Jon Jake Jones
    /> echo "Hello $first"
    Hello Jon
    
    /> read         //如果没有变量时, $REPLY是缺省变量.
    the Chico Nut factory
    /> echo "I guess $REPLY keeps you busy!"
    I guess the Chico Nut factory keeps you busy!
    
    /> read -p "Enter your job titile: "
    Enter your job title: Accountant
    /> echo "I thought you might be an $REPLY".
    I thought you might be an accountant.
    
    /> read -a friends
    Melvin Tim Ernest
    /> echo "Say hi to ${friends[2]}
    Say hi to Ernest.
    /> echo "Say hi to ${friends[$[${#friends[*]}-1]]}"
    Say hi to Ernest.
    
11.    条件结构和流控制:
     1) 条件判断方法: test, [ ], [[ ]], (( )), 当$?为0时表示成功和true, 否则失败和false.
     /> name=tom
     /> test $name != tom
     /> echo $?
     1    //Failure
     
     /> [ $name=Tom ]
     /> echo $?
     0
     
     /> [ $name = [Tt]?? ]  //[]和test不允许使用通配符.
     /> echo $?
     1
     
     /> x=5
     /> y=20
     /> [ $x -gt $y ]
     /> echo $?
     1
     
     /> [ $x -le $y ]
     /> echo $?
     0
     
     /> name=Tom
     /> friend=Joseph
     /> [[ $name == [Tt]om ]]
     /> echo $?
     0
     
     /> [[ $name == [Tt]om && $friend == "Jose" ]]
     /> echo $?
     1
     
     /> x=2
     /> y=3
     /> (( x > 2))
     /> echo $?
     1
     
     /> (( x < 2))
     /> echo $?
     1
     
     /> (( x == 2 && y == 3))
     /> echo $?
     0
     
    2) test命令操作符:
    字符串判断:
    [ string1 = string2 ]    or    两个字符串相等时返回true.
    [ string1 == string2 ]     
          
    [ string1 != string2 ]      两个字符串不等时返回true.
    [ string ]                 string非空时返回true.
    [ -z string ]             为空时返回true.
    [ -n string ]             为非空时返回true.
     
    逻辑判断:(cond1可以包含元字符)
    [ string1 -a string2 ] or     string1和string2都非空时
    [[ cond1 && cond2 ]]     cond1和cond2都为true
         
    [ string1 -o string2 ] or     string1或string2为非空时
    [[ cond1 || cond2 ]]       cond1或cond2为true.
         
    [ ! string ] or              string为空
    [[ !cond ]]               cond为false.
              
    整数判断:
    [ int1 -eq int2 ]          int1等于int2
    [ int1 -ne int2 ]          int1不等于int2
    [ int1 -gt int2 ]           int1大于int2
    [ int1 -ge int2 ]          int1大于等于int2
    [ int1 -lt int2 ]            int1小于int2
    [ int1 -le int2 ]           int1小于等于int2
     
    文件判断
    [ file1 -nt file2 ]         file1比file2新
    [ file1 -ot file2 ]         file1比file2旧
    
    文件检验:
    [ -d $file ] or             表示判断目录是否存在
    [[ -d $file ]]
         
    [ -e $file ] or             表示判断文件是否存在
    [[ -e $file ]]
         
    [ -f $file ] or             表示判断非目录普通文件是否存在
    [[ -f $file ]]
         
    [ -s $file ] or            表示判断文件存在, 并且非0.
    [[ -s $file ]]

    [ -L $file ] or            表示判断为链接符号存在
    [[ -L $file ]]
         
    [ -r $file ] or            表示判断文件存在, 并且可读.
    [[ -r $file ]]
         
    [ -w $file ] or           表示判断文件存在, 并且可写.
    [[ -w $file ]]

    [ -x $file ] or           表示判断文件存在, 并且可执行.
    [[ -x $file ]]

    3) if语句:
    if command     //当command的返回状态为0时执行then后面的语句.
    then
         command
    fi
     
    if test expr
    then
        command
    fi
     
    if [ string/numeric expr ]
    then
        command
    fi
     
    一下两种为new format, 建议使用.
    if [[ string expr ]]     //支持通配符和扩展通配符(shopt -s extglob)
    then
        command
    fi
     
    if (( numeric expr ))
    then
        command
    fi
    
    if command
    then
        command1
    else
        command2
    fi
     
    if command
    then
        command1
    elif command
    then
        command2
    elif command
    then
        command3
    else
        command4
    fi
    
    4) 空语句:
    用冒号表示.
    if expr
    then
        :
    fi
     
    5) case语句:    //;;相当于c语言中break. 
    case variable in
    value1 | value2)     // value1支持通配符和|作为or的关系
        command1
        ;;
    value3)
        command2
        ;;
    [vV]alue4)
        command3
        ;;
    *)
        command4
        ;;
    esac
    
    6) 循环语句:
    IFS 变量表示缺省的分隔符whitespace, tab, newline等. 可以自行修改该值, 但是建议再改之前赋值给其他变量作为备份, 有利于再恢复到缺省值.
    for variable in word_list
    do
        commands
    done
    
    e.g. 
    for pal in Tom Dick Harry Joe
    do
        echo "Hi $pal"
    done
    echo "Out of loop"
     
    for file in testfile[1-5]
    do
        if [[ -f $file ]]
        then
         echo "$file exist"
        fi
    done
    
    for name in $*    // 等同于 for name
    do
        echo "Hi $name"
    done
     
    while command    //这里当条件为true的时候, 或者command的退出为0时执行循环体的命令.
    do
        commands
    done
     
    while (( $num < 10 ))
    do
        echo -n "$num"
        let num+=1
    done
     
    until !command    //这里当条件为false的时候, 或者command的退出为非0时执行循环体的命令.
    do
        commands
    done     
    
    break/continue:    功能等同于c语言中的break和continue.
     
    重定向循环和后台执行的循环:
    while command
    do
        command
    done > file    //重定向到文件
     
    while command
    do
        command
    done | command    //重定向到其他程序的管道
     
    while command
    do
        command
    done &     //后台执行.
     
    shift [n]命令: 用于把脚本的参数列表移出指定的位数, 不指定为左移一位, 一旦发生位移, 被移出的参数将永远被删除. 
    e.g.
     
    [scriptname: doit]
    while (( $# > 0 ))
    do
        echo $*
        shift
    done 
     
    /> doit a b c d e
    a b c d e
    b c d e
    c d e
    d e
    e





