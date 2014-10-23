# directives 

[source](http://www.html-js.com/article/A-practical-guide-to-the-development-of-web-application-written-using-Angular-AngularJS-instruction-a)

指令（directives）是任何AngularJS应用中最重要的成分。尽管AngularJS已经自带了很多指令，你经常会发现需要自己亲手创建一些特别的指令。本文将会带你了解自定义指令并解释如何在现实世界中的Angular项目中使用它们。文章的最后，我们将一起用Angular指令来创建一个简单的笔记小应用。   

# 综述

一个指令就是一个引入新语法的东西。指令是在DOM元素上做的标记，并同时附加了一些特定的行为。例如，静态的HTML并不知道如何来创建并显示一个日期选择插件。为了将这个新语法教给HTML我们需要一条指令。这个指令将会创建一个充当日期选择器的元素。我们将在随后看到如何实现这个指令。   

如果你之前已经编写过Angular应用，那么你已经使用过指令了，不管你有没有意识到这点。你可能已经使用过像是<code class=" language-javascript">ng<span class="token operator">-</span>model</code>，<code class=" language-javascript">ng<span class="token operator">-</span>repeat</code>，<code class=" language-javascript">ng<span class="token operator">-</span>show</code>等等这样的指令。所有这些指令都将特定的功能绑定到了DOM元素之上。例如，<code class=" language-javascript">ng<span class="token operator">-</span>repeat</code>会重复特定的元素，而<code class=" language-javascript">ng<span class="token operator">-</span>show</code>会有条件的展示元素。如果你想要创建一个可拖动元素的话你可能需要创建一个指令。指令背后的基本思想很简单。它通过在元素上绑定事件监听器并且将DOM变形来使HTML变得具有交互性。   

# 从jQuery的角度来看指令

想想你如何使用jQuery来创建一个日期选择器。我们首先在HTML中添加一个普通的input字段然后在jQuery中我们调用$(element).dataPicker()来将其转换为一个日期选择器。但是，考虑一下。当一个设计师想要来检查这个标记时，他/她能够立刻猜出这个字段究竟是干什么用的吗？它仅仅是一个普通的input字段还是一个日期选择器？你必须要查看jQuery来确认这点。Angular的方法是使用指令来扩展HTML。因此，一个日期选择器的指令看上去可能如下所示：   

<code class=" language-javascript"><span class="token operator">&lt;</span>date<span class="token operator">-</span>picker<span class="token operator">&gt;</span><span class="token operator">&lt;</span><span class="token operator">/</span>date<span class="token operator">-</span>picker<span class="token operator">&gt;</span> 
</code>
或者如下所示：   

<code class=" language-javascript"><span class="token operator">&lt;</span>input type<span class="token operator">=</span><span class="token string">'text'</span> data<span class="token operator">-</span>picker<span class="token operator">/</span><span class="token operator">&gt;</span> 
</code> <br>
这种创建UI成分的方法既直观又清楚。你可以看到元素就知道它的用途。   

# 创建自定义指令

一个Angular指令可能以四种形式出现：   

1.一个新的HTML元素（<code class=" language-javascript"><span class="token operator">&lt;</span>date<span class="token operator">-</span>picker<span class="token operator">&gt;</span><span class="token operator">&lt;</span><span class="token operator">/</span>date<span class="token operator">-</span>picker<span class="token operator">&gt;</span></code>） <br>
 2.一个元素上的属性（<code class=" language-javascript"><span class="token operator">&lt;</span>input type<span class="token operator">=</span><span class="token string">'text'</span> date<span class="token operator">-</span>picker<span class="token operator">/</span><span class="token operator">&gt;</span></code>） <br>
 3.作为一个类（<code class=" language-javascript"><span class="token operator">&lt;</span>input type<span class="token operator">=</span><span class="token string">'text'</span> class<span class="token operator">=</span><span class="token string">'date-picker'</span><span class="token operator">/</span><span class="token operator">&gt;</span></code>） 
 4.作为注释（<code class=" language-javascript"><span class="token operator">&lt;</span><span class="token operator">!</span><span class="token operator">--</span>directive<span class="token punctuation">:</span>date<span class="token operator">-</span>picker<span class="token operator">--</span><span class="token operator">&gt;</span></code>）     

当然，我们完全可以决定我们的指令以什么形式出现在HTML中。现在，我们来看看一个典型的Angular指令是如何写成的。它和controller的注册方式类似，但是它会返回一个简单的对象（指令定义），其中那个包含有一些配置指令的属性。下面的代码展示了一个简单和Hello World指令：   

<pre class=" language-javascript"><code class=" language-javascript"><span class="token keyword">var</span> app <span class="token operator">=</span> angular<span class="token punctuation">.</span><span class="token function">module<span class="token punctuation">(</span></span><span class="token string">'myapp'</span><span class="token punctuation">,</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token punctuation">)</span><span class="token punctuation">;</span>   

app<span class="token punctuation">.</span><span class="token function">directive<span class="token punctuation">(</span></span><span class="token string">'helloWorld'</span><span class="token punctuation">,</span><span class="token keyword">function</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">{</span>
    <span class="token keyword">return</span> <span class="token punctuation">{</span>
        restrict<span class="token punctuation">:</span> <span class="token string">'AE'</span><span class="token punctuation">,</span>
        replace<span class="token punctuation">:</span> <span class="token boolean">true</span><span class="token punctuation">,</span>
        template<span class="token punctuation">:</span> <span class="token string">'&lt;h3&gt;Hello World!&lt;/h3&gt;'</span>
    <span class="token punctuation">}</span>
<span class="token punctuation">}</span><span class="token punctuation">)</span><span class="token punctuation">;</span>   
</code></pre>

在上面的代码中，app.diretive()函数在我们的模块中注册了一个新的指令。这个函数的第一个参数是指令的名称。第二个参数是一个返回指令定义对象的函数。如果你的指令对额外的对象/服务(services)例如 $rootScope, $http 或者 $compile 有依赖，它们也可以在其中被注入。这个指令可以作为一个HTML元素来使用，如下所示：   

<pre class=" language-javascript"><code class=" language-javascript"><span class="token operator">&lt;</span>hello<span class="token operator">-</span>world<span class="token operator">/</span><span class="token operator">&gt;</span>   
</code></pre>

或者：   

<pre class=" language-javascript"><code class=" language-javascript"><span class="token operator">&lt;</span>hello<span class="token punctuation">:</span>world<span class="token operator">/</span><span class="token operator">&gt;</span>  
</code></pre>

或者作为一个属性来使用：   

<pre class=" language-javascript"><code class=" language-javascript"><span class="token operator">&lt;</span>div hello<span class="token operator">-</span>world<span class="token operator">&gt;</span><span class="token operator">&lt;</span><span class="token operator">/</span>div<span class="token operator">&gt;</span>   
</code></pre>

或者：   

<pre class=" language-javascript"><code class=" language-javascript"><span class="token operator">&lt;</span>div hello<span class="token punctuation">:</span>world<span class="token operator">/</span><span class="token operator">&gt;</span>   
</code></pre>

如果你想要兼容HTML5，你可以在属性前面加上x-或者data-前缀。因此，下面的标记将会匹配helloWorld指令：   

<pre class=" language-javascript"><code class=" language-javascript"><span class="token operator">&lt;</span>div data‐hello‐world<span class="token operator">&gt;</span><span class="token operator">&lt;</span><span class="token operator">/</span>div<span class="token operator">&gt;</span>    
</code></pre>

或者   

<pre class=" language-javascript"><code class=" language-javascript"><span class="token operator">&lt;</span>di vx‐hello‐world<span class="token operator">&gt;</span><span class="token operator">&lt;</span><span class="token operator">/</span>div<span class="token operator">&gt;</span>      
</code></pre>

## 注意

当匹配指令时，Angular会从元素/属性名之前去除前缀x-或者data-。然后将分隔符 - 或者 : 转换为驼峰表示法已匹配注册的指令。这就是为什么我们的helloWorld指令用在HTML中的时候实际上写成了hello-world。    

尽管上面的这个简单的指令仅仅只是展示了一些静态的文本，其中还是有一些值得我们去探究的有趣的点。我们已经在这个指令定义对象中使用了三个属性。我们来看看这三个属性分别都有什么用：   

<ul>
<li>restrict - 这个属性指明了一个指令应该如何在HTML中使用（记住指令可以以四种方式出现）。在这个例子中我们将它设置为’AE’。因此，这条指令可以作为一个HTML元素或者一个属性来使用。为了允许指令作为一个类来使用我们可以将restrict设置为’AEC’。   </li>
<li>template - 这个实行指明了当指令被Angular编译和链接时生成的HTML标记。它不一定是一个简单的字符串。template可以很复杂，其中经常会涉及其它的指令，表达式（{{}}），等等。在大多数情况下你可能会想要使用templateUrl而不是template。因此，理想情况下你应该首先将模板放置在一个单独的HTML文件中然后让templateUrl指向它。   </li>
<li>replace - 这个属性指明了是否生成的模板会代替绑定指令的元素。在前面的例子中我们在HTML中使用指令为<code class=" language-javascript"><span class="token operator">&lt;</span>hello<span class="token operator">-</span>world<span class="token operator">&gt;</span><span class="token operator">&lt;</span><span class="token operator">/</span>hello<span class="token operator">-</span>world<span class="token operator">&gt;</span></code>，并将replace属性设置为true。因此，在指令编译后，生成的模板代替了<code class=" language-javascript"><span class="token operator">&lt;</span>hello<span class="token operator">-</span>world<span class="token operator">&gt;</span><span class="token operator">&lt;</span><span class="token operator">/</span>hello<span class="token operator">-</span>world<span class="token operator">&gt;</span></code>。最后的输出结果是<code class=" language-javascript"><span class="token operator">&lt;</span>h3<span class="token operator">&gt;</span>Hello World<span class="token operator">!</span><span class="token operator">&lt;</span><span class="token operator">/</span>h3<span class="token operator">&gt;</span></code>。如果你将replace设置为false，默认情况下，输出模板将会被插入到指令被调用的元素中。    </li>
</ul>

# link函数和作用域

有一个指令生成的模板是没有用的除非它在正确的作用域中北编译。默认情况下一个指令并不会得到一个新的子作用域。然而，它可以得到父作用域。这意味着如果一个指令位于在一个控制器中那么它将使用控制器的作用域。   

为了利用作用域，我们可以使用一个叫做link的函数。它可以通过指令定义对象中的link属性来配置。我们现在对helloworld指令做一些修改一遍当用户在一个input字段中输入一个颜色名称时，Hello Wolld文字的背景颜色会自动发生改变。同样，当一个用户点击Hello World文字时，背景颜色会重置为白色。相应的HTML标记如下所示：   

<pre class=" language-javascript"><code class=" language-javascript"><span class="token operator">&lt;</span>body ng<span class="token operator">-</span>controller<span class="token operator">=</span><span class="token string">'MainCtrl'</span><span class="token operator">&gt;</span>   
    <span class="token operator">&lt;</span>input type<span class="token operator">=</span><span class="token string">'text'</span> ng<span class="token operator">-</span>model<span class="token operator">=</span><span class="token string">'color'</span> placeholder<span class="token operator">=</span><span class="token string">'Enter a color'</span> <span class="token operator">/</span> <span class="token operator">&gt;</span>
    <span class="token operator">&lt;</span>hello<span class="token operator">-</span>wolrd<span class="token operator">/</span><span class="token operator">&gt;</span>   
<span class="token operator">&lt;</span><span class="token operator">/</span>body<span class="token operator">&gt;</span>
</code></pre>

修改后的helloWorld指令代码如下所示：   

<pre class=" language-javascript"><code class=" language-javascript">app<span class="token punctuation">.</span><span class="token function">directive<span class="token punctuation">(</span></span><span class="token string">'helloWorld'</span><span class="token punctuation">,</span><span class="token keyword">function</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">{</span>
    <span class="token keyword">return</span> <span class="token punctuation">{</span>
        restrict<span class="token punctuation">:</span> <span class="token string">'AE'</span><span class="token punctuation">,</span>   
        replace<span class="token punctuation">:</span> <span class="token boolean">true</span><span class="token punctuation">,</span>
        template<span class="token punctuation">:</span> <span class="token string">'&lt;p style="background-color:{{color}}"&gt;&lt;/p&gt;'</span><span class="token punctuation">,</span>   
        link<span class="token punctuation">:</span> <span class="token keyword">function</span><span class="token punctuation">(</span>scope<span class="token punctuation">,</span>elem<span class="token punctuation">,</span>attr<span class="token punctuation">)</span><span class="token punctuation">{</span>
            elem<span class="token punctuation">.</span><span class="token function">bind<span class="token punctuation">(</span></span><span class="token string">'click'</span><span class="token punctuation">,</span><span class="token keyword">function</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">{</span>
                elem<span class="token punctuation">.</span><span class="token function">css<span class="token punctuation">(</span></span><span class="token string">'background-color'</span><span class="token punctuation">,</span><span class="token string">'white'</span><span class="token punctuation">)</span><span class="token punctuation">;</span>
            scope<span class="token punctuation">.</span>$<span class="token function">apply<span class="token punctuation">(</span></span><span class="token keyword">function</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">{</span>
                scope<span class="token punctuation">.</span>color <span class="token operator">=</span> <span class="token string">"white"</span><span class="token punctuation">;</span>
            <span class="token punctuation">}</span><span class="token punctuation">)</span><span class="token punctuation">;</span>
            <span class="token punctuation">}</span><span class="token punctuation">)</span><span class="token punctuation">;</span>
            elem<span class="token punctuation">.</span><span class="token function">bind<span class="token punctuation">(</span></span><span class="token string">'mouseover'</span><span class="token punctuation">,</span><span class="token keyword">function</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">{</span>
                elem<span class="token punctuation">.</span><span class="token function">css<span class="token punctuation">(</span></span><span class="token string">'cursor'</span><span class="token punctuation">,</span><span class="token string">'pointer'</span><span class="token punctuation">)</span><span class="token punctuation">;</span>
            <span class="token punctuation">}</span><span class="token punctuation">)</span><span class="token punctuation">;</span>
        <span class="token punctuation">}</span>
    <span class="token punctuation">}</span>
