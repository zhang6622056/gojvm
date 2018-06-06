package testpath

import "fmt"



type Entry interface{

}


type Classpath struct{
	bootClassPath Entry
	extClassPath Entry
	userClassPath Entry
}



func Ba(){
	a := &Classpath{}
	fmt.Println(a)
}