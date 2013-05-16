# ncdu 磁盘分析

ncdu（NCurses Disk Usage）是一个基于 ncurses 界面的磁盘占用分析工具。其特点是快速、简单、且容易使用。


`df` 查看硬盘使用情况

	df -h

`du` 查看当前文件夹的使用情况

	du -h | sort -h 
	# sort  5.9 向下 
	du -h | sort -gr 
	# 当前文件夹
	du -h  --summarize


