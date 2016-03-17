#you

* 增加 php-fpm 进程池

nginx 负载均衡

```
upstream backend {
	server unix:/var/run/php-fpm/php-fpm.sock weight=100 max_fails=5 fail_timeout=5;
	server unix:/var/run/php-fpm/php-fpm2.sock weight=100 max_fails=5 fail_timeout=5;
	server unix:/var/run/php-fpm/php-fpm3.sock weight=100 max_fails=5 fail_timeout=5;
	server unix:/var/run/php-fpm/php-fpm4.sock weight=100 max_fails=5 fail_timeout=5;
}
```
