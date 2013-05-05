# 在windows下nginx与php的启动脚本
**tags:** nginx,fastcgi,windows,启动脚本

## 运行 nginx 开始 脚本
就是用笨一点的方法吧

依次创建及脚本 nginx.bat fastcgi.bat start.vbs stop.bat

### nginx.conf

    echo Starting nginx ...
    c:/nginx/nginx.exe

### fastcgi.bat

    @echo off
    REM Windows 下无效
    REM set PHP_FCGI_CHILDREN=5

    REM 每个进程处理的最大请求数，或设置为 Windows 环境变量
    set PHP_FCGI_MAX_REQUESTS=1000
     
    echo Starting PHP FastCGI...
    c:/php/php-cgi.exe -b 127.0.0.1:9000 -c c:/php/php.ini

### start.vbs 运行脚本
 
    set ws=WScript.CreateObject("WScript.Shell") 
    ws.Run "fastcgi.bat",0 
    ws.Run "nginx.bat",0 

### stop.bat 停止

    @echo off
    echo Stopping nginx...
    taskkill /F /IM nginx.exe > nul
    echo Stopping PHP FastCGI...
    taskkill /F /IM php-cgi.exe > nul
    exit

## php fastcgi在 nginx.conf server 中设置

    #  写入到 server 
    location ~ \.php/?.*$
    {
        fastcgi_pass   127.0.0.1:9000;
        set $path_info "";
        set $real_script_name $fastcgi_script_name;
        if ($fastcgi_script_name ~ "^(.+?\.php)(/.+)$") {
            set $real_script_name $1;
            set $path_info $2;
        }
        fastcgi_param  SCRIPT_FILENAME $document_root$fastcgi_script_name;
        fastcgi_param SCRIPT_NAME $real_script_name;
        fastcgi_param PATH_INFO $path_info;
        fastcgi_index  index.php;
        include        fastcgi_params;
    }


