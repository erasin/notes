# 数据库使用基础

- [配置](#configuration)
- [读 / 写 连接](#read-write-connections)
- [运行查询语句](#running-queries)
- [事务](#database-transactions)
- [同时使用多个数据库系统](#accessing-connections)
- [查询日志](#query-logging)

<a name="configuration"></a>
## 配置

Laravel让连接和使用数据库变得异常简单.数据库配置文件是 `app/config/database.php`.你可以在配置文件中定义所有你的数据库连接,以及指定默认连接.这个文件中已经提供了所有支持的数据库系统连接例子.

目前Laravel支持四种数据库系统: MySQL, Postgres, SQLite, 和 SQL Server.

<a name="read-write-connections"></a>
## 读 / 写 连接 (读写分离配置)

你可能需要使用一个连接来处理 SELECT(读)数据操作,而另一个链接来处理 INSERT, UPDATE, 和 DELETE 这样的操作(即读写分离).Laravel是这项工作变得简单,它会在你直接发送查询语句、使用查询构建器或者Eloquent ORM时自动选择适当的连接(来实现读写分离).

以下是一个读/写 连接配置的例子:

	'mysql' => array(
		'read' => array(
			'host' => '192.168.1.1',
		),
		'write' => array(
			'host' => '196.168.1.2'
		),
		'driver'    => 'mysql',
		'database'  => 'database',
		'username'  => 'root',
		'password'  => '',
		'charset'   => 'utf8',
		'collation' => 'utf8_unicode_ci',
		'prefix'    => '',
	),

注意看两个增加的配置数组: `read` 和 `write`.这两个配置数组值都是一个键值对: `host`(用来标明连接的服务器).其余部分的配置由于 `read` 和 `write` 是相同的,因此他们都在 `mysql` 下面被合并到了一起.因此,我们只需要修改 `read` 和 `write` 配置中我们需要覆盖配置的值即可.在这个例子里, `192.168.1.1` 将会做为"读"连接 而 `192.168.1.2`将会做为"写"连接.数据库的用户名密码,表前缀,字符集和其他的配置项目会在 `mysql` 配置下为两个连接所共用.

<a name="running-queries"></a>
## 执行crud语句

一旦你已经配置好了数据库的连接,你就可以直接使用 `DB` 类来发送sql请求了.

#### 运行一个 Select 查询

	$results = DB::select('select * from users where id = ?', array(1));

这里的 `select` 方法将会一直返回一个 `array`(数组) 结果集.

#### 运行一个 Insert(插入) 请求

	DB::insert('insert into users (id, name) values (?, ?)', array(1, 'Dayle'));

#### 运行一个 Update(更新) 请求

	DB::update('update users set votes = 100 where name = ?', array('John'));

#### 运行一个 Delete(删除) 请求

	DB::delete('delete from users');

> **注意:** `update`(更新) 和 `delete`(删除) 将会返回操作的影响行数(affect rows).

#### 执行非crud操作

	DB::statement('drop table users');

#### 监听数据库操作事件

你可以使用 `DB::listen` 方法来监听query事件:

	DB::listen(function($sql, $bindings, $time)
	{
		//
	});

<a name="database-transactions"></a>
## 事务

你可以用 `transaction` 方法来调用一个事务集合:

	DB::transaction(function()
	{
		DB::table('users')->update(array('votes' => 1));

		DB::table('posts')->delete();
	});

> **注意:** 任何 `transaction` 抛出的异常都将导致事务自动回滚.

有时你需要自己手动开启一个事务:

	DB::beginTransaction();

你可以用过 `rollback` 方法手动回滚事务:

	DB::rollback();

最后,你可以使用 `commit` 方法来提交:

	DB::commit();

<a name="accessing-connections"></a>
## 同时使用多个数据库系统

你可能使用很多的数据库系统,你可以使用 `DB::connection` 方法来选择使用它们:

	$users = DB::connection('foo')->select(...);

你可能需要在数据库系统的层面上操作数据库，使用PDO实例即可:

	$pdo = DB::connection()->getPdo();

使用reconnect方法重新连接一个指定的数据库系统:

	DB::reconnect('foo');


你可以使用 `disconnect` 方法来手动断开数据库连接,防止PDO连接数超过 `max_connections` 的限制:

	DB::disconnect('foo');

<a name="query-logging"></a>
## 查询日志

Laravel默认会为当前请求执行的的所有查询生成日志并保存在内存中。 因此， 在某些特殊的情况下， 比如一次性向数据库中插入大量数据， 就可能导致内存不足。 在这种情况下，你可以通过 `disableQueryLog` 方法来关闭查询日志:

	DB::connection()->disableQueryLog();

调用 `getQueryLog` 方法可以同时获取多个查询执行后的日志:

       $queries = DB::getQueryLog();