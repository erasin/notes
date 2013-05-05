#nginx 安装在Centos中

**OS:** CentOS

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
