# Web.go tutorial {#webgo_tutorial}

## Getting started {#getting_started}

For installation insructions, check out the [quickstart guide][1].

This is the classic hello world web application, written in Go, which this tutorial often references:


    package main

    import (
        "github.com/hoisie/web"
    )

    func hello(val string) string { 
        return "hello " + val 
    } 

    func main() {
        web.Get("/(.*)", hello)
        web.Run("0.0.0.0:9999")
    }

## URL Handling {#url_handling}

An important part of web applications is knowing how to respond to a given URL path. Web.go provides several methods to add handlers based on a route pattern. These include:

*   [web.Get][2] adds a handler for the most common GET method. It takes two parameters: a regular expression containing the pattern to match, and a function that handles the response.
*   [web.Post][3] adds a handler for the POST method, with a signature similar to `web.Get`.
*   [web.Put][4] and [web.Delete][5] can be used to handle the PUT and DELETE methods.
*   [web.Match][6] can match arbitrary HTTP methods. It takes three arguments - the first is the method name, the second is the path regular expression, and the third is the handler.

### Route handlers {#route_handlers}

In the hello example above, HTTP requests with paths that match the regular expression `/(.*)` are appllied to the following function:

    func hello(val string) string { 
        return "hello " + val 
    } 

Because of the wildcard regular expression, everything after the `/` in the path is passed to the argument *val* in the hello function. A handler function is required to have the same number of arguments as regular expression groups in its url pattern.

Although handlers usually return a string, web.go also accepts handler functions that have no return value. These methods are responsible for writing data to the response directly. The following example is equivalent to the one before it:

    package main

    import (
        "github.com/hoisie/web"
    )

    func hello(ctx *web.Context, val string) { 
        ctx.WriteString("hello " + val)
    } 

    func main() {
        web.Get("/(.*)", hello)
        web.Run("0.0.0.0:9999")
    }

Writing to the context variable is explained in detail in a later section.

## The web.Context type {#the_webcontext_type}

Any handler in web.go can optionally have a pointer to a `web.Context` struct as its first argument. This object contains information about the request and provides methods to control the response.

Here is a brief summary of the fields in [web.Context][7]:

*   `ctx.Request` is a struct of type [http.Request][8] which can be use to retrieve details about the HTTP request, including the params, headers, and files.
*   `ctx.Params` is a map[string]string that contains the request parameters. This is a flattened version of ctx.Request.Form, and is available for convenience.
*   `ctx.ResponseWriter` is the [http.ResponseWriter][9] for the connection. Use it to set the response code, headers, or write to the body.
*   `ctx.Server` is the [web.Server][10] object of the current conection. It contains the configuration details of the server.

A web.Context object gives you access to the HTTP response and lets you set a status code, set response headers, and write to the connection directly.

### Setting a response code {#setting_a_response_code}

Set the status code of an HTTP response by using [ctx.WriteHeader][7]. This method takes a numeric status code as an argument. There are also several convinience methods that set the status code:

*   [ctx.NotFound][11] is used to return an error message along with a 404 status code.
*   [ctx.Abort][12] takes a status code and message as arguments. It is useful for 5xx error messages.
*   [ctx.Redirect][13] is commonly used to send 301 or 302 redirects to a url.
*   [ctx.NotModified][14] sets a 304 status code.

You can also set any arbitrary header code using `ctx.WriteHeader(code)`.

You should only call ctx.WriteHeader once. After it is called, setting additional HTTP response headers will have no effect.

### Setting response headers {#setting_response_headers}

There is a convenience method on web.Context, `SetHeader`, that can be used to set HTTP response headers. The first two parameters are the header key and value, and the last parameter, `unique`, is a boolean that determines whether exiting values for this headers should be overwritten.

You can also call `Header` on the to get the actual response header map. This has methods `Add` and `Set` that can be used to set the HTTP headers as well. When `WriteHeader` is called, this map is included in the response.

    package main

    import (
        "github.com/hoisie/web"
    )

    func hello(ctx *web.Context, val string) string {
        ctx.SetHeader("X-Powered-By", "web.go", true)
        ctx.SetHeader("X-Frame-Options", "DENY", true)
        ctx.SetHeader("Connection", "close", true)
        return "hello " + val 
    } 

    func main() {
        web.Get("/(.*)", hello)
        web.Run("0.0.0.0:9999")
    }

### Setting cookies {#setting_cookies}

web.Context has a method `SetCookie` that takes an http.Cookie object and sets it as a response header. It is often used in conjunction with the helper method web.NewCookie. For example, the following code is a modified hello world example that sets the cookie in the response.

    package main

    import (
        "github.com/hoisie/web"
    )

    func hello(ctx *web.Context, val string) string {
        ctx.SetCookie(web.NewCookie("value", val))
        return "hello " + val 
    } 

    func main() {
        web.Get("/(.*)", hello)
        web.Run("0.0.0.0:9999")
    }

