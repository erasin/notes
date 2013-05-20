# gnome-shell


## gnome-shell window 最大化隐藏标题栏
> gnome-shell window max hidden title bar

修改： `/usr/share/themes/Adwaita/metacity-1/metacity-theme-3.xml`

找到 `frame_geometry name="max"` 添加 `has_title="false"`。

	<frame_geometry name="max" has_title="false" title_scale="medium" parent="normal" rounded_top_left="false" rounded_top_right="false">
			 <distance name="left_width" value="0" />
			 <distance name="right_width" value="0" />
			 <distance name="left_titlebar_edge" value="0"/>
			 <distance name="right_titlebar_edge" value="0"/>
			 <distance name="title_vertical_pad" value="0"/>
			 <border name="title_border" left="0" right="0" top="0" bottom="0"/>
			 <border name="button_border" left="0" right="0" top="0" bottom="0"/>
			 <distance name="bottom_height" value="0" />
	</frame_geometry>


重启 `Alt+F2` 运行 `r` 重启 shell

快捷键：

`ALT+F10` 来切换。 gnome 3.8 后 `Super+UP/Down` 来切换 ，或用 `Super+Space` 显示菜单。
