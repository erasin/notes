# mysql 用户管理

```
一、root用户密码的维护：

由于安装MySQL完后,MySQL会自动提供一个不带密码的root用户，为了安全起见给root设置密码：

#mysqladmin -u root password 123 (123为密码，也可以写成:'123'或"123") ；

设置密码后登入时就不能直接输入mysql了，必须跟些参数了,如下：

[root@localhost ~]# mysql -u root -p （-u 后跟登入的用户名，-p 提示要密码登入）

Enter password:(输入密码)

修改密码：

[root@localhost ~] #mysqladmin -u root  -p  password  123456 (password 后跟的是要更新的新密码)

Enter password:(输入原始密码，回车即可)

二、其他用户的增加和删除：

以root用户登入，在mysql中有一张mysql.user表是存储MySQL中所有用户的信息表，所以可以直接增加删除这个表的记录就可增加和删除用户；

1.添加用户（有两种形式）：

A.mysql> grant all on *.* to micxp@"%" identified by "123" ;

mysql>flush privileges; (刷新系统权限表)

(执行完会在mysql.user表插入一条记录，all表示所有权限(包括增 删 改 查等权限)， *.* 表示所有数据库，micxp为添加的用户名，123为密码,%为匹配的所有主机，上面的信息都可以指定如grant select,update on db.* to micxp@localhost identified by '123";)

B.直接对mysql.user添加一条记录

mysql> insert into mysql.user(Host,User,Password) values("localhost","micxp",password("123"));

mysql>flush privileges;

这样就创建了一个名为：micxp密码为：123 (密码是经过加密的 ) 的用户，不过这样没有权限因为只添加了三个字段，也可通过grant添  加权限：

mysql>grant all  on *.* to micxp@localhost identified by '123";

mysql>flush privileges;(刷新系统权限表)
添加完用户 如果要远程登入MySQL,必须跟上主机Ip 如下：

[root@localhost ~]# mysql -u micxp -p -h 192.168.59.123

Enter password:(输入密码)

2.删除用户 ：

mysql>delete from mysql.user where user ='micxp' ;

mysql>flush privileges; (刷新系统权限表)

这样只删除了user表中的信息，建议用drop删除如下

mysql>drop user 'micxp'@'localhost';

```

文2


