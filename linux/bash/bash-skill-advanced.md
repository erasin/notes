


该系列将重点介绍Linux Shell中的高级使用技巧，其主要面向有一定经验的Shell开发者、Linux系统管理员，以及Linux的爱好者。博客中的示例主要来源于网络和一些经典书籍，在经过本人的收集和整理之后，以系列博客的形式呈现给诸位。如果大家有更多更好的Shell脚本经典示例，且愿意在这里与我们一同分享的话，可以以邮件、博客回复等形式与我联系，我将会尽量保证该系列的持续更新。

一、将输入信息转换为大写字符后再进行条件判断：

      我们在读取用户的正常输入后，很有可能会将这些输入信息用于条件判断，那么在进行比较时，我们将不得不考虑这些信息的大小写匹配问题。

      /> cat > test1.sh
      #!/bin/sh
      echo -n "Please let me know your name. "
      read name
      #将变量name的值通过管道输出到tr命令，再由tr命令进行大小写转换后重新赋值给name变量。
      name=`echo $name | tr [a-z] [A-Z]`
      if [[ $name == "STEPHEN" ]]; then
          echo "Hello, Stephen."
      else
          echo "You are not Stephen."
      fi
      CTRL+D
      /> ./test1.sh
      Please let me know your name. stephen
      Hello, Stephen.

二、为调试信息设置输出级别：
    
      我们经常在调试脚本时添加一些必要的调试信息，以便跟踪到程序中的错误。在完成调试后，一般都会选择删除这些额外的调试信息，在过了一段时间之后，如果脚本需要添加新的功能，那么我们将不得不重新进行调试，这样又有可能需要添加这些调试信息，在调试成功之后，这些信息可能会被再次删除。如果我们能够为我们的调试信息添加调试级别，使其只在必要的时候输出，我想这将会是一件非常惬意的事情。
      /> cat > test2.sh
      #!/bin/sh
      if [[ $# == 0 ]]; then
          echo "Usage: ./test2.sh -d debug_level"
          exit 1
      fi
      #1. 读取脚本的命令行选项参数，并将选项赋值给变量argument。
      while getopts d: argument
      do
          #2. 只有到选项为d(-d)时有效，同时将-d后面的参数($OPTARG)赋值给变量debug，表示当前脚本的调试级别。
          case $argument in
          d) debug_level=$OPTARG ;;
          \?) echo "Usage: ./test2.sh -d debug_level"
              exit 1
              ;;
          esac
      done
      #3. 如果debug此时的值为空或者不是0-9之间的数字，给debug变量赋缺省值0.
      if [[ -z $debug_level ||  $debug_level != [0-9] ]]; then
          debug_level=0
      fi
      echo "The current debug_level level is $debug_level."
      echo -n "Tell me your name."
      read name
      name=`echo $name | tr [a-z] [A-Z]`
      if [ $name = "STEPHEN" ];then
          #4. 根据当前脚本的调试级别判断是否输出其后的调试信息，此时当debug_level > 0时输出该调试信息。
          test $debug_level -gt 0 && echo "This is stephen."
          #do something you want here.
      elif [ $name = "ANN" ]; then
          #5. 当debug_level > 1时输出该调试信息。
          test $debug_level -gt 1 && echo "This is ann."
          #do something you want here.
      else
          #6. 当debug_level > 2时输出该调试信息。
          test $debug_level -gt 2 && echo "This is others."
          #do any other else.
      fi
      CTRL+D
      /> ./test2.sh 
      Usage: ./test2.sh -d debug_level
      /> ./test2.sh -d 1
      The current debug level is 1.
      Tell me your name. ann
      /> ./test2.sh -d 2
      The current debug level is 2.
      Tell me your name. ann
      This is ann.

三、判断参数是否为数字：

      有些时候我们需要验证脚本的参数或某些变量的值是否为数字，如果不是则需要需要给出提示，并退出脚本。
      /> cat > test3.sh
      #!/bin/sh
      #1. $1是脚本的第一个参数，这里作为awk命令的第一个参数传入给awk命令。
      #2. 由于没有输入文件作为输入流，因此这里只是在BEGIN块中完成。
      #3. 在awk中ARGV数组表示awk命令的参数数组，ARGV[0]表示命令本身，ARGV[1]表示第一个参数。
      #4. match是awk的内置函数，返回值为匹配的正则表达式在字符串中(ARGV[1])的起始位置，没有找到返回0。
      #5. 正则表达式的写法已经保证了匹配的字符串一定是十进制的正整数，如需要浮点数或负数，仅需修改正则即可。
      #6. awk执行完成后将结果返回给isdigit变量，并作为其初始化值。
      #7. isdigit=`echo $1 | awk '{ if (match($1, "^[0-9]+$") != 0) print "true"; else print "false" }' `
      #8. 上面的写法也能实现该功能，但是由于有多个进程参与，因此效率低于下面的写法。
      isdigit=`awk 'BEGIN { if (match(ARGV[1],"^[0-9]+$") != 0) print "true"; else print "false" }' $1`
      if [[ $isdigit == "true" ]]; then
          echo "This is numeric variable."
          number=$1
      else
          echo "This is not numeric variable."
          number=0
      fi
      CTRL+D
      /> ./test3.sh 12
      This is numeric variable.
      /> ./test3.sh 12r
      This is not numeric variable.

四、判断整数变量的奇偶性：

      为了简化问题和突出重点，这里我们假设脚本的输入参数一定为合法的整数类型，因而在脚本内部将不再进行参数的合法性判断。
      /> cat > test4.sh
      #!/bin/sh
      #1. 这里的重点主要是sed命令中正则表达式的写法，它将原有的数字拆分为两个模式(用圆括号拆分)，一个前面的所有高位数字，另一个是最后一位低位数字，之后再用替换符的方式(\2)，将原有数字替换为只有最后一位的数字，最后将结果返回为last_digit变量。 
      last_digit=`echo $1 | sed 's/\(.*\)\(.\)$/\2/'`
      #2. 如果last_digit的值为0,2,4,6,8，就表示其为偶数，否则为奇数。
      case $last_digit in
      0|2|4|6|8)
          echo "This is an even number." ;;
      *)
          echo "This is not an even number." ;;
      esac
      CTRL+D
      /> ./test4.sh 34
      This is an even number.
      /> ./test4.sh 345
      This is not an even number.
        
五、将Shell命令赋值给指定变量，以保证脚本的移植性：

有的时候当我们在脚本中执行某个命令时，由于操作系统的不同，可能会导致命令所在路径的不同，甚至是命令名称或选项的不同，为了保证脚本具有更好的平台移植性，我们可以将该功能的命令赋值给指定的变量，之后再使用该命令时，直接使用该变量即可。这样在今后增加更多OS时，我们只需为该变量基于新系统赋予不同的值即可，否则我们将不得不修改更多的地方，这样很容易导致因误修改而引发的Bug。

      /> cat > test5.sh
      #!/bin/sh
      #1. 通过uname命令获取当前的系统名称，之后再根据OS名称的不同，给PING变量赋值不同的ping命令的全称。
      osname=`uname -s`
      #2. 可以在case的条件中添加更多的操作系统名称。
      case $osname in
      "Linux")
          PING=/usr/sbin/ping ;;
      "FreeBSD")
          PING=/sbin/ping ;;
      "SunOS")
          PING=/usr/sbin/ping ;;
      *)
          ;;
      esac
      CTRL+D
      /> . ./test5.sh
      /> echo $PING
      /usr/sbin/ping
    
