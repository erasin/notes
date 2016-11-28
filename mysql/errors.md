# errors

kill 和升级 mysql 都有可能导致 mysql 的数据表损坏

修复工具

修复 my
```bash
myisamchk *.MYI 
// or
myisamchk -e *.MYI 
````

mysql 更新处理

```bash
mysql_upgrade -uroot
```

