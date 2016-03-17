
如今，开发人员不断地开发和部署使用 LAMP（Linux®、Apache、MySQL 和 PHP/Perl）架构的应用程序。但是，服务器管理员常常对应用程序本身没有什么控制能力，因为 应用程序是别人编写的。这份 共三部分的系列文章 将讨论许多服务器配置问题，这些配置会影响应用程序的性能。本文是本系列文章的第三部分，也是最后一部分，将重点讨论为实现最高效率而对数据库层进行的调优。

关于 MySQL 调优

有 3 种方法可以加快 MySQL 服务器的运行速度，效率从低到高依次为：

   1. 替换有问题的硬件。
   2. 对 MySQL 进程的设置进行调优。
   3. 对查询进行优化。

替换有问题的硬件通常是我们的第一考虑，主要原因是数据库会占用大量资源。不过这种解决方案也就仅限于此了。实际上，您通常可以让中央处理器（CPU）或 磁盘速度加倍，也可以让内存增大 4 到 8 倍。

第二种方法是对 MySQL 服务器（也称为 mysqld）进行调优。对这个进程进行调优意味着适当地分配内存，并让 mysqld 了解将会承受何种类型的负载。加快磁盘运行速度不如减少所需的磁盘访问次数。类似地，确保 MySQL 进程正确操作就意味着它花费在服务查询上的时间要多于花费在处理后台任务（如处理临时磁盘表或打开和关闭文件）上的时间。对 mysqld 进行调优是本文的重点。

最好的方法是确保查询已经进行了优化。这意味着对表应用了适当的索引，查询是按照可以充分利用 MySQL 功能的方式来编写的。尽管本文并没有包含查询调优方面的内容（很多著作中已经针对这个主题进行了探讨），不过它会配置 mysqld 来报告可能需要进行调优的查询。

虽然已经为这些任务指派了次序，但是仍然要注意硬件和 mysqld 的设置以利于适当地调优查询。机器速度慢也就罢了，我曾经见过速度很快的机器在运行设计良好的查询时由于负载过重而失败，因为 mysqld 被大量繁忙的工作所占用而不能服务查询。

记录慢速查询

在一个 SQL 服务器中，数据表都是保存在磁盘上的。索引为服务器提供了一种在表中查找特定数据行的方法，而不用搜索整个表。当必须要搜索整个表时，就称为表扫描。通常 来说，您可能只希望获得表中数据的一个子集，因此全表扫描会浪费大量的磁盘 I/O，因此也就会浪费大量时间。当必须对数据进行连接时，这个问题就更加复杂了，因为必须要对连接两端的多行数据进行比较。

当然，表扫描并不总是会带来问题；有时读取整个表反而会比从中挑选出一部分数据更加有效（服务器进程中查询规划器用来作出这些决定）。如果索引的使用效率 很低，或者根本就不能使用索引，则会减慢查询速度，而且随着服务器上的负载和表大小的增加，这个问题会变得更加显著。执行时间超过给定时间范围的查询就称 为慢速查询。

您可以配置 mysqld 将这些慢速查询记录到适当命名的慢速查询日志中。管理员然后会查看这个日志来帮助他们确定应用程序中有哪些部分需要进一步调查。清单 1 给出了要启用慢速查询日志需要在 my.cnf 中所做的配置。

清单 1. 启用 MySQL 慢速查询日志
[mysqld]
; enable the slow query log, default 10 seconds
log-slow-queries
; log queries taking longer than 5 seconds
long_query_time = 5
; log queries that don't use indexes even if they take less than long_query_time
; MySQL 4.1 and newer only
log-queries-not-using-indexes
复制代码
这三个设置一起使用，可以记录执行时间超过 5 秒和没有使用索引的查询。请注意有关 log-queries-not-using-indexes 的警告：您必须使用 MySQL 4.1 或更高版本。慢速查询日志都保存在 MySQL 数据目录中，名为 hostname-slow.log。如果希望使用一个不同的名字或路径，可以在 my.cnf 中使用 log-slow-queries = /new/path/to/file 实现此目的。

阅读慢速查询日志最好是通过 mysqldumpslow 命令进行。指定日志文件的路径，就可以看到一个慢速查询的排序后的列表，并且还显示了它们在日志文件中出现的次数。一个非常有用的特性是 mysqldumpslow 在比较结果之前，会删除任何用户指定的数据，因此对同一个查询的不同调用被计为一次；这可以帮助找出需要工作量最多的查询。

对查询进行缓存

