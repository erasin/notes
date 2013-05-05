# Session 类
**title:** codeingniter Session类库  
**tags:** codeingniter,php,session  

Session 类可以使用户在浏览您的网站时，维持他们的状态并跟踪他们的行为。 Session 类将每个用户的 session 信息序列化（serialize）后存储到到 cookie 中（并同时进行加密）。 您还可以将 session 数据存储到数据库中来增强安全性，但是这时要求存储在用户 cookie 中的 session ID 值能与数据库中存储的用户 session ID 值相匹配。程序默认只在 cookie 中存储 session。如果您在要在数据库中存储 session 的话，需要按照下面指示的方法，在您的数据库中创建需要的数据表。 

**注意：** Session类并不使用PHP本身的session，而是使用类自己的session，这样做，可以给开发者提供更大的弹性。

**Note:** 即使没有使用加密会话，你也需要在配置文件里设置一个[加密密钥](encryption)。这将有助于防止伪造会话数据。

## 初始化 Session

Sessions会在每个页面载入后开始运行，所以session类必须首先被[初始化](../general/libraries)。您可以在[控制器](../general/controllers)中初始化，也可以在系统中[自动加载](../general/autoloader)（译者注：在autoload.php设定）。session类的绝大部分都会在后台运行，所以初始化session时，它session数据会被自动读取、创建和更新。

要在您的控制器构造函数中初始化session类，您可以使用 `$this->load->library` 函数:

    $this->load->library('session'); 

一旦被载入, session就可以这样使用： `$this->session`

## Sessions 是怎样工作的？

当页面载入后，session类就会检查用户的cookie中是否存在有效的session数据。如果session数据不存在（或者已经过期），那么就会创建一个新的session并把他保存在cookie中。如果session数据存在，那么他的信息就会被更新，同时cookie也会被同时更新。每次更新都会重新生成`session_id`的值。

对于您来说，需要知道的非常重要的一点就是，session类一旦被初始化，它就会自动运行。对于后面的事情，您可以完全不作理会。正如您将会在下面看到的一样，您可以正常使用session来工作，甚至还可以添加自己的session数据，而在这一切的过程中，读、写和更新的操作都是自动完成的。

## Session 数据是什么？

一个 *session* 是由一个包括下列信息的数组组成的：

