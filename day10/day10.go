//We learned Recursion today!
//Bugün Recursion öğrendik!

//Özyineleme, öğelerin kendine benzer bir şekilde tekrarlanması işlemidir. Aynı kavram programlama dillerinde de geçerlidir. Bir program aynı fonksiyon içinde bir fonksiyonun çağrılmasına izin veriyorsa, buna özyinelemeli fonksiyon çağrısı denir. Aşağıdaki örneğe bir göz atın
//Recursion is the process of repeating items in a self-similar way. The same concept applies in programming languages as well. If a program allows to call a function inside the same function, then it is called a recursive function call. Take a look at the following example 

package main

import  (
	"fmt" )

//is the function from which the application starts.
func main() { 
	var i int = 3
	fmt.Printf("Factorial of %d is %d", i, factorial(i))
 }

func factorial(i int)int {
	if(i <= 1) {
	   return 1
	}
	return i * factorial(i - 1)
 }