<span class="token punctuation">}</span><span class="token punctuation">)</span><span class="token punctuation">;</span>   
</code></pre>

注意到link函数被用在了指令中。它接收三个参数：   

<ul>
<li>scope - 它代表指令被使用的作用域。在上面的例子中它等同于符控制器的作用域。   </li>
<li>elem - 它代表绑定指令的元素的jQlite（jQuery的一个自己）包裹元素。如果你在AngularJS被包含之前就包括了jQuery，那么它将变成jQuery包裹元素。由于该元素已经被jQuery/jQlite包裹，我们没有必要将它包含在$()中来进行DOM操作。   </li>
<li>attars - 它代表绑定指令的元素上的属性。例如，如果你在HTML元素上有一些指令形式为：&lt;hello-world some-attribute&gt;&lt;/hello-world&gt;，你可以在link函数内用attrs.someAttribute来引用这些属性。   </li>
</ul>

link函数主要是用来对DOM元素绑定事件监听器，监视模型属性变化，并更新DOM。在前面的指令代码中，我们绑定了两个监听器，click和mouseover。click处理函数重置了的背景颜色，而mouseover处理函数则将游标改变为pointer。模板中拥有表达式{{color}}，它将随着父作用域中的模型color的变化而变化，从而改变了Hello World的背景色。   

