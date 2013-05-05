# nmap 
**title:** 网络探测和安全扫描程序nmap手册  
**taggs:** linux,nmap,hacker

http://linux.chinaunix.net/doc/safe/2005-01-30/1004.shtml

nmap是一个网络探测和安全扫描程序，系统管理者和个人可以使用这个软件扫描大型的网络，获取那台主机正在运行以及提供什么服务等信息。nmap支持很多扫描技术，例如：UDP、TCP connect()、TCP SYN(半开扫描)、ftp代理(bounce攻击)、反向标志、ICMP、FIN、ACK扫描、圣诞树(Xmas Tree)、SYN扫描和null扫描。从扫描类型一节可以得到细节。nmap还提供了一些高级的特征，例如：通过TCP/IP协议栈特征探测操作系统类型，秘密扫描，动态延时和重传计算，并行扫描，通过并行ping扫描探测关闭的主机，诱饵扫描，避开端口过滤检测，直接RPC扫描(无须端口影射)，碎片扫描，以及灵活的目标和端口设定.

- - - - - - - - - - - - - - -

## 1.名称

nmap-网络探测和安全扫描工具

## 2.语法

    nmap [Scan Type(s)] [Options]

##3.描述


nmap是一个网络探测和安全扫描程序，系统管理者和个人可以使用这个软件扫描大型的网络，获取那台主机正在运行以及提供什么服务等信息。nmap支持很多扫描技术，例如：UDP、TCP connect()、TCP SYN(半开扫描)、ftp代理(bounce攻击)、反向标志、ICMP、FIN、ACK扫描、圣诞树(Xmas Tree)、SYN扫描和null扫描。从扫描类型一节可以得到细节。nmap还提供了一些高级的特征，例如：通过TCP/IP协议栈特征探测操作系统类型，秘密扫描，动态延时和重传计算，并行扫描，通过并行ping扫描探测关闭的主机，诱饵扫描，避开端口过滤检测，直接RPC扫描(无须端口影射)，碎片扫描，以及灵活的目标和端口设定。


为了提高nmap在non-root状态下的性能，软件的设计者付出了很大的努力。很不幸，一些内核界面(例如raw socket)需要在root状态下使用。所以应该尽可能在root使用nmap。


nmap运行通常会得到被扫描主机端口的列表。nmap总会给出well known端口的服务名(如果可能)、端口号、状态和协议等信息。每个端口的状态有：open、filtered、unfiltered。open状态意味着目标主机能够在这个端口使用accept()系统调用接受连接。filtered状态表示：防火墙、包过滤和其它的网络安全软件掩盖了这个端口，禁止 nmap探测其是否打开。unfiltered表示：这个端口关闭，并且没有防火墙/包过滤软件来隔离nmap的探测企图。通常情况下，端口的状态基本都是unfiltered状态，只有在大多数被扫描的端口处于filtered状态下，才会显示处于unfiltered状态的端口。


根据使用的功能选项，nmap也可以报告远程主机的下列特征：使用的操作系统、TCP序列、运行绑定到每个端口上的应用程序的用户名、DNS名、主机地址是否是欺骗地址、以及其它一些东西。

## 4.功能选项


功能选项可以组合使用。一些功能选项只能够在某种扫描模式下使用。nmap会自动识别无效或者不支持的功能选项组合，并向用户发出警告信息。


如果你是有经验的用户，可以略过结尾的示例一节。可以使用nmap -h快速列出功能选项的列表。

### 4.1 扫描类型


**-sT**
:    TCP connect()扫描：这是最基本的TCP扫描方式。connect()是一种系统调用，由操作系统提供，用来打开一个连接。如果目标端口有程序监听， connect()就会成功返回，否则这个端口是不可达的。这项技术最大的优点是，你勿需root权限。任何UNIX用户都可以自由使用这个系统调用。这种扫描很容易被检测到，在目标主机的日志中会记录大批的连接请求以及错误信息。

**-sS**
:    TCP同步扫描(TCP SYN)：因为不必全部打开一个TCP连接，所以这项技术通常称为半开扫描(half-open)。你可以发出一个TCP同步包(SYN)，然后等待回应。如果对方返回SYN|ACK(响应)包就表示目标端口正在监听；如果返回RST数据包，就表示目标端口没有监听程序；如果收到一个SYN|ACK包，源主机就会马上发出一个RST(复位)数据包断开和目标主机的连接，这实际上有我们的操作系统内核自动完成的。这项技术最大的好处是，很少有系统能够把这记入系统日志。不过，你需要root权限来定制SYN数据包。

**-sF -sF -sN**
:    秘密FIN数据包扫描、圣诞树(Xmas Tree)、空(Null)扫描模式：即使SYN扫描都无法确定的情况下使用。一些防火墙和包过滤软件能够对发送到被限制端口的SYN数据包进行监视，而且有些程序比如synlogger和courtney能够检测那些扫描。这些高级的扫描方式可以逃过这些干扰。这些扫描方式的理论依据是：关闭的端口需要对你的探测包回应RST包，而打开的端口必需忽略有问题的包(参考RFC 793第64页)。FIN扫描使用暴露的FIN数据包来探测，而圣诞树扫描打开数据包的FIN、URG和PUSH标志。不幸的是，微软决定完全忽略这个标准，另起炉灶。所以这种扫描方式对Windows95/NT无效。不过，从另外的角度讲，可以使用这种方式来分别两种不同的平台。如果使用这种扫描方式可以发现打开的端口，你就可以确定目标注意运行的不是Windows系统。如果使用-sF、-sX或者-sN扫描显示所有的端口都是关闭的，而使用SYN扫描显示有打开的端口，你可以确定目标主机可能运行的是Windwos系统。现在这种方式没有什么太大的用处，因为nmap有内嵌的操作系统检测功能。还有其它几个系统使用和windows同样的处理方式，包括Cisco、BSDI、HP/UX、MYS、IRIX。在应该抛弃数据包时，以上这些系统都会从打开的端口发出复位数据包。

**-sP**
:    ping扫描：有时你只是想知道此时网络上哪些主机正在运行。通过向你指定的网络内的每个IP地址发送ICMP echo请求数据包，nmap就可以完成这项任务。如果主机正在运行就会作出响应。不幸的是，一些站点例如：microsoft.com阻塞ICMP echo请求数据包。然而，在默认的情况下nmap也能够向80端口发送TCP ack包，如果你收到一个RST包，就表示主机正在运行。nmap使用的第三种技术是：发送一个SYN包，然后等待一个RST或者SYN/ACK包。对于非root用户，nmap使用connect()方法。

在默认的情况下(root用户)，nmap并行使用ICMP和ACK技术。

注意，nmap在任何情况下都会进行ping扫描，只有目标主机处于运行状态，才会进行后续的扫描。如果你只是想知道目标主机是否运行，而不想进行其它扫描，才会用到这个选项。

-sU

UDP扫描：如果你想知道在某台主机上提供哪些UDP(用户数据报协议,RFC768)服务，可以使用这种扫描方法。nmap首先向目标主机的每个端口发出一个0字节的UDP包，如果我们收到端口不可达的ICMP消息，端口就是关闭的，否则我们就假设它是打开的。

有些人可能会想UDP扫描是没有什么意思的。但是，我经常会想到最近出现的solaris rpcbind缺陷。rpcbind隐藏在一个未公开的UDP端口上，这个端口号大于32770。所以即使端口111(portmap的众所周知端口号) 被防火墙阻塞有关系。但是你能发现大于30000的哪个端口上有程序正在监听吗?使用UDP扫描就能！cDc Back Orifice的后门程序就隐藏在Windows主机的一个可配置的UDP端口中。不考虑一些通常的安全缺陷，一些服务例如:snmp、tftp、NFS 使用UDP协议。不幸的是，UDP扫描有时非常缓慢，因为大多数主机限制ICMP错误信息的比例(在RFC1812中的建议)。例如，在Linux内核中 (在net/ipv4/icmp.h文件中)限制每4秒钟只能出现80条目标不可达的ICMP消息，如果超过这个比例，就会给1/4秒钟的处罚。 solaris的限制更加严格，每秒钟只允许出现大约2条ICMP不可达消息，这样，使扫描更加缓慢。nmap会检测这个限制的比例，减缓发送速度，而不是发送大量的将被目标主机丢弃的无用数据包。

不过Micro$oft忽略了RFC1812的这个建议，不对这个比例做任何的限制。所以我们可以能够快速扫描运行Win95/NT的主机上的所有65K个端口。

-sA

