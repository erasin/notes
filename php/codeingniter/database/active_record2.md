# Active Record 类

CodeIgniter 使用的是修改过的 Active Record 数据库模式。 这种模式是以较少的程序代码来实现信息在数据库中的获取，插入，更改。 有时只用一两行的代码就能完成对数据库的操作。 CodeIgniter 不需要每一个数据库表拥有自己的类。它提供了一个更简单的接口。

不只是简单，使用 Active Record 的一个主要的优点是允许你创建独立的数据库应用程序，因为查询语法是由数据库的适配器来产生的。它可以进行更安全的查询，因为系统会自动的对所有的输入值进行转义。
  **注意:** 你如果想要写自己的查询表达式，可以在数据库配置文件中禁止这个类。同时相比 Active Record，使用核心数据库类库和适配器将会使用更少资源。<br />
*   [选择数据][1]
*   [插入数据][2]
*   [更新数据][3]
*   [删除数据][4]
*   [链式方法][5]
*   [Active Record 缓存][6]

<a name="select"> </a> 
# 选择数据

下面的函数帮助你构建 SQL **SELECT**语句。

**备注：如果你正在使用 PHP5，你可以在复杂情况下使用链式语法。本页面底部有具体描述。**

## $this->db->get();

运行选择查询语句并且返回结果集。可以获取一个表的全部数据。

    $query = $this->db->get('mytable');  

    // Produces: SELECT * FROM mytable

第二和第三个参数允许你设置一个结果集每页纪录数(limit)和结果集的偏移(offset)

    $query = $this->db->get('mytable', 10, 20);  

    // Produces: SELECT * FROM mytable LIMIT 20, 10 (in MySQL. Other databases have slightly different syntax)

注意：第二参数是每页纪录数，第三个参数是偏移

你会注意到上面的函数由一个变量<kbd>$query</kbd>执行，这个<kbd>$query</kbd>可以用来显示结果集。

    $query = $this->db->get('mytable');  

    foreach ($query->result() as $row)  
    {  
        echo $row->title;  

    }

请访问[查询结果][7]页查看详细的生成结果的方法。

## $this->db->get_where();

跟上面的函数一样，只是它允许你在函数的第二个参数那里添加一个 where 从句，从而不用使用 db->where() 这个函数：

    $query = $this->db->get_where('mytable', array('id' => $id), $limit, $offset);

请阅读下面的 where 函数了解更多。
  注意: get_where() 在以前的版本中写作 getwhere()，这是已经过时的用法，现已从代码中移除 getwhere()。
## $this->db->select();

允许你在SQL查询中写 SELECT 部分:
    $this->db->select('title, content, date');  

    $query = $this->db->get('mytable');  
    // Produces: SELECT title, content, date FROM mytable
  **注意:** 如果你要查询表中的所有行，你可以不用写这个函数。省略后，CodeIgniter 会认为你要查询 全部行(SELECT *)。
$this->db->select() 可接受一个可选的第二个参数。如果你把它设为FALSE， CodeIgniter 将不会使用反引号保护你的字段或者表名 。这在进行复合查询时很有用。

    $this->db->select("('SELECT SUM(payments.amount) FROM payments WHERE payments.invoice_id=4') AS amount_paid", FALSE);   
    $query = $this->db->get('mytable');  

## $this->db->select_max();

为你的查询编写一个 "SELECT MAX(field)"。你可以选择性的给出第二个参数，以便重命名结果字段名。
    $this->db->select_max('age');  
    $query = $this->db->get('members');  

    // Produces: SELECT MAX(age) as age FROM members  

    $this->db->select_max('age', 'member_age');  

    $query = $this->db->get('members');  
    // Produces: SELECT MAX(age) as member_age FROM members

## $this->db->select_min();

为你的查询编写一个 "SELECT MIN(field)" 。与 _select_max()_ 相似, 你可以选择性地给出第二个参数，用来给结果字段重命名。
    $this->db->select_min('age');  
    $query = $this->db->get('members');  

    // 生成: SELECT MIN(age) as age FROM members

## $this->db->select_avg();

为你的查询编写一个 "SELECT AVG(field)" 。与 _select_max()_ 相似, 你可以选择性地给出第二个参数，用来给结果字段重命名。
    $this->db->select_avg('age');  
    $query = $this->db->get('members');  
    // 生成: SELECT AVG(age) as age FROM members

