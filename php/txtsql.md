#PHP txtSQL安装手册中文版
**title:** PHP txtSQL安装手册中文版  
**tags:** php,txtsql,手册,txtsql手册,中文版,文本数据库,txtdb   
**info:** 其实国内有很多的文本数据开源库，但是现在很难能够找到一个相当不错的txtdb了，这篇文章是我搜索到的 08年的一篇文章。


php开源文本数据库 大多使用的有三个 __php-txt-db-api__ ，__txtdb__ 和 __txtsql__ ，txtdb是04年左右的一个开源文本数据库，现在已经找不到了（估计国外服务有）， php txtdb api是个功能很强的文本数据库， 之后是txtsql。文本数据本身性能缺陷无法避免，却有着很方便于开发小型的数据站点的优势。

欢迎使用txtSQL 2.2快速安装手册。这页将指引你如何开始安装txtSQL。

	1-解压缩下载包
	2-配置类文件
		2.1-目录结构
	3-包含类文件
		3.1-类实例
		3.2-连接到txtSQL
		3.3- 更改密码
		3.4-选择一个数据库
	4-执行SQL指令
		4.1-指令的列表 
		4.2-显示结果
	5-从断开txtSQL连接
	6-差错处理
	7-已发布的txtSQL函数

##1、解压缩下载包

当你打开.zip文件时，你将注意到有两个文件： txtSQL.class.php和txtSQL.core.php。提取两个文件到相同的目录。新建一个任意名字的新目录； 通常，它名为data。这将是包含数据库的目录。它能可以放在服务器上的任何地方，但是它通常位于以上两个文件的同一目录下。确保这个目录权限是0755或者更高。现在返回到.zip文件找到'txtsql.MYI'提取它到我们刚刚建立的数据库目录。（译者注：其实不用这么麻烦，.zip文件已经组织好了，全部解压到服务器上的任意目录，并设置权限就行了）

##2、配置类文件

使用txtSQL的第一步，配置类文件，这样它才能被包含到可能要求它的php文件中。首先，你必须在文本编辑器中打开文件txtSQL.class.php 打开文件时将注意到一个版权声明，其后是一些其它素材。随后有这样一行(缺省是第30行)：

	30. include_once('./txtSQL.core.php');
	
这一行代码使它包括txtSQL的的核心函数和类。方便php找到核心文件，你必须编辑单引号内的内容，让它指向txtSQL.core.php文件。（译者注：这个基本上也不用设置，源文件已经配置好了！只有当你的文件不在同一目录时，才需要这么做）

###2.1、目录结构

一个有效的数据库目录结构应该是这样的：

	+ datafolder (所有数据库的保存目录，比如上面新建的'data' )
	+ database_name （文件夹）
	+ table.FRM (列定义)
	+ table.MYD (行数据)
	+ txtsql
	+ txtsql.MYI (包含在压缩包)
	
基本上，一个数据库是主要的数据库目录下的一子目录。  
同时在数据库目录内部是txtsql数据库，压缩包中的'txtsql.MYI'I。  
在所有的数据库内部，一个数据表由两个文件组成； table.FRM，和table.MYD。.FRM是列定义，另一个是数据行。  

##3、包含类文件

现在我们已经配置完txtSQL2.2，我们能开始使用它。首先使用文本编辑器创造一个空白的php文件。保存为example.php。  
为了简单的说明，假设你把它保存在和'txtSQL.class.php'同样的目录下。  
现在我们必须包括php类，在'example.php中输入：  

以下为引用的内容：

	<?php
	include('./txtSQL.class.php');
	?>

###3.1、类实例

在面向对象编程( OOP)中，当创建类时，一种特殊变量类型--个对象是自动地创造。
我们需要创造指向txtSQL类的一个对象，那么把这些添加到文件：

以下为引用的内容：

	<?php
	include('./txtSQL.class.php');
	$sql = new txtSQL('./data');
	?>


在单引号中的文字，是包含所有数据库的数据目录的路径。这个目录下必须包含一个txtsql(大小写敏感 )的目录，目录下应该有一个'txtsql.MYI'的文件。这个文件包含操作数据库所有用户与和密码。  
这个目录与文件已经在txtSQL压缩包中。一旦路径是正确的，你可以继续向前到下一段。

###3.2、连接数据库

现在我们可以用正确的用户名和密码来连接数据库了。  
默认的用户名是root'，默认的密码是空。（强烈建议在下面的步骤中修改)  
用下面的代码来连接数据库：  

以下为引用的内容：

	<?php
	include('./txtSQL.class.php');
	$sql = new txtSQL('./data');
	$sql->connect($username, $password); // 默认时是 $sql->connect('root', '');
	?>

txtSQl这时会认可你是它的用户，准许你访问数据库和表。
注意：参考手册中有可用的命令清单。

###3.3、更改密码

如果你想更改管理员密码（root)，可以用grant_permissions() 函数，grant_permissions() 函数这样调用：

