# 升级指南

- [从4.1升级到4.2](#upgrade-4.2)
- [从小于等于4.1.x升级到4.1.29](#upgrade-4.1.29)
- [从小于等于4.1.25升级到4.1.26](#upgrade-4.1.26)
- [从4.0升级到4.1](#upgrade-4.1)

<a name="upgrade-4.2"></a>
## 从4.1升级到4.2

### PHP 5.4+

Laravel 4.2需要PHP 5.4.0或更高版本。

### 加密的默认设置

在配置文件“app/config/app.php”里添加一个新的选项“cipher”。 此选项的值应该是“MCRYPT_RIJNDAEL_256”。

	'cipher' => MCRYPT_RIJNDAEL_256

这个设置可以用来控制Laravel加密功能使用默认密钥。

> **Note:** In Laravel 4.2, the default cipher is `MCRYPT_RIJNDAEL_128` (AES), which is considered to be the most secure cipher. Changing the cipher back to `MCRYPT_RIJNDAEL_256` is required to decrypt cookies/values that were encrypted in Laravel <= 4.1

### 软删除模型现在的使用特性

如果你正在使用软删除模型，“softDeletes”属性已经被删除了。现在你应该使用“SoftDeletingTrait”属性，像下面这样：

	use Illuminate\Database\Eloquent\SoftDeletingTrait;

	class User extends Eloquent {
		use SoftDeletingTrait;
	}

你还应该手动添加“deleted_at”列到“dates”属性上：

	class User extends Eloquent {
		use SoftDeletingTrait;

		protected $dates = ['deleted_at'];
	}

所有软删除操作的API保持不变。

> **Note:** The `SoftDeletingTrait` can not be applied on a base model. It must be used on an actual model class.

### 重命名的视图/分页环境

如果你直接的引用了“Illuminate\View\Environment”类或“Illuminate\Pagination\Environment”类，升级你的代码，用“Illuminate\View\Factory”和“Illuminate\Pagination\Factory”来替代。原来的两个类已经被重命名,以更好地反映其功能。

### 分页呈现时的额外参数

如果你扩展了“Illuminate\Pagination\Presenter”类，它的抽象方法“getPageLinkWrapper”的签名已经被改变了，添加了“ref”参数：

	abstract public function getPageLinkWrapper($url, $page, $rel = null);

### Iron.Io Queue Encryption

If you are using the Iron.io queue driver, you will need to add a new `encrypt` option to your queue configuration file:

    'encrypt' => true

<a name="upgrade-4.1.29"></a>
## 从小于等于4.1.x升级到4.1.29 

Laravel 4.1.29改善了所有的数据库驱动程序的列引用。当模型“不”使用“fillable”属性时，它保护您的应用程序免受某个mass assignment漏洞。 如果你在模型上使用了“fillable”属性来防止mass assignemnt漏洞，你的应用程序就是不易受攻击的。然而，如果你正在使用“guarded”，同时传递了一个用户控制的数组给“update”或“save”类型的函数，你应该立即升级到“4.1.29”，因为你的应用程序可能面临“mass assignment”漏洞的风险。

要升级到Laravel 4.1.29，简单的执行“composer update”命令即可。这个版本中没有引入重大的改变。

<a name="upgrade-4.1.26"></a>
## 从小于等于4.1.25升级到4.1.26 

Laravel 4.1.26为cookies引入了安全方面的改进。在此更新之前，如果一个cookie被另一个恶意用户劫持，这个cookie将长期有效，即使此账户真正的所有者进行了重置密码、退出登录等操作.

这项改变需要在你的数据表“users”（或等价的表）添加一个新的列“remember_token”。 在这项改变之后，每次用户登录你的应用程序时都会被分配给一个新的token。当此用户从应用程序退出时，token也将被更新。这项改变的意义是： 当cookie被劫持，简单的退出应用程序也会使cookie失效。

### 升级路线

首先，添加一个新的可为空的列“remember_token”到你的“users”表里， 其类型为VARCHAR(100)，或TEXT，或等价的类型。

接下来，如果你使用了Eloquent认证驱动，用下面三个方法来更新你的“User”类：

	public function getRememberToken()
	{
		return $this->remember_token;
	}

	public function setRememberToken($value)
	{
		$this->remember_token = $value;
	}

	public function getRememberTokenName()
	{
		return 'remember_token';
	}

> **注意：** 这项改变将使所有存在的session失效，所以，所有的用户将被迫在你的应用程序上重新认证。

### 包维护者

两个新方法被加入到了“Illuminate\Auth\UserProviderInterface”接口里。默认的驱动里可以找到实现的示例：

	public function retrieveByToken($identifier, $token);

	public function updateRememberToken(UserInterface $user, $token);

“Illuminate\Auth\UserInterface”接口也添加了“升级路线”里描述的三个新方法。

<a name="upgrade-4.1"></a>
## 从4.0升级到4.1

### 升级你的Composer依赖

要升级你的应用程序到Laravel 4.1，把“composer.json”文件里的“laravel/framework”的版本改成“4.1.*”。

### 替换文件

用仓库里的最新副本替换你的“public/index.php”文件（https://github.com/laravel/laravel/blob/master/public/index.php）。

用仓库里的最新副本替换你的“artisan”文件（https://github.com/laravel/laravel/blob/master/artisan）。

### 添加配置文件和选项

更新你的配置文件“app/config/app.php”里的“aliases”和“providers”数组。在文档（https://github.com/laravel/laravel/blob/master/app/config/app.php）里能找到这两个数组更新过的值。一定要把你自己的定制和包服务加回到providers/aliases数组里。

从仓库里添加新的文件“app/config/remote.php”（https://github.com/laravel/laravel/blob/master/app/config/remote.php）.

在你的文件“app/config/session.php”里添加新的配置项“expire_on_close”，默认值应该是“false”。

在你的文件“app/config/queue.php”里添加新的配置章节“failed”。 下面是此章节的默认值：

	'failed' => array(
		'database' => 'mysql', 'table' => 'failed_jobs',
	),

**（可选）** 在你的文件“app/config/view.php”里把配置项“pagination”更新为“pagination::slider-3”。

### 控制器更新

如果“app/controllers/BaseController.php”文件的顶部有使用声明，把“use Illuminate\Routing\Controllers\Controller;”改为“use Illuminate\Routing\Controller;”。

### 密码提醒更新

为了更加灵活，密码提醒被彻底修改了。 通过运行Artisan命令“php artisan auth:reminders-controller”，你可以检查新的存根控制器。你也可以浏览更新文档（/docs/security#password-reminders-and-reset），并按照文档来更新你的应用程序。

更新你的语言文件“app/lang/en/reminders.php”，使其与文件（https://github.com/laravel/laravel/blob/master/app/lang/en/reminders.php）相匹配。

### 环境检测更新

由于安全的原因，URL域不再被用来检测你的应用程序环境。这些值是容易欺骗的，并允许攻击者修改请求环境。你应该把你的环境检测转换成机器主机名（在命令行界面执行“hostname”命令，适用于Mac，Linux，和Windows）。

### 更简单的日志文件

Laravel现在生成一个单独的日志文件“app/storage/logs/laravel.log”。然而，你还可以在文件“app/start/global.php”里配置此行为。

### 删除重定向末尾斜杠

在你的文件“bootstrap/start.php”里，删除此句调用“$app->redirectIfTrailingSlash()”，这个方法不再需要了，因为它的功能被框架里的“.htaccess”文件负责了。

接下来，把你的Apache的“.htaccess”文件替换成新的（https://github.com/laravel/laravel/blob/master/public/.htaccess）， 这个文件是用来处理尾部斜杠的。

### 当前路由访问

现在通过“Route::current()”来访问当前路由，替换掉原来的“Route::getCurrentRoute()”。

### Composer更新

一旦你完成了上面的更改，你就能运行“composer update”命令来更新你的核心应用程序文件了！如果你收到类加载错误，试着运行启用“--no-scripts”选项的更新命令，就像这样：“composer update --no-scripts”。

### 通配符事件监听器

通配符事件监听器不再附加事件到你的处理函数的参数里。如果你需要找到被触发的事件，你应该使用“Event::firing()”。
