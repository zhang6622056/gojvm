package main

import (
	"flag"
)

//定义类型
type cmd struct{
	help bool
	version bool
	classpath string
	class string
	args [] string
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
	flag.Parse()

	return cmd
}


func main(){
	cmd := parseCmd()

	if cmd.help || cmd.class == ""{
		flag.Usage()
	}
}
