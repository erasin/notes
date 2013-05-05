# 第一个程序

> **First programs in PyQt4 toolkit**

> In this part of the PyQt4 tutorial we will learn some basic functionality.

在这部分我们将会学到一些基本的功能。

## 简单的例子
> **Simple example**

> The code example is very simplistic. It only shows a small window. Yet we can do a lot with this window. We can resize it. Maximize it. Minimize it. This requires a lot of coding. Someone already coded this functionality. Because it repeats in most applications, there is no need to code it over again. So it has been hidden from a programmer. PyQt is a high level toolkit. If we would code in a lower level toolkit, the following code example could easily have dozens of lines.

这个例子十分简单，它仅仅现实一个小窗体。但是我们可以在这个窗体上进行很多操作，我们可以调整大小、最大化、最小化。这些需要很多编码，有人已经把这些功能写好了。因为它在所有的应用程序中重复，因此没有必要再写一次，所以它被隐藏了起来。PyQt是一个高级的工具包。如果我们在一个低级的工具包下写代码，那么接下来的例子将很容易的达到几十行。

    #!/usr/bin/python

    # simple.py

    import sys
    from PyQt4 import QtGui

    app = QtGui.QApplication(sys.argv)

    widget = QtGui.QWidget()
    widget.resize(250, 150)
    widget.setWindowTitle('simple')
    widget.show()

    sys.exit(app.exec_())

> The above code shows a small window on the screen.

上面的代码显示在屏幕上显示一个小窗体。

    import sys
    from PyQt4 import QtGui
  
> Here we provide the necessary imports. The basic GUI widgets are located in **QtGui** module.

我们在这里进行一些必要的import操作。基本的GUI组件在 `QtGui` 模块中。


    app = QtGui.QApplication(sys.argv)
    
> Every PyQt4 application must create an application object. The application object is located in the QtGui module. The sys.argv parameter is a list of arguments from the command line. Python scripts can be run from the shell. It is a way, how we can control the startup of our scripts.

每个PyQt4程序必须创建一个application对象，application在 `QtGui` 模块中， `sys.argv` 参数是命令行中的一组参数。Python脚本可以在shell中运行，这样，我们可以控制脚本的启动。

    widget = QtGui.QWidget()
    
> The QWidget widget is the base class of all user interface objects in PyQt4. We provide the default constructor for QWidget. The default constructor has no parent. A widget with no parent is called a window.

`QWidget` 窗口部件是PyQt4中所有用户界面对象的基类，我们使用 `QWidget` 默认的构造，没有父亲。没有父亲的窗口部件称为窗体。

    widget.resize(250, 150)
    
> The resize() method resizes the widget. It is 250px wide and 150px high.

`resize()` 方法调整了 `widget` 的大小，宽250像素，高150像素。

    widget.setWindowTitle('simple')

> Here we set the title for our window. The title is shown in the titlebar.

这里我们为窗口设置了标题，标题显示在标题栏上。

    widget.show()
    
> The show() method displays the widget on the screen.

`show()` 方法将窗口呈现在屏幕上。

    sys.exit(app.exec_())
    
> Finally, we enter the mainloop of the application. The event handling starts from this point. The mainloop receives events from the window system and dispatches them to the application widgets. The mainloop ends, if we call the exit() method or the main widget is destroyed. The sys.exit() method ensures a clean exit. The environment will be informed, how the application ended.

最后，我们输入应用程序的主事件循环，事件处理从这里开始。主事件循环从窗口系统接收事件并分发到应用程序的窗口部件上。当主事件循环结束，如果我们调用 `exit()` 方法或者主窗口部件被销毁。 `sys.exit()` 方法确保干净的退出。将通知环境应用程序是如何结束的。

> You wonder why the exec_() method has an underscore? Everything has a meaning. This is obviously because the exec is a python keyword. And thus, exec_() was used instead.

你是否想知道为什么 `exec_()` 方法会有下划线，一切皆有含义，这显然是因为exec是Python的关键字，因此，用 `exec_()` 来取代它。

