SDL2 API手册
目录：
（1）SDL_CreateWindow(const char *title,int x,int y,int w,int h,Unit32 flags)
（2）SDL_CreateRenderer(SDL_Window *window,int index,Uint32 flags)
（3）SDL_LoadBMP(const char *file)
（4）SDL_CreateTextureFromSurface(SDL_Renderer *renderer,SDL_Surface *surface)
（5）SDL_FreeSurface(SDL_Surface *surface)
（6）SDL_RenderClear(SDL_Renderer *renderer)
（7）SDL_RenderCopy(SDL_Renderer *renderer,SDL_Texture *texture,const SDL_Rect *srrect,const SDL_Rect *dstrect)
（8）SDL_RenderPresent(SDL_Renderer *renderer)
（9）SDL_DestroyTexture(SDL_Texture *texture)
（10）SDL_DestroyRenderer(SDL_Render *renderer)
（11）SDL_DestroyWindow(SDL_Window *window)
（12）SDL_Delay(Uint32 ms)
（13）SDL_Init(Uint32 flags)
（14）SDL_Quit()
（15）SDL_FillRect(SDL_Surface *dst,const SDL_Rect *rect,Uint32 color)
（16）SDL_MapRGB(const SDL_PixelFormat *format,Uint8 r,Uint8 g,Uint8 b)
（17）SDL_UpdateWindowSurface(SDL_Window *window)


（1）SDL_Window *SDL_CreateWindow(const char *title,int x,int y,int w,int h,Unit32 flags)：创建一个窗口。成功，则返回一个SDL_Window指针；失败，则返回一个NULL。
1）第一个参数title：为窗口标题。
2）第二（或三）个参数x（或y）为：窗口坐标。
3）第四（或五）个参数w（h）为：窗口的长（宽）。
4）第六个参数flags：窗口标签。它可以是：S
DL_WINDOW_FULLSCREEN（全屏），
SDL_WINDOW_OPENGL，
SDL_WINDOW_SHOWN（显示），
SDL_WINDOW_HIDDEN（隐藏），
SDL_WINDOW_BORODERLESS，
SDL_WINDOW_RESIZABLE，
SDL_WINDOW_MAXMIZED（最大化），
SDL_WINDOW_MINIMIZED（最小化），
SDL_WINDOW_INPUT_GRABBED，
SDL_WINDOW_ALLOW_HIGHPL
5）备注：（1）参数title用的是UTF-8编码的，如果你使用了其他编码的字符串，那么你得到的窗口标题可能为乱码。（2）窗口坐标（x，y）也可以是SDL_WINDOWPOS_UNDEFINED
即SDL window position undefined，翻译成中文就是：窗口未被定义，暂时不知道它的用处，经过实验，我想应该是显示在显示器中间。这样做的作用有：不必知道显示器的尺寸，也不用去计算窗口的坐标，即可让窗口显示在中央了。

（2）SDL_Renderer *SDL_CreateRenderer(SDL_Window *window,int index,Uint32 flags):为指定窗口创建一个2D渲染环境（rendering context）。成功，返回一个可用的渲染环境；失败，返回NULL。
1）第一个参数window：指定要渲染的窗口；
2）第二个参数index：用于指定要启动的渲染驱动（rendering driver）；如果是是-1，则启动第一个可用的渲染驱动，或者说SDL自动指定为第三个参数指定的驱动；
3）第三个flags：渲染器标签，如SDL_RENDERER_ACCELERATED，为硬件加速（借助于显卡），其值为2，即可用2代替，只是意思不明了。其他的flags自行百度。

（3）SDL_LoadBMP(const char *file)：载入bmp格式的图片。参数为图片路径，如”C:\Picture\a.bmp”。成功，返回SDL_Surface型的指针（指向加载的图片）；失败，返回NULL。

（4）SDL_Texture* SDL_CreateTextureFromSurface(SDL_Renderer *renderer,SDL_Surface *surface):从一个已存在并还存在（未释放）的表面中创造一个纹理(texture)。（注：要有效的利用硬件加速，就要把SDL_Surface转化为SDL_Texture，但是不会被改变或释放surface）。成功，返回被创建了的纹理；失败则返回0（零）。
1）第一个参数renderer：指定的渲染器。
2）第二个参数surface：去填充纹理的包含了像素的数据（即图片）。

（5）SDL_FreeSurface(SDL_Surface *surface)：释放一个表面。参数surface即为要释放的表面。

（6）int SDL_RenderClear(SDL_Renderer *renderer)：清空指定的渲染器。成功，返回0（零）；失败，返回-1（负一）。

（7）int SDL_RenderCopy(SDL_Renderer *renderer,SDL_Texture *texture,const SDL_Rect *srrect,const SDL_Rect *dstrect)：将纹理（texture）的一部分复制到当前的渲染器。成功，返回0（零）；失败，返回-1（负一）。
1）第一个参数renderer：用于指定一个渲染器。
2）第二个参数texture：用于指定一个纹理。
3）第三个参数srcrect：指定显示texture的哪个部分。NULL为整个texture。
4）第四个参数dstrect：指定texture在renderer的哪个部分显示。NULL为整个renderer。

