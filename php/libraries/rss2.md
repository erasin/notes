# RSS



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

# RSS Library for CodeIgniter

***

## About

I needed to make an RSS feed on one of my projects but couldn't find a solution that I really liked so I built this.  Hopefully someone else can find it handy too.

## Installation

Copy **libraries/rss2.php** to libraries/ in your application folder.

## Initialization

Load the library in whatever controllers you want to use it using:

`$this->load->library('rss2');`

## Usage

RSS feeds are made up of *channels* and *items*.  In this library, each of those are objects and each has their own properties.

### Channels

Create a new channel with:

`$channel = $this->rss2->new_channel();`

You can set properties of the channel using functions like:

	$channel->atom_link(current_url());
	$channel->set_title("Channel Title");
	$channel->set_link(current_url());
	$channel->set_description("Channel Description");

You can set any attribute of the channel using:

`$channel->set_attribute($key, $value, $additional_attributes=false);`

The *$additional_attributes* param allows you to set an additional string within the rendered XML tag.  For instance, the *set_attribute* function would roughly produce the following:

`<key $additional_attributes>value</key>`

### Items

We get a new item by calling

`$item = $channel->new_item();`

You can set item properties by using functions such as

	$item->set_title("Item Title");
	$item->set_link("http://www.example.com/id");
	$item->set_guid("item_guid");
	$item->set_description("Item Description");
	$item->set_author("author.name@example.com (Author Name)");

You can also set any attribute of the item using:

`$item->set_attribute($key, $value, $additional_attributes=false);`

This will be rendered similarly to the equivilent channel function.

We add the item to the channel using

`$channel->add_item($item);`

### Channel Images

You can add images to an entire channel as well.  

**Note:** To generate valid feeds, you should define your channel image before your items.

To start, do

`$image = $channel->new_image();`

You can set various properties of the imag using functions such as

	$image->set_url("http://path-to-image");
	$image->set_title("Image Title");
	$image->set_link("http://www.example.com");
	$image->set_width("100");
	$image->set_width("100");

You add the image to the channel in the same way that you added items.

`$channel->add_item($image);`

### Rendering the Feed

After you've created your channel, channel image, and items you pack them pack into the rss2 object.

`$this->rss2->pack($channel);`

You can specify the response headers using the following

`header($this->rss2->headers());`

And finally, output the feed:

	echo $this->rss2->render();
	exit();


<https://zh.wikipedia.org/wiki/Atom_(%E6%A8%99%E6%BA%96)>

### 实例

~~~{.php}
	public function index()
	{
		$this->load->model('demo_model','demo');
		$data = $this->demo->getlimit(10);

		$this->load->library('rss2');
		# 头部
		$channel = $this->rss2->new_channel();
		$channel->atom_link(current_url());
		$channel->set_title("Channel Title");
		$channel->set_link(current_url());
		$channel->set_description("Channel Description");
		$channel->set_attribute('copyright','归属');
		$channel->set_attribute('language','zh-cn');
		$channel->set_attribute('pubDate',date('Y-m-d'));
		$channel->set_attribute('lastBuildDate',date('Y-m-d H:i:s'));
		$channel->set_attribute('generator','rss 2.0');

		# list item
		$item = $channel->new_item();

		foreach ($data as $v){
			$item->set_title($v['title']);
			$item->set_link(ADMINER_URL."/".$v['id']);
			$item->set_guid($v['id']);
			$item->set_description($v['text']);
			$item->set_author("author.name@example.com (Author Name)");
			$channel->add_item($item);
		}

		$this->rss2->pack($channel);
		header($this->rss2->headers());
		echo $this->rss2->render();
		exit();

		#http://www.google.com/intl/zh-cn/webmasters/add.html
	}
~~~



