# SSH 服务


**SSH** 是 Secure Shell 的缩写，是建立在传输层基础上的安全协议，它本身属于应用层，同时可以为应用层提供安全传输服务。

SSH全称Secure SHell，顾名思义就是非常安全的shell的意思，SSH协议是IETF（Internet Engineering Task Force）的Network Working Group所制定的一种协议。SSH的主要目的是用来取代传统的telnet和R系列命令（rlogin,rsh,rexec等）远程登陆和远程执行命令的工具，实现对远程登陆和远程执行命令加密。防止由于网络监听而出现的密码泄漏，对系统构成威胁。

在出现SSH之前，系统管理员需要登入远程服务器执行系统管理任务时，都是用telnet来实现的，telnet协议采用明文密码传送，在传送过程中对数据也不加密，很容易被不怀好意的人在网络上监听到密码。同样，在SSH工具出现之前R系列命令也很流行（由于这些命令都以字母r开头，故把这些命令合称为R系列命令R是remote的意思），比如rexec是用来执行远程服务器上的命令的，和telnet的区别是telnet需要先登陆远程服务器再实行相关的命令，而R系列命令可以把登陆和执行命令并登出系统的操作整合在一起。这样就不需要为在远程服务器上执行一个命令而特地登陆服务器了。

SSH是一种加密协议，不仅在登陆过程中对密码进行加密传送，而且对登陆后执行的命令的数据也进行加密，这样即使别人在网络上监听并截获了你的数据包，他也看不到其中的内容。OpenSSH已经是目前大多数linux和BSD操作系统（甚至cygwin）的标准组件，因此关于如何安装OpenSSH本文就不再叙述了，如果不出意外，你的系统上必定已经安装好了OpenSSH。

## 基本架构

SSH协议框架中最主要的部分是三个协议：

* 传输层协议（The Transport Layer Protocol）：传输层协议提供服务器认证，数据机密性，信息完整性等的支持。
* 用户认证协议（The User Authentication Protocol）：用户认证协议为服务器提供客户端的身份鉴别。
* 连接协议（The Connection Protocol）：连接协议将加密的信息隧道复用成若干个逻辑通道，提供给更高层的应用协议使用。

同时还有为许多高层的网络安全应用协议提供扩展的支持。

各种高层应用协议可以相对地独立于SSH基本体系之外，并依靠这个基本框架，通过连接协议使用SSH的安全机制。

## SSH的安全验证

在客户端来看，SSH提供两种级别的安全验证。
* 第一种级别（基于密码的安全验证），知道帐号和密码，就可以登录到远程主机，并且所有传输的数据都会被加密。但是，可能会有别的服务器在冒充真正的服务器，无法避免被“中间人”攻击。
* 第二种级别（基于密匙的安全验证），需要依靠密匙，也就是你必须为自己创建一对密匙，并把公有密匙放在需要访问的服务器上。客户端软件会向服务器发出请求，请求用你的密匙进行安全验证。服务器收到请求之后，先在你在该服务器的用户根目录下寻找你的公有密匙，然后把它和你发送过来的公有密匙进行比较。如果两个密匙一致，服务器就用公有密匙加密“质询”（challenge）并把它发送给客户端软件。从而避免被“中间人”攻击。

在服务器端，SSH也提供安全验证。 在第一种方案中，主机将自己的公用密钥分发给相关的客户端，客户端在访问主机时则使用该主机的公开密钥来加密数据，主机则使用自己的私有密钥来解密数据，从而实现主机密钥认证，确定客户端的可靠身份。 在第二种方案中，存在一个密钥认证中心，所有提供服务的主机都将自己的公开密钥提交给认证中心，而任何作为客户端的主机则只要保存一份认证中心的公开密钥就可以了。在这种模式下，客户端必须访问认证中心然后才能访问服务器主机。

## OPENSSH 

