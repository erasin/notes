

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
|   | - controller.js 	# 控制器 
|   | - services.js 	# REST请求服务
|   | - directives.js 	# 指令集
|   ` - filters.js 		# 过滤器
| - js/         	# js 文件
|   | - angular/ 	# angular 包
|   | - ionic/   	# ionic
|   | - pkgs/       # 第三方包
|   | - libs/    	# 其他的js包
|   ` - other... 	
| - css/ 			# 样式
| - imgs/ 			# 图片
| - fonts/ 			# 字体
` - other/      	# 其他
```

##  AngularJs

[supercell]

**入口文件index.html**

```html
<!DOCTYPE html>
<!-- 定义了app位置  -->
<html ng-app="app">
    <head>
        <meta charset="utf-8">
        <meta name="viewport" content="initial-scale=1, maximum-scale=1, user-scalable=no, width=device-width">
		<title ng-bind="$root.title+' - seed'"></title>

		<!-- 这里使用了boostrap css -->
	    <link rel="stylesheet" href="css/bootstrap.css">
	    <link rel="stylesheet" href="css/font-awesome.min.css">
        <link href="css/app.css" rel="stylesheet">

		<!-- 基本资源 -->
        <script src="js/angular/angular.min.js"></script>
        <script src="js/angular/angular-resource.min.js"></script>

		<!-- 服务文件 -->
		<script src="app/services.js"></script>
	    <script src="app/directives.js"></script>
	    <script src="app/filters.js"></script>
		<script src="app/controllers.js"></script>
		<script src="app/app.js"></script>
    </head>
    <body>
    	<!-- 头部 -->
        <div ng-include="'views/header.html'"></div>

	    <div class="container">
	        <div class="row" ng-view></div>
	    </div>
    </body>
</html>
```

在 `app.js`中定义 app

```javascript
// 使用引入的包
var app = angular.module('app', ['ngRoute','app.controllers', 'app.services','app.filters'])
```


