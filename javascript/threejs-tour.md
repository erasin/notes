# three.js 

[jb51](http://www.jb51.net/article/33517.htm)

Three.js是一个伟大的开源WebGL库,WebGL允许JavaScript操作GPU，在浏览器端实现真正意义的3D,本文介绍下THREE.JS基础/浏览器支持/设置场景/构建网格表面/材质/光/渲染循环/通用的对象属性等等，感兴趣的朋友可以了解下哦
Three.js是一个伟大的开源WebGL库，WebGL允许JavaScript操作GPU，在浏览器端实现真正意义的3D。但是目前这项技术还处在发展阶段，资料极为匮乏，爱好者学习基本要通过Demo源码和Three.js本身的源码来学习。

国外网站 aerotwist.com 有六篇较为简单的入门教程，我尝试着将其翻译过来，与大家分享。
我在一些实验项目中使用了Three.js，我发现它对快速上手浏览器3D编程确实很有帮助。通过Three.js，你不仅可以创建相机、物体、光线、材质等等，还可以选择着色器，可以决定使用何种技术（WebGL、Canvas或SVG）在网页上渲染你的3D图形。Three.js是开源的，你甚至可以参与到这个项目中来。但现在，我将把重点放在基础的介绍上，我将向你展示如何使用这个引擎上。

尽管Three.js如此奇妙，但有时候它也会令人抓狂。比如，你将花费大量时间阅读例程，做一些逆向工程（在我的情形下）来确定某个函数的作用，有时还要去GitHub上提问。如果你需要提问，Mr. doob和AlteredQualia是极好的选择。

1.基础
我假定你的三维图形学知识过关，而且也在一定程度上掌握了JavaScript。如果不是这样，那先去学一点吧，否则直接看这篇教程，也许会感到困惑。

在我们的三维世界里，我们有以下这些东西。我会带你一步一步创建它们。
1.场景
2.渲染器
3.相机
4.物体（带有材质的）
当然，你也可以创造些其他的什么东西，我也希望你如此做。
2.浏览器支持
简单地看一下浏览器的支持情况吧。Google家的Chrome浏览器支持Three.js，在我的实验里，无论是对渲染器的支持程度还是JavaScript解释器的运行速度，Chrome都是做得最好的：它支持Canvas、WebGL和SVG，而且运行得非常快。FireFox浏览器排在第二位，它的JavaScript引擎的速度比Chrome慢了半拍，但是对渲染器的支持也很棒，而且FireFox的速度，随着版本更新也越来越快。Opera浏览器正在逐渐增加对WebGL的支持，Mac上的Safari浏览器有一个开启WebGL的选项。总体上，这两个浏览器仅仅支持Canvas渲染。微软家的IE9现在只支持Canvas渲染，而且微软似乎并不乐意支持WebGL这个新特性，所以我们现在肯定不会用IE9来做实验。
3.设置场景
假定你已经选择了一个支持所有渲染技术的浏览器，而且你准备通过Canvas或WebGL来渲染场景（这是更标准化的选择）。Canvas比WebGL有着更广泛地支持，但是WebGL可以直接在GPU上操作，这意味着你的CPU可以专注地处理非渲染类的工作，比如物理引擎或与用户交互等。

无论你选择何种渲染器，你都必须牢记在心的是：JavaScript代码需要优化。三维显示对浏览器来说不是一项轻松的工作（现在能够这样做就很伟大了），所以如果你的渲染太慢了，你需要知道你代码的瓶颈在何处，如果可能，改善它。

说了这么多，我想你已经下载好Three.js源代码，而且将它引入了你的html文档了。那么如何开始创建一个场景呢？就像这样：
复制代码 代码如下:

// 设置场景大小
var WIDTH = 400,
HEIGHT = 300;
// 设置一些相机参数
var VIEW_ANGLE = 45,
ASPECT = WIDTH / HEIGHT,
NEAR = 0.1,
FAR = 10000;
// 获取DOM结构中的元素
// - 假设我们使用了JQuery
var $container = $('#container');
// 创建渲染器、相机和场景
var renderer = new THREE.WebGLRenderer();
var camera =
new THREE.PerspectiveCamera(
VIEW_ANGLE,
ASPECT,
NEAR,
FAR);
var scene = new THREE.Scene();
// 将相机加入场景
scene.add(camera);
// 相机的初始位置为原点
// 将相机拉回来一些（译者注：这样才能看到原点）
camera.position.z = 300;
// 启动渲染器
renderer.setSize(WIDTH, HEIGHT);
// 将渲染器加到DOM结构中
$container.append(renderer.domElement);

你看，简单吧！
4.构建网格表面
现在我们有了一个场景，一个相机和一个渲染器（在我的例子里，当然是一个WebGL渲染器），但我们事实上什么还没画呢。事实上，Three.js提供了载入某几种标准格式3D文件的支持，如果你在Blender，Maya，Cinema4D或是什么其他工具中建模，这简直太棒了。为了简单（毕竟这才刚开始呢！）我们先来考虑基元。基元就是基本的几何表面，比如最基本的球体、平面、立方体、圆柱体。利用Three.js可以很方便地创建这些基元：
复制代码 代码如下:

// 设置球体参数（译者注：球体被划分为16×16的网格，如果后两个参数取4、2，则生成一个八面体，请想象）
var radius = 50,
segments = 16,
rings = 16;
// material覆盖在geometry上，生成mesh
var sphere = new THREE.Mesh(
new THREE.SphereGeometry(
radius,
segments,
rings),
sphereMaterial);
// 将mesh加入到场景中
scene.add(sphere);

好了，但是球体上的材质呢？在代码中我们使用了一个sphereMaterial变量，我们还没定义它呢。那我们就先来看看怎么创建材质吧。
5.材质
毫无疑问，这是Three.js最有用的部分了。这部分提供了几个非常易用的通用材质模型：
1.Basic材质：表示一种不考虑光照的材质，现在只能这么说了。
2.Lambert材质：（译者注：朗伯面，各向同性反射）。
3.Phong材质：（译者注：冯氏面，有光泽的表面，介于镜面反射和朗伯反射之间的反射，描述真实世界的反射）。

除此之外，还有一些其他类型材质，简单起见，就留给你自己探索。事实上，在使用WebGL类型的渲染器时，材质实在太好用了。为什么呢？因为在原生WebGL种你必须亲自为每个渲染编写着色器，而着色器本身就是个巨大的工程：简单地说着色器是使用GLSL语言（OpenGL的着色器语言）写的，用来操作GPU的程序，这意味着你要在数学上模拟光照，反射等等，这很快就变成一项极为复杂的工作。多亏有了Three.js你才可以不必去自己编写着色器，当然，如果你想亲自编写的话，你可以使用MeshShaderMaterial，可见这是很灵活的设定。

现在，让我们用朗伯面材质覆盖球体：
复制代码 代码如下:

// 创建球体表面的材质
var sphereMaterial =
new THREE.MeshLambertMaterial(
{
color: 0xCC0000
});

值得指出的是，创建材质的时候，除了颜色还有很多其他参数可以指定，比如光滑度和环境贴图。你可以需要检索这个Wiki页面来确认哪些是哪些属性可以设置在材质上，或Three.js引擎提供的任何对象上。
6.光
如果你现在就想渲染场景，你会看到一个红色的圆。虽然我们在球体上覆盖了朗伯面材质，但场景里没有光。所以按照默认设定，Three.js会恢复到满环境光，物体的看上去的颜色就是物体表面的颜色。让我们添加一个简单的点光源：
复制代码 代码如下:

// 创建一个点光源
var pointLight =
new THREE.PointLight(0xFFFFFF);
// 设置点光源的位置
pointLight.position.x = 10;
pointLight.position.y = 50;
pointLight.position.z = 130;
// 将点光源加入场景
scene.add(pointLight);

7.渲染循环
显然，关于渲染器的一切都设置好了。万事俱备，我们现在只需要：
复制代码 代码如下:

// 画!
renderer.render(scene, camera);

你很可能像多次渲染，而不是只渲染一次，所以如果你要去做一个循环，你应该使用requestAnimationFrame。这是目前最好的，在浏览器中处理动画的方法，虽然还没有得到最全面的支持，但我强烈建议你去看一看Paul Irish的博客。
8.通用的对象属性
如果你花点时间去浏览一遍Three.js的源代码，你会发现很多对象都继承自Object3D。这个基类包含了很多有用的属性，比如位置、旋转和缩放的信息。特别的，我们的球体是一个Mesh对象，而Mesh对象继承自Object3D对象，但是又增加了些自己的属性：geometry和material。为什么要说这些？因为你一定不会只满足于屏幕中一个什么都不做的圆球，而这些（译者注：基类中的）属性允许你操作Mesh对象更底层的细节和各种各样的材质。
复制代码 代码如下:

// sphere是一个mesh对象
sphere.geometry
// sphere包含了一些点和面的信息
sphere.geometry.vertices // 一个数组
sphere.geometry.faces // 另一个数组
// mesh对象继承自object3d对象
sphere.position // 包含x,y,z
sphere.rotation // 同上
sphere.scale // ... 同上

9.讨厌的秘密
我希望这样说你能很快弄明白：就是如果你修改了，比如说，一个mesh对象的顶点属性vertices，你会发现在渲染循环中，什么都没变。为什么？因为Three.js将mesh对象的信息缓存为某种优化结构了。你真正要做的是给Three.js一个标识，告诉它如果什么东西改变了，需要重新计算缓存中的结构：
复制代码 代码如下:

// 设置geometry为动态的，这样才允许改变其中的顶点
sphere.geometry.dynamic = true;
// 告诉Three.js，需要重新计算顶点
sphere.geometry.__dirtyVertices = true;
// 告诉Three.js，需要重新计算顶点
sphere.geometry.__dirtyNormals = true;

还有更多的标识，但我发现这两个是最有用的。你应该仅仅标识那些确实需要实时计算的属性来避免无谓的运算开销。
10.小结
我希望这篇简单的介绍对你有所帮助。没什么能比得上卷起袖子亲手实践了，我强烈建议你这样做。在浏览器里面运行3D程序很有意思，而且使用像Three.js这样一个引擎免去了很多麻烦，让你一开始就能专注于那些真正cool的事情。



Three.js是一个伟大的开源WebGL库，WebGL允许JavaScript操作GPU，在浏览器端实现真正意义的3D,本文将介绍顶点着色器/片元着色器/着色器变量等等，感兴趣的朋友可以了解下哦
译序
Three.js是一个伟大的开源WebGL库，WebGL允许JavaScript操作GPU，在浏览器端实现真正意义的3D。但是目前这项技术还处在发展阶段，资料极为匮乏，爱好者学习基本要通过Demo源码和Three.js本身的源码来学习。

0.简介
之前我已经给出了一篇《开始使用Three.js》。如果你还没有读过，你可能需要去读一下，本文的基础是在那一篇教程的基础上完成的。

我想讨论一下着色器。在Three.js帮助你免去了很多麻烦之前，原生WebGL就很优秀了。有的时候，你也许会想要完成一些特定的效果，或者想对呈现在你的屏幕上的东西钻研得更深入一些，那么着色器一定会进入你的视野。如果你像我一样，你也同样希望实现一些比上一篇教程中的基础更加有意思的东西。这篇教程中，我会讲解Three.js的基础，这些基础实际上为我们做了很多枯燥的工作。
在开始之前我还要说，这篇教程会有相当多的篇幅在解释着色器的代码，之后会有一篇教程会在着色器代码的基础上前进一点，利用着色器去做点什么。这是因为着色器shaders第一眼看上去并不易懂，需要一些解释。
1.两种着色器
WebGL没有固定的渲染管线，你无法直接使用一个黑盒子式的着色器（译者注：上个世纪的显卡基本都只支持固定渲染管线）；WebGL提供的是可编程的管线，这种方式更强大但也更难理解和使用。长话短说，可编程渲染管线意味着编写程序的人要自己负责获取顶点并将它绘制在屏幕上了。着色器是渲染管线的一部分，有两种着色器：
1.顶点着色器
2.片元着色器
你应当知道的是，这两种着色器都完全运行在显卡的GPU上，我们将需要它们处理的数据从CPU上卸下，装到GPU上，减轻了CPU的复旦。现代的GPU对着色器需要的调用的运算类型都做了大幅优化，这样做很值得。
2.顶点着色器
基元形状，比如一个球体，是由顶点构成的，是吧？顶点着色器被依次传入这些顶点中的一个顶点，然后处理它。如何处理每个顶点是可以自由定制的，但顶点着色器有一个必做的事，就是为一个名为 gl_Position 的变量赋值，该变量是一个4维数组，表示该顶点最终在屏幕上的位置。这本身是个有意思的过程，因为我们实际上在谈论如何将一个三维坐标（一个具有x、y、z值得顶点）转化为，或者说投影到二维的屏幕上。谢天谢地，要是我们使用Three.js之类的工具，我们能够如此方便地访问到 gl_Position 。
3.片元着色器
现在我们有包含顶点的三维物体了，现在要将物体投影到二维屏幕上了，但颜色哪里去了？纹理和光照呢？这正是片元着色器要处理的。
和顶点着色器类似，片元着色器有一项必须完成的任务：设置或消除变量 gl_FragColor ，另一个四维浮点变量，也就是片元点最终的颜色。什么是片元？想象一个具有三个顶点的三角形，片元就是经过这三个顶点计算后的，所有在三角形内部的点。因此，片元值由顶点的值内插生成。如果一个顶点的颜色是红色，相邻顶点的颜色是蓝色，那么我们可以观测到颜色从红色顶点附近渐变，由红色变成紫色，最终在蓝色顶点附近变成蓝色。
4.着色器变量
说到着色器变量，有三种：Uniforma，Attributes和Varyings。当我第一次听到这三个词语时，我很困惑，因为它们和我之前用到的东西完全不匹配。但现在，你可以这样理解它们：
1.Uniforms变量既可以传入顶点着色器，也可以传入片元着色器，它们包含了哪些在整个渲染过程中保持不变的变量，比如，一个点光源的位置。
2.Attributes变量对应于每个顶点，它们只可以传入顶点着色器中，比如每个顶点都具有一个颜色。Attributes变量和顶点的关系是一一对应的。
3.Varyings变量是在顶点着色器中定义，并且准备传入给片元着色器的变量。为了确保这点，我们需要确保在两个着色器中变量的类型和命名完全一致。一个经典的应用是法线向量，因为在计算光照的时候需要用到法线。
在后面一篇教程中，我会使用这三种变量，你也会学习到这三种变量如何真正应用起来得。
现在，我们已经谈过了顶点着色器、片元着色器和三种着色器变量。是时候来看一个我们可以创建的最简单的着色器了。
5.Hello World（译者吐槽：能不能不要秀法语啊）
这儿有一个最简单的顶点着色器：
复制代码 代码如下:

/**
* 每个顶点坐标乘以模型视图矩阵在乘以投影矩阵
* 获得在二维屏幕上的坐标
*/
void main() {
gl_Position = projectionMatrix *
modelViewMatrix *
vec4(position,1.0);
}

一个最简单的片元着色器：
复制代码 代码如下:

/**
* 将任意一个像元色设置为粉红
*/
void main() {
gl_FragColor = vec4(1.0, // R
0.0, // G
1.0, // B
1.0); // A
}

这就是全部了。如果现在直接运行的话，你就可以在屏幕上看到一个“无光”的粉红色形体。不是很复杂，是吗？

在顶点着色器中，我们通过Three.js传入了一些uniforms变量。有两个4×4的矩阵uniforms变量：模型视图矩阵和投影矩阵。你并不需要太了解这两个矩阵是怎么工作的。简单地说，这两个矩阵描述了三维点坐标如何投影成为二维屏幕上的坐标。

事实上，我只介绍了这两段简短的代码段。Three.js在你自己的着色器代码前已经将它们加进来了，所以你不必担心。实话说，Three.js还加了很多东西在你的代码前面，比如光照数据、节点颜色和节点法向量等等。如果没有Three.js你要亲自创建并设置这些对象，真的。
6.使用着色器材质
复制代码 代码如下:

/**
* 假设我们可以使用JQuery
* 将着色器的代码文本从DOM中抽取出来
*/
var vShader = $('vertexshader');
var fShader = $('fragmentshader');
var shaderMaterial =
new THREE.ShaderMaterial({
vertexShader: vShader.text(),
fragmentShader: fShader.text()
});

从这儿开始，Three.js将会编译并运行你的着色器，将其连接在你创建的材质上，材质又依附于你创建的mesh上。它并没有变得比真的更容易。也许是这样吧，但我们在考虑浏览器3D编程，我想你应该预期，这个话题是有一定复杂性的。

我们还可以像着色器材质添加另外两种属性：uniforms和attributes。他们可以是向量、整数或者浮点数，但是如我之前所说，uniforms变量在计算所有点的过程中保持不变，所以它们更加可能是单一的值，而attributes变量是对每个顶点而言的，所以他们应当是数组。在一个mesh中，attribute变量和顶点应当是一一对应的。
7.小结
这篇教程就到这里了，实际上我已经讲得很多了，但是在许多方面我都只是一掠而过。在下一篇教程中我会提供一个复杂的着色器，通过它我将传入一些attributes变量和uniforms变量来做一些模拟光照效果。
我将这篇教程的源码打包了，你可以下载下来作为参考 



译序
Three.js是一个伟大的开源WebGL库，WebGL允许JavaScript操作GPU，在浏览器端实现真正意义的3D。但是目前这项技术还处在发展阶段，资料极为匮乏，爱好者学习基本要通过Demo源码和Three.js本身的源码来学习。

.简介
这是WebGL着色器教程的后半部分，如果你没看过前一篇，阅读这一篇教程可能会使你感到困惑，建议你翻阅前面的教程。

上一篇结束的时候，我们在屏幕中央画了一个好看的粉红色的球体。现在我要开始创建一些更加有意思的东西了。

在这一篇教程中，我们会先花点时间来加入一个动画循环，然后是顶点attributes变量和一个uniform变量。我们还要加一些varying变量，这样顶点着色器就可以向片元着色器传递信息了。最终的结果是哪个粉红色的球体会从顶部开始向两侧“点燃”，然后作有规律的运动。这有一点迷幻，但是会帮助你对着色器中的三种变量有更好的了解：他们互相联系，实现了整个集合体。当然我们会在Three.js的框架中做这些。
1.模拟光照
让我们更新颜色吧，这样球体看起来就不会是个扁平晦暗的圆了。如果我们想看看Three.js是怎样处理光照的，我敢肯定你会发现这比我们需要的要复杂得多，所以我们先模拟光照吧。你应该浏览一下Three.js中那些奇妙的着色器，还有一些来自最近的一个 Chris Milk 和 Google, Rome 的WebGL项目。
回到着色器，我们要更新顶点着色器来向片元着色器传递顶点的法向量。利用一个varying变量：
复制代码 代码如下:

// 创建一个varying变量vNormal，顶点着色器和片元着色器都包含了该变量
varying vec3 vNormal;
void main() {
// 将vNormal设置为normal，后者是Three.js创建并传递给着色器的attribute变量
vNormal = normal;
gl_Position = projectionMatrix *
modelViewMatrix *
vec4(position, 1.0);
}

在片元着色器中，我们将会创建一个相同变量名的变量，然后将法线向量和另一个表示来自右上方光线的向量点乘，并将结果作用于颜色。最后结果的效果有点像平行光。
复制代码 代码如下:

// 和顶点着色器中一样的变量vNormal
varying vec3 vNormal;
void main() {
// 定义光线向量
vec3 light = vec3(0.5,0.2,1.0);
// 确保其归一化
light = normalize(light);
// 计算光线向量和法线向量的点积，如果点积小于0（即光线无法照到），就设为0
float dProd = max(0.0, dot(vNormal, light));
// 填充片元颜色
gl_FragColor = vec4(dProd, // R
dProd, // G
dProd, // B
1.0); // A
}

使用点积的原因是：两个向量的点积表明他们有多么“相似”。如果两个向量都是归一化的，而且他们的方向一模一样，点积的值就是1；如果两个向量的方向恰巧完全相反，点积的值就是-1。我们所做的就是把点积的值拿来作用到光纤上，所以如果这个点在球体的右上方，点积的值就是1，也就是完全照亮了；而在另一边的点，获得的点积值接近0，甚至到了-1。我们将获得的任何负值都设置为0。当你将数据传入之后，你就会看到最基本的光照效果了。

下面是什么？我们会将顶点的坐标掺和进来。
2.Attribut变量
接下来我要通过Attribute变量为每一个顶点传递一个随机数，这个随机数被用来将顶点沿着法线向量推出去一段距离。新的结果有点像一个怪异的不规则物体，每次刷新页面物体都会随机变化。现在，他还不会动（后面我会让他动起来），但是几次刷新就可以很好地观察到，他的形状是随机的。
让我们开始为顶点着色器加入attribute变量吧：
复制代码 代码如下:

attribute float displacement;
varying vec3 vNormal;
void main() {
vNormal = normal;
// 将随机数displacement转化为三维向量，这样就可以和法线相乘了
vec3 newPosition = position +
normal * vec3(displacement);
gl_Position = projectionMatrix *
modelViewMatrix *
vec4(newPosition, 1.0);
}

你看到什么都没变，因为attribute变量displacement还没有被设定你，所以着色器就使用了0作为默认值。这时displacement还没起作用，但我们马上就要在着色器材质中加上attribute变量了，然后Three.js就会自动地把它们绑在一起运行了。

同时也要注意这样一个事实，我将更新后的位置指定给了一个新的三维向量变量，因为原来的位置变量position，就像所有的attribute变量一样，都是只读的。
3.更新着色器材质
现在我们来更新着色器材质，传入一些东西给attribute对象displacement。记住，attribute对象是和顶点一一对应的，所以我们对球体的每一个顶点都有一个值，就像这样：
复制代码 代码如下:

var attributes = {
displacement: {
type: 'f', // 浮点数
value: [] // 空数组
}
};
var vShader = $('#vertexshader');
var fShader = $('#fragmentshader');
// 创建一个包含attribute属性的着色器材质
var shaderMaterial =
new THREE.MeshShaderMaterial({
attributes: attributes,
vertexShader: vShader.text(),
fragmentShader: fShader.text()
});
// 向displacement中填充随机数
var verts = sphere.geometry.vertices;
var values = attributes.displacement.value;
for(var v = 0; v < verts.length; v++) {
values.push(Math.random() * 30);
}

这样，就可以看到一个变形的球体了。最Cool的是：所有这些变形都是在GPU中完成的。
4.动起来
要使这东西动起来，应该怎么做？好吧，应该做这两件事情。
一个uniform变量amplitude，在每一帧控制displacement实际造成了多少位移。我们可以使用正弦或余弦函数来在每一帧中生成它，因为这两个函数的取值范围从-1到1。
一个帧循环。

我们需要将这个uniform变量加入到着色器材质中，同时也需要加入到顶点着色器中。先来看顶点着色器：
复制代码 代码如下:

uniform float amplitude;
attribute float displacement;
varying vec3 vNormal;
void main() {
vNormal = normal;
// 将displacement乘以amplitude，当我们在每一帧中平滑改变amplitude时，画面就动起来了
vec3 newPosition =
position +
normal *
vec3(displacement *
amplitude);
gl_Position = projectionMatrix *
modelViewMatrix *
vec4(newPosition, 1.0);
}

然后更新着色器材质：
复制代码 代码如下:

var uniforms = {
amplitude: {
type: 'f', // a float
value: 0
}
};
var vShader = $('#vertexshader');
var fShader = $('#fragmentshader');
// 创建最终的着色器材质
var shaderMaterial =
new THREE.MeshShaderMaterial({
uniforms: uniforms,
attributes: attributes,
vertexShader: vShader.text(),
fragmentShader: fShader.text()
});

我们的着色器也已经就绪了。但我们好像又倒退了一步，屏幕中又只剩下光滑的球了。别担心，这是因为amplitude值设置为0，因为我们将amplitude乘上了displacement，所以现在看不到任何变化。我们还没设置循环呢，所以amplitude只可能是0.

在我们的JavaScript中，需要将渲染过程打包成一个函数，然后用requestAnimationFrame去调用该函数。在这个函数里，我们更新uniform（译者注：即amplitude）的值。
复制代码 代码如下:

var frame = 0;
function update() {
// amplitude来自于frame的正弦值
uniforms.amplitude.value =
Math.sin(frame);
// 更新全局变量frame
frame += 0.1;
renderer.render(scene, camera);
// 指定下一次屏幕刷新时，调用update
requestAnimFrame(update);
}
requestAnimFrame(update);

5.小结
就是它了！你看到球体正在奇怪地脉动着。关于着色器，还有太多的内容没有讲到呢，但是我希望这篇教程能够对你有一些帮助。现在，当你看到一些其他的着色器时，我希望你能够理解它们，而且你应该有信心去创建自己的着色器了！


译序
Three.js是一个伟大的开源WebGL库，WebGL允许JavaScript操作GPU，在浏览器端实现真正意义的3D。但是目前这项技术还处在发展阶段，资料极为匮乏，爱好者学习基本要通过Demo源码和Three.js本身的源码来学习。
0.简介
嗨，又见面了。这么说我们已经开始学习Three.js了，如果你还没有看过之前三篇教程，建议你先读完。如果你已经读完前面的教程了，你可能会想做一些关于粒子的东西。让我们直面这个话题吧，每个人都爱粒子效果。不管你是否知道，你可以很轻易地创建它们。
1.创建一个粒子系统
Three.js将粒子系统视为一个基本的几何体，因为它就像基本几何体一样，即有形状，又有位置、缩放因子、旋转属性。粒子系统将geometry对象里的每一个点视为一个单独的粒子。为什么这样做？我想基于以下的原因：首先，整个粒子系统地绘制只需要调用一次某个绘制函数，而不是调用上千次；其次，这允许你设定一些全局的参数来影响你的粒子系统内的所有粒子。

即使是粒子系统被视为一个整体的对象，我们仍然可以为每个粒子单独地着色，因为在绘制粒子系统的过程中，Three.js通过attribute变量colour向着色器传递了每一个顶点的颜色。我在本篇教程里并不准备这样做，如果你想知道这是怎样完成的，你可以去GitHub上看Three.js的例程。

粒子系统可能还有一种特殊效果需要引起你的注意：Three.js在粒子系统第一次被渲染的时候，会将其数据缓存下来，之后你无法增加或减少系统中的粒子。如果你不希望看到某个粒子，你可以将它的颜色中的alpha值设置为0，但你无法删除它。所以你应当在创建粒子系统的时候，就将所有可能需要显示的粒子考虑进来。
开始创建一个粒子系统，只需要这么多：
复制代码 代码如下:

// 创建粒子geometry
var particleCount = 1800,
particles = new THREE.Geometry(),
pMaterial =
new THREE.ParticleBasicMaterial({
color: 0xFFFFFF,
size: 20
});
// 依次创建单个粒子
for(var p = 0; p < particleCount; p++) {
// 粒子范围在-250到250之间
var pX = Math.random() * 500 - 250,
pY = Math.random() * 500 - 250,
pZ = Math.random() * 500 - 250,
particle = new THREE.Vertex(
new THREE.Vector3(pX, pY, pZ)
);
// 将粒子加入粒子geometry
particles.vertices.push(particle);
}
// 创建粒子系统
var particleSystem =
new THREE.ParticleSystem(
particles,
pMaterial);
// 将粒子系统加入场景
scene.addChild(particleSystem);

如果你运行：
1.你会发现粒子都是方的
2.粒子都不动
我们一个一个来修复。
2.风格
我们创建一个粒子基本材质时传入了颜色和尺寸。我们可能想做的是传入一张纹理图片用来显示粒子，而这样就可以很好地控制粒子看上去的样式了。

你也看到，粒子是以方块形状绘制的，所以我们也应当使用一张方形的纹理图片。为了看上去效果更好，我还会使用加法混合，但是这样做必须保证纹理图片的背景是黑色的而不是透明的。我理解的原因是：现在加法混合和透明材质之间不兼容。但是没关系，最后看上去会很棒。

我们来更新一下粒子基本材质和粒子系统，加入一些加法混合下透明的粒子。如果你喜欢，你也可以用我的粒子图片。
复制代码 代码如下:

// 创建粒子基本材质
var pMaterial =
new THREE.ParticleBasicMaterial({
color: 0xFFFFFF,
size: 20,
map: THREE.ImageUtils.loadTexture(
"images/particle.png"
),
blending: THREE.AdditiveBlending,
transparent: true
});
// 允许粒子系统对粒子排序，以达到我们想要的效果
particleSystem.sortParticles = true;

这看上去已经好多了。现在来引入一点物理，让粒子们动起来。
3.引入物理
默认情况下，粒子系统在三维空间中不运动，这很好。但我想让他们动起来，而且我要让粒子系统这样运动：让粒子绕着y轴旋转。而且粒子在每个轴的范围都在-250到250之间，所以绕着y轴旋转以为这它们绕着系统地中心旋转。

我还假定，你已经在某个地方有了帧循环的代码，和我在上一篇关于着色器中的教程中类似。所以这里我们只需这样：
复制代码 代码如下:

// 帧循环
function update() {
// 增加一点旋转量
particleSystem.rotation.y += 0.01;
// 绘制粒子系统
renderer.render(scene, camera);
// 设置下一次刷新帧时对update的调用
requestAnimFrame(update);
}

现在我们开始定义单个粒子的运动（译者注：之前的旋转是整个粒子系统的运动）。我们来做个简单的雨点效果，这包含一下几步：
1.给每一个粒子赋一个初始为0的速度
2.在每一帧中，为每一个粒子赋一个随机的重力加速度
3.在每一帧中，通过通过加速度更新速度，通过速度更新位置
4.当一个粒子运动出了视线，重新设置初始位置和速度
听上去很多，其实代码写起来很少。首先，在创建粒子的过程中，我们为每个粒子增加一个水平速度：
复制代码 代码如下:

// 为每个粒子创建一个水平运动速度
particle.velocity = new THREE.Vector3(
0, // x
-Math.random(), // y: 随机数
0); // z

接下来，在帧缓冲中我们传递每个粒子，并且，当粒子离开屏幕底部需要重置时，重置其位置和速度。
复制代码 代码如下:

// 帧循环
function update() {
// 增加旋转量
particleSystem.rotation.y += 0.01;
var pCount = particleCount;
while(pCount--) {
// 获取单个粒子
var particle = particles.vertices[pCount];
// 检查是否需要重置
if(particle.position.y < -200) {
particle.position.y = 200;
particle.velocity.y = 0;
}
// 用随机数更新水平速度分量，并根据速度更新位置
particle.velocity.y -= Math.random() * .1;
particle.position.addSelf(
particle.velocity);
}
// 告诉粒子系统我们改变了粒子位置
particleSystem.geometry.__dirtyVertices = true;
// 画
renderer.render(scene, camera);
// 设置下一次调用
requestAnimFrame(update);
}

虽然不够震撼，但这个粒子至少展示了如何做。你完全应该自己创建一些美妙的粒子效果，然后让我知道。
这里有个警告你应该知道，在帧循环中，我越雷池了：我在一次循环中遍历了所有粒子，这实际上是种很粗放的方式。如果你的帧循环中做了太多的工作（译者注：注意帧循环的js代码是在cpu中运行的，它不像gpu，能一下子并发出成千上万个简单进程），浏览器就会卡顿，事实上如果你用了requestAnimationFrame，它视图每秒刷新60次。所以还是优化你的代码，在帧循环中做尽量少的事情。
4.小结
粒子效果太棒了，是个人都爱粒子效果，而现在你知道如何在Three.js中加入粒子效果了。我希望你能用得顺手，就跟前面一样！ 



 Three.js是一个伟大的开源WebGL库，WebGL允许JavaScript操作GPU，在浏览器端实现真正意义的3D。但是目前这项技术还处在发展阶段，资料极为匮乏，爱好者学习基本要通过Demo源码和Three.js本身的源码来学习。
0.简介
嗨，这是我的第一篇关于如何写出好的代码的文章。和很多开发者一样，我通过实践学习，但同时我也向其他更有经验的开发者们学习。在过去的几个月中，我在canvas标签上花了很多时间，我想如果把这段时间学到的关于WebGL和JavaScript的小技巧都写下来，一定很有意思。有一些很具体，有一些却很笼统，希望你们喜欢！
1.尽快写一个原型
让我们从简单的开始。现在你有个绝妙的注意，那么你应该尽快就程序里最复杂的部分写一个原型，看看这项技术是否可以实现你的想法。WebGL很强大，因为它可以直接操纵显卡里的GPU，但是也别忘了你需要通过JavaScript才能访问显卡，这比显卡内部运算的效率可是低多了。事实上，你的天才想法很可能被这种简单事情击败。
2.使用THREE.JS处理3D
就像我的朋友Hakim一样，我也完全理解我们正使用的技术的底层细节。理解表面之下的东西是很重要的，但是如果你使用three.js，它为你免去了如此多的烦恼。你可以将它用于Canvas，WebGL还有SVG，你也应该找到哪种方式合适你的需求。
3.避免SetInterval
这对所有使用JavaScript创建动画的人来说，都是很重要的一点。为什么？假定你设定每20毫秒后执行一次某个函数，而这个函数需要执行超过20毫秒的时间，那么20ms之后，浏览器不会在乎，而是直接开始下一次执行。至少你可以使用SetTimeout来设置，在某个函数执行完之后，再次执行它。

事实上，有一个更加新潮却还是半成品的函数，叫做requestAnimationFrame，它很棒。它很类似于setTimeOut函数，除了在这两个方面：当标签页失去焦点时，它就不再运行了；现在这个函数还是依赖于浏览器的，标准以后还有可能变化。如果你想要更多的信息，可以访问Paul Irish的博客。
4.使用倒序循环
这是个不错的小技巧，可以让你的循环更快。使用倒序，而且使用while循环。比如，这个循环：
复制代码 代码如下:

for(var a = 0; a < arr.length; a++) {
// 做一些什么
}

它的执行效率不如下面这个循环：
复制代码 代码如下:

// 假设数组arr存在
var aLength = arr.length;
while(aLength--) {
// 做一些什么
}

这可能没帮你省多少开销，因为执行的效率主要还是依赖于你在循环体里面干了什么。但如果你想程序的不必要开销减少到最后一个字节，后一个循环肯定赢。

实话说，主要影响程序执行效率的还是数组缓存的长度。你可以（也确实应该）去看看JSPerf去了解这一点，以及其他影响JavaScript性能的因素。
5.使用纹理
在WebGL里面把物体的任意一个细节都画出来看上去很诱人，但是，如果有可能的话，你应当注意一下你是否能够使用纹理，因为它能够极大地提高性能。在某些特定的情况下，比如阴影或者模糊效果，你也许不得不使用纹理，但在其他时候，你也应该时时关注你是否可以使用纹理。
6.使用缓存
这一点我在自己的试验力试了很多，在帧循环中，你应当避免引用变量、对象或者其他任何东西。基于这点原因，很值得把你的模型、顶点全部缓存起来，这样在渲染动画的时候你就可以快速地访问到它们。
7.禁用选中
我爱这一小段代码，我把它放到任何包含Canvas或WebGL的页面中。
复制代码 代码如下:

// 禁用鼠标选中DOM元素
document.onselectstart = function() {
return false;
};

你也可能只想在Canvas控件中禁用选中。这段是我在那些Canvas占据了整个屏幕的项目中使用的代码。
8.避免在JavaScript中定义CSS
现在，在JavaScript中定义CSS简直太方便了，尤其是你使用JQuery的时候
复制代码 代码如下:

// 尽量不要这样做
$("#someid").css({
position: 'relative',
height: '30px',
width: '300px',
backgroundColor: '#A020F0'
});

问题是这样做之后，你的JavaScript代码中很快就充斥着各种类型的CSS定义，而你同时又使用*.css文件来定义CSS，潜在的问题很难被发现。更好的方法是：使用class模块化CSS，而且只在JavaScript中定义那些不能预知的CSS类。
9.在对象中定义回调函数
我爱下面这段代码，这绝不是我自己想出来的，但它是如此整洁美观。如果你有一大堆回调函数要用，你也许会这样用的：
复制代码 代码如下:

$("#someid").click(function() {
// 回调函数
// 返回false在JQuery中会阻止消息的传递和默认行为的放生
return false;
});

或者，你会回调一个在代码其他地方定义的松散的函数，比如这样
复制代码 代码如下:

$("#someid").click(mySuperFunction);
function mySuperFunction(event) {
// 在这里做很多事情
return false;
}

这样做会有一些问题。第一段代码中，你在某个事件上绑定了匿名函数，你很难将该函数再从事件上解除下来。你当然可以解除某个事件上的所有函数，但你可能在它上面绑定了多个函数，而你只想解除一个。在第二种情况下，你的函数名污染了全局变量空间，代码的可维护性降低了。所以，考虑这样做：
复制代码 代码如下:

$("#someid").click(callbacks.mySuperFunction);
// 所有的回调函数都在callbacks对象中
var callbacks = {
mySuperFunction:function(event) {
// 更多地工作
return false;
}
}
// 解除某个函数的绑定
$("#someid").unbind('click', callbacks.mySuperFunction);

这样做整洁又干净，而且避免了上面提到的两个问题。
10.链式三元运算符
我完全是从Paul Irish的《JQuery，你应该知道的11件事》中学到这个的。这非常好用，你也应该会喜欢。我们经常这样做：
复制代码 代码如下:

// 根据a的值为numberBasedOnA赋值
// 如果a大于5，则赋值200，否则赋值38
var numberBasedOnA = a > 5 ? 200 : 38;

但如果你想这样做，比如，当值为多少时如何，当值大于多少时如何，当值更加大的时候如何，明白吗？在这种情况下，链式三元运算符非常好用：
复制代码 代码如下:

var numberBasedOnA =
a < 5 ? 200 :
a < 7 ? 38 :
a < 11 ? 15 :
a < 15 ? 49 :
64;
// 比这样做更有效率
// when a >=15 


译序
Three.js是一个伟大的开源WebGL库，WebGL允许JavaScript操作GPU，在浏览器端实现真正意义的3D。但是目前这项技术还处在发展阶段，资料极为匮乏，爱好者学习基本要通过Demo源码和Three.js本身的源码来学习。
0.简介
全景图非常酷。使用Three.js做一个属于自己的全景图并不是那么困难。
要做一个全景图，你需要一个软件用来做一张全景图片（译者注：如果你没有那些特殊的设备）。我使用了iPhone上的Microsoft Photosynth软件来制作。
1.环境纹理
首先什么是环境纹理？在WebGL或者OpenGL中他们实际上是种特殊的立方体纹理。一个立方体纹理是对整个场景（虚拟的或现实的）的观察，场景的样子被“贴”在了立方体的内部表面。想象一下，你站在山顶，向前看，向左看，向右看，向上看，向下看，最后向后看。每一次你都看到了这个“立方体”的内部表面，你就站在这个立方体的中心。如果这个立方体足够大，就很难分辨出立方体的棱和角，而给你一种错觉：你处在一个很大的环境里面。如果你还没弄明白，那么维基百科上的cube maps条目会非常有帮助。
这很酷，但是这怎么用？我们可以像做反射和折射一样，而且事实上这两者的函数都已经内建在GLSL，WebGL的着色器语言上了。你只需要传递给着色器6张纹理图片，每张代表立方体的一个内表面，然后告诉WebGL这是个立方体纹理，然后就可以使用上面的效果了。
半轴：这个术语服务于立方体纹理。因为我们通常使用三个轴来描述三维空间：x轴、y轴、z轴，所以用于立方体纹理的图片也用轴的名称来标识了。一共六张图片，每个轴两张图片，正半轴一个，负半轴一个。
2.创建全景图片
创建全景图片的第一步就是走出户外，使用手机上的应用来照一张。我在伦敦的金融区转了一圈，然后在Gherkin照了一张。我获得了下面这张图片：
 
值得指出的是，这个应用将图片做成了贴到球体上的那种。这看上去不错，但是我们现在需要将它贴到一个立方体的内表面上，所以还要处理一下这张图。
3.转化到立方体上
这部分我简短介绍一下。我把刚才获得的那张照片载入到一个3D建模软件中，比如Maya或者Blender，然后将其粘贴到一个球体的内表面上。然后我创建了6个正射投影的相机，每一个都对应于一个半轴。最后我将这6个相机捕捉到的图像保存了下来。具体怎么完成比较复杂，也没必要在这里讲解，所以我写了一个Blender脚本文件，所有的一切都设置好了。

使用这个脚本文件你只需要：
1.将你自己的全景图重命名为 environment.jpg；
2.将全景图和Blender脚本文件放在同一个文件夹下；
3.载入脚本文件；
4.点击右侧的 Animation 按钮；
5.等一会儿，6张图像已经创建好了。
很Cool吧？现在你可以重新命名这些图像，使之与每一个半轴相匹配。比如这样：
•0001.png → pos-z.png
•0002.png → neg-x.png
•0003.png → neg-z.png
•0004.png → pos-x.png
•0005.png → neg-y.png
•0006.png → pos-y.png
4.加入场景
现在我们已经获得了环境纹理，然后将其载入到场景中。Three.js使这变得非常简单：
复制代码 代码如下:

// 纹理图像的url
var urls = [
'path/to/pos-x.png',
'path/to/neg-x.png',
'path/to/pos-y.png',
'path/to/neg-y.png',
'path/to/pos-z.png',
'path/to/neg-z.png'
],
// 打包成我们需要的对象
cubemap = THREE.ImageUtils.
loadTextureCube(urls);
// 设置格式为RGB
cubemap.format = THREE.RGBFormat;

现在只需要将cubemap指定到一个材质中去就可以了！
复制代码 代码如下:

var material = new THREE
.MeshLambertMaterial({
color: 0xffffff,
envMap: cubemap
});

5.小结
就这样了，实现一个全景图很酷，尤其是你可以将你自己的地方制作为WebGL全景图。和往常一样，我打包了这次教程的源码 