# 缓存

- [配置](#configuration)
- [缓存用法](#cache-usage)
- [递增 & 递减](#increments-and-decrements)
- [缓存标签](#cache-tags)
- [数据库缓存](#database-cache)

<a name="configuration"></a>
## 配置

Laravel对不同的缓存系统提供了统一的API. 请在 `app/config/cache.php` 文件中配置缓存信息. 你可以在此文件中指定整个应用程序默认使用哪种缓存驱动. Laravel支持系统以外的流行后端缓存如 [Memcached](http://memcached.org) 和 [Redis](http://redis.io).

缓存配置文件还包含了其他配置项, 都已记录在文件中, 请详细阅读. 默认情况下, Laravel被配置为使用 'file' 缓存驱动, 它将缓存数据序列化的存储在文件系统中, 对于大型的应用, 我们建议你使用内存缓存, 如 Memcached 或 APC.

<a name="cache-usage"></a>
## 缓存用法

#### 将某一条数据存入缓存

	Cache::put('key', 'value', $minutes);

#### 使用Carbon对象并设置过期时间

	$expiresAt = Carbon::now()->addMinutes(10);

	Cache::put('key', 'value', $expiresAt);

#### 当一条数据不在缓存中才存储

	Cache::add('key', 'value', $minutes);

如果该数据实际上 **已经添加** 到缓存中, 那么 `add` 方法将返回 `true` . 否则, 此方法将返回 `false`.

#### 检查缓存中是否存在某个key对应的数据

	if (Cache::has('key'))
	{
		//
	}

#### 从缓存中取得一条key对应的数据

	$value = Cache::get('key');

#### 从缓存中取得数据, 如果不存在, 则返回指定的默认值

	$value = Cache::get('key', 'default');

	$value = Cache::get('key', function() { return 'default'; });

#### 在缓存中永久存储数据

	Cache::forever('key', 'value');

有时候你可能希望从缓存中取得数据, 并且数据不存在时还可以存储一个默认值, 这时可以使用 `Cache::remember` 方法:

	$value = Cache::remember('users', $minutes, function()
	{
		return DB::table('users')->get();
	});

你还可以结合使用 `remember` 和 `forever` 方法:

	$value = Cache::rememberForever('users', function()
	{
		return DB::table('users')->get();
	});

注意: 所有项目都是序列化的存储在缓存中, 所以你可以自由存储任何类型的数据.

#### 在缓存中拉出一条数据

如果你需要从缓存中先取出一条数据, 然后删除它, 你可以使用 `pull` 方法:

	$value = Cache::pull('key');

#### 从缓存中删除某条数据

	Cache::forget('key');

<a name="increments-and-decrements"></a>
## 递增 & 递减

除了 `file` 和 `database` , 其他驱动都支持 `increment` and `decrement` 操作:

#### 递增某一个值

	Cache::increment('key');

	Cache::increment('key', $amount);

#### 递减某一个值

	Cache::decrement('key');

	Cache::decrement('key', $amount);

<a name="cache-tags"></a>
## 缓存标签

> **注意:** 使用 `file` 或 `database` 缓存驱动时不支持缓存标签. 此外, 在使用多个缓存标签时它们将存储为 "forever", 使用一个如 `memcached` 的驱动性能将会是最好, 它会自动清除过时的记录.

#### 访问一个标记的缓存

缓存标签允许你在缓存中标记相关的项目, 刷新指定名称标记的所有缓存. 要访问一个标记的缓存, 可以使用 `tags` 方法.

你可以通过传递一个标记名称的有序列表作为参数, 或是一个标记名称的有序数数组, 来存储一个标记的缓存:

	Cache::tags('people', 'authors')->put('John', $john, $minutes);

	Cache::tags(array('people', 'artists'))->put('Anne', $anne, $minutes);

你可以在任何缓存存储方法中组合使用标签, 包括`remember`, `forever`, 和 `rememberForever`. 你也可以从标记的缓存中访问已缓存的项目, 以及使用其它的缓存方法, 如 `increment` 和 `decrement`.

#### 从标记的缓存中访问项目

通过与保存时所用相同的标签, 作为参数列表来访问一个标记的缓存.

	$anne = Cache::tags('people', 'artists')->get('Anne');

	$john = Cache::tags(array('people', 'authors'))->get('John');

你可以通过一个名称或名称的列表来刷新所有的项目. 例如, 下面的语句将移除标记有任何'people', 'authors', 或两者都有的所有缓存. 所以, 无论是“Anne”和“John”将从缓存中被移除:

	Cache::tags('people', 'authors')->flush();

相比之下, 下面的语句将移除标记中只包含 'authors' 的缓存, 因此 "John" 将被移除, 但不影响 "Anne".

	Cache::tags('authors')->flush();

<a name="database-cache"></a>
## 数据库缓存

当使用'数据库'缓存驱动时, 你将需要设置一个表来存储缓存数据. 以下是一个使用 `Schema` 声明的例子:

	Schema::create('cache', function($table)
	{
		$table->string('key')->unique();
		$table->text('value');
		$table->integer('expiration');
	});
