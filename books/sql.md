# MYSQL SQL 简明速记

## 语法

* 大小写不敏感
* 多行语句时候，句子末尾加`;`。
* 注释
	* `/*...*/`
	* `--- ... ---`
	* `# ...`
* 数据操作语言 (DML)
	* **SELECT**      - 从数据库表中获取数据
	* **UPDATE**      - 更新数据库表中的数据
	* **DELETE**      - 从数据库表中删除数据
	* **INSERT INTO** - 向数据库表中插入数据
* 数据定义语言 (DDL) 。 下面重要的DDL语句
	* **CREATE DATABASE** - 创建新数据库
	* **ALTER DATABASE**  - 修改数据库
	* **CREATE TABLE**    - 创建新表
	* **ALTER TABLE**     - 变更（改变）数据库表
	* **DROP TABLE**      - 删除表
	* **CREATE INDEX**    - 创建索引（搜索键）
	* **DROP INDEX**      - 删除索引

## CREATE 创建 DROP 删除

创建数据库

	CREATE DATABASE db_name;

创建数据表

	CREATE TABLE table_name(
		column_name ...
	);


数据类型		                                       | 描述
-------------------------------------------------------|--------------------------------
integer(size) //int(size)/smallint(size)/tinyint(size) | 仅容纳整数。在括号内规定数字的最大位数。
decimal(size,d)/numeric(size,d)                        | 容纳带有小数的数字。 "size" 规定数字的最大位数。"d" 规定小数点右侧的最大位数。
char(size)                                             | 容纳固定长度的字符串（可容纳字母、数字以及特殊字符）。在括号中规定字符串的长度。
varchar(size)                                          | 容纳可变长度的字符串（可容纳字母、数字以及特殊的字符）。在括号中规定字符串的最大长度。
date(yyyymmdd)                                         | 容纳日期。

创建视图

	CREATE VIEW view_name AS
	SELECT x,y,z
	FROM table_name
	where x=a

删除 DROP

	DROP DATABASE db_name;
	DROP TABLE table_name;
	DROP VIEW view_name;

## DML

选择 **SELECT**  

	SELECT [*[field1,...]] FROM table_name

条件 **WHERE** ：引用与 SELECT,UPDATE,DELETE

	SELECT * FROM table_name WHERE f1=v1 

**AND** 和 **OR** 可在 WHERE 子语句中把两个或多个条件结合起来,也可以把 AND 和 OR 结合起来（使用圆括号来组成复杂的表达式）

	SELECT * FROM table_name WHERE (f1=v1 OR f2 >v2) AND f3=v3

插入 **INSERT** 

	INSERT INTO table_name [(field1,...)] VALUES (value1,...)

更新 **UPDATE**

	UPDATE table_name SET field_name = value[,...] WHERE field2_name = value [...]

删除 **DELETE**

	DELETE FROM table_name WHERE field_name = value

## 高级

### SELECT 的一些功能

**DISTINCT** field 关键词 DISTINCT 用于返回唯一不同的值。

	SELECT DISTINCT field FROM table_name;  -- 选择 field 不同的值 

**LIMIT** number_limit 
**LIMIT** number_start, number_limit

	SELECT * FROM table_name limit 3,5; -- 选择从第三行开始之后的5行

**ALIAS** 可以为列名称和表名称指定别名（Alias）。

	SELECT f1(s) as f1a, alias_name2.f2 AS f2a                    -- 列名别名
	FROM table_name AS alias_name1, table_name2 AS alias_name2;   -- 表别名:w


### WHERE 中的一个功能

**LIKE** / **NOT LIKE** 

	SELECT f1 FROM table_name WHERE f2 LIKE '%value%'

#### 通配符

通配符                    | 描述
--------------------------| -------------
`%`                       | 1个或多个
`_`                       | 1个字符
[charlist]                | 字符列中的任何单一字符
[^charlist] / [!charlist] | 不再字符列中的任何的一个单一字符

	SELECT * FROM table_name WHERE f1 LIKE '[!ALN]%'; -- 选在不以 ALN开头的任何内容


**IN** : WHERE的多个值

	SELECT column_name(s)
	FROM table_name
	WHERE column_name IN (value1,value2,...)

**[NOT] BETWEEN ... AND** 会选取介于两个值之间的数据范围，这些值可以是数值、文本或者日期。

	SELECT column_name(s)
	FROM table_name
	WHERE column_name
	BETWEEN value1 AND value2

### JOIN

用于根据两个或多个表中的列之间的关系，从这些表中查询数据

* **JOIN** / **INNER JOIN** : 如果表中有至少一个匹配，则返回行
* **LEFT JOIN**             : 即使右表中没有匹配，也从左表返回所有的行
* **RIGHT JOIN**            : 即使左表中没有匹配，也从右表返回所有的行
* **FULL JOIN**             : 只要其中一个表中存在匹配，就返回行

**JOIN** 

	SELECT Persons.LastName, Persons.FirstName, Orders.OrderNo
	FROM Persons, Orders
	WHERE Persons.Id_P = Orders.Id_P 

使用 JION ;

	SELECT Persons.LastName, Persons.FirstName, Orders.OrderNo
	FROM Persons
	INNER JOIN Orders
	ON Persons.Id_P = Orders.Id_P
	ORDER BY Persons.LastName

**LEFT JOIN**

	SELECT column_name(s)
	FROM table_name1
	LEFT JOIN table_name2 
	ON table_name1.column_name=table_name2.column_name

**RIGHT JOIN**

	SELECT column_name(s)
	FROM table_name1
	RIGHT JOIN table_name2 
	ON table_name1.column_name=table_name2.column_name

**FULL JOIN**

	SELECT column_name(s)
	FROM table_name1
	FULL JOIN table_name2 
	ON table_name1.column_name=table_name2.column_name

### UNION 操作符

UNION 操作符用于合并两个或多个 SELECT 语句的结果集。

请注意，UNION 内部的 SELECT 语句必须拥有相同数量的列。列也必须拥有相似的数据类型。同时，每条 SELECT 语句中的列的顺序必须相同。
UNION 命令只会选取不同的值。

	SELECT column_name(s) FROM table_name1
	UNION
	SELECT column_name(s) FROM table_name2

**UNION ALL** 命令和 UNION 命令几乎是等效的，不过 UNION ALL 命令会列出所有的值。

###  SELECT INTO 语句

SELECT INTO 语句从一个表中选取数据，然后把数据插入另一个表中。

SELECT INTO 语句常用于创建表的备份复件或者用于对记录进行存档。

您可以把所有的列插入新表：

	SELECT *
	INTO new_table_name [IN externaldatabase] --注意数据库
	FROM old_tablename
	[WHERE column_name = value]

或者只把希望的列插入新表：

	SELECT column_name(s)
	INTO new_table_name [IN externaldatabase] 
	FROM old_tablename


### Constraints 约束 用于创建修改表

* NOT NULL
* UNIQUE
* PRIMARY KEY
* FOREIGN KEY
* CHECK
* DEFAULT









