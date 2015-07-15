# wechat

**欢迎你，开发者**

当你打开此页面，我们期待你成为企业号的开发者。

企业号是微信为企业客户提供的移动应用入口。它帮助企业建立员工、上下游供应链与企业IT系统间的连接。利用企业号，企业或第三方合作伙伴可以帮助企业快速、低成本的实现高质量的移动轻应用，实现生产、管理、协作、运营的移动化。我们有信心，帮助你及每一位企业号开发者，提升开发效率，降低开发成本与难度，确保企业应用的活跃。

我们也坚信企业号能帮助你获得更多的商业机会，服务更多的企业客户，从而不断的提升你的价值。你的成功也将是企业号的成功。

当你成功申领一个企业号后，你可以登录企业号的管理页面，导入通讯录、配置应用、邀请成员关注该企业号，也可以通过应用向成员发送文本、图文、文件、视频、音频等多媒体消息。通过简单的配置，你就可以自动回复成员发送的消息，实现公告通知、知识管理、企业文化建设、手机企业通讯录等基本的企业应用。

你还可以通过本文档所描述的接口，建立企业号同企业应用间的连接，实现更多丰富且个性化的企业移动应用。

# 建立连接

**连接**将使你的企业号更具价值，你可以使用以下**三种方式**，连接你的企业号及企业应用：

1. 企业应用调用企业号提供的接口，管理或查询企业号后台所管理的资源、或给成员发送消息等，以下称**主动调用模式**。
2. 企业号把用户发送的消息或用户触发的事件推送给企业应用，由企业应用处理，以下称**回调模式**。
3. 用户在微信中阅读企业应用下发的H5页面，该页面可以调用微信提供的原生接口，使用微信开放的终端能力，以下称**JSAPI模式**。

通过这三种连接方式的结合，你可以在企业号中建立功能强大的移动轻应用，并依托微信数亿活跃用户，帮助企业方便、快捷地实现应用的部署，并确保应用的活跃度。

## **主动调用**
主动调用是最基本的连接模式，当你的应用调用企业号时，需**使用https协议、Json数据格式、UTF8编码，访问域名为[https://qyapi.weixin.qq.com](//https://qyapi.weixin.qq.com)，数据包不需要加密**。

在每次主动调用企业号接口时需要带上<a href="/wiki/index.php?title=AccessToken" title="AccessToken">AccessToken</a>参数。AccessToken参数由<a href="/wiki/index.php?title=CorpID" title="CorpID">CorpID</a>和<a href="/wiki/index.php?title=Secret" title="Secret">Secret</a>换取。

[CorpID](/wiki/index.php?title=CorpID" title="CorpID)是企业号的标识，每个企业号拥有一个唯一的CorpID；Secret是管理组凭证密钥。

系统管理员可通过管理端的权限管理功能创建管理组，分配管理组对应用、通讯录、接口的访问权限。完成后，管理组即可获得唯一的secret。系统管理员可通过权限管理查看所有管理组的secret，其他管理员可通过设置中的开发者凭据查看。

当企业应用调用企业号接口时，企业号后台为根据此次访问的AccessToken,校验访问的合法性以及所对应的管理组的管理权限以返回相应的结果。

<font color="red">**注：**</font>你应该审慎配置管理组的权限，够用即好，权限过大会增加误操作可能性及信息安全隐患。
### **获取AccessToken**
AccessToken是企业号的全局唯一票据，调用接口时需携带AccessToken。

AccessToken需要用<a href="/wiki/index.php?title=CorpID" title="CorpID">CorpID</a>和<a href="/wiki/index.php?title=Secret" title="Secret">Secret</a>来换取，不同的Secret会返回不同的AccessToken。**正常情况下AccessToken有效期为7200秒，有效期内重复获取返回相同结果，并自动续期**。
<ul>
<li>请求说明</li>
</ul>
Https请求方式: GET

<a rel="nofollow" class="external free" href="https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=id&amp;corpsecret=secrect">https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=id&amp;corpsecret=secrect</a>
<ul>
<li>参数说明</li>
</ul>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th><th>必须</th><th>说明</th></tr>
<tr><td> <a href="/wiki/index.php?title=CorpID" title="CorpID">corpid</a></td><td> 是</td><td> 企业Id</td></tr>
<tr><td> <a href="/wiki/index.php?title=Secret" title="Secret">corpsecret</a></td><td> 是</td><td> 管理组的凭证密钥</td></tr></tbody></table>
<ul>
<li>权限说明</li>
</ul>
每个secret代表了对应用、通讯录、接口的不同权限；不同的管理组拥有不同的secret。
<ul>
<li>返回说明</li>
</ul>
a)正确的Json返回结果:
<pre>{
   "access_token": "accesstoken000001",
}
</pre>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th>
<th style="width:400px">说明</th></tr>
<tr><td> access_token</td><td> 获取到的凭证</td></tr></tbody></table>
b)错误的Json返回示例:
<pre>{
   "errcode": 43003,
   "errmsg": "require https"
}
</pre>

### **主动调用的频率限制**
当你获取到<a href="/wiki/index.php?title=AccessToken" title="AccessToken">AccessToken</a>时，你的应用就可以成功调用企业号后台所提供的各种接口以管理或访问企业号后台的资源或给企业号成员发消息。

为了防止企业应用的程序错误而引发企业号服务器负载异常，默认情况下，每个企业号调用接口都有一定的频率限制，当超过此限制时，调用对应接口会收到相应错误码。

以下是当前默认的频率限制，企业号后台可能会根据运营情况调整此阈值：
<ul>
<li>基础频率</li>
</ul>
每企业调用单个cgi/api不可超过1000次/分，30000次/小时

每ip调用单个cgi/api不可超过2000次/分，60000次/小时
<ul>
<li>发消息频率</li>
</ul>
每企业不可超过帐号上限数*30人次/天
<ul>
<li>创建帐号频率</li>
</ul>
每企业创建帐号数不可超过帐号上限数*3/月

<!--
NewPP limit report
CPU time usage: 0.020 seconds
Real time usage: 0.026 seconds
Preprocessor visited node count: 13/1000000
Preprocessor generated node count: 22/1000000
Post‐expand include size: 0/2097152 bytes
Template argument size: 0/2097152 bytes
Highest expansion depth: 2/40
Expensive parser function count: 0/100
-->
<!-- Saved in parser cache with key db_wiki:pcache:idhash:4-0!*!0!!*!*!* and timestamp 20150115091611 and revision id 646
 -->

<div id="toc" class="toc"><div id="toctitle">## 目录</div>
<ul>
<li class="toclevel-1 tocsection-1"><a href="#.E5.9B.9E.E8.B0.83.E6.A8.A1.E5.BC.8F"><span class="tocnumber">1</span> <span class="toctext"><b>回调模式</a>
<ul>
<li class="toclevel-2 tocsection-2"><a href="#.E5.BC.80.E5.90.AF.E5.BA.94.E7.94.A8.E7.9A.84.E5.9B.9E.E8.B0.83.E6.A8.A1.E5.BC.8F"><span class="tocnumber">1.1</span> <span class="toctext"><b>开启应用的回调模式</a></li>
<li class="toclevel-2 tocsection-3"><a href="#.E4.BD.BF.E7.94.A8.E5.9B.9E.E8.B0.83.E6.A8.A1.E5.BC.8F"><span class="tocnumber">1.2</span> <span class="toctext"><b>使用回调模式</a></li>
<li class="toclevel-2 tocsection-4"><a href="#.E6.8E.A5.E6.94.B6.E6.B6.88.E6.81.AF.E6.97.B6.E7.9A.84.E5.8A.A0.E8.A7.A3.E5.AF.86.E5.A4.84.E7.90.86"><span class="tocnumber">1.3</span> <span class="toctext"><b>接收消息时的加解密处理</a></li>
<li class="toclevel-2 tocsection-5"><a href="#.E8.8E.B7.E5.8F.96.E5.BE.AE.E4.BF.A1.E6.9C.8D.E5.8A.A1.E5.99.A8.E7.9A.84ip.E6.AE.B5"><span class="tocnumber">1.4</span> <span class="toctext"><b>获取微信服务器的ip段</a></li>
</ul></li>
</ul>
</div>
## **回调模式**
在回调模式下，企业不仅可以主动调用企业号接口，还可以接收用户的消息或事件。**接收的信息使用XML数据格式、UTF8编码，并以AES方式加密**。

企业号的每个应用都有自己的回调模式开关。在管理端开启并设置好相关参数后，此应用的回调模式才生效。

针对加解密的处理，微信提供了各种语言的库，企业可以在附录中下载。

### **开启应用的回调模式**
当你开启应用的回调模式时，企业号会要求你填写应用的URL、Token、EncodingAESKey三个参数。

URL是企业应用接收企业号推送请求的访问协议和地址，支持http或https协议。

Token可由企业任意填写，用于生成签名。

EncodingAESKey用于消息体的加密，是AES密钥的Base64编码。

验证URL、Token以及加密的详细处理请参考后续'接收消息时的加解密处理'的章节。

<div class="center"><div class="floatnone"><a href="/wiki/index.php?title=%E6%96%87%E4%BB%B6:Hdms.png" class="image"><img alt="Hdms.png" src="/wiki/images/4/42/Hdms.png" width="600" height="427"></a></div></div>
**验证URL有效性**

当你提交以上信息时，企业号将发送GET请求到填写的URL上，GET请求携带四个参数，**企业在获取时需要做urldecode处理**，否则会验证不成功。
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:120px">参数</th><th>描述</th><th>是否必带</th></tr>
<tr><td> msg_signature</td><td> 微信加密签名，msg_signature结合了企业填写的token、请求中的timestamp、nonce参数、加密的消息体</td><td> 是</td></tr>
<tr><td> timestamp</td><td> 时间戳</td><td> 是</td></tr>
<tr><td> nonce</td><td> 随机数</td><td> 是</td></tr>
<tr><td> echostr</td><td> 加密的随机字符串，以msg_encrypt格式提供。需要解密并返回echostr明文，解密后有random、msg_len、msg、$CorpID四个字段，其中msg即为echostr明文</td><td> 首次校验时必带</td></tr></tbody></table>
企业通过参数msg_signature对请求进行校验，如果确认此次GET请求来自企业号，那么**企业应用对echostr参数解密并原样返回echostr明文(不能加引号)**，则接入验证生效，回调模式才能开启。

后续回调企业时都会在请求URL中带上以上参数（echostr除外），校验方式与首次验证URL一致。

### **使用回调模式**
**企业号在回调企业URL时，会对消息体本身做AES加密，以XML格式POST到企业应用的URL上；企业在被动响应时，也需要对数据加密，以XML格式返回给微信。企业的回复支持文本、图片、语音、视频、图文等格式**。

微信服务器在五秒内收不到响应会断掉连接，并且重新发起请求，总共重试三次。如果在调试中，发现员工无法收到响应的消息，可以检查是否消息处理超时。

当接收成功后，http头部返回200表示接收ok，其他错误码一律当做失败并发起重试

关于重试的消息排重，有msgid的消息推荐使用msgid排重。事件类型消息推荐使用FromUserName + CreateTime排重。

假如企业无法保证在五秒内处理并回复，可以直接回复空串，企业号不会对此作任何处理，并且不会发起重试。这种情况下，可以使用发消息接口进行异步回复。

假设企业回调URL为<a rel="nofollow" class="external text" href="//http://api.3dept.com">http://api.3dept.com</a>。
<ul>
<li>请求说明：</li>
</ul>
<a rel="nofollow" class="external free" href="http://api.3dept.com/?msg_signature=ASDFQWEXZCVAQFASDFASDFSS&amp;timestamp=13500001234&amp;nonce=123412323">http://api.3dept.com/?msg_signature=ASDFQWEXZCVAQFASDFASDFSS&amp;timestamp=13500001234&amp;nonce=123412323</a>
<ul>
<li>回调数据格式：</li>
</ul>
<pre>&lt;xml&gt;
   &lt;ToUserName&gt;&lt;![CDATA[toUser]]&lt;/ToUserName&gt;
   &lt;AgentID&gt;&lt;![CDATA[toAgentID]]&lt;/AgentID&gt;
   &lt;Encrypt&gt;&lt;![CDATA[msg_encrypt]]&lt;/Encrypt&gt;
&lt;/xml&gt;
</pre>

<pre>1.msg_encrypt为经过加密的密文
2.AgentID为接收的应用id，可在应用的设置页面获取
3.ToUserName为企业号的CorpID
</pre>
企业需要对msg_signature进行校验，并解密msg_encrypt，得出msg的原文。
<ul>
<li>被动响应给微信的数据格式：</li>
</ul>
<pre>&lt;xml&gt;
   &lt;Encrypt&gt;&lt;![CDATA[msg_encrypt]]&gt;&lt;/Encrypt&gt;
   &lt;MsgSignature&gt;&lt;![CDATA[msg_signature]]&gt;&lt;/MsgSignature&gt;
   &lt;TimeStamp&gt;timestamp&lt;/TimeStamp&gt;
   &lt;Nonce&gt;&lt;![CDATA[nonce]]&gt;&lt;/Nonce&gt;
&lt;/xml&gt;
</pre>

<pre>1.msg_encrypt为经过加密的密文，算法参见附录
2.MsgSignature为签名，算法参见附录
3.TimeStamp为时间戳，Nonce为随机数，由企业自行生成
</pre>

### **接收消息时的加解密处理**
企业可以直接使用微信提供的库进行加解密的处理，目前提供的有c++/python/php/java/c#等语言版本。代码提供了解密、加密、验证URL三个接口，企业可根据自身需要下载(参见附录)。以下为库函数的使用说明(以c++为例)，更详细的加解密方案请参考附录。

**1、解密函数**
<pre>int DecryptMsg(const string &amp;sMsgSignature, const string &amp;sTimeStamp, const string &amp;sNonce, const string &amp;sPostData, string &amp;sMsg);
</pre>
<ul>
<li>参数说明</li>
</ul>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:120px">参数</th><th>必须</th><th>说明</th></tr>
<tr><td> sMsgSignature</td><td> 是</td><td> 从回调URL中获取的msg_signature参数</td></tr>
<tr><td> sTimeStamp</td><td> 是</td><td> 从回调URL中获取的timestamp参数</td></tr>
<tr><td> sNonce</td><td> 是</td><td> 从回调URL中获取的nonce参数</td></tr>
<tr><td> sPostData</td><td> 是</td><td> 从回调URL中获取的整个post数据</td></tr>
<tr><td> sMsg</td><td> 是</td><td> 用于返回解密后的msg，以xml组织</td></tr></tbody></table>
<ul>
<li>返回说明</li>
</ul>
请参阅附录加解密部分。

**2、加密函数**
<pre>int EncryptMsg(const string &amp;sReplyMsg, const string &amp;sTimeStamp, const string &amp;sNonce, string &amp;sEncryptMsg);
</pre>
<ul>
<li>参数说明</li>
</ul>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:120px">参数</th><th>必须</th><th>说明</th></tr>
<tr><td> sReplyMsg</td><td> 是</td><td> 返回的消息体原文</td></tr>
<tr><td> sTimeStamp</td><td> 是</td><td> 时间戳，调用方生成</td></tr>
<tr><td> sNonce</td><td> 是</td><td> 随机数，调用方生成</td></tr>
<tr><td> sEncryptMsg</td><td> 是</td><td> 用于返回的密文，以xml组织</td></tr></tbody></table>
<ul>
<li>返回说明</li>
</ul>
请参阅附录加解密部分。

**3、验证URL函数**
<pre>int VerifyURL(const string &amp;sMsgSignature, const string &amp;sTimeStamp, const string &amp;sNonce, const string &amp;sEchoStr, string &amp;sReplyEchoStr);
</pre>
<ul>
<li>参数说明</li>
</ul>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:120px">参数</th><th>必须</th><th>说明</th></tr>
<tr><td> sMsgSignature</td><td> 是</td><td> 从回调URL中获取的msg_signature参数</td></tr>
<tr><td> sTimeStamp</td><td> 是</td><td> 从回调URL中获取的timestamp参数</td></tr>
<tr><td> sNonce</td><td> 是</td><td> 从回调URL中获取的nonce参数</td></tr>
<tr><td> sEchoStr</td><td> 是</td><td> 从回调URL中获取的echostr参数。注意，此参数必须是urldecode后的值</td></tr>
<tr><td> sReplyEchoStr</td><td> 是</td><td> 解密后的echostr，用于回包。注意，必须原样返回，不要做加引号或其它处理</td></tr></tbody></table>
<ul>
<li>返回说明</li>
</ul>
请参阅附录加解密部分。

### **获取微信服务器的ip段**
<ul>
<li>请求说明</li>
</ul>
Https请求方式: GET

<a rel="nofollow" class="external free" href="https://qyapi.weixin.qq.com/cgi-bin/getcallbackip?access_token=ACCESS_TOKEN">https://qyapi.weixin.qq.com/cgi-bin/getcallbackip?access_token=ACCESS_TOKEN</a>
<ul>
<li>参数说明</li>
</ul>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:150px">参数</th>
<th style="width:120px">必须</th><th>说明</th></tr>
<tr><td> access_token</td><td> 是</td><td> 调用接口凭证</td></tr>
</tbody></table>
<ul>
<li>返回结果</li>
</ul>
<pre>{
   "ip_list": ["101.226.103.*", "101.226.62.*"]
}
</pre>

<!--
NewPP limit report
CPU time usage: 0.060 seconds
Real time usage: 0.063 seconds
Preprocessor visited node count: 22/1000000
Preprocessor generated node count: 32/1000000
Post‐expand include size: 0/2097152 bytes
Template argument size: 0/2097152 bytes
Highest expansion depth: 2/40
Expensive parser function count: 0/100
-->
<!-- Saved in parser cache with key db_wiki:pcache:idhash:6-0!*!*!!zh-cn!2!* and timestamp 20150115112446 and revision id 636
 -->
# 管理通讯录

企业号通讯录具备完全开放的接口，你的应用可以调用这些接口管理部门、成员和标签。

你的应用也可以使用部门、成员、标签发消息，或更改应用的可见范围。

注意，**每个部门的直属员工上限为1000个**；出于安全考虑，**某些接口需要在管理端有明确的授权**。

<!--
NewPP limit report
CPU time usage: 0.004 seconds
Real time usage: 0.003 seconds
Preprocessor visited node count: 1/1000000
Preprocessor generated node count: 4/1000000
Post‐expand include size: 0/2097152 bytes
Template argument size: 0/2097152 bytes
Highest expansion depth: 1/40
Expensive parser function count: 0/100
-->
<!-- Saved in parser cache with key db_wiki:pcache:idhash:30-0!*!*!*!*!*!* and timestamp 20150115130410 and revision id 313
 -->
## 管理部门
<div id="toc" class="toc"><div id="toctitle">## 目录</div>
<ul>
<li class="toclevel-1 tocsection-1"><a href="#.E5.88.9B.E5.BB.BA.E9.83.A8.E9.97.A8"><span class="tocnumber">1</span> <span class="toctext"><b>创建部门</a></li>
<li class="toclevel-1 tocsection-2"><a href="#.E6.9B.B4.E6.96.B0.E9.83.A8.E9.97.A8"><span class="tocnumber">2</span> <span class="toctext"><b>更新部门</a></li>
<li class="toclevel-1 tocsection-3"><a href="#.E5.88.A0.E9.99.A4.E9.83.A8.E9.97.A8"><span class="tocnumber">3</span> <span class="toctext"><b>删除部门</a></li>
<li class="toclevel-1 tocsection-4"><a href="#.E8.8E.B7.E5.8F.96.E9.83.A8.E9.97.A8.E5.88.97.E8.A1.A8"><span class="tocnumber">4</span> <span class="toctext"><b>获取部门列表</a></li>
</ul>
</div>
### **创建部门**
<ul>
<li>请求说明</li>
</ul>
Https请求方式: POST

<a rel="nofollow" class="external free" href="https://qyapi.weixin.qq.com/cgi-bin/department/create?access_token=ACCESS_TOKEN">https://qyapi.weixin.qq.com/cgi-bin/department/create?access_token=ACCESS_TOKEN</a>

请求包结构体为:
<pre>{
   "name": "广州研发中心",
   "parentid": "1",
   "order": "1"
}
</pre>
<ul>
<li>参数说明</li>
</ul>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:120px">参数</th><th>必须</th><th>说明</th></tr>
<tr><td> access_token</td><td> 是</td><td> 调用接口凭证</td></tr>
<tr><td> name</td><td> 是</td><td> 部门名称。长度限制为1~64个字符</td></tr>
<tr><td> parentid</td><td> 是</td><td> 父亲部门id。根部门id为1</td></tr>
<tr><td> order</td><td> 否</td><td> 在父部门中的次序。从1开始，数字越大排序越靠后</td></tr>
</tbody></table>
<ul>
<li>权限说明</li>
</ul>
管理员须拥有“操作通讯录”的接口权限，以及父部门的管理权限。
<ul>
<li>返回结果</li>
</ul>
<pre>{
   "errcode": 0,
   "errmsg": "created",
   "id": 2
}
</pre>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th><th>说明</th></tr>
<tr><td> errcode</td><td> 返回码</td></tr>
<tr><td> errmsg</td><td> 对返回码的文本描述内容</td></tr>
<tr><td> id</td><td> 创建的部门id</td></tr>
</tbody></table>

### **更新部门**
<ul>
<li>请求说明</li>
</ul>
Https请求方式: POST

<a rel="nofollow" class="external free" href="https://qyapi.weixin.qq.com/cgi-bin/department/update?access_token=ACCESS_TOKEN">https://qyapi.weixin.qq.com/cgi-bin/department/update?access_token=ACCESS_TOKEN</a>

请求包结构体为（如果非必须的字段未指定，则不更新该字段之前的设置值）:
<pre>{
   "id": 2,
   "name": "广州研发中心",
   "parentid": "1",
   "order": "1"
}
</pre>
<ul>
<li>参数说明</li>
</ul>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:150px">参数</th>
<th style="width:120px">必须</th><th>说明</th></tr>
<tr><td> access_token</td><td> 是</td><td> 调用接口凭证</td></tr>
<tr><td> id</td><td> 是</td><td> 部门id</td></tr>
<tr><td> name</td><td> 否</td><td> 更新的部门名称。长度限制为1~64个字符。修改部门名称时指定该参数</td></tr>
<tr><td> parentid</td><td> 否</td><td> 父亲部门id。根部门id为1</td></tr>
<tr><td> order</td><td> 否</td><td> 在父部门中的次序。从1开始，数字越大排序越靠后</td></tr>
</tbody></table>
<ul>
<li>权限说明</li>
</ul>
管理员须拥有“操作通讯录”的接口权限，以及该部门的管理权限。
<ul>
<li>返回结果</li>
</ul>
<pre>{
   "errcode": 0,
   "errmsg": "updated"
}
</pre>

### **删除部门**
<ul>
<li>请求说明</li>
</ul>
Https请求方式: GET

<a rel="nofollow" class="external free" href="https://qyapi.weixin.qq.com/cgi-bin/department/delete?access_token=ACCESS_TOKEN&amp;id=2">https://qyapi.weixin.qq.com/cgi-bin/department/delete?access_token=ACCESS_TOKEN&amp;id=2</a>
<ul>
<li>参数说明</li>
</ul>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:150px">参数</th>
<th style="width:120px">必须</th><th>说明</th></tr>
<tr><td> access_token</td><td> 是</td><td> 调用接口凭证</td></tr>
<tr><td> id</td><td> 是</td><td> 部门id。（注：不能删除根部门；不能删除含有子部门、成员的部门）</td></tr>
</tbody></table>
<ul>
<li>权限说明</li>
</ul>
管理员须拥有“操作通讯录”的接口权限，以及该部门的管理权限。
<ul>
<li>返回结果</li>
</ul>
<pre>{
   "errcode": 0,
   "errmsg": "deleted"
}
</pre>

### **获取部门列表**
<ul>
<li>请求说明</li>
</ul>
Https请求方式: GET

<a rel="nofollow" class="external free" href="https://qyapi.weixin.qq.com/cgi-bin/department/list?access_token=ACCESS_TOKEN">https://qyapi.weixin.qq.com/cgi-bin/department/list?access_token=ACCESS_TOKEN</a>
<ul>
<li>参数说明</li>
</ul>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:150px">参数</th>
<th style="width:120px">必须</th><th>说明</th></tr>
<tr><td> access_token</td><td> 是</td><td> 调用接口凭证</td></tr>
</tbody></table>
<ul>
<li>权限说明</li>
</ul>
管理员须拥有’获取部门列表’的接口权限，以及对部门的查看权限。
<ul>
<li>返回结果</li>
</ul>
<pre>{
   "errcode": 0,
   "errmsg": "ok",
   "department": [
       {
           "id": 2,
           "name": "广州研发中心",
           "parentid": 1
       },
       {
           "id": 3
           "name": "邮箱产品部",
           "parentid": 2
       }
   ]
}
</pre>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:150px">参数</th><th>说明</th></tr>
<tr><td> errcode</td><td> 返回码</td></tr>
<tr><td> errmsg</td><td> 对返回码的文本描述内容</td></tr>
<tr><td> department</td><td> 部门列表数据。以部门的order字段从小到大排列</td></tr>
<tr><td> id</td><td> 部门id</td></tr>
<tr><td> name</td><td> 部门名称</td></tr>
<tr><td> parentid</td><td> 父亲部门id。根部门为1</td></tr>
</tbody></table>

<!--
NewPP limit report
CPU time usage: 0.016 seconds
Real time usage: 0.016 seconds
Preprocessor visited node count: 19/1000000
Preprocessor generated node count: 28/1000000
Post‐expand include size: 0/2097152 bytes
Template argument size: 0/2097152 bytes
Highest expansion depth: 2/40
Expensive parser function count: 0/100
-->
<!-- Saved in parser cache with key db_wiki:pcache:idhash:7-0!*!*!!zh-cn!*!* and timestamp 20150116065127 and revision id 534
 -->
## 管理成员
<div id="toc" class="toc"><div id="toctitle">## 目录</div>
<ul>
<li class="toclevel-1 tocsection-1"><a href="#.E5.88.9B.E5.BB.BA.E6.88.90.E5.91.98"><span class="tocnumber">1</span> <span class="toctext"><b>创建成员</a></li>
<li class="toclevel-1 tocsection-2"><a href="#.E6.9B.B4.E6.96.B0.E6.88.90.E5.91.98"><span class="tocnumber">2</span> <span class="toctext"><b>更新成员</a></li>
<li class="toclevel-1 tocsection-3"><a href="#.E5.88.A0.E9.99.A4.E6.88.90.E5.91.98"><span class="tocnumber">3</span> <span class="toctext"><b>删除成员</a></li>
<li class="toclevel-1 tocsection-4"><a href="#.E6.89.B9.E9.87.8F.E5.88.A0.E9.99.A4.E6.88.90.E5.91.98"><span class="tocnumber">4</span> <span class="toctext"><b>批量删除成员</a></li>
<li class="toclevel-1 tocsection-5"><a href="#.E8.8E.B7.E5.8F.96.E6.88.90.E5.91.98"><span class="tocnumber">5</span> <span class="toctext"><b>获取成员</a></li>
<li class="toclevel-1 tocsection-6"><a href="#.E8.8E.B7.E5.8F.96.E9.83.A8.E9.97.A8.E6.88.90.E5.91.98"><span class="tocnumber">6</span> <span class="toctext"><b>获取部门成员</a></li>
<li class="toclevel-1 tocsection-7"><a href="#.E8.8E.B7.E5.8F.96.E9.83.A8.E9.97.A8.E6.88.90.E5.91.98.28.E8.AF.A6.E6.83.85.29"><span class="tocnumber">7</span> <span class="toctext"><b>获取部门成员(详情)</a></li>
<li class="toclevel-1 tocsection-8"><a href="#.E9.82.80.E8.AF.B7.E6.88.90.E5.91.98.E5.85.B3.E6.B3.A8"><span class="tocnumber">8</span> <span class="toctext"><b>邀请成员关注</a></li>
</ul>
</div>
### **创建成员**
<ul>
<li>请求说明</li>
</ul>
Https请求方式: POST

<a rel="nofollow" class="external free" href="https://qyapi.weixin.qq.com/cgi-bin/user/create?access_token=ACCESS_TOKEN">https://qyapi.weixin.qq.com/cgi-bin/user/create?access_token=ACCESS_TOKEN</a>

请求包结构体为:
<pre>{
   "userid": "zhangsan",
   "name": "张三",
   "department": [1, 2],
   "position": "产品经理",
   "mobile": "15913215421",
   "email": "zhangsan@gzdev.com",
   "weixinid": "zhangsan4dev",
   "extattr": {"attrs":[{"name":"爱好","value":"旅游"},{"name":"卡号","value":"1234567234"}]}
}
</pre>
<ul>
<li>参数说明</li>
</ul>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:120px">参数</th><th>必须</th><th>说明</th></tr>
<tr><td> access_token</td><td> 是</td><td> 调用接口凭证</td></tr>
<tr><td> userid</td><td> 是</td><td> 员工UserID。对应管理端的帐号，企业内必须唯一。长度为1~64个字符</td></tr>
<tr><td> name</td><td> 是</td><td> 成员名称。长度为1~64个字符</td></tr>
<tr><td> department</td><td> 否</td><td> 成员所属部门id列表。注意，每个部门的直属员工上限为1000个</td></tr>
<tr><td> position</td><td> 否</td><td> 职位信息。长度为0~64个字符</td></tr>
<tr><td> mobile</td><td> 否</td><td> 手机号码。企业内必须唯一，mobile/weixinid/email三者不能同时为空</td></tr>
<tr><td> email</td><td> 否</td><td> 邮箱。长度为0~64个字符。企业内必须唯一</td></tr>
<tr><td> weixinid</td><td> 否</td><td> 微信号。企业内必须唯一。（注意：是微信号，不是微信的名字）</td></tr>
<tr><td> extattr</td><td> 否</td><td> 扩展属性。扩展属性需要在WEB管理端创建后才生效，否则忽略未知属性的赋值</td></tr></tbody></table>
<ul>
<li>权限说明</li>
</ul>
管理员须拥有“操作通讯录”的接口权限，以及指定部门的管理权限。
<ul>
<li>返回结果</li>
</ul>
<pre>{
   "errcode": 0,
   "errmsg": "created"
}
</pre>

### **更新成员**
<ul>
<li>请求说明</li>
</ul>
Https请求方式: POST

<a rel="nofollow" class="external free" href="https://qyapi.weixin.qq.com/cgi-bin/user/update?access_token=ACCESS_TOKEN">https://qyapi.weixin.qq.com/cgi-bin/user/update?access_token=ACCESS_TOKEN</a>

请求包示例如下（如果非必须的字段未指定，则不更新该字段之前的设置值）:
<pre>{
   "userid": "zhangsan",
   "name": "李四",
   "department": [1],
   "position": "后台工程师",
   "mobile": "15913215421",
   "email": "zhangsan@gzdev.com",
   "weixinid": "lisifordev",
   "enable": 1,
   "extattr": {"attrs":[{"name":"爱好","value":"旅游"},{"name":"卡号","value":"1234567234"}]}
}
</pre>
<ul>
<li>参数说明</li>
</ul>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:120px">参数</th><th>必须</th><th>说明</th></tr>
<tr><td> access_token</td><td> 是</td><td> 调用接口凭证</td></tr>
<tr><td> userid</td><td> 是</td><td> 员工UserID。对应管理端的帐号，企业内必须唯一。长度为1~64个字符</td></tr>
<tr><td> name</td><td> 否</td><td> 成员名称。长度为0~64个字符</td></tr>
<tr><td> department</td><td> 否</td><td> 成员所属部门id列表。注意，每个部门的直属员工上限为1000个</td></tr>
<tr><td> position</td><td> 否</td><td> 职位信息。长度为0~64个字符</td></tr>
<tr><td> mobile</td><td> 否</td><td> 手机号码。企业内必须唯一，mobile/weixinid/email三者不能同时为空</td></tr>
<tr><td> email</td><td> 否</td><td> 邮箱。长度为0~64个字符。企业内必须唯一</td></tr>
<tr><td> weixinid</td><td> 否</td><td> 微信号。企业内必须唯一。（注意：是微信号，不是微信的名字）</td></tr>
<tr><td> enable</td><td> 否</td><td> 启用/禁用成员。1表示启用成员，0表示禁用成员</td></tr>
<tr><td> extattr</td><td> 否</td><td> 扩展属性。扩展属性需要在WEB管理端创建后才生效，否则忽略未知属性的赋值</td></tr></tbody></table>
<ul>
<li>权限说明</li>
</ul>
管理员须拥有“操作通讯录”的接口权限，以及指定部门、成员的管理权限。

<ul>
<li>返回结果</li>
</ul>
<pre>{
   "errcode": 0,
   "errmsg": "updated"
}
</pre>

### **删除成员**
<ul>
<li>请求说明</li>
</ul>
Https请求方式: GET

<a rel="nofollow" class="external free" href="https://qyapi.weixin.qq.com/cgi-bin/user/delete?access_token=ACCESS_TOKEN&amp;userid=lisi">https://qyapi.weixin.qq.com/cgi-bin/user/delete?access_token=ACCESS_TOKEN&amp;userid=lisi</a>
<ul>
<li>参数说明</li>
</ul>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:120px">参数</th><th>必须</th><th>说明</th></tr>
<tr><td> access_token</td><td> 是</td><td> 调用接口凭证</td></tr>
<tr><td> userid</td><td> 是</td><td> 员工UserID。对应管理端的帐号</td></tr></tbody></table>
<ul>
<li>权限说明</li>
</ul>
管理员须拥有“操作通讯录”的接口权限，以及指定部门、成员的管理权限。
<ul>
<li>返回结果</li>
</ul>
<pre>{
   "errcode": 0,
   "errmsg": "deleted"
}
</pre>