六、获取当前时间距纪元时间(1970年1月1日)所经过的天数：

在获取两个时间之间的差值时，需要考虑很多问题，如闰年、月份中不同的天数等。然而如果我们能够确定两个时间点之间天数的差值，那么再计算时分秒的差值时就非常简单了。在系统提供的C语言函数中，获取的时间值是从1970年1月1日0点到当前时间所流经的秒数，如果我们基于此计算两个时间之间天数的差值，将会大大简化我们的计算公式。

      /> cat > test6.sh
      #!/bin/sh
      #1. 将date命令的执行结果(秒 分 小时 日 月 年)赋值给数组变量DATE。
      declare -a DATE=(`date +"%S %M %k %d %m %Y"`)
      #2. 为了提高效率，这个直接给出1970年1月1日到新纪元所流经的天数常量。
      epoch_days=719591
      #3. 从数组中提取各个时间部分值。
      year=${DATE[5]}
      month=${DATE[4]}
      day=${DATE[3]}
      hour=${DATE[2]}
      minute=${DATE[1]}
      second=${DATE[0]}
      #4. 当月份值为1或2的时候，将月份变量的值加一，否则将月份值加13，年变量的值减一，这样做主要是因为后面的公式中取月平均天数时的需要。
      if [ $month -gt 2 ]; then
          month=$((month+1))
      else
          month=$((month+13))
          year=$((year-1))
      fi
      #5. year变量参与的运算是需要考虑闰年问题的，该问题可以自行去google。
      #6. month变量参与的运算主要是考虑月平均天数。
      #7. 计算结果为当前日期距新世纪所流经的天数。
      today_days=$(((year*365)+(year/4)-(year/100)+(year/400)+(month*306001/10000)+day))
      #8. 总天数减去纪元距离新世纪的天数即可得出我们需要的天数了。
      days_since_epoch=$((today_days-epoch_days))
      echo $days_since_epoch
      seconds_since_epoch=$(((days_since_epoch*86400)+(hour*3600)+(minute*60)+second))
      echo $seconds_since_epoch
      CTRL+D
      /> . ./test6.sh
      15310
      1322829080

      需要说明的是，推荐将该脚本的内容放到一个函数中，以便于我们今后计算类似的时间数据时使用。
 

七、非直接引用变量：

      在Shell中提供了三种为标准(直接)变量赋值的方式：
      1. 直接赋值。
      2. 存储一个命令的输出。
      3. 存储某类型计算的结果。
      然而这三种方式都是给已知变量名的变量赋值，如name=Stephen。但是在有些情况下，变量名本身就是动态的，需要依照运行的结果来构造变量名，之后才是为该变量赋值。这种变量被成为动态变量，或非直接变量。
      /> cat > test7.sh
      #!/bin/sh
      work_dir=`pwd`
      #1. 由于变量名中不能存在反斜杠，因此这里需要将其替换为下划线。
      #2. work_dir和file_count两个变量的变量值用于构建动态变量的变量名。
      work_dir=`echo $work_dir | sed 's/\//_/g'`
      file_count=`ls | wc -l`
      #3. 输出work_dir和file_count两个变量的值，以便确认这里的输出结果和后面构建的命令名一致。
      echo "work_dir = " $work_dir
      echo "file_count = " $file_count
      #4. 通过eval命令进行评估，将变量名展开，如${work_dir}和$file_count，并用其值将其替换，如果不使用eval命令，将不会完成这些展开和替换的操作。最后为动态变量赋值。
      eval BASE${work_dir}_$file_count=$(ls $(pwd) | wc -l)
      #5. 先将echo命令后面用双引号扩住的部分进行展开和替换，由于是在双引号内，仅完成展开和替换操作即可。
      #6. echo命令后面的参数部分，先进行展开和替换，使其成为$BASE_root_test_1动态变量，之后在用该变量的值替换该变量本身作为结果输出。
      eval echo "BASE${work_dir}_$file_count = " '$BASE'${work_dir}_$file_count
      CTRL+D
      /> . ./test7.sh
      work_dir =  _root_test
      file_count =  1
      BASE_root_test_1 = 1
    
八、在循环中使用管道的技巧：

      在Bash Shell中，管道的最后一个命令都是在子Shell中执行的。这意味着在子Shell中赋值的变量对父Shell是无效的。所以当我们将管道输出传送到一个循环结构，填入随后将要使用的变量，那么就会产生很多问题。一旦循环完成，其所依赖的变量就不存在了。
      /> cat > test8_1.sh
      #!/bin/sh
      #1. 先将ls -l命令的结果通过管道传给grep命令作为管道输入。
      #2. grep命令过滤掉包含total的行，之后再通过管道将数据传给while循环。
      #3. while read line命令从grep的输出中读取数据。注意，while是管道的最后一个命令，将在子Shell中运行。
      ls -l | grep -v total | while read line
      do
          #4. all变量是在while块内声明并赋值的。
          all="$all $line"
          echo $line
      done
      #5. 由于上面的all变量在while内声明并初始化，而while内的命令都是在子Shell中运行，包括all变量的赋值，因此该变量的值将不会传递到while块外，因为块外地命令是它的父Shell中执行。
      echo "all = " $all
      CTRL+D
      /> ./test8_1.sh
      -rw-r--r--.  1 root root 193 Nov 24 11:25 outfile
      -rwxr-xr-x. 1 root root 284 Nov 24 10:01 test7.sh
      -rwxr-xr-x. 1 root root 108 Nov 24 12:48 test8_1.sh
      all =

      为了解决该问题，我们可以将while之前的命令结果先输出到一个临时文件，之后再将该临时文件作为while的重定向输入，这样while内部和外部的命令都将在同一个Shell内完成。
      /> cat > test8_2.sh
      #!/bin/sh
      #1. 这里我们已经将命令的结果重定向到一个临时文件中。
      ls -l | grep -v total > outfile
      while read line
      do
          #2. all变量是在while块内声明并赋值的。
          all="$all $line"
          echo $line
          #3. 通过重定向输入的方式，将临时文件中的内容传递给while循环。
      done < outfile
      #4. 删除该临时文件。
      rm -f outfile
      #5. 在while块内声明和赋值的all变量，其值在循环外部仍然有效。
      echo "all = " $all
      CTRL+D
      /> ./test8_2.sh
      -rw-r--r--.  1 root root   0 Nov 24 12:58 outfile
      -rwxr-xr-x. 1 root root 284 Nov 24 10:01 test7.sh
      -rwxr-xr-x. 1 root root 140 Nov 24 12:58 test8_2.sh
      all =  -rwxr-xr-x. 1 root root 284 Nov 24 10:01 test7.sh -rwxr-xr-x. 1 root root 135 Nov 24 13:16 test8_2.sh

      上面的方法只是解决了该问题，然而却带来了一些新问题，比如临时文件的产生容易导致性能问题，以及在脚本异常退出时未能及时删除当前使用的临时文件，从而导致生成过多的垃圾文件等。下面将再介绍一种方法，该方法将同时解决以上两种方法同时存在的问题。该方法是通过HERE-Document的方式来替代之前的临时文件方法。
      /> cat > test8_3.sh
      #!/bin/sh
      #1. 将命令的结果传给一个变量    
      OUTFILE=`ls -l | grep -v total`
      while read line
      do
          all="$all $line"
          echo $line
      done <<EOF
      #2. 将该变量作为该循环的HERE文档输入。
      $OUTFILE
      EOF
      #3. 在循环外部输出循环内声明并初始化的变量all的值。
      echo "all = " $all
      CTRL+D
      /> ./test8_3.sh
      -rwxr-xr-x. 1 root root 284 Nov 24 10:01 test7.sh
      -rwxr-xr-x. 1 root root 135 Nov 24 13:16 test8_3.sh
      all =  -rwxr-xr-x. 1 root root 284 Nov 24 10:01 test7.sh -rwxr-xr-x. 1 root root 135 Nov 24 13:16 test8_3.sh
    