### Writing to the body {#writing_to_the_body}

Web.go gives you the ability to write the body connection directly using `Write` and `WriteString` on web.Context. `WriteHeader(200)` will be called during the first call to `Write` unless it has previously been called.

A `web.Context` object can be used as an `io.Writer` interface. It is useful writing an `io.Reader` object into the body of an HTTP response. This comes up when often writing the contents of a buffer to the response, or when writing the contents of a file or pipe.


    package main

    import (
        "bytes"
        "github.com/hoisie/web"
        "io"
    )

    func hello(ctx *web.Context, val string) {
        var buf bytes.Buffer
        buf.WriteString("hello " + val)
        //copy buf directly into the HTTP response
        io.Copy(ctx, &buf)
    } 

    func main() {
        web.Get("/(.*)", hello)
        web.Run("0.0.0.0:9999")
    }

Using `ctx.Write`, web.go allows you to maintain long-running HTTP connections and send back results periodically. See the [streaming example][15] for more details.

Note, when calling `ctx.Write` and `ctx.WriteString`, the content will be buffered until the handler has completed. Many ResponseWriter objects have a `Flush` method that can be used to flush the content into the HTTP connection. An example of this is also available in the [streaming example][15].

## Templates {#templates}

Templating libraries are very when rendering HTML in web applications. Web.go doesn’t include a templating library. However, there are several good ones available, such as [Go’s http/template library][16] or [mustache][17].

## Static Files {#static_files}

Web.go has the ability to serve static files in a very efficient way. If you place files in the `static` directory of your web application, web.go will serve them if a request matches the name of the file.

For example, if you have a web app in `$HOME/app` being served from `myapp.com`, and there’s a file `$HOME/app/static/image.jpg`, requesting `myapp.com/image.jpg` will serve the static file. A common practice is to have `static/images`, `static/stylesheets`, and `static/javascripts` that contain static files.

Note that Web.go is looking for the `static` directory in both the directory where the web app is located and your current working directory. You can also change web.ServerConfig.StaticDir to specify a directory.

## Shared hosts {#shared_hosts}

Web.go provides methods to run web applications using the SCGI or FastCGI protocols. This enables web.go apps to run in shared hosts environments.

These methods are similar to web.Run:

*   web.RunScgi(addr) serves SCGI requests. (example below)
*   web.RunFcgi(addr) serves FastCGI requests.
*   web.RunTLS(addr, tlsContext) serves HTTPS request. See the [example][18] for details.

For instance, to serve the hello example above running Scgi, just write the following:

    package main

    import (
        "github.com/hoisie/web"
    )

    func hello(val string) string { 
        return "hello " + val 
    } 

    func main() {
        web.Get("/(.*)", hello)
        web.RunScgi("0.0.0.0:6580")
    }

Next you need to configure your web server to pass requests along to port 6580.

## Developing web.go {#developing_webgo}

If you have an issue you’d like to debug in web.go, you’ll want to modify the source code of the project.

By default, when you run `go get package`, the installer will fetch the source code and install the package in the directory specified by `$GOPATH`. If `$GOPATH` is not set, the library is installed in the `$GOROOT/pkg` folder, which defaults to where Go is built. As a developer of Go libraries, it is much more convenient to set a `$GOPATH` variable to a location like `$HOME/golibs` or `$HOME/projects/golibs`. See [how to write Go code][19] for more details. The first step to developing web.go is to ensure that you have an appropriate `$GOPATH` with two sub-directories: `src` and `pkg`.

Next, you should run `cd $GOPATH/src && git clone github.com/hoisie/web`. This creates a clone of the web.go repo in `$GOPATH/src/web`. Next, run `cd GOPATH/src/web && go install`. This will install the `web` package into `$GOPATH/pkg`.

From now on, if you include web.go, you should use `import web`, and this will link to the files that are in `$GOPATH/pkg`. You can now modify the web.go source code in `$GOPATH/src/web`, rebuild with `go install`, and then rebuild your target appication.

 [1]: /index.html
 [2]: /api.html#Get
 [3]: /api.html#Post
 [4]: /api.html#Put
 [5]: /api.html#Delete
 [6]: /api.html#Match
 [7]: /api.html#Context
 [8]: http://golang.org/pkg/net/http/#Request
 [9]: http://golang.org/pkg/net/http/#ResponseWriter
 [10]: /api.html#Server
 [11]: /api.html#Context.NotFound
 [12]: /api.html#Context.Abort
 [13]: /api.html#Context.Redirect
 [14]: /api.html#Context.NotModified
 [15]: https://github.com/hoisie/web/blob/master/examples/streaming.go
 [16]: http://golang.org/pkg/html/template/
 [17]: http://github.com/hoisie/mustache
 [18]: https://github.com/hoisie/web/blob/master/examples/tls.go
 [19]: http://golang.org/doc/code.html