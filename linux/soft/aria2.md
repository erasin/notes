#aria2
**title:** aria2 使用说明  
**title:** aria2,aria2c,wget,linux,下载  
**info:** **aria2**是一款轻量型命令行下载工具，它提供了对多协议和多源地址的支持，目前支持的协议包括**HTTP(S), FTP, BitTorrent (DHT, PEX, MSE/PE), and Metalink**。

aria2可以从多个源地址，并使用多种协议进行下载，并尝试将下载带宽利用率最大化。它可以同时从HTTP(S)/FTP 和 BitTorrent下载一份数据，并且将其上传到bt集群中。通过Metalink的分块检查，aria2可以在下载过程中自动的进行数据校验。

虽然现在有诸如**wget**和**curl**等其他类似产品，但aria2具有两个独特的功能：

1. aria2可以从多个源下载文件(HTTP(S)/FTP/BitTorrent)，
2. aria2可以并发的进行多个源地址的下载。这样用户将不必等待单个文件的下载完成，而且aria2会尽可能快的下载。

除aria2外，也有一些其他可以进行分片下载的工具，它们往往按照线程数分割文件，并行下载，换言之，它们不会对未完成的部分进行自适应性重新分片，当整个流程工作正常时，这个策略是可以的，但一旦存在一个线程运行非常慢，这样整个进程就需要等待该线程的执行。而aria2可以很好的处理这个情况，它可以将文件分割成1M大小的分片，当某个线程运行特别缓慢时，aria2可以使用更快的线程来替换。总之，根据作者的说法，aria2是非常智能和可靠的。

和最初的aria工具具有一个GTK+界面不同，aria2只提供了命令行接口，从而使得对资源的要求更小。通常它的物理内存消耗为4M(HTTP/FTP)到9M(BitTorrent),当进行bt下载且速度为2.8M/s时的cpu消耗约为6%。

## 基本使用

当源地址存在诸如`&`或`*`等其他shell特殊字符时，请使用单引号或者双引号将uri包含起来。

在aria2的1.10.0版中，aria2对每台主机默认使用1个连接和20MB的分块大小，所以不论在-s参数中指定任何值，他都对一台主机只会建立一个连接，需要注意的一点是这个限制是针对单个主机的，当指定了多台主机时，它会对每台主机建立一个连接。如果要使用1.9.x版中的行为，则要使用-x16 -k1M。另见man页面的 `–max-connection-per-server` 和 `–min-split-size`。

_aria2默认会在开始下载前预先分配文件空间以避免可能的文件碎片_，但这会在部分PC上带来50%-90%的cpu消耗。当使用环境为比较新的文件系统，例如ext4，btrfs，xfs或者NTFS时，作者推荐使用**`–file-allocation=falloc`**，这种方式会在瞬间完成大文件（数G）的空间分配并且不会带来额外的性能下降。

如果你既没有使用cutting-edge文件系统，也没有使用linux，并且很在意系统性能，那么可以使用–file-allocation=none来关闭文件的预分配。

如果你忘记了这些参数的全名或者含义，把-h放在选项或者词的前面，比如aria2c -hcrypt，这样aria2就会搜索crypt相关的选项并把他的帮助打印出来，如果aria2发现使用了-h，它会在打印出帮助信息后停止运行。

## 基本用法

### 下载一个文件

    aria2c http://host/image.iso

在1.10.0版中，aria2对每个host使用一个连接，你可以使用`–max-connection-per-server` 或 `-x`来修改

### 使用两个连接从一个源下载文件

    aria2c -x2 http://host/image.iso

`Ctrl+c`可以中止当前的下载，在同样的目录运行同样的命令可以继续原来的下载，你设置可以修改`uri`，如果他们指向的是同一个文件的话。

### 使用两个连接下载文件：

    aria2c -s2 http://host/image.iso http://mirror1/image.iso http://mirror2/image.iso

注：如果命令中uri的数量多于-s的值，如本例所示，aria2将会首先使用前两个uri，将第3个uri作为候补，如果前两个有一个下载失败，就会启用第3个

### 从http或ftp服务器下载：

    aria2c http://host1/file.zip ftp://host2/file.zip

### 从任意源进行并行下载：

    aria2c -Z http://host/file1 file2.torrent file3.metalink

### 如果仅仅指定本地磁盘上的torrent文件或者metalink，是不需要-Z选项的，如：

    aria2c file1.torrent file2.torrent

### 从文件中读取目的文件，并行下载

    aria2c -ifiles.txt -j5

注：-j参数指定了并发下载的数量，在输入文件中可以包含torrent文件和metelink  
注：输入文件中支持添加参数，详见后续的“输入文件”章节

### 退出时保存出错/未完成的下载, 使用`session`

    aria2c -ifiles.txt --save-session=out.txt