## $this->db->select_sum();

为你的查询编写一个 "SELECT SUM(field)" 。与 _select_max()_ 相似, 你可以选择性地给出第二个参数，用来给结果字段重命名。
    $this->db->select_sum('age');  
    $query = $this->db->get('members');  
    // 生成: SELECT SUM(age) as age FROM members

## $this->db->from();

允许你编写查询中的FROM部分:
    $this->db->select('title, content, date');  
    $this->db->from('mytable');  

    $query = $this->db->get();  

    // 生成: SELECT title, content, date FROM mytable
  说明: 正如前面所说，查询中的FROM部分可以在 _$this->db->get()_ 函数中指定，所以你可以根据自己的喜好来选择使用哪个方法。
## $this->db->join();

允许你编写查询中的JOIN部分:
    $this->db->select('*');  
    $this->db->from('blogs');  
    $this->db->join('comments', 'comments.id = blogs.id');  

    $query = $this->db->get();  
    // 生成:   
    // SELECT * FROM blogs  
    // JOIN comments ON comments.id = blogs.id  
如果你想要在查询中使用多个连接，可以多次调用本函数。

如果你需要指定 JOIN 的类型，你可以通过本函数的第三个参数来指定。可选项包括：left, right, outer, inner, left outer, 以及 right outer.
    $this->db->join('comments', 'comments.id = blogs.id', **'left'**);  
    // 生成: LEFT JOIN comments ON comments.id = blogs.id

## $this->db->where();

本函数允许你使用四种方法中的一种来设置 **WHERE** 子句:
  **说明:** 传递给本函数的所有值都会被自动转义，以便生成安全的查询。
1.  **简单的 key/value 方法:** 
        $this->db->where('name', $name);
             // 生成: WHERE name = 'Joe'
    请注意等号已经为你添加。

    如果你多次调用本函数，那么这些条件会被 <var>AND</var> 连接起来:

        $this->db->where('name', $name);  
         $this->db->where('title', $title);  
         $this->db->where('status', $status);
             // WHERE name = 'Joe' AND title = 'boss' AND status = 'active'
2.  **自定义 key/value 方法:** 
    你可以在第一个参数中包含一个运算符，以便控制比较:

        $this->db->where('name !=', $name);  
         $this->db->where('id <', $id);
             // 生成: WHERE name != 'Joe' AND id < 45

3.  **关联数组方法:** 

         $array = array('name' => $name, 'title' => $title, 'status' => $status);    

         $this->db->where($array);
             // 生成: WHERE name = 'Joe' AND title = 'boss' AND status = 'active'
    使用这个方法时你也可以包含运算符:

         $array = array('name !=' => $name, 'id <' => $id, 'date >' => $date);    

         $this->db->where($array);

4.  **自定义字符串:** 
    你可以手动的编写子句:
          $where = "name='Joe' AND status='boss' OR status='active'";    
          $this->db->where($where);

$this->db->where() 接受可选的第三个参数。如果你将它设置为 FALSE, CodeIgniter 将不会为你那些包含反勾号的字段名或表名提供保护。

       $this->db->where('MATCH (field) AGAINST ("value")', NULL, FALSE);  
## $this->db->or_where();

本函数与上面的那个几乎完全相同，唯一的区别是本函数生成的子句是用 OR 来连接的:
    $this->db->where('name !=', $name);  
    $this->db->or_where('id >', $id);

      // 生成: WHERE name != 'Joe' OR id > 50
  说明: or_where() 以前被叫作 orwhere(), 后者已经过时，现已从代码中移除 orwhere()。
## $this->db->where_in();

生成一段 WHERE field IN ('item', 'item') 查询语句，如果合适的话，用 AND 连接起来。
        $names = array('Frank', 'Todd', 'James');  
        $this->db->where_in('username', $names);  
        // 生成: WHERE username IN ('Frank', 'Todd', 'James')

## $this->db->or\_where\_in();

生成一段 WHERE field IN ('item', 'item') 查询语句，如果合适的话，用 OR 连接起来。
        $names = array('Frank', 'Todd', 'James');  

        $this->db->or_where_in('username', $names);  
        // 生成: OR username IN ('Frank', 'Todd', 'James')

