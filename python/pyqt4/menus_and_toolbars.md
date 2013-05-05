# 菜单和工具栏


> **Menus and Toolbars in PyQt4**
> In this part of the PyQt4 tutorial, we will create menus and toolbars.

在PyQt4教程的这部分中，我们将创建菜单和工具栏。
 
## 主窗体
> **Main Window**

> The QMainWindow class provides a main application window. This enables to create the classic application skeleton with a statusbar, toolbars and a menubar.

`QMainWindow` 类提供应用程序主窗口，可以创建一个经典的拥有状态栏、工具栏和菜单栏的应用程序骨架。

## 状态栏
> **Statusbar** 

> The statusbar is a widget that is used for displaying status information.

状态栏是用来显示状态信息的窗口部件。

::

    #!/usr/bin/python

    # statusbar.py 

    import sys
    from PyQt4 import QtGui

    class MainWindow(QtGui.QMainWindow):
        def __init__(self):
            QtGui.QMainWindow.__init__(self)

            self.resize(250, 150)
            self.setWindowTitle('statusbar')

            self.statusBar().showMessage('Ready')


    app = QtGui.QApplication(sys.argv)
    main = MainWindow()
    main.show()
    sys.exit(app.exec_())
  
::

    self.statusBar().showMessage('Ready')
    
> To get the statusbar, we call the statusBar() method of the QApplication class. The showMessage() displays message on the statusbar.

通过掉用 `QMainWindow` 的 `statusBar()` 方法来得到状态栏， `showMessage()` 将消息显示在状态栏上。

![statusbar.png][statusbar]

图：状态栏


## 菜单栏
> Menubar

> A menubar is one of the most visible parts of the GUI application. It is a group of commands located in various menus. While in console applications you had to remember all those arcane commands, here we have most of the commands grouped into logical parts. There are accepted standards that further reduce the amount of time spending to learn a new application.

菜单栏是GUI应用程序最明显的部分之一，这是分布在各个菜单的命令组合，在控制台程序中你需要记住所有那些晦涩难懂的命令，有了这些公认的标准，将进一步缩短学习新应用程序的时间。

::

    #!/usr/bin/python

    # menubar.py 

    import sys
    from PyQt4 import QtGui, QtCore

    class MainWindow(QtGui.QMainWindow):
        def __init__(self):
            QtGui.QMainWindow.__init__(self)

            self.resize(250, 150)
            self.setWindowTitle('menubar')

            exit = QtGui.QAction(QtGui.QIcon('icons/exit.png'), 'Exit', self)
            exit.setShortcut('Ctrl+Q')
            exit.setStatusTip('Exit application')
            self.connect(exit, QtCore.SIGNAL('triggered()'), QtCore.SLOT('close()'))

            self.statusBar()

            menubar = self.menuBar()
            file = menubar.addMenu('&File')
            file.addAction(exit)

    app = QtGui.QApplication(sys.argv)
    main = MainWindow()
    main.show()
    sys.exit(app.exec_())
  
::

    menubar = self.menuBar()
    file = menubar.addMenu('&File')
    file.addAction(exit)
    
> First we create a menubar with the menuBar() method of the QMainWindow class. Then we add a menu with the addMenu() method. In the end we plug the action object into the file menu.
    
首先我们使用`QMainWindow` 类的 `menuBar()` 方法创建一个菜单栏，然后使用 `addMenu()` 方法增加一个菜单项，最后我们插入 `Action` 对象到文件菜单。

note:
:    Mac OS X系统上的菜单栏和其他系统的不同，不再窗口上，而是位于屏幕的顶部，该例子中会自动把退出菜单项和程序的系统菜单进行融合。

![menubar_mac.png][menubar_mac]

图：菜单栏（Mac OS X）
    
![menubar_ubuntu.png][menubar_ubuntu]

图：菜单栏（Ubuntu）

## 工具栏
> **Toolbar**

> Menus group all commands that we can use in an application. Toolbars provide a quick access to the most frequently used commands.
    
菜单集合了应用程序中所有可用的命令，工具栏提供了快速访问最常用的命令功能。

