# html/template


## 简单使用

基本上`template`的使用也就是 ** 创建 New ** -> ** 解析模板 Parse,ParseFile ** -> ** 执行输出 Execute **.

	> t := template.New("Person Info")  // 创建一个 template
	> 1 > t, err := t.Parse(temp1)          // 解析
	> 2 > t, err := t.Parse("temp1.tpl") // 解析模板文件
	> err = t.Execute(os.Stdout, p)        // 输出到 io.Writer 接口

[查看示例](../demo/template/template1.go)

## 模板文件

Go语言的模板通过{{}}来包含需要在渲染时被替换的字段，{{.}}表示当前的对象，访问当前对象的字段通过{{.FieldName}}.

如果模板中输出{{.}}，这个一般应用与字符串对象，默认会调用fmt包输出字符串的内容。


```golang
type Person struct {
	Name    string
	Age     int
	Emails  []string
	Friends []*Friend
}

type Friend struct {
	Name  string
	Email string
}
```

输出普通的字段 `Name` 和 `Age`

```golang
姓名是： {{.Name}}, 年龄是 {{.Age}}.
```

> 字段首字母必须是大写的,否则在渲染的时候就会报错. 

__嵌套字段__

`{{range ...}} ... {{end}}` 这个和Go语法里面的range类似，循环操作数据
`{{with ...}} ... {{end}}` 操作是指当前对象的值，类似上下文的概念


```golang
{{range .Emails}}
	Email: {{.}}
{{end}}

{{with .Friends}}
	{{range .}}
你朋友： {{.Name}}
	{{end}}
{{end}}

```
[查看示例](../demo/template/template1.go)


### if-else 和 模板变量

if 后只跟通道数据,空则为false,不能跟表达式. `{{if .x}} ... {{else}} ... {{end}}`

```golang
{{with .Friends}}
	{{range .}}
		你朋友： {{.Name}}
		{{if .Email}}
			有Email: {{.}}
		{{else}}
			没有EMail地址
		{{end}}
	{{end}}
{{end}}
```

如果要使用模板变量,可以在 `{{range }}`,`{{with}}`,`{{if}}` 中声明变量,其作用域在 `{{end}}`之前. 和golang中有所区别的是变量名要加上`$`前缀,比如

```golang
{{range $index ,$mail := .Email }}
	{{if $index}}
		, 还有 {{$mail}}
	{{else}}
		有邮箱 {{$mail}}
	{{end}}
{{end}}
```

[查看示例](../demo/template/template2.go)


### 模板函数 和 pipelines(通道数据)

在Go语言里面任何`{{}}`里面的都是pipelines数据，利用pipe 和 模板函数则可以实现非常有用的功能.

```golang
// 将输出转化为HTML实体
{{. | html}}
```

此处 `html`作为函数来接受 `{{.}}` 并进行数据处理.


模板在输出对象的字段值时，采用了fmt包把对象转化成了字符串。但是有时候我们的需求可能不是这样的，例如有时候我们为了防止垃圾邮件发送者通过采集网页的方式来发送给我们的邮箱信息，我们希望把@替换成at例如：astaxie at beego.me，如果要实现这样的功能，我们就需要自定义函数来做这个功能。

每一个模板函数都有一个唯一值的名字，然后与一个Go函数关联，通过如下的方式来关联

```golang
type FuncMap map[string]interface{}
```

例如，如果我们想要的email函数的模板函数名是emailDeal，它关联的Go函数名称是EmailDealWith,n那么我们可以通过下面的方式来注册这个函数

```golang
t = t.Funcs(template.FuncMap{"emailDeal": EmailDealWith})
```

在模板包内部已经有内置的实现函数，下面代码截取自模板包里面

```golang
var builtins = FuncMap{
    "and":      and,
    "call":     call,
    "html":     HTMLEscaper,
    "index":    index,
    "js":       JSEscaper,
    "len":      length,
    "not":      not,
    "or":       or,
    "print":    fmt.Sprint,
    "printf":   fmt.Sprintf,
    "println":  fmt.Sprintln,
    "urlquery": URLQueryEscaper,
}
```


## 模板嵌套

对模板进行模块化,更好的组织模板的内容.利用 `define` 和 `template` 来嵌套模板.

定义模板
```
{{define "子模板名称"}} 内容 {{end}}
```
> 

调用则
```
{{template "子模板名称"}}
```

将所有节点加载，然后输出

	ExecuteTemplate(os.Stdout, "子模板", nil)


## Must 检测

模板包里面有一个函数Must，它的作用是检测模板是否正确，例如大括号是否匹配，注释是否正确的关闭，变量是否正确的书写。
