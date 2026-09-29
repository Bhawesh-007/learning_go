package main 
import "fmt"

func main()
{
	 //now this is a prgoram in  which i have to cal the bills 
	 //learn => taking inout , declaring variables and basic operations 
	 var foodAmount float32
	 var gstPercentage  float32
	 var totalPeople int 
	 var servicePercent float32 
	 //now i have to cal the bills according right
	 fmt.Println("tell me the food amount") 
	 fmt.Scan(&foodAmount)
	 fmt.Println("tell me the gst alloted")
	 fmt.Scan(&gstPercentage)
	 fmt.Println("tell me the total people")
	 fmt.Scan(&totalPeople)
	 fmt.Println("tell me the service charge you are alloting ")
	 fmt.Scan(&servicePercent)	 
	 var gstAmount = (foodAmount*(gstPercentage))/100;
	 var serviceCharge = (foodAmount)*(servicePercent)/100;
	 var totalBill  = foodAmount + gstAmount + serviceCharge;
	 var perPerson = totalBill/float32(totalPeople);

	 fmt.Println("THE GST AMOUNT IS " , gstAmount);
	 fmt.Println("the servicecharge is " , serviceCharge )
	 fmt.Println("the total bill produced is " , totalBill)
	 fmt.Println("the per person bill is " , perPerson)





}