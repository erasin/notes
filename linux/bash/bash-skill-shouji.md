1701.大量连接处于 TIME_WAIT 状态的解决方法：

    netstat -nt 看到大量 TIME_WAIT
    dmesg 看到 time wait bucket table overflow
    解决办法：1、 /proc/sys/net/ipv4/tcp_max_tw_buckets 过小了，设成540000以上
              2、 /sys/module/ip_conntrack/parameters/hashsize 过小了，
                  设成/proc/sys/net/ipv4/netfilter/ip_conntrack_count这个值的1/4大小较为合适

1702.编译php出错：error while loading shared libraries: libmysqlclient.so.18: cannot open shared object：

    ln -s /usr/local/mysql/lib/libmysqlclient.so.18  /usr/lib/
    照做后仍然报错，原因是该方法适用于32位系统，64位系统应使用下面的这行
    ln -s /usr/local/mysql/lib/libmysqlclient.so.18  /usr/lib64/
    另外：在编译的时候，不写mysql的路径，而使用mysqld代替，也可解决该问题的出现。

1703.shell打印当前行号：

    echo "$LINENO"
    显示函数在哪一行被调用的：
    line(){
        echo call by line `caller 0 | awk '{print$1}'`
    }
    line

1704.当系统swap空间所剩不多时，可通过在本地磁盘上创建普通文件作为swap ：

    dd一个文件，然后mkswap，然后swapon
    dd if=/dev/zero  of=./swapfile bs=1M count=1000
    mkswap ./swapfile
    swapon ./swapfile  
    即可为系统临时增加1G的swap。
    使用完毕后执行：
    swapoff ./swapfile
    rm -f swapfile

1705.多行结果赋值给变量：用引号防止 shell 做 word split：

    file=`find . -name "*"`
    echo "$file"

1706.wireshark 八进制编码问题：

    google.com/search?q=淘宝网     这个http请求，抓包里面对应的内容是
    GET /search?q=\314\324\261\246\315\370 HTTP/1.1\r\n
    \314\324\261\246\315\370这一段对应的是淘宝网3个字，应该是gbk，gbk是每个汉子两个 
    \nnn   the eight-bit character whose value is the octal value nnn (one to three digits)
    8进制，转换成十六进制看看，再看编码 
    CC D4对应\314\324     echo 0:ccd4|xxd -r|iconv -f gbk -t utf-8
    echo $'\314\324\261\246\315\370'|iconv -f gbk -t utf-8
    前面要加$，这样是表示8进制的  

1707.sort -n -g 按数字排序的区别：

    -g 通用数值排序，-n 数值排序。
    -g 支持的记数方法比 -n 广，-g 使用标准 C 函数 strtod 来转换数值，支持科学记数法，如 1.0e-34，
    不过 -g 比 -n 慢得多。详细区别参考info sort文档或者stackoverflow.com。

1708.关于sort -h排序问题:

    echo "2025K
    1M
    1G"|sort -h        # du、ls、df 等能保证输出一定是k<m<g的格式，不会存在2025k这种形式的。

1709.使用 cURL 测试 Web 站点的响应时间：

    curl -o /dev/null -s -w %{time_connect}:%{time_starttransfer}:%{time_total} http://www.canada.com

1710.设置用户的进程数限制：

    /etc/security/limits.conf
    /etc/profile里面写ulimit -n 65535，只有root才能增加，这样其它用户登录都会报错。
    nofile 是打开文件数，nproc是进程数，soft 指的是当前系统生效的设置值。hard 表明系统中所能设定的最大值。
    soft 的限制不能比hard 限制高。用 - 就表明同时设置了 soft 和 hard 的值

1711.ssh不输出连接信息：

    ssh -t 会输出一个或者 ssh 加 -q 参数  或者 2>/dev/null 屏蔽掉

1712.wget将页面内容输出到标准输出：

    wget -qO- http://shushu.com.cn/tm.php        # -qO- 相当于 -q  -O -

