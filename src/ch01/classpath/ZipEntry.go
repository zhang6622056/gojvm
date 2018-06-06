package classpath

import (
	"path/filepath"
	"archive/zip"
	"io/ioutil"
	"errors"
	"fmt"
)


type ZipEntry struct{
	absPath string
}

//构造函数
func newZipEntry(path string) *ZipEntry {
	absPath,err := filepath.Abs(path)
	if err != nil{
		panic(err)
	}
	return &ZipEntry{absPath}
}



//从zip中读取class
//defer 关键字
func (self *ZipEntry) readClass(className string) ([] byte,Entry,error){
	fmt.Println(className)

	r,err := zip.OpenReader(self.absPath)
	if err != nil{
		return nil,nil,err
	}
	defer r.Close()

	for _,f := range r.File{
		if f.Name == className {
			rc,err := f.Open()
			if err != nil{
				return nil,nil,err
			}
			defer rc.Close()

			data,err := ioutil.ReadAll(rc)
			if err != nil{
				return  nil,nil,err
			}
			return data,nil,err
		}
	}
	return nil,nil,errors.New("class not found: " + className)
}


//string
func (self *ZipEntry) String() string{
	return self.absPath
}

