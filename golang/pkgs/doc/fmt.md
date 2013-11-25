# fmt 

## 占位符

常规 General: 

	%v	相应值的默认格式。在打印结构体时，“加号”标记（%+v）会添加字段名
	%#v	相应值的Go语法表示
	%T	相应值的类型的Go语法表示
	%%	字面上的百分号，并非值的占位符

布尔 Boolean:

	%t	单词 true 或 false。

整数 Integer:

	%b	二进制表示
	%c	相应Unicode码点所表示的字符
	%d	十进制表示
	%o	八进制表示
	%q	单引号围绕的字符字面值，由Go语法安全地转义
	%x	十六进制表示，字母形式为小写 a-f
	%X	十六进制表示，字母形式为大写 A-F
	%U	Unicode格式：U+1234，等同于 "U+%04X"

浮点数及其复合构成 Floating-point and complex constituents:

	%b	无小数部分的，指数为二的幂的科学计数法，与 strconv.FormatFloat
		的 'b' 转换格式一致。例如 -123456p-78
	%e	科学计数法，例如 -1234.456e+78
	%E	科学计数法，例如 -1234.456E+78
	%f	有小数点而无指数，例如 123.456
	%g	根据情况选择 %e 或 %f 以产生更紧凑的（无末尾的0）输出
	%G	根据情况选择 %E 或 %f 以产生更紧凑的（无末尾的0）输出

字符串与字节切片 String and slice of bytes:

	%s	字符串或切片的无解译字节
	%q	双引号围绕的字符串，由Go语法安全地转义
	%x	十六进制，小写字母，每字节两个字符
	%X	十六进制，大写字母，每字节两个字符

指针Pointer：

	%p	十六进制表示，前缀 0x

对数值而言，宽度为该数值占用区域的最小宽度；精度为小数点之后的位数。 但对于 %g/%G 而言，精度为所有数字的总数。例如，对于123.45，格式 %6.2f 会打印123.45，而 %.4g 会打印123.5。%e 和 %f 的默认精度为6；但对于 %g 而言， 它的默认精度为确定该值所必须的最小位数。

其它标记：

	+	总打印数值的正负号；对于%q（%+q）保证只输出ASCII编码的字符。
	-	在右侧而非左侧填充空格（左对齐该区域）
	#	备用格式：为八进制添加前导 0（%#o），为十六进制添加前导 0x（%#x）或
		0X（%#X），为 %p（%#p）去掉前导 0x；如果可能的话，%q（%#q）会打印原始
		（即反引号围绕的）字符串；如果是可打印字符，%U（%#U）会写出该字符的
		Unicode 编码形式（如字符 x 会被打印成 U+0078 'x'）。
	' '	（空格）为数值中省略的正负号留出空白（% d）；
		以十六进制（% x, % X）打印字符串或切片时，在字节之间用空格隔开
	0	填充前导的0而非空格；
		对于数字，这会将填充移到正负号之后

## 函数

* 输出
	- func Errorf(format string, a ...interface{}) error
	- 标准输出 返回字节数
		- func Print(a ...interface{}) (n int, err error)
		- func Printf(format string, a ...interface{}) (n int, err error)
		- func Println(a ...interface{}) (n int, err error)
	- io.Writer 接口输出
		- func Fprint(w io.Writer, a ...interface{}) (n int, err error)
		- func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
		- func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
	- 返回 string 
		- func Sprint(a ...interface{}) string
		- func Sprintf(format string, a ...interface{}) string
		- func Sprintln(a ...interface{}) string
* 扫描
	- func Fscan(r io.Reader, a ...interface{}) (n int, err error)
	- func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)
	- func Fscanln(r io.Reader, a ...interface{}) (n int, err error)
	- func Scan(a ...interface{}) (n int, err error)
	- func Scanf(format string, a ...interface{}) (n int, err error)
	- func Scanln(a ...interface{}) (n int, err error)
	- func Sscan(str string, a ...interface{}) (n int, err error)
	- func Sscanf(str string, format string, a ...interface{}) (n int, err error)
	- func Sscanln(str string, a ...interface{}) (n int, err error)


一组类似的函数通过扫描已格式化的文本来产生值。Scan、Scanf 和 Scanln 从 os.Stdin 中读取；Fscan、Fscanf 和 Fscanln 从指定的 io.Reader 中读取； Sscan、Sscanf 和 Sscanln 从实参字符串中读取。Scanln、Fscanln 和 Sscanln 在换行符处停止扫描，且需要条目紧随换行符之后；Scanf、Fscanf 和 Sscanf 需要输入换行符来匹配格式中的换行符；其它函数则将换行符视为空格。

格式化行为类似于 Printf，但也有如下例外：

	%p 没有实现
	%T 没有实现
	%e %E %f %F %g %G 都完全等价，且可扫描任何浮点数或复合数值
	%s 和 %v 在扫描字符串时会将其中的空格作为分隔符
	标记 # 和 + 没有实现

所有需要被扫描的实参都必须是基本类型或 Scanner 接口的实现。



## 常用



## other

[fmt-zh](http://zh.golanger.com/pkg/fmt/)