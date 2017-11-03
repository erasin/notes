Linux 下 使用 supervisor 构建服务
======================

下载

```
yum install supervisor
# 此处创建自己的服务文件
vim /etc/supervisor/conf.d/{youfile}.conf
```
插入类似的代码

```
[program:shadowsocks]
command=ssserver -c /etc/shadowsocks.json
autorestart=true
user=nobody
```

重启服务

```
service supervisor start
supervisorctl reload
# 在已经启动的情况下
supervisorctl restart {youfile}
```
