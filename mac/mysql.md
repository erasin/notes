# mysql 工具


    brew install mysql

> 建议使用 docker 来创建 mysql。

A "/etc/my.cnf" from another install may interfere with a Homebrew-built
server starting up correctly.

To connect:

    mysql -uroot

To have launchd start mysql at login:

    ln -sfv /usr/local/opt/mysql/*.plist ~/Library/LaunchAgents

Then to load mysql now:

    launchctl load ~/Library/LaunchAgents/homebrew.mxcl.mysql.plist

Or, if you don't want/need launchctl, you can just run:

    mysql.server start

# 删除pkg mysql

```bash
sudo rm /usr/local/mysql

sudo rm -rf /usr/local/mysql*

sudo rm -rf /Library/StartupItems/MySQLCOM

sudo rm -rf /Library/PreferencePanes/My*

#vim /etc/hostconfig and removed the line
# MYSQLCOM=-YES-

sudo rm -rf /Library/Receipts/mysql*
sudo rm -rf /Library/Receipts/MySQL*
sudo rm -rf /var/db/receipts/com.mysql.*
```

