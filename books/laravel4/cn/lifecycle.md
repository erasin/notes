# 请求的生命周期

- [概述](#overview)
- [请求的生命周期](#request-lifecycle)
- [启动文件](#start-files)
- [应用程序事件](#application-events)

<a name="overview"></a>
## 概述

在现实世界中使用工具时，如果理解了工具的工作原理，使用起来就会更加有底气。应用开发也是如此。当你理解了开发工具是如何工作的，使用起来就会更加自如。这篇文档的目标就是提供一个高层次的概述，使你对于Laravel框架的运行方式有一个较好的把握。在更好地了解了整个框架之后，框架的组件和功能就不再显得那么神秘，开发起应用来也更加得心应手。这篇文档包含了关于请求生命周期的高层次概述，以及启动文件和应用程序事件的相关内容。

如果你不能立即理解所有的术语，别灰心，可以先有一个大致的把握，在阅读文档其他章节的过程中继续积累和消化知识。

<a name="request-lifecycle"></a>
## 请求的生命周期

发送给应用程序的所有请求都经由 `public/index.php` 脚本处理。如果使用的是 Apache 服务器，Laravel中包含的 `.htaccess` 文件将对所有请求进行处理并传递给 `index.php`。这是Laravel从接受客户端请求到返回响应给客户端的整个过程的开始。若能对于Laravel的引导过程(bootstrap process)有一个大致的认识，将有助于理解框架，我们不妨先讨论这个。

到目前为止，学习Laravel引导过程所需掌握的最重要的概念就是 **服务提供器**。打开 `app/config/app.php` 配置文件，找到 `providers` 数组，你会发现一个服务提供器的列表。这些提供器充当了Laravel的主要引导机制。在我们深入服务提供器之前，先回到 `index.php`的讨论。当一个请求进入 `index.php` 文件，`bootstrap/start.php` 文件会被加载。这个文件会创建一个 Laravel `Application` 对象，该对象同时作为框架的 [IoC 容器](/docs/ioc)。

`Application` 对象创建完成后，框架会设置一些路径信息并运行 [环境检测](/docs/configuration#environment-configuration) 。然后会执行位于Laravel源码内部的引导脚本，并根据你的配置文件设置时区、错误报告等其他信息。除了配置这些琐碎的配置选项以外，该脚本还会做一件非常重要的事情：注册所有为应用程序配置的服务提供器。

简单的服务提供器只包含一个方法：`register`。当应用程序对象通过自身的 `register` 方法注册某个服务提供器时，会调用该服务提供器的 `register` 方法。服务提供器通过这个方法向 [IoC 容器](/docs/ioc) 注册一些东西。从本质上讲，每个服务提供器都是将一个或多个 [闭包](http://us3.php.net/manual/en/functions.anonymous.php) 绑定到容器中，你可以通过这些闭包访问绑定到应用程序的服务。例如，`QueueServiceProvider` 注册了多个闭包以便使用与 [队列](/docs/queues) 相关的多个类。当然，服务提供器并不局限于向IoC容器注册内容，而是可以用于任何引导性质的任务。服务提供器可以注册事件监听器、视图合成器、Artisan命令等等。

在注册完所有服务提供器后，`app/start` 下的文件会被加载。最后，`app/routes.php` 文件会被加载。一旦 `routes.php` 文件被加载，Request 对象就被发送给应用程序对象，继而被派发到某个路由上。

我们总结一下：

1. 请求进入 `public/index.php` 文件。
2. `bootstrap/start.php` 文件创建应用程序对象并检测环境。
3. 内部的 `framework/start.php` 文件配置相关设置并加载服务提供器。
4. 加载应用程序 `app/start` 目录下的文件。
5. 加载应用程序的 `app/routes.php` 文件。
6. 将 Request 对象发送给应用程序对象，应用程序对象返回一个 Response 对象。
7. 将 Response 对象发回客户端。

你应该已经掌握了 Laravel 应用程序是如何处理发来的请求的。下面我们来看一下启动文件。

<a name="start-files"></a>
## 启动文件

应用程序的启动文件被存放在`app/start`目录中。默认情况下，该目录下包含三个文件：`global.php`、`local.php` 和 `artisan.php`文件。需要获取更多关于`artisan.php`的信息，可以参考文档[Artisan 命令行](/docs/commands#registering-commands)。

`global.php`启动文件默认包含一些基本项目，例如[日志](/docs/errors)的注册以及载入`app/filters.php` 文件。然而，你可以在该文件里做任何你想做的事情。无论在什么环境下，它都将会被自动包含进_每一个_request中。而`local.php` 文件仅在`local`环境下被执行。获取更多关于环境的信息，请查看文档[配置](/docs/configuration)。

当然，如果除了`local`环境你还有其他环境的话，你也可以为针对这些环境创建启动文件。这些文件将在应用程序运行在该环境中时被自动包含。假设你在 `bootstrap/start.php` 文件中配置了一个 `development` 环境，你可以创建一个 `app/start/development.php` 文件，在那个环境下任何进入应用程序的请求都会包含该文件。

### 启动文件里存放什么

启动文件主要用来存放任何“引导”性质的代码。例如，你可以在启动文件中注册视图合成器，配置日志信息，或是进行一些PHP设置等。具体做什么取决于你。当然了，把所有引导代码都丢到启动文件里会使启动文件变得杂乱。对于大型应用而言，或是启动文件显得太杂乱了，请考虑将某些引导代码移至 [服务提供器](/docs/ioc#service-providers) 中。

<a name="application-events"></a>
## 应用程序事件

#### 注册应用程序事件

你还可以通过注册 `before`、`after`、`finish` 和 `shutdown`应用程序事件以便在处理request之前或后做一些操作：

	App::before(function($request)
	{
		//
	});

	App::after(function($request, $response)
	{
		//
	});

这些事件的监听器会在每个到达应用程序的请求处理之前（`before`）或之后（`after`）运行。可以利用这些事件来设置全局过滤器(filter)，或是对于发回客户端的响应(response)统一进行修改。你可以在某个启动文件中或者 [服务提供器](/docs/ioc#service-providers) 中注册这些事件。

你也可以在 `matched` 事件上注册一个监听器，当一个传入请求已经和一个路由相匹配，但还未执行此路由之前，此事件就会被触发：

	Route::matched(function($route, $request)
	{
		//
	});

当来自应用程序的响应发送至客户端后会触发 `finish` 事件。这个事件适合处理应用程序所需的最后的收尾工作。当所有 `finish` 事件的监听器都执行完毕后会立即触发 `shutdown` 事件，如果想在脚本结束前再做一些事情，这是最后的机会。不过在大多数情况下，你都不需要用到这些事件。