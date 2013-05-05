#DEDE CMS
**title:** DEDE CMS手册  
**tags:** cms手册,phpcms


## 全局使用

### arclist|文档列表
* __标签名称：__ arclist
* __标记简介：__ 织梦常用标记，也称为自由列表标记，其中imglist、imginfolist、specart、coolart、autolist都是由该标记所定义的不同属性延伸出来的别名标记。
* __功能说明：__ 获取指定文档列表
* __适用范围：__ 全局使用
* __基本语法：__ 
    >     {dede:arclist  flag='h' typeid='' row='' col='' titlelen='' infolen='' imgwidth='' imgheight='' listtype='' orderby='' keyword='' limit='0,1'}
          <a href='[field:arcurl/]'>[field:title/]</a>
          {/dede:arclist}

* __相关函数：__ 文件\include\taglib\arclist.lib.php第7行
    >     function lib_arclist(&$ctag,&$refObj)

* __参数说明：__ 
    * col='' 分多少列显示（默认为单列），5.3版中本属性可以通过多种方式进行多行显示   
      如果col='1'要多列显示的可用div+css实现   
      以下为通过div+css实现多列的示例：

      >      <style type=text/css>
             div{width:400px;float:left;}
             </style>
             {dede:arclist row='10' titlelen='24' orderby='pubdate' idlist='' col='2'}
                [field:textlink/]([field:pubdate function=MyDate('m-d',@me)/])<br/>
             {/dede:arclist}

      当col>1将使用原来的table多列方式显示

    * row='10' 返回文档列表总数
    * typeid='' 栏目ID,在列表模板和档案模板中一般不需要指定，在首页模板中允许用","分开表示多个栏目；
    * getall='1' 在没有指定这属性的情况下,在栏目页、文章页模板,不会获取以","分开的多个栏目的下级子类
    * titlelen = '30' 标题长度 等同于titlelength
    * infolen='160' 表示内容简介长度 等同于infolength（只控制infos，并不控制description的内容）
    * imgwidth='120' 缩略图宽度
    * imgheight='90' 缩略图高度
    * listtype='all' 栏目类型 image含有缩略图 commend推荐
    * orderby='sortrank' 文档排序方式
        * § orderby='hot' 或 orderby='click' 表示按点击数排列
        * § orderby='sortrank' 或 orderby='pubdate' 按出版时间排列
        * § orderby='near'
        * § orderby=='lastpost' 按最后评论时间
        * § orderby=='scores' 按得分排序
        * § orderby='id' 按文章ID排序
        * § orderby='rand' 随机获得指定条件的文档列表
    * keyword='' 含有指定关键字的文档列表，多个关键字用","分
    * innertext = '' 单条记录样式
    * aid='' 指定文档ID
    * idlist ='' 提取特定文档（文档ID）
    * channelid 频道ID
    * limit='起始ID,记录数'  （起始ID从0开始）表示限定的记录范围（如：limit='1,2'  表示从ID为1的记录开始，取2条记录）
    * flag = 'h' 自定义属性值：头条[h]推荐[c]图片[p]幻灯[f]滚动[s]跳转[j]图文[a]加粗[b]
    * noflag = '' 同flag，但这里是表示不包含这些属性
    * orderway='desc' 值为 desc 或 asc ，指定排序方式是降序还是顺向排序，默认为降序
    * subday='天数' 表示在多少天以内的文档

用arclist调用于附加表字段的方法：

要获取附加表内容，必须符合两个条件  
1. 指定 channelid 属性
2. 指定要获得的字段 addfields='字段1,字段'

如：
>    {dede:arclist addfields='filetype,language,softtype' row='8' channelid='3'}
     [field:textlink /] - [field:softtype /]<br />
     {/dede:arclist}

* __底层模板字段：__
ID(同 id),typeid,sortrank,flag,ismake,channel,arcrank,click,money,title,shorttitle,color,writer, source,litpic(同picname),pubdate,senddate,mid, lastpost,scores,goodpost,badpost,notpost,description(同infos),filename, image, imglink, fulltitle, textlink, typelink,plusurl, memberurl, templeturl,stime(pubdate 的"0000-00-00"格式)

其中：

textlink = `<a href='arcurl'>title</a>`     
typelink = `<a href='typeurl'>typename</a>`           
imglink = `<a href='arcurl'><img src='picname' border='0' width='imgwidth' height='imgheight'></a>`        
image = `<img src='picname' border='0' width='imgwidth' height='imgheight' alt=’titile’>`        

字段调用方法：[field:varname/]

如：

>    {dede:arclist infolen='100'}
     [field:textlink/]
     <br>
     [field:infos/]
     <br>
     {/dede:arclist}
     
注：底层模板里的Field实现也是织梦标记的一种形式，因此支持使用PHP语法，Function扩展等功能。    
如： 给当天发布的内容加上 (new) 标志

>     [field:senddate runphp='yes'] 
      $ntime = time();
      $oneday = 3600 * 24;
      if(($ntime - @me)<$oneday) @me = "<font color='red'>(new)</font>";
      else @me = "";
      [/field:senddate]