九、自链接脚本：

      通常而言，我们是通过脚本的命令行选项来确定脚本的不同行为，告诉它该如何操作。这里我们将介绍另外一种方式来完成类似的功能，即通过脚本的软连接名来帮助脚本决定其行为。
      /> cat > test9.sh
      #!/bin/sh
      #1. basename命令将剥离脚本的目录信息，只保留脚本名，从而确保在相对路径的模式下执行也没有任何差异。
      #2. 通过sed命令过滤掉脚本的扩展名。
      dowhat=`basename $0 | sed 's/\.sh//'`
      #3. 这里的case语句只是为了演示方便，因此模拟了应用场景，在实际应用中，可以为不同的分支执行不同的操作，或将某些变量初始化为不同的值和状态。
      case $dowhat in
      test9)
          echo "I am test9.sh"
          ;;
      test9_1)
          echo "I am test9_1.sh."
          ;;
      test9_2)
          echo "I am test9_2.sh."
          ;;
      *)
          echo "You are illegal link file."
          ;;
      esac
      CTRL+D
      /> chmod a+x test9.sh
      /> ln -s test9.sh test9_1.sh
      /> ln -s test9.sh test9_2.sh
      /> ls -l
      lrwxrwxrwx. 1 root root   8 Nov 24 14:32 test9_1.sh -> test9.sh
      lrwxrwxrwx. 1 root root   8 Nov 24 14:32 test9_2.sh -> test9.sh
      -rwxr-xr-x. 1 root root 235 Nov 24 14:35 test9.sh
      /> ./test9.sh
      I am test9.sh.
      /> ./test9_1.sh
      I am test9_1.sh.
      /> ./test9_2.sh
      I am test9_2.sh.

十、Here文档的使用技巧：

      在命令行交互模式下，我们通常希望能够直接输入更多的信息，以便当前的命令能够完成一定的自动化任务，特别是对于那些支持自定义脚本的命令来说，我们可以将脚本作为输入的一部分传递给该命令，以使其完成该自动化任务。
      #1. 通过sqlplus以dba的身份登录Oracle数据库服务器。
      #2. 在通过登录后，立即在sqlplus中执行oracle的脚本CreateMyTables和CreateMyViews。
      #3. 最后执行sqlplus的退出命令，退出sqlplus。自动化工作完成。
      /> sqlplus "/as sysdba" <<-SQL
      > @CreateMyTables
      > @CreateMyViews
      > exit
      > SQL
         
十一、获取进程的运行时长(单位: 分钟)：

      在进程监控脚本中，我们通常需要根据脚本的参数来确定有哪些性能参数将被收集，当这些性能参数大于最高阈值或小于最低阈值时，监控脚本将根据实际的情况，采取预置的措施，如邮件通知、直接杀死进程等，这里我们给出的例子是收集进程运行时长性能参数。
      ps命令的etime值将给出每个进程的运行时长，其格式主要为以下三种：
      1. minutes:seconds，如20:30
      2. hours:minutes:seconds，如1:20:30
      3. days-hours:minute:seconds，如2-18:20:30
      该脚本将会同时处理这三种格式的时间信息，并最终转换为进程所流经的分钟数。
      /> cat > test11.sh
      #!/bin/sh
      #1. 通过ps命令获取所有进程的pid、etime和comm数据。
      #2. 再通过grep命令过滤，只获取init进程的数据记录，这里我们可以根据需要替换为自己想要监控的进程名。
      #3. 输出结果通常为：1 09:42:09 init
      pid_string=`ps -eo pid,etime,comm | grep "init" | grep -v grep`
      #3. 从这一条记录信息中抽取出etime数据，即第二列的值09:42:09，并赋值给exec_time变量。
      exec_time=`echo $pid_string | awk '{print $2}'`
      #4. 获取exec_time变量的时间组成部分的数量，这里是3个部分，即时:分:秒，是上述格式中的第二种。
      time_field_count=`echo $exec_time | awk -F: '{print NF}'`
      #5. 从exec_time变量中直接提取分钟数，即倒数第二列的数据(42)。
      count_of_minutes=`echo $exec_time | awk -F: '{print $(NF-1)}'`
    
      #6. 判断当前exec_time变量存储的时间数据是属于以上哪种格式。
      #7. 如果是第一种，那么天数和小时数均为0。
      #8. 如果是后两种之一，则需要继续判断到底是第一种还是第二种，如果是第二种，其小时部分将不存在横线(-)分隔符分隔天数和小时数，否则需要将这两个时间字段继续拆分，以获取具体的天数和小时数。对于第二种，天数为0.
      if [ $time_field_count -lt 3 ]; then
          count_of_hours=0
          count_of_days=0
      else
          count_of_hours=`echo $exec_time | awk -F: '{print $(NF-2)}'`
          fields=`echo $count_of_hours | awk -F- '{print NF}'`
          if [ $fields -ne 1 ]; then
              count_of_days=`echo $count_of_hours | awk -F- '{print $1}'`
              count_of_hours=`echo $count_of_hours | awk -F- '{print $2}'`
          else
              count_of_days=0
          fi
      fi
      #9. 通过之前代码获取的各个字段值，计算出该进程实际所流经的分钟数。
      #10. bc命令是计算器命令，可以将echo输出的数学表达式计算为最终的数字值。
      elapsed_minutes=`echo "$count_of_days*1440+$count_of_hours*60+$count_of_minutes" | bc`
      echo "The elapsed minutes of init process is" $elapsed_minutes "minutes."
      CTRL+D
      /> ./test11.sh
      The elapsed minutes of init process is 577 minutes.
    
十二、模拟简单的top命令：
    
      这里用脚本实现了一个极为简单的top命令。为了演示方便，我们在脚本中将很多参数都写成硬代码，你可以根据需要更换这些参数，或者用更为灵活的方式替换现有的实现。
      /> cat > test12.sh
      #!/bin/sh
      #1. 将ps命令的title赋值给一个变量，这样在每次输出时，直接打印该变量即可。
      header=`ps aux | head -n 1`
      #2. 这里是一个无限循环，等价于while true
      #3. 每次循环先清屏，之后打印uptime命令的输出。
      #4. 输出ps的title。
      #5. 这里需要用sed命令删除ps的title行，以避免其参与sort命令的排序。
      #6. sort先基于CPU%倒排，再基于owner排序，最后基于pid排序，最后再将结果输出给head命令，仅显示前20行的数据。
      #7. 每次等待5秒后刷新一次。
     while :
      do
          clear
          uptime
          echo "$header"
          ps aux | sed -e 1d | sort -k3nr -k1,1 -k2n | head -n 20
          sleep 5
      done
      CTRL+D    
      /> ./test12.sh
      21:55:07 up 13:42,  2 users,  load average: 0.00, 0.00, 0.00
      USER       PID %CPU %MEM    VSZ   RSS   TTY      STAT START   TIME   COMMAND
      root      6408     2.0      0.0   4740   932   pts/2    R+    21:45     0:00   ps aux
      root      1755     0.2      2.0  96976 21260   ?        S      08:14     2:08   nautilus
      68        1195     0.0      0.4   6940   4416    ?        Ss    08:13     0:00   hald
      postfix   1399    0.0      0.2  10312  2120    ?        S      08:13     0:00   qmgr -l -t fifo -u
      postfix   6021    0.0      0.2  10244  2080    ?        S      21:33     0:00   pickup -l -t fifo -u
      root         1       0.0      0.1   2828   1364    ?        Ss     08:12    0:02   /sbin/init
      ... ...



