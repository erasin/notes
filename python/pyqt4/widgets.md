#窗口部件
> **PyQt4 Widgets**

> Widgets are basic building blocks of an application. The PyQt4 programming toolkit has a wide range of various widgets. Buttons, check boxes, sliders, list boxes etc. Everything a programmer needs for his job. In this section of the tutorial, we will describe several useful widgets.

窗口部件是应用程序的基本构建块。PyQt4编程工具包拥有范围广泛的各种窗口部件。按钮、选择框、滑块、列表框等等，程序员工作所需要的一切。在教程的这部分中，我们将介绍一些有用的窗口部件。

## QCheckBox {#qcheckbox}

> QCheckBox is a widget that has two states. On and Off. It is a box with a label. Whenever a checkbox is checked or cleared it emits the signal stateChanged().

`QCheckBox` （复选框） 是一个由两种状态的窗口部件。 `On` 和 `Off` 。他是一个带标签的框。每段一个复选框被选中和或者清楚时，都将发射信号 `stateChanged()` 。


    #!/usr/bin/python
    # -*- coding: utf-8 -*-

    # checkbox.py

    import sys
    from PyQt4 import QtGui
    from PyQt4 import QtCore


    class Example(QtGui.QWidget):
  
        def __init__(self):
            super(Example, self).__init__()
        
            self.initUI()
        
        def initUI(self):

            self.setGeometry(300, 300, 250, 150)
            self.setWindowTitle('Checkbox')

            self.cb = QtGui.QCheckBox('Show title', self)
            self.cb.setFocusPolicy(QtCore.Qt.NoFocus)
            self.cb.move(10, 10)
            self.cb.toggle()
            self.connect(self.cb, QtCore.SIGNAL('stateChanged(int)'), 
                self.changeTitle)

        def changeTitle(self, value):
      
            if self.cb.isChecked():
                self.setWindowTitle('Checkbox')
            else:
                self.setWindowTitle('')


    if __name__ == '__main__':
  
        app = QtGui.QApplication(sys.argv)
        ex = Example()
        ex.show()
        app.exec_()

> In our example, we will create a checkbox that will toggle the window title.

在我们的例子中，我们将创建一个复选框来切换窗口标题。

::

    self.cb = QtGui.QCheckBox('Show title', self)
    
> This is the QCheckBox constructor.

构建 `QCheckBox` 。

::

    self.cb.setFocusPolicy(QtCore.Qt.NoFocus)
    
> We disable focus for the QCheckBox. A QCheckBox that has a focus may be selected or unselected with a spacebar.

禁用 `QCheckBox` 的焦点。获由焦点的 `QCheckBox` 可以通过空格选择或者取消选择。

::

    self.cb.toggle()
    
> We set the window title, so we must also check the checkbox. By default, the window title is not set and the checkbox is unchecked.

设置窗口标题，必须选中复选框。默认情况下，不设置窗口标题，复选框设为未选择。

::

    self.connect(self.cb, QtCore.SIGNAL('stateChanged(int)'), self.changeTitle)

> We connect the user defined changeTitle() method to the stateChanged() signal. The changeTitle() method will toggle the window title.

连接用户定义的 `changeTitle()` 到 `stateChanged()` 信号。 `changeTitle()` 方法将切换窗口标题。

![QCheckBox][qcheckbox]

图：QCheckBox

## 切换按钮
> **ToggleButton**

> PyQt4 has no widget for a ToggleButton. To create a ToggleButton, we use a QPushButton in a special mode. ToggleButton is a button that has two states. Pressed and not pressed. You toggle between these two states by clicking on it. There are situations where this functionality fits well.

PyQt4没有切换按钮的窗口部件，为了创建切换按钮，我们使用特殊模式的 `QPushButton` 。切换按钮是指一个两种状态的按钮，按下和非按下。通过点击切换两种状态。在某种状态下来这种方式很合适。

