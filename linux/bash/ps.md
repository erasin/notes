# ps 进程管理


## 关闭进程

关闭进程使用 `kill` 或者 `killall`。

kill 用于关闭进程号 ， killall 用于关闭进程名称

	$ ps aux|grep sleep # 找到进程号 pid

> 

	boc        969  2.9  5.9 3388808 237244 ?      Sl   08:41   7:10 /opt/google/chrome/chrome       
	boc        974  0.2  0.2 316320  9940 ?        S    08:41   0:43 /opt/google/chrome/chrome

> kill 杀死

	$ kill -9 969       # 关闭即可
	$ kill -9 974       # 关闭即可

关闭含有 "chrome" 命令的进程

	$ ps aux|grep chrom[e]|awk '{print $2}'|xargs kill

或者

	$ pgrep chrome|xargs kill


