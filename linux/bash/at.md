# at 调度任务

##启动 atd 服务
arch下
	#systemctl enable at.service
init的话，将atd添加即可

## at命令

|任务|说明|
|----|----|
|创建 at 作业。|使用 at time 命令执行操作|
|显示 at 队列。|使用 atq 命令显示 at 队列。|
|验证 at 作业。|使用 atq 命令确认属于特定用户的 at 作业已提交至队列。|
|显示 at 作业。|使用 at -ljob-id 显示已提交至队列的 at 作业。|	
|删除 at 作业。|使用 at -r [job-id] 命令从队列中删除 at 作业。|
|拒绝访问 at 命令。|要拒绝用户访问 at 命令，请编辑 /etc/cron.d/at.deny 文件。|


## 调度单个系统任务

**at 命令的说明**

提交 at 作业文件需执行以下步骤：
	调用 at 实用程序并指定命令执行时间。
	键入以后要执行的命令或脚本。

**注** - 如果此命令或脚本的输出很重要，请确保将输出定向到一个文件中，以便以后检查。

例如，以下 at 作业将在 7 月的最后一天接近午夜时删除用户帐户 smith 的 core 文件。

	$ at 11:45pm July 31
	at> rm /home/smith/*core*
	at> Press Control-d
	commands will be executed using /bin/csh
	job 933486300.a at Tue Jul 31 23:45:00 2004

编辑 `/etc/at.deny` 按照行添加用户，拒绝包含的用户名访问at

## 创建作业 at

	$ at [-m] _time_ [date]

-m 作业完成后发送邮件
_time_ 指定作业的小时（小时[[:]分钟]），不按照24小时制度可以添加 am/pm ; 其中 1630 和 16:30 表达同一时间 
date 指定月份的前三个字母或者更多，一周内某日，或today tomorrow

	周日-周六  Sunday Monday Tueday Wednesday Thursday Friday Saturday 

以下示例说明 jones 如何调度在星期六凌晨 4:00 执行的大型 at 作业。该作业输出被定向到名为 big.file 的文件中。

	$ at 4 am Saturday
	at> sort -r /usr/dict/words > /export/home/jones/big.file

## 显示列队 atq

atq  或 at -l 

## 删除作业 at -r

使用 at -l 或 atq获得 _job-id_ ，之后
	at -r [job-id]
之后使用 at -l [job-id] 验证删除