::

    #!/usr/bin/python
    # -*- coding: utf-8 -*-

    # togglebutton.py

    import sys
    from PyQt4 import QtGui
    from PyQt4 import QtCore


    class Example(QtGui.QWidget):
  
        def __init__(self):
            super(Example, self).__init__()
        
            self.initUI()
        
        def initUI(self):

            self.color = QtGui.QColor(0, 0, 0)       

            self.red = QtGui.QPushButton('Red', self)
            self.red.setCheckable(True)
            self.red.move(10, 10)

            self.connect(self.red, QtCore.SIGNAL('clicked()'), self.setColor)

            self.green = QtGui.QPushButton('Green', self)
            self.green.setCheckable(True)
            self.green.move(10, 60)

            self.connect(self.green, QtCore.SIGNAL('clicked()'), self.setColor)

            self.blue = QtGui.QPushButton('Blue', self)
            self.blue.setCheckable(True)
            self.blue.move(10, 110)

            self.connect(self.blue, QtCore.SIGNAL('clicked()'), self.setColor)

            self.square = QtGui.QWidget(self)
            self.square.setGeometry(150, 20, 100, 100)
            self.square.setStyleSheet("QWidget { background-color: %s }" % 
                self.color.name())
            
            self.setWindowTitle('ToggleButton')
            self.setGeometry(300, 300, 280, 170)            


        def setColor(self):
      
            source = self.sender()
        
            if source.text() == "Red":
                if self.red.isChecked():
                    self.color.setRed(255)
                else: self.color.setRed(0)    
            
            elif source.text() == "Green":
                if self.green.isChecked():
                    self.color.setGreen(255)
                else: self.color.setGreen(0)
            
            else:
                if self.blue.isChecked():
                    self.color.setBlue(255)
                else: self.color.setBlue(0)
            
            self.square.setStyleSheet("QWidget { background-color: %s }" %
                self.color.name())            



    if __name__ == '__main__':
  
        app = QtGui.QApplication(sys.argv)
        ex = Example()
        ex.show()
        app.exec_()
        
> In our example, we create three ToggleButtons. We also create a QWidget. We set the background color of the QWidget to black. The togglebuttons will toggle the red, green and blue parts of the color value. The background color will depend on which togglebuttons we have pressed.

这个例子中，我们创建了三个切换按钮和一个窗口部件，设置窗口部件的背景为黑色。这些切换按钮将切换颜色值的红绿蓝部分。背景色依赖于我们按下哪个切换按钮。

::

    self.color = QtGui.QColor(0, 0, 0)

> This is the initial color value. No red, green and blue equals to black.

这是初始颜色，没有红绿蓝等于黑色。

::

    self.red = QtGui.QPushButton('Red', self)
    self.red.setCheckable(True)

> To create a ToggleButton, we create a QPushButton and make it checkable by calling setCheckable() method.

为了创建一个切换按钮，我们创建了一个 `QPushButton` 并通过 `setCheckable` 方法使之可选择。

::

    self.connect(self.red, QtCore.SIGNAL('clicked()'), self.setColor)

> We connect a clicked() signal to our user defined method.

连接 `clicked` 信号到自定义的方法。

::

    source = self.sender()
    
> We get the button, which was toggled.

获得切换的按钮。

::

    if source.text() == "Red":
        if self.red.isChecked():
            self.color.setRed(255)
        else: self.color.setRed(0) 

> In case it was a red button, we update the red part of the color accordingly.

如果是红色按钮，我们更新颜色的红色部分。

::

    self.square.setStyleSheet("QWidget { background-color: %s }" %
        self.color.name())

> To change the background color, we use stylesheets.

通过样式表修改背景色。

![ToggleButton][togglebutton]

图：切换按钮

## QSlider

> QSlider is a widget that has a simple handle. This handle can be pulled back and forth. This way we are choosing a value for a specific task. Sometimes using a slider is more natural, than simply providing a number or using a spin box. QLabel displays text or image.

滑块是由一个简单的滑柄的窗口部件。该滑柄可以前后拖动，通过这种方式我们可以为特定任务选择值。有时候使用滑块比简单提供数值或使用微调框(spin box)更自然。 `QLabel` 显示文字或图像。

> In our example we will show one slider and one label. This time, the label will display an image. The slider will control the label.

该例子中我们将显示一个滑块和一个标签。这次，标签将显示一个图像，滑块用来控制标签。

