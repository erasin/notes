#CMUS

## 基本命令
* :set softvol=true - 打开软音量控制
* :add /path/to/music/dir - 添加音乐库目录
* :clear -清除播放列表
* :save playlist.pls - 保存当前播放列表
* :load playlist.pls -加载播放列表
* :colorscheme xterm-white - 修改配色方案

## 操作键

|key|command|note|
|---|-------|----|
|**q**| **quit -i**| 退出|
|**^C**| **echo Type :quit<enter> to exit cmus.**| 退出（^就是ctrl的意思）|
|**b**| **player-next**| 下一首|
|**c**| **player-pause**| 暂停 [快速记忆： c == continue]|
|**x**| **player-play**| 播放|
|**z**| **player-prev**| 前一首|
|**v**| **player-stop**| 停止|
|**^L**| **refresh**| 刷新|
|**n**| **search-next**| 查找下一个|
|**N**| **search-prev**| 查找上一个|
|**.**| **seek +1m**| 快进|
|**l,**| **right seek +5**| 快快进|
|**,**| **seek -1m**| 快退|
|**h,**| **left seek -5**| 快快退|
|**m**| **toggle `aaa_mode`**| 修改`aaa_mode`模式 [快速记忆： m == mode ]（all from library |artist from library | album from library）|
|**C**| **toggle continue**| 持续播放，也就是说这首播放完了继续播放下一首|
|**M**| **toggle `play_library`**| 切换到playlist （问题：如果我M到了playlist是不是R循环播放是对于整个playlist的？）|
|**o**| **toggle `play_sorted`**| 切换成all from sorted library，具体功能我不懂，望高手指点一二|
|**r**| **toggle repeat**| 循环[ 快速记忆： r == repeat ] （问题：循环针对的对象是不是`aaa_mode`？）|
|**^R**| **toggle `repeat_current`**| 对当前循环|
|**t**| **toggle `show_remaining_time`**| 显示剩余时间[ 快速记忆: t == time]（默认是显示播放时间）|
|**s**| **toggle shuffle**| 支持拖拽（？）|
|**F**| **push filter<space>**| 不清楚，只看到命令提示行显示了:fliter ，估计是提示筛选神马的|
|**L**| **push live-filter<space>**| 我也不清楚，命令显示为：:live-filter， 看着想即时筛选？|
|**u**| **update-cache**| 更新缓存（我按了没什么效果）|
|**1**| **view tree**| 主界面|
|**2**| **view sorted**||
|**3**| **view playlist**| 这个估计写错了，应该是view browser|
|**4**| **view queue**| 显示Queue （我也不知道是啥东西）|
|**5**| **view browser**| 这个应该是view playlist，显示播放列表|
|**6**| **view filters**| 显示筛选机制（？）|
|**7**| **view settings**| 显示快捷键|
|**!**| **push shell<space>**| 不明白，因为按下！就到view settings去了|
|**]**| **vol +0 +1**| 音量控制，左声道不变，右声道+1|
|**[**| **vol +1 +0**| 音量控制，左声道+1，右声道+0|
|**+,**| **= vol +10%**| 声音变大10%|
|**-**| **vol -10%**| 声音减小10%|
|**}**| **vol -0 -1**| 音量控制，左声道不变，右声道-1|
|**{**| **vol -1 -0**| 音量控制，左声道+0，右声道+1|
|**enter**| **win-activate**| 不懂解释的意思，但是按下回车就会开始放歌。。。|
|**E**| **win-add-Q**| 传说可以跟别人聊天（？）但是我没试出来|
|**a**| **win-add-l**| 不懂|
|**y**| **win-add-p ms**|是前一首的意思|
|**e**| **win-add-q**| 光标向下|
|**G,**| **end win-bottom**| 光标到整个list的最后|
|**down,**| **j win-down**| 光标向下（vim的快捷键）|
|**p**| **win-mv-after**| 不懂|
|**P**| **win-mv-before color=#FF00FF**| color|
|**tab**| **win-next**| 切换窗口|
|**^F,**| **`page_down` win-page-down**| 光标跳到一屏中的最后|
|**^B,**| **`page_up` win-page-up**| 光标跳到一屏中的嘴上|
|**D,**| **delete win-remove**| 删除|
|**i**| **win-sel-cur**| 打开当前光标所在的track|
|**space**| **win-toggle**| 打开属性列表的子列表|
|**g,**| **home win-top**| 光标跳到最上边|
|**k,**| **up win-up**| 光标向上一行|