### **批量删除成员**
<ul>
<li>请求说明</li>
</ul>
Https请求方式: POST

<a rel="nofollow" class="external free" href="https://qyapi.weixin.qq.com/cgi-bin/user/batchdelete?access_token=ACCESS_TOKEN">https://qyapi.weixin.qq.com/cgi-bin/user/batchdelete?access_token=ACCESS_TOKEN</a>

请求包结构体为:
<pre>{
   "useridlist": ["zhangsan", "lisi"]
}
</pre>
<ul>
<li>参数说明</li>
</ul>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:120px">参数</th><th>必须</th><th>说明</th></tr>
<tr><td> access_token</td><td> 是</td><td> 调用接口凭证</td></tr>
<tr><td> useridlist</td><td> 是</td><td> 员工UserID列表。对应管理端的帐号</td></tr></tbody></table>
<ul>
<li>权限说明</li>
</ul>
管理员须拥有“操作通讯录”的接口权限，以及指定部门、成员的管理权限。
<ul>
<li>返回结果</li>
</ul>
<pre>{
   "errcode": 0,
   "errmsg": "deleted"
}
</pre>

### **获取成员**
<ul>
<li>请求说明</li>
</ul>
Https请求方式: GET

<a rel="nofollow" class="external free" href="https://qyapi.weixin.qq.com/cgi-bin/user/get?access_token=ACCESS_TOKEN&amp;userid=lisi">https://qyapi.weixin.qq.com/cgi-bin/user/get?access_token=ACCESS_TOKEN&amp;userid=lisi</a>
<ul>
<li>参数说明</li>
</ul>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:120px">参数</th><th>必须</th><th>说明</th></tr>
<tr><td> access_token</td><td> 是</td><td> 调用接口凭证</td></tr>
<tr><td> userid</td><td> 是</td><td> 员工UserID。对应管理端的帐号</td></tr></tbody></table>
<ul>
<li>返回结果</li>
</ul>
<pre>{
   "errcode": 0,
   "errmsg": "ok",
   "userid": "zhangsan",
   "name": "李四",
   "department": [1, 2],
   "position": "后台工程师",
   "mobile": "15913215421",
   "email": "zhangsan@gzdev.com",
   "weixinid": "lisifordev",  
   "avatar": <a rel="nofollow" class="external text" href="http://wx.qlogo.cn/mmopen/ajNVdqHZLLA3WJ6DSZUfiakYe37PKnQhBIeOQBO4czqrnZDS79FH5Wm5m4X69TBicnHFlhiafvDwklOpZeXYQQ2icg/0">"http://wx.qlogo.cn/mmopen/ajNVdqHZLLA3WJ6DSZUfiakYe37PKnQhBIeOQBO4czqrnZDS79FH5Wm5m4X69TBicnHFlhiafvDwklOpZeXYQQ2icg/0"</a>,
   "status": 1,
   "extattr": {"attrs":[{"name":"爱好","value":"旅游"},{"name":"卡号","value":"1234567234"}]}
}
</pre>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th><th>说明</th></tr>
<tr><td> errcode</td><td> 返回码</td></tr>
<tr><td> errmsg</td><td> 对返回码的文本描述内容</td></tr>
<tr><td> userid</td><td> 员工UserID。对应管理端的帐号</td></tr>
<tr><td> name</td><td> 成员名称</td></tr>
<tr><td> department</td><td> 成员所属部门id列表</td></tr>
<tr><td> position</td><td> 职位信息</td></tr>
<tr><td> mobile</td><td> 手机号码</td></tr>
<tr><td> email</td><td> 邮箱</td></tr>
<tr><td> weixinid</td><td> 微信号</td></tr>
<tr><td> avatar</td><td> 头像url。注：如果要获取小图将url最后的"/0"改成"/64"即可</td></tr>
<tr><td> status</td><td> 关注状态: 1=已关注，2=已冻结，4=未关注</td></tr>
<tr><td> extattr</td><td> 扩展属性</td></tr></tbody></table>
<ul>
<li>权限说明</li>
</ul>
管理员须拥有’获取成员’的接口权限，以及成员的查看权限。

### **获取部门成员**
<ul>
<li>请求说明</li>
</ul>
Https请求方式: GET

<a rel="nofollow" class="external free" href="https://qyapi.weixin.qq.com/cgi-bin/user/simplelist?access_token=ACCESS_TOKEN&amp;department_id=1&amp;fetch_child=0&amp;status=0">https://qyapi.weixin.qq.com/cgi-bin/user/simplelist?access_token=ACCESS_TOKEN&amp;department_id=1&amp;fetch_child=0&amp;status=0</a>
<ul>
<li>参数说明</li>
</ul>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:120px">参数</th><th>必须</th><th>说明</th></tr>
<tr><td> access_token</td><td> 是</td><td> 调用接口凭证</td></tr>
<tr><td> department_id</td><td> 是</td><td> 获取的部门id</td></tr>
<tr><td> fetch_child</td><td> 否</td><td> 1/0：是否递归获取子部门下面的成员</td></tr>
<tr><td> status</td><td> 否</td><td> 0获取全部员工，1获取已关注成员列表，2获取禁用成员列表，4获取未关注成员列表。status可叠加</td></tr></tbody></table>
<ul>
<li>权限说明</li>
</ul>
管理员须拥有’获取部门成员’的接口权限，以及指定部门的查看权限。
<ul>
<li>返回结果</li>
</ul>
<pre>{
   "errcode": 0,
   "errmsg": "ok",
   "userlist": [
           {
                  "userid": "zhangsan",
                  "name": "李四"
           }
     ]
}
</pre>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th><th>说明</th></tr>
<tr><td> errcode</td><td> 返回码</td></tr>
<tr><td> errmsg</td><td> 对返回码的文本描述内容</td></tr>
<tr><td> userlist</td><td> 成员列表</td></tr>
<tr><td> userid</td><td> 员工UserID。对应管理端的帐号</td></tr>
<tr><td> name</td><td> 成员名称</td></tr></tbody></table>

### **获取部门成员(详情)**
<ul>
<li>请求说明</li>
</ul>
Https请求方式: GET

<a rel="nofollow" class="external free" href="https://qyapi.weixin.qq.com/cgi-bin/user/list?access_token=ACCESS_TOKEN&amp;department_id=1&amp;fetch_child=0&amp;status=0">https://qyapi.weixin.qq.com/cgi-bin/user/list?access_token=ACCESS_TOKEN&amp;department_id=1&amp;fetch_child=0&amp;status=0</a>
<ul>
<li>参数说明</li>
</ul>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:120px">参数</th><th>必须</th><th>说明</th></tr>
<tr><td> access_token</td><td> 是</td><td> 调用接口凭证</td></tr>
<tr><td> department_id</td><td> 是</td><td> 获取的部门id</td></tr>
<tr><td> fetch_child</td><td> 否</td><td> 1/0：是否递归获取子部门下面的成员</td></tr>
<tr><td> status</td><td> 否</td><td> 0获取全部员工，1获取已关注成员列表，2获取禁用成员列表，4获取未关注成员列表。status可叠加</td></tr></tbody></table>
<ul>
<li>权限说明</li>
</ul>
管理员须拥有’获取部门成员’的接口权限，以及指定部门的查看权限。
<ul>
<li>返回结果</li>
</ul>
<pre>{
   "errcode": 0,
   "errmsg": "ok",
   "userlist": [
           {
                  "userid": "zhangsan",
                  "name": "李四",
                  "department": [1, 2],
                  "position": "后台工程师",
                  "mobile": "15913215421",
                  "email": "zhangsan@gzdev.com",
                  "weixinid": "lisifordev",  
                  "avatar":           "<a rel="nofollow" class="external free" href="http://wx.qlogo.cn/mmopen/ajNVdqHZLLA3WJ6DSZUfiakYe37PKnQhBIeOQBO4czqrnZDS79FH5Wm5m4X69TBicnHFlhiafvDwklOpZeXYQQ2icg/0">http://wx.qlogo.cn/mmopen/ajNVdqHZLLA3WJ6DSZUfiakYe37PKnQhBIeOQBO4czqrnZDS79FH5Wm5m4X69TBicnHFlhiafvDwklOpZeXYQQ2icg/0</a>",
                  "status": 1,
                  "extattr": {"attrs":[{"name":"爱好","value":"旅游"},{"name":"卡号","value":"1234567234"}]}
           }
     ]
}
</pre>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th><th>说明</th></tr>
<tr><td> errcode</td><td> 返回码</td></tr>
<tr><td> errmsg</td><td> 对返回码的文本描述内容</td></tr>
<tr><td> userlist</td><td> 成员列表</td></tr>
<tr><td> userid</td><td> 员工UserID。对应管理端的帐号</td></tr>
<tr><td> name</td><td> 成员名称</td></tr>
<tr><td> department</td><td> 成员所属部门id列表</td></tr>
<tr><td> position</td><td> 职位信息</td></tr>
<tr><td> mobile</td><td> 手机号码</td></tr>
<tr><td> email</td><td> 邮箱</td></tr>
<tr><td> weixinid</td><td> 微信号</td></tr>
<tr><td> avatar</td><td> 头像url。注：如果要获取小图将url最后的"/0"改成"/64"即可</td></tr>
<tr><td> status</td><td> 关注状态: 1=已关注，2=已冻结，4=未关注</td></tr>
<tr><td> extattr</td><td> 扩展属性</td></tr></tbody></table>

### **邀请成员关注**
<ul>
<li>接口说明</li>
</ul>
认证号优先使用微信推送邀请关注，如果没有weixinid字段则依次对手机号，邮箱绑定的微信进行推送，全部没有匹配则通过邮件邀请关注。 邮箱字段无效则邀请失败。
非认证号只通过邮件邀请关注。邮箱字段无效则邀请失败。
已关注以及被禁用用户不允许发起邀请关注请求。
<ul>
<li>请求说明</li>
</ul>
Https请求方式: POST

<a rel="nofollow" class="external free" href="https://qyapi.weixin.qq.com/cgi-bin/invite/send?access_token=ACCESS_TOKEN">https://qyapi.weixin.qq.com/cgi-bin/invite/send?access_token=ACCESS_TOKEN</a>

请求包结构体为:
<pre>{
   "userid":"xxxxx",
   "invite_tips":"xxx"
</pre>
}
<ul>
<li>参数说明</li>
</ul>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:120px">参数</th><th>必须</th><th>说明</th></tr>
<tr><td> access_token</td><td> 是</td><td> 调用接口凭证</td></tr>
<tr><td> userid</td><td> 是</td><td> 用户的userid</td></tr>
<tr><td> invite_tips</td><td> 否</td><td> 推送到微信上的提示语（只有认证号可以使用）。当使用微信推送时，该字段默认为“请关注XXX企业号”，邮件邀请时，该字段无效。</td></tr></tbody></table>
<ul>
<li>权限说明</li>
</ul>
管理员须拥有该成员的查看权限
<ul>
<li>返回结果</li>
</ul>
<pre>{
   "errcode": 0,
   "errmsg": "ok",
   "type":1
}
</pre>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th><th>说明</th></tr>
<tr><td> errcode</td><td> 返回码</td></tr>
<tr><td> errmsg</td><td> 对返回码的文本描述内容</td></tr>
<tr><td> type</td><td> 1:微信邀请 2.邮件邀请</td></tr></tbody></table>

<!--
NewPP limit report
CPU time usage: 0.056 seconds
Real time usage: 0.062 seconds
Preprocessor visited node count: 38/1000000
Preprocessor generated node count: 52/1000000
Post‐expand include size: 0/2097152 bytes
Template argument size: 0/2097152 bytes
Highest expansion depth: 2/40
Expensive parser function count: 0/100
-->
<!-- Saved in parser cache with key db_wiki:pcache:idhash:8-0!*!*!!zh-cn!*!* and timestamp 20150115092226 and revision id 665
 -->
## 管理标签
<div id="toc" class="toc"><div id="toctitle">## 目录</div>
<ul>
<li class="toclevel-1 tocsection-1"><a href="#.E5.88.9B.E5.BB.BA.E6.A0.87.E7.AD.BE"><span class="tocnumber">1</span> <span class="toctext"><b>创建标签</a></li>
<li class="toclevel-1 tocsection-2"><a href="#.E6.9B.B4.E6.96.B0.E6.A0.87.E7.AD.BE.E5.90.8D.E5.AD.97"><span class="tocnumber">2</span> <span class="toctext"><b>更新标签名字</a></li>
<li class="toclevel-1 tocsection-3"><a href="#.E5.88.A0.E9.99.A4.E6.A0.87.E7.AD.BE"><span class="tocnumber">3</span> <span class="toctext"><b>删除标签</a></li>
<li class="toclevel-1 tocsection-4"><a href="#.E8.8E.B7.E5.8F.96.E6.A0.87.E7.AD.BE.E6.88.90.E5.91.98"><span class="tocnumber">4</span> <span class="toctext"><b>获取标签成员</a></li>
<li class="toclevel-1 tocsection-5"><a href="#.E5.A2.9E.E5.8A.A0.E6.A0.87.E7.AD.BE.E6.88.90.E5.91.98"><span class="tocnumber">5</span> <span class="toctext"><b>增加标签成员</a></li>
<li class="toclevel-1 tocsection-6"><a href="#.E5.88.A0.E9.99.A4.E6.A0.87.E7.AD.BE.E6.88.90.E5.91.98"><span class="tocnumber">6</span> <span class="toctext"><b>删除标签成员</a></li>
<li class="toclevel-1 tocsection-7"><a href="#.E8.8E.B7.E5.8F.96.E6.A0.87.E7.AD.BE.E5.88.97.E8.A1.A8"><span class="tocnumber">7</span> <span class="toctext"><b>获取标签列表</a></li>
</ul>
</div>
### **创建标签**
<ul>
<li>请求说明</li>
</ul>
Https请求方式: POST

<a rel="nofollow" class="external free" href="https://qyapi.weixin.qq.com/cgi-bin/tag/create?access_token=ACCESS_TOKEN">https://qyapi.weixin.qq.com/cgi-bin/tag/create?access_token=ACCESS_TOKEN</a>

请求包结构体为:
<pre>{
   "tagname": "UI"
}
</pre>
<ul>
<li>参数说明</li>
</ul>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:120px">参数</th><th>必须</th><th>说明</th></tr>
<tr><td> access_token</td><td> 是</td><td> 调用接口凭证</td></tr>
<tr><td> tagname</td><td> 是</td><td> 标签名称。长度为1~64个字符，标签不可与其他同组的标签重名，也不可与全局标签重名</td></tr></tbody></table>
<ul>
<li>权限说明</li>
</ul>
创建的标签属于管理组;默认为未加锁状态。
<ul>
<li>返回结果</li>
</ul>
<pre>{
   "errcode": 0,
   "errmsg": "created"
   "tagid": "1"
}
</pre>

### **更新标签名字**
<ul>
<li>请求说明</li>
</ul>
Https请求方式: POST

<a rel="nofollow" class="external free" href="https://qyapi.weixin.qq.com/cgi-bin/tag/update?access_token=ACCESS_TOKEN">https://qyapi.weixin.qq.com/cgi-bin/tag/update?access_token=ACCESS_TOKEN</a>

请求包示例如下:
<pre>{
   "tagid": "1",
   "tagname": "UI design"
}
</pre>
<ul>
<li>参数说明</li>
</ul>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:120px">参数</th><th>必须</th><th>说明</th></tr>
<tr><td> access_token</td><td> 是</td><td> 调用接口凭证</td></tr>
<tr><td> tagid</td><td> 是</td><td> 标签ID</td></tr>
<tr><td> tagname</td><td> 是</td><td> 标签名称。长度为1~64个字符，标签不可与其他同组的标签重名，也不可与全局标签重名</td></tr></tbody></table>
<ul>
<li>权限说明</li>
</ul>
管理组必须是指定标签的创建者。
<ul>
<li>返回结果</li>
</ul>
<pre>{
   "errcode": 0,
   "errmsg": "updated"
}
</pre>

### **删除标签**
<ul>
<li>请求说明</li>
</ul>
Https请求方式: GET

<a rel="nofollow" class="external free" href="https://qyapi.weixin.qq.com/cgi-bin/tag/delete?access_token=ACCESS_TOKEN&amp;tagid=1">https://qyapi.weixin.qq.com/cgi-bin/tag/delete?access_token=ACCESS_TOKEN&amp;tagid=1</a>
<ul>
<li>参数说明</li>
</ul>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:120px">参数</th><th>必须</th><th>说明</th></tr>
<tr><td> access_token</td><td> 是</td><td> 调用接口凭证</td></tr>
<tr><td> tagid</td><td> 是</td><td> 标签ID</td></tr></tbody></table>
<ul>
<li>权限说明</li>
</ul>
管理组必须是指定标签的创建者，并且标签的成员列表为空。
<ul>
<li>返回结果</li>
</ul>
<pre>{
   "errcode": 0,
   "errmsg": "deleted"
}
</pre>

### **获取标签成员**
<ul>
<li>请求说明</li>
</ul>
Https请求方式: GET

<a rel="nofollow" class="external free" href="https://qyapi.weixin.qq.com/cgi-bin/tag/get?access_token=ACCESS_TOKEN&amp;tagid=1">https://qyapi.weixin.qq.com/cgi-bin/tag/get?access_token=ACCESS_TOKEN&amp;tagid=1</a>
<ul>
<li>参数说明</li>
</ul>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:120px">参数</th><th>必须</th><th>说明</th></tr>
<tr><td> access_token</td><td> 是</td><td> 调用接口凭证</td></tr>
<tr><td> tagid</td><td> 是</td><td> 标签ID</td></tr></tbody></table>
<ul>
<li>权限说明</li>
</ul>
管理组须拥有“获取标签成员”的接口权限，标签须对管理组可见；返回列表仅包含管理组管辖范围的成员。
<ul>
<li>返回结果</li>
</ul>
<pre>{
   "errcode": 0,
   "errmsg": "ok",
   "userlist": [
         {
             "userid": "zhangsan",
             "name": "李四"
         }
     ],
   "partylist": [2]
}
</pre>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th><th>说明</th></tr>
<tr><td> errcode</td><td> 错误码</td></tr>
<tr><td> errmsg</td><td> 错误消息</td></tr>
<tr><td> userlist</td><td> 成员列表</td></tr>
<tr><td> userlist::userid</td><td> 员工UserID</td></tr>
<tr><td> userlist::name</td><td> 员工姓名</td></tr>
<tr><td> partylist</td><td> 部门列表</td></tr></tbody></table>

### **增加标签成员**
<ul>
<li>请求说明</li>
</ul>
Https请求方式: POST

<a rel="nofollow" class="external free" href="https://qyapi.weixin.qq.com/cgi-bin/tag/addtagusers?access_token=ACCESS_TOKEN">https://qyapi.weixin.qq.com/cgi-bin/tag/addtagusers?access_token=ACCESS_TOKEN</a>

请求包示例如下:
<pre>{
   "tagid": "1",
   "userlist":[ "user1","user2"],
   "partylist": [4]
}
</pre>
<ul>
<li>参数说明</li>
</ul>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:120px">参数</th><th>必须</th><th>说明</th></tr>
<tr><td> access_token</td><td> 是</td><td> 调用接口凭证</td></tr>
<tr><td> tagid</td><td> 是</td><td> 标签ID</td></tr>
<tr><td> userlist</td><td> 否</td><td> 企业员工ID列表，注意：userlist、partylist不能同时为空</td></tr>
<tr><td> partylist</td><td> 否</td><td> 企业部门ID列表，注意：userlist、partylist不能同时为空</td></tr></tbody></table>
<ul>
<li>权限说明</li>
</ul>
标签对管理组可见且未加锁，成员属于管理组管辖范围。
<ul>
<li>返回结果</li>
</ul>
a)正确时返回
<pre>{
   "errcode": 0,
   "errmsg": "ok"
}
</pre>
b)若部分userid、partylist非法，则返回
<pre>{
   "errcode": 0,
   "errmsg": "错误消息",
   "invalidlist"："usr1|usr2|usr",
   "invalidparty"：[2,4]
}
</pre>
其中错误消息视具体出错情况而定，分别为：<br>
invalid userlist and partylist faild<br>
invalid userlist faild<br>
invalid partylist faild<br>

c)当包含userid、partylist全部非法时返回
<pre>{
   "errcode": 40070,
   "errmsg": "all list invalid "
}
</pre>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th><th>说明</th></tr>
<tr><td> errcode</td><td> 错误码</td></tr>
<tr><td> errmsg</td><td> 错误消息</td></tr>
<tr><td> invalidlist</td><td> 不在权限内的员工ID列表，以“|”分隔</td></tr>
<tr><td> invalidparty</td><td> 不在权限内的部门ID列表</td></tr></tbody></table>

### **删除标签成员**
<ul>
<li>请求说明</li>
</ul>
Https请求方式: POST

<a rel="nofollow" class="external free" href="https://qyapi.weixin.qq.com/cgi-bin/tag/deltagusers?access_token=ACCESS_TOKEN">https://qyapi.weixin.qq.com/cgi-bin/tag/deltagusers?access_token=ACCESS_TOKEN</a>

请求包如下
<pre>{
   "tagid": "1",
   "userlist":[ "user1","user2"],
   "partylist":[2,4]
}
</pre>
<ul>
<li>参数说明</li>
</ul>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:120px">参数</th><th>必须</th><th>说明</th></tr>
<tr><td> access_token</td><td> 是</td><td> 调用接口凭证</td></tr>
<tr><td> tagid</td><td> 是</td><td> 标签ID</td></tr>
<tr><td> userlist</td><td> 否</td><td> 企业员工ID列表，注意：userlist、partylist不能同时为空</td></tr>
<tr><td> partylist</td><td> 否</td><td> 企业部门ID列表，注意：userlist、partylist不能同时为空</td></tr></tbody></table>
<ul>
<li>权限说明</li>
</ul>
标签对管理组可见且未加锁，成员属于管理组管辖范围。
<ul>
<li>返回结果</li>
</ul>
a)正确时返回
<pre>{
   "errcode": 0,
   "errmsg": "deleted"
}
</pre>
b)若部分userid、partylist非法，则返回
<pre>{
   "errcode": 0,
   "errmsg": "错误消息",
   "invalidlist"："usr1|usr2|usr",
   "invalidparty": [2,4]
}
</pre>
其中错误消息视具体出错情况而定，分别为：<br>
invalid userlist and partylist faild<br>
invalid userlist faild<br>
invalid partylist faild<br>

c)当包含的userid、partylist全部非法时返回
<pre>{
   "errcode": 40031,
   "errmsg": "all list invalid"
}
</pre>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th><th>说明</th></tr>
<tr><td> errcode</td><td> 错误码</td></tr>
<tr><td> errmsg</td><td> 错误消息</td></tr>
<tr><td> invalidlist</td><td> 不在权限内的或者非法的员工ID列表，以“|”分隔</td></tr>
<tr><td> invalidparty</td><td> 不在权限内的部门ID列表</td></tr></tbody></table>

### **获取标签列表**
<ul>
<li>请求说明</li>
</ul>
Https请求方式: GET

<a rel="nofollow" class="external free" href="https://qyapi.weixin.qq.com/cgi-bin/tag/list?access_token=ACCESS_TOKEN">https://qyapi.weixin.qq.com/cgi-bin/tag/list?access_token=ACCESS_TOKEN</a>
<ul>
<li>参数说明</li>
</ul>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:120px">参数</th><th>必须</th><th>说明</th></tr>
<tr><td> access_token</td><td> 是</td><td> 调用接口凭证</td></tr></tbody></table>
<ul>
<li>权限说明</li>
</ul>
标签对管理组可见
<ul>
<li>返回结果</li>
</ul>
<pre>{
   "errcode": 0,
   "errmsg": "ok",
   "taglist":[
      {"tagid":1,"tagname":"a"},
      {"tagid":2,"tagname":"b"}
   ]
}
</pre>
<!--
NewPP limit report
CPU time usage: 0.036 seconds
Real time usage: 0.034 seconds
Preprocessor visited node count: 50/1000000
Preprocessor generated node count: 84/1000000
Post‐expand include size: 0/2097152 bytes
Template argument size: 0/2097152 bytes
Highest expansion depth: 2/40
Expensive parser function count: 0/100
-->
<!-- Saved in parser cache with key db_wiki:pcache:idhash:10-0!*!*!!zh-cn!*!* and timestamp 20150115132052 and revision id 548
 -->
#　管理多媒体文件

企业在使用接口时，对多媒体文件、多媒体消息的获取和调用等操作，是通过media_id来进行的。通过本接口，企业可以上传或下载多媒体文件。

注意，**每个多媒体文件（media_id）会在上传到微信服务器3天后自动删除**，以节省服务器资源。

<!--
NewPP limit report
CPU time usage: 0.000 seconds
Real time usage: 0.001 seconds
Preprocessor visited node count: 1/1000000
Preprocessor generated node count: 4/1000000
Post‐expand include size: 0/2097152 bytes
Template argument size: 0/2097152 bytes
Highest expansion depth: 1/40
Expensive parser function count: 0/100
-->
<!-- Saved in parser cache with key db_wiki:pcache:idhash:37-0!*!*!*!*!*!* and timestamp 20150115124320 and revision id 238
 -->
## 上传媒体文件
### **上传媒体文件**
用于上传图片、语音、视频等媒体资源文件以及普通文件（如doc，ppt），接口返回媒体资源标识ID：media_id。**请注意，media_id是可复用的，同一个media_id可用于消息的多次发送(5天内有效)**。
<ul>
<li>请求说明</li>
</ul>
Https请求方式: POST

<a rel="nofollow" class="external free" href="https://qyapi.weixin.qq.com/cgi-bin/media/upload?access_token=ACCESS_TOKEN&amp;type=TYPE">https://qyapi.weixin.qq.com/cgi-bin/media/upload?access_token=ACCESS_TOKEN&amp;type=TYPE</a>
<ul>
<li>参数说明</li>
</ul>
<table border="1" cellspacing="0" cellpadding="3" align="center" width="640px">
<tbody><tr>
<th style="width:120px">参数</th><th>必须</th><th>说明</th></tr>
<tr><td> access_token</td><td> 是</td><td> 调用接口凭证</td></tr>
<tr><td> type</td><td> 是</td><td> 媒体文件类型，分别有图片（image）、语音（voice）、视频（video），普通文件(file)</td></tr>
<tr><td> media</td><td> 是</td><td> form-data中媒体文件标识，有filename、filelength、content-type等信息</td></tr></tbody></table>
<ul>
<li>权限说明</li>
</ul>
完全公开。所有管理员均可调用，media_id可以共享。
<ul>
<li>返回说明</li>
</ul>
<pre>{
   "type": "image",
   "media_id": "0000001",
   "created_at": "1380000000"
}
</pre>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th><th>说明</th></tr>
<tr><td> type</td><td> 媒体文件类型，分别有图片（image）、语音（voice）、视频（video）,普通文件(file)</td></tr>
<tr><td> media_id</td><td> 媒体文件上传后获取的唯一标识</td></tr>
<tr><td> created_at</td><td> 媒体文件上传时间戳</td></tr></tbody></table>
<ul>
<li>上传的媒体文件限制</li>
</ul>
图片（image）:1MB，支持JPG格式

语音（voice）：2MB，播放长度不超过60s，支持AMR格式

视频（video）：10MB，支持MP4格式

普通文件（file）：10MB

<!--
NewPP limit report
CPU time usage: 0.012 seconds
Real time usage: 0.015 seconds
Preprocessor visited node count: 2/1000000
Preprocessor generated node count: 8/1000000
Post‐expand include size: 0/2097152 bytes
Template argument size: 0/2097152 bytes
Highest expansion depth: 2/40
Expensive parser function count: 0/100
-->
<!-- Saved in parser cache with key db_wiki:pcache:idhash:11-0!*!*!*!*!*!* and timestamp 20150115112143 and revision id 540
 -->
## 获取媒体文件
## **获取媒体文件**
通过media_id获取图片、语音、视频等文件，协议和普通的http文件下载完全相同。
<ul>
<li>请求说明</li>
</ul>
Https请求方式: GET

<a rel="nofollow" class="external free" href="https://qyapi.weixin.qq.com/cgi-bin/media/get?access_token=ACCESS_TOKEN&amp;media_id=MEDIA_ID">https://qyapi.weixin.qq.com/cgi-bin/media/get?access_token=ACCESS_TOKEN&amp;media_id=MEDIA_ID</a>
<ul>
<li>参数说明</li>
</ul>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:120px">参数</th><th>必须</th><th>说明</th></tr>
<tr><td> access_token</td><td> 是</td><td> 调用接口凭证</td></tr>
<tr><td> media_id</td><td> 是</td><td> 媒体文件id</td></tr></tbody></table>
<ul>
<li>权限说明</li>
</ul>
完全公开。所有管理员均可调用，media_id可以共享。
<ul>
<li>返回结果</li>
</ul>
和普通的http下载相同，请根据http头做相应的处理。

a)正确时返回：
<pre>{
   HTTP/1.1 200 OK
   Connection: close
   Content-Type: image/jpeg
   Content-disposition: attachment; filename="MEDIA_ID.jpg"
   Date: Sun, 06 Jan 2013 10:20:18 GMT
   Cache-Control: no-cache, must-revalidate
   Content-Length: 339721

   Xxxx
}
</pre>
b)错误时返回（这里省略了HTTP首部）：
<pre>{
   "errcode": "40004",
   "errmsg": "invalid media_id"
}
</pre>

<!--
NewPP limit report
CPU time usage: 0.008 seconds
Real time usage: 0.007 seconds
Preprocessor visited node count: 3/1000000
Preprocessor generated node count: 10/1000000
Post‐expand include size: 0/2097152 bytes
Template argument size: 0/2097152 bytes
Highest expansion depth: 2/40
Expensive parser function count: 0/100
-->
<!-- Saved in parser cache with key db_wiki:pcache:idhash:12-0!*!*!*!*!*!* and timestamp 20150116060832 and revision id 664
 -->
# 接收消息与事件

将应用设置在回调模式时，企业可以通过回调URL接收员工回复的消息，以及员工关注、点击菜单、上报地理位置等事件。

在接收到事件后，企业可以发送被动响应消息，实现员工与企业的互动。

企业在接收消息，以及发送被动响应消息时，数据包以xml格式组成，以AES方式加密传输。具体可参考'建立连接'中的'回调模式'一节。

<!--
NewPP limit report
CPU time usage: 0.000 seconds
Real time usage: 0.001 seconds
Preprocessor visited node count: 1/1000000
Preprocessor generated node count: 4/1000000
Post‐expand include size: 0/2097152 bytes
Template argument size: 0/2097152 bytes
Highest expansion depth: 1/40
Expensive parser function count: 0/100
-->
<!-- Saved in parser cache with key db_wiki:pcache:idhash:49-0!*!*!*!*!*!* and timestamp 20150116013226 and revision id 513
 -->
## 关注与取消关注
<div id="toc" class="toc"><div id="toctitle">## 目录</div>
<ul>
<li class="toclevel-1 tocsection-1"><a href="#.E5.85.B3.E6.B3.A8.E4.B8.8E.E5.8F.96.E6.B6.88.E5.85.B3.E6.B3.A8"><span class="tocnumber">1</span> <span class="toctext"><b>关注与取消关注</a>
<ul>
<li class="toclevel-2 tocsection-2"><a href="#.E5.91.98.E5.B7.A5.E4.B8.8E.E9.80.9A.E8.AE.AF.E5.BD.95.E4.B8.AD.E7.9A.84.E5.B8.90.E5.8F.B7.E7.BB.91.E5.AE.9A"><span class="tocnumber">1.1</span> <span class="toctext"><b>员工与通讯录中的帐号绑定</a></li>
<li class="toclevel-2 tocsection-3"><a href="#.E4.BA.8C.E6.AC.A1.E9.AA.8C.E8.AF.81"><span class="tocnumber">1.2</span> <span class="toctext"><b>二次验证</a></li>
<li class="toclevel-2 tocsection-4"><a href="#.E5.85.B3.E6.B3.A8.2F.E5.8F.96.E6.B6.88.E5.85.B3.E6.B3.A8.E4.BA.8B.E4.BB.B6.E7.9A.84.E6.8E.A8.E9.80.81"><span class="tocnumber">1.3</span> <span class="toctext"><b>关注/取消关注事件的推送</a></li>
</ul></li>
</ul>
</div>
## **关注与取消关注**
员工在关注企业号时，首先要与企业通讯录中的帐号绑定；如果企业开启了二次验证，那么在绑定成功后还需要经过企业的验证，才可以关注成功。

### **员工与通讯录中的帐号绑定**
员工关注企业号时，会根据员工的微信号、微信绑定的手机或邮箱，与企业通讯录的帐号匹配。如果匹配到，则绑定成功；否则会下发一条图文消息给员工，引导员工在页面上验证手机号或邮箱，验证后即绑定成功。注意，员工的微信版本需要在5.4以上，目前仅支持iOS、Android两个平台。

### **二次验证**
企业在开启二次验证时，必须填写企业二次验证页面的url。当员工绑定通讯录中的帐号后，会收到一条图文消息，引导员工到企业的验证页面验证身份。在跳转到企业的验证页面时，会带上如下参数：code=CODE&amp;state=STATE，企业可以调用oauth2接口，根据code获取员工的userid。

企业在员工验证成功后，调用如下接口即可让员工关注成功。
<ul>
<li>请求说明</li>
</ul>
Https请求方式: GET

<a rel="nofollow" class="external free" href="https://qyapi.weixin.qq.com/cgi-bin/user/authsucc?access_token=ACCESS_TOKEN&amp;userid=USERID">https://qyapi.weixin.qq.com/cgi-bin/user/authsucc?access_token=ACCESS_TOKEN&amp;userid=USERID</a>
<ul>
<li>参数说明</li>
</ul>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:120px">参数</th><th>必须</th><th>说明</th></tr>
<tr><td> access_token</td><td> 是</td><td> 调用接口凭证</td></tr>
<tr><td> userid</td><td> 是</td><td> 员工UserID</td></tr></tbody></table>
<ul>
<li>权限说明</li>
</ul>
管理员须拥有userid对应员工的管理权限。
<ul>
<li>返回结果</li>
</ul>
<pre>{
   "errcode": "0",
   "errmsg": "ok"
}
</pre>