::

    #!/usr/bin/python
    # -*- coding: utf-8 -*-

    # slider.py

    import sys
    from PyQt4 import QtGui
    from PyQt4 import QtCore


    class Example(QtGui.QWidget):
  
        def __init__(self):
            super(Example, self).__init__()
        
            self.initUI()
        
        def initUI(self):

            slider = QtGui.QSlider(QtCore.Qt.Horizontal, self)
            slider.setFocusPolicy(QtCore.Qt.NoFocus)
            slider.setGeometry(30, 40, 100, 30)
            self.connect(slider, QtCore.SIGNAL('valueChanged(int)'), 
                self.changeValue)
        
            self.label = QtGui.QLabel(self)
            self.label.setPixmap(QtGui.QPixmap('mute.png'))
            self.label.setGeometry(160, 40, 80, 30)
        
            self.setWindowTitle('Slider')
            self.setGeometry(300, 300, 250, 150)
        
    
        def changeValue(self, value):

            if value == 0:
                self.label.setPixmap(QtGui.QPixmap('mute.png'))
            elif value > 0 and value <= 30:
                self.label.setPixmap(QtGui.QPixmap('min.png'))
            elif value > 30 and value < 80:
                self.label.setPixmap(QtGui.QPixmap('med.png'))
            else:
                self.label.setPixmap(QtGui.QPixmap('max.png'))


    if __name__ == '__main__':
  
        app = QtGui.QApplication(sys.argv)
        ex = Example()
        ex.show()
        app.exec_()
    
> In our example we simulate a volume control. By dragging the handle of a slider, we change a image on the label.

这个例子中我们模拟音量控制，通过拖动滑块的滑柄来改变标签上的图片。

::

    slider = QtGui.QSlider(QtCore.Qt.Horizontal, self)

> Here we create a horizontal QSlider.

创建一个水平 `QSlider` 。

::

    self.label = QtGui.QLabel(self)
    self.label.setPixmap(QtGui.QPixmap('mute.png'))

> We create a QLabel widget. And set an initial mute image to it.

创建一个 `QLabel` 窗口部件。设置一个促使的无声图形在上面。

::

    self.connect(slider, QtCore.SIGNAL('valueChanged(int)'), 
        self.changeValue)

> We connect the valueChanged signal to the user defined changeValue() method.

连接 `valueChanged` 信号到自定义的 `changeValue()` 方法。

::

    if value == 0:
        self.label.setPixmap(QtGui.QPixmap('mute.png'))
    >.

> Based on the value of the slider, we set an image to the label. In the above code, we set a mute.png image to the label, if the slider is equal to zero.

基于滑块的值，我们设置图形到标签上。在上面的代码，如果滑块的值为0，设置 `mute.png` 图片到标签上。

![QSlider widget][qslider]

图：QSlider 窗口部件

## QProgressBar

> A progress bar is a widget that is used, when we process lengthy tasks. It is animated so that the user knows, that our task is progressing. The QProgressBar widget provides a horizontal or vertical progress bar in PyQt4 toolkit. The task is divided into steps. The programmer can set the minimum and maximum values for the progress bar. The default values are 0, 99.

进度条使用来处理长时间任务的窗口部件，当看到它的动画时，用户就知道我们的任务正在进行中。在PyQt4工具包中， `QProgressBar` 窗口部件提供水平或者垂直的进度条。任务被分成一些阶段。程序员可以为进度条设置最小值和最大值。默认是0，99.

::

    #!/usr/bin/python
    # -*- coding: utf-8 -*-

    # progressbar.py

    import sys
    from PyQt4 import QtGui
    from PyQt4 import QtCore


    class Example(QtGui.QWidget):
  
        def __init__(self):
            super(Example, self).__init__()
        
            self.initUI()
        
        def initUI(self):

            self.pbar = QtGui.QProgressBar(self)
            self.pbar.setGeometry(30, 40, 200, 25)

            self.button = QtGui.QPushButton('Start', self)
            self.button.setFocusPolicy(QtCore.Qt.NoFocus)
            self.button.move(40, 80)

            self.connect(self.button, QtCore.SIGNAL('clicked()'), 
                self.doAction)

            self.timer = QtCore.QBasicTimer()
            self.step = 0
        
            self.setWindowTitle('ProgressBar')
            self.setGeometry(300, 300, 250, 150)
        

        def timerEvent(self, event):
      
            if self.step >= 100:
                self.timer.stop()
                return
            
            self.step = self.step + 1
            self.pbar.setValue(self.step)

        def doAction(self):
      
            if self.timer.isActive():
                self.timer.stop()
                self.button.setText('Start')
            else:
                self.timer.start(100, self)
                self.button.setText('Stop')


    if __name__ == '__main__':

        app = QtGui.QApplication(sys.argv)
        ex = Example()
        ex.show()
        app.exec_()
    
