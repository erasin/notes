# 绘图
> **Drawing in PyQt4**

> Drawing is used, when we want to change or enhance an existing widget. Or if we are creating a custom widget from scratch. To do the drawing, we use the drawing API provided by the PyQt4 toolkit.

当我们想要改变或者增强已存在的窗口部件时，或者准备从零开始创建自定义窗口部件时，可以使用绘图。我们通过使用PyQt4工具包提供的绘图API来绘图。

> The drawing is done within the paintEvent() method. The drawing code is placed between the begin() and end() methods of the QPainter object.

绘图在 `paintEvent()` 方法中进行。绘制代码在 `QPainter` 对象的 `begin()` 和 `end()` 之间。


##　绘制文本
> **Drawing text**

> We begin with drawing some Unicode text onto the window client area.

我们从在窗口客户区绘制一些Unicode文本开始。

::

    #!/usr/bin/python
    # -*- coding: utf-8 -*-

    # drawtext.py

    import sys
    from PyQt4 import QtGui, QtCore


    class Example(QtGui.QWidget):
  
        def __init__(self):
            super(Example, self).__init__()
        
            self.initUI()
        
        def initUI(self):

            self.setGeometry(300, 300, 250, 150)
            self.setWindowTitle('Draw Text')

            self.text = u'\u041b\u0435\u0432 \u041d\u0438\u043a\u043e\u043b\u0430\
      \u0435\u0432\u0438\u0447 \u0422\u043e\u043b\u0441\u0442\u043e\u0439: \n\
      \u0410\u043d\u043d\u0430 \u041a\u0430\u0440\u0435\u043d\u0438\u043d\u0430'


        def paintEvent(self, event):

            qp = QtGui.QPainter()
            qp.begin(self)
            self.drawText(event, qp)
            qp.end()


        def drawText(self, event, qp):
      
            qp.setPen(QtGui.QColor(168, 34, 3))
            qp.setFont(QtGui.QFont('Decorative', 10))
            qp.drawText(event.rect(), QtCore.Qt.AlignCenter, self.text)

    app = QtGui.QApplication(sys.argv)
    ex = Example()
    ex.show()
    app.exec_()

> In our example, we draw some text in Azbuka. The text is vertically and horizontally aligned.

在我们的例子中，我们绘制一些西里尔字母的文本，文正水平和垂直居中对齐。

> Note:
:    译者改成了中文文本。

::

    def paintEvent(self, event):
    
> Drawing is done within a paint event.

在绘制事件中绘画。

::

    qp = QtGui.QPainter()
    qp.begin(self)
    self.drawText(event, qp)
    qp.end()
    
> The QPainter class is responsible for all the low-level painting. All the painting methods go between begin() and end() methods. The actual painting is delegated to the drawText() method.

`QPainter` 类负责所有的低级绘画。所有的绘制方法都在 `begin()` 和 `end()` 方法之间。这里实际的绘制是代理给了 `drawText()` 方法。

::

    paint.setPen(QtGui.QColor(168, 34, 3))
    paint.setFont(QtGui.QFont('Decorative', 10))
    
> Here we define pen and font, which we use to draw the text.

这里我们定义了画笔和字体，用来绘制文本。

::

    paint.drawText(event.rect(), QtCore.Qt.AlignCenter, self.text)
    
> The drawText() method draws text on the window. The rect() method of the paint event returns the rectangle that needs to be updated.

`drawText()` 方法在窗口上绘制文本， 绘制事件的 `rect()` 方法返回需要更新的矩形。

![Drawing Text][drawing-text]

图：绘制文本

## 绘制点
> **Drawing points**

> A point is the most simple graphics object, that can be drawn. It is a small spot on the window.

点是可以绘制的最简单的图形对象，是窗口上的很小的一个区域。