# Compile函数

Compile函数主要用来在link函数运行之前进行一些DOM转化。它接收下面几个参数：   

<ul>
<li>tElement - 指令绑定的元素   </li>
<li>attrs - 元素上声明的属性   </li>
</ul>

这里要注意compile不能够访问scope，而且必须返回一个link函数。但是，如果没有compile函数以依然可以配置link函数。compile函数可以被写成下面的样子：   

<pre class=" language-javascript"><code class=" language-javascript">app<span class="token punctuation">.</span><span class="token function">directive<span class="token punctuation">(</span></span><span class="token string">'test'</span><span class="token punctuation">,</span><span class="token keyword">function</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">{</span>
    <span class="token keyword">return</span> <span class="token punctuation">{</span>
        compile<span class="token punctuation">:</span> <span class="token keyword">function</span><span class="token punctuation">(</span>tElem<span class="token punctuation">,</span>attrs<span class="token punctuation">)</span><span class="token punctuation">{</span>
           <span class="token comment" spellcheck="true"> //在这里原则性的做一些DOM转换   
</span>            <span class="token keyword">return</span> <span class="token keyword">function</span><span class="token punctuation">(</span>scope<span class="token punctuation">,</span>elem<span class="token punctuation">,</span>attrs<span class="token punctuation">)</span><span class="token punctuation">{</span>
            <span class="token comment" spellcheck="true"> //这里编写link函数
</span>            <span class="token punctuation">}</span>
        <span class="token punctuation">}</span>
    <span class="token punctuation">}</span>
<span class="token punctuation">}</span><span class="token punctuation">)</span><span class="token punctuation">;</span>   
</code></pre>

大多数时候，你仅仅只需要编写link函数。这是因为大部分指令都只关心与注册事件监听器，监视器，更新DOM等等，它们在link函数中即可完成。像是ng-repeat这样的指令，需要多次克隆并重复DOM元素，就需要在link函数运行之前使用compile函数。你可能会问威慑呢么要将两个函数分别使用。为什么我们不能只编写一个函数？为了回答这个问题我们需要理解Angular是如何编译指令的！   

# 指令是如何被编译的

当应用在启动时，Angular开始使用$compile服务解析DOM。这项服务会在标记中寻找指令然后将它们各自匹配到注册的适龄。一旦所有的指令都已经被识别完成，Angular就开始执行它们的compile函数。正如前面所提到的，compile函数返回一个link函数，该函数会被添加到稍后执行的link函数队列中。这叫做编译阶段(compile phase)。注意到即使同一个指令有几个实例存在，compile函数也只会运行一次。   

在编译阶段之后就到了链接阶段(link phase)，这时link函数就一个接一个的执行。在这个阶段中模板被生成，指令被运用到正确的作用域，DOM元素上开始有了事件监听器。不像是compile函数，lin函数会对每个指令的实例都执行一次。   

# 改变指令的作用域

默认情况下指令应该访问父作用域。但是我们并不像对所有情况一概而论。如果我们对指令暴露了父控制器的scope，那么指令就可以自由的修改scope属性。在一些情况下你的指令可能想要添加一些只有内部可以使用的属性和函数。如果我们都在父作用域中完成，可能会污染了父作用域。因此，我们有两种选择：   

<ul>
<li>一个子作用域 - 这个作用域会原型继承父作用域。   </li>
<li>一个隔离的作用域 - 一个全新的、不继承、独立存在的作用域。   </li>
</ul>

作用域可以由指令定义对象中的scope属性定义。下面的例子展示了这一点：   

<pre class=" language-javascript"><code class=" language-javascript">app<span class="token punctuation">.</span><span class="token function">directive<span class="token punctuation">(</span></span><span class="token string">'helloWorld'</span><span class="token punctuation">,</span><span class="token keyword">function</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">{</span>
    <span class="token keyword">return</span> <span class="token punctuation">{</span>
        scope<span class="token punctuation">:</span> <span class="token boolean">true</span><span class="token punctuation">,</span><span class="token comment" spellcheck="true"> //使用一个继承父作用域的自作用域   
</span>        restrict<span class="token punctuation">:</span> <span class="token string">'AE'</span><span class="token punctuation">,</span>
        replace<span class="token punctuation">:</span> <span class="token boolean">true</span><span class="token punctuation">,</span>
        template<span class="token punctuation">:</span> <span class="token string">'&lt;h3&gt;Hello World!&lt;/h3&gt;'</span>
    <span class="token punctuation">}</span>
<span class="token punctuation">}</span><span class="token punctuation">)</span><span class="token punctuation">;</span>   
</code></pre>

上面的代码要求Angular为指令提供一个能够原型继承父作用域的子组用于。另一种情形，一个隔离作用域，代码如下所示：   

<pre class=" language-javascript"><code class=" language-javascript">app<span class="token punctuation">.</span><span class="token function">directive<span class="token punctuation">(</span></span><span class="token string">'helloWorld'</span><span class="token punctuation">,</span><span class="token keyword">function</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">{</span>
    <span class="token keyword">return</span> <span class="token punctuation">{</span>
        scope<span class="token punctuation">:</span> <span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">,</span><span class="token comment" spellcheck="true"> //使用一个全新的隔离作用域   
</span>        restrict<span class="token punctuation">:</span> <span class="token string">'AE'</span><span class="token punctuation">,</span>
        replace<span class="token punctuation">:</span> <span class="token boolean">true</span><span class="token punctuation">,</span>
        template<span class="token punctuation">:</span> <span class="token string">'&lt;h3&gt;Hello World!&lt;/h3&gt;'</span>
    <span class="token punctuation">}</span>
<span class="token punctuation">}</span><span class="token punctuation">)</span><span class="token punctuation">;</span>
</code></pre>

上面的指令使用一个不继承父作用域的全新隔离作用域。当你想要创建一个可重用的组件时隔离作用域是一个很好的选择。通过隔离作用域我们确保指令是自包含的兵可以轻松地插入到任何HTML app中。这种做法防止了父作用域被污染，由于它不可访问父作用域。在我们修改后的helloWorld指令中如果你将scope设置为{}，那么代码就不会再正常运行。它将创建一个隔离的作用域然后表达式{{color}}将无法引用隔离作用域中的属性因此值变为undefined。     

隔离作用域并不意味着你一点都不能获取到父作用域中的属性。有一些技巧可以使你访问父作用域中的属性同时监听这些属性的变化。我们将在下一篇文章中提到这种高级技巧。     

<hr>

本文译自A Practical Guide to AngularJS Directives，原文地址<a href="http://www.sitepoint.com/series/a-practical-guide-to-angularjs-directives/">http://www.sitepoint.com/series/a-practical-guide-to-angularjs-directives/</a>    

