//We learned data arrays today!
//Bugün veri dizileri öğrendik!

package main

import  (
	"fmt" )


//is the function from which the application starts.
func main(){

	//This is how we do it to define an array.
	//Array tanımlamak için böyle yaparız.
	var dollarRates = [...]int{17, 18, 17, 16}
    var students = [...]string{"Mustafa", "Fazlı", "Hafize"}

	fmt.Println(students)
	fmt.Println(dollarRates)

    
	
}