# Drawing shapes 绘制图形

## 网格 The grid

<img align="right" alt="" class="internal" src="https://developer.mozilla.org/@api/deki/files/78/=Canvas_default_grid.png" />

> Before we can start drawing, we need to talk about the canvas grid or *coordinate space*. The HTML template on the previous page had a canvas element 150 pixels wide and 150 pixels high. I've drawn this image with the default grid overlayed. Normally 1 unit in the grid corresponds to 1 pixel on the canvas. The origin of this grid is positioned in the top left corner (coordinate (0,0)). All elements are placed relative to this origin. So the position of the top left corner of the blue square becomes x pixels from the left and y pixels from the top (coordinate (x,y)). Later in this tutorial we'll see how we can translate the origin to a different position, rotate the grid and even scale it. For now we'll stick to the default.

在真正开始之前，我们需要先探讨 canvas 的网格（grid）或者坐标空间（*coordinate space*）。在前一页的HTML模板里有一个150像素宽， 150像素高的 canvas 对象。我在画面上叠加上默认网格，如右图。通常网格的1个单元对应 canvas 上的1个像素。网格的原点是定位在左上角（坐标(0,0)）。画面里的所有物体的位置都是相对这个原点。这样，左上角的蓝色方块的位置就是距左边x像素和距上边Y像素（坐标(x, y)）。后面的教程中我们将学会如何把移动原点，旋转以及缩放网格。不过现在我们会使用默认的状态。

## 绘制图形 Drawing shapes

> Unlike [SVG][1], canvas only supports one primitive shape - rectangles. All other shapes must be created by combining one or more paths. Luckily, we have a collection of path drawing functions which make it possible to compose very complex shapes.

不像 [SVG][1]，canvas 只支持一种基本形状——矩形，所以其它形状都是有一个或多个路径组合而成。还好，有一组路径绘制函数让我们可以绘制相当复杂的形状。

### 矩形 Rectangles

> First let's look at the rectangle. There are three functions that draw rectangles on the canvas:

我们首先看看矩形吧，有三个函数用于绘制矩形的：

    fillRect(x,y,width,height) : Draws a filled rectangle
    strokeRect(x,y,width,height) : Draws a rectangular outline
    clearRect(x,y,width,height) : Clears the specified area and makes it fully transparent

> Each of these three functions takes the same parameters. `x` and `y` specify the position on the canvas (relative to the origin) of the top-left corner of the rectangle. `width` and `height` are pretty obvious. Let's see these functions in action.

它们都接受四个参数， `x` 和 `y 指定矩形左上角(相对于原点)的位置，``width` <span style="">和 </span>`height` 是矩形的宽和高。好，实战一下吧。

> Below is the `draw()` function from the previous page, but now I've added the three functions above.

下面就是上页模板里的 `draw() `函数，但添加了上面的三个函数。

#### 绘制矩形的例子 Rectangular shape example

<a class="external" href="http://developer.mozilla.org/samples/canvas-tutorial/2_1_canvas_rect.html">观看示例</a>

    function draw(){
      var canvas = document.getElementById('tutorial');
      if (canvas.getContext){
        var ctx = canvas.getContext('2d');

        ctx.fillRect(25,25,100,100);
        ctx.clearRect(45,45,60,60);
        ctx.strokeRect(50,50,50,50);
      }
    }

<img align="right" alt="" class="internal" src="https://developer.mozilla.org/@api/deki/files/97/=Canvas_rect.png" />

> The result should look something like the image on the right. The `fillRect` function draws a large black square 100x100 pixels. The `clearRect` function removes a 60x60 pixels square from the center and finally the `strokeRect` draws a rectangular outline 50x50 pixels inside the cleared square. In the following pages we'll see two alternative methods for the `clearRect` function and we'll also see how to change the color and stroke style of the rendered shapes.