十三、格式化输出指定用户的当前运行进程：

      在这个例子中，我们通过脚本参数的形式，将用户列表传递给该脚本，脚本在读取参数后，以树的形式将用户列表中用户的所属进程打印出来。
      /> cat > test13.sh
      #!/bin/sh
      #1. 循环读取脚本参数，构造egrep可以识别的用户列表变量(基于grep的扩展正则表达式)。
      #2. userlist变量尚未赋值，则直接使用第一个参数为它赋值。
      #3. 如果已经赋值，且脚本参数中存在多个用户，这里需要在每个用户名之间加一个竖线，在egrep中，竖线是分割的元素之间是或的关系。
      #4. shift命令向左移动一个脚本的位置参数，这样可以使循环中始终操作第一个参数。
      while [ $# -gt 0 ]
      do
          if [ -z "$userlist" ]; then
              userlist="$1"
          else
              userlist="$userlist|$1"
          fi
           shift
      done
      #5. 如果没有用户列表，则搜索所有用户的进程。
      #6. "^ *($userlist) ": 下面的调用方式，该正则的展开形式为"^ *(root|avahi|postfix|rpc|dbus) "。其含义为，以0个或多个空格开头，之后将是root、avahi、postfix、rpc或dbus之中的任何一个字符串，后面再跟随一个空格。
      if [ -z "$userlist" ]; then
          userlist="."
      else
          userlist="^ *($userlist) "
      fi
      #7. ps命令输出所有进程的user和命令信息，将结果传递给sed命令，sed将删除ps的title部分。
      #8. egrep过滤所有进程记录中，包含指定用户列表的进程记录，再将过滤后的结果传递给sort命令。
      #9. sort命令中的-b选项将忽略前置空格，并以user，再以进程名排序，将结果传递个uniq命令。
      #10.uniq命令将合并重复记录，-c选项将会使每条记录前加重复的行数。
      #11.第二个sort将再做一次排序，先以user，再以重复计数由大到小，最后以进程名排序。将结果传给awk命令。
      #12.awk命令将数据进行格式化，并删除重复的user。
      ps -eo user,comm | sed -e 1d | egrep "$userlist" |
          sort -b -k1,1 -k2,2 | uniq -c | sort -b -k2,2 -k1nr,1 -k3,3 |
              awk ' { user = (lastuser == $2) ? " " : $2;
                        lastuser = $2;
                        printf("%-15s\t%2d\t%s\n",user,$1,$3)
              }'
      CTRL+D
      /> ./test13.sh root avahi postfix rpc dbus
      avahi             2      avahi-daemon
      dbus             1      dbus-daemon
      postfix          1      pickup
                          1      qmgr
      root              5      mingetty
                          3      udevd
                          2      sort
                          2      sshd
      ... ...
      rpc               1      rpcbind

十四、用脚本完成which命令的基本功能：

      我们经常会在脚本中调用其他的应用程序，为了保证脚本具有更好的健壮性，以及错误提示的准确性，我们可能需要在执行前验证该命令是否存在，或者说是否可以被执行。这首先要确认该命令是否位于PATH变量包含的目录中，再有就是该文件是否为可执行文件。
      /> cat > test14.sh
      #!/bin/sh
      #1. 该函数用于判断参数1中的命令是否位于参数2所包含的目录列表中。需要说明的是，函数里面的$1和$2是指函数的参数，而不是脚本的参数，后面也是如此。
      #2. cmd=$1和path=$2，将参数赋给有意义的变量名，是一个很好的习惯。
      #3. 由于PATH环境变量中，目录之间的分隔符是冒号，因此这里需要临时将IFS设置为冒号，函数结束后再还原。
      #4. 在for循环中，逐个变量目录列表中的目录，以判断该命令是否存在，且为可执行程序。
      isInPath() {
          cmd=$1        path=$2      result=1
          oldIFS=$IFS   IFS=":"
          for dir in $path
          do
              if [ -x $dir/$cmd ]; then
                  result=0
              fi
          done
          IFS=oldifs
          return $result
      }
      #5. 检查命令是否存在的主功能函数，先判断是否为绝对路径，即$var变量的第一个字符是否为/，如果是，再判断它是否有可执行权限。
      #6. 如果不是绝对路径，通过isInPath函数判断是否该命令在PATH环境变量指定的目录中。
      checkCommand() {
          var=$1
          if [ ! -z "$var" ]; then
              if [ "${var:0:1}" = "/" ]; then
                  if [ ! -x $var ]; then
                      return 1
                  fi
              elif ! isInPath $var $PATH ; then
                  return 2
              fi
          fi
      }
      #7. 脚本参数的合法性验证。
      if [ $# -ne 1 ]; then
          echo "Usage: $0 command" >&2;
      fi
      #8. 根据返回值打印不同的信息。我们可以在这里根据我们的需求完成不同的工作。
      checkCommand $1
      case $? in
      0) echo "$1 found in PATH." ;;
      1) echo "$1 not found or not executable." ;;
      2) echo "$1 not found in PATH." ;;
      esac
      exit 0
      CTRL+D
      /> ./test14.sh echo
      echo found in PATH.
      /> ./test14.sh MyTest
      MyTest not found in PATH.
      /> ./test14.sh /bin/MyTest
      /bin/MyTest not found or not executable.


十五、验证输入信息是否合法：

      这里给出的例子是验证用户输入的信息是否都是数字和字母。需要说明的是，之所以将其收集到该系列中，主要是因为它实现的方式比较巧妙。
      /> cat > test15.sh
      #!/bin/sh
      echo -n "Enter your input: "
      read input
      #1. 事实上，这里的巧妙之处就是先用sed替换了非法部分，之后再将替换后的结果与原字符串比较。这种写法也比较容易扩展。    
      parsed_input=`echo $input | sed 's/[^[:alnum:]]//g'`
      if [ "$parsed_input" != "$input" ]; then
          echo "Your input must consist of only letters and numbers."
      else
          echo "Input is OK."
      fi
      CTRL+D
      /> ./test15.sh
      Enter your input: hello123
      Input is OK.
      /> ./test15.sh
      Enter your input: hello world
      Your input must consist of only letters and numbers.

十六、整数验证：

      整数的重要特征就是只是包含数字0到9和负号(-)。
      /> cat > test16.sh
      #!/bin/sh
      #1. 判断变量number的第一个字符是否为负号(-)，如果只是则删除该负号，并将删除后的结果赋值给left_number变量。
      #2. "${number#-}"的具体含义，可以参考该系列博客中"Linux Shell常用技巧(十一)"，搜索关键字"变量模式匹配运算符"即可。
      number=$1
      if [ "${number:0:1}" = "-" ]; then
          left_number="${number#-}"
      else
          left_number=$number
      fi
      #3. 将left_number变量中所有的数字都替换掉，因此如果返回的字符串变量为空，则表示left_number所包含的字符均为数字。
      nodigits=`echo $left_number | sed 's/[[:digit:]]//g'`
      if [ "$nodigits" != "" ]; then
          echo "Invalid number format!"
      else
          echo "You are valid number."
      fi
      CTRL+D
      /> ./test16.sh -123
      You are valid number.
      /> ./test16.sh 123e
      Invalid number format!
    
