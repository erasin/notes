# 编码范围

正则工具 ： [regexr.com](http://regexr.com/)

* <https://gist.github.com/shingchi/64c04e0dd2cbbfbc1350>

## GBK (GB2312/GB18030)

    x00-xff GBK双字节编码范围
    x20-x7f ASCII
    xa1-xff 中文
    x80-xff 中文

## UTF-8 (Unicode)

    u4e00-u9fa5 (中文)
    x3130-x318F (韩文)
    xAC00-xD7A3 (韩文)
    u0800-u4e00 (日文)
    uff21 – uff5a 英文全角 A-z 
    uff01 - uff09 美式键盘 1-9 上标字符  02 双引号 06 中文省略号……
    uff10 - uff19 全角数字  ０ – ９ 
    uff20 @ 

> 韩文是大于[u9fa5]的字符


正则例子（使用PHP）:

```php
    preg_replace(“/([x80-xff])/”,”",$str);    //GBK中匹配
    preg_replace(“/([u4e00-u9fa5])/”,”",$str);    //UTF8中匹配
```
