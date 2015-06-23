# mysql 监控


```sh
#/bin/sh

#检测mysql server是否正常提供服务
mysqladmin -u sky -ppwd -h localhost ping

#获取mysql当前的几个状态值
mysqladmin -u sky -ppwd -h localhost status

#获取数据库当前的连接信息
mysqladmin -u sky -ppwd -h localhost processlist

#获取当前数据库的连接数
mysql -u root -p123456 -BNe "select host,count(host) from processlist group by host;" information_schema

#显示mysql的uptime
mysql -e"SHOW STATUS LIKE '%uptime%'"|awk '/ptime/{ calc = $NF / 3600;print $(NF-1), calc"Hour" }'

#查看数据库的大小
mysql -u root -p123456-e 'select table_schema,round(sum(data_length+index_length)/1024/1024,4) from information_schema.tables group by table_schema;'

#查看某个表的列信息
mysql -u <user> --password=<password> -e "SHOW COLUMNS FROM <table>" <database> | awk '{print $1}' | tr "\n" "," | sed 's/,$//g'

#执行mysql脚本
mysql -u user-name -p password < script.sql

#mysql dump数据导出
mysqldump -uroot -T/tmp/mysqldump test test_outfile --fields-enclosed-by=\" --fields-terminated-by=,

#mysql数据导入
mysqlimport --user=name --password=pwd test --fields-enclosed-by=\" --fields-terminated-by=, /tmp/test_outfile.txt
LOAD DATA INFILE '/tmp/test_outfile.txt' INTO TABLE test_outfile FIELDS TERMINATED BY '"' ENCLOSED BY ',';

#mysql进程监控
ps -ef | grep "mysqld_safe" | grep -v "grep"
ps -ef | grep "mysqld" | grep -v "mysqld_safe"| grep -v "grep"

#查看当前数据库的状态
mysql -u root -p123456 -e 'show status'


#mysqlcheck 工具程序可以检查(check),修 复( repair),分 析( analyze)和优化(optimize)MySQL Server 中的表
mysqlcheck -u root -p123456 --all-databases

#mysql qps查询  QPS = Questions(or Queries) / Seconds
mysql -u root -p123456 -e 'SHOW /*!50000 GLOBAL */ STATUS LIKE "Questions"'
mysql -u root -p123456 -e 'SHOW /*!50000 GLOBAL */ STATUS LIKE "Queries"'

#mysql Key Buffer 命中率  key_buffer_read_hits = (1 - Key_reads / Key_read_requests) * 100%  key_buffer_write_hits= (1 - Key_writes / Key_write_requests) * 100%
mysql -u root -p123456 -e 'SHOW /*!50000 GLOBAL */ STATUS LIKE "Key%"'

#mysql Innodb Buffer 命中率  innodb_buffer_read_hits=(1-Innodb_buffer_pool_reads/Innodb_buffer_pool_read_requests) * 100%
mysql -u root -p123456 -e 'SHOW /*!50000 GLOBAL */ STATUS LIKE "Innodb_buffer_pool_read%"'

#mysql Query Cache 命中率 Query_cache_hits= (Qcache_hits / (Qcache_hits + Qcache_inserts)) * 100%
mysql -u root -p123456 -e 'SHOW /*!50000 GLOBAL */ STATUS LIKE "Qcache%"'

#mysql Table Cache 状态量
mysql -u root -p123456 -e 'SHOW /*!50000 GLOBAL */ STATUS LIKE "Open%"'

#mysql Thread Cache 命中率  Thread_cache_hits = (1 - Threads_created / Connections) * 100%  正常来说,Thread Cache 命中率要在 90% 以上才算比较合理。
mysql -u root -p123456 -e 'SHOW /*!50000 GLOBAL */ STATUS LIKE "Thread%"'

#mysql 锁定状态:锁定状态包括表锁和行锁两种,我们可以通过系统状态变量获得锁定总次数,锁定造成其他线程等待的次数,以及锁定等待时间信息
mysql -u root -p123456 -e 'SHOW /*!50000 GLOBAL */ STATUS LIKE "%lock%"'

#mysql 复制延时量 在slave节点执行
mysql -u root -p123456 -e 'SHOW SLAVE STATUS'

#mysql Tmp table 状况 Tmp Table 的状况主要是用于监控 MySQL 使用临时表的量是否过多,是否有临时表过大而不得不从内存中换出到磁盘文件上
mysql -u root -p123456 -e 'SHOW /*!50000 GLOBAL */ STATUS LIKE "Created_tmp%"'

#mysql Binlog Cache 使用状况:Binlog Cache 用于存放还未写入磁盘的 Binlog 信 息 。
mysql -u root -p123456 -e 'SHOW /*!50000 GLOBAL */ STATUS LIKE "Binlog_cache%"'

#mysql nnodb_log_waits 量:Innodb_log_waits 状态变量直接反应出 Innodb Log Buffer 空间不足造成等待的次数
mysql -u root -p123456 -e 'SHOW /*!50000 GLOBAL */ STATUS LIKE "Innodb_log_waits'
```