ACK扫描：这项高级的扫描方法通常用来穿过防火墙的规则集。通常情况下，这有助于确定一个防火墙是功能比较完善的或者是一个简单的包过滤程序，只是阻塞进入的SYN包。

这种扫描是向特定的端口发送ACK包(使用随机的应答/序列号)。如果返回一个RST包，这个端口就标记为unfiltered状态。如果什么都没有返回，或者返回一个不可达ICMP消息，这个端口就归入filtered类。注意，nmap通常不输出unfiltered的端口，所以在输出中通常不显示所有被探测的端口。显然，这种扫描方式不能找出处于打开状态的端口。

-sW

对滑动窗口的扫描：这项高级扫描技术非常类似于ACK扫描，除了它有时可以检测到处于打开状态的端口，因为滑动窗口的大小是不规则的，有些操作系统可以报告其大小。这些系统至少包括：某些版本的AIX、Amiga、BeOS、BSDI、Cray、Tru64 UNIX、DG/UX、OpenVMS、Digital UNIX、OpenBSD、OpenStep、QNX、Rhapsody、SunOS 4.x、Ultrix、VAX、VXWORKS。从nmap-hackers邮件3列表的文档中可以得到完整的列表。

-sR

RPC扫描。这种方法和nmap的其它不同的端口扫描方法结合使用。选择所有处于打开状态的端口向它们发出SunRPC程序的NULL命令，以确定它们是否是RPC端口，如果是，就确定是哪种软件及其版本号。因此你能够获得防火墙的一些信息。诱饵扫描现在还不能和RPC扫描结合使用。

-b

FTP反弹攻击(bounce attack):FTP协议(RFC 959)有一个很有意思的特征，它支持代理FTP连接。也就是说，我能够从evil.com连接到FTP服务器target.com，并且可以要求这台 FTP服务器为自己发送Internet上任何地方的文件！1985年，RFC959完成时，这个特征就能很好地工作了。然而，在今天的Internet 中，我们不能让人们劫持FTP服务器，让它向Internet上的任意节点发送数据。如同Hobbit在1995年写的文章中所说的，这个协议"能够用来做投递虚拟的不可达邮件和新闻，进入各种站点的服务器,填满硬盘，跳过防火墙，以及其它的骚扰活动，而且很难进行追踪"。我们可以使用这个特征，在一台代理FTP服务器扫描TCP端口。因此，你需要连接到防火墙后面的一台FTP服务器，接着进行端口扫描。如果在这台FTP服务器中有可读写的目录，你还可以向目标端口任意发送数据(不过nmap不能为你做这些)。

传递给-b功能选项的参数是你要作为代理的FTP服务器。语法格式为：

-b username:password@server:port。

除了server以外，其余都是可选的。如果你想知道什么服务器有这种缺陷，可以参考我在Phrack 51发表的文章。还可以在nmap的站点得到这篇文章的最新版本。

4.2 通用选项


这些内容不是必需的，但是很有用。


-P0

在扫描之前，不必ping主机。有些网络的防火墙不允许ICMP echo请求穿过，使用这个选项可以对这些网络进行扫描。microsoft.com就是一个例子，因此在扫描这个站点时，你应该一直使用-P0或者-PT 80选项。

-PT

扫描之前，使用TCP ping确定哪些主机正在运行。nmap不是通过发送ICMP echo请求包然后等待响应来实现这种功能，而是向目标网络(或者单一主机)发出TCP ACK包然后等待回应。如果主机正在运行就会返回RST包。只有在目标网络/主机阻塞了ping包，而仍旧允许你对其进行扫描时，这个选项才有效。对于非 root用户，我们使用connect()系统调用来实现这项功能。使用-PT <端口号>来设定目标端口。默认的端口号是80，因为这个端口通常不会被过滤。

-PS

对于root用户，这个选项让nmap使用SYN包而不是ACK包来对目标主机进行扫描。如果主机正在运行就返回一个RST包(或者一个SYN/ACK包)。

-PI

设置这个选项，让nmap使用真正的ping(ICMP echo请求)来扫描目标主机是否正在运行。使用这个选项让nmap发现正在运行的主机的同时，nmap也会对你的直接子网广播地址进行观察。直接子网广播地址一些外部可达的IP地址，把外部的包转换为一个内向的IP广播包，向一个计算机子网发送。这些IP广播包应该删除，因为会造成拒绝服务攻击(例如 smurf)。

-PB

这是默认的ping扫描选项。它使用ACK(-PT)和ICMP(-PI)两种扫描类型并行扫描。如果防火墙能够过滤其中一种包，使用这种方法，你就能够穿过防火墙。

-O

这个选项激活对TCP/IP指纹特征(fingerprinting)的扫描，获得远程主机的标志。换句话说，nmap使用一些技术检测目标主机操作系统网络协议栈的特征。nmap使用这些信息建立远程主机的指纹特征，把它和已知的操作系统指纹特征数据库做比较，就可以知道目标主机操作系统的类型。

-I

这个选项打开nmap的反向标志扫描功能。Dave Goldsmith 1996年向bugtap发出的邮件注意到这个协议，ident协议(rfc 1413)允许使用TCP连接给出任何进程拥有者的用户名，即使这个进程并没有初始化连接。例如，你可以连接到HTTP端口，接着使用identd确定这个服务器是否由root用户运行。这种扫描只能在同目标端口建立完全的TCP连接时(例如：-sT扫描选项)才能成功。使用-I选项是，远程主机的 identd精灵进程就会查询在每个打开的端口上监听的进程的拥有者。显然，如果远程主机没有运行identd程序，这种扫描方法无效。

-f

这个选项使nmap使用碎片IP数据包发送SYN、FIN、XMAS、NULL。使用碎片数据包增加包过滤、入侵检测系统的难度，使其无法知道你的企图。不过，要慎重使用这个选项！有些程序在处理这些碎片包时会有麻烦，我最喜欢的嗅探器在接受到碎片包的头36个字节时，就会发生 segmentation faulted。因此，在nmap中使用了24个字节的碎片数据包。虽然包过滤器和防火墙不能防这种方法，但是有很多网络出于性能上的考虑，禁止数据包的分片。

注意这个选项不能在所有的平台上使用。它在Linux、FreeBSD、OpenBSD以及其它一些UNIX系统能够很好工作。

-v

冗余模式。强烈推荐使用这个选项，它会给出扫描过程中的详细信息。使用这个选项，你可以得到事半功倍的效果。使用-d选项可以得到更加详细的信息。

-h

快速参考选项。

-oN

把扫描结果重定向到一个可读的文件logfilename中。

-oM

把扫描结果重定向到logfilename文件中，这个文件使用主机可以解析的语法。你可以使用-oM -来代替logfilename，这样输出就被重定向到标准输出stdout。在这种情况下，正常的输出将被覆盖，错误信息荏苒可以输出到标准错误 stderr。要注意，如果同时使用了-v选项，在屏幕上会打印出其它的信息。

-oS 　　 thIs l0gz th3 r3suLtS of YouR ScanZ iN a s| 　　THe fiL3 U sPecfy 4s an arGuMEnT! U kAn gIv3 the 4rgument -

(wItHOUt qUOteZ) to sh00t output iNT0 stDouT!@!! 莫名其妙，下面是我猜着翻译的，相形字？

把扫描结果重定向到一个文件logfilename中，这个文件使用一种"黑客方言"的语法形式(作者开的玩笑?)。同样，使用-oS -就会把结果重定向到标准输出上。

-resume

某个网络扫描可能由于control-C或者网络损失等原因被中断，使用这个选项可以使扫描接着以前的扫描进行。logfilename是被取消扫描的日志文件，它必须是可读形式或者机器可以解析的形式。而且接着进行的扫描不能增加新的选项，只能使用与被中断的扫描相同的选项。nmap会接着日志文件中的最后一次成功扫描进行新的扫描。

-iL

从inputfilename文件中读取扫描的目标。在这个文件中要有一个主机或者网络的列表，由空格键、制表键或者回车键作为分割符。如果使用-iL -，nmap就会从标准输入stdin读取主机名字。你可以从指定目标一节得到更加详细的信息。

-iR

让nmap自己随机挑选主机进行扫描。

-p <端口范围>

这个选项让你选择要进行扫描的端口号的范围。例如，-p 23表示：只扫描目标主机的23号端口。-p 20-30,139,60000-表示：扫描20到30号端口，139号端口以及所有大于60000的端口。在默认情况下，nmap扫描从1到1024号以及nmap-services文件(如果使用RPM软件包，一般在/usr/share/nmap/目录中)中定义的端口列表。

