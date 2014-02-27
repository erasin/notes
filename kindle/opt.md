OPF文档是epub电子书的核心文件，且是一个标准的XML文件，依据OPF规范，主要由五个部分组成：

1、

    <metadata>,元数据信息，由两个子元素组成：

    <dc-metadata>，其元素构成采用dubline core(DC)的15项核心元素，包括：
    <title>:题名
    <creator>：责任者
    <subject>：主题词或关键词
    <description>：内容描述
    <contributor>：贡献者或其它次要责任者
    <date>：日期
    <type>：类型
    <format>：格式
    <identifier>：标识符
    <source>：来源
    <language>：语种
    <relation>：相关信息
    <coverage>：履盖范围
    <rights>：权限描述
    <x-metadata>，即扩展元素。如果有些信息在上述元素中无法描述，则在此元素中进行扩展。

2、

    <menifest>，文件列表，由于列出OEBPS文档及相关的文档，有一个子元素构成，

    <item id="" href="" media-type="">,该元素由三个属性构成：
    id:表示文件的ID号
    href：文件的相对路径
    media-type：文件的媒体类型
    例如：<item id="page01" href="page01.xhtml" media-type="application/xhtml+xml"/>
    <item id="img000"href="001.png"media-type="image/png" />
    <item id="img001"href="002.jpg"media-type="image/jpeg" />

3、

    <spine toc="ncx">，脊骨，其主要功能是提供书籍的线性阅读次序。由一个子元素构成：

    <itemref idref="">,由一个属性构成：
    idref:即参照menifest列出的ID
    例如：<itemref idref="page01"/>


4、

    <guide>,指南,依次列出电子书的特定页面, 例如封面、目录、序言等, 属性值指向文件保存地址。一般情况下，epub电子书可以不用该元素。

5、
    
    <tour>,导读。可以根据不同的读者水平或者阅读目的, 按一定次序, 选择电子书中的部分页面组成导读。一般情况下，epub电子书可以不用该元素。





## 1.初级使用者

使用须知：kindlegen.exe 是运行于命令行模式的，没有UI界面，因此，必须知道如何进入命令行模式，或者知道如何使用批处理文件。

最简单的转换：kindlegen.exe *.html/.htm/.epub（*是您想转换的文件名）

也就是说，可以直接将网页、epub格式的文件转换成.mobi格式电子书。

如果是网页，则网页代码中title将作为kindle中显示的书名，不会有封面、书目章节跳转功能。

如果是epub，则视epub制作的情况而定。

此种方法最简单，但不能把kindlegen的功能发挥出来，适合于小文件。


## 2.高级使用者

kindlegen 还可转换标准的电子书出版格式.opf文件，注意这不是一个单一的文件，而是一组文件，比较复杂，大家慢慢看。

文件结构：

    /-----------         # root
    |-- ch1.html         # 内容文件
    |-- ch2........
    |-- toc.html         # 书目文件 审定超级链接到内容文件
    |-- toc.ncx          # 书本xml文件 标签和段落
    |-- book_name.opf    # 书的元数据 xml格式
    |-- Cover.png        # 封面 600x800 16 
    |-- css.css          # style
    |-- build.sh         # bulid shell


## 3.关于Kindle Previewer

kindlegen 还有一个姊妹软件叫Kindle Previewer，是Amazon为配合其标准电子书出版软件Kindlegen而开发的，可以模拟显示最终在kindle上显示的效果，包括 Kindle3、Kindle DX、IPad、IPhone等等设备模式，最新版本1.6。比较大，暂时没法放上来，让大家眼馋了。大家如果不闲麻烦，可以将制作好的.mobi文件拷 贝到自己心爱的Kindle上观看，如果觉得麻烦，就在Amazon官网注册，下载（好像也很麻烦）
接下来，详细介绍每个文件的内容。


## 4. 制作

