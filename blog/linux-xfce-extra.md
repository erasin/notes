# Linux Xfce 的使用

## 软件包

xfce4 libxfce4ui libxfcegui4 xfce4-battery-plugin xfce4-clipman-plugin xfce4-cpufreq-plugin xfce4-cpugraph-plugin xfce4-datetime-plugin xfce4-dict xfce4-diskperf-plugin xfce4-eyes-plugin xfce4-fsguard-plugin xfce4-genmon-plugin xfce4-mailwatch-plugin xfce4-mount-plugin xfce4-mpc-plugin xfce4-netload-plugin xfce4-notes-plugin xfce4-power-manager xfce4-quicklauncher-plugin xfce4-sensors-plugin xfce4-smartbookmark-plugin xfce4-systemload-plugin xfce4-time-out-plugin xfce4-timer-plugin xfce4-verve-plugin xfce4-wavelan-plugin xfce4-weather-plugin xfce4-xkb-plugin xfce4-notifyd

## 为xinit启动修改 .xinitrc

	export GTK_IM_MODULE=ibus
	export QT_IM_MODULE=xim
	export XMODIFIERS='@im=ibus'
	export XIM_PROGRAM=&quot;ibus-daemon&quot;
	export XIM_ARGS=&quot;--daemonize --xim&quot;
	export LANG=zh_CN.utf8;
	export LANGUAGE=zh_CN:zh:en;
	exec ck-launch-session dbus-launch --exit-with-session startxfce4
