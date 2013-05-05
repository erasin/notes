#Markdown 中文基础教程文档#



本文来自于：[Markdown basics](http://daringfireball.net/projects/markdown/basics)  
使用转换工具：[ Php Markdown Dingus](http://michelf.com/projects/php-markdown/dingus/ "Markdown 转换html工具")  
使用markdown插件：[ Php Markdown ](http://michelf.com/projects/php-markdown/ "wordpress Markdown plugin 插件")  

## 段落 Paragraphs,换行 br 标题 Headers , 引用块 Blockquotes ##

一个**段落**(`<p>`)是一行活多行文字和一个空行来表示，并且段落开始位置没有空白缩进和空格。

**换行符**（`<br>`）在文字的末尾处使用双空格或多空格，并且下一行非空行。

对于标题(`<h1~2>`)，Markdown 提供了两种方式 *Setext* 和 *atx*  
**Setext-style样式** 使用平衡符号‘========’ 和‘-----------’放在标题文字的下方分别来表示标题1和标题2  
**atx-style** 则使用相应个数的‘`#`’（marks）来标记开始和结束标题文本来对应标题1到标题6. 

**引用块**(`<blockquote>`) email-style '`>`' angle brackets.
<!--more-->
Markdown:

     标题1级别
     ====================
     
     标题2界别
     ---------------------

      这只是一个普通的段落，下面是个空行

     #标题1#
     ##标题2##
     ###标题3###

     > 这是一个引用块blockquote.
     > 
     > 块中的第二个引用段落.
     >
     > ##也可以使用标题##


Output:

     <h1>标题1级别</h1>
     
     <h2>标题2界别</h2>
     
     <p>这只是一个普通的段落，下面是个空行</p>
     
     <h1>标题1</h1>
     
     <h2>标题2</h2>
     
     <h3>标题3</h3>
     
     <blockquote>
       <p>这是一个引用块blockquote.</p>
       
       <p>块中的第二个引用段落.</p>
       
       <h2>也可以使用标题</h2>
     </blockquote>




## Phrase Emphasis 加强语气 ##

Markdown 使用**星号(`*`)**来表示加强语气`<strong>`和斜体字`<em>`

Markdown:

     单星号斜体 *斜体字符*.
     或者单下划线 _一样的斜体字符_.
     
     使用双星号加粗 **强调语气**.
     或者 __使用双下划线加粗__.

Output:

     <p>Some of these words <em>斜体字符</em>.
     或者单下划线<em>一样的斜体字符</em>.</p>
     
     <p>使用双星号加粗 <strong>强调语气</strong>.
     或者 <strong>使用双下划线加粗</strong>.</p>
   


## Lists 列表 ##

**无序列表**使用 (`*`,`+`, 和 `-`) 后加上空格 来标记列表内容，例如:

     *   Candy.
     *   Gum.
     *   Booze.

和:

     +   Candy.
     +   Gum.
     +   Booze.

和:

     -   Candy.
     -   Gum.
     -   Booze.

结果都为无序列表:

     <ul>
     <li>Candy.</li>
     <li>Gum.</li>
     <li>Booze.</li>
     </ul>

使用**数字加英文标点加空格** 来标记**有序列表**的内容:

     1.  Red
     2.  Green
     3.  Blue

Output:

     <ol>
     <li>Red</li>
     <li>Green</li>
     <li>Blue</li>
     </ol>

如果一个列表内容包含多个段落的时候第二段落前使用**4格空格或者1个缩进**符即可。

     *   一个列表项目.
     
         包含多个段落.

     *   另外一个列表内容.

Output:

     <ul>
     <li><p>一个列表项目.</p>
     
     <p>包含多个段落.</p></li>
     <li><p>另外一个列表内容.</p></li>
     </ul>

多级表的列表表示方法和列表项目包含段落相同。

## Links 链接 ##

Markdown 为链接提供了2种方式输出 *inline 内嵌式* 和 *reference 参考式*

**内嵌样式**表示方法为 中括号标记显示的链接文本，中括号后紧跟括号包括的链接，如果链接有title属性，在链接后使用空格+"title属性"即可。  
例如：

     连接到[百度](http://www.baidu.com/)吧!  
     连接到[百度](http://www.baidu.com/ "百度是拥有中国市场最大的搜索引擎")吧!  

输出样式为:

     <p>连接到<a href="http://example.com/">百度</a>吧!<br/>
     连接到<a href="http://example.com/" title="百度是拥有中国市场最大的搜索引擎">百度</a>吧!</p>

**参考样式**，我觉得称为引用样式更好些，一般应用于当多个不同位置地址相同的链接。  
使用的方法为,制作一个链接辞典列表，之后使用中括号引用。  
Markdown：

     搜索的结果上 [Google][1] 要比 [Yahoo][2] 和 [MSN][3] 好的多，不过在中国，[百度][baidu]更加稳定一些.

     [1]:http://google.com/         "Google"
     [2]:http://search.yahoo.com/  "Yahoo Search"
     [3]:http://search.msn.com/     "MSN Search"
     [baidu]:http://www.baidu.com/     "百度引擎"

输出:

     搜索的结果上 <a href="http://google.com/"title="Google">Google</a>
     than from <a href="http://search.yahoo.com/"title="Yahoo Search">Yahoo</a> 
     和 <a href="http://search.msn.com/" title="MSN Search">MSN</a> 好的多，
     不过在中国，<a href="http://www.baidu.com/" title="百度引擎">百度</a>更加稳定一些.

引用辞典中id可以包含数字，英文和空格，但是在使用包含空格中不对其敏感。

## Images 图片 ##

图片的使用基本上和链接相似，在中括号前添加了**叹号(`!`)**。

内联样式:

     ![alt 替代文本](/path/to/img.jpg "Title 标题文本")

参考样式:

     ![alt 替代文本][id]

     [id]: /path/to/img.jpg "标题文本"

上面两种输出的结果都为

     <img src="/path/to/img.jpg" alt="alt 替代文本" title="Title 标题文本" />



## Code 代码 ##

在常规的段落中使用 **反引号** 标记包括文字，可以将任意符号转换为实体字符比如 (`# & >`);这样更为方便了html代码的输出。

     我强力的建议在markdown中不要使用任何 `<blink>` 标记.

     显示实体字符 `&mdash;`
     使用十进制的实体 `&#8212;`.

输出:

     <p>我强力的建议在markdown中不要使用任何<code>&lt;blink&gt;</code> tags.</p>
     
     <p>显示实体字符<code>&amp;mdash;</code> 
     使用十进制的实体 <code>&amp;#8212;</code>.</p>

如果使用**预格式化标签**`<pre>`则是在每一行文字前添加**4格空格**或者1一个缩进符，默认会将实体字符自动转化。

Markdown:

     If you want your page to validate under XHTML 1.0 Strict,
     you've got to put paragraph tags in your blockquotes:

         <blockquote>
              <p>For example.</p>
         </blockquote>

Output:

     <p>If you want your page to validate under XHTML 1.0 Strict,
     you've got to put paragraph tags in your blockquotes:</p>
     
     <pre><code>&lt;blockquote&gt;
         &lt;p&gt;For example.&lt;/p&gt;
     &lt;/blockquote&gt;
     </code></pre>

###其他###
使用 星号空格星号空格星号空格 `( * * * )`多个可以输出 `<hr>`