行中的mysql状态查看

对正在运行的mysql进行监控，其中一个方式就是查看mysql运行状态。



show global status where variable_name in('com_select','com_insert','com_delete','com_update');

show global status where Variable_name in ('Com_commit','Com_delete','Com_insert','Com_rollback','Com_select','Com_update','Questions','uptime',);

(1)QPS(每秒Query量)
QPS = Questions(or Queries) / seconds
mysql > show  global  status like 'Question%';

questions = show global status like 'questions';
uptime = show global status like 'uptime';
qps=questions/uptime



(2)TPS(每秒事务量)
TPS = (Com_commit + Com_rollback) / seconds
mysql > show global status like 'Com_commit';
mysql > show global status like 'Com_rollback';

com_commit = show global status like 'com_commit';
com_rollback = show global status like 'com_rollback';
uptime = show global status like 'uptime';
tps=(com_commit + com_rollback)/uptime


(3)key Buffer 命中率
mysql>show  global   status  like   'key%';
key_buffer_read_hits = (1-key_reads / key_read_requests) * 100%
key_buffer_write_hits = (1-key_writes / key_write_requests) * 100%

(4)InnoDB Buffer命中率
mysql> show status like 'innodb_buffer_pool_read%';
innodb_buffer_read_hits = (1 - innodb_buffer_pool_reads / innodb_buffer_pool_read_requests) * 100%

(5)Query Cache命中率
mysql> show status like 'Qcache%';
Query_cache_hits = (Qcahce_hits / (Qcache_hits + Qcache_inserts )) * 100%;

(6)Table Cache状态量
mysql> show global  status like 'open%';
比较 open_tables  与 opend_tables 值

(7)Thread Cache 命中率
mysql> show global status like 'Thread%';
mysql> show global status like 'Connections';
Thread_cache_hits = (1 - Threads_created / connections ) * 100%

(8)锁定状态
mysql> show global  status like '%lock%';
Table_locks_waited/Table_locks_immediate=0.3%  如果这个比值比较大的话，说明表锁造成的阻塞比较严重
Innodb_row_lock_waits innodb行锁，太大可能是间隙锁造成的

(9)复制延时量
mysql > show slave status
查看延时时间

(10) Tmp Table 状况(临时表状况)
mysql > show status like 'Create_tmp%';
Created_tmp_disk_tables/Created_tmp_tables比值最好不要超过10%，如果Created_tmp_tables值比较大，
可能是排序句子过多或者是连接句子不够优化

(11) Binlog Cache 使用状况
mysql > show status like 'Binlog_cache%';
如果Binlog_cache_disk_use值不为0 ，可能需要调大 binlog_cache_size大小

(12) Innodb_log_waits 量
mysql > show status like 'innodb_log_waits';
Innodb_log_waits值不等于0的话，表明 innodb log  buffer 因为空间不足而等待

比如命令：
>#show global status;
虽然可以使用：
>#show global status like %...%;
来过滤，但是对应长长的list，每一项都代表什么意思，还是有必要弄清楚。

## MySQL监控时常用的的几个MySQL命令
```
status = show status like ‘%%' [例:show status like 'Com_select']
variables = show variables like ‘%%' [例:show variables like 'query_cache_size']
```

1. MySQL查询次数(status)
    * Com_select;Com_update;Com_insert;Com_delete;Com_change_db
2. 查询缓存空间大小:query_cache_size(variables)
    * 查询缓存最大查询数据集大小:query_cache_limit(variables);
    * 缓存中的查询个数:Qcache_inserts(status);
    * 查询缓存命中率：（Qcache_hits/(Qcache_hits+Qcache_inserts)）\*100% (status)
3. 索引缓存命中率
    * 索引缓存空间大小：key_buffer_size (variables)
    * 索引缓存命中率：(Key_reads/Key_read_requests) \*100% (status)
4. 并发连接数
    * 最大充许连接数:max_connections(variables)
    * 实际最大连接数:max_used_connections(status)
    * 当前连接数:Threads_connected(status)
    * 活跃连接数：Threads_running(status)
    * 缓存连接数：Threads_cache(status)
5. 流量统计(status)
    * Bytes_received ,Bytes_sent(status)
6. 连接次数
    * 每秒连接次数:Connections(status)
    * 每秒实际创建连接次数：Threads_created(status)
7. 表锁定统计
    * 立即释放的表锁数：Table_locks_immediate(status)
    * 需要等待的表锁数：Table_locks_waited(status)

    SELECT report FROM sv_report_human_hour ORDER BY id DESC LIMIT 1,1 \G