当ctrl+c或者aria2自己退出时，所有的错误，未完成的下载信息会保存到out.txt中，但通过aria2.addTorrent 和 aria2.addMetalink XML-RPC方式增加的下载不会保存。

### 后续可以使用该文件继续未完成的下载：

    aria2c -i out.txt

## Metalink相关下载示例

### 从远程metalink下载文件：

    aria2c http://host/file.metalink

### 从远程metalink下载文件,但在内存中处理metalink：

    aria2c --follow-metalink=mem http://host/file.metalink

### 通过本地metalink下载：

    aria2c -t10 --lowest-speed-limit=4000 file.metalink

### 使用5个server下载

    aria2c -C5 file.metalink

注：当使用metalink时，-s参数不再起作用，需要使用-C选项

### 通过多个本地metalink文件进行下载：

    aria2c file1.metalink file2.metalink

### 打印metalink的内容

    aria2c -S file.metalink

### 通过序号下载指定文件

    aria2c --select-file=1-4,8 -Mfile.metalink

注：可以通过-S选择来打印metalink文件的内容

### 指定用户偏好从本地metalink下载文件

    aria2c --metalink-location=JP,US --metalink-version=1.1 --metalink-language=en-US file.metali

## BitTorrent相关下载

### 通过远程BitTorrent文件下载

    aria2c http://site/file.torrent

### 通过远程BitTorrent文件下载,但在内存中处理

    aria2c --follow-torrent=mem http://site/file.torrent

### 通过本地torrent文件下载:

    aria2c -u40K /path/to/file.torrent

注：-u, –max-upload-limit用来指定最大上传速度

### 可以同时处理多个torrent文件:

    aria2c /path/to/file1.torrent /path/to/file2.torrent

### 通过BitTorrent Magnet URI下载：

    aria2c "magnet:?xt=urn:btih:248D0A1CD08284299DE78D5C1ED359BB46717D8C&dn=aria2"

注：需要将magnet的地址用单引号或者双引号引起来，因为里面包含’&'字符。当使用magnet时，强烈建议打开dht选项，–enable-dht
将metadata保存成.torrent文件

    aria2c --bt-save-metadata "magnet:?xt=urn:btih:248D0A1CD08284299DE78D5C1ED359BB46717D8C&dn=aria2"

这个命令会将metadata信息保存在248d0a1cd08284299de78d5c1ed359bb46717d8c.torrent文件中

### 自动调整peer节点数量

当所有节点的整体下载速度低于某个值时，aria2可以临时增加节点数量来获取更快的下载速率

    aria2c --bt-request-peer-speed-limit=200K file.torrent

### 开启DHT

    aria2c --enable-dht http://site/file.torrent

注：在1.7.2版本后，dht选项是被默认打开的。当aria2只处理http/ftp下载时，不会打开dht选项。当第一个torrent开始下载时，aria2进行DHT功能的初始化，然后一直运行到aria2退出。

### 开启ipv6的DHT

    aria2c --enable-dht6 --dht-listen-port=6881 --dht-listen-addr6=YOUR_GLOBAL_UNICAST_IPV6_ADDR --enable-async-dns6

注：如果aria2在build时没有使用c-ares，则不需要–enable-async-dns6。aria2在ipv4和ipv6的dht中共享一些端口

### 增加和删除tracker URI：

下面这个例子将从file.torrent中移除所有的tracker的uri，然后使用”http://tracker1/announce” 和 “http://tracker2/announce”

    aria2c --bt-exclude-tracker="*" --bt-tracker="http://tracker1/announce,http://tracker2/announce" file.torrent

### 加密

在默认情况下，aria2可以接收加密/非加密的连接，并且会首先尝试加密握手，如果失败才会去使用传统的BitTorrent握手
下面这个例子中，aria2将只通过加密的握手接受与建立连接。

    aria2c --bt-require-crypto=true http://site/file.torrent

存在两种加密方式，头加密和全连接加密。如果两种都可以被peer提供，aria2将默认使用头加密方式。如果要使用全连接加密，可以使用：

    aria2c --bt-min-crypto-level=arc4 http://site/file.torrent

### 打印torrent文件内容

    aria2c -S file.torrent

### 通过序号选择指定文件下载：

    aria2c --select-file=1-4,8 -Tfile.torrent

注：index信息可以通过-S来获得

### 修改监听端口

    aria2c --listen-port=6881-6883 file.torrent

注：请确保指定端口可以进行tcp的上行和下行通信

### 指定aria2完成下载后的停止条件

    aria2c --seed-time=120 --seed-ratio=1.0 file.torrent

注：本例中，aria2会在下载完成后120分钟，或seed ratio达到1.0时退出程序

### 设置上传速度

    aria2c --max-upload-limit=100K file.torrent

### Seeding已经下载完成的文件

