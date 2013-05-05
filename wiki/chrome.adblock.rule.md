#AdBlock Plus 广告过滤规则介绍
**title:**AdBlock Plus 广告过滤规则介绍  
**tag:**chrome,adblock    
**info**本章节描述的过滤规则属性，对偶尔才写过滤规则的用户来说足矣。

当前的 Adblock Plus 版本允许您通过许多不同的方法来优化过滤规则。本文档就是告诉您如何做。

**声明：**这里给出的过滤规则只是示例，不一定能直接使用。

##基本过滤规则
通常您定义得最琐碎的过滤规则当然是您想阻挡 banner 地址。但是这些地址在您每次打开页面的时候经常会变。例如：可能是 http://example.com/ads/banner123.gif 其中 123 是个随机数字。这里阻挡完整的图片地址是没用的，您需要更通用的过滤规则——类似 `http://example.com/ads/banner*.gif`。 或者更甚 `http://example.com/ads/*`。

_注：_ 不要过多地使用通配符。过滤规则 `http://example.com/*` 虽然可以阻挡所有的 banner， 但也会阻挡 example.com 下其它一些您想看的内容。

##定义例外规则
有时您可能会发现某个过滤规则平时挡广告挡得很好，但在某些情况下，会阻挡一些不该挡的内容。您不想移除这条过滤规则，但也不希望它阻挡不该挡的内容。

这就是例外规则的好处——它们允许您定义过滤规则不被使用的情况。例如，您不满意过滤规则 adv 阻挡了 http://example.com/advice.html，您就可以定义一条例外规则 `@@advice` 。 例外规则和过滤规则没什么两样，您可以使用通配符或正则表达式。您只需在规则前添加 @@ 来声明这是一个例外规则。

例外规则不止可以处理这些。如果一条例外规则以 http:// 或 https://（也可以在前面加上管线符号（|））开始，这会使所有的页面都是例外。例如：如果您的规则是 @@|http://example.com 您浏览 example.com 的页面时，Adblock Plus 就被禁用了，这将不会阻挡任何东西。

##匹配网址开头/结尾
通常 Adblock Plus 处理过滤规则时，会自己假设在过滤规则的开头与结尾都有一个通配符，例如，过滤规则 ad 和 `*ad*` 是一样。 正常情况下这没什么问题，但有时您可能想要定义可以匹配以网址开头或结尾的过滤规则。例如，您想要阻挡所有的 Flash，但如果您添加过滤规则 swf 地址 http://example.com/swf/index.html 同样也将被阻挡。

这个问题的解决方法：使用管线符号（|）来表示地址的最前端或最末端。例如这条过滤规则 swf| 会阻挡 http://example.com/annoyingflash.swf 但不会阻挡 http://example.com/swf/index.html。这条过滤规则 |http://baddomain.example/ 会阻挡 http://baddomain.example/banner.gif 但不会阻挡 http://gooddomain.example/analyze?http://baddomain.example。

有时您想阻挡 http://example.com/banner.gif 以及 https://example.com/banner.gif 和 http://www.example.com/banner.gif。这时只需在过滤规则的域名前面加上两个管线符号（||）：||example.com/banner.gif 将会阻挡上面的地址而不会阻挡 http://badexample.com/banner.gif 或者 http://gooddomain.example/analyze?http://example.com/banner.gif（需要 Adblock Plus 1.1 或更高版本）。

##标记分隔符
通常您需要接受过滤规则的任何分隔符。例如，您可能写这样一个规则阻挡 http://example.com/ 和 http://example.com:8000/ 但不能阻挡 http://example.com.ar/。在这里，符号(^)用作一个分隔符。 http://example.com^（需要 Adblock Plus 1.1 或更高版本）。

分隔符可以是除了字母、数字或者 `_ - . %` 之外的任何字符。 这个地址的结尾也是作为一个分隔符，下面的例子中所有的分隔符以红色标记出：http://example.com:8000/foo.bar?a=12&b=%D1%82%D0%B5%D1%81%D1%82。所以这个地址可以通过这些过滤规则过滤 ^example.com^ 或 ^%D1%82%D0%B5%D1%81%D1%82^ 或 ^foo.bar^ 。