-F

快速扫描模式，只扫描在nmap-services文件中列出的端口。显然比扫描所有65535个端口要快。

-D

使用诱饵扫描方法对目标网络/主机进行扫描。如果nmap使用这种方法对目标网络进行扫描，那么从目标主机/网络的角度来看，扫描就象从其它主机 (decoy1,等)发出的。从而，即使目标主机的IDS(入侵检测系统)对端口扫描发出报警，它们也不可能知道哪个是真正发起扫描的地址，哪个是无辜的。这种扫描方法可以有效地对付例如路由跟踪、response-dropping等积极的防御机制，能够很好地隐藏你的IP地址。

每个诱饵主机名使用逗号分割开，你也可以使用ME选项，它代表你自己的主机，和诱饵主机名混杂在一起。如果你把ME放在第六或者更靠后的位置，一些端口扫描检测软件几乎根本不会显示你的IP地址。如果你不使用ME选项，nmap会把你的IP地址随机夹杂在诱饵主机之中。

注意:你用来作为诱饵的主机应该正在运行或者你只是偶尔向目标发送SYN数据包。很显然，如果在网络上只有一台主机运行，目标将很轻松就会确定是哪台主机进行的扫描。或许，你还要直接使用诱饵的IP地址而不是其域名，这样诱饵网络的域名服务器的日志上就不会留下关于你的记录。

还要注意：一些愚蠢的端口扫描检测软件会拒绝路由试图进行端口扫描的主机。因而，你需要让目标主机和一些诱饵断开连接。如果诱饵是目标主机的网关或者就是其自己时，会给目标主机造成很大问题。所以你需要慎重使用这个选项。

诱饵扫描既可以在起始的ping扫描也可以在真正的扫描状态下使用。它也可以和-O选项组合使用。

使用太多的诱饵扫描能够减缓你的扫描速度甚至可能造成扫描结果不正确。同时，有些ISP会把你的欺骗包过滤掉。虽然现在大多数的ISP不会对此进行限制。

-S <IP_Address>

在一些情况下，nmap可能无法确定你的源地址(nmap会告诉你)。在这种情况下，可以使用这个选项给出你的IP地址。

在欺骗扫描时，也使用这个选项。使用这个选项可以让目标认为是其它的主机对自己进行扫描。

-e

告诉nmap使用哪个接口发送和接受数据包。nmap能够自动对此接口进行检测，如果无效就会告诉你。

-g

设置扫描的源端口。一些天真的防火墙和包过滤器的规则集允许源端口为DNS(53)或者FTP-DATA(20)的包通过和实现连接。显然，如果攻击者把源端口修改为20或者53，就可以摧毁防火墙的防护。在使用UDP扫描时，先使用53号端口；使用TCP扫描时，先使用20号端口。注意只有在能够使用这个端口进行扫描时，nmap才会使用这个端口。例如，如果你无法进行TCP扫描，nmap会自动改变源端口，即使你使用了-g选项。

对于一些扫描，使用这个选项会造成性能上的微小损失，因为我有时会保存关于特定源端口的一些有用的信息。

-r

告诉nmap不要打乱被扫描端口的顺序。

--randomize_hosts

使nmap在扫描之前，打乱每组扫描中的主机顺序，nmap每组可以扫描最多2048台主机。这样，可以使扫描更不容易被网络监视器发现，尤其和--scan_delay 选项组合使用，更能有效避免被发现。

-M

设置进行TCP connect()扫描时，最多使用多少个套接字进行并行的扫描。使用这个选项可以降低扫描速度，避免远程目标宕机。

4.3 适时选项


通常，nmap在运行时，能够很好地根据网络特点进行调整。扫描时，nmap会尽量减少被目标检测到的机会，同时尽可能加快扫描速度。然而，nmap默认的适时策略有时候不太适合你的目标。使用下面这些选项，可以控制nmap的扫描timing：

-T

设置nmap的适时策略。Paranoid:为了避开IDS的检测使扫描速度极慢，nmap串行所有的扫描，每隔至少5分钟发送一个包； Sneaky：也差不多，只是数据包的发送间隔是15秒；Polite：不增加太大的网络负载，避免宕掉目标主机，串行每个探测，并且使每个探测有0.4 秒种的间隔；Normal:nmap默认的选项，在不是网络过载或者主机/端口丢失的情况下尽可能快速地扫描；Aggressive:设置5分钟的超时限制，使对每台主机的扫描时间不超过5分钟，并且使对每次探测回应的等待时间不超过1.5秒钟；b>Insane:只适合快速的网络或者你不在意丢失某些信息，每台主机的超时限制是75秒，对每次探测只等待0.3秒钟。你也可是使用数字来代替这些模式，例如：-T 0等于-T Paranoid，-T 5等于-T Insane。

这些适时模式不能下面的适时选项组合使用。
--host_timeout

设置扫描一台主机的时间，以毫秒为单位。默认的情况下，没有超时限制。
--max_rtt_timeout

设置对每次探测的等待时间，以毫秒为单位。如果超过这个时间限制就重传或者超时。默认值是大约9000毫秒。
--min_rtt_timeout

当目标主机的响应很快时，nmap就缩短每次探测的超时时间。这样会提高扫描的速度，但是可能丢失某些响应时间比较长的包。使用这个选项，可以让nmap对每次探测至少等待你指定的时间，以毫秒为单位。
--initial_rtt_timeout

设置初始探测的超时值。一般这个选项只在使用-P0选项扫描有防火墙保护的主机才有用。默认值是6000毫秒。
--max_parallelism

设置最大的并行扫描数量。--max_parallelism 1表示同时只扫描一个端口。这个选项对其它的并行扫描也有效，例如ping sweep, RPC scan。
--scan_delay

设置在两次探测之间，nmap必须等待的时间。这个选项主要用于降低网络的负载。

4.4 目标设定


在nmap的所有参数中，只有目标参数是必须给出的。其最简单的形式是在命令行直接输入一个主机名或者一个IP地址。如果你希望扫描某个IP地址的一个子网，你可以在主机名或者IP地址的后面加上/掩码。掩码在0(扫描整个网络)到32(只扫描这个主机)。使用/24扫描C类地址，/16扫描B类地址。


除此之外，nmap还有更加强大的表示方式让你更加灵活地指定IP地址。例如，如果要扫描这个B类网络128.210.*.*，你可以使用下面三种方式来指定这些地址:128.210.*.*、128.21-.0-255.0-255或者128.210.0.0/16这三种形式是等价的。

5.例子


本节将由浅入深地举例说明如何使用nmap。

nmap -v target.example.com
扫描主机target.example.com的所有TCP端口。-v打开冗余模式。

nmap -sS -O target.example.com/24
发起对target.example.com所在网络上的所有255个IP地址的秘密SYN扫描。同时还探测每台主机操作系统的指纹特征。需要root权限。

nmap -sX -p 22,53,110,143,4564 128.210.*.1-127
对B类IP地址128.210中255个可能的8位子网的前半部分发起圣诞树扫描。确定这些系统是否打开了sshd、DNS、pop3d、imapd和4564端口。注意圣诞树扫描对Micro$oft的系统无效，因为其协议栈的TCP层有缺陷。

nmap -v --randomize_hosts -p 80 *.*.2.3-5
只扫描指定的IP范围，有时用于对这个Internet进行取样分析。nmap将寻找Internet上所有后两个字节是.2.3、.2.4、.2.5的 IP地址上的WEB服务器。如果你想发现更多有意思的主机，你可以使用127-222，因为在这个范围内有意思的主机密度更大。

host -l company.com | cut -d -f 4 | ./nmap -v -iL -
列出company.com网络的所有主机，让nmap进行扫描。注意：这项命令在GNU/Linux下使用。如果在其它平台，你可能要使用 其它的命令/选项。


- - - - - - - -- - - - - - - -- - - - - -
http://blog.csdn.net/aspirationflow/article/details/7694274

<h1 style="text-align:center">
  <a name="t0"></a>Nmap扫描原理与用法
</h1>
  2012年6月16日
# <a name="t1"></a>1     Nmap介绍


Nmap扫描原理与用法PDF：[下载地址][1]


Nmap是一款开源免费的网络发现（Network Discovery）和安全审计（Security Auditing）工具。软件名字Nmap是Network Mapper的简称。Nmap最初是由Fyodor在1997年开始创建的。随后在开源社区众多的志愿者参与下，该工具逐渐成为最为流行安全必备工具之一。最新版的Nmap6.0在2012年5月21日发布，详情请参见：[www.nmap.org][2]。

