# 同步方案

通道： 使用ssh协议通信,利用ssh的 authorized_keys 特性，使用ssh-keygen生成不带有密码的密钥。
执行： crond来设定执行周期

> 注： 下面  local 指本地，server指同步的服务器，pwd之当前目录路径

## 静态文件同步

视频文件和图片文件比较大，不适合监控具体内容，利用 rsync 同步文件夹，可以快速计算文件夹对比索引同步。

	rsync --log-file="/local/rsync.log" -az --delete -e ssh user@ssh-server-ip:/server/pwd /local/pwd

> 不使用同步删除时，去除  `--delete` 即可。

周期性执行，稳定性待定。

## 站点中心同步

使用git自动管理中心站点代码，及时检测站点文件变动。

### A-server:中心端

建立git库

	$ vim .gitignore # 添加过滤列表
	$ git init 
	$ git add .
	$ git commit -a -m "site source init"

每5m更新提交数据。

	#!/bin/bash
	# file : /pwd/shell/gitcommit.sh
	de=`date +%Y%m%d%H%M%S`
	cd /server/pwd
	git add .
	git commit -a -m "$de"

添加执行

	$ crontab -e
	*/5 * * * * /pwd/shell/gitcommit.sh

### B-client: 同步端 

拷贝站点

	$ git clone user@ip:/pwd/.git

每5m 更新下拉 master 分支 文件数据。

	#!/bin/bash
	# file : /pwd/shell/gitpull.sh
	cd /local/pwd
	git pull

添加执行

	$ crontab -e
	*/5 * * * * /pwd/shell/gitpull.sh

以上执行效果，只有文件被变动的时候才会发生。

好处是，可以检测到具体某个文件中的某行某字符的变动。

> 另外，时间虽然都是每5m，但是根据时间同步，会产生5m的时间差。


