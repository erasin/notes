#Martini源码剖析

<http://fuxiaohei.me/article/25/martini-source-study.html>

[martini][1]是非常优雅的Go Web框架。他基于依赖注入的思想，仿照[Sinatra][2]的路由设计，参考[Express][3]的中间件设计，而且核心微小，扩展方便，非常值得学习。但是由于本身API设计简洁，使很多细节无法从代码理解。所以，我写一点笔记记录`martini`的工作方式。

## Martini核心

我们从最简单的官方实例入手：

    package main

    import "github.com/go-martini/martini"

    func main() {
      m := martini.Classic()
      m.Get("/", func() string {
        return "Hello world!"
      })
      m.Run()
    }

`martini.Martini`是自带的核心结构，负责完成依赖注入和调用的过程。`martini.ClassicMartini`是路由`martini.Router`和`martini.Martini`的组合，实现路由分发和逻辑调用的过程。`m := martini.Classic()`返回的就是`martini.ClassicMartini`。具体在[martini.go#L104][4]:

    func Classic() *ClassicMartini {
        r := NewRouter()
        m := New()
        m.Use(Logger())
        m.Use(Recovery())
        m.Use(Static("public"))
        m.MapTo(r, (*Routes)(nil))
        m.Action(r.Handle)
        return &ClassicMartini{m, r}
    }

里面的`m := New()`定义在[martini.go#L38][5]:

    func New() *Martini {
        m := &Martini{Injector: inject.New(), action: func() {}, logger: log.New(os.Stdout, "[martini] ", 0)}
        m.Map(m.logger)
        m.Map(defaultReturnHandler())
        return m
    }

#### 依赖注入

上面很明显的看到两个奇特方法：`m.Map()`、`m.MapTo()`。这里，需要注意`martini`的一个最重要原则，**注入的任何类型的结构，都是唯一的**。即如：

    type User struct{
        Id int
    }

    m.Map(&User{Id:1})
    m.Map(&User{Id:2})

`martini`在寻找`&User`类型的时候，只能获取到`&User{Id:2}`结构（最后注册的）。`Map`的作用，就是向内部注册对应类型的具体对象或者值。类型索引是`reflect.Type`。从而我们可以理解到`m.New()`的代码中将`m.Logger(*log.Logger)`和`defaultReturnHandler(martini.ReturnHandler)`(括号中是类型索引)注入到内部。

这里出现了一个问题。接口这种类型，是无法用`reflect.Type()`直接获取的（因为传来的都是已经实现接口的具体结构）。解决方法就是`m.MapTo()`。
    
    m.MapTo(r, (*Routes)(nil))

即将`r(martini.router)`按照`martini.Router`接口（注意大小写）类型注入到内部。

    (*Routes)(nil)

也是高明的构造。接口的默认值不是nil，无法直接new。但是指针的默认值是nil，可以直接赋值，比如`var user *User; user = nil`。因此他注册一个接口指针类型的空指针，用`reflect.Type.Elem()`方法就可以获取到指针的内部类型，即接口类型，并以接口类型索引注入到内部。

## 路由过程

#### HTTP处理

`martini.Martini`实现了`http.Handler`方法，实际的HTTP执行过程在代码[martini.go#L68][6]:

    func (m *Martini) ServeHTTP(res http.ResponseWriter, req *http.Request) {
        m.createContext(res, req).run()
    }

这里需要我们关注`m.createContext`，它返回`*martini.context`类型，代码[martini.go#L87][7]：

    func (m *Martini) createContext(res http.ResponseWriter, req *http.Request) *context {
        c := &context{inject.New(), m.handlers, m.action, NewResponseWriter(res), 0}
        c.SetParent(m)
        c.MapTo(c, (*Context)(nil))
        c.MapTo(c.rw, (*http.ResponseWriter)(nil))
        c.Map(req)
        return c
    }

创建`<em>martini.context</em>`*类型；然后`SetParent`设置寻找注入对象的时候同时从`m(`*`martini.Martini)`中寻找（`<em>martini.context</em>`*和``*`martini.Martini`两个独立的`inject`），这样就可以获取`m.Map`注入的数据。

这里叉出来说：从代码看出实际上注入的数据有两层，分别在`<em>martini.context</em>`*和``*`martini.Martini`。`*martini.context`中的是当前请求可以获取的（每个请求都会`m.createContext()`，都是新的对象）;`martini.Martini`是全局的，任何请求都可以获取到。

回到上一段，`c.MapTo`把`<em>martini.context</em>`*按`martini.Context`接口，将`martini.ResponseWriter`按`http.ResponseWriter`接口，把`req(`*`http.Request)`注入到当前上下文。

`context.run`方法定义在[martini.go#L163][8]:

    func (c *context) run() {
        for c.index <= len(c.handlers) {
            _, err := c.Invoke(c.handler())
            if err != nil {
                panic(err)
            }
            c.index += 1

            if c.Written() {
                return
            }
        }
    }

它在循环`c.handlers`(来自`m.handlers`,createContext代码中)。这里想解释三个细节。

`c.Invoke`是`inject.Invoke`方法，内部就是获取`c.hanlder()`返回的`martini.Handler(func)`类型的传入参数`reflect.Type.In()`，根据参数个数和类型去内部找对应的结构，然后拼装成`[]reflect.Value`给函数的`reflect.Value(func).Call()`。

`c.handler()`的返回来自两个方面,`c.hanlders`和`c.action`。`c.handlers`来自`m.Use()`添加,`c.action`来自`r.Handle(*martini.router.Handle)`(见上文`martini.ClassicMartini.New`中的`m.Action(r.Handle)`)。因此，可以发现实际上handlers是有两个列表，一个是`c.handlers([]martini.handler)`和`r.handlers(martini.routerContext.handlers)`。而且前者先执行。也就是说无论`m.Use`写在哪儿，都要比router添加的func先执行。

`c.Written`判断请求是否已经发送。他实际上是判断`martini.ResponseWriter.status`是否大于0。因此只要发送了response status，handlers过程就会停止。

#### 路由调用

从上面可以知道，路由调用过程有两个方面：一是`m.Use()`添加的handlers，二是路由添加比如`m.Get("/",handlers...)`中的handlers。`m.Use`的handlers调用就是上文的`*martini.context.run`方法，不再赘述。路由中的handlers执行是在[router.go#L218][9]:

    func (r *route) Handle(c Context, res http.ResponseWriter) {
        context := &routeContext{c, 0, r.handlers}
        c.MapTo(context, (*Context)(nil))
        context.run()
    }


和[router.go#L315][10]:

    func (r *routeContext) run() {
        for r.index < len(r.handlers) {
            handler := r.handlers[r.index]
            vals, err := r.Invoke(handler)
            if err != nil {
                panic(err)
            }
            r.index += 1

            // if the handler returned something, write it to the http response
            if len(vals) > 0 {
                ev := r.Get(reflect.TypeOf(ReturnHandler(nil)))
                handleReturn := ev.Interface().(ReturnHandler)
                handleReturn(r, vals)
            }

            if r.Written() {
                return
            }
        }
    }

如果你已经理解上文中说明，这个过程和`martini.context.run`是一样的。唯一这里要解释的是`martini.ReturnHandler`。它与很上文中的`m.Map(defaultReturnHandler())`遥相呼应。

## 中间件

从上文不难理解，中间件其实就是`martini.Handler`被`m.Use`添加到`m.handlers`中。这里我们来说明官方的一个中间件`martini.Logger()`，实现代码在[logger.go][11]:


    func Logger() Handler {
        return func(res http.ResponseWriter, req *http.Request, c Context, log *log.Logger) {
            start := time.Now()
            log.Printf("Started %s %s", req.Method, req.URL.Path)

            rw := res.(ResponseWriter)
            c.Next()

            log.Printf("Completed %v %s in %v\n", rw.Status(), http.StatusText(rw.Status()), time.Since(start))
        }
    }

首先看func的传入参数，`http.ResponseWriter`和`*http.Request`来自：

    c := &context{inject.New(), m.handlers, m.action, NewResponseWriter(res), 0}
    // ...
    c.MapTo(c.rw, (*http.ResponseWriter)(nil))
    c.Map(req)

`Context`来自：

    context := &routeContext{c, 0, r.handlers}
    c.MapTo(context, (*Context)(nil))

`*log.Logger`来自：

    m := &Martini{Injector: inject.New(), action: func() {}, logger: log.New(os.Stdout, "[martini] ", 0)}
    m.Map(m.logger)

然后看`rw := res.(ResponseWriter)`。实际上`c.rw`是`NewReponseWriter(res)`返回的`martini.ResponseWriter`类型，一次可以在这里直接转换（注意在外部调用，不是martini包中，要import并写`res.(martini.ResponseWriter)`）。

最后是`c.Next()`方法，源码在[martini.go#L154][12]:

    func (c *context) Next() {
        c.index += 1
        c.run()
    }
    
意思就是index自增，指向下一个handler，`c.run`走完所有handler，然后继续中间件里的`log.Printf...`。

## 总结

martini的对外API很简单，但是内部实现其实比较复杂的。需要仔细的阅读，并且有一定标准库的基础，才能很好的理解他代码的用意。

我这里只是按照自己的理解说明，如果有错误请在评论中指正。

 [1]: http://github.com/go-martini/martini
 [2]: http://www.sinatrarb.com/
 [3]: http://expressjs.com/
 [4]: https://github.com/go-martini/martini/blob/master/martini.go#L104
 [5]: https://github.com/go-martini/martini/blob/master/martini.go#L38
 [6]: https://github.com/go-martini/martini/blob/master/martini.go#L68
 [7]: https://github.com/go-martini/martini/blob/master/martini.go#L87
 [8]: https://github.com/go-martini/martini/blob/master/martini.go#L163
 [9]: https://github.com/go-martini/martini/blob/master/router.go#L218
 [10]: https://github.com/go-martini/martini/blob/master/router.go#L315
 [11]: https://github.com/go-martini/martini/blob/master/logger.go
 [12]: #