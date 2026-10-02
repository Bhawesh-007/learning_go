package main

import (
	"fmt"
	"os"
)

// func main(){
// 	 defer fmt.Println("baad main exceute hoga sabse last main")
// 	 defer fmt.Println'("ab sabse pehle ye hoga lifo principle hota hain")
// 	 //in multiple defer statements case the order operates in the reverse order
// 	 fmt.Println("hello world !! ")

// }
//basic io calls for a file in golang
//for example i have a file in a directory how would i read and write over that fiel
const FilePath string = "/home/bhawesh/redhat/learning_go/test.txt"

func main(){
	  Writefile()
	  Readfile()


 }
func Readfile(){
     file,err := os.ReadFile(FilePath)
     if(err!=nil){
     	panic(err)
     }
     fmt.Println("the content of the file is " , string(file))
}
func Writefile(){
	input := "this the content for the file"
    err := os.WriteFile(FilePath,[]byte(input),0644)
    if(err!=nil){
    	 panic(err)
    }else{
    	fmt.Println("the file was succefully updated!!")
    }
}