ZangDiMiMa.html

    [/font]
    [font=tahoma]<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
    <html xmlns="http://www.w3.org/1999/xhtml">
    <head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <link href="css.css" rel="stylesheet"type="text/css"/>
    <title>藏地密码</title>
    </head>
    <body>
    <h1 id="1">第一部 一部关于西藏的百科全书式小说</h1>
    <mbp:pagebreak/>
    <img src="c1.jpg" width="499px" height="500px" />
    <mbp:pagebreak/>
    <p>这是一个西藏已经开放为全世界的旅游胜地却依旧守口如瓶的秘密。多年之后身在美国宾夕法尼亚州的藏獒专家卓木强巴突 然收到一个陌生人送来的信封，信封里装着两张照片，照片上惊现的远古神兽，促使卓木强巴及导师、世界犬类学专家方新教授亲赴西藏。他们在调查过程中震惊地 发现，照片上的动物竟然和帕巴拉神庙有关……不久之后一支由特种兵、考古学家、生物学家、密修高手等各色人物组成的神秘科考队悄悄从西藏出发开始了一场穿 越全球生死禁地的探险之旅，他们要追寻藏传佛教千年隐秘历史的真相……西藏，到底向我们隐瞒了什么？！</p>
    <h2 id="1.0">序</h2>
    <p>比藏獒更加凶狠的动物是什么？</p>
    ...此处省略4MB文字...
    <p>第七十九章 太可怕的真相</p>
    <p>第八十章 千年前的故事</p>
    <p>第八十一章 以一张照片结束</p>
    </body>
    </html>




<mbp:pagebreak/> 指的是分页，遇到这个标签，kindle将从新的一页开始显示下面的内容。
通过设置id吧标签，让其它文件来超链接。方便书目章节跳转。




