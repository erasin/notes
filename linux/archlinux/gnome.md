#gnome


## 消除gtk3的标题栏

使用 `gnome-tweak-tool` 添加最大化按钮

To remove maximized windows titlebar in GNOME Shell, open "/usr/share/themes/Adwaita/metacity-1/metacity-theme-3.xml" (firstly make a backup of this file!) as root and search for frame_geometry name="max" and edit it so that it looks like this:

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


## totem 

totem 的x264 问题， 请安装`gst-libav`