十七、判断指定的年份是否为闰年：

      这里我们先列出闰年的规则:
      1. 不能被4整除的年一定不是闰年；
      2. 可以同时整除4和400的年一定是闰年；
      3. 可以整除4和100，但是不能整除400的年，不是闰年；
      4. 其他可以整除的年都是闰年。
      #!/bin/sh    
      year=$1
      if [ "$((year % 4))" -ne 0 ]; then
          echo "This is not a leap year."
          exit 1
      elif [ "$((year % 400))" -eq 0 ]; then
          echo "This is a leap year."
          exit 0
      elif [ "$((year % 100))" -eq 0 ]; then
          echo "This is not a leap year."
          exit 1
      else
          echo "This is a leap year."
          exit 0
      fi
      CTRL+D
      /> ./test17.sh 1933
      This is not a leap year.
      /> ./test17.sh 1936
      This is a leap year.
            
十八、将单列显示转换为多列显示：

      我们经常会在显示时将单行的输出，格式化为多行的输出，通常情况下，为了完成该操作，我们将加入更多的代码，将输出的结果存入数组或临时文件，之后再重新遍历它们，从而实现单行转多行的目的。在这里我们介绍一个使用xargs命令的技巧，可以用更简单、更高效的方式来完成该功能。    
      /> cat > test18.sh
      #!/bin/sh
      #1. passwd文件中，有可能在一行内出现一个或者多个空格字符，因此在直接使用cat命令的结果时，for循环会被空格字符切开，从而导致一行的文本被当做多次for循环的输入，这样我们不得不在sed命令中，将cat输出的每行文本进行全局替换，将空格字符替换为%20。事实上，我们当然可以将cat /etc/passwd的输出以管道的形式传递给cut命令，这里之所以这样写，主要是为了演示一旦出现类似的问题该如果巧妙的处理。
      #2. 这里将for循环的输出以管道的形式传递给sort命令，sort命令将基于user排序。
      #3. -xargs -n 2是这个技巧的重点，它将sort的输出进行合并，-n选项后面的数字参数将提示xargs命令将多少次输出合并为一次输出，并传递给其后面的命令。在本例中，xargs会将从sort得到的每两行数据合并为一行，中间用空格符分离，之后再将合并后的数据传递给后面的awk命令。事实上，对于awk而言，你也可以简单的认为xargs减少了对它(awk)的一半调用。
      #4. 如果打算在一行内显示3行或更多的行，可以将-n后面的数字修改为3或其它更高的数字。你还可以修改awk中的print命令，使用更为复杂打印输出命令，以得到更为可读的输出效果。
      for line in `cat /etc/passwd | sed 's/ /%20/g'`
      do
          user=`echo $line | cut -d: -f1`
          echo $user
      done | \
          sort -k1,1 | \
          xargs -n 2 | \
          awk '{print $1, $2}'
      CTRL+D
      /> ./test18.sh
      abrt adm
      apache avahi
      avahi-autoipd bin
      daemon daihw
      dbus ftp
      games gdm
      gopher haldaemon
      halt lp
      mail nobody
      ntp operator
      postfix pulse
      root rtkit
      saslauth shutdown
      sshd sync
      tcpdump usbmuxd
      uucp vcsa