一般情况下，Nmap用于列举网络主机清单、管理服务升级调度、监控主机或服务运行状况。Nmap可以检测目标机是否在线、端口开放情况、侦测运行的服务类型及版本信息、侦测操作系统与设备类型等信息。

Nmap的优点：

1.      **灵活**。支持数十种不同的扫描方式，支持多种目标对象的扫描。

2.      **强大**。Nmap可以用于扫描互联网上大规模的计算机。

3.      **可移植**。支持主流操作系统：Windows/Linux/Unix/MacOS等等；源码开放，方便移植。

4.      **简单**。提供默认的操作能覆盖大部分功能，基本端口扫描nmap targetip，全面的扫描nmap –A targetip。

5.      **自由**。Nmap作为开源软件，在GPL License的范围内可以自由的使用。

6.      **文档丰富**。Nmap官网提供了详细的文档描述。Nmap作者及其他安全专家编写了多部Nmap参考书籍。

7.      **社区支持**。Nmap背后有强大的社区团队支持。

8.      **赞誉有加**。获得很多的奖励，并在很多影视作品中出现（如黑客帝国2、Die Hard4等）。

9.      **流行**。目前Nmap已经被成千上万的安全专家列为必备的工具之一。

 

## <a name="t2"></a>1.1    Zenmap

Zenmap是Nmap官方提供的图形界面，通常随Nmap的安装包发布。Zenmap是用Python语言编写而成的开源免费的图形界面，能够运行在不同操作系统平台上（Windows/Linux/Unix/Mac OS等）。Zenmap旨在为nmap提供更加简单的操作方式。简单常用的操作命令可以保存成为profile，用户扫描时选择profile即可；可以方便地比较不同的扫描结果；提供网络拓扑结构(NetworkTopology)的图形显示功能。
  <img src="http://my.csdn.net/uploads/201206/26/1340719144_5258.JPG" alt="" align="middle" height="623" width="557" />

其中Profile栏位，用于选择“Zenmap默认提供的Profile”或“用户创建的Profile”；Command栏位，用于显示选择Profile对应的命令或者用户自行指定的命令；Topology选项卡，用于显示扫描到的目标机与本机之间的拓扑结构。

## <a name="t3"></a>1.2    功能架构图
  <img src="http://my.csdn.net/uploads/201206/26/1340719324_9785.JPG" alt="" />
Nmap包含四项基本功能：

1.  主机发现（Host Discovery）
2.  端口扫描（Port Scanning）
3.  版本侦测（Version Detection）
4.  操作系统侦测（Operating System Detection）

而这四项功能之间，又存在大致的依赖关系（通常情况下的顺序关系，但特殊应用另外考虑），首先需要进行主机发现，随后确定端口状况，然后确定端口上运行具体应用程序与版本信息，然后可以进行操作系统的侦测。而在四项基本功能的基础上，Nmap提供防火墙与IDS（IntrusionDetection System,入侵检测系统）的规避技巧，可以综合应用到四个基本功能的各个阶段；另外Nmap提供强大的NSE（Nmap Scripting Language）脚本引擎功能，脚本可以对基本功能进行补充和扩展。

 

 

# <a name="t4"></a>2     Nmap基本扫描方法

Nmap主要包括四个方面的扫描功能，主机发现、端口扫描、应用与版本侦测、操作系统侦测。在详细讲解每个具体功能之前，首先可以看看Nmap的典型用法。

## <a name="t5"></a>2.1    用法引入

### <a name="t6"></a>2.1.1    确定端口状况

如果直接针对某台计算的IP地址或域名进行扫描，那么Nmap对该主机进行主机发现过程和端口扫描。该方式执行迅速，可以用于确定端口的开放状况。

命令形式:

<span style="color:#daeef3; background:black">nmap targethost</span>

可以确定目标主机在线情况及端口基本状况。
   <img src="http://my.csdn.net/uploads/201206/27/1340805820_1736.jpg" alt="" />

### <a name="t7"></a>2.1.2    完整全面的扫描

如果希望对某台主机进行完整全面的扫描，那么可以使用nmap内置的-A选项。使用了改选项，nmap对目标主机进行主机发现、端口扫描、应用程序与版本侦测、操作系统侦测及调用默认NSE脚本扫描。

命令形式：

<span style="color:#daeef3; background:black">nmap –T4 –A –v targethost</span>

其中-A选项用于使用进攻性（Aggressive）方式扫描；-T4指定扫描过程使用的时序（Timing），总有6个级别（0-5），级别越高，扫描速度越快，但也容易被防火墙或IDS检测并屏蔽掉，在网络通讯状况良好的情况推荐使用T4；-v表示显示冗余（verbosity）信息，在扫描过程中显示扫描的细节，从而让用户了解当前的扫描状态。
  <img src="http://my.csdn.net/uploads/201206/27/1340805865_9288.jpg" alt="" /><br />
例如，扫描局域网内地址为192.168.1.100的电脑。显而易见，扫描出的信息非常丰富，在对192.168.1.100的扫描报告部分中（以红框圈出），可以看到主机发现的结果“Host is up”；端口扫描出的结果，有996个关闭端口，4个开放端口（在未指定扫描端口时，Nmap默认扫描1000个最有可能开放的端口）；而版本侦测针对扫描到的开放状况进一步探测端口上运行的具体的应用程序和版本信息；OS侦测对该目标主机的设备类型与操作系统进行探测；而绿色框图是nmap调用NSE脚本进行进一步的信息挖掘的显示结果。

 

## <a name="t8"></a>2.2    主机发现

主机发现（Host Discovery），即用于发现目标主机是否在线（Alive，处于开启状态）。

### <a name="t9"></a>2.2.1    主机发现原理

主机发现发现的原理与Ping命令类似，发送探测包到目标主机，如果收到回复，那么说明目标主机是开启的。Nmap支持十多种不同的主机探测方式，比如发送ICMP ECHO/TIMESTAMP/NETMASK报文、发送TCPSYN/ACK包、发送SCTP INIT/COOKIE-ECHO包，用户可以在不同的条件下灵活选用不同的方式来探测目标机。

主机发现基本原理：（以ICMP echo方式为例）
  <img src="http://my.csdn.net/uploads/201207/01/1341105994_4721.jpg" alt="" /><br />
Nmap的用户位于源端，IP地址192.168.0.5，向目标主机192.168.0.3发送ICMP Echo Request。如果该请求报文没有被防火墙拦截掉，那么目标机会回复ICMP Echo Reply包回来。以此来确定目标主机是否在线。

默认情况下，Nmap会发送四种不同类型的数据包来探测目标主机是否在线。

1.      ICMP echo request

2.      a TCP SYN packet to port 443

3.      a TCP ACK packet to port 80

4.      an ICMP timestamp request

依次发送四个报文探测目标机是否开启。只要收到其中一个包的回复，那就证明目标机开启。使用四种不同类型的数据包可以避免因防火墙或丢包造成的判断错误。

### <a name="t10"></a>2.2.2    主机发现的用法

通常主机发现并不单独使用，而只是作为端口扫描、版本侦测、OS侦测先行步骤。而在某些特殊应用（例如确定大型局域网内活动主机的数量），可能会单独专门适用主机发现功能来完成。

不管是作为辅助用法还是专门用途，用户都可以使用Nmap提供的丰富的选项来定制主机发现的探测方式。

<ol start="1">
  <li class="alt">
    <span><span>-sL: List Scan 列表扫描，仅将指定的目标的IP列举出来，不进行主机发现。  </span></span>
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>-sn: Ping Scan 只进行主机发现，不进行端口扫描。  </span>
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>-Pn: 将所有指定的主机视作开启的，跳过主机发现的过程。  </span>
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>-PS/PA/PU/PY[portlist]: 使用TCPSYN/ACK或SCTP INIT/ECHO方式进行发现。  </span>
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>-PE/PP/PM: 使用ICMP echo, timestamp, and netmask 请求包发现主机。-PO[protocollist]: 使用IP协议包探测对方主机是否开启。  </span>
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>-n/-R: -n表示不进行DNS解析；-R表示总是进行DNS解析。  </span>
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>--dns-servers <serv1[,serv2],...>: 指定DNS服务器。  </span>
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>--system-dns: 指定使用系统的DNS服务器  </span>
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>--traceroute: 追踪每个路由节点  </span>
  </li>
</ol>

-sn: Ping Scan 只进行主机发现，不进行端口扫描。

-Pn: 将所有指定的主机视作开启的，跳过主机发现的过程。

