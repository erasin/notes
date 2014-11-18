# 模板

- [控制器布局](#controller-layouts)
- [Blade模板](#blade-templating)
- [Blade模板控制结构](#other-blade-control-structures)
- [扩展Blade](#extending-blade)

<a name="controller-layouts"></a>
## 控制器布局


在Laravel框架中使用模板的一种方法就是通过控制器布局。通过在控制器中指定`layout`属性，指定的视图就会被创建，并作为默认数据，在actions中返回。

#### 在控制器中定义布局（Layouts）

	class UserController extends BaseController {

		/**
		 * The layout that should be used for responses.
		 */
		protected $layout = 'layouts.master';

		/**
		 * Show the user profile.
		 */
		public function showProfile()
		{
			$this->layout->content = View::make('user.profile');
		}

	}

<a name="blade-templating"></a>
## Blade模板

Blade是Laravel框架下的一个简单但又强大的模板引擎。 不同于控制器布局，Blade模板引擎由 _模板继承_ 和 _模板片段_ 驱动。所有的Blade模板文件必须使用 `.blade.php` 文件扩展名。

#### 定义一个Blade布局

	<!-- Stored in app/views/layouts/master.blade.php -->

	<html>
		<body>
			@section('sidebar')
				This is the master sidebar.
			@show

			<div class="container">
				@yield('content')
			</div>
		</body>
	</html>

#### 使用一个Blade布局

	@extends('layouts.master')

	@section('sidebar')
		@parent

		<p>This is appended to the master sidebar.</p>
	@stop

	@section('content')
		<p>This is my body content.</p>
	@stop

注意视图中片段只是简单的替换其`extend`的Blade布局中相应片段。通过在模板片段中使用`@parent`指令，布局的内容可以包含一个子视图，这样你就可以在布局片段中添加诸如侧边栏、底部信息等内容。

有时候，有些片段可能不能确定被定义了，你可以使用`@yield`结构给出一个默认值。如下，第二个值即是默认值。

	@yield('section', 'Default Content');

<a name="other-blade-control-structures"></a>
## 其他 Blade模板 控制结构

#### 输出数据

	Hello, {{{ $name }}}.

	The current UNIX timestamp is {{{ time() }}}.

#### 检测是否存在后输出数据

有时，你可能希望输出一个变量，但又不能确定这个变量是否被设置。直接点，你可能想这么做：

	{{{ isset($name) ? $name : 'Default' }}}

然而，除了写一个三目运算符，Blade有如下简写方法：

	{{{ $name or 'Default' }}}

#### 显示带有大括号的文本

如果你想显示带有大括号的字符串，你可以在文本前放`@`符号，这样会忽略Blade解析行为

	@{{ This will not be processed by Blade }}

当然，用户提供的全部字符串都应该都被转义（主要对html标签，利用htmlentities编码转义）。如果想转义输出，你可以使用三个大括号语法：

	Hello, {{{ $name }}}.

如果不希望数据被转义，可以使用双大括号语法：

	Hello, {{ $name }}.

> **注意：** 一定要小心输出的用户提供的内容。使用三个大括号的语法能够直接输出内容中的HTML标签。

#### If标签

	@if (count($records) === 1)
		I have one record!
	@elseif (count($records) > 1)
		I have multiple records!
	@else
		I don't have any records!
	@endif

	@unless (Auth::check())
		You are not signed in.
	@endunless

#### 循环

	@for ($i = 0; $i < 10; $i++)
		The current value is {{ $i }}
	@endfor

	@foreach ($users as $user)
		<p>This is user {{ $user->id }}</p>
	@endforeach

	@while (true)
		<p>I'm looping forever.</p>
	@endwhile

#### 包含子视图

	@include('view.name')

你也可以传递数组数据到被包含的视图

	@include('view.name', array('some'=>'data'))

#### 覆盖片段

默认的，片段是附加在先前存在的片段上，如果想覆盖一个片段的全部，可以使用`overwrite`标签

	@extends('list.item.container')

	@section('list.item.content')
		<p>This is an item of type {{ $item->type }}</p>
	@overwrite

#### 输出多语言

	@lang('language.line')

	@choice('language.line', 1);

#### 注释

	{{-- This comment will not be in the rendered HTML --}}

<a name="extending-blade"></a>
## 扩展Blade

Blade允许用户定义自己的控制结构。当一个Blade文件被编译后，会调用用户自定义的扩展，用来处理视图内容，从简单的`str_replace`操作，到很复杂的表达式，总之，你可以做任何事情。

Blade的编译器附带了帮助函数`createMatcher`和`createPlainMatcher`，这两个函数可以生成自定义指令。

`createPlainMatcher`函数主要用于没有参数传递的指令，类似`@endif`和`@stop`，而`createMatcher`则用于那些有参数传递的指令。

下面的例子创建`@datetime($var)`指令，它只是简单的对`$var`调用`->format()`方法：

	Blade::extend(function($view, $compiler)
	{
		$pattern = $compiler->createMatcher('datetime');

		return preg_replace($pattern, '$1<?php echo $2->format('m/d/Y H:i'); ?>', $view);
	});