```
MySQL是一个真正的多用户、多线程SQL数据库服务器。MySQL是以一个客户机/服务器结构的实现，它由一个服务器守护程序mysqld和很多不同的客户程序和库组成。由于其源码的开放性及稳定性，且与网站流行编徎语言PHP的完美结合，现在很多站点都利用其当作后端数据库，使其获得了广泛应用。处于安全方面的考虑，需要为每一用户赋于对不同数据库的访问限制，以满足不同用户的要求。
Unix平台的MySQL默认只允许本机才能连接数据库。但是缺省root用户口令是空，所以当务之急是给root用户加上口令。
一、MySQL修改密码方法总结
首先要说明一点的是：一般情况下，修改MySQL密码是需要有mysql里的root权限的，这样一般用户是无法更改密码的，除非请求管理员帮助修改。
方法一
使用phpMyAdmin (图形化管理MySql数据库的工具)，这是最简单的，直接用SQL语句修改mysql数据库库的user表，不过别忘了使用PASSWORD函数，插入用户用Insert命令，修改用户用Update命令，删除用Delete命令。在本节后面有数据表user字段的详细介绍。
方法二
使用mysqladmin。输入
mysqladmin -u root -p111111 password newpasswd

注：111111为原mysql密码，后面的'password newpasswd'是连接数据库成功后要执行的命令。如果p和111111之间存在空格，将会报111111不是执行命令的错误。
执行这个命令后，需要输入root的原密码，这样root的密码将改为newpasswd。同样，把命令里的root改为你的用户名，你就可以改你自己的密码了。 当然如果你的mysqladmin连接不上mysql server，或者你没有办法执行mysqladmin，那么这种方法就是无效的，而且mysqladmin无法把密码清空。
下面的方法都在mysql提示符下使用，且必须有mysql的root权限：
mysql的信息数据存在mysql数据库中。因此，可以直接对此库进行操作。
方法三
mysql> INSERT INTO mysql.user (Host,User,Password) VALUES (\'%\',\'system\', PASSWORD(\'manager\'));
mysql> FLUSH PRIVILEGES
确切地说这是在增加一个用户，用户名为system，密码为manager。注意要使用PASSWORD函数，然后还要使用FLUSH PRIVILEGES来执行确认。
注："%"指这个system用户可以在任意主机上登陆。
方法四
和方法三一样，只是使用了REPLACE语句
mysql> REPLACE INTO mysql.user (Host,User,Password)
VALUES(\'%\',\'system\',PASSWORD(\'manager\'));
mysql> FLUSH PRIVILEGES

方法五
使用SET PASSWORD语句
mysql> SET PASSWORD FOR system@\"%\" = PASSWORD(\'manager\');
你也必须使用PASSWORD()函数，但是不需要使用FLUSH PRIVILEGES来执行确认。

方法六
使用GRANT ... IDENTIFIED BY语句，来进行授权。
mysql> GRANT USAGE ON *.* TO system@\"%\" IDENTIFIED BY \'manager\';
这里PASSWORD()函数是不必要的，也不需要使用FLUSH PRIVILEGES来执行确认。
注：PASSWORD()函数作用是为口令字加密，在程序中MySql自动解释。
flush privileges的意思是强制刷新内存授权表，否则用的还是缓冲中的口令，这时非法用户还可以用root用户及空口令登陆，直到重启MySQL服务器。
二、MySql中访问限制的设置方法
我们采用两种方法来设置用户。
进入到Mysql执行目录下（通常是c:\mysql\bin）。输入mysqld-shareware.exe，输入mysql --user=root mysql ,不然不能添加新用户。进入到mysql>提示符下进行操作。

假设我们要建立一个超级用户，用户名为system，用户口令为manager。
方法一
用Grant 命令授权，输入的代码如下：
mysql>GRANT ALL PRIVILEGES ON *.* TO system@localhost IDENTIFIED BY \'manager\' WITH GRANT OPTION;
应显示:Query OK, 0 rows affected (0.38 sec)

方法二
对用户的每一项权限进行设置：
mysql>INSERT INTO user VALUES(\'localhost\',\'system\',PASSWORD(\'manager\'), \'Y\',\'Y\',\'Y\',\'Y\',\'Y\',\'Y\',\'Y\',\'Y\',\'Y\',\'Y\',\'Y\',\'Y\',\'Y\',\'Y\');
对于3.22.34版本的MySQL,这里共14个\"Y\",其相应的权限如下(按字段顺序排列):
　
权限	表列名称	相应解释	使用范围
select	Select_priv	只有在真正从一个表中检索时才需要select权限	表
insert	Insert_priv	允许您把新行插入到一个存在的表中	表
update	Update_priv	允许你用新值更新现存表中行的列 	表
delete	Delete_priv	允许你删除满足条件的行	表
create	Create_priv	允许你创建新的数据库和表	数据库、表或索引
drop	Drop_priv	抛弃(删除)现存的数据库和表	数据库或表
reload	Reload_priv	允许您告诉服务器再读入授权表	服务器管理
shutdown	Shutdown_priv	可能被滥用(通过终止服务器拒绝为其他用户服务)	服务器管理
process	Process_priv	允许您察看当前执行的查询的普通文本,包括设定或改变口令查询	服务器管理
file	File_priv	权限可以被滥用在服务器上读取任何可读的文件到数据库表	服务器上的文件存取
grant	Grant_priv	允许你把你自己拥有的那些权限授给其他的用户	数据库或表
references	References_priv	允许你打开和关闭记录文件	数据库或表
index	Index_priv	允许你创建或抛弃(删除)索引	表
alter	Alter_priv	允许您改变表格,可以用于通过重新命名表来推翻权限系统	表
如果创建用户时只有select、insert、update和delete权限,则允许用户只能在一个数据库现有的表上实施操作.
下面就可以创建我们要用到的数据库了,我们直接输入. 例如：我们要创建数据库名为XinXiKu，可用如下代码：
mysql>create database XinXiKu;
应显示:Query OK, 1 row affected (0.00 sec)
mysql> desc user;
+-----------------+-----------------+------+-----+---------+-------+
| Field           | Type            | Null | Key | Default | Extra |
+-----------------+-----------------+------+-----+---------+-------+
| Host            | char(60) binary |      | PRI |         |       |
| User            | char(16) binary |      | PRI |         |       |
| Password        | char(16) binary |      |     |         |       |
| Select_priv     | enum('N','Y')   |      |     | N       |       |
| Insert_priv     | enum('N','Y')   |      |     | N       |       |
| Update_priv     | enum('N','Y')   |      |     | N       |       |
| Delete_priv     | enum('N','Y')   |      |     | N       |       |
| Create_priv     | enum('N','Y')   |      |     | N       |       |
| Drop_priv       | enum('N','Y')   |      |     | N       |       |
| Reload_priv     | enum('N','Y')   |      |     | N       |       |
| Shutdown_priv   | enum('N','Y')   |      |     | N       |       |
| Process_priv    | enum('N','Y')   |      |     | N       |       |
| File_priv       | enum('N','Y')   |      |     | N       |       |
| Grant_priv      | enum('N','Y')   |      |     | N       |       |
| References_priv | enum('N','Y')   |      |     | N       |       |
| Index_priv      | enum('N','Y')   |      |     | N       |       |
| Alter_priv      | enum('N','Y')   |      |     | N       |       |
+-----------------+-----------------+------+-----+---------+-------+
17 rows in set (0.01 sec)
演示了通过远程链接数据库：
一、赋予权限
mysql> use mysql
Database changed
mysql> select User,Password,Host from user;
+-----------+------------------+-----------+
| User      | Password         | Host      |
+-----------+------------------+-----------+
| root      | 5fcc735428e45938 | localhost |
|           |                  | localhost | //匿名本机可以登陆，对test
| cactiuser | 5fcc735428e45938 | localhost |  //数据库有权限
| root      |                  | %         |
| cactiuser |                  | %         |
+-----------+------------------+-----------+
5 rows in set (0.01 sec)
mysql> grant all on *.* to root identified by 'hujianping';
Query OK, 0 rows affected (0.00 sec)
mysql> select User,Password,Host from user;
+-----------+------------------+-----------+
| User      | Password         | Host      |
+-----------+------------------+-----------+
| root      | 5fcc735428e45938 | localhost |
|           |                  | localhost |
| cactiuser | 5fcc735428e45938 | localhost |
| root      | 2b4780fb21b27c92 | %         |
| cactiuser |                  | %         | //这个用户很危险地，密码   +-----------+------------------+-----------+ //空，并且可以远程登陆
5 rows in set (0.00 sec)
mysql>
二、在另一台机子上远程链接数据库
E:\bin>mysql -uroot -pxxxxxx -hx.x.x.x
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 6 to server version: 4.0.26-nt
Type 'help;' or '\h' for help. Type '\c' to clear the buffer.
mysql> show database;
三、用空的cactiuser用户远程登陆   (很危险啊）
E:\bin>mysql -ucactiuser -p -hx.x.x.x
Enter password:
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 8 to server version: 4.0.26-nt
Type 'help;' or '\h' for help. Type '\c' to clear the buffer.
mysql> show databases;
+----------+
| Database |
+----------+
| cacti    |
| test     |
+----------+
2 rows in set (0.00 sec)
E:\bin>mysql --user=cactiuser --password="" --host="x.x.x.x"
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 11 to server version: 4.0.26-nt
Type 'help;' or '\h' for help. Type '\c' to clear the buffer.
mysql>
MySQL的5个授权表：user, db, host, tableserials_priv和columnserials_priv提供非常灵活的安全机制，从MySQL 3.22.11开始引入了两条语句GRANT和REVOKE来创建和删除用户权限，可以方便的限制哪个用户可以连接服务器，从哪里连接以及连接后可以做什么操作。作为MySQL管理员，我们必须了解授权表的意义以及如何用GRANT和REVOKE来创建用户、授权和撤权、删除用户。 在3.22.11版本以前的MySQL授权机制不完善，和新版本也有较大的不同，建议升级到最新版本的MySQL。
grant 命令远程测试没有成功。
mysql> grant all on *.* to root@y.y.y.y identified by  'iwmusic';
ERROR 1045: Access denied for user: 'root@220.231.31.203' (Using password: YES)
本地测试是可以的。
```