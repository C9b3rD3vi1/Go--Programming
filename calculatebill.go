/*

Assignment
Complete the calculateFinalBill and the calculateDiscountRate functions.

calculateFinalBill
Use the calculateBaseBill function to get the cost for the messages sent. Then, use the calculateDiscountRate function to get the discount rate. Finally, calculate the final bill by applying the discount to the base bill and return the result.

The discount is a percentage represented by a float e.g. 10% = .1. Remember that any percentage x% is equal to x / 100.

calculateDiscountRate
This function should take the number of messages sent, and return the relevant discount based on the following discount rates:

10% for more than 1000 messages.
20% for more than 5000 messages.
0% for anything less.
*/

package main
import "fmt"



func calculateFinalBill(costPerMessage float64, numMessages int) float64 {
	baseBill := calculateBaseBill(costPerMessage, numMessages) // Calculate base bill
	discountRate := calculateDiscountRate(numMessages)         // Get discount rate
	discountAmount := baseBill * discountRate                   // Calculate discount amount
	finalBill := baseBill - discountAmount                       // Calculate final bill
	return finalBill                                            // Return final bill
}



func calculateDiscountRate(numMessages int) float64 {

	if numMessages >= 5000 {
		return 0.2
		
	} else if numMessages <= 1000 {
		return 0.1
	} else{
		
		return 0.0
	}
		
}



// don't touch below this line
func calculateBaseBill(costPerMessage float64, numMessages int) float64 {
	return costPerMessage * float64(numMessages)
}


func main() {
	costPerMessage := 0.25
	numMessages := 6000 // Example: 6000 messages sent
	finalBill := calculateFinalBill(costPerMessage, numMessages)
	fmt.Printf("Final Bill for %d messages: $%.2f\n", numMessages, finalBill)
}
