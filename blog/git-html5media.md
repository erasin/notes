#HTML5Meadia 最佳web多媒体播放器(兼容IE)

项目名称：html5media   
项目主页：<http://html5media.info/>   
项目git仓库: [git://github.com/etianen/html5media.git](https://github.com/etianen/html5media)

html5media是为了解决不支持html5的浏览器播放html5媒体标签而开发的项目.对于不支持html5的播放器使用flowplayer来播放视频或音频.

<script src="http://api.html5media.info/1.1.4/html5media.min.js"></script>
<video class="video" poster="http://media.html5media.info/poster.jpg" width="618" height="347" controls preload> 
<source src="http://media.html5media.info/video.mp4" media="only screen and (min-device-width: 960px)"></source> 
<source src="http://media.html5media.info/video.iphone.mp4" media="only screen and (max-device-width: 960px)"></source> 
<source src="http://media.html5media.info/video.ogv"></source> 
</video>
<audio class="audio" controls preload> 
<source src="http://media.html5media.info/audio.mp3"></source> 
<source src="http://media.html5media.info/audio.ogg"></source> 
</audio>

事实上,web播放器很多,因为IE市场的关系,多数使用播放器仍为flash播放器.处理代码上并不简洁.

在页面中加入.

     <script src="http://api.html5media.info/1.1.4/html5media.min.js"></script>

之后使用html5的

     <!-- 视频媒体标签 -->
     <video src="video.mp4" width="320" height="200" controls preload></video>
     <!-- 音频媒体标签 -->
     <audio src="audio.mp3" controls preload></audio>

之前我见过一个朋友在使用的时候判断浏览器来执行js以实现兼容,做起来也不怎么易用.

像这类为不支持html5标签而出现js项目,随着HTML5浏览器和HTML5的优势展现会越来越多起来.