以下为引用的内容：

	<?php
	include('./txtSQL.class.php');
	$sql = new txtSQL('./data');
	$sql->connect($username, $password); // default is $sql->connect('root', '');
	$sql->grant_permissions(action, user, pass [, newpass]);
	?>


参数 action（动作）可以是 add（添加）, drop（删除）, or edit（编辑）. newpass（新密码）只有在你编辑（edit)用户时才可用。user（用户）是用你要操作的用户名, pass是它的密码。  
例如, 如果你想改变用户'root'的密码为 'bar' (假设它还是空的), 我们可以这么做：  
以下为引用的内容：

	<?php
	include('./txtSQL.class.php');
	$sql = new txtSQL('./data');
	$sql->connect($username, $password); // default is $sql->connect('root', '');
	$sql->grant_permissions('edit', 'root', '', 'bar');
	?>

或者新建一个用户 'foo' 密码为'bar'

以下为引用的内容：

	<?php
	include('./txtSQL.class.php');
	$sql = new txtSQL('./data');
	$sql->connect($username, $password); // default is $sql->connect('root', '');
	$sql->grant_permissions('add', 'foo', 'bar');
	?>


或者删除一个用户'foo' 密码为 'bar'  
以下为引用的内容：

	<?php
	include('./txtSQL.class.php');
	$sql = new txtSQL('./data');
	$sql->connect($username, $password); // default is $sql->connect('root', '');
	$sql->grant_permissions('drop', 'foo', 'bar');
	?>


注意：你不用删除用户root'，如果没有正确的密码你也不能访问任何数据。

###3.4、选择数据库

像mySQL一样, 在操作一个数据表之前,你必须先说明它在哪一个数据库. 这个步骤不是必须的，因为你可以在操作时指定使用哪一个数据库.  
我们使用下面的语句来选择一个数据库：

以下为引用的内容：

	<?php
	include('./txtSQL.class.php');
	$sql = new txtSQL('./data');
	$sql->connect($username, $password); // default is $sql->connect('root', '');
	$sql->selectdb('test'); //选择了数据库 'test'
	?>


##4、执行指令
通常我们只要使用sql对象的各种方法下执行指令。  
例如：

以下为引用的内容：

	<?php
	include('./txtSQL.class.php');
	$sql = new txtSQL('./data');
	$sql->connect($username, $password); // default is $sql->connect('root', '');
	$sql->selectdb('test'); // 选择了数据库 'test' 
	$results = $sql->select(array(
	'db' => 'test', //这行不是必须的，因为我们已经选定了数据库
	'table' => 'test',
	'where' => array('id = 10', 'and', 'name =~ John Smith'),
	'limit' => array(0, 100)
	));
	?>


###4.1、指令列表

txtSQL2.2支持的指令如下：

以下为引用的内容：

	4.1- List of commands
	showdbs() 
	createdb() 
	dropdb() 
	renamedb() 
	select() 
	insert() 
	update() 
	delete() 
	showtables() 
	createtable() 
	droptable() 
	altertable() 
	describe() 
	
在执行指令之前，你必须连接数据库，不然会产生错误。手册中会用详细的指令说明和实例（随后翻译）。


###4.2、显示结果
$results变量现在包含了表test'中选中行的信息。
你可以用一个循环来实现显示$results中的所有结果。

以下为引用的内容：

	<?php
	include('./txtSQL.class.php');
	$sql = new txtSQL('./data');
	$sql->connect($username, $password); // default is $sql->connect('root', '');
	$sql->selectdb('test'); // database 'test' is now selected
	$results=
	$sql->execute('select',
	array('select' => array('id', 'name'),
		'db' => 'test',
		'table' => 'test',
		'where' => array('id = 10', 'and', 'name =~ John Smith'),
		'limit' => array(0, 100))));
	foreach ( $results as $key => $row )
	{
		print "ID: $row[id], NAME: $row[name]<BR>\n";
	}
	?>


##5-断开txtSQL
用完之后断开数据库是一个好习惯。断开用 disconnect()函数。

以下为引用的内容：

	<?php
	include('./txtSQL.class.php');
	$sql = new txtSQL('./data');
	$sql->connect($username, $password); // default is $sql->connect('root', '');
	$sql->selectdb('test'); // database 'test' is now selected
	$results=
	$sql->execute('select',
	array('select' => array('id', 'name'),
		'db' => 'test',
		'table' => 'test',
		'where' => array('id = 10', 'and', 'name =~ John Smith'),
		'limit' => array(0, 100))));
	foreach ( $results as $key => $row )
	{
		print "ID: $row[id], NAME: $row[name]<BR>\n";
	}
	$sql->disconnect();
	?>


##6-错误处理
txtSQL 包含错误处理能力。主要用以下的函数：

以下为引用的内容：

	strict() 
	get_last_error() 
	last_error() 
	errordump()

txtsql [下载](http://sourceforge.net/projects/txtsql/)  
_note:_文章来源于[chinaz](http://www.chinaz.com/program/2008/1118/45566.shtml)
