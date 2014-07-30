# linux下nginx 优化



## linux下php-fpm 优化


## error

502 error or connection 110  (有可能为mysql连接失败导致或者 selinux)

延长时间检查错误

    fastcgi_connect_timeout 600;
    fastcgi_send_timeout 600;
    fastcgi_read_timeout 600;
    fastcgi_buffer_size 256k;
    fastcgi_buffers 16 256k;
    fastcgi_busy_buffers_size 512k;
    fastcgi_temp_file_write_size 512k;

<http://java-er.com/blog/nginx-php-fpm-speed/>
<http://www.yangjia.org/2012/01/06/nginx%E4%B8%8Ephp-fpm%E4%BC%98%E5%8C%96.html>
<http://blog.haohtml.com/archives/11162>


