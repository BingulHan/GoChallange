//Today we learned the conditions!
//Bugün koşulları öğrendik!

/**

Go supports the usual comparison operators from mathematics:

Less than <
Less than or equal <=
Greater than >
Greater than or equal >=
Equal to ==
Not equal to !=

Additionally, Go supports the usual logical operators:

Logical AND &&
Logical OR ||
Logical NOT !
You can use these operators or their combinations to create conditions for different decisions.

*/

package main

import  (
	"fmt" )


//is the function from which the application starts.
func main(){


	var x int = 10
	var y int = 20

	if ( x > y) {
		fmt.Println("x is greater than y")
	}else if (x == y) {
		fmt.Println("x and y are equal")
	}else {
		fmt.Println("y is greater than x")
	}
	
}