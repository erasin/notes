# db.delete 数据删除
**title:**webpy数据删除操作db.delete  
**tags:**webpy,db.delete  
**info:**数据库操作之删除数据


### 问题

在数据库中删除数据。

### 解决办法

    import web
    
    db = web.database(dbn='postgres', db='mydata', user='dbuser', pw='')
    db.delete('mytable', where="id=10")

上面接受 "using" 和 "vars" 参数。

删除方法返回被删除的影响行数。