1713.iptables端口转发：将本地80端口的请求转发到8080端口，当前主机ip为192.168.2.1

    iptables -t nat -A PREROUTING -d 192.168.2.1 -p tcp -m tcp --dport 80 -j DNAT --to-destination 192.168.2.1:8080
    ssh 或者 haproxy 之类的也可以做端口映射

1714.进程间通信方式：

    PIPE(FIFO)    
    消息队列        
    信号量(Semaphore)        
    共享存储        
    SOCKET    

1715.线程间通信常用的三种方法：

    1.全局变量
        进程中的线程间内存共享，这是比较常用的通信方式和交互方式。
        注：定义全局变量时最好使用volatile来定义，以防编译器对此变量进行优化。
    2.Message消息机制（windows下）
        常用的Message通信的接口主要有两个：PostMessage和PostThreadMessage线程常常要将数据传递给另外一个线程。
        Worker线程可能需要告诉别人说它的工作完成了，GUI线程则可能需要交给Worker线程一件新的工作。
    3.CEvent对象（windows下）
        CEvent为MFC中的一个对象，可以通过对CEvent的触发状态进行改变，从而实现线程间的通信和同步。

1716.本机免key登录：

    ssh-keygen -t dsa -f ~/.ssh/id_dsa
    cat ~/.ssh/id_dsa.pub >> ~/.ssh/authorized_keys

1717.curl抓取http返回码：

    curl -so /dev/null -w '%{http_code}\n' www.sina.com

1718.ssh -D 一句话代理：

    alias proxy='expect -c "spawn ssh -D 127.0.0.1:6789 67.205.5.88 -l username;expect *password* ;send -- \"password\r\" ;  expect eof"'
    alias proxy='expect -c "spawn ssh -D 127.0.0.1:6789 67.205.5.88 -l username;expect *password* ;send -- \"password\r\" ;  interact"'

1719.ps -ww 不限制宽度输出：结果也与终端有关

    09:19:48#tp#~> ps -few|wc -L
    157
    09:19:50#tp#~> ps -feww|wc -L
    1326
    09:19:52#tp#~> echo $COLUMNS
    157

1720./dev/null 被cp，权限被修改：

    mknod -m 666 /dev/null c 1 3        # 创建特殊文件的，比如block或者是管道文件 

1721.每 n 个字符截断换行的几种方法：

    echo 123190287923849241483971837103921|sed 's/..../&\n/g'
    dd cbs=4 conv=unblock；    fold -c4；grep -Eo '.{1,4}'

1722.awk输出到管道或者文件的时候有buffer，一般是4K，到达4K才输出，输出到标准输出没buffer

    i=0;while sleep 0.1;do echo -ne "\r"$((i++));done|awk 'BEGIN{ORS=RS="\r"}{print $0}'|tee a
    可以用这个测试下，如果去掉tee就可以正常显示，或者去掉awk也可以，两个都加上就不行，
    如果把sleep 0.1改成sleep 0.000001，到达4K的时候就会输出了,
    用fflush()可以刷新buffer 
1723.od进行ascii到八进制、10进制、16进制的转换：
    printf "^A" | od -An -o # -b -x -d
1724.curl测试网络状况：
    curl -so /dev/null -w '%{http_code}' -H host:www.host.com 8.8.8.8/a.txt        #测返回码
    curl -sH host:www.host.com 8.8.8.8/a.txt    #下载文件测试
1725.paste -s串行合并文件内容：
    seq 100|paste -sd '    \n'        #每5个换行
    echo {1..100}|xargs -n5            #xargs -n 效率比较低
    echo {1..100}|grep -Po '(\w+ ?){5}'
1726.printf打印ascii的10进制值：
    printf '%d\n' \'c        # \'c 意思就是使用字符作为数字值
    awk 打印 ascii 对应 10 进制值：
    echo a|awk 'BEGIN{for(i=0;i<255;i++)a[sprintf("%c",i)]=i}{print a[$0]}'
    打印ascii的16进制值：
    printf '%02x' "'+"      # echo -n '+' | od -tx1 -An | tr -d ' '
