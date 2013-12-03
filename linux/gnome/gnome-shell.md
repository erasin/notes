# gnome-shell

icon : [Moka](http://snwh.org/moka-icon-theme/#download)

## gnome-shell window 最大化隐藏标题栏
> gnome-shell window max hidden title bar

修改： `/usr/share/themes/Adwaita/metacity-1/metacity-theme-3.xml`

找到 `frame_geometry name="max"` 添加 `has_title="false"`。

	<!-- 找到max 添加 has_title -->
	<frame_geometry name="max" has_title="false" title_scale="medium" parent="normal" rounded_top_left="false" rounded_top_right="false">
	         <distance name="left_width" value="0" />
	         <distance name="right_width" value="0" />
	         <distance name="left_titlebar_edge" value="0"/>
	         <distance name="right_titlebar_edge" value="0"/>
	         <!-- 下两行 -->
	         <distance name="title_vertical_pad" value="0"/>
	         <border name="title_border" left="0" right="0" top="0" bottom="0"/>
	         <border name="button_border" left="0" right="0" top="0" bottom="0"/>
	         <distance name="bottom_height" value="0" />
	</frame_geometry>


重启 `Alt+F2` 运行 `r` 重启 shell

快捷键：

`ALT+F10` 来切换。 gnome 3.8 后 `Super+UP/Down` 来切换 ，或用 `Super+Space` 显示菜单。

## totem 

totem 的x264 问题， 请安装`gst-libav`

## extensions

<https://extensions.gnome.org/>


* [AlternateTab](https://extensions.gnome.org/extension/15/alternatetab/)  by gcampax
* [Applications Menu ](https://extensions.gnome.org/extension/6/applications-menu/) by gcampax
* [Dash to Dock](https://extensions.gnome.org/extension/307/dash-to-dock/)  by michele_g
* [kimpanel](https://extensions.gnome.org/extension/261/kimpanel/)  by csslayer
* [Media player indicator](https://extensions.gnome.org/extension/55/media-player-indicator/)  by eon
* [TopIcons ](https://extensions.gnome.org/extension/495/topicons/) by ag
* [User Themes](https://extensions.gnome.org/extension/19/user-themes/)  by gcampax
* [Weather](https://extensions.gnome.org/extension/613/weather/)  by Neroth

# fcitx

    yaourt -S gnome-shell-extension-kimpanel-git 

    $ gsettings set org.gnome.settings-daemon.plugins.keyboard active false
