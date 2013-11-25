# io/ioutil

实现了 io 接口的几个函数


    func NopCloser(r io.Reader) io.ReadCloser
    func ReadAll(r io.Reader) ([]byte, error)
    func ReadDir(dirname string) ([]os.FileInfo, error)
    func ReadFile(filename string) ([]byte, error)
    func TempDir(dir, prefix string) (name string, err error)
    func TempFile(dir, prefix string) (f *os.File, err error)
    func WriteFile(filename string, data []byte, perm os.FileMode) error
