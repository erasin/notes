
## mysql_affected_rows()

取得前一次 MySQL 操作所影响的记录行数。	3

## mysql_change_user()

不赞成。改变活动连接中登录的用户	3

## mysql_client_encoding()

返回当前连接的字符集的名称	4

## mysql_close()

关闭非持久的 MySQL 连接。	3

## mysql_connect()

打开非持久的 MySQL 连接。	3

## mysql_create_db()

不赞成。新建 MySQL 数据库。使用 mysql_query() 代替。	3

## mysql_data_seek()

移动记录指针。	3

## mysql_db_name()

从对 mysql_list_dbs() 的调用返回数据库名称。	3

## mysql_db_query()



## 不赞成。发送一条 MySQL 查询。

## 使用 mysql_select_db() 和 mysql_query() 代替。

## 3

## mysql_drop_db()



## 不赞成。丢弃（删除）一个 MySQL 数据库。

## 使用 mysql_query() 代替。

## 3

## mysql_errno()

返回上一个 MySQL 操作中的错误信息的数字编码。	3

## mysql_error()

返回上一个 MySQL 操作产生的文本错误信息。	3

## mysql_escape_string()



## 不赞成。转义一个字符串用于 mysql_query。

## 使用 mysql_real_escape_string() 代替。

## 4

## mysql_fetch_array()

从结果集中取得一行作为关联数组，或数字数组，或二者兼有。	3

## mysql_fetch_assoc()

从结果集中取得一行作为关联数组。	4

## mysql_fetch_field()

从结果集中取得列信息并作为对象返回。	3

## mysql_fetch_lengths()

取得结果集中每个字段的内容的长度。	3

## mysql_fetch_object()

从结果集中取得一行作为对象。	3

## mysql_fetch_row()

从结果集中取得一行作为数字数组。	3

## mysql_field_flags()

从结果中取得和指定字段关联的标志。	3

## mysql_field_len()

返回指定字段的长度。	3

## mysql_field_name()

取得结果中指定字段的字段名。	3

## mysql_field_seek()

将结果集中的指针设定为指定的字段偏移量。	3

## mysql_field_table()

取得指定字段所在的表名。	3

## mysql_field_type()

取得结果集中指定字段的类型。	3

## mysql_free_result()

释放结果内存。	3

## mysql_get_client_info()

取得 MySQL 客户端信息。	4

## mysql_get_host_info()

取得 MySQL 主机信息。	4

## mysql_get_proto_info()

取得 MySQL 协议信息。	4

## mysql_get_server_info()

取得 MySQL 服务器信息。	4

## mysql_info()

取得最近一条查询的信息。	4

## mysql_insert_id()

取得上一步 INSERT 操作产生的 ID。	3

## mysql_list_dbs()

列出 MySQL 服务器中所有的数据库。	3

## mysql_list_fields()



## 不赞成。列出 MySQL 结果中的字段。

## 使用 mysql_query() 代替。

## 3

## mysql_list_processes()

列出 MySQL 进程。	4

## mysql_list_tables()



## 不赞成。列出 MySQL 数据库中的表。

## 使用Use mysql_query() 代替。

## 3

## mysql_num_fields()

取得结果集中字段的数目。	3

## mysql_num_rows()

取得结果集中行的数目。	3

## mysql_pconnect()

打开一个到 MySQL 服务器的持久连接。	3

## mysql_ping()

Ping 一个服务器连接，如果没有连接则重新连接。	4

## mysql_query()

发送一条 MySQL 查询。	3

## mysql_real_escape_string()

转义 SQL 语句中使用的字符串中的特殊字符。	4

## mysql_result()

取得结果数据。	3

## mysql_select_db()

选择 MySQL 数据库。	3

## mysql_stat()

取得当前系统状态。	4

## mysql_tablename()

不赞成。取得表名。使用 mysql_query() 代替。	3

## mysql_thread_id()

返回当前线程的 ID。	4

## mysql_unbuffered_query()

向 MySQL 发送一条 SQL 查询（不获取 / 缓存结果）。	4
