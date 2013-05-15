# mplayer


下载电影的时候，我们总希望在全部下载完成之前能够预览一下影片内容，于是发布者时常会放一些影片截图在种子文件中，或者直接贴到网上，也有一些截图是一张图片，但包含很多幅影片在一起，就像下面这张：


有很多软件能够截取影片图像、合并图像，但如果影片太多，比如视频网站为用户上传的图像生成预览图之类的，人工在gui方式下操作就不可取了，我们需要在命令行方式下来截取、合并。



首先，截取影片图像使用最多的就是mplayer或者ffmpeg，我用mplayer比较熟，本文就以此为例了，ffmpeg功能也是非常强大的，但据说支持的文件格式却不丰富。mplayer截取影片图像的基本命令为：

    mplayer -ss START_TIME -noframedrop -nosound -vo jpeg -frames N NAME_OF_VIDEO_FILE   

上例中，-ss指定开始的时间，结合-frames参数，限定从某个时间开始、截取几帧图像。为了体现整个影片的内容，我需要在影片中间隔时间相同的几个点、每个点截取1帧图像，所以按道理应该用-frames 1，但是mplayer这样截图的情况下，第一帧似乎永远都会截取到一个黑屏，所以我常用-frames 2。截取下来的图像保存在了当前目录，名称从00000001.jpg开始依次递增，按照-frames 2，就是取00000002.jpg为结果，删除00000001.jpg即可。经过简单实验，在截取wmv、rmvb影片时，前面的好几帧都会是黑屏，也只能参考上面的做法多取几帧了。

为了取影片中间隔大致相同的几个点，可以用-ss指定时间，也可以用-sb指定开始字节，在我的实际使用中，使用-sb只会得到黑屏，所以通过文件大小来设置间隔点的办法不行，只能用-ss时间间隔了，这就需要首先得到影片的总时间。好在mplayer为我们提供了类似的功能：

    mplayer -identify movie-filename -nosound -vc dummy -vo null   

这样会输出一大堆影片信息，从中截取所需内容即可，在bash shell中，取得影片总时间长度（以秒为单位）的命令如下：

    FILESIZE=`mplayer -identify -nosound -vc dummy -vo null $1 | grep ID_LENGTH | sed -r 's/ID_LENGTH=(:digit:*)(.:digit:*)?/1/g'`   

有了影片的总时长，我们就可以根据所要截取的帧数，计算出每个间隔点的时间位移了。不过要注意一般影片的开始`-ss 0`和结束`-ss TOTAL_TIME_OF_VIDEO`截取下来都会是黑屏，在处理的时候要分别加上和减去若干秒。

