# gunplot 绘图

来自 <http://darksair.org/wiki/Gnuplot.html>

Gnuplot 是一款强大到无敌的跨平台做图工具，可以应付从没事玩玩到写Nature 论文的各种需求，实乃我辈中人之嘲笑 Origin 的利器。Gnuplot 可以绘制宇宙 里存在的各种图表，并可以输出为 EPS, FIG, JPEG, PNG, LaTeX, SVG(!)... 同时，Gnuplot 秉承 UNIX 的优良传统，使用纯文本作为输入，确保它有足够的 表达能力来实现所有功能。最后需要说明的是，Gnuplot 只是名字里恰好有 Gnu， 并不是 GNU 项目的一部分。
Gnuplot 的[官网](http://www.gnuplot.info/)上有完整的[手册](http://www.gnuplot.info/documentation.html)和大量[范例](http://www.gnuplot.info/screenshots/index.html#demos)，值得观摩。

# 使用 Gnuplot

## Gnuplot Interactive Environment

> First Encounter

Gnuplot 有一个交互式命令界面，在 shell 里输入 gnuplot 回车，就会看到 Gnuplot 打印出一些版本信息，然后显示 gnuplot>，这是它的命 令提示符。下面我们来输入第一个 Gnuplot 命令：

	> plot sin(x)

注意第一个 > 字符表示提示符，不要输入。回车后应该会看到 Gnuplot 打开了 一个有工具栏的窗口，里面画着熟悉的正弦曲线。现在试着把这个窗口缩小到一 半大小，怎么样？字看不清了吧？点一下工具栏里的“Apply autoscale”按钮， 让 Gnuplot 重画窗口。

在上面的例子中，我们已经使用了 Gnuplot 的 plot 命令。在 Gnuplot 中，只 存在两个画图的命令：plot 和 splot。其中 plot 用来画二维图形，splot 用来画三维图形。同时我们还见识 了 Gnuplot 的内置数学函数 sin。

再试一个复杂点的：

	> plot sin(x) + sin(1.1*x)

貌似和前面的差不多？那是因为我们画出的范围太小。在 plot 命令中，我们可 以手动指定自变量的范围

	> plot [-50:50] sin(x) + sin(1.1*x)

在 Gnuplot 中，我们使用 [min:max] 这样的语法来定义范围。这次能完整的看 到一拍了。但是有个问题：画出来的图都变成折线了。这是因为Gnuplot 默认取 点太少。对于变化比较平缓的曲线，默认的取点已经足够。如果需要手动指定采 样率的话，可以设置 samples 变量。变量的设置在 Gnuplot 中至关重要，在画 复杂图形的时候，往往会有几十行的变量设置，但画图的命令可能只有最后一行。 我们使用 set 命令设置变量。

	> set samples 1000

Gnuplot 绘图需要使用一系列点，并把它们按顺序用直线连接起来。如果需要绘 制的是一个函数，他就会每隔一定的 x [三维是 (x, y)] 计算出一个函数值，作 为绘图的数据。变量 samples 指定在绘制函数时，我们在绘图区间内取的采样个 数。个数越多曲线越平滑。执行上面的命令后，你会发现曲线并没有变，还是原 来的德性。Gnuplot 不会自动更新曲线（因为如果变量设错了，绘制可能会很 慢），我们需要手动更新。

	> replot

终于，我们的图形平滑了。

在实际使用的时候，绘制函数图形的需求不大，绘制数据才是最重要的。plot 命令可以读取空格（或 tab）分隔的数据文件，比如下面这个：

	# Some data
	1   1
	1.5 2
	2.1 2.5
	3   2.3
	3.2 2.1

注意第一行以 # 开头的是注释，会被 Gnuplot 忽略。把这段文本存为 data.txt，然后在 Gnuplot 里执行

	> reset
	> plot "data.txt"

第一行 reset 命令把我们之前设置的所有变量全部设成默认值。执行后会看到 Gnuplot 画出了 5 个点（第一个在左边的纵坐标轴上，红色的，仔细找）。这 个图很丑，数据之间的关系也不明确。你可能会希望把点和点之间用线连起来。 这可以用两种方式完成：

	> plot "data.txt" with linespoints

选项 with 指定 Gnuplot 对数据的表示方法，“linespoints” 是其中的一种。如果希望以后的所有对数据的 plot 都是用这种方式，可以设定 style 变量：

	> set style data linespoints
	> replot

在 Gnuplot 中有许多 style，比较特殊的是 errorbars 和 candlesticks style，这些 style 需要数据中附加的列来表征多出的维度。

### 获取帮助

Gnuplot 包含大量的命令和变量，一般的地球人不可能把它们的用法全部记住。 当你忘记了一个变量的用法时，可以查阅手册，但一个更快捷的方法是使用 Gnuplot 的在线文档。Gnuplot 的在线文档包含手册的全部内容，并且方便程度 不亚于 Emacs 的 info-mode。假设我们要查 style 命令的用法， 只需执行

	> help style

Gnuplot 就会给出 style 的详细信息。现在假设我们看完后决定试用一种新的 style: impulses。我们对这个 style 一无所知，只是看名字比 较诡异。怎么办？按 q 退出试试，看到了吧？Gnuplot 提示你输入 style 的 subtopic。输入 impulses，就会看到它的介绍了。看完后按 q 退出一层，再按 Ctrl+C 退回提示符。如果一个条目有好几层 subtopics 的话（比如 plot 命 令），Gnuplot 就会一层一层的提示你，知道你狂按 C-c 退到提示符为止。

### 曲线拟合

恐怕大部分人对 Origin 最深刻的体会就是不知道在哪里拟合，即使碰巧打开了 一个拟合的窗口，也会发现里面只能线性拟合，然后郁闷地关上。拟合是 Gnuplot 的强项，凭借先进的 UI，gnuplot 的拟合功能可以在两条命令之内完 成，而且可以拟合为任何函数。假设我们想把上面那个数据文件拟合为抛物线， 只需执行

	> y(x) = a*x**2 + b*x + c
	> fit y(x) "data.txt" via a,b,c

第一条命令定义了一个叫 y 的函数，注意在 Gnuplot 里，求幂的操作符是 `**`。第二条命令就是传说中的 fit 了。其中的 via 是关键字， 后面跟着需要拟合的变量列表。回车后 Gnuplot 会打印出一大票信息，我们只 需要关心结尾的几行，这里把它贴出来：

	Final set of parameters            Asymptotic Standard Error
	=======================            ==========================

	a               = -0.807933        +/- 0.04386      (5.428%)
	b               = 3.87664          +/- 0.1894       (4.887%)
	c               = -2.04687         +/- 0.179        (8.747%)


	correlation matrix of the fit parameters:

				   a      b      c
	a               1.000
	b              -0.992  1.000
	c               0.948 -0.979  1.000

其中 “Final set of parameters” 和 “Asymptotic Standard Error” 是我 们想要的，后面的那个东西...我也不知道是什么 ... -\_-

怎么样？简单吧？我们已经完成了拟合，现在可以把它画出来了。

	> plot "data.txt", y(x)

结果很 PP 吧？ :-) 这个命令虽然简单，也引入了一些新的东西。现在我们知道怎么在一个图里画 100 条曲线了：用 99 个逗号分隔。