可以使用-V选择来播种下载完成的文件，它会首先对文件进行分片的hash校验

    Seeding already downloaded file

如果可以确定下载文件的正确性，可以使用–bt-seed-unverified选项来跳过文件的校验环节

    aria2c --bt-seed-unverified -d/path/to/dir file.torrent

### 还可以同时播种多个torrent文件

    aria2c --bt-seed-unverified -d/path/to/dir file1.torrent file2.torrent

### 通过index指定文件名

为了指定bt下载的文件名称，需要使用-S选线来查看torrent文件中的index信息，例如：

    idx|path/length
    ===+======================
      1|dist/base-2.6.18.iso
       |99.9MiB
    ---+----------------------
      2|dist/driver-2.6.18.iso
       |169.0MiB
    ---+----------------------

下面的命令可以将dist/base-2.6.18.iso 保存为 /tmp/mydir/base.iso，同时将dist/driver-2.6.18.iso 保存成 /tmp/dir/driver.iso

    aria2c --dir=/tmp --index-out=1=mydir/base.iso --index-out=2=dir/driver.iso file.torrent

注：当对http uri中的torrent文件使用–index-out是不成功，它只对本地的torrent文件有效。aria2也不会去记忆–index-out选项内容，在每次的继续下载或播种时都需要手工指明，而且如果没有该选项，它也不会给用户任何提醒，所以需要注意。

### 为进行文件预览进行分片优先下载
优先下载torrent中所有文件的前1MB

    aria2c --bt-prioritize-piece=head file.torrent


##使用代理服务器

### 为所有协议(HTTP(S)/FTP)设定代理服务器

    aria2c --all-proxy='http://proxy:8080' http://host/file

注：–all-proxy选项可以被特定的协议选项覆盖，如–http-proxy, –https-proxy 和 –ftp-proxy

### 设置http代理

    aria2c --http-proxy='http://proxy:8080' http://host/file

### 使用需要认证的代理

    aria2c --http-proxy='http://proxy:8080' --http-proxy-user='username' --http-proxy-passwd='password' http://host/file
    aria2c --http-proxy='http://username:password@proxy:8080' http://host/file

注：username and password 需要进行%转码，如过username是’myid@domain’，则转码后的结果为：’myid%40domain’。

##高级的http特性：

### 导入cookie  
导入Mozilla/Firefox(1.x/2.x) 和 Netscape格式的cookie

    aria2c --load-cookies=cookies.txt http://host/file

导入Firefox3格式的cookie

    aria2c --load-cookies=cookies.sqlite http://host/file

被浏览器或其他程序继续下载

    aria2c -c -s2 http://host/partiallydownloadedfile.zip

##其他高级特性

### 下载速度限制

    aria2c --max-download-limit=100K http://host/file

### 使用-V修复受损的下载

    aria2c -V file.metalink

注：这个选项只适用于BitTorrent或者带有校验的metalink

### 当下载速度下于某个特定值时放弃下载

    aria2c --lowest-speed-limit=10K file.metalink

### URI的参数化支持

可以使用大括号来表达一组列表

    aria2c -P http://{host1,host2,host3}/file.iso

可以使用[]来表示一个数字序列

    aria2c -Z -P http://host/image[000-100].png

注：当所有的URI指向不同的文件时，需要使用-Z选项

甚至可以指定步长

    aria2c -Z -P http://host/image[A-Z:2].png

### 时间戳

保留源文件时间戳

    aria2c -R http://host/file.iso

下载完成后执行特定命令

    aria2c --on-download-complete=COMMAND http://example.org/file.iso

另见：`–on-download-error`, `–on-download-start` 和 `–on-download-stop`，或者执行`aria2c -h#hook`

写入/dev/null

    aria2c -d /dev -o null --allow-overwrite=true http://example.org/file

`–allow-overwrite=true`是为了避免aria2重命名已有的/dev/null

### 输入文件

输入文件中可以包括一系列的URI地址，也可以针对同一个资源设置多个URI：不同的URI写在一行并使用tab分割。  
输入文件的每一行都被当作aria2的命令行参数，因此可以受到-Z和-P选项影响  
另外，选项也可以在每行URI的后面设置，更详尽的描述可以在man page的输入文件部分得到。这种选择的含义与命令行选项的含义一致，不过只适用在他们跟随的那个URI上。

如下，一个uri文件名为uri.txt，其内容如下：

    http://server/file.iso http://mirror/file.iso
      dir=/iso_images
      out=file.img

    http://foo/bar

如果aria2使用 `-i uri.txt -d /tmp`,   
那么file.iso就会被保存为/iso\_images/file.img，
它的下载源为http://server/file.iso 和 http://mirror/file.iso ；
而bar文件则是从http://foo/bar 下载并且保存为 /tmp/bar

*[uri]: 远程或者本地文件地址
