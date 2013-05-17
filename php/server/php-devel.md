# php-devel


phpize 是属于 php-devel 中的东西，主要是设定 php 外挂模块的一些设定

所以安装 php-devel 相关套件就会有 phpize 可以使用 (档案预设存放于 /usr/bin/phpize )

phpize 命令是用来准备 PHP 外挂模块的编译环境的。下面例子中，外挂模块的源程序位于 extname 目录中：

	# cd extname
	# phpize
	# ./configure (注一)
	# make
	# make install

成功的安装将建立 extname.so 并放置于 PHP 的外挂模块目录中 (预设存放于 /usr/lib/php/modules/ 内) 。
需要调整 php.ini，加入 extension=extname.so 这一行之后才能使用此外挂模块。

如在执行　./configure 时出现　not find –with-php-config 时，

可重下以下指令，因 –with-php-config　预设在　/usr/bin/php-config　可找到
`./configure –with-php-config=/usr/bin/php-config`
需要调整 php.ini，加入 `extension=extname.so` 这一行之后才能使用此扩展库。


