# php ctags

创建或编辑 ~/.ctags

    --regex-php=/^[ \t]*[(private| public|static)( \t)]*function[ \t]+([A-Za-z0-9_]+)[ \t]*\(/\1/f, function, functions/
    --regex-php=/^[ \t]*[(private| public|static)]+[ \t]+\$([A-Za-z0-9_]+)[ \t]*/\1/p, property, properties/
    --regex-php=/^[ \t]*(const)[ \t]+([A-Za-z0-9_]+)[ \t]*/\2/d, const, constants/
