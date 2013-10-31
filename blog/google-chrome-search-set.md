# Chrome 设定 

## Mac/linux版修改google默认的引擎地址：

退出Chrome

进入 `~/Library/Application Support/Google/Chrome/Default` , linux 下为 `.config/google-chrome/Default`


打开 Preference 这个文件，找到找到以下内容：

	"last_known_google_url": "http://www.google.com.hk/",
	"last_prompted_google_url": "http://www.google.com.hk/",

把URL里面的`http://google.com.hk`改成`https://google.com`即可


* <https://www.google.com/ncr>
* <https://www.google.com/webhp#>

## 备份和恢复插件参数

输出Chrome扩展参数和数据的方法

在Chrome任意扩展的页面（选项或弹出页等），按F12调出“审查元素”Console控制台，输入如下代码，即可输出：

	var L=localStorage,l=[];for (var i=0; i < L.length; i++){var k=L.key(i);var v=L.getItem(k);if(!v){v=''}l.push('{"k":"' +k+ '","v":"' + escape(v) + '"}');}console.log('['+l.join(',')+']');

导入Chrome扩展参数和数据的方法

同样进入Console控制台，输入如下代码，即可导入

	var b='备份字符串';var jA=eval('('+ b +')');for(var i=0;i<jA.length;i++){var j=jA[i];L[j.k]= unescape(j.v);}