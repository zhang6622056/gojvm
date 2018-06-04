package classpath

import (
	"strings"
	"errors"
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



