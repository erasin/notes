# mysql master slave 主从同步

>下面cli中 `$` 为普通权限，`#` 为root权限。

## 模式1 使用mysql命令来设定slave

__Master__：SerA IP:10.0.2.81     
__Slave__：SerB IP:10.0.2.82

A: my.cnf

	server-id=1
	#master
	log-bin=mysql-bin
	binlog-do-db=dbname       # 需要备份的数据库名, 多个则重复设置 

A: mysql cmd

	mysql> GRANT all privileges ON *.* TO 'backup'@'10.0.2.82' IDENTIFIED BY '123456'
	mysql> flush privileges; 

导出 dbname 数据库

	$ mysqldump -uroot -p dbname > dbname.sql

重启mysql服务器

	# service mysql restart 

A: mysql cmd

	mysql> show grants;          -- 查看权限生效 -- 
	mysql> show master status\G; -- 录下 FILE & Position 的值 -- 

记下 file & Position 的值，在配置 slave时会用到

权限(需要时执行)

	mysql> show grants;  -- 查看权限 --
	mysql> GRANT FILE,SELECT,REPLICATION SLAVE ON *.* TO 'backup'@'10.0.2.82' IDENTIFIED BY '123456';
	mysql> flush privileges; 

B: my.cnf

	server-id=2
	#master 
	log-bin=mysql-bin
	#slave
	replicate-do-db=dbname  # 需要备份的数据库名, 多个则重复设置
	master-connect-retry=60 # 服务器拓机重连 默认60s
	slave-net-timeout=60    # 60s同步

B: mysql cmd

	-- 添加 master 服务器 --
	mysql> stop slave; 
	mysql> change master to 
		master_host='10.0.2.81', 
		master_port=3306, 
		master_user='backup', 
		master_password='123456'
		master_log_file='mysql-bin.000002',
		master_log_pos=98;
	-- 其中 file & pos 到 A 的 master中找 --
	mysql> start slave;
	mysql> show slave status\G;
		-- Slave_IO_Running和Slave_SQL_Running的状态都必须是YES --
	mysql> show master stauts\G;

	-- 导入数据 --
	mysql> stop slave;
	mysql> create database `dbname`;
	mysql> use dbname;
	mysql> source /pwd/dbname.sql;
	mysql> start slave;


如果需要则重启mysql服务器

	# service mysql restart 

## 模式2 使用配置文件来设定slave

A: my.cnf

	default-character-set=utf8
	server-id = 1
	# master
	log-bin=mysql-bin 
	binlog-do-db=dbname           # 需要备份的数据库名, 多个则重复设置
	binlog-ignore-db=dbname       # 不需要备份的数据库

B: my.cnf

	server-id =2                    # 不可重复
	# master
	log-bin=mysql-bin
	relay-log=relay-bin
	relay-log-index=relay-bin-index
	# slave
	master-host ='ip'
	master-user =root
	master-password =passwd
	master-port =3306
	master-connect-retry=30         # s
	slave-net-timeout=60            # 60s同步
	replicate-do-db=dbname          # 需要同步的数据库名, 多个则重复设置


索引偏移：

auto_increment_increment：自增值的自增量  
auto_increment_offset： 自增值的偏移量

设置了两个值之后，改服务器的自增字段值限定为：
`auto_increment_offset + auto_increment_increment*N`  的值，其中N>=0，但是上限还是要受定义字段的类型限制。


比如：

auto_increment_offset=1  
auto_increment_increment=2  
那么ID则是所有的奇数[1,3,5,7，.....]

如果：

auto_increment_offset=5    
auto_increment_increment=10    
那么ID则是所有的奇数[5,15,25,35，.....]




> 其余参考模式1


## password 重置

	$ mysqladmin -uroot -p oldpassword newpasswd
	$ mysql -uroot -p user 
	mysql> UPDATE user SET password=PASSWORD(”new password”) WHERE user=’root’; 

## 局域网访问 权限

	# 给所有的权限
	mysql> grant all privileges on *.* to root@"ip" identified by 'password' with grant option;
	mysql> flush privileges; 

	# 仅仅给同步权限
	mysql> grant replication slave,file on *.* to 'username'@'ip' identified by 'password';
	mysql> flush privileges; 

## 实例 多台循环蛇型同步

环境　linux 3台

IP：

* A: 10.86.1.83
* B: 10.86.1.82
* C: 10.86.1.81

同步数据库： demodb


