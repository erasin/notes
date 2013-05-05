#MongoDB  分布式文档存储数据库MongoDB


__MongoDB__是一个介于关系数据库和非关系数据库之间的产品，是非关系数据库当中功能最丰富，最像关系数据库的。他支持的数据结构非常松散，是类似__json__的__bjson__格式，因此可以存储比较复杂的数据类型。Mongo最大的特点是他支持的查询语言非常强大，其语法有点类似于面向对象的查询语言，几乎可以实现类似关系数据库单表查询的绝大部分功能，而且还支持对数据建立索引。

它的特点是__高性能、易部署、易使用__，存储数据非常方便。主要功能特性有：

* 面向集合存储，易存储对象类型的数据。
* 模式自由。
* 支持动态查询。
* 支持完全索引，包含内部对象。
* 支持查询。
* 支持复制和故障恢复。
* 使用高效的二进制数据存储，包括大型对象（如视频等）。
* 自动处理碎片，以支持云计算层次的扩展性
* 支持RUBY，PYTHON，JAVA，C++，PHP等多种语言。
* 文件存储格式为BSON（一种JSON的扩展）
* 可通过网络访问

所谓“__面向集合__”（Collenction-Orented），意思是数据被分组存储在数据集中，被称为一个集合（Collenction)。每个 集合在数据库中都有一个唯一的标识名，并且可以包含无限数目的文档。集合的概念类似关系型数据库（RDBMS）里的表（table），不同的是它不需要定 义任何模式（schema)。

模式自由（schema-free)，意味着对于存储在mongodb数据库中的文件，我们不需要知道它的任何结构定义。如果需要的话，你完全可以把不同结构的文件存储在同一个数据库里。

存储在集合中的文档，被存储为__键-值__对的形式。键用于唯一标识一个文档，为字符串类型，而值则可以是各中复杂的文件类型。我们称这种存储形式为BSON（Binary Serialized dOcument Format）。

MongoDB服务端可运行在Linux、Windows或OS X平台，支持32位和64位应用，默认端口为27017。推荐运行在64位平台，因为MongoDB

在32位模式运行时支持的最大文件尺寸为2GB。

MongoDB把数据存储在文件中（默认路径为：/data/db），为提高效率使用内存映射文件进行管理。

##从mysql移植到MongoDB服务器

![mysql移植到MongoDB](http://photo2.bababian.com/upload5/20110904/176B2C9237C48064A8F31D7E877AB6BC.jpg)

mysql和Mogo的区别 ：[http://www.mongodb.org/display/DOCS/SQL+to+Mongo+Mapping+Chart]

[MongoDB首页][mongo]
[mongo]:http://www.mongodb.org/ "分布式文档存储数据库 MongoDB"
