#  MacOSX / Linux 用的 Shell 改用 zsh
1. 安装zsh

Debian / Ubuntu Linux 需要安裝: apt-get install zsh

Mac 预设就有 zsh 了~

安装完 zsh 后, chsh -s /bin/zsh 即可.

NOTES:

chsh -s /bin/zsh # 设定为 default shell

相关设定: .zshenv, .zprofile, .zshrc, .zlogin

2. 安裝使用 oh-my-zsh
cd ~/

git clone git://github.com/robbyrussell/oh-my-zsh.git ~/.oh-my-zsh

cp ~/.zshrc ~/.zshrc.orig

cp ~/.oh-my-zsh/templates/zshrc.zsh-template ~/.zshrc

配置zsh时需要修改.zshrc文件。

//修改theme

#export ZSH_THEME="steeef"

export ZSH_THEME="afowler"

更多的themes在.oh-my-zsh/themes中

//修改插件

plugins=(git osx) # 啟用 git, osx 的 plugin

更多plugins可以參考~/.oh-my-zsh/plugins

3. 中文乱码问题
在终端下输入

vim ~/.zshrc

或者使用其他你喜欢的编辑器编辑~/.zshrc

在文件内容末端添加：

export LC_ALL=en_US.UTF-8
export LANG=en_US.UTF-8

接着重启一下终端，或者输入 source ~/.zshrc