# usermod

##语法：

    usermod [-LU][-c <备注>][-d <登入目录>][-e <有效期限>][-f <缓冲天数>][-g <群组>][-G <群组>][-l <帐号名称>][-s <shell>][-u <uid>][用户帐号]

功能说明：修改用户帐号。

_补充说明：usermod可用来修改用户帐号的各项设定。_

## 参数：
* -c<备注> 　修改用户帐号的备注文字。 
* -d登入目录> 　修改用户登入时的目录。 
* -e<有效期限> 　修改帐号的有效期限。 
* -f<缓冲天数> 　修改在密码过期后多少天即关闭该帐号。 
* -g<群组> 　修改用户所属的群组。 
* -G<群组> 　修改用户所属的附加群组。 
* -l<帐号名称> 　修改用户帐号名称。 
* -L 　锁定用户密码，使密码无效。 
* -s<shell> 　修改用户登入后所使用的shell。 
* -u<uid> 　修改用户ID。 
* -U 　解除密码锁定。


##应用举例：

将 newuser2 添加到组 staff 中
    # usermod -G staff newuser2

修改 newuser 的用户名为 newuser1 
    # usermod -l newuser1 newuser

锁定账号 newuser1 
    # usermod -L newuser1

解除对 newuser1 的锁定
    # usermod -U newuser1

