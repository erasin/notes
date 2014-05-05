# golang-Tour

by: **golang.org**

## 01 Hello, 世界

欢迎来到Go 编程语言指南。

该指南被分为三个部分：基础概念、方法和接口，以及并发。

在指南后有若干个练习需要读者完成。

该指南可以进行交互。点击“运行”按钮（或按 Shift + Enter）可以在 远程服务器上 编译并执行程序。 结果展示在代码的下面。

这些例子程序展示了 Go 的各个方面。在指南中的程序可以成为你积累经验的开始。

编辑程序并且再次执行它。

当你准备好继续了，点击右下的“向后”按钮或按 PageDown 键。 也可以使用页面顶端“Go”标志下面的菜单进行导航。

实例

	package main

	import (
		"fmt"
	)

	func main() {
		fmt.Println("string")
	}


## Go Playground

这个指南构建在 Go Playground 之上，这是一个运行在 golang.org 的服务器上的一个 Web 服务。

服务接收 Go 程序的输入，且在沙盒里编译、链接和运行， 然后返回输出。

对于在 Playground 运行的程序的限制是：

除了部分例外，Playground 可以使用大部分的标准库；尤其是关于访问网络和文件系统。因此，在 Playground 中的程序与外部唯一的通讯方式是使用标准输出和标准错误输出。
Playground 的时间从 2009-11-10 23:00:00 UTC（了解这个日期的重要含义是留给读者的练习）。这使得根据可预见的输出来缓存程序变得容易。
对于运行时间、CPU 和内存的使用同样也有限制，并且程序是限制单一线程运行（但是可以有多个 goroutine）。
Playground 使用最后发布的 Go 的稳定版本。

实例

	package main

	import (
	    "fmt"
	    "net"
	    "os"
	    "time"
	)

	func main() {
	    fmt.Println("Welcome to the playground!")

	    fmt.Println("The time is", time.Now())

	    fmt.Println("And if you try to open a file:")
	    fmt.Println(os.Open("filename"))

	    fmt.Println("Or access the network:")
	    fmt.Println(net.Dial("tcp", "google.com"))
	}

## Packages

每个 Go 程序都是由包组成的。

程序运行的入口是包 `main`。

这个程序使用并导入了包 "fmt" 和 `"math/rand"`。

按照惯例，包名与导入路径的最后一个目录一致。例如，`"math/rand"` 包由 package rand 语句开始。

注意： 这个程序的运行环境是固定的，因此 rand.Intn 总是会返回相同的数字。

（为了得到不同的数字，需要生成不同的种子数，参阅 rand.Seed。）

实例

	package main

	import (
	    "fmt"
	    "math/rand"
	)

	func main() {
	    fmt.Println("My favorite number is", rand.Intn(10))
	}

## 导入

这个代码用圆括号组合了导入，这是“factored”导入语句。同样可以编写多个导入语句，例如：

	import "fmt"
	import "math"

实例

	package main

	import (
	    "fmt"
	    "math"
	)

	func main() {
	    fmt.Printf("Now you have %g problems.",
	        math.Nextafter(2, 3))
	}

## 导出名

在导入了一个包之后，就可以用其导出的名称来调用它。

在 Go 中，首字母大写的名称是被导出的。

Foo 和 FOO 都是被导出的名称。名称 foo 是不会被导出的。

执行代码。然后将 math.pi 改名为 math.Pi 再试着执行一下。

实例

	package main

	import (
	    "fmt"
	    "math"
	)

	func main() {
	    fmt.Println(math.pi)
	}

## 函数

函数可以没有参数或接受多个参数。

在这个例子中，`add` 接受两个 int 类型的参数。

注意类型在变量名 _之后_。

（参考 这篇关于 Go 语法定义的文章了解类型以这种形式出现的原因。）

	package main

	import "fmt"

	func add(x int, y int) int {
	    return x + y
	}

	func main() {
	    fmt.Println(add(42, 13))
	}

## 函数（续）

当两个或多个连续的函数命名参数是同一类型，则除了最后一个类型之外，其他都可以省略。

在这个例子中 ， `x int, y int` 被缩写为 `x, y int`

	package main

	import "fmt"

	func add(x, y int) int {
	    return x + y
	}

	func main() {
	    fmt.Println(add(42, 13))
	}


## 多值返回

函数可以返回任意数量的返回值。

这个函数返回了两个字符串。


	package main

	import "fmt"

	func swap(x, y string) (string, string) {
	    return y, x
	}

	func main() {
	    a, b := swap("hello", "world")
	    fmt.Println(a, b)
	}

## 命名返回值

函数接受参数。在 Go 中，函数可以返回多个“结果参数”，而不仅仅是一个值。它们可以像变量那样命名和使用。

如果命名了返回值参数，一个没有参数的 return 语句，会将当前的值作为返回值返回。

	package main

	import "fmt"

	func split(sum int) (x, y int) {
	    x = sum * 4 / 9
	    y = sum - x
	    return
	}

	func main() {
	    fmt.Println(split(17))
	}

## 变量

var 语句定义了一个变量的列表；跟函数的参数列表一样，类型在后面。


	package main

	import "fmt"

	var i int
	var c, python, java bool

	func main() {
	    fmt.Println(i, c, python, java)
	}

## 初始化变量

变量定义可以包含初始值，每个变量对应一个。

如果初始化是使用表达式，则可以省略类型；变量从初始值中获得类型。

	package main

	import "fmt"

	var i, j int = 1, 2
	var c, python, java = true, false, "no!"

	func main() {
	    fmt.Println(i, j, c, python, java)
	}

## 短声明变量

在函数中，`:=` 简洁赋值语句在明确类型的地方，可以用于替代 var 定义。

函数外的每个语法块都必须以关键字开始（`var`、`func`、等等），`:=` 结构不能使用在函数外。

	package main

	import "fmt"

	func main() {
	    var i, j int = 1, 2
	    k := 3
	    c, python, java := true, false, "no!"

	    fmt.Println(i, j, k, c, python, java)
	}

## 基本类型