### **关注/取消关注事件的推送**
员工关注、取消关注企业号的事件，会推送到每个应用在管理端设置的URL；特别的，默认企业小助手可以用于获取整个企业号的关注状况。（以下假设该URL为<a rel="nofollow" class="external free" href="http://api.3dept.com">http://api.3dept.com</a>）。
<ul>
<li>请求说明</li>
</ul>
Http请求方式: POST

<a rel="nofollow" class="external free" href="http://api.3dept.com/?msg_signature=ASDFQWEXZCVAQFASDFASDFSS&amp;timestamp=13500001234&amp;nonce=123412323">http://api.3dept.com/?msg_signature=ASDFQWEXZCVAQFASDFASDFSS&amp;timestamp=13500001234&amp;nonce=123412323</a>
<ul>
<li>参数说明</li>
</ul>
<pre>&lt;xml&gt;
   &lt;ToUserName&gt;&lt;![CDATA[toUser]]&gt;&lt;/ToUserName&gt;
   &lt;FromUserName&gt;&lt;![CDATA[UserID]]&gt;&lt;/FromUserName&gt;
   &lt;CreateTime&gt;1348831860&lt;/CreateTime&gt;
   &lt;MsgType&gt;&lt;![CDATA[event]]&gt;&lt;/MsgType&gt;
   &lt;Event&gt;&lt;![CDATA[subscribe]]&gt;&lt;/Event&gt;
   &lt;AgentID&gt;1&lt;/AgentID&gt;
&lt;/xml&gt;
</pre>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th><th>说明</th></tr>
<tr><td> ToUserName</td><td> 企业号CorpID</td></tr>
<tr><td> FromUserName</td><td> 员工UserID</td></tr>
<tr><td> CreateTime</td><td> 消息创建时间 （整型）</td></tr>
<tr><td> MsgType</td><td> 消息类型，此时固定为：event</td></tr>
<tr><td> Event</td><td> 事件类型，subscribe(订阅)、unsubscribe(取消订阅)</td></tr>
<tr><td> AgentID</td><td> 企业应用的id，整型。可在应用的设置页面获取；如果id为0，则表示是整个企业号的关注/取消关注事件</td></tr></tbody></table>

<!--
NewPP limit report
CPU time usage: 0.008 seconds
Real time usage: 0.011 seconds
Preprocessor visited node count: 18/1000000
Preprocessor generated node count: 28/1000000
Post‐expand include size: 0/2097152 bytes
Template argument size: 0/2097152 bytes
Highest expansion depth: 2/40
Expensive parser function count: 0/100
-->
<!-- Saved in parser cache with key db_wiki:pcache:idhash:13-0!*!*!!zh-cn!*!* and timestamp 20150116075349 and revision id 608
 -->
## 接受普通信息
<div id="toc" class="toc"><div id="toctitle">## 目录</div>
<ul>
<li class="toclevel-1 tocsection-1"><a href="#.E6.8E.A5.E6.94.B6.E6.99.AE.E9.80.9A.E6.B6.88.E6.81.AF"><span class="tocnumber">1</span> <span class="toctext"><b>接收普通消息</a>
<ul>
<li class="toclevel-2 tocsection-2"><a href="#text.E6.B6.88.E6.81.AF"><span class="tocnumber">1.1</span> <span class="toctext"><b>text消息</a></li>
<li class="toclevel-2 tocsection-3"><a href="#image.E6.B6.88.E6.81.AF"><span class="tocnumber">1.2</span> <span class="toctext"><b>image消息</a></li>
<li class="toclevel-2 tocsection-4"><a href="#voice.E6.B6.88.E6.81.AF"><span class="tocnumber">1.3</span> <span class="toctext"><b>voice消息</a></li>
<li class="toclevel-2 tocsection-5"><a href="#video.E6.B6.88.E6.81.AF"><span class="tocnumber">1.4</span> <span class="toctext"><b>video消息</a></li>
<li class="toclevel-2 tocsection-6"><a href="#location.E6.B6.88.E6.81.AF"><span class="tocnumber">1.5</span> <span class="toctext"><b>location消息</a></li>
</ul></li>
</ul>
</div>
## **接收普通消息**
普通消息是指员工向企业号应用发送的消息，包括文本、图片、语音、视频、地理位置等类型。普通消息会推送到每个应用在管理端设置的URL（以下假设该URL为<a rel="nofollow" class="external free" href="http://api.3dept.com">http://api.3dept.com</a>）。
<ul>
<li>请求说明</li>
</ul>
Http请求方式: POST

<a rel="nofollow" class="external free" href="http://api.3dept.com/?msg_signature=ASDFQWEXZCVAQFASDFASDFSS&amp;timestamp=13500001234&amp;nonce=123412323">http://api.3dept.com/?msg_signature=ASDFQWEXZCVAQFASDFASDFSS&amp;timestamp=13500001234&amp;nonce=123412323</a>

### **text消息**
<pre>&lt;xml&gt;
   &lt;ToUserName&gt;&lt;![CDATA[toUser]]&gt;&lt;/ToUserName&gt;
   &lt;FromUserName&gt;&lt;![CDATA[fromUser]]&gt;&lt;/FromUserName&gt;
   &lt;CreateTime&gt;1348831860&lt;/CreateTime&gt;
   &lt;MsgType&gt;&lt;![CDATA[text]]&gt;&lt;/MsgType&gt;
   &lt;Content&gt;&lt;![CDATA[this is a test]]&gt;&lt;/Content&gt;
   &lt;MsgId&gt;1234567890123456&lt;/MsgId&gt;
   &lt;AgentID&gt;1&lt;/AgentID&gt;
&lt;/xml&gt;
</pre>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th><th>说明</th></tr>
<tr><td> ToUserName</td><td> 企业号CorpID</td></tr>
<tr><td> FromUserName</td><td> 员工UserID</td></tr>
<tr><td> CreateTime</td><td> 消息创建时间（整型）</td></tr>
<tr><td> MsgType</td><td> 消息类型，此时固定为：text</td></tr>
<tr><td> Content</td><td> 文本消息内容</td></tr>
<tr><td> MsgId</td><td> 消息id，64位整型</td></tr>
<tr><td> AgentID</td><td> 企业应用的id，整型。可在应用的设置页面查看</td></tr></tbody></table>

### **image消息**
<pre>&lt;xml&gt;
   &lt;ToUserName&gt;&lt;![CDATA[toUser]]&gt;&lt;/ToUserName&gt;
   &lt;FromUserName&gt;&lt;![CDATA[fromUser]]&gt;&lt;/FromUserName&gt;
   &lt;CreateTime&gt;1348831860&lt;/CreateTime&gt;
   &lt;MsgType&gt;&lt;![CDATA[image]]&gt;&lt;/MsgType&gt;
   &lt;PicUrl&gt;&lt;![CDATA[this is a url]]&gt;&lt;/PicUrl&gt;
   &lt;MediaId&gt;&lt;![CDATA[media_id]]&gt;&lt;/MediaId&gt;
   &lt;MsgId&gt;1234567890123456&lt;/MsgId&gt;
   &lt;AgentID&gt;1&lt;/AgentID&gt;
&lt;/xml&gt;
</pre>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th><th>说明</th></tr>
<tr><td> ToUserName</td><td> 企业号CorpID</td></tr>
<tr><td> FromUserName</td><td> 员工UserID</td></tr>
<tr><td> CreateTime</td><td> 消息创建时间（整型）</td></tr>
<tr><td> MsgType</td><td> 消息类型，此时固定为：image</td></tr>
<tr><td> PicUrl</td><td> 图片链接</td></tr>
<tr><td> MediaId</td><td> 图片媒体文件id，可以调用获取媒体文件接口拉取数据</td></tr>
<tr><td> MsgId</td><td> 消息id，64位整型</td></tr>
<tr><td> AgentID</td><td> 企业应用的id，整型。可在应用的设置页面查看</td></tr></tbody></table>

### **voice消息**
<pre>&lt;xml&gt;
   &lt;ToUserName&gt;&lt;![CDATA[toUser]]&gt;&lt;/ToUserName&gt;
   &lt;FromUserName&gt;&lt;![CDATA[fromUser]]&gt;&lt;/FromUserName&gt;
   &lt;CreateTime&gt;1357290913&lt;/CreateTime&gt;
   &lt;MsgType&gt;&lt;![CDATA[voice]]&gt;&lt;/MsgType&gt;
   &lt;MediaId&gt;&lt;![CDATA[media_id]]&gt;&lt;/MediaId&gt;
   &lt;Format&gt;&lt;![CDATA[Format]]&gt;&lt;/Format&gt;
   &lt;MsgId&gt;1234567890123456&lt;/MsgId&gt;
   &lt;AgentID&gt;1&lt;/AgentID&gt;
&lt;/xml&gt;
</pre>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th><th>说明</th></tr>
<tr><td> ToUserName</td><td> 企业号CorpID</td></tr>
<tr><td> FromUserName</td><td> 员工UserID</td></tr>
<tr><td> CreateTime</td><td> 消息创建时间（整型）</td></tr>
<tr><td> MsgType</td><td> 消息类型，此时固定为：voice</td></tr>
<tr><td> MediaId</td><td> 语音媒体文件id，可以调用获取媒体文件接口拉取数据</td></tr>
<tr><td> Format</td><td> 语音格式，如amr，speex等</td></tr>
<tr><td> MsgId</td><td> 消息id，64位整型</td></tr>
<tr><td> AgentID</td><td> 企业应用的id，整型。可在应用的设置页面查看</td></tr></tbody></table>

### **video消息**
<pre>&lt;xml&gt;
   &lt;ToUserName&gt;&lt;![CDATA[toUser]]&gt;&lt;/ToUserName&gt;
   &lt;FromUserName&gt;&lt;![CDATA[fromUser]]&gt;&lt;/FromUserName&gt;
   &lt;CreateTime&gt;1357290913&lt;/CreateTime&gt;
   &lt;MsgType&gt;&lt;![CDATA[video]]&gt;&lt;/MsgType&gt;
   &lt;MediaId&gt;&lt;![CDATA[media_id]]&gt;&lt;/MediaId&gt;
   &lt;ThumbMediaId&gt;&lt;![CDATA[thumb_media_id]]&gt;&lt;/ThumbMediaId&gt;
   &lt;MsgId&gt;1234567890123456&lt;/MsgId&gt;
   &lt;AgentID&gt;1&lt;/AgentID&gt;
&lt;/xml&gt;
</pre>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th><th>说明</th></tr>
<tr><td> ToUserName</td><td> 企业号CorpID</td></tr>
<tr><td> FromUserName</td><td> 员工UserID</td></tr>
<tr><td> CreateTime</td><td> 消息创建时间（整型）</td></tr>
<tr><td> MsgType</td><td> 消息类型，此时固定为：video</td></tr>
<tr><td> MediaId</td><td> 视频媒体文件id，可以调用获取媒体文件接口拉取数据</td></tr>
<tr><td> ThumbMediaId</td><td> 视频消息缩略图的媒体id，可以调用获取媒体文件接口拉取数据</td></tr>
<tr><td> MsgId</td><td> 消息id，64位整型</td></tr>
<tr><td> AgentID</td><td> 企业应用的id，整型。可在应用的设置页面查看</td></tr></tbody></table>

### **location消息**
<pre>&lt;xml&gt;
   &lt;ToUserName&gt;&lt;![CDATA[toUser]]&gt;&lt;/ToUserName&gt;
   &lt;FromUserName&gt;&lt;![CDATA[fromUser]]&gt;&lt;/FromUserName&gt;
   &lt;CreateTime&gt;1351776360&lt;/CreateTime&gt;
   &lt;MsgType&gt;&lt;![CDATA[location]]&gt;&lt;/MsgType&gt;
   &lt;Location_X&gt;23.134521&lt;/Location_X&gt;
   &lt;Location_Y&gt;113.358803&lt;/Location_Y&gt;
   &lt;Scale&gt;20&lt;/Scale&gt;
   &lt;Label&gt;&lt;![CDATA[位置信息]]&gt;&lt;/Label&gt;
   &lt;MsgId&gt;1234567890123456&lt;/MsgId&gt;
   &lt;AgentID&gt;1&lt;/AgentID&gt;
&lt;/xml&gt;
</pre>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th><th>说明</th></tr>
<tr><td> ToUserName</td><td> 企业号CorpID</td></tr>
<tr><td> FromUserName</td><td> 员工UserID</td></tr>
<tr><td> CreateTime</td><td> 消息创建时间（整型）</td></tr>
<tr><td> MsgType</td><td> 消息类型，此时固定为：location</td></tr>
<tr><td> Location_X</td><td> 地理位置纬度</td></tr>
<tr><td> Location_Y</td><td> 地理位置经度</td></tr>
<tr><td> Scale</td><td> 地图缩放大小</td></tr>
<tr><td> Label</td><td> 地理位置信息</td></tr>
<tr><td> MsgId</td><td> 消息id，64位整型</td></tr>
<tr><td> AgentID</td><td> 企业应用的id，整型。可在应用的设置页面查看</td></tr></tbody></table>

<!--
NewPP limit report
CPU time usage: 0.024 seconds
Real time usage: 0.027 seconds
Preprocessor visited node count: 27/1000000
Preprocessor generated node count: 38/1000000
Post‐expand include size: 0/2097152 bytes
Template argument size: 0/2097152 bytes
Highest expansion depth: 2/40
Expensive parser function count: 0/100
-->
<!-- Saved in parser cache with key db_wiki:pcache:idhash:14-0!*!*!!zh-cn!*!* and timestamp 20150116023645 and revision id 510
 -->
## 接收事件
<div id="toc" class="toc"><div id="toctitle">## 目录</div>
<ul>
<li class="toclevel-1 tocsection-1"><a href="#.E6.8E.A5.E6.94.B6.E4.BA.8B.E4.BB.B6"><span class="tocnumber">1</span> <span class="toctext"><b>接收事件</a>
<ul>
<li class="toclevel-2 tocsection-2"><a href="#.E4.B8.8A.E6.8A.A5.E5.9C.B0.E7.90.86.E4.BD.8D.E7.BD.AE.E4.BA.8B.E4.BB.B6"><span class="tocnumber">1.1</span> <span class="toctext"><b>上报地理位置事件</a></li>
<li class="toclevel-2 tocsection-3"><a href="#.E4.B8.8A.E6.8A.A5.E8.8F.9C.E5.8D.95.E4.BA.8B.E4.BB.B6"><span class="tocnumber">1.2</span> <span class="toctext"><b>上报菜单事件</a>
<ul>
<li class="toclevel-3 tocsection-4"><a href="#.E7.82.B9.E5.87.BB.E8.8F.9C.E5.8D.95.E6.8B.89.E5.8F.96.E6.B6.88.E6.81.AF.E7.9A.84.E4.BA.8B.E4.BB.B6.E6.8E.A8.E9.80.81"><span class="tocnumber">1.2.1</span> <span class="toctext"><b>点击菜单拉取消息的事件推送</a></li>
<li class="toclevel-3 tocsection-5"><a href="#.E7.82.B9.E5.87.BB.E8.8F.9C.E5.8D.95.E8.B7.B3.E8.BD.AC.E9.93.BE.E6.8E.A5.E7.9A.84.E4.BA.8B.E4.BB.B6.E6.8E.A8.E9.80.81"><span class="tocnumber">1.2.2</span> <span class="toctext"><b>点击菜单跳转链接的事件推送</a></li>
<li class="toclevel-3 tocsection-6"><a href="#.E6.89.AB.E7.A0.81.E6.8E.A8.E4.BA.8B.E4.BB.B6.E7.9A.84.E4.BA.8B.E4.BB.B6.E6.8E.A8.E9.80.81"><span class="tocnumber">1.2.3</span> <span class="toctext"><b>扫码推事件的事件推送</a></li>
<li class="toclevel-3 tocsection-7"><a href="#.E6.89.AB.E7.A0.81.E6.8E.A8.E4.BA.8B.E4.BB.B6.E4.B8.94.E5.BC.B9.E5.87.BA.E2.80.9C.E6.B6.88.E6.81.AF.E6.8E.A5.E6.94.B6.E4.B8.AD.E2.80.9D.E6.8F.90.E7.A4.BA.E6.A1.86.E7.9A.84.E4.BA.8B.E4.BB.B6.E6.8E.A8.E9.80.81"><span class="tocnumber">1.2.4</span> <span class="toctext"><b>扫码推事件且弹出“消息接收中”提示框的事件推送</a></li>
<li class="toclevel-3 tocsection-8"><a href="#.E5.BC.B9.E5.87.BA.E7.B3.BB.E7.BB.9F.E6.8B.8D.E7.85.A7.E5.8F.91.E5.9B.BE.E7.9A.84.E4.BA.8B.E4.BB.B6.E6.8E.A8.E9.80.81"><span class="tocnumber">1.2.5</span> <span class="toctext"><b>弹出系统拍照发图的事件推送</a></li>
<li class="toclevel-3 tocsection-9"><a href="#.E5.BC.B9.E5.87.BA.E6.8B.8D.E7.85.A7.E6.88.96.E8.80.85.E7.9B.B8.E5.86.8C.E5.8F.91.E5.9B.BE.E7.9A.84.E4.BA.8B.E4.BB.B6.E6.8E.A8.E9.80.81"><span class="tocnumber">1.2.6</span> <span class="toctext"><b>弹出拍照或者相册发图的事件推送</a></li>
<li class="toclevel-3 tocsection-10"><a href="#.E5.BC.B9.E5.87.BA.E5.BE.AE.E4.BF.A1.E7.9B.B8.E5.86.8C.E5.8F.91.E5.9B.BE.E5.99.A8.E7.9A.84.E4.BA.8B.E4.BB.B6.E6.8E.A8.E9.80.81"><span class="tocnumber">1.2.7</span> <span class="toctext"><b>弹出微信相册发图器的事件推送</a></li>
<li class="toclevel-3 tocsection-11"><a href="#.E5.BC.B9.E5.87.BA.E5.9C.B0.E7.90.86.E4.BD.8D.E7.BD.AE.E9.80.89.E6.8B.A9.E5.99.A8.E7.9A.84.E4.BA.8B.E4.BB.B6.E6.8E.A8.E9.80.81"><span class="tocnumber">1.2.8</span> <span class="toctext"><b>弹出地理位置选择器的事件推送</a></li>
<li class="toclevel-3 tocsection-12"><a href="#.E7.94.A8.E6.88.B7.E8.BF.9B.E5.85.A5.E5.BA.94.E7.94.A8.E7.9A.84.E4.BA.8B.E4.BB.B6.E6.8E.A8.E9.80.81"><span class="tocnumber">1.2.9</span> <span class="toctext"><b>用户进入应用的事件推送</a></li>
</ul></li>
</ul></li>
</ul>
</div>
## **接收事件**
事件是指员工在企业号上的某些操作行为，比如关注、上报地理位置、点击菜单等。（关注事件请参考’关注与取消关注’）。事件会推送到每个应用在管理端设置的URL（以下假设该URL为<a rel="nofollow" class="external free" href="http://api.3dept.com">http://api.3dept.com</a>）。
<ul>
<li>请求说明</li>
</ul>
Http请求方式: POST

<a rel="nofollow" class="external free" href="http://api.3dept.com/?msg_signature=ASDFQWEXZCVAQFASDFASDFSS&amp;timestamp=13500001234&amp;nonce=123412323">http://api.3dept.com/?msg_signature=ASDFQWEXZCVAQFASDFASDFSS&amp;timestamp=13500001234&amp;nonce=123412323</a>

### **上报地理位置事件**
员工同意上报地理位置后，每次在进入应用会话时都会上报一次地理位置，或在进入应用会话后每5秒上报一次地理位置。企业可以在管理端修改应用的以上设置。上报地理位置时，微信会将此事件推送到企业应用在管理端设置的URL（以下假设该URL为http://api.3dept.com）。
<ul>
<li>请求说明</li>
</ul>
Http请求方式: POST

<a rel="nofollow" class="external free" href="http://api.3dept.com/?msg_signature=ASDFQWEXZCVAQFASDFASDFSS&amp;timestamp=13500001234&amp;nonce=123412323">http://api.3dept.com/?msg_signature=ASDFQWEXZCVAQFASDFASDFSS&amp;timestamp=13500001234&amp;nonce=123412323</a>
<ul>
<li>参数说明</li>
</ul>
<pre>&lt;xml&gt;
   &lt;ToUserName&gt;&lt;![CDATA[toUser]]&gt;&lt;/ToUserName&gt;
   &lt;FromUserName&gt;&lt;![CDATA[FromUser]]&gt;&lt;/FromUserName&gt;
   &lt;CreateTime&gt;123456789&lt;/CreateTime&gt;
   &lt;MsgType&gt;&lt;![CDATA[event]]&gt;&lt;/MsgType&gt;
   &lt;Event&gt;&lt;![CDATA[LOCATION]]&gt;&lt;/Event&gt;
   &lt;Latitude&gt;23.104105&lt;/Latitude&gt;
   &lt;Longitude&gt;113.320107&lt;/Longitude&gt;
   &lt;Precision&gt;65.000000&lt;/Precision&gt;
   &lt;AgentID&gt;1&lt;/AgentID&gt;
&lt;/xml&gt;
</pre>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th><th>说明</th></tr>
<tr><td> ToUserName</td><td> 企业号CorpID</td></tr>
<tr><td> FromUserName</td><td> 员工UserID</td></tr>
<tr><td> CreateTime</td><td> 消息创建时间（整型）</td></tr>
<tr><td> MsgType</td><td> 消息类型，此时固定为：event</td></tr>
<tr><td> Event</td><td> 事件类型，此时固定为：LOCATION</td></tr>
<tr><td> Latitude</td><td> 地理位置纬度</td></tr>
<tr><td> Longitude</td><td> 地理位置经度</td></tr>
<tr><td> Precision</td><td> 地理位置精度</td></tr>
<tr><td> AgentID</td><td> 企业应用的id，整型。可在应用的设置页面查看</td></tr></tbody></table>

### **上报菜单事件**
用户点击自定义菜单后，微信会把点击事件推送给开发者，请注意，点击菜单弹出子菜单，不会产生上报。另外，扫码、拍照及地理位置的菜单事件，仅支持微信iPhone5.4.1/Android5.4以上版本，旧版本微信用户点击后将没有回应，开发者也不能正常接收到事件推送。
<h4>**点击菜单拉取消息的事件推送**</h4>
推送XML数据包示例：
<pre>&lt;xml&gt;
&lt;ToUserName&gt;&lt;![CDATA[toUser]]&gt;&lt;/ToUserName&gt;
&lt;FromUserName&gt;&lt;![CDATA[FromUser]]&gt;&lt;/FromUserName&gt;
&lt;CreateTime&gt;123456789&lt;/CreateTime&gt;
&lt;MsgType&gt;&lt;![CDATA[event]]&gt;&lt;/MsgType&gt;
&lt;Event&gt;&lt;![CDATA[CLICK]]&gt;&lt;/Event&gt;
&lt;EventKey&gt;&lt;![CDATA[EVENTKEY]]&gt;&lt;/EventKey&gt;
&lt;AgentID&gt;1&lt;/AgentID&gt;
&lt;/xml&gt;
</pre>
参数说明：
<table border="1" cellspacing="0" cellpadding="4" align="center">
<tbody><tr>
<th style="width:180px">参数</th>
<th style="width:470px">描述</th></tr>
<tr><td> ToUserName</td><td> 微信企业号</td></tr>
<tr><td> FromUserName</td><td> 员工UserID</td></tr>
<tr><td> CreateTime</td><td> 消息创建时间 （整型）</td></tr>
<tr><td> MsgType</td><td> 消息类型，event</td></tr>
<tr><td> Event</td><td> 事件类型，CLICK</td></tr>
<tr><td> EventKey</td><td> 事件KEY值，与自定义菜单接口中KEY值对应</td></tr>
<tr><td> AgentID</td><td> 应用代理ID</td></tr>
</tbody></table>

<h4>**点击菜单跳转链接的事件推送**</h4>
推送XML数据包示例：
<pre>&lt;xml&gt;
&lt;ToUserName&gt;&lt;![CDATA[toUser]]&gt;&lt;/ToUserName&gt;
&lt;FromUserName&gt;&lt;![CDATA[FromUser]]&gt;&lt;/FromUserName&gt;
&lt;CreateTime&gt;123456789&lt;/CreateTime&gt;
&lt;MsgType&gt;&lt;![CDATA[event]]&gt;&lt;/MsgType&gt;
&lt;Event&gt;&lt;![CDATA[VIEW]]&gt;&lt;/Event&gt;
&lt;EventKey&gt;&lt;![CDATA[www.qq.com]]&gt;&lt;/EventKey&gt;
&lt;AgentID&gt;1&lt;/AgentID&gt;
&lt;/xml&gt;
</pre>
参数说明：
<table border="1" cellspacing="0" cellpadding="4" align="center">
<tbody><tr>
<th style="width:180px">参数</th>
<th style="width:470px">描述</th></tr>
<tr><td> ToUserName</td><td> 微信企业号</td></tr>
<tr><td> FromUserName</td><td> 员工UserID</td></tr>
<tr><td> CreateTime</td><td> 消息创建时间 （整型）</td></tr>
<tr><td> MsgType</td><td> 消息类型，event</td></tr>
<tr><td> Event</td><td> 事件类型，VIEW</td></tr>
<tr><td> EventKey</td><td> 事件KEY值，设置的跳转URL</td></tr>
<tr><td> AgentID</td><td> 应用代理ID</td></tr>
</tbody></table>

<h4>**扫码推事件的事件推送**</h4>
推送XML数据包示例：
<pre>&lt;xml&gt;&lt;ToUserName&gt;&lt;![CDATA[toUser]]&gt;&lt;/ToUserName&gt;
&lt;FromUserName&gt;&lt;![CDATA[FromUser]]&gt;&lt;/FromUserName&gt;
&lt;CreateTime&gt;1408090502&lt;/CreateTime&gt;
&lt;MsgType&gt;&lt;![CDATA[event]]&gt;&lt;/MsgType&gt;
&lt;Event&gt;&lt;![CDATA[scancode_push]]&gt;&lt;/Event&gt;
&lt;EventKey&gt;&lt;![CDATA[6]]&gt;&lt;/EventKey&gt;
&lt;ScanCodeInfo&gt;&lt;ScanType&gt;&lt;![CDATA[qrcode]]&gt;&lt;/ScanType&gt;
&lt;ScanResult&gt;&lt;![CDATA<a rel="nofollow" class="external autonumber" href="http://weixin.qq.com/r/5HXdxQ-EKFJXrUum9yD2">[1]</a>]&gt;&lt;/ScanResult&gt;
&lt;/ScanCodeInfo&gt;
&lt;AgentID&gt;1&lt;/AgentID&gt;
&lt;/xml&gt;
</pre>
参数说明：
<table border="1" cellspacing="0" cellpadding="4" align="center">
<tbody><tr>
<th style="width:180px">参数</th>
<th style="width:470px">描述</th></tr>
<tr><td> ToUserName</td><td> 微信企业号</td></tr>
<tr><td> FromUserName</td><td> 员工UserID</td></tr>
<tr><td> CreateTime</td><td> 消息创建时间（整型）</td></tr>
<tr><td> MsgType</td><td> 消息类型，event</td></tr>
<tr><td> Event</td><td> 事件类型，scancode_push</td></tr>
<tr><td> EventKey</td><td> 事件KEY值，由开发者在创建菜单时设定</td></tr>
<tr><td> ScanCodeInfo</td><td> 扫描信息</td></tr>
<tr><td> ScanType</td><td> 扫描类型，一般是qrcode</td></tr>
<tr><td> ScanResult</td><td> 扫描结果，即二维码对应的字符串信息</td></tr>
<tr><td> AgentID</td><td> 应用代理ID</td></tr>
</tbody></table>

<h4>**扫码推事件且弹出“消息接收中”提示框的事件推送**</h4>
推送XML数据包示例：
<pre>&lt;xml&gt;&lt;ToUserName&gt;&lt;![CDATA[toUser]]&gt;&lt;/ToUserName&gt;
&lt;FromUserName&gt;&lt;![CDATA[FromUser]]&gt;&lt;/FromUserName&gt;
&lt;CreateTime&gt;1408090606&lt;/CreateTime&gt;
&lt;MsgType&gt;&lt;![CDATA[event]]&gt;&lt;/MsgType&gt;
&lt;Event&gt;&lt;![CDATA[scancode_waitmsg]]&gt;&lt;/Event&gt;
&lt;EventKey&gt;&lt;![CDATA[6]]&gt;&lt;/EventKey&gt;
&lt;ScanCodeInfo&gt;&lt;ScanType&gt;&lt;![CDATA[qrcode]]&gt;&lt;/ScanType&gt;
&lt;ScanResult&gt;&lt;![CDATA<a rel="nofollow" class="external autonumber" href="http://weixin.qq.com/r/5HXdxQ-EKFJXrUum9yD2">[2]</a>]&gt;&lt;/ScanResult&gt;
&lt;/ScanCodeInfo&gt;
&lt;AgentID&gt;1&lt;/AgentID&gt;
&lt;/xml&gt;
</pre>
参数说明：
<table border="1" cellspacing="0" cellpadding="4" align="center">
<tbody><tr>
<th style="width:180px">参数</th>
<th style="width:470px">描述</th></tr>
<tr><td> ToUserName</td><td> 微信企业号</td></tr>
<tr><td> FromUserName</td><td> 员工UserID</td></tr>
<tr><td> CreateTime</td><td> 消息创建时间 （整型）</td></tr>
<tr><td> MsgType</td><td> 消息类型，event</td></tr>
<tr><td> Event</td><td> 事件类型，scancode_waitmsg</td></tr>
<tr><td> EventKey</td><td> 事件KEY值，由开发者在创建菜单时设定</td></tr>
<tr><td> ScanCodeInfo</td><td> 扫描信息</td></tr>
<tr><td> ScanType</td><td> 扫描类型，一般是qrcode</td></tr>
<tr><td> ScanResult</td><td> 扫描结果，即二维码对应的字符串信息</td></tr>
<tr><td> AgentID</td><td> 应用代理ID</td></tr>
</tbody></table>

<h4>**弹出系统拍照发图的事件推送**</h4>
推送XML数据包示例：
<pre>&lt;xml&gt;&lt;ToUserName&gt;&lt;![CDATA[toUser]]&gt;&lt;/ToUserName&gt;
&lt;FromUserName&gt;&lt;![CDATA[FromUser]]&gt;&lt;/FromUserName&gt;
&lt;CreateTime&gt;1408090651&lt;/CreateTime&gt;
&lt;MsgType&gt;&lt;![CDATA[event]]&gt;&lt;/MsgType&gt;
&lt;Event&gt;&lt;![CDATA[pic_sysphoto]]&gt;&lt;/Event&gt;
&lt;EventKey&gt;&lt;![CDATA[6]]&gt;&lt;/EventKey&gt;
&lt;SendPicsInfo&gt;&lt;Count&gt;1&lt;/Count&gt;
&lt;PicList&gt;&lt;item&gt;&lt;PicMd5Sum&gt;&lt;![CDATA[1b5f7c23b5bf75682a53e7b6d163e185]]&gt;&lt;/PicMd5Sum&gt;
&lt;/item&gt;
&lt;/PicList&gt;
&lt;/SendPicsInfo&gt;
&lt;AgentID&gt;1&lt;/AgentID&gt;
&lt;/xml&gt;
</pre>
参数说明：
<table border="1" cellspacing="0" cellpadding="4" align="center">
<tbody><tr>
<th style="width:180px">参数</th>
<th style="width:470px">描述</th></tr>
<tr><td> ToUserName</td><td> 微信企业号</td></tr>
<tr><td> FromUserName</td><td> 员工UserID</td></tr>
<tr><td> CreateTime</td><td> 消息创建时间 （整型）</td></tr>
<tr><td> MsgType</td><td> 消息类型，event</td></tr>
<tr><td> Event</td><td> 事件类型，pic_sysphoto</td></tr>
<tr><td> EventKey</td><td> 事件KEY值，由开发者在创建菜单时设定</td></tr>
<tr><td> SendPicsInfo</td><td> 发送的图片信息</td></tr>
<tr><td> Count</td><td> 发送的图片数量</td></tr>
<tr><td> PicList</td><td> 图片列表</td></tr>
<tr><td> PicMd5Sum</td><td> 图片的MD5值，开发者若需要，可用于验证接收到图片</td></tr>
<tr><td> AgentID</td><td> 应用代理ID</td></tr>
</tbody></table>

