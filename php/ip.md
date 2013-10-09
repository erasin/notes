#IP

##IP138

<http://wap.ip138.com/ip.asp?ip=115.238.95.194>

## 新浪

<http://counter.sina.com.cn/ip?ip=115.238.95.194>

<http://int.dpool.sina.com.cn/iplookup/iplookup.php?format=js&ip=115.238.95.194>

## 搜狐

缺省返回 gbk for js

<http://pv.sohu.com/cityjson>

<http://txt.go.sohu.com/ip/soip >

## 腾讯 

gbk


<http://ip.qq.com/cgi-bin/searchip?searchip1=>

	/**根据腾讯IP分享计划的地址获取IP所在地，比较精确 */
	function getIPLoc_QQ($queryIP){
	    $url = 'http://ip.qq.com/cgi-bin/searchip?searchip1='.$queryIP;
	    $ch = curl_init($url);
	    curl_setopt($ch,CURLOPT_ENCODING ,'gb2312');
	    curl_setopt($ch, CURLOPT_TIMEOUT, 10);
	    curl_setopt($ch, CURLOPT_RETURNTRANSFER, true) ; // 获取数据返回
	    $result = curl_exec($ch);
	    $result = mb_convert_encoding($result, "utf-8", "gb2312"); // 编码转换，否则乱码
	    curl_close($ch);
	    preg_match("@<span>(.*)</span></p>@iU",$result,$ipArray);
	    $loc = $ipArray[1];
	    return $loc;
	}