toc.html书目
[/size][/font]
[font=tahoma][size=12px]<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.1//EN" "http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd">[/size][/font]
[font=tahoma][size=12px]<html xmlns="http://www.w3.org/1999/xhtml">[/size][/font]
[font=tahoma][size=12px]<head>[/size][/font]
[font=tahoma][size=12px]<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />[/size][/font]
[font=tahoma][size=12px]<link href="css.css" rel="stylesheet" type="text/css" />[/size][/font]
[font=tahoma][size=12px]<title>Table of Contents</title>[/size][/font]
[font=tahoma][size=12px]</head>[/size][/font]
[font=tahoma][size=12px]<body>[/size][/font]
[font=tahoma][size=12px]<h1>书目</h1>[/size][/font]
[font=tahoma][size=12px]<hr width="80%" />[/size][/font]
[font=tahoma][size=12px]<h2 style="page-break-before:avoid"><a href="ZangDiMiMa.html#1">第一部</a></h2>[/size][/font]
[font=tahoma][size=12px]<p style="text-indent:0"><a href="ZangDiMiMa.html#1.0">序</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#1.1">第一章 从一张照片说起</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#1.2">第二章 紫麒麟传说</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#1.3">第三章 巴桑的回忆</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#1.4">第四章 横穿可可西里</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#1.5">第五章 史前冰川惊魂记</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#1.6">第六章 笔记之谜</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#1.7">第七章 帕巴拉神庙是否存在？</a></p>[/size][/font]
[font=tahoma][size=12px]<h2 style="page-break-before:avoid"><a href="ZangDiMiMa.html#2">第二部</a></h2>[/size][/font]
[font=tahoma][size=12px]<p style="text-indent:0"><a href="ZangDiMiMa.html#2.1">第八章 出发！亚马逊丛林！</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#2.2">第九章 丛林危机</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#2.3">第十章 深陷原始部落库库尔族</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#2.4">第十一章 叹息丛林：探险家的坟墓</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#2.5">第十二章 洪荒：上帝之手</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#2.6">第十三章 我们被食人族绑架了！</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#2.7">第十四章 玛雅圣城</a></p>[/size][/font]
[font=tahoma][size=12px]        [/size][/font]
[font=tahoma][size=12px]<h2 style="page-break-before:avoid"><a href="ZangDiMiMa.html#3">第三部</a></h2>[/size][/font]
[font=tahoma][size=12px]<p style="text-indent:0"><a href="ZangDiMiMa.html#3.1">第十五章 玛雅迷宫</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#3.2">第十六章 水火地狱机关</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#3.3">第十七章 血池</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#3.4">第十八章 回到西藏</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#3.5">第十九章 玛雅：华夏文明的美洲变种</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#3.6">第二十章 西藏墨脱：最后的秘境</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#3.7">第二十一章 生命之门</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#3.8">第二十二章 地狱之门</a></p>[/size][/font]
[font=tahoma][size=12px]
[/size][/font]
[font=tahoma][size=12px]<h2 style="page-break-before:avoid"><a href="ZangDiMiMa.html#4">第四部</a></h2>[/size][/font]
[font=tahoma][size=12px]<p style="text-indent:0"><a href="ZangDiMiMa.html#4.1">第二十三章 高原雪狼</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#4.2">第二十四章 倒悬空寺</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#4.3">第二十五章 巨门之后</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#4.4">第二十六章 巨佛内部</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#4.5">第二十七章 终极血池</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#4.6">第二十八章 西藏古格</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#4.7">第二十九章 千年一战</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#4.8">第三十章 西藏密教</a></p>[/size][/font]
[font=tahoma][size=12px]
[/size][/font]
[font=tahoma][size=12px]<h2 style="page-break-before:avoid"><a href="ZangDiMiMa.html#5">第五部</a></h2>[/size][/font]
[font=tahoma][size=12px]<p style="text-indent:0"><a href="ZangDiMiMa.html#5.1">第三十一章 雪山仆从</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#5.2">第三十二章 紫麒麟猜想</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#5.3">第三十三章 绝没见过的狼</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#5.4">第三十四章 水晶宫</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#5.5">第三十五章 极南庙</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#5.6">第三十六章 死亡西风带</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#5.7">第三十七章 唐涛的日记</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#5.8">第三十八章 人生的宿命</a></p>[/size][/font]
[font=tahoma][size=12px]
[/size][/font]
[font=tahoma][size=12px]<h2 style="page-break-before:avoid"><a href="ZangDiMiMa.html#6">第六部</a></h2>[/size][/font]
[font=tahoma][size=12px]<p style="text-indent:0"><a href="ZangDiMiMa.html#6.1">第三十九章 希特勒秘闻</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#6.2">第四十章 德军进藏秘密史料</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#6.3">第四十一章 德军进藏秘密地图</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#6.4">第四十二章 希特勒第一次派人进藏之谜</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#6.5">第四十三章 希特勒第二次派人进藏之谜</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#6.6">第四十四章 香巴拉真身之谜</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#6.7">第四十五章冥河：西藏最神秘的河</a></p>[/size][/font]
[font=tahoma][size=12px]
[/size][/font]
[font=tahoma][size=12px]<h2 style="page-break-before:avoid"><a href="ZangDiMiMa.html#7">第七部</a></h2>[/size][/font]
[font=tahoma][size=12px]<p style="text-indent:0"><a href="ZangDiMiMa.html#7.1">第四十六章 大天轮经：藏密最高法典</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#7.2">第四十七章 向下朝香巴拉前进</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#7.3">第四十八章 藏地猜想：喜马拉雅海</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#7.4">第四十九章 再见十三圆桌骑士</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#7.5">第五十章 初入香巴拉，重返古生代</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#7.6">第五十一章 穿越香巴拉原始丛林</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#7.7">第五十二章 被遗忘的古藏戈巴族村</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#7.8">第五十三章 迷失香巴拉深处</a></p>[/size][/font]
[font=tahoma][size=12px]        [/size][/font]
[font=tahoma][size=12px]<h2 style="page-break-before:avoid"><a href="ZangDiMiMa.html#8">第八部</a></h2>[/size][/font]
[font=tahoma][size=12px]<p style="text-indent:0"><a href="ZangDiMiMa.html#8.1">第四十六章 大天轮经：藏密最高法典</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#8.2">第五十四章 喜马拉雅雪人之谜</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#8.3">第五十五章 喜马拉雅雪人，现在在哪儿？</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#8.4">第五十六章 蛊毒患者</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#8.5">第五十七章 穆族遗迹</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#8.6">第五十八章 奇迹之城雀母</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#8.7">第五十九章 神秘王国雅加</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#8.7">第六十章寻 找回归之路的密修者</a><br />[/size][/font]
[font=tahoma][size=12px]<a href="ZangDiMiMa.html#8.8">第六十一章 胡杨队长之死</a></p>[/size][/font]
[font=tahoma][size=12px]</body>[/size][/font]
[font=tahoma][size=12px]</html>[/size][/font]

