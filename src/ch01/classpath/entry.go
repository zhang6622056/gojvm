package classpath

import "os"


const pathListSeparator = string(os.PathListSeparator)

//定义Entry接口
type Entry interface {
	readClass(className string)([] byte,Entry,error)
	String() string
}


//定义新建entry方法
func newEntry(path string) Entry{
	return nil
}

