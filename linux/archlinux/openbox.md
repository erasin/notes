
#安装
openbox openbox-themes obmenu obkey 

    $ mkdir -p ~/.config/openbox
    $ cp /etc/xdg/openbox/{rc.xml,menu.xml,autostart,environment} ~/.config/openbox

修改`~/.xinitrc` 添加
    exec openbox-session

如果使用了 D-Bus 和类似的程序，用下面这行代替:
     exec ck-launch-session openbox-session

menumaker 菜单生成器
你可以通过运行以下命令来生成一个完整的菜单文件 `~/.config/openbox/menu.xml`

     $ mmaker -v OpenBox3
     $ mmaker -vf OpenBox3

gtk-theme-switch2 gtk2的thtme   
修改`~/.config/gtk-3.0/settings.ini` 
    [Settings]
    gtk-application-prefer-dark-theme=0
    gtk-theme-name = Adwaita 
    gtk-fallback-icon-theme = Faenza

dmenu 快速执行

然后把下面的内容加入 ~/.config/openbox/rc.xml 的 <keyboard> 段来开启用快捷键启动 dmenu 
    <keybind key="W-space">
     <action name="Execute">
       <execute>dmenu_run</execute>
     </action>
    </keybind>

feh 壁纸显示 
修改 `~/.config/openbox/autostart`
    feh --bg-scale /path/to/image.file

slim 登陆器

### tint2 面板
配置文件~/.config/tint2/tint2rc  
或者使用 tint2conf 来修改


安装 xcompmgr 为tint2添加透明效果

修改 `~/.config/openbox/autostart`

    # Launch Xcomppmgr and tint2 with openbox
    if which tint2 >/dev/null 2>&1; then
      (sleep 2 && xcompmgr) &
      (sleep 2 && tint2) &
    fi
    
xdotool 键设置
