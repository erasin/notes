#ls

##Base

用法：ls [选项]... [文件]...  
List information about the FILEs (the current directory by default).  
Sort entries alphabetically if none of -cftuvSUX nor --sort is specified.

长选项必须使用的参数对于短选项时也是必需使用的。
	  -a, --all			不隐藏任何以. 开始的项目
	  -A, --almost-all		列出除. 及.. 以外的任何项目
		  --author			与-l 同时使用时列出每个文件的作者
	  -b, --escape			以八进制溢出序列表示不可打印的字符
		  --block-size=SIZE      scale sizes by SIZE before printing them.  E.g.,
								   '--block-size=M' prints sizes in units of
								   1,048,576 bytes.  See SIZE format below.
	  -B, --ignore-backups       do not list implied entries ending with ~
	  -c                         with -lt: sort by, and show, ctime (time of last
								   modification of file status information)
								   with -l: show ctime and sort by name
								   otherwise: sort by ctime, newest first
	  -C                         list entries by columns
		  --color[=WHEN]         colorize the output.  WHEN defaults to 'always'
								   or can be 'never' or 'auto'.  More info below
	  -d, --directory            list directory entries instead of contents,
								   and do not dereference symbolic links
	  -D, --dired                generate output designed for Emacs' dired mode
	  -f                         do not sort, enable -aU, disable -ls --color
	  -F, --classify             append indicator (one of */=>@|) to entries
		  --file-type            likewise, except do not append '*'
		  --format=WORD          across -x, commas -m, horizontal -x, long -l,
								   single-column -1, verbose -l, vertical -C
		  --full-time            like -l --time-style=full-iso
	  -g				类似-l，但不列出所有者
		  --group-directories-first
				在文件前分组目录。此选项可与--sort 一起使用，
				但是一旦使用--sort=none (-U)将禁用分组
	  -G, --no-group		以一个长列表的形式，不输出组名
	  -h, --human-readable		与-l 一起，以易于阅读的格式输出文件大小
					(例如 1K 234M 2G)
		  --si			同上面类似，但是使用1000 为基底而非1024
	  -H, --dereference-command-line
					跟随命令行列出的符号链接
		  --dereference-command-line-symlink-to-dir
					跟随命令行列出的目录的符号链接
		  --hide=PATTERN         	隐藏符合PATTERN 模式的项目
					(-a 或 -A 将覆盖此选项)
		  --indicator-style=WORD  append indicator with style WORD to entry names:
								   none (default), slash (-p),
								   file-type (--file-type), classify (-F)
	  -i, --inode                print the index number of each file
	  -I, --ignore=PATTERN       do not list implied entries matching shell PATTERN
	  -k, --kibibytes            use 1024-byte blocks
	  -l				使用较长格式列出信息
	  -L, --dereference		当显示符号链接的文件信息时，显示符号链接所指示
					的对象而并非符号链接本身的信息
	  -m				所有项目以逗号分隔，并填满整行行宽
	  -n, --numeric-uid-gid		类似 -l，但列出UID 及GID 号
	  -N, --literal			输出未经处理的项目名称 (如不特别处理控制字符)
	  -o				类似 -l，但不列出有关组的信息
	  -p,  --indicator-style=slash	对目录加上表示符号"/"
	  -q, --hide-control-chars   print ? instead of non graphic characters
		  --show-control-chars   show non graphic characters as-is (default
								 unless program is 'ls' and output is a terminal)
	  -Q, --quote-name           enclose entry names in double quotes
		  --quoting-style=WORD   use quoting style WORD for entry names:
								   literal, locale, shell, shell-always, c, escape
	  -r, --reverse			逆序排列
	  -R, --recursive		递归显示子目录
	  -s, --size			以块数形式显示每个文件分配的尺寸
	  -S				根据文件大小排序
		  --sort=WORD		以下是可选用的WORD 和它们代表的相应选项：
					extension -X       status   -c
					none      -U       time     -t
					size      -S       atime    -u
					time      -t       access   -u
					version   -v       use      -u
		  --time=WORD		和-l 同时使用时显示WORD 所代表的时间而非修改时
					间：atime、access、use、ctime 或status；加上
					--sort=time 选项时会以指定时间作为排序关键字
		  --time-style=STYLE     with -l, show times using style STYLE:
								 full-iso, long-iso, iso, locale, +FORMAT.
								 FORMAT is interpreted like 'date'; if FORMAT is
								 FORMAT1<newline>FORMAT2, FORMAT1 applies to
								 non-recent files and FORMAT2 to recent files;
								 if STYLE is prefixed with 'posix-', STYLE
								 takes effect only outside the POSIX locale
	  -t                         sort by modification time, newest first
	  -T, --tabsize=COLS         assume tab stops at each COLS instead of 8
	  -u			同-lt 一起使用：按照访问时间排序并显示
				同-l一起使用：显示访问时间并按文件名排序
				其他：按照访问时间排序
	  -U			不进行排序；按照目录顺序列出项目
	  -v			在文本中进行数字(版本)的自然排序
	  -w, --width=COLS	自行指定萤幕宽度而不使用目前的数值
	  -x			逐行列出项目而不是逐栏列出
	  -X			根据扩展名排序
	  -1			每行只列出一个文件
		  --help		显示此帮助信息并退出
		  --version		显示版本信息并退出

SIZE is an integer and optional unit (example: 10M is 10*1024*1024).  Units are K, M, G, T, P, E, Z, Y (powers of 1024) or KB, MB, ... (powers of 1000).

使用色彩来区分文件类型的功能已被禁用，默认设置和 --color=never 同时禁用了它。  
使用 --color=auto 选项，ls 只在标准输出被连至终端时才生成颜色代码。  
LS_COLORS 环境变量可改变此设置，可使用 dircolors 命令来设置。  