![simple.png][simple]

图: 简单例子

## 应用程序图标
> **An application icon**

> The application icon is a small image, which is usually displayed in the top left corner of the titlebar. In the following example we will show, how we do it in PyQt4. We will also introduce some new methods.

应用程序图标是一个小图像，一般显示在标题栏的左上角，在接下来的例子中，我们将会展示如何在PyQt中实现。我们也会介绍一些新方法。

> note:: 在Mac OS X和Ubuntu中，图标不会显示在标题栏上，因此下面使用的是Windows XP下的截图。

    #!/usr/bin/python

    # icon.py

    import sys
    from PyQt4 import QtGui


    class Icon(QtGui.QWidget):
        def __init__(self, parent=None):
            QtGui.QWidget.__init__(self, parent)

            self.setGeometry(300, 300, 250, 150)
            self.setWindowTitle('Icon')
            self.setWindowIcon(QtGui.QIcon('icons/web.png'))


    app = QtGui.QApplication(sys.argv)
    icon = Icon()
    icon.show()
    sys.exit(app.exec_())
    
> The previous example was coded in a procedural style. Python programming language supports both procedural and object oriented programming styles. Programming in PyQt4 means programming in OOP.

上一个例子用面向过程的方式编写。Python语言支持面向过程和面向对象编程。在PyQt4中编程意味着OOP编程。

    class Icon(QtGui.QWidget):
         def __init__(self, parent=None):
             QtGui.QWidget.__init__(self, parent)
    
> The three most important things in object oriented programming are classes, data and methods. Here we create a new class called Icon. The Icon class inherits from QtGui.QWidget class. This means, that we must call two constructors. The first one for the Icon class and the second one for the inherited class.

在面向对象编程中三个最重要东西是类、数据和方法。这里我们创建一个名为 `Icon` 的新类，继承自 `QtGui.QWidget` 类，意味着我们必须调用两个构造方法，第一个为 `Icon` 类，第二个为继承的父类。

    self.setGeometry(300, 300, 250, 150)
    self.setWindowTitle('Icon')
    self.setWindowIcon(QtGui.QIcon('icons/web.png'))
    
> All three classes have been inherited from the QtGui.QWidget class. The setGeometry() does two things. It locates the window on the screen and sets the size of the window. The first two parameters are the x and y positions of the window. The third is the width and the fourth is the height of the window. The last method sets the application icon. To do this, we have created a QIcon object. The QIcon receives the path to our icon to be displayed.

所有这三个类，均继承自 `QtGui.QWidget` ， `setGeometry()` 方法干了两件事，它定义了窗体在屏幕上的位置，并设置窗体的大小。开始的两个参数是窗体的x和y的位置，第三个是宽度，第四个是高度。最后一个方法设置了应用程序图标。通过这个，我们创建了一个 `QIcon` 对象， `QIcon` 接受我们想要显示的图标的路径。

![icon.png][icons]

图: 图标（Windows）

## 显示工具提示
> **Showing a tooltip**

> We can provide a balloon help for any of our widgets.

我们可以为任何组件提供气泡帮助。

    #!/usr/bin/python

    # tooltip.py

    import sys
    from PyQt4 import QtGui
    from PyQt4 import QtCore


    class Tooltip(QtGui.QWidget):
        def __init__(self, parent=None):
            QtGui.QWidget.__init__(self, parent)

            self.setGeometry(300, 300, 250, 150)
            self.setWindowTitle('Tooltip')

            self.setToolTip('This is a <b>QWidget</b> widget')
            QtGui.QToolTip.setFont(QtGui.QFont('OldEnglish', 10))


    app = QtGui.QApplication(sys.argv)
    tooltip = Tooltip()
    tooltip.show()
    sys.exit(app.exec_())
    
> In this example, we show a tooltip for a QWidget widget.

在这个例子中，我们为一个 `QWidget` 窗口部件显示提示。

    self.setToolTip('This is a <b>QWidget</b> widget')
    
