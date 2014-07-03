#GO语言标准库概览

<http://my.oschina.net/itfanr/blog/212654>

本文翻译自Dr.Dobb's的"A Brief Tour of the Go Standard Library"一文。

在Go语言五周系列教程的最后一部分中，我们将带领大家一起来浏览一下Go语言丰富的标准库。

Go标准库包含了大量包，提供了丰富广泛的功能特性。这里提供了概览仅仅是有选择性的且非常简单。本文发表后，标准库的内容还可能继续增加，因此 建议大家最好是通过在线查阅库API或使用godoc（包含在Go发布包中）来获取最新信息以及全面了解每个包所具备的功能。


exp包(试验性的)是那些未来可能被加入标准库的包起步的地方，因此除非你想参加这些包的开发(通过测试、讨论、提交补丁)，否则不应该使用其 下面的包。exp包通常只存在于从Google Go源码树上签出的源码包中，但一般不会包含在预构建好的包中。其他包可以放心使用，虽然在写下本文的这一刻，很多包依旧不够完整。



## Archive(归档)和Compression(压缩)包

Go支持读写tarball和.zip文件。与此相关的包为archive/tar和archive/zip；以及用于压缩tarball的 compress/gzip和compress/bzip2。

Go同样也支持其他压缩格式；例如用于TIFF图像和PDF文件的Lempel-Ziv-Welch (compress/lzw)格式。



## Bytes(字节)和String(字符串)相关包

bytes和strings包中有很多相同的函数，但前者操作的是[]byte类型的值，而后者操作的是string类型的值。strings包 提供了所有最有用的功能函数，诸如查找子字符串、替换子字符串、拆分字符串、剔除字符串以及大小写变换等。strconv包提供了数字和布尔类型 与string类型相互转换的功能。



## fmt包

提供了大量有用的print和scan函数，它们在本系列教程的第一和第二部分已有相关介绍。



## unicode包

提供一些用于确定字符属性的函数，诸如判断一个字符是否是可打印的，或是否是一个数字。unicode/utf8与 unicode/utf16这两个包提供了rune(即，Unicode码点/字符)的编码和解码功能。



## text/template和html/template包

可以被用于创建模板，这些模板可基于填入的数据生成文本形式的输出(例如HTML)。 这里是一个小且简单的有关text/template包使用的例子。

```go
type GiniIndex struct {
    Country string
    Index float64
}
gini := []GiniIndex{{"Japan", 54.7}, {"China", 55.0}, {"U.S.A.", 80.1}}
giniTable := template.New("giniTable")
giniTable.Parse(
'<TABLE>' +
    '{{range .}}' +
    '{{printf "<TR><TD>%s</TD><TD>%.1f%%</TD></TR>"' +
    '.Country .Index}}'+
    '{{end}}' +
    '</TABLE>')
err := giniTable.Execute(os.Stdout, gini)
```

输出：

```html
<TABLE>
<TR><TD>Japan</TD><TD>54.7%</TD></TR>
<TR><TD>China</TD><TD>55.0%</TD></TR>
<TR><TD>U.S.A.</TD><TD>80.1%</TD></TR>
</TABLE>
```


template.New()函数用给定的名字创建了一个新的template.Template。模板名字用于识别模板，尤其是嵌入在其他模板 中时。template.Template.Parse()函数用于解析一个模板(通常从一个.html文件中)，解析后模板即可用。 template.Template.Execute()函数执行这个模板，将结果输出到给定的io.Writer，并且从其第二个参数那里读取 用于生成模板的数据。在这个例子中，我将结果输出到os.Stdout，并将GiniIndex类型的gini切片作为数据传入(我将输出拆分为 多行以便让结果更加清晰)。

在模板内部，行为(action)包含在双大括号中`({{和}})。{{range}} … {{end}}`可用于迭代访问一个切片中的每个元素。这里我将切片中的每个GiniIndex设置为点(`.`)；即是当前的元素。我们可以通过在名字访问导 出字段，当然名字前面需要用.来指明当前元素。{{printf}}的行为与fmt.Printf()函数类似，但用空格替换括号以及用于分隔参 数的逗号。

