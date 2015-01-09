# 配置

- [引言](#introduction)
- [环境配置](#environment-configuration)
- [提供者配置](#provider-configuration)
- [敏感信息保护配置](#protecting-sensitive-configuration)
- [维护模式](#maintenance-mode)

<a name="introduction"></a>
## 引言

Laravel 框架的所有配置文件都存储于 `app/config` 目录。每个文件中的每个选项都做了详细的记录，以便随时翻阅文件，熟悉提供给你的选项。

有时你需要在程序执行阶段访问配置的值。你可以使用 `Config` 类：

#### 访问一个配置的值

	Config::get('app.timezone');

你也可以指定一个默认值，如果配置选项不存在它将被返回：

	$timezone = Config::get('app.timezone', 'UTC');

#### 设置一个配置的值

注意“点”语法风格可以用于访问不同文件里的值，你也可以在程序执行阶段设置配置的值：

	Config::set('database.default', 'sqlite');

在程序执行阶段设置的配置的值仅作用于当前请求，并不会延续到后续请求中。

<a name="environment-configuration"></a>
## 环境配置

基于应用程序的运行环境拥有不同配置值通常是非常有益的。例如，相对于生产服务器你希望在你本地开发设备上使用一个不同的缓存驱动。使用基于环境的配置可以很容易的做到这点。

在 `config` 目录里简单的创建一个与你的环境同名的文件夹，比如 `local`。接着，创建你希望在这个环境中指定选项被覆盖的配置文件。例如，在 local 环境下覆盖缓存驱动，你将要在 `app/config/local` 里创建一个 `cache.php` 文件并包含以下内容：

	<?php

	return array(

		'driver' => 'file',

	);

> **注意：** 不要使用 'testing' 作为环境名称。这是为单元测试预留的。

注意，你不必指定基础配置文件中的 _每一个_ 选项，仅需指定你希望覆盖的选项。环境配置文件将 "叠加" 在基础配置文件之上。

接着，我们需要告知框架如何判定自己运行于哪个环境中。默认的环境始终是 `production`。然而，你可以在安装程序的根目录下 `bootstrap/start.php` 文件中设置其他环境。在这个文件中你将找到一个 `$app->detectEnvironment` 的调用。向这个方法传入的数组用于确定当前环境。必要时你可以增加其他环境和机器名。

    <?php

    $env = $app->detectEnvironment(array(

        'local' => array('your-machine-name'),

    ));

在这个例子中，'local' 是环境名称而 'your-machine-name' 是你本地服务器的主机名。在 Linux 和 Mac 上，你可以使用 `hostname` 终端命令来确定你的主机名。

如果你需要更灵活的环境检测，你可以通过向 `detectEnvironment` 方法传入一个 `匿名函数`，这允许你按照自己的方式执行环境检测：

	$env = $app->detectEnvironment(function()
	{
		return $_SERVER['MY_LARAVEL_ENV'];
	});

#### 访问当前的应用环境

你可以通过 `environment` 方法访问当前的应用环境：

	$environment = App::environment();

你也可以通过向 `environment` 方法传递参数来检测环境是否与给定的值匹配：

	if (App::environment('local'))
	{
		// 当前为 local 运行环境
	}

	if (App::environment('local', 'staging'))
	{
		// 当前为 local 或 staging 运行环境
	}

<a name="provider-configuration"></a>
### 提供者配置

当使用环境配置，你可能想要 "追加" 环境 [服务提供者](ioc.md#service-providers) 到你的基础 `app` 配置文件中。然而，如果你尝试这么做，你需要注意这个环境 `app` 提供者将会完全覆盖你的基础 `app` 配置文件中的值。要强制追加提供者，需要在你的环境 `app` 配置文件中使用 `append_config` 辅助函数：

	'providers' => append_config(array(
		'LocalOnlyServiceProvider',
	))

<a name="protecting-sensitive-configuration"></a>
## 敏感信息保护配置

对于 "真实" 的应用程序，保持你所有的敏感配置信息位于配置文件之外，这是明智的。诸如数据库密码，第三方 API 密钥，加密密钥等尽可能的放置于配置文件之外。所以，要放在哪里呢？谢天谢地，Laravel 提供了一个非常简单的方案来保护这些配置项，使用 "点" 风格的文件。

首先，[设置你的应用程序](configuration.md#environment-configuration) 识别你的机器是在 `local` 环境下。接着，在你项目的根目录创建一个 `.env.local.php` 文件，这通常与包含 `composer.json` 文件的目录相同。这个 `.env.local.php` 必须返回一个键值对数组，就像一个典型的 Laravel 配置文件：

	<?php

	return array(

		'TEST_STRIPE_KEY' => 'super-secret-sauce',

	);

这个文件中所有返回的键值对，将会自动通过 `$_ENV` 和 `$_SERVER` PHP "超全局变量" 变为可用。现在你可以在你的配置文件中引用这些全局变量：

	'key' => $_ENV['TEST_STRIPE_KEY']

确保在你的 `.gitignore` 文件中增加了对 `.env.local.php` 文件的忽略规则。这将允许你团队的其他开发者创建他们自己的本地环境配置，以及从源头隐藏你的敏感配置项。

现在，在你的生产服务器上，你项目的根目录里创建一个 `.env.php` 文件，包含你生产环境所对应的值。就像 `.env.local.php` 文件，生产环境 `.env.php` 文件不应该被包含在源码中。

> **注意：** 你可以为每一个应用程序支持的环境创建一个文件。例如，在 `development` 环境下将载入 `.env.development.php` 文件，如果它存在的话。

<a name="maintenance-mode"></a>
## 维护模式

当你的应用程序处于维护模式中，所有进入到你应用程序的路由都将显示一个自定义的视图。这使得当你的应用程序更新或进行维护时，可以很容易的 "禁用" 你的应用程序。在你的 `app/start/global.php` 文件中已经准备了一个 `App::down` 方法的调用。当你的应用程序处于维护模式中时，该方法的响应将发送给用户。

要启用维护模式，可以简单的执行 `down` Artisan 命令：

	php artisan down

要禁用维护模式，则使用 `up` 命令：

	php artisan up

当你的应用程序处于维护模式时，若需显示一个自定义视图，你可以在应用程序的 `app/start/global.php` 文件中添加如下代码：

	App::down(function()
	{
		return Response::view('maintenance', array(), 503);
	});

如果传递给 `down` 方法的闭包返回 `NULL`，那么在此次请求中将忽略维护模式。

### 维护模式 & 队列

在你的应用程序处于维护模式期间，不会有 [队列工作](queues.md) 被处理。一旦应用程序退出维护模式，这些工作将继续正常处理。
