#ctags php

在用vim查看代码的时候，代码跳转是必不可少的，一般可以用ctags来生成tag文件供vim读取、跳转。但是，ctags默认设置对php支持的不是很好（主要是类的function和property支持的不好）。在网上找了半天，然后自己试验了几次，终于找到一个相对好用的设置，如下：

（以下命令可放到~/.bash_profile中）

	alias phptags='ctags --langmap=php:.engine.inc.module.theme.php  --php-kinds=cdf  --languages=php'

以下设置放到 ~/.ctags文件中

	$ cat ~/.ctags
	--regex-php=/^[ \t]*[(private|public|static)( \t)]*function[ \t]+([A-Za-z0-9_]+)[ \t]*\(/\1/f, function, functions/
	--regex-php=/^[ \t]*[(private|public|static)]+[ \t]+\$([A-Za-z0-9_]+)[ \t]*/\1/p, property, properties/
	--regex-php=/^[ \t]*(const)[ \t]+([A-Za-z0-9_]+)[ \t]*/\2/d, const, constants/

使用时，在代码目录中：

	$ phptags -R

就可以生成比较可用的tags文件了。

