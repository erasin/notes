#列出所有服务器地址，nginx 自动均衡分发请求到各个服务器。  
upstream frontends {    
    ip_hash;  
    server 192.168.199.1:8088;
    server 192.168.199.2:8089;
}
server {
    listen      80; 
    server_name mydomain.com www.mydomain.com;
    location / {
        proxy_pass_header Server;
        proxy_set_header Host $http_host;
        proxy_redirect off;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Scheme $scheme;
        proxy_pass http://frontends;
    }
     
    #静态资源交由nginx管理
    location /static {
        root        /var/www/mydomain/web;
        expires     1d;
        add_header  Cache-Control public;
        access_log  off;
    }
}

//this host ip 192.168.199.1
func main() {
    ...
    http.ListenAndServe(":8088", nil)
    os.Exit(0)
}
 
...
//other
//this host ip 192.168.199.2
func main() {
    ...
    http.ListenAndServe(":8089", nil)
    os.Exit(0)
}