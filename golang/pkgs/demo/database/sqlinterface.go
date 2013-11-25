// database/sql 接口
package sql

//数据库驱动的接口
type Driver interface {
	// 返回一个数据库的Conn连接接口
	// 返回的Conn只能用来进行一次goroutine的操作，也就是说不能把这个Conn应用于Go的多个goroutine里面
	Open(name string) (Conn, error)
}

// Conn  数据库连接接口定义
// Conn只能应用在一个goroutine里面，不能使用在多个goroutine里面，详情请参考上面的说明。
type Conn interface {
	// 返回与当前连接相关的执行Sql语句的准备状态，可以进行查询、删除等操作。
	Prepare(query string) (Smt, error)

	// 关闭当前的连接，执行释放连接拥有的资源等清理工作。因为驱动实现了database/sql里面建议的conn pool，所以你不用再去实现缓存conn之类的，这样会容易引起问题。
	Close() error

	// 返回一个代表事务处理的Tx，通过它你可以进行查询,更新等操作，或者对事务进行回滚、递交。
	Begin() (Tx, error)
}

// Stmt是一种准备好的状态，和Conn相关联，而且只能应用于一个goroutine中，不能应用于多个goroutine。
type Stmt interface {
	// 关闭当前的链接状态，但是如果当前正在执行query，query还是有效返回rows数据。
	Close() error

	// 返回当前预留参数的个数，当返回>=0时数据库驱动就会智能检查调用者的参数。当数据库驱动包不知道预留参数的时候，返回-1。
	NumInput() int

	// 执行Prepare准备好的sql，传入参数执行update/insert等操作，返回Result数据
	Exec(args []Value) (Result, error)

	//执行Prepare准备好的sql，传入需要的参数执行select操作，返回Rows结果集
	Query(args []Valuse) (Rows, error)
}

// 事务接口
type Tx interface {
	// 递交
	Commit() error

	// 回滚
	Rollback() error
}

// 如果这个接口没有定义，那么在调用DB.Exec,就会首先调用Prepare返回Stmt，然后执行Stmt的Exec，然后关闭Stmt。
type Execer interface {
	Exec(query string, args []Value) (Result, error)
}

// 这个是执行Update/Insert等操作返回的结果接口定义
type Result interface {
	// 返回由数据库执行插入操作得到的自增ID号
	LastInsertId() (int64, error)

	// 返回query操作影响的数据条目数
	RowsAffected() (int64, error)
}

// 执行查询返回的结果集接口定义
type Rows interface {
	// 返回查询数据库表的字段信息，这个返回的slice和sql查询的字段一一对应，而不是返回整个表的所有字段。
	Columns() []string

	// 关闭Rows迭代器
	Close() error

	// 返回下一条数据，把数据赋值给dest。dest里面的元素必须是driver.
	// Value的值除了string，返回的数据里面所有的string都必须要转换成[]byte
	// 如果最后没数据了，Next函数最后返回io.EOF。
	Next(dest []Value) error
}

// 底层实现 Result
type RowsAffected int64

func (RowsAffected) LastInsertId() (int64, error)
func (v RowsAffected) RowsAffected() (int64, error)

// Value 是个空接口，也就是说任何数据都可
// drive的Value是驱动必须能够操作的Value，Value要么是nil，要么是下面的任意一种
// int64  float64 bool []byte  time.Time string [*]除了Rows.Next返回的不能是string.
type Value interface{}

//ValueConverter接口定义了如何把一个普通的值转化成driver.Value的接口
type ValueConverter interface {
	ConvertValue(v interface{}) (Value, error)
}

// Valuer接口定义了返回一个driver.Value的方式
type Valuer interface {
	// 多数类型都实现了该方法
	Value() (Value, error)
}

type DB struct {
	Driver   driver.Driver
	dsn      string
	mu       sync.Mutex // protects freeConn and closed
	freeConn []driver.Conn
	closed   bool
}
