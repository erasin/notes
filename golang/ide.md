# ide for golang 


vscode 

gocode: 			go get -u -v github.com/nsf/gocode
godef: 				go get -u -v github.com/rogpeppe/godef
golint: 			go get -u -v github.com/golang/lint/golint
go-find-references: go get -u -v github.com/lukehoban/go-find-references
go-outline: 		go get -u -v github.com/lukehoban/go-outline
goreturns: 			go get -u -v sourcegraph.com/sqs/goreturns
gorename: 			go get -u -v golang.org/x/tools/cmd/gorename




# atom 
http://www.open-open.com/lib/view/open1425047100812.html
go 环境安装
这一部分是最重要的，如果没有它，每次build的时候出现 too many errors 心里真的是非常难过的。

环境配置：（golint,gooracle,mercurial）

安装mercurial: brew install mercurial

这个东西是用来做版本管理的，也是下载代码的工具类似git，貌似google的项目用的挺多的。

安装golint：


$ go get github.com/golang/lint$
go install github.com/golang/lint
安装gooracle


go get golang.org/x/tools/cmd/oracle
安装goimport


go get golang.org/x/tools/cmd/goimports
安装gocode


go get -u github.com/nsf/gocode
安装 godef


go get -v code.google.com/p/rog-go/exp/cmd/godef
go install -v code.google.com/p/rog-go/exp/cmd/godef

github.com/sqs/goreturns

gocode 提供代码补全

godef 代码跳转

gofmt 自动代码整理

golint 代码语法检查

goimports 自动整理imports

oracle 代码callgraph查询（plugin中还在todolist中，但是不配置一直报错。实在烦。）
