# PHP 坑

## json_encode JSON输出

对于json数据函数，对数组中字符串中的空格等字符使用 `urlencode`处理下再 json_encode