[font=tahoma][size=12px]
复制代码





toc.ncx

[/size][/size][/font]
[font=tahoma]<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE ncx PUBLIC "-//NISO//DTD ncx 2005-1//EN" "http://www.daisy.org/z3986/2005/ncx-2005-1.dtd">
<ncx xmlns="http://www.daisy.org/z3986/2005/ncx/" version="2005-1" xml:lang="en-US">
<head>
<meta name="dtb:uid" content="BookId"/>
<meta name="dtb:depth" content="2"/>
<meta name="dtb:totalPageCount" content="0"/>
<meta name="dtb:maxPageNumber" content="0"/>
</head>
<docTitle><text>藏地密码</text></docTitle>
<docAuthor><text>何马</text></docAuthor>
        <navMap>
                <navPoint class="toc" id="toc" playOrder="1">
                        <navLabel>
                                <text>Table of Contents</text>
                        </navLabel>
                        <content src="toc.html"/>
                </navPoint>
                <navPoint class="chapter" id="chapter1" playOrder="2">
                  <navLabel>
                                <text>第一部</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#1"/>
          </navPoint>
                <navPoint class="chapter" id="chapter00" playOrder="3">
                  <navLabel>
                        <text>序</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#1.0"/>
                </navPoint>
                <navPoint class="chapter" id="chapter01" playOrder="4">
                  <navLabel>
                        <text>第一章 从一张照片说起</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#1.1"/>
                </navPoint>
                <navPoint class="chapter" id="chapter02" playOrder="5">
                  <navLabel>
                        <text>第二章 紫麒麟传说</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#1.2"/>
                </navPoint>
                <navPoint class="chapter" id="chapter03" playOrder="6">
                  <navLabel>
                        <text>第三章 巴桑的回忆</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#1.3"/>
                </navPoint>
                <navPoint class="chapter" id="chapter04" playOrder="7">
                  <navLabel>
                        <text>第四章 横穿可可西里</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#1.4"/>
                </navPoint>
                <navPoint class="chapter" id="chapter05" playOrder="8">
                  <navLabel>
                        <text>第五章 史前冰川惊魂记</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#1.5"/>
                </navPoint>
                <navPoint class="chapter" id="chapter06" playOrder="9">
                  <navLabel>
                        <text>第六章 笔记之谜</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#1.6"/>
                </navPoint>
                <navPoint class="chapter" id="chapter07" playOrder="10">
                  <navLabel>
                        <text>第七章 帕巴拉神庙是否存在？</text>
                  </navLabel>
                                <content src="ZangDiMiMa.html#1.7"/>
                </navPoint>

                <navPoint class="chapter" id="chapter2" playOrder="11">
                  <navLabel>
                                <text>第二部</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#2"/>
          </navPoint>
                <navPoint class="chapter" id="chapter08" playOrder="12">
                  <navLabel>
                        <text>第八章 出发！亚马逊丛林！</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#2.1"/>
                </navPoint>
                <navPoint class="chapter" id="chapter09" playOrder="13">
                  <navLabel>
                        <text>第九章 丛林危机</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#2.2"/>
                </navPoint>
                <navPoint class="chapter" id="chapter10" playOrder="14">
                  <navLabel>
                        <text>第十章 深陷原始部落库库尔族</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#2.3"/>
                </navPoint>
                <navPoint class="chapter" id="chapter11" playOrder="15">
                  <navLabel>
                        <text>第十一章 叹息丛林：探险家的坟墓</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#2.4"/>
                </navPoint>
                <navPoint class="chapter" id="chapter12" playOrder="16">
                  <navLabel>
                        <text>第十二章 洪荒：上帝之手</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#2.5"/>
                </navPoint>
                <navPoint class="chapter" id="chapter13" playOrder="17">
                  <navLabel>
                        <text>第十三章 我们被食人族绑架了！</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#2.6"/>
                </navPoint>
                <navPoint class="chapter" id="chapter14" playOrder="18">
                  <navLabel>
                        <text>第十四章 玛雅圣城</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#2.7"/>
                </navPoint>

                <navPoint class="chapter" id="chapter3" playOrder="19">
                  <navLabel>
                                <text>第三部</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#3"/>
          </navPoint>
                <navPoint class="chapter" id="chapter15" playOrder="20">
                  <navLabel>
                        <text>第十五章 玛雅迷宫</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#3.1"/>
                </navPoint>
                <navPoint class="chapter" id="chapter16" playOrder="21">
                  <navLabel>
                        <text>第十六章 水火地狱机关</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#3.2"/>
                </navPoint>
                <navPoint class="chapter" id="chapter17" playOrder="22">
                  <navLabel>
                        <text>第十七章 血池</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#3.3"/>
                </navPoint>
                <navPoint class="chapter" id="chapter18" playOrder="23">
                  <navLabel>
                        <text>第十八章 回到西藏</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#3.4"/>
                </navPoint>
                <navPoint class="chapter" id="chapter19" playOrder="24">
                  <navLabel>
                        <text>第十九章 玛雅：华夏文明的美洲变种</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#3.5"/>
                </navPoint>
                <navPoint class="chapter" id="chapter20" playOrder="25">
                  <navLabel>
                        <text>第二十章 西藏墨脱：最后的秘境</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#3.6"/>
                </navPoint>
                <navPoint class="chapter" id="chapter21" playOrder="26">
                  <navLabel>
                        <text>第二十一章 生命之门</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#3.7"/>
                </navPoint>
                <navPoint class="chapter" id="chapter22" playOrder="27">
                  <navLabel>
                        <text>第二十二章 地狱之门</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#3.8"/>
                </navPoint>

                <navPoint class="chapter" id="chapter4" playOrder="28">
                  <navLabel>
                                <text>第四部</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#4"/>
          </navPoint>
                <navPoint class="chapter" id="chapter23" playOrder="29">
                  <navLabel>
                        <text>第二十三章 高原雪狼</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#4.1"/>
                </navPoint>
                <navPoint class="chapter" id="chapter24" playOrder="30">
                  <navLabel>
                        <text>第二十四章 倒悬空寺</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#4.2"/>
                </navPoint>
                <navPoint class="chapter" id="chapter25" playOrder="31">
                  <navLabel>
                        <text>第二十五章 巨门之后</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#4.3"/>
                </navPoint>
                <navPoint class="chapter" id="chapter26" playOrder="32">
                  <navLabel>
                        <text>第二十六章 巨佛内部</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#4.4"/>
                </navPoint>
                <navPoint class="chapter" id="chapter27" playOrder="33">
                  <navLabel>
                        <text>第二十七章 终极血池</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#4.5"/>
                </navPoint>
                <navPoint class="chapter" id="chapter28" playOrder="34">
                  <navLabel>
                        <text>第二十八章 西藏古格</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#4.6"/>
                </navPoint>
                <navPoint class="chapter" id="chapter29" playOrder="35">
                  <navLabel>
                        <text>第二十九章 千年一战</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#4.7"/>
                </navPoint>
                <navPoint class="chapter" id="chapter30" playOrder="36">
                  <navLabel>
                        <text>第三十章 西藏密教</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#4.8"/>
                </navPoint>

                <navPoint class="chapter" id="chapter5" playOrder="37">
                  <navLabel>
                                <text>第五部</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#5"/>
          </navPoint>
                <navPoint class="chapter" id="chapter31" playOrder="38">
                  <navLabel>
                        <text>第三十一章 雪山仆从</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#5.1"/>
                </navPoint>
                <navPoint class="chapter" id="chapter32" playOrder="39">
                  <navLabel>
                        <text>第三十二章 紫麒麟猜想</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#5.2"/>
                </navPoint>
                <navPoint class="chapter" id="chapter33" playOrder="40">
                  <navLabel>
                        <text>第三十三章 绝没见过的狼</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#5.3"/>
                </navPoint>
                <navPoint class="chapter" id="chapter34" playOrder="41">
                  <navLabel>
                        <text>第三十四章 水晶宫</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#5.4"/>
                </navPoint>
                <navPoint class="chapter" id="chapter35" playOrder="42">
                  <navLabel>
                        <text>第三十五章 极南庙</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#5.5"/>
                </navPoint>
                <navPoint class="chapter" id="chapter36" playOrder="43">
                  <navLabel>
                        <text>第三十六章 死亡西风带</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#5.6"/>
                </navPoint>

