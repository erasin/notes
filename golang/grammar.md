#Golang语法与代码格式速记

    // Description: Golang语法与代码格式速记
    // Author: cxy
    // Date: 2013-04-01
    // Version: 0.3
    // TODO 说明


    // TODO package

    // Go是采用语法解析器自动在每行末尾增加分号，所以在写代码的时候可以把分号省略。
    // Go编程中只有几个地方需要手工增加分号，如: for循环使用分号把初始化，条件和遍历元素分开。在一行中有多条语句时，需要增加分号。
    // 不能把控制语句(if, for, switch, or select)、函数、方法 的左大括号单独放在一行， 如果你这样作了语法解析器会在大括号之前插入一个分号，导致编译错误。

    // 引用包名与导入路径的最后一个目录一致
    import "fmt"
    import "math/rand"
    fmt.Println(rand.Intn(10))    // 0到10之间的非负伪随机数

    // 用圆括号组合导入包，这是“factored”导入语句
    import ("fmt"; "math")
    import (
        "fmt"
        "math"
    )
    // 导入包可以定义别名，防止同名称的包冲突
    import n "net/http"
    import (
        控制台 "fmt"
        m "math"
    )
    控制台.Println(m.Pi)

    // 首字母大写的名称是被导出的, 首字母小写的名称只能在同一包内访问(同包跨文件也能访问)
    var In int // In is public
    var in byte // in is private
    var 看不见 string // 看不见 is private
    const Com bool = false // Com is public
    const 还是看不见 uint8 = 1 // 还是看不见 is private
    type Integer int // Integer is public
    type ブーリアン *bool // ブーリアン is private
    func Export() {} // Export is public
    func 导入() {} // 导入 is private
    func (me *Integer) valueOf(s string) int {} // valueOf is private
    func (i ブーリアン) String() string {} // String is public

    // Go 的基本类型：
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
    // byte 是 uint8 的别名
    // rune 是 int32 的别名，代表一个Unicode码点

    // 变量声明, 使用var关键字    (Go中只能使用var声明变量，无需显式初始化值）
    var i int    // i = 0
    var s string    // s = ""    (Go中的string不存在nil(null)值，默认零值就是空串 "" 或 ``)
    var e error    // e = nil, error是Go的内建接口类型，不是基本类型。

    // var 语句声明了一个变量的列表，类型在变量名之后
    var a,b,c int    // a = 0, b = 0, c = 0
    var (
        a int    // a = 0
        b string    // b = ""
        c uint    // c = 0
    )

    // 变量定义时初始化赋值，每个变量对应一个值
    var a int = 0
    var a,b int = 0, 1

    // 初始化使用表达式时，可以省略类型，变量从初始值中获得类型
    var a = 'A'    // a int32
    c := 1 + 2i    // c complex128
    var a,b = 0, "B"    // a int, b string
    a, b := 0, "B"    // a int, b string
    c := `formatted
     string`    // c string

    // := 结构不能使用在函数外，函数外的每个语法块都必须以关键字开始

    // 常量可以是字符、字符串、布尔或数字类型的值，数值常量是高精度的值
    const x int = 3
    const (
        a byte = 'A'
        b string = "B"
        c bool = true
        d int = 4
        e float32 = 5.1
        f complex64 = 6 + 6i
    )

    // 未指定类型的常量由常量值决定其类型
    const a = 0    // a int
    const (
        b = 2.3    // b float64
        c = true    // c bool
    )

    // 自动枚举常量 iota
    // iota的枚举值可以赋值给数值兼容类型
    // 每个常量单独声明时, iota不会自动递增(无意义)
    const a int = iota    // a = 0
    const b int = iota    // b = 0
    const c byte = iota    // c = 0
    const d uint64 = iota    // d = 0

    // 常量组合声明时, iota每次引用会逐步自增, 初始值为0,步进值为1
    const (
        a uint8 = iota    // a = 0
        b int16 = iota    // b = 1
        c rune = iota    // c = 2
        d float64 = iota    // d = 3
        e uintptr = iota    // e = 4
    )

    // 枚举的常量都为同一类型时, 可以使用简单序列格式.
    const (
        a = iota    // a int32 = 0
        b            // b int32 = 1
        c            // c int32 = 2
    )

    // 枚举序列中的未指定类型的常量会跟随序列前面最后一次出现类型定义的类型
    const (
        a byte = iota    // a uint8 = 0
        b                // b uint8 = 1
        c                // c uint8 = 2
        d rune = iota    // d int32 = 3
        e                // e int32 = 4
        f                // f int32 = 5
    )

    // iota自增值只在一个常量定义组合中有效,跳出常量组合定义后iota值归0
    const (
        a = iota    // a int32 = 0
        b            // b int32 = 1
        c            // c int32 = 2
    )
    const (
        e = iota    // e int32 = 0    (iota重新初始化并自增)
        f            // f int32 = 1
    )

    // 定制iota序列初始值与步进值 (通过数学公式实现)
    const (
        a = (iota + 2) * 3    // a int32 = 0    (a=(0+2)*3) 初始值为6,步进值为3
        b                    // b int32 = 3    (b=(1+2)*3)
        c                    // c int32 = 6    (c=(2+2)*3)
        d                    // d int32 = 9    (d=(3+2)*3)
    )

    // 数组声明带有长度信息，数组的长度固定
    var a [3]int = [3]int{0, 1, 2}    // a = [0 1 2]
    var b [3]int = [3]int{}    // b = [0 0 0]
    var c = [3]int{}    // c = [0 0 0]
    d := [3]int{}    // d = [0 0 0]
    fmt.Printf("%T\t%#v\t%d\t%d\n", d, d, len(d), cap(d))    // [3]int    [3]int{0, 0, 0}    3    3
    // 使用自动计算数组初始数据的长度
    var a = []int{0, 1, 2}
    x := [][3]int{{0, 1, 2}, {3, 4, 5}}

    // slice 指向数组的值，并且同时包含了长度信息
    var a []int
    fmt.Printf("%T\t%#v\t%d\t%d\n", a, a, len(a), cap(a))    // []int    []int(nil)    0    0
    var a = new([]int)
    fmt.Printf("%T\t%#v\t%d\t%d\n", a, a, len(*a), cap(*a))    // *[]int    &[]int(nil)    0    0
    var b = make([]int, 0)
    fmt.Printf("%T\t%#v\t%d\t%d\n", b, b, len(b), cap(b))    // []int    []int{}    0    0
    var c = make([]int, 3, 10)
    fmt.Printf("%T\t%#v\t%d\t%d\n", c, c, len(c), cap(c))    // []int    []int{}    3    10
    var d []int = []int{0, 1, 2}
    fmt.Printf("%T\t%#v\t%d\t%d\n", d, d, len(d), cap(d))    // []int    []int{0, 1, 2}    3    3

    // slice 可以重新切片，创建一个新的 slice 值指向相同的数组
    s := []int{0, 1, 2, 3, 4}
    fmt.Println(s[1,3])    // [1 2]    (截取从开始索引到结束索引-1 之间的片段)
    fmt.Println(s[:4])    // [0 1 2 3]
    fmt.Println(s[1:])    // [1 2 3 4]
    fmt.Println(s[1:1])    // []

    // 向slice中添加元素
    s := make([]string, 3)
    s = append(s, "a")


    // map 在使用之前必须用 make 来创建（不是 new）；一个值为 nil 的 map 是空的，并且不能赋值
    var m map[int]int
    m[0] = 0    // × runtime error: assignment to entry in nil map
    fmt.Printf("type: %T\n", m)    // map[int]int
    fmt.Printf("value: %#v\n", m)    // map[int]int(nil)
    fmt.Printf("value: %v\n", m)    // map[]
    fmt.Println("is nil: ", nil == m)    // true
    fmt.Println("length: ", len(m))    // 0，if m is nil, len(m) is zero.

    var m map[int]int = make(map[int]int)
    m[0] = 0    // 插入或修改元素
    fmt.Printf("type: %T\n", m)        // map[int]int
    fmt.Printf("value: %#v\n", m)        // map[int]int(0:0)
    fmt.Printf("value: %v\n", m)        // map[0:0]
    fmt.Println("is nil: ", nil == m)    // false
    fmt.Println("length: ", len(m))        // 1

    m = map[int]int{
    0:0,
    1:1,    // 最后的逗号是必须的
    }
    m = map[string]S{
    "a":S{0,1},
    "b":{2,3},    // 类型名称可省略
    }
    a := m["a"]    // 取值
    a, ok := m["a"]    // 取值, 并通过ok(bool)判断key对应的元素是否存在.
    delete(m, "a")    // 删除key对应的元素.

    // 结构体（struct）就是一个字段的集合， type 定义跟其字面意思相符
    type S struct {
        A int
        B, c string
    }
    type (
        A struct {
            s *S
        }
        B struct {
            A    // 组合
        }
    )
    // 结构体文法表示通过结构体字段的值作为列表来新分配一个结构体。
    var s S = S{0, "1", "2"}
    // 使用 Name: 语法可以仅列出部分字段。（字段名的顺序无关。）
    var s S = S{A: 0, B: "1"}
    var s S = S{}
    // 特殊的前缀 & 构造了指向结构体文法的指针。
    var s *S = &S{0, "1", "2"}

    // 表达式 new(T) 分配了一个零初始化的 T 值，并返回指向它的指针
    var s *S = new(S)
    // 有指针，但是没有指针运算，结构体字段使用点号来访问
    // 结构体字段可以通过结构体指针来访问。通过指针间接的访问是透明的
    fmt.Println(s.A)
    fmt.Println((*s).A)

    // TODO interface
    type IF interface {
        a()
    }

    // TODO chanel

    // TODO error

    // if 语句 小括号 ( )是可选的，而 { } 是必须的。
    if (i < 0)        // 编译错误.
        println(i)

    if i < 0        // 编译错误.
        println(i)

    if (i < 0) {    // 编译通过.
        println(i)
    }
    if i < 0 {
        println(i)
    } else {
        println(i)
    }

    // 可以在条件之前执行一个简单的语句，由这个语句定义的变量的作用域仅在 if/else 范围之内
    if (i := 0; i < 1) {    // 编译错误.
        println(i)
    }

    if i := 0; (i < 1) {    // 编译通过.
        println(i)
    }

    if i := 0; i < 0 {    // 使用gofmt格式化代码会自动移除代码中不必要的小括号( )
        println(i)
    } else if i == 0 {
        println(i)
    } else {
        println(i)
    }

    // if语句作用域范围内定义的变量会覆盖外部同名变量,(与方法函数内局部变量覆盖全局变量相同)
    a, b := 0, 1
    if a, b := 3, 4; a > 1 && b > 2 {
        println(a, b)    // 3 4
    }
    println(a, b)    // 0 1


    // 只有一种循环结构，for 循环。可以让前置、后置语句为空，或者全为空
    for i := 0; i < 10; i++ {}
    for i := 0; i < 10; {}
    for ; i < 10; i++ {}
    for ; i < 10; {}
    for i < 10 {}
    for ; ; {}
    for {}

    // 小括号 ( )是可选的，而 { } 是必须的。
    for (i := 0; i < 10; i++) {}    // 编译错误.
    for i := 0; (i < 10); i++ {}    // 编译通过.
    for (i < 10) {}    // 编译通过.

    // TODO continue

    // TODO for range

    // TODO switch
    // TODO fallthrough break
    // TODO type assertion

    // TODO select

    // TODO goto

    // 函数可以没有参数或接受多个参数
    func f() {}
    func f(a int) {}
    func f(a int, b byte) {}
    func f(a int) {}    // 可变参数
    func f(a int, b bool, c string) {}
    // 函数可以返回任意数量的返回值
    func f() int {
        return 0
    }
    func f() int, string {
        return 0, "A"
    }
    // 函数返回结果参数，可以像变量那样命名和使用
    func f() a int, b string {
        a = 1
        b = "B"
        return    // 或者 return a, b
    }

    // 当两个或多个连续的函数命名参数是同一类型，则除了最后一个类型之外，其他都可以省略
    func f(a,b,c int) {}
    func f() a,b,c int {}
    func f(a,b,c int) x,y,z int {}

    // 函数也是值，可以将函数赋值给变量
    var f (func(i int) int) = func(i int) int {
        return i
    }
    fmt.Println(f(3))    // 3
    var f func() int = func() int {
        return 0
    }
    fmt.Println(f())    // 0
    var f func() = func() {}
    var f = func() {}
    f := func() {}

    // TODO defer

    // TODO 方法

    // TODO 内建函数
    append 
    cap 
    close 
    complex 
    copy 
    delete 
    imag 
    len 
    make 
    new 
    panic 
    print 
    println 
    real 
    recover

    // TODO 并发
    go func() {}