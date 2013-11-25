# OS

操作系统接口


## 常用

* 接口
	* FileInfo
* 函数
* 结构

## 参数 Args

`var Args[]string` 程序开始运行时，获得的参数组成的部分为 `Args`, 默认有1个参数。 当运行时输出参数则从 `Args[1]` 开始。

> 建议使用 `flag.Args()` 来处理参数.

```go
fmt.Printf("%T , %v\n", os.Args, os.Args)
```

执行

``` bash
go run args.go canshu
```

结果为 

	[]string , [/tmp/go-build457611479/command-line-arguments/_obj/exe/args canshu]

常用 `len(os.Args) < 2 ` 来测试是否有参数输入。


## 文件操作



### 权限




## File

可以返回



## FileInfo interface 
可以共过 File.Stat() 或 File.Readdir 来获取 
当前File实例的 FileInfo.



另外两个函数也提供了FileInfo 的返回值.

	Lstat(name string)(fi FileInfo,error error)

和