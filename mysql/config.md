autocommit={0|1}
设定MySQL事务是否自动提交，1表示立即提交，0表示需要显式提交。作用范围为全局或会话，可用于配置文件中(但在5.5.8之前的版本中不可用于配置文件)，属于动态变量。
automatic_sp_privileges={0|1}
设定MySQL服务器是否为存储例程的创建赋予其创建存储例程上的EXECUTE和ALTER ROUTINE权限，默认为1(赋予此两个权限给其创建者)。作用范围为全局。
back_log=#
当MySQL的主线程在短时间内收到大量连接请求时，其会花些时间检测已经有线程并为新请求启动新线程，back_log参数的值即为短时间内到达的请求中有多少可以被接受并等待主MySQL线程进行后续处理。作用范围为全局，可以用配置文件，非动态变量。
basedir=PATH, -b PATH
用于指定MySQL的安装目录，所有其它的常用相对路径都相对于此处的路径而言。作用范围为全局，可用于配置文件中，但属于非动态变量。
bind-address=ADDR
指定mysqld服务监听的IP地址，默认为0.0.0.0，表示本机已配置的所有IP地址。作用范围为全局，可用于配置文件中，但属于非动态变量。
binlog-format={ROW|STATEMENT|MIXED}
指定二进制日志的类型，默认为STATEMENT。如果设定了二进制日志的格式，却没有启用二进制日志，则MySQL启动时会产生警告日志信息并记录于错误日志中。作用范围为全局或会话，可用于配置文件，且属于动态变量。
buld_insert_buffer_size
MyISAM引擎使用一个特殊的树状结构的缓存来加速批量插入操作，如INSERT…SELECT, INSERT…VALUES(…),(…),…和LOAD DATA INFILE命令完成的插入操作。对于每个线程来说，此buffer的大小是独立的，其配置的数值单位为字节，有效取值范围为0至“2^CPU字长”次方，默认大小为8MB。作用范围为全局或会话，可用于配置文件，属动态变量。
chroot=PATH, -r PATH
设定MySQL基于chroot模式工作时的工作目录，在安全问题尤为重要的环境中，这是推荐使用的机制。但此时，LOAD DATA INFILE等命令的工作可能会受到影响。可用于配置文件。
console
仅用于Windows平台的选项，用于实现将错误日志信息发送至标准输入和错误输出，即使配置了–log-error选项，此功能也一样有效。
concurrent_insert={NEVER|AUTO|ALWAYS} 或分别使用{0|1|2}
设定是否允许在MyISAM表上并行执行INSERT和SELECT语句。作用范围为全局，可用于配置文件，属动态变量。
connect_timeout=#
mysqld服务器端在响应“失败的握手操作”信息给客户端之前所等待的秒数，默认为10秒。作用范围为全局，可用于配置文件，属动态变量。
core-file
当MySQL进程宕掉时将信息保存为一个core文件，在Linux平台上，core文件通常被保存至当前进程的工作目录中，并命名为core.pid，其文件名后缀pid为当前进程的进程号；对MySQL而言，保存目录为数据文件目录。
datadir=PATH, -h PATH
指定MySQL服务的数据目录。作用范围为全局，可用于配置文件中，但属于非动态变量。
default_storage_engine={Engine_Name}
设定MySQL服务器的默认存储引擎。MySQL 5.5.5版本之前默认为MyISAM，之后的版本默认为InnoDB。作用范围为全局，可用于配置文件，属动态变量。
delay-key-write={ON|OFF|ALL}
仅用于MyISAM表，且要求在创建表时使用了DELAY_KEY_WRITE选项。在启用时，key buffer不会在每一次索引更新时都予以清空，而是在表关闭时才执行key buffer清空操作。OFF表示忽略DELAY_KEY_WRITE，ON表示MySQL接受CREATE TABLE时使用的任何DELAY_KEY_WRITE选项，ALL表示所有新打开的表遵循此特性。
error-count
上一条SQL语句导致的错误信息的数目，此为只读变量。
event-scheduler={ON|OFF|DISABLED}
设定MySQL服务器是否启用以及否启动Event Scheduler。OFF表示停止，此为默认值；ON表示启动，其处于运行状态并执行所有的调度事务；DISABLED表示禁用Event Scheduler，即其不能切换为启动状态。作用范围为全局，可用于配置文件，属动态变量。
expire_logs_days={0..99}
设定二进制日志的过期天数，超出此天数的二进制日志文件将被自动删除。默认为0，表示不启用过期自动删除功能。如果启用此功能，自动删除工作通常发生在MySQL启动时或FLUSH日志时。作用范围为全局，可用于配置文件，属动态变量。
external_user=name
在MySQL服务器上基于认证插件进行用户认证时，此插件会把发起连接请求用户当作另一个用户以达到权限检查的目的，这样可以使得外部用户作第二用户的代理用户，并拥有第二用户的所有权限。当使用MySQL的内部认证机制或没有插件为其设定值时，此变量的值为NULL。作用范围为会话级别，不可用于配置文件，属非动态变量。
flush={ON|OFF}
设定MySQL服务器是否单独为每个SQL语句执行数据同步(将数据写入磁盘)。正常情况下，MySQL为每个语句执行数据同步工作，并将后续的同步过程交由操作系统完成。默认为OFF。作用范围为全局，可用于配置文件，属动态变量。
flush-time={0..}
为非0值时，MySQL服务器会将所有打开的表每隔flush_time指定的时长进行关闭，使用其释放所有资源并将数据同步至磁盘中。只有在系统资源极其稀缺的情况下才需要启用此功能。默认值是0，即为禁用此功能。作用范围为全局，可用于配置文件，属动态变量。
foreign-key-checks={0|1}
设定是否为InnoDB表查检外键约束，默认为1，即检查。在不确保按原有顺序重新装载所有InnoDB表时，禁用此功能会避免外键约束的副作用。
general_log={ON|OFF}
设定是否启用查询日志，默认值为取决于在启动mysqld时是否使用了–general_log选项。如若启用此项，其输出位置则由–log_output选项进行定义，如果log_output的值设定为NONE，即使用启用查询日志，其也不会记录任何日志信息。作用范围为全局，可用于配置文件，属动态变量。
general_log_file=FILE_NAME
查询日志的日志文件名称，默认为“hostname.log"。作用范围为全局，可用于配置文件，属动态变量。
group_concat_max_len={4..}
设定GROUP_CONCAT()函数返回值的最大长度，默认为1024。有效取值范围为4至“2^CPU字长”次方。作用范围为全局或会话级别，用于配置文件，属动态变量。
have-compress={YES|NO}
zlib压缩库是否能为MySQL服务器所用。当其值为NO时，COMPRESS()和UNCOMPRESS()函数均不可用。
have_crypt={YES|NO}
crypt()系统调用是否可为MySQL服务器所用。当其值为NO时，ENCRYPT()函数则不可用。
have_csv={YES|NO}
mysqld支持CSV引擎时为YES，否则为NO。
have_dynamic_loading={YES|NO}
mysqld支持动态加载插件时为YES，否则为NO。
have_geometry={YES|NO}
mysqld支持空间数据类型时为YES，否则为NO。
have_innodb={YES|NO}
mysqld支持InnoDB存储引擎时为YES，否则为NO。
have_openssl={YES|NO}
此为have_ssl选项的别名；
have_ssl={YES|NO}
mysqld支持SSL连接时为YES，否则为NO。DISABLED表示mysqld编译时启用了对SSL的支持，但在启动mysqld时没能使用正确的ssl-xxx类(如ssl_cert)的选项。
have_partitioning={YES|NO}
mysqld是否支持partitioning，此选项已经基本废弃，且在MySQL-5.6中已经移除，使用SHOW ENGINES可获取此相关信息。
have_profiling={YES|NO}
mysqld支持语句性能分析时则为YES，否则为NO。如果支持profiling功能，则–profiling变量则用于控制是否启动此功能。
have_query_cache={YES|NO}
mysqld支持查询缓存则为YES，否则为NO。
have_rtree_keys={YES|NO}
mysqld支持RTREE索引则为YES，否则为NO。RTREE索引用于MyISAM表的空间索引。
have_symlink={YES|NO}
mysqld支持符号链接则为YES，否则为NO。在Unix主机上，此功能对数据目录和索引目录有用。
hostname=STRING
mysqld服务器启动时将主机名称赋值给此变量。作用范围为全局，属非动态变量。
identity
last_insert_id变量的同义词，其存在的主要目的是为了兼容其它数据库系统。会话级别的变量。
init_connect=STRING
设定在每个客户端与mysqld建立连接时事先执行的一个或多个(彼此间用分号隔开)SQL语句，但对于具有SUPER权限的用户来说，此功能无效。例如，在5.5.8之前的MySQL中尚未出现autocommit变量，此时若要为每位用户默认禁用autocommit功能，就可以在mysqld的配置文件中使用init_connect='SET autocommit=0'来实现，当然也可以使用SET GLOBAL init_connect='SET autocommit=0';命令完成。作用范围为全局级别，可用于配置文件，属动态变量。
init-file=/PATH/TO/SOMEFILE
定义在mysqld启动时使用的初始化文件，此文件每行包含一个单独的SQL语句(不能有注释，不需要且不能使用语句结束符)，并会在mysqld启动时逐个执行。
insert_id
为某表中设定了AUTO_INCREMENT的字段执行INSERT或ALTER_TABLE语句时将使用此变量的值。主要为二进制日志所用。
interactive_timeout=#
mysqld进程等待一个已经建立连接的交互式客户端的后续命令之前所经过的秒数，默认为28800。作用范围为全局或会话级别，用于配置文件，属动态变量。
join_buffer_size=#
mysqld用于平面索引扫描(plain index scans)、范围索引扫描或不使用索引的全表扫描时所能够使用的最小缓冲。正常情况下，添加索引是加快连接执行速度的有效手段，而无法添加索引时，增大join_buffer_size的值可以加快完全连接的执行速度。两表之间的每个完全连接会使用一个单独的join buffer，多表之间的非基于索引的复杂完全连接则有可能使用多个join buffer。将此变量值设定的大过每个匹配的行的大小等并不能带来太多的益处，因此，不应该在全局范围内将此值设定的过大。建议使用较小的全局设定，只为需要的会话中使用较大连接时设定较大值。其最大值取决平台，如32bit平台上的最大值为4G。
keep_files_on_create={ON|OFF}
此项默认值为OFF。创建MyISAM类型的表时，mysqld会在数据目录中为其创建一个.MYD文件和一个.MYI文件，如果数据目录中已经存在一个同名的文件，默认设定为覆盖操作，当设定此变量为OFF时，则会返回一个错误信息。作用范围为全局或会话级别，可用于配置文件，属动态变量。
key_buffer_size=#
所有线程共享的、用于MyISAM表的索引缓冲空间大小，其也通常被称作key cache。在32位平台上其最大值为4G，64位平台上允许使用更大的值，但较为有效的值取决于可用物理RAM资源的大小和每进程可用RAM大小的限制。mysqld启动时会尽可能分配接近于指定大小的RAM空间给key_buffer_size，而非一定是指定大小。增大此值可以加速读写操作时对索引的处理速度，因此，在一个以MyISAM为主要表类型的应用场景中可以将此值设定到物理内存空间的25%，然而，比此者再大的值则反而可能引起系统性能下降了，比如设定到物理内容间的50%时则可能带来严重的性能问题。更何况，还需要考虑其它类型存储引擎对内存的需要。
当向表中同时插入多行数据时，使用LOCK TABLES会加速其执行过程。当然，也可以观察SHOW STATUS命令输出中的Key_read_requests, Key_reads, Key_write_requests和Key_writes值也判定mysqld的性能表现。正常情况下，Key_reads/Key_read_requests的比值应该小于0.01，而Key_writes/Key_write_requests的比值通常会接近于1，不过在启用了DELAY_KEY_WRITE选项的场景中，这个比值可能会更小。
key_cache_block_size=#
MyISAM存储引擎的索引存放于“.MYI”文件中，每个“.MYI”文件由文件头和实际的索引数据共同组成。在“.MYI”的相关概念中，其逻辑上表现为多个Index Block，但并非物理结构。在物理上，索引是以文件块(File Block)的形式来存放在磁盘上面的。在Key Cache中缓存的索引信息是以缓存块(Cache Block)的形式组织存放的，缓存块是一组相同大小的存储空间，和“.MYI”文件物理存储的Block(File Block)类似。
在一条查询语句通过索引检索表数据的时候，首先会检查索引缓存(key_buffer_cache)中是否已经存储了需要的索引信息，如果没有，则会读取“.MYI”文件，将相应的索引数据读入Key Cache中的内存空间中，并存储为缓存块格式。此时，如果整个Key Cache中已经没有空闲的缓存块空间可以使用的话，mysqld将会通过LRU算法将某些缓存块予以清除。
key_cache_block_size参数则用于设定cache block的大小，默认为1024。作用范围为全局级别，可用于配置文件，属动态变量。
key_cache_division_limit=#
实际上，在MySQL的Key Cache中所使用的LRU算法并不像传统的算法一样仅仅只是通过访问频率以及最后访问时间来通过一个唯一的链表实现，而是将其分成了两部分。一部分用来存放使 用比较频繁的Hot Cache Lock(Hot Chain)，被称作Hot Area，另外一部分则用来存放使用不太频繁的Warm Cache Block(Warm Chain)，也被称作Warm Area。这样做的目的主要是为了保护使用比较频繁的Cache Block更不容易被换出。而key_cache_division_limit参数则正是用于告诉MySQL该如何划分整个Cache Chain划分为Hot Chain和Warm Chain两部分，参数值为Warm Chain占整个Chain 的百分比值。设置范围1～100，系统默认为100，也就是只有Warm Chain。
key_cache_age_threshold=#
控制Hot Area中的Cache Block何时该被降级到Warm Area中。系统默认值为300，最小可以设置为100。值越小，被降级的可能性越大。
large_files_support={YES|NO}
mysqld是否在编译时的编译选项中指定了支持大文件。其作用域为全局，为非动态变量。
large_pages={YES|NO}
Linux平台上专用的参数，用于设定mysqld是否支持使用大内存页。使用大内存而可以提高TLB的命中率，进行提高系统性能。其作用域为全局，可以用于配置文件中，为非动态变量。
large_page_size=#
Linux平台上专用的参数，用于设定mysqld使用的大内存页的大小，一般为4MB，在其它平台上此参数的值为0，即为禁用。其作用域为全局，为非动态变量。
last_insert_id
此参数的值由LAST_INSERT_ID()函数返回，在更新表的语句中使用LAST_INSERT_ID()时其对应的确切值会存储于二进制日志中。
lc_messages=STRING
错误信息的区域设定(即语言区域)，mysqld将此值转换为语言名称，并结合lc_messages_dir参数指定的路径中的区域相关的语言文件来返回错误信息。作用范围为全局或会话级别，可用于配置文件，属动态变量。
lc_messages_dir=/PATH/TO/SOME_DIR
错误信息的存储目录，通过与lc_messages参数设定的语言区域来返回错误信息。作用范围为全局或会话级别，可用于配置文件，属动态变量。
lc_time_names=STRING
设定基于语言区域来显示日、月及其简写方式等日期信息的语言区域，其值如en_US等，但与系统的locale无关。此设定将影响DATE_FORMAT(), DAYNAME()和MONTHNAME()函数的输出结果。作用范围为全局或会话级别，属动态变量。
local_infile={YES|NO}
设定mysqld是否支持使用LOAD DATA INFILE语句。默认为ON。作用范围为全局级别，属动态变量。
local_wait_timeout=#
以秒为单位设定所有SQL语句等待获取元数据锁(metadata lock)的超时时长，默认为31536000(1年)，有效取值范围为0-31536000。其影响的SQL语句包括用于表、视图、存储过程和存储函数的DML和DDL语句，以及LOCK TABLES、FLUSH TABLES WITH READ LOCK和HANDLER语句等。但其作用的所有对象不包括mysql数据库中的系统表及用于记录日志的表上的GRANT或REVOKE语句，但包括SELECT或UPDATE语句。
另外，此超时时长分别应用于每个元数据锁，因此，一个语句可能会持有多个元数据锁，那么其最后的生效超时时长有可能会长时这个设定值。作用范围为全局或会话级别，可用于配置文件，属动态变量。
locked_in_memory={YES|NO}
mysqld是否使用–memlock选项锁定在了内存中。作用范围为全局级别，属非动态变量。
log={YES|NO}
是否启用记录所有语句的日志信息于一般查询日志(general query log)中，默认通常为OFF。MySQL 5.6已经弃用此选项。
log-bin={YES|NO}
是否启用二进制日志，如果为mysqld设定了–log-bin选项，则其值为ON，否则则为OFF。其仅用于显示是否启用了二进制日志，并不反应log-bin的设定值。作用范围为全局级别，属非动态变量。
log_bin_trust_function_creators={TRUE|FALSE}
此参数仅在启用二进制日志时有效，用于控制创建存储函数时如果会导致不安全的事件记录二进制日志条件下是否禁止创建存储函数。默认值为0，表示除非用户除了CREATE ROUTING或ALTER ROUTINE权限外还有SUPER权限，否则将禁止创建或修改存储函数，同时，还要求在创建函数时必需为之使用DETERMINISTIC属性，再不然就是附带READS SQL DATA或NO SQL属性。设置其值为1时则不启用这些限制。作用范围为全局级别，可用于配置文件，属动态变量。
log_error=/PATH/TO/ERROR_LOG_FILENAME
定义错误日志文件。作用范围为全局或会话级别，可用于配置文件，属非动态变量。
log_output={TABLE|FILE|NONE}
定义一般查询日志和慢查询日志的保存方式，可以是TABLE、FILE、NONE，也可以是TABLE及FILE的组合(用逗号隔开)，默认为TABLE。如果组合中出现了NONE，那么其它设定都将失效，同时，无论是否启用日志功能，也不会记录任何相关的日志信息。作用范围为全局级别，可用于配置文件，属动态变量。
log_query_not_using_indexes={ON|OFF}
设定是否将没有使用索引的查询操作记录到慢查询日志。作用范围为全局级别，可用于配置文件，属动态变量。
log_slave_updates|
用于设定复制场景中的从服务器是否将从主服务器收到的更新操作记录进本机的二进制日志中。本参数设定的生效需要在从服务器上启用二进制日志功能。
log_slow_queries={YES|NO}
是否记录慢查询日志。慢查询是指查询的执行时间超出long_query_time参数所设定时长的事件。MySQL 5.6将此参数修改为了slow_query_log。作用范围为全局级别，可用于配置文件，属动态变量。
log_warnings=#
设定是否将警告信息记录进错误日志。默认设定为1，表示启用；可以将其设置为0以禁用；而其值为大于1的数值时表示将新发起连接时产生的“失败的连接”和“拒绝访问”类的错误信息也记录进错误日志。
long_query_time=#
设定区别慢查询与一般查询的语句执行时间长度。这里的语句执行时长为实际的执行时间，而非在CPU上的执行时长，因此，负载较重的服务器上更容易产生慢查询。其最小值为0，默认值为10，单位是秒钟。它也支持毫秒级的解析度。作用范围为全局或会话级别，可用于配置文件，属动态变量。
low_priority_updates={TRUE|FALSE}
设定是否降低更新操作的优先级，仅对只支持表级别锁的存储引擎有效，如MyISAM、MEMORY或MERGE。其值为1则表示所有的INSERT、UPDATE、DELETE或LOCK TABLE WRITE语句只能在没有等待执行的SELECT或LOCK TABLE READ语句时才能执行。作用范围为全局或会话级别，可用于配置文件，属动态变量。
lower_case_file_system={ON|OFF}
用于描述数据目录所在的文件系统是否区分文件名称字符大小写，OFF表示区分大小写，ON表示不区分大小写。此变量是只读的，其是否区分大小写取决于文件系统。
lower_case_table_name={0|1|2}
设定是否区分表、表别名或者数据库名称中的字符大小写。0表示区分大小写，1表示不区分大小写且一律存储为小写字符，2表示按给定的大小写进行存储但不区分大小写。对于本身不支持区分文件名称大小写功能的文件系统来讲，不应该设定为0值；而在设定为0值的系统上，对于MyISAM存储引擎来说，使用大小写不对应的名称进行访问可能会导致索引文件崩溃。Windows系统上，其默认值为1，Mac OS X上其默认值为2。
对于InnoDB存储引擎来说，应该将其设定为1，无论基于什么平台。同时，也不应该为主从复制集群中的主机使用不同的设定，不然，其可能导致复制失败。作用范围为全局级别，可用于配置文件，属非动态变量。
max_allowed_packet={YES|NO}
设定单个报文或任何中间字符串(intermediate string)的最大长度，单位是字节。报文消息缓冲由net_buffer_length参数进行设定，但其最终可以按需增长至max_allowed_packet参数设定的大小。此参数的默认值较小，在使用了BLOB列或长字符串的场景中，应该增大其值至能容纳最大BLOB数据的长度。协议本身限定此值最大为1G，参数只接受1024整数倍的数值，非1024的整数倍将会被自动圆整至离其最近的1024整数倍的数值。
最终生效的报文长度还取决于客户端的设定。客户端程序如mysql或mysqldump的默认设定为1G。作用范围为全局级别，可用于配置文件，属动态变量。
max_connect_errors=#
设定客户端连接至mysqld时的最大错误尝试次数。在某客户端尝试连接当前mysqld的错误次数连接达到max_connect_errors所设定的值时，其后续的连接尝试将被直接阻止。管理员可以通过FLUSH HOSTS语句或mysqladmin flush-hosts命令清空主机缓存(host cache)来解除对此前阻止主机的访问锁定。如果某客户端的在其错误尝试次数达到此参数设定的值之前成功建立了连接，其错误尝试次数的计数器将会被清空。作用范围为全局级别，可用于配置文件，属动态变量。
max_connections=#
设定mysqld允许客户端同时发起的最大并发连接数。增加此值将增加mysqld进程需要同时访问的文件描述符数目。作用范围为全局级别，可用于配置文件，属动态变量。
max_delayed_threads=#
设定为INSERT DELAYED语句所能够启动的最大线程数。如果当前相关的线程数目已经达到此参数所设定的值，后续的INSERT DELAYED语句将无视其DELAYED属性。如果将其值高精为0，mysqld将不为INSERT DELAYED创建任何线程，即禁用DELAYED功能。作用范围为全局级别，可用于配置文件，属动态变量。
max_error_count=#
设定为SHOW ERRORS或SHOW WARNINGS语句所保留的关于错误、警告或注意信息条目的最大数。作用范围为全局或会话级别，可用于配置文件，属动态变量。
max_heap_table_size=#
设定每个用户创建的MEMORY表所能够使用的最大内存空间。修改其值对当前已经创建的MEMORY表没有影响，除非使用CREATE TABLE、ALTER TABLE或TRUNCATE TABLE对表进行了重建。其最在小值为16384，单位是字节。最大值受限于平台字长，如32位平台为4G。此参数与tmp_table_size参数联合使用可用于限制内部内存表的大小。另外，max_heap_table_size不会被复制。作用范围为全局或会话级别，可用于配置文件，属动态变量。
max_insert_delayed_threads=#
max_delayed_threads的同义词。作用范围为全局级别，动态变量。
max_join_size=#
设定SELECT语句执行时所能够检查的行数(单表)或行组合(多表查询)的最大值。此参数可以阻止对键的错误使用而导致的需要执行较长时间的查询操作，因此，其尤其适用于有用户经常使用不带WHERE子句的查询场景中。有效取值范围为1-18446744073709551615，默认为18446744073709551615，这可以被理解为不限制。作用范围为全局或会话级别，可用于配置文件，属动态变量。
max_length_for_sort_data=#
filesort算法改进版所能够使用的字段最大长度值。有效取值范围是4-8388608。MySQL的filesort算法有两个版本，即原始版本和修改版本，字段长度大于max_length_for_sort_data设定的将使用原始版本，小于此参数值的则使用修改版本在排序缓冲(sort buffer)中完成排序。在使用超出字段超出指定长度时使用修改版本算法，由于可能需要更多的I/O操作，将会导致修改版算法执行速度更慢，而不是更快。作用范围为全局或会话级别，可用于配置文件，属动态变量。
max_long_data_size=#
设定可以由mysql_stmt_send_long_data()这个C API函数所传送的参数值的最大长度，如果没有在mysqld启动时设定，其默认为max_allowed_packet变量的值。MySQL 5.6已经弃用此变量。 作用范围为全局级别，可用于配置文件，属非动态变量。
max_prepared_stmt_count={0..1048576}
设定mysqld所允许的所有连接会话中执行的准备语句的总数。大量的准备语句同时执行会消耗大量的内存资源，这会带来潜在的“拒绝服务”的风险，因此，应该根据生产需要仔细设定此参数的值。如果新设定的值低于目前已经打开的准备语句总数，其不会对原有的语句产生影响，但不再接受新的执行请求，直到有新的空余额度。默认值是16382，0表示禁用准备语句。作用范围为全局级别，可用于配置文件，属动态变量。
max_relay_log_size={4096..1073741824}
设定从服务器上中继日志的体积上限，到达此限度时其会自动进行中继日志滚动。此参数值为0时，mysqld将使用max_binlog_size参数同时为二进制日志和中继日志设定日志文件体积上限。作用范围为全局级别，可用于配置文件，属动态变量。
max_seeks_for_key={1 .. 18446744073709547520}
设定基于某key执行查询时所允许的最大查找次数。在通过扫描索引的方式在某表中搜寻与查询条件匹配的行时，无论其索引的基数是什么，MySQL优化器都会假定其查找次数不需要超过此参数设定的数值。较小的值可以强制MySQL更倾向于索引扫描而非表扫描。作用范围为全局或会话级别，可用于配置文件，属动态变量。
max_sort_length={4 .. 8388608}
设定mysqld执行数值排序时所使用的字节数，剩余的将被忽略。作用范围为全局级别，可用于配置文件，属动态变量。
max_sp_recoursion_depth={0 .. 255}
设定存储过程可被递归调用的最大次数。递归调用会增大对线程栈空间的需要，因此，增大此参数的值，可能还需要在启动时调整thread_stack参数的值。默认值为0，表示禁止递归；最大值为255。作用范围为全局级别，可用于配置文件，属动态变量。
max_user_connections={ 0 .. 4294967295 }
设定单个用户允许同时向mysqld发起的最大并发连接请求个数。默认值为0，表示无上限。可为mysqld为此参数指定全局(GLOBAL)参数值，也可将某用户帐号的此参数值设定为只读以为其设定有效并发上限(通过GRANT语句实现)。这意味着，如果用户的此参数有个非0值，则以此限定为准；否则，mysqld会将用户的此参数值设定为全局值。作用范围为全局或会话级别，可用于配置文件，属动态变量。
max_write_lock_count=#
mysqld已施加的写锁个数达到此参数值指定的个数时，将允许处理一些挂起的读请求。其最小值为1，最大值取决于平台字长。作用范围为全局级别，属非动态变量。
metadata_locks_cache_size={1 .. 1048576}
设定mysqld元数据锁缓存的上限。此缓存可用来避免创建或销毁同步对象(synchronization object)，这对于此类操作代价较高的操作系统(如Windows XP)来说尤为有用。默认值为1024。作用范围为全局级别，属非动态变量。
min_examined_row_limit=#
所检查的行数低于此参数设定的数值的查询操作将不被记入慢查询日志。默认值为0，最大值取决于平台字长。作用范围为全局或会话级别，可用于配置文件，属动态变量。
myisam_data_point_size={2 .. 7}
创建MyISAM表时如果没有设定MAX_ROWS选项，则通过此参数设定其默认指针的大小，单位为字节。默认值是6。作用范围为全局级别，可用于配置文件，属动态变量。
myisam_max_sort_file_size=#
设定在MySQL在使用REPARE TABLE、ALTER TABLE或LOAD DATA INFILE命令时为MyISAM表重新创建索引所能够使用的临时文件的体积上限，单位是字节。如果临时文件的大小大过了此上限值，则mysqld会使用key cache创建索引。默认值是2G，而如果MyISAM索引文件本身大过此值且其所在的文件系统有足够的空闲空间，增大此值会提升MySQL性能。作用范围为全局级别，可用于配置文件，属动态变量。
myisam_mmap_size=#
设定基于内存映射压缩MyISAM文件时可以使用的内存上限。在有着很多压缩格式MyISAM表的场景中，降低此值可以帮助减少出现内存交换的可能性。其最小值为7，默认值和最大值取决于平台位数。作用范围为全局级别，可用于配置文件，属非动态变量。
myisam_recover_options
用于保存mysqld命令行选项–myisam-revover-options的值。此选项用于设定MyISAM存储引擎的恢复模式，其可接受的值有OFF、DEFAULT(恢复模式，但无备份、强制或快速检测)、BACKUP(如果恢复过程中数据发生了改变，则将原表文件备份为table_name-datatime.BAK)、FORCE(强制恢复，哪怕会丢失数据)和QUICK(快速修复)，可以以逗号为分隔为此选项同时指定多个值。也可以不为其提供任何参数值，默认为DEFAULT，而“”则表示为OFF，即禁用恢复模式。
如果启用恢复模式，mysqld每次打开一个MyISAM表时都会检测其是否标记为损坏或非正常关闭。如果损坏，则mysqld会尝试修复它；如果为非正常关闭，mysqld将会对其进行检测。
myisam_repair_threads=#
在通过排序修复过程中为MyISAM表创建索引的线程的个数，默认为1；如果给定大于1的值，则可以启动多个并行创建索引的线程(每个索引只能由一个线程创建)。最大值取决于平台位数。作用范围为全局或会话级别，可用于配置文件，属动态变量。
myisam_sort_buffer_size=#
在REPAIR TABLE过程中，或通过CREATE INDEX/ALTER TABLE为MyISAM表添加索引时为了对索引排序所能够使用的缓冲空间大小。最小值为4，32位系统上所能够使用的最大值为4294967295，即4G；64位系统上可以使用更大的空间。作用范围为全局或会话级别，可用于配置文件，属动态变量。
myisam_stats_method={nulls_equal|nulls_unequal|nulls_ignored}
定义在为MyISAM表收集索引分布相关的统计信息时处理NULL值的方式。nulls_equal表示所有的NULL值都视作相同值，nulls_unequal表示所有的NULL值都视作不同值，nulls_ignored表示所有的NULL值都将被忽略。作用范围为全局或会话级别，可用于配置文件，属动态变量。
myisam_use_mmap={ON|OFF}
在读写MyISAM表时能否使用内存映射。默认值为OFF。作用范围为全局级别，可用于配置文件，属动态变量。
net_buffer_length={1024 .. 1048576}
每个客户端线程都有一个连接缓冲(connection buffer)和一个结果缓冲(result buffer)，此参数可以设定这两个缓冲的大小，但它们都可以按需动态增长至max_allowed_packet参数所设定的大小。但每个SQL语句完成后，结果缓冲都会收缩至net_buffer_length参数所定义的大小。一般说来不需要修改此参数的值，除非是内存资源较吃紧的场景中可以将其调小至客户端预期的SQL语句长度。如SQL语句超出此长度，连接缓冲会自动调节其值。其默认值为16384，单位是字节。作用范围为全局或会话级别，可用于配置文件，属动态变量，但对于会话级别来说，此变量是只读的。
net_read_timeout=#
设定mysqld等待从客户端接收更多数据的超时时长，默认值为30。作用范围为全局和会话级别，可用于配置文件，属动态变量。
net_write_timeout=#
设定mysqld等待向客户端传输数据的超时时长，默认值为60。作用范围为全局和会话级别，可用于配置文件，属动态变量。
net_retry_timeout=#
设定mysqld与客户端的通讯中断时，其中止与客户端的连接之前需要重试的次数。默认值为10，其最大取值取决于平台位数。作用范围为全局和会话级别，可用于配置文件，属动态变量。
new={ON|OFF}
用于MySQL 4.0中以启用支持MySQL 4.1版本上的某些新特性的能力，但仍然可以保持向后兼容。在MySQL 5.5无须设置，故其值为OFF。
old={ON|OFF}
用于定义兼容老版本MySQL的变量，默认是禁用的，但可以在mysqld启动时将其启用以兼容较老的MySQL版本。
innodb_adaptive_flushing={ON|OFF}
设定是否允许MySQL服务器根据工作负载动态调整刷写InnoDB buffer pool中的脏页的速率。动态调整刷写速率的目的在于避免出现IO活动尖峰。默认值为ON。作用范围为全局级别，可用于选项文件，属动态变量。
innodb_adaptive_hash_index={ON|OFF}
设定是否启用InnoDB的自适应hash索引。基准测试结果显示，自适应hash索引并非对所有工作负载都有益，因此需要根据实际应用场景的测试结果选择更合适的值。此特性默认已启用，可以通过命令行选项–skip-innodb_adaptive_hash_index将其禁用。作用范围是全局，可用于选项文件，属动态变量。
innodb_additional_mem_pool_size={2097152 .. 4294967295}
设定innodb存储引擎为了存储数据字典和其它内部数据结构的内在池大小，单位是字节。表的个数越多，此参数的值就应该设定的越大；当InnoDB用完此内存池的空间，它就会向操作系统申请内存空间，并将向错误日志记录警告信息。默认大小是8MB。作用范围为全局，可用于选项文件，属非动态变量。
innodb_autoextend_increment={1 .. 1000}
当共享表空间没有多余的存储空间时，如果其允许自动增长，此变量可用于设定其单次增长的空间大小，单位是MB，默认值是8。设置了变量innodb_file_per_table的值为1时InnoDB会为每张表使用一个单独的表空间文件，而innodb_autoextend_increment变量不会对此种表空间产生影响。作用范围为全局级别，可用于选项文件，属动态变量。
innodb_autoinc_lock_mode={0|1|2}
设定用于生成“自动增长(auto_increment字段)”值的锁模型。其可接受的值有0、1和2，分别用于表示"traditional"、"consecutive"和"interleaved"锁模型。默认值为1。作用范围为全局，可用于选项文件，属非动态变量。
innodb_buffer_pool_instances=#
设定将InnoDB的buffer pool分隔为多少个区域。对于有着数GB空间的buffer pool来说，将其分隔为多个区域可以降低不同的线程对缓存页面的读写操作时资源争用系数，进行增强其并发能力。在buffer pool中，读取或存入页面时所选择的区域是基于hash算法随机进行的。每个buffer pool管理自己的空闲列表、列表刷写、LRU以及其它跟buffer pool相关的数据结构，并通过各自的互斥锁进行保护。
此变量仅在变量innodb_buffer_pool_size的值大于1G时才能发挥功用，缓冲池的整体空间将由各buffer pool实例分割使用。出于最佳效用的目的，建议配合使用innodb_buffer_pool_instances和innodb_buffer_pool_size变量以使得每个buffer pool实例的都至少有1G的空间。作用范围为全局，可用于选项文件，属非动态变量。
innodb_buffer_pool_size=#
设定InnoDB缓存表数据和索引的内存缓冲区大小，单位是字节。其默认值为128MB，最大值依赖于CPU架构。在一个较繁忙的服务器上，当缓冲池(buffer pool)大于1G时，设定innodb_buffer_pool_instances的值大于1可提其升伸缩能力。innodb_buffer_pool_size变量的值越大，MySQL服务器完成数据访问时就需要越少的IO，因此，在一个有够较大内存且为MySQL服务专用的服务器上，可以将此值设置为物理内存的80%。但如果出现如下情况，建议缩小此变量的值：(1)物理内存资源紧张导致内存页面换出；(2)InnoDB会为缓冲和控制结构(buffers and control structures)预留额外的内存，因此事实上其占用的内存空间可能会比指定的数值大10%左右，这不可能超出对内存资源分配的预估；(3)内存地址空间必须连续，这在基于DLL库使用特殊地址空间的Windows系统上可能会出现意外情况；(4)缓冲池的初始化所需要时长与为其指定的空间大小成正比，例如有10G缓冲池的x86_64的Linux系统上，初始化时间大约要6秒钟。作用范围为全局，可用于选项文件，属非动态变量。
innodb_change_buffering=#
当在表上执行INSERT、UPDATE或DELETE操作时，索引中尤其是第二索引中的数据未必按序存储，这就可能引发随机IO以完成第二索引的更新操作。此变量用来设定InnoDB是否启用修改缓冲(change buffering)以及使用何种类型的修改缓冲。修改缓冲是一种优化方式，它能够通过延迟写入操作至第二索引将IO操作转换为顺序模式。其接受的值有inserts(缓冲insert操作)、deletes(缓冲delete-marking操作)、changes(缓冲insert和delete-marking操作)、purges(缓冲purge操作)、all(缓冲insert、delete-marking和purge操作)和none(不缓冲任何操作)。默认值是all。MySQL 5.5.4之前的版本只接受inserts和none两种值。作用范围为全局，可用于选项文件，属动态变量。
innodb_checksums={ON|OFF}
InnoDB能够使用校验和(checksum)来验正从磁盘读取的所有页面数据的完整性，从而提高对硬件或数据文件损坏的容错能力。默认为启用，然而，在少数情况下或许需要禁用这种特性，这可以通过使用–skip-innodb-checksums命令行选项实现。作用范围为全局，可用于选项文件，属非动态变量。
innodb_commit_concurrency={0 .. 1000}
设定InnoDB可同时运行的“提交”操作线程的数量。0表示无限制。此变量不能在运行时将其从“零值”修改为“非零值”，但可以从一个“非零值”修改为其它值。作用范围为全局，可用于选项文件，属非动态变量。
innodb_concurrency_tickets=#
在一个线程进入(enter)InnoDB时，其将会获取一定数量的“自由卷轴”(free tickets)并凭这些卷轴自由出入InnoDB(即免检)，直到其卷轴耗尽；而后的线程将被置于等待队列中，并可能需要再次接受并发上限限制检查。此变量则正是用于设定可同时进入InnoDB的线程并发数，即线程的“自由卷轴”数量。默认值是500。作用范围为全局，可用于选项文件，属动态变量。
innodb_data_file_path=IBDATA_FILE
指定InnoDB的各个数据文件及其大小，文件多于一个时彼此间用分号隔开。数据文件路径可以为相对路径，其相对于innodb_data_home_dir变量所指向的目录；而文件大小的表示可以以K(KB)、M(MB)、G(GB)为单位，但这些文件的大小之和至少要达到10MB。在没有显式设定innodb_data_file_path变量的情况下，MySQL服务器会在数据目录中自动创建一个可自动增长、初始大小为10MB的名为ibdata1的数据文件。单个数据文件的大小上限取决于操作系统，这意味着可以使用操作系统所支持的最大单个文件大小以为其数据文件的体积上限。InnoDB还支持使用裸设备作为数据文件。作用范围为全局，可用于选项文件，属非动态变量。
innodb_data_home_dir=/PATH/TO/DIR
InnoDB所有共享表空间数据文件的目录路径。默认值为MySQL的数据目录。可以将此变量的值设置为空，然后在innodb_data_file_path中为每个数据文件使用绝对路径。此变量不影响变量innodb_file_per_table启用状态下的每表表空间的数据文件。作用范围为全局，可用于选项文件，属非动态变量。
innodb_doublewirte={ON|OFF}
设定InnoDB是否使用双写缓冲。默认为启用。InnoDB在对页面进行部分写入的时候使用双写缓冲，以防止数据损坏。双写缓冲是表空间中一个特殊的保留区域，其大小足够在一个连续区间容纳100个页面。当InnoDB把页面从缓冲池刷写至磁盘时，它会先把这些页面刷到双写缓冲中，然后再保存至真正的目标位置。因此，双写缓冲本质上是最近写入页面的备份，其可确保每次写入的原子性和可持续性。在有些情况下双写缓冲是不必要的，例如在从服务器上就可以将之禁用；此外，一些文件系统(如ZFS)自身也会实现此功能，那么InnoDB就不用做重复的工作了。作用范围为全局，可用于选项文件，属非动态变量。
innodb_fast_shutdown={0|1|2}
设定InnoDB关闭模式。其可接受的值中，“0”表示慢速关闭，这意味着InnoDB关闭之前会完成完全清写(full purge)和修改缓冲合并(insert buffer merge)操作；“1”是默认值，它表示InnoDB在关闭时会跳过模式0中进行的这些操作，这也是其之所以称作“快速关闭”的原因；“2”表示InnoDB仅刷写日志信息并执行冷(cold)关闭，此时并没有事务丢失，只是下次启动MySQL服务时需要花费较长的时间进行故障恢复(crash recovery)。
执行慢速关闭时其过程可能会持续数分钟的时间，甚至在有些极端情况下，比如有着大量数据缓冲的场景，此过程时长会以小时来计。一般情况下仅在对MySQL进行主版本升级时才需要进行慢速关闭以使得数据文件能够为完全适应新版本而准备妥当。通常也只能遇到紧急状况或出于调试的目的才需要将此变量的值设定为2，以便让处于有可能损坏风险中的数据执行最快速度的关闭。作用范围为全局，可用于选项文件，属动态变量。
innodb_file_format={Antelope|Barracuda}
设定新建InnoDB表的文件格式。其可接受的参数有Antelope和Barracuda，但这仅对基于变量innodb_file_per_file的每表表空间文件有影响。某些InnoDB特性如表压缩功能仅有Barracuda文件格式支持。作用范围为全局，可用于选项文件，属动态变量。
innodb_file_format_check={ON|OFF}
用于设定InnoDB是否在MySQL服务器启动时检查共享表空间的文件格式标签。检查标签时如果其高于当前InnoDB版本所支持的能力，InnoDB就会产生错误并拒绝启动；否则，对MySQL 5.5.5 及后来的版本来说InnoDB则会设置变量innodb_file_format_max的值为共享表空间的文件格式标签，而对于MySQL 5.5.5之前的版本来说，InnoDB会将共享表空间的文件格式设置为变量innodb_file_format_check的值。作用范围为全局，可用于选项文件，属非动态变量。
innodb_file_format_max={Antelope|Barracuda}
在MySQL服务启动时，InnoDB会将变量innodb_file_format_max的值设置为共享表空间的文件格式标签(比如，Antelope或Barracuda)。如果MySQL服务器创建或打开了一个有着更高级格式的表，此变量的值则会被设置为那个更高级的格式。作用范围为全局，可用于选项文件，属动态变量。
innodb_file_per_table={ON|OFF}
设定InnoDB表是否使用每表表空间数据文件(以.ibd结尾)分别存储每个表的数据和索引。如果使用了每表表空间数据文件，其将不再使用系统表空间(即共享表空间)。InnoDB表的某些特性，如压缩表等仅对每表表空间生效。作用范围为全局，可用于选项文件，属动态变量。
innodb_flush_log_at_trx_commit={0|1|2}
设定InnoDB同步日志缓冲区(log buffer)数据至日志文件中的方式，以及刷写日志文件至磁盘的方式。其可接受的值中，“0”表示将日志缓冲区每秒一次地写入日志文件，并同时将日志文件刷写至磁盘中，但事务提交时不会采取任何动作；“1”是默认值，表示在有事务提交时将日志缓冲区写入日志文件，并同时将日志文件刷写至磁盘；“2”表示每事务提交或每秒一次将日志缓冲区写入日志文件，但不会同时执行日志文件的刷写操作。当然，由于操作系统进程调度的原因，每秒一次的日志写入或刷写操作并不能得到100%的保证。
完全兼容ACID的场景需要将此变量值设置为1，由于要执行每事务的日志刷写操作，其会阻止I/O调用，直到写操作完成，故其会显著降低InnoDB每秒钟可以提交的事务数。设置为“2”可获得比“1”更好的性能，而且仅在操作系统崩溃时才会丢失最后一秒钟的数据，因此数据安全性也有着不错的表现。设置为“0”则有可能会导致事务最后一秒钟的数据丢失，于是整个事务的数据安全性将无法保证，但其通常有着最好的性能。为了在最大程序上保证复制的InnoDB事务持久性和一致性，应该设置变量innodb_flush_log_at_trx_commit=1以及设置变量sync_binlog=1。
然而需要注意的是，有些磁盘自身也有缓存，这可能会给事务操作带来额外的潜在风险。可以使用hdparm工具或供应商的自有工具等禁用磁盘自身的缓存。当然，高性能事务的最佳配置是把此变量的值设置为1，并且将日志文件放在有备用电池的写入缓存的RAID上。作用范围为全局，可用于选项文件，属动态变量。
innodb_flush_method={O_DSYNC|O_DIRECT}
设定InnoDB实际与文件系统进行交互的方式。除了写操作之外，它还可以影响InnoDB如何读取数据。设置innodb_flush_method变量的值为O_DSYNC时，InnoDB使用O_SYNC标志来打开和刷写日志文件，而使用fsync()来刷写数据文件。O_SYNC会使得所有的写入操作都是同步的，即只有在数据被写入磁盘之后才会返回，但不会在操作系统层面禁止缓存，因此，它不会避免双缓冲，并且不会直接写入磁盘。fsync()会同时刷数据和元数据(而fdatasync()只刷写数据)，它比fdatasync()产生更多的IO操作，而且操作系统会缓存一些数据在自己的缓存中(这将导致双缓冲)。如文件系统可以智能地处理I/O需求，双缓冲可能不是坏事儿，但如果MySQL设置了innodb_file_per_table变量的值为1，则会导致第个表空间文件都单独使用fsync()函数，其写入操作就不可能会被合并了。
设置innodb_flush_method变量的值为O_DIRECT时，InnoDB使用O_DIRECT标志打开数据文件，而使用fsync()刷写数据和日志文件。O_DIRECT标志会导致操作系统既不缓存数据，也不预读数据，它完全禁止了操作系统的缓存并且使所有的读写动作直接至存储设备，避免了双缓冲。然而，其不能禁止硬件层面(如RAID卡)的缓存和预读功能，而且启用硬件层面的缓存和预读功能也是保证InnoDB使用了O_DIRECT标志时仍能保持良好性能的惟一途径。
作用范围为全局，可用于选项文件，属非动态变量。
innodb_force_load_corrupted={ON|OFF}
设定InnoDB在启动时是否装载标记为“已损坏(corrupted)”的表。仅应该在troubleshooting的场景中启用该功能以修复无法访问的表，在troubleshooting任务完成后应该禁用此功能并重启MySQL服务。作用范围为全局，可用于选项文件，属非动态变量。
innodb_force_recovery={0|1|2|3|4|5|6}
设定InnoDB的故障恢复模式。InnoDB出现了“页面损坏(page corruption)”时，通常大部分数据仍然完好，于是可以通过SELECT…INTO OUTFILE命令备份出数据以降低损失程度。然而，某些“损坏”类的故障可能会导致SELECT * FROM tbl_name命令无法执行或InnoDB后台操作崩溃，甚至会导致InnoDB的前滚操作。这种情况下，就可以使用innodb_force_recovery变量强制InnoDB存储引擎在启动时不执行后台操作，以便能将数据备份出来。
innodb_force_recovery可接受的值中，“0”为默认值，表示执行正常启动，即不启用“强制修复”模式。而非零值中，某数值会包含比其小的所有数值的预防措施，然而其也较可能给B-tree索引及其它的数据结构带来更多的损坏。故此，在此变量值为非零值时，其会阻止用户使用INSERT、UPDATE或DELETE操作，但是会允许执行SELECT、CREATE TABLE或DROP TABLE类的操作。以下是关于其它非零值功能的说明：
1(SRV_FORCE_IGNORE_CORRUPT)：即使出现了页面损坏也照常运行MySQL服务，其会在SELECT * FROM tbl_name语句执行时尝试跳过损坏的索引记录和页面。
2(SRV_FORCE_NO_BACKGROUND)：禁止启动主线程(master thread)，其会在执行清写(purge)操作时防止出现崩溃(crash)。
3(SRV_FORCE_NO_TRX_UNDO)：在故障恢复(crash recovery)后不执行事务的回滚操作。
4(SRV_FORCE_NO_IBUF_MERGE)：禁止执行修改缓冲(insert buffer)合并操作。
5(SRV_FORCE_NO_UNDO_LOG_SCAN)：在启动数据库服务时不检查撤消日志(undo logs)，这会导致InnoDB将未完成的事务视为已提交。
6(SRV_FORCE_NO_LOG_REDO)：不执行重做日志(redo log)的前滚操作。此时，仅能执行不带WHERE、ORDER BY或其它子句的SELECT * FROM tbl_name操作，因为复杂查询在遇到损坏的数据结构时会中止并退出。
innodb_io_capacity=#
设定InnoDB后台任务(如从缓冲池刷写页面或合并修改缓冲中的数据等)可执行的I/O操作上限。其最小值为100，默认值为200，最大值取决于CPU架构。对于有着较大I/O负载的服务器来讲，应该为其指定更大的值以便能够更好更快的执行后台维护任务。然而，在实践中，此变量的值应该尽可能接近MySQL服务器每秒钟执行的I/O操作数量(即IOPS)，甚至于让其低至以不影响后台任务执行为目标的最低限度。因为，如果此值过高的话，数据会被频繁地从缓冲中移入移出，这会降低缓存池的在系统性能提升方面的效用。单个5400RPM或7200RPM磁盘仅能完成大约100个IOPS，因此，此种情况下应该将此变量值降低至100；而对于有着多块磁盘或更强性能的存储设备(如固态磁盘)的应用场景，可以按需提高此变量的值。作用范围为全局，可用于选项文件，属动态变量。
innodb_large_prefix={ON|OFF}
设定对于使用了DYNAMIC或COMPRESSED行格式的InnoDB表来说，是否能够使用大于767字节长度的索引前缀。然而，创建此种类型的表还需要设定innodb_file_format的值为barracuda，以及innodb_file_per_table的值为ture。同时，此设定对使用了REDUNDANT和COMPACT行格式的表的索引长度限定来说是不起作用的。作用范围为全局，可用于选项文件，属动态变量。
innodb_lock_wait_timeout={1 .. 1073741824}
设定InnoDB中某事务试图访问一个由其它InnoDB事务加锁的行时其最长的等待时间，单位为秒钟，默认值为50。在超时情况发生时，InnoDB会返回一个1205类型的错误信息，并对当前语句(非整个事务)执行回滚操作；如果需要在此种情况下对整个事务进行回滚，则需要在MySQL服务启动时使用–innodb_rollback_on_timeout选项。
对于OLTP系统或有着较多交互式应用的程序来说，应该降低此变量值以使得用户较快地获取到反馈信息，或使得系统较块地将此更新操作提交到队列中以便延后处理。对于批处理应用较多的场景来说，如数据仓库，应该增加此变量的值以等待其它较大的插入或更新操作完成。
此变量仅对InnoDB的行锁产生作用，MySQL的表锁并非在InnoDB中实现，所以此超时时长对表锁没有影响。而且，由于InnoDB会能立即探测到死锁的发生并会对其中的一修整务执行回滚操作，因此此超时时长也不应用于死锁。作用范围为全局或会话级别，可用于选项文件，属动态变量。
innodb_locks_unsafe_for_binlog={ON|OFF}
设定InnnoDB是否在搜索和索引扫描中使用间隙锁(gap locking)。InnoDB使用行级锁(row-level locking)，通常情况下，InnoDB在搜索或扫描索引的行锁机制中使用“下一键锁定(next-key locking)”算法来锁定某索引记录及其前部的间隙(gap)，以阻塞其它用户紧跟在该索引记录之前插入其它索引记录。站在这个角度来说，行级锁也叫索引记录锁(index-record lock)。
默认情况下，此变量的值为OFF，意为禁止使用非安全锁，也即启用间隙锁功能。将其设定为ON表示禁止锁定索引记录前的间隙，也即禁用间隙锁，InnoDB仅使用索引记录锁(index-record lock)进行索引搜索或扫描，不过，这并不禁止InnoDB在执行外键约束检查或重复键检查时使用间隙锁。
启用innodb_locks_unsafe_for_binlog的效果类似于将MySQL的事务隔离级别设定为READ-COMMITTED，但二者并不完全等同：innodb_locks_unsafe_for_binlog是全局级别的设定且只能在服务启动时设定，而事务隔离级别可全局设定并由会话级别继承，然而会话级别也以按需在运行时对其进行调整。类似READ-COMMITTED事务隔离级别，启用innodb_locks_unsafe_for_binlog也会带来“幻影问题(phantom problem)”，但除此之外，它还能带来如下特性：
(1)对UPDATE或DELETE语句来说，InnoDB仅锁定需要更新或删除的行，对不能够被WHERE条件匹配的行施加的锁会在条件检查后予以释放。这可以有效地降低死锁出现的概率；
(2)执行UPDATE语句时，如果某行已经被其它语句锁定，InnoDB会启动一个“半一致性(semi-consistent)”读操作从MySQL最近一次提交版本中获得此行，并以之判定其是否能够并当前UPDATE的WHERE条件所匹配。如果能够匹配，MySQL会再次对其进行锁定，而如果仍有其它锁存在，则需要先等待它们退出。
innodb_log_buffer_size={262144 .. 4294967295}
设定InnoDB用于辅助完成日志文件写操作的日志缓冲区大小，单位是字节，默认为8MB。较大的事务可以借助于更大的日志缓冲区来避免在事务完成之前将日志缓冲区的数据写入日志文件，以减少I/O操作进而提升系统性能。因此，在有着较大事务的应用场景中，建议为此变量设定一个更大的值。作用范围为全局级别，可用于选项文件，属非动态变量。
innodb_log_file_size={108576 .. 4294967295}
设定日志组中每个日志文件的大小，单位是字节，默认值是5MB。较为明智的取值范围是从1MB到缓存池体积的1/n，其中n表示日志组中日志文件的个数。日志文件越大，在缓存池中需要执行的检查点刷写操作就越少，这意味着所需的I/O操作也就越少，然而这也会导致较慢的故障恢复速度。作用范围为全局级别，可用于选项文件，属非动态变量。
innodb_log_files_in_group={2 .. 100}
设定日志组中日志文件的个数。InnoDB以循环的方式使用这些日志文件。默认值为2。作用范围为全局级别，可用于选项文件，属非动态变量。
innodb_log_group_home_dir=/PATH/TO/DIR
设定InnoDB重做日志文件的存储目录。在缺省使用InnoDB日志相关的所有变量时，其默认会在数据目录中创建两个大小为5MB的名为ib_logfile0和ib_logfile1的日志文件。作用范围为全局级别，可用于选项文件，属非动态变量。
innodb_max_dirty_pages_pct={0 .. 99}
设定InnoDB的缓冲池中脏页比例的上限，默认为75。当缓存池中的脏页比例接近或达到此变量定义的比值时，InnoDB的主线程会将刷写部分脏页中的数据至对应的文件中。作用范围为全局级别，可用于选项文件，属动态变量。
innodb_max_purge_lag={0 .. 4294967295}
InnoDB事务系统会维持一个有索引记录被添加了删除标记(delete-marked)的事务的列表，此列表的长度即为清写延迟(purge_lag)。此变量用于设定当发生清写延迟时，其队列长度达到多大时开始延迟INSERT、UPDATE或DELETE操作。当puge_lag超过innodb_max_purge_lag时，将延迟这些操作((purge_lag/innodb_max_purge_lag)*10)-5毫秒。默认值为0，表示从不延迟这些操作。需要进行操作延迟与否是在purge操作刚开始时计算的，并且每隔10秒钟会重新计算一次。基于历史地原因，purge操作无法启动时是不会有任何操作延迟的情况发生。作用范围为全局级别，可用于选项文件，属动态变量。
innodb_mirrored_log_groups=#
设定日志组镜像的个数。其值应该为1。
innodb_old_blocks_pct={5 .. 95}
InnoDB以“列表”结构管理缓存池，并使用修改版的LRU算法对其执行维护操作。当需要空间以保存新块(new block)时，InnoDB会清理最近最少使用的块并将新块加入到列表中。“中点插入策略(midpoint insertion policy)”将整个列表看作两个子列表：在列表首部是那些最近被访问过的新块(new/young block)子列表，尾部是那些最近较少被访问到的旧块(lod block)子列表。而LRU算法和中点插入策略用于保证将最近经常被访问到的块置于新块子列表，InnoDB新读入的块将被置于旧块子列表的前头，并根据需要将旧块子列表中的块移除。而某个被再次访问到的旧块则会被移至新块子列表的首部。表扫描操作可能会一次性地向缓存池中读入大量的数据块并可能导致一大批旧块被移出。
此变量正是用于设置被视作旧块子列表的长度所占据整个列表长度的比例，默认值是37，即缓存池的八分之三。作用范围为全局级别，可用于选项文件，属动态变量。
innodb_old_blocks_time=#
用于设定缓冲池中旧块子列表中的某旧块在其第一次又被访问到时，其至少需要在旧块子列表中再呆上多长时间(单位为毫秒)才会被转移至新块子列表。默认值是0，表示立即转移至新块子列表，哪怕其刚刚被转移至旧块子列表。而非零值则明确定义旧块列表中的块在其第一次被访问到时至少需要在旧块子列表中等待转移的时长。此变量通常要结合innodb_old_blocks_pct使用。作用范围为全局级别，可用于选项文件，属动态变量。
innodb_open_files=#
设定MySQL可同时打开的.ibd表空间文件的数量上限。此变量仅在使用多表空间文件时生效，其最小值为10，默认值为300。此变量的限定仅应用于InnoDB表的.ibd文件，跟MySQL服务器选项–open-files-limit没有关系，更不会影响表缓存的操作。作用范围为全局级别，可用于选项文件，属非动态变量。
innodb_purge_batch_size={1 .. 5000}
清写(purge)是指将缓存池中的脏页同步至持久性存储设备中的操作，以重做日志的记录为单位。此变量则用于定义清写操作的粒度，即多少个重做日志记录组合起来可以触发一次清写操作，默认值为20。此变量通常用于跟innodb_purge_threads=1一起对进行性能调优，但一般场景中都不需要修改它。作用范围为全局级别，可用于选项文件，属非动态变量。
innodb_purge_threads={0|1}
设定InnoDB执行清写操作的线程数量。默认值为0，表示清写操作由InnoDB的主线程自己完成，这可以降低内部资源竞争发生的概率，进而增强MySQL服务伸缩能力。不过，随着InnoDB内部各式各样的竞争越来越多，这种设置带来的性能优势已几乎不值一提。作用范围为全局级别，可用于选项文件，属动态变量。
innodb_read_ahead_threshold={0 .. 64}
设定InnoDB预读页面至缓冲池时的线性预读敏感度，也即InnoDB的读操作至少从一个盘区(extent，包含64个页面)中读取多个页面时才会为读取整个盘区中后续的页面初始化一个异步读操作。默认值为56。作用范围为全局级别，可用于选项文件，属动态变量。
innodb_read_io_threads={1 .. 64}
设定InnoDB为读操作启动的I/O线程数量，默认为4个。作用范围为全局级别，可用于选项文件，属非动态变量。
innodb_replication_delay={0 .. 4294967295}
设定在从服务器(slave)上运行的线程数达到innodb_thread_concurrency变量定义的并发上限时复制线程需要延迟的时长。默认为0，表示不延迟。作用范围为全局级别，可用于选项文件，属动态变量。
innodb_rollback_on_timeout={ON|OFF}
设定事务执行过程超时时事务回滚的方式。在MySQL 5.5中，默认为OFF，表示仅回滚事务中的最后一个语句。如果设置为ON，则表示中止事务执行并回滚整个事务。作用范围为全局级别，可用于选项文件，属非动态变量。
innodb_rollback_segments={1 .. 128}
设定InnoDB在系统表空间中为每个事务使用多少个回滚段(rollback segment)，默认为128个。如果较少的回滚段可以提升系统性能，则应该降低此变量的值。作用范围为全局级别，可用于选项文件，属动态变量。
innodb_spin_wait_delay={0 .. 4294967295}
自旋(spin)是一种通过不间断地测试来查看一个资源是否变为可用状态的等待操作，用于仅需要等待很短的时间等待所需资源的场景。使用自旋这种“空闲循环(busy-loop)”来完成资源等待的方式要比通过上下文切换使线程转入睡眠状态的方式要高效得多。但如果自旋了一个很短的时间后其依然无法获取资源，则仍然会转入前述第二种资源等待方式。此变量则正是用于定义InnoDB自旋操作的空闲循环转数，默认为6转。作用范围为全局级别，可用于选项文件，属动态变量。
innodb_stats_method={nulls_equal|nulls_unequal|null_ignored}
设定MySQL为InnoDB表收集分布的索引值的统计数据时如何处理NULL类型的数据。其可接受的值有三个，null_equals意指将所有的NULL值视为相同，并为之创建一个值组(value group)以保存NULL类值的个数；nulls_unequal意指将所有的NULL值视为不同，并为每个NULL单独创建一个大小为1的值组；nulls_ignored表示所有的NULL值都被忽略。这些用于生成表统计数据的方法会影响到优化器为执行查询如何选择选择索引。作用范围为全局级别，可用于选项文件，属动态变量。
innodb_stats_on_metadata={OFF|ON}
设定使用SHOW TABLE STATUS或者SHOW INDEX这两个元数据语句时，或访问INFORMATION_SCHEMA中的TABLES或STATISTICS表时，InnoDB是否更新统计数据。默认为更新。禁用此功能可以加速访问有着大量的表或索引的数据库，也可能提升InnoDB表上查询操作执行计划(execution plan)的稳定性。作用范围为全局级别，可用于选项文件，属动态变量。
innodb_strict_mod={ON|OFF}
为防止无视SQL语句书写或语法中的错误或无视操作模式与SQL语句各种组合中的无心之过，InnoDB提供了所谓的严格模式。严格模式中，前述的问题一旦出现将会导致InnoDB产生一个错误，而非警告和一系列特定的处理操作。此参数则正是用于定义是否启用InnoDB的严格模式，默认为OFF。
innodb_support_xa={TRUE|FLASE}
存储引擎事务在存储引擎内部被赋予了ACID属性，分布式(XA)事务是一种高层次的事务，它利用“准备”然后“提交”(prepare-then-commit)两段式的方式将ACID属性扩展到存储引擎外部，甚至是数据库外部。然而，“准备”阶段会导致额外的磁盘刷写操作。XA需要事务协调员，它会通知所有的参与者准备提交事务(阶段1)。当协调员从所有参与者那里收到“就绪”信息时，它会指示所有参与者进行真正的“提交”操作。
此变量正是用于定义InnoDB是否支持两段式提交的分布式事务，默认为启用。事实上，所有启用了二进制日志的并支持多个线程同时向二进制日志写入数据的MySQL服务器都需要启用分布式事务，否则，多个线程对二进制日志的写入操作可能会以与原始次序不同的方式完成，这将会在基于二进制日志的恢复操作中或者是从服务器上创建出不同原始数据的结果。因此，除了仅有一个线程可以改变数据以外的其它应用场景都不应该禁用此功能。而在仅有一个线程可以修改数据的应用中，禁用此功能是安全的并可以提升InnoDB表的性能。作用范围为全局和会话级别，可用于选项文件，属动态变量。
innodb_sync_spin_loops={0 .. 4294967295}
设定一个线程在等待InnoDB释放某个互斥量(mutex)之前自旋的转数，当自旋操作达到这个转数但互斥量仍未被释放时此线程将被挂起。默认值为30。作用范围为全局级别，可用于选项文件，属动态变量。
innodb_table_locks={ON|OFF}
InnoDB在存储引擎级别支持行级锁，而MySQL在服务器级别还支持使用表级锁。此变量则正是用来定义InnoDB是否在其内部支持使用MySQL表级锁。默认值为1或ON，表示如果autocommit变量的值为0(即禁止自动提交)，在InnoDB表上显式使用LOCK TABLES语句将使得InnoDB在存储引擎内部锁定此表。使用0或OFF值，则意味着显式使用LOCKS TABLE…WRITE语句不会在存储引擎级别产生影响，但对其它显式使用的LOCK TABLES…WRITE或LOCK TABLES…READ语句依然会有影响。作用范围为全局和会话级别，可用于选项文件，属动态变量。
innodb_thread_concurrency={0…1000}
设定InnoDB可在其内部并发运行的操作系统线程数量上限。多出的线程将被放置于FIFO队列进行等待，且不被计入并发运行线程数量。对于不用的应用场景来说，其理想的取值取决于硬件环境和工作负载，一般推荐为CPU个数的2倍加上磁盘的个数。默认值为0，表示无上限(不检查并发数量)，这意味着InnoDB可以按需要使用任意数量的并发线程，并会禁用SHOW ENGINE INNODB STATUS中的queries inside InnoDB和queries in queue counters两个计数器。作用范围为全局级别，可用于选项文件，属动态变量。
innodb_thread_sleep_delay=#
设定InnoDB线程在加入InnoDB队列之前的睡眠时长，单位是毫秒，默认值为10000。0值表示禁止睡眠而直接加入队列。作用范围为全局级别，可用于选项文件，属动态变量。
innodb_use_native_aio={ON|OFF}
设定InnoDB是否使用Linux的异步I/O子系统，因此，其仅应用于Linux系统平台，且MySQL启动后不能更改其值。InnoDB默认会启用此功能，而InnoDB如果因为Linux的异步I/O子系统的问题而无法正常启动，可以在选项文件中将此变量设置为OFF并重新启动之。事实上，就算变量值为ON，如果MySQL服务启动探测到了潜在的问题如联合的临时目录路径、tmpfs文件系统以及Linux内核不支持在tmpfs上使用AIO机制时也会自动关闭此变量。作用范围为全局级别，可用于选项文件，属非动态变量。
innodb_use_sys_malloc={ON|OFF}
设定InnoDB使用操作系统的(ON)还是自有的(OFF)内存分配器。默认值为ON。
innodb_version=STRING
InnoDB存储引擎的版本号，只读变量。
innodb_write_io_threads={1 .. 64}
设定InnoDB用于完成写操作的I/O线程数量，默认为4个。
log-slave-updates
用来做多级复制，让从服务器同时做主服务器
附：InnoDB的数据字典指的是跟踪InnoDB相关的对象如表、索引或表中的字段等的元数据信息，这些元数据存储在InnoDB的系统表空间中(system tablespace)。历史地原因，它跟.frm文件中的某些数据有重叠的地方。