出来的结果应该和右边的是一样的。`fillRect` 函数画了一个大的黑色矩形（100x100），`clearRect` 函数清空了中间 60x60 大小的方块，然后`strokeRect` 函数又在清空了的空间内勾勒出一个 50x50 的矩形边框。在接下去的页面里，我们会看到和 `clearRect` 函数差不多另外两个方法，以及如何去改变图形的填充和边框颜色。

> Unlike the path functions we'll see in the next section, all three rectangle functions draw immediately to the canvas.

与下一节的路径函数不一样，这三个函数的效果会立刻在 canvas 上反映出来。

## 绘制路径 Drawing paths

> To make shapes using paths, we need a couple of extra steps.

不像画矩形那样的直截了当，绘制路径是需要一些额外的步骤的。

    beginPath()
    closePath()
    stroke()
    fill()

> The first step to create a path is calling the `beginPath` method. Internally, paths are stored as a list of sub-paths (lines, arcs, etc) which together form a shape. Every time this method is called, the list is reset and we can start drawing new shapes.

第一步是用 `beginPath` 创建一个路径。在内存里，路径是以一组子路径（直线，弧线等）的形式储存的，它们共同构成一个图形。每次调用 `beginPath`，子路径组都会被重置，然后可以绘制新的图形。

> The second step is calling the methods that actually specify the paths to be drawn. We'll see these shortly.

第二步就是实际绘制路径的部分，很快我们就会看到。

> The third, and an optional step, would be to call the `closePath` method. This method tries to close the shape by drawing a straight line from the current point to the start. If the shape has already been closed or there's only one point in the list, this function does nothing.

第三步是调用 `closePath` 方法，它会尝试用直线连接当前端点与起始端点来关闭路径，但如果图形已经关闭或者只有一个点，它会什么都不做。这一步不是必须的。

> The final step will be calling the `stroke` and/or `fill` methods. Calling one of these will actually draw the shape to the canvas. `stroke` is used to draw an outlined shape, while `fill` is used to paint a solid shape.

最后一步是调用 `stroke` 或 `fill 方法，这时，图形才是实际的绘制到 canvas` 上去。`stroke` 是绘制图形的边框，`fill` 会用填充出一个实心图形。

**Note:** When calling the `fill` method any open shapes will be closed automatically and it isn't necessary to use the `closePath` method.

**注意：**当调用` fill `时，开放的路径会自动闭合，而无须调用` closePath` 。

> The code for a drawing simple shape (a triangle) would look something like this.

画一个简单图形（如三角形）的代码如下。

    ctx.beginPath();
    ctx.moveTo(75,50);
    ctx.lineTo(100,75);
    ctx.lineTo(100,25);
    ctx.fill();

### moveTo

> One very useful function, which doesn't actually draw anything, but is part of the path list described above, is the `moveTo` function. You can probably best think of this as lifting a pen or pencil from one spot on a piece of paper and placing it on the next.

`moveTo` 是一个十分有用的方法，虽然并不能用它来画什么，但却是绘制路径的实用方法的一部分。你可以把它想象成是把笔提起，并从一个点移动到另一个点的过程。

    moveTo(x, y)

> The `moveTo` function takes two arguments - `x` and `y`, - which are the coordinates of the new starting point.

它接受 `x` 和 `y` （新的坐标位置）作为参数。

<img align="right" alt="" class="internal" src="https://developer.mozilla.org/@api/deki/files/107/=Canvas_smiley.png" />

> When the canvas is initialized or the `beginPath` method is called, the starting point is set to the coordinate (0,0). In most cases we would use the `moveTo` method to place the starting point somewhere else. We could also use the `moveTo` method to draw unconnected paths. Take a look at the smiley face on the right. I've marked the places where I used the `moveTo` method (the red lines).

当 canvas 初始化或者调用 `beginPath` 的时候，起始坐标设置就是原点(0,0)。大多数情况下，我们用 `moveTo` 方法将起始坐标移至其它地方，或者用于绘制不连续的路径。看看右边的笑脸，红线就是使用 `moveTo 移动的轨迹。`