使用 apt-get , yum ,pacman 或者在 [OPENSSH官网](<http://openssh.org/> 'openssh') 来下载编译openssh。 

OpenSSH软件包包含以下命令：

* sshd ―― SSH服务端程序
* sftp-server ―― SFTP服务端程序（类似FTP但提供数据加密的一种协议）
* scp ―― 非交互式sftp-server的客户端，用来向服务器上传/下载文件
* sftp ―― 交互式sftp-server客户端，用法和ftp命令一样。
* slogin ――　ssh的别名
* ssh ――　SSH协议的客户端程序，用来登入远程系统或远程执行命令
* ssh-add ――    SSH代理相关程序，用来向SSH代理添加dsa　key
* ssh-agent ――    ssh代理程序
* ssh-keyscan ――　ssh　public key 生成器

## ssh

    ssh -l 远程服务器用户名 远程服务器ip地址 -p 远程服务器ssh端口（默认22）

在第一次使用SSH2协议向服务器建立连接 时，，SecureCRT首先会检查Server的公钥是否在本地数据库存放，如果没有，则不会把你的用户名、口令传输过去，它会将服务器端的公钥取回来 （可以直接从Server取，也可以从第三方获取），并提示：

    The host key database does not contain an entry for the
    hostname 172.16.200.244, which resolved to 172.16.200.244,
    port 22. If you have received this message more
    than once for 172.16.200.244, this may mean that 172.16.200.244
    is an “alias” which resolves to different hosts.
    It is recommended you verify your host key before accepting.
    Server’s host key fingerprint (MD5 hash):
    24:0f:36:5e:43:ad:f5:b8:1b:ae:ac:f7:9f:c2:c0:4c

当 你选择确认保存，则SecureCRT会把Server的公钥保存到本地公钥数据库，然后会重新让你输入用户名和密码，再次安全地登录服务器，因为这一次 会使用Server的公钥来加密用户名和口令。之后的登录和交互就会一直使用本地数据库保存的Server的公钥来加密传输。


SSH最常用的使用方式是代替telnet进行远程登陆。不同于telnet的密码登陆，SSH还同时支持Publickey、Keybord Interactive、GSSAPI等多种登入方式，不像telnet那样只有输入系统密码一种途径。目前最常用的登陆方式还是传统的Password方式和Publickey方式登陆。下面以Redhat　AS4为例，举例说明这两种登陆方式的用法。

    [root@mail ~]# ssh 172.18.6.227
    The authenticity of host ’172.18.6.227 (172.18.6.227)’ can’t be established.
    RSA key fingerprint is 43:80:f2:e1:9b:b6:6e:c0:e2:dd:57:8f:ed:89:b3:81.
    Are you sure you want to continue connecting (yes/no)? yes
    Warning: Permanently added ’172.18.6.227′ (RSA) to the list of known hosts.
    root@172.18.6.227‘s password: 
    Last login: Thu Jul 12 18:47:47 2007 from 172.18.6.130
    [root@qmail ~]#

第一次登陆后，ssh就会把登陆的ssh指纹存放在用户home目录的.ssh目录的know_hosts文件中，如果远程系统重装过系统，ssh指纹已经改变，你需要把 .ssh 目录下的know_hosts中的相应指纹删除，再登陆回答yes，方可登陆。请注意.ssh目录是开头是”.”的隐藏目录，需要ls –a参数才能看到。而且这个目录的权限必须是700,并且用户的home目录也不能给其他用户写权限，否则ssh服务器会拒绝登陆。如果发生不能登陆的问题，请察看服务器上的日志文件/var/log/secure。通常能很快找到不能登陆的原因。

ssh远程执行命令：

    [root@mail ~]# ssh 172.18.6.227 ls -l /
    root@172.18.6.227‘s password: 
    total 1244
    drwxr-xr-x    2 root root    4096 Jun 26 04:02 bin
    drwxr-xr-x    4 root root    4096 Mar 29 11:17 boot
    drwxr-xr-x    2 root root    4096 Jan 25 11:26 command
    drwxr-xr-x   15 root root    4096 Jun 12 20:09 data
    drwxr-xr-x    9 root root    5360 Jul  2 13:38 dev
    drwxr-xr-x   87 root root   12288 Jul 11 04:02 etc
    drwxr-xr-x   20 root root    4096 Apr 10 10:54 home
    drwxr-xr-x    2 root root    4096 Aug 13  2004 initrd

输入正确的密码后，ssh会链接远程服务器的sshd服务器程序，然后执行远程服务器上的
ls –l /命令　，并把输入结果传到本地服务器。相当于你先登陆到远程服务器，然后再实行命令ls –l /，最后再登出服务器。需要提醒的是，如果你需要登陆服务器并执行不止一个命令，必须要把命令用单引号或双引号引起来：

    ssh 172.18.6.227 “cd /root && ls “

ssh的远程实行命令的功能是用来代替原始的R系列命令的，在ssh出现之前系统管理员们不得不用rexec, rsh等不安全的远程执行命令工具来完成同样的操作。这个功能在管理大批机器的时候是非常有用的，比如我要重启10.0.0.0/24网段内所有的服务器，只要输入一条命令：

    for i in $(seq 1 254) ; do  ssh 10.0.0.${i} reboot ; done

就可以完成重启所有服务器的操作，也许你会说，这要虽然不需要再登陆每一台服务器了，但是还是要每次输入密码，多麻烦啊。别急，下面要讲的用ssh public key方式登陆就是要解决问题。

采用public key登录：

openssh的ssh-keygen命令用来产生这样的私钥和公钥。

    [root@mail ~]# ssh-keygen -b 1024 -t dsa -C gucuiwen@myserver.com
    Generating public/private dsa key pair.
    #提示正在生成，如果选择4096长度，可能需要较长时间
    Enter file in which to save the key (/root/.ssh/id_dsa): 
    ＃询问把公钥和私钥放在那里，回车用默认位置即可
    Enter passphrase (empty for no passphrase): 
    ＃询问输入私钥密语，为了实现自动登陆，应该不要密语，直接回车
    Enter same passphrase again: 
    ＃再次提示输入密语，再次直接回车
    Your identification has been saved in /root/.ssh/id_dsa.
    Your public key has been saved in /root/.ssh/id_dsa.pub.
    ＃提示公钥和私钥已经存放在/root/.ssh/目录下
    The key fingerprint is:
    71:e5:cb:15:d3:8c:05:ed:05:84:85:32:ce:b1:31:ce gucuiwen@myserver.com
    ＃提示key的指纹

说明：

-b 1024　采用长度为1024字节的公钥/私钥对，最长4096字节，一般1024或2048就可以了，太长的话加密解密需要的时间也长。  
-t dsa　　采用dsa加密方式的公钥/私钥对，除了dsa还有rsa方式，rsa方式最短不能小于768字节长度。  
-C gucuiwen@myserver.com 对这个公钥/私钥对的一个注释和说明，一般用所有人的邮件代替。可以省略不写，更多其他参数请man ssh-keygen。

    [root@mail ~]# ls -l /root/.ssh
    total 16
    -rw——-  1 root root 668 Jul 12 20:07 id_dsa
    -rw-r–r–  1 root root 611 Jul 12 20:07 id_dsa.pub
    -rw-r–r–  1 root root 222 Jul 12 19:37 known_hosts

产生的公钥/私钥文件在用户home目录的.ssh目录下，其中id_dsa.pub是公钥，把产生的公钥上传到需要登陆的服务器的对应用户目录的home目录的.ssh目录下，再一次强调用户自己的目录(home目录)必须不能有其他人可写的权限，.ssh目录的权限必须是700，即除了用户自己，其他人没有任何读写察看该目录的权限，否则ssh服务器会拒绝登陆。ssh默认的公钥文件是用户home目录下的.ssh目录下的 authorized_keys 文件，因此需要把产生的公钥以这个文件名放到服务器的/root/.ssh/目录下，这个文件中可以存放多个客户端的公钥文件，就好比一个大门上可以上很多锁，可以有不同的钥匙来尝试开锁，只要有一个锁被打开了，门就可以打开了。放到服务器上应该是这样子的：

私钥必须是600权限，否则ssh服务器会拒绝用户登陆。

大致就是这个样子了。现把`/etc/ssh/ssh_config` 和 `/etc/ssh/sshd_config`的配置说下。

`/etc/ssh/ssh_config:`

    Host *
    选项“Host”只对能够匹配后面字串的计算机有效。“*”表示所有的计算机。

    ForwardAgent no
    “ForwardAgent”设置连接是否经过验证代理（如果存在）转发给远程计算机。

    ForwardX11 no
    “ForwardX11”设置X11连接是否被自动重定向到安全的通道和显示集（DISPLAY set）。

    RhostsAuthentication no
    “RhostsAuthentication”设置是否使用基于rhosts的安全验证。

    RhostsRSAAuthentication no
    “RhostsRSAAuthentication”设置是否使用用RSA算法的基于rhosts的安全验证。

    RSAAuthentication yes
    “RSAAuthentication”设置是否使用RSA算法进行安全验证。

    PasswordAuthentication yes
    “PasswordAuthentication”设置是否使用口令验证。

    FallBackToRsh no
    “FallBackToRsh”设置如果用ssh连接出现错误是否自动使用rsh。

    UseRsh no
    “UseRsh”设置是否在这台计算机上使用“rlogin/rsh”。

    BatchMode no
    “BatchMode”如果设为“yes”，passphrase/password（交互式输入口令）的提示将被禁止。当不能交互式输入口令的时候，这个选项对脚本文件和批处理任务十分有用。

    CheckHostIP yes
    “CheckHostIP”设置ssh是否查看连接到服务器的主机的IP地址以防止DNS欺骗。建议设置为“yes”。

    StrictHostKeyChecking no
    “StrictHostKeyChecking”如果设置成“yes”，ssh就不会自动把计算机的密匙加入“$HOME/.ssh/known_hosts”文件，并且一旦计算机的密匙发生了变化，就拒绝连接。

    IdentityFile ~/.ssh/identity
    “IdentityFile”设置从哪个文件读取用户的RSA安全验证标识。

    Port 22
    “Port”设置连接到远程主机的端口。

    Cipher blowfish
    “Cipher”设置加密用的密码。

    EscapeChar ~
    “EscapeChar”设置escape字符。

`/etc/ssh/sshd_config:`

    Port 22
    “Port”设置sshd监听的端口号。

    ListenAddress 192.168.1.1
    “ListenAddress”设置sshd服务器绑定的IP地址。

    HostKey /etc/ssh/ssh_host_key

    “HostKey”设置包含计算机私人密匙的文件。

    ServerKeyBits 1024
    “ServerKeyBits”定义服务器密匙的位数。

    LoginGraceTime 600
    “LoginGraceTime”设置如果用户不能成功登录，在切断连接之前服务器需要等待的时间（以秒为单位）。

    KeyRegenerationInterval 3600
    “KeyRegenerationInterval”设置在多少秒之后自动重新生成服务器的密匙（如果使用密匙）。重新生成密匙是为了防止用盗用的密匙解密被截获的信息。

    PermitRootLogin no
    “PermitRootLogin”设置root能不能用ssh登录。这个选项一定不要设成“yes”。

    IgnoreRhosts yes
    “IgnoreRhosts”设置验证的时候是否使用“rhosts”和“shosts”文件。

    IgnoreUserKnownHosts yes
    “IgnoreUserKnownHosts”设置ssh daemon是否在进行RhostsRSAAuthentication安全验证的时候忽略用户的“$HOME/.ssh/known_hosts”

    StrictModes yes
    “StrictModes”设置ssh在接收登录请求之前是否检查用户家目录和rhosts文件的权限和所有权。这通常是必要的，因为新手经常会把自己的目录和文件设成任何人都有写权限。

    X11Forwarding no
    “X11Forwarding”设置是否允许X11转发。

    PrintMotd yes
    “PrintMotd”设置sshd是否在用户登录的时候显示“/etc/motd”中的信息。

    SyslogFacility AUTH
    “SyslogFacility”设置在记录来自sshd的消息的时候，是否给出“facility code”。

    LogLevel INFO
    “LogLevel”设置记录sshd日志消息的层次。INFO是一个好的选择。查看sshd的man帮助页，已获取更多的信息。

    RhostsAuthentication no
    “RhostsAuthentication”设置只用rhosts或“/etc/hosts.equiv”进行安全验证是否已经足够了。

    RhostsRSAAuthentication no
    “RhostsRSA”设置是否允许用rhosts或“/etc/hosts.equiv”加上RSA进行安全验证。

    RSAAuthentication yes
    “RSAAuthentication”设置是否允许只有RSA安全验证。

    PasswordAuthentication yes
    “PasswordAuthentication”设置是否允许口令验证。

    PermitEmptyPasswords no
    “PermitEmptyPasswords”设置是否允许用口令为空的帐号登录。

    AllowUsers admin
    “AllowUsers”的后面可以跟着任意的数量的用户名的匹配串（patterns）或user@host这样的匹配串，这些字符串用空格隔开。主机名可以是DNS名或IP地址。

将SSH2兼容格式的公钥转换成为Openssh兼容格式

    ssh-keygen -i -f Identity.pub >> /root/.ssh/authorized_keys2

当服务器端公钥改变的时候认证的IP，要将本地的 `know_hosts` 相关的节点删除

	ssh-keygen -R 服务器IP或者网址 

否则会出现 `WARNING: REMOTE HOST IDENTIFICATION HAS CHANGED!`

## scp 拷贝

参见 ·[scp 命令](<./sftp>)

scp 就是 secure copy, 是用来进行远程文件拷贝的 . 数据传输使用 ssh1, 并且和 ssh1 使用相同的认证方式 , 提供相同的安全保证。与 rcp 不同的是 ,scp 会要求你输入密码，如果需要的话。

命令基本格式：

    scp [可选参数] file_source file_target
    scp 本地用户名 @IP 地址 : 文件名 1 远程用户名 @IP 地址 : 文件名 2

常用的参数:

* **-v** 和大多数 linux 命令中的 -v 意思一样 , 用来显示进度 . 可以用来查看连接 , 认证 , 或是配置错误 .
* **-C** 使能压缩选项 .
* **-P** 选择端口 . 注意 -p 已经被 rcp 使用 .
* **-4** 强行使用 IPV4 地址 .
* **-6** 强行使用 IPV6 地址 .
* **-r** Recursively copy entire directories.


## sftp 默认ftp

## sshfs 挂载远端





