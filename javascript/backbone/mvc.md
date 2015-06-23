# 为什么Backbone是纯净MVC？

http://www.html-js.com/article/Bidirectional-binding-cornerstone-fool-Ann-BackboneEvents--pure-MVC-framework

在这个大前端时代，各路MV*框架如雨后春笋搬涌现出来,在infoQ上有一篇 12种JavaScript MVC框架之比较,胜出的是Ember.js,
当然这只是 Gordon L. Hempton的一家之言（Ember.js确实有其强大之处），关于孰强孰弱，大家肯定有自己心中的No.1。
关于到底有多少种前端MVC 框架，愚安我肯定是不知道的，除了上面提到的12种以外，还有很多国内国外的MV*框架，大家造轮子的热情也无比高涨，
各种demo活跃在 各大技术社区。一时间有句调侃的话——前端MV*哪家强，不服写个TodoList，
这里有一个目前主流的MV*框架写的Todolist的example，叫做 Helping you select an MV* framework大家可以稍作了解。

我们知道，MV*框架的优势在于，在结构上其可以组织良好的结构化、模块化代码；在逻辑上，实现以下功能：

* 构建DOM
* 实现视图逻辑
* 在模型与视图间进行同步
* 管理复杂的UI交互操作
* 管理状态和路由
* 创建与连接组件

在诸多的此类框架中，笔者真正在生产环节使用过的聊聊无几，如，Angular，Backbone，Ember，React，其余的我就不敢多言了。
其中,React.js使用的是一种叫做virtual dom的概念，让我眼前一亮。Angular.js采用一种预编译技术，
将dom中的元素与Controler的scope 结合起来，然后采取脏轮训的方式监听二者的变化，实现模型数据与dom间的双向绑定,实时更新。
而Ember.js作为Ruby on Rails框架开发团队的 又一力作，其野心可以从其类库的强大看出，单纯的Ember.js文件就有足足141kb,
而且Ember在视图层提供了数据绑定的功能，可以轻松实现 页面数据与模型的数据绑定。

而今天愚安要说的Backbone.js相比以上三者，就显的弱小多了，其文件大小只有18kb（无依赖未压缩）。为了单纯的实现一个MVC结构，
Backbone 并没有像其他框架那样，花大力气增强自己的工具类库。其在操作dom和ajax上完全依赖jQuery，在工具类上完全依赖underscore。

正是因为如此，Backbone的结构十分简洁清晰，易于扩展，所以Backbone得开源社区十分活跃，插件数量在所有MV*框架中鹤立鸡群。
所以，加上注释也只有1700 余行的Backbone是一个纯净的MVC框架。

Backbone的结构

打开Backbone的官网,我们发现构成Backbone的模块只有Events,Model,Collection,Router,History,Sync,
View,noConflict几部分组成。

折叠后的Backbone关键代码如下：
```js
Backbone.VERSION = '1.1.2';//版本
Backbone.$ = $;
//出让对Backbone命名空间的所有权
Backbone.noConflict = function() {
};
//Events事件
var Events = Backbone.Events = {
};
_.extend(Backbone, Events);
//Model模型
var Model = Backbone.Model = function(attributes, options) {
};
_.extend(Model.prototype, Events, {
});
//Collection集合
var Collection = Backbone.Collection = function(models, options) {
};
_.extend(Collection.prototype, Events, {
});
//View视图
var View = Backbone.View = function(options) {
};
_.extend(View.prototype, Events, {
});
//sync同步方法
Backbone.sync = function(method, model, options) {
};
//贴出只是为了佐证Backbone的ajax是使用jQuery的ajax,而不是像Angular.js那样实现自己的$http
Backbone.ajax = function() {
    return Backbone.$.ajax.apply(Backbone.$, arguments);
};
//Router路由
var Router = Backbone.Router = function(options) {
};
_.extend(Router.prototype, Events, {
});
//History浏览历史(window.history)
var History = Backbone.History = function() {
};
_.extend(History.prototype, Events, {
});
Backbone.history = new History;
//在underscore基础上实现的关键性的继承方法,这个也很关键
var extend = function(protoProps, staticProps) {
});
Model.extend = Collection.extend = Router.extend = View.extend = History.extend = extend;
```

不得不承认Backbone的代码真的非常简洁清晰。

Events如何实现从Model到View的绑定？

通过上面的简单折叠代码，我们可以看出,不管是Model,Collection,Router,History,History,甚至是Backbone本身， 都或通过原型链或直接继承了Backboe.Events。这也是在Backbone的代码编写顺序上Backboe.Events会放在最前面的原因。

那么Backboe.Events到底做了些什么呢？还是贴代码最有说服力：

```js
var Events = Backbone.Events = {
    //绑定一个事件到`callback`回调函数上。通过 `"all"`可以绑定这个回调函数到所有事件上
    on: function(name, callback, context) {
    },
    //绑定一个仅会被触发一次的事件。在这个事件的回调函数被调用一次之后，这个回调函数将被移除
    once: function(name, callback, context) {
    },
    //移除一个或多个的事件回调
    off: function(name, callback, context) {
    },
    //触发一个或多个事件，调用对应的回调函数。
    trigger: function(name) {
    },
    //`on`和`once`的控制反转版本。告诉当前对象去监听另一个对象的事件
    listenTo: function(obj, name, callback) {
    },
    listenToOnce: function(obj, name, callback) {
    },
    //告诉当前对象停止对指定对象的指定事件的监听，或停止所有监听
    stopListening: function(obj, name, callback) {
    }
  };
```
基于这样的一个Events对象的实现，Backbone可以轻松实现了很多功能，如在Model.set(key,value)
时触发一个change事件,视图层在扑捉到 这个事件的时候，对dom做出相应的更新，这样就实现了Model层到View层的绑定。例如：

