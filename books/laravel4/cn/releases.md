# 版本说明

- [Laravel 4.2]（#laravel-4.2）
- [Laravel 4.1]（#laravel-4.1）

<a name="laravel-4.2"></a>
## Laravel 4.2

通过在4.2版本的安装目录下运行命令“php artisan changes”来获得此版本完整的变更列表, 也可以查看Github上的变更文件（https://github.com/laravel/framework/blob/4.2/src/Illuminate/Foundation/changes.json）。 这些更改记录只包含当前版本主要的功能改进和变更。

> **注意:** 在4.2版本的发布周期中，很多小的bug修复和功能改进被合并进了Laravel 4.1的各个发布版本中。所以，也一定要检查Laravel 4.1的变更列表！

### 最低要求PHP 5.4

Laravel 4.2 要求 PHP 5.4 或更高版本. 这个PHP版本的升级需求使得我们能够使用PHP的新特性，例如，为Laravel Cashier（/docs/billing）等工具提供的更具表现力的接口。 与PHP 5.3相比，PHP 5.4 在速度和执行效率上也有显著的提高。

### Laravel Forge

Laravel Forge是一个新的基于web的应用程序，它提供一种简单的方式来创建和管理你自己云端的PHP服务，这些服务包括Linode、DigitalOcean、Rackspace和Amazon EC2。Forge支持Nginx配置自动化，访问SSH key，Cron job的自动化，通过NewRelic或Papertrail进行服务监控，“Push To Deploy”，配置Laravel队列工人，此外，Forge提供最简单且最经济的方式来运行你的所有Laravel应用程序。

默认情况下，Laravel 4.2安装目录下的文件“app/config/database.php”是用来配置Forge的，使得新的应用能更方便的部署到这个平台上来。

在Forge的官方网站（https://forge.laravel.com）上，可以找到更多的有关Laravel Forge的信息。

### Laravel Homestead

Laravel Homestead是一个官方的Vagrant环境，它被用来开发强健的Laravel和PHP应用程序。在箱子被打包派发之前，绝大部分的箱子供应需求会被处理，它允许箱子极其快速的启动。Homestead包括Nginx 1.6、PHP 5.5.12、MySQL、Postgres、Redis、Memcached、Beanstalk、Node、Gulp、Grunt和Bower。Homestaed包含一个简单的配置文件“Homestead.yaml”，它能在单个箱子里管理多个Laravel应用程序。

默认情况下，Laravel 4.2安装目录包含一个配置文件“app/config/local/database.php”，它是用来配置使用箱子以外的Homestead数据库的，这使得Laravel初始化安装目录和配置更加方便。

官方文档已经包括Homestead文档（/docs/homestead）。

### Laravel Cashier

Laravel Cashier是一个简单的、富于表现力的库，它用来管理条形码的订阅计费。随着Laravel 4.2的引入，尽管安装组件本身仍然是可选的，但是我们在Laravel主文档里包含了Cashier文档。这个版本的Cashier修复了很多bug，支持多币种，并兼容最新的条形码API。

### 守护进程队列工人（Daemon Queue Workers）

Artisan的“queue:work”命令现在支持“--daemon”选项，它用来启动一个“守护进程模式”的工人，这个模式使得工人在不需要重启框架的情况下继续工作。这会显著的减少CPU的使用率，其代价只是使得应用程序的部署过程稍显复杂。

在队列文档（/docs/queues#daemon-queue-workers）里能找到更多的队列工人相关的信息。

### 邮件API驱动

Laravel 4.2为“邮件”功能引入了新的Mailgun和Mandrill接口驱动程序。对很多应用程序来说，比起SMTP方式，此接口提供了一种更快和更可靠发送电子邮件的方法。新的驱动程序采用Guzzle 4 HTTP库。

### 软删除特性（Soft Deleting Traits）

一种为“软删除”和其它“全局范围”而生的更干净的架构通过PHP 5.4的特性被引入进来。这种新架构考虑到了更简单的构建全局特性等功能，也可考虑到了一个更干净的框架本身的关注点分离问题。