<div class="entry-content"><img src="http://htmljs.b0.upaiyun.com/uploads/1392950189422-angular_0.jpg" alt="enter image description here" title="">
在上一篇文章中我们学到了一些关于AngularJS指令的基础知识。在结尾处我们也学到了如何为指令创建一个隔离作用域。本文将接着上一篇文章的结尾向下讲。首先，我们将学习如何从一个隔离作用域中获取父作用域的属性。然后，我们将通过探索controller函数以及transclusions来讨论如何为一个指令选择正确的作用域。在本文中，我们也会编写一个简单的笔记应用。   

# 隔离空间和父作用域属性间的绑定

通常情况下，隔离一个指令的作用域是很方便的，尤其是当你在操纵许多作用域模型的时候。但是，你也可能想要在指令中获取一些父作用域的属性。好消息是Angular可以通过绑定来让你有选择的将作用域属性传递给指令。还记得我们将指令放入隔离作用域之后代码就无法正常工作了吗？现在就让我们来解决这个问题。   

假设变量app已经被初始化并指向了Angular module。指令的代码如下所示：   

<pre class=" language-javascript"><code class=" language-javascript">app<span class="token punctuation">.</span><span class="token function">directive<span class="token punctuation">(</span></span><span class="token string">'helloWorld'</span><span class="token punctuation">,</span><span class="token keyword">function</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">{</span>
    <span class="token keyword">return</span> <span class="token punctuation">{</span>
        scope<span class="token punctuation">:</span> <span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">,</span>
        restrict<span class="token punctuation">:</span> <span class="token string">'AE'</span><span class="token punctuation">,</span>
        replace<span class="token punctuation">:</span> <span class="token boolean">true</span><span class="token punctuation">,</span>
        template<span class="token punctuation">:</span> <span class="token string">'&lt;p style="background-color:{{color}}"&gt;Hello World&lt;/p&gt;'</span><span class="token punctuation">,</span>
        link<span class="token punctuation">:</span> <span class="token keyword">function</span><span class="token punctuation">(</span>scope<span class="token punctuation">,</span>elem<span class="token punctuation">,</span>attrs<span class="token punctuation">)</span><span class="token punctuation">{</span>
            elem<span class="token punctuation">.</span><span class="token function">bind<span class="token punctuation">(</span></span><span class="token string">'click'</span><span class="token punctuation">,</span><span class="token keyword">function</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">{</span>
                elem<span class="token punctuation">.</span><span class="token function">css<span class="token punctuation">(</span></span><span class="token string">'background-color'</span><span class="token punctuation">,</span><span class="token string">'white'</span><span class="token punctuation">)</span><span class="token punctuation">;</span>
                scope<span class="token punctuation">.</span>$<span class="token function">apply<span class="token punctuation">(</span></span><span class="token keyword">function</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">{</span>
                    scope<span class="token punctuation">.</span>color <span class="token operator">=</span> <span class="token string">"white"</span><span class="token punctuation">;</span>
                <span class="token punctuation">}</span><span class="token punctuation">)</span><span class="token punctuation">;</span>
            <span class="token punctuation">}</span><span class="token punctuation">)</span><span class="token punctuation">;</span>
            elem<span class="token punctuation">.</span><span class="token function">bind<span class="token punctuation">(</span></span><span class="token string">'mousover'</span><span class="token punctuation">,</span><span class="token keyword">function</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">{</span>
                elem<span class="token punctuation">.</span><span class="token function">css<span class="token punctuation">(</span></span><span class="token string">'cusor'</span><span class="token punctuation">,</span><span class="token string">'pointer'</span><span class="token punctuation">)</span><span class="token punctuation">;</span>
            <span class="token punctuation">}</span><span class="token punctuation">)</span><span class="token punctuation">;</span>
        <span class="token punctuation">}</span>
    <span class="token punctuation">}</span>
<span class="token punctuation">}</span><span class="token punctuation">)</span><span class="token punctuation">;</span>   
</code></pre>

HTML标记，连同指令一起使用的代码如下所示：   

<pre class=" language-javascript"><code class=" language-javascript"><span class="token operator">&lt;</span>body ng<span class="token operator">-</span>controller<span class="token operator">=</span><span class="token string">"MainCtrl"</span><span class="token operator">&gt;</span>
    <span class="token operator">&lt;</span>input type<span class="token operator">=</span><span class="token string">"text"</span> ng<span class="token operator">-</span>model<span class="token operator">=</span><span class="token string">"color"</span> placeholder<span class="token operator">=</span><span class="token string">"Enter a color"</span><span class="token operator">/</span><span class="token operator">&gt;</span>   
    <span class="token operator">&lt;</span>hello<span class="token operator">-</span>world<span class="token operator">/</span><span class="token operator">&gt;</span>
<span class="token operator">&lt;</span><span class="token operator">/</span>body<span class="token operator">&gt;</span> 
</code></pre>

这些代码目前来说并不能正常运行。因为我们使用了一个隔离作用域，在template中的表达式{{color}}依赖于作用域（不是父作用域）。但是input元素中的ng-model指令指向的是父作用域中的属性color。因此，我们需要一种方法来绑定隔离作用域和父作用域中的属性。在Angular中，这种绑定可以在HTML中设置指令元素的属性以及配置指令定义对象中的scope属性来完成。下面我们就来探索设置这种绑定的几种方法。   

# 方法1：使用<code class=" language-javascript">@</code>来进行单项文本绑定

在指令定义中，如下面的代码所示，我们已经指明了隔离作用域中的属性color应该绑定到属性colorAttr，该属性被运用于HTML的指令中。如果你查看这个标记，你可以看到表达式{{color}}被赋值给了color-attr。当表达式的值发生变化时，属性color-attr也会发生变化。这反过来改变了隔离作用域中的属性，color。   

<pre class=" language-javascript"><code class=" language-javascript">app<span class="token punctuation">.</span><span class="token function">directive<span class="token punctuation">(</span></span><span class="token string">'helloWorld'</span><span class="token punctuation">,</span><span class="token keyword">function</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">{</span>
    <span class="token keyword">return</span> <span class="token punctuation">{</span>
        scope<span class="token punctuation">:</span> <span class="token punctuation">{</span>
            color<span class="token punctuation">:</span> <span class="token string">'@colorAttr'</span>
        <span class="token punctuation">}</span><span class="token punctuation">,</span>
        <span class="token punctuation">.</span><span class="token punctuation">.</span><span class="token punctuation">.</span>
       <span class="token comment" spellcheck="true"> //配置的余下部分
</span>    <span class="token punctuation">}</span>
<span class="token punctuation">}</span><span class="token punctuation">)</span><span class="token punctuation">;</span>   
</code></pre>

更新的标记如下所示：   

<pre class=" language-javascript"><code class=" language-javascript"><span class="token operator">&lt;</span>body ng<span class="token operator">-</span>controller<span class="token operator">=</span><span class="token string">'MainCtrl'</span><span class="token operator">&gt;</span>
    <span class="token operator">&lt;</span>input type<span class="token operator">=</span><span class="token string">'text'</span> ng<span class="token operator">-</span>model<span class="token operator">=</span><span class="token string">'color'</span> placeholder<span class="token operator">=</span><span class="token string">'Enter a color'</span><span class="token operator">/</span><span class="token operator">&gt;</span>   
    <span class="token operator">&lt;</span>hello<span class="token operator">-</span>world color<span class="token operator">-</span>attr<span class="token operator">=</span><span class="token string">'{{color}}'</span><span class="token operator">/</span><span class="token operator">&gt;</span>
<span class="token operator">&lt;</span><span class="token operator">/</span>body<span class="token operator">&gt;</span>   
</code></pre>