##注释
任何以感叹号 (!) 开始的规则，都被视为注释。在过滤规则的列表中，仍然会显示这些规则，但会用灰色的字来显示，而不是黑色。Adblock Plus 在判断规则时，会忽略这些注释，所以我们可以写下任何我们想写的东西。您可以在一条规则上面写下这条规则是做什么用的。也可以在过滤列表的上方写上作者信息（大多数过滤列表的作者已经这样做了）。

##进阶功能
本章节描述的特性通常只有高级用户和维护过滤列表的作者才会看。普通用户可跳过。

##指定过滤规则选项
Adblock Plus 允许您指定某些选项来改变某条规则的行为。您列举这些选项的时候将它们放在美元符号 ($) 后面并用逗号 (,) 分割这些选项，放在过滤规则的最后面，例如：

>`*/ads/*$script,match-case`

这里的 */ads/* 是真实的过滤规则 script 和 match-case 是其指定的选项。下面是目前支持的选项：

* 类型选项：判定过滤规则（或例外规则）过滤元素的类型。过滤规则可以指定多个类型选项来过滤指定的元素类型。可以指定的类型包括：
	* script —— 外部脚本，由 HTML script 标签加载
	* image —— 正常图片，通常由 HTML 的 img 标签所载入
	* background —— 背景图片，通常用 CSS 指定
	* stylesheet —— 外部 CSS 样式文件
	* object —— 由浏览器插件处理的内容，例如 Flash 或 Java
	* xbl —— XBL 绑定（通常由 -moz-binding CSS 属性加载）
	* ping —— link pings
	* xmlhttprequest —— XMLHttpRequest 对象
	* object-subrequest —— 插件的请求，比如Flash
	* dtd —— 通过 XML 文档加载的 DTD 文件
	* subdocument —— 内嵌的页面，通常通过 HTML 的框架方式内嵌
	* document —— 网页本身（只适用于 例外规则 ）
	* elemhide —— 只适用于例外规则，类似于document 但是只禁用页面上的隐藏规则而不是所有规则（需要Adblock Plus 1.2 或更高版本）
	* other —— 其他不在上面的类型的请求
* 反转类型选项：指定过滤规则不应用的元素类型。可以指定的类型选项： ~script, ~image, ~background, ~stylesheet, ~object, ~xbl, ~ping, ~xmlhttprequest, ~object-subrequest, ~dtd, ~subdocument, ~document, ~elemhide, ~other
* third-party/first-party 请求限制：如果指定了 third-party 选项， 则过滤规则只适用于来源与当前正在浏览的页面的不同的请求。类似地，~third-party 适用于来源与当前浏览页面相同的请求。
* 域名限定：选项 domain=example.com 指过滤规则只适用于 "example.com" 下的页面 。多个域名可以用 "|" 分隔： 过滤规则 domain=example.com|example.net 将只适用于 "example.com" 或 "example.net" 的页面。如果一个域名是前面有"~"，则该过滤规则不适用于这个域名的页面。例如： domain=~example.com 指过滤规则适用于除了 example.com 之外的任何域名的页面而 domain=example.com|~foo.example.com 限定了过滤规则适用于 "example.com" 但不包括 "foo.example.com" 。
* match-case —— 使过滤规则只适用于匹配地址，例如：过滤规则 */BannerAd.gif$match-case 会阻挡 http://example.com/BannerAd.gif 但不会阻挡 http://example.com/bannerad.gif。
* collapse —— 这个选项将覆盖全局“瓦解可过滤的元素”，并确过滤规则总是瓦解元素。类似地，~collapse 选项将确保过滤规则不瓦解元素 。
* donottrack —— 对有该选项的阻挡规则匹配到且有该选项的例外规则未匹配到的地址会发送一个 Do-Not-Track 头 (需要Adblock Plus 1.3.5 或更高版本)。 为了向后兼容，使用此选项时建议使用矛盾的组合类型选项，防止此规则在早期版本的 Adblock Plus 中阻挡任何东西： *$donottrack,image,~image

##使用正则表达式
如果您想更好地控制您的过滤规则，什么匹配，什么不匹配，您可以使用正则表达式。例如过滤规则 /banner\d+/ 会匹配 banner123 和 banner321 而不会匹配 banners。 您可以查看正则表达式的文档来学习如何写正则表达式。

注： 由于性能原因，建议尽可能避免使用正则表达式。