> To try this for yourself, you can use the code snippet below. Just paste it into the `draw` function we saw earlier.

试一试下面的代码，粘贴到之前用过的 `draw` 函数内在看看效果吧。

#### `moveTo` 的使用示例

<a class="external" href="http://developer.mozilla.org/samples/canvas-tutorial/2_2_canvas_moveto.html">观看示例</a>

    ctx.beginPath();
    ctx.arc(75,75,50,0,Math.PI*2,true); // Outer circle
    ctx.moveTo(110,75);
    ctx.arc(75,75,35,0,Math.PI,false);   // Mouth (clockwise)
    ctx.moveTo(65,65);
    ctx.arc(60,65,5,0,Math.PI*2,true);  // Left eye
    ctx.moveTo(95,65);
    ctx.arc(90,65,5,0,Math.PI*2,true);  // Right eye
    ctx.stroke();

    //thegoneheart 完整例子

    ctx.beginPath();
    ctx.arc(75,75,50,0,Math.PI*2,true); // Outer circle
    ctx.moveTo(110,75);
    ctx.arc(75,75,35,0,Math.PI,false);   // Mouth (clockwise)
    ctx.moveTo(65,65);
    ctx.arc(60,65,5,0,Math.PI*2,true);  // Left eye
    ctx.moveTo(95,65);
    ctx.arc(90,65,5,0,Math.PI*2,true);  // Right eye
    ctx.stroke();
                
    ctx.beginPath();
    ctx.moveTo(40,75);
    ctx.lineTo(60,65);
    ctx.lineTo(90,65);
    ctx.moveTo(110,75);
    ctx.lineTo(125,75);
    ctx.stroke();

**Note**: remove the `moveTo` methods to see the connecting lines.  
**Note**: For a description of the `arc` function and its parameters look below.

**注意：**你可以注释 `moveTo` 方法来观察那些连接起来的线。  
**注意：**`arc` 方法的用法见下面。

### 绘制各种线条 Lines

> For drawing straight lines we use the `lineTo` method.

我们用` lineTo` 方法来画直线。

    lineTo(x, y)


> This method takes two arguments - `x` and `y`, - which are the coordinates of the line's end point. The starting point is dependent on previous drawn paths, where the end point of the previous path is the starting point for the following, etc. The starting point can also be changed by using the `moveTo` method.

`lineTo` 方法接受终点的坐标（x，y）作为参数。起始坐标取决于前一路径，前一路径的终点即当前路径的起点，起始坐标也可以通过 `moveTo` 方法来设置。

#### `lineTo` 的使用示例

<img align="right" alt="" class="internal" src="https://developer.mozilla.org/@api/deki/files/86/=Canvas_lineTo.png" />

> In the example below two triangles are drawn, one filled and one outlined. (The result can be seen in the image on the right). First the `beginPath` method is called to begin a new shape path. We then use the `moveTo` method to move the starting point to the desired position. Below this two lines are drawn which make up two sides of the triangle.

示例（如右图）画的是两个三角形，一个实色填充，一个勾边。首先调用 `beginPath` 方法创建一个新路径，然后用`moveTo` 方法将起始坐标移至想要的位置，然后画两条直线来构成三角形的两条边。

> You'll notice the difference between the filled and stroked triangle. This is, as mentioned above, because shapes are automatically closed when a path is filled. If we would have done this for the stroked triangle only two lines would have been drawn, not a complete triangle.

可以注意到 fill 和 strok 绘三角形的区别，上面也提到过，使用 fill 路径会自动闭合，但使用 stroke 不会，如果不关闭路径，勾画出来的只有两边。

