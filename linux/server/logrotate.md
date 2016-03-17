# 系统日志分割

logrotate  分割日志处理 

	cd /etc/logrotate.d/ 

php-fpm 

	/var/log/php-fpm/*log {
		missingok
		notifempty
		sharedscripts
		delaycompress
		postrotate
		/bin/kill -SIGUSR1 `cat /var/run/php-fpm/php-fpm.pid 2>/dev/null` 2>/dev/null || true
		endscript
	}

nginx 

	/home/logs/*log {
		daily
		rotate 10
		missingok
		notifempty
		compress
		sharedscripts
		postrotate
			/bin/kill -USR1 $(cat /var/run/nginx.pid 2>/dev/null) 2>/dev/null || :
		endscript
	}

