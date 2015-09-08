

pm = dynamic
pm.max_children = 8192
pm.start_servers = 128
pm.min_spare_servers = 128
pm.max_spare_servers = 1024
pm.max_requests = 500000


Nginx 是非阻塞IO & IO复用模型，通过操作系统提供的类似 epoll 的功能，可以在一个线程里处理多个客户端的请求。
Nginx 的进程就是线程，即每个进程里只有一个线程，但这一个线程可以服务多个客户端。

PHP-FPM 是阻塞的单线程模型，pm.max_children 指定的是最大的进程数量，
pm.max_requests 指定的是每个进程处理多少个请求后重启(因为 PHP 偶尔会有内存泄漏，所以需要重启).
PHP-FPM 的每个进程也只有一个线程，但是一个进程同时只能服务一个客户端。


nginx是异步IO模型？为什么我觉得是非阻塞IO & IO复用模型（基本上事件库都是如此），
然后在特定机器上选择最优的IO复用机制（譬如在linux机器选择epoll，FreeBSD使用kqueue，默认select），
而且如果nginx配置的是以多进程方式运行（fork出n个worker_process），似乎只要强调“进程”的概念即可，再扯“线程” 容易混淆；最后向楼主补充一句：worker_connections设置数可以更大，但是要小于一个进程允许打开的最大描述符数，
在配置文件中可以使用worker_rlimit_nofile来突破默认的1024












http://www.ha97.com/4339.html
http://www.cnblogs.com/qq78292959/p/4034359.html