-PS/PA/PU/PY[portlist]: 使用TCPSYN/ACK或SCTP INIT/ECHO方式进行发现。

-PE/PP/PM: 使用ICMP echo, timestamp, and netmask 请求包发现主机。-PO[protocollist]: 使用IP协议包探测对方主机是否开启。

-n/-R: -n表示不进行DNS解析；-R表示总是进行DNS解析。

--dns-servers &lt;serv1[,serv2],...>: 指定DNS服务器。

--system-dns: 指定使用系统的DNS服务器

--traceroute: 追踪每个路由节点</pre>其中，比较常用的使用的是-sn，表示只单独进行主机发现过程；-Pn表示直接跳过主机发现而进行端口扫描等高级操作（如果已经确知目标主机已经开启，可用该选项）；-n，如果不想使用DNS或reverse DNS解析，那么可以使用该选项。 

### <a name="t11"></a>2.2.3    使用演示

**探测scanme.nmap.org**

下面以探测[scanme.nmap.org][3] 的主机为例，简单演示主机发现的用法。

命令如下：

<span style="color:#daeef3; background:black">nmap –sn –PE –PS80,135 –PU53 scanme.nmap.org</span>
  <img src="http://my.csdn.net/uploads/201207/01/1341106101_4407.jpg" alt="" /><br />


使用Wireshark抓包，我们看到，scanme.nmap.org 的IP地址182.140.147.57发送了四个探测包：ICMPEcho，80和135端口的TCP SYN包，53端口的UDP包（DNS domain）。而收到ICMP Echo的回复与80端口的回复。从而确定了scanme.nmap.org主机正常在线。
  <img src="http://my.csdn.net/uploads/201207/01/1341106154_2093.jpg" alt="" /><br />
**探测局域网内活动主机**

扫描局域网192.168.1.100-192.168.1.120范围内哪些IP的主机是活动的。

命令如下：

<span style="color:#daeef3; background:black">nmap –sn 192.168.1.100-120</span>
  <img src="http://my.csdn.net/uploads/201207/01/1341106184_1310.jpg" alt="" /><br />
从结果中，可以看到这个IP范围内有三台主机处于活动状态。

从Wireshark抓取的包中，可以看到发送的探测包的情况：
  <img src="http://my.csdn.net/uploads/201207/01/1341106204_2446.jpg" alt="" /><br />
在局域网内，Nmap是通过ARP包来询问IP地址上的主机是否活动的，如果收到ARP回复包，那么说明主机在线。

例如，某条ARP回复的报文详细信息如下：


   <img src="http://my.csdn.net/uploads/201207/01/1341106231_4458.jpg" alt="" />
## <a name="t12"></a>2.3    端口扫描

端口扫描是Nmap最基本最核心的功能，用于确定目标主机的TCP/UDP端口的开放情况。

默认情况下，Nmap会扫描1000个最有可能开放的TCP端口。

Nmap通过探测将端口划分为6个状态：

1.  open：端口是开放的。
2.  closed：端口是关闭的。
3.  filtered：端口被防火墙IDS/IPS屏蔽，无法确定其状态。
4.  unfiltered：端口没有被屏蔽，但是否开放需要进一步确定。
5.  open|filtered：端口是开放的或被屏蔽。
6.  closed|filtered ：端口是关闭的或被屏蔽。

### <a name="t13"></a>2.3.1    端口扫描原理

Nmap在端口扫描方面非常强大，提供了十多种探测方式。

#### <a name="t14"></a>2.3.1.1    TCP SYN scanning

这是Nmap默认的扫描方式，通常被称作半开放扫描（Half-open scanning）。该方式发送SYN到目标端口，如果收到SYN/ACK回复，那么判断端口是开放的；如果收到RST包，说明该端口是关闭的。如果没有收到回复，那么判断该端口被屏蔽（Filtered）。因为该方式仅发送SYN包对目标主机的特定端口，但不建立的完整的TCP连接，所以相对比较隐蔽，而且效率比较高，适用范围广。

TCP SYN探测到端口关闭：
  <img src="http://my.csdn.net/uploads/201207/01/1341106252_9646.jpg" alt="" /><br />
TCP SYN探测到端口开放：
  <img src="http://my.csdn.net/uploads/201207/01/1341106266_3589.jpg" alt="" /><br />
#### <a name="t15"></a>2.3.1.2    TCP connect scanning 

TCP connect方式使用系统网络API connect向目标主机的端口发起连接，如果无法连接，说明该端口关闭。该方式扫描速度比较慢，而且由于建立完整的TCP连接会在目标机上留下记录信息，不够隐蔽。所以，TCP connect是TCP SYN无法使用才考虑选择的方式。

TCP connect探测到端口关闭：
  <img src="http://my.csdn.net/uploads/201207/01/1341106280_8116.jpg" alt="" /><br />
TCP connect探测到端口开放：
  <img src="http://my.csdn.net/uploads/201207/01/1341106296_3110.jpg" alt="" /><br />
#### <a name="t16"></a>2.3.1.3    TCP ACK scanning

向目标主机的端口发送ACK包，如果收到RST包，说明该端口没有被防火墙屏蔽；没有收到RST包，说明被屏蔽。该方式只能用于确定防火墙是否屏蔽某个端口，可以辅助TCP SYN的方式来判断目标主机防火墙的状况。

TCP ACK探测到端口被屏蔽：
  <img src="http://my.csdn.net/uploads/201207/01/1341106327_8291.jpg" alt="" /><br />
TCP ACK探测到端口未被屏蔽：
  <img src="http://my.csdn.net/uploads/201207/01/1341106346_8325.jpg" alt="" /><br />
#### <a name="t17"></a>2.3.1.4    TCP FIN/Xmas/NULL scanning

这三种扫描方式被称为秘密扫描（Stealthy Scan），因为相对比较隐蔽。FIN扫描向目标主机的端口发送的TCP FIN包或Xmas tree包/Null包，如果收到对方RST回复包，那么说明该端口是关闭的；没有收到RST包说明端口可能是开放的或被屏蔽的（open|filtered）。

其中Xmas tree包是指flags中<span style="color:#333333">FIN URG PUSH</span><span style="color:#333333">被置为</span><span style="color:#333333">1</span><span style="color:#333333">的</span><span style="color:#333333">TCP</span><span style="color:#333333">包；</span><span style="color:#333333">NULL</span><span style="color:#333333">包是指所有</span><span style="color:#333333">flags</span><span style="color:#333333">都为</span><span style="color:#333333"></span><span style="color:#333333">的</span><span style="color:#333333">TCP</span><span style="color:#333333">包。</span>

TCP FIN探测到主机端口是关闭的：
  <img src="http://my.csdn.net/uploads/201207/01/1341106363_1149.jpg" alt="" /><br />
TCP FIN探测到主机端口是开放或屏蔽的：
   <img src="http://my.csdn.net/uploads/201207/01/1341106375_7898.jpg" alt="" />
#### <a name="t18"></a>2.3.1.5    UDP scanning

UDP扫描方式用于判断UDP端口的情况。向目标主机的UDP端口发送探测包，如果收到回复“ICMP port unreachable”就说明该端口是关闭的；如果没有收到回复，那说明UDP端口可能是开放的或屏蔽的。因此，通过反向排除法的方式来断定哪些UDP端口是可能出于开放状态。

UDP端口关闭：
  <img src="http://my.csdn.net/uploads/201207/01/1341106392_8631.jpg" alt="" /><br />
UDP端口开放或被屏蔽：
  <img src="http://my.csdn.net/uploads/201207/01/1341106404_7106.jpg" alt="" /><br />
#### <a name="t19"></a>2.3.1.6    其他方式

除上述几种常用的方式之外，Nmap还支持多种其他探测方式。例如使用SCTP INIT/COOKIE-ECHO方式来探测SCTP的端口开放情况；使用IP protocol方式来探测目标主机支持的协议类型（TCP/UDP/ICMP/SCTP等等）；使用idle scan方式借助僵尸主机（zombie host，也被称为idle host，该主机处于空闲状态并且它的IPID方式为递增。详细实现原理参见：http://nmap.org/book/idlescan.html）来扫描目标在主机，达到隐蔽自己的目的；或者使用FTP bounce scan，借助FTP允许的代理服务扫描其他的主机，同样达到隐藏自己的身份的目的。

 

### <a name="t20"></a>2.3.2    端口扫描用法

端口扫描用法比较简单，Nmap提供丰富的命令行参数来指定扫描方式和扫描端口。

具体可以参见如下描述。

#### <a name="t21"></a>2.3.2.1    扫描方式选项

