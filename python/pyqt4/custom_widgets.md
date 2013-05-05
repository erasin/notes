# 自定义窗口部件
> **Custom Widgets in PyQt4**

> Have you ever looked at an application and wondered, how a particular GUI item was created? Probably every wannabe programmer has. Then you were looking at a list of widgets provided by your favourite GUI library. But you couldn't find it. Toolkits usually provide only the most common widgets like buttons, text widgets, sliders etc. No toolkit can provide all possible widgets.

你是否曾经看着应用程序并思考特定的GUI项是如何产生的？大概每位程序员都这样过。然后你能看到你喜欢的GUI库提供的一系列窗口部件，但是你无法找到它。工具包通常仅仅提供最常用的窗口部件，比如按钮、文本组件、滑块等等。没有工具包能够提供一切可能的组件。

> There are actually two kinds of toolkits. Spartan toolkits and heavy weight toolkits. The FLTK toolkit is a kind of a spartan toolkit. It provides only the very basic widgets and assumes, that the programemer will create the more complicated ones himself. PyQt4 is a heavy weight one. It has lots of widgets. Yet it does not provide the more specialized widgets. For example a speed meter widget, a widget that measures the capacity of a CD to be burned (found e.g. in nero). Toolkits also don't have usually charts.

实际上有两种工具包，轻量级和重量级。FLTK工具包是一种轻量级的工具包，它仅仅提供非常基本的组件并假设程序员能够自己创建更复杂的组件。PyQt4属于重量级，它有很多窗口部件，但是并不提供非常专业化的窗口部件。比如速度计窗口部件，用来度量烧录的CD的容量（可在Nero中找到）。也没有包含常用的图表。

> Programmers must create such widgets by themselves. They do it by using the drawing tools provided by the toolkit. There are two possibilities. A programmer can modify or enhance an existing widget. Or he can create a custom widget from scratch.

程序员必须自己创建这些窗口部件，通过工具包提供的绘画工具来创建。有两种方法，修改或增强已有的组件，或者从零开始创建。


##　烧录窗口部件
> **Burning widget** 

> This is a widget that we can see in Nero, K3B or other CD/DVD burning software.

这是我们在Nero、K3B或其他CD烧录软件中看到的窗口部件。

::

    #!/usr/bin/python
    # -*- coding: utf-8 -*-

    """
    ZetCode PyQt4 tutorial

    In this program, we create a custom
    Burning widget. 

    author: Jan Bodnar
    website: zetcode.com
    last edited: December 2010
    """


    import sys
    from PyQt4 import QtGui, QtCore


    class BurningWidget(QtGui.QWidget):
  
        def __init__(self):      
            super(BurningWidget, self).__init__()
        
            self.initUI()
        
        def initUI(self):
        
            self.setMinimumSize(1, 30)
            self.value = 75
            self.num = [75, 150, 225, 300, 375, 450, 525, 600, 675]
        
            self.connect(self, QtCore.SIGNAL("updateBurningWidget(int)"), 
                self.setValue)


        def setValue(self, value):

            self.value = value


        def paintEvent(self, e):
      
            qp = QtGui.QPainter()
            qp.begin(self)
            self.drawWidget(qp)
            qp.end()
      
      
        def drawWidget(self, qp):
      
            font = QtGui.QFont('Serif', 7, QtGui.QFont.Light)
            qp.setFont(font)

            size = self.size()
            w = size.width()
            h = size.height()

            step = int(round(w / 10.0))


            till = int(((w / 750.0) * self.value))
            full = int(((w / 750.0) * 700))

            if self.value >= 700:
                qp.setPen(QtGui.QColor(255, 255, 255))
                qp.setBrush(QtGui.QColor(255, 255, 184))
                qp.drawRect(0, 0, full, h)
                qp.setPen(QtGui.QColor(255, 175, 175))
                qp.setBrush(QtGui.QColor(255, 175, 175))
                qp.drawRect(full, 0, till-full, h)
            else:
                qp.setPen(QtGui.QColor(255, 255, 255))
                qp.setBrush(QtGui.QColor(255, 255, 184))
                qp.drawRect(0, 0, till, h)


            pen = QtGui.QPen(QtGui.QColor(20, 20, 20), 1, 
                QtCore.Qt.SolidLine)
            
            qp.setPen(pen)
            qp.setBrush(QtCore.Qt.NoBrush)
            qp.drawRect(0, 0, w-1, h-1)

            j = 0

            for i in range(step, 10*step, step):
          
                qp.drawLine(i, 0, i, 5)
                metrics = qp.fontMetrics()
                fw = metrics.width(str(self.num[j]))
                qp.drawText(i-fw/2, h/2, str(self.num[j]))
                j = j + 1
            

    class Example(QtGui.QWidget):
  
        def __init__(self):
            super(Example, self).__init__()

            self.initUI()
        
        def initUI(self):

            slider = QtGui.QSlider(QtCore.Qt.Horizontal, self)
            slider.setFocusPolicy(QtCore.Qt.NoFocus)
            slider.setRange(1, 750)
            slider.setValue(75)
            slider.setGeometry(30, 40, 150, 30)

            self.wid = BurningWidget()

            self.connect(slider, QtCore.SIGNAL('valueChanged(int)'), 
                self.changeValue)
            hbox = QtGui.QHBoxLayout()
            hbox.addWidget(self.wid)
            vbox = QtGui.QVBoxLayout()
            vbox.addStretch(1)
            vbox.addLayout(hbox)
            self.setLayout(vbox)

            self.setGeometry(300, 300, 300, 220)
            self.setWindowTitle('Burning')

        def changeValue(self, value):
             
            self.wid.emit(QtCore.SIGNAL("updateBurningWidget(int)"), value)
            self.wid.repaint()


    def main():
  
        app = QtGui.QApplication(sys.argv)
        ex = Example()
        ex.show()
        app.exec_()  


    if __name__ == '__main__':
        main()

