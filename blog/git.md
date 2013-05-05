#Git



"**Git**  ----  *The stupid content tracker*, 傻瓜内容跟踪器。Linus 是这样给我们介绍 Git 的。

Git 是用于 Linux 内核开发的版本控制工具。与常用的版本控制工具 CVS, Subversion 等不同， 它采用了分布式版本库的方式，不必服务器端软件支持，使源代码的发布和交流极其方便。 Git 的速度很快，这对于诸如 Linux kernel 这样的大项目来说自然很重要。 Git 最为出色的是它的合并跟踪（merge tracing）能力。

实际上内核开发团队决定开始开发和使用 Git 来作为内核开发的版本控制系统的时候， 世界开源社群的反对声音不少，最大的理由是 Git 太艰涩难懂，从 Git 的内部工作机制来说，的确是这样。 但是随着开发的深入，Git 的正常使用都由一些友好的脚本命令来执行，使 Git 变得非常好用， 即使是用来管理我们自己的开发项目，Git 都是一个友好，有力的工具。 现在，越来越多的著名项目采用 Git 来管理项目开发，例如：wine, U-boot 等，详情看[http://www.kernel.org/git](http://www.kernel.org/git)

作为开源自由原教旨主义项目，Git 没有对版本库的浏览和修改做任何的权限限制。 它只适用于 Linux / Unix 平台，没有 Windows 版本，目前也没有这样的开发计划。

本文将以 Git 官方文档 Tutorial， core-tutorial 和 Everyday GIT 作为蓝本翻译整理，但是暂时去掉了对 Git 内部工作机制的阐述， 力求简明扼要，并加入了作者使用 Git 的过程中的一些心得体会，注意事项，以及更多的例子。 建议你最好通过你所使用的 Unix / Linux 发行版的安装包来安装 Git, 你可以在线浏览本文 ，也可以通过下面的命令来得到本文最新的版本库，并且通过后面的学习用 Git 作为工具参加到本文的创作中来。
>$ git-clone http://www.bitsun.com/git/gittutorcn.git

##git 中文相关教程##
**阅读**来自于 github上的开源书籍  《[Pro Git 中文版本](/progit/menu)》  
一个简单的Git教程来自于<a href="http://www.linuxsir.org/main/doc/git/gittutorcn.htm" target="_blank">http://www.linuxsir.org/main/doc/git/gittutorcn.htm</a>, 这只简单的介绍git的使用方法
"
