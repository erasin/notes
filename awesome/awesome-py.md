作者：严林
链接：https://www.zhihu.com/question/24590883/answer/89226375
来源：知乎
著作权归作者所有，转载请联系作者获得授权。

**网络**
[Scapy**](//link.zhihu.com/?target=http%3A//secdev.org/projects/scapy), [Scapy3k**](//link.zhihu.com/?target=http%3A//github.com/phaethon/scapy): 发送，嗅探，分析和伪造网络数据包。可用作交互式包处理程序或单独作为一个库。
[pypcap**](//link.zhihu.com/?target=http%3A//code.google.com/p/pypcap/), [Pcapy**](//link.zhihu.com/?target=http%3A//oss.coresecurity.com/projects/pcapy.html), [pylibpcap**](//link.zhihu.com/?target=http%3A//pylibpcap.sourceforge.net/): 几个不同 libpcap 捆绑的python库
[libdnet**](//link.zhihu.com/?target=http%3A//code.google.com/p/libdnet/): 低级网络路由，包括端口查看和以太网帧的转发
[dpkt**](//link.zhihu.com/?target=https%3A//github.com/kbandla/dpkt): 快速，轻量数据包创建和分析，面向基本的 TCP/IP 协议
[Impacket**](//link.zhihu.com/?target=http%3A//oss.coresecurity.com/projects/impacket.html): 伪造和解码网络数据包，支持高级协议如 NMB 和 SMB
[pynids**](//link.zhihu.com/?target=http%3A//jon.oberheide.org/pynids/): libnids 封装提供网络嗅探，IP 包碎片重组，TCP 流重组和端口扫描侦查 
[Dirtbags py-pcap**](//link.zhihu.com/?target=http%3A//dirtbags.net/py-pcap.html): 无需 libpcap 库支持读取 pcap 文件
[flowgrep**](//link.zhihu.com/?target=http%3A//monkey.org/%7Ejose/software/flowgrep/): 通过正则表达式查找数据包中的 Payloads
[Knock Subdomain Scan**](//link.zhihu.com/?target=https%3A//github.com/guelfoweb/knock): 通过字典枚举目标子域名
[SubBrute**](//link.zhihu.com/?target=https%3A//github.com/TheRook/subbrute): 快速的子域名枚举工具
[Mallory**](//link.zhihu.com/?target=https%3A//bitbucket.org/IntrepidusGroup/mallory): 可扩展的 TCP/UDP 中间人代理工具，可以实时修改非标准协议
[Pytbull**](//link.zhihu.com/?target=http%3A//pytbull.sourceforge.net/): 灵活的 IDS/IPS 测试框架（附带超过300个测试样例）
**调试和逆向工程**
[Paimei**](//link.zhihu.com/?target=https%3A//github.com/OpenRCE/paimei): 逆向工程框架，包含 [PyDBG**](//link.zhihu.com/?target=https%3A//github.com/OpenRCE/pydbg), PIDA,pGRAPH
[Immunity Debugger**](//link.zhihu.com/?target=http%3A//debugger.immunityinc.com/): 脚本 GUI 和命令行调试器 
[mona.py**](//link.zhihu.com/?target=https%3A//www.corelan.be/index.php/2011/07/14/mona-py-the-manual/): Immunity Debugger 中的扩展，用于代替 pvefindaddr
[IDAPython**](//link.zhihu.com/?target=https%3A//github.com/idapython/src): IDA pro 中的插件，集成 Python 编程语言，允许脚本在 IDA Pro 中执行
[PyEMU**](//link.zhihu.com/?target=https%3A//github.com/codypierce/pyemu): 全脚本实现的英特尔32位仿真器，用于恶意软件分析 
[pefile**](//link.zhihu.com/?target=https%3A//github.com/erocarrera/pefile): 读取并处理 PE 文件
[pydasm**](//link.zhihu.com/?target=https%3A//github.com/axcheron/pydasm): Python 封装的 [libdasm**](//link.zhihu.com/?target=https%3A//github.com/alexeevdv/libdasm)
[PyDbgEng**](//link.zhihu.com/?target=http%3A//pydbgeng.sourceforge.net/): Python 封装的微软 Windows 调试引擎
[uhooker**](//link.zhihu.com/?target=http%3A//oss.coresecurity.com/projects/uhooker.htm): 截获 DLL 或内存中任意地址可执行文件的 API 调用
[diStorm**](//link.zhihu.com/?target=http%3A//www.ragestorm.net/distorm/): AMD64 下的反汇编库
[python-ptrace**](//link.zhihu.com/?target=http%3A//python-ptrace.readthedocs.org/): Python 写的使用 ptrace 的调试器
[vdb/vtrace**](//link.zhihu.com/?target=https%3A//github.com/joonty/vdebug): vtrace 是用 Python 实现的跨平台调试 API, vdb 是使用它的调试器
[Androguard**](//link.zhihu.com/?target=https%3A//github.com/androguard/androguard): 安卓应用程序的逆向分析工具
[Capstone**](//link.zhihu.com/?target=http%3A//www.capstone-engine.org/): 一个轻量级的多平台多架构支持的反汇编框架。支持包括ARM,ARM64,MIPS和x86/x64平台。
[PyBFD**](//link.zhihu.com/?target=https%3A//github.com/Groundworkstech/pybfd/): GNU 二进制文件描述(BFD)库的 Python 接口
**Fuzzing**
[Sulley**](//link.zhihu.com/?target=https%3A//github.com/OpenRCE/sulley): 一个模糊器开发和模糊测试的框架，由多个可扩展的构件组成的
[Peach Fuzzing Platform**](//link.zhihu.com/?target=http%3A//peachfuzz.sourceforge.net/): 可扩展的模糊测试框架(v2版本 是用 Python 语言编写的)
[antiparser**](//link.zhihu.com/?target=http%3A//antiparser.sourceforge.net/): 模糊测试和故障注入的 API
[TAOF**](//link.zhihu.com/?target=http%3A//sourceforge.net/projects/taof/): (The Art of Fuzzing, 模糊的艺术)包含 ProxyFuzz, 一个中间人网络模糊测试工具
[untidy**](//link.zhihu.com/?target=http%3A//untidy.sourceforge.net/): 针对 XML 模糊测试工具
[Powerfuzzer**](//link.zhihu.com/?target=http%3A//www.powerfuzzer.com/): 高度自动化和可完全定制的 Web 模糊测试工具
[SMUDGE**](//link.zhihu.com/?target=http%3A//www.fuzzing.org/wp-content/SMUDGE.zip): 纯 Python 实现的网络协议模糊测试
[Mistress**](//link.zhihu.com/?target=http%3A//www.packetstormsecurity.org/fuzzer/mistress.rar): 基于预设模式，侦测实时文件格式和侦测畸形数据中的协议
[Fuzzbox**](//link.zhihu.com/?target=https%3A//isecpartners.com/tools/application-security/fuzzbox.aspx): 媒体多编码器的模糊测试
[Forensic Fuzzing Tools**](//link.zhihu.com/?target=https%3A//isecpartners.com/tools/application-security/forensic-fuzzing-tools.aspx): 通过生成模糊测试用的文件，文件系统和包含模糊测试文件的文件系统，来测试取证工具的鲁棒性
[Windows IPC Fuzzing Tools**](//link.zhihu.com/?target=https%3A//isecpartners.com/tools/application-security/windows-ipc-fuzzing-tools.aspx): 使用 Windows 进程间通信机制进行模糊测试的工具
[WSBang**](//link.zhihu.com/?target=https%3A//www.isecpartners.com/tools/application-security/wsbang.aspx): 基于 Web 服务自动化测试 SOAP 安全性
[Construct**](//link.zhihu.com/?target=http%3A//construct.readthedocs.org/): 用于解析和构建数据格式(二进制或文本)的库
[fuzzer.py(feliam)**](//link.zhihu.com/?target=http%3A//sites.google.com/site/felipeandresmanzano/fuzzer.py%3Fattredirects%3D0): 由 Felipe Andres Manzano 编写的简单模糊测试工具
[Fusil**](//link.zhihu.com/?target=http%3A//fusil.readthedocs.org/): 用于编写模糊测试程序的 Python 库
**Web**
[Requests**](//link.zhihu.com/?target=http%3A//python-requests.org/): 优雅，简单，人性化的 HTTP 库
[HTTPie**](//link.zhihu.com/?target=http%3A//httpie.org/): 人性化的类似 cURL 命令行的 HTTP 客户端
[ProxMon**](//link.zhihu.com/?target=https%3A//www.isecpartners.com/tools/application-security/proxmon.aspx): 处理代理日志和报告发现的问题
[WSMap**](//link.zhihu.com/?target=https%3A//www.isecpartners.com/tools/application-security/wsmap.aspx): 寻找 Web 服务器和发现文件
[Twill**](//link.zhihu.com/?target=http%3A//twill.idyll.org/): 从命令行界面浏览网页。支持自动化网络测试
[Ghost.py**](//link.zhihu.com/?target=http%3A//jeanphix.me/Ghost.py/): Python 写的 WebKit Web 客户端
[Windmill**](//link.zhihu.com/?target=http%3A//www.getwindmill.com/): Web 测试工具帮助你轻松实现自动化调试 Web 应用
[FunkLoad**](//link.zhihu.com/?target=http%3A//funkload.nuxeo.org/): Web 功能和负载测试
[spynner**](//link.zhihu.com/?target=https%3A//github.com/makinacorpus/spynner): Python 写的 Web浏览模块支持 Javascript/AJAX
[python-spidermonkey**](//link.zhihu.com/?target=https%3A//github.com/davisp/python-spidermonkey): 是 Mozilla JS 引擎在 Python 上的移植，允许调用 Javascript 脚本和函数
[mitmproxy**](//link.zhihu.com/?target=http%3A//mitmproxy.org/): 支持 SSL 的 HTTP 代理。可以在控制台接口实时检查和编辑网络流量
[pathod/pathoc**](//link.zhihu.com/?target=http%3A//pathod.net/): 变态的 HTTP/S 守护进程，用于测试和折磨 HTTP 客户端
**取证**
[Volatility**](//link.zhihu.com/?target=http%3A//www.volatilityfoundation.org/): 从 RAM 中提取数据
[Rekall**](//link.zhihu.com/?target=http%3A//www.rekall-forensic.com/): Google 开发的内存分析框架
[LibForensics**](//link.zhihu.com/?target=http%3A//code.google.com/p/libforensics/): 数字取证应用程序库
[TrIDLib**](//link.zhihu.com/?target=http%3A//mark0.net/code-tridlib-e.html): Python 实现的从二进制签名中识别文件类型
[aft**](//link.zhihu.com/?target=https%3A//github.com/agnivesh/aft): 安卓取证工具集恶意软件分析
[pyew**](//link.zhihu.com/?target=https%3A//github.com/joxeankoret/pyew): 命令行十六进制编辑器和反汇编工具，主要用于分析恶意软件
[Exefilter**](//link.zhihu.com/?target=http%3A//www.decalage.info/exefilter): 过滤 E-mail，网页和文件中的特定文件格式。可以检测很多常见文件格式，也可以移除文档内容。
[pyClamAV**](//link.zhihu.com/?target=http%3A//xael.org/norman/python/pyclamav/index.html): 增加你 Python 软件的病毒检测能力
[jsunpack-n**](//link.zhihu.com/?target=https%3A//github.com/urule99/jsunpack-n): 通用 JavaScript 解释器，通过模仿浏览器功能来检测针对目标浏览器和浏览器插件的漏洞利用
[yara-python**](//link.zhihu.com/?target=https%3A//github.com/plusvic/yara/tree/master/yara-python): 对恶意软件样本进行识别和分类
[phoneyc**](//link.zhihu.com/?target=https%3A//github.com/honeynet/phoneyc): 纯 Python 实现的蜜罐
[CapTipper**](//link.zhihu.com/?target=https%3A//github.com/omriher/CapTipper): 分析，研究和重放 PCAP 文件中的 HTTP 恶意流量
**PDF**
[peepdf**](//link.zhihu.com/?target=http%3A//eternal-todo.com/tools/peepdf-pdf-analysis-tool): Python 编写的PDF文件分析工具，可以帮助检测恶意的PDF文件
[Didier Stevens’ PDF tools**](//link.zhihu.com/?target=http%3A//blog.didierstevens.com/programs/pdf-tools): 分析，识别和创建 PDF 文件(包含[PDFiD**](//link.zhihu.com/?target=http%3A//blog.didierstevens.com/programs/pdf-tools/%23pdfid)，[pdf-parser**](//link.zhihu.com/?target=http%3A//blog.didierstevens.com/programs/pdf-tools/%23pdf-parser)，[make-pdf**](//link.zhihu.com/?target=http%3A//blog.didierstevens.com/programs/pdf-tools/%23make-pdf) 和 mPDF)
[Opaf**](//link.zhihu.com/?target=http%3A//code.google.com/p/opaf/): 开放 PDF 分析框架，可以将 PDF 转化为 XML 树从而进行分析和修改。
[Origapy**](//link.zhihu.com/?target=http%3A//www.decalage.info/python/origapy): Ruby 工具 [Origami**](//link.zhihu.com/?target=http%3A//www.security-labs.org/origami/) 的 Python 接口，用于审查 PDF 文件 
[pyPDF2**](//link.zhihu.com/?target=http%3A//mstamy2.github.io/PyPDF2/): Python PDF 工具包包含：信息提取，拆分，合并，制作，加密和解密等等
[PDFMiner**](//link.zhihu.com/?target=http%3A//www.unixuser.org/%7Eeuske/python/pdfminer/index.html): 从 PDF 文件中提取文本
[python-poppler-qt4**](//link.zhihu.com/?target=https%3A//github.com/wbsoft/python-poppler-qt4): Python 写的 [Poppler**](//link.zhihu.com/?target=http%3A//poppler.freedesktop.org/) PDF 库，支持 Qt4
杂项
[InlineEgg**](//link.zhihu.com/?target=http%3A//oss.coresecurity.com/projects/inlineegg.html): 使用 Python 编写的具有一系列小功能的工具箱
[Exomind**](//link.zhihu.com/?target=http%3A//corelabs.coresecurity.com/index.php%3Fmodule%3DWiki%26action%3Dview%26type%3Dtool%26name%3DExomind): 是一个利用社交网络进行钓鱼攻击的工具
[RevHosts**](//link.zhihu.com/?target=http%3A//www.securityfocus.com/tools/3851): 枚举指定 IP 地址包含的虚拟主句
[simplejson**](//link.zhihu.com/?target=https%3A//github.com/simplejson/simplejson/): JSON 编码和解码器，例如使用 [Google’s AJAX API**](//link.zhihu.com/?target=http%3A//dcortesi.com/2008/05/28/google-ajax-search-api-example-python-code/)
[PyMangle**](//link.zhihu.com/?target=http%3A//code.google.com/p/pymangle/): 命令行工具和一个创建用于渗透测试使用字典的库
[Hachoir**](//link.zhihu.com/?target=https%3A//bitbucket.org/haypo/hachoir/wiki/Home): 查看和编辑二进制流
**其他有用的库和工具**
[IPython**](//link.zhihu.com/?target=http%3A//ipython.scipy.org/): 增强的交互式 Python shell
[Beautiful Soup**](//link.zhihu.com/?target=http%3A//www.crummy.com/software/BeautifulSoup/): HTML 解析器
[matplotlib**](//link.zhihu.com/?target=http%3A//matplotlib.sourceforge.net/): 制作二维图
[Mayavi**](//link.zhihu.com/?target=http%3A//code.enthought.com/projects/mayavi/): 三维科学数据的可视化与绘图
[RTGraph3D**](//link.zhihu.com/?target=http%3A//www.secdev.org/projects/rtgraph3d/): 在三维空间中创建动态图
[Twisted**](//link.zhihu.com/?target=http%3A//twistedmatrix.com/): Python 语言编写的事件驱动的网络框架
[Suds**](//link.zhihu.com/?target=https%3A//fedorahosted.org/suds/): 一个轻量级的基于SOAP的python客户端
[M2Crypto**](//link.zhihu.com/?target=http%3A//chandlerproject.org/bin/view/Projects/MeTooCrypto):  Python 语言对 OpenSSL 的封装
[NetworkX**](//link.zhihu.com/?target=http%3A//networkx.lanl.gov/): 图库(边, 节点)
[Pandas**](//link.zhihu.com/?target=http%3A//pandas.pydata.org/): 基于 Numpy 构建的含有更高级数据结构和工具的数据分析包
[pyparsing**](//link.zhihu.com/?target=http%3A//pyparsing.wikispaces.com/): 通用解析模块
[lxml**](//link.zhihu.com/?target=http%3A//lxml.de/): 使用 Python 编写的库，可以迅速、灵活地处理 XML
[Whoosh**](//link.zhihu.com/?target=https%3A//bitbucket.org/mchaput/whoosh/): 纯python实现的全文搜索组件
[Pexpect**](//link.zhihu.com/?target=https%3A//github.com/pexpect/pexpect): 控制和自动化程序
[Sikuli**](//link.zhihu.com/?target=http%3A//groups.csail.mit.edu/uid/sikuli/): 使用 [Jython**](//link.zhihu.com/?target=http%3A//www.jython.org/) 脚本自动化基于截图进行视觉搜索
[PyQt**](//link.zhihu.com/?target=http%3A//www.riverbankcomputing.co.uk/software/pyqt) 和[PySide**](//link.zhihu.com/?target=http%3A//www.pyside.org/): Python 捆绑的 Qt 应用程序框架和 GUI 库