<a class="external" href="http://developer.mozilla.org/samples/canvas-tutorial/2_3_canvas_lineto.html">观看示例</a>

    // 填充三角形
    ctx.beginPath();
    ctx.moveTo(25,25);
    ctx.lineTo(105,25);
    ctx.lineTo(25,105);
    ctx.fill();

    // 勾边三角形
    ctx.beginPath();
    ctx.moveTo(125,125);
    ctx.lineTo(125,45);
    ctx.lineTo(45,125);
    ctx.closePath();
    ctx.stroke(); 

### 弧线 Arcs

> For drawing arcs or circles we use the `arc` method. The specification also describes the `arcTo` method, which is supported by Safari but hasn't been implemented in the current Gecko browsers.

我们用 `arc` 方法来绘制弧线或圆。标准说明中还包含 `arcTo` 方法，当前 Safari 是支持的，但基于 Gecko 的浏览器还未实现。

    arc(x, y, radius, startAngle, endAngle, anticlockwise)

> This method takes five parameters: `x` and `y` are the coordinates of the circle's center. Radius is self explanatory. The `startAngle` and `endAngle` parameters define the start and end points of the arc in radians. The starting and closing angle are measured from the x axis. The `anticlockwise` parameter is a boolean value which when `true` draws the arc anticlockwise, otherwise in a clockwise direction.

方法接受五个参数：x，y 是圆心坐标，radius 是半径，`startAngle` 和 `endAngle` 分别是起末弧度（以 x 轴为基准），`anticlockwise` 为 true 表示逆时针，反之顺时针。

> **Warning**: In the Firefox beta builds, the last parameter is `clockwise`. The final release will support the function as described above. All scripts that use this method in its current form will need to be updated once the final version is released.

警告：在 Firefox 的 beta 版本里，最后一个参数是 `clockwise，而最终版本`不是。因此如果是从 beta 升级至发行版需要做相应修改。

> **Note**: Angles in the `arc` function are measured in radians, not degrees. To convert degrees to radians you can use the following JavaScript expression: `var radians = (Math.PI/180)*degrees`.

注意：`arc` 方法里用到的角度是以弧度为单位而不是度。度和弧度直接的转换可以用这个表达式：`var radians = (Math.PI/180)*degrees;。`

#### `arc` 的使用示例

<img align="right" alt="" class="internal" src="https://developer.mozilla.org/@api/deki/files/55/=Canvas_arc.png" />

> The following example is a little more complex than the ones we've seen above. I've drawn 12 different arcs all with different angles and fills. If I would have written this example just like the smiley face above, firstly this would have become a very long list of statements and secondly, when drawing arcs, I would need to know every single starting point. For arcs of 90, 180 and 270 degrees, like the ones I used here, this wouldn't be to much of a problem, but for more complex ones this becomes way too difficult.

这个示例比之前见到过的要复杂一些，画了12个不同的弧形，有不同夹角和填充状态的。如果我用上面画笑脸的方式来画这些弧形，那会是一大段的代码，而且，画每一个弧形时我都需要知道其圆心位置。像我这里画 90，180 和 270 度的弧形看起来不是很麻烦，但是如果图形更复杂一些,则实现起来会越来越困难。

> The two `for` loops are for looping through the rows and columns of arcs. For every arc I start a new path using `beginPath`. Below this I've written out all the parameters as variables, so it's easier to read what's going on. Normally this would be just one statement. The `x` and `y` coordinates should be clear enough. `radius` and `startAngle` are fixed. The `endAngle` starts of as 180 degrees (first column) and is increased with steps of 90 degrees to form a complete circle (last column). The statement for the `clockwise` parameter results in the first and third row being drawn as clockwise arcs and the second and fourth row as counterclockwise arcs. Finally, the `if` statement makes the top half stroked arcs and the bottom half filled arcs.