1727.- 表示标准输入或者标准输出，用法举例：
    tar cvf - * | (cd  /dest/dir && tar xvfp -)            # - 需要程序自己支持，例如tar、cat，-可以理解为临时文件形式存在
    -在命令符当中表示标准输入或者标准输出。
    在tar cvf - *当中，表示标准输出，因为f参数后面跟的是输出的，
    这个命令的意思就是把当前目录的所有文件，tar打包到标准输出当中
    然后通过管道，这个命令的标准输出作为后面 (cd  /dest/dir && tar xvfp -)
    命令的标准输入，cd /dest/dir没有什么好说的，切换到那个目录。
    然后tar xvfp表示解压缩文件，但是没有指明是那个具体的文件？所以就是-了，
    表示标准输入当中获取文件，然后解压缩到当当前目录。
    总体来看，这个命令就是拷贝文件了。
1728.sub()与match()的搭配使用,可逐次取出原字串中合乎指定條件的所有子字符串：
    awk 'BEGIN {
        data = "p12-P34 P56-p61"
        while( match( data ,/[0-9]+/) >0) {
           print substr(data,RSTART, RLENGTH)
           sub(/[0-9]+/,"",data)
        }
    }'        #类似python re中的group或者正则中的捕获组
1729.巧用find的内置命令：delete、ls等：
    find -name "AAA" -ls    #同时减少了-exec执行fork带来的开销。
    find  . -name  AAA  \( -type f -exec ls -l {} \; -o -type d -exec ls -ld {} \; \)
1730.显示所有 ipv4：
    ip -o -4 a s
1731.printf -v 直接赋值变量：
    printf -v var '%%%02x' 111
1732.awk  FIELDWIDTHS 将字符串按指定域宽度划分：
    echo '23456.23478.58924.6' | awk -vFIELDWIDTHS="3 6 5 5" -vOFS="\t" '{$1=$1;print}'
    echo ${A:0:3} ${A:3:6} ${A:9:5} ${A:14:5} 
1733.查看 tcp 连接中 keepalive 的数量：
    netstat -town
1734.模拟骰子产生1~6的随机数：
    echo $((0x$(head -c5 /dev/random|xxd -ps)%6+1))
1735.pkill杀进程是模糊匹配，小心误杀：
    pkill aaa ，所有进程名里包含aaa的都会被干掉
    你可以用pgrep看看，进程名，不包含参数，你还可以指定tty、用户名什么的  
    如果要精确匹配：pgrep -lx 类似 grep 的 -x 按行匹配
    -f 是不仅仅在进程里面查，参数也会查 ，就是你pkill -f apache，会把java -u apache也杀掉。默认是只看进程名是否匹配 
    pgrep -fl 99
    27659 sleep 99
    -fx 是完整匹配整个命令行（进程名+参数） ，-x 是完整匹配进程名 
1736.pkill 踢出who里的登录终端或者ip：
    pkill -kill -t pts/2
1737.得到10个字符的随机密码：
    tr -dc A-Za-z0-9_ < /dev/urandom | head -c 10 | xargs
1738.perl列出当前目录下所有的符号链接：
    perl -e 'print map { -l and $_ .= "\n" } <*>'
1739.nnohup退出终端会断掉：
    shopt |grep huponexit
    grep -iR hup .bashrc .profile /etc/bash* /etc/profile*找下，反正我这没有，我是off
    默认值都是off的，开了就会出现他的那种问题，shell退出不应该给这个shell的子进程发HUP信号的，可以trap下。
1740.利用nc远程传输文件：
    tar -cf - * | nc 172.20.51.51 5555
