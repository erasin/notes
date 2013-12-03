#31种选择器的应用

原文<http://my.oschina.net/u/1403186/blog/179641>

1. *

    * { margin: 0; padding: 0; }

星号选择器用于选取页面中的所有元素，可用于快速清除所有元素的 margin 与 padding，但最好只在测试的时候使用，而不要正式用在 CSS 文件中，否则会大大加重浏览器负担。此外，星号选择器也可以给父层的所有子元素设定样式，重复一遍，尽量少用这种方式：

    #container * { border: 1px solid black; }

兼容 IE6+

2. #X

    #container { width: 960px; margin: auto; }

id 选择器，最常见的选择器用法之一，不可重复使用。

兼容 IE6+

3. .X

    .error { color: red; }

class 选择器，也是最常见的选择器用法之一，与 id 选择器不同的是 class 选择器可同时选取多个元素，而 id 选择器只能给一个独一无二的元素设定样式。

兼容 IE6+

4. X Y

    li a { text-decoration: none; }

后代选择器 (descendant selector)，选取 X 元素内的所有 Y 元素，比如上面这段代码将选取 li 标签内的所有链接。

兼容 IE6+

5. X

    a { color: red; } ul { margin-left: 0; }

标签选择器 (type selector)，用于选取 HTML 标签 (tag)。

兼容 IE6+

6. X:visited and X:link

    a:link { color: red; } a:visted { color: purple; }

:link 伪类选择器 (pseudo class selector) 用于选取所有未点击过的链接，而 :visited 则用于选取所有已访问过的链接。

兼容 IE6+

7. X + Y

    ul + p { color: red; }

临近选择器 (adjacent selector)，选取紧邻在 X 元素后面出现的第一个元素，比如上面这段代码将选取 ul 元素后出现的第一个元素，也就是 p 元素。

兼容性 IE6+

8. X > Y

    div#container > ul { border: 1px solid black; }

在第 4 条中，后代选择器 X Y 选取父层 X 内的所有 Y 元素；子选择器 X > Y 则只选取直接出现在父层 X 内的 Y 元素。比如下面的 HTML 结构中，#container > ul 选取直接出现在 div#container 内的 ul 元素，不包含嵌套在 li 内的 ul 元素：

    <div id="container"> 
        <ul> 
            <li> List Item
             <ul> 
                <li> Child </li> 
            </ul> 
        </li> 
        <li> List Item </li> 
        <li> List Item </li> 
        <li> List Item </li> 
    </ul> </div>

兼容 IE7+

9. X ~ Y

    ul ~ p { color: red; }

同样也是临近选择器，前面第 7 条 X + Y 选取紧邻在 X 后出现的第一个元素，而 X ~ Y 将选取 X 元素后出现的所有同级元素。上面这段代码将选取 ul 元素后出现的所有同级 p 元素，而不是像 ul + p 这样选取第一个出现的 p 元素。

兼容 IE7+

10. X[title]

    a[title] { color: green; }

属性选择器 (attributes selector)，根据元素使用的属性进一步缩小选取范围，上面这段代码将选取所有使用了 title 属性的链接，或者 a[title="title content"]{color:green} 再进一步缩小选取范围。

兼容 IE7+