（8）SDL_RenderPresent(SDL_Renderer *renderer):刷新屏幕。（或者说，Update the screen with rendering performed）。

（9）SDL_DestroyTexture(SDL_Texture *texture):销毁纹理（texture）。参数texture用于指定要销毁的纹理。

（10）SDL_DestroyRenderer(SDL_Render *renderer)：销毁渲染器（renderer）。参数renderer用于指定要销毁的渲染器。

（11）SDL_DestroyWindow(SDL_Window *window)：销毁窗口（window）。参数window用于指定要销毁的窗口。

（12）SDL_Delay(Uint32 ms)：等待（暂停）ms毫秒。参数ms指定等待秒数。如：SDL_Delay(3000)，意思是暂停3000毫（3000ms），即3秒（3s）。

（13）int SDL_Init(Uint32 flags)：启动用flag指定的SDL子系统。参数flags用于指定要启动的子系统。
Flags可为如下参数：
1）SDL_INIT_TIMER：启动时间系统。
2）SDL_INIT_AUDIO：启动声频系统。
3）SDL_INIT_VIDEO：启动视频系统。
4）SDL_INIT_EVENTS：启动事件系统。
5）SDL_INIT_JOYSTICK：启动游戏杆（joystick）系统。（就是打游戏用的那个遥控器）。
6）SDL_INIT_HAPTIC：启动触觉系统。
7）SDL_INIT_GAMECONTROLLER：启动游戏控件。（英语上是game controller不知道准不准确）
8）SDL_INIT_EVERYTING：启动所有子系统。
9）备注：如果你不想启动所有子系统而又想启动不少于两个的子系统，你可以使用 | 将它们隔开。如：SDL_INIT_AUDIO | SDL_INIT_VIDEO。这里就只启动了声频和视频系统。实际上SDL_INIT_EVERYTING是一个宏，它是被这样定义的：#define SDL_INIT_EVERYTING (SDL_INIT_TIMER | SDL_INIT_AUDIO | SDL_INIT_VIDEO | SDL_INIT_EVENTS | SDL_INIT_JOYSTICK | SDL_INIT_HAPTIC | SDL_INIT_GAMECONTROLLER )  。如此一来，用SDL_INIT_EVERYTING就可以启动所有子系统了。

（14）SDL_Quit()：退出所有启动了的SDL子系统。

（15）SDL_Surface *SDL_GetWindowSurface(SDL_Window *window)：得到一个有窗口的SDL表面（SDL surface）。成功，返回一个带有帧缓冲器（frame buffer）窗口；失败，返回NULL。参数window可用SDL_CreateWindow得到。
1）备注：1）这个函数会为参数window创建一个优化了的新的表面。
 2）新的表面会在参数window指定的窗口销毁时释放。
 3）此外，不能再创建的表面使用3D或渲染API。

（15）int SDL_FillRect(SDL_Surface *dst,const SDL_Rect *rect,Uint32 color)：将窗口的指定部分填满（填充）某颜色。成功，返回0（零）；失败，返回-1（负一）。
1）参数dst：指定要填充的表面。
2）参数rect：指定填充窗口的哪个部分。如果为零（NULL），则为整个表面。
3）参数color：是指定的表面的像素像素格式。也可以用SDL_MapRGB()函数得到。如：直接写成SDL_MapRGB(dst->format,0xF,0,0)。
4）举例（略去了变量声明）：SDL_FillRect(dst,NULL,SDL_MapRGB(dst->format,0xFF,0,0))。意思是将dst这个表面的整个表面（NULL）染成红色（0xF,0,0）。
5）看了SDL_MapRGB()可能会更明白些。

（16）Uint32 SDL_MapRGB(const SDL_PixelFormat *format,Uint8 r,Uint8 g,Uint8 b)：将RGB（红、绿、蓝）三色的值映射到一个给定的表面的格式（format）中。
1）参数format：一般是一个表面里的format（参见SDL_FillRect()里的举例，就在上面）。
2）参数（r，g，b）：指定红绿蓝三色的数值。如(0xFF,0,0)表示红色（只有r有数值，其他为零，所以为红），(0xFF,0xFF,0xFF)白色（所有颜色数值为最大，所以为白色）。(0,0,0)黑色（所有颜色的数值为零，所以为黑）。
3）科普：为什么是红绿蓝三色（RGB）？在物理课上，我们学过色光三原色，即红光、绿光、蓝光。把这三种光按不同比例混合，即可得到所有颜色的光。三种光等量混合可得白色。而黑色不是因为存在黑光，而是因为没有光。（只在可见光范围内讨论）平时我们看到的黑色是因为周围的物体发出或反射的进入我们的眼睛，使得我们感觉起来那里“黑的”。这样你就明白黑夜为什么是黑的了，因为没有光，或光很弱无法引起你的视觉。物理课没上好可能不好理解这里，不过不要紧，你只要知道通过不同的色光数值的组合可以得到不同的颜色即可。

（17）int SDL_UpdateWindowSurface(SDL_Window *window)：将窗口表面复制到窗口表面。成功，返回0（零）；失败，返回-1（负一）。简单的理解可以为：刷新一下窗口。