使用同步方向　　A 83 -> B 82 -> C 81 -> A83

**1. 授权用户** ,为了方便使用统一帐号和密码: demouser 123456

分别开始在ＡＢＣ三台服务器上授权用户登录访问　Ａ要给Ｂ，Ｂ给Ｃ依次设定

	A->B:
	mysql>　grant replication slave,file on *.* to 'demouser'@'10.86.1.82' identified by '123456';
	mysql> flush privileges; 

	B->C:	
	mysql>　grant replication slave,file on *.* to 'demouser'@'10.86.1.81' identified by '123456';
	mysql> flush privileges; 

	C->A:	
	mysql>　grant replication slave,file on *.* to 'demouser'@'10.86.1.83' identified by '123456';
	mysql> flush privileges; 

**2. 修改配置** , 修改`my.cnf`文件添加


A: 83 服务器

	# master
	log-bin=mysql-bin
	server-id = 1
	binlog-do-db=demodb
	binlog-ignore-db=mysql

	log-slave-updates

	# slave
	replicate-do-db=demodb
	replicate-ignore-db=mysql
	slave-skip-errors=all
	sync_binlog=1

	# index
	auto_increment_increment=1
	auto_increment_offset=1

	master-host = 10.86.1.81
	master-user = demouser
	master-password = '123456' 
	master-port = 3306

B: 82 服务器

	# master
	log-bin=mysql-bin
	server-id = 2
	binlog-do-db=demodb
	binlog-ignore-db=mysql

	log-slave-updates

	# slave
	replicate-do-db=demodb
	replicate-ignore-db=mysql
	slave-skip-errors=all
	sync_binlog=1

	# index
	auto_increment_increment=2
	auto_increment_offset=2

	master-host = 10.86.1.83
	master-user = demouser
	master-password = '123456' 
	master-port = 3306

C: 81 服务器

	# master
	log-bin=mysql-bin
	server-id = 3
	binlog-do-db=demodb
	binlog-ignore-db=mysql

	log-slave-updates

	# slave
	replicate-do-db=demodb
	replicate-ignore-db=mysql
	slave-skip-errors=all
	sync_binlog=1

	# index
	auto_increment_increment=2
	auto_increment_offset=3

	master-host = 10.86.1.82
	master-user = demouser
	master-password = '123456' 
	master-port = 3306

> 注意参考 `auto_increment_increment` 和 `auto_increment_offset` 

**3. 重启服务 检查错误**

	service mysql restart

登录msyql查看 slave 状态 ,执行  `show slave status\G;`

	Slave_IO_Running: Yes
    Slave_SQL_Running: Yes

以上两个参数。

极有可能出现的问题：

**Slave_IO_Running: NO**  同步日志和索引问题

比如 B 出现错误 则查看其master A

	A服务器：
	mysql> show master status\G;
	*************************** 1. row ***************************
	            File: mysql-bin.000016
	        Position: 41489841
	    Binlog_Do_DB: demodb
	Binlog_Ignore_DB: mysql
	1 row in set (0.00 sec)

	ERROR:
	No query specified

    B:服务器
    mysql> show slave status\G;
	*************************** 1. row ***************************
               Slave_IO_State: Waiting for master to send event
                  Master_Host: 10.86.1.83
                  Master_User: root
                  Master_Port: 3306
                Connect_Retry: 60
              Master_Log_File: mysql-bin.000024             # 该位置出现错误           
          Read_Master_Log_Pos: 41523164
               Relay_Log_File: web82-relay-bin.000004
                Relay_Log_Pos: 41071442
        Relay_Master_Log_File: mysql-bin.000024
             Slave_IO_Running: No
            Slave_SQL_Running: Yes
              Replicate_Do_DB: demodb
          Replicate_Ignore_DB: mysql

在A上获取日志(File)和pos(41489841)后B上执行修改

	mysql> STOP SLAVE;CHANGE MASTER TO MASTER_LOG_FILE='mysql-bin.000016', MASTER_LOG_POS=41489841;START SLAVE ;

不出问题的话查看 `show slave status\G;` 的 `Slave_IO_Running` 应该ok

**Slave_SQL_Running: NO** 错误

	mysql> SLAVE STOP; SET GLOBAL SQL_SLAVE_SKIP_COUNTER=1; START SLAVE; 

**4. 测试**
插入表到其他服务器查看

	mysql> CREATE TABLE `testc` ( `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY, `title` int NOT NULL);

