# vim 配置 比较全

    " GLOBAL 环境配置
    " ------------------------------------------------------------------

    " 中文帮助
    set helplang=cn

    " 保留历史 
    set history=500

    " 行控制
    set linebreak               " 英文单词在换行时不被截断
    set nocompatible            " 设置不兼容VI
    "set textwidth=80           " 设置每行80个字符自动换行，加上换行符
    set wrap                    " 设置自动折行
    "set nowrap                 " 设置不自动换行

    " 标签页
    set tabpagemax=15           " 最多15个标签
    "set showtabline=2          " 总是显示标签栏

    " 关闭错误时的声音提示
    set noerrorbells
    set novisualbell
    set t_vb=                   " close visual bell

    " 行号与标尺
    set ruler                   " 右下角显示光标位置的状态行
    set number                  " 显示行号
    set cursorline              " 突出显示当前行
    "set rulerformat=%15(%c%V\ %p%%%)

    " 制表符(设置所有的tab和缩进为4个空格)
    set tabstop=4               " 设置tab键的宽度
    set shiftwidth=4            " 换行时行间交错使用4个空格
    "set cindent shiftwidth=4   " 自动缩进4空格
    set softtabstop=4
    set expandtab               " 使用空格来替换tab
    set smarttab                " 开启新行时使用智能 tab 缩进

    set list                     " 显示Tab符，
    set listchars=tab:\|\ ,      " 使用一高亮竖线代替 把符号显示为 |
    "set listchars=tab:>-,trail:-
    "set listchars=tab:\|\ ,nbsp:%,trail:-

    highlight LeaderTab guifg=#666666   " 设定行首tab为灰色
    match LeaderTab /\t/        " 匹配行首tab

    " 缩进
    set autoindent              " 设置自动缩进
    set smartindent             " 设置智能缩进

    " 搜索
    set hlsearch                " 开启高亮显示结果
    set incsearch               " 开启实时搜索功能
    "set noincsearch             " 关闭显示查找匹配过程
    "set magic     " Set magic on, for regular expressions
    "set showmatch " Show matching bracets when text indicator is over them
    "set mat=2     " How many tenths of a second to blink
    set ignorecase              " 搜索时无视大小写
    set nowrapscan              " 搜索到文件两端时不重新搜索

    " 状态栏显示目前所执行的指令
    set showcmd

    " 自动重新读入
    set autoread                " 当文件在外部被修改，自动更新该文件


    " 备份与缓存
    set nobackup
    set nowb
    "set noswapfile
    set writebackup             " 设置无备份文件

    " 自动完成
    set complete=.,w,b,k,t,i
    set completeopt=longest,menu    " 只在下拉菜单中显示匹配项目，并且会自动插入所有匹配项目的相同文本


    set showmatch               " 显示括号配对情况
    set iskeyword+=_,$,@,%,#,-  " 带有如下符号的单词不要被换行分割
    "set wildmenu "打开 wildmenu 选项，启动具有菜单项提示的命令行自动完成。
    "set matchpairs=(:),{:},[:],<:>
    "set whichwrap=b,s,<,>,[,]

    set backspace=2             " 设置退格键可用
    set mouse=a                 " 鼠标可用
    set ve=block                "光标可以定位在没有实际字符的地方
    "set fullscreen             " 启动后自动全屏

    set nocompatible            " 关闭兼容模式
    set hidden                  " 允许在有未保存的修改时切换缓冲区
    set relativenumber          " 行标跟随
    set clipboard+=unnamed      " 共享外部剪贴板
    set autochdir               " 设定文件浏览器目录为当前目录

    colorscheme default         " 配色方案

    filetype indent on          " 针对不同的文件类型采用不同的缩进格式
    filetype plugin on          " 针对不同的文件类型加载对应的插件
    filetype plugin indent on
    filetype on                 " for taglist

    " 默认为 UTF-8 编码
    " --------------------------------------------------------------- 
    if has("multi_byte")
        set encoding=utf-8
        " English messages only
        "language messages zh_CN.utf-8

        if has('win32')
            language english
            let &termencoding=&encoding " 处理consle输出乱码
        endif

        set fencs=utf-8,gbk,chinese,latin1
        set formatoptions+=mM
        set nobomb " 不使用 Unicode 签名

        if v:lang =~? '^\(zh\)\|\(ja\)\|\(ko\)'
            set ambiwidth=double
        endif
    else
        echoerr "Sorry, this version of (g)vim was not compiled with +multi_byte"
    endif

    " 编码设置
    " -------------------------------------------------------------
    "set fileencoding=utf-8
    set encoding=utf-8 "缓冲编码
    set fileencodings=utf-8,cp936,gb18030,big5,gbk,euc-jp,latin1
    set fileencoding=utf-8
    set termencoding=utf-8 "编码转换
    "set fileformats=unix

    " 命令行与状态行
    " -----------------------------------------------------------------------
    set laststatus=2                          " 开启状态栏信息
    set cmdheight=1                           " 命令行的高度，默认为1，这里设为2
    " 状态行显示的内容 [包括系统平台、文件类型、坐标、所占比例、时间等]
    set statusline=%{fugitive#statusline()}\ %F%m%r%h%w\ [FORMAT=%{&ff}]\ [TYPE=%Y]\ [POS=%l,%v][%p%%]\ %y%r%m%*%=\ %{strftime(\"%d/%m/%y\ -\ %H:%M\")}

    " line color
    " now set it up to change the status line based on mode
    if version >= 700
      au InsertEnter * hi StatusLine term=reverse ctermbg=5 gui=undercurl guisp=Magenta
      au InsertLeave * hi StatusLine term=reverse ctermfg=0 ctermbg=2 gui=bold,reverse
    endif

    function! InsertStatuslineColor(mode)
      if a:mode == 'i'
        hi statusline guibg=magenta
      elseif a:mode == 'r'
        hi statusline guibg=blue
      else
        hi statusline guibg=red
      endif
    endfunction

    au InsertEnter * call InsertStatuslineColor(v:insertmode)
    au InsertChange * call InsertStatuslineColor(v:insertmode)
    au InsertLeave * hi statusline guibg=green

    " GUI MODEL 图形界面
    " ---------------------------------------------
    if has("gui_running")
        colorscheme neon 
        "set guifont=Monospace\ 11
        "set gfw=Zhenyuan:h12:cGB2312
        "au GUIEnter * simalt ~x " 窗口启动时自动最大化
        set guioptions-=m        " 隐藏菜单栏
        set guioptions-=T        " 隐藏工具栏
        set guioptions-=L        " 隐藏左侧滚动条
        set guioptions-=r        " 隐藏右侧滚动条
        "set guioptions-=b       " 隐藏底部滚动条
        "set showtabline=0       " 隐藏Tab栏
        "colorscheme tango2 
        set lines=42
        set columns=136
        "set guifontset=-dt-interface
        set encoding=utf-8
        "language messages zh_CN.UTF-8  "解决consle输出乱码 
        set fdm=indent           "按照缩进进行折叠 manual

        "Toggle Menu and Toolbar 使用F2隐藏/显示菜单
        map <silent> <F2> :if &guioptions =~# 'T' <Bar>
                \set guioptions-=T <Bar>
                \set guioptions-=m <bar>
            \else <Bar>
                \set guioptions+=T <Bar>
                \set guioptions+=m <Bar>
            \endif<CR>

        if has("win32")
            " Windows 兼容配置
            source $VIMRUNTIME/mswin.vim

            " F11 最大化
            map <f11> :call libcallnr('fullscreen.dll', 'ToggleFullScreen', 0)<cr>

            " 字体配置
            "exec 'set guifont='.iconv('Courier_New', &enc, 'gbk').':h10:cANSI'
            "exec 'set guifontwide='.iconv('微软雅黑', &enc, 'gbk').':h10'
            set guifont=YaHei_Consolas_Hybrid:h12:cANSI
            set guifontwide=YaHei_Consolas_Hybrid:h12
        endif

        if has("unix") && !has('gui_macvim')
            set guifont=Ubuntu\ Mono\ 11
        endif

    endif

    " Floding 折叠
    " ------------------------------------------------
    "set foldmethod=syntax " 选择代码折叠类型
    "set foldlevel=100 " 禁止自动折叠
    "
    set foldenable " Enables folding.
    " 折叠方法  
    " manual    手工折叠  
    " indent    使用缩进表示折叠  
    " expr      使用表达式定义折叠  
    " syntax    使用语法定义折叠  
    " diff      对没有更改的文本进行折叠  
    " marker    使用标记进行折叠, 默认标记是 {{{ 和 }}} 
    set foldmethod=manual " Set fold method to 'manual'.
    ""set fdc=0 " Show where folds start and end, when they are opened.
    "nnoremap @=((foldclosed(line('.')) < 0 )? 'zc':'zo')

    :set foldtext=MyFoldText()
    :function MyFoldText()
    :  let line = getline(v:foldstart)
    :  let sub = substitute(line, '/\*\|\*/\|{{{\d\=', '', 'g')
    :  return v:folddashes . sub
    :endfunction

    " 语法高亮 
    " ----------------------------------------------
    syntax enable               " 打开语法高亮
    syntax on                   " 开启文件类型侦测
    au BufRead,BufNewFile *.txt setlocal ft=txt
    au BufRead,BufNewFile *.tpl setlocal ft=smarty 		" 支持 Smarty
    au BufRead,BufNewFile *.as setlocal ft=actionscript	" 支持 ActionScript
    "au BufRead,BufNewFile *.js set ft=javascript syntax=jquery
    au BufRead,BufNewFile jquery.*.js set ft=javascript syntax=jquery
    "autocmd BufRead *.as set filetype=actionscript
    "autocmd BufRead *.js set filetype=javascript
    autocmd BufRead,BufNewFile *.txtfmt set filetype=txtfmt
    "autocmd BufRead,BufNewFile *.txt set filetype=txtfmt
    autocmd BufRead,BufNewFile *.mxml set filetype=mxml
    autocmd BufRead,BufNewFile *.bash set filetype=bash
    autocmd BufRead,BufNewFile *.php set filetype=php
    autocmd BufRead,BufNewFile *.c set filetype=c

    au BufNewFile,BufRead,BufEnter,WinEnter,FileType *.m,*.h setf objc  " 增加 Objective-C 语法支持
    " fencview 自动编码识别     :FencView   查看文件编码和更改文件编码
    let g:fencview_autodetect=1

    " dict
    autocmd FileType javascript set dictionary=~/.vim/dict/javascript.dict

    " 每行超过80个的字符用下划线标示
    au BufRead,BufNewFile *.asm,*.c,*.cpp,*.java,*.cs,*.sh,*.lua,*.pl,*.pm,*.py,*.rb,*.hs,*.vim 2match Underlined /.\%81v/

    " 将指定文件的换行符转换成 UNIX 格式
    au FileType php,javascript,html,css,python,vim,vimwiki set ff=unix

    " 关闭VIM的时候保存会话，按F122读取会话
    set sessionoptions=buffers,sesdir,help,tabpages,winsize
    au VimLeave * mks! ~/Session.vim
    nmap <F7> :so ~/Session.vim<CR>

    " 自动刷新firefox
    autocmd BufWriteCmd *.html,*.js,*.css,*.gtpl :call Refresh_firefox()
    function! Refresh_firefox()
        if &modified
            write
            silent !echo ‘vimYo = content.window.pageYOffset;
                        \ vimXo = content.window.pageXOffset;
                        \ BrowserReload();
                        \ content.window.scrollTo(vimXo,vimYo);
                        \ repl.quit();’ |
                        \ nc localhost 4242 2>&1 > /dev/null
        endif
    endfunction

    " 自动载入VIM配置文件
    autocmd! bufwritepost vimrc source $MYVIMRC

    " 自动匹配括号
    " -----------------------------------
    :inoremap ( ()<ESC>i
    :inoremap ) <c-r>=ClosePair(')')<CR>
    :inoremap { {}<ESC>i
    :inoremap } <c-r>=ClosePair('}')<CR>
    :inoremap [ []<ESC>i
    :inoremap ] <c-r>=ClosePair(']')<CR>
    ":inoremap < <><ESC>i
    ":inoremap > <c-r>=ClosePair('>')<CR>
    :inoremap " ""<ESC>i
    :inoremap ' ''<ESC>i
    :inoremap ` ``<ESC>i

    function ClosePair(char)
        if getline('.')[col('.') - 1] == a:char
            return "\<Right>"
        else
            return a:char
        endif
    endf

    " 快捷键
    " --------------------------------------

    "设置','为leader快捷键
    let mapleader = ","
    let g:mapleader = ","

    " Ctrl + S 保存文件
    map <C-s> <ESC>:w<CR>
    imap <C-s> <ESC>:w<CR>a
    vmap <C-s> <ESC>:w<CR>

    " Ctrl + C 选中状态下复制
    vnoremap <C-c> "+y

    " Ctrl + V 粘贴剪切板中的内容
    map <C-v> "+p
    imap <C-v> <esc>"+pa
    vmap <C-v> d"+P

    "设置快速保存和退出
    "快速保存为,s
    "快速退出（保存）为,w
    "快速退出（不保存）为,q
    nmap <leader>s :w!<cr>
    nmap <leader>w :wq!<cr>
    nmap <leader>q :q!<cr>

    nmap <C-t>   :tabnew<cr>
    nmap <C-p>   :tabprevious<cr>
    nmap <C-n>   :tabnext<cr>
    nmap <C-k>   :tabclose<cr>
    nmap <C-Tab> :tabnext<cr>

    "切换buffer
    nmap bn :bn<cr>
    nmap bp :bp<cr>

    " 插入模式按 F4 插入当前时间
    imap <f4> <C-r>=GetDateStamp()<cr>

    " 返回当前时期
    func! GetDateStamp()
        return strftime('%Y-%m-%d')
    endfunction

    " 选中一段文字并全文搜索这段文字
    vnoremap  *  y/<C-h>=escape(@", '\\/.*$^~[]')<CR><CR>
    vnoremap  #  y?<C-h>=escape(@", '\\/.*$^~[]')<CR><CR>

    "----------------------------------------------------------------------------
    " Plugin TagList
    " ---------------------------------------------------------------------------
    " TagList 键入:Tlist开启
    " if MySys() == "windows"                "设定windows系统中ctags程序的位置
    " 	let Tlist_Ctags_Cmd = 'ctags'
    " elseif MySys() == "linux"              "设定windows系统中ctags程序的位置
        let Tlist_Ctags_Cmd = '/usr/bin/ctags'
    " endif
    map t :TlistToggle<cr>
    let Tlist_Show_One_File = 1            "不同时显示多个文件的tag，只显示当前文件的
    let Tlist_Compact_Format = 1 		   " 使用小窗口.
    let Tlist_Exit_OnlyWindow = 1          "如果taglist窗口是最后一个窗口，则退出vim
    let Tlist_Enable_Fold_Column = 0 	   " 使taglist插件不显示左边的折叠行
    let Tlist_Use_Right_Window = 1         "在右侧窗口中显示taglist窗口
    "let Tlist_Use_Left_Window = 1          " Split to the left side of the screen.
    let Tlist_GainFocus_On_ToggleOpen = 1
    let Tlist_Inc_Winwidth = 0 			   " 防止taglist改变终端窗口的大小
    let Tlist_WinWidth = 25 			   " taglist窗口宽度
    let Tlist_Sort_Type = 'name'           " 排序 name.

    " ---------------------------------------------------------------------------
    " Plugin NERDTree
    " ----------------------------------------------------------------------------
    " plugin - NERD_tree.vim 以树状方式浏览系统中的文件和目录
    " :NERDtree 打开NERD_tree         :NERDtreeClose    关闭NERD_tree
    " o 打开关闭文件或者目录         t 在标签页中打开
    " T 在后台标签页中打开           ! 执行此文件
    " p 到上层目录                   P 到根目录
    " K 到第一个节点                 J 到最后一个节点
    " u 打开上层目录                 m 显示文件系统菜单（添加、删除、移动操作）
    " r 递归刷新当前目录             R 递归刷新当前根目录
    "-----------------------------------------------------------------
    let loaded_netrwPlugin = 0 					" 不启动netrw
    let NERDTreeCaseSensitiveSort=1 			" 让文件排列更有序
    let NERDTreeHijackNetrw = 0 				" 输入:e filename不再显示netrw,而是显示nerdtree
    let NERDTreeChDirMode = 2 					" 改变tree目录的同时不改变工程的目录
    let NERDTreeWinPos = 'left' 				" NERDTree显示位置在窗口右侧
    let NERDTreeWinSize = 25 					" NERDTREE的大小
    let NERDTreeIgnore = [ '^\.svn$', '\~$' ] 	" 忽略.svn的显示
    " F3 NERDTree 切换
    map <F3> :NERDTreeToggle<CR>
    imap <F3> <ESC>:NERDTreeToggle<CR>

    "-----------------------------------------------------------------
    " plugin - NERD_commenter.vim 注释代码用的
    " 将 mapleader 设置为 `,`
    " [count],cc 光标以下count行逐行添加注释(7,cc)
    " [count],cu 光标以下count行逐行取消注释(7,cu)
    " [count],cm 光标以下count行尝试添加块注释(7,cm)
    " ,cA 在行尾插入 /* */,并且进入插入模式。 这个命令方便写注释
    " ,c<space> toggle 注释
    " 注：count参数可选，无则默认为选中行或当前行
    "-----------------------------------------------------------------
    let NERDSpaceDelims=1       " 让注释符与语句之间留一个空格
    let NERDCompactSexyComs=1   " 多行注释时样子更好看

    " Plugin MiniBufExplorer 多个文件切换 可使用鼠标双击相应文件名进行切换
    ""let g:miniBufExplMapWindowNavVim = 1
    ""let g:miniBufExplMapWindowNavArrows = 1
    ""let g:miniBufExplMapCTabSwitchBufs = 1
    ""let g:miniBufExplModSelTarget = 1

    " --------------------------------------------------------------------
    " plugin Load_Template 根据文件后缀自动加载模板，使用:LoadTemplate呼出
    " --------------------------------------------------------------------
    " let g:template_path = '~/.vim/templates'

    " ---------------------------------------------------------------------
    " plugin indent guides 对齐线
    " ---------------------------------------------------------------------
    " <Leader>ig     toggle
    let g:indent_guides_guide_size=1    "设置宽度

    " ---------------------------------------------------------------------
    " 处理文件
    " ---------------------------------------------------------------------
    " **** PHP **** {{{
    " 不显示PHP变量再Taglist中
    let tlist_php_settings = 'php;c:class;d:constant;f:function'
    " 高亮显示sql语句
    let php_sql_query = 1 
    " }}}

    " **** Python *** {{{
    " For lines that end with \n\ or \ and continue on the next one.
    "let g:pyindent_continue = '&sw - &sw'
    let g:pyindent_continue = 0
    autocmd Filetype python set completefunc=pythoncomplete#Complete
    " }}}


    " 一键保存和编译 
    " -----------------------------------------
    " Ctrl + G 一键保存、编译
    " Ctrl + R 一键保存、运行
    "
    " 编译C源文件
    func! CompileGcc()
        exec "w"
        let compilecmd="!gcc -Wall -std=c99 "
        let compileflag="-o %<"
        exec compilecmd." % ".compileflag
    endfunc

    " 编译C++源文件
    func! CompileCpp()
        exec "w"
        let compilecmd="!g++ -Wall "
        let compileflag="-o %<"
        exec compilecmd." % ".compileflag
    endfunc

    " 编译Haskell源文件
    func! CompileHaskell()
        exec "w"
        let compilecmd="!ghc --make "
        let compileflag="-o %<"
        exec compilecmd." % ".compileflag
    endfunc

    " 编译Java源文件
    func! CompileJava()
        exec "w"
        exec "!javac %"
    endfunc

    " 编译C#源文件
    func! CompileCs()
        exec "w"
        exec "!csc %"
    endfunc

    " 编译Gas源文件
    func! CompileGas()
        exec "w"
        exec "!gcc -Wall -ggdb -o %< %"
    endfunc

    " 运行Shell源文件
    func! RunShell()
        exec "w"
        exec "!sh %"
    endfunc

    " 运行Lua源文件
    func! RunLua()
        exec "w"
        exec "!lua %"
    endfunc

    " 运行Perl源文件
    func! RunPerl()
        exec "w"
        exec "!perl %"
    endfunc

    " 运行Python源文件
    func! RunPython()
        exec "w"
        exec "!python %"
    endfunc

    " 运行Ruby源文件
    func! RunRuby()
        exec "w"
        exec "!ruby %"
    endfunc

    " 检测PHP程序语法
    func! RunPHP()
        exec "w"
        exec "!php -l %"
    endfunc

    " vala
    func! CompileVala()
        exec "w"
        exec "!valac --pkg gtk+-3.0 %"
        exec "!./%<"
    endfunc

    " 根据文件类型自动选择相应的编译函数
    func! CompileCode()
        exec "w"
        if &filetype == "c"
            exec "call CompileGcc()"
        elseif &filetype == "cpp"
            exec "call CompileCpp()"
        elseif &filetype == "haskell"
            exec "call CompileHaskell()"
        elseif &filetype == "java"
            exec "call CompileJava()"
        elseif &filetype == "cs"
            exec "call CompileCs()"
        elseif &filetype == "asm"
            exec "call CompileGas()"
        elseif &filetype == "sh"
            exec "call RunShell()"
        elseif &filetype == "lua"
            exec "call RunLua()"
        elseif &filetype == "perl"
            exec "call RunPerl()"
        elseif &filetype == "python"
            exec "call RunPython()"
        elseif &filetype == "ruby"
            exec "call RunRuby()"
        elseif &filetype == "php"
            exec "call RunPHP()"
        elseif &filetype == "vala"
            exec "call CompileVala()"
        endif
    endfunc

    " 运行可执行文件
    func! RunResult()
        exec "w"
        if &filetype == "c"
            exec "! %<"
        elseif &filetype == "cpp"
            exec "! %<"
        elseif &filetype == "haskell"
            exec "! %<"
        elseif &filetype == "java"
            exec "!java %<"
        elseif &filetype == "cs"
            exec "! %<"
        elseif &filetype == "asm"
            exec "! %<"
        elseif &filetype == "sh"
            exec "!sh %<.sh"
        elseif &filetype == "lua"
            exec "!lua %<.lua"
        elseif &filetype == "perl"
            exec "!perl %<.pl"
        elseif &filetype == "python"
            exec "!python %<.py"
        elseif &filetype == "ruby"
            exec "!ruby %<.rb"
        endif
    endfunc


    " Ctrl + G 一键保存、编译
    " Ctrl + R 一键保存、运行
    map <C-g> :call CompileCode()<CR>
    imap <C-g> <ESC>:call CompileCode()<CR>
    vmap <C-g> <ESC>:call CompileCode()<CR>

    map <C-r> :call RunResult()<CR>
    imap <C-r> <ESC>:call RunResult()<CR>
    vmap <C-r> <ESC>:call RunResult()<CR>

    " 对齐线高亮显示 
    " 利用自带的 cc 即 colorcolum
    " -----------------------------------------------------
    " 使用  ,ch 来toggle对齐线
    map ,ch :call SetColorColumn()<CR>
    function! SetColorColumn()
        let col_num = virtcol(".")
        let cc_list = split(&cc, ',')
        if count(cc_list, string(col_num)) <= 0
            execute "set cc+=".col_num
        else
            execute "set cc-=".col_num
        endif
    endfunction