::

    #!/usr/bin/python
    # -*- coding: utf-8 -*-

    # points.py

    import sys, random
    from PyQt4 import QtGui, QtCore


    class Example(QtGui.QWidget):
  
        def __init__(self):
            super(Example, self).__init__()

            self.setGeometry(300, 300, 250, 150)
            self.setWindowTitle('Points')

        def paintEvent(self, e):
      
            qp = QtGui.QPainter()
            qp.begin(self)
            self.drawPoints(qp)
            qp.end()
        
        def drawPoints(self, qp):
      
            qp.setPen(QtCore.Qt.red)
            size = self.size()
        
            for i in range(1000):
                x = random.randint(1, size.width()-1)
                y = random.randint(1, size.height()-1)
                qp.drawPoint(x, y)

    app = QtGui.QApplication(sys.argv)
    ex = Example()
    ex.show()
    app.exec_()

> In our example, we draw randomly 1000 red points on the client area.

在这个例子中，我们在客户区随机地绘制1000个红点。

::

    paint.setPen(QtCore.Qt.red)
    
> We set the pen to red color. We use a predefined color constant.

使用预定义的颜色常量，把画笔设为红色。

::

    size = self.size()
    
> Each time we resize the window, a paint event is generated. We get the current size of the window with the size() method. We use the size of the window to distribute the points all over the client area of the window.

每次我们缩放窗口，都将产生绘制事件。通过 `size()` 方法得到窗口的尺寸，使用窗口尺寸来把点分布到窗口的客户区。

::

    paint.drawPoint(x, y)
    
> We draw the point with the drawPoint() method.

使用 `drawPoint()` 方法绘制点。

![Point][point]

图：点


## 颜色
> **Colors**

> A color is an object representing a combination of Red, Green, and Blue (RGB) intensity values. Valid RGB values are in the range 0 to 255. We can define a color in various ways. The most common are RGB decimal values or hexadecimal values. We can also use an RGBA value, which stands for Red, Green, Blue, Alpha. Here we add some extra information, regarding transparency. Alpha value of 255 defines full opacity, 0 is for full transparency, eg the color is invisible.

颜色是指一个代表红（Red）、绿（Green）、蓝（Blue）（RGB）强度值组合的对象，有效的RGB值在0~255之间。我们可以用多种方式定义颜色，最常用的是RGB十进制或者十六进制值。也可以使用RGBA值，表示红（Red）、绿（Green）、蓝（Blue）和透明度（Alpha）。这里我们增加了额外的信息——关于透明度。Alpha值是255表明完全不透明，0是全透明，即颜色不可见。

::

    #!/usr/bin/python
    # -*- coding: utf-8 -*-

    # colors.py

    import sys, random
    from PyQt4 import QtGui, QtCore


    class Example(QtGui.QWidget):
  
        def __init__(self):
            super(Example, self).__init__()

            self.setGeometry(300, 300, 350, 280)
            self.setWindowTitle('Colors')

        def paintEvent(self, e):
      
            qp = QtGui.QPainter()
            qp.begin(self)
        
            self.drawRectangles(qp)
        
            qp.end()
        
        def drawRectangles(self, qp):

            color = QtGui.QColor(0, 0, 0)
            color.setNamedColor('#d4d4d4')
            qp.setPen(color)

            qp.setBrush(QtGui.QColor(255, 0, 0, 80))
            qp.drawRect(10, 15, 90, 60)

            qp.setBrush(QtGui.QColor(255, 0, 0, 160))
            qp.drawRect(130, 15, 90, 60)

            qp.setBrush(QtGui.QColor(255, 0, 0, 255))
            qp.drawRect(250, 15, 90, 60)

            qp.setBrush(QtGui.QColor(10, 163, 2, 55))
            qp.drawRect(10, 105, 90, 60)

            qp.setBrush(QtGui.QColor(160, 100, 0, 255))
            qp.drawRect(130, 105, 90, 60)

            qp.setBrush(QtGui.QColor(60, 100, 60, 255))
            qp.drawRect(250, 105, 90, 60)

            qp.setBrush(QtGui.QColor(50, 50, 50, 255))
            qp.drawRect(10, 195, 90, 60)

            qp.setBrush(QtGui.QColor(50, 150, 50, 255))
            qp.drawRect(130, 195, 90, 60)

            qp.setBrush(QtGui.QColor(223, 135, 19, 255))
            qp.drawRect(250, 195, 90, 60)


    app = QtGui.QApplication(sys.argv)
    ex = Example()
    ex.show()
    app.exec_()