*   唯一的用户Session ID (这是一个平均信息量统计出来的非常坚固的随机字符串，使用MD5加密，默认是每五分钟就重新生成一次。
*   用户的 IP 地址
*   用户浏览器信息（取前120个字符）
*   最新的一个活跃时间戳.

以上数据将会用以下数组格式序列化并存到cookie里：

    [array]  
    (  
         'session_id'    => random hash,  
         'ip_address'    => 'string - user IP address',  
         'user_agent'    => 'string - user agent data',  
         'last_activity' => timestamp  
    ) 

如果你将加密设置开启，serialized 的数组会先被加密，然后存入cookie中。这会让数据不容易被看到和修改，从而提高安全性。从[这里](encryption)可以找到更多关于加密的信息。Session类会自动负责初始化和数据加密。

注意: 默认情况下, Session Cookie 每隔 5 分钟才会更新一次,这样会减少对处理器的负荷。如果你重复的装载页面， 你会发现"上次活动"的时间在五分钟，或多余五分钟的时候才会变化，也就是 cookie 上次被写入的时间。 这个时间可以通过设置 `application/config/config.php` 文件里的 `$config['sess_time_to_update']` 行来改变。

## 取得 Session 数据

可以通过如下的函数来得到 session 数组的任何信息:

    $this->session->userdata('item'); 

`item` 是数组里的相对应数据的索引。例如，想要获得 session ID， 你要使用如下的代码:

    $session_id = $this->session->userdata('session_id'); 

**注意:** 如果你的目标数据不存在的话，这个函数会返回 FALSE (布尔值boolean)。

## 添加自定义的 Session 数据

session 数组的一个非常有用的用途是你可以向它里面添加你自己的数据，这些数据会被保存在用户的 cookie 中。这样做的原因是什么呢？看看这个例子:

假设，有个特定用户登陆到你的网站， 当他通过检测后 你可以添加他的用户名和电子邮件到 session cookie 中，这些信息可以在不去访问数据库的情况下，当成全局量来使用。

通过以下函数，你可以传递一个新的用户数组到 session 数组中:

    $this->session->set_userdata($array); 

`$array` 是一个结合数组，用来存储你的新数据。例如 :

    $newdata = array(  
                   'username'  => 'johndoe',  
                   'email'     => 'johndoe@some-site.com',  
                   'logged_in' => TRUE  
               );  
    $this->session->set_userdata($newdata);

如果使用下面 set_userdata()函数的写法,可以每次只添加一个用户数据。

    $this->session->set_userdata('some_name', 'some_value');

**注意:** Cookies 只能存储 4KB 的数据, 使用时要小心超出它的容量。特别指出的是，加密会产生比原数据更长的数据字符串，所以一定要当心你要存放数据的大小。

## 取得所有 Session 数据

用下面的这种方式可以得到一个所有Session用户数据的数组:

    $this->session->all_userdata() 

代码将返回一个类似这样的关联数组:
    (
        [session_id] => 4a5a5dca22728fb0a84364eeb405b601
        [ip_address] => 127.0.0.1
        [user_agent] => Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_7;
        [last_activity] => 1303142623
    )

## 删除 Session 数据

正如使用 `set_userdata()` 是用来添加信息到 session 中，而通过向 `unset_userdata()` 函数中传递 session key 可以用来删除这些信息。例如, 你想要从 session 信息里去掉 'some_name': 

    $this->session->unset_userdata('some_name');

也可以给这个函数传一个要删除项的关联数组。

    $array_items = array('username' => '', 'email' => '');  
    $this->session->unset_userdata($array_items);

## 闪出数据

CodeIgniter 支持 "闪出数据", 或者说Session数据只对下次服务器请求可用, 然后会自动清除。这应该会非常有用，往往应用在信息或状态提示中（例如：“记录2已删除”）。

注意: 闪出数据变量名以“`flash_`”开头，所以在你自己的变量名中要避免使用这个前缀。

要添加闪出数据：

    $this->session->set_flashdata('item', 'value');

你也可以使用和 `set_userdata()` 同样的方式向 `set_flashdata()` 传递一个数组。

要读取一个闪出数据变量：

    $this->session->flashdata('item');

如果你发现你需要在一个附加的请求中保留一个闪出数据，你可以使用 keep_flashdata() 这个函数。

    $this->session->keep_flashdata('item');

## 将 Session 数据存入数据库

由于Session数据数组是附带一个Session ID保存在用户cookie里的，你无法验证它，除非你把session数据存储在数据库中。在一些不需要或很少需要安全保护的应用中，session ID 或许并不需要。但如果你的应用需要安全保护，验证是必须的。Otherwise, an old session could be restored by a user modifying their cookies.

当session 数据在数据库中可用时，每当从用户cookie中发现一个有效的session，一个数据库查询就会被执行以匹配它。如果 session ID 不相配，session 就会被销毁。Session ID永远不会更新，它们只会在一个新的会话创建时生成。

为了存储session，你必须先创建一个数据表。这是 session 类所需的基本结构（用于MySQL的）：

    CREATE TABLE IF NOT EXISTS `ci_sessions` (
        session_id varchar(40) DEFAULT '0' NOT NULL,
        ip_address varchar(16) DEFAULT '0' NOT NULL,
        user_agent varchar(120) NOT NULL,
        last_activity int(10) unsigned DEFAULT 0 NOT NULL,
        user_data text DEFAULT '' NOT NULL,
        PRIMARY KEY (session_id), KEY `last_activity_idx` (`last_activity`)
    ); 

**注意:** 默认情况下这个表叫做 `ci_sessions`, 但是你可以给它指定任意名字，只要你更新了 `application/config/config.php` 文件以确保它包含了你所起的名字。 一旦你创建了数据表，你就可以像下面这样在config.php文件中启用数据库选项：

    $config['sess_use_database'] = TRUE; 

一旦启用了，Session类就会在数据库中存储session数据。

同时确保你已经在配置文件中指定了数据表名：

    $config['sess_table_name'] = 'ci_sessions'; 

**注意:** Session类已经内置了清除过期session的垃圾回收机制，因此你不需要编写你自己的事务来做这个。

## 销毁 Session 

要清除当前 session: 

    $this->session->sess_destroy(); 

**注意:** 此函数应该是最后被调用的。即使闪出变量已不再有效。如果你只想让某几项而不是所有项被销毁，请使用 `unset_userdata()`.

## Session 的参数

你可以在`application/config/config.php` 文件中找到以下的 Session 相关的参数:

|参数                       |默认       |选项                       |描述
|---------------------------|-----------|---------------------------|------------------------------------------------
|**sess_cookie_name**       |ci_session |无                         |你想要保存 Session Cookie 的名字。
|**sess_expiration**        |7200       |无                         |session 持续的秒数。默认是2个小时(7200秒)。如果将这个数值设为: 0，就可以得到 永久 session。
|**sess_expire_on_close**   |FALSE      |TRUE/FALSE (boolean)       |这个选项决定当浏览器窗口关闭时是否自动使session过期。
|**sess_encrypt_cookie**    |FALSE      |TRUE/FALSE (布尔值boolean) |是否对 session 数据加密.
|**sess_use_database**      |FALSE      |TRUE/FALSE (布尔值boolean) |是否将 session 数据存放入数据库中。在开启这个选项前，你要先创建一个数据库表。
|**sess_table_name**        |ci_sessions|任何有效的 SQL 表名        |session 数据库表的名字。
|**sess_time_to_update**    |300        |时间以秒计算               |这个选项控制 session 类多久会产生一个新的session 和 session id。
|**sess_match_ip**          |FALSE      |TRUE/FALSE (布尔值boolean) |是否通过用户的IP地址来读取 session 的数据。 注意 ，有些网络运行商 ISPs 会动态的改变IP, 所以将这个选项设为 FALSE， 才有可能得到永久的 session。
|**sess_match_useragent**   |TRUE       |TRUE/FALSE (布尔值boolean) |是否要按照对应的 User Agent 来读取 session 数据。

翻译贡献者: aykirk, bearcat001, Drice, Fanbin, guns1985, Hex, mchipengfei, mental, noproblem, qixingyue, shishirui, xwjie

最后修改: 2012-02-08 20:07:53

