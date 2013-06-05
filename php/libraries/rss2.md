# Feed 订阅

采用 Rss 2.0 协议

## 规范

~~~{.xml}
	<rss version="2.0">
　　<channel>
　　<title>网站标题</title>
　　<link>网站首页地址</link>
　　<description>描述</description>
　　<copyright>授权信息</copyright>
　　<language>使用的语言（zh-cn表示简体中文）</language>
　　<pubDate>发布的时间</pubDate>
　　<lastBuildDate>最后更新的时间</lastBuildDate>
　　<generator>生成器</generator>
　　<item>
　　<title>标题</title>
　　<link>链接地址</link>
　　<description>内容简要描述</description>
　　<pubDate>发布时间</pubDate>
　　<category>所属目录</category>
　　<author>作者</author>
　　</item>
　　</channel>
　　</rss>
~~~



## 文件结构

	libraries/Rss2.php          // 输出处理类
	static/css/feed.css         // 样式
	App/controllers/feed.php    // 控制器
	App/views/feed_xsl.php      // xsl 文件

> 其中 `feed_xsl.php` 可以分离为独立的 xsl 文件，注意要和 feed 地址保持同一域名。

## Rss2.php

	$this->load->library('Rss2');

### set_xsl

设定xsl模板地址

	set_xsl($xsl);

### set_cannel

设定订阅信息

	set_channel($title,$link,$description,$copyright,$pubDate,$lastBulidDate,$language='zh-cn')

参数为：

* 标题
* feed URL
* 简介
* CopyRight 
* 更新日期
* 最后生成日期
* 语言 (默认 'zh-cn')


### item_add

添加信息

	item_add($title,$link,$pubDate,$author,$description)

参数为

* 标题
* 文章链接
* 跟新日期时间
* 作者
* 内容简介

### headers

	headers()

输出文件头

	Content-Type: application/xml

### render

生成并返回xml文件字符串流。

	render();

### timestamp2pubDate

将时间戳转化为	`pubDate` 格式。

	timestamp2pubDate($timestamp)

## XSL 

如果使用模板文件 ,则输出时

	header('Content-Type: application/xslt');

注意修改 `views/feed_xsl.php` 中的 css 路径

## 实例

~~~{.php}
	/**
	 * Class Feed extends CI_Controller 
	 * @author 
	 */
	class Feed extends CI_Controller
	{
		public function index()
		{
			$this->load->library('Rss2');
			$this->rss2->set_xsl(site_url('demo/xsl'));
			$this->rss2->set_channel(
				'信息测试',               // title
				current_url(),            // url
				'简介',                   // description
				'所属权',                 // copyright
				date('D, d M Y H:i:s O'), // pubDate
				date('D, d M Y H:i:s O'), // lastBulidDate
				'zh-cn'                   // 可忽略
			);

			$this->load->model('demo_model','demo');
			$data = $this->demo->getlimit(10);

			foreach ($data as $v){
				$this->rss2->item_add(
					$v['title'],              // title
					ADMINER_URL.$v['id'],     // link
					date('D, d M Y H:i:s O'), // pubDate
					'adminer',                // author
					$v['text']             // description
				);
			}
			header( $this->rss2->headers());
			echo $this->rss2->render();
		}

		public function xsl()
		{
			header('Content-Type: application/xslt');
			$data['feed_url'] = site_url('demo/feed');
			$this->load->view('feed_xsl',$data);
		}
~~~


