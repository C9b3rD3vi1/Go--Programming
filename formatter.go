
/* Assignment
Complete the reformat function. It takes a message string and a formatter function as input:

Apply the given formatter three times to the message
Add a prefix of TEXTIO: to the result
Return the final string
For example, if the message is "General Kenobi" and the formatter adds a period to the end of the string, the final result should be

TEXTIO: General Kenobi...
Copy icon

*/

package main

import "fmt"


func reformat(message string, formatter func(string) string) string {
	result := message

	for i := 0; i < 3; i++ {
		result = formatter(result)
	}
	return "TEXTIO: " + result
}



func formatter(message string) string {
	return message + "."
}



func main() {
	message := "General Kenobi"
	finalResult := reformat(message, formatter) // Call reformat here
	fmt.Println(finalResult) // Print the result directly
}
