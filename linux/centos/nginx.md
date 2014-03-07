#nginx 安装在Centos中

**OS:** CentOS

## 源文件：

<http://wiki.nginx.org/Install>

To add nginx yum repository, create a file named /etc/yum.repos.d/nginx.repo and paste one of the configurations below:

### CentOS:

    [nginx]
    name=nginx repo
    baseurl=http://nginx.org/packages/centos/$releasever/$basearch/
    gpgcheck=0
    enabled=1

### RHEL:

    [nginx]
    name=nginx repo
    baseurl=http://nginx.org/packages/rhel/$releasever/$basearch/
    gpgcheck=0
    enabled=1

Due to differences between how CentOS, RHEL, and Scientific Linux populate the $releasever variable, it is necessary to manually replace $releasever with either "5" (for 5.x) or "6" (for 6.x), depending upon your OS version.


### Ubuntu 10.04:

    deb http://nginx.org/packages/ubuntu/ lucid nginx
    deb-src http://nginx.org/packages/ubuntu/ lucid nginx

### Debian 6:

    deb http://nginx.org/packages/debian/ squeeze nginx
    deb-src http://nginx.org/packages/debian/ squeeze nginx

### Ubuntu PPA

This PPA is maintained by volunteers and is not distributed by nginx.org. It has some additional compiled-in modules and may be more fitting for your environment.

You can get the latest stable version of Nginx from the Nginx PPA on Launchpad: You will need to have root privileges to perform the following commands.

For Ubuntu 10.04 and newer:

    sudo -s
    nginx=stable # use nginx=development for latest development version
    add-apt-repository ppa:nginx/$nginx
    apt-get update 
    apt-get install nginx

If you get an error about add-apt-repository not existing, you will want to install python-software-properties. For other Debian/Ubuntu based distributions, you can try the lucid variant of the PPA which is the most likely to work on older package sets.

    sudo -s
    nginx=stable # use nginx=development for latest development version
    echo "deb http://ppa.launchpad.net/nginx/$nginx/ubuntu lucid main" > /etc/apt/sources.list.d/nginx-$nginx-lucid.list
    apt-key adv --keyserver keyserver.ubuntu.com --recv-keys C300EE8C
    apt-get update 
    apt-get install nginx

创建脚本 vi /etc/init.d/nginxd
    
    #!/bin/bash
    #
    # chkconfig: - 85 15
    # description: Nginx is a World Wide Web server.
    # processname: nginx

    nginx=/usr/local/nginx/sbin/nginx
    conf=/usr/local/nginx/conf/nginx.conf
    case $1 in
        start)
            echo -n "Starting Nginx"
            $nginx -c $conf
            echo " done"
        ;;
        stop)
            echo -n "Stopping Nginx"
            killall -9 nginx
            echo " done"
        ;;
        test)
            $nginx -t -c $conf
        ;;
        reload)
            echo -n "Reloading Nginx"
            ps auxww | grep nginx | grep master | awk '{print $2}' | xargs kill -HUP
            echo " done"
        ;;
        restart)
            $0 stop
            $0 start
        ;;
        show)
            ps -aux|grep nginx
        ;;
        *)
            echo -n "Usage: $0 {start|restart|reload|stop|test|show}"
        ;;
    esac

设置权限与开机启动

    chmod 755 /etc/init.d/nginxd
    chkconfig nginxd on
