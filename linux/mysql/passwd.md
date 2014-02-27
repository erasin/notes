#RESET PASSWORD

First things first. Log in as root and stop the mysql daemon. 

sudo /etc/init.d/mysql stop 

Now lets start up the mysql daemon and skip the grant tables which store the passwords.

sudo mysqld_safe --skip-grant-tables&

(press Ctrl+C now to disown the process and start typing commands again)

You should see mysqld start up successfully. If not, well you have bigger issues. Now you should be able to connect to mysql without a password.

sudo mysql --user=root mysql

update user set Password=PASSWORD('new-password');
flush privileges;
exit; 

Now kill your running mysqld then restart it normally. 

sudo killall mysqld_safe&
(press Ctrl+C now to disown the process and start typing commands again)
/etc/init.d/mysql start

You should be good to go. Try not to forget your password again.