我们将这种方法成为单向绑定因为使用这种方法你只可以将字符串传递到属性（使用表达式，{{}}）。当父作用域属性发生变化时，你的隔离作用域模型也发生变化。你甚至可以在指令内部监视这个作用域属性并且在变化发生时触发任务。然而，反之则不成立！你不能通过操纵隔离作用域来改变父作用域。   

## 注意：

如果隔离作用域属性和元素属性名相同你可以这样来写指令定义：   

app.directive('helloWorld',function(){
    return {
        scope: {
            color: '@'
        },
        ...
        //配置的余下部分
    }
});   

指令被调用的HTML这样写：   

<pre class=" language-javascript"><code class=" language-javascript"><span class="token operator">&lt;</span>hello<span class="token operator">-</span>world color<span class="token operator">=</span><span class="token string">"{{color}}"</span><span class="token operator">/</span><span class="token operator">&gt;</span>   
</code></pre>

# 方法2：使用<code class=" language-javascript"><span class="token operator">=</span></code>进行双向绑定

我们将指令定义修改为下面的样子：   

<pre class=" language-javascript"><code class=" language-javascript">app<span class="token punctuation">.</span><span class="token function">directive<span class="token punctuation">(</span></span><span class="token string">'helloWorld'</span><span class="token punctuation">,</span><span class="token keyword">function</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">{</span>
    <span class="token keyword">return</span> <span class="token punctuation">{</span>
        scope<span class="token punctuation">:</span> <span class="token punctuation">{</span>
            color<span class="token punctuation">:</span> <span class="token string">'='</span>
        <span class="token punctuation">}</span><span class="token punctuation">,</span>
        <span class="token punctuation">.</span><span class="token punctuation">.</span><span class="token punctuation">.</span>
       <span class="token comment" spellcheck="true"> //配置的余下部分
</span>    <span class="token punctuation">}</span>
<span class="token punctuation">}</span><span class="token punctuation">)</span><span class="token punctuation">;</span>   
</code></pre>

将HTML改成下面的样子：   

<pre class=" language-javascript"><code class=" language-javascript"><span class="token operator">&lt;</span>body ng<span class="token operator">-</span>controller<span class="token operator">=</span><span class="token string">'MainCtrl'</span><span class="token operator">&gt;</span>
    <span class="token operator">&lt;</span>input type<span class="token operator">=</span><span class="token string">'text'</span> ng<span class="token operator">-</span>model<span class="token operator">=</span><span class="token string">'color'</span> placeholder<span class="token operator">=</span><span class="token string">'Enter a color'</span><span class="token operator">/</span><span class="token operator">&gt;</span>   
    <span class="token operator">&lt;</span>hello<span class="token operator">-</span>world color<span class="token operator">=</span><span class="token string">'color'</span><span class="token operator">/</span><span class="token operator">&gt;</span>
<span class="token operator">&lt;</span><span class="token operator">/</span>body<span class="token operator">&gt;</span>   
</code></pre>

和<code class=" language-javascript">@</code>不一样，这种技巧让你可以将一个实际的作用域模型赋值给一个属性而不是一个普通的字符串。结果是你可以传递将简单地字符串到复杂的数组传递到隔离作用域。同时，一个双向绑定被建立了。无论什么时候父作用域属性发生了变化，相应的隔离作用域的属性也会发生变化，反之亦然。通常情况下，你可以监视作用域属性的变化。   

# 方法3：使用<code class=" language-javascript"><span class="token operator">&amp;</span></code>来执行父作用域中的函数

有时我们需要从隔离作用域的指令中调用父作用域中的定义的函数。为了引用在外部作用域中定义的函数我们使用<code class=" language-javascript"><span class="token operator">&amp;</span></code>。假设我们现在想从指令中调用一个叫做sayHello()的函数。下面的代码解释了怎样实现这一点。   

<pre class=" language-javascript"><code class=" language-javascript">app<span class="token punctuation">.</span><span class="token function">directive<span class="token punctuation">(</span></span><span class="token string">'sayHello'</span><span class="token punctuation">,</span><span class="token keyword">function</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">{</span>
    <span class="token keyword">return</span> <span class="token punctuation">{</span>
        scope<span class="token punctuation">:</span> <span class="token punctuation">{</span>
            sayHelloIsolated<span class="token punctuation">:</span> <span class="token string">'$amp;'</span>
        <span class="token punctuation">}</span><span class="token punctuation">,</span>
        <span class="token punctuation">.</span><span class="token punctuation">.</span><span class="token punctuation">.</span>
       <span class="token comment" spellcheck="true"> //配置的余下部分
</span>    <span class="token punctuation">}</span>
<span class="token punctuation">}</span><span class="token punctuation">)</span><span class="token punctuation">;</span>   
</code></pre>

HTML中的指令这样使用：   

<pre class=" language-javascript"><code class=" language-javascript"><span class="token operator">&lt;</span>body ng‐controller<span class="token operator">=</span><span class="token string">"MainCtrl"</span><span class="token operator">&gt;</span>
<span class="token operator">&lt;</span>input type<span class="token operator">=</span><span class="token string">"text"</span> ng‐model<span class="token operator">=</span><span class="token string">"color"</span> placeholder<span class="token operator">=</span><span class="token string">"Enter a color"</span><span class="token operator">/</span><span class="token operator">&gt;</span>
<span class="token operator">&lt;</span>say‐hello sayHelloIsolated<span class="token operator">=</span><span class="token string">"sayHello()"</span><span class="token operator">/</span><span class="token operator">&gt;</span>
<span class="token operator">&lt;</span><span class="token operator">/</span>body<span class="token operator">&gt;</span>   
</code></pre>

如果你是一个Angular新手，你可能会对究竟该选择什么作用域感到困惑。默认情况下一个指令不会创建一个新的作用域而是使用父作用域。但是在许多情形中这并不是我们想要的。如果你的指令需要经常操纵父作用域属性并同时创建了新的作用域，它可能会污染父作用域。让所有的指令使用同一个父作用域并不是一个好主意，因为任何人都可以修改我们的作用域属性。因此，下列的指导原则可能会帮助你为你的指令选择正确的作用域。   

1.父作用域（scope: false）- 这是默认情况下的选择。如果你的指令不需要早总父作用域属性你可能不需要一个新的作用域。在这种情形中，使用父作用域是没问题的。   

2.子作用域（scope: true）- 这为一个指令创建了一个原型继承于父作用域的子组用于。如果你在作用域中设置的属性和函数与其他的指令和父作用域没什么联系，你应该创建一个新的子作用域。这样一来你也拥有了在父作用域中定义的所有作用域属性和函数。   

3.隔离作用域（scope:{}）- 这像是一个沙盒！如果你创建的指令是自包含并且可重用的你需要这种作用域。你的指令可能会创建许多作用域属性和函数用于内部使用，并且永远不应该被外界看到。如果你处于这种情况下，最好使用一个隔离作用域。正如你所期望的，隔离作用域并不会继承父作用域。   

# Transclusion

Transclusion是一个能够让我们用任何内容包裹一个指令的特性。我们接下来可以用正确的作用域提取并编译这个指令，最后将他放置在指令模板的特定位置。如果你在指令定义中设置transclusion: true，一个新的transcluded空间将会被创建，它同时也原型继承了父作用域。如果你想要你的指令放置在隔离控作用于中并包含一块任意内容并在父作用域中执行它，Transclusion将会被用到。   

假设我们有一个如下所示的指令：   

