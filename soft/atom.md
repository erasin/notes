# atom

- cmd+o
- cmd+t,cmd+p 查找打开当前项目的文件
- cmd+b 切换已经打开的标签
- cmd+f 当前文件查询
- cmd+shif+f 项目文件内全局查询

## plugins

- atom-beautify 代码美化，php美化需要配置php路径
- color-picker 取色器
- autocomplete-paths  系统文件路径自动提示
- vim-mode vim模式
- ex-mode  vim模式扩充
- file-icons 文件图标
- git-plus  git工具
- atom-alignment  自动对齐工具
- emmet html快速编辑
- go-plus go语言插件
- merge-conflicts，在 Atom 里面处理合并产生冲突的文件。
- Remote-FTP@0.6.2  带有树形视图的FTP，保存时自动上传，基于项目(文件夹)的配置
- atom-beautify@0.24.1 代码美化，php美化需要配置php路径
- atom-color-highlight@3.0.9 显示颜色
- atom-ctags@3.2.0 添加了一个快捷键Ctrl+R，让你拥有类似于Ctrl+T (项目内文件快速跳转) 的体验。
- atom-html-preview@0.1.6 预览html，我使用时发现没有css样式
- autocomplete-paths@1.0.2 路径提示，实测从系统根目录开始提示，所以无用
- autocomplete-plus@2.7.1 自动提示接口?
- autocomplete-snippets@1.3.0 snippet提示
- color-picker@1.7.0 取色？
- linter@0.12.1 语法错误提示接口
- linter-php@0.0.12 php文件中的语法错误提示
- merge-conflicts@1.3.1  ？
- minimap@4.7.6 小地图
- open-last-project@0.4.1 启动atom自动打开上一次关闭时的项目
- pretty-json@0.3.3 json美化
- project-manager@1.15.5 通过配置可以快捷的打开不同项目(文件夹)
- random-tips@0.4.0 状态栏有一些关于编码的经验
- remote-edit@1.7.2 个人觉得Remote-FTP更加直观方便
- script@2.19.0 直接运行代码，需要配置完整的路径
- seti-syntax@0.4.0 UI
- seti-ui@0.6.3 UI
- sloc@0.1.3 显示代码行数
- symbols-tree-view@0.9.2  代码大纲视图
- todo-show@0.8.0 查找项目(项目)所有的带有todo的注释
- vim-mode@0.45.0 部分vim的快捷键，但我用不习惯
- livereload 相当重要
- filetype-color@0.1.4
- linter-coffeelint@0.3.2
- linter-jshint@0.1.2
- minimap@4.10.0
- monokai@0.14.0
- remote-edit@1.8.2
- remote-sync@3.1.2
- script@2.25.2
- symbols-tree-view@0.9.3
- unity-ui@2.0.11
- valign@1.0.2
- vim-mode@0.54.0
- project-manager 项目管理 跳转管理
    - 配合tree view （ cmd - \） 打开关闭project file
    - 搜索项目文件 （ cmd -t ）
    - 打开项目列表 （ ctrl-cmd-p ）
- atom-runner 修改源代码中的go run运行为go: 'sh [绝对路径到你的全局run脚上] '+atom.project.getRepo().project.path 

## go环境

[go-plus](https://github.com/joefitzgerald/go-plus)
Go (Golang) Package For Atom That Adds Autocomplete, Formatting, Syntax Checking, Linting and Vetting

```
apm install go-plus
```

go plugin : gocode ,gofmt,goimports,goreturns,vet, gooracle, mercurial , golint

- gocode 提供代码补全
- godef 代码跳转
- gofmt 自动代码整理
- golint 代码语法检查
- goimports 自动整理imports
- oracle 代码callgraph查询（plugin中还在todolist中，但是不配置一直报错。实在烦。）

```sh
go get github.com/nsf/gocode
go get golang.org/x/tools/cmd/...
go get github.com/golang/lint/golint


brew install mercurial

```




[godef](https://github.com/litgh/atom-godef)

先安装

```sh
go get -v code.google.com/p/rog-go/exp/cmd/godef
```
