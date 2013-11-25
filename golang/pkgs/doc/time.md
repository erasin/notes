# time

	type Time struct {}

生成

	// 创建
	func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
	// 解析
	func Parse(layout, value string) (Time, error)
	// 当前时间
	func Now() Time
	// 时间戳转换
	func Unix(sec int64, nsec int64) Time
	// 格式化
	func (t Time) Format(layout string) string
	// string
	func (t Time) String() string

> time.Now() 使用的 Location 是 Local


取值

	func (t Time) Date() (year int, month Month, day int)
	func (t Time) Year() int
	func (t Time) Month() Month
	func (t Time) Day() int
	func (t Time) Hour() int
	func (t Time) Minute() int
	func (t Time) Nanosecond() int
	// 年中第N天
	func (t Time) YearDay() int
	// 时间戳
	func (t Time) Unix() int64
	func (t Time) UnixNano() int64


> 取时间戳 January 1, 1970 UTC.注意 Location

## var

	// 时间格式
	const (
	    ANSIC       = "Mon Jan _2 15:04:05 2006"
	    UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	    RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	    RFC822      = "02 Jan 06 15:04 MST"
	    RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	    RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	    RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	    RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	    RFC3339     = "2006-01-02T15:04:05Z07:00"
	    RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	    Kitchen     = "3:04PM"
	    // Handy time stamps.
	    Stamp      = "Jan _2 15:04:05"
	    StampMilli = "Jan _2 15:04:05.000"
	    StampMicro = "Jan _2 15:04:05.000000"
	    StampNano  = "Jan _2 15:04:05.000000000"
	)

	type Month int
	// 月份
	const (
	    January Month = 1 + iota
	    February
	    March
	    April
	    May
	    June
	    July
	    August
	    September
	    October
	    November
	    December
	)
	
	// 微妙 - 时
	const (
	    Nanosecond  Duration = 1
	    Microsecond          = 1000 * Nanosecond
	    Millisecond          = 1000 * Microsecond
	    Second               = 1000 * Millisecond
	    Minute               = 60 * Second
	    Hour                 = 60 * Minute
	)


## Location

UTC 加上时区，Local为本地

	var Local *Location = &localLoc
	var UTC *Location = &utcLoc

时区的转换，使用对应的 

	// 获取使用方式
	func (t Time) Location() *Location
	// 本地	
	func (t Time) Local() Time
	// 带有时区
	func (t Time) UTC() Time




取当天起始时间 

	time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local).Unix()
	time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC).Unix()

二者取出时间戳结果比较下


	