...省略部分代码...                

                <navPoint class="chapter" id="chapter60" playOrder="69">
                  <navLabel>
                        <text>第六十章寻 找回归之路的密修者</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#8.7"/>
                </navPoint>
                <navPoint class="chapter" id="chapter61" playOrder="70">
                  <navLabel>
                        <text>第六十一章 胡杨队长之死</text>
                  </navLabel>
                  <content src="ZangDiMiMa.html#8.8"/>
                </navPoint>

        </navMap>
</ncx></ncx>
复制代码

这些代码符合XML要求，其中class可以为chapter，也可为section，即在章节内可再细分段落（如果有必要的话），要注意id是唯一的，playOrder是需要按顺序递增的。






ZangDiMiMa.opf

[/font]
[font=tahoma]<?xml version="1.0" encoding="utf-8"?>
<package unique-identifier="uid">
        <metadata>
                <dc-metadata xmlns:dc="http://purl.org/metadata/dublin_core" xmlns:oebpackage="http://openebook.org/namespaces/oeb-package/1.0/">
                        <dc:Title>藏地密码</dc:Title>
                        <dc:Language>zh</dc:Language>
                        <dc:Identifier id="uid">010B5D2ACA</dc:Identifier>
                        <dc:Creator>何马</dc:Creator>
                        <dc:Subject>小说</dc:Subject>
                </dc-metadata>
                <x-metadata>
                        <output encoding="utf-8"></output>
                        <EmbeddedCover>Cover.gif</EmbeddedCover>
                </x-metadata>
        </metadata>
        <manifest>
                <item id="My_Table_of_Contents" media-type="application/x-dtbncx+xml" href="toc.ncx"/>
                <item id="item0" media-type="application/xhtml+xml" href="toc.html"></item>
                <item id="item1" media-type="text/x-oeb1-document" href="ZangDiMiMa.html"></item>
        </manifest>
        <spine toc="My_Table_of_Contents">
                <itemref idref="item0"/>
                <itemref idref="item1"/>
        </spine>
        <tours></tours>
        <guide>
                <reference type="toc" title="Table of Contents" href="toc.html"></reference>
        </guide>
</package>
复制代码




在这个例子中，只有一个HTML文件，就是ZangDiMiMa.html，所以只需要



[/font]
[font=tahoma]<item id="item1" media-type="text/x-oeb1-document" href="ZangDiMiMa.html"></item>[/font]
[font=tahoma]
复制代码



，否则，需要把每一个要加入的文件都写入，如



[/font]
[font=tahoma]<item id="item2" media-type="text/x-oeb1-document" href="ZangDiMiMa2.html"></item><item id="item3" media-type="text/x-oeb1-document" href="ZangDiMiMa3.html"></item>[/font]
[font=tahoma]
[/font]
[font=tahoma]
[/font]
[font=tahoma]build.bat[/font]
[font=tahoma][code][/font]
[font=tahoma]kindlegen.exe ZangDiMiMa.opf -c1
PAUSE
EXIT
[/font]
[font=tahoma]
复制代码

简单的批处理文件，双击后运行Kindlegen，省得敲键盘。
其中-c1是默认的压缩等级，也可以不用。