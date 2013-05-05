
* 项目名称：**Pro Git**
* 项目地址：[https://github.com/progit/progit](https://github.com/progit/progit "progit github仓库地址")
* 项目首页：[http://progit.org/](http://progit.org/book/zh "progit 中文官方")
* 仓库地址：[git://github.com/progit/progit.git][progit]

复制仓库：  
>git clone [git://github.com/progit/progit.git][progit]

在根目录下建立 [pro2html.sh][pro2html]

    # bluid html from markdown
    # it's for https://github.com/progit/progit
    # git://github.com/progit/progit.git
    # install markdown
    # 2011-06-10
    # put the file in / of object

	if [ $# -ne 1 ]; then
		read -p 'put a language shortname：' lang
	else
		lang=$1
	fi
	out=progit_$lang.html
	echo $lang
	echo $out
	touch $out
	echo '&lt;!doctype html&gt;&lt;html&gt;&lt;head&gt;&lt;meta http-equiv=content-type content=&quot;text/html; charset=utf8&quot;&gt;
	&lt;/head&gt;&lt;style type=&quot;text/css&quot;&gt;&lt;!--body{margin:0 20px;font-size:14px;}pre{margin:1em 0;font-size:13px;background-color:#eee;border:1px solid #ddd;padding:5px;line-height:1.5em;color:#444;overflow:auto;border-radius:3px;}code{background:#eee;}h1{text-align:center;margin-top:30px;color:green}h2{color:#6EA2F8}h3{color:#EE7D2F}.page{border:2px solid #333;background:#333;}img{max-width:600px;display:block;text-align:center;border-radius:5px;}--&gt;&lt;/style&gt;&lt;body&gt;' &gt;&gt; $out
	for i in `ls $lang/`
		do
		list=`find $lang/$i -iname &quot;*.markdown&quot;`
		#html=$lang/${i/\//\.html}
		html=$lang/$i.html
		markdown -v -o 'html4' $list -f $html
		cat $html &gt;&gt; $out
		rm -f $html
		echo '&lt;div class=&quot;page&quot;&gt;&lt;/div&gt;' &gt;&gt; $out
		echo &quot;get $list; markdown2html add to $out&quot;
	done
	echo '&lt;/body&gt;&lt;/html&gt;' &gt;&gt; $out;


使用

>./[pro2html.sh][pro2html] zh

来生成[ pro git 中文页面 progit_zh.html][progit_zh]就可以查看了。  
或者下载源码： [pro2html.sh][pro2html]

[progit]:git://github.com/progit/progit.git "progit github仓库"
[progit_zh]:http://opengit.org/markdown/progit/progit_zh.html "pro git 中文手册"
[pro2html]:http://opengit.org/download/source/pro2html.sh "转换markdown到html"