Go 的基本类型有Basic types

	┌──────┬─────────┬────────┬─────────┬───────────┬────────────┐
	│ bool │ string  │        │         │           │            │
	├──────┼─────────┼────────┼─────────┼───────────┼────────────┤
	│ int  │ int8    │ int16  │ int32   │ int64     │            │
	│      │         │        │ rune    │           │            │
	├──────┼─────────┼────────┼─────────┼───────────┼────────────┤
	│ uint │ uint8   │ uint16 │ uint32  │ uint64    │ uintptr    │
	│      │ byte    │        │         │           │            │
	├──────┼─────────┼────────┼─────────┼───────────┼────────────┤
	│      │         │        │ float32 │ float64   │            │
	├──────┼─────────┼────────┼─────────┼───────────┼────────────┤
	│      │         │        │         │ complex64 │ complex128 │
	└──────┴─────────┴────────┴─────────┴───────────┴────────────┘

**示例**

	package main

	import (
	    "fmt"
	    "math/cmplx"
	)

	var (
	    ToBe   bool       = false
	    MaxInt uint64     = 1<<64 - 1
	    z      complex128 = cmplx.Sqrt(-5 + 12i)
	)

	func main() {
	    const f = "%T(%v)\n"
	    fmt.Printf(f, ToBe, ToBe)
	    fmt.Printf(f, MaxInt, MaxInt)
	    fmt.Printf(f, z, z)
	}



## 类型转换

表达式 T(v) 将值 v 转换为类型 `T`。

一些关于数值的转换：

	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)

或者，更加简单的形式：

	i := 42
	f := float64(i)
	u := uint(f)

与 C 不同的是 Go 的在不同类型之间的项目赋值时需要显式转换。 试着移除例子中 float64 或 int 的转换看看会发生什么。


	package main

	import (
	    "fmt"
	    "math"
	)

	func main() {
	    var x, y int = 3, 4
	    var f float64 = math.Sqrt(float64(3*3 + 4*4))
	    var z int = int(f)
	    fmt.Println(x, y, z)
	}

## 常量

常量的定义与变量类似，只不过使用 const 关键字。

常量可以是字符、字符串、布尔或数字类型的值。

常量不能使用 := 语法定义。


	package main

	import "fmt"

	const Pi = 3.14

	func main() {
	    const World = "世界"
	    fmt.Println("Hello", World)
	    fmt.Println("Happy", Pi, "Day")

	    const Truth = true
	    fmt.Println("Go rules?", Truth)
	}


## 数值常量

数值常量是高精度的 _值_。

一个未指定类型的常量由上下文来决定其类型。

也尝试一下输出 `needInt(Big)` 吧。

	package main

	import "fmt"

	const (
	    Big   = 1 << 100
	    Small = Big >> 99
	)

	func needInt(x int) int { return x*10 + 1 }
	func needFloat(x float64) float64 {
	    return x * 0.1
	}

	func main() {
	    fmt.Println(needInt(Small))
	    fmt.Println(needFloat(Small))
	    fmt.Println(needFloat(Big))
	}

## for

Go 只有一种循环结构——`for` 循环。

基本的 for 循环除了没有了 `( )` 之外（甚至强制不能使用它们），看起来跟 C 或者 Java 中做的一样，而 `{ }` 是必须的。

	package main

	import "fmt"

	func main() {
	    sum := 0
	    for i := 0; i < 10; i++ {
	        sum += i
	    }
	    fmt.Println(sum)
	}


## for（续）

跟 C 或者 Java 中一样，可以让前置、后置语句为空。

	package	main

	import "fmt"

	func main() {
	    sum := 1
	    for ; sum < 1000; {
	        sum += sum
	    }
	    fmt.Println(sum)
	}

## for 是 Go 的 “while”

基于此可以省略分号：C 的 while 在 Go 中叫做 `for`。

	package main

	import "fmt"

	func main() {
	    sum := 1
	    for sum < 1000 {
	        sum += sum
	    }
	    fmt.Println(sum)
	}

## 死循环

如果省略了循环条件，循环就不会结束，因此可以用更简洁地形式表达死循环。

	package main

	func main() {
	    for {
	    }
	}

## if

if 语句除了没有了 `( )` 之外（甚至强制不能使用它们），看起来跟 C 或者 Java 中的一样，而 `{ }` 是必须的。

（耳熟吗？）


	package main

	import (
	    "fmt"
	    "math"
	)

	func sqrt(x float64) string {
	    if x < 0 {
	        return sqrt(-x) + "i"
	    }
	    return fmt.Sprint(math.Sqrt(x))
	}

	func main() {
	    fmt.Println(sqrt(2), sqrt(-4))
	}

## if 的便捷语句

跟 for 一样，`if` 语句可以在条件之前执行一个简单的语句。

由这个语句定义的变量的作用域仅在 if 范围之内。

（在最后的 return 语句处使用 v 看看。）


	package main

	import (
	    "fmt"
	    "math"
	)

	func pow(x, n, lim float64) float64 {
	    if v := math.Pow(x, n); v < lim {
	        return v
	    }
	    return lim
	}

	func main() {
	    fmt.Println(
	        pow(3, 2, 10),
	        pow(3, 3, 20),
	    )
	}

## if 和 else

在 if 的便捷语句定义的变量同样可以在任何对应的 else 块中使用。

	package main

	import (
	    "fmt"
	    "math"
	)

	func pow(x, n, lim float64) float64 {
	    if v := math.Pow(x, n); v < lim {
	        return v
	    } else {
	        fmt.Printf("%g >= %g\n", v, lim)
	    }
	    // 这里开始就不能使用 v 了
	    return lim
	}

	func main() {
	    fmt.Println(
	        pow(3, 2, 10),
	        pow(3, 3, 20),
	    )
	}

## 练习：循环和函数

作为练习函数和循环的简单途径，用牛顿法实现开方函数。

在这个例子中，牛顿法是通过选择一个初始点 z 然后重复这一过程求 Sqrt(x) 的近似值：

![newton](img/newton.png)

