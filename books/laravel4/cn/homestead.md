# Laravel Homestead

- [前言](#introduction)
- [包含的软件](#included-software)
- [安装和设置](#installation-and-setup)
- [日常使用](#general-usage)
- [端口](#ports)

<a name="introduction"></a>
## 前言

Laravel努力的在整个PHP开发当中提供令人愉快的开发体验，包括你自己本地的开发环境。Vagrant（http://vagrantup.com）提供一种简单又优雅的方式来管理和装备虚拟机。

Laravel Homestead是一个官方的、预封装的Vagrant“箱子”，它提供给你一个奇妙的开发环境而不需要你在本机上安装PHP、web服务器、和其它的服务器软件。不用再担心搞乱你的操作系统！Vagrant箱子是完全可支配的。如果出现故障，你可以在几分种内完成销毁和重建箱子！

Homestead能运行在所有的Windows、Mac和Linux上，它包含了Nginx、PHP 5.5、MySQL、Postgres、Redis、Memcached和你开发神奇的Laravel应用程序需要的所有其它软件。

Homestead is currently built and tested using Vagrant 1.6.

<a name="included-software"></a>
## 包含的软件

- Ubuntu 14.04
- PHP 5.5
- Nginx
- MySQL
- Postgres
- Node （With Bower, Grunt, and Gulp）
- Redis
- Memcached
- Beanstalkd
- [Laravel Envoy](ssh.md#envoy-task-runner)
- Fabric + HipChat Extension

<a name="installation-and-setup"></a>
## 安装和设置

### 安装VirtualBox和Vagrant

在启动Homestead环境之前，你必须安装VirtualBox（https://www.virtualbox.org/wiki/Downloads）和Vagrant（http://www.vagrantup.com/downloads.html）。这两个软件为所有主流的操作系统提供了简单易用的可视化安装界面。

### 添加Vagrant箱子

一旦VirtualBox和Vagrant安装完成，你应该添加“laravel/homestead”箱子到你的Vagrant安装目录下，在终端使用下面的命令，这将花费几分钟的时间来下载箱子，这取决于你的网速：

	vagrant box add laravel/homestead

### 克隆Homestead仓库

一旦箱子被添加到Vagrant安装目录下，你应该克隆或下载这个仓库。因为Homestead箱子做为主机将为你的所有Laravel（和PHP）项目提供服务，所以可以考虑克隆这个仓库到主Homestead目录，这个目录存放着你所有的Laravel项目。

	git clone https://github.com/laravel/homestead.git Homestead

### 设置你的SSH密钥

接下来，你应该编辑仓库里的“Homestead.yaml”文件。在这个文件里，你可以配置公共SSH密钥的路径，也可以配置主机与Homestead虚拟机的共享目录。

还没有SSH密钥？在Mac和Linux机器上，通常你可以使用下面的命令创建一个SSH密钥对：

	ssh-keygen -t rsa -C "your@email.com"

在Windows机器上，你可以安装Git（http://git-scm.com/）工具，并使用Git自带的“Git Bash”命令行工具执行上面的命令。或者，你可以使用PuTTY（http://www.chiark.greenend.org.uk/~sgtatham/putty/download.html）工具或PuTTYgen（http://www.chiark.greenend.org.uk/~sgtatham/putty/download.html）工具。

一旦你创建了一个SSH密钥，就可以在“Homestead.yaml”文件里为“authorize”属性指定密钥的路径。

### 配置共享目录

“Homestead.yaml”文件里的“folders”属性列出所有你想与Homestead环境共享的目录。当这些目录中的文件发生了改变，它们将在本机和Homestead环境之间保持同步。你可以根据需要配置尽可能多的共享目录！

### 配置Nginx站点

不熟悉Nginx？没关系。Homestead环境里的“sites”属性允许你轻松地将一个“域”映射到一个目录。“Homestead.yaml”文件里包含一个示例站点配置。再强调一遍，你可以根据需要添加尽可能多的站点到Homestead环境里。Homestead能够为你的每一个Laravel项目提供一个方便的虚拟环境！

### Bash Aliases

想在你的Homestead箱子里加入Bash aliases，只需要在Homestead的根目录里简单的添加“aliases”文件即可。

### 启动Vagrant箱子

一旦你按照意愿编辑了“Homestead.yaml”文件，就可以在终端上的“Homestead”目录下执行“vagrant up”命令。Vagrant将启动虚拟机，并自动配置共享目录和Nginx站点！

不要忘记把你的Nginx站点的“域”添加到机器里的“hosts”文件里！“hosts”文件将把对本地域的请求重定向到Homestead环境里。在Mac和Linux机器上，这个文件位于“/etc”目录。在Windows机器上，它位于“C:\Windows\System32\drivers\etc”目录。你添加到此文件的内容就像下面这样：

	127.0.0.1  homestead.app

一旦你把域添加到“hosts”文件，你就可以通过浏览器在8000端口上访问此站点！

	http://homestead.app:8000

想知道如何连接数据库，请接着看！

<a name="daily-usage"></a>
## 日常使用

### 通过SSH连接

想通过SSH连接到Homestead环境，你应该使用“Homestead.yaml”文件里指定的SSH密钥在端口2222上连接到“127.0.0.1”。你也可以在“Homestead”目录下简单的执行“vagrant ssh”命令连接到“Homestead”环境。

如果你想更方便，把下面的内容添加到“~/.bash_aliases”或“~/.bash_profile”文件里，这将会提供给你很大的帮助:

	alias vm='ssh vagrant@127.0.0.1 -p 2222'

### 连接到数据库

“homestead”数据库是为箱子外面的MySQL和Postres配置的。为了更加方便，Laravel的本地数据库配置默认设置为使用这个数据库。

想通过你主机上的Navicat或Sequel Pro连接MySQL或Postgres，你应该使用端口33060（MySQL）或54320（Postgres）来连接“127.0.0.1”。这两个数据库的用户名和密码都是“homestead” / “secret”。

> **注意：** 当从主机连接数据库时，你应该只使用非标准的端口。在你的Laravel配置文件中，你将使用默认的3306和5432端口，因为Laravel运行在虚拟机当中。

### 添加额外站点

一旦你的Homestead环境被分配并运行，你可能想为Laravel应用程序添加额外的Nginx站点。在一个Homestead环境中，你可以按意愿运行尽可能多的Laravel应用程序。有两种方法可以做到这一点。首先，你可以简单的添加站点到“Homestead.yaml”文件里，先对箱子执行“vagrant destroy”命令，然后再执行“vagrant provision”命令。

或者，你可以使用Homestead环境里的“serve”脚本。想使用“serve”脚本，先SSH到Homestead环境并运行下面的命令：

	serve domain.app /home/vagrant/Code/path/to/public/directory

> **注意：** 在执行“serve”命令后，不要忘记添加新站点到你机器的“hosts”文件里！

<a name="ports"></a>
## 端口

下面的端口被转发到你的Homestead环境:

- **SSH:** 2222 -> 转发到 22
- **HTTP:** 8000 -> 转发到 80
- **MySQL:** 33060 -> 转发到 3306
- **Postgres:** 54320 -> 转发到 5432