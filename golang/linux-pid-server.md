# go 服务器常用


1. 使用 nohup 守护进程
    语法：nohup Command [ Arg ... ] [　& ]
2. 使用 grace 热气动服务器支持
    github.com/facebookgo/grace/gracehttp

nohup command > myout.file 2>&1 &

使用 jobs 查看任务。

　　使用 fg %n　关闭。