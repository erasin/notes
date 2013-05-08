
## array()

创建数组。	3

__定义和用法__

array() 创建数组，带有键和值。如果在规定数组时省略了键，则生成一个整数键，这个 key 从 0 开始，然后以 1 进行递增。

要用 array() 创建一个关联数组，可使用 `=>` 来分隔键和值。

要创建一个空数组，则不传递参数给 array()：

    $new = array();

注意：array() 实际上是一种语言结构 (language construct)，通常用来定义直接量数组，但它的用法和函数的用法很相似，所以我们把它也列到手册中。

__语法__

    array(key => value) 

参数    | 描述
--------|----------
key     | 可选。规定 key，类型是数值或字符串。如果未设置，则生成整数类型的 key。
value   | 必需。规定值。


## array_change_key_case()

返回其键均为大写或小写的数组。	4

__定义和用法__

`array_change_key_case()` 函数将数组的所有的 KEY 都转换为大写或小写。

数组的数字索引不发生变化。如果未提供可选参数（即第二个参数），则__默认转换为小写字母__。

__语法__

    array_change_key_case(array,case)

参数    | 描述
--------|--------
array   | 必需。规定要使用的数组。
case    | 可选。可能的值：
        | CASE_LOWER - 默认值。  以小写字母返回数组的键。
        | CASE_UPPER - 以大写字母返回数组的键。

> 注释：如果在运行该函数时两个或多个键相同，则最后的元素会覆盖其他元素（参见例子 1）。

例子 1

    <?php
    $a=array("a"=>"Cat","b"=>"Dog","c"=>"Horse","B"=>"Bird");
    print_r(array_change_key_case($a,CASE_UPPER));
    ?> 

输出：

    Array ( [A] => Cat [B] => Bird [C] => Horse )

## array_chunk()

把一个数组分割为新的数组块。	4

定义和用法
array_chunk() 函数把一个数组分割为新的数组块。
其中每个数组的单元数目由 size 参数决定。最后一个数组的单元数目可能会少几个。
可选参数 preserve_key 是一个布尔值，它指定新数组的元素是否有和原数组相同的键（用于关联数组），还是从 0 开始的新数字键（用于索引数组）。默认是分配新的键。
语法
array_chunk(array,size,preserve_key)
参数	描述
array	必需。规定要使用的数组。
size	必需。规定每个新数组包含多少个元素。
preserve_key	
可选。可能的值：
true - 保留原始数组中的键名。
false - 默认。每个结果数组使用从零开始的新数组索引。

## array_combine()

通过合并两个数组来创建一个新数组。	5

定义和用法
array_combine() 函数通过合并两个数组来创建一个新数组，其中的一个数组是键名，另一个数组的值为键值。
如果其中一个数组为空，或者两个数组的元素个数不同，则该函数返回 false。
语法
array_combine(array1,array2)
参数	描述
array1	必需。规定键名。
array2	必需。规定值。
提示和注释
注释：两个参数必须有相同数目的元素。

## array_count_values()

用于统计数组中所有值出现的次数。	4
定义和用法
array_count_values() 函数用于统计数组中所有值出现的次数。
本函数返回一个数组，其元素的键名是原数组的值，键值是该值在原数组中出现的次数。
语法
array_count_values(array)
参数	描述
array	必需。规定输入的数组。
例子
<?php
$a=array("Cat","Dog","Horse","Dog");
print_r(array_count_values($a));
?>
输出：
Array ( [Cat] => 1 [Dog] => 2 [Horse] => 1 )


## array_diff()

返回两个数组的差集数组。	4

定义和用法
array_diff() 函数返回两个数组的差集数组。该数组包括了所有在被比较的数组中，但是不在任何其他参数数组中的键值。
在返回的数组中，键名保持不变。
语法
array_diff(array1,array2,array3...)
参数	描述
array1	必需。与其他数组进行比较的第一个数组。
array2	必需。与第一个数组进行比较的数组。
array3	可选。与第一个数组进行比较的数组。
提示和注释
提示：可用一个或任意多个数组与第一个数组进行比较。
注释：仅有值用于比较。
例子
<?php
$a1=array(0=>"Cat",1=>"Dog",2=>"Horse");
$a2=array(3=>"Horse",4=>"Dog",5=>"Fish");
print_r(array_diff($a1,$a2));
?>
输出：
Array ( [0] => Cat )

