#man  bash
${VAR_NAME#word}
  取得VAR_NAME中路径的基名，
  这里的word表示分隔符，而“#”表示是从左往右的方向分割，加一个“#”表示尽量向右匹配，
      所以“*”是在word的左侧
      这里相当于basename  COMMAND,但是这里可以针对任何变量。

# Path=/etc/rc.d/init.d/network 
# echo ${Path#/}
etc/rc.d/init.d/network
# echo ${Path##/}
etc/rc.d/init.d/network
#echo ${Path##*/}  (从左向右，以"/"为分割符，尽量截掉到最右端，常用）
network

${VAR_NAME%word}  常用来取得路径中文件的路径，
      这里的word仍表示分隔符，而%表示从右往左的方向分割，加一个%表示尽量向左。
      所以“*”的匹配需要从word右侧。
# echo ${Path%/*} 
/etc/rc.d/init.d

${VAR_NAME:-WORD} 
    使用默认值 
    当变量为空值时，就使用WORD当做默认值显示。 
${VAR_NAME:=WORD} 
    设定默认值 
    当变量为空值时，就设定WORD当做默认值显示并保存到变量。 
${VAR_NAME:?WORD} 
    当变量为Null或者未设置时，就显示WORD中的错误信息。 
${VAR_NAME:+WORD} 
    使用替换值 
    当变量为Null或者未设置时，什么都不显示。 
    当变量不为空时，就显示WORD替换原有的内容。

if判读语句中的取得奇偶数或者某个数的倍数时，需要使用取模运算：
  比如被除数为I,取值范围{1..999}，除数为9，取得9的倍数的数值为ZHI
(1)    ZHI=$(bc  <<<  "scale=0;$I%9“

(2)    ZHI=$[$I%9]

$0 是脚本本身的名字
$1,$2,$3....为位置变量
$#，表示除$0之外的参数的个数。
$@，除$0之外的参数的列表，这个列表时分开的
$* ，除$0之外的参数的列表，但是这个列表是一个整体
$？，命令执行的状态返回值
$$ 是脚本运行的当前进程ID号
例子：计算参数列表的和
#!/bin/bash
declare -i SUM=0
for I in $@;do
      SUM+=$I
done
echo $SUM
#./sum.sh  1 2 3
#6


在for循环中计算某个命令执行的次数，比如成功增加用户的次数，这时可以先在前面定义
  declare  -i COUNT=0
  然后再在for循环中命令执行成功或失败后，执行COUNT+=1

 

for 、until、while中的continue、break
  continue是返回重新循环。

  break是直接打断循环

在进行bash字符串测试时，除“>"、"<"、“==”
也使用[[  $option1  =~  $option2  ]]进行模式匹配，
              模式：
              一般做行首、行尾锚定；不要加引号；
              但是不能做词首、词尾锚定
              可以使用正则表达式
查找 计算机cpu的公司，这里使用模式匹配
#!/bin/bash
#
Vendor=`grep "vendor_id" /proc/cpuinfo  | uniq | cut -d: -f2`
或者A=$(grep -o "Genuine.*" /proc/cpuinfo | head -1)
if [[ "$Vendor" =~ [[:space:]]*GenuineIntel$ ]]; then
 echo "Intel"
else
 echo "AMD"
fi


cat  <<  EOF
m|M) show memory usages;
d|D) show disk usages;
q|Q) quit
EOF
在执行时就可以执行并显示上面的信息了！


遍历文件内容：
while  read  LINE;do
  state  1

  state  2

done  <  /path/to/somefile
使用read读取文件的一行，然后保存到LINE变量中，然后执行语句，循环后再读取第二行进行覆盖。
如果用户的ID号为偶数，则显示其名称和shell；对所有用户执行此操作
while read LINE; do
  Uid=`echo $LINE | cut -d: -f3` 
  if [ $[$Uid%2] -eq 0 ]; then
    echo $LINE | cut -d: -f1,7 
  fi
done < /etc/passwd

使用until或者while时，可以将前面没有进入循环的变量值重新复制，然后纳入循环。
每隔5秒查看Hadoop用户是否登录，如果登录，显示其登录并退出；
否则，显示当前时间，并说明hadoop尚未登录：
#!/bin/bash 
who | grep "^hadoop" &> /dev/null
RetVal=$? 
while [ $RetVal -ne 0 ]; do
  date 
  sleep 5
  who | grep "^hadoop" &> /dev/null  #这里既是重新对返回值赋值 
  RetVal=$? 
done 
echo "hadoop is here."