> In our example, we draw 9 colored rectangles. The first row shows a red color, with different alpha values.

在例子中，我们绘制了9个有色矩形，第一行显示红色，具有不同的透明度。

::

    color = QtGui.QColor(0, 0, 0)
    color.setNamedColor('#d4d4d4')
    
> Here we define a color using hexadecimal notation.

这里我们使用十六进制符号定义颜色。

::

    paint.setBrush(QtGui.QColor(255, 0, 0, 80));
    paint.drawRect(10, 15, 90, 60)
    
> Here we define a brush and draw a rectangle. A brush is an elementary graphics object, which is used to draw the background of a shape. The drawRect() method accepts four parameter. The first two are x, y values on the axis. The third and fourth parameters are width and height of the rectangle. The method draws a rectangle using current pen and current brush.

这里我们定义一个画刷并绘制一个矩形，画刷从一个初级图形对象，用来绘制图形的背景。 `drawRect()` 方法接受四个参数。头两个是坐标轴的x和y，第三和第四个是矩形的宽高，该方法使用当前的画笔和画刷绘制一个矩形。

![Colors][colors]

图：颜色

## QPen

> QPen is an elementary graphics object. It is used to draw lines, curves and outlines of rectangles, ellipses, polygons or other shapes.

`QPen` 是初级图形对象，用来绘制线条、曲线和矩形、椭圆、多边形或其他形状的轮廓。

::

    #!/usr/bin/python
    # -*- coding: utf-8 -*-

    # penstyles.py

    import sys
    from PyQt4 import QtGui, QtCore


    class Example(QtGui.QWidget):
  
        def __init__(self):
            super(Example, self).__init__()

            self.setGeometry(300, 300, 280, 270)
            self.setWindowTitle('penstyles')

        def paintEvent(self, e):
      
            qp = QtGui.QPainter()

            qp.begin(self)        
            self.doDrawing(qp)        
            qp.end()
        
        def doDrawing(self, qp):

            pen = QtGui.QPen(QtCore.Qt.black, 2, QtCore.Qt.SolidLine)

            qp.setPen(pen)
            qp.drawLine(20, 40, 250, 40)

            pen.setStyle(QtCore.Qt.DashLine)
            qp.setPen(pen)
            qp.drawLine(20, 80, 250, 80)

            pen.setStyle(QtCore.Qt.DashDotLine)
            qp.setPen(pen)
            qp.drawLine(20, 120, 250, 120)

            pen.setStyle(QtCore.Qt.DotLine)
            qp.setPen(pen)
            qp.drawLine(20, 160, 250, 160)

            pen.setStyle(QtCore.Qt.DashDotDotLine)
            qp.setPen(pen)
            qp.drawLine(20, 200, 250, 200)

            pen.setStyle(QtCore.Qt.CustomDashLine)
            pen.setDashPattern([1, 4, 5, 4])
            qp.setPen(pen)
            qp.drawLine(20, 240, 250, 240)
        

    app = QtGui.QApplication(sys.argv)
    ex = Example()
    ex.show()
    app.exec_()

> In our example, we draw six lines. The lines are drawn in six different pen styles. There are five predefined pen styles. We can create also custom pen styles. The last line is drawn using custom pen style.

该例子中，我们绘制了6条线，使用了不同的画笔样式。其中5个预定义样式，我们也可以创建自定义样式，最后一个使用了自定义的样式。

::

    pen = QtGui.QPen(QtCore.Qt.black, 2, QtCore.Qt.SolidLine)

> We create a QPen object. The color is black. The width is set to 2 pixels, so that we can see the differences between the pen styles. The QtCore.Qt.SolidLine is one of the predefined pen styles.

创建一个 `QPen` 对象，颜色为黑色，宽度为2个像素，以便能够看到各种画笔样式间的不同。 `QtCore.Qt.SolidLine` 是其中一种预定义的画笔样式。

