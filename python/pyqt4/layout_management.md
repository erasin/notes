# 布局管理

> **Layout management in PyQt4**

> Important thing in programming is the layout management. Layout management is the way how we place the widgets on the window. The management can be done in two ways. We can use absolute positioning or layout classes.
    
编程中的一个重要事情是布局管理，布局管理是如何在窗体上摆放窗口部件。可以有两种方式进行管理：绝对定位或使用布局类。

## 绝对定位
> **Absolute positioning**

> The programmer specifies the position and the size of each widget in pixels. When you use absolute positioning, you have to understand several things.

程序员用像素指定每个控件的位置和尺寸。使用绝对定位时，你必须理解几件事情。

> * the size and the position of a widget do not change, if you resize a window
    * applications might look different on various platforms
    * changing fonts in your application might spoil the layout
    * if you decide to change your layout, you must completely redo your layout, which is tedious and time consuming

* 如果你调整窗体的大小，组件的尺寸和位置并不会改变
* 在不同的平台上，程序可能看起来不一样
* 改变程序的字体可能破坏布局
* 如果你决定改变你的布局，你必须完全重做你的布局，这将是乏味并且浪费时间的

::

    #!/usr/bin/python
    # -*- coding: utf-8 -*-

    # absolute.py

    import sys
    from PyQt4 import QtGui


    class Example(QtGui.QWidget):
  
        def __init__(self):
            super(Example, self).__init__()
        
            self.initUI()
        
        def initUI(self):
            label1 = QtGui.QLabel('Zetcode', self)
            label1.move(15, 10)

            label2 = QtGui.QLabel('tutorials for programmers', self)
            label2.move(35, 40)

            self.setWindowTitle('Absolute')
            self.resize(250, 150)

    app = QtGui.QApplication(sys.argv)
    ex = Example()
    ex.show()
    sys.exit(app.exec_())
    
> We simply call the move() method to position our widgets. In our case QLabel-s. We position them by providing the x and the y coordinates. The beginning of the coordinate system is at the left top corner. The x values grow from left to right. The y values grow from top to bottom.

我们简单的调用 `move()` 方法来定位组件。在我们的 `QLabel` 例子中，我们用x和y坐标来定位。坐标系统从左上角开始，x值从左到右增长，y值从上到下增长。

> Figure: Absolute positioning

![absolute.png][absolute]

图：绝对定位

    
## 框布局
> **Box Layout**

> Layout management with layout classes is much more flexible and practical. It is the preferred way to place widgets on a window. The basic layout classes are QHBoxLayout and QVBoxLayout. They line up widgets horizontally and vertically.

使用布局类管理布局更灵活、更实用。这是在窗体上摆放组件的首选方式。基本的布局类是 `QHBoxLayout` 和 `QVBoxLayout` ，它们可以横向和纵向排列窗口部件。

> Imagine that we wanted to place two buttons in the right bottom corner. To create such a layout, we will use one horizontal and one vertical box. To create the neccessary space, we will add a stretch factor.

假设我们想要摆放两个按钮到右下角。为了创建这样一个布局，我们需要一个水平框和一个垂直框。我们通过增加 **延展因素** 来创建必要的间隔。

::

    #!/usr/bin/python
    # -*- coding: utf-8 -*-

    # boxlayout.py

    import sys
    from PyQt4 import QtGui


    class Example(QtGui.QWidget):
  
        def __init__(self):
            super(Example, self).__init__()
        
            self.initUI()
        
        def initUI(self):
            okButton = QtGui.QPushButton("OK")
            cancelButton = QtGui.QPushButton("Cancel")

            hbox = QtGui.QHBoxLayout()
            hbox.addStretch(1)
            hbox.addWidget(okButton)
            hbox.addWidget(cancelButton)

            vbox = QtGui.QVBoxLayout()
            vbox.addStretch(1)
            vbox.addLayout(hbox)
        
            self.setLayout(vbox)
        
            self.setWindowTitle('box layout')
            self.resize(300, 150)

    app = QtGui.QApplication(sys.argv)
    ex = Example()
    ex.show()
    sys.exit(app.exec_())
    
::

    okButton = QtGui.QPushButton("OK")
    cancelButton = QtGui.QPushButton("Cancel")
    
> Here we create two push buttons.

这里我们创建两个按钮。

::

    hbox = QtGui.QHBoxLayout()
    hbox.addStretch(1)
    hbox.addWidget(okButton)
    hbox.addWidget(cancelButton)
    
> We create a horizontal box layout. Add a stretch factor and both buttons.

我们创建一个水平框布局，增加一个延展因素和两个按钮。

::

    vbox = QtGui.QVBoxLayout()
    vbox.addStretch(1)
    vbox.addLayout(hbox)

> To create the necessary layout, we put a horizontal layout into a vertical one.

为了创建所需的布局，我们把水平布局放到垂直布局中。

::

    self.setLayout(vbox)
    
> Finally, we set the main layout of the window.

最后，我们设置窗体的主布局。

> Figure: box layout

![boxlayout.png][boxlayout]

图：框布局

    -----------
    
## QGridLayout
> **QGridLayout**