## array_diff_assoc()

比较键名和键值，并返回两个数组的差集数组。	4
定义和用法
array_diff_assoc() 函数返回两个数组的差集数组。该数组包括了所有在被比较的数组中，但是不在任何其他参数数组中的键和值。
和 array_diff() 函数 不同，本函数要求键名和键值都进行比较。返回的数组中键名保持不变。
语法
array_diff_assoc(array1,array2,array3...)
参数	描述
array1	必需。与其他数组进行比较的第一个数组。
array2	必需。与第一个数组进行比较的数组。
array3	可选。与第一个数组进行比较的数组。可以有多个。
提示和注释
提示：可用一个或任意多个数组与第一个数组进行比较。
注释：键和值都用于比较。
例子
<?php
$a1=array(0=>"Cat",1=>"Dog";,2=>"Horse");
$a2=array(0=>"Rat",1=>"Horse";,2=>"Dog");
$a3=array(0=>"Horse",1=>"Dog",2=>"Cat");
print_r(array_diff_assoc($a1,$a2,$a3));
?>
输出：
Array ( [0] => Cat [2] => Horse )


## array_diff_key()

比较键名，并返回两个数组的差集数组。	5
定义和用法
array_diff_key() 函数返回一个数组，该数组包括了所有在被比较的数组中，但是不在任何其他参数数组中的键。
语法
array_diff_key(array1,array2,array3...)
参数	描述
array1	必需。与其他数组进行比较的第一个数组。
array2	必需。与第一个数组进行比较的数组。
array3	可选。与第一个数组进行比较的数组。可以有多个。
提示和注释
提示：可用一个或任意多个数组与第一个数组进行比较。
注释：仅仅键名用于比较。
例子
<?php
$a1=array(0=>"Cat",1=>"Dog",2=>"Horse");
$a2=array(2=>"Bird",3=>"Rat",4=>"Fish");
$a3=array(5=>"Horse",6=>"Dog",7=>"Bird");
print_r(array_diff_key($a1,$a2,$a3));
?>
输出：
Array ( [0] => Cat [1] => Dog )
## array_diff_uassoc()

通过用户提供的回调函数做索引检查来计算数组的差集。	5
定义和用法
array_diff_uassoc() 函数使用用户自定义的回调函数 (callback) 做索引检查来计算两个或多个数组的差集。返回一个数组，该数组包括了在 array1 中但是不在任何其他参数数组中的值。
注意，与 array_diff() 函数 不同的是，键名也要进行比较。
参数 function 是用户自定义的用来比较两个数组的函数，该函数必须带有两个参数 - 即两个要进行对比的键名。因此与函数 array_diff_assoc() 的行为正好相反，后者是用内部函数进行比较的。
返回的数组中键名保持不变。
语法
array_diff_uassoc(array1,array2,array3...,function)
参数	描述
array1	必需。与其他数组进行比较的第一个数组。
array2	必需。与第一个数组进行比较的数组。
array3	可选。与第一个数组进行比较的数组。可以有多个。
function	必需。用户自定义函数的名称。
例子 1
<?php
function myfunction($v1,$v2) 
{
if ($v1===$v2)
	{
	return 0;
	}
if ($v1>$v2)
	{
	return 1;
	}
else
	{
	return -1;
	}
}
$a1=array(0=>"Dog",1=>"Cat",2=>"Horse");
$a2=array(3=>"Dog",1=>"Cat",5=>"Horse");
print_r(array_diff_uassoc($a1,$a2,"myfunction"));
?>
输出：
Array ( [0] => Dog [2] => Horse )
例子 2
如何为该函数分配多个数组：
<?php
function myfunction($v1,$v2) 
{
if ($v1===$v2)
	{
	return 0;
	}
if ($v1>$v2)
	{
	return 1;
	}
else
	{
	return -1;
	}
}
$a1=array(0=>"Dog",1=>"Cat",2=>"Horse");
$a2=array(3=>"Dog",1=>"Cat",5=>"Horse");
$a3=array(6=>"Bird",0=>"Dog",5=>"Horse");
print_r(array_diff_uassoc($a1,$a2,$a3,"myfunction"));
?>
输出：
Array ( [2] => Horse )
## array_diff_ukey()

用回调函数对键名比较计算数组的差集。	5

## array_fill()