<ol start="1">
  <li class="alt">
    <span><span>-sS/sT/sA/sW/sM:指定使用 TCP SYN/Connect()/ACK/Window/Maimon scans的方式来对目标主机进行扫描。  </span></span>
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>  -sU: 指定使用UDP扫描方式确定目标主机的UDP端口状况。  </span>
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>  -sN/sF/sX: 指定使用TCP Null, FIN, and Xmas scans秘密扫描方式来协助探测对方的TCP端口状态。  </span>
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>  --scanflags <flags>: 定制TCP包的flags。  </span>
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>  -sI <zombiehost[:probeport]>: 指定使用idle scan方式来扫描目标主机（前提需要找到合适的zombie host）  </span>
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>  -sY/sZ: 使用SCTP INIT/COOKIE-ECHO来扫描SCTP协议端口的开放的情况。  </span>
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>  -sO: 使用IP protocol 扫描确定目标机支持的协议类型。  </span>
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>  -b <FTP relay host>: 使用FTP bounce scan扫描方式  </span>
  </li>
</ol>

  -sU: 指定使用UDP扫描方式确定目标主机的UDP端口状况。

  -sN/sF/sX: 指定使用TCP Null, FIN, and Xmas scans秘密扫描方式来协助探测对方的TCP端口状态。

  --scanflags &lt;flags>: 定制TCP包的flags。

  -sI &lt;zombiehost[:probeport]>: 指定使用idle scan方式来扫描目标主机（前提需要找到合适的zombie host）

  -sY/sZ: 使用SCTP INIT/COOKIE-ECHO来扫描SCTP协议端口的开放的情况。

  -sO: 使用IP protocol 扫描确定目标机支持的协议类型。

  -b &lt;FTP relay host>: 使用FTP bounce scan扫描方式</pre>
#### <a name="t22"></a>2.3.2.2    端口参数与扫描顺序

<ol start="1">
  <li class="alt">
    <span><span>-p 
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>实例: -p22; -p1-65535; -p U:53,111,137,T:21-25,80,139,8080,S:9（其中T代表TCP协议、U代表UDP协议、S代表SCTP协议）  </span>
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>-F: Fast mode – 快速模式，仅扫描TOP 100的端口  </span>
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>-r: 不进行端口随机打乱的操作（如无该参数，nmap会将要扫描的端口以随机顺序方式扫描，以让nmap的扫描不易被对方防火墙检测到）。  </span>
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>--top-ports <number>:扫描开放概率最高的number个端口（nmap的作者曾经做过大规模地互联网扫描，以此统计出网络上各种端口可能开放的概率。以此排列出最有可能开放端口的列表，具体可以参见文件：nmap-services。默认情况下，nmap会扫描最有可能的1000个TCP端口）  </span>
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>--port-ratio <ratio>: 扫描指定频率以上的端口。与上述--top-ports类似，这里以概率作为参数，让概率大于--port-ratio的端口才被扫描。显然参数必须在在0到1之间，具体范围概率情况可以查看nmap-services文件。  </span>
  </li>
</ol>

实例: -p22; -p1-65535; -p U:53,111,137,T:21-25,80,139,8080,S:9（其中T代表TCP协议、U代表UDP协议、S代表SCTP协议）

-F: Fast mode – 快速模式，仅扫描TOP 100的端口

-r: 不进行端口随机打乱的操作（如无该参数，nmap会将要扫描的端口以随机顺序方式扫描，以让nmap的扫描不易被对方防火墙检测到）。

--top-ports &lt;number>:扫描开放概率最高的number个端口（nmap的作者曾经做过大规模地互联网扫描，以此统计出网络上各种端口可能开放的概率。以此排列出最有可能开放端口的列表，具体可以参见文件：nmap-services。默认情况下，nmap会扫描最有可能的1000个TCP端口）

--port-ratio &lt;ratio>: 扫描指定频率以上的端口。与上述--top-ports类似，这里以概率作为参数，让概率大于--port-ratio的端口才被扫描。显然参数必须在在0到1之间，具体范围概率情况可以查看nmap-services文件。</pre>  

### <a name="t23"></a>2.3.3    端口扫描演示

这里，我们以扫描局域网内192.168.1.100主机为例。

命令如下：

<span style="color:#daeef3; background:black">nmap –sS –sU –T4 –top-ports 300 192.168.1.100</span>

参数-sS表示使用TCP SYN方式扫描TCP端口；-sU表示扫描UDP端口；-T4表示时间级别配置4级；--top-ports 300表示扫描最有可能开放的300个端口（TCP和UDP分别有300个端口）。
  <img src="http://my.csdn.net/uploads/201207/01/1341106491_4668.jpg" alt="" /><br />
从上图中，我们看到扫描结果，横线处写明有共有589端口是关闭的；红色框图中列举出开放的端口和可能是开放的端口。

## <a name="t24"></a>2.4    版本侦测

版本侦测，用于确定目标主机开放端口上运行的具体的应用程序及版本信息。

Nmap提供的版本侦测具有如下的优点：

*   高速。并行地进行套接字操作，实现一组高效的探测匹配定义语法。
*   尽可能地确定应用名字与版本名字。
*   支持TCP/UDP协议，支持文本格式与二进制格式。 
*   支持多种平台服务的侦测，包括Linux/Windows/Mac OS/FreeBSD等系统。
*   如果检测到SSL，会调用openSSL继续侦测运行在SSL上的具体协议（如HTTPS/POP3S/IMAPS）。
*   如果检测到SunRPC服务，那么会调用brute-force RPC grinder进一步确定RPC程序编号、名字、版本号。
*   支持完整的IPv6功能，包括TCP/UDP，基于TCP的SSL。
*   通用平台枚举功能（CPE）
*   广泛的应用程序数据库（nmap-services-probes）。目前Nmap可以识别几千种服务的签名，包含了180多种不同的协议。

### <a name="t25"></a>2.4.1    版本侦测原理

简要的介绍版本的侦测原理。

版本侦测主要分为以下几个步骤：

1.  首先检查open与open|filtered状态的端口是否在排除端口列表内。如果在排除列表，将该端口剔除。
2.  如果是TCP端口，尝试建立TCP连接。尝试等待片刻（通常6秒或更多，具体时间可以查询文件nmap-services-probes中Probe TCP NULL q||对应的totalwaitms）。通常在等待时间内，会接收到目标机发送的“WelcomeBanner”信息。nmap将接收到的Banner与nmap-services-probes中NULL probe中的签名进行对比。查找对应应用程序的名字与版本信息。
3.  如果通过“Welcome Banner”无法确定应用程序版本，那么nmap再尝试发送其他的探测包（即从nmap-services-probes中挑选合适的probe），将probe得到回复包与数据库中的签名进行对比。如果反复探测都无法得出具体应用，那么打印出应用返回报文，让用户自行进一步判定。
4.  如果是UDP端口，那么直接使用nmap-services-probes中探测包进行探测匹配。根据结果对比分析出UDP应用服务类型。
5.  如果探测到应用程序是SSL，那么调用openSSL进一步的侦查运行在SSL之上的具体的应用类型。
6.  如果探测到应用程序是SunRPC，那么调用brute-force RPC grinder进一步探测具体服务。

### <a name="t26"></a>2.4.2    版本侦测的用法

版本侦测方面的命令行选项比较简单。

<ol start="1">
  <li class="alt">
    <span><span>-sV: 指定让Nmap进行版本侦测  </span></span>
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>--version-intensity <level>: 指定版本侦测强度（0-9），默认为7。数值越高，探测出的服务越准确，但是运行时间会比较长。  </span>
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>--version-light: 指定使用轻量侦测方式 (intensity 2)  </span>
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>--version-all: 尝试使用所有的probes进行侦测 (intensity 9)  </span>
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>--version-trace: 显示出详细的版本侦测过程信息。  </span>
  </li>
</ol>

--version-intensity &lt;level>: 指定版本侦测强度（0-9），默认为7。数值越高，探测出的服务越准确，但是运行时间会比较长。

--version-light: 指定使用轻量侦测方式 (intensity 2)

--version-all: 尝试使用所有的probes进行侦测 (intensity 9)

--version-trace: 显示出详细的版本侦测过程信息。</pre>
### <a name="t27"></a>2.4.3    版本侦测演示

命令：

<span style="color:#daeef3; background:black">nmap –sV 192.168.1.100</span>

对主机192.168.1.100进行版本侦测。
  <img src="http://my.csdn.net/uploads/201207/01/1341106542_9477.jpg" alt="" /><br />
