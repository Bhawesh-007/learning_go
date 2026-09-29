// now what i have to do here in this excercise is that
// create a go program that would take pids as input and just give me a sorted list of those
// pids
package main

import (
	"fmt"
)
// func main(){
// 	var  n int 
// 	fmt.Scan(&n)
// 	var mem  []int
// 	for i := 1 ; i<=n ;i++ {
// 		 var pid int
// 		 fmt.Scan(&pid)
// 		 mem = append(mem, pid)

// 	}

// 	highest_mem := mem[0]
// 	tot_mem :=0
	
// 	for i:= 0 ; i<n;i++{
// 		if mem[i]>highest_mem{
// 			highest_mem = mem[i]
// 		}
// 		tot_mem += mem[i]
// 	}
// 	fmt.Println("max memory used is " , highest_mem)
// 	fmt.Println("total memory used is " , tot_mem)



// }
func main(){
	  logs := []string{
	"INFO system boot started",
	"INFO network manager started",
	"WARN disk usage high",
	"ERROR sshd failed login",
	"INFO docker started",
	"ERROR nginx crashed",
	"WARN memory pressure detected",
	"INFO backup completed",
	}
	//task => a slice from 2 to 5 right 
	nlogs := logs[2:5] // always remember that end index is excluded in slicing
	fmt.Println(nlogs)
	window := make([]string ,len(logs[2:5])) // make can be used to use a data strutucture which is already
	//initialized and already mem has been allocated to it 

	//make creates a ready-to-use slice with given length and optional capacity.
	copy(window , logs[2:5])
	fmt.Println(window)
}