text/template和html/template包自身支持一种复杂的模板语言，包括许多action，迭代和条件分支，支持变量和方法 调用，以及其它一些。除此之外，html/templage包还对代码注入免疫。



## 集合包

切片是Go语言提供了最高效的集合类型，但有些时候使用一个更为特定的集合类型更有用或有必要。在多数情况下，内置的map类型已经足够了，但Go标准库还是提供了container包，其中包含了各种不同的集合包。



### container/heap包

提供了操作**heap(堆)**的函数，这里heap必须是一个自定义类型的值，该类型必须满足定义在heap包中 heap.Interface。一个heap(严格地说是一个min-heap)按特定顺序维护其中的值 – 即第一个元素总是heap中最小的（对于max-heap，应该是最大的）- 这就是熟知的heap属性。heap.Interface中嵌入了sort.Interface以及Push()和Pop方法。

我们可以很容易地创建一个满足heap.Interface的自定义heap类型。下面是一个正在使用的heap的例子：

```go
ints := &IntHeap{5, 1, 6, 7, 9, 8, 2, 4}
heap.Init(ints) // Heapify
ints.Push(9) // IntHeap.Push() doesn't preserve the heap property
ints.Push(7)
ints.Push(3)
heap.Init(ints) // Must reheapify after heap-breaking changes
for ints.Len() > 0 {
    fmt.Printf("%v ", heap.Pop(ints))
}
fmt.Println() // prints: 1 2 3 4 5 6 7 7 8 9 9
```

下面是完整的自定义heap实现。

```go
type IntHeap []int

func (ints *IntHeap) Less(i, j int) bool {
    return (*ints)[i] < (*ints)[j]
}

func (ints *IntHeap) Swap(i, j int) {
    (*ints)[i], (*ints)[j] = (*ints)[j], (*ints)[i]
}

func (ints *IntHeap) Len() int {
    return len(*ints)
}

func (ints *IntHeap) Pop() interface{} {
    x := (*ints)[ints.Len()-1]
    *ints = (*ints)[:ints.Len()-1]
    return x
}

func (ints *IntHeap) Push(x interface{}) {
    *ints = append(*ints, x.(int))
}
```


对于多数情况这个实现都足以应付了。我们可以将IntHeap类型换为type IntHeap struct { ints []int}，这样代码更漂亮，我们可以在方法内部使用ints.ints，而不再是*ints了。



### container/list包

提供了双向链表。元素以interface{}类型的值加入链表。从链表中获取的元素的类型为list.Element，其原始值可通过list.Element.Value访问到。

```go
items := list.New()
for _, x := range strings.Split("ABCDEFGH", "") {
    items.PushFront(x)
}
items.PushBack(9)
for element := items.Front(); element != nil;
    element = element.Next() {
    switch value := element.Value.(type) {
        case string:
            fmt.Printf("%s ", value)
        case int:
            fmt.Printf("%d ", value)
    }
}
fmt.Println() // prints: H G F E D B A 9
```

