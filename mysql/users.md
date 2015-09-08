# mysql users
./mysqladmin flush-hosts -uroot -p


create user 'root'@'115.238.95.194' identified by password password('boc2009!@#')



grant all privileges on `era_%`.* to 'wangweixin'@'%' with grant option;



'boc123456web';
boc2009!@#


grant all privileges on *.* to root@'115.238.95.194' identified by 'boc2009!@#';


--
create user 'wangweixin'@'%' identified by password 'boc123456web';
FLUSH PRIVILEGES;


grant select,insert,update,delete,create,drop on
era_bocms to 'wangweixin'@'115.238.95.194' identified by 'boc123456web';


grant all privileges on era*.* to wangweixin@% identified by 'boc123456web';

--
boc123456web

use mysql;

-- 查看当前用户
select User();

-- 创建用户指定IP访问
create user era@10.0.0.2 identified by password '123456';

-- 获取用户信息
select user,host,password from user where user='era';
1.       CREATE USER

语法：

CREATE USER 'username'@'host' IDENTIFIED BY 'password';

   例子: CREATE USER 'dog'@'localhost' IDENTIFIED BY '123456';

               CREATE USER 'pig'@'192.168.1.101_' IDENDIFIED BY '123456';

               CREATE USER 'pig'@'%' IDENTIFIED BY '123456';

               CREATE USER 'pig'@'%' IDENTIFIED BY '';

               CREATE USER 'pig'@'%';

     实例1：

       mysql> create user jss;

        这样创建的用户，可以从任意安装了mysql客户端，并能够访问目标服务器的机器上创建连接，无须密码.例如，从ip:10.0.0.99的客户端执行连接：

         mysql -ujss -h 172.16.1.110

        查看该用户：

         mysql> select user,host,password from user where user='jss';

                SELECT USER();    //显示当前用户

     实例2：

        mysql> create user jss_ps identified by 'jss';

       用户连接时，必须指定密码，那就可以在创建用户时，通过指定identified by子句来设定密码

       用密码登陆：

         mysql -ujss_ps -p -h 172.16.1.110

      如果希望指定的用户只能从某台指定的域(domain)或主机访问，可以在创建用户时指定host，例如，指定用户只能从10.0.0.99访问

mysql> create user jss_ip@10.0.0.99 identified by password '123456';



2.       使用GRANT语句

语法：mysql> grant 权限1,权限2,...权限n on 数据库名称.表名称 to 用户名@用户地址 identified by '连接口令';

权限1,权限2,...权限n代表

select,insert,update,delete,create,drop,index,alter,grant,references,reload,shutdown,process,file等14个权限

实例：

  mysql>grant select,insert,update,delete,create,drop on vtdc.employee to joe@10.163.225.87 identified by '123';

给来自10.163.225.87的用户joe分配可对数据库vtdc的employee表进行select,insert,update,delete,create,drop等操作的权限，并设定口令为123。

mysql>grant all privileges on vtdc.* to joe@10.163.225.87 identified by '123';

给来自10.163.225.87的用户joe分配可对数据库vtdc所有表进行所有操作的权限，并设定口令为123。

mysql>grant all privileges on *.* to joe@10.163.225.87 identified by '123';

给来自10.163.225.87的用户joe分配可对所有数据库的所有表进行所有操作的权限，并设定口令为123。

mysql>grant all privileges on *.* to joe@localhost identified by '123';

给本机用户joe分配可对所有数据库的所有表进行所有操作的权限，并设定口令为123。

3.       直接向mysql.user表插入记录:

mysql> insert into user (host,user,password) values ('%','jss_insert',password('jss'));

mysql>flush privileges;   //刷新系统权限表

4.       修改mysql用户密码方式：

a.       使用mysqladmin语法：mysqladmin -u用户名 -p旧密码 password 新密码

例如：mysqladmin -u root -p 123 password 456；

b.       直接修改user表的用户口令：

语法：update mysql.user set password=password('新密码') where User="phplamp" and Host="localhost";

实例：update user set password=password('54netseek') where user='root';

      flush privileges;

c.       使用SET PASSWORD语句修改密码：语法：

SET PASSWORD FOR 'username'@'host' = PASSWORD('newpassword');

如果是当前登陆用户用SET PASSWORD = PASSWORD("newpassword");

实例：

set password for root@localhost=password('');

SET PASSWORD FOR name=PASSWORD('new password');

SET PASSWORD FOR 'pig'@'%' = PASSWORD("123456");

5.        删除用户和撤销权限：

a.       取消一个账户和其权限

Drop USER user;

drop user username@'%'

drop user username@localhost

b.       取消授权用户:

语法：REVOKE privilege ON databasename.tablename FROM 'username'@'host';

例子: REVOKE SELECT ON *.* FROM 'pig'@'%';

  REVOKE SELECT ON test.user FROM 'pig'@'%';

  revoke all on *.* from sss@localhost ;

  revoke all on user.* from 'admin'@'%';

      SHOW GRANTS FOR 'pig'@'%';     //查看授权

c.       删除用户：

语法: Delete from user where user = "user_name" and host = "host_name" ;

例子：delete from user where user='sss' and host='localhost';



二、数据库表

1.查看所有数据库： 数据库目录：/usr/local/mysql/data

   mysql> SHOW DATABASES;   //显示数据库

   mysql> USE abccs         //进入数据库

   mysql> SHOW TABLES;      //显示表

   mysql> DESCRIBE mytable; //显示表结构

   mysql> CREATE DATABASE abccs;    //创建一个数据库

   mysql> CREATE TABLE mytable (name VARCHAR(20), sex CHAR(1), birth DATE, birthaddr VARCHAR(20));   //创建表

   mysql> insert into mytable values (‘abccs’,‘f’,‘1977-07-07’,‘china’);                     //插入表数据

   使用文本方式插入数据：

    {

      mysql.txt内容：abccs f 1977-07-07 china 　

                     mary f 1978-12-12 usa

                     tom m 1970-09-02 usa

      mysql> LOAD DATA LOCAL INFILE "mytable.txt" INTO TABLE pet;    //导入TXT文件数据

     }



2.删除数据库：

  mysql> drop database drop_database;   //删除一个已经确定存在的数据库

         alter table 表名 ENGINE=存储引擎名；  //修改表的存储引擎

         alter table 表名 drop 属性名； //删除字段

         alter table 旧表名 rename to 新表名；  //修改表名

         alter table 表名 modify 属性名 数据类型；  //修改字段数据类型

         alter table 表名 change 旧属性名 新属性名 新数据类型； //修改字段名

         alter table 表名 drop FOREING KEY 外键别名； //删除子表外键约束

         增加表字段：

         { alter table example add phone VACGAR(20); //增加无约束的字段

           alter table example add age INT(4) NOT NULL; //增加万增约束的字段

           alter table example add num INT(8) PRIMARY KEY FIRST;  //表的第一个位置增加字段

           alter table example add address VARCHAR(30) NOT NULL AFTER phone;  //表的指定位置之后增加字段

           alter table example modify name VARCHAR(20) FIRST; //把字段修改到第一位

           alter table example modify num INT(8) ATER phone；//把字段修改到指定字段之后

         }
