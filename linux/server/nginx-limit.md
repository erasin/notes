# limit


<http://nginx.org/en/docs/http/ngx_http_limit_conn_module.html>


```conf
http {
    limit_conn_zone $binary_remote_addr zone=addr:10m;

    ...

    server {

        ...

        location /download/ {
            limit_conn addr 1;
        }
```


