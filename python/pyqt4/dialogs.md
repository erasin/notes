# PyQt4的对话框
> **Dialogs in PyQt4**

> Dialog windows or dialogs are an indispensable part of most modern GUI applications. A dialog is defined as a conversation between two or more persons. In a computer application a dialog is a window which is used to "talk" to the application. A dialog is used to input data, modify data, change the application settings etc. Dialogs are important means of communication between a user and a computer program.

对话框窗体或对话框是现代GUI应用不可或缺的一部分。dialog定义为两个或多个人之间的交谈。在计算机程序中dialog是一个窗体，用来和程序“交谈”。对话框用来输入数据、修改数据、改变程序设置等等。对话框是用户和计算机程序沟通的重要手段。

## QInputDialog

> The QInputDialog provides a simple convenience dialog to get a single value from the user. The input value can be a string, a number or an item from a list.

`QInputDialog` 提供一个简单的对话框，以便从用户获取单个值。输入值可以是一个字符串，一个数字或者列表的一项。

::

    #!/usr/bin/python
    # -*- coding: utf-8 -*-

    # inputdialog.py

    import sys
    from PyQt4 import QtGui
    from PyQt4 import QtCore


    class Example(QtGui.QWidget):
  
        def __init__(self):
            super(Example, self).__init__()

            self.initUI()
        
        def initUI(self):

            self.button = QtGui.QPushButton('Dialog', self)
            self.button.setFocusPolicy(QtCore.Qt.NoFocus)

            self.button.move(20, 20)
            self.connect(self.button, QtCore.SIGNAL('clicked()'), 
                self.showDialog)
            self.setFocus()
        
            self.label = QtGui.QLineEdit(self)
            self.label.move(130, 22)
        
            self.setWindowTitle('InputDialog')
            self.setGeometry(300, 300, 350, 80)
        
    
        def showDialog(self):
            text, ok = QtGui.QInputDialog.getText(self, 'Input Dialog', 
                'Enter your name:')
        
            if ok:
                self.label.setText(str(text))


    if __name__ == '__main__':

        app = QtGui.QApplication(sys.argv)
        ex = Example()
        ex.show()
        app.exec_()

> The example has a button and a line edit widget. The button shows the input dialog for getting text values. The entered text will be displayed in the line edit widget.

这个例子有一个按钮和一个单行编辑器，按钮显示输入对话框来获取值，输入的值将会显示在单行编辑器中。

::

    text, ok = QtGui.QInputDialog.getText(self, 'Input Dialog', 
        'Enter your name:')    

> This line displays the input dialog. The first string is a dialog title, the second one is a message within the dialog. The dialog returns the entered text and a boolean value. If we clicked ok button, the boolean value is true, otherwise false.

这行代码显示输入对话框，第一个字符串是对话框标题，第二个是对话框里面的消息。对话框返回输入的文本那一个布尔值。如果我们点击 ok 按钮，布尔值是 `True` ，否则为 `False` 。

![Input Dialog][input-dialog]

图：输入对话框

## QColorDialog

> The QColorDialog provides a dialog widget for specifying colors.

颜色对话框为定制颜色提供一个对话框组件。

::

    #!/usr/bin/python
    # -*- coding: utf-8 -*-

    # colordialog.py

    import sys
    from PyQt4 import QtGui
    from PyQt4 import QtCore


    class Example(QtGui.QWidget):
  
        def __init__(self):
            super(Example, self).__init__()
        
            self.initUI()
        
        def initUI(self):

            color = QtGui.QColor(0, 0, 0) 

            self.button = QtGui.QPushButton('Dialog', self)
            self.button.setFocusPolicy(QtCore.Qt.NoFocus)
            self.button.move(20, 20)

            self.connect(self.button, QtCore.SIGNAL('clicked()'), 
                self.showDialog)
            self.setFocus()

            self.widget = QtGui.QWidget(self)
            self.widget.setStyleSheet("QWidget { background-color: %s }" 
                % color.name())
            self.widget.setGeometry(130, 22, 100, 100)
        
            self.setWindowTitle('ColorDialog')
            self.setGeometry(300, 300, 250, 180)
        

        def showDialog(self):
      
            col = QtGui.QColorDialog.getColor()

            if col.isValid():
                self.widget.setStyleSheet("QWidget { background-color: %s }"
                    % col.name())


    if __name__ == '__main__':
  
        app = QtGui.QApplication(sys.argv)
        ex = Example()
        ex.show()
        app.exec_()
    
> The application example shows a push button and a QWidget. The widget background is set to black color. Using the QColorDialog, we can change its background.

该例子显示一个按钮和一个 `QWidge` 对象。组件的背景设为黑色，我们可以用 `QColorDialog` 来改变它的背景。

::

    color = QtGui.QColorDialog.getColor()
    
> This line will pop up the QColorDialog.

这行代码将会弹出一个 `QColorDialog` 。

::

    if col.isValid():
        self.widget.setStyleSheet("QWidget { background-color: %s }"
            % col.name())

> We check if the color is valid. If we click on the cancel button, no valid color is returned. If the color is valid, we change the background color using stylesheets.

检查颜色是否有效，如果点击了取消按钮，将返回无效的颜色。如果颜色有些，我们使用样式修改背景颜色。

![Color dialog][color-dialog]

图：颜色对话框