很多 LAMP 应用程序都严重依赖于数据库，但却会反复执行相同的查询。每次执行查询时，数据库都必须要执行相同的工作 —— 对查询进行分析，确定如何执行查询，从磁盘中加载信息，然后将结果返回给客户机。MySQL 有一个特性称为查询缓存，它将（后面会用到的）查询结果保存在内存中。在很多情况下，这会极大地提高性能。不过，问题是查询缓存在默认情况下是禁用的。

将 query_cache_size = 32M 添加到 /etc/my.conf 中可以启用 32MB 的查询缓存。

监视查询缓存

在启用查询缓存之后，重要的是要理解它是否得到了有效的使用。MySQL 有几个可以查看的变量，可以用来了解缓存中的情况。清单 2 给出了缓存的状态。

清单 2. 显示查询缓存的统计信息
mysql> SHOW STATUS LIKE 'qcache%';
+-------------------------+------------+
| Variable_name           | Value      |
+-------------------------+------------+
| Qcache_free_blocks      | 5216       |
| Qcache_free_memory      | 14640664   |
| Qcache_hits             | 2581646882 |
| Qcache_inserts          | 360210964  |
| Qcache_lowmem_prunes    | 281680433  |
| Qcache_not_cached       | 79740667   |
| Qcache_queries_in_cache | 16927      |
| Qcache_total_blocks     | 47042      |
+-------------------------+------------+
8 rows in set (0.00 sec)
复制代码
这些项的解释如表 1 所示。
变 量名	说明
                        Qcache_free_blocks                    	缓存中相邻内存块的个数。数目大说明可能有碎片。 FLUSH QUERY CACHE 会对缓存中的碎片进行整理，从而得到一个空闲块。
                        Qcache_free_memory                    	缓存中的空闲内存。
                        Qcache_hits                    	每次查询在缓存中命 中时就增大。
                        Qcache_inserts                    	每次插入一个查询时就增大。命中次数除以插入次数就是不中比率；用 1 减去这个值就是命中率。在上面这个例子中，大约有 87% 的查询都在缓存中命中。
                        Qcache_lowmem_prunes                    	缓存出现内存不足并且必须要 进行清理以便为更多查询提供空间的次数。这个数字最好长时间来看；如果这个数字在不断增长，就表示可能碎片非常严重，或者内存很少。（上面的 free_blocks 和            free_memory 可以告诉您属于哪种情况）。
                        Qcache_not_cached                    	不适合进行缓 存的查询的数量，通常是由于这些查询不是 SELECT 语句。
                        Qcache_queries_in_cache                    	当前缓存的查询（和响应）的数量。
                        Qcache_total_blocks                    	缓 存中块的数量。


通常，间隔几秒显示这些变量就可以看出区别，这可以帮助确定缓存是否正在有效地使用。运行 FLUSH STATUS 可以重置一些计数器，如果服务器已经运行了一段时间，这会非常有帮助。

使用非常大的查询缓存，期望可以缓存所有东西，这种想法非常诱人。由于 mysqld 必须要对缓存进行维护，例如当内存变得很低时执行剪除，因此服务器可能会在试图管理缓存时而陷入困境。作为一条规则，如果 FLUSH QUERY CACHE 占用了很长时间，那就说明缓存太大了。

强制限制

您可以在 mysqld 中强制一些限制来确保系统负载不会导致资源耗尽的情况出现。清单 3 给出了 my.cnf 中与资源有关的一些重要设置。

清单 3. MySQL 资源设置
max_connections=500
wait_timeout=10
max_connect_errors = 100
复制代码
连接最大个数是在第一行中进行管理的。与 Apache 中的 MaxClients 类似，其想法是确保只建立服务允许数目的连接。要确定服务器上目前建立过的最大连接数，请执行 SHOW STATUS LIKE 'max_used_connections'。

第 2 行告诉 mysqld 终止所有空闲时间超过 10 秒的连接。在 LAMP 应用程序中，连接数据库的时间通常就是 Web 服务器处理请求所花费的时间。有时候，如果负载过重，连接会挂起，并且会占用连接表空间。如果有多个交互用户或使用了到数据库的持久连接，那么将这个值设 低一点并不可取！

最后一行是一个安全的方法。如果一个主机在连接到服务器时有问题，并重试很多次后放弃，那么这个主机就会被锁定，直到 FLUSH HOSTS 之后才能运行。默认情况下，10 次失败就足以导致锁定了。将这个值修改为 100 会给服务器足够的时间来从问题中恢复。如果重试 100 次都无法建立连接，那么使用再高的值也不会有太多帮助，可能它根本就无法连接。