## $this->db->where\_not\_in();

生成一段 WHERE field NOT IN ('item', 'item') 查询语句，如果合适的话，用 AND 连接起来。
        $names = array('Frank', 'Todd', 'James');  
        $this->db->where_not_in('username', $names);  

        // 生成: WHERE username NOT IN ('Frank', 'Todd', 'James')

## $this->db->or\_where\_not_in();

生成一段 WHERE field NOT IN ('item', 'item') 查询语句，如果合适的话，用 OR 连接起来。
        $names = array('Frank', 'Todd', 'James');  
        $this->db->or_where_not_in('username', $names);  
        // 生成: OR username NOT IN ('Frank', 'Todd', 'James')

## $this->db->like();

本函数允许你生成 **LIKE** 子句，在做查询时非常有用。
  **说明:** 传递给本函数的所有值都会被自动转义。
1.  **简单 key/value 方式:** 
        $this->db->like('title', 'match');
             // 生成: WHERE title LIKE '%match%' 

    如果你多次调用本函数，那么这些条件将由 <var>AND</var> 连接起来:

        $this->db->like('title', 'match');  
         $this->db->like('body', 'match');
         // WHERE title LIKE '%match%' AND  body LIKE '%match%' 如果你想要控制通配符(%)所出现的位置，你可以使用可选的第三个参数。可用的选项是 'before', 'after' 以及 'both' (这是默认值)。 

        $this->db->like('title', 'match', 'before');

          // 生成: WHERE title LIKE '%match'   

         $this->db->like('title', 'match', 'after');   

        // 生成: WHERE title LIKE 'match%'   

         $this->db->like('title', 'match', 'both');   
        // 生成: WHERE title LIKE '%match%'  如果你不想使用百分号(%)，你可以给第三个可选的参数传递一个'none'。 
      $this->db->like('title', 'match', 'none');   
    // Produces: WHERE title LIKE 'match'
2.  **关联数组方式:** 
         $array = array('title' => $match, 'page1' => $match, 'page2' => $match);    

         $this->db->like($array);
             // WHERE title LIKE '%match%' AND page1 LIKE '%match%' AND page2 LIKE '%match%'

## $this->db->or_like();

本函数与上面那个函数几乎完全相同，唯一的区别是多个实例之间是用 OR 连接起来的:
    $this->db->like('title', 'match');  

    $this->db->or_like('body', $match);

      // WHERE title LIKE '%match%' OR body LIKE '%match%'
  说明: or_like() 曾经被称为 orlike(), 后者已经过时，现已从代码中移除 orlike()。
## $this->db->not_like();

本函数与 **like()** 几乎完全相同，唯一的区别是它生成 NOT LIKE 语句:

     $this->db->not_like('title', 'match');  

    // WHERE title NOT LIKE '%match%

## $this->db->or\_not\_like();

本函数与 **not_like()** 几乎完全相同，唯一的区别是多个实例之间是用 OR 连接起来的:

     $this->db->like('title', 'match');  

    $this->db->or_not_like('body', 'match');   

    // WHERE title  LIKE '%match%' OR body NOT LIKE '%match%'

## $this->db->group_by();

允许你编写查询语句中的 GROUP BY 部分:

    $this->db->group_by("title");
        // 生成: GROUP BY title
你也可以把多个值作为数组传递过去:

    $this->db->group_by(array("title", "date"));

      // 生成: GROUP BY title, date
  说明: group_by() 曾经被称为 groupby(), 后者已经过时，现已从代码中移除 groupby()。
## $this->db->distinct();  
为查询语句添加 "DISTINCT" 关键字:

    $this->db->distinct();  
        $this->db->get('table');  

        // 生成: SELECT DISTINCT * FROM table

## $this->db->having();

允许你为你的查询语句编写 HAVING 部分。有两种语法形式，一个或两个参数都可以:

    $this->db->having('user_id = 45');

    // 生成: HAVING user_id = 45  

    $this->db->having('user_id',  45);   
    // 生成: HAVING user_id = 45  

你也可以把多个值通过数组传递过去:

    $this->db->having(array('title =' => 'My Title', 'id <' => $id));   

     // 生成: HAVING title = 'My Title' AND id < 45