##元素隐藏
###基本规则
有时您可能会发现无法阻挡某些内嵌在网页中的文字广告。如果查看源码的话，可能发现类似这样的代码：

>	 <div class="textad">
	Cheapest tofu, only here and now!
	</div>
	<div id="sponsorad">
	Really cheap tofu, click here!
	</div>
	<textad>
	Only here you get the best tofu!
	</textad>

因为您必须下载页面的内容，所以您也必须下载这些广告。对于这种情况，您可以做的就是把这些广告藏起来，这样您就不会看到他们了。这也就是元素隐藏的意义所在。

上面代码中的第一则广告是在一个 class 属性为“textad”的 div 容器内。过滤规则 `##div.textad` 。 这里的 `##` 表明这是一条元素隐藏规则，剩下的就是定义需要隐藏元素的选择器，同样的，您可以通过他们的 id 属性来隐藏 `##div#sponsorad` 会隐藏第二个广告。您不需要指定元素的名称， 过滤规则 `##*#sponsorad` 同样也可以。您也可以仅指定要阻挡的元素名称来隐藏，例如：`##textad` 可以隐藏第三则广告。

在不查看页面源码的情况下，Element Hiding Helper 扩展 可以帮助选择正确的元素并写出相应的规则。基础的HTML知识还是很有用的。

_注：_元素隐藏规则与普通过滤规则的工作方式有很大的差别。元素隐藏规则不支持通配符。

###限定在特定域名的规则
通常您只想要隐藏特定网站的特定广告，而不希望规则会作用于其他网站。例如，过滤规则 `##*.sponsor` 可能会把某些网站的有效代码也隐藏了。但如果你把它写成 `example.com##*.sponsor` 就只会在 http://example.com/ 和 http://something.example.com/ 生效了，而不是 http://example.org/。 你也可以指定多个域名——只要用逗号（,）分隔即可：`domain1.example,domain2.example,domain3.example##*.sponsor` 。

如果在域名之前有 "~"，该过滤规则不适用于这个域名的页面（需要 AdBlock Plus 1.1或更高版本）。例如， `~example.com##*.sponsor` 将适用于除了 "example.com" 之外的域名，`example.com,~foo.example.com##*.sponsor` 适用于 "example.com" 但不适用于 "foo.example.com" 子域名。

_注：_由于元素隐藏实现方式的关系，您只可以将隐藏规则限定在完整的域名。您不能使用网址的其他部份，也不可用 domain 代替 domain.example,domain.test 。

_注：_ 限定域名的元素隐藏规则也可用来隐藏浏览器的使用界面。例如，过滤规则 `browser##menuitem#javascriptConsole` 会隐藏 Firefox 工具菜单中的 JavaScript 控制台。

###属性选择符
一些广告隐藏起来并不容易——它们广告不仅没有 id 也没有 class 属性。您可以使用其他属性来隐藏，例如 `##`table[width="80%"] 可以隐藏 width 属性值为 80% 的表格元素。 如果您不想指定属性的完整值，`##div[title*="adv"]` 会隐藏所有 title 属性包含 adv 字符的 div 元素。您还可以检查属性的开始和结束字符，例如 `##div[title^="adv"][title$="ert"]` 会隐藏 titile 属性以 adv 开始并且以 ert 结束的 div 元素。正如您所见，你可以使用多个条件 —— table[width="80%"][bgcolor="white"] 会匹配到 width 属性为 80%、bgcolor 属性为 white 的表格元素。

###高级选择符
通常情况下，Firefox 支持的 CSS 选择器都可用于元素隐藏。例如：下面的过滤规则会隐藏 class 的属性为 adheader 的 div 元素相邻的元素： `##div.adheader + *`。完整的 CSS 列表请查看 W3C CSS 规范 （Firefox 目前并没有支持所有的选择器）。

_注：_这个功能只是给高级用户使用的，您可以很舒服地通过 CSS 选择符去使用它。Adblock Plus 无法检查您添加的选择器的语法是否正确，如果您使用无效的 CSS 语法，可能会破坏其它已有的有效过滤规则。建议使用 JavaScript 控制台检查是否有 CSS 错误。

###简单元素隐藏语法
Adblock Plus 支持简单元素隐藏语法（例如： #div(id=foo)）只是为了向后兼容性。使用这个语法是不好的，CSS 选择器才是首选。对这个语法的支持可能在以后的某个时间就不支持了。