<h4>**弹出拍照或者相册发图的事件推送**</h4>
推送XML数据包示例：
<pre>&lt;xml&gt;&lt;ToUserName&gt;&lt;![CDATA[toUser]]&gt;&lt;/ToUserName&gt;
&lt;FromUserName&gt;&lt;![CDATA[FromUser]]&gt;&lt;/FromUserName&gt;
&lt;CreateTime&gt;1408090816&lt;/CreateTime&gt;
&lt;MsgType&gt;&lt;![CDATA[event]]&gt;&lt;/MsgType&gt;
&lt;Event&gt;&lt;![CDATA[pic_photo_or_album]]&gt;&lt;/Event&gt;
&lt;EventKey&gt;&lt;![CDATA[6]]&gt;&lt;/EventKey&gt;
&lt;SendPicsInfo&gt;&lt;Count&gt;1&lt;/Count&gt;
&lt;PicList&gt;&lt;item&gt;&lt;PicMd5Sum&gt;&lt;![CDATA[5a75aaca956d97be686719218f275c6b]]&gt;&lt;/PicMd5Sum&gt;
&lt;/item&gt;
&lt;/PicList&gt;
&lt;/SendPicsInfo&gt;
&lt;AgentID&gt;1&lt;/AgentID&gt;
&lt;/xml&gt;
</pre>
参数说明：
<table border="1" cellspacing="0" cellpadding="4" align="center">
<tbody><tr>
<th style="width:180px">参数</th>
<th style="width:470px">描述</th></tr>
<tr><td> ToUserName</td><td> 微信企业号</td></tr>
<tr><td> FromUserName</td><td> 员工UserID</td></tr>
<tr><td> CreateTime</td><td> 消息创建时间 （整型）</td></tr>
<tr><td> MsgType</td><td> 消息类型，event</td></tr>
<tr><td> Event</td><td> 事件类型，pic_photo_or_album</td></tr>
<tr><td> EventKey</td><td> 事件KEY值，由开发者在创建菜单时设定</td></tr>
<tr><td> SendPicsInfo</td><td> 发送的图片信息</td></tr>
<tr><td> Count</td><td> 发送的图片数量</td></tr>
<tr><td> PicList</td><td> 图片列表</td></tr>
<tr><td> PicMd5Sum</td><td> 图片的MD5值，开发者若需要，可用于验证接收到图片</td></tr>
<tr><td> AgentID</td><td> 应用代理ID</td></tr>
</tbody></table>

<h4>**弹出微信相册发图器的事件推送**</h4>
推送XML数据包示例：
<pre>&lt;xml&gt;&lt;ToUserName&gt;&lt;![CDATA[toUser]]&gt;&lt;/ToUserName&gt;
&lt;FromUserName&gt;&lt;![CDATA[FromUser]]&gt;&lt;/FromUserName&gt;
&lt;CreateTime&gt;1408090816&lt;/CreateTime&gt;
&lt;MsgType&gt;&lt;![CDATA[event]]&gt;&lt;/MsgType&gt;
&lt;Event&gt;&lt;![CDATA[pic_weixin]]&gt;&lt;/Event&gt;
&lt;EventKey&gt;&lt;![CDATA[6]]&gt;&lt;/EventKey&gt;
&lt;SendPicsInfo&gt;&lt;Count&gt;1&lt;/Count&gt;
&lt;PicList&gt;&lt;item&gt;&lt;PicMd5Sum&gt;&lt;![CDATA[5a75aaca956d97be686719218f275c6b]]&gt;&lt;/PicMd5Sum&gt;
&lt;/item&gt;
&lt;/PicList&gt;
&lt;/SendPicsInfo&gt;
&lt;AgentID&gt;1&lt;/AgentID&gt;
&lt;/xml&gt;
</pre>
参数说明：
<table border="1" cellspacing="0" cellpadding="4" align="center">
<tbody><tr>
<th style="width:180px">参数</th>
<th style="width:470px">描述</th></tr>
<tr><td> ToUserName</td><td> 微信企业号</td></tr>
<tr><td> FromUserName</td><td> 员工UserID</td></tr>
<tr><td> CreateTime</td><td> 消息创建时间 （整型）</td></tr>
<tr><td> MsgType</td><td> 消息类型，event</td></tr>
<tr><td> Event</td><td> 事件类型，pic_weixin</td></tr>
<tr><td> EventKey</td><td> 事件KEY值，由开发者在创建菜单时设定</td></tr>
<tr><td> SendPicsInfo</td><td> 发送的图片信息</td></tr>
<tr><td> Count</td><td> 发送的图片数量</td></tr>
<tr><td> PicList</td><td> 图片列表</td></tr>
<tr><td> PicMd5Sum</td><td> 图片的MD5值，开发者若需要，可用于验证接收到图片</td></tr>
<tr><td> AgentID</td><td> 应用代理ID</td></tr>
</tbody></table>

<h4>**弹出地理位置选择器的事件推送**</h4>
推送XML数据包示例：
<pre>&lt;xml&gt;&lt;ToUserName&gt;&lt;![CDATA[toUser]]&gt;&lt;/ToUserName&gt;
&lt;FromUserName&gt;&lt;![CDATA[FromUser]]&gt;&lt;/FromUserName&gt;
&lt;CreateTime&gt;1408091189&lt;/CreateTime&gt;
&lt;MsgType&gt;&lt;![CDATA[event]]&gt;&lt;/MsgType&gt;
&lt;Event&gt;&lt;![CDATA[location_select]]&gt;&lt;/Event&gt;
&lt;EventKey&gt;&lt;![CDATA[6]]&gt;&lt;/EventKey&gt;
&lt;SendLocationInfo&gt;&lt;Location_X&gt;&lt;![CDATA[23]]&gt;&lt;/Location_X&gt;
&lt;Location_Y&gt;&lt;![CDATA[113]]&gt;&lt;/Location_Y&gt;
&lt;Scale&gt;&lt;![CDATA[15]]&gt;&lt;/Scale&gt;
&lt;Label&gt;&lt;![CDATA[ 广州市海珠区客村艺苑路 106号]]&gt;&lt;/Label&gt;
&lt;Poiname&gt;&lt;![CDATA[]]&gt;&lt;/Poiname&gt;
&lt;/SendLocationInfo&gt;
&lt;AgentID&gt;1&lt;/AgentID&gt;
&lt;/xml&gt;
</pre>
参数说明：
<table border="1" cellspacing="0" cellpadding="4" align="center">
<tbody><tr>
<th style="width:180px">参数</th>
<th style="width:470px">描述</th></tr>
<tr><td> ToUserName</td><td> 微信企业号</td></tr>
<tr><td> FromUserName</td><td> 员工UserID</td></tr>
<tr><td> CreateTime</td><td> 消息创建时间 （整型）</td></tr>
<tr><td> MsgType</td><td> 消息类型，event</td></tr>
<tr><td> Event</td><td> 事件类型，location_select</td></tr>
<tr><td> EventKey</td><td> 事件KEY值，由开发者在创建菜单时设定</td></tr>
<tr><td> SendLocationInfo</td><td> 发送的位置信息</td></tr>
<tr><td> Location_X</td><td> X坐标信息</td></tr>
<tr><td> Location_Y</td><td> Y坐标信息</td></tr>
<tr><td> Scale</td><td> 精度，可理解为精度或者比例尺、越精细的话 scale越高</td></tr>
<tr><td> Label</td><td> 地理位置的字符串信息</td></tr>
<tr><td> Poiname</td><td> 朋友圈POI的名字，可能为空</td></tr>
<tr><td> AgentID</td><td> 应用代理ID</td></tr>
</tbody></table>

<h4>**用户进入应用的事件推送**</h4>
本事件只有在应用的回调模式中打开上报开关时上报
推送XML数据包示例：
<pre>&lt;xml&gt;&lt;ToUserName&gt;&lt;![CDATA[toUser]]&gt;&lt;/ToUserName&gt;
&lt;FromUserName&gt;&lt;![CDATA[FromUser]]&gt;&lt;/FromUserName&gt;
&lt;CreateTime&gt;1408091189&lt;/CreateTime&gt;
&lt;MsgType&gt;&lt;![CDATA[event]]&gt;&lt;/MsgType&gt;
&lt;Event&gt;&lt;![CDATA[enter_agent]]&gt;&lt;/Event&gt;
&lt;EventKey&gt;&lt;![CDATA[]]&gt;&lt;/EventKey&gt;
&lt;AgentID&gt;1&lt;/AgentID&gt;
&lt;/xml&gt;
</pre>
参数说明：
<table border="1" cellspacing="0" cellpadding="4" align="center">
<tbody><tr>
<th style="width:180px">参数</th>
<th style="width:470px">描述</th></tr>
<tr><td> ToUserName</td><td> 微信企业号</td></tr>
<tr><td> FromUserName</td><td> 员工UserID</td></tr>
<tr><td> CreateTime</td><td> 消息创建时间 （整型）</td></tr>
<tr><td> MsgType</td><td> 消息类型，event</td></tr>
<tr><td> Event</td><td> 事件类型，enter_agent</td></tr>
<tr><td> EventKey</td><td> 事件KEY值，此事件该值为空</td></tr>
</tbody></table>
<!--
NewPP limit report
CPU time usage: 0.052 seconds
Real time usage: 0.059 seconds
Preprocessor visited node count: 47/1000000
Preprocessor generated node count: 54/1000000
Post‐expand include size: 0/2097152 bytes
Template argument size: 0/2097152 bytes
Highest expansion depth: 2/40
Expensive parser function count: 0/100
-->
<!-- Saved in parser cache with key db_wiki:pcache:idhash:15-0!*!*!!zh-cn!*!* and timestamp 20150115203008 and revision id 613
 -->
## 被动响应信息
<div id="toc" class="toc"><div id="toctitle">## 目录</div>
<ul>
<li class="toclevel-1 tocsection-1"><a href="#.E8.A2.AB.E5.8A.A8.E5.93.8D.E5.BA.94.E6.B6.88.E6.81.AF"><span class="tocnumber">1</span> <span class="toctext"><b>被动响应消息</a>
<ul>
<li class="toclevel-2 tocsection-2"><a href="#text.E6.B6.88.E6.81.AF"><span class="tocnumber">1.1</span> <span class="toctext"><b>text消息</a></li>
<li class="toclevel-2 tocsection-3"><a href="#image.E6.B6.88.E6.81.AF"><span class="tocnumber">1.2</span> <span class="toctext"><b>image消息</a></li>
<li class="toclevel-2 tocsection-4"><a href="#voice.E6.B6.88.E6.81.AF"><span class="tocnumber">1.3</span> <span class="toctext"><b>voice消息</a></li>
<li class="toclevel-2 tocsection-5"><a href="#video.E6.B6.88.E6.81.AF"><span class="tocnumber">1.4</span> <span class="toctext"><b>video消息</a></li>
<li class="toclevel-2 tocsection-6"><a href="#news.E6.B6.88.E6.81.AF"><span class="tocnumber">1.5</span> <span class="toctext"><b>news消息</a></li>
</ul></li>
</ul>
</div>
## **被动响应消息**
**企业响应的消息同样应该经过加密，并带上msg_signature、timestamp、nonce及密文**，其中timestamp、nonce由企业指定，msg_signature、密文经特定算法生成，具体算法参见附录。

以下是标准的回包：
<pre>&lt;xml&gt;
   &lt;Encrypt&gt;&lt;![CDATA[msg_encrypt]]&gt;&lt;/Encrypt&gt;
   &lt;MsgSignature&gt;&lt;![CDATA[msg_signature]]&gt;&lt;/MsgSignature&gt;
   &lt;TimeStamp&gt;timestamp&lt;/TimeStamp&gt;
   &lt;Nonce&gt;&lt;![CDATA[nonce]]&gt;&lt;/Nonce&gt;
&lt;/xml&gt;
</pre>
以下是各类型消息的明文XML结构：

### **text消息**
<pre>&lt;xml&gt;
   &lt;ToUserName&gt;&lt;![CDATA[toUser]]&gt;&lt;/ToUserName&gt;
   &lt;FromUserName&gt;&lt;![CDATA[fromUser]]&gt;&lt;/FromUserName&gt;
   &lt;CreateTime&gt;1348831860&lt;/CreateTime&gt;
   &lt;MsgType&gt;&lt;![CDATA[text]]&gt;&lt;/MsgType&gt;
   &lt;Content&gt;&lt;![CDATA[this is a test]]&gt;&lt;/Content&gt;
&lt;/xml&gt;
</pre>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th><th>说明</th></tr>
<tr><td> ToUserName</td><td> 员工UserID</td></tr>
<tr><td> FromUserName</td><td> 企业号CorpID</td></tr>
<tr><td> CreateTime</td><td> 消息创建时间（整型）</td></tr>
<tr><td> MsgType</td><td> 消息类型，此时固定为：text</td></tr>
<tr><td> Content</td><td> 文本消息内容</td></tr></tbody></table>

### **image消息**
<pre>&lt;xml&gt;
   &lt;ToUserName&gt;&lt;![CDATA[toUser]]&gt;&lt;/ToUserName&gt;
   &lt;FromUserName&gt;&lt;![CDATA[fromUser]]&gt;&lt;/FromUserName&gt;
   &lt;CreateTime&gt;1348831860&lt;/CreateTime&gt;
   &lt;MsgType&gt;&lt;![CDATA[image]]&gt;&lt;/MsgType&gt;
   &lt;Image&gt;
       &lt;MediaId&gt;&lt;![CDATA[media_id]]&gt;&lt;/MediaId&gt;
   &lt;/Image&gt;
&lt;/xml&gt;
</pre>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th><th>说明</th></tr>
<tr><td> ToUserName</td><td> 员工UserID</td></tr>
<tr><td> FromUserName</td><td> 企业号CorpID</td></tr>
<tr><td> CreateTime</td><td> 消息创建时间（整型）</td></tr>
<tr><td> MsgType</td><td> 消息类型，此时固定为：image</td></tr>
<tr><td> MediaId</td><td> 图片文件id，可以调用上传媒体文件接口获取</td></tr></tbody></table>

### **voice消息**
<pre>&lt;xml&gt;
   &lt;ToUserName&gt;&lt;![CDATA[toUser]]&gt;&lt;/ToUserName&gt;
   &lt;FromUserName&gt;&lt;![CDATA[fromUser]]&gt;&lt;/FromUserName&gt;
   &lt;CreateTime&gt;1357290913&lt;/CreateTime&gt;
   &lt;MsgType&gt;&lt;![CDATA[voice]]&gt;&lt;/MsgType&gt;
   &lt;Voice&gt;
       &lt;MediaId&gt;&lt;![CDATA[media_id]]&gt;&lt;/MediaId&gt;
   &lt;/Voice&gt;
&lt;/xml&gt;
</pre>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th><th>说明</th></tr>
<tr><td> ToUserName</td><td> 员工UserID</td></tr>
<tr><td> FromUserName</td><td> 企业号CorpID</td></tr>
<tr><td> CreateTime</td><td> 消息创建时间（整型）</td></tr>
<tr><td> MsgType</td><td> 消息类型，此时固定为：voice</td></tr>
<tr><td> MediaId</td><td> 语音文件id，可以调用上传媒体文件接口获取</td></tr></tbody></table>

### **video消息**
<pre>&lt;xml&gt;
   &lt;ToUserName&gt;&lt;![CDATA[toUser]]&gt;&lt;/ToUserName&gt;
   &lt;FromUserName&gt;&lt;![CDATA[fromUser]]&gt;&lt;/FromUserName&gt;
   &lt;CreateTime&gt;1357290913&lt;/CreateTime&gt;
   &lt;MsgType&gt;&lt;![CDATA[video]]&gt;&lt;/MsgType&gt;
   &lt;Video&gt;
       &lt;MediaId&gt;&lt;![CDATA[media_id]]&gt;&lt;/MediaId&gt;
       &lt;Title&gt;&lt;![CDATA[title]]&gt;&lt;/Title&gt;
       &lt;Description&gt;&lt;![CDATA[description]]&gt;&lt;/Description&gt;
   &lt;/Video&gt;
&lt;/xml&gt;
</pre>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th><th>说明</th></tr>
<tr><td> ToUserName</td><td> 员工UserID</td></tr>
<tr><td> FromUserName</td><td> 企业号CorpID</td></tr>
<tr><td> CreateTime</td><td> 消息创建时间（整型）</td></tr>
<tr><td> MsgType</td><td> 消息类型，此时固定为：video</td></tr>
<tr><td> MediaId</td><td> 视频文件id，可以调用上传媒体文件接口获取</td></tr>
<tr><td> Title</td><td> 视频消息的标题</td></tr>
<tr><td> Description</td><td> 视频消息的描述</td></tr></tbody></table>

### **news消息**
<pre>&lt;xml&gt;
   &lt;ToUserName&gt;&lt;![CDATA[toUser]]&gt;&lt;/ToUserName&gt;
   &lt;FromUserName&gt;&lt;![CDATA[fromUser]]&gt;&lt;/FromUserName&gt;
   &lt;CreateTime&gt;12345678&lt;/CreateTime&gt;
   &lt;MsgType&gt;&lt;![CDATA[news]]&gt;&lt;/MsgType&gt;
   &lt;ArticleCount&gt;2&lt;/ArticleCount&gt;
   &lt;Articles&gt;
       &lt;item&gt;
           &lt;Title&gt;&lt;![CDATA[title1]]&gt;&lt;/Title&gt;
           &lt;Description&gt;&lt;![CDATA[description1]]&gt;&lt;/Description&gt;
           &lt;PicUrl&gt;&lt;![CDATA[picurl]]&gt;&lt;/PicUrl&gt;
           &lt;Url&gt;&lt;![CDATA[url]]&gt;&lt;/Url&gt;
       &lt;/item&gt;
       &lt;item&gt;
           &lt;Title&gt;&lt;![CDATA[title]]&gt;&lt;/Title&gt;
           &lt;Description&gt;&lt;![CDATA[description]]&gt;&lt;/Description&gt;
           &lt;PicUrl&gt;&lt;![CDATA[picurl]]&gt;&lt;/PicUrl&gt;
           &lt;Url&gt;&lt;![CDATA[url]]&gt;&lt;/Url&gt;
       &lt;/item&gt;
   &lt;/Articles&gt;
&lt;/xml&gt;
</pre>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th><th>说明</th></tr>
<tr><td> ToUserName</td><td> 员工UserID</td></tr>
<tr><td> FromUserName</td><td> 企业号CorpID</td></tr>
<tr><td> CreateTime</td><td> 消息创建时间（整型）</td></tr>
<tr><td> MsgType</td><td> 消息类型，此时固定为：news</td></tr>
<tr><td> ArticleCount</td><td> 图文条数，默认第一条为大图。图文数不能超过10，否则将会无响应</td></tr>
<tr><td> Title</td><td> 图文消息标题</td></tr>
<tr><td> Description</td><td> 图文消息描述</td></tr>
<tr><td> PicUrl</td><td> 图片链接，支持JPG、PNG格式，较好的效果为大图360*200，小图200*200</td></tr>
<tr><td> Url</td><td> 点击图文消息跳转链接</td></tr></tbody></table>

<!--
NewPP limit report
CPU time usage: 0.024 seconds
Real time usage: 0.027 seconds
Preprocessor visited node count: 28/1000000
Preprocessor generated node count: 40/1000000
Post‐expand include size: 0/2097152 bytes
Template argument size: 0/2097152 bytes
Highest expansion depth: 2/40
Expensive parser function count: 0/100
-->
<!-- Saved in parser cache with key db_wiki:pcache:idhash:16-0!*!*!!zh-cn!*!* and timestamp 20150116061005 and revision id 512
 -->
# 发送消息
企业可以主动发消息给员工，**消息量不受限制**。

调用接口时，使用Https协议、JSON数据包格式，数据包不需做加密处理。

目前支持文本、图片、语音、视频、文件、图文等消息类型。除了news类型，其它类型的消息可在发送时加上保密选项，保密消息会被打上水印，并且只有接收者才能阅读。

<!--
NewPP limit report
CPU time usage: 0.000 seconds
Real time usage: 0.003 seconds
Preprocessor visited node count: 1/1000000
Preprocessor generated node count: 4/1000000
Post‐expand include size: 0/2097152 bytes
Template argument size: 0/2097152 bytes
Highest expansion depth: 1/40
Expensive parser function count: 0/100
-->
<!-- Saved in parser cache with key db_wiki:pcache:idhash:33-0!*!*!*!*!*!* and timestamp 20150115102709 and revision id 216
 -->
## 发送接口说明
## **发送接口说明**
<ul>
<li>请求说明</li>
</ul>
Https请求方式: POST

<a rel="nofollow" class="external free" href="https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=ACCESS_TOKEN">https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=ACCESS_TOKEN</a>
<ul>
<li>参数说明</li>
</ul>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:120px">参数</th><th>必须</th><th>说明</th></tr>
<tr><td> access_token</td><td> 是</td><td> 调用接口凭证</td></tr></tbody></table>
<ul>
<li>权限要求</li>
</ul>
需要管理员对应用有使用权限，对收件人touser、toparty、totag有查看权限，否则本次调用失败。
<ul>
<li>返回结果</li>
</ul>
如果对应用或收件人、部门、标签任何一个无权限，则本次发送失败；如果收件人、部门或标签不存在，发送仍然执行，但返回无效的部分。
<pre>{
   "errcode": 0,
   "errmsg": "ok",
   "invaliduser": "UserID1",
   "invalidparty":"PartyID1",
   "invalidtag":"TagID1"
}
</pre>

<!--
NewPP limit report
CPU time usage: 0.004 seconds
Real time usage: 0.006 seconds
Preprocessor visited node count: 2/1000000
Preprocessor generated node count: 8/1000000
Post‐expand include size: 0/2097152 bytes
Template argument size: 0/2097152 bytes
Highest expansion depth: 2/40
Expensive parser function count: 0/100
-->
<!-- Saved in parser cache with key db_wiki:pcache:idhash:17-0!*!*!*!*!*!* and timestamp 20150115160657 and revision id 343
 -->
## 消息类型及数据结构
<div id="toc" class="toc"><div id="toctitle">## 目录</div>
<ul>
<li class="toclevel-1 tocsection-1"><a href="#text.E6.B6.88.E6.81.AF"><span class="tocnumber">1</span> <span class="toctext"><b>text消息</a></li>
<li class="toclevel-1 tocsection-2"><a href="#image.E6.B6.88.E6.81.AF"><span class="tocnumber">2</span> <span class="toctext"><b>image消息</a></li>
<li class="toclevel-1 tocsection-3"><a href="#voice.E6.B6.88.E6.81.AF"><span class="tocnumber">3</span> <span class="toctext"><b>voice消息</a></li>
<li class="toclevel-1 tocsection-4"><a href="#video.E6.B6.88.E6.81.AF"><span class="tocnumber">4</span> <span class="toctext"><b>video消息</a></li>
<li class="toclevel-1 tocsection-5"><a href="#file.E6.B6.88.E6.81.AF"><span class="tocnumber">5</span> <span class="toctext"><b>file消息</a></li>
<li class="toclevel-1 tocsection-6"><a href="#news.E6.B6.88.E6.81.AF"><span class="tocnumber">6</span> <span class="toctext"><b>news消息</a></li>
<li class="toclevel-1 tocsection-7"><a href="#mpnews.E6.B6.88.E6.81.AF"><span class="tocnumber">7</span> <span class="toctext"><b>mpnews消息</a></li>
</ul>
</div>
### **text消息**
<pre>{
   "touser": "UserID1|UserID2|UserID3",
   "toparty": " PartyID1 | PartyID2 ",
   "totag": " TagID1 | TagID2 ",
   "msgtype": "text",
   "agentid": "1",
   "text": {
       "content": "Holiday Request For Pony(<a rel="nofollow" class="external free" href="http://xxxxx">http://xxxxx</a>)"
   },
   "safe":"0"
}
</pre>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:120px">参数</th>
<th style="width:120px">必须</th><th>说明</th></tr>
<tr><td> touser</td><td> 否</td><td> UserID列表（消息接收者，多个接收者用‘|’分隔）。特殊情况：指定为@all，则向关注该企业应用的全部成员发送</td></tr>
<tr><td> toparty</td><td> 否</td><td> PartyID列表，多个接受者用‘|’分隔。当touser为@all时忽略本参数</td></tr>
<tr><td> totag</td><td> 否</td><td> TagID列表，多个接受者用‘|’分隔。当touser为@all时忽略本参数</td></tr>
<tr><td> msgtype</td><td> 是</td><td> 消息类型，此时固定为：text</td></tr>
<tr><td> agentid</td><td> 是</td><td> 企业应用的id，整型。可在应用的设置页面查看</td></tr>
<tr><td> content</td><td> 是</td><td> 消息内容</td></tr>
<tr><td> safe</td><td> 否</td><td> 表示是否是保密消息，0表示否，1表示是，默认0</td></tr></tbody></table>

### **image消息**
<pre>{
   "touser": "UserID1|UserID2|UserID3",
   "toparty": " PartyID1 | PartyID2 ",
   "msgtype": "image",
   "agentid": "1",
   "image": {
       "media_id": "MEDIA_ID"
   },
   "safe":"0"
}
</pre>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:120px">参数</th>
<th style="width:120px">必须</th><th>说明</th></tr>
<tr><td> touser</td><td> 否</td><td> UserID列表（消息接收者，多个接收者用‘|’分隔）。特殊情况：指定为@all，则向关注该企业应用的全部成员发送</td></tr>
<tr><td> toparty</td><td> 否</td><td> PartyID列表，多个接受者用‘|’分隔。当touser为@all时忽略本参数</td></tr>
<tr><td> totag</td><td> 否</td><td> TagID列表，多个接受者用‘|’分隔。当touser为@all时忽略本参数</td></tr>
<tr><td> msgtype</td><td> 是</td><td> 消息类型，此时固定为：image</td></tr>
<tr><td> agentid</td><td> 是</td><td> 企业应用的id，整型。可在应用的设置页面查看</td></tr>
<tr><td> media_id</td><td> 是</td><td> 图片媒体文件id，可以调用上传媒体文件接口获取</td></tr>
<tr><td> safe</td><td> 否</td><td> 表示是否是保密消息，0表示否，1表示是，默认0</td></tr></tbody></table>

### **voice消息**
<pre>{
   "touser": "UserID1|UserID2|UserID3",
   "toparty": " PartyID1 | PartyID2 ",
   "totag": " TagID1 | TagID2 ",
   "msgtype": "voice",
   "agentid": "1",
   "voice": {
       "media_id": "MEDIA_ID"
   },
   "safe":"0"
}
</pre>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:120px">参数</th>
<th style="width:120px">必须</th><th>说明</th></tr>
<tr><td> touser</td><td> 否</td><td> UserID列表（消息接收者，多个接收者用‘|’分隔）。特殊情况：指定为@all，则向关注该企业应用的全部成员发送</td></tr>
<tr><td> toparty</td><td> 否</td><td> PartyID列表，多个接受者用‘|’分隔。当touser为@all时忽略本参数</td></tr>
<tr><td> totag</td><td> 否</td><td> TagID列表，多个接受者用‘|’分隔。当touser为@all时忽略本参数</td></tr>
<tr><td> msgtype</td><td> 是</td><td> 消息类型，此时固定为：voice</td></tr>
<tr><td> agentid</td><td> 是</td><td> 企业应用的id，整型。可在应用的设置页面查看</td></tr>
<tr><td> media_id</td><td> 是</td><td> 语音文件id，可以调用上传媒体文件接口获取</td></tr>
<tr><td> safe</td><td> 否</td><td> 表示是否是保密消息，0表示否，1表示是，默认0</td></tr></tbody></table>

### **video消息**
<pre>{
   "touser": "UserID1|UserID2|UserID3",
   "toparty": " PartyID1 | PartyID2 ",
   "totag": " TagID1 | TagID2 ",
   "msgtype": "video",
   "agentid": "1",
   "video": {
       "media_id": "MEDIA_ID",
       "title": "Title",
       "description": "Description"
   },
   "safe":"0"
}
</pre>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:120px">参数</th>
<th style="width:120px">必须</th><th>说明</th></tr>
<tr><td> touser</td><td> 否</td><td> UserID列表（消息接收者，多个接收者用‘|’分隔）。特殊情况：指定为@all，则向关注该企业应用的全部成员发送</td></tr>
<tr><td> toparty</td><td> 否</td><td> PartyID列表，多个接受者用‘|’分隔。当touser为@all时忽略本参数</td></tr>
<tr><td> totag</td><td> 否</td><td> TagID列表，多个接受者用‘|’分隔。当touser为@all时忽略本参数</td></tr>
<tr><td> msgtype</td><td> 是</td><td> 消息类型，此时固定为：video</td></tr>
<tr><td> agentid</td><td> 是</td><td> 企业应用的id，整型。可在应用的设置页面查看</td></tr>
<tr><td> media_id</td><td> 是</td><td> 视频媒体文件id，可以调用上传媒体文件接口获取</td></tr>
<tr><td> title</td><td> 否</td><td> 视频消息的标题</td></tr>
<tr><td> description</td><td> 否</td><td> 视频消息的描述</td></tr>
<tr><td> safe</td><td> 否</td><td> 表示是否是保密消息，0表示否，1表示是，默认0</td></tr></tbody></table>

### **file消息**
<pre>{
   "touser": "UserID1|UserID2|UserID3",
   "toparty": " PartyID1 | PartyID2 ",
   "totag": " TagID1 | TagID2 ",
   "msgtype": "file",
   "agentid": "1",
   "file": {
       "media_id": "MEDIA_ID"
   },
   "safe":"0"
}
</pre>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:120px">参数</th>
<th style="width:120px">必须</th><th>说明</th></tr>
<tr><td> touser</td><td> 否</td><td> UserID列表（消息接收者，多个接收者用‘|’分隔）。特殊情况：指定为@all，则向关注该企业应用的全部成员发送</td></tr>
<tr><td> toparty</td><td> 否</td><td> PartyID列表，多个接受者用‘|’分隔。当touser为@all时忽略本参数</td></tr>
<tr><td> totag</td><td> 否</td><td> TagID列表，多个接受者用‘|’分隔。当touser为@all时忽略本参数</td></tr>
<tr><td> msgtype</td><td> 是</td><td> 消息类型，此时固定为：file</td></tr>
<tr><td> agentid</td><td> 是</td><td> 企业应用的id，整型。可在应用的设置页面查看</td></tr>
<tr><td> media_id</td><td> 是</td><td> 媒体文件id，可以调用上传媒体文件接口获取</td></tr>
<tr><td> safe</td><td> 否</td><td> 表示是否是保密消息，0表示否，1表示是，默认0</td></tr></tbody></table>

### **news消息**
<pre>{
   "touser": "UserID1|UserID2|UserID3",
   "toparty": " PartyID1 | PartyID2 ",
   "totag": " TagID1 | TagID2 ",
   "msgtype": "news",
   "agentid": "1",
   "news": {
       "articles":[
           {
               "title": "Title",
               "description": "Description",
               "url": "URL",
               "picurl": "PIC_URL"
           },
           {
               "title": "Title",
               "description": "Description",
               "url": "URL",
               "picurl": "PIC_URL"
           }
       ]
   }
}
</pre>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:120px">参数</th>
<th style="width:120px">必须</th><th>说明</th></tr>
<tr><td> touser</td><td> 否</td><td> UserID列表（消息接收者，多个接收者用‘|’分隔）。特殊情况：指定为@all，则向关注该企业应用的全部成员发送</td></tr>
<tr><td> toparty</td><td> 否</td><td> PartyID列表，多个接受者用‘|’分隔。当touser为@all时忽略本参数</td></tr>
<tr><td> totag</td><td> 否</td><td> TagID列表，多个接受者用‘|’分隔。当touser为@all时忽略本参数</td></tr>
<tr><td> msgtype</td><td> 是</td><td> 消息类型，此时固定为：news</td></tr>
<tr><td> agentid</td><td> 是</td><td> 企业应用的id，整型。可在应用的设置页面查看</td></tr>
<tr><td> articles</td><td> 是</td><td> 图文消息，一个图文消息支持1到10条图文</td></tr>
<tr><td> title</td><td> 否</td><td> 标题</td></tr>
<tr><td> description</td><td> 否</td><td> 描述</td></tr>
<tr><td> url</td><td> 否</td><td> 点击后跳转的链接。企业可根据url里面带的code参数校验员工的真实身份。具体参考“9 微信页面跳转员工身份查询”</td></tr>
<tr><td> picurl</td><td> 否</td><td> 图文消息的图片链接，支持JPG、PNG格式，较好的效果为大图640*320，小图80*80。如不填，在客户端不显示图片</td></tr></tbody></table>

### **mpnews消息**
注：mpnews消息与news消息类似，不同的是图文消息内容存储在微信后台，并且支持保密选项。
<pre>{
   "touser": "UserID1|UserID2|UserID3",
   "toparty": " PartyID1 | PartyID2 ",
   "totag": " TagID1 | TagID2 ",
   "msgtype": "mpnews",
   "agentid": "1",
   "mpnews": {
       "articles":[
           {
               "title": "Title",
               "thumb_media_id": "id",
               "author": "Author",
               "content_source_url": "URL",
               "content": "Content",
               "digest": "Digest description",
               "show_cover_pic": "0"
           },
           {
               "title": "Title",
               "thumb_media_id": "id",
               "author": "Author",
               "content_source_url": "URL",
               "content": "Content",
               "digest": "Digest description",
               "show_cover_pic": "0"
           }
       ]
   }
   "safe":"0"
}
</pre>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:120px">参数</th>
<th style="width:120px">必须</th><th>说明</th></tr>
<tr><td> touser</td><td> 否</td><td> UserID列表（消息接收者，多个接收者用‘|’分隔）。特殊情况：指定为@all，则向关注该企业应用的全部成员发送</td></tr>
<tr><td> toparty</td><td> 否</td><td> PartyID列表，多个接受者用‘|’分隔。当touser为@all时忽略本参数</td></tr>
<tr><td> totag</td><td> 否</td><td> TagID列表，多个接受者用‘|’分隔。当touser为@all时忽略本参数</td></tr>
<tr><td> msgtype</td><td> 是</td><td> 消息类型，此时固定为：mpnews</td></tr>
<tr><td> agentid</td><td> 是</td><td> 企业应用的id，整型。可在应用的设置页面查看</td></tr>
<tr><td> articles</td><td> 是</td><td> 图文消息，一个图文消息支持1到10个图文</td></tr>
<tr><td> title</td><td> 是</td><td> 图文消息的标题</td></tr>
<tr><td> thumb_media_id</td><td> 是</td><td> 图文消息缩略图的media_id, 可以在上传多媒体文件接口中获得。此处thumb_media_id即上传接口返回的media_id</td></tr>
<tr><td> author</td><td> 否</td><td> 图文消息的作者</td></tr>
<tr><td> content_source_url</td><td> 否</td><td> 图文消息点击“阅读原文”之后的页面链接</td></tr>
<tr><td> content</td><td> 是</td><td> 图文消息的内容，支持html标签</td></tr>
<tr><td> digest</td><td> 否</td><td> 图文消息的描述</td></tr>
<tr><td> show_cover_pic</td><td> 否</td><td> 是否显示封面，1为显示，0为不显示</td></tr>
<tr><td> safe</td><td> 否</td><td> 表示是否是保密消息，0表示否，1表示是，默认0</td></tr></tbody></table>

