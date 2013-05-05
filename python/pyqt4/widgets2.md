
# 窗口组件 II
> **PyQt4 Widgets II**

> Here we will continue introducing PyQt4 widgets. We will cover QPixmap, QLineEdit, QSplitter and QComboBox.

我们继续介绍PyQt4的窗口组件。我们将涵盖 `QPixmap` ， `QLineEdit` ， `QSplitter` ， `QComboBox` 。

## QPixmap

> QPixmap is one of the widgets used to work with images. It is optimized for showing images on screen. In our code example, we will use QPixmap to display an image on the window.

`QPixmap` 是处理图像的窗口部件之一，非常适合在屏幕上显示图像。在我们的代码示例里，我们使用 `QPixmap` 在窗口中显示图像。

::

    #!/usr/bin/python
    # -*- coding: utf-8 -*-

    # ZetCode PyQt4 tutorial
    #
    # In this example, we show
    # an image on the window.
    # 
    # author: Jan Bodnar
    # website: zetcode.com
    # last edited: December 2010


    from PyQt4 import QtGui

    class Example(QtGui.QWidget):
  
        def __init__(self):
            super(Example, self).__init__()

            self.initUI()

        def initUI(self):

            hbox = QtGui.QHBoxLayout(self)
            pixmap = QtGui.QPixmap("rotunda.jpg")

            label = QtGui.QLabel(self)
            label.setPixmap(pixmap)

            hbox.addWidget(label)
            self.setLayout(hbox)

            self.setWindowTitle("Rotunda in Skalica")
            self.move(250, 200)


    def main():
      
        app = QtGui.QApplication([])
        exm = Example()
        exm.show()
        app.exec_()
  

    if __name__ == '__main__':
        main()
    
> In our example, we display an image on the window.

该例子中，我们在窗口中显示图像。

::

    pixmap = QtGui.QPixmap("rotunda.jpg")

> We create a QPixmap object. It takes the name of the file as a parameter.

创建一个 `QPixmap` 对象，用文件名作为参数。

::

    label = QtGui.QLabel(self)
    label.setPixmap(pixmap)

> We put the pixmap into the QLabel widget.

把 `pixmap` 放入 `QLabel` 窗口部件。

> Figure: QPixmap

图：QPixmap

## QLineEdit

> QLineEdit is a widget that allows to enter and edit a single line of plain text. There are undo/redo, cut/paste and drag & drop functions available for QLineEdit widget.

`QLineEdit` 窗口部件用来输入或者编辑单行纯文本，有撤销/重做，剪切/粘贴和拖放功能。

::

    #!/usr/bin/python
    # -*- coding: utf-8 -*-

    # ZetCode PyQt4 tutorial
    #
    # This example shows text which 
    # is entered in a QLineEdit
    # in a QLabel widget.
    # 
    # author: Jan Bodnar
    # website: zetcode.com
    # last edited: December 2010



    from PyQt4 import QtGui
    from PyQt4 import QtCore


    class Example(QtGui.QWidget):
        def __init__(self, parent=None):
            QtGui.QWidget.__init__(self, parent)

            self.initUI()


        def initUI(self):

            self.label = QtGui.QLabel(self)
            edit = QtGui.QLineEdit(self)
        
            edit.move(60, 100)
            self.label.move(60, 40)

            self.connect(edit, QtCore.SIGNAL('textChanged(QString)'), 
                self.onChanged)

            self.setWindowTitle('QLineEdit')
            self.setGeometry(250, 200, 350, 250)
        

        def onChanged(self, text):
            self.label.setText(text)
            self.label.adjustSize()


    def main():
  
        app = QtGui.QApplication([])
        exm = Example()
        exm.show()
        app.exec_()  


    if __name__ == '__main__':
        main()
        
> This example shows a line edit widget and a label. The text that we key in the line edit is displayed immediately in the label widget.

该例子现在一个单行编辑器和一个标签。在单行编辑器中键入的文字会立即显示在标签中。

::

    edit = QtGui.QLineEdit(self)
    
> The QLineEdit widget is created.

创建 `QLineEdit` 。

::

    self.connect(edit, QtCore.SIGNAL('textChanged(QString)'), 
        self.onChanged)

> If the text in the line edit widget changes, we call the onChanged() method.

如果单行编辑器中的文本发生变化，调用 `onChanged()` 方法。

::

    def onChanged(self, text):
        self.label.setText(text)
        self.label.adjustSize()

> Inside the onChanged() method, we set the typed text to the label widget. We call the adjustSize() method to adjust the size of the label to the length of the text.

在 `onChanged()` 方法中，我们把输入的文字设置到标签中。并调用 `adjustSize()` 方法调整标签的尺寸为文本的长度。

![QLineEdit][qlineedit]

图：QLineEdit

## QSplitter

> QSplitter lets the user control the size of child widgets by dragging the boundary between the children. In our example, we show three QFrame widgets organized with two splitters.

`QSplitter` 使得用户可以通过拖动子窗口部件的边界来控制子窗口部件的尺寸。在我们的例子中，我们显示由两个分离器组织的三个 `QFrame` 窗口部件。

