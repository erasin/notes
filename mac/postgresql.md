# Mac上使用homebrew安装PostgreSql 2014.10.30

## 安装
brew 安装 postgresql ：

    $ brew install postgresql

查看安装的版本：

    $ pg_ctl -V
    pg_ctl (PostgreSQL) 9.3.5

安装成功之后，安装路径为：/usr/local/var/postgres

接下来，初始化数据库：

```
$ initdb /usr/local/var/postgres
The files belonging to this database system will be owned by user "june".
This user must also own the server process.

The database cluster will be initialized with locale "zh_CN.UTF-8".
The default database encoding has accordingly been set to "UTF8".
initdb: could not find suitable text search configuration for locale "zh_CN.UTF-8"
The default text search configuration will be set to "simple".

Data page checksums are disabled.

creating directory /usr/local/var/postgres ... ok
creating subdirectories ... ok
selecting default max_connections ... 100
selecting default shared_buffers ... 128MB
creating configuration files ... ok
creating template1 database in /usr/local/var/postgres/base/1 ... ok
initializing pg_authid ... ok
initializing dependencies ... ok
creating system views ... ok
loading system objects' descriptions ... ok
creating collations ... ok
creating conversions ... ok
creating dictionaries ... ok
setting privileges on built-in objects ... ok
creating information schema ... ok
loading PL/pgSQL server-side language ... ok
vacuuming database template1 ... ok
copying template1 to template0 ... ok
copying template1 to postgres ... ok
syncing data to disk ... ok

WARNING: enabling "trust" authentication for local connections
You can change this by editing pg_hba.conf or using the option -A, or
--auth-local and --auth-host, the next time you run initdb.

Success. You can now start the database server using:

    postgres -D /usr/local/var/postgres
or
    pg_ctl -D /usr/local/var/postgres -l logfile start
```

配置开机登陆（可选）：

    $ mkdir -p ~/Library/LaunchAgents
    $ cp /usr/local/Cellar/postgresql/9.3.5_1/homebrew.mxcl.postgresql.plist ~/Library/LaunchAgents/
    $ launchctl load -w ~/Library/LaunchAgents/homebrew.mxcl.postgresql.plist

手动启动 postgresql

    $ pg_ctl -D /usr/local/var/postgres -l /usr/local/var/postgres/server.log start

查看状态：

    pg_ctl -D /usr/local/var/postgres -l /usr/local/var/postgres/server.log status

停止：

    $ pg_ctl -D /usr/local/var/postgres -l /usr/local/var/postgres/server.log stop -s -m fast

查看进程：

```
$ ps auxwww | grep postgres
june            56126   0.0  0.0  2432772    644 s000  S+    5:01下午   0:00.00 grep postgres
june            56058   0.0  0.0  2467360    584   ??  Ss    5:00下午   0:00.00 postgres: stats collector process
june            56057   0.0  0.0  2611808   1744   ??  Ss    5:00下午   0:00.00 postgres: autovacuum launcher process
june            56056   0.0  0.0  2611676    696   ??  Ss    5:00下午   0:00.00 postgres: wal writer process
june            56055   0.0  0.0  2611676    944   ??  Ss    5:00下午   0:00.01 postgres: writer process
june            56054   0.0  0.0  2611676    756   ??  Ss    5:00下午   0:00.00 postgres: checkpointer process
june            56044   0.0  0.2  2611676  14096 s000  S     5:00下午   0:00.02 /usr/local/Cellar/postgresql/9.3.5_1/bin/postgres -D /usr/local/var/postgres
```

创建用户和数据库：

    #createuser will prompt you for a password, enter it twice.
    $ createuser  -P test
    $ createdb -Otest -Eutf8 test_db
    $ psql
    postgres=# GRANT ALL PRIVILEGES ON test TO test;
    postgres=# \q

进入命令行模式：

    $ psql -U test test_db -h localhost -W

如果出现 FATAL: Ident authentication failed for user，是因为：

    This is because by default PostgreSQL uses ‘ident’ authentication i.e it checks if the username exists on the system. You need to change authentication mode to ‘trust’ as we do not want to add a system user.
    Modify the settings in "pg_hba.conf" to use 'trust' authentication.

请修改 /usr/local/var/postgres/pg_hba.conf 为：

    host    all             all             127.0.0.1/32            trust
    # IPv6 local connections:
    host    all             all             ::1/128                 trust

安装 pgadmin，下载地址：http://www.pgadmin.org/download/macosx.php

卸载
$ brew uninstall postgres
如果配置了开机登陆：

$ launchctl unload -w ~/Library/LaunchAgents/homebrew.mxcl.postgresql.plist
$ rm -rf ~/Library/LaunchAgents/homebrew.mxcl.postgresql.plist