------------

在Apache, PHP, MySQL的体系架构中，MySQL对于性能的影响最大，也是关键的核心部分。对于Discuz!论坛程序也是如此，MySQL的设置是否合理优化，直接影响到论坛的速度和承载量！同时，MySQL也是优化难度最大的一个部分，不但需要理解一些MySQL专业知识，同时还需要长时间的观察统计并且根据经验进行判断，然后设置合理的参数。 下面我们了解一下MySQL优化的一些基础，MySQL的优化我分为两个部分，一是服务器物理硬件的优化，二是MySQL自身(my.cnf)的优化。

一、服务器硬件对MySQL性能的影响

①磁盘寻道能力（磁盘I/O）,以目前高转速SCSI硬盘(7200转/秒)为例，这种硬盘理论上每秒寻道7200次，这是物理特性决定的，没有办法改变。MySQL每秒钟都在进行大量、复杂的查询操作，对磁盘的读写量可想而知。所以，通常认为磁盘I/O是制约MySQL性能的最大因素之一，对于日均访问量在100万PV以上的Discuz!论坛，由于磁盘I/O的制约，MySQL的性能会非常低下！解决这一制约因素可以考虑以下几种解决方案： 使用RAID-0+1磁盘阵列，注意不要尝试使用RAID-5，MySQL在RAID-5磁盘阵列上的效率不会像你期待的那样快。