### arclistsg|独立单表模型列表
* __标签名称：__ arclistsg
* __功能说明：__ 单表独立模型的文档列表调用标记 
* __适用范围：__ 全局使用 
* __基本语法：__
    >     {dede:arclistsg  flag='h' typeid='' row='' col='' titlelen='' orderway='' keyword='' limit='0,1'} 
          <a href='[field:arcurl/]'>[field:title/]</a> 
          {/dede:arclistsg}
 
* __相关函数：__ 文件\include\taglib\arclistsg.lib.php第2行
    >     function lib_arclistsg(&$ctag,&$refObj)
 
* __参数说明：__
    * col='' 分多少列显示（默认为单列），5.3版中本属性无效，要多列显示的可用div+css实现
    * row='10' 返回文档列表总数
    * typeid='' 栏目ID,在列表模板和档案模板中一般不需要指定，在封面模板中允许用","分开表示多个栏目；
    * titlelen = '30' 标题长度 等同于titlelength
    * orderwey='desc'或'asc' 排序方向
    * keyword= 含有指定关键字的文档列表，多个关键字用","分
    * innertext = "[field:title/]" 单条记录样式(innertext是放在标签之间的代码)
    * arcid='' 指定文档ID
    * idlist ='' 提取特定文档（文档ID）
    * channelid = '' 频道ID
    * limit='' 起始,结束 表示限定的记录范围（如：limit='1,2'）
    * flag = 'h' 自定义属性值：头条[h]推荐[c]图片[p]幻灯[f]滚动[s]跳转[j]图文[a]加粗[b]
    * subday='天数' 表示在多少天以内的文档

* __底层模板字段：__     
ID(同 id),typeid, channel, arcrank, mid, click, title, senddate, flag, litpic(同picname), userip, lastpost, scores, goodpost, badpost, textlink

* __使用实例：__    
这个标签用于调用单表模型的内容，在V5.3系统中加入了单表模型的概念，脱离了以前的主从表的数据表关联结构，一般我们在添加内容模型的时候会选择是否为单表模型还是主从表模型。
我们可以进入系统后台[核心]-[内容模型管理]，查看系统现有的内容模型，在系统默认环境下，只有分类信息是单表模型，一般添加单表模型，模型的id号都小于-1，即分类信息模型id号为-8。

在使用这个标签的时候，使用方法同arclist大致相同，我们来通过以下的标签内容来调用分类信息的内容：

>     {dede:arclistsg channelid='-8' limit='0,2'} 
      <a href='[field:arcurl/]'>[field:title/]</a> <br/>
      {/dede:arclistsg}
      
这个标签就是调用分类信息的内容，从id为0开始调用2条记录，在系统后台的[全局标签测试]中显示


### channel|频道标签
* __标签名称：__ channel 
* __标记简介：__ 织梦常用标记，通常用于网站顶部以获取站点栏目信息，方便网站会员分类浏览整站信息 
* __功能说明：__ 用于获取栏目列表 
* __适用范围：__ 全局使用 
* __基本语法：__ 
    >      {dede:channel type='top' row='8' currentstyle="<li><a href='~typelink~' class='thisclass'>~typename~</a> </li>"} 
            <li><a href='[field:typelink/]'>[field:typename/]</a> </li> 
           {/dede:channel}
       
* __相关函数：__ 文件\include\taglib\channel.lib.php第2行
    >     function lib_channel(&$ctag,&$refObj)

* __参数说明：__
    * typeid = '0' 栏目ID
    * reid = '0' 上级栏目ID
    * row = '100' 调用栏目数
    * col = '1' 分多少列显示（默认为单列）
    * type = 'son | sun' son表示下级栏目,self表示同级栏目,top顶级栏目
    * currentstyle = '' 应用样式

* __底层模板字段：__ 
    ID(同 id),typeid, typelink, typename, typeurl,typedir(仅表示栏目的网址)     
    例：
    >     {dede:channel type='top'} <a href='[field:typelink /]'>[field:typename/]</a> {/dede:channel}
 
    注：在没有指定typeid的情况下，type标记与模板的环境有关，如，模板生成到栏目一，那么type='son'就表示栏目一的所有子类    
    使用实例：这个标签是全局常用标记，主要用于显示页面的栏目分类，我们可以查看默认模板\templets\default\head.htm中的相关代码：
    >     {dede:channel type='top' currentstyle="<li class='thisclass'><a href='~typelink~'>~typename~</a> </li>"}
          <li><a href='[field:typeurl/]'>[field:typename/]</a></li>
          {/dede:channel}

    这里的栏目可以通过后台进行设置，栏目显示顺序按照排序的高低进行，我们可以在系统后台[核心]-[栏目管理]中进行栏目设置：   
    我们查看这个文章封面\templets\default\index_article.htm的模板标签
    >     <ul>
          {dede:channel type='son' currentstyle="<li class='thisclass'><a href='~typelink~'><span>~typename~</span></a> </li>"}
          <li><a href='[field:typeurl/]'><span>[field:typename/]</span></a></li>
          {/dede:channel}
          </ul>
    这里就使用了 type='son'这个属性用来显示子栏目。