如果你正在使用一个由CodeIgniter进行转义保护的数据库，为了避免内容转义，你可以传递可选的第三个参数，并将其设置为FALSE。

    $this->db->having('user_id',  45);   
    // 生成: HAVING `user_id` = 45 (在诸如MySQL等数据库中)

      $this->db->having('user_id',  45, FALSE);   
    // 生成: HAVING user_id = 45

## $this->db->or_having();

与 having() 函数几乎完全一样，唯一的区别是多个子句之间是用 "OR" 分隔的。

## $this->db->order_by();

帮助你设置一个 ORDER BY 子句。第一个参数是你想要排序的字段名。第二个参数设置结果的顺序，可用的选项包括 <kbd>asc</kbd> (升序)或 <kbd>desc</kbd>(降序), 或 <kbd>random</kbd>(随机)。

    $this->db->order_by("title", "desc");

      // 生成: ORDER BY title DESC
你也可以在第一个参数中传递你自己的字符串:

    $this->db->order_by('title desc, name asc');

      // 生成: ORDER BY title DESC, name ASC
或者，多次调用本函数就可以排序多个字段。

    $this->db->order_by("title", "desc");  
        $this->db->order_by("name", "asc");   

        // 生成: ORDER BY title DESC, name ASC

  说明: order_by() 曾经被称为 orderby(), 后者已经过时，现已从代码中移除 orderby()。

  说明: 目前 Oracle 和 MSSQL 的驱动还不支持随机排序，将被默认设置为 'ASC'(升序)。
## $this->db->limit();

限制查询所返回的结果数量:
    $this->db->limit(10);  

    // 生成: LIMIT 10

第二个参数设置的是结果偏移量。
    $this->db->limit(10, 20);  

    // 生成: LIMIT 20, 10 (仅限MySQL中。其它数据库有稍微不同的语法)

## $this->db->count\_all\_results();

允许你获得某个特定的Active Record查询所返回的结果数量。可以使用Active Record限制函数，例如 where(), or\_where(), like(), or\_like() 等等。范例：

    echo $this->db->count_all_results('<var>my_table</var>');  

    // 生成一个整数，例如 25  

    $this->db->like('title', 'match');  

    $this->db->from('<var>my_table</var>');  
    echo $this->db->count_all_results();  
    // 生成一个整数，例如 17 

<a name="insert"> </a> 
# 插入数据

## $this->db->insert();

生成一条基于你所提供的数据的SQL插入字符串并执行查询。你可以向函数传递 **数组** 或一个 **对象**。下面是一个使用数组的例子：
    $data = array(  
                   'title' => 'My title' ,  
                   'name' => 'My Name' ,  
                   'date' => 'My date'  

                );  

    $this->db->insert('mytable', $data);

    // Produces: INSERT INTO mytable (title, name, date) VALUES ('My title', 'My name', 'My date')

第一个参数包含表名，第二个是一个包含数据的关联数组。

下面是一个使用对象的例子：
    /*  

        class Myclass {  
            var  $title = 'My Title';  
            var  $content = 'My Content';  
            var  $date = 'My Date';  
        }  
    */  

    $object = new Myclass;  

    $this->db->insert('mytable', $object);
    // Produces: INSERT INTO mytable (title, content, date) VALUES ('My Title', 'My Content', 'My Date')

第一个参数包含表名，第二个是一个对象。
  **注意:** 所有的值已经被自动转换为安全查询。
## $this->db->insert_batch();

生成一条基于你所提供的数据的SQL插入字符串并执行查询。你可以向函数传递 **数组** 或一个 **对象**。下面是一个使用数组的例子：
    $data = array(  
       array(  
          'title' => 'My title' ,  
          'name' => 'My Name' ,  
          'date' => 'My date'  
       ),  
       array(  
          'title' => 'Another title' ,  
          'name' => 'Another Name' ,  
          'date' => 'Another date'  
       )  
    );  

    $this->db->insert_batch('mytable', $data);

    //生成： INSERT INTO mytable (title, name, date) VALUES ('My title', 'My name', 'My date'),  ('Another title', 'Another name', 'Another date')

第一个参数包含表名，第二个是一个包含数据的关联数组。
  **注意:** 所有的值已经被自动转换为安全查询。
## $this->db->set();

本函数使您能够设置_inserts_(插入)或_updates_(更新)值。

