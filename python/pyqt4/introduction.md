# PyQt4工具包简介

> **Introduction to PyQt4 toolkit**


## 关于本教程
> **About this tutorial**

> This is an introductory PyQt4 tutorial. The purpose of this tutorial is to get you started with the PyQt4 toolkit. The tutorial has been created and tested on Linux.

本教程是用来介绍PyQt，目的是使您开始使用PyQt工具包，该教程在Linux下测试通过。

>译者根据最新版本的PyQt4做了一些修改，并在Mac OS X下测试通过。


## 关于PyQt
> **About PyQt4**

> PyQt4 is a toolkit for creating GUI applications. It is a blending of Python programming language and the successfull Qt library. Qt library is one of the most powerful libraries on this planet. The official home site for PyQt4 is on www.riverbankcomputing.co.uk It was developed by Phil Thompson.

PyQt是用来创建GUI应用程序的工具包，它把Python和成功的Qt绑定在一起，Qt库是这个星球上最强大的库之一，如果不是最强大的话。PyQt的官方网站是 [www.riverbankcomputing.co.uk](<http://www.riverbankcomputing.co.uk>) ， 它由 **Phil Thompson** 开发。

> PyQt4 is implemented as a set of Python modules. It has over 300 classes and almost 6000 functions and methods. It is a multiplatform toolkit. It runs on all major operating systems. Including Unix, Windows and Mac. PyQt4 is dual licenced. Developers can choose between GPL and commercial licence. Previously, GPL version was available only on Unix. Starting from PyQt version 4, GPL licence is available on all supported platforms.

PyQt作为一组Python模块的实现。有超过300个类和超过6000个方法。它是个跨平台的工具包，它可以运行在所有的主流操作系统上。包括Unix、Windows、Mac。它具有双重授权，开发者可以选择GPL或者商业授权。以前GPL版本只能在Unix上获得。从PyQt4开始，可以在所有支持的平台上获得GPL协议。

> Because there are a lot of classes available, they have been divided into several modules.

由于有大量可用的类，它们被分成多个模块。

![modules][modules]

图：PyQt4模块
   
> The QtCore module contains the core non-gui functionality. This module is used for working with time, files and directories, various data types, streams, urls, mime types, threads or processes. The QtGui module contains the graphical components and related classes. These include for example buttons, windows, status bars, toolbars, sliders, bitmaps, colors, fonts etc. The QtNetwork module contains the classes for network programming. These classes allow to write TCP/IP and UDP clients and servers. They make the network programming easier and more portable. The QtXml contains classes for working with xml files. This module provides implementation for both SAX and DOM APIs. The QtSvg module provides classes for displaying the contents of SVG files. Scalable Vector Graphics (SVG) is a language for describing two-dimensional graphics and graphical applications in XML. The QtOpenGL module is used for rendering 3D and 2D graphics using the OpenGL library. The module enables seamless integration of the Qt GUI libary and the OpenGL library. The QtSql module provides classes for working with databases.
   
* **QtCore** 模块包括了核心的非GUI功能，该模块用来对时间、文件、目录、各种数据类型、流、网址、媒体类型、线程或进程进行处理。
* **QtGui** 模块包括图形化窗口部件和及相关类。包括如按钮、窗体、状态栏、滑块、位图、颜色、字体等等。
* **QtHelp** 模块包含了用于创建和查看可查找的文档的类。
* **QtNetwork** 模块包括网络编程的类。这些类可以用来编写TCP/IP和UDP的客户端和服务器。它们使得网络编程更容易和便捷。
* **QtOpenGL** 模块使用OpenGL库来渲染3D和2D图形。该模块使得Qt GUI库和OpenGL库无缝集成。
* **QtScript** 模块包含了使PyQt应用程序使用JavaScript解释器编写脚本的类。
* **QtSql** 模块提供操作数据库的类。
* **QtSvg** 模块提供了显示SVG文件内容的类。可缩放矢量图形(SVG)是一种用XML描述二维图形和图形应用的语言。
* **QtTest** 模块包含了对PyQt应用程序进行单元测试的功能。（PyQt没有实现完全的Qt单元测试框架，相反，它假设使用标准的Python单元测试框架来实现模拟用户和GUI进行交互。）
* **QtWebKit** 模块实现了基于开源浏览器引擎WebKit的浏览器引擎。
* **QtXml** 包括处理XML文件的类，该模块提供了SAX和DOM API的接口。
* **QtXmlPatterns** 模块包含的类实现了对XML和自定义数据模型的XQuery和XPath的支持。
* **phonon** 模块包含的类实现了跨平台的多媒体框架，可以在PyQt应用程序中使用音频和视频内容。
* **QtMultimedia** 模块提供了低级的多媒体功能，开发人员通常使用 **phonon** 模块。
* **QtAssistant** 模块包含的类允许集成 **Qt Assistant** 到PyQt应用程序中，提供在线帮助。
* **QtDesigner** 模块包含的类允许使用PyQt扩展 **Qt Designer** 。
* **Qt** 模块综合了上面描述的模块中的类到一个单一的模块中。这样做的好处是你不用担心哪个模块包含哪个特定的类，坏处是加载进了整个Qt框架，从而增加了应用程序的内存占用。
* **uic** 模块包含的类用来处理.ui文件，该文件由Qt Designer创建，用于描述整个或者部分用户界面。它包含的加载.ui文件和直接渲染以及从.ui文件生成Python代码为以后执行的类。

## Python

> Python is a dynamic object-oriented programming language. It is a general purpose programming language. It can be used for many kinds of software development. The design purpose of the Python language emphasizes programmer productivity and code readability. Python was initially developed by Guido van Rossum. It was first released in 1991. Python was inspired by ABC, Haskell, Java, Lisp, Icon and Perl programming languages. Python is a high level, general purpose, multiplatform, interpreted language. Python is a minimalistic language. One of it's most visible features is that it does not use semicolons nor brackets. Python uses indentation instead. There are two main branches of Python currently. Python 2.x and Python 3.x. Python 3.x breaks backward compatibility with previous releases of Python. It was created to correct some design flaws of the language and make the language more clean. The most recent version of Python 2.x is 2.7.1, and of Python 3.x 3.1.3. This tutorial covers Python 2.x versions. Most of the code is written in Python 2.x versions. It will take some time till the software base and programmers will migrate to Python 3.x. Today, Python is maintained by a large group of volunteers worldwide. Python is open source software.

![python-logo][python-logo]

Python是个成功的脚本语言。它最初由 **Guido van Rossum** 开发，在1991年第一次发布。Python由ABC和Haskell语言所启发。Python是一个高级的、通用的、跨平台、解释型的语言。一些人更倾向于称之为动态语言。它很易学，Python是一种简约的语言。它的最明显的一个特征是，不使用分号或括号，Python使用缩进。最近的版本是2.7(3.2)，2011年二月发布。现在，Python由来自世界各地的庞大的志愿者维护。

> Python is an ideal start for those, who want to learn programming.

Python是那些想要学习编程的人的理想的入门语言。

> Python programming language supports several programming styles. It does not force a programmer to a specific paradigm. Python supports object oriented and procedural programming. There is also a limited support for functional programming.

Python支持多种编程模式，它不强制程序员使用特定的模式。Python支持面向对象和面向过程编程，也有限的支持函数式编程。

> The official web site for the Python programming language is python.org

Python的官方网站是： [python.org](<http://python.org>) 。

> There are currently several widely used programming languages. The following list is based on the TIOBE Programming Community Index. The numbers are from November 2010.

这里多种广泛使用的编程语言，下面的列表是根据 `TIOBE <http://www.tiobe.com/index.php/content/paperinfo/tpci/index.html>`_ 编程语言排行榜_2011年3月的排行。

 ![tiobe][tiobe]

> Java is the most widely used programming language. Java excels in creating portable mobile applications, programming various appliances and in creating enterprise applications. Every fourth application is programmed in C/C++. They are standard for creating operating systems and various desktop applications. C/C++ are the most widely used system programming languages. Most famous desktop applications were created in C++. May it be MS Office, Macromedia Flash, Adobe Photoshop or 3D Max. These two languages also dominate the game programming business.

> PHP dominates over the Web. While Java is used mainly by large organizations, PHP is used by smaller companies and individuals. PHP is used to create dynamic web applications.

> C# is the main programming language of the Microsoft .NET platform. C# is followed in .NET by Visual Basic. It represents of the popularity of the RAD. (Rapid Application Development.)

> Perl, Python and Ruby are the most widely used scripting languages. They share many similarities. They are close competitors.

> The Objective C is the main programming language of the Apple ecosystem.


## Python 图形工具包
> **Python toolkits**

Python程序员有三种不错的选择来创建图形界面：PyGTK、wxPython、PyQt。根据实际情况来选择工具包。

note::
:    现在Nokia推出了PySide，Qt的Python绑定又多了一个选择。

[modules]:http://jimmykuu.sinaapp.com/static/PyQt4_Tutorial/html/_images/modules.png
[python-logo]:http://python.org/images/python-logo.gif
[tiobe]:http://jimmykuu.sinaapp.com/static/PyQt4_Tutorial/html/_images/tiobe.png