这里使用两个 `for` 循环来画多行多列的弧形。每一个弧形都用 `beginPath `方法创建一个新路径。然后为了方便阅读和理解，我把所有参数都写成变量形式。显而易见，x 和 y 作为圆心坐标。 `radius` <span style="">和 </span>`startAngle` 都是固定，`endAngle` 从 180 度半圆开始，以 90 度方式递增至圆。`anticlockwise` 则取决于奇偶行数。最后，`通过 if` 语句判断使前两行表现为勾边，而后两行为填充效果。

<a class="external" href="http://developer.mozilla.org/samples/canvas-tutorial/2_4_canvas_arc.html">观看示例</a>

    for (i=0;i<4;i++){
       for(j=0;j<3;j++){    //chinese_xu 原始代码
        ctx.beginPath();
        var x              = 25+j*50;               // x coordinate
        var y              = 25+i*50;               // y coordinate
        var radius         = 20;                    // Arc radius
        var startAngle     = 0;                     // Starting point on circle
        var endAngle       = Math.PI+(Math.PI*j)/2;  // End point on circle ---//修复错误标点
        var anticlockwise  = i%2==0 ? false : true; // clockwise or anticlockwise

        ctx.arc(x,y,radius,startAngle,endAngle, anticlockwise);

        if (i>1){
          ctx.fill();
        } else {
          ctx.stroke();
        }
      }
    }
    //chinese_xu 原始代码并没有按照1/4圆递增来画。
    //修改后输出4行4列，要把画布扩大到200*200观看
    for (i=0;i<4;i++){
       for(j=0;j<4;j++){    
        ctx.beginPath();
        var x              = 25+j*50;               // x coordinate
        var y              = 25+i*50;               // y coordinate
        var radius         = 20;                    // Arc radius
        var startAngle     = 0;                     // Starting point on circle
        var endAngle       = Math.PI*(2-j/2);   // End point on circle
        var anticlockwise  = i%2==0 ? false : true; // clockwise or anticlockwise

        ctx.arc(x,y,radius,startAngle,endAngle, anticlockwise);

        if (i>1){
          ctx.fill();
        } else {
          ctx.stroke();
        }
      }
    }

### 贝塞尔和二次方曲线 Bezier and quadratic curves

> The next type of paths available are <a class="external" href="http://en.wikipedia.org/wiki/B%C3%A9zier_curve">Bézier curves</a>, available in the cubic and quadratic varieties. These are generally used to draw complex organic shapes.

 接下来要介绍的路径是 <a class="external" href="http://zh.wikipedia.org/w/index.php?title=%E8%B2%9D%E8%8C%B2%E6%9B%B2%E7%B7%9A&variant=zh-hans">贝塞尔曲线</a> ，它可以是二次和三次方的形式，一般用于绘制复杂而有规律的形状。

    quadraticCurveTo(cp1x, cp1y, x, y) // BROKEN in Firefox 1.5 (see work around below)
    bezierCurveTo(cp1x, cp1y, cp2x, cp2y, x, y)

<img align="right" alt="" class="internal" src="https://developer.mozilla.org/@api/deki/files/76/=Canvas_curves.png" />

> The difference between these can best be described using the image on the right. A quadratic Bézier curve has a start and an end point (blue dots) and just one *control point* (red dot) while a cubic Bézier curve uses two control points.

上面两行代码的区别见右图。它们都有一个起点一个终点（图中的蓝点），但二次方贝塞尔曲线只有一个（红色）控制点点）而三次方贝塞尔曲线有两个。

> The `x` and `y` parameters in both these methods are the coordinates of the end point. `cp1x` and `cp1y` are the coordinates of the first control point, and `cp2x` and `cp2y` are the coordinates of the second control point.

`参数 x` 和 `y` 是终点坐标，`cp1x` <span style="">和 </span>`cp1y 是第一个控制点的坐标，``cp2x` <span style="">和 </span>`cp2y` 是第二个的。

> Using quadratic and cubic Bézier curves can be quite challenging, because unlike vector drawing software like Adobe Illustrator, we don't have direct visual feedback as to what we're doing. This makes it pretty hard to draw complex shapes. In the following example, we'll be drawing some simple organic shapes, but if you have the time and, most of all, the patience, much more complex shapes can be created.