<pre class=" language-javascript"><code class=" language-javascript">app<span class="token punctuation">.</span><span class="token function">directive<span class="token punctuation">(</span></span><span class="token string">'outputText'</span><span class="token punctuation">,</span><span class="token keyword">function</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">{</span>
    <span class="token keyword">return</span> <span class="token punctuation">{</span>
        transclude<span class="token punctuation">:</span> <span class="token boolean">true</span><span class="token punctuation">,</span>
        scope<span class="token punctuation">:</span> <span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">,</span>
        template<span class="token punctuation">:</span> <span class="token string">'&lt;div ng-transclude&gt;&lt;/div&gt;'</span>
    <span class="token punctuation">}</span><span class="token punctuation">;</span>
<span class="token punctuation">}</span><span class="token punctuation">)</span><span class="token punctuation">;</span>
</code></pre>

它这样被使用：  

<pre class=" language-javascript"><code class=" language-javascript"><span class="token operator">&lt;</span>div output<span class="token operator">-</span>text<span class="token operator">&gt;</span>
    <span class="token operator">&lt;</span>p<span class="token operator">&gt;</span>Hello <span class="token punctuation">{</span><span class="token punctuation">{</span>name<span class="token punctuation">}</span><span class="token punctuation">}</span><span class="token operator">&lt;</span><span class="token operator">/</span>p<span class="token operator">&gt;</span>
<span class="token operator">&lt;</span><span class="token operator">/</span>div<span class="token operator">&gt;</span>   
</code></pre>

ng-transclude指明了应该在什么地方放置transcluded内容。在这个例子中，DOM内容Hello {{name}}被提取出来放到了<div></div>中。重要的一点是记住表达式{{name}}依赖于父作用域而不是隔离作用域中的属性进行插值。   

# <code class=" language-javascript">transclude<span class="token punctuation">:</span> <span class="token string">'element'</span></code>和<code class=" language-javascript">transclude<span class="token punctuation">:</span> <span class="token boolean">true</span></code>之间的差别

有时我们需要针对指令而不仅仅是内容来transcluded元素。在这种情形下，<code class=" language-javascript">transclude<span class="token punctuation">:</span> <span class="token string">'element'</span></code>就应该被使用了。和<code class=" language-javascript">transclude<span class="token punctuation">:</span> <span class="token boolean">true</span></code>不同的是，它包含了我们在指令中用ng-transclude标记的元素。作为transclusion的结果你的link函数得到了一个预先绑定到正确的指令作用域的transcluded链接函数。这个链接函数同时也传递了另一个DOM内容被transcluded的DOM元素的副本。你可以使用它来修改这个副本并将其添加到DOM中。一些指令像是ng-repeat就是用了这种技巧来重复DOM元素。   

同时需要注意的是在使用<code class=" language-javascript">transclude<span class="token punctuation">:</span> element</code>时，指令绑定的元素被转换为了一个HTML注释。因此，如果你将<code class=" language-javascript">transclude<span class="token punctuation">:</span> element</code>和<code class=" language-javascript">replace<span class="token punctuation">:</span> <span class="token boolean">false</span></code>合并起来，指令模板本质上来说将innerHTML放入了注释中 – 这意味着实际上什么都没有发生。相反，如果你选择<code class=" language-javascript">replace<span class="token punctuation">:</span> <span class="token boolean">true</span></code>指令模板将会取代HTML注释，代码将会如我们所期待的运行。将<code class=" language-javascript">replace<span class="token punctuation">:</span> <span class="token boolean">false</span></code>和<code class=" language-javascript">transclude<span class="token punctuation">:</span> <span class="token string">'element'</span></code>是一件好事，如果你想要重复DOM元素并且不想保持元素的第一个实例（它被转换为一个注释）。   

# controller函数和require

如果你想要允许其他指令和你的指令进行交流，那么你需要使用指令的controller函数。在某些情形中你也许需要通过合并两个指令来创建一个特别的UI成分。例如你可以将一个controller函数绑定到一个指令上，如下所示：   

<pre class=" language-javascript"><code class=" language-javascript">app<span class="token punctuation">.</span><span class="token function">directive<span class="token punctuation">(</span></span><span class="token string">'outerDirective'</span><span class="token punctuation">,</span><span class="token keyword">function</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">{</span>
    <span class="token keyword">return</span> <span class="token punctuation">{</span>
        scope<span class="token punctuation">:</span> <span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">,</span>
        restrict<span class="token punctuation">:</span> <span class="token string">'AE'</span><span class="token punctuation">,</span>
        controller<span class="token punctuation">:</span> <span class="token keyword">function</span><span class="token punctuation">(</span>$scope<span class="token punctuation">,</span>$compile<span class="token punctuation">,</span>$http<span class="token punctuation">)</span><span class="token punctuation">{</span>
           <span class="token comment" spellcheck="true"> //$scope是针对指令的一个合适的作用域   
</span>            this<span class="token punctuation">.</span>addChild <span class="token operator">=</span> <span class="token keyword">function</span><span class="token punctuation">(</span>nestedDirective<span class="token punctuation">)</span><span class="token punctuation">{</span><span class="token comment" spellcheck="true"> //this指代controller   
</span>            console<span class="token punctuation">.</span><span class="token function">log<span class="token punctuation">(</span></span><span class="token string">'Got the message from nested directive'</span> <span class="token operator">+</span> nestedDirective<span class="token punctuation">.</span>message<span class="token punctuation">)</span><span class="token punctuation">;</span>
            <span class="token punctuation">}</span>
        <span class="token punctuation">}</span>
    <span class="token punctuation">}</span>
<span class="token punctuation">}</span><span class="token punctuation">)</span><span class="token punctuation">;</span>    
</code></pre>

这段代码绑定了一个叫做outerDirective的controller到指令上。当另一个指令想要发生信息交流时，它需要声明它需要这个指令的controller实例。代码如下所示：   

<pre class=" language-javascript"><code class=" language-javascript">app<span class="token punctuation">.</span><span class="token function">directive<span class="token punctuation">(</span></span><span class="token string">'innerDirective'</span><span class="token punctuation">,</span><span class="token keyword">function</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">{</span>
    <span class="token keyword">return</span> <span class="token punctuation">{</span>
        scope<span class="token punctuation">:</span> <span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">,</span>
        restrict<span class="token punctuation">:</span> <span class="token string">'AE'</span><span class="token punctuation">,</span>
        require<span class="token punctuation">:</span> <span class="token string">'^outerDirective'</span><span class="token punctuation">,</span>
        link<span class="token punctuation">:</span> <span class="token keyword">function</span><span class="token punctuation">(</span>scope<span class="token punctuation">,</span>elem<span class="token punctuation">,</span>attrs<span class="token punctuation">,</span>controllerInstance<span class="token punctuation">)</span><span class="token punctuation">{</span>
           <span class="token comment" spellcheck="true"> //第四个参数是你require的controller实例
</span>            scope<span class="token punctuation">.</span>message <span class="token operator">=</span> <span class="token string">'Hi, Parent directive'</span><span class="token punctuation">;</span>
            controllerInstance<span class="token punctuation">.</span><span class="token function">addChild<span class="token punctuation">(</span></span>scope<span class="token punctuation">)</span><span class="token punctuation">;</span>
        <span class="token punctuation">}</span>
    <span class="token punctuation">}</span><span class="token punctuation">;</span>
<span class="token punctuation">}</span><span class="token punctuation">)</span><span class="token punctuation">;</span>   
</code></pre>

HTML标记应该如下所示：   

<pre class=" language-javascript"><code class=" language-javascript"><span class="token operator">&lt;</span>outer‐directive<span class="token operator">&gt;</span> 
    <span class="token operator">&lt;</span>inner‐directive<span class="token operator">&gt;</span><span class="token operator">&lt;</span><span class="token operator">/</span>inner‐directive<span class="token operator">&gt;</span>