十九、将文件的输出格式化为指定的宽度：

      在这个技巧中，不仅包含了如何获取和文件相关的详细信息，如行数，字符等，而且还可以让文件按照指定的宽度输出。这种应用在输出帮助信息、License相关信息时还是比较有用的。
      /> cat > test19.sh
      #!/bin/sh
      #1. 这里我们将缺省宽度设置为75，如果超过该宽度，将考虑折行显示，否则直接在一行中全部打印输出。这里只是为了演示方便，事实上，你完全可以将该值作为脚本或函数的参数传入，那样你将会得到更高的灵活性。    
      my_width=75
      #2. for循环的读取列表来自于脚本的参数。
      #3. 在获取lines和chars变量时，sed命令用于过滤掉多余的空格字符。
      #4. 在if的条件判断中${#line}用于获取line变量的字符长度，这是Shell内置的规则。
      #5. fmt -w 80命令会将echo输出的整行数据根据其命令选项指定的宽度(80个字符)进行折行显示，再将折行后的数据以多行的形式传递给sed命令。
      #6. sed在收到fmt命令的格式化输出后，将会在折行后的第一行头部添加两个空格，在其余行的头部添加一个加号和一个空格以表示差别。
      for input; do
          lines=`wc -l < $input | sed 's/ //g'`
          chars=`wc -c < $input | sed 's/ //g'`
          owner=`ls -l $input | awk '{print $3}'`
          echo "-------------------------------------------------------------------------------"
          echo "File $input ($lines lines, $chars characters, owned by $owner):"
          echo "-------------------------------------------------------------------------------"
          while read line; do
              if [ ${#line} -gt $my_width ]; then
                  echo "$line" | fmt -w 80 | sed -e '1s/^/  /' -e '2,$s/^/+ /'
              else
                  echo "  $line"
              fi
          done < $input
          echo "-------------------------------------------------------------------------------"
      done | more
      CTRL+D
      /> ./test19.sh testfile
      -------------------------------------------------------------------------------
      File testfile.dat (3 lines, 645 characters, owned by root):
      -------------------------------------------------------------------------------
         The PostgreSQL Global Development Group today released updates for all
      + active branches of the PostgreSQL object-relational database system,
      + including versions 9.1.2, 9.0.6, 8.4.10, 8.3.17 and 8.2.23. Users of any of
      + the several affected features in this release, including binary replication,
      + should update their PostgreSQL installations as soon as possible.
         This is also the last update for PostgreSQL 8.2, which is now End-Of-Life
      + (EOL). Users of version 8.2 should plan to upgrade their PostgreSQL
      + installations to 8.3 or later within the next couple of months. For more
      + information, see our Release Support Policy.
         This is just a test file.
      -------------------------------------------------------------------------------
        
二十、监控指定目录下磁盘使用空间过大的用户：

      在将Linux用作文件服务器时，所有的注册用户都可以在自己的主目录下存放各种类型和大小的文件。有的时候，有些用户的占用空间可能会明显超过其他人，这时就需要管理员可以及时发现这一异常使用状况，并根据实际情况作出应对处理。
      /> cat > test20.sh
      #!/bin/sh
      #1. 该脚本仅用于演示一种处理技巧，其中很多阈值都是可以通过脚本参来初始化的，如limited_qutoa和dirs等变量。
      limited_quota=200
      dirs="/home /usr /var"
      #2. 以冒号作为分隔符，截取passwd文件的第一和第三字段，然后将输出传递给awk命令。
      #3. awk中的$2表示的是uid，其中1-99是系统保留用户，>=100的uid才是我们自己创建的用户，awk通过print输出所有的用户名给for循环。
      #4. 注意echo命令的输出是由八个单词构成，同时由于-n选项，echo命令并不输出换行符。
      #5. 之所以使用find命令，也是为了考虑以点(DOT)开头的隐藏文件。这里的find将在指定目录列表内，搜索指定用户的，类型为普通文件的文件。并通过-ls选项输出找到文件的详细信息。其中输出的详细信息的第七列为文件大小列。
      #6. 通过awk命令累加find输出的第七列，最后再在自己的END块中将sum的值用MB计算并输出。该命令的输出将会与上面echo命令的输出合并作为for循环的输出传递给后面的awk命令。这里需要指出的是，该awk的输出就是后面awk命令的$9，因为echo仅仅输出的8个单词。
      #7. 从for循环管道获取数据的awk命令，由于awk命令执行的动作是用双引号括起的，所以表示域字段的变量的前缀$符号，需要用\进行转义。变量$limited_quota变量将会自动完成命令替换，从而构成该awk命令的最终动作参数。
      for name in `cut -d: -f1,3 /etc/passwd | awk -F: '$2 > 99 {print $1}'`
      do
          echo -n "User $name exceeds disk quota. Disk Usage is: "
          find $dirs -user $name -type f -ls |\
                awk '{ sum += $7 } END { print sum / (1024*1024) " MB" }'
      done | awk "\$9 > $limited_quota { print \$0 }"
      CTRL+D
      /> ./test20.sh    

二十一、编写一个更具可读性的df命令输出脚本：

      这里我们将以awk脚本的方式来实现df -h的功能。
      /> cat > test21.sh
      #!/bin/sh
      #1. $$表示当前Shell进程的pid。    
      #2. trap信号捕捉是为了保证在Shell正常或异常退出时，仍然能够将该脚本创建的临时awk脚本文件删除。
      awk_script_file="/tmp/scf_tmp.$$"
      trap "rm -f $awk_script_file" EXIT
      #3. 首先需要说明的是，'EOF'中的单引号非常重要，如果忽略他将无法通过编译，这是因为awk的命令动作必须要用单引号扩住。
      #4. awk脚本的show函数中，int(mb * 100) / 100这个技巧是为了保证输出时保留小数点后两位。
      cat << 'EOF' > $awk_script_file
      function show(size) {
          mb = size / 1024;
          int_mb = (int(mb * 100)) / 100;
          gb = mb / 1024;
          int_gb = (int(gb * 100)) / 100;
          if (substr(size,1,1) !~ "[0-9]" || substr(size,2,1) !~ "[0-9]") {
              return size;
          } else if (mb < 1) {
              return size "K";
          } else if (gb < 1) {
              return int_mb "M";
          } else {
              return int_gb "G";
          }
      }
      #5. 在BEGIN块中打印重定义的输出头信息。
      BEGIN {
            printf "%-20s %7s %7s %7s %8s %s\n","FileSystem","Size","Used","Avail","Use%","Mounted"
      }
      #6. !/Filesystem/ 表示过滤掉包含Filesystem的行，即df输出的第一行。其余行中，有个域字段可以直接使用df的输出，有的需要通过show函数的计算，以得到更为可读的显示结果。
      !/Filesystem/ {
          size = show($2);
          used = show($3);
          avail = show($4);
          printf "%-20s %7s %7s %7s %8s %s\n",$1,size,used,avail,$5,$6
      }
      EOF
      df -k | awk -f $awk_script_file
      CTRL+D
      /> ./test12.sh
      FileSystem              Size       Used      Avail     Use%   Mounted
      /dev/sda2              3.84G    2.28G     1.36G      63%   /
      tmpfs                 503.57M     100K 503.47M        1%   /dev/shm
      /dev/sda1             48.41M  35.27M  10.63M      77%   /boot
      /dev/sda3              14.8G 171.47M  13.88G        2%   /home
    
二十二、编写一个用于添加新用户的脚本：

      之所以在这里选择这个脚本，没有更多的用意，只是感觉这里的有些技巧和常识还是需要了解的，如/etc/passwd、/etc/shadow、/etc/group的文件格式等。
      /> cat > test22.sh
      #!/bin/sh
      #1. 初始化和用户添加相关的变量。    
      passwd_file="/etc/passwd"
      shadow_file="/etc/shadow"
      group_file="/etc/group"
      home_root_dir="/home"
      #2. 只有root用户可以执行该脚本。    
      if [ "$(whoami)" != "root" ]; then
          echo "Error: You must be roor to run this command." >&2
          exit 1
      fi
    
      echo "Add new user account to $(hostname)"
      echo -n "login: "
      read login
      #3. 去唯一uid，即当前最大uid值加一。
      uid=`awk -F: '{ if (big < $3 && $3 < 5000) big = $3 } END {print big + 1}' $passwd_file`
      #4. 设定新用户的主目录变量
      home_dir=$home_root_dir/$login
      gid=$uid
      #5. 提示输入和创建新用户相关的信息，如用户全名和主Shell。
      echo -n "full name: "
      read fullname
      echo -n "shell: "
      read shell
      #6. 将输入的信息填充到passwd、group和shadow三个关键文件中。
      echo "Setting up account $login for $fullname..."
      echo ${login}:x:${uid}:${gid}:${fullname}:${home_dir}:$shell >> $passwd_file
      echo ${login}:*:11647:0:99999:7::: >> $shadow_file
      echo "${login}:x:${gid}:$login" >> $group_file
      #7. 创建主目录，同时将新用户的profile模板拷贝到新用户的主目录内。
      #8. 设定该主目录的权限，再将其下所有文件的owner和group设置为新用户。
      #9. 为新用户设定密码。
      mkdir $home_dir
      cp -R /etc/skel/.[a-zA-Z]* $home_dir
      chmod 755 $home_dir
      find $home_dir -print | xargs chown ${login}:${login}
      passwd $login
      exit 0
      CTRL+D
      /> ./test22.sh
      Add new user account to bogon
      login: stephen
      full name: Stephen Liu
      shell: /bin/shell
      Setting up account stephen for Stephen Liu...
      Changing password for user stephen.
      New password:
      Retype new password:
      passwd: all authentication tokens updated successfully.
    
二十三、kill指定用户或指定终端的用户进程：

      这是一台运行Oracle数据库的Linux服务器，现在我们需要对Oracle做一些特殊的优化工作，由于完成此项工作需要消耗更多的系统资源，因此我们不得不杀掉一些其他用户正在运行的进程，以便节省出更多的系统资源，让本次优化工作能够尽快完成。
      /> cat > test23.sh
      #!/bin/sh
      user=""
      tty=""
      #1. 通过读取脚本的命令行选项获取要kill的用户或终端。-t后面的参数表示终端，-u后面的参数表示用户。这两个选项不能同时使用。
      #2. case中的代码对脚本选项进行验证，一旦失败则退出脚本。
      while getopts u:t: opt; do
          case $opt in
          u) if [ "$tty" != "" ]; then
                 echo "-u and -t can not be set at one time."
                 exit 1
              fi
              user=$OPTARG
              ;;
          t)  if [ "$user" != "" ]; then
                 echo "-u and -t can not be set at one time."
                 exit 1
              fi
              tty=$OPTARG
              ;;
          ?) echo "Usage: $0 [-u user|-t tty]" >&2
              exit 1
          esac
      done
      #3. 如果当前选择的是基于终端kill，就用$tty来过滤ps命令的输出，否则就用$user来过滤ps命令的输出。
      #4. awk命令将仅仅打印出pid字段，之后传递给sed命令，sed命令删除ps命令输出的头信息，仅保留后面的进程pids作为输出，并初始化pids数组。
      if [ ! -z "$tty" ]; then
          pids=$(ps cu -t $tty | awk "{print \$2}" | sed '1d')
      else
          pids=$(ps cu -U $user | awk "{print \$2}" | sed '1d')
      fi
      #5. 判断数组是否为空，空则表示没有符合要求的进程，直接退出脚本。
      if [ -z "$pids" ]; then
          echo "No processes matches."
          exit 1
      fi
      #6. 遍历pids数组，逐个kill指定的进程。
      for pid in $pids; do
          echo "Killing process[pid = $pid]... ..."
          kill -9 $pid
      done
      exit 0
      CTRL+D
      /> ./test23.sh -t pts/1
      Killing process[pid = 11875]... ...
      Killing process[pid = 11894]... ...
      /> ./test23.sh -u stephen
      Killing process[pid = 11910]... ...
      Killing process[pid = 11923]... ...
        
