# ps 进程管理


## 关闭进程

关闭进程使用 `kill` 或者 `killall`。

kill 用于关闭进程号 ， killall 用于关闭进程名称

	$ ps aux|grep sleep # 找到进程号 pid
	$ kill -9 pid       # 关闭即可

关闭含有 "sleep" 命令的进程

	$ ps -o pid,cmd|grep [s]leep|awk '{print $2}'|xargs kill

或者

	$ pgrep sleep|xargs kill


