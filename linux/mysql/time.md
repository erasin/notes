# 时间

## mysql 

UNIX时间戳转换为日期用函数： FROM_UNIXTIME()

    select FROM_UNIXTIME(1156219870);

日期转换为UNIX时间戳用函数： UNIX_TIMESTAMP()

    Select UNIX_TIMESTAMP(’2006-11-04 12:23:00′);

例：mysql查询当天的记录数：

    $sql=”select * from message Where DATE_FORMAT(FROM_UNIXTIME(chattime),’%Y-%m-%d’) = DATE_FORMAT(NOW(),’%Y-%m-%d’) order by id desc”;

## php 

UNIX时间戳转换为日期用函数： date()

    date('Y-m-d H:i:s', 1156219870);

日期转换为UNIX时间戳用函数：strtotime()

    strtotime('2010-03-24 08:15:42');
