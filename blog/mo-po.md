

**一、直接修改.mo文件**。 事实上，在<a href="http://www.ninthday.net/tag/linux" target="_blank">Linux</a>世界里有软件可以直接打开.mo文件以供编辑，这款软件叫——**Virtaal**，它是一款图形化的翻译工具。可以直接从源里安装该软件： Fedora： 
>   sudo yum install virtaal -y

**二、先反编译成.po文件。** 这要使用到GNU Gettext下的两个软件：**msgunfmt**和**msgfmt**。首先也是先安装需要用到的软件： Fedora： 
    sudo yum install gettext -y

    msgunfmt ./*.mo -o ./out.po

即可反编译生成.po文件，然后用Gedit打开该.po文件进行编辑。编辑好后，再执行下面的语句进行重新打包： 

	msgfmt ./out.po -o ./out.mo

安装opedit编辑po文件，保存时默认编译mo文件

Windows下在“命令提示符”界面操作步骤一样。只是把程序名后面加上".exe"。
