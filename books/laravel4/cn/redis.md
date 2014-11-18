# Redis

- [简介](#introduction)
- [配置](#configuration)
- [使用方法](#usage)
- [流水线(Pipelining)](#pipelining)

<a name="Introduction"></a>
## 简介

[Redis](http://redis.io) 是一个开源的先进键值存储工具(其实是Nosql数据库的一种). 它常用来进行一些数据结构的存储,这些数据结构包括 [strings(字符串)](http://redis.io/topics/data-types#strings), [hashes(哈希表)](http://redis.io/topics/data-types#hashes), [lists(双向链表)](http://redis.io/topics/data-types#lists), [sets(无序集合)](http://redis.io/topics/data-types#sets), and [sorted sets(有序集合)](http://redis.io/topics/data-types#sorted-sets).

> **注意:** 如果你的Redis PHP扩展是使用PECL安装的,那你需要将 `app/config/app.php` 文件中的 Redis别名(alias)改成别的名字.

<a name="configuration"></a>
## 配置

Redis的配置信息存储在 **app/config/database.php** 文件内. 在这个文件里你可以看见一个 **redis** 数组,它就是你应用中使用Redis的相关配置:

	'redis' => array(

		'cluster' => true,

		'default' => array('host' => '127.0.0.1', 'port' => 6379),

	),

这个 `defaule`(默认) 的服务器配置可以在做开发的时候使用. 你也可以任意的按照你的需要和使用环境来进行配置. 你可以简单的给每个Redis服务器起一个名字并指定服务器(host)和端口(port). 

上面那个 `cluster`(集群) 选项是告诉Laravel框架Redis的客户端要访问多个节点,允许你在内存中建立一个记录这些节点的节点池. 然而,请注意客户端不会处理故障转移,也就是说缓存的数据可能会从另一个服务器上读取(这个不是特别理解,可能是说一旦一个节点上的服务器挂了,有可能配置了集群这个选项后就能在集群中其他服务器节点上读取到你要的缓存数据了).

如果你的 Redis服务器端使用了用户验证,你需要增加一个 `password` 键值对选项到配置的数组中来配置密码.

<a name="usage"></a>
## 使用方法

你可以用 `Redis::connection` 方法来获取一个Redis连接实例:

	$redis = Redis::connection();

这样做可以给你一个默认的Redis服务器连接实例.如果你需要其他服务器连接实例,那么你就要在连接函数中声明你具体需要使用的那个连接配置名称(这个名称就是前面你在配置里除了 default 外自定义的其他连接名称):

	$redis = Redis::connection('other');

一旦你有了一个Redis的客户端连接实例,你就可以发送你的 [Redis commands(Redis操作命令)](http://redis.io/commands) 来操作Redis数据库了.Laravel使用魔术方法来发送命令到Redis服务器(即你可以按照实际需要发送任何Redis能处理的命令过去):

	$redis->set('name', 'Taylor');

	$name = $redis->get('name');

	$values = $redis->lrange('names', 5, 10);

注意命令对应的参数可以简单的放在魔术方法调用的函数括号里面发送.当然,你可能不喜欢使用魔术方法来调用那些命令,你也可以通过 `command` 方法向服务器发送命令:

	$values = $redis->command('lrange', array(5, 10));

当你只是想简单执行一下default(默认)连接的命令时,你可以使用 `Redis` 类的静态魔术方法:

	Redis::set('name', 'Taylor');

	$name = Redis::get('name');

	$values = Redis::lrange('names', 5, 10);

> **注意:** Redis [cache(缓存)](/docs/cache) 和 [session](/docs/session) 驱动程序已经包含在Laravel框架内.

<a name="pipelining"></a>
## 流水线(Pipelining)

流水线模式(Pipelining)是当你需要一次向服务器发送多个命令的时候使用的方法.你需要使用 `pipeline` 命令来开头:

#### 流水线模式可以一次向Redis服务器发送多条指令

	Redis::pipeline(function($pipe)
	{
		for ($i = 0; $i < 1000; $i++)
		{
			$pipe->set("key:$i", $i);
		}
	});