用给定的值填充数组。	4

## array_filter()

用回调函数过滤数组中的元素。	4

## array_flip()

交换数组中的键和值。	4

## array_intersect()

计算数组的交集。	4

## array_intersect_assoc()

比较键名和键值，并返回两个数组的交集数组。	4

## array_intersect_key()

使用键名比较计算数组的交集。	5

## array_intersect_uassoc()

带索引检查计算数组的交集，用回调函数比较索引。	5

## array_intersect_ukey()

用回调函数比较键名来计算数组的交集。	5

## array_key_exists()

检查给定的键名或索引是否存在于数组中。	4

## array_keys()

返回数组中所有的键名。	4

## array_map()

将回调函数作用到给定数组的单元上。	4

## array_merge()

把一个或多个数组合并为一个数组。	4

## array_merge_recursive()

递归地合并一个或多个数组。	4

## array_multisort()

对多个数组或多维数组进行排序。	4

## array_pad()

用值将数组填补到指定长度。	4

## array_pop()

将数组最后一个单元弹出（出栈）。	4

## array_product()

计算数组中所有值的乘积。	5

## array_push()

将一个或多个单元（元素）压入数组的末尾（入栈）。	4

## array_rand()

从数组中随机选出一个或多个元素，并返回。	4

## array_reduce()

用回调函数迭代地将数组简化为单一的值。	4

## array_reverse()

将原数组中的元素顺序翻转，创建新的数组并返回。	4

## array_search()

在数组中搜索给定的值，如果成功则返回相应的键名。	4

## array_shift()

删除数组中的第一个元素，并返回被删除元素的值。	4

## array_slice()

在数组中根据条件取出一段值，并返回。	4

## array_splice()

把数组中的一部分去掉并用其它值取代。	4

## array_sum()

计算数组中所有值的和。	4

## array_udiff()

用回调函数比较数据来计算数组的差集。	5

## array_udiff_assoc()

带索引检查计算数组的差集，用回调函数比较数据。	5

## array_udiff_uassoc()

带索引检查计算数组的差集，用回调函数比较数据和索引。	5

## array_uintersect()

计算数组的交集，用回调函数比较数据。	5

## array_uintersect_assoc()

带索引检查计算数组的交集，用回调函数比较数据。	5

## array_uintersect_uassoc()

带索引检查计算数组的交集，用回调函数比较数据和索引。	5

## array_unique()

删除数组中重复的值。	4

## array_unshift()

在数组开头插入一个或多个元素。	4

## array_values()

返回数组中所有的值。	4

## array_walk()

对数组中的每个成员应用用户函数。	3

## array_walk_recursive()

对数组中的每个成员递归地应用用户函数。	5

## arsort()

对数组进行逆向排序并保持索引关系。	3

## asort()

对数组进行排序并保持索引关系。	3

## compact()

建立一个数组，包括变量名和它们的值。	4

## count()

计算数组中的元素数目或对象中的属性个数。	3

## current()

返回数组中的当前元素。	3

## each()

返回数组中当前的键／值对并将数组指针向前移动一步。	3

## end()

将数组的内部指针指向最后一个元素。	3

## extract()

从数组中将变量导入到当前的符号表。	3

## in_array()

检查数组中是否存在指定的值。	4

## key()

从关联数组中取得键名。	3

## krsort()

对数组按照键名逆向排序。	3

## ksort()

对数组按照键名排序。	3

## list()

把数组中的值赋给一些变量。	3

## natcasesort()

用“自然排序”算法对数组进行不区分大小写字母的排序。	4

## natsort()

用“自然排序”算法对数组排序。	4

## next()

将数组中的内部指针向前移动一位。	3

## pos()

current() 的别名。	3

## prev()

将数组的内部指针倒回一位。	3

## range()

建立一个包含指定范围的元素的数组。	3

## reset()

将数组的内部指针指向第一个元素。	3

## rsort()

对数组逆向排序。	3

## shuffle()

把数组中的元素按随机顺序重新排列。	3

## sizeof()

count() 的别名。	3

## sort()

对数组排序。	3

## uasort()

使用用户自定义的比较函数对数组中的值进行排序并保持索引关联。	3

## uksort()

使用用户自定义的比较函数对数组中的键名进行排序。	3

## usort()

使用用户自定义的比较函数对数组中的值进行排序。	3