> To create a tooltip, we call the setTooltip() method. We can use rich text formatting.

通过调用 `setTooltip()` 方法来创建提示，我们使用富文本格式。

    QtGui.QToolTip.setFont(QtGui.QFont('OldEnglish', 10))
    
> Because the default QToolTip font looks bad, we change it.

因为默认的 `QToolTip` 字体太难看，我们改变了它。

note:: 
:    在Mac系统上，修改字体并没有变化。

![tooltip.png][tooltip]

图: 工具提示
    

## 关闭窗体
> **Closing a window**

> The obvious way to how to close a window is to click on the x mark on the titlebar. In the next example, we will show, how we can programatically close our window. We will briefly touch signals and slots.

关闭一个窗体很明显的方式是点击标题栏上的x标记，在接下来的例子中，我们将展示如何用编程的方式关闭窗体，我们将简略的接触到信号和槽。

> The following is the constructor of a QPushButton, that we will use in our example.

接下来是构造一个 `QPushButton` ，我们将用在例子中。


    QPushButton(string text, QWidget parent = None)
    
> The text parameter is a text that will be displayed on the button. The parent is the ancestor, onto which we place our button. In our case it is QWidget.

参数 `text` 是显示在按钮上的文本， `parent` 是放置按钮的父亲，在这里是 `QWidget` 。

    #!/usr/bin/python

    # quitbutton.py

    import sys
    from PyQt4 import QtGui, QtCore

    class QuitButton(QtGui.QWidget):
        def __init__(self, parent=None):
            QtGui.QWidget.__init__(self, parent)

            self.setGeometry(300, 300, 250, 150)
            self.setWindowTitle('Quit button')

            quit = QtGui.QPushButton('Close', self)
            quit.setGeometry(10, 10, 64, 35)

            self.connect(quit, QtCore.SIGNAL('clicked()'), 
                QtGui.qApp, QtCore.SLOT('quit()'))


    app = QtGui.QApplication(sys.argv)
    qb = QuitButton()
    qb.show()
    sys.exit(app.exec_())
    
::

    quit = QtGui.QPushButton('Close', self)
    quit.setGeometry(10, 10, 60, 35)

> We create a push button and position it on the QWidget just like we have positioned the QWidget on the screen.

创建一个按钮并放置在 `QWidget` 上，就像把 `QWidget` 放置在屏幕上一样。

::

    self.connect(quit, QtCore.SIGNAL('clicked()'),
        QtGui.qApp, QtCore.SLOT('quit()'))

> The event processing system in PyQt4 is built with the signal & slot mechanism. If we click on the button, the signal clicked() is emitted. The slot can be a PyQt slot or any python callable. The QtCore.QObject.connect() method connects signals with slots. In our case the slot is a predefined PyQt quit() slot. The communication is done between two objects. The sender and the receiver. The sender is the push button, the receiver is the application object.

PyQt4中的事件处理系统是建立信号和槽机制。如果点击按钮，信号 `clicked()` 将会发射，槽可以是PyQt的槽或者任何的Python调用。 `QtCore.QObject.connect()` 把信号和槽连接起来。在这个例子中槽是PyQt预定义的 `quit()` 槽，发射方和接收方两个对象间进行通讯，发射方为按钮，接收方为application对象。

![quit_button.png][quit_button]

图: 退出按钮
    
    
## 消息框
> **Message Box**
    
> By default, if we click on the x button on the titlebar, the QWidget is closed. Sometimes we want to modify this default behaviour. For example, if we have a file opened in an editor to which we did some changes. We show a message box to confirm the action.

