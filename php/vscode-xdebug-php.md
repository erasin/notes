# vscode php xdebug

下载插件 php debug

配置追加 `php.validate.executablePath` 配置php的位置。

开启远程debug, xdebug 的默认远程端口是 9000，php-fpm 也是 9000 ， 所以在这里使用了 10001;

```ext-xdebug.conf
[xdebug]
zend_extension="/usr/local/opt/php56-xdebug/xdebug.so"
xdebug.remote-enable = 1
xdebug.remote-autostart = 1
xdebug.remote_port = 10001
```

