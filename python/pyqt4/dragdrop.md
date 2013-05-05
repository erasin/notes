# 拖放
> **Drag and Drop in PyQt4**

> In this part of the PyQt4 tutorial, we will talk about drag & drop operations.

在PyQt4教程的这部分中，我们讨论拖放操作。

> In computer graphical user interfaces, drag-and-drop is the action of (or support for the action of) clicking on a virtual object and dragging it to a different location or onto another virtual object. In general, it can be used to invoke many kinds of actions, or create various types of associations between two abstract objects. (Wikipedia)

拖放（Drag-and-drop）指的是图形用户界面（Graphical user interface）中，在一个虚拟的对象上按着鼠标键将之拖曳到另一个地方或另一虚拟对象之上的动作（或是支持着这样的界面的技术）。一般而言，这可以用来产生很多动作，或是在两个抽象对象当中产生各式各样的连接。

> Drag and drop functionality is one of the most visible aspects of the graphical user interface. Drag and drop operation enables users to do complex things intuitively.

拖放操作功能是图形用户界面最明显的方面之一。拖放操作能使用户直观的做复杂的事情。

> Usually, we can drag and drop two things. Data or some graphical objects. If we drag an image from one application to another, we drag and drop binary data. If we drag a tab in Firefox and move it to another place, we drag and drop a graphical component.

通常，我们可以拖放两类东西：数据或者一些图形对象。如果我们从一个程序拖动图片到另一个，我们实际拖动了二进制数据。如果我们在Firefox中拖动标签页到另一个地方，我们拖动了一个可视化组件。

## 简单拖动
> **Simple Drag and Drop**

> In the first example, we will have a QLineEdit and a QPushButton. We will drag plain text from the line edit widget and drop it onto the button widget.

在第一个例子中，有一个 `QLineEdit` 和一个 `QPushButton` 。我们将从单行编辑器中拖动纯文本到按钮上。

::

    #!/usr/bin/python
    # -*- coding: utf-8 -*-

    """
    ZetCode PyQt4 tutorial

    This is a simple drag and
    drop example. 

    author: Jan Bodnar
    website: zetcode.com
    last edited: December 2010
    """

    import sys
    from PyQt4 import QtGui

    class Button(QtGui.QPushButton):
  
        def __init__(self, title, parent):
            super(Button, self).__init__(title, parent)
        
            self.setAcceptDrops(True)

        def dragEnterEvent(self, e):
      
            if e.mimeData().hasFormat('text/plain'):
                e.accept()
            else:
                e.ignore() 

        def dropEvent(self, e):
            self.setText(e.mimeData().text()) 


    class Example(QtGui.QWidget):
  
        def __init__(self):
            super(Example, self).__init__()
        
            self.initUI()
        
        def initUI(self):

            edit = QtGui.QLineEdit('', self)
            edit.setDragEnabled(True)
            edit.move(30, 65)

            button = Button("Button", self)
            button.move(190, 65)
        
            self.setWindowTitle('Simple Drag & Drop')
            self.setGeometry(300, 300, 300, 150)


    def main():
  
        app = QtGui.QApplication(sys.argv)
        ex = Example()
        ex.show()
        app.exec_()  
  

    if __name__ == '__main__':
        main()
        
> Simple drag & drop operation.

简单拖放操作。

::

    class Button(QtGui.QPushButton):
  
        def __init__(self, title, parent):
            super(Button, self).__init__(title, parent)

> In order to drop text on the QPushButton widget, we must reimplement some methods. So we create our own Button class, which will inherit from the QPushButton class.

为了把文本放到 `QPushButton` 窗口部件上，必须重新实现一些方法。我们创建自己的 `Button` 类，从 `QPushButton` 继承。


::

    self.setAcceptDrops(True)
    
> We enable drop events for the widget.

为窗口部件设置可以接受放入事件。

::

    def dragEnterEvent(self, e):
  
        if e.mimeData().hasFormat('text/plain'):
            e.accept()
        else:
            e.ignore() 

> First we reimplement the dragEnterEvent() method. We inform about the data type, we will accept. In our case it is plain text.

首先重新实现 `dragEnterEvent()` 方法。接受了解的数据类型，在该例子中是纯文本。

::

    def dropEvent(self, e):
        self.setText(e.mimeData().text()) 
        
> By reimplementing the dropEvent() method, we will define, what we will do upon the drop event. Here we change the text of the button widget.

在重新实现 `dropEvent()` 方法中，我们定义在拖入事件中处理什么任务。这里我们修改按钮的文字。

::

    edit = QtGui.QLineEdit('', self)
    edit.setDragEnabled(True)
    
> The QLineEdit widget has a built-in support for drag operations. All we need to do is to call setDragEnabled() method to activate it.

`QLineEdit` 窗口部件有内置拖动操作支持，我们需要做的是调用 `setDragEnabled` 方法并激它。

> Figure: Simple Drag & Drop

图：简单拖放


