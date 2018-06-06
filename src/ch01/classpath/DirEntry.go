package classpath

import (
	"path/filepath"
	"io/ioutil"
)

//struct代表实现类
type DirEntry struct{
	absDir string
}

//构造函数
func newDirEntry(path string) *DirEntry{
	//转换绝对路径
	absDir,err := filepath.Abs(path)
	if err != nil{
		//终止其后执行的code
		panic(err)
	}
	//创建entry实例
	return &DirEntry{absDir}
}



//1-传入对象本身 self *DirEntry
//2-readClass(className string) 方法名和传入
//3-返回值 ...([] byte,Entry,error)
func (self *DirEntry) readClass(className string)([] byte,Entry,error){
	fileName := filepath.Join(self.absDir,className)
	data,err := ioutil.ReadFile(fileName)
	return data,self,err
}



//类似tostring方法
func (self *DirEntry) String() string{
	return self.absDir
}
