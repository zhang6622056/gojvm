package classpath

import (
	"path/filepath"
	"os"
)





//读取class
func (self *Loadclass) ReadClass(className string) ([] byte,Entry ,error){
	className = className + ".class"
	if data,entry,err := self.bootClassPath.readClass(className); err == nil{
		return data,entry,err
	}
	if data,entry,err := self.extClassPath.readClass(className); err == nil{
		return data,entry,err
	}
	return self.userClassPath.readClass(className)
}




type Loadclass struct{
	bootClassPath Entry
	extClassPath Entry
	userClassPath Entry
}

//classpath 转换Entry对象
func Parse(jreOption,cpOption string) *Loadclass{
	cp := &Loadclass{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClassPath(cpOption)
	return cp
}

//转换bootClassPath和ExtClassPath
func (self *Loadclass) parseBootAndExtClasspath(jreOption string){
	jreDir := getJreDir(jreOption)
	// jre/lib/*
	jreLibPath := filepath.Join(jreDir,"lib","*")
	self.bootClassPath = newWildCardEntry(jreLibPath)

	//jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir,"lib","ext","*")
	self.extClassPath = newWildCardEntry(jreExtPath)
}


//获取JAVA_HOME路径
func getJreDir(jreOption string) string{
	if jreOption != "" && exists(jreOption){
		return jreOption
	}
	if exists(". /jre"){
		return ". /jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != ""{
		return filepath.Join(jh,"jre")
	}
	panic("Can not find jre folder!")
}


//转换UserClassPath
func (self *Loadclass) parseUserClassPath(cpOption string){
	if cpOption == ""{
		cpOption = "."
	}
	self.userClassPath = newEntry(cpOption)
}

//判断目录是否存在
func exists(path string) bool{
	if _,err := os.Stat(path) ; err != nil{
		if os.IsNotExist(err){
			return false
		}
	}
	return true
}


//如果没有设置用户classpath，那么默认读取当前路径
func (self *Loadclass) parseUserClasspath(cpOption string){
	if cpOption == ""{
		cpOption = "."
	}
	self.userClassPath = newEntry(cpOption)
}



////输入classpath路径地址
//func (self *Loadclass) String() string{
//	return self.userClassPath.String()
//}