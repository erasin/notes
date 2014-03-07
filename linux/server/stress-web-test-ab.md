# 压力测试 和 网速测试


## 压力测试

### 工具 

ab[^1] , gunplot[^2]

[^1]: apache 下压力测试工具，功能强大
[^2]: 函数画图工具


### ab 测试

使用脚本循环测试并打印出测试 响应，等待时间的曲线图。

生成文件:   
`boc-ab-*-list.txt` 访问日志  
`boc-ab-*-over.txt` 打印日志  
`boc-ab-*-list.png` 访问响应时间曲线图

测试脚本：

	#!/bin/bash
	# filename abgo.sh

	# 地址列表
	url1="url1"
	url1="url2"
	url1="url3"
	url1="url4"

	 # run 函数 执行并发测验并记录
	 # 参数1 执行序号，用于生成文件序号
	 # 参数2 访问链接数 
	 # 参数3 并发数
	 # 参数4 时间限制 [可选]
	function run {

		i=$1
		n=$2
		c=$3
		t=$4

		#the_day=` date +%Y-%m-%d ` 
		the_day=`date` 
		declare -i j=0

		for k in $url1 $url2 $url3 $url4; do
			j=$j+1

			echo $j
			echo $k

			touch boc-ab-${i}-url${j}-list.txt
			touch boc-ab-${i}-url${j}-over.txt

			file_list=boc-ab-${i}-url${j}-list.txt 
			file_over=boc-ab-${i}-url${j}-over.txt
			file_png=boc-ab-${i}-url${j}-list.png
			#file_png_tile="${the_day} connectioner:${c} Request:${2}"
			file_png_tile="${the_day} 并发数:${c} 发送请求数:${2}"

			date >> $file_over 
			echo "==============================================================" >> $file_over 
			echo "" >> $file_over 


			if [[ $t > 0 ]]; then
				echo "限定执行时间：$t > ab -n $n -c $c -t $t $k"
				echo "限定执行时间：$t 
	ab -n $n -c $c -t $t $k
	---------------------------------------" >> $file_over
				ab -n $n -c $c -t $t -g $file_list $k | tee >> $file_over 
			else
				echo "执行： > ab -n $n -c $c $k"
				echo "执行：
	ab -n $n -c $c $k
	---------------------------------------" >> $file_over
				ab -n $n -c $c -g $file_list $k | tee >> $file_over 

			fi

			echo "" >> $file_over 
			echo "==============================================================" >> $file_over
			date >> $file_over

			# 打印图表
			gnuplot <<- EOF
				datafile="$file_list"
				set title  "$file_png_tile \n $k"
				#set xlabel "Request"
				#set ylabel "ms"
				set xlabel "处理请求数"
				set ylabel "时耗(毫秒)"
				set terminal png
				set terminal pngcairo size 960,600 enhanced
				set output "${file_png}"
				plot datafile using 7 with lines title "ctime", datafile using 8 with lines title "dtime", datafile using 9 with lines title "ttime", datafile using 10 with lines title "wait" 
			EOF
			
			#此处修改eog为自己图片查看器
			eog $file_png & 
		done
	}

	echo "第1次; 100并发 100访问"
	run 1 100 100

	echo "第2次; 200并发 200访问"
	run 2 200 200 

	echo "第3次; 100并发 1000访问"
	run 3 1000 100

	echo "第4次; 200并发 1000访问"
	run 4 1000 200 

将获得的 `*-1-*-over.txt` 中的 进度的响应请求时间 即 `Percentage of the requests served within a certain time (ms)` 下的数据进行统计对比。
> 有兴趣的同学用 sed 来调出下试试。

进度 | 1-u1 | 1-u2 | 1-u3  | 1-u4
-----|------|------|-------|-------
50%  | 310  | 310  | 2780  | 1214
66%  | 453  | 453  | 5032  | 1498
75%  | 1030 | 1030 | 11715 | 1642
80%  | 1154 | 1154 | 12206 | 1722
90%  | 5643 | 5643 | 12559 | 1898
95%  | 5880 | 5880 | 12597 | 5219
98%  | 5897 | 5897 | 13266 | 5267
99%  | 5900 | 5900 | 13282 | 5267
100% | 5900 | 5900 | 13282 | 5267

保存文件命名为 `boc-ab-1.txt`,对应第一次测试。其他的测试相同。

gunplot 脚本, 生成对比图片

	#!/bin/bash 
	#filename contrast2png.sh
	i=$1
	file_title="对比： $i - ${2}"
	file_png="boc-ab-${i}.png"
	file_data="boc-ab-${i}.txt"

	gnuplot <<- EOF
		set terminal png; 
		set terminal pngcairo size 960,600 enhanced
		set output "${file_png}"
		datafile="$file_data"
		set title  "$file_title"
		set xlabel "运行进度（%）"; 
		set ylabel "响应(毫秒)"; 
		set grid
		set xtics ("" 0, "50" 1, "66" 2, "75" 3, "80" 4, "90" 5, "95" 6, "98" 7, "99" 8, "100" 9)
		plot datafile using 2 with lines title "url1", datafile using 3 with lines title "url2", datafile using 4 with lines title "url3", datafile using 5 with lines title "url4" 
	EOF

执行 `./contrast2png.sh 1 ` 来生成 `boc-ab-1.png` 图片,用来直观对比四个地址的访问状态。


## 网速测试

工具： crond, wget , sed, paste , gnuplot[^2] ,vim   
场景： 两个站点，同样大小的文件(用于下载)。

先使用 wget下载设定好的文件，查看要对比的两个站点的所使用的大概时间，比如 2M大小文件使用时间约1分钟。

设定 `crontab -e ` 采集数据。

	*/5 * * * * /home/boc/wgetit.sh

脚本：

	#!/bin/bash
	# filename wgetit.sh
	n=`date +%Y-%m-%d-%H-%M-%S`

	wget --limit-rate=1000 site1/file1.flv -d -a /home/boc/${n}-t1
	wget --limit-rate=1000 site1/file2.flv -d -a /home/boc/${n}-t2

	rm -rvf *.flv

将wget导出的数据整合，生产出对比图。

	#!/bin/bash
	# filename speed2png.sh
	# wget url -d -a `date +%Y-%m-%D-%H-%M-%s-t1`

	function readend() {

		for i in `find . -iname "$2"`; do
			echo $i
			n=`sed -n "$=" $i`
			((n--))
			sed -n "${n} p" $i >> time${1}.txt
		done

	}

    # 生成 对应的最终时间表
	readend 1 "2013*t1"
	readend 2 "2013*t2"

使用vim 处理 time1.txt 和 time2.txt,去除多余部分
> 可以在上面脚本中使用 sed 直接处理，正则无敌

	:%s/^.*(//g
	:%s/\ KB.*$//g

合并结果并生成图

    # 合并两列
	paste -d "\t" time1.txt time2.txt >> time3.txt 

    # 生成图标
	gnuplot <<- EOF
		set terminal png; 
		set terminal pngcairo size 960,600 enhanced
		set output "speed.png"
		datafile="time3.txt"
		set title  "网速状态"
		set xlabel "时间/5分"; 
		set ylabel "网速（Kb/s)"; 
		set grid
		plot datafile using 1 with lines title "site1", datafile using 2 with lines title "site2"
	EOF


*[Kb/s]: 千字节每秒
*[date]: 时间函数