### 做出专业的图

上面的图虽然很 PP，但是如果你做个这样的图拿给导师，一定会被暴扁一顿然 后到教室/实验室后面罚站 100 天。因为这个图还缺少图例、坐标名称和标题。 当然，严格地说图例已经有了，只是很搓而已。先把这个图例搞好：

	> set key right bottom
	> plot "data.txt" title "Some Data", y(x) t "~x^2"

一行设置图例的位置在右下角，第二行和前面的 plot 命令的不同 就是给每一个曲线都设置了一个标题，显示在图例中。注意在 Gnuplot 中，很多 关键字都有简写，比如这个 title 选项就可以简写为 t。有人说了，你那个 x^2 太呆了，就不能搞成 x2 么？当然可以了！Gnuplot 怎么可能那么弱 智... 不过这个问题要放到后面讲 terminal 的时候再说。下面设置坐标名称和 标题：

	> set xlabel "t"
	> set ylabel "A"
	> set title "A Dumb Figure"

这个都能看懂，就不说了。赶快 replot 一下，好看多了吧？赶快 拿给导师。“叔叔... 这是我最新画的图，你看漂漂不？”

## Gnuplot Scripting

### Terminal

现在我们已经做出了 PP 图，下一步是什么？有人说了，发给审稿人显摆一下^^。 发个毛啊！你存成 EPS 了么？！所以下一步是把图保存成某种格式。怎么保存呢？ 你是不是开始在工具栏里找“保存”按钮了？如果是的话，说明你已经被 M$ 洗 脑了。在 Gnuplot 中，我们设置 terminal 变量（可以简写为 term）来指定“保存”的格式。这个变量告诉 Gnuplot 你要以什 么方式/格式来展示图形，默认值为 wxt，表示使用一个 WxWidget 窗口来显示图形。要设置格式为 EPS，我们只需把这个变量设为 postscript （可以简写为 post）：

	> set term post eps
	> replot