**它可以用来代替那种直接传递数组给插入和更新函数的方式:** 

    $this->db->set('name', $name);

    $this->db->insert('mytable');

    // 生成: INSERT INTO mytable (name) VALUES ('{$name}')

如果你多次调用本函数，它们会被合理地组织起来，这取决于你执行的是插入操作还是更新操作:

    $this->db->set('name', $name);  

    $this->db->set('title', $title);  
    $this->db->set('status', $status);  
    $this->db->insert('mytable'); 

**set()** 也接受可选的第三个参数($escape)，如果此参数被设置为 FALSE，就可以阻止数据被转义。为了说明这种差异，这里将对 包含转义参数 和 不包含转义参数 这两种情况的 set() 函数做一个说明。

    $this->db->set('field', 'field+1', FALSE);  

        $this->db->insert('mytable');   
        // 得到 INSERT INTO mytable (field) VALUES (field+1)  

        $this->db->set('field', 'field+1');  
        $this->db->insert('mytable');   

        // 得到 INSERT INTO mytable (field) VALUES ('field+1')

你也可以将一个关联数组传递给本函数:
    $array = array('name' => $name, 'title' => $title, 'status' => $status);    

    $this->db->set($array);  
    $this->db->insert('mytable');
或者一个对象也可以:
    /*  
        class Myclass {  
            var  $title = 'My Title';  
            var  $content = 'My Content';  
            var  $date = 'My Date';  
        }  
    */  
    $object = new Myclass;  

    $this->db->set($object);  
    $this->db->insert('mytable');
<a name="update"> </a> 
# 更新数据

## $this->db->update();

根据你提供的数据生成并执行一条update(更新)语句。你可以将一个**数组**或者**对象**传递给本函数。这里是一个使用数组的例子:
    $data = array(  
                   'title' => $title,  
                   'name' => $name,  
                   'date' => $date  

                );  

    $this->db->where('id', $id);  
    $this->db->update('mytable', $data);

    // 生成:  
    // UPDATE mytable   
    // SET title = '{$title}', name = '{$name}', date = '{$date}'  
    // WHERE id = $id

或者你也可以传递一个对象:
    /*  
        class Myclass {  
            var  $title = 'My Title';  
            var  $content = 'My Content';  
            var  $date = 'My Date';  
        }  
    */  
    $object = new Myclass;  

    $this->db->where('id', $id);  
    $this->db->update('mytable', $object);
    // 生成:  
    // UPDATE mytable   
    // SET title = '{$title}', name = '{$name}', date = '{$date}'  
    // WHERE id = $id
  **说明:** 所有值都会被自动转义，以便生成安全的查询。
你会注意到 _$this->db->where()_ 函数的用法，它允许你设置 WHERE 子句。你可以有选择性地将这一信息直接以字符串的形式传递给 update 函数:

    $this->db->update('mytable', $data, "id = 4");

或者是一个数组:

    $this->db->update('mytable', $data, array('id' => $id));

在进行更新时，你还可以使用上面所描述的 _$this->db->set()_ 函数。

## $this->db->update_batch();

生成一条update命令是以你提供的数据为基础的，并执行查询。你可以传递一个数组或对象的参数给update_batch()函数。下面是一个使用一个数组作为参数的示例：Generates an update string based on the data you supply, and runs the query. You can either pass an **array** or an **object** to the function. Here is an example using an array:
    $data = array(  
       array(  
          'title' => 'My title' ,  
          'name' => 'My Name 2' ,  
          'date' => 'My date 2'  
       ),  
       array(  
          'title' => 'Another title' ,  
          'name' => 'Another Name 2' ,  
          'date' => 'Another date 2'  
       )  
    );  

    $this->db->update_batch('mytable', $data, 'title');

    // Produces:   
    // UPDATE `mytable` SET `name` = CASE  
    // WHEN `title` = 'My title' THEN 'My Name 2'  
    // WHEN `title` = 'Another title' THEN 'Another Name 2'  
    // ELSE `name` END,  
    // `date` = CASE   
    // WHEN `title` = 'My title' THEN 'My date 2'  
    // WHEN `title` = 'Another title' THEN 'Another date 2'  
    // ELSE `date` END  
    // WHERE `title` IN ('My title','Another title')

参数1:表名 参数2:如上所示的二维数组 参数3:键名.
  **提示:** 所有的值都会自动进行安全性过滤.