1741.正则贪婪匹配陷阱：最左边的贪婪匹配优先级最高，后面的匹配都是基于正则的回溯机制，所以后面的都属于非贪婪匹配了。
    *虽然贪婪，但也可能发空匹配，因为他可以为0，而+至少匹配一次，限制较强。
    echo "房屋信息        房屋信息        1--district     190_360"|sed -r 's/(.*)([0-9]+)_([0-9]+)/=\1==\2===\3/'
    =房屋信息        房屋信息        1--district     19==0===360
    echo "房屋信息        房屋信息        1--district     190_360"|sed -r 's/(.*)([0-9]*)_([0-9]*)/=\1==\2===\3/'
    =房屋信息        房屋信息        1--district     190=====360
    echo "房屋信息        房屋信息        1--district     190_360"|sed -r 's/(.+)([0-9]+)_([0-9]*)/=\1==\2===\3/'
    =房屋信息        房屋信息        1--district     19==0===360
    echo "房屋信息        房屋信息        1--district     190_360"|sed -r 's/(.*) ([0-9]+)_([0-9]*)/=\1==\2===\3/'
    =房屋信息        房屋信息        1--district    ==190===360
1742.rsync -b 更新前先备份：
    rsync -b，然后--suffix指定后缀或者--backup-dir指定备份的目录，mv肯定比cp快，rsync的-b也是mv的 
1743.windows下查询域名对应的所有服务器：
    nslookup -qt=mx web.qq.com
1744.vi一个远程文件：
    vi scp://username@host//path/to/somefile        #注意双斜线
1745.删除所有非打印字符：
    tr -dc '[:print:]' < filename
1746.PS1 设置putty等ssh工具的标签页 title：
    PS1='\u@\h:\w\$ '
    # If this is an xterm set the title to user@host:dir
    case "$TERM" in
        xterm*|rxvt*)
            PS1="\[\e]0;\u@\h: \w\a\]$PS1"
            ;;
        *)
            ;;
    esac
    或者：
    PROMPT_COMMAND='echo -ne "\033]0;${USER}@${HOSTNAME%%.*}:${PWD}"; echo -ne "\007"'
1747.tail -f 管道的buffer不能实时输出的问题：
    unbuffer -p tail -f .sh_histoy |awk '{print $0}'
1748.利用 bc 计算器 obase 参数，实现十进制转二进制：
    echo 'obase=2;98'|bc
1749.为 man 指定阅读器 pager ：less，并搜索指定字符串 034
    PAGER='less -iRsp034' man ascii
1750.vim中获取只读文件的sudo权限写入：:w !sudo tee %
    命令:w !{cmd}，让 vim 执行一个外部命令{cmd}，然后把当前缓冲区的内容从 stdin 传入。
    tee 是一个把 stdin 保存到文件的小工具。
    而 %，是vim当中一个只读寄存器的名字，总保存着当前编辑文件的文件路径。
    所以执行这个命令，就相当于从vim外部修改了当前编辑的文件，然后从定向到当前文件。
    %! sudo tee % > /dev/null
    %　        #VI/VIM编辑的文件内容
    !　        #管道
    sudo    #以root权限操作
    tee　    #
    %　　    #VI/VIM编辑的文件
    > /dev/null　    #这里可省略
1751.查看linux版本的几种方法：
    cat /proc/version; uname -a; lsb_release -a; cat /etc/issue; 
    cat /etc/redhat-release; rpm -q redhat-release
    file /bin/bash; file /bin/cat
1752.ls --time-style自定义时间格式显示：
    ls -lth --time-style=+"%F %T"
1753.去除vi打开文件时带的颜色及控制乱码字符：col -b
1754.sh 虽然链接到了 /bin/bash，但是 /bin/sh 执行效果不完全等同 /bin/bash，而是 /bin/bash --posix
    有很多系统里 gzip gunzip zcat 都指向一个inode, 同一程序对不同加载名作不同动作的用法在*nix中太常见了
1755.正则捕获示例：抓取网页上的 URL 链接：注意设置 LC_ALL 语言环境和字符集
    curl -s www.58.com.cn | sed -rn "s#.*(http://[0-9A-Za-z]*\.[^/\"?' ]*).*#\1#p"|sort|uniq -c|sort -k1rn
    curl -s www.qq.com | LC_ALL=C sed -rn "s#.*(http://\w*\.[[:alnum:].]*).*#\1#p"|sort|uniq -c|sort -k1rn
    上面的 sed 有一行上不能匹配多个网址的 bug，因为sed需要完全匹配后才能print，下面用 grep 修复该问题
    curl -s www.360buy.com | grep -Po '(https?://\w*\.[[:alnum:].]*)' |sort|uniq -c|sort -k1rn
    注：posix 字符类 [[:alnum:]] 等价于 [0-9A-Za-z]，
    \w 为 [[:alnum:]_]的同义词，指单词字符，相当于 [0-9A-Za-z_]，关于这点 man grep \w 中有前后矛盾的描述。
