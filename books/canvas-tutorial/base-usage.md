# 基本用法 Basic usage

> **使用基础**

## canvas 元素

> Let's start this tutorial by looking at the `<canvas>` element itself.

让我们从`<canvas>`元素的定义开始吧。

    <canvas id="tutorial" width="150" height="150"></canvas>

> This looks a lot like the `<img>` element, the only difference is that it doesn't have the src and alt attributes. The `<canvas>` element has only two attributes - width and height. These are both optional and can also be set using DOM properties or CSS rules. When no width and height attributes are specified, the canvas will initially be 300 pixels wide and 150 pixels high.The element can be sized arbitrarily by CSS, but during rendering the image is scaled to fit its layout size.   (If your  renderings seem distorted, try specifying your width and height attributes explicitly in the `<canvas>` attributes, and not with CSS.)

`<canvas>`看起来很像`<img>`，唯一不同就是它不含 src 和 alt 属性。它只有两个属性，width 和 height，两个都是可选的，并且都可以 DOM 或者 CSS 来设置。如果不指定width 和 height，默认的是宽300像素，高150像素。虽然可以通过 CSS 来调整canvas的大小，但渲染图像会缩放来适应布局的（如果你发现渲染结果看上去变形了，不必一味依赖CSS，可以尝试显式指定canvas的width 和 height 属性值）。

> The id attribute isn't specific to the `<canvas>` element but is one of default HTML attributes which can be applied to (almost) every HTML element (like class for instance). It's always a good idea to supply an id because this makes it much easier to identify it in our script.

id  属性不是`<canvas>`专享的，就像标准的HTML标签一样，任何一个HTML元素都可以指定其 id 值。一般，为元素指定 id 是个不错的主意，这样使得在脚本中应用更加方便。

> The `<canvas>` element can be styled just like any normal image (margin, border, background, etc). These rules however don't affect the actual drawing on the canvas. We'll see how this is done later in this tutorial. When no styling rules are applied to the canvas it will initially be fully transparent.

`<canvas>`元素可以像普通图片一样指定其样式（边距，边框，背景等等）。然而这些样式并不会对canvas实际生成的图像产生什么影响。下面我们会看到如何应用样式。如果不指定样式，canvas默认是全透明的。

### 替用内容

> Because the `<canvas>` element is still relatively new and isn't implemented in some browsers (such as Firefox 1.0 and Internet Explorer), we need a means of providing fallback content when a browser doesn't support the element.

因为 `<canvas>` 相对较新，有些浏览器并没实现，如Firefox 1.0 和 Internet Explorer，所以我们需要为那些不支持canvas的浏览器提供替用显示内容。

> Luckily this is very straightforward: we just provide alternative content inside the canvas element. Browsers who don't support it will ignore the element completely and render the fallback content, others will just render the canvas normally.
For instance we could provide a text description of the canvas content or provide a static image of the dynamically rendered content. This can look something like this:

我们只需要直接在canvas元素内插入替用内容即可。不支持canvas的浏览器会忽略canvas元素而直接渲染替用内容，而支持的浏览器则会正常地渲染canvas。例如，我们可以把一些文字或图片填入canvas内，作为替用内容：

    <canvas id="stockGraph" width="150" height="150">
      current stock price: $3.15 +0.15
    </canvas>

    <canvas id="clock" width="150" height="150">
      <img src="images/clock.png" width="150" height="150"/>
    </canvas>

### 结束标签 `</canvas>` 是必须的

> In the Apple Safari implementation, `<canvas>` is an element implemented in much the same way `<img>` is; it does not have an end tag. However, for `<canvas>` to have widespread use on the web, some facility for fallback content must be provided. Therefore, Mozilla's implementation requires an end tag (`</canvas>`).

在Apple Safari里，`<canvas>`的实现跟`<img>`很相似，它没有结束标签。然而，为了使 `<canvas>` 能在web的世界里广泛适用，需要给替用内容提供一个容身之所，因此，在Mozilla的实现里结束标签(`</canvas>`)是必须的。

> If fallback content is not needed, a simple `<canvas id="foo" ...></canvas>` will be fully compatible with both Safari and Mozilla -- Safari will simply ignore the end tag.

如果没有替用内容，`<canvas id="foo" ...></canvas>` 对 Safari 和 Mozilla 是完全兼容的—— Safari 会简单地忽略结束标签。

> If fallback content is desired, some CSS tricks must be employed to mask the fallback content from Safari (which should render just the canvas), and also to mask the CSS tricks themselves from IE (which should render the fallback content).

如果有替用内容，那么可以用一些 CSS 技巧来为并且仅为 Safari 隐藏替用内容，因为那些替用内容是需要在 IE 里显示但不需要在 Safari 里显示。

## 渲染上下文（Rendering Context）

