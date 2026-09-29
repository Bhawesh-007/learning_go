package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// func main()
// {
// 	 //now this is a prgoram in  which i have to cal the bills
// 	 //learn => taking inout , declaring variables and basic operations
// 	 var foodAmount float32
// 	 var gstPercentage  float32
// 	 var totalPeople int
// 	 var servicePercent float32
// 	 //now i have to cal the bills according right
// 	 fmt.Println("tell me the food amount")
// 	 fmt.Scan(&foodAmount)
// 	 fmt.Println("tell me the gst alloted")
// 	 fmt.Scan(&gstPercentage)
// 	 fmt.Println("tell me the total people")
// 	 fmt.Scan(&totalPeople)
// 	 fmt.Println("tell me the service charge you are alloting ")
// 	 fmt.Scan(&servicePercent)
// 	 var gstAmount = (foodAmount*(gstPercentage))/100;
// 	 var serviceCharge = (foodAmount)*(servicePercent)/100;
// 	 var totalBill  = foodAmount + gstAmount + serviceCharge;
// 	 var perPerson = totalBill/float32(totalPeople);

// 	 fmt.Println("THE GST AMOUNT IS " , gstAmount);
// 	 fmt.Println("the servicecharge is " , serviceCharge )
// 	 fmt.Println("the total bill produced is " , totalBill)
// 	 fmt.Println("the per person bill is " , perPerson)
// }

//noow what  i  will do is instead implement a linux resource using go
//Objective
//Create a Go program that:
//1. Takes a Linux process PID as input.
//2. Reads process information from /proc/<pid>/status.
//3. Extracts important fields.
//4. Prints a clean report.
//5. Warns if memory usage is high.
//so what i making is a simple program in go which takes a linux pid as a input
//and gives information about that process
// i would input a linux process pid
// then how would it acess aobut that pid



func main() {
	var pid string

	fmt.Println("Enter the pid:")
	fmt.Scan(&pid)

	path := "/proc/" + pid + "/status"

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error in opening the path")
		return
	}
	defer file.Close()

	var name string
	var state string
	var processID string
	var threads string
	var vmSize string
	var vmRSS string
	var vmRSSValue int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "Name:") {
			name = strings.TrimSpace(strings.TrimPrefix(line, "Name:"))

		} else if strings.HasPrefix(line, "State:") {
			state = strings.TrimSpace(strings.TrimPrefix(line, "State:"))

		} else if strings.HasPrefix(line, "Pid:") {
			processID = strings.TrimSpace(strings.TrimPrefix(line, "Pid:"))

		} else if strings.HasPrefix(line, "Threads:") {
			threads = strings.TrimSpace(strings.TrimPrefix(line, "Threads:"))

		} else if strings.HasPrefix(line, "VmSize:") {
			vmSize = strings.TrimSpace(strings.TrimPrefix(line, "VmSize:"))

		} else if strings.HasPrefix(line, "VmRSS:") {
			vmRSS = strings.TrimSpace(strings.TrimPrefix(line, "VmRSS:"))

			parts := strings.Fields(vmRSS)
			if len(parts) > 0 {
				input, err := strconv.Atoi(parts[0])
				if err == nil {
					vmRSSValue = input
				}
			}
		}
	}

	if scanner.Err() != nil {
		fmt.Println("Error while reading process file")
		return
	}

	memoryLevel := "LOW"

	if vmRSS == "" {
		memoryLevel = "NOT AVAILABLE"
	} else if vmRSSValue > 200000 {
		memoryLevel = "HIGH"
	} else if vmRSSValue >= 50000 {
		memoryLevel = "MEDIUM"
	}

	fmt.Println()
	fmt.Println("Process Report")
	fmt.Println("--------------")
	fmt.Println("PID:", processID)
	fmt.Println("Name:", name)
	fmt.Println("State:", state)
	fmt.Println("Threads:", threads)

	if vmSize == "" {
		fmt.Println("VmSize: Not available")
	} else {
		fmt.Println("VmSize:", vmSize)
	}

	if vmRSS == "" {
		fmt.Println("VmRSS: Not available")
	} else {
		fmt.Println("VmRSS:", vmRSS)
	}

	fmt.Println()
	fmt.Println("Memory Level:", memoryLevel)
}