## QFontDialog

> The QFontDialog is a dialog widget for selecting font.

`QFontDialog` 是一个用来选择字体的对话框组件。

::

    #!/usr/bin/python
    # -*- coding: utf-8 -*-

    # fontdialog.py

    import sys
    from PyQt4 import QtGui
    from PyQt4 import QtCore


    class Example(QtGui.QWidget):
  
        def __init__(self):
            super(Example, self).__init__()

            self.initUI()
        
        def initUI(self):

            hbox = QtGui.QHBoxLayout()

            button = QtGui.QPushButton('Dialog', self)
            button.setFocusPolicy(QtCore.Qt.NoFocus)
            button.move(20, 20)

            hbox.addWidget(button)

            self.connect(button, QtCore.SIGNAL('clicked()'), self.showDialog)
        
            self.label = QtGui.QLabel('Knowledge only matters', self)
            self.label.move(130, 20)

            hbox.addWidget(self.label, 1)
            self.setLayout(hbox)
        
            self.setWindowTitle('FontDialog')
            self.setGeometry(300, 300, 250, 110)
        
    
        def showDialog(self):

            font, ok = QtGui.QFontDialog.getFont()
            if ok:
                self.label.setFont(font)
        

    if __name__ == '__main__':
  
        app = QtGui.QApplication(sys.argv)
        ex = Example()
        ex.show()
        app.exec_()

In our example, we have a button and a label. With QFontDialog, we change the font of the label.

该例子中，我们有一个按钮和一个标签。我们用 `QFontDialog` 改变标签的字体。

::

    hbox.addWidget(self.label, 1)
    
> We add the label to the horizontal box. We set the stretch factor to 1. When we select a different font, the text may become larger. Otherwise the label might not be fully visible.

我们把标签加入到水平框布局中。设置延展因素为1，当我们选择不同的字体时，文字可能变得更大。否则标签可能显示不完全。

::
    font, ok = QtGui.QFontDialog.getFont()
    
> Here we pop up the font dialog.

这里弹出一个字体对话框。

::

    if ok:
        self.label.setFont(font)

> If we clicked ok, the font of the label was changed.

如果点击确定，标签的字体将会改变。

![Font dialog][font-dialog]

图：字体对话框

## QFileDialog

> The QFileDialog is a dialog that allows users to select files or directories. The files can be selected for both opening and saving.

`QFileDialog` 允许用户选择文件或文件夹，可选择文件来打开和保存。

::

    #!/usr/bin/python
    # -*- coding: utf-8 -*-

    # openfiledialog.py

    import sys
    from PyQt4 import QtGui
    from PyQt4 import QtCore


    class Example(QtGui.QMainWindow):
  
        def __init__(self):
            super(Example, self).__init__()
        
            self.initUI()
        
        def initUI(self):

            self.textEdit = QtGui.QTextEdit()
            self.setCentralWidget(self.textEdit)
            self.statusBar()
            self.setFocus()

            openFile = QtGui.QAction(QtGui.QIcon('open.png'), 'Open', self)
            openFile.setShortcut('Ctrl+O')
            openFile.setStatusTip('Open new File')
            self.connect(openFile, QtCore.SIGNAL('triggered()'), self.showDialog)

            menubar = self.menuBar()
            fileMenu = menubar.addMenu('&File')
            fileMenu.addAction(openFile)
        
            self.setGeometry(300, 300, 350, 300)
            self.setWindowTitle('OpenFile')        
    
        def showDialog(self):

            filename = QtGui.QFileDialog.getOpenFileName(self, 'Open file',
                        '/home')
            fname = open(filename)
            data = fname.read()
            self.textEdit.setText(data)

    app = QtGui.QApplication(sys.argv)
    ex = Example()
    ex.show()
    app.exec_()

> The example shows a menubar, centrally set text edit widget and a statusbar. The statusbar is shown only for desing purposes. The the menu item shows the QFileDialog which is used to select a file. The contents of the file are loaded into the text edit widget.

这个例子显示一个菜单，中间放置一个文本编辑框，还有一个状态栏。状态机仅为了设计目的显示。菜单项显示 `QFileDialog` 来选择文件，文件的内容加载进文本编辑器。

::

    class Example(QtGui.QMainWindow):
  
        def __init__(self):
            super(Example, self).__init__()

> The example is based on the QMainWindow widget, because we centrally set the text edit widget.

这个例子建立在 `QMainWindow` 组件上，因为我们需要在中间设置文本编辑器。

::
    filename = QtGui.QFileDialog.getOpenFileName(self, 'Open file',
                  '/home')
                  
> We pop up the QFileDialog. The first string in the getOpenFileName() method is the caption. The second string specifies the dialog working directory. By default, the file filter is set to All files (*).

我们弹出 `QFileDialog` ， `getOpenFileName` 方法的第一个字符串是标题，第二个字符串指定对话框的工作目录，文件过滤默认设置 `All files(*)` 。

::

    fname = open(filename)
    data = fname.read()
    self.textEdit.setText(data)

> The selected file name is read and the contents of the file are set to the text edit widget.

读取选择的文件，并把文件内容放入文本编辑器。

![File dialog][file-dialog]

图：文件对话框

> In this part of the PyQt4 tutorial, we worked with dialogs.

在PyQt4教程的这部分中，我们介绍了对话框。
