# 路由

- [基本路由](#basic-routing)
- [路由参数](#route-parameters)
- [路由过滤器](#route-filters)
- [命名路由](#named-routes)
- [路由组](#route-groups)
- [子域名路由](#sub-domain-routing)
- [路由前缀](#route-prefixing)
- [路由与模型绑定](#route-model-binding)
- [抛出 404 错误](#throwing-404-errors)
- [控制器路由](#routing-to-controllers)

<a name="basic-routing"></a>
## 基本路由

应用中的大多数路由都会定义在 `app/routes.php` 文件中。最简单的Laravel路由由URI和闭包回调函数组成。

#### 基本 GET 路由

	Route::get('/', function()
	{
		return 'Hello World';
	});

#### 基本 POST 路由

	Route::post('foo/bar', function()
	{
		return 'Hello World';
	});

#### 为多个动作注册同一个路由

	Route::match(array('GET', 'POST'), '/', function()
	{
		return 'Hello World';
	});

#### 注册一个可以响应任何HTTP动作的路由

	Route::any('foo', function()
	{
		return 'Hello World';
	});

#### 仅支持HTTPS的路由

	Route::get('foo', array('https', function()
	{
		return 'Must be over HTTPS';
	}));

实际开发中经常需要根据路由生成 URL，`URL::to`方法就可以满足此需求：

	$url = URL::to('foo');

<a name="route-parameters"></a>
## 路由参数

	Route::get('user/{id}', function($id)
	{
		return 'User '.$id;
	});

#### 可选路由参数

	Route::get('user/{name?}', function($name = null)
	{
		return $name;
	});

#### 带有默认值的可选路由参数

	Route::get('user/{name?}', function($name = 'John')
	{
		return $name;
	});

#### 用正则表达式限定的路由参数

	Route::get('user/{name}', function($name)
	{
		//
	})
	->where('name', '[A-Za-z]+');

	Route::get('user/{id}', function($id)
	{
		//
	})
	->where('id', '[0-9]+');

#### 传递参数限定的数组

当然，必要的时候你还可以传递一个包含参数限定的数组作为参数：

	Route::get('user/{id}/{name}', function($id, $name)
	{
		//
	})
	->where(array('id' => '[0-9]+', 'name' => '[a-z]+'))

#### 定义全局模式

如果希望在全局范围用指定正则表达式限定路由参数，可以使用 `pattern` 方法：

	Route::pattern('id', '[0-9]+');

	Route::get('user/{id}', function($id)
	{
		// Only called if {id} is numeric.
	});

#### 访问路由参数

如果想在路由范围外访问路由参数，可以使用 `Route::input` 方法：

	Route::filter('foo', function()
	{
		if (Route::input('id') == 1)
		{
			//
		}
	});

<a name="route-filters"></a>
## 路由过滤器

路由过滤器提供了非常方便的方法来限制对应用程序中某些功能访问，例如对于需要验证才能访问的功能就非常有用。Laravel框架自身已经提供了一些过滤器，包括 `auth`过滤器、`auth.basic`过滤器、`guest`过滤器以及`csrf`过滤器。这些过滤器都定义在`app/filter.php`文件中。

#### 定义一个路由过滤器

	Route::filter('old', function()
	{
		if (Input::get('age') < 200)
		{
			return Redirect::to('home');
		}
	});

如果过滤器返回了response，那么该response将被认为对应的是此次request，路由将不会被执行，并且，此路由中所有定义在此过滤器之后的代码也都不会被执行。

#### 为路由绑定过滤器

	Route::get('user', array('before' => 'old', function()
	{
		return 'You are over 200 years old!';
	}));

#### 将过滤器绑定为控制器Action

	Route::get('user', array('before' => 'old', 'uses' => 'UserController@showProfile'));

#### 为路由绑定多个过滤器

	Route::get('user', array('before' => 'auth|old', function()
	{
		return 'You are authenticated and over 200 years old!';
	}));

#### 通过数据绑定多个过滤器

	Route::get('user', array('before' => array('auth', 'old'), function()
	{
		return 'You are authenticated and over 200 years old!';
	}));

#### 指定过滤器参数

	Route::filter('age', function($route, $request, $value)
	{
		//
	});

	Route::get('user', array('before' => 'age:200', function()
	{
		return 'Hello World';
	}));

After filters receive a `$response` as the third argument passed to the filter:

	Route::filter('log', function($route, $request, $response)
	{
		//
	});

#### 基于模式的过滤器

你也可以只针对URI为一组路由指定过滤器。

	Route::filter('admin', function()
	{
		//
	});

	Route::when('admin/*', 'admin');

上述案例中，`admin`过滤器将会应用到所有以`admin/`开头的路由中。星号是通配符，将会匹配任意多个字符的组合。

还可以针对HTTP动作限定模式过滤器：

	Route::when('admin/*', 'admin', array('post'));

#### 过滤器类

过滤器的高级用法中，还可以使用类来替代闭包函数。由于过滤器类是通过[IoC container](/docs/ioc)实现解析的，所有，你可以在这些过滤器中利用依赖注入（dependency injection）的方法实现更好的测试能力。

#### 注册过滤器类

	Route::filter('foo', 'FooFilter');

默认情况下，FooFilter类里的filter方法将被调用：

	class FooFilter {

		public function filter()
		{
			// Filter logic...
		}

	}

如果你不想使用filter方法，那就指定另一个方法：

	Route::filter('foo', 'FooFilter@foo');

<a name="named-routes"></a>
## 命名路由

重定向和生成URL时，使用命名路由会更方便。你可以为路由指定一个名字，如下所示：

	Route::get('user/profile', array('as' => 'profile', function()
	{
		//
	}));

还可以为 controller action指定路由名称：

	Route::get('user/profile', array('as' => 'profile', 'uses' => 'UserController@showProfile'));

现在，你可以使用路由名称来创建URL和重定向：

	$url = URL::route('profile');

	$redirect = Redirect::route('profile');

可以使用`currentRouteName`方法来获取当前运行的路由名称:

	$name = Route::currentRouteName();

<a name="route-groups"></a>
## 路由组

有时你可能需要为一组路由应用过滤器。使用路由组就可以避免单独为每个路由指定过滤器了：

	Route::group(array('before' => 'auth'), function()
	{
		Route::get('/', function()
		{
			// Has Auth Filter
		});

		Route::get('user/profile', function()
		{
			// Has Auth Filter
		});
	});

你也可以在组数组中使用`namespace`参数来指定此组里的控制器都在一个给定的命名空间里：

	Route::group(array('namespace' => 'Admin'), function()
	{
		//
	});

<a name="sub-domain-routing"></a>
## 子域名路由

Laravel中的路由功能还支持通配符子域名，你可以在域名中指定通配符参数:

#### 注册子域名路由

	Route::group(array('domain' => '{account}.myapp.com'), function()
	{

		Route::get('user/{id}', function($account, $id)
		{
			//
		});

	});

<a name="route-prefixing"></a>
## 路由前缀

可以通过`prefix`属性为组路由设置前缀：

	Route::group(array('prefix' => 'admin'), function()
	{

		Route::get('user', function()
		{
			//
		});

	});

<a name="route-model-binding"></a>
## 路由与模型绑定

模型绑定，为在路由中注入模型实例提供了便捷的途径。例如，你可以向路由中注入匹配用户ID的整个模型实例，而不是仅仅注入用户ID。首先，使用 `Route::model` 方法指定要被注入的模型：

#### 给模型绑定参数

	Route::model('user', 'User');

然后，定义一个包含`{user}`参数的路由：

	Route::get('profile/{user}', function(User $user)
	{
		//
	});

由于我们已将`{user}`参数绑定到了`User`模型，因此可以向路由中注入一个`User`实例。例如，对`profile/1`的访问将会把ID为1的`User`实例注入到路由中。

> **注意：** 如果在数据库中无法匹配到对应的模型实例，404错误将被抛出。

如果你希望自定义"not found"行为，可以通过传递一个闭包函数作为 `model` 方法的第三个参数：

	Route::model('user', 'User', function()
	{
		throw new NotFoundHttpException;
	});

如果你想自己实现路由参数的解析，只需使用`Route::bind`方法即可：

	Route::bind('user', function($value, $route)
	{
		return User::where('name', $value)->first();
	});

<a name="throwing-404-errors"></a>
## 抛出 404 错误

有两种从路由中手动触发404错误的方法。首先，你可以使用`App::abort`方法：

	App::abort(404);

其次，你可以抛出`Symfony\Component\HttpKernel\Exception\NotFoundHttpException`异常。

更多关于处理404异常以及错误发生时自定义response的信息可以查看[错误](/docs/errors#handling-404-errors)文档。

<a name="routing-to-controllers"></a>
## 控制器路由

Laravel不光提供了利用闭包函数处理路由的功能，还可以路由到控制器，甚至支持创建 [resource controllers](/docs/controllers#resource-controllers)。

参见文档 [Controllers](/docs/controllers) 以获取更多信息。
