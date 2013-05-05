#7项Web开发者需要了解的HTML5和CSS3新技术

Web 开发者需要经常更新他们的知识，学习新的技术，如果他们还想继续在 Web 开发领域混并混得还不错的话。下面将为你展示 7 项新的Web开发技术，作为一个Web开发人员，你需要了解、熟悉并学会的技术。

##CSS3 media queries
目前，大量的智能手机设备的涌现，同时各种不同尺寸屏幕的设备，如平板电脑之类的出现，对Web开发带来了前所未有的挑战，如何让 Web 页面能适应各种尺寸的屏幕让很多 Web 开发人员相当的纠结。幸运的是 CSS3 规范可帮我们轻松的解决此事，你可以根据不同尺寸的屏幕定义不同的 CSS 样式。

例如，下面的代码只在屏幕显示区域大小为 767px 的时候才有效：

>	 @media screen and (max-width:767px){
		#container{
			width:320px;
		} 

>		header h1#logo a{
			width:320px;
			height:44px;
			background:url(image-small.jpg) no-repeat 0 0;
		}
	}

更详细的信息请阅读: [Create an adaptable website layout with CSS3 media queries](http://www.catswhocode.com/blog/create-an-adaptable-website-layout-with-css3-media-queries)<!--more-->

##Font resizing with REMs
CSS3 引入新的字体尺寸单位 __rem (root rm)__

em 单位是相对于父节点的 font-size ，会有一些组合的问题，而 rem 是相对于根节点（或者是 html 节点），意思就是说你可以在 html 节点定义一个单独的字体大小，然后所有其他元素使用 rem 相对于这个字体的百分比进行设置。

>	 html { font-size: 62.5%; }
	body { font-size: 1.4rem; } /* =14px */
	h1   { font-size: 2.4rem; } /* =24px */

更多关于 rem 的内容请看: [Font resizing with REMs](http://snook.ca/archives/html_and_css/font-size-with-rem)

##Cache pages for offline usage
HTML5 引入了一个强大的特性：__离线缓存__。该特性可让你告诉浏览器缓存某些页面，使得用户可以在离线的情况下再次访问该页面。

要缓存页面非常简单，首先在你网站的 .htaccess 文件中添加如下一行：

>	 AddType text/cache-manifest .manifest

然后你可创建一个文件如 offline.manifest ，包含如下内容：

>	 CACHE MANIFEST

>	 CACHE
	index.html  
	style.css  
	image.jpg  

最后，在 html 节点中增加：

>	 <html manifest=""/offline.manifest"">

就这么多。
详情阅读: [How to create offline HTML5 web apps in 5 easy steps](http://www.catswhocode.com/blog/how-to-create-offline-html5-web-apps-in-5-easy-steps)

##Server-side JavaScript
JavaScript 现在已经是非常流行的Web客户端编程语言了，但JavaScript也越来越多的出现在服务器端了，通过强大的 JavaScript 服务器端环境：[Jaxer](http://www.oschina.net/p/jaxer), [Node.js](http://www.oschina.net/p/nodejs) and [Narwhal](http://narwhaljs.org/).

下面代码显示如何用 __Node.js__ 创建一个简单的 Hello World 程序

>	 var sys = require(""sys"");
	sys.puts(""Hello World!"");

更详细内容请阅读: [Learning Server-Side JavaScript with Node.js](http://net.tutsplus.com/tutorials/javascript-ajax/learning-serverside-javascript-with-node-js/)

##HTML5 drag & drop
HTML5 让网页上的拖放变得非常简单，我们只需要简单的定义 __draggable=""true""__ 属性即可，如下所示：

>	  <div id=""columns"">
	  <div class=""column"" draggable=""true""><header>A</header></div>
	  <div class=""column"" draggable=""true""><header>B</header></div>
	  <div class=""column"" draggable=""true""><header>C</header></div>
	 </div>

有了这些 draggable=true 的元素，我们只需要编写一些简单的 JavaScript 代码来处理拖放，这里不再详细描述处理过程，如果你感兴趣，可以阅读这里。

提示：如果你希望阻止可拖放元素被选中，可使用以下 CSS 规则：

>	 [draggable] {
	  -moz-user-select: none;
	  -khtml-user-select: none;
	  -webkit-user-select: none;
	  user-select: none;
	}

__More info__: [Cross Browser HTML5 Drag and Drop](http://www.useragentman.com/blog/2010/01/10/cross-browser-html5-drag-and-drop/)

##Forms, the HTML5 way
__HTML5__ 规范在表单定义方面引入很多新特性，包含很多新的表单组件，例如日期选择、数字调整、使用正则表达式对输入框进行验证等等（email、tel、link）

下面代码显示了一些新的表单元素：

	 <form>
		<label for=""range-slider"">Slider</label>
		<input type=""range"" name=""range-slider"" id=""range-slider"" class=""slider""
		min=""0"" max=""20"" step=""1"" value=""0"">

		<label for=""numeric-spinner"">Numeric spinner</label>
		<input type=""number"" name=""numeric-spinner"" id=""numeric-spinner"" value=""2"">

		<label for=""date-picker"">Date picker</label>
		<input type=""date"" name=""date-picker"" id=""date-picker"" value=""2010-10-06"">

		<label for=""color-picker"">Color picker</label>
		<input type=""color"" name=""color-picker"" id=""color-picker"" value=""ff0000"">

		<label for=""text-field"">Text field with placeholder</label>
		<input type=""text"" name=""text-field"" id=""text-field"" placeholder=""Insert your text 
		here"">

		<label for=""url-field"">Url field</label>
		<input type=""url"" id=""url-field"" name=""url-field"" 
		placeholder=""http://net.tutsplus.com/"" required>

		<label for=""email-field"">Email field</label>
		<input type=""email"" id=""email-field"" name=""email-field"" 
		placeholder=""contact@ghinda.net"" required>

		<button type=""submit"" class=""ui-button ui-widget ui-state-default 
		ui-corner-all ui-button-text-only"" role=""button"" aria-disabled=""false"">
		<span class=""ui-button-text"">Submit form</span>
		</button>
	</form>

__More info__: [How to Build Cross-Browser HTML5 Forms](http://net.tutsplus.com/tutorials/html-css-techniques/how-to-build-cross-browser-html5-forms/)

##CSS animations
很多现在的浏览器都支持 CSS 动画，是的，CSS 已经允许你创建一些简单的动画，而无需 JavaScript 的支持。

下面代码显示如何让背景色改变：

> 	 #logo {
		margin: 15px 15px 0 15px;
		background: red;
		float: left;

>		/* Firefox 4+ */
		-moz-animation-name: colour-change;
		-moz-animation-timing-function: linear;
		-moz-animation-iteration-count: infinite;
		-moz-animation-duration: 30s;

>		/* Webkit */
		-webkit-animation-name: colour-change;
		-webkit-animation-timing-function: linear;
		-webkit-animation-iteration-count: infinite;
		-webkit-animation-duration: 30s;
	}

>	 @-moz-keyframes colour-change {
		0% {
			background: red;
		}
		33% {
			background: green;
		}
		66% {
			background: blue;
		}
	}

>	 @-webkit-keyframes colour-change {
		0% {
			background: red;
		}
		33% {
			background: green;
		}
		66% {
			background: blue;
		}
	}

[来源](http://www.oschina.net/news/21533/7-new-techniques-every-web-developer-should-know)
本文译自：[http://www.catswhocode.com/](http://www.catswhocode.com/blog/7-new-techniques-every-web-developer-should-know)
