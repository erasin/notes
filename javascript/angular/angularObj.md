<!--
Title: Use AngularJs To Create A APP
TimeLine: Sat 28 Jun 2014
--> 

# 使用 AngularJs 创建一个项目

最为宽松的情况下也可以创建单页面APP，使用 HTML5+JS+CSS 和 RESTFUL 服务创建。

工具：

* AngularJs包 [1.2.17](/)
* ionic 包 [1.0bata](http://code.ionicframework.com/1.0.0-beta.8/ionic-v1.0.0-beta.8.zip)
* 现代浏览器　chrome/firefox/IE9+
* 文本编辑器或IDE 
	- Sublime/vim
* SERVER

## 部署项目

```text
AppName/
| - index.html  	# 入口
| - views/ 			# 模板文件
| - app/			# 主程序位置
|   | - app.js  		# 路由器 
| 	| - controller.js 	# 控制器 
| 	| - services.js 	# REST请求服务
| 	| - directives.js 	# 指令集
| 	` - filters.js 		# 过滤器
| - js/         	# js 文件
| 	| - angular/ 	# angular 包
| 	| - ionic/   	# ionic
| 	| - pkgs/       # 第三方包
| 	| - libs/    	# 其他的js包
| 	` - other... 	
| - css/ 			# 样式
| - imgs/ 			# 图片
| - fonts/ 			# 字体
` - other/      	# 其他
```

## 



