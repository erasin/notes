Golang中的正则表达式

【用法】

单一：

    .                   匹配任意一个字符，如果设置 s = true，则可以匹配换行符

    [字符类]            匹配“字符类”中的一个字符，“字符类”见后面的说明
    [^字符类]           匹配“字符类”外的一个字符，“字符类”见后面的说明

    \小写Perl标记       匹配“Perl类”中的一个字符，“Perl类”见后面的说明
    \大写Perl标记       匹配“Perl类”外的一个字符，“Perl类”见后面的说明

    [:ASCII类名:]       匹配“ASCII类”中的一个字符，“ASCII类”见后面的说明
    [:^ASCII类名:]      匹配“ASCII类”外的一个字符，“ASCII类”见后面的说明

    \pUnicode普通类名   匹配“Unicode类”中的一个字符(仅普通类)，“Unicode类”见后面的说明
    \PUnicode普通类名   匹配“Unicode类”外的一个字符(仅普通类)，“Unicode类”见后面的说明

    \p{Unicode类名}     匹配“Unicode类”中的一个字符，“Unicode类”见后面的说明
    \P{Unicode类名}     匹配“Unicode类”外的一个字符，“Unicode类”见后面的说明
复合：

    xy             匹配 xy（x 后面跟随 y）
    x|y            匹配 x 或 y (优先匹配 x)
重复：

    x*             匹配零个或多个 x，优先匹配更多(贪婪)
    x+             匹配一个或多个 x，优先匹配更多(贪婪)
    x?             匹配零个或一个 x，优先匹配一个(贪婪)
    x{n,m}         匹配 n 到 m 个 x，优先匹配更多(贪婪)
    x{n,}          匹配 n 个或多个 x，优先匹配更多(贪婪)
    x{n}           只匹配 n 个 x
    x*?            匹配零个或多个 x，优先匹配更少(非贪婪)
    x+?            匹配一个或多个 x，优先匹配更少(非贪婪)
    x??            匹配零个或一个 x，优先匹配零个(非贪婪)
    x{n,m}?        匹配 n 到 m 个 x，优先匹配更少(非贪婪)
    x{n,}?         匹配 n 个或多个 x，优先匹配更少(非贪婪)
    x{n}?          只匹配 n 个 x
分组：

    (子表达式)            被捕获的组，该组被编号 (子匹配)
    (?P<命名>子表达式)    被捕获的组，该组被编号且被命名 (子匹配)
    (?:子表达式)          非捕获的组 (子匹配)
    (?标记)               在组内设置标记，非捕获，标记影响当前组后的正则表达式
    (?标记:子表达式)      在组内设置标记，非捕获，标记影响当前组内的子表达式

    标记的语法是：
    xyz  (设置 xyz 标记)
    -xyz (清除 xyz 标记)
    xy-z (设置 xy 标记, 清除 z 标记)

    可以设置的标记有：
    i              不区分大小写 (默认为 false)
    m              多行模式：让 ^ 和 $ 匹配整个文本的开头和结尾，而非行首和行尾(默认为 false)
    s              让 . 匹配 \n (默认为 false)
    U              非贪婪模式：交换 x* 和 x*? 等的含义 (默认为 false)
位置标记：

    ^              如果标记 m=true 则匹配行首，否则匹配整个文本的开头（m 默认为 false）
    $              如果标记 m=true 则匹配行尾，否则匹配整个文本的结尾（m 默认为 false）
    \A             匹配整个文本的开头，忽略 m 标记
    \b             匹配单词边界
    \B             匹配非单词边界
    \z             匹配整个文本的结尾，忽略 m 标记
转义序列：

    \a             匹配响铃符    （相当于 \x07）
                   注意：正则表达式中不能使用 \b 匹配退格符，因为 \b 被用来匹配单词边界，
                   可以使用 \x08 表示退格符。
    \f             匹配换页符    （相当于 \x0C）
    \t             匹配横向制表符（相当于 \x09）
    \n             匹配换行符    （相当于 \x0A）
    \r             匹配回车符    （相当于 \x0D）
    \v             匹配纵向制表符（相当于 \x0B）
    \123           匹配 8  進制编码所代表的字符（必须是 3 位数字）
    \x7F           匹配 16 進制编码所代表的字符（必须是 3 位数字）
    \x{10FFFF}     匹配 16 進制编码所代表的字符（最大值 10FFFF  ）
    \Q...\E        匹配 \Q 和 \E 之间的文本，忽略文本中的正则语法

    \\             匹配字符 \
    \^             匹配字符 ^
    \$             匹配字符 $
    \.             匹配字符 .
    \*             匹配字符 *
    \+             匹配字符 +
    \?             匹配字符 ?
    \{             匹配字符 {
    \}             匹配字符 }
    \(             匹配字符 (
    \)             匹配字符 )
    \[             匹配字符 [
    \]             匹配字符 ]
    \|             匹配字符 |