::

    #!/usr/bin/python

    # toolbar.py 

    import sys
    from PyQt4 import QtGui, QtCore

    class MainWindow(QtGui.QMainWindow):
        def __init__(self):
            QtGui.QMainWindow.__init__(self)

            self.resize(250, 150)
            self.setWindowTitle('toolbar')

            self.exit = QtGui.QAction(QtGui.QIcon('icons/exit.png'), 'Exit', self)
            self.exit.setShortcut('Ctrl+Q')
            self.connect(self.exit, QtCore.SIGNAL('triggered()'), QtCore.SLOT('close()'))

            self.toolbar = self.addToolBar('Exit')
            self.toolbar.addAction(self.exit)


    app = QtGui.QApplication(sys.argv)
    main = MainWindow()
    main.show()
    sys.exit(app.exec_())
    
::

    self.exit = QtGui.QAction(QtGui.QIcon('icons/exit.png'), 'Exit', self)
    self.exit.setShortcut('Ctrl+Q')

> GUI applications are controlled with commands. These commands can be launched from a menu, a context menu, a toolbar or with a shortcut. PyQt simplifies development with the introduction of actions. An action object can have menu text, an icon, a shortcut, status text, "What's This?" text and a tooltip. In our example, we define an action object with an icon, a tooltip and a shortcut.
    
GUI程序通过命令来控制，这些命令可以通过菜单、上下文菜单、工具栏或者快捷键来启动。PyQt引入了行为(Action)来简化开发。行为对象可以拥有菜单标题、图标、快捷键、状态栏内容、aasda“这是什么？”内容以及提示。在我们的例子中，我们定义了一个行为对象，包括图标、提示和快捷键。

::

    self.connect(self.exit, QtCore.SIGNAL('triggered()'), QtCore.SLOT('close()'))
    
> Here we connect the action's triggered() signal to the predefined close() signal.
    
这里我们连接行为的 `triggered()` 信号到预定义的 `close()` 槽。

::

    self.toolbar = self.addToolBar('Exit')
    self.toolbar.addAction(self.exit)
    
>
	Here we create a toolbar and plug and action object into it.

这里我们我们创建了一个工具栏并把行为加了进去。

![toolbar.png][toolbar]

图: 工具栏


## 放在一起
> **Putting it together**

> In the last example of this section, we will create a menubar, toolbar and a statusbar. We will also create a central widget.

在这部分的最后一个例子中，我们将会创建一个菜单栏、工具栏和状态栏。

::

    #!/usr/bin/python

    # mainwindow.py 

    import sys
    from PyQt4 import QtGui, QtCore

    class MainWindow(QtGui.QMainWindow):
        def __init__(self):
            QtGui.QMainWindow.__init__(self)

            self.resize(350, 250)
            self.setWindowTitle('mainwindow')

            textEdit = QtGui.QTextEdit()
            self.setCentralWidget(textEdit)

            exit = QtGui.QAction(QtGui.QIcon('icons/exit.png'), 'Exit', self)
            exit.setShortcut('Ctrl+Q')
            exit.setStatusTip('Exit application')
            self.connect(exit, QtCore.SIGNAL('triggered()'), QtCore.SLOT('close()'))

            self.statusBar()

            menubar = self.menuBar()
            file = menubar.addMenu('&File')
            file.addAction(exit)

            toolbar = self.addToolBar('Exit')
            toolbar.addAction(exit)


    app = QtGui.QApplication(sys.argv)
    main = MainWindow()
    main.show()
    sys.exit(app.exec_())
    
::

    textEdit = QtGui.QTextEdit()
    self.setCentralWidget(textEdit)
    
> Here we create a text edit widget. We set it to be the central widget of the QMainWindow. The central widget will occupy all space that is left.

这里我们创建了一个文本编辑控件，把它设置成 `QMainWinow` 的中心组件。中心组件将会占据所有留下的空间。

> Figure: QMainWindow

![mainwindow.png][mainwindow]

图：QMainWindow（Ubuntu）

> In this part of the PyQt4 tutorial, we created menus and toolbars.
    
在PyQt4教程的这部分，我们创建了菜单和工具栏。

[statusbar]:http://jimmykuu.sinaapp.com/static/PyQt4_Tutorial/html/_images/statusbar.png
[menubar_mac]:http://jimmykuu.sinaapp.com/static/PyQt4_Tutorial/html/_images/menubar_mac.png
[menubar_ubuntu]:http://jimmykuu.sinaapp.com/static/PyQt4_Tutorial/html/_images/menubar_ubuntu.png
[toolbar]:http://jimmykuu.sinaapp.com/static/PyQt4_Tutorial/html/_images/toolbar.png
[mainwindow]:http://jimmykuu.sinaapp.com/static/PyQt4_Tutorial/html/_images/mainwindow.png


