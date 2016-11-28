# /var/log




* /log/messages — 包括整体系统信息，其中也包含系统启动期间的日志。此外，mail，cron，daemon，kern和auth等内容也记录在var/log/messages日志中。
* /log/dmesg — 包含内核缓冲信息（kernel ring buffer）。在系统启动时，会在屏幕上显示许多与硬件有关的信息。可以用dmesg查看它们。
* /log/auth.log — 包含系统授权信息，包括用户登录和使用的权限机制等。
* /log/boot.log — 包含系统启动时的日志。
* /log/daemon.log — 包含各种系统后台守护进程日志信息。
* /log/dpkg.log – 包括安装或dpkg命令清除软件包的日志。
* /log/kern.log – 包含内核产生的日志，有助于在定制内核时解决问题。
* /log/lastlog — 记录所有用户的最近信息。这不是一个ASCII文件，因此需要用lastlog命令查看内容。
* /log/maillog * /log/mail.log — 包含来着系统运行电子邮件服务器的日志信息。例如，sendmail日志信息就全部送到这个文件中。
* /log/user.log — 记录所有等级用户信息的日志。
* /log/Xorg.x.log — 来自X的日志信息。
* /log/alternatives.log – 更新替代信息都记录在这个文件中。
* /log/btmp – 记录所有失败登录信息。使用last命令可以查看btmp文件。例如，”last -f * /log/btmp | more“。
* /log/cups — 涉及所有打印信息的日志。
* /log/anaconda.log — 在安装Linux时，所有安装信息都储存在这个文件中。
* /log/yum.log — 包含使用yum安装的软件包信息。
* /log/cron — 每当cron进程开始一个工作时，就会将相关信息记录在这个文件中。
* /log/secure — 包含验证和授权方面信息。例如，sshd会将所有信息记录（其中包括失败登录）在这里。
* /log/wtmp或* /log/utmp — 包含登录信息。使用wtmp可以找出谁正在登陆进入系统，谁使用命令显示这个文件或信息等。
* /log/faillog – 包含用户登录失败信息。此外，错误登录命令也会记录在本文件中。
  除了上述Log文件以外， * /log还基于系统的具体应用包含以下一些子目录：
* /log/httpd/或* /log/apache2 — 包含服务器access_log和error_log信息。
* /log/lighttpd/ — 包含light HTTPD的access_log和error_log。
* /log/mail/ –  这个子目录包含邮件服务器的额外日志。
* /log/prelink/ — 包含.so文件被prelink修改的信息。
* /log/audit/ — 包含被 Linux audit daemon储存的信息。
* /log/samba/ – 包含由samba存储的信息。
* /log/sa/ — 包含每日由sysstat软件包收集的sar文件。
* /log/sssd/ – 用于守护进程安全服务。
  除了手动存档和清除这些日志文件以外，还可以使用logrotate在文件达到一定大小后自动删除。可以尝试用vi，tail，grep和less等命令查看这些日志文件。
