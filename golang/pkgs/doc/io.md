# io

几个常用接口

	type Reader interface {
	    Read(p []byte) (n int, err error)
	}

	type Writer interface {
	    Write(p []byte) (n int, err error)
	}

	// 偏移引导
	type Seeker interface {
	    Seek(offset int64, whence int) (ret int64, err error)
	}

	type Closer interface {
	    Colse() (err error)
	}

	