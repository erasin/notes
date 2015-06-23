# 备份与恢复

备份MySQL数据库的命令
mysqldump -hhostname -uusername -ppassword databasename > backupfile.sql
备份MySQL数据库为带删除表的格式备份MySQL数据库为带删除表的格式，能够让该备份覆盖已有数据库而不需要手动删除原有数据库。
mysqldump ---add-drop-table -uusername -ppassword databasename > backupfile.sql
直接将MySQL数据库压缩备份
mysqldump -hhostname -uusername -ppassword databasename | gzip > backupfile.sql.gz
备份MySQL数据库某个(些)表
mysqldump -hhostname -uusername -ppassword databasename specific_table1 specific_table2 > backupfile.sql
同时备份多个MySQL数据库
mysqldump -hhostname -uusername -ppassword --databases databasename1 databasename2 databasename3 > multibackupfile.sql
仅仅备份数据库结构
mysqldump --no-data --databases databasename1 databasename2 databasename3 > structurebackupfile.sql
备份服务器上所有数据库
mysqldump --all-databases  allbackupfile.sql
还原MySQL数据库的命令
mysql -hhostname -uusername -ppassword databasename < backupfile.sql
还原压缩的MySQL数据库
gunzip < backupfile.sql.gz | mysql -uusername -ppassword databasename
将数据库转移到新服务器
mysqldump \-uusername \-ppassword databasename \| mysql \--host=*.*.*.\* \-C databasename

压缩备份

备份并用gzip压缩：
mysqldump < mysqldump options> | gzip > outputfile.sql.gz
从gzip备份恢复：
gunzip < outputfile.sql.gz | mysql < mysql options>
备份并用bzip压缩：
mysqldump < mysqldump options> | bzip2 > outputfile.sql.bz2
从bzip2备份恢复:
bunzip2 < outputfile.sql.bz2 | mysql < mysql options>
