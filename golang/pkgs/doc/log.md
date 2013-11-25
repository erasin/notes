# log

轻量级别日志输出，Mutex同步。

**打印方式**

Print: 信息显示

	func Print(v ...interface{})
	func Printf(format string, v ...interface{})
	func Println(v ...interface{}) 

Fatal: Print 后 os.Exit(1)

	func Fatal(v ...interface{})
	func Fatalf(format string, v ...interface{})
	func Fatalln(v ...interface{}) 

Panic: Print Panic

	func Panic(v ...interface{})
	func Panicf(format string, v ...interface{})
	func Panicln(v ...interface{}) 

**其他**

前缀记录

	func SetPrefix(prefix string)
	func Prefix() string

flag 设定显示的时间等信息
	
	func SetFlags(flag int)
	func Flags() int

flag 常量
	
	const (
	    // Bits or'ed together to control what's printed. There is no control over the
	    // order they appear (the order listed here) or the format they present (as
	    // described in the comments).  A colon appears after these items:
	    //	2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
	    Ldate         = 1 << iota     // the date: 2009/01/23
	    Ltime                         // the time: 01:23:23
	    Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	    Llongfile                     // full file name and line number: /a/b/c/d.go:23
	    Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	    LstdFlags     = Ldate | Ltime // initial values for the standard logger
	)

输出日志

	func SetOutput(w io.Writer)

记录器

	func New(out io.Writer, prefix string, flag int) *Logger

[demo](../demo/log/log.go)