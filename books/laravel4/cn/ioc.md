# IoC 容器


- [介绍](#introduction)
- [基本用例](#basic-usage)
- [哪里去注册绑定呢](#where-to-register)
- [自动解析](#automatic-resolution)
- [实际用例](#practical-usage)
- [服务提供器](#service-providers)
- [容器事件](#container-events)

<a name="introduction"></a>
## Introduction

Laravel使用IoC（Inversion of Control，控制倒转，这是一个设计模式，可以先查看下百科）容器这个强有力的工具管理类依赖。依赖注入（也是一种设计模式，一般用于实现IoC）是一个不用编写固定代码来处理类之间依赖的方法，相反的，这些依赖是在运行时注入的，这样允许处理依赖时具有更大的灵活性。

理解 Laravel IoC容器是构建强大应用程序所必要的，也有助于Laravel 核心本身。

<a name="basic-usage"></a>
## 基本用例

#### 绑定一个类型到容器

IoC 容器有两种方法来解决依赖关系：通过闭包回调或者自动解析。首先，我们来探究一下闭包回调。首先，需要绑定一个“类型”到容器中：

	App::bind('foo', function($app)
	{
		return new FooBar;
	});

#### 从容器中取得一个类型

	$value = App::make('foo');

当执行 `App::make` 方法，闭包函数被执行并返回结果。


#### 绑定一个”共享“类型到容器

有时，你只想将绑定到容器的类型处理一次，然后接下来从容器中取得的都应该是相同实例：

	App::singleton('foo', function()
	{
		return new FooBar;
	});

#### 绑定一个已经存在的类型实例到容器

你也可以使用`instance`方法，将一个已经存在的对象接口绑定到容器中：

	$foo = new Foo;

	App::instance('foo', $foo);

<a name="where-to-register"></a>
## 哪里去注册绑定呢

IoC绑定，很像事件句柄或者路由过滤，通常在"bootstrap code(引导代码)"之后完成。换句话说，它们在你的应用程序准备处理请求，也即是在一个路由或者控制器被实际执行之前执行。和其他引导代码一样，`start`文件通常作为IoC绑定注册一种方法。另外，你可以创建一个`app/ioc.php`（文件名不一定一样）文件，并在`start`文件中包含它。

如果你的应用程序有很大量IoC绑定，或者你想根据不同的分类将IoC绑定分割到不同的文件，你可以尝试在[服务提供器](#service-providers)中进行绑定

<a name="automatic-resolution"></a>
## 自动解析

#### 取得一个类

IoC容器足够强大，在许多场景下不需要任何配置就能取得类。例如

	class FooBar {

		public function __construct(Baz $baz)
		{
			$this->baz = $baz;
		}

	}

	$fooBar = App::make('FooBar');


注意我们虽然没有在容器中注册`FooBar`类，容器仍然可以取得该类，甚至自动注入`Baz`依赖！

当某个类型没有绑定到容器，IoC容器将使用 PHP 的反射工具来检查类和读取构造器的类型提示。使用这些信息，容器可以自动构建类实例。

#### 绑定一个接口实现

然而，在某些情况下，一个类可能依赖某个接口实现，而不是一个 “具体的类”。当在这种情况下，`App::bind`方法必须通知容器注入哪个接口实现：

	App::bind('UserRepositoryInterface', 'DbUserRepository');

现在考虑下这个控制器：

	class UserController extends BaseController {

		public function __construct(UserRepositoryInterface $users)
		{
			$this->users = $users;
		}

	}

由于我们将 `UserRepositoryInterface` 绑定了具体类，`DbUserRepository` 在该控制器创建时将会被自动注入到该控制器。
<a name="practical-usage"></a>
## 实际用例

Laravel 提供了几个方法使用 IoC 容器增强应用程序可扩展性和可测试性。一个主要的例子是取得控制器。所有控制器都通过 IoC 容器取得，意味着可以在控制器构造方法中对依赖的类型提示，它们将自动被注入。

#### 对控制器的依赖关系做类型提示

	class OrderController extends BaseController {

		public function __construct(OrderRepository $orders)
		{
			$this->orders = $orders;
		}

		public function getIndex()
		{
			$all = $this->orders->all();

			return View::make('orders', compact('all'));
		}

	}

在这个例子中，`OrderRepository` 将会自动注入到控制器。意味着当 [单元测试](/docs/4.2/testing) 模拟请求时，`OrderRepository` 将会绑定到容器以及注入到控制器中，允许无痛与数据库层交互。

#### IoC 使用的其他例子
	
[过滤器](/docs/4.2/routing#route-filters), [composers](/docs/4.2/responses#view-composers), 和 [事件句柄](/docs/4.2/events#using-classes-as-listeners)也能够从IoC容器中获取到。当注册它们的时候，只需要把它们使用的类名简单给出即可：

	Route::filter('foo', 'FooFilter');

	View::composer('foo', 'FooComposer');

	Event::listen('foo', 'FooHandler');

<a name="service-providers"></a>
## 服务提供器

服务器提供器是将一组相关 IoC 注册到单一路径的有效方法。将它们看做是一种引导组件的方法。在服务器提供器里，你可以注册自定义的验证驱动器，使用 IoC 容器注册应用程序仓库类，甚至是自定义 `Artisan` 命令。

事实上，大多数核心 Laravel 组件包含服务提供器。应用程序所有注册在服务提供器的均列在 `app/config/app.php` 配置文件的 `providers` 数组中。

#### 定义服务提供器

要创建服务提供器，只需继承 `Illuminate\Support\ServiceProvider` 类并且定义一个 `register` 方法：


	use Illuminate\Support\ServiceProvider;

	class FooServiceProvider extends ServiceProvider {

		public function register()
		{
			$this->app->bind('foo', function()
			{
				return new Foo;
			});
		}

	}

注意在 `register` 方法，应用程序通过 `$this->app` 属性访问 IoC 容器。一旦你已经创建了提供器并且想将它注册到应用程序中， 只需简单的放入 app 配置文件里 `providers` 数组中。

#### 运行时注册服务提供器

你也可以使用 `App::register` 方法在运行时注册服务提供器：

	App::register('FooServiceProvider');

<a name="container-events"></a>
## 容器事件

#### 注册获取事件监听者
	
容器在每次获取对象时都触发一个事件。你可以通过使用 `resolving` 方法来监听该事件：

	App::resolvingAny(function($object)
	{
		//
	});

	App::resolving('foo', function($foo)
	{
		//
	});

注意获取到的对象将会传入回调函数中。

译者：mpandar（马胜盼）