::

    pen.setStyle(QtCore.Qt.CustomDashLine)
    pen.setDashPattern([1, 4, 5, 4])
    qp.setPen(pen)
    
> Here we define a custom pen style. We set a QtCore.Qt.CustomDashLine pen style and call a setDashPattern() method. The list of numbers defines a style. There must be an even number of numbers. Odd numbers define a dash, even numbers space. The greater the number, the greater the space or the dash. Our pattern is 1px dash 4px space 5px dash 4px space etc.

这里我们自定义了一个画笔样式，设置 `QtCore.Qt.CustomDashLine` 样式，并调用 `setDashPattern()` 方法。用一列数字定义样式，必须是一个偶数序列，技术定义破折号，而是定义间隔，数字越大，间隔或破折号越大。我们的样式是1像素的破折号、3像素的间隔、5像素的破折号和4像素的间隔。

[Pen Styles][pen-styles]

图：画笔样式

## QBrush

> QBrush is an elementary graphics object. It is used to paint the background of graphics shapes, such as rectangles, ellipses or polygons. A brush can be of three different types. A predefined brush a gradient or a texture pattern.

`QBrush` 是初级图形对象，用来绘制图形的背景，如：矩形，椭圆或多边形。画刷有三种类型。可以是预定义的渐变或纹理图案。

::

    #!/usr/bin/python
    # -*- coding: utf-8 -*-

    # brushes.py

    import sys
    from PyQt4 import QtGui, QtCore


    class Example(QtGui.QWidget):
  
        def __init__(self):
            super(Example, self).__init__()

            self.setGeometry(300, 300, 355, 280)
            self.setWindowTitle('Brushes')

        def paintEvent(self, e):
      
            qp = QtGui.QPainter()
        
            qp.begin(self)
            self.drawBrushes(qp)
            qp.end()
        
        def drawBrushes(self, qp):

            brush = QtGui.QBrush(QtCore.Qt.SolidPattern)
            qp.setBrush(brush)
            qp.drawRect(10, 15, 90, 60)

            brush.setStyle(QtCore.Qt.Dense1Pattern)
            qp.setBrush(brush)
            qp.drawRect(130, 15, 90, 60)

            brush.setStyle(QtCore.Qt.Dense2Pattern)
            qp.setBrush(brush)
            qp.drawRect(250, 15, 90, 60)

            brush.setStyle(QtCore.Qt.Dense3Pattern)
            qp.setBrush(brush)
            qp.drawRect(10, 105, 90, 60)

            brush.setStyle(QtCore.Qt.DiagCrossPattern)
            qp.setBrush(brush)
            qp.drawRect(10, 105, 90, 60)

            brush.setStyle(QtCore.Qt.Dense5Pattern)
            qp.setBrush(brush)
            qp.drawRect(130, 105, 90, 60)

            brush.setStyle(QtCore.Qt.Dense6Pattern)
            qp.setBrush(brush)
            qp.drawRect(250, 105, 90, 60)

            brush.setStyle(QtCore.Qt.HorPattern)
            qp.setBrush(brush)
            qp.drawRect(10, 195, 90, 60)

            brush.setStyle(QtCore.Qt.VerPattern)
            qp.setBrush(brush)
            qp.drawRect(130, 195, 90, 60)

            brush.setStyle(QtCore.Qt.BDiagPattern)
            qp.setBrush(brush)
            qp.drawRect(250, 195, 90, 60)


    app = QtGui.QApplication(sys.argv)
    ex = Example()
    ex.show()
    app.exec_()

> In our example, we draw nine different rectangles.

该例子中，我们绘制了9个不同的矩形。

::

    brush = QtGui.QBrush(QtCore.Qt.SolidPattern)
    qp.setBrush(brush)
    qp.drawRect(10, 15, 90, 60)

> We define a brush object. Set it to the painter object. And draw the rectangle calling the drawRect() method.

定义画刷对象，并设置到绘画者对象，调用 `drawRect` 方法绘制矩形。

![Brushes][brushes]

图：画刷

> In this part of the PyQt4 tutorial, we did some basic painting.

在PyQt4教程的这部分中，我们做了一些基本的绘画。