默认情况下，当我们点击标题栏上的关闭按钮， `QWdiget` 将关闭。有时候我们需要改变默认行为，比如：当我们在编辑器中打开一个文件并且做了一些修改，显示一个消息框来确认这个动作。


    #!/usr/bin/python

    # messagebox.py

    import sys
    from PyQt4 import QtGui

    class MessageBox(QtGui.QWidget):
        def __init__(self, parent=None):
            QtGui.QWidget.__init__(self, parent)

            self.setGeometry(300, 300, 250, 150)
            self.setWindowTitle('message box')


        def closeEvent(self, event):
            reply = QtGui.QMessageBox.question(self, 'Message',
                "Are you sure to quit?", QtGui.QMessageBox.Yes, QtGui.QMessageBox.No)

            if reply == QtGui.QMessageBox.Yes:
                event.accept()
            else:
                event.ignore()

    app = QtGui.QApplication(sys.argv)
    qb = MessageBox()
    qb.show()
    sys.exit(app.exec_())

> If we close the QWidget, the QCloseEvent is generated. To modify the widget behaviour we need to reimplement the closeEvent() event handler.

如果我们关闭 `QWidget` ，将会产生一个 `QCloseEvent` 事件，我们需要重新实现 `closeEvent()` 事件来改变组件的行为。

    reply = QtGui.QMessageBox.question(self, 'Message',
        "Are you sure to quit?", QtGui.QMessageBox.Yes, QtGui.QMessageBox.No)

> We show a message box with two buttons. Yes and No. The first string appears on the titlebar. The second string is the message text displayed by the dialog. The return value is stored in the reply variable.

我们显示一个两个按钮的消息框，Yes和No。第一个字符串显示在标题栏上，第二个字符串是信息文本，显示在对话框中，返回值存储在 `relay` 变量中。

note:
:    Mac OS X上不显示标题


    if reply == QtGui.QMessageBox.Yes:
        event.accept()
    else:
        event.ignore()
        
> Here we test the return value. If we clicked Yes button, we accept the event which leads to the closure of the widget and to the termination of the application. Otherwise we ignore the close event.

这里我们测试返回值，如果点击Yes按钮，表明接受这个事件，将会关闭窗口部件并进回到应用程序的结尾，反之则忽略该关闭事件。

![message_box.png][message_box]

图：消息框
    
## 屏幕居中显示窗体
> **Centering window on the screen**

> The following script shows, how we can center a window on the desktop screen.

下面的脚本我们将展示如何在屏幕中央显示窗体。

    #!/usr/bin/python

    # center.py

    import sys
    from PyQt4 import QtGui

    class Center(QtGui.QWidget):
        def __init__(self, parent=None):
            QtGui.QWidget.__init__(self, parent)

            self.setWindowTitle('center')
            self.resize(250, 150)
            self.center()

        def center(self):
            screen = QtGui.QDesktopWidget().screenGeometry()
            size =  self.geometry()
            self.move((screen.width()-size.width())/2, (screen.height()-size.height())/2)


    app = QtGui.QApplication(sys.argv)
    qb = Center()
    qb.show()
    sys.exit(app.exec_())

::

    self.resize(250, 150)
    
> Here we resize the QWidget to be 250px wide and 150px heigh.

调整这个 `QWidget` 的大小为250像素宽和150像素高。

    screen = QtGui.QDesktopWidget().screenGeometry()

> We figure out the screen resolution of our monitor.

计算屏幕分辨率。

    size = self.geometry()
    
> Here we get the size of our QWidget.

获得 `QWidget` 的尺寸。

    self.move((screen.width()-size.width())/2, (screen.height()-size.height())/2)

> Here we move the window to the center of the screen.

把窗体移动到屏幕中央。

> In this part of the PyQt4 tutorial, we covered some basics.

在PyQt4教程的这一部分，我们涉及一些基本知识。

[simple]:http://jimmykuu.sinaapp.com/static/PyQt4_Tutorial/html/_images/simple.png
[icon]:http://jimmykuu.sinaapp.com/static/PyQt4_Tutorial/html/_images/icon.png
[tooltip]:http://jimmykuu.sinaapp.com/static/PyQt4_Tutorial/html/_images/tooltip.png
[quit_button]:http://jimmykuu.sinaapp.com/static/PyQt4_Tutorial/html/_images/quit_button.png
[message_box]:http://jimmykuu.sinaapp.com/static/PyQt4_Tutorial/html/_images/message_box.png