可以将“命名字符类”作为“字符类”的元素：

    [\d]           匹配数字 (相当于 \d)
    [^\d]          匹配非数字 (相当于 \D)
    [\D]           匹配非数字 (相当于 \D)
    [^\D]          匹配数字 (相当于 \d)
    [[:name:]]     命名的“ASCII 类”包含在“字符类”中 (相当于 [:name:])
    [^[:name:]]    命名的“ASCII 类”不包含在“字符类”中 (相当于 [:^name:])
    [\p{Name}]     命名的“Unicode 类”包含在“字符类”中 (相当于 \p{Name})
    [^\p{Name}]    命名的“Unicode 类”不包含在“字符类”中 (相当于 \P{Name})
【说明】

“字符类”取值如下（“字符类”包含“Perl类”、“ASCII类”、“Unicode类”）：

x                    单个字符
A-Z                  字符范围(包含首尾字符)
\小写字母            Perl类
[:ASCII类名:]        ASCII类
\p{Unicode脚本类名}  Unicode类 (脚本类)
\pUnicode普通类名    Unicode类 (普通类)
“Perl 类”取值如下：

\d             数字 (相当于 [0-9])
\D             非数字 (相当于 [^0-9])
\s             空白 (相当于 [\t\n\f\r ])
\S             非空白 (相当于[^\t\n\f\r ])
\w             单词字符 (相当于 [0-9A-Za-z_])
\W             非单词字符 (相当于 [^0-9A-Za-z_])
“ASCII 类”取值如下