<!--
NewPP limit report
CPU time usage: 0.048 seconds
Real time usage: 0.052 seconds
Preprocessor visited node count: 204/1000000
Preprocessor generated node count: 466/1000000
Post‐expand include size: 0/2097152 bytes
Template argument size: 0/2097152 bytes
Highest expansion depth: 2/40
Expensive parser function count: 0/100
-->
<!-- Saved in parser cache with key db_wiki:pcache:idhash:18-0!*!*!!zh-cn!*!* and timestamp 20150116073930 and revision id 538
 -->

# 自定义菜单

企业号的每个应用都可以拥有自己的菜单，企业可以调用接口来创建、删除、获取应用菜单。

注意，在操作应用的菜单时，**应用必须处于回调模式；菜单最多为两级，一级菜单最多为3个，二级菜单最多为5个**。

<!--
NewPP limit report
CPU time usage: 0.004 seconds
Real time usage: 0.003 seconds
Preprocessor visited node count: 1/1000000
Preprocessor generated node count: 4/1000000
Post‐expand include size: 0/2097152 bytes
Template argument size: 0/2097152 bytes
Highest expansion depth: 1/40
Expensive parser function count: 0/100
-->
<!-- Saved in parser cache with key db_wiki:pcache:idhash:38-0!*!*!*!*!*!* and timestamp 20150115105245 and revision id 239
 -->
## 创建应用菜单
## **创建应用菜单**
目前自定义菜单最多包括3个一级菜单，每个一级菜单最多包含5个二级菜单。一级菜单最多4个汉字，二级菜单最多7个汉字，多出来的部分将会以“...”代替。请注意，<b>**创建自定义菜单后，由于微信客户端缓存，需要24小时微信客户端才会展现出来。**</b>建议测试时可以尝试取消关注企业号后再次关注，则可以看到创建后的效果。
<br>
自定义菜单接口可实现多种类型按钮，如下：
<table border="1" cellspacing="0" cellpadding="4" align="center" width="840px">
<tbody><tr>
<th style="width:120px">字段值</th>
<th style="width:180px">功能名称</th><th>说明</th></tr>
<tr><td> click</td><td> 点击推事件</td><td> 用户点击click类型按钮后，微信服务器会通过消息接口推送消息类型为event    的结构给开发者（参考消息接口指南），并且带上按钮中开发者填写的key值，开发者可以通过自定义的key值与用户进行交互；</td></tr>
<tr><td> view</td><td> 跳转URL</td><td> 用户点击view类型按钮后，微信客户端将会打开开发者在按钮中填写的网页URL，可与网页授权获取用户基本信息接口结合，获得用户基本信息。</td></tr>
<tr><td> scancode_push</td><td> 扫码推事件</td><td> 用户点击按钮后，微信客户端将调起扫一扫工具，完成扫码操作后显示扫描结果（如果是URL，将进入URL），且会将扫码的结果传给开发者，开发者可以下发消息。</td></tr>
<tr><td> scancode_waitmsg</td><td> 扫码推事件且弹出“消息接收中”提示框</td><td> 用户点击按钮后，微信客户端将调起扫一扫工具，完成扫码操作后，将扫码的结果传给开发者，同时收起扫一扫工具，然后弹出“消息接收中”提示框，随后可能会收到开发者下发的消息。</td></tr>
<tr><td> pic_sysphoto</td><td> 弹出系统拍照发图</td><td> 用户点击按钮后，微信客户端将调起系统相机，完成拍照操作后，会将拍摄的相片发送给开发者，并推送事件给开发者，同时收起系统相机，随后可能会收到开发者下发的消息。</td></tr>
<tr><td> pic_photo_or_album</td><td> 弹出拍照或者相册发图</td><td> 用户点击按钮后，微信客户端将弹出选择器供用户选择“拍照”或者“从手机相册选择”。用户选择后即走其他两种流程。</td></tr>
<tr><td> pic_weixin</td><td> 弹出微信相册发图器</td><td> 用户点击按钮后，微信客户端将调起微信相册，完成选择操作后，将选择的相片发送给开发者的服务器，并推送事件给开发者，同时收起相册，随后可能会收到开发者下发的消息。</td></tr>
<tr><td> location_select</td><td> 弹出地理位置选择器</td><td>用户点击按钮后，微信客户端将调起地理位置选择工具，完成选择操作后，将选择的地理位置发送给开发者的服务器，同时收起位置选择工具，随后可能会收到开发者下发的消息。</td></tr></tbody></table>
<br>
<b>**请注意，除click和view外所有事件，仅支持微信iPhone5.4.1/Android5.4以上版本，旧版本微信用户点击后将没有回应，开发者也不能正常接收到事件推送。**</b>
<ul>
<li>请求说明</li>
</ul>
Https请求方式: POST

<a rel="nofollow" class="external free" href="https://qyapi.weixin.qq.com/cgi-bin/menu/create?access_token=ACCESS_TOKEN&amp;agentid=1">https://qyapi.weixin.qq.com/cgi-bin/menu/create?access_token=ACCESS_TOKEN&amp;agentid=1</a>

click和view类型请求包如下：
<pre>{
   "button":[
       {
           "type":"click",
           "name":"今日歌曲",
           "key":"V1001_TODAY_MUSIC"
       },
       {
           "name":"菜单",
           "sub_button":[
               {
                   "type":"view",
                   "name":"搜索",
                   "url":"<a rel="nofollow" class="external free" href="http://www.soso.com/">http://www.soso.com/</a>"
               },
               {
                   "type":"click",
                   "name":"赞一下我们",
                   "key":"V1001_GOOD"
               }
           ]
      }
   ]
}
</pre>
<b>其他新增按钮类型的请求示例</b>
<pre>{
    "button": [
        {
            "name": "扫码",
            "sub_button": [
                {
                    "type": "scancode_waitmsg",
                    "name": "扫码带提示",
                    "key": "rselfmenu_0_0",
                    "sub_button": [ ]
                },
                {
                    "type": "scancode_push",
                    "name": "扫码推事件",
                    "key": "rselfmenu_0_1",
                    "sub_button": [ ]
                }
            ]
        },
        {
            "name": "发图",
            "sub_button": [
                {
                    "type": "pic_sysphoto",
                    "name": "系统拍照发图",
                    "key": "rselfmenu_1_0",
                   "sub_button": [ ]
                 },
                {
                    "type": "pic_photo_or_album",
                    "name": "拍照或者相册发图",
                    "key": "rselfmenu_1_1",
                    "sub_button": [ ]
                },
                {
                    "type": "pic_weixin",
                    "name": "微信相册发图",
                    "key": "rselfmenu_1_2",
                    "sub_button": [ ]
                }
            ]
        },
        {
            "name": "发送位置",
            "type": "location_select",
            "key": "rselfmenu_2_0"
        }
    ]
}
</pre>
<ul>
<li>参数说明</li>
</ul>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:120px">参数</th><th>必须</th><th>说明</th></tr>
<tr><td> access_token</td><td> 是</td><td> 调用接口凭证</td></tr>
<tr><td> agentid</td><td> 是</td><td> 企业应用的id，整型。可在应用的设置页面查看</td></tr>
<tr><td> button</td><td> 是</td><td> 一级菜单数组，个数应为1~3个</td></tr>
<tr><td> sub_button</td><td> 否</td><td> 二级菜单数组，个数应为1~5个</td></tr>
<tr><td> type</td><td> 是</td><td> 菜单的响应动作类型</td></tr>
<tr><td> name</td><td> 是</td><td> 菜单标题，不超过16个字节，子菜单不超过40个字节</td></tr>
<tr><td> key</td><td> click等点击类型必须</td><td> 菜单KEY值，用于消息接口推送，不超过128字节</td></tr>
<tr><td> url</td><td> view类型必须</td><td> 网页链接，员工点击菜单可打开链接，不超过256字节</td></tr></tbody></table>
<ul>
<li>权限说明</li>
</ul>
管理员须拥有应用的管理权限，并且应用必须设置在回调模式。

返回结果
<pre>{
   "errcode":0,
   "errmsg":"ok"
}
</pre>

<!--
NewPP limit report
CPU time usage: 0.024 seconds
Real time usage: 0.030 seconds
Preprocessor visited node count: 3/1000000
Preprocessor generated node count: 10/1000000
Post‐expand include size: 0/2097152 bytes
Template argument size: 0/2097152 bytes
Highest expansion depth: 2/40
Expensive parser function count: 0/100
-->
<!-- Saved in parser cache with key db_wiki:pcache:idhash:19-0!*!*!*!*!*!* and timestamp 20150116041244 and revision id 494
 -->
## 删除菜单
<ul>
<li>请求说明</li>
</ul>
Https请求方式：GET

<a rel="nofollow" class="external free" href="https://qyapi.weixin.qq.com/cgi-bin/menu/delete?access_token=ACCESS_TOKEN&amp;agentid=1">https://qyapi.weixin.qq.com/cgi-bin/menu/delete?access_token=ACCESS_TOKEN&amp;agentid=1</a>
<ul>
<li>参数说明</li>
</ul>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:120px">参数</th><th>必须</th><th>说明</th></tr>
<tr><td> access_token</td><td> 是</td><td> 调用接口凭证</td></tr>
<tr><td> agentid</td><td> 是</td><td> 企业应用的id，整型。可在应用的设置页面查看</td></tr></tbody></table>
<ul>
<li>权限说明</li>
</ul>
管理员须拥有应用的管理权限，并且应用必须设置在回调模式。
<ul>
<li>返回结果</li>
</ul>
<pre>{
   "errcode":0,
   "errmsg":"ok"
}
</pre>

<!--
NewPP limit report
CPU time usage: 0.008 seconds
Real time usage: 0.010 seconds
Preprocessor visited node count: 2/1000000
Preprocessor generated node count: 6/1000000
Post‐expand include size: 0/2097152 bytes
Template argument size: 0/2097152 bytes
Highest expansion depth: 1/40
Expensive parser function count: 0/100
-->
<!-- Saved in parser cache with key db_wiki:pcache:idhash:20-0!*!*!*!*!*!* and timestamp 20150116070907 and revision id 374
 -->
## 获取菜单列表
<ul>
<li>请求说明</li>
</ul>
Https请求方式：GET

<a rel="nofollow" class="external free" href="https://qyapi.weixin.qq.com/cgi-bin/menu/get?access_token=ACCESS_TOKEN&amp;agentid=1">https://qyapi.weixin.qq.com/cgi-bin/menu/get?access_token=ACCESS_TOKEN&amp;agentid=1</a>
<ul>
<li>参数说明</li>
</ul>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:120px">参数</th><th>必须</th><th>说明</th></tr>
<tr><td> access_token</td><td> 是</td><td> 调用接口凭证</td></tr>
<tr><td> agentid</td><td> 是</td><td> 企业应用的id，整型。可在应用的设置页面查看</td></tr></tbody></table>
<ul>
<li>权限说明</li>
</ul>
管理员须拥有应用的管理权限，并且应用必须设置在回调模式。
<ul>
<li>返回结果</li>
</ul>
返回结果与菜单创建的参数一致。

<!--
NewPP limit report
CPU time usage: 0.008 seconds
Real time usage: 0.009 seconds
Preprocessor visited node count: 2/1000000
Preprocessor generated node count: 6/1000000
Post‐expand include size: 0/2097152 bytes
Template argument size: 0/2097152 bytes
Highest expansion depth: 1/40
Expensive parser function count: 0/100
-->
<!-- Saved in parser cache with key db_wiki:pcache:idhash:21-0!*!*!*!*!*!* and timestamp 20150116052726 and revision id 375
 -->
# OAuth2验证接口
企业应用中的URL链接（包括自定义菜单或者消息中的链接），可以通过OAuth2.0验证接口来获取员工的身份信息。

通过此接口获取用户身份会有一定的时间开销。对于频繁获取用户身份的场景，建议采用如下方案：

1、企业应用中的URL链接直接填写企业自己的页面地址

2、用户跳转到企业页面时，企业校验是否有代表用户身份的cookie，此cookie由企业生成

3、如果没有获取到cookie，重定向到OAuth验证链接，获取用户身份后，由企业生成代表用户身份的cookie

4、根据cookie获取用户身份，进入相应的页面

注意，此URL的域名，**必须完全匹配企业应用设置项中的'可信域名'**，否则获取用户信息时会返回50001错误码。

<!--
NewPP limit report
CPU time usage: 0.000 seconds
Real time usage: 0.002 seconds
Preprocessor visited node count: 1/1000000
Preprocessor generated node count: 4/1000000
Post‐expand include size: 0/2097152 bytes
Template argument size: 0/2097152 bytes
Highest expansion depth: 1/40
Expensive parser function count: 0/100
-->
<!-- Saved in parser cache with key db_wiki:pcache:idhash:45-0!*!*!*!*!*!* and timestamp 20150115151455 and revision id 518
 -->
## 企业获取code
## **企业获取code**
企业如果需要员工在跳转到企业网页时带上员工的身份信息，需构造如下的链接：

<a rel="nofollow" class="external free" href="https://open.weixin.qq.com/connect/oauth2/authorize?appid=CORPID&amp;redirect_uri=REDIRECT_URI&amp;response_type=code&amp;scope=SCOPE&amp;state=STATE#wechat_redirect">https://open.weixin.qq.com/connect/oauth2/authorize?appid=CORPID&amp;redirect_uri=REDIRECT_URI&amp;response_type=code&amp;scope=SCOPE&amp;state=STATE#wechat_redirect</a>
<ul>
<li>参数说明</li>
</ul>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:120px">参数</th><th>必须</th><th>说明</th></tr>
<tr><td> appid</td><td> 是</td><td> 企业的CorpID</td></tr>
<tr><td> redirect_uri</td><td> 是</td><td> 授权后重定向的回调链接地址，请使用urlencode对链接进行处理</td></tr>
<tr><td> response_type</td><td> 是</td><td> 返回类型，此时固定为：code</td></tr>
<tr><td> scope</td><td> 是</td><td> 应用授权作用域，此时固定为：snsapi_base</td></tr>
<tr><td> state</td><td> 否</td><td> 重定向后会带上state参数，企业可以填写a-zA-Z0-9的参数值</td></tr>
<tr><td> #wechat_redirect</td><td> 是</td><td> 微信终端使用此参数判断是否需要带上身份信息</td></tr></tbody></table>
员工点击后，页面将跳转至 redirect_uri/?code=CODE&amp;state=STATE，企业可根据code参数获得员工的userid。

<!--
NewPP limit report
CPU time usage: 0.008 seconds
Real time usage: 0.014 seconds
Preprocessor visited node count: 2/1000000
Preprocessor generated node count: 8/1000000
Post‐expand include size: 0/2097152 bytes
Template argument size: 0/2097152 bytes
Highest expansion depth: 2/40
Expensive parser function count: 0/100
-->
<!-- Saved in parser cache with key db_wiki:pcache:idhash:22-0!*!*!*!*!*!* and timestamp 20150115144653 and revision id 309
 -->
## 根据Code获取成员信息
<ul>
<li>请求说明</li>
</ul>
Https请求方式：GET

<a rel="nofollow" class="external free" href="https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo?access_token=ACCESS_TOKEN&amp;code=CODE&amp;agentid=AGENTID">https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo?access_token=ACCESS_TOKEN&amp;code=CODE&amp;agentid=AGENTID</a>
<ul>
<li>参数说明</li>
</ul>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:120px">参数</th><th>必须</th><th>说明</th></tr>
<tr><td> access_token</td><td> 是</td><td> 调用接口凭证</td></tr>
<tr><td> code</td><td> 是</td><td> 通过员工授权获取到的code，每次员工授权带上的code将不一样，code只能使用一次，5分钟未被使用自动过期</td></tr>
<tr><td> agentid</td><td> 是</td><td> 跳转链接时所在的企业应用ID</td></tr></tbody></table>
<ul>
<li>权限说明</li>
</ul>
管理员须拥有agent的使用权限；agentid必须和跳转链接时所在的企业应用ID相同。
<ul>
<li>返回结果</li>
</ul>
a)正确时返回示例如下：
<pre>{
   "UserId":"USERID",
   "DeviceId":"DEVICEID"
}
</pre>
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th><th>说明</th></tr>
<tr><td> UserId</td><td> 员工UserID</td></tr>
<tr><td> DeviceId</td><td> 手机设备号(由微信在安装时随机生成)</td></tr></tbody></table>
出错时返回示例如下：
<pre>{
   "errcode": "40029",
   "errmsg": "invalid code"
}
</pre>

<!--
NewPP limit report
CPU time usage: 0.012 seconds
Real time usage: 0.011 seconds
Preprocessor visited node count: 2/1000000
Preprocessor generated node count: 6/1000000
Post‐expand include size: 0/2097152 bytes
Template argument size: 0/2097152 bytes
Highest expansion depth: 1/40
Expensive parser function count: 0/100
-->
<!-- Saved in parser cache with key db_wiki:pcache:idhash:51-0!*!*!*!*!*!* and timestamp 20150116034148 and revision id 520
 -->
# 微信JS接口
<div style="text-align:center;font-size:24px;"><b>微信JS-SDK说明文档</b></div>

<div id="toc" class="toc"><div id="toctitle">## 目录</div>
<ul>
<li class="toclevel-1 tocsection-1"><a href="#.E6.A6.82.E8.BF.B0"><span class="tocnumber">1</span> <span class="toctext"><b>概述</a>
<ul>
<li class="toclevel-2 tocsection-2"><a href="#.E4.BD.BF.E7.94.A8.E8.AF.B4.E6.98.8E"><span class="tocnumber">1.1</span> <span class="toctext"><b>使用说明</a>
<ul>
<li class="toclevel-3 tocsection-3"><a href="#.E6.AD.A5.E9.AA.A4.E4.B8.80.EF.BC.9A.E5.BC.95.E5.85.A5JS.E6.96.87.E4.BB.B6"><span class="tocnumber">1.1.1</span> <span class="toctext"><b>步骤一：引入JS文件</a></li>
<li class="toclevel-3 tocsection-4"><a href="#.E6.AD.A5.E9.AA.A4.E4.BA.8C.EF.BC.9A.E9.80.9A.E8.BF.87config.E6.8E.A5.E5.8F.A3.E6.B3.A8.E5.85.A5.E6.9D.83.E9.99.90.E9.AA.8C.E8.AF.81.E9.85.8D.E7.BD.AE"><span class="tocnumber">1.1.2</span> <span class="toctext"><b>步骤二：通过config接口注入权限验证配置</a></li>
<li class="toclevel-3 tocsection-5"><a href="#.E6.AD.A5.E9.AA.A4.E4.B8.89.EF.BC.9A.E9.80.9A.E8.BF.87ready.E6.8E.A5.E5.8F.A3.E5.A4.84.E7.90.86.E6.88.90.E5.8A.9F.E9.AA.8C.E8.AF.81"><span class="tocnumber">1.1.3</span> <span class="toctext"><b>步骤三：通过ready接口处理成功验证</a></li>
<li class="toclevel-3 tocsection-6"><a href="#.E6.AD.A5.E9.AA.A4.E5.9B.9B.EF.BC.9A.E9.80.9A.E8.BF.87error.E6.8E.A5.E5.8F.A3.E5.A4.84.E7.90.86.E5.A4.B1.E8.B4.A5.E9.AA.8C.E8.AF.81"><span class="tocnumber">1.1.4</span> <span class="toctext"><b>步骤四：通过error接口处理失败验证</a></li>
</ul></li>
<li class="toclevel-2 tocsection-7"><a href="#.E6.8E.A5.E5.8F.A3.E8.B0.83.E7.94.A8.E8.AF.B4.E6.98.8E"><span class="tocnumber">1.2</span> <span class="toctext"><b>接口调用说明</a></li>
</ul></li>
<li class="toclevel-1 tocsection-8"><a href="#.E5.9F.BA.E7.A1.80.E6.8E.A5.E5.8F.A3"><span class="tocnumber">2</span> <span class="toctext"><b>基础接口</a>
<ul>
<li class="toclevel-2 tocsection-9"><a href="#.E5.88.A4.E6.96.AD.E5.BD.93.E5.89.8D.E5.AE.A2.E6.88.B7.E7.AB.AF.E7.89.88.E6.9C.AC.E6.98.AF.E5.90.A6.E6.94.AF.E6.8C.81.E6.8C.87.E5.AE.9AJS.E6.8E.A5.E5.8F.A3"><span class="tocnumber">2.1</span> <span class="toctext"><b>判断当前客户端版本是否支持指定JS接口</a></li>
</ul></li>
<li class="toclevel-1 tocsection-10"><a href="#.E5.88.86.E4.BA.AB.E6.8E.A5.E5.8F.A3"><span class="tocnumber">3</span> <span class="toctext"><b>分享接口</a>
<ul>
<li class="toclevel-2 tocsection-11"><a href="#.E8.8E.B7.E5.8F.96.E2.80.9C.E5.88.86.E4.BA.AB.E5.88.B0.E6.9C.8B.E5.8F.8B.E5.9C.88.E2.80.9D.E6.8C.89.E9.92.AE.E7.82.B9.E5.87.BB.E7.8A.B6.E6.80.81.E5.8F.8A.E8.87.AA.E5.AE.9A.E4.B9.89.E5.88.86.E4.BA.AB.E5.86.85.E5.AE.B9.E6.8E.A5.E5.8F.A3"><span class="tocnumber">3.1</span> <span class="toctext"><b>获取“分享到朋友圈”按钮点击状态及自定义分享内容接口</a></li>
<li class="toclevel-2 tocsection-12"><a href="#.E8.8E.B7.E5.8F.96.E2.80.9C.E5.88.86.E4.BA.AB.E7.BB.99.E6.9C.8B.E5.8F.8B.E2.80.9D.E6.8C.89.E9.92.AE.E7.82.B9.E5.87.BB.E7.8A.B6.E6.80.81.E5.8F.8A.E8.87.AA.E5.AE.9A.E4.B9.89.E5.88.86.E4.BA.AB.E5.86.85.E5.AE.B9.E6.8E.A5.E5.8F.A3"><span class="tocnumber">3.2</span> <span class="toctext"><b>获取“分享给朋友”按钮点击状态及自定义分享内容接口</a></li>
<li class="toclevel-2 tocsection-13"><a href="#.E8.8E.B7.E5.8F.96.E2.80.9C.E5.88.86.E4.BA.AB.E5.88.B0QQ.E2.80.9D.E6.8C.89.E9.92.AE.E7.82.B9.E5.87.BB.E7.8A.B6.E6.80.81.E5.8F.8A.E8.87.AA.E5.AE.9A.E4.B9.89.E5.88.86.E4.BA.AB.E5.86.85.E5.AE.B9.E6.8E.A5.E5.8F.A3"><span class="tocnumber">3.3</span> <span class="toctext"><b>获取“分享到QQ”按钮点击状态及自定义分享内容接口</a></li>
<li class="toclevel-2 tocsection-14"><a href="#.E8.8E.B7.E5.8F.96.E2.80.9C.E5.88.86.E4.BA.AB.E5.88.B0.E8.85.BE.E8.AE.AF.E5.BE.AE.E5.8D.9A.E2.80.9D.E6.8C.89.E9.92.AE.E7.82.B9.E5.87.BB.E7.8A.B6.E6.80.81.E5.8F.8A.E8.87.AA.E5.AE.9A.E4.B9.89.E5.88.86.E4.BA.AB.E5.86.85.E5.AE.B9.E6.8E.A5.E5.8F.A3"><span class="tocnumber">3.4</span> <span class="toctext"><b>获取“分享到腾讯微博”按钮点击状态及自定义分享内容接口</a></li>
</ul></li>
<li class="toclevel-1 tocsection-15"><a href="#.E5.9B.BE.E5.83.8F.E6.8E.A5.E5.8F.A3"><span class="tocnumber">4</span> <span class="toctext"><b>图像接口</a>
<ul>
<li class="toclevel-2 tocsection-16"><a href="#.E6.8B.8D.E7.85.A7.E6.88.96.E4.BB.8E.E6.89.8B.E6.9C.BA.E7.9B.B8.E5.86.8C.E4.B8.AD.E9.80.89.E5.9B.BE.E6.8E.A5.E5.8F.A3"><span class="tocnumber">4.1</span> <span class="toctext"><b>拍照或从手机相册中选图接口</a></li>
<li class="toclevel-2 tocsection-17"><a href="#.E9.A2.84.E8.A7.88.E5.9B.BE.E7.89.87.E6.8E.A5.E5.8F.A3"><span class="tocnumber">4.2</span> <span class="toctext"><b>预览图片接口</a></li>
<li class="toclevel-2 tocsection-18"><a href="#.E4.B8.8A.E4.BC.A0.E5.9B.BE.E7.89.87.E6.8E.A5.E5.8F.A3"><span class="tocnumber">4.3</span> <span class="toctext"><b>上传图片接口</a></li>
<li class="toclevel-2 tocsection-19"><a href="#.E4.B8.8B.E8.BD.BD.E5.9B.BE.E7.89.87.E6.8E.A5.E5.8F.A3"><span class="tocnumber">4.4</span> <span class="toctext"><b>下载图片接口</a></li>
</ul></li>
<li class="toclevel-1 tocsection-20"><a href="#.E9.9F.B3.E9.A2.91.E6.8E.A5.E5.8F.A3"><span class="tocnumber">5</span> <span class="toctext"><b>音频接口</a>
<ul>
<li class="toclevel-2 tocsection-21"><a href="#.E5.BC.80.E5.A7.8B.E5.BD.95.E9.9F.B3.E6.8E.A5.E5.8F.A3"><span class="tocnumber">5.1</span> <span class="toctext"><b>开始录音接口</a></li>
<li class="toclevel-2 tocsection-22"><a href="#.E5.81.9C.E6.AD.A2.E5.BD.95.E9.9F.B3.E6.8E.A5.E5.8F.A3"><span class="tocnumber">5.2</span> <span class="toctext"><b>停止录音接口</a></li>
<li class="toclevel-2 tocsection-23"><a href="#.E7.9B.91.E5.90.AC.E5.BD.95.E9.9F.B3.E8.87.AA.E5.8A.A8.E5.81.9C.E6.AD.A2.E6.8E.A5.E5.8F.A3"><span class="tocnumber">5.3</span> <span class="toctext"><b>监听录音自动停止接口</a></li>
<li class="toclevel-2 tocsection-24"><a href="#.E6.92.AD.E6.94.BE.E8.AF.AD.E9.9F.B3.E6.8E.A5.E5.8F.A3"><span class="tocnumber">5.4</span> <span class="toctext"><b>播放语音接口</a></li>
<li class="toclevel-2 tocsection-25"><a href="#.E6.9A.82.E5.81.9C.E6.92.AD.E6.94.BE.E6.8E.A5.E5.8F.A3"><span class="tocnumber">5.5</span> <span class="toctext"><b>暂停播放接口</a></li>
<li class="toclevel-2 tocsection-26"><a href="#.E5.81.9C.E6.AD.A2.E6.92.AD.E6.94.BE.E6.8E.A5.E5.8F.A3"><span class="tocnumber">5.6</span> <span class="toctext"><b>停止播放接口</a></li>
<li class="toclevel-2 tocsection-27"><a href="#.E7.9B.91.E5.90.AC.E8.AF.AD.E9.9F.B3.E6.92.AD.E6.94.BE.E5.AE.8C.E6.AF.95.E6.8E.A5.E5.8F.A3"><span class="tocnumber">5.7</span> <span class="toctext"><b>监听语音播放完毕接口</a></li>
<li class="toclevel-2 tocsection-28"><a href="#.E4.B8.8A.E4.BC.A0.E8.AF.AD.E9.9F.B3.E6.8E.A5.E5.8F.A3"><span class="tocnumber">5.8</span> <span class="toctext"><b>上传语音接口</a></li>
<li class="toclevel-2 tocsection-29"><a href="#.E4.B8.8B.E8.BD.BD.E8.AF.AD.E9.9F.B3.E6.8E.A5.E5.8F.A3"><span class="tocnumber">5.9</span> <span class="toctext"><b>下载语音接口</a></li>
</ul></li>
<li class="toclevel-1 tocsection-30"><a href="#.E6.99.BA.E8.83.BD.E6.8E.A5.E5.8F.A3"><span class="tocnumber">6</span> <span class="toctext"><b>智能接口</a>
<ul>
<li class="toclevel-2 tocsection-31"><a href="#.E8.AF.86.E5.88.AB.E9.9F.B3.E9.A2.91.E5.B9.B6.E8.BF.94.E5.9B.9E.E8.AF.86.E5.88.AB.E7.BB.93.E6.9E.9C.E6.8E.A5.E5.8F.A3"><span class="tocnumber">6.1</span> <span class="toctext"><b>识别音频并返回识别结果接口</a></li>
</ul></li>
<li class="toclevel-1 tocsection-32"><a href="#.E8.AE.BE.E5.A4.87.E4.BF.A1.E6.81.AF"><span class="tocnumber">7</span> <span class="toctext"><b>设备信息</a>
<ul>
<li class="toclevel-2 tocsection-33"><a href="#.E8.8E.B7.E5.8F.96.E7.BD.91.E7.BB.9C.E7.8A.B6.E6.80.81.E6.8E.A5.E5.8F.A3"><span class="tocnumber">7.1</span> <span class="toctext"><b>获取网络状态接口</a></li>
</ul></li>
<li class="toclevel-1 tocsection-34"><a href="#.E5.9C.B0.E7.90.86.E4.BD.8D.E7.BD.AE"><span class="tocnumber">8</span> <span class="toctext"><b>地理位置</a>
<ul>
<li class="toclevel-2 tocsection-35"><a href="#.E4.BD.BF.E7.94.A8.E5.BE.AE.E4.BF.A1.E5.86.85.E7.BD.AE.E5.9C.B0.E5.9B.BE.E6.9F.A5.E7.9C.8B.E4.BD.8D.E7.BD.AE.E6.8E.A5.E5.8F.A3"><span class="tocnumber">8.1</span> <span class="toctext"><b>使用微信内置地图查看位置接口</a></li>
<li class="toclevel-2 tocsection-36"><a href="#.E8.8E.B7.E5.8F.96.E5.9C.B0.E7.90.86.E4.BD.8D.E7.BD.AE.E6.8E.A5.E5.8F.A3"><span class="tocnumber">8.2</span> <span class="toctext"><b>获取地理位置接口</a></li>
</ul></li>
<li class="toclevel-1 tocsection-37"><a href="#.E7.95.8C.E9.9D.A2.E6.93.8D.E4.BD.9C"><span class="tocnumber">9</span> <span class="toctext"><b>界面操作</a>
<ul>
<li class="toclevel-2 tocsection-38"><a href="#.E9.9A.90.E8.97.8F.E5.8F.B3.E4.B8.8A.E8.A7.92.E8.8F.9C.E5.8D.95.E6.8E.A5.E5.8F.A3"><span class="tocnumber">9.1</span> <span class="toctext"><b>隐藏右上角菜单接口</a></li>
<li class="toclevel-2 tocsection-39"><a href="#.E6.98.BE.E7.A4.BA.E5.8F.B3.E4.B8.8A.E8.A7.92.E8.8F.9C.E5.8D.95.E6.8E.A5.E5.8F.A3"><span class="tocnumber">9.2</span> <span class="toctext"><b>显示右上角菜单接口</a></li>
<li class="toclevel-2 tocsection-40"><a href="#.E5.85.B3.E9.97.AD.E5.BD.93.E5.89.8D.E7.BD.91.E9.A1.B5.E7.AA.97.E5.8F.A3.E6.8E.A5.E5.8F.A3"><span class="tocnumber">9.3</span> <span class="toctext"><b>关闭当前网页窗口接口</a></li>
<li class="toclevel-2 tocsection-41"><a href="#.E6.89.B9.E9.87.8F.E9.9A.90.E8.97.8F.E5.8A.9F.E8.83.BD.E6.8C.89.E9.92.AE.E6.8E.A5.E5.8F.A3"><span class="tocnumber">9.4</span> <span class="toctext"><b>批量隐藏功能按钮接口</a></li>
<li class="toclevel-2 tocsection-42"><a href="#.E6.89.B9.E9.87.8F.E6.98.BE.E7.A4.BA.E5.8A.9F.E8.83.BD.E6.8C.89.E9.92.AE.E6.8E.A5.E5.8F.A3"><span class="tocnumber">9.5</span> <span class="toctext"><b>批量显示功能按钮接口</a></li>
<li class="toclevel-2 tocsection-43"><a href="#.E9.9A.90.E8.97.8F.E6.89.80.E6.9C.89.E9.9D.9E.E5.9F.BA.E7.A1.80.E6.8C.89.E9.92.AE.E6.8E.A5.E5.8F.A3"><span class="tocnumber">9.6</span> <span class="toctext"><b>隐藏所有非基础按钮接口</a></li>
<li class="toclevel-2 tocsection-44"><a href="#.E6.98.BE.E7.A4.BA.E6.89.80.E6.9C.89.E5.8A.9F.E8.83.BD.E6.8C.89.E9.92.AE.E6.8E.A5.E5.8F.A3"><span class="tocnumber">9.7</span> <span class="toctext"><b>显示所有功能按钮接口</a></li>
</ul></li>
<li class="toclevel-1 tocsection-45"><a href="#.E5.BE.AE.E4.BF.A1.E6.89.AB.E4.B8.80.E6.89.AB"><span class="tocnumber">10</span> <span class="toctext"><b>微信扫一扫</a>
<ul>
<li class="toclevel-2 tocsection-46"><a href="#.E8.B0.83.E8.B5.B7.E5.BE.AE.E4.BF.A1.E6.89.AB.E4.B8.80.E6.89.AB.E6.8E.A5.E5.8F.A3"><span class="tocnumber">10.1</span> <span class="toctext"><b>调起微信扫一扫接口</a></li>
</ul></li>
<li class="toclevel-1 tocsection-47"><a href="#.E9.99.84.E5.BD.951-JS-SDK.E4.BD.BF.E7.94.A8.E6.9D.83.E9.99.90.E7.AD.BE.E5.90.8D.E7.AE.97.E6.B3.95"><span class="tocnumber">11</span> <span class="toctext"><b>附录1-JS-SDK使用权限签名算法</a></li>
<li class="toclevel-1 tocsection-48"><a href="#.E9.99.84.E5.BD.952-.E6.89.80.E6.9C.89JS.E6.8E.A5.E5.8F.A3.E5.88.97.E8.A1.A8"><span class="tocnumber">12</span> <span class="toctext"><b>附录2-所有JS接口列表</a></li>
<li class="toclevel-1 tocsection-49"><a href="#.E9.99.84.E5.BD.953-.E6.89.80.E6.9C.89.E8.8F.9C.E5.8D.95.E9.A1.B9.E5.88.97.E8.A1.A8"><span class="tocnumber">13</span> <span class="toctext"><b>附录3-所有菜单项列表</a></li>
<li class="toclevel-1 tocsection-50"><a href="#.E9.99.84.E5.BD.954-.E4.BD.8D.E7.BD.AE.E7.AD.BE.E5.90.8D.E7.94.9F.E6.88.90.E7.AE.97.E6.B3.95"><span class="tocnumber">14</span> <span class="toctext"><b>附录4-位置签名生成算法</a></li>
<li class="toclevel-1 tocsection-51"><a href="#.E9.99.84.E5.BD.955-.E5.B8.B8.E8.A7.81.E9.94.99.E8.AF.AF.E5.8F.8A.E8.A7.A3.E5.86.B3.E6.96.B9.E6.B3.95"><span class="tocnumber">15</span> <span class="toctext"><b>附录5-常见错误及解决方法</a></li>
<li class="toclevel-1 tocsection-52"><a href="#.E9.99.84.E5.BD.955-DEMO.E9.A1.B5.E9.9D.A2.E5.92.8C.E7.A4.BA.E4.BE.8B.E4.BB.A3.E7.A0.81"><span class="tocnumber">16</span> <span class="toctext"><b>附录5-DEMO页面和示例代码</a></li>
<li class="toclevel-1 tocsection-53"><a href="#.E9.99.84.E5.BD.956-.E9.97.AE.E9.A2.98.E5.8F.8D.E9.A6.88"><span class="tocnumber">17</span> <span class="toctext"><b>附录6-问题反馈</a></li>
</ul>
</div>
## 概述
微信JS-SDK是微信公众平台面向网页开发者提供的基于微信内的网页开发工具包。
通过使用微信JS-SDK，网页开发者可借助微信高效地使用拍照、选图、语音、位置等手机系统的能力，同时可以直接使用微信分享、扫一扫等微信特有的能力，为微信用户提供更优质的网页体验。
此文档面向网页开发者介绍微信JS-SDK如何使用及相关注意事项。
### 使用说明
在使用微信JS-SDK对应的JS接口前，需确保已获得使用对应JS接口的权限，可在下表中根据自己的帐号角色查看。
企业号帐号角色分为体验号、注册号和认证号，其中体验号与注册号的权限一致，认证号则拥有更多的JS-SDK权限，具体详见下方表格：
<table border="1" cellspacing="0" cellpadding="4" align="center" width="800px">
<tbody><tr>
<th style="width:180px"> 功能</th>
<th style="width:520px"> 接口</th>
<th style="width:150px"> 注册号和体验号</th>
<th style="width:150px"> 认证号</th></tr>
<tr>
<th rowspan="1"> 基础接口</th><td> 判断当前客户端版本是否支持指定JS接口</td><td> 有</td><td> 有</td></tr>
<tr>
<th rowspan="4"> 分享接口</th><td> 获取“分享到朋友圈”按钮点击状态及设置分享内容接口</td><td> 无</td><td> 有</td></tr>
<tr><td> 获取“分享给朋友”按钮点击状态及设置分享内容接口</td><td> 无</td><td> 有</td></tr>
<tr><td> 获取“分享到QQ”按钮点击状态及设置分享内容接口</td><td> 无</td><td> 有</td></tr>
<tr><td> 获取“分享到腾讯微博”按钮点击状态及设置分享内容接</td><td> 无</td><td> 有</td></tr>
<tr>
<th rowspan="4"> 图像接口</th><td> 本地选图或拍照接口</td><td> 有</td><td> 有</td></tr>
<tr><td> 图片预览接口</td><td> 有</td><td> 有</td></tr>
<tr><td> 上传图片接口</td><td> 有</td><td> 有</td></tr>
<tr><td> 下载图片接口</td><td> 有</td><td> 有</td></tr>
<tr>
<th rowspan="7"> 音频接口</th><td> 开始录音接口</td><td> 有</td><td> 有</td></tr>
<tr><td> 停止录音接口</td><td> 有</td><td> 有</td></tr>
<tr><td> 播放音频接口</td><td> 有</td><td> 有</td></tr>
<tr><td> 暂停播放接口</td><td> 有</td><td> 有</td></tr>
<tr><td> 停止播放接口</td><td> 有</td><td> 有</td></tr>
<tr><td> 上传语音接口</td><td> 有</td><td> 有</td></tr>
<tr><td> 下载语音接口</td><td> 有</td><td> 有</td></tr>
<tr>
<th rowspan="1"> 智能接口</th><td> 识别音频并返回识别结果接口</td><td> 有</td><td> 有</td></tr>
<tr>
<th rowspan="1"> 设备信息</th><td> 获取网络状态接口</td><td> 有</td><td> 有</td></tr>
<tr>
<th rowspan="2"> 地理位置</th><td> 查看地理位置地图接口</td><td> 有</td><td> 有</td></tr>
<tr><td> 获取地理位置接口</td><td> 有</td><td> 有</td></tr>
<tr>
<th rowspan="7"> 地理位置</th><td> 隐藏右上角菜单接口</td><td> 有</td><td> 有</td></tr>
<tr><td> 显示右上角菜单接口</td><td> 有</td><td> 有</td></tr>
<tr><td> 关闭当前窗口接口</td><td> 有</td><td> 有</td></tr>
<tr><td> 批量隐藏菜单项接口</td><td> 有</td><td> 有</td></tr>
<tr><td> 批量显示菜单项接口</td><td> 有</td><td> 有</td></tr>
<tr><td> 隐藏所有非基本菜单项接口</td><td> 有</td><td> 有</td></tr>
<tr><td> 显示所有被隐藏的非基本菜单项接口</td><td> 有</td><td> 有</td></tr>
<tr>
<th rowspan="1"> 微信扫一扫</th><td> 扫一扫接口</td><td> 有</td><td> 有</td></tr></tbody></table>
<font color="red">
注意： 所有的JS接口只能在企业号应用的可信域名下调用，可在企业号应用中心里设置应用可信域名。
</font>
<h4>步骤一：引入JS文件</h4>
在需要调用JS接口的页面引入如下JS文件，（支持https）：<a rel="nofollow" class="external free" href="http://res.wx.qq.com/open/js/jweixin-1.0.0.js">http://res.wx.qq.com/open/js/jweixin-1.0.0.js</a>
备注：支持使用 AMD/CMD 标准模块加载方法加载
<h4>步骤二：通过config接口注入权限验证配置</h4>
<font color="#FF0000">
所有需要使用JS-SDK的页面必须先注入配置信息，否则将无法调用（同一个url仅需调用一次，对于变化url的SPA的web app可在每次url变化时进行调用）。
</font>
<pre>wx.config({
    debug: true, // 开启调试模式,调用的所有api的返回值会在客户端alert出来，若要查看传入的参数，可以在pc端打开，参数信息会通过log打出，仅在pc端时才会打印。
    appId: '', // 必填，企业号的唯一标识，此处填写企业号corpid
    timestamp: , // 必填，生成签名的时间戳
    nonceStr: '', // 必填，生成签名的随机串
    signature: '',// 必填，签名，见附录1
    jsApiList: [] // 必填，需要使用的JS接口列表，所有JS接口列表见附录2
});
</pre>

