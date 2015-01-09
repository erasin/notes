# 控制器

- [基本控制器](#basic-controllers)
- [控制器过滤器](#controller-filters)
- [RESTful风格的控制器](#restful-controllers)
- [资源控制器](#resource-controllers)
- [处理缺失的方法](#handling-missing-methods)

<a name="basic-controllers"></a>
## 基本控制器

与在单个`routes.php`文件中定义所有的路由逻辑不同，你可能想要使用多个控制器类来组织控制器逻辑行为。控制器可以将与之相关的路由逻辑组织进一个类中，同时控制器也能使用一些更加高级的框架特性例如自动[依赖注入](ioc.md)。


控制器一般放在`app/controllers`目录下，这个目录默认会在你的`composer.json`文件的`classmap`选项中被注册。但是，从技术上来说，控制器可以存放在任意目录或者子目录中。路由声明并不依赖于控制器类文件在磁盘中的具体为止。因此，只要Composer能够知道在怎样自动加载控制器类，具体的控制器文件可以放在任何你想存放的地方。

下面是一个基本控制器类的例子：

	class UserController extends BaseController {

		/**
		 * 显示给定用户的资料。
		 */
		public function showProfile($id)
		{
			$user = User::find($id);

			return View::make('user.profile', array('user' => $user));
		}

	}

所有的控制器都应该扩展`BaseController`类。`BaseController`类文件也存放在`app/controllers`目录中，它主要用于存放一些共享的控制器逻辑。`BaseController`类继承了框架中的`Controller`类。现在，我们可以像下面的代码一样路由到这整个控制器的行为：  


	Route::get('user/{id}', 'UserController@showProfile');

如果你选择使用PHP命名空间嵌套或者组织你的控制器，可以在定义路由时简单是使用完整的类名:

	Route::get('foo', 'Namespace\FooController@method');

> **注意:** 由于我们使用[Composer](http://getcomposer.com)来自动加载我们的PHP类，控制器可以存放在文件系统中的任何地方，只要composer知道应该在哪里加载它们。在你的应用中，控制器不会被强制要求以某种特定的文件结构存放。路由到该控制器的逻辑和文件系统之前完全解耦。

你也可以在控制器路由中指定一个名词: 

	Route::get('foo', array('uses' => 'FooController@method',
											'as' => 'name'));

为了给控制器行为生成一个URL，你可以使用`URL::action`方法或者`action`辅助方法:  

	$url = URL::action('FooController@method');

	$url = action('FooController@method');

你可以使用`currentRouteAction`方法来获得目前正在运行的控制器名称:

	$action = Route::currentRouteAction();

<a name="controller-filters"></a>
## 控制器过滤器

[过滤器](routing.md#route-filters) 可以在控制器的路由中指定，它和"一般的"路由很相似：

	Route::get('profile', array('before' => 'auth',
				'uses' => 'UserController@showProfile'));

然而，你也可以在你的控制器内部指定过滤器:

	class UserController extends BaseController {

		/**
		 * 生成一个UserController实例
		 */
		public function __construct()
		{
			$this->beforeFilter('auth', array('except' => 'getLogin'));

			$this->beforeFilter('csrf', array('on' => 'post'));

			$this->afterFilter('log', array('only' =>
								array('fooAction', 'barAction')));
		}

	}

你可以使用一个闭包来内联的指定控制器过滤器:

	class UserController extends BaseController {

		/**
		 * 生成一个UserController实例
		 */
		public function __construct()
		{
			$this->beforeFilter(function()
			{
				//
			});
		}

	}

如果你想要在控制器中将另一个方法作为过滤器使用，你可以使用`@`语法来定义这个过滤器:

	class UserController extends BaseController {

		/**
		 * 生成一个UserController实例
		 */
		public function __construct()
		{
			$this->beforeFilter('@filterRequests');
		}

		/**
		 * 过滤进入该控制器的请求
		 */
		public function filterRequests($route, $request)
		{
			//
		}

	}

<a name="restful-controllers"></a>
## RESTful风格的控制器

Laravel允许你使用简单的、REST风格的命名规则，轻松的在一个控制器中使用一个路由来处理每一种行为。首先，使用`Route::controller`方法来定义路由:

	Route::controller('users', 'UserController');

`controller`方法接收两个参数。第一个参数是控制器处理的基本URI，第二个参数是控制器的类名。其次，为你的控制器添加方法，并在方法名称之前加上对应的HTTP动词：

	class UserController extends BaseController {

		public function getIndex()
		{
			//
		}

		public function postProfile()
		{
			//
		}

	}

`index`方法将会想要由控制器处理的根URI，在上面的例子中，即`users`。

如果你的控制器行为包含多个词，你可以在URI中使用"破折号"语法。例如，在我们的`UserController`类中，下面的控制器行为将会响应`users/admin-profile`URI:

	public function getAdminProfile() {}

<a name="resource-controllers"></a>
## 资源控制器

使用资源控制器能够围绕资源构建RESTful风格的控制器。例如，你可能想要创建一个控制器来管理存放在你的应用中的"照片"。通过使用Artisan命令行工具的`controller:make`命令以及`Route::resource`方法，我们可以快速的构建出这样一个控制器。

执行下面的命令，通过命令行创建一个控制器:

	php artisan controller:make PhotoController

现在我们可以为这个控制器注册一个resourceful风格的路由:

	Route::resource('photo', 'PhotoController');

这个简单的路由声明创建了多个路由来处理与照片资源相关的一系列RESTful行为。与此同时，生成的控制器也已经具备了一系列的方法来处理每一种路由。

#### 资源控制器处理的行为

动词       | 路径                         | 动作         | 路由名称
----------|-----------------------------|--------------|---------------------
GET       | /resource                   | index        | resource.index
GET       | /resource/create            | create       | resource.create
POST      | /resource                   | store        | resource.store
GET       | /resource/{resource}        | show         | resource.show
GET       | /resource/{resource}/edit   | edit         | resource.edit
PUT/PATCH | /resource/{resource}        | update       | resource.update
DELETE    | /resource/{resource}        | destroy      | resource.destroy

有时你只想要处理资源行为的一个子集:

	php artisan controller:make PhotoController --only=index,show

	php artisan controller:make PhotoController --except=index

同时，你也也可以在路由中制定一个行为的子集:

	Route::resource('photo', 'PhotoController',
					array('only' => array('index', 'show')));

	Route::resource('photo', 'PhotoController',
					array('except' => array('create', 'store', 'update', 'destroy')));

默认情况下，所有的资源控制器行为都有一个路由名称；然而，你可以通过在选项中使用一个`name`数组重载这些名称:

	Route::resource('photo', 'PhotoController',
					array('names' => array('create' => 'photo.build')));

#### 为资源控制器添加额外的路由

如果你需要为你的资源控制器添加默认资源路由之外的路由，你应该在调用`Route::resource`方法之前定义这些路由:

	Route::get('photos/popular');
	Route::resource('photos', 'PhotoController');

<a name="handling-missing-methods"></a>
## 处理缺失的方法

你可以在控制器中定义一个万能的方法，用来处理不能被其他所有方法处理的请求。这个方法的名称必须为`missingMethod`，它接收这个未被处理的请求的方法和参数数组作为参数。

#### 定义一个万能方法

	public function missingMethod($parameters = array())
	{
		//
	}