你是不是看到了一坨乱七八糟的代码？对了。那个就是 Gnuplot 生成的 postscript 代码。term 的选项 eps 指示 Gnuplot 生成一个 Encapsulated Postscript 文件。当然我们会希望把它保存到一个文件里，这个 可以通过设置 output 变量实现：

	> set output "test.eps"
	> replot

用一个 postscript 查看器/打印机看看生成的文件。

### 第一个 Gnuplot Script

现在我们已经设置了很多变量，如果你不幸关闭了 Gnuplot，想在下次启动时实 现现在的效果，就要把这些变量重新设置一遍。能不能把这些东西都存到文件里， 让 Gnuplot 一下一起搞定呢？当然可以，而且我一般都是这样使用 Gnuplot 的。 Gnuplot 文件就是一个纯文本，可以用普通的文本编辑器编写。格式也很简单， 只要把我们在 Gnuplot 提示符后面输入的东西照抄一遍就行了。下面是一个简 单的 Gnuplot script:

	# test.gnu
	set term post eps color
	set output "test.eps"
	set size 0.5, 0.5
	set xlabel "t"
	set ylabel "A"
	set title "A Dumb Figure"
	plot "data.txt"

在这个文件中，有一个我们没见过的变量 size，这个变量的作用 是设置输出图形的大小为设置前的一定比例。比如 Gnuplot 的 Postscript 格式 的默认大小是 10in x 7in，设置 size 为 0.5,0.5 以后的大小变为 5in x 3.5in。变量 terminal 的选项 color 指示 Gnuplot 生成带颜色的图形，和在窗 口中看到的一样；也可以指定为 monochrome，不同的曲线将以不 同的线型加以区分。要执行这个文件，只需在 shell 中输入 gnuplot test.gnu 即可。

### Teminal, Enhanced

现在我们可以开始解决原来的 "x^2" 问题了。Gnuplot 提供了一种 enhanced termianl 机制来输入这些数学符号，开启后输入 x^2 就会输出成 x2。开启 enhance 的方法是在 terminal 变 量的后面加上选项 enhanced。比如

	set term postscript eps color enhanced
	set output "test.eps"
	set title "x^2 Against {/Symbol b}"
	plot "data.txt" t "x^2"

其中 {/Symbol b} 输出为一个 β。Enhanced termianl 还有很多 其他用法，基本可以满足对数学符号的各种需要。使用方法简单，功能强大，非 Origin 之流可比。详细的解释请见 Gnuplot 源码包中附带的 `ps_guide.ps`。

## Gnuplot for Real

在这一节，我们将使用一个实际科研中碰到的例子来讲解 Gnuplot 的使用方法。 考虑这个数据文件，它表征了一个磁性系统的标度行为，在理想情况下应该满足 y = Axβ 的函数关系， 其中 A 为常数。我们的目标是 —— 没有蛀牙！噢，错了，是求出指数 β，并 把它清晰的表示出来。

