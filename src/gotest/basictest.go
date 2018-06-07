package main

import (
	"path/filepath"
	"fmt"
	"os"
	"strings"
	"archive/zip"
	"io/ioutil"
)



//-根本含义
func main(){
	bootjre := "D:\\Java\\jdk8\\jre"
	class := "java/lang/Object.class"
	bootClassPath := filepath.Join(bootjre,"lib","*")
	baseDir := bootClassPath[:len(bootClassPath)-1]


	//遍历目录下的jar文件
	oswalk := func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path,"jar"){
			abspath,errb := filepath.Abs(path)
			if errb != nil{
				return errb
			}
			rc,erro := zip.OpenReader(abspath)
			if erro != nil{
				return erro
			}
			file := rc.File
			for _,f := range file{
				if f.Name == class{
					rcp,erra := f.Open()
					if erra != nil{
						return erra
					}
					data,erru := ioutil.ReadAll(rcp)

					if erru != nil{
						return erru
					}
					fmt.Println(data)
				}
			}
		}
		return nil
	}
	filepath.Walk(baseDir,oswalk)
}



