#VIM

## 保存选择的内容
	:start,end w file
	:start,end w > file
	:start,end w >> file
使用view模式进行选择后直接使用 :w file（同上） 也可以保存选取内容

## 批量修改 
    选择文件
        :args *.c
    执行命令
        :argdo set fileencoding=utf-8
    保存
        :wa

## 编码转换

    :set fenc=编码  

## 查询

使用 nohl  取消高亮
	
	:nohl
	:nohlsearch
