Bash数组
1.描述
  Bash的数组，其元素的个数没有限制。数组的索引由0开始，但不一定要
连续(可以跳号)。索引也可以算术表达式。bash仅支持一维数组。
1.1 declare
  名称：设置变量和属性(可以用来声明一个数组)
  用法：
 
      declare [-aAfFilrtux] [-p] [Name[=Value]...]
  选项：
      -f 显示函数名和函数定义
      -F 仅显示函数名
      下面是用来设置属性的选项
       -a  表示Name为一个索引数组
       -A  表示Name为一个关系数组
       -i  表述Name为一个整数
       -l  将Name的值转换为小写
       -u  将Name的值转换为大写
       -r  表示Name为一个只读变量
       -x  表示Name为输出变量(输出为全局变量)
       -t  表示Name具有'trace'属性
      使用'+'替代'-'来关闭指定的属性
      当用在一个函数内部时，declare等价local
  实例：
 
      declare -A myarray  #声明一个关系数组myarray
      myarray[1]=val1
      myarray[item2]=val2
1.2 数组的用法
   [1].建立数组
       格式：
           数组名[索引]=值
           ......
       例:
           A[0]=val1
           A[1]=val2
           A[2]=val3
    [2].引用数组的元素值
 
        格式：
           ${数组名[索引]}
         例：
           echo ${A[0]}
    [3].一次设置一个数组的多个元素的值
 
         格式：
           数组名=([索引]=val [索引]=val2 ...)或
           数组名=(val1 val2 val3 ...)
         实例：
           B=(1 2 3 4 5)
           B=([0]=1 [1]=2 [2]=3 [3]=4 [4]=5)
    [4].一次取出数组的所有元素
          格式:
            ${数组名[@]}  (取出的每个元素以空白符隔开)
            ${数组名[*]}  (取出的元素组成一个字符串)
    [5].获得数组的元素的个数
          格式：
            ${#数组名[@]} 或 ${#数组名[*]}
    [6].取得某个元素的长度
          格式:
            ${#数组名[索引]}
    [7].取消数组或数组元素的设置
          格式：
            unset 数组名
            unset 数组名[索引]
 或 ${#数组名[*]}
    [6].取得某个元素的长度
          格式:
            ${#数组名[索引]}
    [7].取消数组或数组元素的设置
          格式：
            unset 数组名
            unset 数组名[索引]
1.3 mapfile
   名称：从标准输入中读取到一个索引数组中
   用法：
 
       mapfile [-n Count] [-O origin] [-s Count] [-t] [-u FD]
               [-C Callback] [-c Quantum] [Array]
   描述：
       从标准输入中读取并输出到一个数组中，也可以通过指定-u FD来从一个文件中读取
  若没指定数组名，则默认为MAPFILE
   选项：
      -n Count:最多读取Count行，若Count为0则表示读取所有行
      -O Origin:数组的索引从Origin开始，默认为0
      -s Count:忽略前Count行
      -u FD:从指定的文件描述符FD中读取，默认为从标准输入中读取
      -t:从读取的每一行后面移除换行符
   参数：
      Array :指定数组名，默认为MAPFILE
   实例：
      mapfile myarray < file.txt
      #将file.txt的内容保存为数组myarray.数组的每个元素为一行



摘要：
1、array=(value1 value2 ...... valueN)    #赋值
2、read -a array                                  #读入数组
3、${array[@]}    ${!array[@]}     ${#array[@]}     ${array[@]:n:m}      #数组信息，数组下标，数组长度，去数组位置
4、# array=($(ls | grep rpm))        #命令执行结果放入数组
5、for i in ${a[@]} ; do      #遍历数组
6、set| grep array             #利用set查看数组赋值情况

1、数组的声明、赋值和查看
   BASH只支持一维数组，但参数个数无限制。如果要构造一个二维数组，就需要自己想办法，在第三小节会有举例。bash数组的下标从0开始。
   array=(value1 value2 ...... valueN)               #从下标0开始依次赋值
   array=([1]=value1 [2]=value2 [0]=value0)   #指定下标赋值
   declare -a array=(value1 value2 ...... valueN)   #声明+赋值，也可以只声明
   unixtype=('Debian' 'Red Hat' 'Fedora')           #如果元素有空格，就要用引号
   
   unset array                     #清除数组
   unset array[1]                 #清除数组的指定元素     

read赋值举例
# read -a array                  #-a表示从标准输入读入数组，遇到换行为止
1 2 3 4 5
# echo "${array[@]}"
1 2 3 4 5

查看举例
# echo ${array[0]}
red hat
# echo "${array[0]}"                    #注意加不加引号的区别
red    hat
# echo "${array[@]}"
red    hat fedora
# echo ${array[@]}                     #查看数组的信息，注意引号的区别
red hat fedora
# set | grep array
array=([0]="red    hat" [1]="fedora")        #通过set查看变量和grep结合查看数组的赋值情况

2、数组常用的变量
   array[0]  array[1]  ...... #是数组的每个元素，根据下标指定，类似C语言，读取时相当于变量${array[1]}
   ${#array[@]}  #数组长度
   ${array[@]}    #数组的所有元素
   for i in ${array[@]};do ...  #遍历数组，这时i就是array里的某个元素
   或者for idx in ${!array[@]}   #这时idx就是array的某个下标
   ${array[@]:n:m}   #从数组的n位置开始取m个元素

   变量使用的举例如下：
# array=('red    hat' 'fedora')            #赋值，第一个元素中有多个空格
# echo ${array[@]}
red hat fedora                                  #打印数组，第一个元素的空格变成一个了
# echo "${array[@]}"                      
red    hat fedora                               #加引号保持原貌
# echo ${#array[@]}                        
2                                                     #数组长度
# echo ${!array[@]}
0 1                                                   #数组下标
# echo ${array[@]:1:1}
fedora                                               #获取数组元素

3、数组的常用操作
（1）命令执行结果放入数组
  # array=($(ls | grep rpm))          #建议采用($())的方式
# echo ${array[@]}
bind-9.8.2-0.17.rc1.el6_4.4.x86_64.rpm bind-chroot-9.8.2-0.17.rc1.el6_4.4.x86_64.rpm bind-devel-9.8.2-0.17.rc1.el6_4.4.x86_64.rpm bind-libs-9.8.2-0.17.rc1.el6_4.4.x86_64.rpm bind-sdb-9.8.2-0.17.rc1.el6_4.4.x86_64.rpm bind-utils-9.8.2-0.17.rc1.el6_4.4.x86_64.rpm

# array=(`ls | grep rpm`)            #效果相同，这个例子采用反向单引号
# echo ${array[@]}
bind-9.8.2-0.17.rc1.el6_4.4.x86_64.rpm bind-chroot-9.8.2-0.17.rc1.el6_4.4.x86_64.rpm bind-devel-9.8.2-0.17.rc1.el6_4.4.x86_64.rpm bind-libs-9.8.2-0.17.rc1.el6_4.4.x86_64.rpm bind-sdb-9.8.2-0.17.rc1.el6_4.4.x86_64.rpm bind-utils-9.8.2-0.17.rc1.el6_4.4.x86_64.rpm

（2）读入字符串，给数组赋值
i=0
n=5
while [ "$i" -lt $n ] ; do                     #遍历5个输入
  echo "Please input strings ... `expr $i + 1`"            
  read array[$i]                                #数组赋值
  b=${array[$i]}
  echo "$b"
  i=`expr $i + 1`                              #i递增
done

（3）字符串的字母逐个放入数组，并输出
chars='abcdefghijklmnopqrstuvwxyz'
i=0
while [ "$i" -lt ${#chars} ] ; do    # ${#char}是字符串长度
   array[$i]=${chars:$i:1}            #从$i取1个字节
done

（4）判断一个变量是否在数组中
一个很简洁的写法是：
echo ${array[@]} | grep -wq "${member}"
  if [ $? -eq $SUCCESS ];then

但是这会带来一个问题，如果array的元素里面带有空格，就会误认为是一个元素，因此遍历比较是更稳妥的选择。
for i in ${array[@]};do
   if [ "$i" = "${member}" ];then
   ....
   fi
done

（5）构建二维数组
a=('1 2 3' '4 5 6' '7 8 9')             #赋值，每个元素中都有空格
for i in ${a[@]} ; do
   b=($i)                                    #赋值给b，这样b也是一个数组
   for j in ${b[@]};do                  #相当于对二元数组操作
   ......
   done
done

（6）文件内容读入数组
# cat /etc/shells | tr "\n" " " >/tmp/tmp.file                      #回车变空格
# read -a array < /tmp/tmp.file                                       #读入数组
# set| grep array
array=([0]="/bin/sh" [1]="/bin/bash" [2]="/sbin/nologin" [3]="/bin/tcsh" [4]="/bin/csh" [5]="/bin/dash")


追加 append

  ARRAY=()
  ARRAY+=('foo')
  ARRAY+=('bar')