<h4>步骤三：通过ready接口处理成功验证</h4>
<pre>wx.ready(function(){
    // config信息验证后会执行ready方法，所有接口调用都必须在config接口获得结果之后，config是一个客户端的异步操作，所以如果需要在页面加载时就调用相关接口，则须把相关接口放在ready函数中调用来确保正确执行。对于用户触发时才调用的接口，则可以直接调用，不需要放在ready函数中。
});
</pre>

<h4>步骤四：通过error接口处理失败验证</h4>
<pre>wx.error(function(res){
    // config信息验证失败会执行error函数，如签名过期导致验证失败，具体错误信息可以打开config的debug模式查看，也可以在返回的res参数中查看，对于SPA可以在这里更新签名。
});
</pre>
### 接口调用说明
所有接口通过wx对象(也可使用jWeixin对象)来调用，参数是一个对象，除了每个接口本身需要传的参数之外，还有以下通用参数：
<ol>
<li> success：接口调用成功时执行的回调函数。</li>
<li> fail：接口调用失败时执行的回调函数。</li>
<li> complete：接口调用完成时执行的回调函数，无论成功或失败都会执行。</li>
<li> cancel：用户点击取消时的回调函数，仅部分有用户取消操作的api才会用到。</li>
<li> trigger: 监听Menu中的按钮点击时触发的方法，该方法仅支持Menu中的相关接口。</li>
</ol>
<br>
以上几个函数都带有一个参数，类型为对象，其中除了每个接口本身返回的数据之外，还有一个通用属性errMsg，其值格式如下：
<ol>
<li> 调用成功时："xxx:ok" ，其中xxx为调用的接口名</li>
<li> 用户取消时："xxx:cancel"，其中xxx为调用的接口名</li>
<li> 调用失败时：其值为具体错误信息</li>
</ol>
## 基础接口
### 判断当前客户端版本是否支持指定JS接口
<pre>wx.checkJsApi({
    jsApiList: ['chooseImage'] // 需要检测的JS接口列表，所有JS接口列表见附录2,
    success: function(res) {
        // 以键值对的形式返回，可用的api值true，不可用为false
        // 如：{"checkResult":{"chooseImage":true},"errMsg":"checkJsApi:ok"}
    });
</pre>
备注：checkJsApi接口是客户端6.0.2新引入的一个预留接口，第一期开放的接口均可不使用checkJsApi来检测。
## 分享接口
请注意不要有诱导分享等违规行为，对于诱导分享行为将永久回收企业号接口权限，详细规则请查看：<a rel="nofollow" class="external text" href="http://kf.qq.com/faq/131117ne2MV7141117JzI32q.html">朋友圈管理常见问题</a> 。
### 获取“分享到朋友圈”按钮点击状态及自定义分享内容接口
<pre>wx.onMenuShareTimeline({
    title: '', // 分享标题
    link: '', // 分享链接
    imgUrl: '', // 分享图标
    success: function () {
        // 用户确认分享后执行的回调函数
    },
    cancel: function () {
        // 用户取消分享后执行的回调函数
    }
});
</pre>
### 获取“分享给朋友”按钮点击状态及自定义分享内容接口
<pre>wx.onMenuShareAppMessage({
    title: '', // 分享标题
    desc: '', // 分享描述
    link: '', // 分享链接
    imgUrl: '', // 分享图标
    type: '', // 分享类型,music、video或link，不填默认为link
    dataUrl: '', // 如果type是music或video，则要提供数据链接，默认为空
    success: function () {
        // 用户确认分享后执行的回调函数
    },
    cancel: function () {
        // 用户取消分享后执行的回调函数
    }
});
</pre>
### 获取“分享到QQ”按钮点击状态及自定义分享内容接口
<pre>wx.onMenuShareQQ({
    title: '', // 分享标题
    desc: '', // 分享描述
    link: '', // 分享链接
    imgUrl: '' // 分享图标
    success: function () {
       // 用户确认分享后执行的回调函数
    },
    cancel: function () {
       // 用户取消分享后执行的回调函数
    }
});
</pre>
### 获取“分享到腾讯微博”按钮点击状态及自定义分享内容接口
<pre>wx.onMenuShareWeibo({
    title: '', // 分享标题
    desc: '', // 分享描述
    link: '', // 分享链接
    imgUrl: '' // 分享图标
    success: function () {
       // 用户确认分享后执行的回调函数
    },
    cancel: function () {
        // 用户取消分享后执行的回调函数
    }
});
</pre>
## 图像接口
### 拍照或从手机相册中选图接口
<pre>wx.chooseImage({
    success: function (res) {
        var localIds = res.localIds; // 返回选定照片的本地ID列表，localId可以作为img标签的src属性显示图片
    }
});
</pre>
### 预览图片接口
<pre>wx.previewImage({
    current: '', // 当前显示的图片链接
    urls: [] // 需要预览的图片链接列表
});
</pre>
### 上传图片接口
<pre>wx.uploadImage({
    localId: '', // 需要上传的图片的本地ID，由chooseImage接口获得
    isShowProgressTips: 1// 默认为1，显示进度提示
    success: function (res) {
        var serverId = res.serverId; // 返回图片的服务器端ID
    }
});
</pre>
备注：可用企业微信获取媒体文件接口下载上传的语音，此处获得的 serverId 即 media_id，参考文档
<a rel="nofollow" class="external free" href="http://qydev.weixin.qq.com/wiki/index.php?title=%E8%8E%B7%E5%8F%96%E5%AA%92%E4%BD%93%E6%96%87%E4%BB%B6">http://qydev.weixin.qq.com/wiki/index.php?title=%E8%8E%B7%E5%8F%96%E5%AA%92%E4%BD%93%E6%96%87%E4%BB%B6</a>

### 下载图片接口
<pre>wx.downloadImage({
    serverId: '', // 需要下载的图片的服务器端ID，由uploadImage接口获得
    isShowProgressTips: 1// 默认为1，显示进度提示
    success: function (res) {
        var localId = res.localId; // 返回图片下载后的本地ID
    }
});
</pre>
## 音频接口
### 开始录音接口
<pre>wx.startRecord();
</pre>
### 停止录音接口
<pre>wx.stopRecord({
    success: function (res) {
        var localId = res.localId;
    }
});
</pre>
### 监听录音自动停止接口
<pre>wx.onVoiceRecordEnd({
    // 录音时间超过一分钟没有停止的时候会执行 complete 回调
    complete: function (res) {
        var localId = res.localId;
    }
});
</pre>
### 播放语音接口
<pre>wx.playVoice({
    localId: '' // 需要播放的音频的本地ID，由stopRecord接口获得
});
</pre>

### 暂停播放接口
<pre>wx.pauseVoice({
    localId: '' // 需要暂停的音频的本地ID，由stopRecord接口获得
});
</pre>
### 停止播放接口
<pre>wx.stopVoice({
    localId: '' // 需要停止的音频的本地ID，由stopRecord接口获得
});
</pre>
### 监听语音播放完毕接口
<pre>wx.onVoicePlayEnd({
    serverId: '', // 需要下载的音频的服务器端ID，由uploadVoice接口获得
    success: function (res) {
        var localId = res.localId; // 返回音频的本地ID
    }
});
</pre>

### 上传语音接口
<pre>wx.uploadVoice({
    localId: '', // 需要上传的音频的本地ID，由stopRecord接口获得
    isShowProgressTips: 1// 默认为1，显示进度提示
        success: function (res) {
        var serverId = res.serverId; // 返回音频的服务器端ID
    }
});
</pre>
备注：可用企业微信获取媒体文件接口下载上传的语音，此处获得的 serverId 即 media_id，参考文档
<a rel="nofollow" class="external free" href="http://qydev.weixin.qq.com/wiki/index.php?title=%E8%8E%B7%E5%8F%96%E5%AA%92%E4%BD%93%E6%96%87%E4%BB%B6">http://qydev.weixin.qq.com/wiki/index.php?title=%E8%8E%B7%E5%8F%96%E5%AA%92%E4%BD%93%E6%96%87%E4%BB%B6</a>

### 下载语音接口
<pre>wx.downloadVoice({
    serverId: '', // 需要下载的音频的服务器端ID，由uploadVoice接口获得
    isShowProgressTips: 1// 默认为1，显示进度提示
    success: function (res) {
        var localId = res.localId; // 返回音频的本地ID
    }
});
</pre>

## 智能接口
### 识别音频并返回识别结果接口
<pre>wx.translateVoice({
   localId: '', // 需要识别的音频的本地Id，由录音相关接口获得
    isShowProgressTips: 1, // 默认为1，显示进度提示
    success: function (res) {
        alert(res.translateResult); // 语音识别的结果
    }
});
</pre>
## 设备信息
### 获取网络状态接口
<pre>wx.getNetworkType({
    success: function (res) {
        var networkType = res.networkType; // 返回网络类型2g，3g，4g，wifi
    }
});
</pre>

## 地理位置
### 使用微信内置地图查看位置接口
<pre>wx.openLocation({
    latitude: 0, // 纬度，浮点数，范围为90 ~ -90
    longitude: 0, // 经度，浮点数，范围为180 ~ -180。
    name: '', // 位置名
    address: '', // 地址详情说明
    scale: 1, // 地图缩放级别,整形值,范围从1~28。默认为最大
    infoUrl: '' // 在查看位置界面底部显示的超链接,可点击跳转
});
</pre>
### 获取地理位置接口
<pre>wx.getLocation({
    timestamp: 0, // 位置签名时间戳，仅当需要兼容6.0.2版本之前时提供
    nonceStr: '', // 位置签名随机串，仅当需要兼容6.0.2版本之前时提供
    addrSign: '', // 位置签名，仅当需要兼容6.0.2版本之前时提供，详见附录4
    success: function (res) {
       var longitude = res.longitude; // 纬度，浮点数，范围为90 ~ -90
       var latitude = res.latitude; // 经度，浮点数，范围为180 ~ -180。
        var speed = res.speed; // 速度，以米/每秒计
        var accuracy = res.accuracy; // 位置精度
    }
});
</pre>

## 界面操作
### 隐藏右上角菜单接口
<pre>wx.hideOptionMenu();
</pre>
### 显示右上角菜单接口
<pre>wx.showOptionMenu();
</pre>
### 关闭当前网页窗口接口
<pre>wx.closeWindow();
</pre>
### 批量隐藏功能按钮接口
<pre>wx.hideMenuItems({
    menuList: [] // 要隐藏的菜单项，所有menu项见附录3
});
</pre>
### 批量显示功能按钮接口
<pre>wx.showMenuItems({
    menuList: [] // 要显示的菜单项，所有menu项见附录3
});
</pre>
### 隐藏所有非基础按钮接口
<pre>wx.hideAllNonBaseMenuItem();
</pre>
### 显示所有功能按钮接口
<pre>wx.showAllNonBaseMenuItem();
</pre>
## 微信扫一扫
### 调起微信扫一扫接口
<pre>wx.scanQRCode({
    desc: 'scanQRCode desc',
    needResult: 0, // 默认为0，扫描结果由微信处理，1则直接返回扫描结果，
    scanType: ["qrCode","barCode"], // 可以指定扫二维码还是一维码，默认二者都有
    success: function (res) {
    var result = res.resultStr; // 当needResult 为 1 时，扫码返回的结果
}
});
</pre>

## 附录1-JS-SDK使用权限签名算法
<b>jsapi_ticket</b>
生成签名之前必须先了解一下jsapi_ticket，jsapi_ticket是企业号用于调用微信JS接口的临时票据。正常情况下，<font color="red">jsapi_ticket的有效期为7200秒</font>，通过access_token来获取。由于获取jsapi_ticket的api调用次数非常有限，频繁刷新jsapi_ticket会导致api调用受限，影响自身业务，<font color="red">开发者必须在自己的服务全局缓存jsapi_ticket</font> 。
<ol>
<li> 参考以下文档获取access_token（<font color="red">有效期7200秒，开发者必须在自己的服务全局缓存access_token</font>）：<a rel="nofollow" class="external free" href="http://qydev.weixin.qq.com/wiki/index.php?title=%E4%B8%BB%E5%8A%A8%E8%B0%83%E7%94%A8">http://qydev.weixin.qq.com/wiki/index.php?title=%E4%B8%BB%E5%8A%A8%E8%B0%83%E7%94%A8</a></li>
<li> 用第一步拿到的access_token 采用http GET方式请求获得jsapi_ticket（<font color="red">有效期7200秒，开发者必须在自己的服务全局缓存jsapi_ticket</font>）：<a rel="nofollow" class="external free" href="https://qyapi.weixin.qq.com/cgi-bin/get_jsapi_ticket?access_token=ACCESS_TOKEN">https://qyapi.weixin.qq.com/cgi-bin/get_jsapi_ticket?access_token=ACCESS_TOKEN</a></li>
</ol>
<br>
成功返回如下JSON：
<pre>{
"errcode":0,
"errmsg":"ok",
"ticket":"bxLdikRXVbTPdHSM05e5u5sUoXNKd8-41ZO3MhKoyN5OfkWITDGgnr2fwJ0m9E8NYzWKVZvdVtaUgWvsdshFKA",
"expires_in":7200
}
</pre>
获得jsapi_ticket之后，就可以生成JS-SDK权限验证的签名了。

<b>签名算法</b>
签名生成规则如下：参与签名的字段包括noncestr（随机字符串）, 有效的jsapi_ticket, timestamp（时间戳）, url（当前网页的URL，<font color="red">不包含#及其后面部分</font>） 。对所有待签名参数按照字段名的ASCII 码从小到大排序（字典序）后，使用URL键值对的格式（即key1=value1&amp;key2=value2…）拼接成字符串string1。这里需要注意的是所有参数名均为小写字符。对string1作sha1加密，字段名和字段值都采用原始值，不进行URL 转义。
<br>
即signature=sha1(string1)。
示例：
<ul>
<li> noncestr=Wm3WZYTPz0wzccnW</li>
<li> jsapi_ticket=sM4AOVdWfPE4DxkXGEs8VMCPGGVi4C3VM0P37wVUCFvkVAy_90u5h9nbSlYy3-Sl-HhTdfl2fzFy1AOcHKP7qg</li>
<li> timestamp=1414587457</li>
<li> url=<a rel="nofollow" class="external free" href="https://qy.weixin.qq.com">https://qy.weixin.qq.com</a></li>
</ul>
<br>
步骤1. 对所有待签名参数按照字段名的ASCII 码从小到大排序（字典序）后，使用URL键值对的格式（即key1=value1&amp;key2=value2…）拼接成字符串string1：
<pre>jsapi_ticket=sM4AOVdWfPE4DxkXGEs8VMCPGGVi4C3VM0P37wVUCFvkVAy_90u5h9nbSlYy3-Sl-HhTdfl2fzFy1AOcHKP7qg&amp;noncestr=Wm3WZYTPz0wzccnW&amp;timestamp=1414587457&amp;url=https://qy.weixin.qq.com
</pre>
<br>
步骤2. 对string1进行sha1签名，得到signature：
<pre>e9b72681621e836c6babbceafc7f5df0be59932d
</pre>
注意事项
<ol>
<li> 签名用的noncestr和timestamp必须与wx.config中的nonceStr和timestamp相同。</li>
<li> 签名用的url必须是调用JS接口页面的完整URL。</li>
<li> <font color="red">出于安全考虑，开发者必须在服务器端实现签名的逻辑</font>。</li>
</ol>

## 附录2-所有JS接口列表
版本1.0.0接口
<ul>
<li> onMenuShareTimeline</li>
<li> onMenuShareAppMessage</li>
<li> onMenuShareQQ</li>
<li> onMenuShareWeibo</li>
<li> startRecord</li>
<li> stopRecord</li>
<li> onVoiceRecordEnd</li>
<li> playVoice</li>
<li> pauseVoice</li>
<li> stopVoice</li>
<li> onVoicePlayEnd</li>
<li> uploadVoice</li>
<li> downloadVoice</li>
<li> chooseImage</li>
<li> previewImage</li>
<li> uploadImage</li>
<li> downloadImage</li>
<li> translateVoice</li>
<li> getNetworkType</li>
<li> openLocation</li>
<li> getLocation</li>
<li> hideOptionMenu</li>
<li> showOptionMenu</li>
<li> hideMenuItems</li>
<li> showMenuItems</li>
<li> hideAllNonBaseMenuItem</li>
<li> showAllNonBaseMenuItem</li>
<li> closeWindow</li>
<li> scanQRCode</li>
</ul>
## 附录3-所有菜单项列表
基本类
<ul>
<li> 举报: "menuItem:exposeArticle"</li>
<li> 调整字体: "menuItem:setFont"</li>
<li> 日间模式: "menuItem:dayMode"</li>
<li> 夜间模式: "menuItem:nightMode"</li>
<li> 刷新: "menuItem:refresh"</li>
<li> 查看企业号（已添加）: "menuItem:profile"</li>
<li> 查看企业号（未添加）: "menuItem:addContact"</li>
</ul>
传播类
<ul>
<li> 发送给朋友: "menuItem:share:appMessage"</li>
<li> 分享到朋友圈: "menuItem:share:timeline"</li>
<li> 分享到QQ: "menuItem:share:qq"</li>
<li> 分享到Weibo: "menuItem:share:weiboApp"</li>
<li> 收藏: "menuItem:favorite"</li>
<li> 分享到FB: "menuItem:share:facebook"</li>
</ul>
保护类
<ul>
<li> 调试: "menuItem:jsDebug"</li>
<li> 编辑标签: "menuItem:editTag"</li>
<li> 删除: "menuItem:delete"</li>
<li> 复制链接: "menuItem:copyUrl"</li>
<li> 原网页: "menuItem:originPage"</li>
<li> 阅读模式: "menuItem:readMode"</li>
<li> 在QQ浏览器中打开: "menuItem:openWithQQBrowser"</li>
<li> 在Safari中打开: "menuItem:openWithSafari"</li>
<li> 邮件: "menuItem:share:email"</li>
<li> 一些特殊企业号: "menuItem:share:brand"</li>
</ul>

## 附录4-位置签名生成算法
addrSign的生成规则与JS-SDK权限验证的签名生成规则相同（参考附录1），只是参与签名参数有所不同。参与addrSign的签名参数有：corpid、url（当前网页url）、timestamp、noncestr、accesstoken（用户授权凭证，请参照oauth2.0 协议获取）。
## 附录5-常见错误及解决方法
调用config 接口的时候传入参数 debug: true 可以开启debug模式，页面会alert出错误信息。以下为常见错误及解决方法：
<ol>
<li> invalid url domain当前页面所在域名与使用的corpid没有绑定（可在该企业号的应用可信域名中配置域名）。</li>
<li> invalid signature签名错误。建议按如下顺序检查：
<ol>
<li> 确认签名算法正确，可用 <a rel="nofollow" class="external free" href="http://mp.weixin.qq.com/debug/cgi-bin/sandbox?t=jsapisign">http://mp.weixin.qq.com/debug/cgi-bin/sandbox?t=jsapisign</a> 页面工具进行校验。</li>
<li> 确认config中noncestr, timestamp与用以签名中的对应noncestr, timestamp一致。</li>
<li> 确认url是页面完整的url，包括GET参数部分。</li>
<li> 确认 config 中的 corpid 与用来获取 jsapi_ticket 的 corpid 一致。</li>
</ol></li>
<li> the permission value is offline verifying这个错误是因为config没有正确执行，或者是调用的JSAPI没有传入config的jsApiList参数中。建议按如下顺序检查：
<ol>
<li> 确认config正确通过。</li>
<li> 如果是在页面加载好时就调用了JSAPI，则必须写在wx.ready的回调中。</li>
<li> 确认config的jsApiList参数包含了这个JSAPI。</li>
</ol></li>
<li> permission denied该应用没有权限使用这个JSAPI。</li>
</ol>
## 附录5-DEMO页面和示例代码
<b>DEMO页面</b>：
<a rel="nofollow" class="external free" href="http://demo.open.weixin.qq.com/jssdk">http://demo.open.weixin.qq.com/jssdk</a>
<div class="center"><div class="floatnone"><a href="/wiki/index.php?title=%E6%96%87%E4%BB%B6:Jssdk_demo001.png" class="image"><img alt="Jssdk demo001.png" src="/wiki/images/thumb/7/7b/Jssdk_demo001.png/250px-Jssdk_demo001.png" width="250" height="247" srcset="/wiki/images/7/7b/Jssdk_demo001.png 1.5x, /wiki/images/7/7b/Jssdk_demo001.png 2x"></a></div></div>
<br>
<b>示例代码</b>：
<a rel="nofollow" class="external free" href="http://demo.open.weixin.qq.com/jssdk/sample.zip">http://demo.open.weixin.qq.com/jssdk/sample.zip</a>
备注：链接中包含<font color="red">php、java、nodejs以及python</font>的示例代码供第三方参考，第三方<font color="red">切记要对获取的accesstoken以及jsapi_ticket进行缓存</font>以确保不会触发频率限制。
## 附录6-问题反馈
邮箱地址：weixin-open@qq.com
邮件主题：【微信JS-SDK反馈】
邮件内容说明：
用简明的语言描述问题所在，并交代清楚遇到该问题的场景，可附上截屏图片，微信团队会尽快处理你的反馈。
<!--
NewPP limit report
CPU time usage: 0.100 seconds
Real time usage: 0.100 seconds
Preprocessor visited node count: 546/1000000
Preprocessor generated node count: 958/1000000
Post‐expand include size: 0/2097152 bytes
Template argument size: 0/2097152 bytes
Highest expansion depth: 2/40
Expensive parser function count: 0/100
-->
<!-- Saved in parser cache with key db_wiki:pcache:idhash:73-0!*!*!!zh-cn!2!* and timestamp 20150116082602 and revision id 682
 -->
# 第三方应用授权
## **目的与场景**

目前，企业号管理员要通过繁琐的操作才能与第三方应用提供商对接。通过第三方授权接口，能够减少企业号管理员的操作步骤，提升第三方应用的接入效率。

第三方应用授权方案包含以下两个场景：

一）应用提供商注册应用

你，作为应用提供商若要使用此方案，只需在企业号官网的开发者中心注册成为第三方应用提供商，创建应用套件，并在应用套件中配置好相应的应用。然后在你的官网发布应用套件，以便于企业号管理员访问即可。

另外，你必须具备以下几个条件：

1、拥有通过认证的，能证明第三方应用提供商身份的企业号。

2、具有互联网上部署及发布你的应用的能力。

本方案所说的**应用套件**，是第三方应用授权的主体，它可以包含多个第三方所提供的同一类型的应用。目前一个第三方最多可以注册五个应用套件，一个应用套件最多可以包含十五个应用。

二）企业号管理员授权应用

企业号的管理员浏览你的官网，发现适合他的应用套件后，即可发起一键授权，系统将展示企业号第三方授权页面，管理员根据授权页面的引导，确认授权内容，完成授权操作。

授权完成之后，企业号就可使用应用提供商所提供的应用服务了。一切将变得简单自然。以下章节将对每一个操作过程做具体的介绍和说明。

<!--
NewPP limit report
CPU time usage: 0.008 seconds
Real time usage: 0.011 seconds
Preprocessor visited node count: 2/1000000
Preprocessor generated node count: 8/1000000
Post‐expand include size: 0/2097152 bytes
Template argument size: 0/2097152 bytes
Highest expansion depth: 2/40
Expensive parser function count: 0/100
-->
<!-- Saved in parser cache with key db_wiki:pcache:idhash:67-0!*!*!*!*!*!* and timestamp 20150115090458 and revision id 607
 -->
## 应用提供注册应用
### **注册成为第三方应用提供商**
**注册成为应用提供商，必须输入以下信息：**
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">信息项</th><th>要求及说明</th></tr>
<tr><td> 企业Logo</td><td> 应用提供商的企业Logo，小于2M，640*240，背景为白色</td></tr>
<tr><td> 企业简称</td><td> 使用对外宣传的企业简称，能代表企业的名字，2-16个字</td></tr>
<tr><td> 企业简介</td><td> 描述企业所提供的服务，4-120个字</td></tr>
<tr><td> 企业官网</td><td> 应用服务商的企业官网</td></tr></tbody></table>
注册条件：a）拥有一个已经过认证的企业号 b）用系统管理员身份进行申请

### **创建应用套件**
开发者完成注册之后，即可创建应用套件。应用套件是第三方应用授权的主体，接口的开发都与应用套件息息相关，请开发者仔细阅读下方内容。

**基本信息：**
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">信息项</th><th>要求及说明</th></tr>
<tr><td> 套件Logo</td><td> 应用套件的Logo，小于2M，640*240，在授权页会被用于展示。</td></tr>
<tr><td> 介绍网站</td><td> 介绍该应用套件网站或者页面</td></tr>
<tr><td> 应用套件介绍</td><td> 描述该应用套件所提供的服务，4-120个字</td></tr>
<tr><td> 服务行业</td><td> 该应用套件所服务的行业对象，一个套件只能属于一个服务行业。</td></tr>
<tr><td> 套件标签</td><td> 套件提供的服务类型，如OA办公、CRM、HR、ERP等。一个套件只能拥有一个标签。</td></tr></tbody></table>
**注意：**

**1）你应谨慎选择所填写的服务行业，后续的版本中，用户可以在企业号中通过服务行业搜索相应的套件。**

**2）你可以创建或者选择其他开发者已创建的标签。你应该谨慎选择套件标签，用户往往会在企业号中通过标签查找相关联的套件。**

**开发信息：**
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">套件参数内容</th><th>说明</th></tr>
<tr><td> 发起授权域名</td><td> 在该域名下发起的授权请求才可被通过，企业点击授权链接时，企业号会检查该域名是否已登记。</td></tr>
<tr><td> 授权完成回调域名</td><td> 在第三方应用授权流程中，授权成功后会回调该域名，返回临时code。你需用此code换取永久授权码，请尽量将此域名与发起授权域名保持一致。</td></tr>
<tr><td> 服务事件接收URL</td><td> 系统将会把此套件的授权变更事件以及ticket参数推送给此URL。(ticket说明详见API接口说明)</td></tr>
<tr><td> Token</td><td> 可任意填写，用于生成签名,校验回调请求的合法性。</td></tr>
<tr><td> EncodingAESKey</td><td> 回调消息加解密参数，是AES密钥的Base64编码，用于解密回调消息内容对应的密文。</td></tr>
<tr><td> 白名单IP列表</td><td> 应用套件调用企业号第三方应用API时的合法IP列表，只有白名单内的IP才能正常调用企业号API，后续IP若有修改，需要及时进行列表更新。</td></tr>
<tr><td> 特殊权限</td><td> 是否需要通讯录管理权限，需要应用提供商在注册套件时进行申报。若勾选通讯录管理权限，在授权页会出现相应需要授权内容。如果应用不会修改企业通讯录的内容，则无需勾选。</td></tr></tbody></table>
创建完成之后，系统会告知开发者该应用套件的Suiteid和Suitesecret。<a rel="nofollow" class="external text" href="http://qydev.weixin.qq.com/wiki/index.php?title=%E7%AC%AC%E4%B8%89%E6%96%B9%E5%BA%94%E7%94%A8%E6%8E%A5%E5%8F%A3%E8%AF%B4%E6%98%8E">（详见第三方应用接口说明）</a>

### **在应用套件里添加应用**
当你创建完应用套件后，需要在套件配置应用，应用的信息填写相对简单。

**基本信息：**
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">信息项</th><th>要求及说明</th></tr>
<tr><td> 应用Logo</td><td> 应用的Logo，小于2M，640*240，在授权页会被用于展示。</td></tr>
<tr><td> 应用名称</td><td> 应用的名称，2-16个字。</td></tr>
<tr><td> 功能介绍</td><td> 描述该应用的功能与特色，4-120个字内</td></tr></tbody></table>
**开发信息：**
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">应用参数内容</th><th>说明</th></tr>
<tr><td> callbackurl</td><td> 用于接收托管企业号应用的用户消息，URL支持使用$CORPID$模板参数表示corpid，推送事件时会替换为企业的corpid。</td></tr>
<tr><td> 特殊权限</td><td> 应用的特殊权限为上报地理位置的功能开关，若在创建应用时勾选了此特殊权限，在授权会提示用户。</td></tr></tbody></table>
<!--
NewPP limit report
CPU time usage: 0.020 seconds
Real time usage: 0.026 seconds
Preprocessor visited node count: 10/1000000
Preprocessor generated node count: 16/1000000
Post‐expand include size: 0/2097152 bytes
Template argument size: 0/2097152 bytes
Highest expansion depth: 2/40
Expensive parser function count: 0/100
-->
<!-- Saved in parser cache with key db_wiki:pcache:idhash:62-0!*!*!!*!*!* and timestamp 20150116030015 and revision id 603
 -->
## 企业号管理员授权应用
### **应用提供商网站发起授权**
应用提供商构造授权链接，该链接中需要提供套件suieid、预授权码和服务事件接收的url，预授权码通过接口get_preauthcode获取<a rel="nofollow" class="external text" href="http://qydev.weixin.qq.com/wiki/index.php?title=%E7%AC%AC%E4%B8%89%E6%96%B9%E5%BA%94%E7%94%A8%E6%8E%A5%E5%8F%A3%E8%AF%B4%E6%98%8E">（详见第三方应用接口说明）</a>。

同时，为了保证授权体验的一致性，企业号也提供了统一的授权按钮样式，开发者可在<a rel="nofollow" class="external text" href="http://qydev.weixin.qq.com/wiki/index.php?title=%E8%AE%BE%E8%AE%A1%E8%B5%84%E6%BA%90%E4%B8%8B%E8%BD%BD">（设计资源下载）</a>中找到。

企业号系统管理员点击授权链接进入授权页，此时会校验授权链接所在的域名是否已被登记。

### **展示授权页面，选择授权内容**
由于应用套件里有多个应用，管理员需要选择第三方提供的应用与企业号内的应用进行关联，且企业号内的一个应用只能授权给一个第三方应用提供商。

若该应用套件在注册时申报了需要通讯录管理权限，此时在授权页会提示企业管理员，并让管理员选择部门标签或者成员进行授权。

### **完成授权操作**
企业管理员完成授权之后，页面会回跳第三方应用提供商的网站（即应用套件的授权完成回调域名），此时会带上一个临时code，开发者使用临时code通过接口query_auth换取永久授权码。<a rel="nofollow" class="external text" href="http://qydev.weixin.qq.com/wiki/index.php?title=%E7%AC%AC%E4%B8%89%E6%96%B9%E5%BA%94%E7%94%A8%E6%8E%A5%E5%8F%A3%E8%AF%B4%E6%98%8E">（详见第三方应用接口说明）</a>

授权完成后，授权方的企业号管理后台在权限管理里会出现第三方管理组类型，系统管理员也可进行取消授权操作与修改通讯录授权内容，企业取消授权与授权内容的改变也会实时通知应用提供商。

应用提供商在授权成功之后，会得到权限情况如下表所示：
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:120px"> 授权主体</th>
<th colspan="2" style="width:360px"> 权限内容</th>
<th style="width:120px"> 权限情况</th></tr>
<tr>
<th rowspan="9"> 应用套件</th>
<th rowspan="8" style="width:120px"> 应用</th><td> 应用头像</td><td> 读写</td></tr>
<tr><td> 应用简介</td><td> 读写</td></tr>
<tr><td> 地理位置上报开关</td><td> 读写（应用创建时申报）</td></tr>
<tr><td> 可信域名</td><td> 读写</td></tr>
<tr><td> 用户状态变更</td><td> 读写</td></tr>
<tr><td> 应用id</td><td> 只读</td></tr>
<tr><td> 应用可见范围</td><td> 只读</td></tr>
<tr><td> 使用应用发消息、自定义菜单、oAuth 2.0等接口</td><td> 读写</td></tr>
<tr>
<th style="width:120px"> 通讯录</th><td> 授权的通讯录（包括成员、部门、标签）</td><td> 读写</td></tr></tbody></table>
<!--
NewPP limit report
CPU time usage: 0.020 seconds
Real time usage: 0.024 seconds
Preprocessor visited node count: 10/1000000
Preprocessor generated node count: 16/1000000
Post‐expand include size: 0/2097152 bytes
Template argument size: 0/2097152 bytes
Highest expansion depth: 2/40
Expensive parser function count: 0/100
-->
<!-- Saved in parser cache with key db_wiki:pcache:idhash:63-0!*!*!!*!*!* and timestamp 20150116032244 and revision id 604
 -->
## 第三方应用接口说明
<div id="toc" class="toc"><div id="toctitle">## 目录</div>
<ul>
<li class="toclevel-1 tocsection-1"><a href="#.E7.AC.AC.E4.B8.89.E6.96.B9.E5.BA.94.E7.94.A8.E6.8E.A5.E5.8F.A3.E8.AF.B4.E6.98.8E"><span class="tocnumber">1</span> <span class="toctext"><b>第三方应用接口说明</a>
<ul>
<li class="toclevel-2 tocsection-2"><a href="#.E6.8E.A5.E5.8F.A3.E6.A6.82.E8.BF.B0"><span class="tocnumber">1.1</span> <span class="toctext"><b>接口概述</a></li>
<li class="toclevel-2 tocsection-3"><a href="#.E6.8E.88.E6.9D.83.E6.B5.81.E7.A8.8B.E8.AF.B4.E6.98.8E"><span class="tocnumber">1.2</span> <span class="toctext"><b>授权流程说明</a></li>
<li class="toclevel-2 tocsection-4"><a href="#.E8.8E.B7.E5.8F.96.E5.BA.94.E7.94.A8.E5.A5.97.E4.BB.B6.E4.BB.A4.E7.89.8C"><span class="tocnumber">1.3</span> <span class="toctext"><b>获取应用套件令牌</a></li>
<li class="toclevel-2 tocsection-5"><a href="#.E8.8E.B7.E5.8F.96.E9.A2.84.E6.8E.88.E6.9D.83.E7.A0.81"><span class="tocnumber">1.4</span> <span class="toctext"><b>获取预授权码</a></li>
<li class="toclevel-2 tocsection-6"><a href="#.E8.8E.B7.E5.8F.96.E4.BC.81.E4.B8.9A.E5.8F.B7.E7.9A.84.E6.B0.B8.E4.B9.85.E6.8E.88.E6.9D.83.E7.A0.81"><span class="tocnumber">1.5</span> <span class="toctext"><b>获取企业号的永久授权码</a></li>
<li class="toclevel-2 tocsection-7"><a href="#.E8.8E.B7.E5.8F.96.E4.BC.81.E4.B8.9A.E5.8F.B7.E7.9A.84.E6.8E.88.E6.9D.83.E4.BF.A1.E6.81.AF"><span class="tocnumber">1.6</span> <span class="toctext"><b>获取企业号的授权信息</a></li>
<li class="toclevel-2 tocsection-8"><a href="#.E8.8E.B7.E5.8F.96.E4.BC.81.E4.B8.9A.E5.8F.B7.E5.BA.94.E7.94.A8"><span class="tocnumber">1.7</span> <span class="toctext"><b>获取企业号应用</a></li>
<li class="toclevel-2 tocsection-9"><a href="#.E8.AE.BE.E7.BD.AE.E4.BC.81.E4.B8.9A.E5.8F.B7.E5.BA.94.E7.94.A8"><span class="tocnumber">1.8</span> <span class="toctext"><b>设置企业号应用</a></li>
<li class="toclevel-2 tocsection-10"><a href="#.E8.8E.B7.E5.8F.96.E4.BC.81.E4.B8.9A.E5.8F.B7access_token"><span class="tocnumber">1.9</span> <span class="toctext"><b>获取企业号access_token</a></li>
</ul></li>
</ul>
</div>
## **第三方应用接口说明**

### **接口概述**
应用提供商拥有自己的API集合，主要用于完成授权流程以及实现授权后的相关功能。
套件API（即接口）如下：
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">功能</th><th>API名称</th></tr>
<tr><td> 获取应用套件令牌</td><td> get_suite_token</td></tr>
<tr><td> 获取预授权码</td><td> get_pre_auth_code</td></tr>
<tr><td> 获取企业号的永久授权码</td><td> get_permanent_code</td></tr>
<tr><td> 获取企业号应用</td><td> get_agent</td></tr>
<tr><td> 设置企业号应用</td><td> set_agent</td></tr>
<tr><td> 获取企业号access_token</td><td> get_corp_token</td></tr>
<tr><td> 获取企业授权信息</td><td> get_auth_info</td></tr></tbody></table>
应用套件在获得企业号授权后，应用提供商可以获取企业号access_token，调用企业授权给应用套件的API，企业号API的使用方式详见

<a rel="nofollow" class="external free" href="http://qydev.weixin.qq.com/wiki/index.php?title=%E9%A6%96%E9%A1%B5">http://qydev.weixin.qq.com/wiki/index.php?title=%E9%A6%96%E9%A1%B5</a>

**特别注意，所有API调用需要验证来源IP。只有在应用套件申请信息中填写的合法IP列表内才能合法调用，其他一律拒绝。**

### **授权流程说明**

<div class="center"><div class="floatnone"><a href="/wiki/index.php?title=%E6%96%87%E4%BB%B6:Sqlc1.png" class="image"><img alt="Sqlc1.png" src="/wiki/images/thumb/8/82/Sqlc1.png/600px-Sqlc1.png" width="600" height="271" srcset="/wiki/images/8/82/Sqlc1.png 1.5x, /wiki/images/8/82/Sqlc1.png 2x"></a></div></div>

下面对其进行详细介绍：

**1）企业进入应用提供商网站**

指的是，企业系统管理员进入应用提供商网站，如www.ABC.com。

**2）获取预授权码**

预授权码是应用套件实现授权托管的安全凭证，通过suite_id，suite_secret和suite_ticket获取，相关接口为get_pre_auth_code。

**3）应用提供商引导企业系统管理员进入应用套件授权页**

应用提供商可以在自己的网站中放置“微信企业号应用授权”的入口，引导企业号管理员进入应用套件授权页。网址为:

<a rel="nofollow" class="external free" href="https://qy.weixin.qq.com/cgi-bin/loginpage?suite_id=$suite_id$&amp;pre_auth_code=$pre_auth_code$&amp;redirect_uri=$redirect_uri$&amp;state=$state$">https://qy.weixin.qq.com/cgi-bin/loginpage?suite_id=$suite_id$&amp;pre_auth_code=$pre_auth_code$&amp;redirect_uri=$redirect_uri$&amp;state=$state$</a>

该网址中应用提供商需要提供suite_id、预授权码、授权完成回调URI和state

**4）企业号管理员确认并同意授权托管给应用提供商**

企业号管理员进入套件授权页后，设置授权内容，确认并同意将自己的企业号应用或通讯录授权托管给应用提供商，完成授权流程。

**5）授权成功，返回临时授权码**

授权流程完成后，会进入回调URI，并在URI参数中返回授权码、过期时间以及state参数(redirect_uri?auth_code=xxx&amp;expires_in=1200&amp;state=xx)

**6）利用临时授权码获取永久授权码以及授权信息**

在得到临时授权码后，应用提供商可以使用临时授权码换取永久授权码以及授权信息，后续可以通过永久授权码调用企业号相关API（能调用哪些API，取决于用户将哪些权限集授权给了应用提供商）。

### **获取应用套件令牌**
该API用于获取应用套件令牌（suite_access_token）。

**注1：**由于应用提供商可能托管了大量的企业号，其安全问题造成的影响会更加严重，故API中除了合法来源IP校验之外，还额外增加了1项安全策略：

获取suite_access_token时，还额外需要suite_ticket参数**（请永远使用最新接收到的suite_ticket）**。suite_ticket由企业号后台定时推送给应用套件，并定时更新。

**注2：**通过本接口获取的accesstoken不会自动续期，每次获取都会自动更新。

**接口调用请求说明**

https请求方式: POST

<a rel="nofollow" class="external free" href="https://qyapi.weixin.qq.com/cgi-bin/service/get_suite_token">https://qyapi.weixin.qq.com/cgi-bin/service/get_suite_token</a>

**POST数据示例**
<pre>{
    "suite_id":"id_value" ,
    "suite_secret": "secret_value",
    "suite_ticket": "ticket_value"
}
</pre>
**请求参数说明**
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th><th>说明</th></tr>
<tr><td> suite_id</td><td> 应用套件id</td></tr>
<tr><td> suite_secret</td><td> 应用套件secret</td></tr>
<tr><td> suite_ticket</td><td> 微信后台推送的ticket</td></tr></tbody></table>
**返回结果示例**
<pre>{
    "suite_access_token":"61W3mEpU66027wgNZ_MhGHNQDHnFATkDa9-2llqrMBjUwxRSNPbVsMmyD-yq8wZETSoE5NQgecigDrSHkPtIYA",
    "expires_in":7200
}
</pre>

**结果参数说明**
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th><th>说明</th></tr>
<tr><td> suite_access_token</td><td> 应用套件access_token</td></tr>
<tr><td> expires_in</td><td> 有效期</td></tr></tbody></table>

### **获取预授权码**
该API用于获取预授权码。预授权码用于企业号授权时的应用提供商安全验证。

**接口调用请求说明**

https请求方式: POST

<a rel="nofollow" class="external free" href="https://qyapi.weixin.qq.com/cgi-bin/service/get_pre_auth_code?suite_access_token=xxx">https://qyapi.weixin.qq.com/cgi-bin/service/get_pre_auth_code?suite_access_token=xxx</a>

**POST数据示例**
<pre>{
    "suite_id":"id_value",
    "appid":[id1,id2,id3]
}
</pre>
**请求参数说明**
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th><th>说明</th></tr>
<tr><td> suite_id</td><td> 应用套件id</td></tr>
<tr><td> appid</td><td> 应用id，本参数选填，表示用户能对本套件内的哪些应用授权，不填时默认用户有全部授权权限</td></tr></tbody></table>
**返回结果示例**
<pre>{
    "errcode":"0" ,
    "errmsg":"ok" ,
    "pre_auth_code":"Cx_Dk6qiBE0Dmx4EmlT3oRfArPvwSQ-oa3NL_fwHM7VI08r52wazoZX2Rhpz1dEw",
    "expires_in":1200
}
</pre>
**结果参数说明**
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th><th>说明</th></tr>
<tr><td> pre_auth_code</td><td> 预授权码</td></tr>
<tr><td> expires_in</td><td> 有效期</td></tr></tbody></table>

### **获取企业号的永久授权码**
该API用于使用临时授权码换取授权方的永久授权码，并换取授权信息、企业access_token。

**注：临时授权码使用一次后即失效**

**接口调用请求说明**

https请求方式: POST

<a rel="nofollow" class="external free" href="https://qyapi.weixin.qq.com/cgi-bin/service/get_permanent_code?suite_access_token=xxxx">https://qyapi.weixin.qq.com/cgi-bin/service/get_permanent_code?suite_access_token=xxxx</a>

**POST数据示例**
<pre>{
    "suite_id":"id_value" ,
    "auth_code": "auth_code_value"
}
</pre>
**请求参数说明**
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th><th>说明</th></tr>
<tr><td> suite_id</td><td> 应用套件id</td></tr>
<tr><td> auth_code</td><td> 临时授权码会在授权成功时附加在redirect_uri中跳转回应用提供商网站。</td></tr></tbody></table>
**返回结果示例**
<pre>{
    "access_token": "xxxxxx",
    "expires_in": 7200,
    "permanent_code": "xxxx",
    "auth_corp_info":
    {
        "corpid": "xxxx",
        "corp_name": "name",
        "corp_type": "verified",
        "corp_round_logo_url": "xxxxxx",
        "corp_square_logo_url": "yyyyy",
        "corp_user_max": "50",
        "corp_agent_max": "30"
    },
    "auth_info":
    {
    "agent"&nbsp;:
        [
            {
                "agentid":"1",
                "name":"NAME",
                "square_logo_url":"xxxxxx",
                "round_logo_url":"yyyyyy",
                "appid":"1",
                "api_group":["get_location"]
            },
            {
                "agentid":"2",
                "name":"NAME2",
                "square_logo_url":"xxxxxx",
                "round_logo_url":"yyyyyy",
                "appid":"5",
                "api_group":[]
            }
        ],
        "department":
        [
            {
                "id":"2",
                "name":"PARTYNAME",
                "parentid":"1",
                "writable":"true"
            }
        ]
    }
}
</pre>

**结果参数说明**
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th><th>说明</th></tr>
<tr><td> access_token</td><td> 授权方（企业）access_token</td></tr>
<tr><td> expires_in</td><td> 授权方（企业）access_token超时时间</td></tr>
<tr><td> permanent_code</td><td> 企业号永久授权码</td></tr>
<tr><td> corp_info</td><td> 授权方企业信息</td></tr>
<tr><td> corpid</td><td> 授权方企业号id</td></tr>
<tr><td> corp_name</td><td> 授权方企业号名称</td></tr>
<tr><td> corp_type</td><td> 授权方企业号类型，认证号：verified, 注册号：unverified，体验号：test</td></tr>
<tr><td> corp_round_logo_url</td><td> 授权方企业号圆形头像</td></tr>
<tr><td> corp_square_logo_url</td><td> 授权方企业号方形头像</td></tr>
<tr><td> corp_user_max</td><td> 授权方企业号用户规模</td></tr>
<tr><td> corp_agent_max</td><td> 授权方企业号应用规模</td></tr>
<tr><td> auth_info</td><td> 授权信息</td></tr>
<tr><td> agent</td><td> 授权的应用信息</td></tr>
<tr><td> agentid</td><td> 授权方应用id</td></tr>
<tr><td> agent:name</td><td> 授权方应用名字</td></tr>
<tr><td> square_logo_url</td><td> 授权方应用方形头像</td></tr>
<tr><td> round_logo_url</td><td> 授权方应用圆形头像</td></tr>
<tr><td> appid</td><td> 服务商套件中的对应应用id</td></tr>
<tr><td> api_group</td><td> 授权方应用敏感权限组，目前仅有get_location，表示是否有权限设置应用获取地理位置的开关</td></tr>
<tr><td> department</td><td> 授权的通讯录部门</td></tr>
<tr><td> department:id</td><td> 部门id</td></tr>
<tr><td> department:name</td><td> 部门名称</td></tr>
<tr><td> department:parentid</td><td> 父部门id</td></tr>
<tr><td> department:writable</td><td> 是否具有该部门的写权限</td></tr></tbody></table>

### **获取企业号的授权信息**
该API用于通过永久授权码换取企业号的授权信息。 永久code的获取，是通过临时授权码使用get_permanent_code 接口获取到的permanent_code。

**接口调用请求说明**

https请求方式: POST

<a rel="nofollow" class="external free" href="https://qyapi.weixin.qq.com/cgi-bin/service/get_auth_info?suite_access_token=xxxx">https://qyapi.weixin.qq.com/cgi-bin/service/get_auth_info?suite_access_token=xxxx</a>

**接口调用请求说明**

https请求方式: POST

<a rel="nofollow" class="external free" href="https://qyapi.weixin.qq.com/cgi-bin/service/get_auth_info?suite_access_token=xxxx">https://qyapi.weixin.qq.com/cgi-bin/service/get_auth_info?suite_access_token=xxxx</a>

**POST数据示例**
<pre>{
    "suite_id":"suite_id_value",
    "auth_corpid": "auth_corpid_value",
    "permanent_code": "code_value"
}
</pre>
**请求参数说明**
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th><th>说明</th></tr>
<tr><td> suite_id</td><td> 应用套件id</td></tr>
<tr><td> auth_corpid</td><td> 授权方corpid</td></tr>
<tr><td> permanent_code</td><td> 永久授权码，通过get_permanent_code获取</td></tr></tbody></table>
**返回结果示例**
<pre>{
    "auth_corp_info": {
        "corpid": "xxxx",
        "corp_name": "name",
        "corp_type": "verified",
        "corp_round_logo_url": "xxxxxx",
        "corp_square_logo_url": "yyyyy",
        "corp_user_max": "50",
        "corp_agent_max": "30"
    },
    "auth_info": {
        "agent"&nbsp;: [
        {
            "agentid":"1",
            "name":"NAME",
            "round_logo_url":"xxxxxx",
            "square_logo_url ":"yyyyyy",
            "app_id":"1",
            "api_group":["get_location"]
        },
        {
            "agentid":"2",
            "name":"NAME2",
            "round_logo_url":"xxxxxx",
            "square_logo_url ":"yyyyyy",
            "app_id":"5",
            "api_group":[]
        }
        ],
        "department":[
        {
            "id":"2",
            "name":"PARTYNAME",
            "parentid":"1",
            "writable":"true"
        }
        ]
    }
}
</pre>
**结果参数说明**
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th><th>说明</th></tr>
<tr><td> auth_corp_info</td><td> 授权方企业信息</td></tr>
<tr><td> corpid</td><td> 授权方企业号id</td></tr>
<tr><td> corp_name</td><td> 授权方企业号名称</td></tr>
<tr><td> corp_type</td><td> 授权方企业号类型，认证号：verified, 注册号：unverified，体验号：test</td></tr>
<tr><td> corp_round_logo_url</td><td> 授权方企业号圆形头像</td></tr>
<tr><td> corp_square_logo_url</td><td> 授权方企业号方形头像</td></tr>
<tr><td> corp_user_max</td><td> 授权方企业号用户规模</td></tr>
<tr><td> corp_agent_max</td><td> 授权方企业号应用规模</td></tr>
<tr><td> auth_info</td><td> 授权信息</td></tr>
<tr><td> agent</td><td> 授权的应用信息</td></tr>
<tr><td> agentid</td><td> 授权方应用id</td></tr>
<tr><td> agent:name</td><td> 授权方应用名字</td></tr>
<tr><td> square_logo_url</td><td> 授权方应用方形头像</td></tr>
<tr><td> round_logo_url</td><td> 授权方应用圆形头像</td></tr>
<tr><td> appid</td><td> 服务商套件中的对应应用id</td></tr>
<tr><td> api_group</td><td> 授权方应用敏感权限组，目前仅有get_location，表示是否有权限设置应用获取地理位置的开关</td></tr>
<tr><td> department</td><td> 授权的通讯录部门</td></tr>
<tr><td> department:id</td><td> 部门id</td></tr>
<tr><td> department:name</td><td> 部门名称</td></tr>
<tr><td> department:parentid</td><td> 父部门id</td></tr>
<tr><td> department:writable</td><td> 是否具有该部门的写权限</td></tr></tbody></table>

### **获取企业号应用**
该API用于获取授权方的企业号某个应用的基本信息，包括头像、昵称、帐号类型、认证类型、可见范围等信息

**接口调用请求说明**

https请求方式: POST

```html
<a rel="nofollow" class="external free" href="https://qyapi.weixin.qq.com/cgi-bin/service/get_agent?suite_access_token=xxxx">https://qyapi.weixin.qq.com/cgi-bin/service/get_agent?suite_access_token=xxxx</a>
```

**POST数据示例**

```html
<pre>{
    "suit_id":"suit_id_value" ,
    "auth_corpid": "auth_corpid_value",
    "permanent_code": " permanent_code _value",
    "agentid ": "1"
}
</pre>
**请求参数说明**
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th><th>说明</th></tr>
<tr><td> suite_id</td><td> 应用套件id</td></tr>
<tr><td> auth_corpid</td><td> 授权方corpid</td></tr>
<tr><td> permanent_code</td><td> 永久授权码，从get_permanent_code接口中获取</td></tr>
<tr><td> agentid</td><td> 授权方应用id</td></tr></tbody></table>
**返回结果示例**
<pre>{
    "errcode":"0" ,
    "errmsg":"ok" ,
    "agentid":"1" ,
    "name":"NAME" ,
    "square_logo_url":"xxxxxxxx" ,
    "round_logo_url":"yyyyyyyy" ,
    "description":"desc" ,
    "allow_userinfos":{
        "user":[
            {
                "userid":"id1",
                "status":"1"
            },
            {
                "userid":"id2",
                "status":"1"
            },
            {
                "userid":"id3",
                "status":"1"
            }
        ]
    },
    "allow_partys": {
        "partyid": [1]
    }
    "allow_tags": {
        "tagid": [1,2,3]
    }
    "close":0 ,
    "redirect_domain":"www.qq.com",
    "report_location_flag":0,
    "isreportuser":0
}
</pre>
```

**结果参数说明**

```html
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th><th>说明</th></tr>
<tr><td> agentid</td><td> 授权方企业应用id</td></tr>
<tr><td> name</td><td> 授权方企业应用名称</td></tr>
<tr><td> square_logo_url</td><td> 授权方企业应用方形头像</td></tr>
<tr><td> round_logo_url</td><td> 授权方企业应用圆形头像</td></tr>
<tr><td> description</td><td> 授权方企业应用详情</td></tr>
<tr><td> allow_userinfos</td><td> 授权方企业应用可见范围（人员），其中包括userid和关注状态state</td></tr>
<tr><td> allow_partys</td><td> 授权方企业应用可见范围（部门）</td></tr>
<tr><td> allow_tags</td><td> 授权方企业应用可见范围（标签）</td></tr>
<tr><td> close</td><td> 授权方企业应用是否被禁用</td></tr>
<tr><td> redirect_domain</td><td> 授权方企业应用可信域名</td></tr>
<tr><td> report_location_flag</td><td> 授权方企业应用是否打开地理位置上报  0：不上报；1：进入会话上报；2：持续上报</td></tr>
<tr><td> isreportuser</td><td> 是否接收用户变更通知。0：不接收；1：接收</td></tr></tbody></table>
```

### **设置企业号应用**
该API用于设置授权方的企业应用的选项设置信息，如：地理位置上报等。注意，获取各项选项设置信息，需要有授权方的授权。

**接口调用请求说明**

https请求方式: POST

<a rel="nofollow" class="external free" href="https://qyapi.weixin.qq.com/cgi-bin/service/set_agent?suite_access_token=xxxx">https://qyapi.weixin.qq.com/cgi-bin/service/set_agent?suite_access_token=xxxx</a>

**POST数据示例**
<pre>{
    "suite_id":"id_value",
    "auth_corpid": "auth_corpid_value",
    "permanent_code ": "code_value",
    "agent":
    {
        "agentid": "5",
        "report_location_flag": "0",
        "logo_mediaid": "xxxxx",
        "name": "NAME",
        "description": "DESC",
        "redirect_domain": "xxxxxx",
        "isreportuser":0
    }
}
</pre>
**请求参数说明**
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th><th>说明</th></tr>
<tr><td> suite_id</td><td> 应用套件id</td></tr>
<tr><td> auth_corpid</td><td> 授权方corpid</td></tr>
<tr><td> permanent_code</td><td> 永久授权码，从get_permanent_code接口中获取</td></tr>
<tr><td> agent</td><td> 要设置的企业应用的信息</td></tr>
<tr><td> agentid</td><td> 企业应用的id</td></tr>
<tr><td> report_location_flag</td><td> 企业应用是否打开地理位置上报 0：不上报；1：进入会话上报；2：持续上报</td></tr>
<tr><td> logo_mediaid</td><td> 企业应用头像的mediaid，通过多媒体接口上传图片获得mediaid，上传后会自动裁剪成方形和圆形两个头像</td></tr>
<tr><td> name</td><td> 企业应用名称</td></tr>
<tr><td> description</td><td> 企业应用详情</td></tr>
<tr><td> redirect_domain</td><td> 企业应用可信域名</td></tr>
<tr><td> isreportuser</td><td> 是否接收用户变更通知。0：不接收；1：接收</td></tr></tbody></table>
**返回结果示例**
<pre>{
    "errcode":"0",
    "errmsg":"ok"
}
</pre>

### **获取企业号access_token**
应用提供商在取得企业号的永久授权码并完成对企业号应用的设置之后，便可以开始通过调用企业接口（详见企业接口文档）来运营这些应用。其中，调用企业接口所需的access_token获取方法如下。

**接口调用请求说明**

https请求方式: POST

<a rel="nofollow" class="external free" href="https://qyapi.weixin.qq.com/cgi-bin/service/get_corp_token?suite_access_token=xxxx">https://qyapi.weixin.qq.com/cgi-bin/service/get_corp_token?suite_access_token=xxxx</a>

**POST数据示例**
<pre>{
    "suite_id":"suite_id_value",
    "auth_corpid": "auth_corpid_value",
    "permanent_code": "code_value"
}
</pre>
**请求参数说明**
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th><th>说明</th></tr>
<tr><td> suite_id</td><td> 应用套件id</td></tr>
<tr><td> auth_corpid</td><td> 授权方corpid</td></tr>
<tr><td> permanent_code</td><td> 永久授权码，通过get_permanent_code获取</td></tr></tbody></table>
**返回结果示例**
<pre>{
    "access_token": "xxxxxx",
    "expires_in": 7200,
}
</pre>

**结果参数说明**
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th><th>说明</th></tr>
<tr><td> access_token</td><td> 授权方（企业）access_token</td></tr>
<tr><td> expires_in</td><td> 授权方（企业）access_token超时时间</td></tr></tbody></table>
<!--
NewPP limit report
CPU time usage: 0.184 seconds
Real time usage: 0.196 seconds
Preprocessor visited node count: 38/1000000
Preprocessor generated node count: 44/1000000
Post‐expand include size: 0/2097152 bytes
Template argument size: 0/2097152 bytes
Highest expansion depth: 2/40
Expensive parser function count: 0/100
-->
<!-- Saved in parser cache with key db_wiki:pcache:idhash:64-0!*!*!!zh-cn!2!* and timestamp 20150115100238 and revision id 638
 -->
## 第三方回调协议
### **推送suite_ticket协议**
微信服务器会向应用提供商申请时填写的套件事件接收 URL定时推送ticket：

<a rel="nofollow" class="external free" href="https://127.0.0.1/suite/receive?msg_signature=3a7b08bb8e6dbce3c9671d6fdb69d15066227608&amp;timestamp=1403610513&amp;nonce=380320359">https://127.0.0.1/suite/receive?msg_signature=3a7b08bb8e6dbce3c9671d6fdb69d15066227608&amp;timestamp=1403610513&amp;nonce=380320359</a>

**POST数据示例**
<pre>{
    &lt;xml&gt;
        &lt;SuiteId&gt;&lt;![CDATA[wxfc918a2d200c9a4c]]&gt;&lt;/SuiteId&gt;
        &lt;InfoType&gt; &lt;![CDATA[suite_ticket]]&gt;&lt;/InfoType&gt;
        &lt;TimeStamp&gt;1403610513&lt;/TimeStamp&gt;
        &lt;SuiteTicket&gt;&lt;![CDATA[asdfasfdasdfasdf]]&gt;&lt;/SuiteTicket&gt;
    &lt;/xml&gt;
}
</pre>
应用提供商在收到ticket推送后需要返回字符串success。

**字段说明**
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th><th>说明</th></tr>
<tr><td> SuiteId</td><td> 应用套件的SuiteId</td></tr>
<tr><td> InfoType</td><td> suite_ticket</td></tr>
<tr><td> TimeStamp</td><td> 时间戳</td></tr>
<tr><td> SuiteTicket</td><td> Ticket内容</td></tr></tbody></table>
为了加强安全性，postdata中的xml将使用应用套件申请时的加解密key来进行加密，具体请见

<a rel="nofollow" class="external free" href="http://qydev.weixin.qq.com/wiki/index.php?title=%E5%8A%A0%E8%A7%A3%E5%AF%86%E6%96%B9%E6%A1%88%E7%9A%84%E8%AF%A6%E7%BB%86%E8%AF%B4%E6%98%8E">http://qydev.weixin.qq.com/wiki/index.php?title=%E5%8A%A0%E8%A7%A3%E5%AF%86%E6%96%B9%E6%A1%88%E7%9A%84%E8%AF%A6%E7%BB%86%E8%AF%B4%E6%98%8E</a>

注意需要将corpid替换为suiteid，并忽略AgentID参数

### **变更授权的通知**
当授权方（即授权企业号）在企业号管理端的授权管理中，修改了对套件方的授权托管后，微信服务器会向应用提供商的套件事件接收 URL（创建套件时填写）推送变更授权通知。

<a rel="nofollow" class="external free" href="https://127.0.0.1/suite/receive?msg_signature=3a7b08bb8e6dbce3c9671d6fdb69d15066227608&amp;timestamp=1403610513&amp;nonce=380320359">https://127.0.0.1/suite/receive?msg_signature=3a7b08bb8e6dbce3c9671d6fdb69d15066227608&amp;timestamp=1403610513&amp;nonce=380320359</a>

**POST数据示例**
<pre>{
    &lt;xml&gt;
        &lt;SuiteId&gt;&lt;![CDATA[wxfc918a2d200c9a4c]]&gt;&lt;/SuiteId&gt;
        &lt;InfoType&gt;&lt;![CDATA[change_auth]]&gt;&lt;/InfoType&gt;
        &lt;TimeStamp&gt;1403610513&lt;/TimeStamp&gt;
        &lt;AuthCorpId&gt;&lt;![CDATA[wxf8b4f85f3a794e77]]&gt;&lt;/AuthCorpId&gt;
    &lt;/xml&gt;
}
</pre>
应用提供商在收到推送消息后需要返回字符串success

**字段说明**
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th><th>说明</th></tr>
<tr><td> SuiteId</td><td> 应用套件的SuiteId</td></tr>
<tr><td> InfoType</td><td> change_auth</td></tr>
<tr><td> TimeStamp</td><td> 时间戳</td></tr>
<tr><td> AuthCorpId</td><td> 授权方企业号的corpid</td></tr></tbody></table>
为了加强安全性，postdata中的xml将使用应用套件申请时的加解密key来进行加密，具体请见：

<a rel="nofollow" class="external free" href="http://qydev.weixin.qq.com/wiki/index.php?title=%E5%8A%A0%E8%A7%A3%E5%AF%86%E6%96%B9%E6%A1%88%E7%9A%84%E8%AF%A6%E7%BB%86%E8%AF%B4%E6%98%8E">http://qydev.weixin.qq.com/wiki/index.php?title=%E5%8A%A0%E8%A7%A3%E5%AF%86%E6%96%B9%E6%A1%88%E7%9A%84%E8%AF%A6%E7%BB%86%E8%AF%B4%E6%98%8E</a>

注意需要将corpid替换为suiteid，并忽略AgentID参数

### **取消授权的通知**
当授权方（即授权企业号）在企业号管理端的授权管理中，取消了对套件方的授权托管后，微信服务器会向应用提供商的套件事件接收 URL（创建套件时填写）推送取消授权通知。

<a rel="nofollow" class="external free" href="https://127.0.0.1/suite/receive?msg_signature=3a7b08bb8e6dbce3c9671d6fdb69d15066227608&amp;timestamp=1403610513&amp;nonce=380320359">https://127.0.0.1/suite/receive?msg_signature=3a7b08bb8e6dbce3c9671d6fdb69d15066227608&amp;timestamp=1403610513&amp;nonce=380320359</a>

**POST数据示例**
<pre>{
    &lt;xml&gt;
        &lt;SuiteId&gt;&lt;![CDATA[wxfc918a2d200c9a4c]]&gt;&lt;/ SuiteId&gt;
        &lt;InfoType&gt;&lt;![CDATA[cancel_auth]]&gt;&lt;/InfoType&gt;
        &lt;TimeStamp&gt;1403610513&lt;/TimeStamp&gt;
        &lt;AuthCorpId&gt;&lt;![CDATA[wxf8b4f85f3a794e77]]&gt;&lt;/AuthCorpId&gt;
    &lt;/xml&gt;
}
</pre>
应用提供商在收到推送消息后需要返回字符串success

**字段说明**
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">参数</th><th>说明</th></tr>
<tr><td> SuiteId</td><td> 应用套件的SuiteId</td></tr>
<tr><td> InfoType</td><td> cancel_auth</td></tr>
<tr><td> TimeStamp</td><td> 时间戳</td></tr>
<tr><td> AuthCorpId</td><td> 授权方企业号的corpid</td></tr></tbody></table>
为了加强安全性，postdata中的xml将使用应用套件申请时的加解密key来进行加密，具体请见

<a rel="nofollow" class="external free" href="http://qydev.weixin.qq.com/wiki/index.php?title=%E5%8A%A0%E8%A7%A3%E5%AF%86%E6%96%B9%E6%A1%88%E7%9A%84%E8%AF%A6%E7%BB%86%E8%AF%B4%E6%98%8E">http://qydev.weixin.qq.com/wiki/index.php?title=%E5%8A%A0%E8%A7%A3%E5%AF%86%E6%96%B9%E6%A1%88%E7%9A%84%E8%AF%A6%E7%BB%86%E8%AF%B4%E6%98%8E</a>

注意需要将corpid替换为suiteid，并忽略AgentID参数

<!--
NewPP limit report
CPU time usage: 0.016 seconds
Real time usage: 0.022 seconds
Preprocessor visited node count: 10/1000000
Preprocessor generated node count: 16/1000000
Post‐expand include size: 0/2097152 bytes
Template argument size: 0/2097152 bytes
Highest expansion depth: 2/40
Expensive parser function count: 0/100
-->
<!-- Saved in parser cache with key db_wiki:pcache:idhash:65-0!*!*!!*!*!* and timestamp 20150116031901 and revision id 581
 -->
## 设计资源下载
## **设计资源下载**

为了保证授权体验的一致性，企业号提供了两种授权按钮样式，开发者可通过“鼠标右键另存为”保存按钮图片

授权按钮蓝底：

大  
<div class="center"><div class="floatnone"><a href="/wiki/index.php?title=%E6%96%87%E4%BB%B6:Ld-d.png" class="image"><img alt="Ld-d.png" src="/wiki/images/7/73/Ld-d.png" width="240" height="50"></a></div></div>

中  
<div class="center"><div class="floatnone"><a href="/wiki/index.php?title=%E6%96%87%E4%BB%B6:Ld-z.png" class="image"><img alt="Ld-z.png" src="/wiki/images/b/b1/Ld-z.png" width="195" height="40"></a></div></div>

小  
<div class="center"><div class="floatnone"><a href="/wiki/index.php?title=%E6%96%87%E4%BB%B6:Ld-x.png" class="image"><img alt="Ld-x.png" src="/wiki/images/8/8d/Ld-x.png" width="168" height="30"></a></div></div>

授权按钮白底：

大  
<div class="center"><div class="floatnone"><a href="/wiki/index.php?title=%E6%96%87%E4%BB%B6:Bd-d.png" class="image"><img alt="Bd-d.png" src="/wiki/images/b/b3/Bd-d.png" width="240" height="50"></a></div></div>

中  
<div class="center"><div class="floatnone"><a href="/wiki/index.php?title=%E6%96%87%E4%BB%B6:Bd-z.png" class="image"><img alt="Bd-z.png" src="/wiki/images/9/90/Bd-z.png" width="195" height="40"></a></div></div>

小  
<div class="center"><div class="floatnone"><a href="/wiki/index.php?title=%E6%96%87%E4%BB%B6:Bd-x.png" class="image"><img alt="Bd-x.png" src="/wiki/images/c/ca/Bd-x.png" width="168" height="30"></a></div></div>

<!--
NewPP limit report
CPU time usage: 0.124 seconds
Real time usage: 0.150 seconds
Preprocessor visited node count: 2/1000000
Preprocessor generated node count: 8/1000000
Post‐expand include size: 0/2097152 bytes
Template argument size: 0/2097152 bytes
Highest expansion depth: 2/40
Expensive parser function count: 0/100
-->
<!-- Saved in parser cache with key db_wiki:pcache:idhash:66-0!*!*!*!*!2!* and timestamp 20150116032244 and revision id 582
 -->
# 附录

附录包含了企业号回调企业时加解密的详细方案、库和示例代码的下载，以及企业号api接口返回的错误码。

<!--
NewPP limit report
CPU time usage: 0.000 seconds
Real time usage: 0.001 seconds
Preprocessor visited node count: 1/1000000
Preprocessor generated node count: 4/1000000
Post‐expand include size: 0/2097152 bytes
Template argument size: 0/2097152 bytes
Highest expansion depth: 1/40
Expensive parser function count: 0/100
-->
<!-- Saved in parser cache with key db_wiki:pcache:idhash:36-0!*!*!*!*!*!* and timestamp 20150115123641 and revision id 195
 -->
## 加解密方案的详细说明
<div id="toc" class="toc"><div id="toctitle">## 目录</div>
<ul>
<li class="toclevel-1 tocsection-1"><a href="#.E5.85.B3.E4.BA.8E.E5.8A.A0.E8.A7.A3.E5.AF.86.E6.96.B9.E6.A1.88.E7.9A.84.E8.AF.A6.E7.BB.86.E8.AF.B4.E6.98.8E"><span class="tocnumber">1</span> <span class="toctext"><b>关于加解密方案的详细说明</a>
<ul>
<li class="toclevel-2 tocsection-2"><a href="#.E6.9C.AF.E8.AF.AD.E5.8F.8A.E8.AF.B4.E6.98.8E"><span class="tocnumber">1.1</span> <span class="toctext"><b>术语及说明</a></li>
<li class="toclevel-2 tocsection-3"><a href="#.E6.B6.88.E6.81.AF.E4.BD.93.E7.AD.BE.E5.90.8D"><span class="tocnumber">1.2</span> <span class="toctext"><b>消息体签名</a></li>
<li class="toclevel-2 tocsection-4"><a href="#.E5.8A.A0.E8.A7.A3.E5.AF.86.E6.96.B9.E6.A1.88.E8.AF.B4.E6.98.8E"><span class="tocnumber">1.3</span> <span class="toctext"><b>加解密方案说明</a></li>
</ul></li>
</ul>
</div>
## **关于加解密方案的详细说明**
### **术语及说明**
开启回调模式时，有以下术语需要了解：

1.msg_signature是签名，用于验证调用者的合法性。具体算法见以下'消息体签名'章节

2.EncodingAESKey用于消息体的加密，长度固定为43个字符，从a-z, A-Z, 0-9共62个字符中选取，是AESKey的Base64编码。解码后即为32字节长的AESKey

3.AESKey=Base64_Decode(EncodingAESKey + “=”)，是AES算法的密钥，长度为32字节。AES采用CBC模式，数据采用PKCS#7填充；IV初始向量大小为16字节，取AESKey前16字节。具体详见：<a rel="nofollow" class="external text" href="//http://tools.ietf.org/html/rfc2315">http://tools.ietf.org/html/rfc2315</a>

4.msg为消息体明文，格式为XML

5.msg_encrypt = Base64_Encode( AES_Encrypt[random(16B) + msg_len(4B) + msg + $CorpID] )，是对明文消息msg加密处理后的Base64编码。其中random为16字节的随机字符串；msg_len为4字节的msg长度，网络字节序；msg为消息体明文；$CorpID为企业号的标识

### **消息体签名**
为了验证调用者的合法性，微信在回调url中增加了消息签名，以参数msg_signature标识，企业需要验证此参数的正确性后再解密。验证步骤：

1.企业计算签名：dev_msg_signature=sha1(sort(token、timestamp、nonce、msg_encrypt))。sort的含义是将参数按照字母字典排序，然后从小到大拼接成一个字符串

2.比较dev_msg_signature和msg_signature是否相等，相等则表示验证通过

在被动响应消息时，企业同样需要用如上方法生成签名并传给微信
### **加解密方案说明**
<ul>
<li>对明文msg加密的过程如下：</li>
</ul>
msg_encrypt = Base64_Encode( AES_Encrypt[random(16B) + msg_len(4B) + msg + $CorpID] )

AES加密的buf由16个字节的随机字符串、4个字节的msg长度、明文msg和$CorpID组成。其中msg_len为msg的字节数，网络字节序；$CorpID为企业号的CorpID。经AESKey加密后，再进行Base64编码，即获得密文msg_encrypt。
<ul>
<li>对应于加密方案，解密方案如下：</li>
</ul>
1.对密文BASE64解码：aes_msg=Base64_Decode(msg_encrypt)

2.使用AESKey做AES解密：rand_msg=AES_Decrypt(aes_msg)

3.验证解密后$CorpID、msg_len

4.去掉rand_msg头部的16个随机字节，4个字节的msg_len,和尾部的$CorpID即为最终的消息体原文msg

<!--
NewPP limit report
CPU time usage: 0.008 seconds
Real time usage: 0.011 seconds
Preprocessor visited node count: 17/1000000
Preprocessor generated node count: 26/1000000
Post‐expand include size: 0/2097152 bytes
Template argument size: 0/2097152 bytes
Highest expansion depth: 2/40
Expensive parser function count: 0/100
-->
<!-- Saved in parser cache with key db_wiki:pcache:idhash:47-0!*!*!!zh-cn!*!* and timestamp 20150115123643 and revision id 450
 -->
## 加解密库下载与返回码
### **加解密库的返回码**
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">返回码</th><th>说明</th></tr>
<tr><td> 0</td><td> 请求成功</td></tr>
<tr><td> -40001</td><td> 签名验证错误</td></tr>
<tr><td> -40002</td><td> xml解析失败</td></tr>
<tr><td> -40003</td><td> sha加密生成签名失败</td></tr>
<tr><td> -40004</td><td> AESKey 非法</td></tr>
<tr><td> -40005</td><td> corpid 校验错误</td></tr>
<tr><td> -40006</td><td> AES 加密失败</td></tr>
<tr><td> -40007</td><td> AES 解密失败</td></tr>
<tr><td> -40008</td><td> 解密后得到的buffer非法</td></tr>
<tr><td> -40009</td><td> base64加密失败</td></tr>
<tr><td> -40010</td><td> base64解密失败</td></tr>
<tr><td> -40011</td><td> 生成xml失败</td></tr></tbody></table>

### **加解密库下载及示例**
<ul>
<li><a rel="nofollow" class="external text" href="//qydev.weixin.qq.com/c++.zip">c++库(9月22日更新,点击下载)</a></li>
</ul>
注意事项：

1.WXBizMsgCrypt.h声明了WXBizMsgCrypt类，提供用户接入企业微信的三个接口。WXBizMsgCrypt.cpp文件提供了三个接口的实现。Sample.cpp文件提供了如何使用这三个接口的示例。

2.WXBizMsgCrypt类封装了VerifyURL, DecryptMsg, EncryptMsg三个接口，分别用于开发者验证回调url，收到用户回复消息的解密以及开发者回复消息的加密过程。使用方法可以参考Sample.cpp文件。

3.加解密协议请参考企业微信官方文档。

4.加解密过程使用了开源的openssl和tinyxml2库，请开发者自行安装之后使用。

&nbsp;&nbsp;&nbsp;*openssl的版本号是openssl-1.0.1h，<a rel="nofollow" class="external free" href="http://www.openssl.org/">http://www.openssl.org/</a>

&nbsp;&nbsp;&nbsp;*tinyxml2的版本号是tinyxml2-2.1.0，<a rel="nofollow" class="external free" href="https://github.com/leethomason/tinyxml2">https://github.com/leethomason/tinyxml2</a>

<ul>
<li><a rel="nofollow" class="external text" href="//qydev.weixin.qq.com/python.zip">python库(9月22日更新,点击下载)</a></li>
</ul>
注意事项：

1.WXBizMsgCrypt.py文件封装了WXBizMsgCrypt接口类，提供了用户接入企业微信的三个接口，Sample.py文件提供了如何使用这三个接口的示例，ierror.py提供了错误码。

2.WXBizMsgCrypt封装了VerifyURL, DecryptMsg, EncryptMsg三个接口，分别用于开发者验证回调url、接收消息的解密以及开发者回复消息的加密过程。使用方法可以参考Sample.py文件。

3.本代码用到了pycrypto第三方库，请开发者自行安装此库再使用。

<ul>
<li><a rel="nofollow" class="external text" href="//qydev.weixin.qq.com/php.zip">php库(9月25日更新,点击下载)</a></li>
</ul>
注意事项：

1.WXBizMsgCrypt.php文件提供了WXBizMsgCrypt类的实现，是用户接入企业微信的接口类。Sample.php提供了示例以供开发者参考。errorCode.php, pkcs7Encoder.php, sha1.php, xmlparse.php文件是实现这个类的辅助类，开发者无须关心其具体实现。

2.WXBizMsgCrypt类封装了VerifyURL, DecryptMsg, EncryptMsg三个接口，分别用于开发者验证回调url、接收消息的解密以及开发者回复消息的加密过程。使用方法可以参考Sample.php文件。

<ul>
<li><a rel="nofollow" class="external text" href="//qydev.weixin.qq.com/java.zip">java库(9月24日更新,点击下载)</a></li>
</ul>
注意事项：

1.com\qq\weixin\mp\aes目录下是用户需要用到的接入企业微信的接口，其中WXBizMsgCrypt.java文件提供的WXBizMsgCrypt类封装了用户接入企业微信的三个接口，其它的类文件用户用于实现加解密，用户无须关心。sample.java文件提供了接口的使用示例。

2.WXBizMsgCrypt封装了VerifyURL, DecryptMsg, EncryptMsg三个接口，分别用于开发者验证回调url、接收消息的解密以及开发者回复消息的加密过程。使用方法可以参考Sample.java文件。

3.请开发者使用jdk1.6或以上的版本。针对org.apache.commons.codec.binary.Base64，需要导入jar包commons-codec-1.9（或comm ons-codec-1.8等其他版本），我们有提供，官方下载地址：

<a rel="nofollow" class="external free" href="http://commons.apache.org/proper/commons-codec/download_codec.cgi">http://commons.apache.org/proper/commons-codec/download_codec.cgi</a>

4.异常java.security.InvalidKeyException:illegal Key Size的解决方案：

在官方网站下载JCE无限制权限策略文件（请到官网下载对应的版本， 例如JDK7的下载地址：<a rel="nofollow" class="external free" href="http://www.oracle.com/technetwork/java/javase/downloads/jce-7-download-432124.html">http://www.oracle.com/technetwork/java/javase/downloads/jce-7-download-432124.html</a> )：

下载后解压，可以看到local_policy.jar和US_export_policy.jar以及readme.txt。如果安装了JRE，将两个jar文件放到%JRE_HOME% \lib\security目录下覆盖原来的文件，如果安装了JDK，将两个jar文件放到%JDK_HOME%\jre\lib\security目录下覆盖原来文件。

<ul>
<li><a rel="nofollow" class="external text" href="//qydev.weixin.qq.com/csharp.zip">c#库(9月22日更新,点击下载)</a></li>
</ul>
注意事项：

1.Cryptography.cs文件封装了AES加解密过程，用户无须关心具体实现。WXBizMsgCrypt.cs文件提供了用户接入企业微信的三个接口，Sample.cs文件提供了如何使用这三个接口的示例。

2.WXBizMsgCrypt.cs封装了VerifyURL, DecryptMsg, EncryptMsg三个接口，分别用于开发者验证回调url、接收消息的解密以及开发者回复消息的加密过程。使用方法可以参考Sample.cs文件。

<!--
NewPP limit report
CPU time usage: 0.024 seconds
Real time usage: 0.028 seconds
Preprocessor visited node count: 9/1000000
Preprocessor generated node count: 16/1000000
Post‐expand include size: 0/2097152 bytes
Template argument size: 0/2097152 bytes
Highest expansion depth: 2/40
Expensive parser function count: 0/100
-->
<!-- Saved in parser cache with key db_wiki:pcache:idhash:50-0!*!*!!*!*!* and timestamp 20150116031151 and revision id 537
 -->
## 全局返回码说明

企业号每次调用接口时，可能获得正确或错误的返回码，企业可以根据返回码信息调试接口，排查错误。

```
**全局返回码说明如下：**
<table border="1" cellspacing="0" cellpadding="4" align="center" width="640px">
<tbody><tr>
<th style="width:240px">返回码</th><th>说明</th></tr>
<tr><td> -1</td><td> 系统繁忙</td></tr>
<tr><td> 0</td><td> 请求成功</td></tr>
<tr><td> 40001</td><td> 获取access_token时Secret错误，或者access_token无效</td></tr>
<tr><td> 40002</td><td> 不合法的凭证类型</td></tr>
<tr><td> 40003</td><td> 不合法的UserID</td></tr>
<tr><td> 40004</td><td> 不合法的媒体文件类型</td></tr>
<tr><td> 40005</td><td> 不合法的文件类型</td></tr>
<tr><td> 40006</td><td> 不合法的文件大小</td></tr>
<tr><td> 40007</td><td> 不合法的媒体文件id</td></tr>
<tr><td> 40008</td><td> 不合法的消息类型</td></tr>
<tr><td> 40013</td><td> 不合法的corpid</td></tr>
<tr><td> 40014</td><td> 不合法的access_token</td></tr>
<tr><td> 40015</td><td> 不合法的菜单类型</td></tr>
<tr><td> 40016</td><td> 不合法的按钮个数</td></tr>
<tr><td> 40017</td><td> 不合法的按钮类型</td></tr>
<tr><td> 40018</td><td> 不合法的按钮名字长度</td></tr>
<tr><td> 40019</td><td> 不合法的按钮KEY长度</td></tr>
<tr><td> 40020</td><td> 不合法的按钮URL长度</td></tr>
<tr><td> 40021</td><td> 不合法的菜单版本号</td></tr>
<tr><td> 40022</td><td> 不合法的子菜单级数</td></tr>
<tr><td> 40023</td><td> 不合法的子菜单按钮个数</td></tr>
<tr><td> 40024</td><td> 不合法的子菜单按钮类型</td></tr>
<tr><td> 40025</td><td> 不合法的子菜单按钮名字长度</td></tr>
<tr><td> 40026</td><td> 不合法的子菜单按钮KEY长度</td></tr>
<tr><td> 40027</td><td> 不合法的子菜单按钮URL长度</td></tr>
<tr><td> 40028</td><td> 不合法的自定义菜单使用员工</td></tr>
<tr><td> 40029</td><td> 不合法的oauth_code</td></tr>
<tr><td> 40031</td><td> 不合法的UserID列表</td></tr>
<tr><td> 40032</td><td> 不合法的UserID列表长度</td></tr>
<tr><td> 40033</td><td> 不合法的请求字符，不能包含\uxxxx格式的字符</td></tr>
<tr><td> 40035</td><td> 不合法的参数</td></tr>
<tr><td> 40038</td><td> 不合法的请求格式</td></tr>
<tr><td> 40039</td><td> 不合法的URL长度</td></tr>
<tr><td> 40040</td><td> 不合法的插件token</td></tr>
<tr><td> 40041</td><td> 不合法的插件id</td></tr>
<tr><td> 40042</td><td> 不合法的插件会话</td></tr>
<tr><td> 40048</td><td> url中包含不合法domain</td></tr>
<tr><td> 40054</td><td> 不合法的子菜单url域名</td></tr>
<tr><td> 40055</td><td> 不合法的按钮url域名</td></tr>
<tr><td> 40056</td><td> 不合法的agentid</td></tr>
<tr><td> 40057</td><td> 不合法的callbackurl</td></tr>
<tr><td> 40058</td><td> 不合法的红包参数</td></tr>
<tr><td> 40059</td><td> 不合法的上报地理位置标志位</td></tr>
<tr><td> 40060</td><td> 设置上报地理位置标志位时没有设置callbackurl</td></tr>
<tr><td> 40061</td><td> 设置应用头像失败</td></tr>
<tr><td> 40062</td><td> 不合法的应用模式</td></tr>
<tr><td> 40063</td><td> 红包参数为空</td></tr>
<tr><td> 40064</td><td> 管理组名字已存在</td></tr>
<tr><td> 40065</td><td> 不合法的管理组名字长度</td></tr>
<tr><td> 40066</td><td> 不合法的部门列表</td></tr>
<tr><td> 40067</td><td> 标题长度不合法</td></tr>
<tr><td> 40068</td><td> 不合法的标签ID</td></tr>
<tr><td> 40069</td><td> 不合法的标签ID列表</td></tr>
<tr><td> 40070</td><td> 列表中所有标签（用户）ID都不合法</td></tr>
<tr><td> 40071</td><td> 不合法的标签名字，标签名字已经存在</td></tr>
<tr><td> 40072</td><td> 不合法的标签名字长度</td></tr>
<tr><td> 40073</td><td> 不合法的openid</td></tr>
<tr><td> 40074</td><td> news消息不支持指定为高保密消息</td></tr>
<tr><td> 40077</td><td> 不合法的预授权码</td></tr>
<tr><td> 40078</td><td> 不合法的临时授权码</td></tr>
<tr><td> 40079</td><td> 不合法的授权信息</td></tr>
<tr><td> 40080</td><td> 不合法的suitesecret</td></tr>
<tr><td> 40082</td><td> 不合法的suitetoken</td></tr>
<tr><td> 40083</td><td> 不合法的suiteid</td></tr>
<tr><td> 40084</td><td> 不合法的永久授权码</td></tr>
<tr><td> 40085</td><td> 不合法的suiteticket</td></tr>
<tr><td> 40086</td><td> 不合法的第三方应用appid</td></tr>
<tr><td> 41001</td><td> 缺少access_token参数</td></tr>
<tr><td> 41002</td><td> 缺少corpid参数</td></tr>
<tr><td> 41003</td><td> 缺少refresh_token参数</td></tr>
<tr><td> 41004</td><td> 缺少secret参数</td></tr>
<tr><td> 41005</td><td> 缺少多媒体文件数据</td></tr>
<tr><td> 41006</td><td> 缺少media_id参数</td></tr>
<tr><td> 41007</td><td> 缺少子菜单数据</td></tr>
<tr><td> 41008</td><td> 缺少oauth code</td></tr>
<tr><td> 41009</td><td> 缺少UserID</td></tr>
<tr><td> 41010</td><td> 缺少url</td></tr>
<tr><td> 41011</td><td> 缺少agentid</td></tr>
<tr><td> 41012</td><td> 缺少应用头像mediaid</td></tr>
<tr><td> 41013</td><td> 缺少应用名字</td></tr>
<tr><td> 41014</td><td> 缺少应用描述</td></tr>
<tr><td> 41015</td><td> 缺少Content</td></tr>
<tr><td> 41016</td><td> 缺少标题</td></tr>
<tr><td> 41017</td><td> 缺少标签ID</td></tr>
<tr><td> 41018</td><td> 缺少标签名字</td></tr>
<tr><td> 41021</td><td> 缺少suiteid</td></tr>
<tr><td> 41022</td><td> 缺少suitetoken</td></tr>
<tr><td> 41023</td><td> 缺少suiteticket</td></tr>
<tr><td> 41024</td><td> 缺少suitesecret</td></tr>
<tr><td> 41025</td><td> 缺少永久授权码</td></tr>
<tr><td> 42001</td><td> access_token超时</td></tr>
<tr><td> 42002</td><td> refresh_token超时</td></tr>
<tr><td> 42003</td><td> oauth_code超时</td></tr>
<tr><td> 42004</td><td> 插件token超时</td></tr>
<tr><td> 42007</td><td> 预授权码失效</td></tr>
<tr><td> 42008</td><td> 临时授权码失效</td></tr>
<tr><td> 42009</td><td> suitetoken失效</td></tr>
<tr><td> 43001</td><td> 需要GET请求</td></tr>
<tr><td> 43002</td><td> 需要POST请求</td></tr>
<tr><td> 43003</td><td> 需要HTTPS</td></tr>
<tr><td> 43004</td><td> 需要接收者关注</td></tr>
<tr><td> 43005</td><td> 需要好友关系</td></tr>
<tr><td> 43006</td><td> 需要订阅</td></tr>
<tr><td> 43007</td><td> 需要授权</td></tr>
<tr><td> 43008</td><td> 需要支付授权</td></tr>
<tr><td> 43009</td><td> 需要员工已关注</td></tr>
<tr><td> 43010</td><td> 需要处于回调模式</td></tr>
<tr><td> 43011</td><td> 需要企业授权</td></tr>
<tr><td> 44001</td><td> 多媒体文件为空</td></tr>
<tr><td> 44002</td><td> POST的数据包为空</td></tr>
<tr><td> 44003</td><td> 图文消息内容为空</td></tr>
<tr><td> 44004</td><td> 文本消息内容为空</td></tr>
<tr><td> 45001</td><td> 多媒体文件大小超过限制</td></tr>
<tr><td> 45002</td><td> 消息内容超过限制</td></tr>
<tr><td> 45003</td><td> 标题字段超过限制</td></tr>
<tr><td> 45004</td><td> 描述字段超过限制</td></tr>
<tr><td> 45005</td><td> 链接字段超过限制</td></tr>
<tr><td> 45006</td><td> 图片链接字段超过限制</td></tr>
<tr><td> 45007</td><td> 语音播放时间超过限制</td></tr>
<tr><td> 45008</td><td> 图文消息超过限制</td></tr>
<tr><td> 45009</td><td> 接口调用超过限制</td></tr>
<tr><td> 45010</td><td> 创建菜单个数超过限制</td></tr>
<tr><td> 45015</td><td> 回复时间超过限制</td></tr>
<tr><td> 45016</td><td> 系统分组，不允许修改</td></tr>
<tr><td> 45017</td><td> 分组名字过长</td></tr>
<tr><td> 45018</td><td> 分组数量超过上限</td></tr>
<tr><td> 45024</td><td> 账号数量超过上限</td></tr>
<tr><td> 46001</td><td> 不存在媒体数据</td></tr>
<tr><td> 46002</td><td> 不存在的菜单版本</td></tr>
<tr><td> 46003</td><td> 不存在的菜单数据</td></tr>
<tr><td> 46004</td><td> 不存在的员工</td></tr>
<tr><td> 47001</td><td> 解析JSON/XML内容错误</td></tr>
<tr><td> 48002</td><td> Api禁用</td></tr>
<tr><td> 48003</td><td> suitetoken无效</td></tr>
<tr><td> 48004</td><td> 授权关系无效</td></tr>
<tr><td> 50001</td><td> redirect_uri未授权</td></tr>
<tr><td> 50002</td><td> 员工不在权限范围</td></tr>
<tr><td> 50003</td><td> 应用已停用</td></tr>
<tr><td> 50004</td><td> 员工状态不正确（未关注状态）</td></tr>
<tr><td> 50005</td><td> 企业已禁用</td></tr>
<tr><td> 60001</td><td> 部门长度不符合限制</td></tr>
<tr><td> 60002</td><td> 部门层级深度超过限制</td></tr>
<tr><td> 60003</td><td> 部门不存在</td></tr>
<tr><td> 60004</td><td> 父亲部门不存在</td></tr>
<tr><td> 60005</td><td> 不允许删除有成员的部门</td></tr>
<tr><td> 60006</td><td> 不允许删除有子部门的部门</td></tr>
<tr><td> 60007</td><td> 不允许删除根部门</td></tr>
<tr><td> 60008</td><td> 部门名称已存在</td></tr>
<tr><td> 60009</td><td> 部门名称含有非法字符</td></tr>
<tr><td> 60010</td><td> 部门存在循环关系</td></tr>
<tr><td> 60011</td><td> 管理员权限不足，（user/department/agent）无权限</td></tr>
<tr><td> 60012</td><td> 不允许删除默认应用</td></tr>
<tr><td> 60013</td><td> 不允许关闭应用</td></tr>
<tr><td> 60014</td><td> 不允许开启应用</td></tr>
<tr><td> 60015</td><td> 不允许修改默认应用可见范围</td></tr>
<tr><td> 60016</td><td> 不允许删除存在成员的标签</td></tr>
<tr><td> 60017</td><td> 不允许设置企业</td></tr>
<tr><td> 60019</td><td> 不允许设置应用地理位置上报开关</td></tr>
<tr><td> 60020</td><td> 访问ip不在白名单之中</td></tr>
<tr><td> 60102</td><td> UserID已存在</td></tr>
<tr><td> 60103</td><td> 手机号码不合法</td></tr>
<tr><td> 60104</td><td> 手机号码已存在</td></tr>
<tr><td> 60105</td><td> 邮箱不合法</td></tr>
<tr><td> 60106</td><td> 邮箱已存在</td></tr>
<tr><td> 60107</td><td> 微信号不合法</td></tr>
<tr><td> 60108</td><td> 微信号已存在</td></tr>
<tr><td> 60109</td><td> QQ号已存在</td></tr>
<tr><td> 60110</td><td> 部门个数超出限制</td></tr>
<tr><td> 60111</td><td> UserID不存在</td></tr>
<tr><td> 60112</td><td> 成员姓名不合法</td></tr>
<tr><td> 60113</td><td> 身份认证信息（微信号/手机/邮箱）不能同时为空</td></tr>
<tr><td> 60114</td><td> 性别不合法</td></tr>
<tr><td> 60023</td><td> 应用已授权予第三方，不允许通过分级管理员修改菜单</td></tr></tbody></table>
<!--
NewPP limit report
CPU time usage: 0.020 seconds
Real time usage: 0.019 seconds
Preprocessor visited node count: 1/1000000
Preprocessor generated node count: 4/1000000
Post‐expand include size: 0/2097152 bytes
Template argument size: 0/2097152 bytes
Highest expansion depth: 1/40
Expensive parser function count: 0/100
-->
<!-- Saved in parser cache with key db_wiki:pcache:idhash:25-0!*!*!*!*!*!* and timestamp 20150116090425 and revision id 683
 -->
```