::

    #!/usr/bin/python
    # -*- coding: utf-8 -*-

    # ZetCode PyQt4 tutorial
    #
    # This example shows
    # how to use QSplitter widget.
    # 
    # author: Jan Bodnar
    # website: zetcode.com
    # last edited: December 2010


    from PyQt4 import QtGui, QtCore


    class Example(QtGui.QWidget):
  
        def __init__(self):
            super(Example, self).__init__()

            self.initUI()


        def initUI(self):

            hbox = QtGui.QHBoxLayout(self)

            topleft = QtGui.QFrame(self)
            topleft.setFrameShape(QtGui.QFrame.StyledPanel)
 
            topright = QtGui.QFrame(self)
            topright.setFrameShape(QtGui.QFrame.StyledPanel)

            bottom = QtGui.QFrame(self)
            bottom.setFrameShape(QtGui.QFrame.StyledPanel)

            splitter1 = QtGui.QSplitter(QtCore.Qt.Horizontal)
            splitter1.addWidget(topleft)
            splitter1.addWidget(topright)

            splitter2 = QtGui.QSplitter(QtCore.Qt.Vertical)
            splitter2.addWidget(splitter1)
            splitter2.addWidget(bottom)

            hbox.addWidget(splitter2)
            self.setLayout(hbox)

            self.setWindowTitle('QSplitter')
            QtGui.QApplication.setStyle(QtGui.QStyleFactory.create('Cleanlooks'))
            self.setGeometry(250, 200, 350, 250)
        

    def main():

        app = QtGui.QApplication([])
        ex = Example()
        ex.show()
        app.exec_()


    if __name__ == '__main__':
        main()
        
> In our example we have three frame widgets and two splitters.

例子中有三个框架窗口部件和两个分离器。

::

    topleft = QtGui.QFrame(self)
    topleft.setFrameShape(QtGui.QFrame.StyledPanel)

> We use a styled frame in order to see boundaries between the QFrame widgets.

为了看到 `QFrame` 之间的边界，我们使用带样式的框架。

::

    splitter1 = QtGui.QSplitter(QtCore.Qt.Horizontal)
    splitter1.addWidget(topleft)
    splitter1.addWidget(topright)

> We create a QSplitter widget and add two frames into it.

创建一个 `QSplitter` 并加入两个 框架。

::

    splitter2 = QtGui.QSplitter(QtCore.Qt.Vertical)
    splitter2.addWidget(splitter1)

> We can also add splitter to another splitter widget.

也可以把一个分离器加到另一个分离器窗口部件中。

::

    QtGui.QApplication.setStyle(QtGui.QStyleFactory.create('Cleanlooks'))

> We use a Cleanlooks style. In some styles the frames are not visible.

使用 `Cleanlooks` 样式。 在某些样式中，框架是不可见的。
    
![QSplitter widget][qsplitter]

图：QSplitter 窗口部件

## QComboBox

> The QComboBox is a widget that allows the user to choose from a list of options.

`QComboBox` 窗口部件允许用户从列表清单中选择。

::

    #!/usr/bin/python
    # -*- coding: utf-8 -*-

    # ZetCode PyQt4 tutorial
    #
    # In this example, we show how to 
    # use the QComboBox widget.
    # 
    # author: Jan Bodnar
    # website: zetcode.com
    # last edited: December 2010


    from PyQt4 import QtGui, QtCore


    class Example(QtGui.QWidget):
  
        def __init__(self):
            super(Example, self).__init__()

            self.initUI()


        def initUI(self):

            self.label = QtGui.QLabel("Ubuntu", self)

            combo = QtGui.QComboBox(self)
            combo.addItem("Ubuntu")
            combo.addItem("Mandriva")
            combo.addItem("Fedora")
            combo.addItem("Red Hat")
            combo.addItem("Gentoo")

            combo.move(50, 50)
            self.label.move(50, 150)

            self.connect(combo, QtCore.SIGNAL('activated(QString)'), 
                self.onActivated)

            self.setGeometry(250, 200, 350, 250)
            self.setWindowTitle('QComboBox')

        def onActivated(self, text):
      
            self.label.setText(text)
            self.label.adjustSize()


    def main():
  
        app = QtGui.QApplication([])
        ex = Example()
        ex.show()
        app.exec_()    


    if __name__ == '__main__':
        main()
    
> The example shows a QComboBox and a QLabel. The combo box has a list of six options. These are the names of Linux distros. The label widget shows the selected option from the combo box.

这个例子中显示一个 `QComboBox` 和一个 `QLabel` 。组合框有5个选项的列表，他们是Linux发行版的名称。标签显示从组合框选择的内容。

::

    combo = QtGui.QComboBox(self)
    combo.addItem("Ubuntu")
    combo.addItem("Mandriva")
    combo.addItem("Fedora")
    combo.addItem("Red Hat")
    combo.addItem("Gentoo")

> We create a QComboBox widget and add five options into it.

创建一个 `QComboBox` 窗口部件并增加5个选项。

::

    self.connect(combo, QtCore.SIGNAL('activated(QString)'), 
        self.onActivated)

> Upon an item selection, we call the onActivated() method.

当一个选项被选择，我们调用 `onActivated()` 方法。

::

    def onActivated(self, text):

        self.label.setText(text)
        self.label.adjustSize()
        
> Inside the method, we set the text of the chosen item to the label widget. We adjust the size of the label.

在该方法中，我们把选择项设置到标签中，并调整标签的尺寸。

![QComboBox][qcombobox]

图：QComboBox

> In this part of the PyQt4 tutorial, we covered other four PyQt4 widgets.

在PyQt4教程的这部分中，我们涵盖了其他4个PyQt4窗口部件。