二十四、判断用户输入(是/否)的便捷方法：

      对于有些交互式程序，经常需要等待用户的输入，以便确定下一步的执行流程。通常而言，当用户输入"Y/y/Yes/yes"时表示确认当前的行为，而当输入"N/n/No/no"时则表示否定当前的行为。基于这种规则，我们可以实现一个便捷确认方式，即只判断输入的首字母，如果为Y或y，表示确认，如为N或n，则为否定。
      /> cat > test24.sh
      #!/bin/sh
      echo -n "Please type[y/n/yes/no]: "
      read input
      #1. 先转换小写到大写，再通过cut截取第一个字符。
      ret=`echo $input | tr '[a-z]' '[A-Z]' | cut -c1`
    
      if [ $ret = "Y" ]; then
          echo "Your input is Y."
      elif [ $ret = "N" ]; then
          echo "Your input is N."
      else
          echo "Your input is error."
      fi
      CTRL+D
      /> ./test24.sh
      Please type[y/n/yes/no]: y
      Your input is Y.
      /> ./test24.sh
      Please type[y/n/yes/no]: n
      Your input is N.  

二十五、通过FTP下载指定的文件：

      相比于手工调用FTP命令下载文件，该脚本提供了更为方便的操作方式。
      /> cat > test25.sh
      #!/bin/sh
      #1. 测试脚本参数数量的有效性。    
      if [ $# -ne 2 ]; then
          echo "Usage: $0 ftp://... username" >&2
          exit 1
      fi
      #2. 获取第一个参数的前六个字符，如果不是"ftp://"，则视为非法FTP URL格式。这里cut的-c选项表示按照字符的方式截取第一到第六个字符。
      header=`echo $1 | cut -c1-6`
      if [ "$header" != "ftp://" ]; then
          echo "$0: Invalid ftp URL." >&2
          exit 1
      fi
      #3. 合法ftp URL的例子：ftp://ftp.myserver.com/download/test.tar
      #4. 针对上面的URL示例，cut命令通过/字符作为分隔符，这样第三个域字段表示server(ftp.myserver.com)。
      #5. 在截取filename时，cut命令也是通过/字符作为分隔符，但是"-f4-"将获取从第四个字段开始的后面所有字段(download/test.tar)。
      #6. 通过basename命令获取filename的文件名部分。
      server=`echo $1 | cut -d/ -f3`
      filename=`echo $1 | cut -d/ -f4-`
      basefile=`basename $filename`
      ftpuser=$2
      #7. 这里需要调用stty -echo，以便后面的密码输入不会显示，在输入密码之后，需要再重新打开该选项，以保证后面的输入可以恢复显示。
      #8. echo ""，是模拟一次换换。
      echo -n "Password for $ftpuser: "
      stty -echo
      read password
      stty echo
      echo ""
      #9. 通过HERE文档，批量执行ftp命令。
      echo ${0}: Downloading $baseile from server $server.
      ftp -n << EOF
      open $server
      user $ftpuser $password
      get $filename $basefile
      quit
      EOF
      #10.Shell内置变量$?表示上一个Shell进程的退出值，0表示成功执行，其余值均表示不同原因的失败。
      if [ $? -eq 0 ]; then
          ls -l $basefile
      fi
      exit 0
      CTRL+D
      /> ./test25.sh  ftp://ftp.myserver.com/download/test.tar stephen
      Password for stephen:
      ./test25.sh: Downloading from server ftp.myserver.com.
      -rwxr-xr-x. 1 root root 678 Dec  9 11:46 test.tar

二十六、文件锁定：

      在工业应用中，有些来自于工业设备的文件将会被放到指定的目录下，由于这些文件需要再被重新格式化后才能被更高层的软件进行处理。而此时负责处理的脚本程序极有可能是多个实例同时运行，因此这些实例之间就需要一定的同步，以避免多个实例同时操作一个文件而造成的数据不匹配等问题的发生。文件锁定命令可以帮助我们实现这一同步逻辑。
      /> cat > test26.sh
      #!/bin/sh
      #1. 这里需要先确认flock命令是否存在。
      if [ -z $(which flock) ]; then
          echo "flock doesn't exist."
          exit 1
      fi
      #2. flock中的-e选项表示对该文件加排它锁，-w选项表示如果此时文件正在被加锁，当前的flock命令将等待20秒，如果能锁住该文件，就继续执行，否则退出该命令。
      #3. 这里锁定的文件是/var/lock/lockfile1，-c选项表示，如果成功锁定，则指定其后用双引号括起的命令，如果是多个命令，可以用分号分隔。
      #4. 可以在两个终端同时启动该脚本，然后观察脚本的输出，以及lockfile1文件的内容。
      flock -e -w 20 /var/lock/lockfile1 -c "sleep 10;echo `date` | cat >> /var/lock/lockfile1"
      if [ $? -ne 0 ]; then
          echo "Fail."
          exit 1
      else
          echo "Success."
          exit 0
      fi
      CTRL+D
    
二十七、用小文件覆盖整个磁盘：

      假设我们现在遇到这样一个问题，公司的关键资料copy到测试服务器上了，在直接将其删除后，仍然担心服务器供应商可以将其恢复，即便是通过fdisk进行重新格式化，也仍然存在被恢复的风险，鉴于此，我们需要编写一个脚本，创建很多小文件(5MB左右)，之后不停在关键资料所在的磁盘中复制该文件，以使Linux的inode无法再被重新恢复，为了达到这里效果，我们需要先构造该文件，如：
      /> find . -name "*" > testfile
      /> ls -l testfile
      -rwxr-xr-x. 1 root root 5123678 Dec  9 11:46 testfile
      /> cat > test27.sh
      #!/bin/sh
      #1. 初始化计数器变量，其中max的值是根据当前需要填充的磁盘空间和testfile的大小计算出来的。
      counter=0
      max=2000000
      remainder=0
      #2. 每次迭代counter变量都自增一，以保证每次生成不同的文件。当该值大于最大值时退出。
      #3. 对计数器变量counter按1000取模，这样可以在每生成1000个文件时打印一次输出，以便看到覆盖的进度，输出时间则便于预估还需要多少时间可以完成。
      #4. 创建单独的、用于存放这些覆盖文件的目录。
      #5. 生成临时文件，如果写入失败打印出提示信息。
      while true
      do
          ((counter=counter+1))
          if [ #counter -ge $max ]; then
              break
          fi
          ((remainder=counter%1000))
          if [ $remainder -eq 0 ]; then
              echo -e "counter = $counter\t date = " $(date)
          fi
          mkdir -p /home/temp2
          cat < testfile > "/home/temp/myfiles.$counter"
          if [[ $? -ne 0 ]]; then
              echo "Failed to wrtie file."
              exit 1
          fi
      done
      echo "Done"
      CTRL+D
      /> ./test27.sh
      counter = 1000        Fri Dec  9 17:25:04 CST 2011
      counter = 2000        Fri Dec  9 17:25:24 CST 2011
      counter = 3000        Fri Dec  9 17:25:54 CST 2011
      ... ...
      与此同时，可以通过执行下面的命令监控磁盘空间的使用率。
      /> watch -n 2 'df -h'
      Every 2.0s: df -h                                       Fri Dec  9 17:31:56 2011
    
      Filesystem            Size   Used Avail Use% Mounted on
      /dev/sda2             3.9G  2.3G  1.4G  63% /
      tmpfs                  504M  100K  504M   1% /dev/shm
      /dev/sda1              49M   36M   11M  77% /boot
      /dev/sda3              15G  172M   14G   2% /home
      我们也可以在执行的过程中通过pidstat命令监控脚本进程的每秒读写块数。    
 
二十八、统计当前系统中不同运行状态的进程数量：

      在Linux系统中，进程的运行状态主要分为四种：运行时、睡眠、停止和僵尸。下面的脚本将统计当前系统中，各种运行状态的进程数量。
      /> cat > test28.sh
      #!/bin/sh
      #1. 初始化计数器变量，分别对应于运行时、睡眠、停止和僵尸。
      running=0
      sleeping=0
      stopped=0
      zombie=0
      #2. 在/proc目录下，包含很多以数字作为目录名的子目录，其含义为，每个数字对应于一个当前正在运行进程的pid，该子目录下包含一些文件用于描述与该pid进程相关的信息。如1表示init进程的pid。那么其子目录下的stat文件将包含和该进程运行状态相关的信息。
      #3. cat /proc/1/stat，通过该方式可以查看init进程的运行状态，同时也可以了解该文件的格式，其中第三个字段为进程的运行状态字段。
      #4. 通过let表达式累加各个计数器。
      for pid in /proc/[1-9]*
      do
          ((procs=procs+1))
          stat=`awk '{print $3}' $pid/stat`
          case $stat in
              R) ((running=runing+1));;
              S) ((sleeping=sleeping+1));;
              T) ((stopped=stopped+1));;
              Z) ((zombie=zombie+1));
          esac
      done
      echo -n "Process Count: "
      echo -e "Running = $running\tSleeping = $sleeping\tStopped = $stopped\tZombie = $zombie."
      CTRL+D
      /> ./test28.sh
      Process Count: Running = 0      Sleeping = 136  Stopped = 0     Zombie = 0.
    