缓冲区和缓存

MySQL 支持超过 100 个的可调节设置；但是幸运的是，掌握少数几个就可以满足大部分需要。查找这些设置的正确值可以通过 SHOW STATUS 命令查看状态变量，从中可以确定 mysqld 的运作情况是否符合我们的预期。给缓冲区和缓存分配的内存不能超过系统中的现有内存，因此调优通常都需要进行一些妥协。

MySQL 可调节设置可以应用于整个 mysqld 进程，也可以应用于单个客户机会话。

服务器端的设置

每个表都可以表示为磁盘上的一个文件，必须先打开，后读取。为了加快从文件中读取数据的过程，mysqld 对这些打开文件进行了缓存，其最大数目由 /etc/mysqld.conf 中的 table_cache 指定。清单 4 给出了显示与打开表有关的活动的方式。

清单 4. 显示打开表的活动
mysql> SHOW STATUS LIKE 'open%tables';
+---------------+-------+
| Variable_name | Value |
+---------------+-------+
| Open_tables   | 5000  |
| Opened_tables | 195   |
+---------------+-------+
2 rows in set (0.00 sec)
复制代码
清单 4 说明目前有 5,000 个表是打开的，有 195 个表需要打开，因为现在缓存中已经没有可用文件描述符了（由于统计信息在前面已经清除了，因此可能会存在 5,000 个打开表中只有 195 个打开记录的情况）。如果 Opened_tables 随着重新运行 SHOW STATUS 命令快速增加，就说明缓存命中率不够。如果 Open_tables 比 table_cache 设置小很多，就说明该值太大了（不过有空间可以增长总不是什么坏事）。例如，使用 table_cache = 5000 可以调整表的缓存。

与表的缓存类似，对于线程来说也有一个缓存。 mysqld 在接收连接时会根据需要生成线程。在一个连接变化很快的繁忙服务器上，对线程进行缓存便于以后使用可以加快最初的连接。

清单 5 显示如何确定是否缓存了足够的线程。

清单 5. 显示线程使用统计信息
mysql> SHOW STATUS LIKE 'threads%';
+-------------------+--------+
| Variable_name     | Value  |
+-------------------+--------+
| Threads_cached    | 27     |
| Threads_connected | 15     |
| Threads_created   | 838610 |
| Threads_running   | 3      |
+-------------------+--------+
4 rows in set (0.00 sec)
复制代码
此处重要的值是 Threads_created，每次 mysqld 需要创建一个新线程时，这个值都会增加。如果这个数字在连续执行 SHOW STATUS 命令时快速增加，就应该尝试增大线程缓存。例如，可以在 my.cnf 中使用 thread_cache = 40 来实现此目的。

关键字缓冲区保存了 MyISAM 表的索引块。理想情况下，对于这些块的请求应该来自于内存，而不是来自于磁盘。清单 6 显示了如何确定有多少块是从磁盘中读取的，以及有多少块是从内存中读取的。

清单 6. 确定关键字效率
mysql> show status like '%key_read%';
+-------------------+-----------+
| Variable_name     | Value     |
+-------------------+-----------+
| Key_read_requests | 163554268 |
| Key_reads         | 98247     |
+-------------------+-----------+
2 rows in set (0.00 sec)
复制代码
Key_reads 代表命中磁盘的请求个数， Key_read_requests 是总数。命中磁盘的读请求数除以读请求总数就是不中比率 —— 在本例中每 1,000 个请求，大约有 0.6 个没有命中内存。如果每 1,000 个请求中命中磁盘的数目超过 1 个，就应该考虑增大关键字缓冲区了。例如，key_buffer = 384M 会将缓冲区设置为 384MB。

临时表可以在更高级的查询中使用，其中数据在进一步进行处理（例如 GROUP BY 字句）之前，都必须先保存到临时表中；理想情况下，在内存中创建临时表。但是如果临时表变得太大，就需要写入磁盘中。清单 7 给出了与临时表创建有关的统计信息。