使用二次方和三次方的贝塞尔曲线是相当有挑战的，因为不像在矢量绘图软件 Adobe Illustrator 里那样有即时的视觉反馈。因为用它来画复杂图形是比较麻烦的。但如果你有时间，并且最重要是有耐心，再复杂的图形都可以绘制出来的。下面我们来画一个简单而又规律的图形。

> There's nothing very difficult in these examples. In both cases we see a succession of curves being drawn which finally result in a complete shape.

这些例子都比较简单。我们绘制的都是完整的图形。

#### `quadraticCurveTo` 的使用示例

<a class="external" href="http://developer.mozilla.org/samples/canvas-tutorial/2_5_canvas_quadraticcurveto.html">查看示例</a> <img align="right" alt="" class="internal" src="https://developer.mozilla.org/@api/deki/files/94/=Canvas_quadratic.png" />

    // Quadratric curves example
    ctx.beginPath();
    ctx.moveTo(75,25);
    ctx.quadraticCurveTo(25,25,25,62.5);
    ctx.quadraticCurveTo(25,100,50,100);
    ctx.quadraticCurveTo(50,120,30,125);
    ctx.quadraticCurveTo(60,120,65,100);
    ctx.quadraticCurveTo(125,100,125,62.5);
    ctx.quadraticCurveTo(125,25,75,25);
    ctx.stroke();

It is possible to convert any quadratic Bézier curve to a cubic Bézier curve by correctly computing both cubic Bézier control points from the single quadratic Bézier control point, although the reverse is NOT true. An exact conversion of a cubic Bézier curve to a quadratic Bézier curve is only possible if the cubic term is zero, more commonly a subdivision method is used to approximate a cubic Bézier using multiple quadratic Bézier curves.

通过计算，可以由二次曲线的单个控制点得出相应三次方曲线的两个控制点，因此二次方转三次方是可能的，但是反之不然。仅当三次方程中的三次项为零是才可能转换为二次的贝塞尔曲线。通常地可以用多条二次方曲线通过细分算法来近似模拟三次方贝塞尔曲线。

#### `bezierCurveTo` 的使用示例

<a class="external" href="http://developer.mozilla.org/samples/canvas-tutorial/2_6_canvas_beziercurveto.html">查看示例</a> <img align="right" alt="" class="internal" src="https://developer.mozilla.org/@api/deki/files/59/=Canvas_bezier.png" />


    // Bezier curves example
    ctx.beginPath();
    ctx.moveTo(75,40);
    ctx.bezierCurveTo(75,37,70,25,50,25);
    ctx.bezierCurveTo(20,25,20,62.5,20,62.5);
    ctx.bezierCurveTo(20,80,40,102,75,120);
    ctx.bezierCurveTo(110,102,130,80,130,62.5);
    ctx.bezierCurveTo(130,62.5,130,25,100,25);
    ctx.bezierCurveTo(85,25,75,37,75,40);
    ctx.fill();

### Firefox 1.5 quadraticCurveTo() bug 的应对方案

> There is a bug in the Firefox 1.5 implementation of quadatricCurveTo(). It does NOT draw a quadratic curve, as it is just calling the same cubic curve function bezierCurveTo() calls, and repeating the single quadratic control point (x,y) coordinate twice. For this reason quadraticCurveTo() will yield incorrect results. If you require the use of quadraticCurveTo() you must convert your quadratic Bézier curve to a cubic Bézier curve yourself, so you can use the working bezierCurveTo() method.