在一开始，我们还不知道该怎样画这个图，先在 Gnuplot 的交互命令行里试着 plot 一下吧^^ 噼里啪啦~~，感觉还行，只是左边挤得太 扁了，什么都看不出来。如果你对这种幂函数比较了解，应该知道如果使用双对 数坐标画图，幂函数的曲线是一条直线，斜率即为幂指数，貌似符合我们的要求， 那就试试双对数坐标吧~~

	> set logscale xy
	> replot

变量 logscale 的用法就不说了，作为作业吧^^。怎么样，画出来以后有没有惊 艳的感觉？现在我们开始拟合：

	> y1(x) = A1*x**b1
	> y2(x) = A2*x**b2
	> fit [0.01:0.3] y1(x) "data.txt" via A1,b1
	> fit [3:30] y2(x) "data.txt" via A2,b2
	> plot "data.txt", y1(x), y2(x)

看起来还不错，现在我们可以开始写 Gnuplot 脚本了。开头先把变量设了。

	set term postscript eps enhanced color
	set size 0.8,0.8
	set output "figure.eps"
	set logscale xy
	set style data linespoints
	set xrange [0.007:40]
	set yrange [1:50]
	set key left top

其中 xrange 和 yrange 设置绘图的范围。现在拟合，并把曲线画出来。

	y1(x) = A1*x**b1
	y2(x) = A2*x**b2
	fit [0.01:0.3] y1(x) "data.txt" via A1,b1
	fit [3:30] y2(x) "data.txt" via A2,b2
	plot y1(x) t "{/Symbol b} in low-{/Symbol w} regime" lt 2, \
		 y2(x) t "{/Symbol b} in high-{/Symbol w} regime" lt 3, \
		 "data.txt" t "Exponential" lw 2 lt 1

注意在 Gnuplot 中，如果一条命令分成了多于一行，要在除最后一行的行末加一 个 \。lw 参数设置曲线宽度为 2，lt 设置曲线的 linetype，一般来说 lt 1 是醒目的红色，以上的 plot 参数保证 data.txt 的 曲线在两条直线的上方，并且是红色的实线。现在图画出来了，效果不错，但是 光看这个图还是不知道指数 β 到底是多少，最好是直接把它指明。Gnuplot 提 供了 label 变量（！居然是个变量！），用来在指定的位置添加一段文字。

	set label 1 sprintf("{/Symbol b} = %.2f", b1) at 0.09,10 right
	set label 2 sprintf("{/Symbol b} = %.2f", b2) at 4.5,8 right

在这里我们使用了 Gnuplot 的内置函数 sprintf，用法就不说啦，相信有大脑 的同学都明白。注意这段要添加在 plot 命令的前面，否则无效，以下所有的变 量设置都是如此。

现在我们的图已经很像回事了，可以最后添加一些装饰~~

	set xlabel "{/Symbol w}"
	set ylabel "A"
	set title "Exponential relation between A and {/Symbol w}"

大功告成，你可以拿着这个图对用 Origin 的人显摆一下 :-)。

现在，我们对 Gnuplot 提一个难题：把坐标上的数值改成 10n 这种形式，就是说 1 这个刻度显示成 100，10 显示成 101，点点点。这种形式看起来很清晰，很爽快， 很装逼。我们的思路是让 Gnuplot 自己对每一条数据的两个分量都求以 10 为底 的对数，然后把坐标刻度的显示方式变一下。这个好像很变态，其实在Gnuplot 里是一件再简单不过的事：只要设一下 format 变量和 plot 的 using 选项就行 了：

	set format "10^{%.1f}"
	plot "data.txt" using (log10($1)):(log10($2))

这里把那两条拟合的直线省略了。选项 using 的使用比较复杂，这里很难说清楚， 请参阅 Gnuplot 手册。注意这时的坐标应该是线性的，而且我们应该调整绘图范 围：

	unset logscale
	set xrange [-2.2:2]
	set yrange [0:1.6]

命令 unset 删去一个变量的值。注意如果一个变量的默认值不是空的，那么 unset 和设为默认值是不一样的。如果你把 xtics 这个变量给 unset 了，画出 来的图就会没有 x 轴刻度，而不是把刻度设为默认的样子。

## EOF

噢，不对，还没完。这里再补充一句：Origin sucks!

真正的 EOF

