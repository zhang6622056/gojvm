package classpath

import (
	"strings"
	"errors"
	"fmt"
	"path/filepath"
	"os"
)

type CompositeEntry []Entry


//将若干个path合并到一个数组内部
func newCompositeEntry(pathList string) CompositeEntry{
	compositeEntry := [] Entry{}
	for _,path := range strings.Split(pathList,pathListSeparator){
		entry := newEntry(path)
		compositeEntry = append(compositeEntry,entry)
	}
	return compositeEntry
}

//通过循环遍历每一个子路径，读取。
func (self CompositeEntry) readClass(className string) ([] byte,Entry,error){

	fmt.Println("self")
	fmt.Println(self)

	for _,entry := range self{
		data,from,err := entry.readClass(className)
		if err == nil{
			return data,from,nil
		}
	}
	return nil,nil,errors.New("class not found:" + className)
}


//tostring
func (self CompositeEntry) String() string{
	strs := make([] string,len(self))
	for i,entry := range self{
		strs[i] = entry.String()
	}
	return strings.Join(strs,pathListSeparator)
}



func newWildCardEntry(path string) CompositeEntry{
	baseDir := path[: len(path) - 1]	//去掉*
	compositeEntry := [] Entry{}

	walkFn := func(path string,info os.FileInfo,err error) error {
		fmt.Println(1111)
		fmt.Println(path)

		if err != nil{
			return err
		}
		if info.IsDir() && path == baseDir {
			return filepath.SkipDir
		}
		if strings.HasPrefix(path,".jar") || strings.HasSuffix(path,".JAR"){
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry,jarEntry)
		}
		return nil
	}

	filepath.Walk(baseDir,walkFn)
	return compositeEntry
}