②CPU 对于MySQL应用，推荐使用S.M.P.架构的多路对称CPU，例如：可以使用两颗Intel Xeon 3.6GHz的CPU，现在我较推荐用4U的服务器来专门做数据库服务器，不仅仅是针对于mysql。

③物理内存对于一台使用MySQL的Database Server来说，服务器内存建议不要小于2GB，推荐使用4GB以上的物理内存，不过内存对于现在的服务器而言可以说是一个可以忽略的问题，工作中遇到了高端服务器基本上内存都超过了16G。

二、MySQL自身因素当解决了上述服务器硬件制约因素后，让我们看看MySQL自身的优化是如何操作的。 对MySQL自身的优化主要是对其配置文件my.cnf中的各项参数进行优化调整。下面我们介绍一些对性能影响较大的参数。 由于my.cnf文件的优化设置是与服务器硬件配置息息相关的， 因而我们指定一个假想的服务器硬件环境：CPU: 2颗Intel Xeon 2.4GHz 内存: 4GB DDR 硬盘: SCSI 73GB(很常见的2U服务器 ) 。

下面，我们根据以上硬件配置结合一份已经优化好的my.cnf进行说明：



```
[mysqld] 
port = 3306 
serverid = 1 
socket = /tmp/mysql.sock 
skip-locking 
#避免MySQL的外部锁定，减少出错几率增强稳定性。 
skip-name-resolve 
#禁止MySQL对外部连接进行DNS解析，使用这一选项可以消除MySQL进行DNS解析的时间。但需要注意，如果开启该选项，则所有远程主机连接授权都要使用IP地址方式，否则MySQL将无法正常处理连接请求！
back_log = 384 
#back_log参数的值指出在MySQL暂时停止响应新请求之前的短时间内多少个请求可以被存在堆栈中。 如果系统在一个短时间内有很多连接，则需要增大该参数的值，该参数值指定到来的TCP/IP连接的侦听队列的大小。不同的操作系统在这个队列大小上有它自己的限制。 试图设定back_log高于你的操作系统的限制将是无效的。默认值为50。对于Linux系统推荐设置为小于512的整数。
key_buffer_size = 256M 
#key_buffer_size指定用于索引的缓冲区大小，增加它可得到更好的索引处理性能。对于内存在4GB左右的服务器该参数可设置为256M或384M。注意：该参数值设置的过大反而会是服务器整体效率降低！
max_allowed_packet = 4M 
thread_stack = 256K 
table_cache = 128K 
sort_buffer_size = 6M 
#查询排序时所能使用的缓冲区大小。注意：该参数对应的分配内存是每连接独占，如果有100个连接，那么实际分配的总共排序缓冲区大小为100 × 6 ＝ 600MB。所以，对于内存在4GB左右的服务器推荐设置为6-8M。
read_buffer_size = 4M 
#读查询操作所能使用的缓冲区大小。和sort_buffer_size一样，该参数对应的分配内存也是每连接独享。
join_buffer_size = 8M 
#联合查询操作所能使用的缓冲区大小，和sort_buffer_size一样，该参数对应的分配内存也是每连接独享。
myisam_sort_buffer_size = 64M 
table_cache = 512 
thread_cache_size = 64 
query_cache_size = 64M 
#指定MySQL查询缓冲区的大小。可以通过在MySQL控制台观察，如果Qcache_lowmem_prunes的值非常大，则表明经常出现缓冲不够的情况；如果Qcache_hits的值非常大，则表明查询缓冲使用非常频繁，如果该值较小反而会影响效率，那么可以考虑不用查询缓冲；Qcache_free_blocks，如果该值非常大，则表明缓冲区中碎片很多。
tmp_table_size = 256M 
max_connections = 768 
#指定MySQL允许的最大连接进程数。如果在访问论坛时经常出现Too Many Connections的错误提 示，则需要增大该参数值。
max_connect_errors = 10000000 
wait_timeout = 10 
#指定一个请求的最大连接时间，对于4GB左右内存的服务器可以设置为5-10。 
thread_concurrency = 8 
#该参数取值为服务器逻辑CPU数量*2，在本例中，服务器有2颗物理CPU，而每颗物理CPU又支持H.T超线程，所以实际取值为4*2=8
skip-networking 
#开启该选项可以彻底关闭MySQL的TCP/IP连接方式，如果WEB服务器是以远程连接的方式访问MySQL数据库服务器则不要开启该选项！否则将无法正常连接！
table_cache=1024 
#物理内存越大,设置就越大.默认为2402,调到512-1024最佳 
innodb_additional_mem_pool_size=4M 
#默认为2M 
innodb_flush_log_at_trx_commit=1 
#设置为0就是等到innodb_log_buffer_size列队满后再统一储存,默认为1 
innodb_log_buffer_size=2M 
#默认为1M 
innodb_thread_concurrency=8 
#你的服务器CPU有几个就设置为几,建议用默认一般为8 
key_buffer_size=256M 
#默认为218，调到128最佳 
tmp_table_size=64M 
#默认为16M，调到64-256最挂 
read_buffer_size=4M 
#默认为64K 
read_rnd_buffer_size=16M 
#默认为256K 
sort_buffer_size=32M 
#默认为256K 
thread_cache_size=120 
#默认为60 
query_cache_size=32M
```



很多情况需要具体情况具体分析

一、如果Key_reads太大，则应该把my.cnf中Key_buffer_size变大，保持Key_reads/Key_read_requests至少1/100以上，越小越好。

二、如果Qcache_lowmem_prunes很大，就要增加Query_cache_size的值。