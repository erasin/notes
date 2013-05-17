# yum 工具
**title:** Centos/Fedora下yum工具  

[Linux][1]系统维护中令管理员很头疼的就是软件包之间的依赖性了，往往是你要安装A软件，但是编译的时候告诉你X软件安装之前需要B软件，而当你安装Y软件的时 候，又告诉你需要Z库了，好不容易安装好Z库，发现版本还有问题等。 

由于历史原因，RPM软件包管理系统对软件之间的依存关系没有内部定义，造成安装RPM软件时经常出现令人无法理解的软件依赖问题。其实开源社区早就对这个问题尝试进行解决了，不同的发行版推出了各自的工具，比如Yellow Dog的YUM，Debian的APT等。开发这些工具的目的都是为了要解决安装R PM时的依赖性问题，而不是额外再建立一套安装模式。这些软件也被开源软件爱好者们逐渐移植到别的发行版上。目前，APT和YUM都可以运行在Red Hat系统上。目前yum是Red Hat/Fedora系统上默认安装的更新系统。 

## yum特点

*   自动解决包的倚赖性问题能更方便的添加/删除/更新RPM包 
*   便于管理大量系统的更新问题 
*   可以同时配置多个资源库(Repository) 
*   简洁的配置文件(/etc/yum.conf) 
*   保持与RPM数据库的一致性 
*   有一个比较详细的log，可以查看何时升级安装了什么软件包等 
*   使用方便 

yum是[CentOS][2]/Fedora系统自带的，因此它能使用CentOS官方的软件源，完成各种官方发布的各种升级。对于第三方软件源的支持，yum也不差，大多数支持apt的repository，也能支持yum 。 

##配置yum客户的更新源

在.repo文件中，配置语法是分段的，每一段配置一个软件仓库，配置语法如下： 

    name=Some name for this repository
    baseurl=url://server1/path/to/repository/
    url://server2/path/to/repository/
    url://server3/path/to/repository/
    mirrorlist=url://path/to/mirrorlist/repository/
    enabled=0/1
    gpgcheck=0/1
    gpgkey=A URL pointing to the ASCII-armoured GPG key file for the repository`</pre>`

其中： 

*   repositoryid ： 用于指定一个仓库 
*   name： 用于指定易读的仓库名称 
*   baseurl ： 用于指定本仓库的URL，可以是如下的几种类型： 
    *   http — 用于指定远程 HTTP 协议的源 
    *   ftp — 用于指定远程 FTP 协议的源 
    *   file — 用于本地镜像或 NFS 挂装文件系统 
*   mirrorlist ： 用于指定仓库的镜像站点 
*   enabled ： 用于指定是否使用本仓库，默认值为1，即可用 
*   gpgcheck ： 用于指定是否检查软件包的 GPG 签名 
*   gpgkey ： 用于指定GPG签名文件的URL 

在name baseurl中经常使用如下的变量： 

*   $releasever — 当前系统的版本号 
*   $basearch — 当前系统的平台架构 
*   文件中以“#”开头的行是注释行 
*   若指定mirrorlist，系统将从CentOS的镜像站点中选择离您最近的仓库 
*   并非所有的国内镜像都在CentOS的镜像站点列表中，所以我们可以直接使用baseurl直接指定离您最近的仓库 
*   baseurl 可以指定多个 UR L，系统会依次检查您列出的仓库，以便在某个服务器宕机时可以使用另外的服务器 
*   为了加快更新，在确保更新服务器及线路良好的情况下，在baseurl中只指定一个URL既可 

##设置网络更新源

下面是一个CentOS-Base.repo文件的实例，在此文件中没有设置mirrorlist ，使用baseurl只指定了一个URL。

    name=CentOS-$releasever - Base
    baseurl=http://centos.candishosting.com.cn/$releasever/os/$basearch/
    gpgcheck=1
    gpgkey=http://mirror.centos.org/centos/RPM-GPG-KEY-CentOS-5
    [updates]
    name=CentOS-$releasever - Updates
    baseurl=http://centos.candishosting.com.cn/$releasever/updates/$basearch/
    gpgcheck=1
    gpgkey=http://mirror.centos.org/centos/RPM-GPG-KEY-CentOS-5
    [addons]
    name=CentOS-$releasever - Addons
    baseurl=http://centos.candishosting.com.cn/$releasever/addons/$basearch/
    gpgcheck=1
    gpgkey=http://mirror.centos.org/centos/RPM-GPG-KEY-CentOS-5
    [extras]
    name=CentOS-$releasever - Extras
    baseurl=http://centos.candishosting.com.cn/$releasever/extras/$basearch/
    gpgcheck=1
    gpgkey=http://mirror.centos.org/centos/RPM-GPG-KEY-CentOS-5
    [centosplus]
    name=CentOS-$releasever - Plus
    baseurl=http://centos.candishosting.com.cn/$releasever/centosplus/$basearch/
    gpgcheck=1
    enabled=0
    gpgkey=http://mirror.centos.org/centos/RPM-GPG-KEY-CentOS-5

