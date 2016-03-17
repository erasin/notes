# brew install php 

brew install php --with-pear

pear install PHP_CodeSniffer




To enable PHP in Apache add the following to httpd.conf and restart Apache:
    LoadModule php5_module    /usr/local/opt/php55/libexec/apache2/libphp5.so

The php.ini file can be found in:
    /usr/local/etc/php/5.5/php.ini

✩✩✩✩ PEAR ✩✩✩✩

If PEAR complains about permissions, 'fix' the default PEAR permissions and config:
    chmod -R ug+w /usr/local/Cellar/php55/5.5.18/lib/php
    pear config-set php_ini /usr/local/etc/php/5.5/php.ini

✩✩✩✩ Extensions ✩✩✩✩

If you are having issues with custom extension compiling, ensure that
you are using the brew version, by placing /usr/local/bin before /usr/sbin in your PATH:

      PATH="/usr/local/bin:$PATH"

PHP55 Extensions will always be compiled against this PHP. Please install them
using --without-homebrew-php to enable compiling against system PHP.

✩✩✩✩ PHP CLI ✩✩✩✩

If you wish to swap the PHP you use on the command line, you should add the following to ~/.bashrc,
~/.zshrc, ~/.profile or your shell's equivalent configuration file:

      export PATH="$(brew --prefix homebrew/php/php55)/bin:$PATH"

To have launchd start php55 at login:
    ln -sfv /usr/local/opt/php55/*.plist ~/Library/LaunchAgents
Then to load php55 now:
    launchctl load ~/Library/LaunchAgents/homebrew.mxcl.php55.plist
