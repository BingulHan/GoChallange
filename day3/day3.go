//Today I processed fixed and immutable data ( const )
//Bugün sabit ve değiştirilemez verileri işledim ( const)

package main

import  (
	"fmt" )

//a constant data is defined like this	
const PI = 3.14

//is the function from which the application starts.
func main(){


	var circumference = 10

	_ = circumference 


	var result = PI * float64(circumference*circumference)
	_ = result



	fmt.Println(float64(result))
    
	
}