在这里例子中，我们将8个单字母字符串推入一个新链表的前端，将一个整型数推入尾端。接下来，我们迭代访问链表中的元素并打印元素的值。我们不是真的需要 type switch，因为我们可以使用fmt.Printf(%v ", element.Value)打印元素值，但如果我们要做的不仅仅是打印的话，如果列表中包含的元素类型不同，我们将需要type switch。当然，如果所有的元素都具有同一类型，我们可以使用type assertion，例如对string类型元素，我们使用element.Value.(string)。

除了上述提到的方法之外，list.List类型还提供了其他许多方法，包括Back()，Init()（用于清理链 表），InsertAfter()，InsertBefore()，Len()，MoveToBack()，MoveToFront()，PushBackList() （用于将一个链表推入另外一个链表的尾端），以及Remove()。



标准库还提供了**container/ring**包，这个包实现了一个环形链表。

然而所有的集合类型都将数据存储在内存中，Go还提供了database/sql包，该包提供了一个通用的SQL数据库接口。当与真实数据库交互时，特定的数据库驱动包必须被单独安装。这些包，以及其他许多集合包都放在了Go Bashboard上。



## 文件，操作系统以及相关包

标准库提供了许多支持文件和目录操作以及与操作系统交互的包。在许多情况下，这些包提供了操作系统无关的抽象使得创建跨平台Go应用更为简单。



### os(操作系统)包

提供了与操作系统交互相关的函数，诸如改变当前工作目录，修改文件模式和所有权，获取和设置环境变量，创建和删除文件和目录等。

此外，该包还提供了创建和打开文件(os.Create()和os.Open())、获取文件属性(例如，通过os.FileInfo类型)，以及在之前系列文章中我们所见过的函数。

一旦文件被打开，尤其是对于那些文本文件，通过一个buffer来访问该文件是非常常见的情况(将读取的行存入字符串而不是byte切片)。我们需要的这 个功能由bufio包提供。除了用bufio.Reader和bufio.Writer进行读写字符串外，我们还可以读(不读)rune，读(不读)单字 节，读多字节以及写rune和写单字节以及多字节。



### io(input/output)包

提供了大量的函数用于与io.Reader和io.Writer一起工作(这两个接口都可以被os.File类型值满 足)。例如，我们曾用过io.Copy()函数将数据从一个reader拷贝到一个writer中。这个包还包含了用于创建同步的内存管道(pipe)的函数。



### io/iotuil包

提供了一些非常易用的函数。其中，这个包提供的ioutil.ReadAll()函数用于读取一个io.Reader的所有数据，并 将数据放入一个[]byte中返回；ioutil.ReadFile()函数所做的事情类似，只是参数由一个io.Reader换成了一个字符串(文件 名)；ioutil.TempFile()函数返回一个临时文件(一个os.File)；ioutil.WriteFile()函数向一个给定名字的文件 中写入由[]byte承载的数据。



### path包

提供的函数用于操作Unix样式路径，例如Linux和Mac OS X路径，用于处理URL路径，git“引用”，FTP文件等。path/filepath包提供提供了与path相同的函数- 许多其他的 – 函数被设计用于提供平台中立的路径处理。这个包还提供了filepath.Walk()函数用于递归地对给定路径下的所有文件和目录进行迭代访问。



### runtime包

包含了许多函数和类型用于访问Go的运行时系统。这里面的大多数都是高级功能，在日常创建标准Go程序时不应该使用到这些功能。但是，一 些包中的常量可能十分有用 – 例如，字符串runtime.GOOS(其值例如，"darwin," "freebsd," "linux," 或 "windows")和字符串runtime.GOARCH(其值例如386," "amd64,"或 "arm")。runtime.GOROOT()函数返回GOROOT环境变量的值(或者如果该环境变量没有设置，返回Go构建根目 录)，runtime.Version()返回Go版本(以一个字符串形式)。runtime.GOMAXPROCS()和 runtime.NumCPU()函数保证Go使用机器的所有处理器，在Go的文档中有详尽解释。



## 文件格式相关包

Go提供出色的文件处理功能，既可用于文本文件(使用7-bit ASCII编码或UTF-8和UTF-16 Unicode编码），也可用于二进制文件。Go提供了专门的包，用于处理JSON和XML文件以及它自己专有的快速、简洁以及方便的Go二进制格式。此 外，Go提供了csv包用于读取CSV(逗号分隔的值)文件。这个包将这些文件视为记录(每行算作一个记录)，么个记录由多个(逗号分隔的)字段组成。这 个包用途非常广泛，例如，可以用它修改分隔符(从逗号改为tab或其他字符)，以及其他诸如如何读写记录和字段的方面。



### encoding包

包含许多子包，其中的encoding/binary包我们曾用于读写二进制数据。其他包提供了针对各种格式的编解码功能 – 例如，encoding/base64包可以用于编码和解码我们日常常用的URL。



## 图像相关包

Go的image包提供了一些高层次的函数和类型，用于创建和持有图像数据。它还提供了一些包，可用于不同种类标准图像文件格式的编解码，例如image/jpeg和image/png。



### image/draw包

提供了一些基本的绘图函数。第三方的freetype包加入了更多绘图函数。freetype自身可以使用任意指定TrueType字体绘制文本，freetype/raster包可以绘制线条以及立方和二次曲线。



## 数学包

math/big包提供了无限大(实际受限于内存)整型数(big.Int)以及有理数(big.Rat)。math包提供了所有标准数学函数(基于float64)以及一些标准常量。math/cmplx包提供一些用于复数计算的标准函数(基于complex128)。



## 其他杂项包

除了这些可以被粗略分组的包外，标准库还包含了许多相对独立的包。



### crypto包

提供了使用MD5, SHA-1, SHA-224, SHA-256, SHA-384以及SHA-512算法的Hash(每个算法由一个包提供，例如crypto/sha512)。此外，crypto还提供了用于加密和解密 的子包，这些包使用了不同算法，诸如AES、DES等等。每个包都对应相应的名字(例如，crypto/aes和crypto/des)。



### exec包

用于运行外部程序。我们也可以使用os.StartProcess来完成这件事，但exec.Cmd类型用起来更加简单。



### flag包

提供了一个命令行解析器。它接受X11风格的命令行选项(例如，-width，非GNU风格的-w以及–width)。这个包产生一个非常基 本的usage消息并且没有提供除值类型之外的任何校验(因此，这个包可以用于指定一个int型选项，而不是用于检查接受哪些值)。还有一些候选包可以在 Go Bashboard中找到。



### log包

提供了一些函数，用于记录日志信息(默认输出到os.Stdout)、结束程序或抛出异常(panick)并携带一条日志信息。log包输出目标 可以使用log.SetOutput()函数变更为任何io.Writer。日志信息以一个时间戳加后续消息的格式输出；时间戳的分割符可以在调用第一个 log函数之前通过log.SetFlags(0)设置。通过log.New()函数我们还可以创建自定义的logger。



### math/rand包

提供许多有用的伪随机数生成函数，包括返回一个随机整型数的rand.Int()以及rand.Intn(n)，后者返回[0,n]范围内的一个随机整数。crypto/rand包中有一个函数，可用于产生加密的强伪随机数字。regexp包提供快速且强大的正则式引擎，并支持RE2引擎的语法。



### sort包

提供了许多方便易用的函数，用于对ints、float64以及string类型的切片进行排序，并且提供基于有序切片的高效(二分查找)的查找。它还提供了用于自定义数据的通用sort.Sort()和sort.Search函数。



### time包

提供了用于测量时间、解析和格式化日期，日期/时间以及时间值的函数。time.After()函数可用于在特定纳秒后，向通道 (channel)发送当前时间。time.Tick()和time.NewTicker()函数可用于提供一个通道，它会返回在特定时间间隔后将 'tick'发送到该通道上。time.Time结构具有一些方法，可提供当前时间，将data/time格式化为一个字符串以及解析data /time。



## 网络包net/http包

Go标准库中有许多包用于支持网络以及相关方面的编程。net包提供的函数和类型可用于使用Unix域以及网络socket通信、TCP/IP和UDP编程。


这个包还提供了用于域名解析的函数。net/http包充分利用了net包，并提供了解析HTTP请求和应答的功能，并提供了一个基本的HTTP客户端。net/http包也包含一个易于扩展的HTTP server。net/url包提供了URL解析和查询转义。



标准库中还包含其他一些其他高层次的网络包。一个是net/rpc(远程过程调用)包，它允许一个服务端提供导出可被客户端调用的方法的对象。另外一个是net/smtp(简单邮件传输协议)包，可用于发送email。



## Reﬂect包

reflect包提供了运行时反射(或称为自省)；即，在运行时访问和与任意类型的值交互的能力。

这个包还提供了一些有用的工具函数，诸如reflect.DeepEqual()用于比较任意两个值 – 例如，切片，我们无法用==和!=操作符对其进行比较。

Go中的每个值都有两个属性：它的实际值与类型。reflect.TypeOf()函数可以告诉我们任意值的类型。

```go
x := 8.6
y := float32(2.5)
fmt.Printf("var x %v = %v\n", reflect.TypeOf(x), x)
fmt.Printf("var y %v = %v\n", reflect.TypeOf(y), y)
```

输出

```go
var x float64 = 8.6
var y float32 = 2.5
```

这里我们使用reflection输出两个浮点变量和它们的类型，类似Go的变量声明。

当将reflect.ValueOf函数用于一个值时，该函数返回一个reflect.Value，它持有值但它本身却不是那个值。如果我们想访问那个被持有的值，我们必须使用reflect.Value的一个方法。

```go
word := "Chameleon"
value := reflect.ValueOf(word)
text := value.String()
fmt.Println(text)
```

输出：

```go
Chameleon
```

reflect.Value类型拥有很多可以提取底层类型的方法，包括 reflect.Value.Bool(), reflect.Value.Complex(), reflect.Value.Float(), reflect.Value.Int(),以及reflect.Value.String()。

reflect包也可以与集合类型一起使用，比如切片和map，也可以与struct一起使用；它甚至可以访问结构体tag的文本（这种能力被用到了JSON和XML的编码和解码中）。

```go
type Contact struct {
    Name string "check:len(3,40)"
    Id int "check:range(1,999999)"
}
person := Contact{"Bjork", 0xDEEDED}
personType := reflect.TypeOf(person)
if nameField, ok := personType.FieldByName("Name"); ok {
    fmt.Printf("%q %q %q\n", nameField.Type, nameField.Name, nameField.Tag)
}
```

输出：

```
"string" "Name" "check:len(3,40)"
```


reflect.Value持有的真实值如果是"可设置的"，那么它可以被改变。是否具备可设置能力可以通过reflect.Value.CanSet()来获知，该函数返回一个布尔值。

```go
presidents := []string{"Obama", "Bushy", "Clinton"}
sliceValue := reflect.ValueOf(presidents)
value = sliceValue.Index(1)
value.SetString("Bush")
fmt.Println(presidents)
```

输出：

```text
[Obama Bush Clinton]
```

虽然Go的字符串是不可改变的，但给定[]string中的任意一个元素都可以被另外一个字符串所替代，这就是我们在这里所做的。(顺利成章地，在这个特定的例子中，最容易的修改方法应该是presidents[1] = "Bush"，而且完全没有用到自省特性)。

你无法改变不可改变的值本身，但如果我们得到原值的地址，我们可以将原不可改变的值替换为另一个新值。

```go
count := 1
if value = reflect.ValueOf(count); value.CanSet() {
    value.SetInt(2) // 将抛出异常，我们不能设置int
}
fmt.Print(count, " ")

value = reflect.ValueOf(&count)
// 不能在值上调用SetInt()，因为值是一个*int，而不是一个int
pointee := value.Elem()
pointee.SetInt(3) // OK. 我们可以通过值指针替换
fmt.Println(count)
```

输出：


```text
1 3
```

这小段代码的输出表明如果条件表达式求值结果为false，其分支语句将不会被执行。虽然我们无法重新设置那些不可改变的值，诸如ints、 float64或字符串，但我们可以使用`reflect.Value.Elem()`方法来获取一个`reflectValue`，通过它我们可以重新设置该地 址上的值，这就是我们在这段代码结尾处所做的。



我们还可以使用反射来调用任意函数和方法。这里是一个例子，例子用两次调用了自定义函数TitleCase，一次是用传统的方式，一次则是用反射。

```go
caption := "greg egan's dark integers"
title := TitleCase(caption)
fmt.Println(title)

titleFuncValue := reflect.ValueOf(TitleCase)
values := titleFuncValue.Call(
[]reflect.Value{reflect.ValueOf(caption)})
title = values[0].String()
fmt.Println(title)
```

输出：

```
Greg Egan's Dark Integers
Greg Egan's Dark Integers
```

reflect.Value.Call()方法接受以及返回一个类型[]reflect.Value的切片。在这个例子中，我们传入一个单一值(作为一个长度为1的切片)，并获取到一个单一的结果值。



我们可以用同样的方法调用方法 – 事实上，我们甚至可以查询一个方法是否存在，并且在它确实存在的情况下再调用它。

```go
a := list.New() // a.Len() == 0
b := list.New()
b.PushFront(1) // b.Len() == 1
c := stack.Stack{}
c.Push(0.5)
c.Push(1.5) // c.Len() == 2
d := map[string]int{"A": 1, "B": 2, "C": 3} // len(d) == 3
e := "Four" // len(e) == 4
f := []int{5, 0, 4, 1, 3} // len(f) == 5
fmt.Println(Len(a), Len(b), Len(c), Len(d), Len(e), Len(f))
```

输出：

```
0 1 2 3 4 5
```


这里我们创建了两个链表(使用container/list包)，我们给其中一个加入一个元素。我们还创建了一个stack，并向其中加入两个元素。我们 接下来创建了一个map，一个字符串以及一个int类型切片，它们长度各不相同。我们使用Len()函数获取了它们的长度。

```go
func Len(x interface{}) int {
    value := reflect.ValueOf(x)
    switch reflect.TypeOf(x).Kind() {
        case reflect.Array, reflect.Chan, reflect.Map,
                reflect.Slice, reflect.String:
            return value.Len()
        default:
            if method := value.MethodByName("Len"); method.IsValid() {
                values := method.Call(nil)
                return int(values[0].Int())
            }
        }
    panic(fmt.Sprintf("'%v' does not have a length", x))
}
```


这个函数返回传入值的长度或当值类型不支持长度概念时引发异常。



我们开始获得reflect.Value类型值，因为我们后续需要这个值。接下来我们根据reflect.Kind做switch判断。如果value的 kind是某支持内建len()函数的内建类型的话，我们可以在该值上直接调用reflect.Value.Len()函数。否则，我们要么得到一个不支 持长度概念的类型，要么是一个拥有Len()方法的类型。我们使用reflect.Value.MethodByName()方法来获取这个方法-或者获 取一个无效的reflect.Value。如果这个方法有效，我们就调用它。

这个例子用没有任何参数传入，因为传统Len()方法不接收任何参数。当我们使用reflect.Value.MethodByName()方法获取一个 方法时，返回的reflect.Value既持有方法，又持有这个value。因此当我们调用reflect.Value.Call()时，这个 value将传入并作为接收者。



`reflect.Value.Int()`方法返回一个int64类型值；我们这里已将其转换成一个普通的int以匹配通用的Len()函数的返回值类型。



如果一个传入的值不支持内建的len()函数并且没有Len()方法，通用的Len()将引发异常。我们本可以采用其他方式处理这个错误情况 – 例如，返回-1一表明"不支持长度"，或返回一个整型值和一个错误码。

Go的reflect包十分灵活，允许我们在运行时根据程序的动态状态做一些事情。但是，这里引用Rob Pike的观点，反射是“一个强大的工具，需谨慎并尽量避免使用，除非非常必要。(Rob Pick撰写了一篇非常有趣和实用的有关Go反射的博客文章)。



## 结论

这篇文章给Go语言五周系列教程做了一个收尾。此时此刻，你应该对这门语言，其工具以及它的标准库有了一个很好的感性认识了。- 甚至是如何在Google App Engine上编写Go程序以运行一个Web应用。我希望你能认可这一点：Go是一门非常有趣的语言，它提供了一种编写可移植的、本地代码的愉快的方式。



&copy; 2012, bigwhite. 版权所有.