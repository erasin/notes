# php ctags

创建或编辑 ~/.ctags

    --regex-php=/^[ \t]*[(private| public|static)( \t)]*function[ \t]+([A-Za-z0-9_]+)[ \t]*\(/\1/f, function, functions/
    --regex-php=/^[ \t]*[(private| public|static)]+[ \t]+\$([A-Za-z0-9_]+)[ \t]*/\1/p, property, properties/
    --regex-php=/^[ \t]*(const)[ \t]+([A-Za-z0-9_]+)[ \t]*/\2/d, const, constants/

使用ctags 创建

    ctags -R -f .tags


## subl 中的 keybinding

    ==============================  ================  ===========  ======================
    Command                         Key Binding       Alt Binding  Mouse Binding
    ==============================  ================  ===========  ======================
    rebuild_ctags                   ctrl+t, ctrl+r
    navigate_to_definition          ctrl+t, ctrl+t    ctrl+>       ctrl+shift+left_click
    jump_prev                       ctrl+t, ctrl+b    ctrl+<       ctrl+shift+right_click
    show_symbols                    alt+s
    show_symbols (all files)        alt+shift+s
    show_symbols (suffix)           ctrl+alt+shift+s