清单 7. 确定临时表的使用
mysql> SHOW STATUS LIKE 'created_tmp%';
+-------------------------+-------+
| Variable_name           | Value |
+-------------------------+-------+
| Created_tmp_disk_tables | 30660 |
| Created_tmp_files       | 2     |
| Created_tmp_tables      | 32912 |
+-------------------------+-------+
3 rows in set (0.00 sec)
复制代码
每次使用临时表都会增大 Created_tmp_tables；基于磁盘的表也会增大 Created_tmp_disk_tables。对于这个比率，并没有什么严格的规则，因为这依赖于所涉及的查询。长时间观察 Created_tmp_disk_tables 会显示所创建的磁盘表的比率，您可以确定设置的效率。 tmp_table_size 和 max_heap_table_size 都可以控制临时表的最大大小，因此请确保在 my.cnf 中对这两个值都进行了设置。

每个会话的设置

下面这些设置针对于每个会话。在设置这些数字时要十分谨慎，因为它们在乘以可能存在的连接数时候，这些选项表示大量的内存！您可以通过代码修改会话中的这 些数字，或者在 my.cnf 中为所有会话修改这些设置。

当 MySQL 必须要进行排序时，就会在从磁盘上读取数据时分配一个排序缓冲区来存放这些数据行。如果要排序的数据太大，那么数据就必须保存到磁盘上的临时文件中，并再 次进行排序。如果 sort_merge_passes 状态变量很大，这就指示了磁盘的活动情况。清单 8 给出了一些与排序相关的状态计数器信息。

清单 8. 显示排序统计信息
mysql> SHOW STATUS LIKE "sort%";
+-------------------+---------+
| Variable_name     | Value   |
+-------------------+---------+
| Sort_merge_passes | 1       |
| Sort_range        | 79192   |
| Sort_rows         | 2066532 |
| Sort_scan         | 44006   |
+-------------------+---------+
4 rows in set (0.00 sec)
复制代码
如果 sort_merge_passes 很大，就表示需要注意 sort_buffer_size。例如， sort_buffer_size = 4M 将排序缓冲区设置为 4MB。

MySQL 也会分配一些内存来读取表。理想情况下，索引提供了足够多的信息，可以只读入所需要的行，但是有时候查询（设计不佳或数据本性使然）需要读取表中大量数 据。要理解这种行为，需要知道运行了多少个 SELECT 语句，以及需要读取表中的下一行数据的次数（而不是通过索引直接访问）。实现这种功能的命令如清单 9 所示。

清单 9. 确定表扫描比率
mysql> SHOW STATUS LIKE "com_select";
+---------------+--------+
| Variable_name | Value  |
+---------------+--------+
| Com_select    | 318243 |
+---------------+--------+
1 row in set (0.00 sec)

mysql> SHOW STATUS LIKE "handler_read_rnd_next";
+-----------------------+-----------+
| Variable_name         | Value     |
+-----------------------+-----------+
| Handler_read_rnd_next | 165959471 |
+-----------------------+-----------+
1 row in set (0.00 sec)
复制代码
Handler_read_rnd_next / Com_select 得出了表扫描比率 —— 在本例中是 521:1。如果该值超过 4000，就应该查看 read_buffer_size，例如 read_buffer_size = 4M。如果这个数字超过了 8M，就应该与开发人员讨论一下对这些查询进行调优了！

3 个必不可少的工具

尽管在了解具体设置时，SHOW STATUS 命令会非常有用，但是您还需要一些工具来解释 mysqld 所提供的大量数据。我发现有 3 个工具是必不可少的；在 参考资料 一节中您可以找到相应的链接。

大部分系统管理员都非常熟悉 top 命令，它为任务所消耗的 CPU 和内存提供了一个不断更新的视图。 mytop 对 top 进行了仿真；它为所有连接上的客户机以及它们正在运行的查询提供了一个视图。mytop 还提供了一个有关关键字缓冲区和查询缓存效率的实时数据和历史数据，以及有关正在运行的查询的统计信息。这是一个很有用的工具，可以查看系统中（比如 10 秒钟之内）的状况，您可以获得有关服务器健康信息的视图，并显示导致问题的任何连接。

mysqlard 是一个连接到 MySQL 服务器上的守护程序，负责每 5 分钟搜集一次数据，并将它们存储到后台的一个 Round Robin Database 中。有一个 Web 页面会显示这些数据，例如表缓存的使用情况、关键字效率、连接上的客户机以及临时表的使用情况。尽管 mytop 提供了服务器健康信息的快照，但是 mysqlard 则提供了长期的健康信息。作为奖励，mysqlard 使用自己搜集到的一些信息针对如何对服务器进行调优给出一些建议。

搜集 SHOW STATUS 信息的另外一个工具是 mysqlreport。其报告要远比 mysqlard 更加复杂，因为需要对服务器的每个方面都进行分析。这是对服务器进行调优的一个非常好的工具，因为它对状态变量进行适当计算来帮助确定需要修正哪些问题。