```js
var View = Backbone.View.extend({
    initialize:function(){
        this.listenTo(this.model,'change:name',this.onNameChange);
    },
    onNameChange:function(){
        this.$('.name').text(this.model.get('name'));
    }
    template: '<span class="name"></span>'
});
var m = new Backbone.Model({name:'Jack'});
var v = new View({model:m});
m.set('name','John');
```

那么，这里的set为什么会触发change事件呢？具体实现我还是贴一下源码和自己的中文注释：
```js

Backbone.Model.prototype.set = function (key, val, options) {
      var attr, attrs, unset, changes, silent, changing, prev, current;
      if (key == null) return this;
      //格式化参数
      if (typeof key === 'object') {
        attrs = key;
        options = val;
      } else {
        (attrs = {})[key] = val;
      }
      options || (options = {});
      //执行当前对象的验证方法
      if (!this._validate(attrs, options)) return false;
      //提取属性和可选项
      unset           = options.unset;
      silent          = options.silent;
      changes         = [];
      changing        = this._changing;
      this._changing  = true;
      //标记当前Model是否改变，并记录改变的属性及其变化前后的值
      if (!changing) {
        this._previousAttributes = _.clone(this.attributes);
        this.changed = {};
      }
      current = this.attributes, prev = this._previousAttributes;
      //若改变的属性为id，则同时改变当前对象的id
      if (this.idAttribute in attrs) this.id = attrs[this.idAttribute];
      //遍历`set`的属性，更新或删除对应属性的当前值
      for (attr in attrs) {
        val = attrs[attr];
        if (!_.isEqual(current[attr], val)) changes.push(attr);
        if (!_.isEqual(prev[attr], val)) {
          this.changed[attr] = val;
        } else {
          delete this.changed[attr];
        }
        unset ? delete current[attr] : current[attr] = val;
      }
      //若非沉默更新（传参时options.silent=true），触发change：attr事件
      //attr为各个对应被set的属性的key，并传当前值到回调函数
      if (!silent) {
        if (changes.length) this._pending = options;
        for (var i = 0, l = changes.length; i < l; i++) {
          this.trigger('change:' + changes[i], this, current[changes[i]], options);
        }
      }
      //change可以递归嵌套到change事件中
      if (changing) return this;
      if (!silent) {
        while (this._pending) {
          options = this._pending;
          this._pending = false;
          this.trigger('change', this, options);
        }
      }
      this._pending = false;
      this._changing = false;
      return this;
    }
```
所以若想，监听到Model的属性变化，改变Model的属性值时，必须采用Model.set()方法，
而不能简单的使用Model.attributes[key] = value。 其实这个是一个兼容的做法，
我们知道Backbone对低版本浏览器的支持非常好，如果不考虑这些的话，完全可以使用更高级的API

```js
Object.observe(this.attributes, function(changes){
    if (!silent) {
        if (changes.length) this._pending = options;
        for (var i = 0, l = changes.length; i < l; i++) {
          this.trigger('change:' + changes[i].name, this, current[changes[i].name], options);
    }
}.bind(this));
```
当然，这只是愚安我的一点意淫，没有什么实际意义。

实际上，Backbone内部事件除了change以外还有很多，这里简单列举一下：

* “add” (model, collection, options) — 当一个model被add到一个collection时
* “remove” (model, collection, options) — 当一个model从一个collection移除时
* “reset” (collection, options) — 当一个collection的实体内容已经被替换掉时
* “sort” (collection, options) — 当一个collection的内容被重新排序时
* “change” (model, options) — 当一个model的属性被改变时
* “change:[attribute]” (model, value, options) — 当一个model的指定属性被改变时
* “destroy” (model, collection, options) — 当一个model被销毁时
* “request” (model_or_collection, xhr, options) — 当一个model或collection开始发起一个向服务端的请求时
* “sync” (model_or_collection, resp, options) — 当一个model或collection已经成功与服务端同步时
* “error” (model_or_collection, resp, options) — 当model或collection对服务端的请求已经失败时
* “invalid” (model, error, options) — 当一个model的验证失败时
* “route:[name]” (params) — 当路由的一个指定path被匹配时由路由器触发
* “route” (route, params) — 当任意路由被匹配时由路由器触发
* “route” (router, route, params) — 当任意路由被匹配时由history触发
* “all” — 任意事件,事件名称作为第一个参数传递

Events如何实现从View到Model的绑定？

上面我们已经知道基于强大的Backbone.Events,我们可以轻松的实现model到view的绑定，反之呢？

Backbone没有类似Angular.js的预编译机制，也没有View-Model的概念，从View到Model的绑定依赖于原生DOM事件的监听，完整双向绑定如：

```js
var View = Backbone.View.extend({
    initialize:function(){
        this.listenTo(this.model,'change:name',this.onNameChange);
    },
    onNameChange:function(){
        this.$('.name').text(this.model.get('name'));
    }
    template: '<input type="text" class="name-input"><span class="name"></span>',
    events: {'input .name-input': '_changeName'},
    _changeName: function(e){
        var value = e.currentTarget.value;
        this.model.set('name', value);
        return false;
    }
});
var m = new Backbone.Model({name:'Jack'});
var v = new View({model:m});
m.set('name','John');
```

需要注意的是View层的events字典，其实就是DOM事件，而不是Backbone.Events。
而且，纯净的Backbone这里用的是jQuery的jQuery的on 方法进行绑定的。

好的，愚安又贴了很多源码，和一些自己对文档的不成翻译，有没有干货，见仁见智了。
另外，本人是非常推荐刚接触前端框架的童鞋，以Backbone做为开始的。就像学习PHP的MVC框架，
我非常推荐以codeigniter作为开始的，没有强大的封装，但有着最基本纯净的MVC思想。