在 Firefox 1.5 里，quadatricCurveTo() 的实现是有 bug 的，它不是直接绘制二次方曲线，而是调用 bezierCurveTo() ，其中两个控制点都是二次方曲线的那个单控制点。因此，它会绘制出不正确的曲线。如果必须使用到 quadraticCurveTo()，你需要自行去将二次方曲线转换成三次方的，这样就可以用 bezierCurveTo() 方法了。

    var currentX, currentY;  // set to last x,y sent to lineTo/moveTo/bezierCurveTo or quadraticCurveToFixed()

    function quadraticCurveToFixed( cpx, cpy, x, y ) {
      /*
       For the equations below the following variable name prefixes are used:
         qp0 is the quadratic curve starting point (you must keep this from your last point sent to moveTo(), lineTo(), or bezierCurveTo() ).
         qp1 is the quadatric curve control point (this is the cpx,cpy you would have sent to quadraticCurveTo() ).
         qp2 is the quadratic curve ending point (this is the x,y arguments you would have sent to quadraticCurveTo() ).
       We will convert these points to compute the two needed cubic control points (the starting/ending points are the same for both
       the quadratic and cubic curves.

       The equations for the two cubic control points are:
         cp0=qp0 and cp3=qp2
         cp1 = qp0 + 2/3 *(qp1-qp0)
         cp2 = cp1 + 1/3 *(qp2-qp0) 

       In the code below, we must compute both the x and y terms for each point separately. 

        cp1x = qp0x + 2.0/3.0*(qp1x - qp0x);
        cp1y = qp0y + 2.0/3.0*(qp1y - qp0y);
        cp2x = cp1x + (qp2x - qp0x)/3.0;
        cp2y = cp1y + (qp2y - qp0y)/3.0;

       We will now 
         a) replace the qp0x and qp0y variables with currentX and currentY (which *you* must store for each moveTo/lineTo/bezierCurveTo)
         b) replace the qp1x and qp1y variables with cpx and cpy (which we would have passed to quadraticCurveTo)
         c) replace the qp2x and qp2y variables with x and y.
       which leaves us with: 
      */
      var cp1x = currentX + 2.0/3.0*(cpx - currentX);
      var cp1y = currentY + 2.0/3.0*(cpy - currentY);
      var cp2x = cp1x + (x - currentX)/3.0;
      var cp2y = cp1y + (y - currentY)/3.0;

      // and now call cubic Bezier curve to function 
      bezierCurveTo( cp1x, cp1y, cp2x, cp2y, x, y );

      currentX = x;
      currentY = y;
    }
     

### 矩形路径 Rectangles

> Besides the three methods we saw above which draw rectangular shapes directly to the canvas, we also have a method `rect` which adds a rectangular path to the path list.

除了上面提到的三个方法可以直接绘制矩形之外，我们还有一个 `rect` 方法是用于绘制矩形路径的。

    rect(x, y, width, height)

> This method takes four arguments. The `x` and `y` parameters define the coordinate of the top left corner of the new rectangular path. `width` and `height` define the width and the height of the rectangle.

它接受四个参数，`x`和 `y` 是其左上角坐标，`width` 和 `height` 是其宽和高。  

> When this method is executed, the `moveTo` method is automatically called with the parameters (0,0) (i.e. it resets the starting point to its default location).

当它被调用时，`moveTo` 方法会自动被调用，参数为(0,0)，于是起始坐标又恢复成初始原点了。

### 综合 Making combinations

> In all examples on this page I've only used one type of path function per shape. However there's absolutely no limitation to the amount or type of paths you can use to create a shape. So in this last example I've tried to combine all of the path functions to make a set of very famous game characters.

上面所用到的例子都只用到了一种类型的路径，当然 canvas 不会限制所使用的路径类型的多少。所以，我们来看一个路径大杂烩。

#### 综合样例

> I'm not going to run through this complete script, but the most important things to note are the function `roundedRect` and the use of the `fillStyle` property. It can be very usefull and time saving to define your own functions to draw more complex shapes. In this script it would have taken me twice as many lines of code as I have now.  
We will look at the `fillStyle` property in greater depth later in this tutorial. Here I'm using it to change the fill color from the default black, to white, and back again.