二十九、浮点数验证：

     浮点数数的重要特征就是只是包含数字0到9、负号(-)和点(.)，其中负号只能出现在最前面，点(.)只能出现一次。
      /> cat > test29.sh
      #!/bin/sh
      #1. 之前的一个条目已经介绍了awk中match函数的功能，如果匹配返回匹配的位置值，否则返回0。
      #2. 对于Shell中的函数而言，返回0表示成功，其他值表示失败，该语义等同于Linux中的进程退出值。调用者可以通过内置变量$?获取返回值，或者作为条件表达式的一部分直接判断。
      validint() {
          ret=`echo $1 | awk '{start = match($1,/^-?[0-9]+$/); if (start == 0) print "1"; else print "0"}'`
          return $ret
      }
    
      validfloat() {
          fvalue="$1"
          #3. 判断当前参数中是否包含小数点儿。如果包含则需要将其拆分为整数部分和小数部分，分别进行判断。
          if [ ! -z  $(echo $fvalue | sed 's/[^.]//g') ]; then
              decimalpart=`echo $fvalue | cut -d. -f1`
              fractionalpart=`echo $fvalue | cut -d. -f2`
              #4. 如果整数部分不为空，但是不是合法的整型，则视为非法格式。
              if [ ! -z $decimalpart ]; then
                  if ! validint "$decimalpart" ; then
                      echo "decimalpart is not valid integer."
                      return 1
                  fi
              fi
              #5. 判断小数部分的第一个字符是否为-，如果是则非法。
              if [ "${fractionalpart:0:1}" = "-" ]; then
                  echo "Invalid floating-point number: '-' not allowed after decimal point." >&2
                  return 1
              fi
              #6. 如果小数部分不为空，同时也不是合法的整型，则视为非法格式。
              if [ "$fractionalpart" != "" ]; then
                  if ! validint "$fractionalpart" ; then
                      echo "fractionalpart is not valid integer."
                      return 1
                  fi
              fi
              #7. 如果整数部分仅为-，或者为空，如果此时小数部分也是空，则为非法格式。
              if [ "$decimalpart" = "-" -o -z "$decimalpart" ]; then
                  if [ -z $fractionalpart ]; then
                      echo "Invalid floating-point format." >&2
                      return 1
                  fi
              fi
          else
              #8. 如果当前参数仅为-，则视为非法格式。
              if [ "$fvalue" = "-" ]; then
                  echo "Invalid floating-point format." >&2
                  return 1
              fi
              #9. 由于参数中没有小数点，如果该值不是合法的整数，则为非法格式。
              if ! validint "$fvalue" ; then
                  echo "Invalid floating-point format." >&2
                  return 1
              fi
          fi
          return 0
      }    
      if validfloat $1 ; then
          echo "$1 is a valid floating-point value."
      fi
      exit 0
      CTRL+D
      /> ./test29.sh 47895       
      47895 is a valid floating-point value.
      /> ./test29.sh 47895.33
      47895.33 is a valid floating-point value.
      /> ./test29.sh 47895.3e
      fractionalpart is not valid integer.
      /> ./test29.sh 4789t.34
      decimalpart is not valid integer.   


三十、统计英文文章中每个单词出现的频率：
    
      这个技巧的主要目的是显示如何更好的使用awk命令的脚本。
      /> cat > test30.sh
      #!/bin/sh
      #1. 通过当前脚本的pid，生成awk脚本的临时文件名。
      #2. 捕捉信号，在脚本退出时删除该临时文件，以免造成大量的垃圾临时文件。
      awk_script_file="/tmp/scf_tmp.$$"
      trap "rm -f $awk_script_file" EXIT
      #3. while循环将以当前目录下的testfile作为输入并逐行读取，在读取到末尾时退出循环。
      #4. getline读取到每一行将作为awk的正常输入。在内层的for循环中，i要从1开始，因为$0表示整行。NF表示域字段的数量。
      #5. 使$i作为数组的键，如果$i的值匹配正则表达式"^[a-zA-Z]+$"，我们将其视为单词而不是标点。每次遇到单词时，其键值都会递增。
      #6. 最后通过awk脚本提供的特殊for循环，遍历数组的键值数据。
      cat << 'EOF' > $awk_script_file
      BEGIN {
          while (getline < "./testfile" > 0) {
              for (i = 1; i <= NF; ++i) {
                  if (match($i,"^[a-zA-Z]+$") != 0)
                      arr[$i]++
              }
          }
          for (word in arr) {
              printf "word = %s\t count = %s\n",word,arr[word]
          }
      }
      EOF
      awk -f $awk_script_file
      CTRL+D
      /> cat testfile
      hello world liu liu , ,
      stephen liu , ?
      /> ./test30.sh
      word = hello      count = 1
      word = world     count = 1
      word = stephen count = 1
      word = liu         count = 3
































