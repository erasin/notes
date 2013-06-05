# 同步方案

通道： 使用ssh协议通信,利用ssh的 `authorized_keys` 特性，使用ssh-keygen生成不带有密码的密钥。
执行： crond来设定执行周期

> 注： 下面  local 指本地，server指同步的服务器，pwd 指当前目录路径

## 静态文件同步

视频文件和图片文件比较大，不适合监控具体内容，利用 rsync 同步文件夹，可以快速计算文件夹对比索引同步。

	rsync --log-file="/local/rsync.log" -az --delete -e ssh user@ssh-server-ip:/server/pwd /local/pwd

> 不使用同步删除时，去除  `--delete` 即可。

## 站点中心同步

使用git自动管理中心站点代码，检测站点文件变动，防止客户端的文件篡改以及添加删除等。

### A-server:中心端

在站点根目录中添加过滤列表文件 `.gitignore` 

    *.log
    *.html
    upload/
    log/

则在检测变动的时候，可以过滤 所有 log，html后缀和 upload,log文件夹。

如果没有库索引则建立git库

	$ git init 
	$ git add .
	$ git commit -a -m "site source init"

> **建议** 建立分支 `develop` 开发分支，本地运行成功后推送到服务器，然后合并到`master`供同步端拉取

### B-client: 同步端 

拷贝站点

	$ git clone user@ip:/pwd/.git

每5m 更新下拉 master 分支 文件数据。

流程：

![tamper](/upload/server/git-tamper.png)


脚本功能实现

1. 检测文件变动
2. 无有变动则拉取服务器A的变动内容
3. 有变动的时候则将变动文件提交到临时git分支中，并删除。 之后拉取服务器A的变动内容
4. 同时将被篡改的内容部分写入到日志文件中去。

- - - - - - - 

    #!/bin/bash
    # file : /pwd/shell/gitpull.sh

    # 修改地址
    cd `pwd`
    # 修改日志地址
    gitlog='/pwd/git.log'

    today=`date +%Y%m%d`
    gitstatus=`git status`
    echo '检测库状态'
    if [[ $gitstatus == *"working directory clean"* ]];then 
        echo '无任何改变'
        echo '拉取'
        git pull >> $gitlog
    else 
        echo '有文件变动'
        echo '将改变提交到分支 '$today
        git checkout -b $today 

        echo '写入日志'
        touch $gitlog
        echo " " >> $gitlog 
        date '+%Y-%m-%d %H:%M' >> $gitlog
        echo "==========================================" >> $gitlog
        git diff >> $gitlog
        echo "==========================================" >> $gitlog
        echo " " >> $gitlog 

        git add .
        git commit -a -m 'its change'
        echo '回到 master'
        git checkout master
        echo '删除分支'
        git branch -D $today
        echo '拉取'
        git pull >> $gitlog
    fi

添加执行

	$ crontab -e
	*/5 * * * * /pwd/shell/gitpull.sh


> 会产生5m的时间差。

## mysql 文件备份

> 非均衡负载，可查阅[mysql 主从同步](mysql-master-slave.md)

	#!/bin/bash
	# filename: bakmysql.sh

	# set path 文件存储位置

	bakpath='/pwd/backup/'
	logfile='/pwd/backup/bak.log'

	themonth=${bakpath}`date +%y%m`'/'
	theday=${themonth}`date +%d`'/' 
	thetime=${theday}`date +%Y%m%d%H%M`'.dbname.sql.bak' 

	declare -i month=`date +%y%m`
	month=$month-2

	echo '' >> $logfile
	echo `date +%D\ %T` >> $logfile

	for i in `ls -F $bakpath`; do
		if [[ -d $bakpath$i ]]; then
			if [[ $i < $month ]]; then
				echo 'delete out 2 month '$i >> $logfile
				rm -rf ${bakpath}$i
			fi
		fi
	done

	if [[ -d $themonth ]]; then
		echo $themonth 'is exist;' >> $logfile
	else
		echo 'makedir ' $theday >> $logfile
		mkdir -p $theday
	fi

	if [[ -d $theday ]]; then
		echo $theday 'is exist;' >> $logfile
	else
		echo 'makedir ' $theday >> $logfile
		mkdir $theday
	fi

	echo 'bakup mysqldatabase dbname - '${thetime} >> $logfile
	mysqldump -uroot -p123456 dbname > $thtime 

	# -h host


添加每 30m 执行 

    $ crontab -e
    */30 * * * * /pwd/shell/mysqlbak.sh