目前是用作一台纯数据库服务器，单机多实例设置主从，然后hugepages为32G

现在不知道该如何去优化这个my.cnf，晚上查到一个最大内存的使用量查询方法，结果出来的结果是5G都不到，

贴上mysqld1的配置，求优化

skip-external-locking

key_buffer_size = 1024M

max_allowed_packet = 32M

table_open_cache = 512

sort_buffer_size = 2M

read_buffer_size = 2M

read_rnd_buffer_size = 8M

myisam_sort_buffer_size = 128M

thread_cache_size = 8

query_cache_size = 64M

thread_concurrency = 24

skip-name-resolve

skip_slave_start = 1

log-bin = /path/binlogs/mysqld-bin

log-bin-index =  /path/binlogs/mysqld-bin.index

server-id = 1

binlog_format=mixed

expire-logs-days = 21


===>

skip-external-locking

key_buffer_size = 20480M

max_allowed_packet = 32M

table_open_cache = 2048

sort_buffer_size = 256M

read_buffer_size = 256M

read_rnd_buffer_size = 128M

myisam_sort_buffer_size = 512M

thread_cache_size = 32

query_cache_size = 2048M

thread_concurrency = 24

max_connections = 4096

tmp_table_size = 64M

skip-name-resolve

skip_slave_start = 1

log-bin = /path/binlogs/mysqld-bin

log-bin-index =  /path/binlogs/mysqld-bin.index

server-id = 1

binlog_format=mixed

expire-logs-days = 21



