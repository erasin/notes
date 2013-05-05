#MPD CONFIG 配置

## 播放器配置 ~/.mpd/mpd.conf 文件

注意修改`[name]` 和 `[yourname]` 为自己的目录名称

>     music_directory		"/home/[home]/Music"
	playlist_directory		"/home/[home]/.mpd/playlists"
	db_file				"/home/[home]/.mpd/mpd.db"
	#log_file				"/var/lib/mpd/mpd.log"
	#error_file			"/var/lib/mpd/mpd.error"
	pid_file				"/home/[home]/.mpd/mpd.pid"
	state_file				"/home/[home]/.mpd/mpdstate"
	user					"[yourname]"
	bind_to_address		"127.0.0.1"
	filesystem_charset		"UTF-8"
	#id3v1_encoding		"gbk"
	#metadata_to_use		"artist,album,title,track,name,genre,date,composer,performer,disc"
	mixer_type			"software"
	audio_output {
		type                    	"pulse"
		name                   	"pulse audio"
	}

## 客户端

[mpd](https://wiki.archlinux.org/index.php/Music_Player_Daemon_(%E7%AE%80%E4%BD%93%E4%B8%AD%E6%96%87))推荐播放器 mpc + gmpc + mpdris2(for dbus) + gnome shell 音乐插件[下载](https://github.com/eonpatapon/gnome-shell-extensions-mediaplayer)"
