# 重新安装MAC

官方：

您可以使用“磁盘工具”抹掉您的电脑，然后重新安装 Mac OS X。在您抹掉和重新安装前，备份重要文件。如果您在便携式电脑上进行安装，请确定您的电源适配器已连接并已接上电源。
重要： 若要重新安装 Mac OS X，需要连接到互联网。

选取苹果菜单 >“重新启动”，然后在电脑重新启动时按住 Command 键 (⌘) 和 R 键。
选择“磁盘工具”，然后点按“继续”。
从左侧的列表中选择您的启动磁盘，然后点按“抹掉”标签。
从“格式”弹出式菜单中，选择“Mac OS 扩展格式（日志式）”，键入磁盘的名称，然后点按“抹掉”。
在磁盘已被抹掉后，请选取“磁盘工具”>“退出磁盘工具”。
如果未连接到互联网，请从菜单栏（在屏幕右上角）右边角的“Wi-Fi”菜单中选取一个网络。
若要重新安装 Mac OS X，请点按“继续”，然后按照屏幕指示执行操作。 有关重新安装 Mac OS X 的更多信息，请参阅此帮助主题： 重新安装 Mac OS X
Mac OS X Lion 附带了内建的恢复磁盘，您可以用它来重新安装 Mac OS X、修复磁盘、从 Time Machine 备份中恢复等等。若要使用恢复磁盘，请在按住 Command (⌘) 键和 R 键的同时重新启动电脑。如果您想，您还可以创建外部恢复磁盘。有关更多信息，请参阅此 Apple 支持文章：

## 步骤

### 备份

备份 配置文件

```
~/.ssh/
~/.bashrc
~/.bash_profile
~/.vim/
~/.zshrc
~/.oh-my-zsh
~/.filezilla/
~/.aria2/
```

备份资料

```
~/Documents
~/Downloads
~/work
~/...

~/Library/Containers/com.tencent.qq/Data/Library/Application Support/QQ
~/Library/Application Support/Steam
```

###  格式系统重新安装

## 软件安装

* `oh-my-zsh` 

* `xcode`  命令行支持

* `brew` 

  ```
  ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
  ```

* brew install 

  ```
  brew install cask
  brew install aria2 ffmpeg mpv ncdu iterm2-beta tree ranger mp3splt freetype
  brew install go lua sqlite node
  brew install sdl2 sdl2_gfx sdl2_image sdl2_mixer sdl2_ttf 
  brew cask install macvim google-chrome typora atom java limechat filezilla lantern virtualbox flux calibre
  ```

* 桌面应用

  * sketchBook
  * Nutstore
  * NeteaseMusic
  * sequel pro
  * YouDaoDict
  * 富途牛牛
  * movist
  * google jpanese ime
  * twitter
  * QQ
  * 钉钉
  * wechat
  * xmind
  * paw
  * Microsoft remote desktop
  * [截图](https://itunes.apple.com/cn/app/jie-tu-jietu/id1059334054?mt=12&ign-mpt=uo%3D4)
  * steam and battle.net

