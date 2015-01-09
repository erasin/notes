# Facades

- [介绍](#introduction)
- [说明](#explanation)
- [实际用例](#practical-usage)
- [创建 Facades](#creating-facades)
- [模拟 Facades](#mocking-facades)
- [Facade 类参考](#facade-class-reference)

<a name="introduction"></a>
## 介绍

Facades（一种设计模式，通常翻译为外观模式）提供了一个"static"（静态）接口去访问注册到[IoC 容器](4.2/ioc.md)中的类 。Laravel 含有很多 facades，你可能不知道其实你在某些地方已经使用过它们了。Laravel的"facades"就像"静态代理方法"，为IoC容器中的类服务，相比传统的静态方法，在保持了更强的可测试性和灵活性的同时，它提供更简洁，更富有表现力的语法。

有时候， 你可能想为你的应用程序或包创建自己的 facades， 所以，让我们来讨论一下如何开发和使用这些类。

> **注意：** 在深入 facades 之前，我们强烈建议你多了解一下 Laravel [IoC 容器](4.2/ioc.md)。

<a name="explanation"></a>
## 说明

在 Laravel 应用程序中， facade 是提供从容器中访问对象的类。`Facade` 类实现了该机制。Laravel的facades，和其他自定义的facades，都需要继承基类`Facade`

你的 facade 类只需要实现一个方法： `getFacadeAccesor` 。 `getFacadeAccessor` 方法的工作是定义如何从容器中取得对象。 Facades 基类构建了 `__callStatic()` 魔术方法来从 facade 延迟访问取得对象。

因此，当你使用facade调用，类似`Cache::get`，Laravel会从IoC容器取得Cache管理类并调用`get`方法。在技术层上说，Laravel Facades是一种便捷的语法使用Laravel IoC容器作为服务定位器.

<a name="practical-usage"></a>
## 实际用例

在以下示例中，调用Laravel缓存系统， 咋一看该代码，可能会认为`get`静态方法是在`Cache`类执行。

	$value = Cache::get('key');

然后，如果我们查看`Illuminate\Support\Facades\Cache`类， 你会发现该类没有任何`get`静态方法：

	class Cache extends Facade {

		/**
		 * Get the registered name of the component.
		 *
		 * @return string
		 */
		protected static function getFacadeAccessor() { return 'cache'; }

	}

`Cache`类继承`Facade`这个基类，并且定义了个`getFacadeAccessor()`方法。注意，该方法的工作是返回绑定到`IoC`的名字。

当用户引用任何在`Cache` facade 中的静态方法， Laravel 从 IoC 容器绑定中取得 `cache`，并且执行请求的对象方法（在该例子中为`get`）。

所以，我们 `Cache::get` 执行可以重写为：

	$value = $app->make('cache')->get('key');

<a name="creating-facades"></a>
## 创建Facades

要为自己的应用程序或者包创建一个facade是非常简单的。你只需要做三件事情：

- 一个 IoC 绑定。
- 一个 facade 类。
- 一个 facade 别名配置。

让我们来看个例子。这里，我们定义一个`PaymentGateway\Payment`类。

	namespace PaymentGateway;

	class Payment {

		public function process()
		{
			//
		}

	}

这个类可以放在`app/models`目录，或者其他任何Composer能够自动载入的位置。

我们需要能够在 IoC 容器中取得该类。所以，让我们增加一个绑定：

	App::bind('payment', function()
	{
		return new \PaymentGateway\Payment;
	});

最好注册该绑定的位置是创建一个新的名为`PaymentServiceProvider`[服务提供器](4.2/ioc.md#service-providers)，并且将该绑定加入到 `register` 方法。接下来你就可以配置 Laravel `app/config/app.php` 配置文件来加载该服务提供器。

接下来，我们就可以创建我们的 facade 类：

	use Illuminate\Support\Facades\Facade;

	class Payment extends Facade {

		protected static function getFacadeAccessor() { return 'payment'; }

	}

最后，如果你想，我们可以为我们 facade 设置一个别名到 `app/config/app.php` 配置文件里的 `aliases` 数组。现在，我们能够调用 Payment 类实例的 `process` 方法。 

	Payment::process();

### 关于自动载入别名一些注意点

在`aliases`数组中的有些类接口可能是不可行的，因为[PHP不会尝试去自动载入未定义的类型约定](https://bugs.php.net/bug.php?id=39003)。假设`\ServiceWrapper\ApiTimeoutException`别名是`ApiTimeoutException`，  如果`catch(ApiTimeoutException $e)` 是在`\ServiceWrapper`命名空间之外使用，它将会永远不会被捕获到, 即使真的有一个异常被抛出。 一个类似的问题存在在那些进行类型约定的别名类上。唯一的变通方法是摒弃别名，在每个你想使用类型约定的文件的开始的地方使用`use`。

<a name="mocking-facades"></a>
## 模拟Facades

单元测试是 facades 工作的重要体现。事实上，可测试性是 facedes 存在的主要原因。要了解更多信息，查看文档[模拟 facades](4.2/testing.md#mocking-facades)部分。

<a name="facade-class-reference"></a>
## Facade 类参考

下面你会找到所有的facade以及其包含的类。这是一个非常有用的工具，可以根据给定的facade快速定位到API文档。适用于[IoC 绑定](ioc.md) 的也同时给出了其key。


Facade  |  Class  |  IoC Binding
------------- | ------------- | -------------
App  |  [Illuminate\Foundation\Application](http://laravel.com/api/4.2/Illuminate/Foundation/Application.html)  | `app`
Artisan  |  [Illuminate\Console\Application](http://laravel.com/api/4.2/Illuminate/Console/Application.html)  |  `artisan`
Auth  |  [Illuminate\Auth\AuthManager](http://laravel.com/api/4.2/Illuminate/Auth/AuthManager.html)  |  `auth`
Auth (Instance)  |  [Illuminate\Auth\Guard](http://laravel.com/api/4.2/Illuminate/Auth/Guard.html)  |
Blade  |  [Illuminate\View\Compilers\BladeCompiler](http://laravel.com/api/4.2/Illuminate/View/Compilers/BladeCompiler.html)  |  `blade.compiler`
Cache  |  [Illuminate\Cache\Repository](http://laravel.com/api/4.2/Illuminate/Cache/Repository.html)  |  `cache`
Config  |  [Illuminate\Config\Repository](http://laravel.com/api/4.2/Illuminate/Config/Repository.html)  |  `config`
Cookie  |  [Illuminate\Cookie\CookieJar](http://laravel.com/api/4.2/Illuminate/Cookie/CookieJar.html)  |  `cookie`
Crypt  |  [Illuminate\Encryption\Encrypter](http://laravel.com/api/4.2/Illuminate/Encryption/Encrypter.html)  |  `encrypter`
DB  |  [Illuminate\Database\DatabaseManager](http://laravel.com/api/4.2/Illuminate/Database/DatabaseManager.html)  |  `db`
DB (Instance)  |  [Illuminate\Database\Connection](http://laravel.com/api/4.2/Illuminate/Database/Connection.html)  |
Event  |  [Illuminate\Events\Dispatcher](http://laravel.com/api/4.2/Illuminate/Events/Dispatcher.html)  |  `events`
File  |  [Illuminate\Filesystem\Filesystem](http://laravel.com/api/4.2/Illuminate/Filesystem/Filesystem.html)  |  `files`
Form  |  [Illuminate\Html\FormBuilder](http://laravel.com/api/4.2/Illuminate/Html/FormBuilder.html)  |  `form`
Hash  |  [Illuminate\Hashing\HasherInterface](http://laravel.com/api/4.2/Illuminate/Hashing/HasherInterface.html)  |  `hash`
HTML  |  [Illuminate\Html\HtmlBuilder](http://laravel.com/api/4.2/Illuminate/Html/HtmlBuilder.html)  |  `html`
Input  |  [Illuminate\Http\Request](http://laravel.com/api/4.2/Illuminate/Http/Request.html)  |  `request`
Lang  |  [Illuminate\Translation\Translator](http://laravel.com/api/4.2/Illuminate/Translation/Translator.html)  |  `translator`
Log  |  [Illuminate\Log\Writer](http://laravel.com/api/4.2/Illuminate/Log/Writer.html)  |  `log`
Mail  |  [Illuminate\Mail\Mailer](http://laravel.com/api/4.2/Illuminate/Mail/Mailer.html)  |  `mailer`
Paginator  |  [Illuminate\Pagination\Factory](http://laravel.com/api/4.2/Illuminate/Pagination/Factory.html)  |  `paginator`
Paginator (Instance)  |  [Illuminate\Pagination\Paginator](http://laravel.com/api/4.2/Illuminate/Pagination/Paginator.html)  |
Password  |  [Illuminate\Auth\Reminders\PasswordBroker](http://laravel.com/api/4.2/Illuminate/Auth/Reminders/PasswordBroker.html)  |  `auth.reminder`
Queue  |  [Illuminate\Queue\QueueManager](http://laravel.com/api/4.2/Illuminate/Queue/QueueManager.html)  |  `queue`
Queue (Instance) |  [Illuminate\Queue\QueueInterface](http://laravel.com/api/4.2/Illuminate/Queue/QueueInterface.html)  |
Queue (Base Class) |  [Illuminate\Queue\Queue](http://laravel.com/api/4.2/Illuminate/Queue/Queue.html)  |
Redirect  |  [Illuminate\Routing\Redirector](http://laravel.com/api/4.2/Illuminate/Routing/Redirector.html)  |  `redirect`
Redis  |  [Illuminate\Redis\Database](http://laravel.com/api/4.2/Illuminate/Redis/Database.html)  |  `redis`
Request  |  [Illuminate\Http\Request](http://laravel.com/api/4.2/Illuminate/Http/Request.html)  |  `request`
Response  |  [Illuminate\Support\Facades\Response](http://laravel.com/api/4.2/Illuminate/Support/Facades/Response.html)  |
Route  |  [Illuminate\Routing\Router](http://laravel.com/api/4.2/Illuminate/Routing/Router.html)  |  `router`
Schema  |  [Illuminate\Database\Schema\Blueprint](http://laravel.com/api/4.2/Illuminate/Database/Schema/Blueprint.html)  |
Session  |  [Illuminate\Session\SessionManager](http://laravel.com/api/4.2/Illuminate/Session/SessionManager.html)  |  `session`
Session (Instance)  |  [Illuminate\Session\Store](http://laravel.com/api/4.2/Illuminate/Session/Store.html)  |
SSH  |  [Illuminate\Remote\RemoteManager](http://laravel.com/api/4.2/Illuminate/Remote/RemoteManager.html)  |  `remote`
SSH (Instance)  |  [Illuminate\Remote\Connection](http://laravel.com/api/4.2/Illuminate/Remote/Connection.html)  |
URL  |  [Illuminate\Routing\UrlGenerator](http://laravel.com/api/4.2/Illuminate/Routing/UrlGenerator.html)  |  `url`
Validator  |  [Illuminate\Validation\Factory](http://laravel.com/api/4.2/Illuminate/Validation/Factory.html)  |  `validator`
Validator (Instance)  |  [Illuminate\Validation\Validator](http://laravel.com/api/4.2/Illuminate/Validation/Validator.html) |
View  |  [Illuminate\View\Factory](http://laravel.com/api/4.2/Illuminate/View/Factory.html)  |  `view`
View (Instance)  |  [Illuminate\View\View](http://laravel.com/api/4.2/Illuminate/View/View.html)  |


译者：mpandar（马胜盼）