在整个例子里，最值得注意的是 `roundedRect` 函数的使用和 `fillStyle` 属性的设置。`<span style="">自定义函数对于封装复杂图形</span>`的绘制是非常有用的。在这个例子里使用自定义函数就省掉了大约一半的代码。

在接下来的例子里会深入探讨 `fillStyle` 属性的使用。这里是用它来改变填充颜色，从默认的黑色，到白色，然后再回到黑色。

<a class="external" href="http://developer.mozilla.org/samples/canvas-tutorial/2_7_canvas_combined.html">查看示例</a>

    function draw() {
      var ctx = document.getElementById('canvas').getContext('2d');
      roundedRect(ctx,12,12,150,150,15);
      roundedRect(ctx,19,19,150,150,9);
      roundedRect(ctx,53,53,49,33,10);
      roundedRect(ctx,53,119,49,16,6);
      roundedRect(ctx,135,53,49,33,10);
      roundedRect(ctx,135,119,25,49,10);

      ctx.beginPath();
      ctx.arc(37,37,13,Math.PI/7,-Math.PI/7,false); //chiensexu  本来是true呵呵，反了
      ctx.lineTo(31,37);
      ctx.fill();
      for(i=0;i<8;i++){
        ctx.fillRect(51+i*16,35,4,4);
      }
      for(i=0;i<6;i++){
        ctx.fillRect(115,51+i*16,4,4);
      }
      for(i=0;i<8;i++){
        ctx.fillRect(51+i*16,99,4,4);
      }
      ctx.beginPath();
      ctx.moveTo(83,116);
      ctx.lineTo(83,102);
      ctx.bezierCurveTo(83,94,89,88,97,88);
      ctx.bezierCurveTo(105,88,111,94,111,102);
      ctx.lineTo(111,116);
      ctx.lineTo(106.333,111.333);
      ctx.lineTo(101.666,116);
      ctx.lineTo(97,111.333);
      ctx.lineTo(92.333,116);
      ctx.lineTo(87.666,111.333);
      ctx.lineTo(83,116);
      ctx.fill();
      ctx.fillStyle = "white";
      ctx.beginPath();
      ctx.moveTo(91,96);
      ctx.bezierCurveTo(88,96,87,99,87,101);
      ctx.bezierCurveTo(87,103,88,106,91,106);
      ctx.bezierCurveTo(94,106,95,103,95,101);
      ctx.bezierCurveTo(95,99,94,96,91,96);
      ctx.moveTo(103,96);
      ctx.bezierCurveTo(100,96,99,99,99,101);
      ctx.bezierCurveTo(99,103,100,106,103,106);
      ctx.bezierCurveTo(106,106,107,103,107,101);
      ctx.bezierCurveTo(107,99,106,96,103,96);
      ctx.fill();
      ctx.fillStyle = "black";
      ctx.beginPath();
      ctx.arc(101,102,2,0,Math.PI*2,true);
      ctx.fill();
      ctx.beginPath();
      ctx.arc(89,102,2,0,Math.PI*2,true);
      ctx.fill();
    }
    function roundedRect(ctx,x,y,width,height,radius){
      ctx.beginPath();
      ctx.moveTo(x,y+radius);
      ctx.lineTo(x,y+height-radius);
      ctx.quadraticCurveTo(x,y+height,x+radius,y+height);
      ctx.lineTo(x+width-radius,y+height);
      ctx.quadraticCurveTo(x+width,y+height,x+width,y+height-radius);
      ctx.lineTo(x+width,y+radius);
      ctx.quadraticCurveTo(x+width,y,x+width-radius,y);
      ctx.lineTo(x+radius,y);
      ctx.quadraticCurveTo(x,y,x,y+radius);
      ctx.stroke();
    }

 [1]: https://developer.mozilla.org/en/SVG "en/SVG"

