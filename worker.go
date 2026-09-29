package main 
import (
	 "fmt"
	 "os" 
	 "time"
)
func main(){
	fmt.Println("worker started")
	fmt.Println("get pid ", os.Getpid())

	 //now i will create a worker program which will just do a count 
	 //form 1 to 1e9 lets do
	 for i := 1 ; i<=60 ; i++{
	 	 fmt.Println("Running process .. ", i);
	 	 time.Sleep(1*time.Second);
	 } 


}