> In our example, we have a QSlider and a custom widget. The slider controls the custom widget. This widget shows graphically the total capacity of a medium and the free space available to us. The minimum value of our custom widget is 1, the maximum is 750. If we reach value 700, we begin drawing in red colour. This normally indicates overburning.

该例子中，我们又一个 `QSlider` 和一个自定义窗口部件，滑块用来控制自定义窗口部件。该窗口部件图形化的显示媒体的总量和可用的空余空间。自定义的车口不见的最小值是1，最大值是750。如果到达700，将开始绘制红色。通常表示超刻。

> The burning widget is placed at the bottom of the window. This is achieved using one QHBoxLayout and one QVBoxLayout

烧录部件通常放在窗体的下部，使用一个 `QHBoxLayout` 和一个 `QVBoxLayout` 来达到目的。

::

    class BurningWidget(QtGui.QWidget):
  
        def __init__(self):      
            super(BurningWidget, self).__init__()

> The burning widget it based on the QLabel widget.

烧录窗口部件基于 `QLabel` 窗口部件。

::

    self.setMinimumSize(1, 30)

> We change the minimum size (height) of the widget. The default value is a bit small for us.

修改窗口部件最小值（高度），默认值对于我们有点小。

::

    font = QtGui.QFont('Serif', 7, QtGui.QFont.Light)
    paint.setFont(font)

> We use a smaller font than the default one. That better suits our needs.

使用比默认值更小的字体，更适合我们的需求。

::

    size = self.size()
    w = size.width()
    h = size.height()

    step = int(round(w / 10.0))

    till = int(((w / 750.0) * self.value))
    full = int(((w / 750.0) * 700))

> We draw the widget dynamically. The greater the window, the greater the burning widget. And vice versa. That is why we must calculate the size of the widget onto which we draw the custom widget. The till parameter determines the total size to be drawn. This value comes from the slider widget. It is a proportion of the whole area. The full parameter determines the point, where we begin to draw in red color. Notice the use of floating point arithmetics. This is to achieve greater precision.

我们动态的绘制窗口部件，当窗口变大时，烧录窗口部件也跟着变大，反之亦然。这正是为什么我们必须计算部件的尺寸到我们自定义的部件。 `till` 参数决定绘制的所有尺寸。 `value` 从滑块获得，这是整个区域的比例。 `full` 参数决定我们开始绘制红色的点。注意使用浮点算术，为了实现更高的精度。

> The actual drawing consists of three steps. We draw the yellow or red and yellow rectangle. Then we draw the vertical lines, which divide the widget into several parts. Finally, we draw the numbers, which indicate the capacity of the medium.

实际的绘制由三步组成。我们绘制黄色或红色和黄色矩形，然后绘制垂直线条，把部件分割多部分，最后绘制数字，用来指示媒体的容量。

::

    metrics = qp.fontMetrics()
    fw = metrics.width(str(self.num[j]))
    qp.drawText(i-fw/2, h/2, str(self.num[j]))

> We use font metrics to draw the text. We must know the width of the text in order to center it around the vertical line.

我们使用自体度量来绘制文本，我们必须知道文本的宽度，这样才能在垂直线的中间绘制文本。

::

    def changeValue(self, value):
          
        self.wid.emit(QtCore.SIGNAL("updateBurningWidget(int)"), value)
        self.wid.repaint()

> When we move the slider, the changeValue() method is called. Inside the method, we send a custom updateBurningWidget(int) signal with a parameter. The parameter is the current value of the slider. The value is later used to calculate the capacity of the Burning widget to be drawn.

当我们移动滑块， `changeValue()` 方法被调用，在该方法中我们发送一个自定义的 `updateBurningWidget(int)` 信号及相应参数，参数是滑块的当前值。该值稍后用来计算烧录部件的容量并绘制。

![The burning widget][buring]

图：烧录窗口部件

> In this part of the PyQt4 tutorial, we created a custom widget.

在PyQt4教程的这部分中，我们创建了一个自定义窗口部件。


[buring]: 