[:alnum:]      字母数字 (相当于 [0-9A-Za-z])
[:alpha:]      字母 (相当于 [A-Za-z])
[:ascii:]      ASCII 字符集 (相当于 [\x00-\x7F])
[:blank:]      空白占位符 (相当于 [\t ])
[:cntrl:]      控制字符 (相当于 [\x00-\x1F\x7F])
[:digit:]      数字 (相当于 [0-9])
[:graph:]      图形字符 (相当于 [!-~] 相当于 [A-Za-z0-9!"#$%&'()*+,\-./:;<=>?@[\\\]^_`{|}~])
[:lower:]      小写字母 (相当于 [a-z])
[:print:]      可打印字符 (相当于 [ -~] 相当于 [ [:graph:]])
[:punct:]      标点符号 (相当于 [!-/:-@[-`{-~])
[:space:]      空白字符(相当于 [\t\n\v\f\r ])
[:upper:]      大写字母(相当于 [A-Z])
[:word:]       单词字符(相当于 [0-9A-Za-z_])
[:xdigit:]     16 進制字符集(相当于 [0-9A-Fa-f])
“Unicode 类”取值如下—普通类：

C                 -其他-          (other)
Cc                控制字符        (control)
Cf                格式            (format)
Co                私人使用区      (private use)
Cs                代理区          (surrogate)
L                 -字母-          (letter)
Ll                小写字母        (lowercase letter)
Lm                修饰字母        (modifier letter)
Lo                其它字母        (other letter)
Lt                首字母大写字母  (titlecase letter)
Lu                大写字母        (uppercase letter)
M                 -标记-          (mark)
Mc                间距标记        (spacing mark)
Me                关闭标记        (enclosing mark)
Mn                非间距标记      (non-spacing mark)
N                 -数字-          (number)
Nd                十進制数字      (decimal number)
Nl                字母数字        (letter number)
No                其它数字        (other number)
P                 -标点-          (punctuation)
Pc                连接符标点      (connector punctuation)
Pd                破折号标点符号  (dash punctuation)
Pe                关闭的标点符号  (close punctuation)
Pf                最后的标点符号  (final punctuation)
Pi                最初的标点符号  (initial punctuation)
Po                其他标点符号    (other punctuation)
Ps                开放的标点符号  (open punctuation)
S                 -符号-          (symbol)
Sc                货币符号        (currency symbol)
Sk                修饰符号        (modifier symbol)
Sm                数学符号        (math symbol)
So                其他符号        (other symbol)
Z                 -分隔符-        (separator)
Zl                行分隔符        (line separator)
Zp                段落分隔符      (paragraph separator)
Zs                空白分隔符      (space separator)
“Unicode 类”取值如下—脚本类：

Arabic                  阿拉伯文
Armenian                亚美尼亚文
Balinese                巴厘岛文
Bengali                 孟加拉文
Bopomofo                汉语拼音字母
Braille                 盲文
Buginese                布吉文
Buhid                   布希德文
Canadian_Aboriginal     加拿大土著文
Carian                  卡里亚文
Cham                    占族文
Cherokee                切诺基文
Common                  普通的，字符不是特定于一个脚本
Coptic                  科普特文
Cuneiform               楔形文字
Cypriot                 塞浦路斯文
Cyrillic                斯拉夫文
Deseret                 犹他州文
Devanagari              梵文
Ethiopic                衣索比亚文
Georgian                格鲁吉亚文
Glagolitic              格拉哥里文
Gothic                  哥特文
Greek                   希腊
Gujarati                古吉拉特文
Gurmukhi                果鲁穆奇文
Han                     汉文
Hangul                  韩文
Hanunoo                 哈鲁喏文
Hebrew                  希伯来文
Hiragana                平假名（日语）
Inherited               继承前一个字符的脚本
Kannada                 坎那达文
Katakana                片假名（日语）
Kayah_Li                克耶字母
Kharoshthi              卡罗须提文
Khmer                   高棉文
Lao                     老挝文
Latin                   拉丁文
Lepcha                  雷布查文
Limbu                   林布文
Linear_B                B类线形文字（古希腊）
Lycian                  利西亚文
Lydian                  吕底亚文
Malayalam               马拉雅拉姆文
Mongolian               蒙古文
Myanmar                 缅甸文
New_Tai_Lue             新傣仂文
Nko                     Nko文
Ogham                   欧甘文
Ol_Chiki                桑塔利文
Old_Italic              古意大利文
Old_Persian             古波斯文
Oriya                   奥里亚文
Osmanya                 奥斯曼亚文
Phags_Pa                八思巴文
Phoenician              腓尼基文
Rejang                  拉让文
Runic                   古代北欧文字
Saurashtra              索拉什特拉文（印度县城）
Shavian                 萧伯纳文
Sinhala                 僧伽罗文
Sundanese               巽他文
Syloti_Nagri            锡尔赫特文
Syriac                  叙利亚文
Tagalog                 塔加拉文
Tagbanwa                塔格巴努亚文
Tai_Le                  德宏傣文
Tamil                   泰米尔文
Telugu                  泰卢固文
Thaana                  塔安那文
Thai                    泰文
Tibetan                 藏文
Tifinagh                提非纳文
Ugaritic                乌加里特文
Vai                     瓦伊文
Yi                      彝文
【注意】

　　对于 [a-z] 这样的正则表达式，如果要在 [] 中匹配 - ，可以将 - 放在 [] 的开头或结尾，例如 [-a-z] 或 [a-z-]

　　可以在 [] 中使用转义字符：\f、\t、\n、\r、\v、\377、\xFF、\x{10FFFF}、\、^、\$、.、*、+、\?、{、}、(、)、[、]、\|（具体含义见上面的说明）

　　如果在正则表达式中使用了分组，则在执行正则替换的时候，“替换内容”中可以使用 $1、${1}、$name、${name} 这样的“分组引用符”获取相应的分组内容。其中 $0 代表整个匹配项，$1 代表第 1 个分组，$2 代表第 2 个分组，……。

　　如果“分组引用符”是 $name 的形式，则在解析的时候，name 是取尽可能长的字符串，比如：$1x 相当于 ${1x}，而不是${1}x，再比如：$10 相当于 ${10}，而不是 ${1}0。

　　由于 $ 字符会被转义，所以要在“替换内容”中使用 $ 字符，可以用 \$ 代替。

　　上面介绍的正则表达式语法是“Perl 语法”，除了“Perl 语法”外，Go 语言中还有另一种“POSIX 语法”，“POSIX 语法”除了不能使用“Perl 类”之外，其它都一样。

============================================================

示例：

func main() {

text := `Hello 世界！123 Go.`

// 查找连续的小写字母
reg := regexp.MustCompile(`[a-z]+`)
fmt.Printf("%q\n", reg.FindAllString(text, -1))
// ["ello" "o"]

// 查找连续的非小写字母
reg = regexp.MustCompile(`[^a-z]+`)
fmt.Printf("%q\n", reg.FindAllString(text, -1))
// ["H" " 世界！123 G" "."]

// 查找连续的单词字母
reg = regexp.MustCompile(`[\w]+`)
fmt.Printf("%q\n", reg.FindAllString(text, -1))
// ["Hello" "123" "Go"]

// 查找连续的非单词字母、非空白字符
reg = regexp.MustCompile(`[^\w\s]+`)
fmt.Printf("%q\n", reg.FindAllString(text, -1))
// ["世界！" "."]

// 查找连续的大写字母
reg = regexp.MustCompile(`[[:upper:]]+`)
fmt.Printf("%q\n", reg.FindAllString(text, -1))
// ["H" "G"]

// 查找连续的非 ASCII 字符
reg = regexp.MustCompile(`[[:^ascii:]]+`)
fmt.Printf("%q\n", reg.FindAllString(text, -1))
// ["世界！"]

// 查找连续的标点符号
reg = regexp.MustCompile(`[\pP]+`)
fmt.Printf("%q\n", reg.FindAllString(text, -1))
// ["！" "."]

// 查找连续的非标点符号字符
reg = regexp.MustCompile(`[\PP]+`)
fmt.Printf("%q\n", reg.FindAllString(text, -1))
// ["Hello 世界" "123 Go"]

// 查找连续的汉字
reg = regexp.MustCompile(`[\p{Han}]+`)
fmt.Printf("%q\n", reg.FindAllString(text, -1))
// ["世界"]

// 查找连续的非汉字字符
reg = regexp.MustCompile(`[\P{Han}]+`)
fmt.Printf("%q\n", reg.FindAllString(text, -1))
// ["Hello " "！123 Go."]

// 查找 Hello 或 Go
reg = regexp.MustCompile(`Hello|Go`)
fmt.Printf("%q\n", reg.FindAllString(text, -1))
// ["Hello" "Go"]

// 查找行首以 H 开头，以空格结尾的字符串
reg = regexp.MustCompile(`^H.*\s`)
fmt.Printf("%q\n", reg.FindAllString(text, -1))
// ["Hello 世界！123 "]

// 查找行首以 H 开头，以空白结尾的字符串（非贪婪模式）
reg = regexp.MustCompile(`(?U)^H.*\s`)
fmt.Printf("%q\n", reg.FindAllString(text, -1))
// ["Hello "]

// 查找以 hello 开头（忽略大小写），以 Go 结尾的字符串
reg = regexp.MustCompile(`(?i:^hello).*Go`)
fmt.Printf("%q\n", reg.FindAllString(text, -1))
// ["Hello 世界！123 Go"]

// 查找 Go.
reg = regexp.MustCompile(`\QGo.\E`)
fmt.Printf("%q\n", reg.FindAllString(text, -1))
// ["Go."]

// 查找从行首开始，以空格结尾的字符串（非贪婪模式）
reg = regexp.MustCompile(`(?U)^.* `)
fmt.Printf("%q\n", reg.FindAllString(text, -1))
// ["Hello "]

// 查找以空格开头，到行尾结束，中间不包含空格字符串
reg = regexp.MustCompile(` [^ ]*$`)
fmt.Printf("%q\n", reg.FindAllString(text, -1))
// [" Go."]

// 查找“单词边界”之间的字符串
reg = regexp.MustCompile(`(?U)\b.+\b`)
fmt.Printf("%q\n", reg.FindAllString(text, -1))
// ["Hello" " 世界！" "123" " " "Go"]

// 查找连续 1 次到 4 次的非空格字符，并以 o 结尾的字符串
reg = regexp.MustCompile(`[^ ]{1,4}o`)
fmt.Printf("%q\n", reg.FindAllString(text, -1))
// ["Hello" "Go"]

// 查找 Hello 或 Go
reg = regexp.MustCompile(`(?:Hell|G)o`)
fmt.Printf("%q\n", reg.FindAllString(text, -1))
// ["Hello" "Go"]

// 查找 Hello 或 Go，替换为 Hellooo、Gooo
reg = regexp.MustCompile(`(?PHell|G)o`)
fmt.Printf("%q\n", reg.ReplaceAllString(text, "${n}ooo"))
// "Hellooo 世界！123 Gooo."

// 交换 Hello 和 Go
reg = regexp.MustCompile(`(Hello)(.*)(Go)`)
fmt.Printf("%q\n", reg.ReplaceAllString(text, "$3$2$1"))
// "Go 世界！123 Hello."

// 特殊字符的查找
reg = regexp.MustCompile("[\\f\\t\\n\\r\\v\\123\\x7F\\x{10FFFF}\\\\\\^\\$\\.\\*\\+\\?\\{\\}\\(\\)\\[\\]\\|]")
fmt.Printf("%q\n", reg.ReplaceAllString("\f\t\n\r\v\123\x7F\U0010FFFF\\^$.*+?{}()[]|", "-"))
// "----------------------"
}