为了做到这个，只需要重复计算 10 次，并且观察不同的值（1，2，3，……）是如何逐步逼近结果的。 然后，修改循环条件，使得当值停止改变（或改变非常小）的时候退出循环。观察迭代次数是否变化。结果与 [math.Sqrt](http://golang.org/pkg/math/#Sqrt) 接近吗？

提示：定义并初始化一个浮点值，向其提供一个浮点语法或使用转换：

	z := float64(1)
	z := 1.0

**示例**

	package main

	import (
	    "fmt"
	)

	func Sqrt(x float64) float64 {
	}

	func main() {
	    fmt.Println(Sqrt(2))
	}

## 结构体

一个结构体（`struct`）就是一个字段的集合。

（而 type 的含义跟其字面意思相符。）

	package main

	import "fmt"

	type Vertex struct {
	    X int
	    Y int
	}

	func main() {
	    fmt.Println(Vertex{1, 2})
	}

## 结构体字段

结构体字段使用点号来访问。

	package main

	import "fmt"

	type Vertex struct {
	    X int
	    Y int
	}

	func main() {
	    v := Vertex{1, 2}
	    v.X = 4
	    fmt.Println(v.X)
	}

## 指针

Go 有指针，但是没有指针运算。

结构体字段可以通过结构体指针来访问。通过指针间接的访问是透明的。

	package main

	import "fmt"

	type Vertex struct {
	    X int
	    Y int
	}

	func main() {
	    p := Vertex{1, 2}
	    q := &p
	    q.X = 1e9
	    fmt.Println(p)
	}


## 结构体文法

结构体文法表示通过结构体字段的值作为列表来新分配一个结构体。

使用 Name: 语法可以仅列出部分字段。（字段名的顺序无关。）

特殊的前缀 & 构造了指向结构体的指针。


	package main

	import "fmt"

	type Vertex struct {
	    X, Y int
	}

	var (
	    p = Vertex{1, 2}  // 类型为 Vertex
	    q = &Vertex{1, 2} // 类型为 *Vertex
	    r = Vertex{X: 1}  // Y:0 被省略
	    s = Vertex{}      // X:0 和 Y:0
	)

	func main() {
	    fmt.Println(p, q, r, s)
	}

## new 函数

表达式 new(T) 分配了一个零初始化的 T 值，并返回指向它的指针。

`var t *T = new(T)` 或 `t := new(T)`

	package main

	import "fmt"

	type Vertex struct {
	    X, Y int
	}

	func main() {
	    v := new(Vertex)
	    fmt.Println(v)
	    v.X, v.Y = 11, 9
	    fmt.Println(v)
	}


## 数组

类型 [n]T 是一个有 n 个类型为 T 的值的数组。

表达式

	var a [10]int

定义变量 a 是一个有十个整数的数组。

数组的长度是其类型的一部分，因此数组不能改变大小。 这看起来是一个制约，但是请不要担心；Go 提供了更加便利的方式来使用数组。

	package main

	import "fmt"

	func main() {
	    var a [2]string
	    a[0] = "Hello"
	    a[1] = "World"
	    fmt.Println(a[0], a[1])
	    fmt.Println(a)
	}

## slice

一个 slice 会指向一个数组，并且包含了长度信息。

`[]T` 是一个元素类型为 T 的 slice。

	package main

	import "fmt"

	func main() {
	    p := []int{2, 3, 5, 7, 11, 13}
	    fmt.Println("p ==", p)

	    for i := 0; i < len(p); i++ {
	        fmt.Printf("p[%d] == %d\n", i, p[i])
	    }
	}



## 对 slice 切片

slice 可以重新切片，创建一个新的 slice 值指向相同的数组。

表达式

	s[lo:hi]

表示从 lo 到 hi-1 的 slice 元素，含两端。因此

	s[lo:lo]

是空的，而

	s[lo:lo+1]

有一个元素。

	package main

	import "fmt"

	func main() {
	    p := []int{2, 3, 5, 7, 11, 13}
	    fmt.Println("p ==", p)
	    fmt.Println("p[1:4] ==", p[1:4])

	    // 省略下标代表从 0 开始
	    fmt.Println("p[:3] ==", p[:3])

	    // 省略上标代表到 len(s) 结束
	    fmt.Println("p[4:] ==", p[4:])
	}

## 构造 slice

slice 由函数 make 创建。这会分配一个零长度的数组并且返回一个 slice 指向这个数组：

	a := make([]int, 5)  // len(a)=5

为了指定容量，可传递第三个参数到 `make`：

	b := make([]int, 0, 5) // len(b)=0, cap(b)=5

	b = b[:cap(b)] // len(b)=5, cap(b)=5
	b = b[1:]      // len(b)=4, cap(b)=4

示例

	package main

	import "fmt"

	func main() {
	    a := make([]int, 5)
	    printSlice("a", a)
	    b := make([]int, 0, 5)
	    printSlice("b", b)
	    c := b[:2]
	    printSlice("c", c)
	    d := c[2:5]
	    printSlice("d", d)
	}

	func printSlice(s string, x []int) {
	    fmt.Printf("%s len=%d cap=%d %v\n",
	        s, len(x), cap(x), x)
	}


## 空 slice

slice 的零值是 `nil`。

一个 nil 的 slice 的长度和容量是 0。

（了解更多关于 slice 的内容，参阅文章slice：[使用和内幕](http://golang.org/doc/articles/gos_declaration_syntax.html)。）

	package main

	import "fmt"

	func main() {
	    var z []int
	    fmt.Println(z, len(z), cap(z))
	    if z == nil {
	        fmt.Println("nil!")
	    }
	}

## range

for 循环的 range 格式可以对 slice 或者 map 进行迭代循环。

	package main

	import "fmt"

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	func main() {
	    for i, v := range pow {
	        fmt.Printf("2**%d = %d\n", i, v)
	    }
	}

## range（续）

可以将值赋值给 _ 来忽略序号和值。

如果只需要索引值，去掉“, value”的部分即可。

	package main

	import "fmt"

	func main() {
	    pow := make([]int, 10)
	    for i := range pow {
	        pow[i] = 1 << uint(i)
	    }
	    for _, value := range pow {
	        fmt.Printf("%d\n", value)
	    }
	}

## 练习：slice

实现 `Pic`。它返回一个 slice 的长度 `dy`，和 slice 中每个元素的长度的 8 位无符号整数 `dx`。当执行这个程序，它会将整数转换为灰度（好吧，蓝度）图片进行展示。

图片的实现已经完成。可能用到的函数包括 >x^y ， (x+y)/2 和 `x*y`。

（需要使用循环来分配 [][]uint8 中的每个 `[]uint8`。）

（使用 uint8(intValue) 在类型之间进行转换。）


	package main

	import "code.google.com/p/go-tour/pic"

	func Pic(dx, dy int) [][]uint8 {
	}

	func main() {
	    pic.Show(Pic)
	}

## map

map 映射键到值。

map 在使用之前必须用 make 而不是 new 来创建；值为 nil 的 map 是空的，并且不能赋值。

	package main

	import "fmt"

	type Vertex struct {
	    Lat, Long float64
	}

	var m map[string]Vertex

	func main() {
	    m = make(map[string]Vertex)
	    m["Bell Labs"] = Vertex{
	        40.68433, -74.39967,
	    }
	    fmt.Println(m["Bell Labs"])
	}

## map 的文法

map 的文法跟结构体文法相似，不过必须有键名。

	package main

	import "fmt"

	type Vertex struct {
	    Lat, Long float64
	}

	var m = map[string]Vertex{
	    "Bell Labs": Vertex{
	        40.68433, -74.39967,
	    },
	    "Google": Vertex{
	        37.42202, -122.08408,
	    },
	}

	func main() {
	    fmt.Println(m)
	}

## map 的文法（续）

如果顶级的类型只有类型名的话，可以在文法的元素中省略键名。

	package main

	import "fmt"

	type Vertex struct {
	    Lat, Long float64
	}

	var m = map[string]Vertex{
	    "Bell Labs": {40.68433, -74.39967},
	    "Google":    {37.42202, -122.08408},
	}

	func main() {
	    fmt.Println(m)
	}

## 修改 map

在 map m 中插入或修改一个元素：

	m[key] = elem

获得元素：

	elem = m[key]

删除元素：

	delete(m, key)

通过双赋值检测某个键存在：

	elem, ok = m[key]

如果 key 在 m 中，`ok` 为 true 。否则， ok 为 `false`，并且 elem 是 map 的元素类型的零值。

同样的，当从 map 中读取某个不存在的键时，结果是 map 的元素类型的零值。

	package main

	import "fmt"

	func main() {
	    m := make(map[string]int)

	    m["Answer"] = 42
	    fmt.Println("The value:", m["Answer"])

	    m["Answer"] = 48
	    fmt.Println("The value:", m["Answer"])

	    delete(m, "Answer")
	    fmt.Println("The value:", m["Answer"])

	    v, ok := m["Answer"]
	    fmt.Println("The value:", v, "Present?", ok)
	}

## 练习：map

实现 `WordCount`。它应当返回一个含有 s 中每个 “词” 个数的 map。函数 wc.Test 针对这个函数执行一个测试用例，并打印成功或者失败。

你会发现 [strings.Fields](http://golang.org/pkg/strings/#Fields) 很有帮助。

	package main

	import (
	    "code.google.com/p/go-tour/wc"
	)

	func WordCount(s string) map[string]int {
	    return map[string]int{"x": 1}
	}

	func main() {
	    wc.Test(WordCount)
	}


## 函数为值

函数也是值。Function values

	package main

	import (
	    "fmt"
	    "math"
	)

	func main() {
	    hypot := func(x, y float64) float64 {
	        return math.Sqrt(x*x + y*y)
	    }

	    fmt.Println(hypot(3, 4))
	}

## 函数的闭包

Go 函数可以是闭包的。闭包是一个函数值，它来自函数体的外部的变量引用。函数可以对这个引用值进行访问和赋值；换句话说这个函数被“绑定”在这个变量上。

例如，函数 adder 返回一个闭包。每个闭包都被绑定到其各自的 sum 变量上。

	package main

	import "fmt"

	func adder() func(int) int {
	    sum := 0
	    return func(x int) int {
	        sum += x
	        return sum
	    }
	}

	func main() {
	    pos, neg := adder(), adder()
	    for i := 0; i < 10; i++ {
	        fmt.Println(
	            pos(i),
	            neg(-2*i),
	        )
	    }
	}

## 练习：斐波纳契闭包

现在来通过函数做些有趣的事情。

实现一个 fibonacci 函数，返回一个函数（一个闭包）可以返回连续的斐波纳契数。


	package main

	import "fmt"

	// fibonacci 函数会返回一个返回 int 的函数。
	func fibonacci() func() int {
	}

	func main() {
	    f := fibonacci()
	    for i := 0; i < 10; i++ {
	        fmt.Println(f())
	    }
	}

结果


	package main

	import "fmt"
	import "log"

	func fibonacci() func() int {
		f0 := 0
		f1 := 1
		return func() (f2 int) {
			f2 = f0 + f1
			f0 = f1
			f1 = f2
			return
		}
	}

	func getFibonacci(n int) (fi []int) {
		// func getFibonacci(n int) (fi int) {

		defer func() {
			if x := recover(); x != nil {
				log.Println("Panic:", x)
			} else {
				log.Println("没有 Panic!,输出正确的值")
			}
		}()

		f := fibonacci()

		// 必须给出长度 否则溢出
		fi = make([]int, n)

		for i := 0; i < n; i++ {
			fi[i] = int(f())
		}
		return
	}

	func main() {
		f := fibonacci()
		for i := 0; i < 10; i++ {
			fmt.Println(f())
		}

		fmt.Printf("fibonacci 12 is %d \n", getFibonacci(12))
	}



## switch

你可能已经猜到 switch 可能的形式了。

除非使用 fallthrough 语句作为结尾，否则 case 部分会自动终止。

	package main

	import (
	    "fmt"
	    "runtime"
	)

	func main() {
	    fmt.Print("Go runs on ")
	    switch os := runtime.GOOS; os {
	    case "darwin":
	        fmt.Println("OS X.")
	    case "linux":
	        fmt.Println("Linux.")
	    default:
	        // freebsd, openbsd,
	        // plan9, windows...
	        fmt.Printf("%s.", os)
	    }
	}

## switch 的执行顺序

switch 的条件从上到下的执行，当匹配成功的时候停止。

（例如，

	switch i {
		case 0:
		case f():
	}

当 i==0 时不会调用 `f`。）


注意：Go playground 中的时间总是从 2009-11-10 23:00:00 UTC 开始， 如何校验这个值作为一个练习留给读者完成。



	package main

	import (
	    "fmt"
	    "time"
	)

	func main() {
	    fmt.Println("When's Saturday?")
	    today := time.Now().Weekday()
	    switch time.Saturday {
	    case today + 0:
	        fmt.Println("Today.")
	    case today + 1:
	        fmt.Println("Tomorrow.")
	    case today + 2:
	        fmt.Println("In two days.")
	    default:
	        fmt.Println("Too far away.")
	    }
	}

## 没有条件的 switch

没有条件的 switch 同 `switch true` 一样。

这一构造使得可以用更清晰的形式来编写长的 if-then-else 链。


	package main

	import (
	    "fmt"
	    "time"
	)

	func main() {
	    t := time.Now()
	    switch {
	    case t.Hour() < 12:
	        fmt.Println("Good morning!")
	    case t.Hour() < 17:
	        fmt.Println("Good afternoon.")
	    default:
	        fmt.Println("Good evening.")
	    }
	}

## 进阶练习：复数立方根

让我们通过 complex64 和 complex128 来探索一下 Go 内建的复数。对于立方根，牛顿法需要大量循环：

![newton3](img/newton3.png)

找到 2 的立方根，确保算法能够工作。在 math/cmplx 包中有 Pow 函数。

	package main

	import "fmt"

	func Cbrt(x complex128) complex128 {
	}

	func main() {
	    fmt.Println(Cbrt(2))
	}

答案

	package main

	import (
		"fmt"
		"math/cmplx"
	)

	func Cbrt(x complex128) (y complex128) {
		y = cmplx.Pow(x, 3)
		return
	}

	func main() {
		fmt.Println(Cbrt(2))
	}


## 方法和接口

下面的一组幻灯包含了方法和接口，以及定义对象和其行为的结构。

## 方法

Go 没有类。然而，仍然可以在结构体类型上定义方法。

方法接收者 出现在 func 关键字和方法名之间的参数中。


	package main

	import (
	    "fmt"
	    "math"
	)

	type Vertex struct {
	    X, Y float64
	}

	func (v *Vertex) Abs() float64 {
	    return math.Sqrt(v.X*v.X + v.Y*v.Y)
	}

	func main() {
	    v := &Vertex{3, 4}
	    fmt.Println(v.Abs())
	}

## 方法（续）

事实上，可以对包中的 任意 类型定义任意方法，而不仅仅是针对结构体。

不能对来自其他包的类型或基础类型定义方法。

	package main

	import (
	    "fmt"
	    "math"
	)

	type MyFloat float64

	func (f MyFloat) Abs() float64 {
	    if f < 0 {
	        return float64(-f)
	    }
	    return float64(f)
	}

	func main() {
	    f := MyFloat(-math.Sqrt2)
	    fmt.Println(f.Abs())
	}

## 接收者为指针的方法

方法可以与命名类型或命名类型的指针关联。

刚刚看到的两个 Abs 方法。一个是在 *Vertex 指针类型上，而另一个在 MyFloat 值类型上。 有两个原因需要使用指针接收者。首先避免在每个方法调用中拷贝值（如果值类型是大的结构体的话会更有效率）。其次，方法可以修改接收者指向的值。

尝试修改 Abs 的定义，同时 Scale 方法使用 Vertex 代替 *Vertex 作为接收者。

当 v 是 Vertex 的时候 Scale 方法没有任何作用。`Scale` 修改 `v`。当 v 是一个值（非指针），方法看到的是 Vertex 的副本，并且无法修改原始值。

Abs 的工作方式是一样的。只不过，仅仅读取 `v`。所以读取的是原始值（通过指针）还是那个值的副本并没有关系。

	package main

	import (
	    "fmt"
	    "math"
	)

	type Vertex struct {
	    X, Y float64
	}

	func (v *Vertex) Scale(f float64) {
	    v.X = v.X * f
	    v.Y = v.Y * f
	}

	func (v *Vertex) Abs() float64 {
	    return math.Sqrt(v.X*v.X + v.Y*v.Y)
	}

	func main() {
	    v := &Vertex{3, 4}
	    v.Scale(5)
	    fmt.Println(v, v.Abs())
	}

## 接口

接口类型是由一组方法定义的集合。

接口类型的值可以存放实现这些方法的任何值。

Note: The code on the left fails to compile.

Vertex doesn't satisfy Abser because the Abs method is defined only on *Vertex, not Vertex.


	package main

	import (
	    "fmt"
	    "math"
	)

	type Abser interface {
	    Abs() float64
	}

	func main() {
	    var a Abser
	    f := MyFloat(-math.Sqrt2)
	    v := Vertex{3, 4}

	    a = f  // a MyFloat 实现了 Abser
	    a = &v // a *Vertex 实现了 Abser

	    // 下面一行，v 是一个 Vertex（而不是 *Vertex）
	    // 所以没有实现 Abser。
	    a = v

	    fmt.Println(a.Abs())
	}

	type MyFloat float64

	func (f MyFloat) Abs() float64 {
	    if f < 0 {
	        return float64(-f)
	    }
	    return float64(f)
	}

	type Vertex struct {
	    X, Y float64
	}

	func (v *Vertex) Abs() float64 {
	    return math.Sqrt(v.X*v.X + v.Y*v.Y)
	}

## 隐式接口

类型通过实现那些方法来实现接口。

没有显式声明的必要。

隐式接口解藕了实现接口的包和定义接口的包：互不依赖。

因此，也就无需在每一个实现上增加新的接口名称，这样同时也鼓励了明确的接口定义。

包 io 定义了 Reader 和 `Writer`；其实不一定要这么做。

	package main

	import (
	    "fmt"
	    "os"
	)

	type Reader interface {
	    Read(b []byte) (n int, err error)
	}

	type Writer interface {
	    Write(b []byte) (n int, err error)
	}

	type ReadWriter interface {
	    Reader
	    Writer
	}

	func main() {
	    var w Writer

	    // os.Stdout 实现了 Writer
	    w = os.Stdout

	    fmt.Fprintf(w, "hello, writer\n")
	}

## 错误

错误是可以用字符串描述自己的任何东西。主要思路是由预定义的内建接口类型 `error`，和方法 `Error`，返回字符串：

	type error interface {
	    Error() string
	}

当用 fmt 包的多种不同的打印函数输出一个 error 时，会自动的调用该方法。

	package main

	import (
	    "fmt"
	    "time"
	)

	type MyError struct {
	    When time.Time
	    What string
	}

	func (e *MyError) Error() string {
	    return fmt.Sprintf("at %v, %s",
	        e.When, e.What)
	}

	func run() error {
	    return &MyError{
	        time.Now(),
	        "it didn't work",
	    }
	}

	func main() {
	    if err := run(); err != nil {
	        fmt.Println(err)
	    }
	}

## 练习：错误

从之前的练习中复制 Sqrt 函数，并修改使其返回 error 值。

Sqrt 接收到一个负数时，应当返回一个非 nil 的错误值。复数同样也不被支持。

创建一个新类型

	type ErrNegativeSqrt float64

为其实现

	func (e ErrNegativeSqrt) Error() string

使其成为一个 `error`， 该方法就可以让 ErrNegativeSqrt(-2).Error() 返回 `"cannot Sqrt negative number: -2"`。

注意： 在 Error 方法内调用 `fmt.Print(e)` 将会让程序陷入死循环。可以通过先转换 e 来避免这个问题：`fmt.Print(float64(e))`。请思考这是为什么呢？

修改 Sqrt 函数，使其接受一个负数时，返回 ErrNegativeSqrt 值。

	package main

	import (
	    "fmt"
	)

	func Sqrt(f float64) (float64, error) {
	    return 0, nil
	}

	func main() {
	    fmt.Println(Sqrt(2))
	    fmt.Println(Sqrt(-2))
	}

## Web 服务器

包 http 通过任何实现了 http.Handler 的值来响应 HTTP 请求：

	package http
	type Handler interface {
	    ServeHTTP(w ResponseWriter, r *Request)
	}

在这个例子中，类型 Hello 实现了 `http.Handler`。

访问 <http://localhost:4000/> 会看到来自程序的问候。

注意： 这个例子无法在基于 web 的指南用户界面运行。为了尝试编写 web 服务器，可能需要[安装 Go](http://golang.org/doc/install/)。

	package main

	import (
	    "fmt"
	    "net/http"
	)

	type Hello struct{}

	func (h Hello) ServeHTTP(
	    w http.ResponseWriter,
	    r *http.Request) {
	    fmt.Fprint(w, "Hello!")
	}

	func main() {
	    var h Hello
	    http.ListenAndServe("localhost:4000", h)
	}


## 练习：HTTP 处理

实现下面的类型，并在其上定义 ServeHTTP 方法。在 web 服务器中注册它们来处理指定的路径。

	type String string
	type Struct struct {
	    Greeting string
	    Punct    string
	    Who      string
	}

例如，可以使用如下方式注册处理方法：

	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})

基本

	package main

	import (
	    "net/http"
	)

	func main() {
	    // your http.Handle calls here
	    http.ListenAndServe("localhost:4000", nil)
	}


## 图片

Package image 定义了 Image 接口：

	package image

	type Image interface {
		ColorModel() color.Model
		Bounds() Rectangle
		At(x, y int) color.Color
	}

（参阅[文档](http://golang.org/pkg/image/#Image)了解全部信息。）

同样，`color.Color` 和 color.Model 也是接口，但是通常因为直接使用预定义的实现 image.RGBA 和 image.RGBAModel 而被忽视了。这些接口和类型由[image/color](http://golang.org/pkg/image/color/) 包定义。


	package main

	import (
	    "fmt"
	    "image"
	)

	func main() {
	    m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	    fmt.Println(m.Bounds())
	    fmt.Println(m.At(0, 0).RGBA())
	}


## 练习：图片

还记得之前编写的图片生成器吗？现在来另外编写一个，不过这次将会返回 `image.Image` 来代替 `slice` 的数据。

自定义的 Image 类型，要实现[必要的方法](http://golang.org/pkg/image/#Image)，并且调用 `pic.ShowImage`。

Bounds 应当返回一个 `image.Rectangle`，例如 `image.Rect(0, 0, w, h)`。

ColorModel 应当返回 `image.RGBAModel`。

At 应当返回一个颜色；在这个例子里，在最后一个图片生成器的值 v 匹配 `color.RGBA{v, v, 255, 255}`。

	package main

	import (
	    "code.google.com/p/go-tour/pic"
	    "image"
	)

	type Image struct{}

	func main() {
	    m := Image{}
	    pic.ShowImage(m)
	}

## 练习：Rot13 读取器

一般的模式是 [io.Reader](http://golang.org/pkg/io/#Reader) 包裹另一个 `io.Reader`，用某些途径修改特定的流。

例如，[gzip.NewReader](http://golang.org/pkg/compress/gzip/#NewReader) 函数输入一个 `io.Reader`（gzip 的数据流）并且返回一个同样实现了 io.Reader 的 `*gzip.Reader`（解压缩后的数据流）。

实现一个实现了 io.Reader 的 `rot13Reader`，用 [ROT13](http://en.wikipedia.org/wiki/ROT13) 修改数据流中的所有的字母进行密文替换。

rot13Reader 已经提供。通过实现其 Read 方法使得它匹配 `io.Reader`。

	package main

	import (
	    "io"
	    "os"
	    "strings"
	)

	type rot13Reader struct {
	    r io.Reader
	}

	func main() {
	    s := strings.NewReader(
	        "Lbh penpxrq gur pbqr!")
	    r := rot13Reader{s}
	    io.Copy(os.Stdout, &r)
	}

## 并发

接下来的章节涵盖了 Go 的并发机制。

## goroutine

goroutine 是由 Go 运行时环境管理的轻量级线程。

	go f(x, y, z)

开启一个新的 goroutine 执行

	f(x, y, z)

f ， x ， y 和 z 是当前 goroutine 中定义的，但是在新的 goroutine 中运行 `f`。

goroutine 在相同的地址空间中运行，因此访问共享内存必须进行同步。sync 提供了这种可能，不过在 Go 中并不经常用到，因为有其他的办法。（在接下来的内容中会涉及到。）

	package main

	import (
	    "fmt"
	    "time"
	)

	func say(s string) {
	    for i := 0; i < 5; i++ {
	        time.Sleep(100 * time.Millisecond)
	        fmt.Println(s)
	    }
	}

	func main() {
	    go say("world")
	    say("hello")
	}

	package main


## channel

channel 是有类型的管道，可以用 channel 操作符 <- 对其发送或者接收值。

	ch <- v    // 将 v 送入 channel ch。
	v := <-ch  // 从 ch 接收，并且赋值给 v。

（“箭头”就是数据流的方向。）

和 map 与 slice 一样，channel 使用前必须创建：

	ch := make(chan int)

默认情况下，在另一端准备好之前，发送和接收都会阻塞。这使得 goroutine 可以在没有明确的锁或竞态变量的情况下进行同步。


	import "fmt"

	func sum(a []int, c chan int) {
	    sum := 0
	    for _, v := range a {
	        sum += v
	    }
	    c <- sum // 将和送入 c
	}

	func main() {
	    a := []int{7, 2, 8, -9, 4, 0}

	    c := make(chan int)
	    go sum(a[:len(a)/2], c)
	    go sum(a[len(a)/2:], c)
	    x, y := <-c, <-c // 从 c 中获取

	    fmt.Println(x, y, x+y)
	}

## 缓冲 channel

channel 可以是 _带缓冲的_。为 make 提供第二个参数作为缓冲长度来初始化一个缓冲 channel：

	ch := make(chan int, 100)

向缓冲 channel 发送数据的时候，只有在缓冲区满的时候才会阻塞。当缓冲区清空的时候接受阻塞。

修改例子使得缓冲区被填满，然后看看会发生什么。

	package main

	import "fmt"

	func main() {
	    c := make(chan int, 2)
	    c <- 1
	    c <- 2
	    fmt.Println(<-c)
	    fmt.Println(<-c)
	}

## range 和 close

发送者可以 close 一个 channel 来表示再没有值会被发送了。接收者可以通过赋值语句的第二参数来测试 channel 是否被关闭：当没有值可以接收并且 channel 已经被关闭，那么经过

	v, ok := <-ch

之后 ok 会被设置为 `false`。

循环 `for i := range c` 会不断从 channel 接收值，直到它被关闭。

**注意**： 只有发送者才能关闭 channel，而不是接收者。向一个已经关闭的 channel 发送数据会引起 panic。

 **还要注意**： channel 与文件不同；通常情况下无需关闭它们。只有在需要告诉接收者没有更多的数据的时候才有必要进行关闭，例如中断一个 `range`。

	package main

	import (
	    "fmt"
	)

	func fibonacci(n int, c chan int) {
	    x, y := 0, 1
	    for i := 0; i < n; i++ {
	        c <- x
	        x, y = y, x+y
	    }
	    close(c)
	}

	func main() {
	    c := make(chan int, 10)
	    go fibonacci(cap(c), c)
	    for i := range c {
	        fmt.Println(i)
	    }
	}

## select

select 语句使得一个 goroutine 在多个通讯操作上等待。

select 会阻塞，直到条件分支中的某个可以继续执行，这时就会执行那个条件分支。当多个都准备好的时候，会随机选择一个。

## 默认选择

当 select 中的其他条件分支都没有准备好的时候，`default` 分支会被执行。

为了非阻塞的发送或者接收，可使用 default 分支：

	select {
	case i := <-c:
	    // 使用 i
	default:
		// 从 c 读取会阻塞
	}

demo:

	package main

	import (
	    "fmt"
	    "time"
	)

	func main() {
	    tick := time.Tick(100 * time.Millisecond)
	    boom := time.After(500 * time.Millisecond)
	    for {
	        select {
	        case <-tick:
	            fmt.Println("tick.")
	        case <-boom:
	            fmt.Println("BOOM!")
	            return
	        default:
	            fmt.Println("    .")
	            time.Sleep(50 * time.Millisecond)
	        }
	    }
	}

## 练习：等价二叉树

可以用多种不同的二叉树的叶子节点存储相同的数列值。例如，这里有两个二叉树保存了序列 1，1，2，3，5，8，13。

![tree](img/tree.png)

用于检查两个二叉树是否存储了相同的序列的函数在多数语言中都是相当复杂的。这里将使用 Go 的并发和 channel 来编写一个简单的解法。

这个例子使用了 tree 包，定义了类型：

	type Tree struct {
	    Left  *Tree
	    Value int
	    Right *Tree
	}


## 练习：等价二叉树

1. 实现 Walk 函数。

2. 测试 Walk 函数。

函数 tree.New(k) 构造了一个随机结构的二叉树，保存了值 `k`，`2k`，`3k`，...，`10k`。 创建一个新的 channel ch 并且对其进行步进：

	go Walk(tree.New(1), ch)

然后从 channel 中读取并且打印 10 个值。应当是值 1，2，3，...，10。

3. 用 Walk 实现 Same 函数来检测是否 t1 和 t2 存储了相同的值。

4. 测试 Same 函数。

`Same(tree.New(1), tree.New(1))` 应当返回 true，而 `Same(tree.New(1), tree.New(2))` 应当返回 false。

	package main

	import "code.google.com/p/go-tour/tree"

	// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
	func Walk(t *tree.Tree, ch chan int)

	// Same 检测树 t1 和 t2 是否含有相同的值。
	func Same(t1, t2 *tree.Tree) bool

	func main() {
	}


## 练习：Web 爬虫

在这个练习中，将会使用 Go 的并发特性来并行执行 web 爬虫。

修改 Crawl 函数来并行的抓取 URLs，并且保证不重复。

	package main

	import (
	    "fmt"
	)

	type Fetcher interface {
	    // Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	    Fetch(url string) (body string, urls []string, err error)
	}

	// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
	func Crawl(url string, depth int, fetcher Fetcher) {
	    // TODO: 并行的抓取 URL。
	    // TODO: 不重复抓取页面。
	        // 下面并没有实现上面两种情况：
	    if depth <= 0 {
	        return
	    }
	    body, urls, err := fetcher.Fetch(url)
	    if err != nil {
	        fmt.Println(err)
	        return
	    }
	    fmt.Printf("found: %s %q\n", url, body)
	    for _, u := range urls {
	        Crawl(u, depth-1, fetcher)
	    }
	    return
	}

	func main() {
	    Crawl("http://golang.org/", 4, fetcher)
	}

	// fakeFetcher 是返回若干结果的 Fetcher。
	type fakeFetcher map[string]*fakeResult

	type fakeResult struct {
	    body string
	    urls []string
	}

	func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	    if res, ok := f[url]; ok {
	        return res.body, res.urls, nil
	    }
	    return "", nil, fmt.Errorf("not found: %s", url)
	}

	// fetcher 是填充后的 fakeFetcher。
	var fetcher = fakeFetcher{
	    "http://golang.org/": &fakeResult{
	        "The Go Programming Language",
	        []string{
	            "http://golang.org/pkg/",
	            "http://golang.org/cmd/",
	        },
	    },
	    "http://golang.org/pkg/": &fakeResult{
	        "Packages",
	        []string{
	            "http://golang.org/",
	            "http://golang.org/cmd/",
	            "http://golang.org/pkg/fmt/",
	            "http://golang.org/pkg/os/",
	        },
	    },
	    "http://golang.org/pkg/fmt/": &fakeResult{
	        "Package fmt",
	        []string{
	            "http://golang.org/",
	            "http://golang.org/pkg/",
	        },
	    },
	    "http://golang.org/pkg/os/": &fakeResult{
	        "Package os",
	        []string{
	            "http://golang.org/",
	            "http://golang.org/pkg/",
	        },
	    },
	}

# golang编码规范

by **golang.org**

## gofmt

大部分的格式问题可以通过gofmt解决，gofmt自动格式化代码，保证所有的go代码一致的格式。

正常情况下，采用Sublime编写go代码时，插件GoSublilme已经调用gofmt对代码实现了格式化。

## 注释

在编码阶段同步写好变量、函数、包注释，注释可以通过godoc导出生成文档。

注释必须是完整的句子，以需要注释的内容作为开头，句点作为结尾。

程序中每一个被导出的（大写的）名字，都应该有一个文档注释。

**包注释**

每个程序包都应该有一个包注释，一个位于package子句之前的块注释或行注释。

包如果有多个go文件，只需要出现在一个go文件中即可。

	//Package regexp implements a simple library 
	//for regular expressions.
	package regexp

**可导出类型**

第一条语句应该为一条概括语句，并且使用被声明的名字作为开头。

	// Compile parses a regular expression and returns, if successful, a Regexp
	// object that can be used to match against text.
	func Compile(str string) (regexp *Regexp, err error) {

## 命名

使用短命名，长名字并不会自动使得事物更易读，文档注释会比格外长的名字更有用。

**包名**

包名应该为小写单词，不要使用下划线或者混合大小写。

**接口名**

单个函数的接口名以"er"作为后缀，如Reader,Writer

接口的实现则去掉“er”

	type Reader interface {
	    Read(p []byte) (n int, err error)
	}

两个函数的接口名综合两个函数名

	type WriteFlusher interface {
	    Write([]byte) (int, error)
	    Flush() error
	}

三个以上函数的接口名，类似于结构体名

	type Car interface {
	    Start([]byte) 
	    Stop() error
	    Recover()
	}

**混合大小写**

采用驼峰式命名

	MixedCaps 大写开头，可导出
	mixedCaps 小写开头，不可导出

## 控制结构

**if**

if接受初始化语句，约定如下方式建立局部变量

	if err := file.Chmod(0664); err != nil {
	    return err
	}

**for**

采用短声明建立局部变量

	sum := 0
	for i := 0; i < 10; i++ {
	    sum += i
	}

**range**

如果只需要第一项（key），就丢弃第二个：

	for key := range m {
	    if key.expired() {
	        delete(m, key)
	    }
	}

如果只需要第二项，则把第一项置为下划线

	sum := 0
	for _, value := range array {
	    sum += value
	}

**return**

尽早return：一旦有错误发生，马上返回

	f, err := os.Open(name)
	if err != nil {
	    return err
	}
	d, err := f.Stat()
	if err != nil {
	    f.Close()
	    return err
	}
	codeUsing(f, d)

## 函数（必须）

函数采用命名的多值返回

传入变量和返回变量以小写字母开头

	func nextInt(b []byte, pos int) (value, nextPos int) {

> 在godoc生成的文档中，带有返回值的函数声明更利于理解

## 错误处理

* error作为函数的值返回,必须对error进行处理
* 错误描述如果是英文必须为小写，不需要标点结尾
* 采用独立的错误流进行处理

不要采用这种方式

	if err != nil {
        // error handling
    } else {
        // normal code
    }

而要采用下面的方式

	if err != nil {
        // error handling
        return // or continue, etc.
    }
    // normal code

如果返回值需要初始化，则采用下面的方式

	x, err := f()
	if err != nil {
	    // error handling
	    return
	}
	// use x

## panic

尽量不要使用panic，除非你知道你在做什么

## import

对import的包进行分组管理，而且标准库作为第一组

	package main

	import (
	    "fmt"
	    "hash/adler32"
	    "os"

	    "appengine/user"
	    "appengine/foo"

	    "code.google.com/p/x/y"
	    "github.com/foo/bar"
	)

[goimports](https://godoc.org/code.google.com/p/go.tools/cmd/goimports) 实现了自动格式化

## 缩写

采用全部大写或者全部小写来表示缩写单词

比如对于url这个单词，不要使用

	UrlPony

而要使用

	urlPony 或者 URLPony

## 参数传递

* 对于少量数据，不要传递指针
* 对于大量数据的struct可以考虑使用指针
* 传入参数是map，slice，chan不要传递指针

因为map，slice，chan是引用类型，不需要传递指针的指针

## 接受者

**名称**

统一采用单字母'p'而不是this，me或者self

	type T struct{}

	func (p *T)Get(){}

**类型**

对于go初学者，接受者的类型如果不清楚，统一采用指针型

	func (p *T)Get(){}

而不是

	func (p T)Get(){}

在某些情况下，出于性能的考虑，或者类型本来就是引用类型，有一些特例

如果接收者是map,slice或者chan，不要用指针传递

**Map**

	//Map
	package main

	import (
	    "fmt"
	)

	type mp map[string]string

	func (m mp) Set(k, v string) {
	    m[k] = v
	}

	func main() {
	    m := make(mp)
	    m.Set("k", "v")
	    fmt.Println(m)
	}

**Channel**

	//Channel
	package main

	import (
	    "fmt"
	)

	type ch chan interface{}

	func (c ch) Push(i interface{}) {
	    c <- i
	}

	func (c ch) Pop() interface{} {
	    return <-c
	}

	func main() {
	    c := make(ch, 1)
	    c.Push("i")
	    fmt.Println(c.Pop())
	}

如果需要对slice进行修改，通过返回值的方式重新赋值

**Slice**

	//Slice
	package main

	import (
	    "fmt"
	)

	type slice []byte

	func main() {
	    s := make(slice, 0)
	    s = s.addOne(42)
	    fmt.Println(s)
	}

	func (s slice) addOne(b byte) []byte {
	    return append(s, b)
	}

如果接收者是含有sync.Mutex或者类似同步字段的结构体，必须使用指针传递避免复制

	package main

	import (
	    "sync"
	)

	type T struct {
	    m sync.Mutex
	}

	func (t *T) lock() {
	    t.m.Lock()
	}

	/*
	Wrong !!!
	func (t T) lock() {
	    t.m.Lock()
	}
	*/

	func main() {
	    t := new(T)
	    t.lock()
	}

如果接收者是大的结构体或者数组，使用指针传递会更有效率。

	package main

	import (
	    "fmt"
	)

	type T struct {
	    data [1024]byte
	}

	func (t *T) Get() byte {
	    return t.data[0]
	}

	func main() {
	    t := new(T)
	    fmt.Println(t.Get())
	}