1756.sed单引号转义问题：
    sed单引号转义只支持16进制 \x27, 而且外层一定要加单引号或者双引号引起来：
    echo "1'1"|sed -r 's/\x27//'    # 不加引号那shell 会解释\x 成为 x，所以要养成加引号的习惯
    不支持8进制：echo "1'1"|sed -r 's/\047//' ，而 awk 是两者都支持的。
    或者sed中用 "'" 表示单引号
1757.用 bc -l 加载数学库做小数计算：
    echo 2/70|bc -l        #    echo 2/70|awk '{printf "%f", 2/70}'
1758.locale 字符集语言环境的设置：
    echo http://wwAw.HA~12o.com|LC_ALL=C sed -rn "s#.*(http://([a-z]*)\.[^/\"?' ]*).*#\2#p"
    echo Aab|LC_ALL=C sed -r 's/[A-Z]//g'            
    注：有时候莫名其妙的问题一般与 LC_ALL=C 有关，C 是 POSIX，和具体语言无关。
    其实最好是C，这样sort排序会快很多，还有遇到[a-z]这种情况比较安全，还有 [.-/] 这种
1759.巧用 sed // 上次匹配功能排除边界值：
    echo -e "a\n2\nc"|sed '/a/,/c/{//!d}'    # // 相当于awk中 !/^a|^c/