## 拖放按钮
> **Drag & drop a button widget**

> In the following example, we will demonstrate, how to drag & drop a button widget.

在接下来的例子中，我们将演示如何拖放一个按钮窗口部件。

::

    #!/usr/bin/python
    # -*- coding: utf-8 -*-

    """
    ZetCode PyQt4 tutorial

    In this program, we can press
    on a button with a left mouse
    click or drag and drop the button 
    with  the right mouse click. 

    author: Jan Bodnar
    website: zetcode.com
    last edited: December 2010
    """

    import sys
    from PyQt4 import QtGui
    from PyQt4 import QtCore


    class Button(QtGui.QPushButton):
  
        def __init__(self, title, parent):
            super(Button, self).__init__(title, parent)

        def mouseMoveEvent(self, e):

            if e.buttons() != QtCore.Qt.RightButton:
                return

            mimeData = QtCore.QMimeData()

            drag = QtGui.QDrag(self)
            drag.setMimeData(mimeData)
            drag.setHotSpot(e.pos() - self.rect().topLeft())

            dropAction = drag.start(QtCore.Qt.MoveAction)


        def mousePressEvent(self, e):
      
            QtGui.QPushButton.mousePressEvent(self, e)
            if e.button() == QtCore.Qt.LeftButton:
                print 'press'


    class Example(QtGui.QWidget):
  
        def __init__(self):
            super(Example, self).__init__()

            self.initUI()
        
        def initUI(self):

            self.setAcceptDrops(True)

            self.button = Button('Button', self)
            self.button.move(100, 65)

            self.setWindowTitle('Click or Move')
            self.setGeometry(300, 300, 280, 150)

        def dragEnterEvent(self, e):
      
            e.accept()

        def dropEvent(self, e):

            position = e.pos()
            self.button.move(position)

            e.setDropAction(QtCore.Qt.MoveAction)
            e.accept()

    def main():
  
        app = QtGui.QApplication(sys.argv)
        ex = Example()
        ex.show()
        app.exec_()  


    if __name__ == '__main__':
        main()
    
> In our code example, we have a QPushButton on the window. If we click on the button with a left mouse button, we print 'press' to the console. By right clicking and moving the button, we perform a drag & drop operation on the button widget.

该例子中，在窗口上有一个 `QPushButton` ，如果我们用鼠标左键按下按钮，在控制台打印“press”。用右键按下并移动按钮，对按钮执行拖放操作。

::

    class Button(QtGui.QPushButton):
  
        def __init__(self, title, parent):
            super(Button, self).__init__(title, parent)

> We create a Button class, which will derive from the QPushButton. We also reimplement two methods of the QPushButton. mouseMoveEvent() and mousePressEvent(). The mouseMoveEvent() method is the place, where the drag & drop operation begins.

创建一个 `Button` 类，从 `QPushButton` 继承。重新实现 `QPushButton` 的两个方法： `mouseMoveEvent()` 和 `mousePressEvent()` 。 `mouseMoveEvent()` 方法是拖放操作开始的地方。

::

    if event.buttons() != QtCore.Qt.RightButton:
        return

> Here we decide, that we can perform drag & drop only with a right mouse button. The left mouse button is reserved for clicking on the button.

这里我们确定只能使用鼠标右键执行拖放操作，鼠标左键盘保留给点击操作。

::

    mimeData = QtCore.QMimeData()

    drag = QtGui.QDrag(self)
    drag.setMimeData(mimeData)
    drag.setHotSpot(event.pos() - self.rect().topLeft())

> Here we create a QDrag object.

创建一个 `QDrag` 对象。

::

    dropAction = drag.start(QtCore.Qt.MoveAction)

> The start() method of the drag object starts the drag & drop operation.

拖动对象的 `starts()` 方法开始拖动操作。

::

    def mousePressEvent(self, e):
  
        QtGui.QPushButton.mousePressEvent(self, e)
        if e.button() == QtCore.Qt.LeftButton:
            print 'press'
        
> We print 'press' to the console, if we left click on the button with the mouse. Notice that we call mousePressEvent() method on the parent as well. Otherwise we would not see the button being pushed.

如果按下鼠标左键，在控制台打印“press”。注意我们同样调用了父类的 `mousePressEvent()` 方法，否则我们将看不到按钮被按下。

::

    position = e.pos()

    self.button.move(position)

> In the dropEvent() method we code, what happens after we release the mouse button and finish the drop operation. We find out the current mouse pointer position and move a button accordingly.

在 `dropEvent()` 方法中编写在释放鼠标并且结束拖入操作后将发生什么。找出当前鼠标点的位置并把按钮移到相应位置。

::

    e.setDropAction(QtCore.Qt.MoveAction)
    e.accept()

> We specify the type of the drop action. In our case it is a move action.

指定拖入动作的类型，这里是移动动作。

This part of the PyQt4 tutorial was dedicated to drag and drop.

在PyQt4教程的这部分中，我们专注于拖放操作。