11. X[href="foo"]

    a[href="http://net.tutsplus.com"] { color: #1f6053; /* nettuts green */ }

上面这段代码将选取所有跳转到 http://net.tutsplus.com 的链接，这些链接将显示为绿色，其他链接不受影响。

只是这种方式很严格不能相差一个字符，下面将会逐一介绍更灵活的用法。

兼容 IE7+

12. X[href*="nettuts"]

    a[href*="tuts"] { color: #1f6053; /* nettuts green */ }

* 表示只要属性值中包含双引号内的内容就满足选取要求，这段代码将选取跳转到 nettuts.com，net.tutsplus.com，或者 tutsplus.com 等链接。

兼容 IE7+

13. X[href^="http"]

    a[href^="http"] { background: url(path/to/external/icon.png) no-repeat; padding-left: 10px; }

^ 表示只要属性值以双引号内的内容开头就满足选取要求，这段代码也常用来给页面中所有外部接设定样式。

兼容 IE7+

14. X[href$=".jpg"]

    a[href$=".jpg"] { color: red; }

$ 表示只要属性值以双引号内的内容结尾就满足选取要求，这段代码将选取所有跳转到 jpg 图片的链接。

兼容 IE7+

15. X[data-*="foo"]

上面第 14 条提到了如何选取所有跳转到 jpg 图片的链接，若要选取跳转到图片的所有链接可以用下面的方法：

    a[href$=".jpg"], a[href$=".jpeg"], a[href$=".png"], a[href$=".gif"] { color: red; }

或者，先给图片链接添加 data- 属性（注：HTML5 Custom Data Attributes）

    <a href="path/to/image.jpg" data-filetype="image"> Image Link </a>

然后再通过属性选择器选取：

    a[data-filetype="image"] { color: red; }

兼容 IE7+

16. X[foo~="bar"]

    a[data-info~="external"] { color: red; } a[data-info~="image"] { border: 1px solid black; }

如果属性值中有用空格分隔的一连串属性值，~ 可以选取其中一个属性值，比如：

    <a href="path/to/image.jpg" data-info="external image"> Click Me, Fool </a>

借助 ~ 选取包含 external 或者 image 属性值的元素：

    /* Target data-info attr that contains the value "external" */ 
    a[data-info~="external"] { color: red; } 
    /* And which contain the value "image" */ 
    a[data-info~="image"] { border: 1px solid black; }

兼容 IE7+

17. X:checked

    input[type=radio]:checked { border: 1px solid black; }

:checked 伪类选择器用于选取所有标记为 checked 的元素，比如单选框 (radio button) 或复选框 (checkbox)。

兼容 IE9+

18. X:after

:before 与 :after 是两个令人兴奋的伪类选择器，几乎每天都有人发明出一些新用法，这里简单介绍一下如何用它清除浮动：

    .clearfix:after { content: ""; display: block; clear: both; visibility: hidden; font-size: 0; height: 0; }
    .clearfix { *display: inline-block; _height: 1%; }

这种方式通过 :after 在元素后面添加一块区域，然后将其隐藏，可以弥补 overflow: hidden; 的缺陷。

根据 CSS3 选择器标准，理论上伪类选择器应该使用双冒号，但实际上浏览器也支持单冒号的形式，所以可以继续使用单冒号。

兼容 IE8+

19. X:hover

    div:hover { background: #e3e3e3; }

最常用的伪类选择器，不多解释了，只是需要注意 IE6 不支持将 :hover 作用于除 a 链接外的其他元素。

    a:hover { border-bottom: 1px solid black; }

另外提醒一点：border-bottom: 1px solid black; 的效果要比 text-decoration: underline; 好看一些。

兼容 IE6+ （在 IE6 中 :hover 只能作用于链接）

20. X:not(selector)

    div:not(#container) { color: blue; }

:not 伪类选择器有时会起到很重要的作用，假设现在要选取除 #contaienr 外的所有 div 元素，就可以用上面的代码实现。

兼容 IE9+

21. X::pseudoElement

    p::first-line { font-weight: bold; font-size: 1.2em; }

通过伪元素（使用双冒号 ::）可以给元素的某一部分设定样式，比如第一行、或者第一个字母。需要注意的是，这只对块级元素 (block level elements) 生效。

提示：伪元素 (pseudo element) 使用双冒号 ::

选取段落的第一个字母

    p::first-letter { float: left; font-size: 2em; font-weight: bold; font-family: cursive; padding-right: 2px; }

这段代码将选取页面中所有 p 元素，然后再选取其中的第一个字母。

选取段落的第一行

    p::first-line { font-weight: bold; font-size: 1.2em; }

与上面的例子类似，通过 ::first-line 选取页面的第一行。

为了兼容 CSS1 与 CSS2 中的伪元素（比如 :first-line, :first-letter, :before 以及 :after），浏览器接受单冒号与双冒号两种格式，但对于 CSS3 中最新引入的伪元素，必须使用双冒号。

兼容 IE6+

22. X:nth-child(n)

    li:nth-child(3) { color: red; }

:nth-child(n) 用于选取 stack 中的某一个元素，只接受整数作参数（参数从 1 开始计数），如果你想选取第二个 li 元素，只需这样写 li:nth-child(2)。

也可以设定可变的参数，比如 li:nth-child(4n) 将选取第 4, 8 , 12… 个元素（4*n, n=1, n++）。

兼容 IE9+

23. X:nth-last-child(n)

    li:nth-last-child(2) { color: red; }

除了正序（从上往下）选择，也可以使用 :nth-last-child(n) 倒序（从下往上）选择第几个元素，其他用法与第 22 条完全一样。

兼容 IE9+

24. X:nth-of-type(n)

    ul:nth-of-type(3) { border: 1px solid black; }

:nth-of-type(n) 的作用不是选取子元素 (child element)，而是选取同类元素 (type of element)。想象一下 HTML 文件中包含 5 个 ul 元素，现在要选取第三个，只需使用上面的代码，而不用再单独这个 ul 添加 id。

关于 :nth-child 与 :nth-of-type 的区别，具体请查看 CSS-Tricks 网站的解释，简单来说，如果父层内只包含一种元素（比如都是 p 元素）那两种用法是等效的，如果父层包含多种元素（比如 p 元素与同级的其他元素），需要选取第几个 p 元素时应该用 :nth-of-type。

兼容 IE9+

25. X:nth-last-of-type(n)

    ul:nth-last-of-type(3) { border: 1px solid black; }

与第 24 条用法相同，倒序选取同类元素。

兼容 IE9+

26. X:first-child

    ul li:first-child { border-top: none; }

选取父层内的第一个子元素。

兼容 IE7+

27. X:last-child

    ul > li:last-child { color: green; }

与第 26 条用法相同，区别在于 :last-child 选取父层元素内的最后一个子元素。

:first-child 与 :last-child 通常用来清除边框 (border)，比如 <ul></ul> 内每个 <li></li> 都使用了 border-top 与 border-bottom 边框，结果是，第一个元素的上方与最后一个元素的下方会是单边框效果。这种情况可以用 :first-child 与 :last-child 清除上下的边框，而不用给第一个元素添加id="first" 或者给最后一个元素添加 id="last"。

兼容 IE9+

28. X:only-child

    div p:only-child { color: red; }

这个伪类选择器不常用，它可以选取包含唯一指定子元素的父层。比如上面的代码中将选取下面第一个 div 元素，而不是第二个 div 中的 p 元素。

    <div>
        <p> My paragraph here. </p>
    </div> 
    <div> 
        <p> Two paragraphs total. </p> 
        <p> Two paragraphs total. </p> 
    </div>

兼容 IE9+

29. X:only-of-type

    li:only-of-type { font-weight: bold; }

这个选择器会选取某个元素，并且这个元素在其父层内没有其他同级同类元素（不一定是唯一元素）。比如，要选取所有只包含一个 li 元素的 ul 元素该怎么做呢？如果使用 ul li 将选取所有 li 元素，应该使用 only-of-type。

兼容 IE9+

30. X:first-of-type

first-of-type 伪类可以选取某种元素的第一个同类元素。

为了更好地理解它的用法，现在思考一下如何在下面的 HTML 结构中选取 List Item 2 ？

    <div> 
        <p> My paragraph here. </p>
        <ul> 
           <li> List Item 1 </li>
           <li> List Item 2 </li>
       </ul> 
       <ul> 
         <li> List Item 3 </li> 
         <li> List Item 4 </li> 
       </ul> 
    </div>

方法一

    ul:first-of-type > li:nth-child(2) { font-weight: bold; }

这段代码的意思是：首先选取第一个 ul 元素；然后选取其中的所有直接子元素，也就是 li；最后选取第二个子元素。

方法二

    p + ul li:last-child { font-weight: bold; }

找到 p 元素后第一个出现的 ul 元素，然后选取其中的最后一个子元素。

方法三

    ul:first-of-type li:nth-last-child(1) { font-weight: bold; }

找到第一个 ul 元素，然后从上往下选取第一个子元素。

兼容 IE9+

31. 伪类选择器叠用

有些伪类选择器或者伪元素是可以叠用的，例如：

    #post p:nth-of-type(2):first-letter { float: left; margin: 0 5px 0 1em; width: 1em; height: 1em; font-size: 2em; }