从结果中，我们可以看到996个端口是关闭状态，对于4个open的端口进行版本侦测。图中红色为版本信息。红色线条划出部分是版本侦测得到的附加信息，因为从应用中检测到微软特定的应用服务，所以推断出对方运行的Windows的操作系统。

## <a name="t28"></a>2.5    OS侦测

操作系统侦测用于检测目标主机运行的操作系统类型及设备类型等信息。

Nmap拥有丰富的系统数据库nmap-os-db，目前可以识别2600多种操作系统与设备类型。

### <a name="t29"></a>2.5.1    OS侦测原理

Nmap使用TCP/IP协议栈指纹来识别不同的操作系统和设备。在RFC规范中，有些地方对TCP/IP的实现并没有强制规定，由此不同的TCP/IP方案中可能都有自己的特定方式。Nmap主要是根据这些细节上的差异来判断操作系统的类型的。

具体实现方式如下：

1.  Nmap内部包含了2600多已知系统的指纹特征（在文件nmap-os-db文件中）。将此指纹数据库作为进行指纹对比的样本库。
2.  分别挑选一个open和closed的端口，向其发送经过精心设计的TCP/UDP/ICMP数据包，根据返回的数据包生成一份系统指纹。
3.  将探测生成的指纹与nmap-os-db中指纹进行对比，查找匹配的系统。如果无法匹配，以概率形式列举出可能的系统。

### <a name="t30"></a>2.5.2    OS侦测用法

OS侦测的用法简单，Nmap提供的命令比较少。

<ol start="1">
  <li class="alt">
    <span><span>-O: 指定Nmap进行OS侦测。  </span></span>
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>--osscan-limit: 限制Nmap只对确定的主机的进行OS探测（至少需确知该主机分别有一个open和closed的端口）。  </span>
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>--osscan-guess: 大胆猜测对方的主机的系统类型。由此准确性会下降不少，但会尽可能多为用户提供潜在的操作系统。  </span>
  </li>
</ol>

--osscan-limit: 限制Nmap只对确定的主机的进行OS探测（至少需确知该主机分别有一个open和closed的端口）。

--osscan-guess: 大胆猜测对方的主机的系统类型。由此准确性会下降不少，但会尽可能多为用户提供潜在的操作系统。</pre>

### <a name="t31"></a>2.5.3    OS侦测演示

命令：

<span style="color:#daeef3; background:black">nmap –O 192.168.1.100</span>
  <img src="http://my.csdn.net/uploads/201207/01/1341106565_4209.jpg" alt="" /><br />
从上图中可看到，指定-O选项后先进行主机发现与端口扫描，根据扫描到端口来进行进一步的OS侦测。获取的结果信息有设备类型，操作系统类型，操作系统的CPE描述，操作系统细节，网络距离等。

# <a name="t32"></a>3     Nmap高级用法

## <a name="t33"></a>3.1    防火墙/IDS规避

防火墙与IDS规避为用于绕开防火墙与IDS（入侵检测系统）的检测与屏蔽，以便能够更加详细地发现目标主机的状况。

Nmap提供了多种规避技巧，通常可以从两个方面考虑规避方式：数据包的变换（Packet Change）与时序变换（Timing Change）。

### <a name="t34"></a>3.1.1    规避原理

#### <a name="t35"></a>3.1.1.1    分片（Fragmentation）

将可疑的探测包进行分片处理（例如将TCP包拆分成多个IP包发送过去），某些简单的防火墙为了加快处理速度可能不会进行重组检查，以此避开其检查。

#### <a name="t36"></a>3.1.1.2    IP诱骗（IP decoys）

在进行扫描时，将真实IP地址和其他主机的IP地址（其他主机需要在线，否则目标主机将回复大量数据包到不存在的主机，从而实质构成了拒绝服务攻击）混合使用，以此让目标主机的防火墙或IDS追踪检查大量的不同IP地址的数据包，降低其追查到自身的概率。注意，某些高级的IDS系统通过统计分析仍然可以追踪出扫描者真实IP地址。

#### <a name="t37"></a>3.1.1.3    IP伪装（IP Spoofing）

顾名思义，IP伪装即将自己发送的数据包中的IP地址伪装成其他主机的地址，从而目标机认为是其他主机在与之通信。需要注意，如果希望接收到目标主机的回复包，那么伪装的IP需要位于统一局域网内。另外，如果既希望隐蔽自己的IP地址，又希望收到目标主机的回复包，那么可以尝试使用idle scan或匿名代理（如TOR）等网络技术。

#### <a name="t38"></a>3.1.1.4    指定源端口

某些目标主机只允许来自特定端口的数据包通过防火墙。例如FTP服务器配置为：允许源端口为21号的TCP包通过防火墙与FTP服务端通信，但是源端口为其他端口的数据包被屏蔽。所以，在此类情况下，可以指定Nmap将发送的数据包的源端口都设置特定的端口。

#### <a name="t39"></a>3.1.1.5    扫描延时

某些防火墙针对发送过于频繁的数据包会进行严格的侦查，而且某些系统限制错误报文产生的频率（例如，Solaris 系统通常会限制每秒钟只能产生一个ICMP消息回复给UDP扫描），所以，定制该情况下发包的频率和发包延时可以降低目标主机的审查强度、节省网络带宽。

#### <a name="t40"></a>3.1.1.6    其他技术
  Nmap还提供多种规避技巧，比如指定使用某个网络接口来发送数据包、指定发送包的最小长度、指定发包的MTU、指定TTL、指定伪装的MAC地址、使用错误检查和（badchecksum）。

  更多信息http://nmap.org/book/man-bypass-firewalls-ids.html
 

### <a name="t41"></a>3.1.2    规避用法

<ol start="1">
  <li class="alt">
    <span><span>-f; --mtu <val>: 指定使用分片、指定数据包的MTU.  </span></span>
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>-D <decoy1,decoy2[,ME],...>: 用一组IP地址掩盖真实地址，其中ME填入自己的IP地址。  </span>
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>-S <IP_Address>: 伪装成其他IP地址  </span>
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>-e <iface>: 使用特定的网络接口  </span>
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>-g/--source-port 
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>--data-length <num>: 填充随机数据让数据包长度达到Num。  </span>
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>--ip-options <options>: 使用指定的IP选项来发送数据包。  </span>
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>--ttl <val>: 设置time-to-live时间。  </span>
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>--spoof-mac <mac address/prefix/vendor name>: 伪装MAC地址  </span>
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>--badsum: 使用错误的checksum来发送数据包（正常情况下，该类数据包被抛弃，如果收到回复，说明回复来自防火墙或IDS/IPS）。  </span>
  </li>
</ol>

-D &lt;decoy1,decoy2[,ME],...>: 用一组IP地址掩盖真实地址，其中ME填入自己的IP地址。

-S &lt;IP_Address>: 伪装成其他IP地址

-e &lt;iface>: 使用特定的网络接口

-g/--source-port &lt;portnum>: 使用指定源端口

--data-length &lt;num>: 填充随机数据让数据包长度达到Num。

--ip-options &lt;options>: 使用指定的IP选项来发送数据包。

--ttl &lt;val>: 设置time-to-live时间。

--spoof-mac &lt;mac address/prefix/vendor name>: 伪装MAC地址

--badsum: 使用错误的checksum来发送数据包（正常情况下，该类数据包被抛弃，如果收到回复，说明回复来自防火墙或IDS/IPS）。</pre>
### <a name="t42"></a>3.1.3    规避演示

使用命令：
  <span style="color:white">nmap -v -F -Pn -D192.168.1.100,192.168.1.102,ME -e eth0 -g 3355 192.168.1.1</span>
其中，-F表示快速扫描100个端口；-Pn表示不进行Ping扫描；-D表示使用IP诱骗方式掩盖自己真实IP（其中ME表示自己IP）；-e eth0表示使用eth0网卡发送该数据包；-g 3355表示自己的源端口使用3355；192.168.1.1是被扫描的目标IP地址。
  <img src="http://my.csdn.net/uploads/201207/01/1341106610_5800.jpg" alt="" /><br />
我们可以从Wireshark中看到数据包的流动情况：对于每个探测包，Nmap都使用-D选项指定的IP地址发送不同的数据包，从而达到扰乱对方防火墙/IDS检查的目的（更好的方式-D选项中嵌入RND随机数，这样更具有迷惑性）。当探测到80端口时候，目标主机向我们回复了SYN/ACK包回来（当然也向其他诱骗的IP回复SYN/ACK包，我们无法接收到），证明80端口是开放的。


   <img src="http://my.csdn.net/uploads/201207/01/1341106628_5988.jpg" alt="" />
