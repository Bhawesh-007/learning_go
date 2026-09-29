 package main
 import (
 	"fmt"
 	"bufio"
 	"os"
 	)
// func main(){
// 	 fmt.Println("hello go!!!")
// }
func main(){
	 reader := bufio.NewReader(os.Stdin)
	 fmt.Println("enter the rating" )
	 //now theres is a try and catch method of reading the input instead we say
	 //comma , err method 
	 input,_ := reader.ReadSstring('\n') //read string until next line comes 
	 fmt.Println("thanks for rating" , input)
}