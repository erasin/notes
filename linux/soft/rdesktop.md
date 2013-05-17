# rdesktop

远程链接windows服务器

## 语法
	rdesktop [options] server[:port]

参数 

   -u: 用户名 username 
   -d: 服务器IP或域名 domain
   -s: shell
   -c: working directory
   -p: 密码 password (- to prompt)
   -n: client hostname
   -k: 键盘布局 keyboard layout on server (en-us, de, sv, etc.)
   -g: 分辨率 desktop geometry (WxH)
   -f: 全屏模式 full-screen mode
   -b: force bitmap updates
   -L: local codepage
   -A: enable SeamlessRDP mode
   -B: use BackingStore of X-server (if available)
   -e: disable encryption (French TS)
   -E: disable encryption from client to server
   -M: do not map logical mouse buttons to physical
   -m: do not send motion events
   -C: use private colour map
   -D: hide window manager decorations
   -K: keep window manager key bindings
   -S: caption button size (single application mode)
   -T: window title
   -N: enable numlock syncronization
   -X: embed into another window with a given id.
   -a: 色彩位数 connection colour depth
   -z: enable rdp compression
   -x: RDP5 experience (m[odem 28.8], b[roadband], l[an] or hex nr.)
   -P: use persistent bitmap caching
   -r: 映射硬盘 enable specified device redirection (this flag can be repeated)
         '-r comport:COM1=/dev/ttyS0': enable serial redirection of /dev/ttyS0 to COM1
             or      COM1=/dev/ttyS0,COM2=/dev/ttyS1
         '-r disk:floppy=/mnt/floppy': enable redirection of /mnt/floppy to 'floppy' share
             or   'floppy=/mnt/floppy,cdrom=/mnt/cdrom'
         '-r clientname=<client name>': Set the client name displayed
             for redirected disks
         '-r lptport:LPT1=/dev/lp0': enable parallel redirection of /dev/lp0 to LPT1
             or      LPT1=/dev/lp0,LPT2=/dev/lp1
         '-r printer:mydeskjet': enable printer redirection
             or      mydeskjet="HP LaserJet IIIP" to enter server driver as well
         '-r sound:[local[:driver[:device]]|off|remote]': enable sound redirection
                     remote would leave sound on server
                     available drivers for 'local':
                     alsa:	ALSA output driver, default device: default
                     oss:	OSS output driver, default device: /dev/dsp or $AUDIODEV
                     libao:	libao output driver, default device: system dependent
         '-r clipboard:[off|PRIMARYCLIPBOARD|CLIPBOARD]': enable clipboard
                      redirection.
                      'PRIMARYCLIPBOARD' looks at both PRIMARY and CLIPBOARD
                      when sending data to server.
                      'CLIPBOARD' looks at only CLIPBOARD.
   -0: attach to console
   -4: use RDP version 4
   -5: use RDP version 5 (default)

## 实例

服务器 demo.com or IP
用户： theusername
密码： thepassword
窗口大小: 1200x800
挂载： /home/disk

	rdesktop -d 24 demo.com -utheusername -pthepassword -g 1200x800 -r disk:mydisk=/home/disk




