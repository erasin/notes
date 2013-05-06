# mysql master slave 主从同步

>下面cli中 `$` 为普通权限，`#` 为root权限。

## 模式1

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

## 模式2 

A: my.cnf

	default-character-set=utf8
	server-id = 1
	# master
	log-bin=mysql-bin 
	binlog-do-db=dbname           # 需要备份的数据库名, 多个则重复设置
	binlog-ignore-db=dbname               # 不需要备份的数据库

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
	replicate-do-db=dbname          # 需要备份的数据库名, 多个则重复设置


> 其余参考模式1


## password 重置

	$ mysqladmin -uroot -p oldpassword newpasswd
	$ mysql -uroot -p user 
	mysql> UPDATE user SET password=PASSWORD(”new password”) WHERE user=’root’; 


## 局域网访问 权限

	mysql> grant all privileges on *.* to root@"ip" identified by 'password' with grant option;
	sql> flush privileges; 