## 设置本地更新源

为了使用安装光盘作为更新源，可以修改CentOS-Media.repo 文件，下面是一个配置实例。 

    name=CentOS-$releasever - Media
    baseurl=file:///media/CentOS/
            file:///media/cdrom/
            file:///media/cdrecorder/
    gpgcheck=1
    enabled=1
    gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-beta

*   为了使用 c5-media仓库，需将CentO S-Base .repo文件中的base仓库使用enabled=0设置成不可用。 
*   若本地磁盘空间有足够空间，您也可以将安装光盘复制到本地磁盘的一个目录中 

## yum命令简介

下面是一些较常见的用法。 


 | 命令                            | 功能
 |---------------------------------|---------------------------------------------------------------
 | yum check-update                | 检查可更新的所有软件包
 | yum update                      | 下载更新系统已安装的所有软件包
 | yum upgrade                     | 大规模的版本升级,与yum update不同的是,连旧的淘汰的包也升级
 | yum install                     | 安装新软件包
 | yum update                      | 更新指定的软件包
 | yum remove                      | 卸载指定的软件包
 | yum groupinstall `<groupnames>` | 安装指定软件组中的软件包
 | yum groupupdate `<groupnames>`  | 更新指定软件组中的软件包
 | yum groupremove `<groupnames>`  | 卸载指定软件组中的软件包
 | yum grouplist                   | 查看系统中已经安装的和可用的软件组
 | yum list                        | 列出资源库中所有可以安装或更新以及已经安装的rpm包
 | yum list `<regex>`              | 列出资源库中与正则表达式匹配的可以安装或更新以及已经安装的rpm包
 | yum list available              | 列出资源库中所有可以安装的rpm包
 | yum list available `<regex>`    | 列出资源库中与正则表达式匹配的所有可以安装的rpm包
 | yum list updates                | 列出资源库中所有可以更新的rpm包
 | yum list updates `<regex>`      | 列出资源库中与正则表达式匹配的所有可以更新的rpm包
 | yum list installed              | 列出资源库中所有已经安装的rpm包
 | yum list installed `<regex>`    | 列出资源库中与正则表达式匹配的所有已经安装的rpm包
 | yum list extras                 | 列出已经安装的但是不包含在资源库中的rpm包
 | yum list extras `<regex>`       | 列出与正则表达式匹配的已经安装的但是不包含在资源库中的rpm包
 | yum list recent                 | 列出最近被添加到资源库中的软件包
 | yum search `<regex>`            | 检测所有可用的软件的名称、描述、概述和已列出的维护者，查找与[正则表达式][regex]匹配的值
 | yum provides `<regex>`          | 检测软件包中包含的文件以及软件提供的功能，查找与正则表达式匹配的值
 | yum clean headers               | 清除缓存中的rpm头文件
 | yum clean packages              | 清除缓存中rpm包文件
 | yum clean all                   | 清除缓存中的rpm头文件和包文件
 | yum deplist                     | 显示软件包的依赖信息

*   当第一次使用yum或yum资源库有更新时，yum会自动下载所有所需的headers放置于 /var/cache /yum 目录下，所需时间可能较长。 
*   还可以使用 yum info 命令列出包信息，yum info 可用的参数与 yum list 的相同。 
*   yum 命令还可以使用 -y 参数用于用 yes 回答命令运行时所提出的问题。

 [1]: /linux "Linux"
 [2]: /linux/centos "centos linux发行版本"
 [regex]: /wiki/regex "正则表达式"

 *[yum]:Centos/Fedora中仓库工具
 *[regex]:正则表达式