## <a name="t43"></a>3.2    NSE脚本引擎

NSE脚本引擎（Nmap Scripting Engine）是Nmap最强大最灵活的功能之一，允许用户自己编写脚本来执行自动化的操作或者扩展Nmap的功能。

NSE使用Lua脚本语言，并且默认提供了丰富的脚本库，目前已经包含14个类别的350多个脚本。

NSE的设计初衷主要考虑以下几个方面：

*   网络发现（Network Discovery）
*   更加复杂的版本侦测（例如skype软件）
*   漏洞侦测(Vulnerability Detection)
*   后门侦测(Backdoor Detection)
*   漏洞利用(Vulnerability Exploitation)

 

### <a name="t44"></a>3.2.1    NSE创建脚本方法

下面以daytime.nse脚本为例说明一下NSE格式。
  <img src="http://my.csdn.net/uploads/201207/01/1341106665_4408.jpg" alt="" /><br />
NSE的使用Lua脚本，并且配置固定格式，以减轻用户编程负担。通常的一个脚本分为几个部分：

description字段：描述脚本功能的字符串，使用双层方括号表示。

comment字段：以--开头的行，描述脚本输出格式

author字段：描述脚本作者

license字段：描述脚本使用许可证，通常配置为Nmap相同的license

categories字段：描述脚本所属的类别，以对脚本的调用进行管理。

**rule字段**：描述脚本执行的规则，也就是确定触发脚本执行的条件。在Nmap中有四种类型的规则，prerule用于在Nmap没有执行扫描之前触发脚本执行，这类脚本并不需用到任何Nmap扫描的结果；hostrule用在Nmap执行完毕主机发现后触发的脚本，根据主机发现的结果来触发该类脚本；portrule用于Nmap执行端口扫描或版本侦测时触发的脚本，例如检测到某个端口时触发某个脚本执行以完成更详细的侦查。postrule用于Nmap执行完毕所有的扫描后，通常用于扫描结果的数据提取和整理。在上述实例中，只有一个portrule，说明该脚本在执行端口扫描后，若检测到TCP 13号端口开放，那么触发该脚本的执行。

**action字段**：脚本执行的具体内容。当脚本通过rule字段的检查被触发执行时，就会调用action字段定义的函数。

### <a name="t45"></a>3.2.2    NSE脚本用法

Nmap提供不少脚本使用的命令行参数。

<ol start="1">
  <li class="alt">
    <span><span>-sC: 等价于 --script=default，使用默认类别的脚本进行扫描。  </span></span>
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>--script=<Lua scripts>: <Lua scripts>使用某个或某类脚本进行扫描，支持通配符描述  </span>
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>--script-args=<n1=v1,[n2=v2,...]>: 为脚本提供默认参数  </span>
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>--script-args-file=filename: 使用文件来为脚本提供参数  </span>
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>--script-trace: 显示脚本执行过程中发送与接收的数据  </span>
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>--script-updatedb: 更新脚本数据库  </span>
  </li>
  <li class="">
    <span>  </span>
  </li>
  <li class="alt">
    <span>--script-help=<Lua scripts>: 显示脚本的帮助信息，其中<Luascripts>部分可以逗号分隔的文件或脚本类别。  </span>
  </li>
</ol>

--script=&lt;Lua scripts>: &lt;Lua scripts>使用某个或某类脚本进行扫描，支持通配符描述

--script-args=&lt;n1=v1,[n2=v2,...]>: 为脚本提供默认参数

--script-args-file=filename: 使用文件来为脚本提供参数

--script-trace: 显示脚本执行过程中发送与接收的数据

--script-updatedb: 更新脚本数据库

--script-help=&lt;Lua scripts>: 显示脚本的帮助信息，其中&lt;Luascripts>部分可以逗号分隔的文件或脚本类别。</pre>
### <a name="t46"></a>3.2.3    NSE用法演示

配合脚本扫描192.168.1.1，查看能否获得有用的信息。

命令如下：

<span style="color:#daeef3; background:black">nmap –sV –p 80 –v –script default,http*192.168.1.1</span>
  <img src="http://my.csdn.net/uploads/201207/01/1341106687_9192.jpg" alt="" /><br />
从上图中，我们可以看到Nmap扫描到对方80端口是开放的，然后使用了大量的名字为http开头的脚本对其进行扫描。扫描过程发现在http-auth脚本执行，出现了“Basic relm=TP-LINK Wireless N router WR740”字样（红线划出部分），这里已经挖掘对方的设备类型与具体版本信息。如果我们知道更多关于WR740已知的漏洞，那么就可以进行更进一步的渗透测试了。

# <a name="t47"></a>4     参考资料

## <a name="t48"></a>4.1    书籍

**Nmap Network Scanning**

Nmap创始人Fyodor编写的Nmap的权威指南，非常详尽地描述Nmap的实现原理及使用方法。Nmap官方文档正是来自该书部分章节。

[**Secrets of Network Cartography**][4]

该书对Nmap的实现原理及使用场景有比较丰富的介绍。

**Nmap in the Enterprise: Your Guide to Network Scanning**

这本书描述Nmap在企业领域的运用。**  
**

[**Nmap mindmap.pdf**][5]

这nmap使用方法的思维导图（一页纸的图片），对Nmap用法整理很完整。

** **

## <a name="t49"></a>4.2    网站

官网：[www.nmap.org][2]

安全工具排名：<http://sectools.org/>

 

 [1]: http://aspirationflowspace.googlecode.com/files/Nmap%E6%89%AB%E6%8F%8F%E5%8E%9F%E7%90%86%E4%B8%8E%E7%94%A8%E6%B3%95.pdf
 [2]: http://www.nmap.org
 [3]: http://www.163.com
 [4]: http://aspirationflowspace.googlecode.com/files/Syngress%20-%20Secrets%20of%20Network%20Cartography%20-%20A%20Comprehensive%20Guide%20to%20Nmap%20-%20Complete.pdf
 [5]: http://aspirationflowspace.googlecode.com/files/nmap-mindmap.pdf

- - -- - - -- - - - - -- - - - -
这里是 10 条 nmap 的技巧，运行于 CLI 环境下，如果你更喜欢用 GUI 工具的话，请用 Zenmap 。

1) 获取远程主机的系统类型及开放端口

    nmap -sS -P0 -sV -O <target>

这里的 < target > 可以是单一 IP, 或主机名，或域名，或子网

    -sS TCP SYN 扫描 (又称半开放,或隐身扫描)
    -P0 允许你关闭 ICMP pings.
    -sV 打开系统版本检测
    -O 尝试识别远程操作系统

其它选项:

    -A 同时打开操作系统指纹和版本检测
    -v 详细输出扫描情况.
    nmap -sS -P0 -A -v < target >

2) 列出开放了指定端口的主机列表

    nmap -sT -p 80 -oG – 192.168.1.* | grep open

3) 在网络寻找所有在线主机

    nmap -sP 192.168.0.*
    
或者也可用以下命令:

    nmap -sP 192.168.0.0/24

指定 subnet

4) Ping 指定范围内的 IP 地址

    nmap -sP 192.168.1.100-254

5) 在某段子网上查找未占用的 IP

    nmap -T4 -sP 192.168.2.0/24 && egrep "00:00:00:00:00:00" /proc/net/arp

6) 在局域网上扫找 Conficker 蠕虫病毒

    nmap -PN -T4 -p139,445 -n -v --script=smb-check-vulns --script-args safe=1 192.168.0.1-254

7) 扫描网络上的恶意接入点 （rogue APs）.

    nmap -A -p1-85,113,443,8080-8100 -T4 --min-hostgroup 50 --max-rtt-timeout 2000 --initial-rtt-timeout 300 --max-retries 3 --host-timeout 20m --max-scan-delay 1000 -oA wapscan 10.0.0.0/8

8 ) 使用诱饵扫描方法来扫描主机端口

    sudo nmap -sS 192.168.0.10 -D 192.168.0.2

9) 为一个子网列出反向 DNS 记录

    nmap -R -sL 209.85.229.99/27 | awk '{if($3=="not")print"("$2") no PTR";else print$3" is "$2}' | grep '('

10) 显示网络上共有多少台 Linux 及 Win 设备?
    sudo nmap -F -O 192.168.0.1-255 | grep "Running: " > /tmp/os; echo "$(cat /tmp/os | grep Linux | wc -l) Linux device(s)"; echo "$(cat /tmp/os | grep Windows | wc -l) Window(s) device"