在文档（/docs/eloquent#soft-deleting）里能找到更多的关于“软删除特性”的信息。

### 方便的认证和提醒特性（Convenient Auth & Remindable Traits）

现在，Laravel 4.2的安装目录默认使用简单的特性来包含认证和密码提醒用户接口所需要的属性。这提供了一个更干净的默认的箱子之外的“用户”模型文件。

### 简单的分页

一个新方法“simplePaginate”被加入到查询和Eloquent构建器里，当在分页页面里使用简单的“上一个”和“下一个”链接时，使用此方法查询就会更高效。

### 迁移确认（Migration Confirmation）

现在，在生产中的破坏性操作将被要求确认。使用“--force”选项的迁移命令可以强制执行而不会有任何提示。

<a name="laravel-4.1"></a>
## Laravel 4.1

### 全部变更列表

通过在4.1版本的安装目录下运行命令“php artisan changes”来获得此版本完整的变更列表, 也可以查看Github上的变更文件（https://github.com/laravel/framework/blob/4.1/src/Illuminate/Foundation/changes.json）。这些更改记录只包含当前版本主要的功能改进和变更。

### 新SSH组件

一个全新的“SSH”组件被引入到此版本中。其特性允许你很容易的SSH到远程服务器并运行命令。想了解更多信息，请查阅SSH组件文档（/docs/ssh）。

新的“php artisan tail”命令使用了新SSH组件。更多信息，请查阅“tail”命令文档（http://laravel.com/docs/ssh#tailing-remote-logs）。

### Boris In Tinker

命令“php artisan tinker”使用的是Boris REPL（https://github.com/d11wtq/boris），首先你的系统得支持Boris REPL。要使用这个特性必须安装PHP扩展“readline”和“pcntl”。如果你没有这两个扩展，将使用4.0版本的shell.

### Eloquent改进点

一种新的“hasManyThrough”关系已经被加入到Eloquent里。要学习怎样使用，请查阅Eloquent文档（/docs/eloquent#has-many-through）。

一个新方法“whereHas”也已经被引入，此方法允许基于关系约束的索引模型，相关文档（/docs/eloquent#querying-relations）。

### 数据库读写连接

在数据库层有可供使用的自动处理单独读写的连接，包括队列构建器和Eloquent。更多相关信息，请查阅文档（/docs/database#read-write-connections）。

### 队列优先权

现在支持通过传递给命令“queue:listen”一组以逗号分隔的列表来设定队列的优先级。

### 失败的队列工作处理

现在，在命令“queue:listen”中使用“--tries”选项时，队列功能会自动处理失败的工作。在队列文档（/docs/queues#failed-jobs）里可以找到处理失败工作的更多信息。

### 缓存标记

缓存章节（sections）已经被标记（tags）取代。缓存标记允许你把多个标记赋值给一个缓存项，也允许把所有的缓存项赋值给一个标记。在缓存文档（/docs/cache#cache-tags）里可以找到使用缓存标记相关的更多信息。

### 灵活的密码提醒

密码提醒引擎已经被改成当验证密码的时候为开发者提供很大的灵活性，引擎会闪存状态信息到会话里。有关使用增强的密码提醒的更多信息，请查阅文档（/docs/security#password-reminders-and-reset）。

### 改善的路由引擎

Laravel 4.1以完全重写的路由层为特色。虽然接口是相同的，但是路由注册比4.0版本要快100%。整个引擎被大大的简化了，而且依赖于Symfony的路由被最小化到路由表达式的编译里了。

### 改善的会话引擎

在这个版本中，我们也引入了一个全新的会话引擎。 类似于路由的改进，新的会话层更简洁更快。我们不再使用Symfony（PHP的）会话处理功能，而是使用一种更简单更容易维护的定制的解决方案。

### Doctrine DBAL

如果你在迁移（migrations）中用到了“renameColumn”方法，你需要在“composer.json”文件中添加“doctrine/dbal”依赖。这个包不再默认包含在Laravel中。