很多用户希望顶部导航连接便于SEO，需要去除超链接中的“index.html”，我们可以在这里对标签进行一个修改：
>     {dede:channel type='top'} 
      <li><a href='[field:typeurl function='str_replace("index.html","",@me)'/]'>[field:typename/]</a></li> 
      {/dede:channel} 

### channelartlist|频道文档
* __标签名称：__ channelartlist 
* __标记简介：__ 
* __功能说明：__ 获取当前频道的下级栏目的内容列表标签 
* __适用范围：__ 全局使用 
* __基本语法：__     
    >      {dede:channelartlist row=6}
           <dl>
           <dt><a href='{dede:field name='typeurl'/}'>{dede:field name='typename'/}</a></dt>
           <dd>
           {dede:arclist titlelen='42' row='10'}    <ul class='autod'> 
               <li><a href="[field:arcurl /]">[field:title /]</a></li>
                <li>([field:pubdate function="MyDate('m-d',@me)"/])</li>
              </ul>
          {/dede:arclist}
          </dl>
          {/dede:channelartlist}
          
* __相关函数：__ 文件\include\taglib\channelartlist.lib.php第6行
    >     function lib_channelartlist(&$ctag,&$refObj)
    
* __参数说明：__
    typeid = '0' 频道ID,多个请用","分隔 row = '20' 获取的栏目返回值 其他说明： 除了宏标记外，channelArtlist 是唯一一个可以直接嵌套其它标记的标记，不过仅限于嵌套 {dede:type}{/dede:type} 和 {dede:arclist}{/dede:arclist} 两个标记。 
* __底层模板变量：__ 
    包含{dede:type}{/dede:type}及{dede:arclist}{/dede:arclist}下面所有底层模板变量。  
* __使用实例：__
    这个标签是系统中不多的一个支持嵌套的标签，这个标签通常使用在首页（含封面首页），用于输出一组栏目内容列表，我们可以看到默认模板首页：   
    这部分的内容就是通过这个标签进行显示的，我们可以查看首页模板\templets\default\index.htm120行 
    
    >     {dede:channelartlist}
         <dl class="tbox">
         <dt>
           <strong><a href="{dede:field name='typeurl'/}">{dede:field name='typename'/}</a></strong>
           <span class="more"><a href="{dede:field name='typeurl'/}">更多...</a></span>
         </dt>
         <dd>
         <ul class="d1 ico3">
         {dede:arclist titlelen='60' row='8'}
         <li>
           <span class="date">[field:pubdate function="MyDate('m-d',@me)"/]</span>
           <a href="[field:arcurl /]">[field:title /]</a>
         </li>
         {/dede:arclist}
         </ul>
         </dd>
         </dl>
         {/dede:channelartlist}

### feedback|会员评论内容
* __标签名称：__ feedback 
* __功能说明：__ 用于调用最新评论 
* __适用范围：__ 全局使用 
* __基本语法：__ 
    >     {dede:feedback}
         <ul>
          <li class='fbtitle'>[field:username function="(@me=='guest' ? '游客' : @me)"/] 对 [field:title/] 的评论：</li>
          <li class='fbmsg'> <a href="plus/feedback.php?aid=[field:aid/]" class='fbmsg'>[field:msg /]</a></li>
         </ul>
        {/dede:feedback}
        
* __相关函数：__  文件\include\taglib\feedback.lib.php第15行
    >     function lib_feedback(&$ctag,&$refObj)
* __参数说明：__ 
    * row='12' 调用评论条数   
    * titlelen='24' 标题长度   
    * infolen='100' 评论长度  
* __使用实例：__
    这个标签主要调用系统的会员评论信息，我们在默认模板首页可以查看到相关的内容：  
    我们可以查看首页的模板\templets\default\index.htm，在大约151行有如下代码：
    
    >      {dede:feedback row='5' titlelen='24' infolen='80'}
           <li> <small><a href="#" class="username">[field:username function="(@me=='guest' ? '游客' : @me)"/]</a> 评论 <a href="[field:global.cfg_phpurl/]/feedback.php?aid=[field:aid/]" class="title">[field:title/]</a></small>
           <p>[field:msg/]</p>
           </li>
           {/dede:feedback}
          
这里调用的是会员最新评论的内容，当然我们也可以在系统后台的[核心]-[评论管理]中对现有的评论进行审核和编辑.

在[系统]-[系统基本参数] 的”互动设置“中也有评论的相关设置选项，可以控制会员评论。


### field|常用变量

### fieldlist|变量列表

### hotwords|热门关键词

### hotwords|热门关键词

- - - - - - - - - - - - - - - - -
## `article_*.htm` 页面模板

### adminname|责任编辑
* __标签名称：__ adminname
* __功能说明：__ 获得责任编辑名称
* __适用范围：__ 仅内容模板 article_*.htm
* __基本语法：__  

>      {dede:adminname /}

* __相关函数：__ 文件\include\taglib\adminname.lib.php第7行 
   
>      function lib_adminname(&$ctag,&$refObj)


- - - - - - - - - - - - - - - - -
## `list_*.htm`  列表模板