<a name="delete"> </a> 
# 删除数据

## $this->db->delete();

生成并执行一条DELETE(删除)语句。
    $this->db->delete('mytable', array('id' => $id));

    // 生成:  
    // DELETE FROM mytable   
    // WHERE id = $id

第一个参数是表名，第二个参数是where子句。你可以不传递第二个参数，使用 _where()_ 或者 _or_where()_ 函数来替代它:

     $this->db->where('id', $id);  
        $this->db->delete('mytable');   

        // 生成:  
        // DELETE FROM mytable   
        // WHERE id = $id

如果你想要从一个以上的表中删除数据，你可以将一个包含了多个表名的数组传递给delete()函数。

    $tables = array('table1', 'table2', 'table3');  
    $this->db->where('id', '5');  
    $this->db->delete($tables);

如果你想要删除表中的全部数据，你可以使用 _truncate()_ 函数，或者 _empty_table()_ 函数。

## $this->db->empty_table();

生成并执行一条DELETE(删除)语句。 
        $this->db->empty_table('mytable');   

    // 生成  
    // DELETE FROM mytable

## $this->db->truncate();

生成并执行一条TRUNCATE(截断)语句。

     $this->db->from('mytable');   
    $this->db->truncate();   
    // 或   
    $this->db->truncate('mytable');   
    // 生成:  
    // TRUNCATE TABLE mytable   

  **说明:** 如果 TRUNCATE 命令不可用，truncate() 将会以 "DELETE FROM table" 的方式执行。
# delete方法貌似现在没有办法接收排序参数

假设我想删除早期的用户登录日志信息，我的语法可能会是根据登录时间正序排列，然后跟上limit已经有的总日志条数减去想保留的条数，delete方法没有这个参数用来接收，那么我只能直接使用query传递我的sql了

# <a name="chaining"> </a>链式方法

链式方法允许你以连接多个函数的方式简化你的语法。考虑一下这个范例:
    _$this->db_<kbd>-></kbd><var>select</var>('title')<kbd>-></kbd><var>from</var>('mytable')<kbd>-></kbd><var>where</var>('id', $id)<kbd>-></kbd><var>limit</var>(10, 20);  
    $query = $this->db->get();
  **说明:** 链式方法只能在PHP 5下面运行。
 

# <a name="caching"> </a>Active Record 缓存

尽管不是 "真正的" 缓存，Active Record 允许你将查询的某个特定部分保存(或"缓存")起来，以便在你的脚本执行之后重用。一般情况下，当一次Active Record调用结束，所有已存储的信息都会被重置，以便下一次调用。如果开启缓存，你就可以使信息避免被重置，方便你进行重用。

缓存调用是累加的。如果你调用了两次有缓存的 select()，然后再调用两次没有缓存的 select()，这会导致 select() 被调用4次。有三个可用的缓存函数:

## $this->db->start_cache()

本函数必须被用来开启缓存。所有类型正确的(下面给出了支持的查询类型) Active Record 查询都会被存储起来供以后使用。

## $this->db->stop_cache()

本函数可以被用来停止缓存。

## $this->db->flush_cache()

本函数从Active Record 缓存中删除全部项目。

这里是一个使用范例:

    $this->db->start_cache();  
    $this->db->select('field1');  
    $this->db->stop_cache();    
    $this->db->get('tablename');  

    //Generates: SELECT `field1` FROM (`tablename`)  

    $this->db->select('field2');  
    $this->db->get('tablename');  

    //Generates:  SELECT `field1`, `field2` FROM (`tablename`)  

    $this->db->flush_cache();  

    $this->db->select('field2');  
    $this->db->get('tablename');  

    //Generates:  SELECT `field2` FROM (`tablename`)
  **说明:** 下列语句能够被缓存: select, from, join, where, like, group_by, having, order_by, set
 
  翻译贡献者: atans, aykirk, Drice, Fanbin, Hex, houlianshan, huanxiangwu, imjie, k7gxn56, loiynet, longjianghu, ripen, thankwsx, yinzhili, zhangminghua218

  最后修改: 2012-10-25 16:29:28
 [1]: #select
 [2]: #insert
 [3]: #update
 [4]: #delete
 [5]: #chaining
 [6]: #caching
 [7]: results.html