<span class="token operator">&lt;</span><span class="token operator">/</span>outer‐directive<span class="token operator">&gt;</span> 
</code></pre>

<code class=" language-javascript">require<span class="token punctuation">:</span> <span class="token string">'^outerDirective'</span></code>告诉Angular在元素和它的父元素上搜索这个controller。在这个例子中北发现的controller实例被传递给了link函数的第四个参数。在我们的例子中我们将嵌套指令的作用域传递给了父作用域。   

# 一个笔记应用

在这一部分中我们将使用指令创建一个简单地笔记应用。我们将使用HTML5中的localStorage来存储数据。我们将会创建一个指令来渲染出一个记事本。一个用户可以查看他/她创建的笔记列表。当他点击add new按钮时，记事本变为了可编辑状态并允许一个创建一个新笔记。当点击back按钮时笔记会被自动存储。笔记使用一个叫做noteFactory的factory存储，连同来自localStorage的帮助。factory代码非常直观。因此，我们在此只专注于指令代码。   

## 第一步

我们首先来注册一个notepad指令。   

<pre class=" language-javascript"><code class=" language-javascript">app<span class="token punctuation">.</span><span class="token function">directive<span class="token punctuation">(</span></span><span class="token string">'notepad'</span><span class="token punctuation">,</span><span class="token keyword">function</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">{</span>
    <span class="token keyword">return</span> <span class="token punctuation">{</span>
        restrict<span class="token punctuation">:</span> <span class="token string">'AE'</span><span class="token punctuation">,</span>
        scope<span class="token punctuation">:</span> <span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">,</span>
        link<span class="token punctuation">:</span> <span class="token keyword">function</span><span class="token punctuation">(</span>scope<span class="token punctuation">,</span>elem<span class="token punctuation">,</span>attrs<span class="token punctuation">)</span><span class="token punctuation">{</span>
        <span class="token punctuation">}</span><span class="token punctuation">,</span>
        templateUrl<span class="token punctuation">:</span> <span class="token string">'templateurl,html'</span>
    <span class="token punctuation">}</span><span class="token punctuation">;</span>
<span class="token punctuation">}</span><span class="token punctuation">)</span><span class="token punctuation">;</span>
</code></pre>

注意关于该指令的几点：   

<ul>
<li>作用域是隔离的，因为我们想要指令被重用。该指令将会拥有许多外界无法访问的属性和函数。   </li>
<li>有restrict性质的设置，该指令可以被用作一个属性或者是元素。   </li>
<li>link函数初始化为空函数   </li>
<li>指令从templateurl.html获取它的模板。   </li>
</ul>

## 第二步

下面的HTML是该指令的模板   

<pre class=" language-javascript"><code class=" language-javascript"><span class="token operator">&lt;</span>div class<span class="token operator">=</span><span class="token string">"note‐area"</span> ng‐show<span class="token operator">=</span><span class="token string">"!editMode"</span><span class="token operator">&gt;</span> 
   <span class="token operator">&lt;</span>ul<span class="token operator">&gt;</span>
      <span class="token operator">&lt;</span>li ng‐repeat<span class="token operator">=</span><span class="token string">"note in notes|orderBy:'id'"</span><span class="token operator">&gt;</span>
      <span class="token operator">&lt;</span>a href<span class="token operator">=</span><span class="token string">"#"</span> ng‐click<span class="token operator">=</span><span class="token string">"openEditor(note.id)"</span><span class="token operator">&gt;</span><span class="token punctuation">{</span><span class="token punctuation">{</span>note<span class="token punctuation">.</span>title<span class="token punctuation">}</span><span class="token punctuation">}</span><span class="token operator">&lt;</span><span class="token operator">/</span>a<span class="token operator">&gt;</span>
     <span class="token operator">&lt;</span><span class="token operator">/</span>li<span class="token operator">&gt;</span>
<span class="token operator">&lt;</span><span class="token operator">/</span>ul<span class="token operator">&gt;</span>
<span class="token operator">&lt;</span><span class="token operator">/</span>div<span class="token operator">&gt;</span>
<span class="token operator">&lt;</span>div id<span class="token operator">=</span><span class="token string">"editor"</span> ng‐show<span class="token operator">=</span><span class="token string">"editMode"</span> class<span class="token operator">=</span><span class="token string">"note‐area"</span> contenteditable<span class="token operator">=</span><span class="token string">"true"</span> ng‐bind<span class="token operator">=</span><span class="token string">"noteText"</span><span class="token operator">&gt;</span><span class="token operator">&lt;</span><span class="token operator">/</span> <span class="token operator">&lt;</span>span<span class="token operator">&gt;</span><span class="token operator">&lt;</span>
 a href<span class="token operator">=</span><span class="token string">"#"</span> ng‐click<span class="token operator">=</span><span class="token string">"save()"</span> ng‐show<span class="token operator">=</span><span class="token string">"editMode"</span><span class="token operator">&gt;</span>Back<span class="token operator">&lt;</span><span class="token operator">/</span>a<span class="token operator">&gt;</span><span class="token operator">&lt;</span><span class="token operator">/</span>span<span class="token operator">&gt;</span>
<span class="token operator">&lt;</span>span<span class="token operator">&gt;</span><span class="token operator">&lt;</span>a href<span class="token operator">=</span><span class="token string">"#"</span> ng‐click<span class="token operator">=</span><span class="token string">"openEditor()"</span> ng‐show<span class="token operator">=</span><span class="token string">"!editMode"</span><span class="token operator">&gt;</span>Add Note<span class="token operator">&lt;</span><span class="token operator">/</span>a<span class="token operator">&gt;</span><span class="token operator">&lt;</span><span class="token operator">/</span>span<span class="token operator">&gt;</span>  
</code></pre>

需要注意的重点有以下几点：   

<ul>
<li>note对象封装了title,id和content。   </li>
<li>ng-repeat被用来循环遍历notes并且更具自动生成的id升序排列。   </li>
<li>我们用一个叫做editMode的属性来表明我们所处的模式。在编辑模式之下该属性为true并且可编辑的div是可见的。用户可以在此编写笔记。   </li>
<li>如果editMode是false，我们将处于浏览模式并且展示notes。   </li>
<li>两个按钮也可以更具editMode来显示/消失   </li>
<li>hg-click指令被用来响应按钮点击时间。这些方法，和editMode属性一样，都会被添加到作用域中。   </li>
<li>可编辑的div被绑定到了noteText上，它将保存用户输入的文本。如果你想要编辑一个已存在的笔记，这个模型将会将div初始化为笔记内容。   </li>
</ul>

## 第三步

我们现在在我们的作用域中创建一个叫做restore()的函数来初始化我们应用的一些控制。当link函数运行并且save按钮被点击时该函数将会被调用。   

<pre class=" language-javascript"><code class=" language-javascript">scope<span class="token punctuation">.</span>restore <span class="token operator">=</span> <span class="token keyword">function</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">{</span>
    scope<span class="token punctuation">.</span>editMode <span class="token operator">=</span> <span class="token boolean">false</span><span class="token punctuation">;</span>
    scope<span class="token punctuation">.</span>index <span class="token operator">=</span> <span class="token operator">-</span><span class="token number">1</span><span class="token punctuation">;</span>
    scope<span class="token punctuation">.</span>noteText <span class="token operator">=</span> <span class="token string">''</span><span class="token punctuation">;</span>
<span class="token punctuation">}</span>   
</code></pre>

