package main

import "fmt"



type Entry interface{

}


type Classpath struct{
	bootClassPath Entry
	extClassPath Entry
	userClassPath Entry
}



func b(){
	a := &Classpath{}
	fmt.Println(a)
}



func main(){
	b()
}