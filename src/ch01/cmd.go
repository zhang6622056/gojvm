package main

import (
	"flag"
	"ch01/classpath"
	"fmt"
	"strings"
)

//定义类型
type cmd struct{
	help bool
	version bool
	classpath string
	class string
	args [] string
	XjreOption string
}


//返回值在后，加*
func parseCmd() *cmd{
	//创建结构体实例
	cmd := &cmd{}

	//绑定命令行工具
	flag.BoolVar(&cmd.help,"h",false,"")
	flag.BoolVar(&cmd.version,"version",false,"1.0.0")
	flag.StringVar(&cmd.classpath,"classpath","","the path of the classpath")
	flag.StringVar(&cmd.class,"class","","the path of java file")
	flag.StringVar(&cmd.XjreOption,"Xjre","","the path of jre")
	flag.Parse()

	return cmd
}


func main(){
	cmd := parseCmd()
	startJVM(cmd)
}




//启动jvm
func startJVM(cmd *cmd){
	cp := classpath.Parse(cmd.XjreOption,cmd.classpath)

	fmt.Println("cp")
	fmt.Println(cp)


	className := strings.Replace(cmd.class,".","/",-1)
	classData,_,err := cp.ReadClass(className)
	if err != nil{
		fmt.Printf("Could not find or load main class %s\n",cmd.class)
		return
	}
	fmt.Printf("class data:%v\n",classData)
}