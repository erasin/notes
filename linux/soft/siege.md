#siege 网站压力测试
**title:** linux下siege 网站压力测试机
**tags:** siege,apache,nginx

##安装Siege

    wget ftp://ftp.joedog.org/pub/siege/siege-latest.tar.gz
    tar zxvf siege-latest.tar.gz
    cd siege-2.67
    ./configure
    make && make install

安装完成后，就可以开始进行压力测试了

修改siege配置 `~/.siegerc`

调整：
    verbose = false
    concurrent = 50
    delay = 1
    internet = true
    benchmark = true

参数介绍：

* **-c NUM** 设置并发的用户（连接）数量，比如-c10,设置并发10个连接。默认的连接数量可以到`~/.siegerc`中查看，指令为concurrent = x，前面咱们已经调整了默认并发连接为50。
* **-r NUM(repetitions)**，重复数量，即每个连接发出的请求数量，设置这个的话，就不需要设置-t了。对应.siegerc配置文件中的reps = x指令
* **-t NUM(time)**，持续时间，即测试持续时间，在NUM时间后结束，单位默认为分，比如`-t10`，那么测试时间为10分钟，`-t10s`，则测试时间为10秒钟。对应.siegerc中的指令为time = x指令
* **-b (benchmark)**,基准测试，如果设置这个参数的话，那么delay时间为0。在.siegerc中咱们修改为默认开启。
* **-f url.txt (file)**,这是url列表文件。对应.siegerc配置文件中的file = x指令

测试结果分析,用500并发重复测试50次bbs.url里的url列表的结果：

    $ siege -c 500 -r 50 -f url
    ** SIEGE 2.67
    ** Preparing 500 concurrent users for battle.
    The server is now under siege..      done.
    Transactions:                  25000 hits           #意思是总共完成了25000次测试
    Availability:                 100.00 %              #测试的有效性100%
    Elapsed time:                  65.52 secs           #用时65.52秒
    Data transferred:              83.65 MB             #传输了83.65MB数据
    Response time:                  0.57 secs           #响应时间
    Transaction rate:             381.56 trans/sec      #每秒传输381.56次
    Throughput:                     1.28 MB/sec         #数据吞吐量每秒1.28MB
    Concurrency:                  216.02                #实际并发访问
    Successful transactions:       21707                #成功的传输
    Failed transactions:               0                #失败的传输
    Longest transaction:            5.83                #每次传输所花最长时间
    Shortest transaction:           0.00                #每次传输所花最短时间

**SEE** 如果你的WEB服务器用的是Apache，请不要将并发数设为大于200。

## apache ab for windows
 
在Apache的bin中文件夹中有 `ab`/`ab.exe`  命令，也可以用于测试

命令是 ab -c 请求数 -n 线程数 网址（建议输入目标网页）。。

例：ab -n 10 -c 10 http://192.168.1.199/index.html