我们在link函数内部创建这个函数。editMode和noteText已经解释过了。index用来追踪哪一个笔记正在被编辑。如果我们想要创建一个新笔记，index为-1.如果我们正在编辑一个已经存在的笔记，index为note对象的id。   

## 第四步

现在我们需要创建两个作用域函数来处理编辑和存储动作：   

<pre class=" language-javascript"><code class=" language-javascript">scope<span class="token punctuation">.</span>openEditor <span class="token operator">=</span> <span class="token keyword">function</span><span class="token punctuation">(</span>index<span class="token punctuation">)</span><span class="token punctuation">{</span>
    scope<span class="token punctuation">.</span>editMode <span class="token operator">=</span> <span class="token boolean">true</span><span class="token punctuation">;</span>

    <span class="token keyword">if</span><span class="token punctuation">(</span>index <span class="token operator">!</span><span class="token operator">==</span> undefined<span class="token punctuation">)</span><span class="token punctuation">{</span>
        scope<span class="token punctuation">.</span>noteText <span class="token operator">=</span> notesFactory<span class="token punctuation">.</span><span class="token function">get<span class="token punctuation">(</span></span>index<span class="token punctuation">)</span><span class="token punctuation">.</span>content<span class="token punctuation">;</span>
        scope<span class="token punctuation">.</span>index <span class="token operator">=</span> index<span class="token punctuation">;</span>
    <span class="token punctuation">}</span> <span class="token keyword">else</span> <span class="token punctuation">{</span>
        scope<span class="token punctuation">.</span>noteText <span class="token operator">=</span> undefined<span class="token punctuation">;</span>
    <span class="token punctuation">}</span>
<span class="token punctuation">}</span><span class="token punctuation">;</span>

scope<span class="token punctuation">.</span>save <span class="token operator">=</span> <span class="token keyword">function</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">{</span>
    <span class="token keyword">if</span><span class="token punctuation">(</span>scope<span class="token punctuation">.</span>noteText <span class="token operator">!</span><span class="token operator">==</span> <span class="token string">''</span><span class="token punctuation">)</span><span class="token punctuation">{</span>
        <span class="token keyword">var</span> note <span class="token operator">=</span> <span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">;</span>

        note<span class="token punctuation">.</span>title <span class="token operator">=</span> scope<span class="token punctuation">.</span>noteText<span class="token punctuation">.</span>length <span class="token operator">&gt;</span> <span class="token number">10</span> <span class="token operator">?</span> scope<span class="token punctuation">.</span>noteText<span class="token punctuation">.</span><span class="token function">substring<span class="token punctuation">(</span></span><span class="token number">0</span><span class="token punctuation">,</span><span class="token number">10</span><span class="token punctuation">)</span> <span class="token operator">+</span> <span class="token string">'...'</span> <span class="token punctuation">:</span> scope<span class="token punctuation">.</span>noteText<span class="token punctuation">;</span>
        note<span class="token punctuation">.</span>content <span class="token operator">=</span> scope<span class="token punctuation">.</span>noteText<span class="token punctuation">;</span>
        note<span class="token punctuation">.</span>id <span class="token operator">=</span> scope<span class="token punctuation">.</span>index <span class="token operator">!</span><span class="token operator">=</span> <span class="token operator">-</span><span class="token number">1</span> <span class="token operator">?</span> scope<span class="token punctuation">.</span>index <span class="token punctuation">:</span> localStorage<span class="token punctuation">.</span>length<span class="token punctuation">;</span>
        scope<span class="token punctuation">.</span>notes <span class="token operator">=</span> notesFactory<span class="token punctuation">.</span><span class="token function">put<span class="token punctuation">(</span></span>note<span class="token punctuation">)</span><span class="token punctuation">;</span>
    <span class="token punctuation">}</span>

    scope<span class="token punctuation">.</span><span class="token function">restore<span class="token punctuation">(</span></span><span class="token punctuation">)</span><span class="token punctuation">;</span>
<span class="token punctuation">}</span>
</code></pre>

需要注意的重点有以下几个：   

<ul>
<li>openEditor为编辑器做准备。如果你正在编辑一个笔记，它将会获得笔记内容并通过ng-bind来更新可编辑的div。  </li>
<li>如果我们正在创建一个新的笔记我们需要将noteText设置为undefined，以便在我们保存笔记的时候监视器能被触发。   </li>
<li>如果函数的参数index是undefined，它意味着用户将要创建一个新笔记。   </li>
<li>save函数在notesFactory的帮助下来保存笔记。在保存之后，它刷新了notes数组以便监视器能检测到一个变化，notes列表能够被更新。   </li>
<li>save函数在最后调用restore()函数来重置控制以便我们可以从编辑模式回退到浏览模式。   </li>
</ul>

## 第五步

当link函数运行时我们初始化notes数组并且将一个keydown时间绑定到了可编辑的div上以便我们的noteText模型可以与div内容保持同步。我们使用noteText来保存笔记内容。   

<pre class=" language-javascript"><code class=" language-javascript"><span class="token keyword">var</span> editor <span class="token operator">=</span> elem<span class="token punctuation">.</span><span class="token function">find<span class="token punctuation">(</span></span><span class="token string">'#editor'</span><span class="token punctuation">)</span><span class="token punctuation">;</span>  

scope<span class="token punctuation">.</span><span class="token function">restore<span class="token punctuation">(</span></span><span class="token punctuation">)</span><span class="token punctuation">;</span><span class="token comment" spellcheck="true"> //初始化我们的应用控制   
</span>scope<span class="token punctuation">.</span>notes <span class="token operator">=</span> notesFactory<span class="token punctuation">.</span><span class="token function">getAll<span class="token punctuation">(</span></span><span class="token punctuation">)</span><span class="token comment" spellcheck="true">;//载入笔记  
</span>
editor<span class="token punctuation">.</span><span class="token function">bind<span class="token punctuation">(</span></span><span class="token string">'keyup keydown'</span><span class="token punctuation">,</span><span class="token keyword">function</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">{</span>
    scope<span class="token punctuation">.</span>noetText <span class="token operator">=</span> editor<span class="token punctuation">.</span><span class="token function">text<span class="token punctuation">(</span></span><span class="token punctuation">)</span><span class="token punctuation">.</span><span class="token function">trim<span class="token punctuation">(</span></span><span class="token punctuation">)</span><span class="token punctuation">;</span>
<span class="token punctuation">}</span><span class="token punctuation">)</span><span class="token punctuation">;</span>  
</code></pre>

## 第六步

最后，我们就可以像使用其他HTML元素一样使用这条指令开始记笔记。   

<pre class=" language-javascript"><code class=" language-javascript"><span class="token operator">&lt;</span>h1 class<span class="token operator">=</span><span class="token string">'title'</span><span class="token operator">&gt;</span>The Note Making App<span class="token operator">&lt;</span><span class="token operator">/</span>h1<span class="token operator">&gt;</span> 
<span class="token operator">&lt;</span>notepad<span class="token operator">/</span><span class="token operator">&gt;</span> 
</code></pre>

# 总结

我们注意到了任何可以用jQuery完成的事情，我们都可以用Angular指令完成并且代码量更少。因此，在使用jQuery之前，想一想同样的事情是不是不可用一种不涉及DOM操纵的更好的方式完成。你完全可以用Angular来减少使用哪个jQuery的次数。   

<hr>

本文译自A Practical Guide to AngularJS Directives (Part Two)，原文地址<a href="http://www.sitepoint.com/practical-guide-angularjs-directives-part-two/">http://www.sitepoint.com/practical-guide-angularjs-directives-part-two/</a>   