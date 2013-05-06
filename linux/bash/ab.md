#ab


格式： 

	./ab [options] [http://]hostname[:port]/path


参数  |              | 描述
----- | ------------ | -----------------------
-A    | attribute    | 对服务器提供BASIC认证信任。 用户名和密码由一个:隔开，并以base64编码形式发送。 无论服务器是否需要(即, 是否发送了401认证需求代码)，此字符串都会被发送。
-n    | requests     | 在测试会话中所执行的请求个数。 默认时，仅执行一个请求，但通常其结果不具有代表意义。
-c    | concurrency  | 一次产生的请求个数。默认是一次一个。
-t    | timelimit    | 测试所进行的最大秒数。其内部隐含值是-n 50000。它可以使对服务器的测试限制在一个固定的总时间以内
-C    | attribute    | Add cookie, eg. 'Apache=1234'. (repeatable)
-H    | attribute    | Add Arbitrary header line, eg. 'Accept-Encoding: gzip'
-i    |              | 执行HEAD请求，而不是GET。
-k    |              | 启用HTTP KeepAlive功能，即, 在一个HTTP会话中执行多个请求。 默认时，不启用KeepAlive功能.
-p    | POST-file    | 包含了需要POST的数据的文件.
-X    | proxy[:port] | 对请求使用代理服务器。
-T    | content-type | POST数据所使用的Content-type头信息
-q    |              | 如果处理的请求数大于150， ab每处理大约10%或者100个请求时，会在stderr输出一个进度计数。 此-q标记可以抑制这些信息。
-w    |              | 以HTML表的格式输出结果。默认时，它是白色背景的两列宽度的一张表。
-e    | csv-file     | 产生一个以逗号分隔的(CSV)文件， 其中包含了处理每个相应百分比的请求所需要(从1%到100%)的相应百分比的(以微妙为单位)时间。
-g    | gnuplot-file | 把所有测试结果写入一个'gnuplot'或者TSV (以Tab分隔的)文件。
-h    |              | 显示使用方法。



一般我们用 -c 和 -n 参数就可以了. 例如:

	./ab -c 1000 -n 1000 http://127.0.0.1/index.php

这个表示同时处理1000个请求并运行1000次index.php文件.

Server Software: Apache/2.0.54          // 平台apache 版本2.0.54  
Server Hostname: 127.0.0.1              // 服务器主机名  
Server Port: 80                         // 服务器端口  
Document Path: /index.html.zh-cn.gb2312 // 测试的页面文档  
Document Length: 1018 bytes             // 文档大小  

Concurrency Level: 1000                // 并发数  
Time taken for tests: 8.188731 seconds // 整个测试持续的时间  
Complete requests: 1000                // 完成的请求数量  
Failed requests: 0                     // 失败的请求数量  
Write errors: 0

Total transferred: 1361581 bytes                                    // 整个场景中的网络传输量  
HTML transferred: 1055666 bytes                                     // 整个场景中的HTML内容传输量  
Requests per second: 122.12 [#/sec] (mean)                          // 大家最关心的指标之一，相当于 LR 中的 每秒事务数 ，后面括号中的 mean 表示这是一个平均值  
Time per request: 8188.731 [ms] (mean)                              // 大家最关心的指标之二，相当于 LR 中的 平均事务响应时间 ，后面括号中的 mean 表示这是一个平均值  
Time per request: 8.189 [ms] (mean, across all concurrent requests) // 每个请求实际运行时间的平均值  
Transfer rate: 162.30 [Kbytes/sec] received                         // 平均每秒网络上的流量，可以帮助排除是否存在网络流量过大导致响应时间延长的问题  

Connection Times (ms)  
min mean[+/-sd] median max  
Connect:|4 646 1078.7 89 3291  
Processing: 165 992 493.1 938 4712  
Waiting: 118 934 480.6 882 4554  
Total: 813 1638 1338.9 1093 7785  

//网络上消耗的时间的分解，各项数据的具体算法还不是很清楚

	Percentage of the requests served within a certain time (ms)
	50% 1093
	66% 1247
	75% 1373
	80% 1493
	90% 4061
	95% 4398
	98% 5608
	99% 7368
	100% 7785 (longest request)

//整个场景中所有请求的响应情况。在场景中每个请求都有一个响应时间，其中50％的用户响应时间小于1093 毫秒，60％ 的用户响应时间小于1247 毫秒，最大的响应时间小于7785 毫秒

由于对于并发请求，cpu实际上并不是同时处理的，而是按照每个请求获得的时间片逐个轮转处理的，所以基本上第一个Time per request时间约等于第二个Time per request时间乘以并发请求数



使用-g参数其结果可以输出到指定文件 

	ctime：connection time 
	dtime: processing time 
	ttime: total time, = connection time + processing time 
	wait：wait time 


	$ gnuplot 
		datafile="http_benchmark.txt"
		set terminal png 
		set output "http_benchmark.png" 
		set xlabel "request" 
		set ylabel "ms" 
		plot datafile using 7 with lines title "ctime", \ 
		datafile using 8 with lines title "dtime", \ 
		datafile using 9 with lines title "ttime", \ 
		datafile using 10 with lines title "wait" 


[gnuplot](http://darksair.org/wiki/Gnuplot.html 'gnuplot 教程')
