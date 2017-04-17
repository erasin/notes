
使用 -x 来检查错误
```
GOOS=windows GOARCH=amd64 go build -x
```

使用 cgo 环境， 使用 Mingw-w64 交叉编译 (brew 安装)
```
CGO_ENABLED=1 GOOS=window64 GOARCH=amd64 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ go build -x
```

#### mac 下 brew 安装 Mingw-w64

```
brew tap cosmo0920/mingw_w64
brew install 
```
