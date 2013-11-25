# flag

对命令行参数解析.

处理方式: flag 设定 -> 解析 -> 处理修改

[查看例子](../demo/flag/flag1.go)

常用的是 flag.String(), Bool(), Int()

```go
func Int(name string, value int, usage string) *int
func String(name string, value string, usage string) *string
func Bool(name string, value bool, usage string) *bool 
```

例子:

```golang
filename := flag.String("file", "default.txt", "to input a file name!")
flag.Parse() // 解析
fmt.Println(*file)
```

记得要使用 `flag.Parse()` 来显式解析.

如果使用存储对象,则使用对应 flag.StringVar() 

```go
func IntVar(p *int, name string, value int, usage string)
func String(p *string, name string, value string, usage string) *string
func Bool(p *bool, name string, value bool, usage string) *bool 
func Var(value Value, name string, usage string)
```

设定好flag参数后, 运行时使用`--help`会有默认的帮助信息显示.  
使用 `flag.Usage = func(){}` 来自定义帮助信息.

```go
flag.Usage = func(){
	fmt.Fprintf(os.Stderr,"Progame ver...\n")
	fmt.Fprintf(os.Stderr,"Usage of %s\n", os.Args[0])
	flag.PrintDefaults()
}
```

使用 `Args() string` 来获取非flag参数列表, `Arg(i int)string`则可以直接返回该值.   
相关的函数 `Nflag() int` 和 `NArg() int` 可以获得flag 的 flag参数处理的数量和非flag参数处理的数量.

`func Set(name, value string) error` 可以修改已经处理后的flag参数.

## 其他函数 

`func VisitAll(fn func(*Flag))` 历遍已经定过的flag参数.

`func Lookup(name string) *Flag` 来直接获得 flag.

flag 结构

```go
type Flag struct {
    Name     string // name as it appears on command line
    Usage    string // help message
    Value    Value  // value as set
    DefValue string // default value (as text); for usage message
}
```

> *flag 对象 的 value 要 `.String()` 

