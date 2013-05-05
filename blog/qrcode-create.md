#二维码生成 qrencode

## 安装：
>	sudo apt-get install qrencode

## 使用：

>	qrencode -o `[`filename.png`] `‘`[`text 1="to" 2="encode" language="/url/information"`][`/text`][`/text`]`‘

比如要生成本站的二维码
>	qrencode -o wowubuntu.png 'http://wowubuntu.com'

输出图形如下，如果你的手机上安装了二维码识别软件的庆，你可以用手机进行拍摄并识别了。 

想自定义尺寸的话，加上 -s 参数，比如 -s 6 表示尺寸为 6x6 平方像表大小，如下。

>	qrencode -o ~/Desktop/google.png -s 6 'http://wowubuntu.com'

除此之外，你还可以使用更多其它参数，详细用法请 man qrencode 。