> `<canvas>` creates a fixed size drawing surface that exposes one or more rendering contexts, which are used to create and manipulate the content shown. We'll focus on the 2D rendering context, which is the only currently defined rendering context. In the future, other contexts may provide different types of rendering; for example, it is likely that a 3D context based on [OpenGL ES][opengl] will be added.

`<canvas>` 创建的固定尺寸的绘图画面开放了一个或多个渲染上下文（rendering context），我们可以通过它们来控制要显示的内容。我们专注于2D 渲染上，这也是目前唯一的选择，可能在将来会添加基于[OpenGL ES][opengl] 的 3D 上下文。

> The `<canvas>` is initially blank, and to display something a script first needs to access the rendering context and draw on it. The canvas element has a DOM method called getContext, used to obtain the rendering context and its drawing functions. getContext() takes one parameter, the type of context.

`<canvas>` 初始化是空白的，要在上面用脚本画图首先需要其渲染上下文（rendering context），它可以通过 canvas 元素对象的 getContext 方法来获取，同时得到的还有一些画图用的函数。getContext() 接受一个用于描述其类型的值作为参数。

    var canvas = document.getElementById('tutorial');
    var ctx = canvas.getContext('2d');

> In the first line we retrieve the canvas DOM node using the getElementById method. We can then access the drawing context using the getContext method.

上面第一行通过 getElementById 方法取得 canvas 对象的 DOM 节点。然后通过其 getContext 方法取得其画图操作上下文。

### 检查浏览器的支持

> The fallback content is displayed in browsers which do not support `<canvas>`; scripts can also check for support when they execute. This can easily be done by testing for the getContext method. Our code snippet from above becomes something like this:

除了在那些不支持  的浏览器上显示替用内容，还可以通过脚本的方式来检查浏览器是否支持 canvas 。方法很简单，判断 getContext 是否存在即可。

    var canvas = document.getElementById('tutorial');
    if (canvas.getContext){
      var ctx = canvas.getContext('2d');
      // drawing code here
    } else {
      // canvas-unsupported code here
    }

## 代码模板

> Here is a minimalistic template, which we'll be using as a starting point for later examples. You can download this file to work with on your system.

我们会用下面这个最简化的代码模板来（后续的示例需要用到）作为开始，你可以 下载文件 到本地备用。

    <html>
      <head>
        <title>Canvas tutorial</title>
        <script type="text/javascript">
          function draw(){
            var canvas = document.getElementById('tutorial');
            if (canvas.getContext){
              var ctx = canvas.getContext('2d');
            }
          }
        </script>
        <style type="text/css">
          canvas { border: 1px solid black; }
        </style>
      </head>
      <body onload="draw();">
        <canvas id="tutorial" width="150" height="150"></canvas>
      </body>
    </html>

> If you look at the script you'll see I've made a function called draw, which will get executed once the page finishes loading (via the onload attribute on the body tag). This function could also have been called from a setTimeout, setInterval, or any other event handler function just as long the page has been loaded first.

细心的你会发现我准备了一个名为 draw 的函数，它会在页面装载完毕之后执行一次（通过设置 body 标签的 onload 属性），它当然也可以在 setTimeout，setInterval，或者其他事件处理函数中被调用。

### 一个简单的例子

> To start off, here's a simple example that draws two intersecting rectangles, one of which has alpha transparency. We'll explore how this works in more detail in later examples.

作为开始，来一个简单的吧——绘制两个交错的矩形，其中一个是有alpha透明效果。我们会在后面的示例中详细的让你了解它是如何运作的。

[观看示例](https://developer.mozilla.org/@api/deki/files/2931/=simple_example_(1).html)

    <html>
     <head>
      <script type="application/x-javascript">
        function draw() {
          var canvas = document.getElementById("canvas");
          if (canvas.getContext) {
            var ctx = canvas.getContext("2d");

            ctx.fillStyle = "rgb(200,0,0)";
            ctx.fillRect (10, 10, 55, 50);

            ctx.fillStyle = "rgba(0, 0, 200, 0.5)";
            ctx.fillRect (30, 30, 55, 50);
          }
        }
      </script>
     </head>
     <body onload="draw();">
       <canvas id="canvas" width="150" height="150"></canvas>
     </body>
    </html>


<blockquote>
<canvas id="canvas" width="150" height="150"></canvas>
<script type="application/x-javascript">
function draw() {
  var canvas = document.getElementById("canvas");
  if (canvas.getContext) {
    var ctx = canvas.getContext("2d");

    ctx.fillStyle = "rgb(200,0,0)";
    ctx.fillRect (10, 10, 55, 50);

    ctx.fillStyle = "rgba(0, 0, 200, 0.5)";
    ctx.fillRect (30, 30, 55, 50);
  }
}
draw();
</script>
</blockquote>

[opengl]:http://en.wikipedia.org/wiki/OpenGL_ES "opengl"
