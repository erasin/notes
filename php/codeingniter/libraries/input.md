# 输入类

输入类有两个目的：

1.  为了安全，预处理输入数据。
2.  提供helper的一些方法，取得输入数据，并预处理输入数据。

**说明:** 系统自动加载此类，不用手动加载。

## 安全过滤（Security Filtering）

当触发一个[控制器][1]的时候，安全过滤（Security Filtering）功能自动启动。做以下事情：

*   如果 $config['allow\_get\_array'] 的值为FALSE(默认为 TRUE), 销毁全局GET数组。
*   当 register_globals 被设置为 on 的时候，销毁所有的全局变量。
*   过滤 GET/POST/COOKIE 数组键，只允许字母-数字(以及一些其它的)字符。
*   可以过滤跨站脚本攻击 (Cross-site Scripting Hacks) 此功能可全局打开（enabled globally），或者按要求打开。
*   换行符统一换为 \n（Windows 下为 \r\n）

## 跨站脚本（XSS）过滤

输入类有能力阻止跨站脚本攻击。如果你想让过滤器遇到 POST 或者 COOKIE 数据时自动运行，你可以通过打开你的 <kbd>application/config/config.php</kbd> 文件进行如下设置实现:

    $config['global_xss_filtering'] = TRUE; 

请参考 [安全类][2] 文档以获得更多信息在你的应用中使用跨站脚本过滤。

## 使用 POST, COOKIE, 或 SERVER 数据

CodeIgniter 有3个 helper方法可以让用户取得POST, COOKIE 或 SERVER 的内容。用这些方法比直接使用php方法($_POST['something'])的好处是不用先检查此项目是不是存在。 直接使用php方法，必须先做如下检验：

    if ( ! isset($_POST['something']))
    {
        $something = FALSE;
    }
    else
    {
        $something = $_POST['something'];
    }

用CodeIgniter内建的方法，你可以这样：

    $something = $this->input->post('something'); 

这3个方法是：

*   $this->input->post()
*   $this->input->cookie()
*   $this->input->server()

## $this->input->post()

第一个参数是所要取得的post中的数据：

    $this->input->post('some_data'); 

如果数据不存在，方法将返回 FALSE (布尔值)。

第二个参数是可选的，如果想让取得的数据经过跨站脚本过滤（XSS Filtering），把第二个参数设为TRUE。

    $this->input->post('some_data', TRUE); 

不设置任何参数，该方法将以一个数组的形式返回全部POST过来的数据。

把第一个参数设置为NULL，第二个参数设置为 TRUE (boolean)，该方法将经过跨站脚本过滤，返回一个包含全部POST数据的数组。

如果POST没有传递任何数据，该方法将返回 FALSE (boolean)

    
      $this->input->post(NULL, TRUE); // 经过跨站脚本过滤 返回全部 POST 数据
      $this->input->post(); // 不经过跨站脚本过滤 返回全部 POST 数据
     

## $this->input->get()

此方法类似post方法，用来取得get数据：

    $this->input->get('some_data', TRUE); 

如果没有设置参数将返回GET的数组

如果第一参数为NULL，且第二参数为True，则返回经过跨站脚本过滤（XSS Filtering）的数组。

如果没有设从GET中取到数据将返回 FALSE (boolean)

    
      $this->input->get(NULL, TRUE); // 经过跨站脚本过滤 返回全部 GET 数据
      $this->input->get(); // 不经过跨站脚本过滤 返回全部 GET 数据
     
## $this->input->get_post()

这个方法将会搜索POST和GET方式的数据流，首先以POST方式搜索，然后以GET方式搜索:

    $this->input->get_post('some_data', TRUE); 
## $this->input->cookie()

此方法类似post方法，用来取得cookie数据：

    $this->input->cookie('some_data', TRUE); 
## $this->input->server()

此方法类似上面两个方法，用来取得server数据：

    $this->input->server('some_data'); 
## $this->input->set_cookie()

设置一个 Cookie 的值。这个函数接收两种形式的参数：数组形式和参数形式：

#### 数组形式

用这种形式的话，第一个参数传递的是一个关联数组：

    $cookie = array(
        'name'   => 'The Cookie Name',
        'value'  => 'The Value',
        'expire' => '86500',
        'domain' => '.some-domain.com',
        'path'   => '/',
        'prefix' => 'myprefix_',
        'secure' => TRUE
    );

    $this->input->set_cookie($cookie);
    

**说明：**

只有 name 和 value 是必须的。可以通过将 expire 设置成空来实现删除 Cookie 的操作。

Cookie 的过期时间是以**秒**为单位来设置的，他是通过将 Cookie 的存续的时间值加上当前系统时间来得到的。切记，expire 的值仅仅设置为Cookie 需要存续的时间长短，请不要将当前的系统时间加上存续时间后再赋给变量。如果将 expire 设置成零，那么 Cookie 仅在浏览器关闭的时候失效。

如果需要设置全站范围内使用的cookie，无论你怎么请求都可以,那么你要把你的网站域名赋给$domain变量，并且需要以英文的句号"."开头，如: .your-domain.com

path通常是不需要设置的，该方法设置path为网站的根目录。

prefix(前缀)只有在为了避免和其它服务器上的相同命名的cookies冲突是才需要使用。 

secure(安全)设置选项只有在你想把他设置成安全的cookie时，才需要把secure设置为 TRUE(boolean).

#### 参数形式

你可以通过一个个单独的参数来设置cookies, 如果你喜欢的话：

    $this->input->set_cookie($name, $value, $expire, $domain, $path, $prefix, $secure); 

## $this->input->ip_address()

返回当前用户的IP。如果IP地址无效，返回0.0.0.0的IP：

    echo $this->input->ip_address(); 

## $this->input->valid_ip(<var>$ip</var>)

测试输入的IP地址是不是有效，返回布尔值TRUE或者FALSE。 注意：$this->input->ip_address()自动测试输入的IP地址本身格式是不是有效。

    if ( ! $this->input->valid_ip($ip))
    {

         echo 'Not Valid';
    }
    else
    {
         echo 'Valid';
    } 

## $this->input->user_agent()

返回当前用户正在使用的浏览器的user agent信息。 如果不能得到数据，返回FALSE。

    echo $this->input->user_agent(); 

查看[User Agent Class][3]了解在user agent字符串中该方法的更多扩展信息。

## $this->input->request_headers()

在不支持[apache\_request\_headers()][4]的非Apache环境非常有用。返回请求头（header）数组。

    $headers = $this->input->request_headers(); 

## $this->input->get\_request\_header();

返回请求头（request header）数组中某一个元素的值

    $this->input->get_request_header('some-header', TRUE); 

## $this->input->is\_ajax\_request()

检查服务器头HTTP_X_REQUESTED_WITH是否被设置，并返回布尔值。

    $this->input->is_ajax_request() 
    
## $this->input->is\_cli\_request()

检查看常量STDIN是否被设置, 这只是一个检查PHP是否以命令行方式运行的应急方法。

    $this->input->is_cli_request()

翻译贡献者: architectcom, Hex, hk\_yuhe, IT不倒翁, loiynet, qiutao520, soyota, sunjiaxi, xjflyttp, yinzhili, yzheng624, 暗夜星辰, 月夜之人

最后修改: 2012-10-25 09:33:26

[1]: ../general/controllers
[2]: security
[3]: user_agent
[4]: http://php.net/apache_request_headers
