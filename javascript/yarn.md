# yarn

yarn 为 npm 的替代工具。

### yarn 换 [taobao 源](https://npm.taobao.org/)

```
    # 查看当前源
    $ yarn config get registry 
    https://registry.yarnpkg.com
    $ yarn config set registry 'https://registry.npm.taobao.org'  
    yarn config v0.20.0
    success Set "registry" to "https://registry.npm.taobao.org".
    ✨  Done in 0.07s.
    # 可以通过查看配置文件检查
    $ cat ~/.yarnrc 
```
