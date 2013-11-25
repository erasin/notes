# 分类

IO类
:   io - 基本的IO接口
    io/ioutil - 方便的IO操作函数集
    fmt - 格式化IO
    bufio - 缓存IO

OS
:   os - 平台无关的操作系统功能实现
    path - 操作路径
    path/filepath - 操作文件名路径
    flag  - 参数解析

string
:    strings - 字符串操作
    strconv - 基本数据类型和字符串之间转换
    regexp - 正则表达式
    unicode - Unicode码点、UTF-8/16编码 数据结构与算法
    container - 容器数据类型：heap、list和ring
    bytes - byte slice 便利操作
    index/suffixarray - 后缀数组实现子字符串查询
    sort - 排序算法

math
:    math - 基本数学函数
    math/big - 大数实现
    math/cmplx - 复数基本函数操作
    math/rand - 伪随机数生成器

time - 日期和时间操作、显示

log - 日志

database
:    database/sql - SQL/SQL-Like 数据库操作接口
    encoding/json - json 解析
    encoding/xml - xml 解析
    encoding/gob - golang 自定义二进制格式
    encoding/base64 - 
    cvs - 逗号分隔值文件

cmopress
:    compress/zlib - gnu zlib压缩
    compress/gzip - 读写gnu zip文件
    compress/bzip2 - bzip2压缩
    archive/tar - tar归档访问
    archive/zip - zip归档访问

net
:   net
    net/http - http 服务
    html/template - html 文档
    net/url - url 解析
    net/smtp - 邮件服务
    net/rpc - RPC 服务
    net/rpc/jsonrpc - JSON RPC
    code.google.com/p/go.net/websocket - google提供 socket库