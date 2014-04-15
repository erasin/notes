# 绑定域名的问题

demo1

	upstream frontends {
	    server 127.0.0.1:8888;
	}

	server {
	    server_name golang.tc www.golang.tc;
	    location / {
	        proxy_pass_header Server;
	        proxy_set_header Host $http_host;
	        proxy_redirect off;
	        proxy_set_header X-Real-IP $remote_addr;
	        proxy_set_header X-Scheme $scheme;
	        proxy_pass http://frontends;
	    }
	    location /static {
	        root        /var/go/gopher;
	        expires     1d;
	        add_header  Cache-Control public;
	        access_log  off;
	    }
	}

demo2 

	server {
        listen       80;
        server_name  www.a.com;
        charset utf-8;
        access_log  /home/a.com.access.log  main;
        location / {
            proxy_pass http://127.0.0.1:80;
        }
    }

	server {
        listen       80;
        server_name  www.b.com;
        charset utf-8;
        access_log  /home/b.com.access.log  main;
        location / {
            proxy_pass http://127.0.0.1:81;
        }
    }