# go 简介

[来源](http://coolshell.cn/articles/8460.html)

周末天气不好，只能宅在家里，于是就顺便看了一下Go语言，觉得比较有意思，所以写篇文章介绍一下。**我想写一篇你可以在乘坐地铁或公交车上下班时就可以初步了解一门语言的文章**。所以，下面的文章主要是以代码和注释为主。只需要你对C语言，Unix，Python有一点基础，我相信你会在30分钟左右读完并对Go语言有一些初步了解的。

![go](http://coolshell.cn//wp-content/uploads/2012/11/go2.jpg)

## Hello World

文件名 hello.go

``` java
package main //声明本文件的package名

import &quot;fmt&quot; //import语言的fmt库——用于输出

func main() {
    fmt.Println(&quot;hello world&quot;)
}
```

## 运行

你可以有两种运行方式，

解释执行（实际是编译成a.out再执行）
``` bash
$go run hello.go
hello world
```

编译执行
``` bash
$go build hello.go

$ls
hello hello.go

$./hello
hello world
```

## 自己的package

你可以使用GOPATH环境变量，或是使用相对路径来import你自己的package。

Go的规约是这样的：

* **在import中，你可以使用相对路径，如 ./或 ../ 来引用你的package**
* **如果没有使用相对路径，那么，go会去找$GOPATH/src/目录。**

使用相对路径

``` java
import &quot;./haoel&quot;  //import当前目录里haoel子目录里的所有的go文件
```

使用GOPATH路径
``` java
import &quot;haoel&quot;  //import 环境变量 $GOPATH/src/haoel子目录里的所有的go文件
```

## fmt输出格式

fmt包和libc里的那堆使用printf， scanf，fprintf，fscanf 很相似。下面的东西对于C程序员不会陌生。

注意：Println不支持，Printf才支持%式的输出：

``` java
package main

import &quot;fmt&quot;
import &quot;math&quot;

func main() {
    fmt.Println(&quot;hello world&quot;)

    fmt.Printf(&quot;%t\n&quot;, 1==2)
    fmt.Printf(&quot;二进制：%b\n&quot;, 255)
    fmt.Printf(&quot;八进制：%o\n&quot;, 255)
    fmt.Printf(&quot;十六进制：%X\n&quot;, 255)
    fmt.Printf(&quot;十进制：%d\n&quot;, 255)
    fmt.Printf(&quot;浮点数：%f\n&quot;, math.Pi)
    fmt.Printf(&quot;字符串：%s\n&quot;, &quot;hello world&quot;)
}
```

当然，也可以使用如\n\t\r这样的和C语言一样的控制字符
## 变量和常量

变量的声明很像 javascript，使用 var关键字。注意：**go是静态类型的语言**，下面是代码：

``` jscript
//声明初始化一个变量
var  x int = 100
var str string = &quot;hello world&quot;&lt;/pre&gt;
//声明初始化多个变量
var  i, j, k int = 1, 2, 3

//不用指明类型，通过初始化值来推导
var b = true //bool型
```

还有一种定义变量的方式（这让我想到了Pascal语言，但完全不一样）

``` java
x := 100 //等价于 var x int = 100;
```

常量很简单，使用const关键字：

``` java
const s string = &quot;hello world&quot;
const pi float32 = 3.1415926
```

## 数组

直接看代码（注意其中的for语句，和C很相似吧，就是没有括号了）

``` java
func main() {
    var a [5]int
    fmt.Println(&quot;array a:&quot;, a)

    a[1] = 10
    a[3] = 30
    fmt.Println(&quot;assign:&quot;, a)

    fmt.Println(&quot;len:&quot;, len(a))

    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println(&quot;init:&quot;, b)

    var c [2][3]int
    for i := 0; i &lt; 2; i++ {
        for j := 0; j &lt; 3; j++ {
            c[i][j] = i + j
        }
    }
    fmt.Println(&quot;2d: &quot;, c)
}
```

运行结果：

``` bash
array a: [0 0 0 0 0]
assign: [0 10 0 30 0]
len: 5
init: [1 2 3 4 5]
2d:  [[0 1 2] [1 2 3]]
```

## 数组的切片操作

这个很Python了。

``` java
a := [5]int{1, 2, 3, 4, 5}

b := a[2:4] // a[2] 和 a[3]，但不包括a[4]
fmt.Println(b)

b = a[:4] // 从 a[0]到a[4]，但不包括a[4]
fmt.Println(b)

b = a[2:] // 从 a[2]到a[4]，且包括a[2]
fmt.Println(b)
```

## 分支循环语句

**if语句**

注意：if 语句没有圆括号，而必需要有花括号

``` java
//if 语句
if x % 2 == 0 {
    //...
}
//if - else
if x % 2 == 0 {
    //偶数...
} else {
    //奇数...
}

//多分支
if num &lt; 0 {
    //负数
} else if num == 0 {
    //零
} else {
    //正数
}
```

**switch 语句**

注意：switch语句没有break，还可以使用逗号case多个值

``` cpp
switch i {
    case 1:
        fmt.Println(&quot;one&quot;)
    case 2:
        fmt.Println(&quot;two&quot;)
    case 3:
        fmt.Println(&quot;three&quot;)
    case 4,5,6:
        fmt.Println(&quot;four, five, six&quot;)
    default:
        fmt.Println(&quot;invalid value!&quot;)
}
```

**for 语句**

前面你已见过了，下面再来看看for的三种形式：（注意：Go语言中没有while）

``` cpp
//经典的for语句 init; condition; post
for i := 0; i&lt;10; i++{
     fmt.Println(i)
}

//精简的for语句 condition
i := 1
for i&lt;10 {
    fmt.Println(i)
    i++
}

//死循环的for语句 相当于for(;;)
i :=1
for {
    if i&gt;10 {
        break
    }
    i++
}
```

## 关于分号

从上面的代码我们可以看到代码里没有分号。其实，**和C一样，Go的正式的语法使用分号来终止语句。和C不同的是，这些分号由词法分析器在扫描源代码过程中使用简单的规则自动插入分号，因此输入源代码多数时候就不需要分号了**。

规则是这样的：如果在一个新行前方的最后一个标记是一个标识符（包括像`int`和`float64`这样的单词）、一个基本的如数值这样的文字、或以下标记中的一个时，会自动插入分号：

```
break continue fallthrough return ++ -- ) }
```

通常Go程序仅在`for`循环语句中使用分号，以此来分开初始化器、条件和增量单元。如果你在一行中写多个语句，也需要用分号分开。

**注意**：**无论任何时候，你都不应该将一个控制结构（(`if`、`for`、`switch`或`select`）的左大括号放在下一行。如果这样做，将会在大括号的前方插入一个分号，这可能导致出现不想要的结果**。
## map

map在别的语言里可能叫哈希表或叫dict，下面是和map的相关操作的代码，代码很容易懂

``` cpp
func main(){
    m := make(map[string]int) //使用make创建一个空的map

    m[&quot;one&quot;] = 1
    m[&quot;two&quot;] = 2
    m[&quot;three&quot;] = 3

    fmt.Println(m) //输出 map[three:3 two:2 one:1] (顺序在运行时可能不一样)
    fmt.Println(len(m)) //输出 3

    v := m[&quot;two&quot;] //从map里取值
    fmt.Println(v) // 输出 2

    delete(m, &quot;two&quot;)
    fmt.Println(m) //输出 map[three:3 one:1]

    m1 := map[string]int{&quot;one&quot;: 1, &quot;two&quot;: 2, &quot;three&quot;: 3}
    fmt.Println(m1) //输出 map[two:2 three:3 one:1] (顺序在运行时可能不一样)

    for key, val := range m1{
        fmt.Printf(&quot;%s =&gt; %d \n&quot;, key, val)
        /*输出：(顺序在运行时可能不一样)
            three =&gt; 3
            one =&gt; 1
            two =&gt; 2*/
    }
}
```

## 指针

Go语言一样有指针，看代码

``` cpp
var i int = 1
var pInt *int = &amp;i
//输出：i=1     pInt=0xf8400371b0       *pInt=1
fmt.Printf(&quot;i=%d\tpInt=%p\t*pInt=%d\n&quot;, i, pInt, *pInt)

*pInt = 2
//输出：i=2     pInt=0xf8400371b0       *pInt=2
fmt.Printf(&quot;i=%d\tpInt=%p\t*pInt=%d\n&quot;, i, pInt, *pInt)

i = 3
//输出：i=3     pInt=0xf8400371b0       *pInt=3
fmt.Printf(&quot;i=%d\tpInt=%p\t*pInt=%d\n&quot;, i, pInt, *pInt)
```

Go具有两个分配内存的机制，分别是内建的函数new和make。他们所做的事不同，所应用到的类型也不同，这可能引起混淆，但规则却很简单。
## **内存分配 **

**new** 是一个分配内存的内建函数，但不同于其他语言中同名的new所作的工作，**它只是将内存清零，而不是初始化内存**。new(T)为一个类型为T的新项目分配了值为零的存储空间并返回其地址，也就是一个类型为*T的值。用Go的术语来说，就是**它返回了一个指向新分配的类型为T的零值的指针**。

`**make**(T, `<em>args</em>`)`函数的目的与`new(T)`不同。它仅用于创建切片、map和chan（消息管道），并返回类型`T`（不是`*T`）的一个**被初始化了的**（不是**零**）实例。这种差别的出现是由于这三种类型实质上是对在使用前必须进行初始化的数据结构的引用。例如，切片是一个具有三项内容的描述符，包括指向数据（在一个数组内部）的指针、长度以及容量，在这三项内容被初始化之前，切片值为`nil`。对于切片、映射和信道，`make`初始化了其内部的数据结构并准备了将要使用的值。如：

下面的代码分配了一个整型数组，长度为10，容量为100，并返回前10个数组的切片

``` cpp
make([]int, 10, 100)
```

以下示例说明了`new`和`make`的不同。

``` jscript
var p *[]int = new([]int)   // 为切片结构分配内存；*p == nil；很少使用
var v  []int = make([]int, 10) // 切片v现在是对一个新的有10个整数的数组的引用

// 不必要地使问题复杂化：
var p *[]int = new([]int)
fmt.Println(p) //输出：&amp;[]
*p = make([]int, 10, 10)
fmt.Println(p) //输出：&amp;[0 0 0 0 0 0 0 0 0 0]
fmt.Println((*p)[2]) //输出： 0

// 习惯用法:
v := make([]int, 10)
fmt.Println(v) //输出：[0 0 0 0 0 0 0 0 0 0]

```

## 函数

老实说，我对Go语言这种反过来声明变量类型和函数返回值的做法有点不满（保持和C一样的不可以吗? 呵呵）

``` java

package main
import &quot;fmt&quot;

func max(a int, b int) int { //注意参数和返回值是怎么声明的

    if a &gt; b {
        return a
    }
    return b
}

func main(){
    fmt.Println(max(4, 5))
}

```

**函数返回多个值**

Go中很多Package 都会返回两个值，一个是正常值，一个是错误，如下所示：

``` java

package main
import &quot;fmt&quot;

func main(){
    v, e := multi_ret(&quot;one&quot;)
    fmt.Println(v,e) //输出 1 true

    v, e = multi_ret(&quot;four&quot;)
    fmt.Println(v,e) //输出 0 false

    //通常的用法(注意分号后有e)
    if v, e = multi_ret(&quot;four&quot;); e {
    	// 正常返回
    }else{
    	// 出错返回
    }
}

func multi_ret(key string) (int, bool){
    m := map[string]int{&quot;one&quot;: 1, &quot;two&quot;: 2, &quot;three&quot;: 3}

    var err bool
    var val int

    val, err = m[key]

    return val, err
}

```

**函数不定参数**

例子很清楚了，我就不多说了

``` cpp

func sum(nums ...int) {
    fmt.Print(nums, &quot; &quot;)  //输出如 [1, 2, 3] 之类的数组
    total := 0
    for _, num := range nums { //要的是值而不是下标
        total += num
    }
    fmt.Println(total)
}
func main() {
    sum(1, 2)
    sum(1, 2, 3)

    //传数组
    nums := []int{1, 2, 3, 4}
    sum(nums...)
}
```

**函数闭包**

nextNum这个函数返回了一个匿名函数，这个匿名函数记住了nextNum中i+j的值，并改变了i,j的值，于是形成了一个闭包的用法

``` cpp
func nextNum() func() int {
    i,j := 1,1
    return func() int {
        var tmp = i+j
        i, j = j, tmp
        return tmp
    }
}
//main函数中是对nextNum的调用，其主要是打出下一个斐波拉契数
func main(){
    nextNumFunc := nextNum()
    for i:=0; i&lt;10; i++ {
    	fmt.Println(nextNumFunc())
    }
}

```

**函数的递归**

和c基本是一样的

``` cpp
func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}

func main() {
    fmt.Println(fact(7))
}
```

## 结构体

Go的结构体和C的基本上一样，不过在初始化时有些不一样，Go支持带名字的初始化。

``` cpp
type Person struct {
    name string
    age  int
    email string
}

func main() {
    //初始化
    person := Person{&quot;Tom&quot;, 30, &quot;tom@gmail.com&quot;}
    person = Person{name:&quot;Tom&quot;, age: 30, email:&quot;tom@gmail.com&quot;}

    fmt.Println(person) //输出 {Tom 30 tom@gmail.com}

    pPerson := &amp;person

    fmt.Println(pPerson) //输出 &amp;{Tom 30 tom@gmail.com}

    pPerson.age = 40
    person.name = &quot;Jerry&quot;
    fmt.Println(person) //输出 {Jerry 40 tom@gmail.com}
}
```

## 结构体方法

不多说了，看代码吧。

注意：Go语言中没有public, protected, private的关键字，所以，**如果你想让一个方法可以被别的包访问的话，你需要把这个方法的第一个字母大写。这是一种约定**。

``` cpp
type rect struct {
    width, height int
}

func (r *rect) area() int { //求面积
    return r.width * r.height
}

func (r *rect) perimeter() int{ //求周长
    return 2*(r.width + r.height)
}

func main() {
    r := rect{width: 10, height: 15}

    fmt.Println(&quot;面积: &quot;, r.area())
    fmt.Println(&quot;周长: &quot;, r.perimeter())

    rp := &amp;r
    fmt.Println(&quot;面积: &quot;, rp.area())
    fmt.Println(&quot;周长: &quot;, rp.perimeter())
}
```

## 接口和多态

接口意味着多态，下面是一个经典的例子，不用多说了，自己看代码吧。

``` cpp

//---------- 接 口 --------//
type shape interface {
	area() float64 //计算面积
	perimeter() float64 //计算周长
}

//--------- 长方形 ----------//
type rect struct {
    width, height float64
}

func (r *rect) area() float64 { //面积
	return r.width * r.height
}

func (r *rect) perimeter() float64 { //周长
	return 2*(r.width + r.height)
}

//----------- 圆  形 ----------//
type circle struct {
	radius float64
}

func (c *circle) area() float64 { //面积
	return math.Pi * c.radius * c.radius
}

func (c *circle) perimeter() float64 { //周长
	return 2 * math.Pi * c.radius
}

// ----------- 接口的使用 -----------//
func interface_test() {
    r := rect {width:2.9, height:4.8}
    c := circle {radius:4.3}

    s := []shape{&amp;r, &amp;c} //通过指针实现

    for _, sh := range s {
        fmt.Println(sh)
    	fmt.Println(sh.area())
    	fmt.Println(sh.perimeter())
    }
}

```

## 错误处理 &#8211; Error接口

函数错误返回可能是C/C++时最让人纠结的东西的，Go的多值返回可以让我们更容易的返回错误，其可以在返回一个常规的返回值之外，还能轻易地返回一个详细的错误描述。通常情况下，错误的类型是error，它有一个内建的接口。

``` cpp
type error interface {
    Error() string
}
```

还是看个示例吧：

``` java
package main

import &quot;fmt&quot;
import &quot;errors&quot;

//自定义的出错结构
type myError struct {
    arg  int
    errMsg string
}
//实现Error接口
func (e *myError) Error() string {
    return fmt.Sprintf(&quot;%d - %s&quot;, e.arg, e.errMsg)
}

//两种出错
func error_test(arg int) (int, error) {
    if arg &lt; 0  {
         return -1, errors.New(&quot;Bad Arguments - negtive!&quot;)
     }else if arg &gt;256 {
        return -1, &amp;myError{arg, &quot;Bad Arguments - too large!&quot;}
    }
    return arg*arg, nil
}

//相关的测试
func main() {
    for _, i := range []int{-1, 4, 1000} {
        if r, e := error_test(i); e != nil {
            fmt.Println(&quot;failed:&quot;, e)
        } else {
            fmt.Println(&quot;success:&quot;, r)
        }
    }
}
```

程序运行后输出：

``` bash

failed: Bad Arguments - negtive!
success: 16
failed: 1000 - Bad Arguments - too large!

```

## 错误处理 &#8211; Defer

下面的程序对于每一个熟悉C语言的人来说都不陌生（有资源泄露的问题），C++使用RAII来解决这种问题。

``` cpp
func CopyFile(dstName, srcName string) (written int64, err error) {
    src, err := os.Open(srcName)
    if err != nil {
        return
    }

    dst, err := os.Create(dstName)
    if err != nil {
        return
    }

    written, err = io.Copy(dst, src)
    dst.Close()
    src.Close()
    return
}
```

Go语言引入了Defer来确保那些被打开的文件能被关闭。如下所示：（这种解决方式还是比较优雅的）

``` cpp; highlight: [6,12]
func CopyFile(dstName, srcName string) (written int64, err error) {
    src, err := os.Open(srcName)
    if err != nil {
        return
    }
    defer src.Close()

    dst, err := os.Create(dstName)
    if err != nil {
        return
    }
    defer dst.Close()

    return io.Copy(dst, src)
}
```

Go的defer语句预设一个函数调用（延期的函数），该调用在函数执行defer返回时立刻运行。该方法显得不同常规，但却是处理上述情况很有效，无论函数怎样返回，都必须进行资源释放。

我们再来看一个defer函数的示例：

``` cpp
for i := 0; i &lt; 5; i++ {
    defer fmt.Printf(&quot;%d &quot;, i)
}
```

被延期的函数以后进先出（LIFO）的顺行执行，因此以上代码在返回时将打印4 3 2 1 0。

总之，我个人觉得defer的函数行为有点怪异，我现在还没有完全搞清楚。
## 错误处理 &#8211; Panic/Recover

对于不可恢复的错误，Go提供了一个内建的panic函数，它将创建一个运行时错误并使程序停止（相当暴力）。该函数接收一个任意类型（往往是字符串）作为程序死亡时要打印的东西。当编译器在函数的结尾处检查到一个panic时，就会停止进行常规的return语句检查。

下面的仅仅是一个示例。实际的库函数应避免panic。如果问题可以容忍，最好是让事情继续下去而不是终止整个程序。

``` cpp
var user = os.Getenv(&quot;USER&quot;)

func init() {
    if user == &quot;&quot; {
        panic(&quot;no value for $USER&quot;)
    }
}
```

当panic被调用时，它将立即停止当前函数的执行并开始逐级解开函数堆栈，同时运行所有被defer的函数。如果这种解开达到堆栈的顶端，程序就死亡了。但是，也可以使用内建的recover函数来重新获得Go程的控制权并恢复正常的执行。 对recover的调用会通知解开堆栈并返回传递到panic的参量。由于仅在解开期间运行的代码处在被defer的函数之内，recover仅在被延期的函数内部才是有用的。

你可以简单地理解为recover就是用来捕捉Painc的，防止程序一下子就挂掉了。

下面是一个例程，很简单了，不解释了

``` cpp
func g(i int) {
    if i&gt;1 {
        fmt.Println(&quot;Panic!&quot;)
        panic(fmt.Sprintf(&quot;%v&quot;, i))
    }

}

func f() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println(&quot;Recovered in f&quot;, r)
        }
    }()

    for i := 0; i &lt; 4; i++ {
        fmt.Println(&quot;Calling g with &quot;, i)
        g(i)
        fmt.Println(&quot;Returned normally from g.&quot;)
     }
}

func main() {
    f()
    fmt.Println(&quot;Returned normally from f.&quot;)
}
```

运行结果如下：（我们可以看到Painc后的for循环就没有往下执行了，但是main的程序还在往下走）

``` bash
Calling g with  0
Returned normally from g.
Calling g with  1
Returned normally from g.
Calling g with  2
Panic!
Recovered in f 2
Returned normally from f.

```

你习惯这种编程方式吗？我觉得有点诡异。呵呵。


希望你看到这篇文章的时候还是在公交车和地铁上正在上下班的时间，我希望我的这篇文章可以让你利用这段时间了解一门语言。当然，希望你不会因为看我的文章而错过站。呵呵。

如果你还不了解Go语言的语法，还请你移步先看一下上篇——《**<a title="Go语言简介（上）：语法" href="http://coolshell.cn/articles/8460.html" target="_blank">Go语言简介（上）：语法</a>**》

<img src="http://coolshell.cn//wp-content/uploads/2012/11/google-go-language.jpg" alt="" title="google-go-language" width="450" height="272" class="aligncenter size-full wp-image-8531" />
## goroutine

GoRoutine主要是使用go关键字来调用函数，你还可以使用匿名函数，如下所示：


``` java; highlight: [9,11]
package main
import &quot;fmt&quot;

func f(msg string) {
    fmt.Println(msg)
}

func main(){
    go f(&quot;goroutine&quot;)

    go func(msg string) {
        fmt.Println(msg)
    }(&quot;going&quot;)
}
```

我们再来看一个示例，下面的代码中包括很多内容，包括时间处理，随机数处理，还有goroutine的代码。如果你熟悉C语言，你应该会很容易理解下面的代码。

你可以简单的把go关键字调用的函数想像成pthread_create。下面的代码使用for循环创建了3个线程，每个线程使用一个随机的Sleep时间，然后在routine()函数中会输出一些线程执行的时间信息。

``` java

package main

import &quot;fmt&quot;
import &quot;time&quot;
import &quot;math/rand&quot;

func routine(name string, delay time.Duration) {

    t0 := time.Now()
    fmt.Println(name, &quot; start at &quot;, t0)

    time.Sleep(delay)

    t1 := time.Now()
    fmt.Println(name, &quot; end at &quot;, t1)

    fmt.Println(name, &quot; lasted &quot;, t1.Sub(t0))
}

func main() {

    //生成随机种子
    rand.Seed(time.Now().Unix())

    var name string
    for i:=0; i&lt;3; i++{
        name = fmt.Sprintf(&quot;go_%02d&quot;, i) //生成ID
        //生成随机等待时间，从0-4秒
        go routine(name, time.Duration(rand.Intn(5)) * time.Second)
    }

    //让主进程停住，不然主进程退了，goroutine也就退了
    var input string
    fmt.Scanln(&amp;input)
    fmt.Println(&quot;done&quot;)
}

```

运行的结果可能是：

``` bash

go_00  start at  2012-11-04 19:46:35.8974894 +0800 +0800
go_01  start at  2012-11-04 19:46:35.8974894 +0800 +0800
go_02  start at  2012-11-04 19:46:35.8974894 +0800 +0800
go_01  end at  2012-11-04 19:46:36.8975894 +0800 +0800
go_01  lasted  1.0001s
go_02  end at  2012-11-04 19:46:38.8987895 +0800 +0800
go_02  lasted  3.0013001s
go_00  end at  2012-11-04 19:46:39.8978894 +0800 +0800
go_00  lasted  4.0004s

```

## goroutine的并发安全性

关于goroutine，我试了一下，无论是Windows还是Linux，基本上来说是用操作系统的线程来实现的。不过，goroutine有个特性，也就是说，**如果一个goroutine没有被阻塞，那么别的goroutine就不会得到执行**。这并不是真正的并发，如果你要真正的并发，你需要在你的main函数的第一行加上下面的这段代码：

``` cpp
import &quot;runtime&quot;
...
runtime.GOMAXPROCS(4)
```

还是让我们来看一个有并发安全性问题的示例（注意：我使用了C的方式来写这段Go的程序）

这是一个经常出现在教科书里卖票的例子，我启了5个goroutine来卖票，卖票的函数sell_tickets很简单，就是随机的sleep一下，然后对全局变量total_tickets作减一操作。

``` java
package main

import &quot;fmt&quot;
import &quot;time&quot;
import &quot;math/rand&quot;
import &quot;runtime&quot;

var total_tickets int32 = 10;

func sell_tickets(i int){
    for{
        if total_tickets &gt; 0 { //如果有票就卖
            time.Sleep( time.Duration(rand.Intn(5)) * time.Millisecond)
            total_tickets-- //卖一张票
            fmt.Println(&quot;id:&quot;, i, &quot;  ticket:&quot;, total_tickets)
        }else{
            break
        }
    }
}

func main() {
    runtime.GOMAXPROCS(4) //我的电脑是4核处理器，所以我设置了4
    rand.Seed(time.Now().Unix()) //生成随机种子

    for i := 0; i &lt; 5; i++ { //并发5个goroutine来卖票
         go sell_tickets(i)
    }
    //等待线程执行完
    var input string
    fmt.Scanln(&amp;input)
    fmt.Println(total_tickets, &quot;done&quot;) //退出时打印还有多少票
}
```

这个程序毋庸置疑有并发安全性问题，所以执行起来你会看到下面的结果：

``` bash
$go run sell_tickets.go
id: 0   ticket: 9  
id: 0   ticket: 8  
id: 4   ticket: 7  
id: 1   ticket: 6  
id: 3   ticket: 5  
id: 0   ticket: 4  
id: 3   ticket: 3  
id: 2   ticket: 2  
id: 0   ticket: 1  
id: 3   ticket: 0  
id: 1   ticket: -1  
id: 4   ticket: -2  
id: 2   ticket: -3  
id: 0   ticket: -4  
-4 done
```

可见，我们需要使用上锁，我们可以使用互斥量来解决这个问题。下面的代码，我只列出了修改过的内容：

``` java; highlight: [5,9,13,19]
 package main
import &quot;fmt&quot;
import &quot;time&quot;
import &quot;math/rand&quot;
import &quot;sync&quot;
import &quot;runtime&quot;

var total_tickets int32 = 10;
var mutex = &amp;sync.Mutex{} //可简写成：var mutex sync.Mutex

func sell_tickets(i int){
    for total_tickets&gt;0 {
        mutex.Lock()
        if total_tickets &gt; 0 {
            time.Sleep( time.Duration(rand.Intn(5)) * time.Millisecond)
            total_tickets--
            fmt.Println(i, total_tickets)
        }
        mutex.Unlock()
    }
}
.......
......

```

## 原子操作

说到并发就需要说说原子操作，相信大家还记得我写的那篇《<a title="无锁队列的实现" href="http://coolshell.cn/articles/8239.html" target="_blank">无锁队列的实现</a>》一文，里面说到了一些CAS &#8211; CompareAndSwap的操作。Go语言也支持。你可以看一下相当的文档

我在这里就举一个很简单的示例：下面的程序有10个goroutine，每个会对cnt变量累加20次，所以，最后的cnt应该是200。如果没有atomic的原子操作，那么cnt将有可能得到一个小于200的数。

下面使用了atomic操作，所以是安全的。

``` java; highlight: [5,13,18]
package main

import &quot;fmt&quot;
import &quot;time&quot;
import &quot;sync/atomic&quot;

func main() {
    var cnt uint32 = 0
    for i := 0; i &lt; 10; i++ {
        go func() {
            for i:=0; i&lt;20; i++ {
                time.Sleep(time.Millisecond)
                atomic.AddUint32(&amp;cnt, 1)
            }
        }()
    }
    time.Sleep(time.Second)//等一秒钟等goroutine完成
    cntFinal := atomic.LoadUint32(&amp;cnt)//取数据
    fmt.Println(&quot;cnt:&quot;, cntFinal)
}
```

这样的函数还有很多，参看<a href="http://golang.org/pkg/sync/atomic/" target="_blank">go的atomic包文档</a>（被墙）
## Channel 信道

Channal是什么？Channal就是用来通信的，就像Unix下的管道一样，在Go中是这样使用Channel的。

下面的程序演示了一个goroutine和主程序通信的例程。这个程序足够简单了。

``` java; highlight: [7,10]

package main

import &quot;fmt&quot;

func main() {
    //创建一个string类型的channel
    channel := make(chan string)

    //创建一个goroutine向channel里发一个字符串
    go func() { channel &lt;- &quot;hello&quot; }()

    msg := &lt;- channel
    fmt.Println(msg)
}
```

**指定channel的buffer**

指定buffer的大小很简单，看下面的程序：

``` java; highlight: [5]
package main
import &quot;fmt&quot;

func main() {
    channel := make(chan string, 2)

    go func() {
        channel &lt;- &quot;hello&quot;
        channel &lt;- &quot;World&quot;
    }()

    msg1 := &lt;-channel
    msg2 := &lt;-channel
    fmt.Println(msg1, msg2)
}
```

**Channel的阻塞**

注意，channel默认上是阻塞的，也就是说，如果Channel满了，就阻塞写，如果Channel空了，就阻塞读。于是，我们就可以使用这种特性来同步我们的发送和接收端。

下面这个例程说明了这一点，代码有点乱，不过我觉得不难理解。

``` java
package main

import &quot;fmt&quot;
import &quot;time&quot;

func main() {

    channel := make(chan string) //注意: buffer为1

    go func() {
        channel &lt;- &quot;hello&quot;
        fmt.Println(&quot;write \&quot;hello\&quot; done!&quot;)

        channel &lt;- &quot;World&quot; //Reader在Sleep，这里在阻塞
        fmt.Println(&quot;write \&quot;World\&quot; done!&quot;)

        fmt.Println(&quot;Write go sleep...&quot;)
        time.Sleep(3*time.Second)
        channel &lt;- &quot;channel&quot;
        fmt.Println(&quot;write \&quot;channel\&quot; done!&quot;)
    }()

    time.Sleep(2*time.Second)
    fmt.Println(&quot;Reader Wake up...&quot;)

    msg := &lt;-channel
    fmt.Println(&quot;Reader: &quot;, msg)

    msg = &lt;-channel
    fmt.Println(&quot;Reader: &quot;, msg)

    msg = &lt;-channel //Writer在Sleep，这里在阻塞
    fmt.Println(&quot;Reader: &quot;, msg)
}
```

上面的代码输出的结果如下：

``` bash

Reader Wake up...
Reader:  hello
write &quot;hello&quot; done!
write &quot;World&quot; done!
Write go sleep...
Reader:  World
write &quot;channel&quot; done!
Reader:  channel

```

**Channel阻塞的这个特性还有一个好处是，可以让我们的goroutine在运行的一开始就阻塞在从某个channel领任务，这样就可以作成一个类似于线程池一样的东西。关于这个程序我就不写了。我相信你可以自己实现的。**

**多个Channel的select**

``` java
package main
import &quot;time&quot;
import &quot;fmt&quot;

func main() {
    //创建两个channel - c1 c2
    c1 := make(chan string)
    c2 := make(chan string)

    //创建两个goruntine来分别向这两个channel发送数据
    go func() {
        time.Sleep(time.Second * 1)
        c1 &lt;- &quot;Hello&quot;
    }()
    go func() {
        time.Sleep(time.Second * 1)
        c2 &lt;- &quot;World&quot;
    }()

    //使用select来侦听两个channel
    for i := 0; i &lt; 2; i++ {
        select {
        case msg1 := &lt;-c1:
            fmt.Println(&quot;received&quot;, msg1)
        case msg2 := &lt;-c2:
            fmt.Println(&quot;received&quot;, msg2)
        }
    }
}
```

注意：上面的select是阻塞的，所以，才搞出ugly的for i &lt;2这种东西**。<br />
**

**Channel select阻塞的Timeout**

解决上述那个for循环的问题，一般有两种方法：一种是阻塞但有timeout，一种是无阻塞。我们来看看如果给select设置上timeout的。

``` cpp; highlight: [8]

    for {
        timeout_cnt := 0
        select {
        case msg1 := &lt;-c1:
            fmt.Println(&quot;msg1 received&quot;, msg1)
        case msg2 := &lt;-c2:
            fmt.Println(&quot;msg2 received&quot;, msg2)
        case  &lt;-time.After(time.Second * 30)：
            fmt.Println(&quot;Time Out&quot;)
            timout_cnt++
        }
        if time_cnt &gt; 3 {
            break
        }
    }

```

上面代码中高亮的代码主要是用来让select返回的，注意 case中的time.After事件。

**Channel的无阻塞**

好，我们再来看看无阻塞的channel，其实也很简单，就是在select中加入default，如下所示：

``` cpp; highlight: [8]

    for {
        select {
        case msg1 := &lt;-c1:
            fmt.Println(&quot;received&quot;, msg1)
        case msg2 := &lt;-c2:
            fmt.Println(&quot;received&quot;, msg2)
        default: //default会导致无阻塞
            fmt.Println(&quot;nothing received!&quot;)
            time.Sleep(time.Second)
        }
    }

```

**Channel的关闭**

关闭Channel可以通知对方内容发送完了，不用再等了。参看下面的例程：

``` java; highlight: [19,27]
package main

import &quot;fmt&quot;
import &quot;time&quot;
import &quot;math/rand&quot;

func main() {

    channel := make(chan string)
    rand.Seed(time.Now().Unix())

    //向channel发送随机个数的message
    go func () {
        cnt := rand.Intn(10)
        fmt.Println(&quot;message cnt :&quot;, cnt)
        for i:=0; i&lt;cnt; i++{
            channel &lt;- fmt.Sprintf(&quot;message-%2d&quot;, i)
        }
        close(channel) //关闭Channel
    }()

    var more bool = true
    var msg string
    for more {
        select{
        //channel会返回两个值，一个是内容，一个是还有没有内容
        case msg, more = &lt;- channel:
            if more {
                fmt.Println(msg)
            }else{
                fmt.Println(&quot;channel closed!&quot;)
            }
        }
    }
}
```

## 定时器

Go语言中可以使用time.NewTimer或time.NewTicker来设置一个定时器，这个定时器会绑定在你的当前channel中，通过channel的阻塞通知机器来通知你的程序。

下面是一个timer的示例。

``` java; highlight: [9]
package main

import &quot;time&quot;
import &quot;fmt&quot;

func main() {
    timer := time.NewTimer(2*time.Second)

    &lt;- timer.C
    fmt.Println(&quot;timer expired!&quot;)
}
```

上面的例程看起来像一个Sleep，是的，不过Timer是可以Stop的。你需要注意Timer只通知一次。如果你要像C中的Timer能持续通知的话，你需要使用Ticker。下面是Ticker的例程：

``` java; highlight: [9]
package main

import &quot;time&quot;
import &quot;fmt&quot;

func main() {
    ticker := time.NewTicker(time.Second)

    for t := range ticker.C {
        fmt.Println(&quot;Tick at&quot;, t)
    }
}
```

上面的这个ticker会让你程序进入死循环，我们应该放其放在一个goroutine中。下面这个程序结合了timer和ticker

``` java
package main

import &quot;time&quot;
import &quot;fmt&quot;

func main() {

    ticker := time.NewTicker(time.Second)

    go func () {
        for t := range ticker.C {
            fmt.Println(t)
        }
    }()

    //设置一个timer，10钞后停掉ticker
    timer := time.NewTimer(10*time.Second)
    &lt;- timer.C

    ticker.Stop()
    fmt.Println(&quot;timer expired!&quot;)
}
```

## Socket编程

下面是我尝试的一个Echo Server的Socket代码，感觉还是挺简单的。
Server端
``` java; highlight: [12,19,24,33,36]

package main

import (
    &quot;net&quot;
    &quot;fmt&quot;
    &quot;io&quot;
)

const RECV_BUF_LEN = 1024

func main() {
    listener, err := net.Listen(&quot;tcp&quot;, &quot;0.0.0.0:6666&quot;)//侦听在6666端口
    if err != nil {
        panic(&quot;error listening:&quot;+err.Error())
    }
    fmt.Println(&quot;Starting the server&quot;)

    for {
        conn, err := listener.Accept() //接受连接
        if err != nil {
            panic(&quot;Error accept:&quot;+err.Error())
        }
        fmt.Println(&quot;Accepted the Connection :&quot;, conn.RemoteAddr())
        go EchoServer(conn)
    }
}

func EchoServer(conn net.Conn) {
    buf := make([]byte, RECV_BUF_LEN)
    defer conn.Close()

    for {
        n, err := conn.Read(buf);
        switch err {
            case nil:
                conn.Write( buf[0:n] )
            case io.EOF:
                fmt.Printf(&quot;Warning: End of data: %s \n&quot;, err);
                return
            default:
                fmt.Printf(&quot;Error: Reading data : %s \n&quot;, err);
                return
        }
     }
}

```

Client端
``` java; highlight: [12,23,31]

package main

import (
    &quot;fmt&quot;
    &quot;time&quot;
    &quot;net&quot;
)

const RECV_BUF_LEN = 1024

func main() {
    conn,err := net.Dial(&quot;tcp&quot;, &quot;127.0.0.1:6666&quot;)
    if err != nil {
        panic(err.Error())
    }
    defer conn.Close()

    buf := make([]byte, RECV_BUF_LEN)

    for i := 0; i &lt; 5; i++ {
        //准备要发送的字符串
        msg := fmt.Sprintf(&quot;Hello World, %03d&quot;, i)
        n, err := conn.Write([]byte(msg))
        if err != nil {
            println(&quot;Write Buffer Error:&quot;, err.Error())
            break
        }
        fmt.Println(msg)

        //从服务器端收字符串
        n, err = conn.Read(buf)
        if err !=nil {
            println(&quot;Read Buffer Error:&quot;, err.Error())
            break
        }
        fmt.Println(string(buf[0:n]))

        //等一秒钟
        time.Sleep(time.Second)
    }
}

```

## 系统调用

Go语言那么C，所以，一定会有一些系统调用。Go语言主要是通过两个包完成的。一个是<a href="http://golang.org/pkg/os/" target="_blank">os包</a>，一个是<a href="http://golang.org/pkg/syscall/" target="_blank">syscall包</a>。（注意，链接被墙）

这两个包里提供都是Unix-Like的系统调用，
<ul>
<li>syscall里提供了什么Chroot/Chmod/Chmod/Chdir&#8230;，Getenv/Getgid/Getpid/Getgroups/Getpid/Getppid&#8230;，还有很多如Inotify/Ptrace/Epoll/Socket/&#8230;的系统调用。</li>
</ul>
<ul>
<li>os包里提供的东西不多，主要是一个跨平台的调用。它有三个子包，Exec（运行别的命令）, Signal（捕捉信号）和User（通过uid查name之类的）</li>
</ul>

syscall包的东西我不举例了，大家可以看看《Unix高级环境编程》一书。

os里的取几个例：

**环境变量**

``` java
package main

import &quot;os&quot;
import &quot;strings&quot;

func main() {
    os.Setenv(&quot;WEB&quot;, &quot;http://coolshell.cn&quot;) //设置环境变量
    println(os.Getenv(&quot;WEB&quot;)) //读出来

    for _, env := range os.Environ() { //穷举环境变量
        e := strings.Split(env, &quot;=&quot;)
        println(e[0], &quot;=&quot;, e[1])
    }
}

```

## 执行命令行

下面是一个比较简单的示例

``` java

package main
import &quot;os/exec&quot;
import &quot;fmt&quot;
func main() {
    cmd := exec.Command(&quot;ping&quot;, &quot;127.0.0.1&quot;)
    out, err := cmd.Output()
    if err!=nil {
        println(&quot;Command Error!&quot;, err.Error())
        return
    }
    fmt.Println(string(out))
}
```

正规一点的用来处理标准输入和输出的示例如下：

``` java
package main

import (
	&quot;strings&quot;
	&quot;bytes&quot;
	&quot;fmt&quot;
	&quot;log&quot;
	&quot;os/exec&quot;
)

func main() {
	cmd := exec.Command(&quot;tr&quot;, &quot;a-z&quot;, &quot;A-Z&quot;)
	cmd.Stdin = strings.NewReader(&quot;some input&quot;)
	var out bytes.Buffer
	cmd.Stdout = &amp;out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(&quot;in all caps: %q\n&quot;, out.String())
}
```

## 命令行参数

Go语言中处理命令行参数很简单：(使用os的Args就可以了)

``` cpp; highlight: [2]
func main() {
    args := os.Args
    fmt.Println(args) //带执行文件的
    fmt.Println(args[1:]) //不带执行文件的
}
```

在Windows下，如果运行结果如下：

`C:\Projects\Go>go run args.go aaa bbb ccc ddd<br />
[C:\Users\haoel\AppData\Local\Temp\go-build742679827\command-line-arguments\_<br />
obj\a.out.exe aaa bbb ccc ddd]<br />
[aaa bbb ccc ddd]`

那么，如果我们要搞出一些像 mysql -uRoot -hLocalhost -pPwd 或是像 cc -O3 -Wall -o a a.c 这样的命令行参数我们怎么办？Go提供了一个package叫flag可以容易地做到这一点

``` java; highlight: [13]
package main
import &quot;flag&quot;
import &quot;fmt&quot;

func main() {

    //第一个参数是“参数名”，第二个是“默认值”，第三个是“说明”。返回的是指针
    host := flag.String(&quot;host&quot;, &quot;coolshell.cn&quot;, &quot;a host name &quot;)
    port := flag.Int(&quot;port&quot;, 80, &quot;a port number&quot;)
    debug := flag.Bool(&quot;d&quot;, false, &quot;enable/disable debug mode&quot;)

    //正式开始Parse命令行参数
    flag.Parse()

    fmt.Println(&quot;host:&quot;, *host)
    fmt.Println(&quot;port:&quot;, *port)
    fmt.Println(&quot;debug:&quot;, *debug)
}
```

执行起来会是这个样子：

``` bash
#如果没有指定参数名，则使用默认值
$ go run flagtest.go
host: coolshell.cn
port: 80
debug: false

#指定了参数名后的情况
$ go run flagtest.go -host=localhost -port=22 -d
host: localhost
port: 22
debug: true

#用法出错了（如：使用了不支持的参数，参数没有=）
$ go build flagtest.go
$ ./flagtest -debug -host localhost -port=22
flag provided but not defined: -debug
Usage of flagtest:
  -d=false: enable/disable debug mode
  -host=&quot;coolshell.cn&quot;: a host name
  -port=80: a port number
exit status 2
```

感觉还是挺不错的吧。
## 一个简单的HTTP Server

代码胜过千言万语。呵呵。这个小程序让我又找回以前用C写CGI的时光了。（Go的官方文档是《**<a href="http://golang.org/doc/articles/wiki/" target="_blank">Writing Web Applications</a>**》）

``` java
package main

import (
    &quot;fmt&quot;
    &quot;net/http&quot;
    &quot;io/ioutil&quot;
    &quot;path/filepath&quot;
)

const http_root = &quot;/home/haoel/coolshell.cn/&quot;

func main() {
    http.HandleFunc(&quot;/&quot;, rootHandler)
    http.HandleFunc(&quot;/view/&quot;, viewHandler)
    http.HandleFunc(&quot;/html/&quot;, htmlHandler)

    http.ListenAndServe(&quot;:8080&quot;, nil)
}

//读取一些HTTP的头
func rootHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, &quot;rootHandler: %s\n&quot;, r.URL.Path)
    fmt.Fprintf(w, &quot;URL: %s\n&quot;, r.URL)
    fmt.Fprintf(w, &quot;Method: %s\n&quot;, r.Method)
    fmt.Fprintf(w, &quot;RequestURI: %s\n&quot;, r.RequestURI )
    fmt.Fprintf(w, &quot;Proto: %s\n&quot;, r.Proto)
    fmt.Fprintf(w, &quot;HOST: %s\n&quot;, r.Host) 
}

//特别的URL处理
func viewHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, &quot;viewHandler: %s&quot;, r.URL.Path)
}

//一个静态网页的服务示例。（在http_root的html目录下）
func htmlHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Printf(&quot;htmlHandler: %s\n&quot;, r.URL.Path)
    
    filename := http_root + r.URL.Path
    fileext := filepath.Ext(filename)

    content, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Printf(&quot;   404 Not Found!\n&quot;)
        w.WriteHeader(http.StatusNotFound)
        return
    }
    
    var contype string
    switch fileext {
        case &quot;.html&quot;, &quot;htm&quot;:
            contype = &quot;text/html&quot;
        case &quot;.css&quot;:
            contype = &quot;text/css&quot;
        case &quot;.js&quot;:
            contype = &quot;application/javascript&quot;
        case &quot;.png&quot;:
            contype = &quot;image/png&quot;
        case &quot;.jpg&quot;, &quot;.jpeg&quot;:
            contype = &quot;image/jpeg&quot;
        case &quot;.gif&quot;:
            contype = &quot;image/gif&quot;
        default: 
            contype = &quot;text/plain&quot;
    }
    fmt.Printf(&quot;ext %s, ct = %s\n&quot;, fileext, contype)
    
    w.Header().Set(&quot;Content-Type&quot;, contype)
    fmt.Fprintf(w, &quot;%s&quot;, content)
    
}
```

Go的功能库有很多，大家自己慢慢看吧。**我再吐个槽——Go的文档真不好读。例子太少了**。

先说这么多吧。这是我周末两天学Go语言学到的东西，写得太仓促了，而且还有一些东西理解不到位，还大家请指正！

（全文完）