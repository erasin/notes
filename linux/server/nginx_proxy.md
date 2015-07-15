#nginx 反向代理

```
upstream proxyName{
    server 192.168.1.4:8099;
}
```
location / {
    proxy_pass http://proxyName;
    proxy_redirect          default ;
    proxy.conf
    proxy_set_header        Host            $host;
    proxy_set_header        X-Real-IP       $remote_addr;
    proxy_set_header        X-Forwarded-For $proxy_add_x_forwarded_for;
    client_max_body_size    10m;
    client_body_buffer_size 128k;
    proxy_connect_timeout   90;
    proxy_send_timeout      90;
    proxy_read_timeout      90;
    proxy_buffers           32 4k;
}