[client]
port = 3306
socket = /tmp/mysql.sock
[mysqld]
port = 3306
socket = /tmp/mysql.sock
basedir = /usr/local/mysql
datadir = /data/mysql
pid-file = /data/mysql/mysql.pid
user = mysql
bind-address = 0.0.0.0
server-id = 1 #表示是本机的序号为1,一般来讲就是master的意思
skip-name-resolve
# 禁止MySQL对外部连接进行DNS解析，使用这一选项可以消除MySQL进行DNS解析的时间。但需要注意，如果开启该选项，
# 则所有远程主机连接授权都要使用IP地址方式，否则MySQL将无法正常处理连接请求
#skip-networking
back_log = 600
# MySQL能有的连接数量。当主要MySQL线程在一个很短时间内得到非常多的连接请求，这就起作用，
# 然后主线程花些时间(尽管很短)检查连接并且启动一个新线程。back_log值指出在MySQL暂时停止回答新请求之前的短时间内多少个请求可以被存在堆栈中。
# 如果期望在一个短时间内有很多连接，你需要增加它。也就是说，如果MySQL的连接数据达到max_connections时，新来的请求将会被存在堆栈中，
# 以等待某一连接释放资源，该堆栈的数量即back_log，如果等待连接的数量超过back_log，将不被授予连接资源。
# 另外，这值（back_log）限于您的操作系统对到来的TCP/IP连接的侦听队列的大小。
# 你的操作系统在这个队列大小上有它自己的限制（可以检查你的OS文档找出这个变量的最大值），试图设定back_log高于你的操作系统的限制将是无效的。
max_connections = 1000
# MySQL的最大连接数，如果服务器的并发连接请求量比较大，建议调高此值，以增加并行连接数量，当然这建立在机器能支撑的情况下，因为如果连接数越多，介于MySQL会为每个连接提供连接缓冲区，就会开销越多的内存，所以要适当调整该值，不能盲目提高设值。可以过'conn%'通配符查看当前状态的连接数量，以定夺该值的大小。
max_connect_errors = 6000
# 对于同一主机，如果有超出该参数值个数的中断错误连接，则该主机将被禁止连接。如需对该主机进行解禁，执行：FLUSH HOST。
open_files_limit = 65535
# MySQL打开的文件描述符限制，默认最小1024;当open_files_limit没有被配置的时候，比较max_connections*5和ulimit -n的值，哪个大用哪个，
# 当open_file_limit被配置的时候，比较open_files_limit和max_connections*5的值，哪个大用哪个。
table_open_cache = 128
# MySQL每打开一个表，都会读入一些数据到table_open_cache缓存中，当MySQL在这个缓存中找不到相应信息时，才会去磁盘上读取。默认值64
# 假定系统有200个并发连接，则需将此参数设置为200*N(N为每个连接所需的文件描述符数目)；
# 当把table_open_cache设置为很大时，如果系统处理不了那么多文件描述符，那么就会出现客户端失效，连接不上
max_allowed_packet = 4M
# 接受的数据包大小；增加该变量的值十分安全，这是因为仅当需要时才会分配额外内存。例如，仅当你发出长查询或MySQLd必须返回大的结果行时MySQLd才会分配更多内存。
# 该变量之所以取较小默认值是一种预防措施，以捕获客户端和服务器之间的错误信息包，并确保不会因偶然使用大的信息包而导致内存溢出。
binlog_cache_size = 1M
# 一个事务，在没有提交的时候，产生的日志，记录到Cache中；等到事务提交需要提交的时候，则把日志持久化到磁盘。默认binlog_cache_size大小32K
max_heap_table_size = 8M
# 定义了用户可以创建的内存表(memory table)的大小。这个值用来计算内存表的最大行数值。这个变量支持动态改变
tmp_table_size = 16M
# MySQL的heap（堆积）表缓冲大小。所有联合在一个DML指令内完成，并且大多数联合甚至可以不用临时表即可以完成。
# 大多数临时表是基于内存的(HEAP)表。具有大的记录长度的临时表 (所有列的长度的和)或包含BLOB列的表存储在硬盘上。
# 如果某个内部heap（堆积）表大小超过tmp_table_size，MySQL可以根据需要自动将内存中的heap表改为基于硬盘的MyISAM表。还可以通过设置tmp_table_size选项来增加临时表的大小。也就是说，如果调高该值，MySQL同时将增加heap表的大小，可达到提高联接查询速度的效果
read_buffer_size = 2M
# MySQL读入缓冲区大小。对表进行顺序扫描的请求将分配一个读入缓冲区，MySQL会为它分配一段内存缓冲区。read_buffer_size变量控制这一缓冲区的大小。
# 如果对表的顺序扫描请求非常频繁，并且你认为频繁扫描进行得太慢，可以通过增加该变量值以及内存缓冲区大小提高其性能
read_rnd_buffer_size = 8M
# MySQL的随机读缓冲区大小。当按任意顺序读取行时(例如，按照排序顺序)，将分配一个随机读缓存区。进行排序查询时，
# MySQL会首先扫描一遍该缓冲，以避免磁盘搜索，提高查询速度，如果需要排序大量数据，可适当调高该值。但MySQL会为每个客户连接发放该缓冲空间，所以应尽量适当设置该值，以避免内存开销过大
sort_buffer_size = 8M
# MySQL执行排序使用的缓冲大小。如果想要增加ORDER BY的速度，首先看是否可以让MySQL使用索引而不是额外的排序阶段。
# 如果不能，可以尝试增加sort_buffer_size变量的大小
join_buffer_size = 8M
# 联合查询操作所能使用的缓冲区大小，和sort_buffer_size一样，该参数对应的分配内存也是每连接独享
thread_cache_size = 8
# 这个值（默认8）表示可以重新利用保存在缓存中线程的数量，当断开连接时如果缓存中还有空间，那么客户端的线程将被放到缓存中，
# 如果线程重新被请求，那么请求将从缓存中读取,如果缓存中是空的或者是新的请求，那么这个线程将被重新创建,如果有很多新的线程，
# 增加这个值可以改善系统性能.通过比较Connections和Threads_created状态的变量，可以看到这个变量的作用。(–>表示要调整的值)
# 根据物理内存设置规则如下：
# 1G  —> 8
# 2G  —> 16
# 3G  —> 32
# 大于3G  —> 64
query_cache_size = 8M
#MySQL的查询缓冲大小（从4.0.1开始，MySQL提供了查询缓冲机制）使用查询缓冲，MySQL将SELECT语句和查询结果存放在缓冲区中，
# 今后对于同样的SELECT语句（区分大小写），将直接从缓冲区中读取结果。根据MySQL用户手册，使用查询缓冲最多可以达到238%的效率。
# 通过检查状态值'Qcache_%'，可以知道query_cache_size设置是否合理：如果Qcache_lowmem_prunes的值非常大，则表明经常出现缓冲不够的情况，
# 如果Qcache_hits的值也非常大，则表明查询缓冲使用非常频繁，此时需要增加缓冲大小；如果Qcache_hits的值不大，则表明你的查询重复率很低，
# 这种情况下使用查询缓冲反而会影响效率，那么可以考虑不用查询缓冲。此外，在SELECT语句中加入SQL_NO_CACHE可以明确表示不使用查询缓冲
query_cache_limit = 2M
#指定单个查询能够使用的缓冲区大小，默认1M
key_buffer_size = 4M
#指定用于索引的缓冲区大小，增加它可得到更好处理的索引(对所有读和多重写)，到你能负担得起那样多。如果你使它太大，
# 系统将开始换页并且真的变慢了。对于内存在4GB左右的服务器该参数可设置为384M或512M。通过检查状态值Key_read_requests和Key_reads，
# 可以知道key_buffer_size设置是否合理。比例key_reads/key_read_requests应该尽可能的低，
# 至少是1:100，1:1000更好(上述状态值可以使用SHOW STATUS LIKE 'key_read%'获得)。注意：该参数值设置的过大反而会是服务器整体效率降低
ft_min_word_len = 4
# 分词词汇最小长度，默认4
transaction_isolation = REPEATABLE-READ
# MySQL支持4种事务隔离级别，他们分别是：
# READ-UNCOMMITTED, READ-COMMITTED, REPEATABLE-READ, SERIALIZABLE.
# 如没有指定，MySQL默认采用的是REPEATABLE-READ，ORACLE默认的是READ-COMMITTED
log_bin = mysql-bin
binlog_format = mixed
expire_logs_days = 30 #超过30天的binlog删除
log_error = /data/mysql/mysql-error.log #错误日志路径
slow_query_log = 1
long_query_time = 1 #慢查询时间 超过1秒则为慢查询
slow_query_log_file = /data/mysql/mysql-slow.log
performance_schema = 0
explicit_defaults_for_timestamp
#lower_case_table_names = 1 #不区分大小写
skip-external-locking #MySQL选项以避免外部锁定。该选项默认开启
default-storage-engine = InnoDB #默认存储引擎
innodb_file_per_table = 1
# InnoDB为独立表空间模式，每个数据库的每个表都会生成一个数据空间
# 独立表空间优点：
# 1．每个表都有自已独立的表空间。
# 2．每个表的数据和索引都会存在自已的表空间中。
# 3．可以实现单表在不同的数据库中移动。
# 4．空间可以回收（除drop table操作处，表空不能自已回收）
# 缺点：
# 单表增加过大，如超过100G
# 结论：
# 共享表空间在Insert操作上少有优势。其它都没独立表空间表现好。当启用独立表空间时，请合理调整：innodb_open_files
innodb_open_files = 500
# 限制Innodb能打开的表的数据，如果库里的表特别多的情况，请增加这个。这个值默认是300
innodb_buffer_pool_size = 64M
# InnoDB使用一个缓冲池来保存索引和原始数据, 不像MyISAM.
# 这里你设置越大,你在存取表里面数据时所需要的磁盘I/O越少.
# 在一个独立使用的数据库服务器上,你可以设置这个变量到服务器物理内存大小的80%
# 不要设置过大,否则,由于物理内存的竞争可能导致操作系统的换页颠簸.
# 注意在32位系统上你每个进程可能被限制在 2-3.5G 用户层面内存限制,
# 所以不要设置的太高.
innodb_write_io_threads = 4
innodb_read_io_threads = 4
# innodb使用后台线程处理数据页上的读写 I/O(输入输出)请求,根据你的 CPU 核数来更改,默认是4
# 注:这两个参数不支持动态改变,需要把该参数加入到my.cnf里，修改完后重启MySQL服务,允许值的范围从 1-64
innodb_thread_concurrency = 0
# 默认设置为 0,表示不限制并发数，这里推荐设置为0，更好去发挥CPU多核处理能力，提高并发量
innodb_purge_threads = 1
# InnoDB中的清除操作是一类定期回收无用数据的操作。在之前的几个版本中，清除操作是主线程的一部分，这意味着运行时它可能会堵塞其它的数据库操作。
# 从MySQL5.5.X版本开始，该操作运行于独立的线程中,并支持更多的并发数。用户可通过设置innodb_purge_threads配置参数来选择清除操作是否使用单
# 独线程,默认情况下参数设置为0(不使用单独线程),设置为 1 时表示使用单独的清除线程。建议为1
innodb_flush_log_at_trx_commit = 2
# 0：如果innodb_flush_log_at_trx_commit的值为0,log buffer每秒就会被刷写日志文件到磁盘，提交事务的时候不做任何操作（执行是由mysql的master thread线程来执行的。
# 主线程中每秒会将重做日志缓冲写入磁盘的重做日志文件(REDO LOG)中。不论事务是否已经提交）默认的日志文件是ib_logfile0,ib_logfile1
# 1：当设为默认值1的时候，每次提交事务的时候，都会将log buffer刷写到日志。
# 2：如果设为2,每次提交事务都会写日志，但并不会执行刷的操作。每秒定时会刷到日志文件。要注意的是，并不能保证100%每秒一定都会刷到磁盘，这要取决于进程的调度。
# 每次事务提交的时候将数据写入事务日志，而这里的写入仅是调用了文件系统的写入操作，而文件系统是有 缓存的，所以这个写入并不能保证数据已经写入到物理磁盘
# 默认值1是为了保证完整的ACID。当然，你可以将这个配置项设为1以外的值来换取更高的性能，但是在系统崩溃的时候，你将会丢失1秒的数据。
# 设为0的话，mysqld进程崩溃的时候，就会丢失最后1秒的事务。设为2,只有在操作系统崩溃或者断电的时候才会丢失最后1秒的数据。InnoDB在做恢复的时候会忽略这个值。
# 总结
# 设为1当然是最安全的，但性能页是最差的（相对其他两个参数而言，但不是不能接受）。如果对数据一致性和完整性要求不高，完全可以设为2，如果只最求性能，例如高并发写的日志服务器，设为0来获得更高性能
innodb_log_buffer_size = 2M
# 此参数确定些日志文件所用的内存大小，以M为单位。缓冲区更大能提高性能，但意外的故障将会丢失数据。MySQL开发人员建议设置为1－8M之间
innodb_log_file_size = 32M
# 此参数确定数据日志文件的大小，更大的设置可以提高性能，但也会增加恢复故障数据库所需的时间
innodb_log_files_in_group = 3
# 为提高性能，MySQL可以以循环方式将日志文件写到多个文件。推荐设置为3
innodb_max_dirty_pages_pct = 90
# innodb主线程刷新缓存池中的数据，使脏数据比例小于90%
innodb_lock_wait_timeout = 120
# InnoDB事务在被回滚之前可以等待一个锁定的超时秒数。InnoDB在它自己的锁定表中自动检测事务死锁并且回滚事务。InnoDB用LOCK TABLES语句注意到锁定设置。默认值是50秒
bulk_insert_buffer_size = 8M
# 批量插入缓存大小， 这个参数是针对MyISAM存储引擎来说的。适用于在一次性插入100-1000+条记录时， 提高效率。默认值是8M。可以针对数据量的大小，翻倍增加。
myisam_sort_buffer_size = 8M
# MyISAM设置恢复表之时使用的缓冲区的尺寸，当在REPAIR TABLE或用CREATE INDEX创建索引或ALTER TABLE过程中排序 MyISAM索引分配的缓冲区
myisam_max_sort_file_size = 10G
# 如果临时文件会变得超过索引，不要使用快速排序索引方法来创建一个索引。注释：这个参数以字节的形式给出
myisam_repair_threads = 1
# 如果该值大于1，在Repair by sorting过程中并行创建MyISAM表索引(每个索引在自己的线程内)
interactive_timeout = 28800
# 服务器关闭交互式连接前等待活动的秒数。交互式客户端定义为在mysql_real_connect()中使用CLIENT_INTERACTIVE选项的客户端。默认值：28800秒（8小时）
wait_timeout = 28800
# 服务器关闭非交互连接之前等待活动的秒数。在线程启动时，根据全局wait_timeout值或全局interactive_timeout值初始化会话wait_timeout值，
# 取决于客户端类型(由mysql_real_connect()的连接选项CLIENT_INTERACTIVE定义)。参数默认值：28800秒（8小时）
# MySQL服务器所支持的最大连接数是有上限的，因为每个连接的建立都会消耗内存，因此我们希望客户端在连接到MySQL Server处理完相应的操作后，
# 应该断开连接并释放占用的内存。如果你的MySQL Server有大量的闲置连接，他们不仅会白白消耗内存，而且如果连接一直在累加而不断开，
# 最终肯定会达到MySQL Server的连接上限数，这会报'too many connections'的错误。对于wait_timeout的值设定，应该根据系统的运行情况来判断。
# 在系统运行一段时间后，可以通过show processlist命令查看当前系统的连接状态，如果发现有大量的sleep状态的连接进程，则说明该参数设置的过大，
# 可以进行适当的调整小些。要同时设置interactive_timeout和wait_timeout才会生效。
[mysqldump]
quick
max_allowed_packet = 16M #服务器发送和接受的最大包长度
[myisamchk]
key_buffer_size = 8M
sort_buffer_size = 8M
read_buffer = 4M
write_buffer = 4M
