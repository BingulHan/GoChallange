
//Today we learned about infinite loops and for loops!
//Bugün sonsuz döngü ve for döngülerini öğrendik!

package main

import  (
	"fmt" )


//is the function from which the application starts.
func main(){

	//The Infinite Loop
	/*
	for true  {
		fmt.Printf("This loop will run forever.\n");
	}
    */

	for a := 0; a < 10; a++ {
		fmt.Printf("value of a: %d\n", a)
	}

	
	fmt.Println("-------")


	var a int = 1;
	var b int = 100;


	for a < b {
		a++
		fmt.Printf("value of a: %d\n", a)
	 }

}