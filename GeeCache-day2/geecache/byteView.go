package geecache

//用于储存缓存的值
type ByteView struct {
	b []byte
}

//返回当前内存的大小 用于判断是否需要淘汰
func (v ByteView) Len() int {
	return len(v.b)
}

//用于返回一个只读的切片 防止缓存值被外部程序修改
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

// 	用于将缓存值转换成字符串 用于调试
func (v ByteView) String() string {
	return string(v.b)
}

//用于复制一个切片 防止外部篡改缓存的值
func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
