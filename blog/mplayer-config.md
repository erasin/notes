#mplayer 配置

mplayer 配置文件 ~/.mplayer/config  
字幕以及其他

    [default]
    #-----------------------------------------------------------
    # 字幕
    #-----------------------------------------------------------
    subcp=&quot;cp936&quot;
    #subcp=enca:zh:ucs-2
    #font=&quot;Sans&quot;
    font=&quot;/usr/share/fonts/wenquanyi/wqy-zenhei/wqy-zenhei.ttc&quot;
    subfont-autoscale=1
    subfont-osd-scale=4
    subfont-text-scale=4
    subalign=2
    subpos=96
    spuaa=20
    subfont-encoding=&quot;unicode&quot;
    #unicode=1
    utf8=1
    ass=1    
    #---------------------------------------------------------
    # other 
    #---------------------------------------------------------
    # ao
    #flip=yes    			#倒立
    #vf=eq2=1.0:-0.8		#负片    
    # terminal 
    msgcolor=1				#终端色彩显示
    msgmodule=1				#终端显示模块名称    
    # 播放器前端选项
    use-filename-title=1	#显示标题
    #noborder=1				#无边框
    #shuffle=1				#随机播放
    softvol-max=200			#调整音量 200为2倍
    #volstep=5
    #border=0				#边框窗口装饰
    geometry=50%:50%		#摆放在中间
    #ontop=1				#前置    
    [gnome-mplayer]
    msglevel=all=5