1760.shell中转换16进制到10进制：
     i=A; echo "obase=10;ibase=16;$i" | bc
     a=99 b=e8; (( 16#a > 16#b ))&&echo "a > b"||echo 'a <= b'
     echo $(printf "%d" 0xD)
1761.rsync 只同步目录（危险慎用！）
    "只" 同步目录，子目录的子目录也可以同步，但子目录里的文件不同步。
    rsync -av --delete -f '+ */' -f '- *' SRC/ DEST/    # -f '+ */'是同步目录，-f '- *'是不同步文件。
    想把子目录的内容也同步：
    rsync -av --delete -f '+ */' -f '- /*' SRC/ DEST/
1762.vim 粘贴板简介: 
    vim 有 12个粘贴板 0、1、2、…、9、a、"、＋；用:reg命令可以查看各个粘贴板里的内容。
    在vim中简单用y只是复制到 " (双引号)粘贴板里，p也这样。
    Ny完成复制到N号剪贴板，其中N为粘贴板号(注意是按一下双引号然后按粘贴板号最后按y)，
    例如要把内容复制到粘贴板a，选中内容后按 "ay 就可以了，有两点需要说明一下：
    +号粘贴板是系统粘贴板，用"+y将内容复制到该粘贴板后可以使用Ctrl＋V将其粘贴到其他文档（如firefox、gedit）中，
    同理，要把在其他地方用Ctrl＋C或右键复制的内容复制到vim中，需要在正常模式下按 "+p，
    要将vim某个粘贴板里的内容粘贴进来，需要退出编辑模式，在正常模式按"Np，其中N为粘贴板号，
    如上所述，可以按"5p将5号粘贴板里的内容粘贴进来，也可以按"+p将系统全局粘贴板里的内容粘贴进来。
1763.查看本机外网 ip：
    curl ifconfig.me    
1764.awk 处理浮点数四舍五入的问题：
    ieee754 http://zh.wikipedia.org/zh-cn/IEEE_754
    $ echo "3.445"|awk '{printf "%0.2f\n",$0}' 
    3.44
    There is no ieee754 number for 3.445000000000000000000...
    $ echo "3.445"|awk '{printf "%0.20f\n",$0}' 
    3.44499999999999984013
    $ echo "3.4450000000000001"|awk '{printf "%0.20f\n",$0}' 
    3.44500000000000028422
    $ echo "3.44500000000000001"|awk '{printf "%0.20f\n",$0}' 
    3.44499999999999984013
1765.过滤掉 /etc/passwd 中的自动生成用户
    awk -F: '$0!~/nologin/&&($3=0||$3>=500)' /etc/passwd
1766.利用taskset有效控制cpu资源
    taskset -c 0 sh job.sh        #开启一个只用0标记的cpu核心的新进程(job.sh是你的工作脚本)
    taskset -pc 0 23328            #查找现有的进程号，调整该进程cpu核心使用情况
    可在top中进行负载check，最后你可以在你的工作脚本中加入该指令来合理利用现有的cpu资源
1767.关于 [^A-Za-z] 是否等价 [^A-z] 的问题：
    [^a-zA-Z] 是正则
    [!a-zA-Z] 是通配符
    用的地方不一样，而且A-z包含几个非字母字符，且用的时候还需要注意字符集
1768.校对大量文件的完整性：
    rsync 返回码或者 find打印大小，不要用du，du与FS的block大小设置有关，异构的FS中不准确
1769.登录linux，执行配置文件的顺序：
    /etc/profile -> [~/.bash_profile | ~/.bash_login | ~/.profile] -> ~/.bashrc -> /etc/bashrc -> ~/.bash_logout
    在登录时，首先会执行/etc/profile，其次是用户目录下的.bash_profile，.bash_profile不存在则读取.bash_login，
    若.bash_profile和.bash_login都不存在，最后才读取.profile
1770.wget 带上浏览器 agent 下载：
    wget -c $url -O test2.flv -U    # --user-agent=AGENT
1771.ed 是真正的 inplace editing，sed 是假的 inplace editing，中间是生成了临时文件的。
1772.查看文件编码的 2 种方式：file -i 文件名    # vim 中 :set fileencoding
1773.根据进程名显示进程号： pidof java
1774.如何限制用户的最小密码长度：修改 /etc/login.defs 里面的PASS_MIN_LEN的值
1775.如何使新用户首次登陆后强制修改密码：useradd -p '' testuser; chage -d 0 testuser
1776.vi 编辑标准输入输出：crontab -l | vi -
1777.awk 利用协同进程监听指定端口：
    awk 'BEGIN{service="/inet/tcp/2000/0/0"; service |& getline; print $0; close(service)}'
1778. awk ‘FS=”:” {print $1,$3,$6}’ 明显不等价于 awk -F: ‘{print $1,$3,$6}’ ，
    前者是读取第一行记录并将记录分割成字段之后才进行FS的设置,FS=”:” 作为了模式，通常这样的写法是错误的.
    例如 awk 'FS=":" {print $1,$3,$6}' /etc/passwd 会把第一个整行打印出来
1779.awk gensub 删除字段及其FS：
    echo 1 2 3 4 5|awk '{$4="";print gensub(FS,"",4)}'
    echo 1 2 3 4 5|awk '{sub($4FS,"")}1'
    echo "1 2 3    4    5"|awk '{$4="";$0=$0;$4=$4;print}'        
    注意 $1=$1;$4=$4 缺一不可，否则会有残留FS输出，$4=$4 在重构该域时会把周围多个 FS 合并，就像 echo 1    2 一样
1780.设置 PS1，根据上一个命令返回状态码控制显示红色还是绿色：
    `[ $? -eq 0 ]||echo \[\e[31m\]`
1781.判断程序输入来自文件还是标准输入：
    用 < 的话也是来自 - 的，不是来自文件
    15:18:17#tp#~> awk 'END{print FILENAME}' <a
    -
    15:18:23#tp#~> cat a|awk 'END{print FILENAME}'
    -
    15:18:58#tp#~> 
1782.查看dd命令的执行进度:
    dd if=/dev/zero of=/data3/test bs=1k count=1024000        # a 终端执行
    while killall -USR1 dd; do sleep 5; done                # b 终端执行
    然后就能看到间隔5秒一次的输出结果了，所执行的killall命令循环会在dd命令执行结束之后退出。
    为什么dd命令的进度输出要killall命令来激发呢？其实答案很简单，dd的进度输出就是这么设计的，
    再详细点就是dd命令在执行的时候接收到SIGUSR1信号的输出当前的读写进度，而killall命令在这里起的作用就是给dd进程发送这个信号。
    对于没有那些不与终端交互的守护进程(Deamon Process)，通常的设计是在接收到SIGHUP信号之后就会重新读取配置文件。比如下面的：
    # 当你对smb.conf修改之后，运行下面的命令，修改就能生效了
    # 而不用再来一个smb restart
    killall -HUP smbd
1783.atime、ctime、mtime 区别与联系：
    atime：访问时间（access time），指的是文件最后被读取的时间，可以使用touch命令更改为当前时间；
    ctime：变更时间（change time），指的是文件属性最后被变更的时间，变更动作可以使chmod、chgrp、mv等等；
    mtime：修改时间（modify time），指的是文件内容最后被修改的时间，修改动作可以使echo重定向、vi等等；
    ctime 变动时不影响 mtime
    mtime 变动时 ctime 一起变。
1784.bc中十进制、二进制互转，ibase影响你所有的输入，
    14:53:49#tp#~> echo 'ibase=10;obase=2;11'|bc -l
    1011
    14:53:53#tp#~> echo 'ibase=2;obase=10;11'|bc -l    # 这里obase=10就是二进制 
    11
    14:53:58#tp#~> echo 'ibase=2;obase=1010;11'|bc -l
    3 
    14:55:21#tp#~> echo 'ibase=2;obase=10;11'|bc -l
    11
    14:58:59#tp#~> echo 'obase=10;ibase=2;11'|bc -l    # ibase影响你所有的输入，换个位置就不一样了 
    3
    14:59:09#tp#~>
1785.nginx 的 logRotation 机制：
    mv access.log access.log.0
    kill -USR1 `cat master.nginx.pid`    # 发信号让进程写 access.log，而不是 access.log.0 老日志
    sleep 1
    gzip access.log.0    # do something with access.log.0
1786.linux 文件系统的一些限制：
    一级子目录限制个数：32000 （真正可以容纳的是31998个子目录，. 和 ..）
    文件名字符长度 255，中文 127
    文件个数：貌似没有限制
    注：不同的文件系统可能限制不一样，修改 kernel include file 然后 recompile 可以修改限制
1787.关于 find 不管找到与否都返回 0 的问题：因为默认是 find -print
    find -type f -name "123"
1788.管道是异步（并行）执行的：
    20:28:29#tp#~> sleep 999|sleep 998 &
    [1] 9239
    20:28:34#tp#~> ps -eo lstart,cmd|grep 'sleep 99'
    Mon Mar 18 20:28:33 2013 sleep 999
    Mon Mar 18 20:28:33 2013 sleep 998
    Mon Mar 18 20:28:35 2013 grep --color=auto sleep 99
    20:28:36#tp#~>
    sleep 999和sleep 998是同时启动的，并不是sleep 999执行完成，再执行sleep 998的 
    管道前后的进程并发执行，不过这个还是会有一个时间差，这个看cpu的调度和负载。
    正因为如此，ps -fe|grep grep 不一定会有结果，而 ps -fe|grep init 一定有结果
    ps -fe|grep grep，可能是零个，也可能一个，不过大部分情况都是一个
    ps -fe|grep init，可能是一个，也可能是两个，不过大部分情况都是两个
1789.管道其实在操作内存：
    echo|ls -l /proc/self/fd/，结果里面的0 -> pipe:[1148917]就是管道文件 
    或者ls -l /proc/self/fd/|cat，这样的话1是管道 
    如果没有管道的话，应该是/dev/pts/4什么的 
    self是你当前的pid，fd是这个进程的文件描述符，1148917 是匿名管道的文件名
1790.强大的进程匹配工具 pgrep：可依据多个属性匹配指定的进程
    如：pgrep -u root sshd
    pgrep -fl bash 等价 ps -eo pid,cmd|grep [b]ash
1791.nc 测网卡带宽：
    nc -l -p 1235 </dev/zero
    nc 127.0.0.1 1235 >/dev/null 
    直接从/dev/zero读数据到1235端口，然后我的第二个nc从1235接手数据写到/dev/null，然后ifstat测试当前流量
    全部走内存，没有磁盘io，可以用eth测试，但是eth有危险，万一把网卡堵死，命令都执行不了。
1792.当 FS 为单个空格时，无论 RS 是否被重新定义，FS 始终都是 [ \t\n]+
    seq 3|awk '{$0=$0"\n#\n";print $1,$2}'
1793.注意大括号命令组的用法：大括号前后空格 和 每条命令分号结束，否则提示命令未结束
    { sleep 1; continue; } 
1794.lsof && test -s 检查文件写入完成，或者先写临时文件再 mv：
    chkFileWriteOver(){
        while [[ $c -le 3600 ]]
        do  
            ((c++))
            # if file length > 0, it's ok.
            [[ -s $1 ]] || { sleep 1; continue; }
            lsof $1
            # if file not locked, it's ok.
            [[ $? -eq 1 ]] && echo "---------- File is OK ..." && break
            sleep 1
        done
        # here, if file is exception, you can send a email to somebody.
        [[ $c -eq 3600 ]] && echo "---------- File is not exsit ..."
    }
    chkFileWriteOver fileName 
1795.shell的解释和编译引起的性能问题：
    for i in {1..10000}; do str="$str $i"; done
    # 这里是for慢，不是{1..10000}慢，你echo {1..10000}很快的 
    解析 1次循环0.001秒 10次循环0.01秒 1000次循环1秒
    编译 1次循环0.001秒 10次循环0.002秒 1000次循环0.003秒 
    编译其实就是那个意思，并不是那个比例，解析的话几次就是几倍 
    可以认为编译 T(1)，解释 T(n)
1796.find的路径问题：-newer 读取文件时间默认是在当前路径，因此需要显示 cd 或者 -newer 绝对路径
    find /tmp/mstat ! -newer `date -I -d'3 days ago'`
    find: 2013-03-24: 没有那个文件或目录
    cd /tmp/mstat; find /tmp/mstat ! -newer `date -I -d'3 days ago'`
    /tmp/mstat/2013-03-24
1797.sed 模式匹配//与模式替换s///的语法区别：
    当都是 / 做分隔符并无区别，但是当不是 / 的时候是有语法区别的：
    sed '\@/usr/local/test.txt@d'     # 第一个字符需要转义，man sed 语法如此，因为还可以被解析为行号
    sed 's@/usr/local/test.txt@@'    # 只能被解析为字符串内容，没有歧义
1798.iptables 的保存：
    /etc/init.d/iptables save就好了，保存到 /etc/sysconfig/iptables，启动自己会去读取的 
    也可以手动 iptables-save > a.txt 或者 iptables-restore < a.txt
1799.curl 获取网页的状态码：
    http_code=`curl -I -s www.baidu.com -w %{http_code}` && echo ${http_code:9:3}
1800.awk中的隐式数据类型转换：
    awk 'BEGIN{a="7868";if(a>=6211&&a<=10075)print "ok"}'
    awk 'BEGIN{a="56";if(a>=3&&a<=7189)print "ok"}'
    结果不一致是由于在gawk中，数字和字符串同时比较会发生隐式类型转换，
    会把数字转换成字符串，然后按字符串排序（字典序），要想用数字比较，
    需要显示转换类型：a="7868"+0; 或者 a=int("7868");
    需要注意的是这样他只会转换以数字开头到第一个字母前的字符串
    如：
    echo |awk '{print "a12b"+0,"|",int("a12b"),"|","12b1"+0,"|",int("12b1")}'
    0 | 0 | 12 | 12
    awk的sorti也是使用字符串做索引排序，需要注意