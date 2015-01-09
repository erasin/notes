# 数据库迁移 & 数据填充

- [简介](#introduction)
- [创建数据迁移](#creating-migrations)
- [运行数据迁移](#running-migrations)
- [回滚数据迁移](#rolling-back-migrations)
- [数据库填充](#database-seeding)

<a name="introduction"></a>
## 简介

Migrations是一种数据库版本控制功能.它允许团队开发者修改数据库结构,并使其保持最新状态.Migrations通常和[结构生成器](schema.md) 配合使用来管理您的应用程序结构.

<a name="creating-migrations"></a>
## 创建数据迁移

使用 Artisan 命令行的 `migrate:make`  命令创建一个迁移：(在命令行模式下使用)

	php artisan migrate:make create_users_table

所有的迁移都被存放在 `app/database/migrations` 文件夹下,文件以时间戳命名以方便Laravel框架按时间来界定这些文件顺序.

您可以在创建迁移的时候使用 `--path` 选项，用来指定迁移文件存放的路径.该路径是你安装框架根目录的相对路径:

	php artisan migrate:make foo --path=app/migrations

`--table` 和 `--create` 选项用来指定表名以及是否创建一个新表：

	php artisan migrate:make add_votes_to_user_table --table=users

	php artisan migrate:make create_users_table --create=users

<a name="running-migrations"></a>
## 运行数据迁移

#### 运行所有迁移(使你的所有表保持最新)

	php artisan migrate

#### 运行某个路径下的所有迁移(指定迁移文件路径)

	php artisan migrate --path=app/foo/migrations

#### 运行某个包下的所有迁移(安装或升级某个扩展包对应数据库时候使用)

	php artisan migrate --package=vendor/package

> **注意:** 如果在运行迁移的时候收到一个 "class not found" 的错误，请尝试运行 `composer dump-autoload` 命令.

### 在生产环境中强制使用数据迁移

有些迁移操作具有破坏性,会导致你丢失数据库中原有数据.为了防止你运行这样的命令造成不必要的破坏,这些命令运行的时候会询问你是否确定要这样做.如果你想运行这样的命令而不出现提示,可以使用 `--force` 选项:

	php artisan migrate --force

<a name="rolling-back-migrations"></a>
## 回滚数据迁移(即使回滚,原有数据也被破坏了,只能回滚表结构,所以别拿这个功能当救命稻草)

#### 回滚最后一次迁移

	php artisan migrate:rollback

#### 回滚所有迁移

	php artisan migrate:reset

#### 回滚所有迁移并重新运行数据迁移

	php artisan migrate:refresh

	php artisan migrate:refresh --seed

<a name="database-seeding"></a>
## 数据库填充

Laravel 可以非常简单的使用数据填充类(seed classes)帮你生成一些测试数据放到数据库中去.所有的数据填充类(seed classes)都存放在 `app/database/seeds` 路径下.数据填充类(seed classes)你可以随便命名,但最好遵循一些合理的约定,例如 `UserTableSeeder` 等. `DatabaseSeeder`是一个已经生成好的默认类(它将默认被执行,你也可以把这个当作例子).在这个类(DatabaseSeeder)中,你可以使用 `call` 方法来运行其他数据填充类,这样你就能控制数据填充的顺序了.

#### 数据库填充类的例子

	class DatabaseSeeder extends Seeder {

		public function run()
		{
			$this->call('UserTableSeeder');

			$this->command->info('User table seeded!');
		}

	}

	class UserTableSeeder extends Seeder {

		public function run()
		{
			DB::table('users')->delete();

			User::create(array('email' => 'foo@bar.com'));
		}

	}

使用 Artisan 命令行的 `db:seed` 命令填充数据库：

	php artisan db:seed

默认情况下 `db:seed` 命令运行的是 `DatabaseSeeder` 类, 这个类中可以(像上面的例子中那样)调用其他seed类. 但是你也可以使用 `--class` 选项来单独运行指定数据库填充类 :

	php artisan db:seed --class=UserTableSeeder

您也可以使用 `migrate:refresh` 命令,这将回滚并重新运行所有数据库迁移:

	php artisan migrate:refresh --seed
