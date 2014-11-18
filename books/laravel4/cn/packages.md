# Package Development

- [简介](#introduction)
- [创建包](#creating-a-package)
- [包结构](#package-structure)
- [服务提供器](#service-providers)
- [延迟加载服务提供器](#deferred-providers)
- [包约定](#package-conventions)
- [开发流程](#development-workflow)
- [包路由](#package-routing)
- [包配置](#package-configuration)
- [包视图](#package-views)
- [包迁移](#package-migrations)
- [包Assets](#package-assets)
- [发布包](#publishing-packages)

<a name="introduction"></a>
## 简介

包是向Laravel中添加功能最重要的途径。它通过一种强大的方式几乎可以包含任意功能，比如处理日期的扩展包[Carbon](https://github.com/briannesbitt/Carbon)，完整的[BDD](http://baike.baidu.com/view/1384794.htm)测试框架扩展包[Behat](https://github.com/Behat/Behat).

当然，还有很多不同类型的包。有些包是独立的，这意味着它们可以在任何框架中工作，而不仅仅是Laravel。上面提到的Carbon和Behat就是独立的包。要在Laravel中使用这些包只需要在`composer.json`文件中指明。

另一方面，有些包仅支持Laravel。在上一个Laravel版本中，这些类型的包我们称为"bundles"。这些包可以增强Laravel应用的路由、控制器、视图、配置和迁移。由于开发独立的包不需要专门的过程，因此，本手册主要涵盖针对Laravel开发独立的包。

所有Laravel包都是通过[Packagist](http://packagist.org)和[Composer](http://getcomposer.org)发布的，因此很有必要学习这些功能强大的PHP包发布管理工具。

<a name="creating-a-package"></a>
## 创建包

为Laravel创建一个包的最简单方式是使用Artisan的`workbench`命令。首先，你需要在`app/confg/workbench.php`文件中配置一些参数。在该文件中，你会看到`name`和`email`两个参数，这些值是用来为您新创建的包生成`composer.json`文件的。一旦你提供了这些值，就可以开始构建一个新包了！

#### 执行Artisan的workbench命令

	php artisan workbench vendor/package --resources

厂商名称（vendor name）是用来区分不同作者构建了相同名字的包。例如，如果我（Taylor Otwell）创建了个名为"Zapper"的包，厂商名就可以叫做`Taylor`，包名可以叫做`Zapper`。默认情况下，workbench命令建的包是不依赖任何框架的；不过，`resources`命令将会告诉workbench生成关联Laravel的一些特殊目录，例如migrations、views、config等。

一旦执行了`workbench`命令，新创建的包就会出现在Laravel安装跟目录下的`workbench`目录中。接下来就应该为你创建的包注册`ServiceProvider`了。你可以通过在`app/config/app.php`文件里的`provides`数组中添加该包。这将通知Laravel在应用程序开始启动时加载该包。服务提供者（Service providers）使用`[Package]ServiceProvider`样式的命名方式。所以，以上案例中，你需要将`Taylor\Zapper\ZapperServiceProvider`添加到providers数组。

一旦注册了provider，你就可以开始写代码了！然而，在此之前，建议你查看以下部分来了解更多关于包结构和开发流程的知识。

> **注意:** 如果你的Service providers提示无法找到, 在项目根目录执行`php artisan dump-autoload` 

<a name="package-structure"></a>
## 包结构

执行`workbench`命令之后，你的包结构将被初始化，并能够与laravel框架完美融合：

#### 包的基本目录结构

	/src
		/Vendor
			/Package
				PackageServiceProvider.php
		/config
		/lang
		/migrations
		/views
	/tests
	/public

让我们来深入了解该结构。`src/Vendor/Package`目录是所有`class`的主目录，包括`ServiceProvider`。`config`、`lang`、`migrations`和`views`目录，就如你所猜测，包含了你创建的包的相应资源。包可以包含这些资源中的任意几个，就像一个"常规"的应用。

<a name="service-providers"></a>
## 服务提供器

服务提供器只是包的引导类。默认情况下，他们包含两个方法：`boot`和`register`。你可以在这些方法内做任何事情，例如：包含路由文件、注册IoC容器的绑定、监听事件或者其他任何你想做的事情。

`register`方法在服务提供器注册时被立即调用，而`boot`方法仅在请求被路由到之前调用。因此，如果服务提供器中的动作（action）依赖另一个已经注册的服务提供器，或者你正在覆盖另一个服务提供其绑定的服务，就应该使用`boot`方法。

当使用`workbench`命令创建包时，`boot`方法已经包含了如下的动作：

	$this->package('vendor/package');

该方法告诉Laravel如何为应用程序加载视图、配置或其他资源。通常情况下，你没有必要改变这行代码，因为它会根据workbench的默认约定将包设置好的。

默认情况下，一旦注册了一个包，那么它的资源可以通过"package"方法在`vendor/package`中找到。你也可以向`package` 方法中传入第二个参数来重写这个方法。例如

	//向 `package` 方法中传入一个自定义的命名空间
	$this->package('vendor/package', 'custom-namespace');

	//现在，这个包的资源现在可以通过这个自定义的命名空间来访问
	$view = View::make('custom-namespace::foo');

Laravel并没有为service provider提供“默认”的存放地点。您可以根据自己的喜好，将它们放置在任何地方，您也可以将它们统一组织在一个`Providers`命名空间里，并放置在应用的`app`目录下。这些文件可以被放置在任何地方，只需保证Composer的[自动加载](http://getcomposer.org/doc/01-basic-usage.md#autoloading)组件知道如何加载这些类。

如果你改变了你的包得资源的位置，比如配置文件或者视图，你需要在`package`函数中传递第三个参数，指定你的资源位置

	$this->package('vendor/package', null, '/path/to/resources');

<a name="deferred-providers"></a>
## 延迟加载服务提供器

如果你写了一个服务提供器，但没有注册任何资源，比如配置文件或视图等，你可以选择"延迟"加载你的提供器。延迟加载的服务提供器在这个服务真正被IoC容器需要的时候才会被加载和注册。如果这个服务提供器没有被当前的请求路由需要，那么这个提供器永远不会被加载

如果想延时载入你的服务提供器，只需要在提供器重设置`defer`属性为`true`:

	protected $defer = true;


接下来你需要重写继承自基类`Illuminate\Support\ServiceProvider`的`provides`方法，并返回绑定了你的提供器，对应的IoC容器中类型的数组集合。例如，你的提供器在IoC容器注册了`package.service` 和 `package.another-service`两个类型，你的`provides`的方法看起来应该是这个样子：

	public function provides()
	{
		return array('package.service', 'package.another-service');
	}

<a name="package-conventions"></a>
## 包约定


要使用包中的资源，例如配置或视图，需要用双冒号语法：

#### 从包中载入视图

	return View::make('package::view.name');

#### 获取包的某个配置项

	return Config::get('package::group.option');

> **注意:** 如果你包中包含迁移，请为迁移名（migration name）添加包名作为前缀，以避免与其他包中的类名冲突。


<a name="development-workflow"></a>
## 开发流程

当开发一个包时，能够使用应用程序上文是相当有用的，这样将允许你很容易的解决视图模板的等问题。所以，我们开始，安装一个全新的Laravel框架，使用`workbench`命令创建包结构。

在使用`workbench`命令创建包后。你可以在`workbench/[vendor]/[package]`目录使用`git init`，并在workbench中直接`git push`！这将允许你在应用程序上下文中方便开发而不用为反复使用`composer update`命令苦扰。


当包存放在`workbench`目录时，你可能担心Composer如何知道自动加载包文件。当workbench目录存在，Laravel将智能扫描该目录，在应用程序开始时加载它们的Composer自动加载文件！

如果你需要重新生成包的自动加载文件，你可以使用`php artisan dump-autoload`命令，这个命令将会重新为您的整个项目生成自动加载文件，也包括你的工作台(workbenches)中的包

#### 运行Artisan的自动加载命令

	php artisan dump-autoload

<a name="package-routing"></a>
## 包路由

在之前的Laravel版本中，`handlers`用来指定哪个URI包会响应。然而，在Laravel4中，一个包可以响应任意URI。要在包中加载路由文件，只需在服务提供器的`boot`方法`include`它。

#### 在服务提供器中包含路由文件

	public function boot()
	{
		$this->package('vendor/package');

		include __DIR__.'/../../routes.php';
	}


> **注意:** 如果你的包中使用了控制器, 你需要确保正确配置了`composer.json`文件的auto-load字段.

<a name="package-configuration"></a>
## 包配置

#### 访问包配置文件

有时创建的包可能会需要配置文件。这些配置文件应该和应用程序配置文件相同方法定义。并且，当使用 `$this->package`方法来注册服务提供器时，那么就可以使用“双冒号”语法来访问：

	Config::get('package::file.option');

#### 访问包单一配置文件

然而，如果你包仅有一个配置文件，你可以简单命名为`config.php`。当你这么做时，你可以直接访问该配置项，而不需要特别指明文件名：
	Config::get('package::option');

#### Registering A Resource Namespace Manually
#### 手动注册资源的命名空间

有时候，你可能希望在`$this->package`特有方法之外注册包的资源，比如视图。通常只有在这些资源不在约定的位置的时候才应该这么做。如果需要手动注册资源，你可以使用`View`, `Lang`, 和 `Config` 中的`addNamespace`方法实现:

	View::addNamespace('package', __DIR__.'/path/to/views');

一旦这个资源命名空间被注册，你就可以使用这个空间名，并使用"双冒号"语法去访问这个资源

	return View::make('package::view.name');

`View`, `Lang`, 和 `Config`类中的`addNamespace`方法的参数都是一样的

### 级联配置文件

但其他开发者安装你的包时，他们也许需要覆盖一些配置项。然而，如果从包源代码中改变值，他们将会在下次使用Composer更新包时又被覆盖。替代方法是使用artisan命令`config:publish`：

	php artisan config:publish vendor/package

当执行该命令，配置文件就会拷贝到`app/config/packages/vendor/package`，开发者就可以安全的更改配置项了。

> **注意:** 开发者也可以为该包创建指定环境下的配置文件，替换某些配置项然后并放置在`app/config/packages/vendor/package/environment`.

<a name="package-views"></a>
## 包视图

如果在你的应用程序中使用某个包，你可能会定制这个包的视图。你可以很轻松的将这个包的视图文件导入到你自己的`app/views`目录，只需使用`view:publish`这个Artisan命令

	php artisan view:publish vendor/package

这个命令会将包的视图文件迁移到`app/views/packages`目录，如果这个目录不存在，将会被创建。一旦这些视图文件被公开创建，你就可以按照你自己的喜好去修改他们，这些导出的视图会比包自己的视图文件优先载入。

<a name="package-migrations"></a>
## 包迁移

#### 为工作台(Workbench)的包创建迁移

你可以很容易在包中创建和运行迁移。要为工作台里的包创建迁移，使用`--bench`选项：

	php artisan migrate:make create_users_table --bench="vendor/package"

#### 为工作台(Workbench)的包运行迁移

	php artisan migrate --bench="vendor/package"

#### 为已安装的包执行迁移

要为已经通过Composer安装在vendor目录下的包执行迁移，你可以直接使用`--package`：

	php artisan migrate --package="vendor/package"

<a name="package-assets"></a>
## 包Assets

#### 将包Assets移动到Public

有些包可能含有assets，例如JavaScript，CSS，和图片。然而，我们无法链接到`vendor`或`workbench`目录里的`assets`，所以我们需要可以将这些`assets`移入应用程序的`public`目录。`asset:publish`命令可以实现：

	php artisan asset:publish

	php artisan asset:publish vendor/package

如果这个包仍在`workbench`中，那么请使用`--bench`指令:

	php artisan asset:publish --bench="vendor/package"

这个命令将会把assets移入与“供应商”和“包名”相对应的`public/packages`目录下面。因此，包名为`userscape/kudos`的assets将会被移至`public/packages/userscape/kudos`。通过使用这个assets发布方法，可以让您安全的在包中的view内访问assets路径。

<a name="publishing-packages"></a>
## 发布包

当你创建的包准备发布时，你应该将包提交到 [Packagist](http://packagist.org) 仓库。如果你的包只针对Laravel，最好在包的`composer.json`文件中添加`laravel`标签

还有，在发布的版本中添加tag，以便开发者能当请求你的包在他们`composer.json`文件中依赖稳定版本。如果稳定版本还没有好，考虑直接在Composer中使用`branch-alias`。

一旦你的包发布，放舒心，继续在由`workbench`创建的包中，结合应用程序上下文进行开发。对于在发布包后继续进行开发，这是相当便利的。


一些组织使用他们私有分支包为他们自己开发者。如果你对这感兴趣，查看Composer团队构建的[Satis](http://github.com/composer/satis) 文档。


译者：mpandar（马胜盼）
