# 边框 Borders

用CSS3，你可以创建圆角边框，添加阴影框，并作为边界的形象而不使用设计程序，如Photoshop。

浏览器支持

属性            |  IE 9+ | firefox | chrome | safari | opera
----------------|--------|--------|--------|---------|--------
border-radius   | <i class="fa fa-check green"></i> | <i class="fa fa-check green"></i> | <i class="fa fa-check green"></i>| <i class="fa fa-check green"></i>| <i class="fa fa-check green"></i> 
box-shadow      | <i class="fa fa-check green"></i> | <i class="fa fa-check green"></i> | <i class="fa fa-check green"></i>| <i class="fa fa-check green"></i>| <i class="fa fa-check green"></i>
border-image    | <i class="fa fa-times red"></i> | <i class="fa fa-check green"></i> | <i class="fa fa-check green"></i>| <i class="fa fa-check green"></i>| -o- <i class="fa fa-check green"></i> 

Internet Explorer 9+ 支持 border-radius 和 box-shadow.  
Firefox, Chrome, 和 Safari 支持所有最新的 border 属性.  
注意： 前缀是-webkit-的Safari支持阴影边框。  
前缀是-o-的Opera支持边框图像。

## 圆角 border-radius

一个用于设置所有四个边框- *-半径属性的速记属性   

**语法:** 
    border-radius: 1-4 length|% / 1-4 length|%;

> 注意: 每个半径的四个值的顺序是：左上角，右上角，右下角，左下角。如果省略左下角，右上角是相同的。如果省略右下角，左上角是相同的。如果省略右上角，左上角是相同的。 类似于**padding、margin**之类 。

值   |描述
-----|-------
length  |定义弯道的形状。
%   |使用%定义角落的形状。

**JavaScript 语法:** 
    object object.style.borderRadius="5px"

```css
.border-radius{
    border:1px solid #333;
    border-radius: 5px;
    
    /**等价于**/
    border-top-left-radius:5px;
    border-top-right-radius:5px;
    border-bottom-right-radius:5px;
    border-bottom-left-radius:5px;

}
```

结果为

<div style="width:100px;text-align:center;border:1px solid #333;border-radius: 5px;">test</div>

## 阴影 border-shadow

附加一个或多个下拉框的阴影

**语法**
    box-shadow: h-shadow v-shadow blur spread color inset;

**JavaScript 语法:** 
    object.style.boxShadow="10px 10px 5px #888888"

> 注意：boxShadow 属性把一个或多个下拉阴影添加到框上。该属性是一个用逗号分隔阴影的列表，每个阴影由 2-4 个长度值、一个可选的颜色值和一个可选的 inset 关键字来规定。省略长度的值是 0。 

值  |说明
----|-----
h-shadow   |必需的。水平阴影的位置。允许负值
v-shadow   |必需的。垂直阴影的位置。允许负值
blur   |可选。模糊距离
spread |可选。阴影的大小
color  |可选。阴影的颜色。在CSS颜色值寻找颜色值的完整列表
inset  |可选。从外层的阴影（开始时）改变阴影内侧阴影


```css
.border-shadow{
    box-shadow: 8px 8px 5px #ccc;
}
```
<div style="width:100px;text-align:center;border:1px solid #333;border-radius: 5px;box-shadow: 10px 10px 5px #ddd;">test</div>

### 图片边框
设置所有边框图像的速记属性。

**语法**
    border-image: source slice width outset repeat;

**JavaScript 语法:**  
    object.style.borderImage="url(border.png) 30 30 round"

值  |描述
----|-----
border-image-source|用于指定要用于绘制边框的图像的位置
border-image-slice |图像边界向内偏移
border-image-width |图像边界的宽度
border-image-outset|用于指定在边框外部绘制 border-image-area 的量
border-image-repeat|这个例子演示了如何创建一个border-image 属性的按钮。


```css
.border-image{
    border-image:url(border.png) 30 30 round;
    -webkit-border-image:url(border.png) 30 30 round; /* Safari 5 and older */
    -o-border-image:url(border.png) 30 30 round; /* Opera */
}
```

<div style="text-align:center;border:1px solid #333;border-radius: 5px;box-shadow: 10px 10px 5px #ddd;-webkit-border-image:url(http://1.su.bdimg.com/icon/weather/aladdin/png_18/a0_night.png) 30 30 round;">test</div>

...