> The most universal layout class is the grid layout. This layout divides the space into rows and columns. To create a grid layout, we use the QGridLayout class.

最常用的布局类是网格布局，网格布局把空间划分成行和列。我们使用 `QGridLayout` 类来创建网格布局。

::

    #!/usr/bin/python
    # -*- coding: utf-8 -*-

    # gridlayout1.py

    import sys
    from PyQt4 import QtGui


    class Example(QtGui.QWidget):
  
        def __init__(self):
            super(Example, self).__init__()
        
            self.initUI()
        
        def initUI(self):
            self.setWindowTitle('grid layout')

            names = ['Cls', 'Bck', '', 'Close', '7', '8', '9', '/',
                '4', '5', '6', '*', '1', '2', '3', '-',
                '0', '.', '=', '+']

            grid = QtGui.QGridLayout()

            j = 0
            pos = [(0, 0), (0, 1), (0, 2), (0, 3),
                    (1, 0), (1, 1), (1, 2), (1, 3),
                    (2, 0), (2, 1), (2, 2), (2, 3),
                    (3, 0), (3, 1), (3, 2), (3, 3 ),
                    (4, 0), (4, 1), (4, 2), (4, 3)]

            for i in names:
                button = QtGui.QPushButton(i)
                if j == 2:
                    grid.addWidget(QtGui.QLabel(''), 0, 2)
                else: grid.addWidget(button, pos[j][0], pos[j][1])
                j = j + 1

            self.setLayout(grid)

    app = QtGui.QApplication(sys.argv)
    ex = Example()
    ex.show()
    sys.exit(app.exec_())
    
> In our example, we create a grid of buttons. To fill one gap, we add one QLabel widget.

在该例子中，我们创建了一个按钮格，增加一个 `QLabel` 窗口部件来填补一个空白。

::

    grid = QtGui.QGridLayout()
    Here we create a grid layout.

::

    if j == 2:
        grid.addWidget(QtGui.QLabel(''), 0, 2)
    else: 
        grid.addWidget(button, pos[j][0], pos[j][1])

> To add a widget to a grid, we call the addWidget() method. The arguments are the widget, the row and the column number.

调用 `addWidget()` 方法来把窗口部件加到网格中，参数是部件（ `widget` ），行（ `row` ）和列（ `column` ）数字。

> Figure: grid layout

![gridlayout1.png][gridlayout1]

图：网格布局

> Widgets can span multiple columns or rows in a grid. In the next example we illustrate this.

组件可以在表格中跨越多列或多行，在下一个例子中我们将演示这个。


::

    #!/usr/bin/python
    # -*- coding: utf-8 -*-

    # gridlayout2.py

    import sys
    from PyQt4 import QtGui


    class Example(QtGui.QWidget):
  
        def __init__(self):
            super(Example, self).__init__()
        
            self.initUI()
        
        def initUI(self):
            title = QtGui.QLabel('Title')
            author = QtGui.QLabel('Author')
            review = QtGui.QLabel('Review')

            titleEdit = QtGui.QLineEdit()
            authorEdit = QtGui.QLineEdit()
            reviewEdit = QtGui.QTextEdit()

            grid = QtGui.QGridLayout()
            grid.setSpacing(10)

            grid.addWidget(title, 1, 0)
            grid.addWidget(titleEdit, 1, 1)

            grid.addWidget(author, 2, 0)
            grid.addWidget(authorEdit, 2, 1)

            grid.addWidget(review, 3, 0)
            grid.addWidget(reviewEdit, 3, 1, 5, 1)
        
            self.setLayout(grid)
        
            self.setWindowTitle('grid layout')
            self.resize(350, 300)
        
    app = QtGui.QApplication(sys.argv)
    ex = Example()
    ex.show()
    sys.exit(app.exec_())
    
> We create a window in which we have three labels, two line edits and one text edit widget. The layout is done with the QGridLayout.

我们创建了一个窗体，上面有三个标签，两个单行编辑框和一个文本编辑框组件，使用 `QGridLayout` 来布局。

::

    grid = QtGui.QGridLayout()
    grid.setSpacing(10)
    
> We create a grid layout and set spacing between widgets.

我们创建了一个表格布局，并且设置组件间的间隔。

::

    grid.addWidget(reviewEdit, 3, 1, 5, 1)

> If we add a widget to a grid, we can provide row span and column span of the widget. In our case, we make the reviewEdit widget span 5 rows.

如果我们增加一个窗口部件到网格中，我们可以提供窗口部件的行跨度和列跨度。在这个例子中，我们设置 `reviewEdit` 占用5行。

> Figure: grid layout2

![gridlayout2.png][gridlayout2]

图：表格布局2

> This part of the PyQt4 tutorial was dedicated to layout management.

PyQt4教程这的部分我们致力于布局管理。

[absolute]:http://jimmykuu.sinaapp.com/static/PyQt4_Tutorial/html/_images/absolute.png
[boxlayout]:404
[gridlayout1]:http://jimmykuu.sinaapp.com/static/PyQt4_Tutorial/html/_images/gridlayout1.png
[gridlayout2]:http://jimmykuu.sinaapp.com/static/PyQt4_Tutorial/html/_images/gridlayout2.png