> In our example we have a horizontal progress bar and a push button. The push button starts and stops the progress bar.

这个例子中有一个水平进度表和一个按钮，按钮开始或结束进度条。

::

    self.pbar = QtGui.QProgressBar(self)

> This is a QProgressBar constructor.

构建 `QProgressBar` 。

::

    self.timer = QtCore.QBasicTimer()

> To activate the progress bar, we use the timer object.

我们使用定时器对象激活进度条。

::

    self.timer.start(100, self)
    
> To launch the timer events, we call the start() method. This method has two parameters. The timeout and the object, which will receive the events.

通过调用 `start()` 方法加载定时器事件，该方法有两个参数，超时时间（ `timeout` ）和接受事件的对象（ `object` ）。

::

    def timerEvent(self, event):

        if self.step >= 100:
            self.timer.stop()
            return
        self.step = self.step + 1
        self.pbar.setValue(self.step)

> Each QObject and its descendants has a QObject.timerEvent event handler. In order to react to timer events, we reimplement the event handler.

每个 `QObject` 和它的继承者都有个 `QObject.timerEvent` 事件处理程序。为了应对定时器事件，我们重新实现了该事件处理程序。

![ProgressBar][progressbar]

图：进度条

## QCalendarWidget

> The QCalendarWidget provides a monthly based calendar widget. It allows a user to select a date in a simple and intuitive way.

`QCalendarWidget` 提供基于月份的日历窗口部件，它允许用户简单并且直观的选择日期。

::

    #!/usr/bin/python
    # -*- coding: utf-8 -*-

    # calendar.py

    import sys
    from PyQt4 import QtGui
    from PyQt4 import QtCore


    class Example(QtGui.QWidget):
  
        def __init__(self):
            super(Example, self).__init__()
        
            self.initUI()
        
        def initUI(self):

            self.cal = QtGui.QCalendarWidget(self)
            self.cal.setGridVisible(True)
            self.cal.move(20, 20)
            self.connect(self.cal, QtCore.SIGNAL('selectionChanged()'), 
                self.showDate)

        
            self.label = QtGui.QLabel(self)
            date = self.cal.selectedDate()
            self.label.setText(str(date.toPyDate()))
            self.label.move(130, 260)
        
            self.setWindowTitle('Calendar')  
            self.setGeometry(300, 300, 350, 300)
                
        def showDate(self):
      
            date = self.cal.selectedDate()
            self.label.setText(str(date.toPyDate()))


    if __name__ == '__main__':

        app = QtGui.QApplication(sys.argv)
        ex = Example()
        ex.show()
        app.exec_()

> The example has a calendar widget and a label widget. The currently selected date is displayed in the label widget.

该例子中有一个日历窗口部件和一个标签。当前选择的日期显示在标签中。

::

    self.cal = QtGui.QCalendarWidget(self)
    
> We construct a calendar widget.

构建日历窗口部件。

::

    self.connect(self.cal, QtCore.SIGNAL('selectionChanged()'), 
        self.showDate)

> If we select a date from the widget, a selectionChanged() signal is emitted. We connect this method to the user defined showDate() method.

如果从日历上选择一个日期， `selectionChanged()` 信号将会发射。我们连接该方法到自定义的 `showDate()` 方法上。

::

    def showDate(self):
        date = self.cal.selectedDate()
        self.label.setText(str(date.toPyDate()))

> We retrieve the selected date calling the selectedDate() method. Then we transform the date object into string and set it to the label widget.

通过调用 `selectedDate()` 方法获得日期，然后转换日期对象到字符串并设置到标签上。

![Calendar widget][]

图：日历窗口部件

> In this part of the PyQt4 tutorial, we covered several widgets.

在PyQt4教程的这部分中，我们介绍了一些窗口部件。