截取工作完成后，我们拥有了一堆000000xx.jpg文件，如果能把这些文件都放到一个文件中，每行2张，成为一张大图片，在发布的时候会很方便。所以，我们使用imagemagick(<http://www.imagemagick.org/script/index.php)%E4%B8%AD%E7%9A%84montage%E5%91%BD%E4%BB%A4%E6%9D%A5%E5%AE%9E%E7%8E%B0>：

    montage -geometry +0+0 -tile 2 *.jpg montage.jpg   

`-geometry +0+0`是设定使用原始图片大小，-tile 2参数设定每行放2张图片，最后一个参数是要生成的目标文件名，现在，我们就能够得到像刚才那张一样的图片了。

原理已经讲清楚了，可以自己写一个bash脚本来方便调用，我在网上找到了一个很不错的例子(<http://www.linuxquestions.org/questions/showthread.php?t=361072)%EF%BC%8C%E5%8F%AF%E4%BB%A5%E5%9C%A8%E8%BF%99%E4%B8%AA%E5%9F%BA%E7%A1%80%E4%B8%8A%E8%BF%9B%E8%A1%8C%E4%BF%AE%E6%94%B9%EF%BC%8C%E8%BF%87%E7%A8%8B%E4%B8%8D%E5%86%8D%E8%AF%A6%E8%BF%B0%E4%BA%86>。

下面再列一些在网上找到的其他mplayer、mencoder、ffmpeg的使用实例：

mplayer获取影片信息

    mplayer -identify movie-filename -nosound -vc dummy -vo null

从所有输出中可以grep到如下信息：

     - filetype: ASF file format detected.
     - dimensions and format: VIDEO: [MP43] 320×240 24bpp 1000.000 fps 0.0 kbps ( 0.0 kbyte/s)
     - video format: ID_VIDEO_FORMAT=MP43
     - width (dimensions): ID_VIDEO_WIDTH=320
     - height (dimensions): ID_VIDEO_HEIGHT=240
     - length in seconds: ID_LENGTH=98.00

参考8(<http://gallery.menalto.com/node/40548>)

## mencoder图片做成电影

    #用当前目录中的所有JPEG文件创建DivX4文件：
    mencoder *.jpg -mf on:w=800:h=600:fps=25 -ovc divx4 -o output.avi
    #用当前目录中的一些JPEG文件创建DivX4文件：
    mencoder -mf on:w=800:h=600:fps=25 -ovc divx4 -o output.avi *.jpg
    #用当前目录中的所有JPEG文件创建Motion JPEG(MJPEG)文件：
    mencoder -mf on:w=800:h=600:fps=25 -ovc copy -o output.avi *.jpg
    #用当前目录中的所有PNG文件创建一个非压缩的文件：
    mencoder -mf on:w=800:h=600:fps=25:type=png -ovc rawrgb -o output.avi *.png

简单用法：

    mencoder *.jpg -mf on:fps=15 -o output.avi -ovc xvid

参考6 参考7(<http://huangjiahua.livejournal.com/99358.html>)

## ffmpeg屏幕录像

    ffmpeg -vcodec mpeg4 -b 1000 -r 10 -g 300 -vd x11:0,0 -s 1024×768 ~/test.avi

其中，`-vd x11:0,0` 指录制所使用的偏移为 x=0 和 y=0，-s 1024×768 指录制视频的大小为 1024×768。录制的视频文件为 test.avi，将保存到用户主目录中。其他选项可查阅其说明文档。

如果你只想录制一个应用程序窗口或者桌面上的一个固定区域，那么可以指定偏移位置和区域大小。使用xwininfo -frame命令可以完成查找上述参数。

你也可以重新调整视频尺寸大小，如：

    ./ffmpeg -vcodec mpeg4 -b 1000 -r 10 -g 300 -i ~/test.avi -s 800×600 ~/test-800-600.avi。

参考5(<http://linuxtoy.org/archives/ffmpeg.html>)

## mplayer对video进行截屏 截图(wmv mpeg mov flv all works)

    mplayer 78.mov -ss 1 -nosound -vo jpeg:outdir=./ -frames 2

我截的第一张图不知为何全部都是黑屏

参考4(<http://www.linuxfans.org/nuke/modules.php?name=Forums&file=viewtopic&t=165254>)

## 转换为flv文件

    mencoder NOW.wmv -ffourcc FLV1 -of lavf -ovc lavc -lavcopts vcodec=flv:acodec=mp3:abitrate=56 -srate 22050 -oac mp3lame -o NOW.flv

    ffmpeg -i a.asf -ab 56 -ar 22050 -b 500 -r 15 -s 320×240 asf.flv

参考3(<http://www.roading.net/blog/article.asp?id=114>)

## 使用ffmpeg抓图

    ffmpeg -i test2.asf -y -f image2 -ss 08.010 -t 0.001 -s 352×240 b.jpg
    jpg: ffmpeg -i test.asf -y -f image2 -t 0.001 -s 352×240 -ss a.jpg //注意-ss就是要提取视频文件中指定时间的图像
    jpg: ffmpeg -i asf.flv -y -f image2 -t 1 asf.jpg
    gif: ffmpeg -i test.asf -vframes 30 -y -f gif a.gif

参考3 参考2(<http://www.killflash.net/blog/article.asp?id=77>)


## 如何合并几个视频片段

    mencoder -oac copy -ovc copy -idx -o output.avi video1.avi video2.avi video3.avi

* 其中，-oac copy 选项告诉 mencoder 要正确拷贝音频流。而 -ovc copy 选项则是拷贝视频流。
* 如果在视频文件中没有找到索引的话，那么 -idx 选项会要求 mencoder 建立它。
* -o 选项指定输出文件的名称。
* 最后几个参数为需要合并的几个视频片段。

参考1 (<http://linuxtoy.org/archives/join_several_videos.html>)

## 大杂烩

### 服务器端转换工具(Server-Side-FLV-Conversion)

场景:想把 MPG 或 AVI 上传到你的服务器并自动转换成 FLV 吗?

1. FFmpeg (<http://sourceforge.net/projects/ffmpeg>) | 教程一 (<http://soenkerohde.com/tutorials/ffmpeg>) | 教程二 (<http://klaus.geekserver.net/flash/streaming.html)(Google> Video 使用的就是这个东东.)
2. Flix Engine (<http://www.on2.com/developer/flix-engine-sdk>) | 教程 (<http://www.flexauthority.com/articlesIndex.cfm>) | 范例 (<http://www.flexauthority.com/Samples/FlixEngine/index.html>)
3. Turbine Video Engine (<http://www.blue-pacific.com/products/turbinevideosdk/default.htm>)
4. Video to Flash Console (<http://www.geovid.com/Video_to_Flash_Console>)

### 录像/实时广播(Record/Broadcast)

场景:想制作一个语音视频Blog满足自恋的欲望吗？

1 ,RED5 (<http://www.osflash.org/red5>)
2, Flash Media Server (<http://www.macromedia.com/go/fms>)


在线编码,分享视频(Online Encode & Share)

场景:想不花钱就可以在线分享你的视频吗?

1, Google Video (<http://video.google.com/>)
2, You Tube (<http://www.youtube.com/>)

本地 FLV 文件播放器(FLV Player)

场景:拿到了 FLV 文件不知道怎么播放了．

1, martijndevisser FLV Player (<http://www.martijndevisser.com/2005/10/flv_player_updated.html>)
2, FlashGuru FLV Player (<http://www.flashguru.co.uk/free-tool-flash-video-player>)
3, FCZone FLV Player (<http://fczone.com/2006/01/fms-media-player.cfm>)

在线 FLV 文件播放器(Online FLV Player)

场景:知道一个在线FLV地址,又懒得下载和安装播放器．

1,Loadr (<http://dengjie.com/loadr>)
2,Google Player Generator (<http://dengjie.com/loadr/r.swf?file=/temp/google_player.swf&clr=000FFF>)
更多相关软件看这篇文章:Flash 网站的视频策略 (<http://www.macromedia.com/cfusion/knowledgebase/index.cfm?id=tn_14571>)

此文章转自 shadow

## ffmpeg 参数

利用ffmpeg+mencoder视频转换的总结
<http://www.yitian130.com/article.asp?id=69>

flv视频转换和flash播放的解决方案笔记
<http://blog.verycd.com/dash/showentry=35982>

## Youtube技术原理

1. 网页文件允许上传视频文件（这个和上传其他文件一样的）（作者用的是python的架构）
2. 后台调用ffmpeg对上传的视频进行压缩，输出flv格式的文件。这个开源程序win32和linux都有实现，所以可以适应不用的主机环境。
3. 使用flvtools处理flv文件，标记上时长、帧速、关键帧等元数据，这样的flash文件才可以拖放。
4. 使用 ffmpeg 产生flv文件的缩略，和大图像文件的缩略图是一个道理。
5. 使用适当的flv播放器在网页中播放服务器端生成的flv文件。

更多详细：<http://www.gotonx.com/bbs/simple/index.php?t6322.html>

安装和使用ffmpeg转换视频为flv文件（windows和linux）

1、环境winxp-sp2下：
从 <http://ffdshow.faireal.net/mirror/ffmpeg/> 下载
最新版本的 FFMpeg.exe直接用就行（须rar解压）。

以下的东西是为对ffmpeg无法解析的文件格式(wmv9，rm，rmvb等)转换用的,
[从http://mediacoder.sourceforge.net/download_zh.htm下载]()
最新版本的mediacoder的安装后；找到其中的mencoder.exe；drv43260.dll和pncrt.dll三个文件。

2、环境linuxas4。3下：

a、先装mp3在linux下的包：lame-3.97.tar.gz；

    tar -xvzf lame-3.97.tar.gz;
    cd lame-3.97;
    //(默认是装在/usr/local下);
    //--prefix=/usr/此参数一定要(便于调用os的其它系统包)
    //--enable-shared此参数一定要
    ./configure --enable-shared --prefix=/usr/;
    make;
    make install;

b、支持3gp格式，这也是现在好多手机支持的格式，因为手机用户是我们的主要用户，所以也得支持编译

编译的时候加上--enable-amr_nb --enable-amr_wb参数就行，根据编译系统的提示，所以我们得下载一些编译3gp所需得文件。

wget <http://www.3gpp.org/ftp/Specs/archive/26_series/26.204/26204-510.zip>
解压以后把里面的文件都拷贝到libavcodec/amrwb_float/

wget <http://www.3gpp.org/ftp/Specs/archive/26_series/26.104/26104-510.zip>
解压以后把里面的文件都拷贝到libavcodec/amr_float/

c、mpg4 aac格式支持，由于服务器还针对手机用户服务，所以，类似aac，mpg4铃声格式的支持，我们也得做。这里我们安装faad2和faac就行
[下载请到http://www.audiocoding.com/modules/mydownloads/]()

    tar zxvf faad2-2.5.tar.gz
    cd faad2
    echo > plugins/Makefile.am
    echo > plugins/xmms/src/Makefile.am
    sed -i '/E_B/d' configure.in
    autoreconf -vif
    ./configure --prefix=/usr
    make &&
    make install

    tar zxvf faac-1.25.tar.gz
    cd faac
    sed -i '/[2de].M/d' configure.in
    echo "AC_OUTPUT(common/Makefile common/mp4v2/Makefile libfaac/Makefile frontend/Makefile include/Makefile Makefile)" >> configure.in
    autoreconf -vif
    ./configure --prefix=/usr
    make &&
    make install

d、支持xvid; x264，现在最流行的两种高质量的压缩格式
xvid的编译安装
    wget <http://downloads.xvid.org/downloads/xvidcore-1.1.2.tar.gz>
    tar zvxf xvidcore-1.1.2.tar.gz

    cd xvidcore-1.1.2/build/generic
    ./configure --prefix=/usr --enable-shared
    make
    make install

x264的获取同样是采用svn方式:

    svn co <svn://svn.videolan.org/x264/trunk> x264

[linux下须从http://www.kernel.org/pub/software/devel/nasm/binaries/linux/下载nasm-0.98.39-1.i386.rpm]()

在linux下安装就行了。。。

    rpm -ivh nasm-0.98.39-1.i386.rpm（如-ivh不行就用-Uvh）

    cd x264
    ./configure --prefix=/usr --enable-shared
    make
    make install

e、安装ffmpeg:

    //as4.3系统已经支持ac3编码，只要加--enable-a52 --enable-gpl参数就行
    //我加--enable-shared参数没有成功
    ./configure --prefix=/opt/ffmpeg/ --enable-mp3lame --enable-amr_nb --enable-amr_wb --enable-a52 --enable-xvid --enable-x264 --enable-faad --enable-faac --enable-gpl --enable-pthreads;

    make clean;//一定要；否则有可能没声音。
    make;
    make install;

在相应windows和linux目录下（有ffmpeg文件的;以下用linux下说明）：
3、使用ffmpeg转换视频为flv文件：

    ./ffmpeg -i "/opt/input/1.mpg" -y -ab 32 -ar 22050 -b 800000 -s 640*480 /opt/output/1.flv"

ffmpeg能解析的格式：（asx，asf，mpg，wmv，3gp，mp4，mov，avi，flv等）

对ffmpeg无法解析的文件格式(wmv9，rm，rmvb等),

可以先用别的工具（mencoder）转换为avi(ffmpeg能解析的)格式.

./mencoder /input/a.rmvb -oac lavc -lavcopts acodec=mp3:abitrate=64 -ovc xvid -xvidencopts bitrate=600 -of avi -o /output/a.avi

在执行./ffmpeg -i "/opt/input/a.avi" -y -ab 32 -ar 22050 -b 800000 -s 640*480 /opt/output/a.flv"就可以转了。

4、视频抓图:

    ./ffmpeg -i "/opt/input/a.flv" -y -f image2 -t 1 -s 300*200 "/opt/output/1.jpg" //获取静态图

    ./ffmpeg -i "/opt/input/a.mpg" -vframes 30 -y -f gif "/output/1.gif" //获取动态图;

不提倡抓gif文件；因为抓出的gif文件大而播放不流畅。

用mencoder在线转换视频格式并控制视频品质
<http://www.bestlovesky.com/index.php/632.html>

经验小结：

1，以上两个软件可以在 Windwos 与 Linux 平台上使用。不过，做为服务端的应用，最好是采用 Linux 系统。

2，在实际的使用当中，ffmpeg 在 Redhat与Suse Linux 下面都运行正常，但是我在Suse Linux 10 上面压缩rm与 rmvb 视频时，压出来的片子声音不正常。在其他参数都完全一样的情况下使用Redhat 压缩，居然是正常的。看来 Suse 对企业应用比较好，但是对多媒体的应用（偏向个人方面）还是 Redhat 比较好一些。

3，ffmpeg 也可以压缩视频，不过，效果可是比mencoder 差好些。所以，基本上我只使用 ffmpeg 来抓取视频中的图片。可以使用 ffmpeg-php 这个开源项目程序来抓取任何一帧的图片，这样，我们就可以很方便地大致了解这个视频的内容了。

4，ffmpeg 压缩一个 wmv 文件，可能使用不到一分钟，但是 mencoder 却压缩了好几分钟，由于 mencoder 需要计算更多东西，所以，需要花更多的时间。

5，mencoder 支持的视频格式非常多，如常见的 wmv,avi,mpg,rm,rmvb,mov,3gp,mp4 等，大约有上百种，不过，我还无法一一测试，估计也是没有问题的。而ffmpeg 不支持 rm与rmvb 格式。

6，做为视频压缩，对机器的要求是比较高的，对系统资源的占用比较大，主要是对CPU与磁盘IO要求高。前两天压缩一个视频，使用 宝德 PR2700D 用了4分钟，使用宝德 PR4800 用了12分钟，使用一台 Dell 2950 上面的 Vmware 虚拟机使用了大约8分钟。综合来看，最好是CPU强一些，内存不要小于2GB，磁盘的IO要快一些。
